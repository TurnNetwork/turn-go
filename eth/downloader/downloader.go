// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package downloader contains the manual full chain synchronisation.
package downloader

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	ethereum "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/core/snapshotdb"
	"github.com/PlatONnetwork/PlatON-Go/eth/protocols/snap"
	"github.com/PlatONnetwork/PlatON-Go/trie"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/ethdb"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/metrics"
	"github.com/PlatONnetwork/PlatON-Go/params"
)

const (
	PPOSStorageKVSizeFetch = 100 // the kv size send to peer
	KeyFastSyncStatus      = "FastSyncStatus"
	FastSyncBegin          = 0
	FastSyncFail           = 1
	FastSyncDel            = 2
)

var (
	MaxBlockFetch   = 128 // Amount of blocks to be fetched per retrieval request
	MaxHeaderFetch  = 192 // Amount of block headers to be fetched per retrieval request
	MaxSkeletonSize = 128 // Number of header fetches to need for a skeleton assembly
	MaxReceiptFetch = 256 // Amount of transaction receipts to allow fetching per request
	MaxStateFetch   = 384 // Amount of node state values to allow fetching per request

	rttMinEstimate   = 2 * time.Second  // Minimum round-trip time to target for download requests
	rttMaxEstimate   = 20 * time.Second // Maximum round-trip time to target for download requests
	rttMinConfidence = 0.1              // Worse confidence factor in our estimated RTT value
	ttlScaling       = 3                // Constant scaling factor for RTT -> TTL conversion
	ttlLimit         = time.Minute      // Maximum TTL allowance to prevent reaching crazy timeouts

	qosTuningPeers   = 5    // Number of peers to tune based on (best peers)
	qosConfidenceCap = 10   // Number of peers above which not to modify RTT confidence
	qosTuningImpact  = 0.25 // Impact that a new tuning target has on the previous value

	maxQueuedHeaders         = 32 * 1024                         // [eth/62] Maximum number of headers to queue for import (DOS protection)
	maxHeadersProcess        = 2048                              // Number of header download results to import at once into the chain
	maxResultsProcess        = 2048                              // Number of content download results to import at once into the chain
	maxForkAncestry   uint64 = params.LightImmutabilityThreshold // Maximum chain reorganisation (locally redeclared so tests can reduce it)

	reorgProtHeaderDelay = 2 // Number of headers to delay delivering to cover mini reorgs

	fsHeaderCheckFrequency = 100             // Verification frequency of the downloaded headers during fast sync
	fsHeaderSafetyNet      = 2048            // Number of headers to discard in case a chain violation is detected
	fsHeaderForceVerify    = 24              // Number of headers to verify before and after the pivot to accept it
	fsHeaderContCheck      = 3 * time.Second // Time interval to check for header continuations during state download
)

var (
	errBusy                    = errors.New("busy")
	errUnknownPeer             = errors.New("peer is unknown or unhealthy")
	errBadPeer                 = errors.New("action from bad peer ignored")
	errStallingPeer            = errors.New("peer is stalling")
	errNoPeers                 = errors.New("no peers to keep download active")
	errTimeout                 = errors.New("timeout")
	errEmptyHeaderSet          = errors.New("empty header set by peer")
	errPeersUnavailable        = errors.New("no peers available or all tried for download")
	errInvalidAncestor         = errors.New("retrieved ancestor is invalid")
	errInvalidChain            = errors.New("retrieved hash chain is invalid")
	errInvalidBody             = errors.New("retrieved block body is invalid")
	errInvalidReceipt          = errors.New("retrieved receipt is invalid")
	errCancelStateFetch        = errors.New("state data download canceled (requested)")
	errCancelContentProcessing = errors.New("content processing canceled (requested)")
	errCanceled                = errors.New("syncing canceled (requested)")
	errNoSyncActive            = errors.New("no sync active")
	errTooOld                  = errors.New("peer's protocol version too old")
)

type Downloader struct {
	// WARNING: The `rttEstimate` and `rttConfidence` fields are accessed atomically.
	// On 32 bit platforms, only 64-bit aligned fields can be atomic. The struct is
	// guaranteed to be so aligned, so take advantage of that. For more information,
	// see https://golang.org/pkg/sync/atomic/#pkg-note-BUG.
	rttEstimate   uint64 // Round trip time to target for download requests
	rttConfidence uint64 // Confidence in the estimated RTT (unit: millionths to allow atomic ops)

	mode uint32         // Synchronisation mode defining the strategy used (per sync cycle), use d.getMode() to get the SyncMode
	mux  *event.TypeMux // Event multiplexer to announce sync operation events

	queue *queue   // Scheduler for selecting the hashes to download
	peers *peerSet // Set of active peers from which download can proceed

	stateDB    ethdb.Database  // Database to state sync into (and deduplicate via)
	stateBloom *trie.SyncBloom // Bloom filter for fast trie node and contract code existence checks
	snapshotDB snapshotdb.DB

	// Statistics
	syncStatsChainOrigin uint64 // Origin block number where syncing started at
	syncStatsChainHeight uint64 // Highest block number known when syncing started
	syncStatsState       stateSyncStats
	syncStatsLock        sync.RWMutex // Lock protecting the sync stats fields

	lightchain LightChain
	blockchain BlockChain

	// Callbacks
	dropPeer peerDropFn // Drops a peer for misbehaving

	// Status
	synchroniseMock func(id string, hash common.Hash) error // Replacement for synchronise during testing
	synchronising   int32
	notified        int32
	committed       int32
	ancientLimit    uint64 // The maximum block number which can be regarded as ancient data.

	// Channels
	headerCh          chan dataPack        // Channel receiving inbound block headers
	bodyCh            chan dataPack        // Channel receiving inbound block bodies
	receiptCh         chan dataPack        // Channel receiving inbound receipts
	bodyWakeCh        chan bool            // Channel to signal the block body fetcher of new tasks
	receiptWakeCh     chan bool            // Channel to signal the receipt fetcher of new tasks
	headerProcCh      chan []*types.Header // Channel to feed the header processor new tasks
	pposInfoCh        chan dataPack        // Channel receiving inbound ppos storage
	pposStorageCh     chan dataPack        // Channel receiving inbound ppos storage
	pposStorageDoneCh chan struct{}        // Channel to signal termination completion
	originAndPivotCh  chan dataPack        // Channel receiving origin and pivot block

	// State sync
	pivotHeader *types.Header // Pivot block header to dynamically push the syncing state root
	pivotLock   sync.RWMutex  // Lock protecting pivot header reads from updates

	snapSync   bool         // Whether to run state sync over the snap protocol
	SnapSyncer *snap.Syncer // TODO(karalabe): make private! hack for now

	// for stateFetcher
	stateSyncStart chan *stateSync
	trackStateReq  chan *stateReq
	stateCh        chan dataPack // Channel receiving inbound node state data

	// Cancellation and termination
	cancelPeer string         // Identifier of the peer currently being used as the master (cancel on drop)
	cancelCh   chan struct{}  // Channel to cancel mid-flight syncs
	cancelLock sync.RWMutex   // Lock to protect the cancel channel and peer in delivers
	cancelWg   sync.WaitGroup // Make sure all fetcher goroutines have exited.

	quitCh   chan struct{} // Quit channel to signal termination
	quitLock sync.Mutex    // Lock to prevent double closes

	// Testing hooks
	syncInitHook     func(uint64, uint64)  // Method to call upon initiating a new sync run
	bodyFetchHook    func([]*types.Header) // Method to call upon starting a block body fetch
	receiptFetchHook func([]*types.Header) // Method to call upon starting a receipt fetch
	chainInsertHook  func([]*fetchResult)  // Method to call upon inserting a chain of blocks (possibly in multiple invocations)
}

// LightChain encapsulates functions required to synchronise a light chain.
type LightChain interface {
	// HasHeader verifies a header's presence in the local chain.
	HasHeader(common.Hash, uint64) bool

	// GetHeaderByHash retrieves a header from the local chain.
	GetHeaderByHash(common.Hash) *types.Header

	// CurrentHeader retrieves the head header from the local chain.
	CurrentHeader() *types.Header

	// InsertHeaderChain inserts a batch of headers into the local chain.
	InsertHeaderChain([]*types.Header, int) (int, error)

	// SetHead rewinds the local chain to a new head.
	SetHead(uint64) error
}

// BlockChain encapsulates functions required to sync a (full or fast) blockchain.
type BlockChain interface {
	LightChain

	// HasBlock verifies a block's presence in the local chain.
	HasBlock(common.Hash, uint64) bool

	// GetBlockByHash retrieves a block from the local chain.
	GetBlockByHash(common.Hash) *types.Block

	// CurrentBlock retrieves the head block from the local chain.
	CurrentBlock() *types.Block

	// CurrentFastBlock retrieves the head fast block from the local chain.
	CurrentFastBlock() *types.Block

	// FastSyncCommitHead directly commits the head block to a certain entity.
	FastSyncCommitHead(common.Hash) error

	// InsertChain inserts a batch of blocks into the local chain.
	InsertChain(types.Blocks) (int, error)

	// InsertReceiptChain inserts a batch of receipts into the local chain.
	InsertReceiptChain(types.Blocks, []types.Receipts, uint64) (int, error)
}

// New creates a new downloader to fetch hashes and blocks from remote peers.
func New(stateDb ethdb.Database, snapshotDB snapshotdb.DB, stateBloom *trie.SyncBloom, mux *event.TypeMux, chain BlockChain, lightchain LightChain, dropPeer peerDropFn, decodeExtra decodeExtraFn) *Downloader {
	if lightchain == nil {
		lightchain = chain
	}
	dl := &Downloader{
		stateDB:          stateDb,
		stateBloom:       stateBloom,
		mux:              mux,
		queue:            newQueue(blockCacheMaxItems, blockCacheInitialItems, decodeExtra),
		peers:            newPeerSet(),
		rttEstimate:      uint64(rttMaxEstimate),
		rttConfidence:    uint64(1000000),
		blockchain:       chain,
		lightchain:       lightchain,
		dropPeer:         dropPeer,
		headerCh:         make(chan dataPack, 1),
		bodyCh:           make(chan dataPack, 1),
		receiptCh:        make(chan dataPack, 1),
		bodyWakeCh:       make(chan bool, 1),
		receiptWakeCh:    make(chan bool, 1),
		headerProcCh:     make(chan []*types.Header, 1),
		pposStorageCh:    make(chan dataPack, 1),
		pposInfoCh:       make(chan dataPack, 1),
		originAndPivotCh: make(chan dataPack, 1),
		quitCh:           make(chan struct{}),
		stateCh:          make(chan dataPack),
		SnapSyncer:       snap.NewSyncer(stateDb, stateBloom),
		stateSyncStart:   make(chan *stateSync),
		syncStatsState: stateSyncStats{
			processed: rawdb.ReadFastTrieProgress(stateDb),
		},
		trackStateReq: make(chan *stateReq),
		snapshotDB:    snapshotDB,
	}
	go dl.qosTuner()
	go dl.stateFetcher()
	return dl
}

// Progress retrieves the synchronisation boundaries, specifically the origin
// block where synchronisation started at (may have failed/suspended); the block
// or header sync is currently at; and the latest known block which the sync targets.
//
// In addition, during the state download phase of fast synchronisation the number
// of processed and the total number of known states are also returned. Otherwise
// these are zero.
func (d *Downloader) Progress() ethereum.SyncProgress {
	// Lock the current stats and return the progress
	d.syncStatsLock.RLock()
	defer d.syncStatsLock.RUnlock()

	current := uint64(0)
	mode := d.getMode()
	switch {
	case d.blockchain != nil && mode == FullSync:
		current = d.blockchain.CurrentBlock().NumberU64()
	case d.blockchain != nil && mode == SnapSync:
		current = d.blockchain.CurrentFastBlock().NumberU64()
	case d.lightchain != nil:
		current = d.lightchain.CurrentHeader().Number.Uint64()
	default:
		log.Error("Unknown downloader chain/mode combo", "light", d.lightchain != nil, "full", d.blockchain != nil, "mode", mode)
	}
	return ethereum.SyncProgress{
		StartingBlock: d.syncStatsChainOrigin,
		CurrentBlock:  current,
		HighestBlock:  d.syncStatsChainHeight,
		PulledStates:  d.syncStatsState.processed,
		KnownStates:   d.syncStatsState.processed + d.syncStatsState.pending,
	}
}

// Synchronising returns whether the downloader is currently retrieving blocks.
func (d *Downloader) Synchronising() bool {
	return atomic.LoadInt32(&d.synchronising) > 0
}

// RegisterPeer injects a new download peer into the set of block source to be
// used for fetching hashes and blocks from.
func (d *Downloader) RegisterPeer(id string, version uint, peer Peer) error {
	var logger log.Logger
	if len(id) < 16 {
		// Tests use short IDs, don't choke on them
		logger = log.New("peer", id)
	} else {
		logger = log.New("peer", id[:8])
	}
	logger.Trace("Registering sync peer")
	if err := d.peers.Register(newPeerConnection(id, version, peer, logger)); err != nil {
		logger.Error("Failed to register sync peer", "err", err)
		return err
	}
	d.qosReduceConfidence()

	return nil
}

// RegisterLightPeer injects a light client peer, wrapping it so it appears as a regular peer.
func (d *Downloader) RegisterLightPeer(id string, version uint, peer LightPeer) error {
	return d.RegisterPeer(id, version, &lightPeerWrapper{peer})
}

// UnregisterPeer remove a peer from the known list, preventing any action from
// the specified peer. An effort is also made to return any pending fetches into
// the queue.
func (d *Downloader) UnregisterPeer(id string) error {
	// Unregister the peer from the active peer set and revoke any fetch tasks
	var logger log.Logger
	if len(id) < 16 {
		// Tests use short IDs, don't choke on them
		logger = log.New("peer", id)
	} else {
		logger = log.New("peer", id[:8])
	}
	logger.Trace("Unregistering sync peer")
	if err := d.peers.Unregister(id); err != nil {
		logger.Error("Failed to unregister sync peer", "err", err)
		return err
	}
	d.queue.Revoke(id)

	return nil
}

// Synchronise tries to sync up our local block chain with a remote peer, both
// adding various sanity checks as well as wrapping it with various log entries.
func (d *Downloader) Synchronise(id string, head common.Hash, bn *big.Int, mode SyncMode) error {
	log.Debug("Synchronise from other peer", "peerID", id, "head", head, "bn", bn, "mode", mode)
	err := d.synchronise(id, head, bn, mode)
	switch err {
	case nil, errBusy, errCanceled:
		return err
	}
	if errors.Is(err, errInvalidChain) || errors.Is(err, errBadPeer) || errors.Is(err, errTimeout) ||
		errors.Is(err, errStallingPeer) || errors.Is(err, errEmptyHeaderSet) || errors.Is(err, errPeersUnavailable) ||
		errors.Is(err, errTooOld) || errors.Is(err, errInvalidAncestor) {
		log.Warn("Synchronisation failed, dropping peer", "peer", id, "err", err)
		if d.dropPeer == nil {
			// The dropPeer method is nil when `--copydb` is used for a local copy.
			// Timeouts can occur if e.g. compaction hits at the wrong time, and can be ignored
			log.Warn("Downloader wants to drop peer, but peerdrop-function is not set", "peer", id)
		} else {
			d.dropPeer(id)
		}
		return err
	}
	log.Warn("Synchronisation failed, retrying", "err", err)
	return err
}

// synchronise will select the peer and use it for synchronising. If an empty string is given
// it will use the best peer possible and synchronize if its TD is higher than our own. If any of the
// checks fail an error will be returned. This method is synchronous
func (d *Downloader) synchronise(id string, hash common.Hash, bn *big.Int, mode SyncMode) error {
	// Mock out the synchronisation if testing
	if d.synchroniseMock != nil {
		return d.synchroniseMock(id, hash)
	}
	// Make sure only one goroutine is ever allowed past this point at once
	if !atomic.CompareAndSwapInt32(&d.synchronising, 0, 1) {
		return errBusy
	}
	defer atomic.StoreInt32(&d.synchronising, 0)

	// Post a user notification of the sync (only once per session)
	if atomic.CompareAndSwapInt32(&d.notified, 0, 1) {
		log.Info("Block synchronisation started")
	}
	// If we are already full syncing, but have a fast-sync bloom filter laying
	// around, make sure it doesn't use memory any more. This is a special case
	// when the user attempts to fast sync a new empty network.
	if mode == FullSync && d.stateBloom != nil {
		d.stateBloom.Close()
	}
	// If snap sync was requested, create the snap scheduler and switch to fast
	// sync mode. Long term we could drop fast sync or merge the two together,
	// but until snap becomes prevalent, we should support both. TODO(karalabe).
	if mode == SnapSync {
		if !d.snapSync {
			log.Warn("Enabling snapshot sync prototype")
			d.snapSync = true
		}
		mode = FastSync
	}
	// Reset the queue, peer set and wake channels to clean any internal leftover state
	d.queue.Reset(blockCacheMaxItems, blockCacheInitialItems)
	d.peers.Reset()

	for _, ch := range []chan bool{d.bodyWakeCh, d.receiptWakeCh} {
		select {
		case <-ch:
		default:
		}
	}
	for _, ch := range []chan dataPack{d.headerCh, d.bodyCh, d.receiptCh, d.pposInfoCh, d.pposStorageCh, d.originAndPivotCh} {
		for empty := false; !empty; {
			select {
			case <-ch:
			default:
				empty = true
			}
		}
	}
	for empty := false; !empty; {
		select {
		case <-d.headerProcCh:
		default:
			empty = true
		}
	}
	// Create cancel channel for aborting mid-flight and mark the master peer
	d.cancelLock.Lock()
	d.cancelCh = make(chan struct{})
	d.cancelPeer = id
	d.cancelLock.Unlock()

	defer d.Cancel() // No matter what, we can't leave the cancel channel open

	// Atomically set the requested sync mode
	atomic.StoreUint32(&d.mode, uint32(mode))

	// Retrieve the origin peer and initiate the downloading process
	p := d.peers.Peer(id)
	if p == nil {
		return errUnknownPeer
	}
	return d.syncWithPeer(p, hash, bn)
}

func (d *Downloader) getMode() SyncMode {
	return SyncMode(atomic.LoadUint32(&d.mode))
}

// syncWithPeer starts a block synchronization based on the hash chain from the
// specified peer and head hash.
func (d *Downloader) syncWithPeer(p *peerConnection, hash common.Hash, bn *big.Int) (err error) {
	d.mux.Post(StartEvent{})
	defer func() {
		// reset on error
		if err != nil {
			d.mux.Post(FailedEvent{err})
		} else {
			d.mux.Post(DoneEvent{})
		}
	}()
	if p.version < 64 {
		return fmt.Errorf("%w: advertized %d < required %d", errTooOld, p.version, 64)
	}
	mode := d.getMode()

	log.Debug("Synchronising with the network", "peer", p.id, "eth", p.version, "head", hash, "bn", bn, "mode", mode)
	defer func(start time.Time) {
		log.Debug("Synchronisation terminated", "elapsed", common.PrettyDuration(time.Since(start)))
	}(time.Now())

	var (
		latest          *types.Header
		originh, pivoth *types.Header
		origin          uint64
	)

	originh, pivoth, err = d.findOrigin(p)
	if err != nil {
		log.Error("findOrigin error", "err", err.Error())
		return err
	}
	origin = originh.Number.Uint64()

	// Ensure our origin point is below any fast sync pivot point
	if mode == FastSync {
		pivotNumber := pivoth.Number.Uint64()
		if pivotNumber <= origin {
			origin = pivotNumber - 1
		}
		// Write out the pivot into the database so a rollback beyond it will
		// reenable fast sync
		rawdb.WriteLastPivotNumber(d.stateDB, pivotNumber)
	}
	log.Info("synchronising findOrigin", "peer", p.id, "origin", origin, "pivot", pivoth.Number)
	// Ensure our origin point is below any fast sync pivot point
	d.committed = 1
	if mode == SnapSync {
		if pivoth.Number.Uint64() > origin {
			// fetch latest ppos storage cache from remote peer
			latest, pivoth, err = d.fetchPPOSInfo(p)
			if err != nil {
				return err
			}
			d.committed = 0
		} else {
			log.Info("no need synchronising", "peer", p.id, "origin", origin, "pivot", pivoth.Number)
			d.committed = 0
			return
		}
	} else {
		// Look up the sync boundaries: the common ancestor and the target block
		latest, err = d.fetchHeight(p)
		if err != nil {
			return err
		}

	}
	height := latest.Number.Uint64()

	//origin, err := d.findAncestor(p, height)
	//if err != nil {
	//	return err
	//}

	d.syncStatsLock.Lock()
	if d.syncStatsChainHeight <= origin || d.syncStatsChainOrigin > origin {
		d.syncStatsChainOrigin = origin
	}
	d.syncStatsChainHeight = height
	d.syncStatsLock.Unlock()

	if mode == SnapSync {
		// Set the ancient data limitation.
		// If we are running fast sync, all block data older than ancientLimit will be
		// written to the ancient store. More recent data will be written to the active
		// database and will wait for the freezer to migrate.
		//
		// If there is a checkpoint available, then calculate the ancientLimit through
		// that. Otherwise calculate the ancient limit through the advertised height
		// of the remote peer.
		//
		// The reason for picking checkpoint first is that a malicious peer can give us
		// a fake (very high) height, forcing the ancient limit to also be very high.
		// The peer would start to feed us valid blocks until head, resulting in all of
		// the blocks might be written into the ancient store. A following mini-reorg
		// could cause issues.
		//todo checkpoint mod need add
		if height > maxForkAncestry+1 {
			d.ancientLimit = height - maxForkAncestry - 1
		} else {
			d.ancientLimit = 0
		}
		frozen, _ := d.stateDB.Ancients() // Ignore the error here since light client can also hit here.
		// If a part of blockchain data has already been written into active store,
		// disable the ancient style insertion explicitly.
		if origin >= frozen && frozen != 0 {
			d.ancientLimit = 0
			log.Info("Disabling direct-ancient mode", "origin", origin, "ancient", frozen-1)
		} else if d.ancientLimit > 0 {
			log.Debug("Enabling direct-ancient mode", "ancient", d.ancientLimit)
		}
		// Rewind the ancient store and blockchain if reorg happens.
		// we don't have lightchain,so should not Rollback hash
		/*if origin+1 < frozen {
			var hashes []common.Hash
			if err := d.lightchain.SetHead(origin + 1); err != nil {
				return err
			}
			d.lightchain.Rollback(hashes)
		}*/
	}
	// Initiate the sync using a concurrent header and content retrieval algorithm
	d.queue.Prepare(origin+1, mode)
	if d.syncInitHook != nil {
		d.syncInitHook(origin, height)
	}

	fetchers := []func() error{
		func() error { return d.fetchHeaders(p, origin+1, pivoth.Number.Uint64()) }, // Headers are always retrieved
		func() error { return d.fetchBodies(origin + 1) },                           // Bodies are retrieved during normal and fast sync
		func() error { return d.fetchReceipts(origin + 1) },                         // Receipts are retrieved during fast sync
		func() error { return d.processHeaders(origin+1, pivoth.Number.Uint64(), bn) },
	}
	if mode == SnapSync {
		if err := d.snapshotDB.SetEmpty(); err != nil {
			p.log.Error("set  snapshotDB empty fail")
			return errors.New("set  snapshotDB empty fail:" + err.Error())
		}
		if err := d.snapshotDB.SetCurrent(pivoth.Hash(), *pivoth.Number, *pivoth.Number); err != nil {
			p.log.Error("set snapshotdb current fail", "err", err)
			return errors.New("set current fail")
		}
		fetchers = append(fetchers, func() error { return d.processFastSyncContent(latest, pivoth.Number.Uint64()) })
		fetchers = append(fetchers, func() error { return d.fetchPPOSStorage(p, pivoth) })
	} else if mode == FullSync {
		fetchers = append(fetchers, d.processFullSyncContent)
	}
	if err := d.spawnSync(fetchers); err != nil {
		return err
	}
	return nil
}

// origin is the this chain  current block header ,compare remote  the same num of header in remote,
// the pivot header is the remote peer  snapshotDB base num
func (d *Downloader) findOrigin(p *peerConnection) (*types.Header, *types.Header, error) {
	var current *types.Header
	mode := d.getMode()
	if mode == FullSync {
		current = d.blockchain.CurrentBlock().Header()
	} else if mode == SnapSync {
		current = d.blockchain.CurrentFastBlock().Header()
	} else {
		current = d.lightchain.CurrentHeader()
	}
	currentNumber := current.Number.Uint64()
	go p.peer.RequestOriginAndPivotByCurrent(currentNumber)

	ttl := d.requestTTL()
	timeout := time.After(ttl)
	for {
		select {
		case <-d.cancelCh:
			return nil, nil, errCanceled

		case packet := <-d.originAndPivotCh:
			// Discard anything not from the origin peer
			if packet.PeerId() != p.id {
				p.log.Error("Received headers from incorrect peer", "peer", packet.PeerId())
				return nil, nil, errUnknownPeer
			}
			// Make sure the peer actually gave something valid
			headers := packet.(*headerPack).headers
			if len(headers) == 0 {
				p.log.Error("Empty head header set")
				return nil, nil, errEmptyHeaderSet
			}
			if len(headers) != 2 {
				p.log.Error("header length wrong", "len", len(headers))
				return nil, nil, errors.New("header length wrong")
			}
			if headers[0] == nil {
				p.log.Error("not find  current block")
				return nil, nil, errors.New("not find  current block")
			}
			if headers[0].Number.Uint64() != currentNumber || headers[0].Hash() != current.Hash() {
				p.log.Error("retrieved hash chain is invalid", "current num", currentNumber, "remote current num", headers[0].Number.Uint64(), "current hash", current.Hash(), "remote current hash", headers[0].Hash())
				return nil, nil, errInvalidChain
			}
			if headers[1] == nil {
				return nil, nil, errors.New("not find pivot")
			}
			return headers[0], headers[1], nil
		case <-timeout:
			p.log.Error("Waiting for head header timed out", "elapsed", ttl)
			return nil, nil, errTimeout
		case <-d.bodyCh:
		case <-d.receiptCh:
		case <-d.headerCh:
		case <-d.pposStorageCh:
			// Out of bounds delivery, ignore
		}
	}

}

// Latest is the  remote currentHeader, pivot is remote snapshotDB base num
func (d *Downloader) fetchPPOSInfo(p *peerConnection) (latest *types.Header, pivot *types.Header, err error) {
	p.log.Debug("Retrieving latest ppos info cache from remote peer")
	var current *types.Block

	current = d.blockchain.CurrentFastBlock()

	timeout := time.NewTimer(0) // timer to dump a non-responsive active peer
	<-timeout.C                 // timeout channel should be initially empty
	defer timeout.Stop()
	go p.peer.RequestPPOSStorage()

	var ttl time.Duration
	ttl = d.requestTTL()
	timeout.Reset(ttl)
	for {
		select {
		case <-d.cancelCh:
			return nil, nil, errCanceled
		case <-timeout.C:
			p.log.Error("Waiting for ppos storage timed out", "elapsed", ttl)
			return nil, nil, errTimeout
		case packet := <-d.pposInfoCh:
			// Discard anything not from the origin peer
			if packet.PeerId() != p.id {
				p.log.Error("Received ppos storage from incorrect peer", "peer", packet.PeerId())
				return nil, nil, errors.New("received ppos storage from incorrect peer")
			}
			pposDada := packet.(*pposInfoPack)
			if pposDada.pivot == nil {
				p.log.Error("pivot should not be nil")
				return nil, nil, errors.New("pivot should not be nil")
			}
			pivotNumber := pposDada.pivot.Number
			if pivotNumber.Cmp(pposDada.latest.Number) > 0 {
				p.log.Error("pivotNumber is larger than latestNumber", "pivotNumber", pivotNumber.Uint64(), "latestNumber", pposDada.latest.Number.Uint64())
				return nil, nil, errors.New("pivotNumber is larger than latestNumber")
			}

			if current.NumberU64() >= pposDada.pivot.Number.Uint64() {
				p.log.Error("current is larger than pposDada.pivot", "current", current.NumberU64(), "pposDada.pivot", pposDada.pivot)
				return nil, nil, errors.New("pivotNumber is larger than latestNumber")
			}
			latest = pposDada.latest
			return latest, pposDada.pivot, nil
		case <-d.bodyCh:
		case <-d.receiptCh:
		// Out of bounds delivery, ignore
		case <-d.originAndPivotCh:
		case <-d.pposStorageCh:
		}
	}
}

// setFastSyncStatus set status to snapshot db when fast sync begin
// if  the user close platon when sync not finish,set status fail
// if the sync is complete,will del the key
func (d *Downloader) setFastSyncStatus(status uint16) error {
	key := []byte(KeyFastSyncStatus)
	log.Debug("set  fast sync status", "status", status)
	switch status {
	case FastSyncDel:
		if err := d.snapshotDB.DelBaseDB(key); err != nil {
			log.Error("del  fast sync status  from snapshotdb  fail", "err", err)
			return err
		}
	case FastSyncBegin, FastSyncFail:
		syncStatus := [2][]byte{
			key,
			common.Uint16ToBytes(status),
		}
		if err := d.snapshotDB.WriteBaseDB([][2][]byte{syncStatus}); err != nil {
			log.Error("save  fast sync status to snapshotdb  fail", "err", err)
			return err
		}
	default:
		return errors.New("status is not supports")
	}
	return nil
}

func (d *Downloader) fetchPPOSStorage(p *peerConnection, pivot *types.Header) (err error) {
	log.Debug("Retrieving latest ppos storage cache from remote peer", "pivot number", pivot.Number)
	timeout := time.NewTimer(0) // timer to dump a non-responsive active peer
	<-timeout.C                 // timeout channel should be initially empty
	defer timeout.Stop()
	var ttl time.Duration
	ttl = d.requestTTL()
	timeout.Reset(ttl)
	defer func() {
		close(d.pposStorageDoneCh)
	}()
	d.pposStorageDoneCh = make(chan struct{})

	if err := d.setFastSyncStatus(FastSyncBegin); err != nil {
		return err
	}

	var count int64
	for {
		select {
		case <-d.cancelCh:
			return errCanceled
		case <-timeout.C:
			log.Error("Waiting for ppos storage timed out", "elapsed", ttl)
			return errTimeout
		case packet := <-d.pposStorageCh:
			// Discard anything not from the origin peer
			if packet.PeerId() != p.id {
				log.Error("Received ppos storage from incorrect peer", "peer", packet.PeerId())
				return errors.New("received ppos storage from incorrect peer")
			}
			pposDada := packet.(*pposStoragePack)

			count += int64(len(pposDada.kvs))
			if uint64(count) != pposDada.kvNum {
				p.log.Error("received ppos storage from incorrect kvNum", "kvNum", pposDada.kvNum, "count", count)
				return errors.New("received ppos storage from incorrect kvNum")
			}

			if err := d.snapshotDB.WriteBaseDB(pposDada.KVs()); err != nil {
				p.log.Error("write to base db fail", "err", err)
				return errors.New("write to base db fail")
			}
			if pposDada.last {
				log.Info("fetchPPOSStorage has finish")
				return nil
			}
			ttl = d.requestTTL()
			timeout.Reset(ttl)
		//case <-d.bodyCh:
		//case <-d.receiptCh:
		// Out of bounds delivery, ignore
		case <-d.originAndPivotCh:
		case <-d.pposInfoCh:

		}
	}
}

// spawnSync runs d.process and all given fetcher functions to completion in
// separate goroutines, returning the first error that appears.
func (d *Downloader) spawnSync(fetchers []func() error) error {
	errc := make(chan error, len(fetchers))
	d.cancelWg.Add(len(fetchers))
	for _, fn := range fetchers {
		fn := fn
		go func() { defer d.cancelWg.Done(); errc <- fn() }()
	}
	// Wait for the first error, then terminate the others.
	var (
		err    error
		failed bool
	)
	for i := 0; i < len(fetchers); i++ {
		if i == len(fetchers)-1 {
			// Close the queue when all fetchers have exited.
			// This will cause the block processor to end when
			// it has processed the queue.
			d.queue.Close()
		}
		if err = <-errc; err != nil {
			failed = true
			if err != errCanceled {
				break
			}
		}
	}
	mode := d.getMode()
	if mode == SnapSync {
		if failed {
			if err := d.setFastSyncStatus(FastSyncFail); err != nil {
				return err
			}
		} else {
			if err := d.setFastSyncStatus(FastSyncDel); err != nil {
				return err
			}
		}
	}
	d.queue.Close()
	d.Cancel()
	return err
}

// cancel aborts all of the operations and resets the queue. However, cancel does
// not wait for the running download goroutines to finish. This method should be
// used when cancelling the downloads from inside the downloader.
func (d *Downloader) cancel() {
	// Close the current cancel channel
	d.cancelLock.Lock()
	defer d.cancelLock.Unlock()

	if d.cancelCh != nil {
		select {
		case <-d.cancelCh:
			// Channel was already closed
		default:
			close(d.cancelCh)
		}
	}
}

// Cancel aborts all of the operations and waits for all download goroutines to
// finish before returning.
func (d *Downloader) Cancel() {
	d.cancel()
	d.cancelWg.Wait()

	d.ancientLimit = 0
	log.Debug("Reset ancient limit to zero")
}

// Terminate interrupts the downloader, canceling all pending operations.
// The downloader cannot be reused after calling Terminate.
func (d *Downloader) Terminate() {
	// Close the termination channel (make sure double close is allowed)
	d.quitLock.Lock()
	select {
	case <-d.quitCh:
	default:
		close(d.quitCh)
	}
	if d.stateBloom != nil {
		d.stateBloom.Close()
	}
	d.quitLock.Unlock()

	// Cancel any pending download requests
	d.Cancel()
}

// fetchHeight retrieves the head header of the remote peer to aid in estimating
// the total time a pending synchronisation would take.
func (d *Downloader) fetchHeight(p *peerConnection) (*types.Header, error) {
	p.log.Debug("Retrieving remote chain height")

	// Request the advertised remote head block and wait for the response
	head, _ := p.peer.Head()
	go p.peer.RequestHeadersByHash(head, 1, 0, false)

	ttl := d.requestTTL()
	timeout := time.After(ttl)
	for {
		select {
		case <-d.cancelCh:
			return nil, errCanceled

		case packet := <-d.headerCh:
			// Discard anything not from the origin peer
			if packet.PeerId() != p.id {
				log.Debug("Received headers from incorrect peer", "peer", packet.PeerId())
				break
			}
			// Make sure the peer actually gave something valid
			headers := packet.(*headerPack).headers
			if len(headers) != 1 {
				p.log.Warn("Multiple headers for single request", "headers", len(headers))
				return nil, fmt.Errorf("%w: multiple headers (%d) for single request", errBadPeer, len(headers))
			}
			head := headers[0]
			p.log.Debug("Remote head header identified", "number", head.Number, "hash", head.Hash())
			return head, nil

		case <-timeout:
			p.log.Debug("Waiting for head header timed out", "elapsed", ttl)
			return nil, errTimeout

		case <-d.bodyCh:
		case <-d.receiptCh:
		case <-d.originAndPivotCh:

			// Out of bounds delivery, ignore
		}
	}
}

// fetchHeaders keeps retrieving headers concurrently from the number
// requested, until no more are returned, potentially throttling on the way. To
// facilitate concurrency but still protect against malicious nodes sending bad
// headers, we construct a header chain skeleton using the "origin" peer we are
// syncing with, and fill in the missing headers using anyone else. Headers from
// other peers are only accepted if they map cleanly to the skeleton. If no one
// can fill in the skeleton - not even the origin peer - it's assumed invalid and
// the origin is dropped.
func (d *Downloader) fetchHeaders(p *peerConnection, from uint64, pivot uint64) error {
	p.log.Debug("Directing header downloads", "origin", from, "pivot", pivot)
	defer p.log.Debug("Header download terminated")

	// Create a timeout timer, and the associated header fetcher
	skeleton := true            // Skeleton assembly phase or finishing up
	request := time.Now()       // time of the last skeleton fetch request
	timeout := time.NewTimer(0) // timer to dump a non-responsive active peer
	<-timeout.C                 // timeout channel should be initially empty
	defer timeout.Stop()

	var ttl time.Duration
	getHeaders := func(from uint64) {
		request = time.Now()

		ttl = d.requestTTL()
		timeout.Reset(ttl)

		if skeleton {
			p.log.Trace("Fetching skeleton headers", "count", MaxHeaderFetch, "from", from)
			go p.peer.RequestHeadersByNumber(from+uint64(MaxHeaderFetch)-1, MaxSkeletonSize, MaxHeaderFetch-1, false)
		} else {
			p.log.Trace("Fetching full headers", "count", MaxHeaderFetch, "from", from)
			go p.peer.RequestHeadersByNumber(from, MaxHeaderFetch, 0, false)
		}
	}
	// Start pulling the header chain skeleton until all is done
	getHeaders(from)

	for {
		select {
		case <-d.cancelCh:
			return errCanceled

		case packet := <-d.headerCh:
			// Make sure the active peer is giving us the skeleton headers
			if packet.PeerId() != p.id {
				log.Debug("Received skeleton from incorrect peer", "peer", packet.PeerId())
				break
			}
			headerReqTimer.UpdateSince(request)
			timeout.Stop()

			// If the skeleton's finished, pull any remaining head headers directly from the origin
			if packet.Items() == 0 && skeleton {
				log.Trace("skeleton's finished", "from", from)
				skeleton = false
				getHeaders(from)
				continue
			}
			// If no more headers are inbound, notify the content fetchers and return
			if packet.Items() == 0 {
				// Don't abort header fetches while the pivot is downloading
				if atomic.LoadInt32(&d.committed) == 0 && pivot <= from {
					p.log.Debug("No headers, waiting for pivot commit", "P", pivot, "F", from)
					select {
					case <-time.After(fsHeaderContCheck):
						getHeaders(from)
						continue
					case <-d.cancelCh:
						return errCanceled
					}
				}
				// Pivot done (or not in fast sync) and no more headers, terminate the process
				p.log.Debug("No more headers available")
				select {
				case d.headerProcCh <- nil:
					return nil
				case <-d.cancelCh:
					return errCanceled
				}
			}
			headers := packet.(*headerPack).headers

			// If we received a skeleton batch, resolve internals concurrently
			if skeleton {
				filled, proced, err := d.fillHeaderSkeleton(from, headers)
				if err != nil {
					p.log.Debug("Skeleton chain invalid", "err", err)
					return fmt.Errorf("%w: %v", errInvalidChain, err)
				}
				headers = filled[proced:]
				from += uint64(proced)
			}
			// Insert all the new headers and fetch the next batch
			if len(headers) > 0 {
				p.log.Trace("Scheduling new headers", "count", len(headers), "from", from)
				if headers[0].Number.Cmp(big.NewInt(int64(from))) == 0 {
					select {
					case d.headerProcCh <- headers:
					case <-d.cancelCh:
						return errCanceled
					}
					from += uint64(len(headers))
				} else {
					p.log.Warn("headers not compare, continue fetchHeaders", "count", len(headers), "from", from, "first header", headers[0].Number)
				}
			}
			getHeaders(from)

		case <-timeout.C:
			if d.dropPeer == nil {
				// The dropPeer method is nil when `--copydb` is used for a local copy.
				// Timeouts can occur if e.g. compaction hits at the wrong time, and can be ignored
				p.log.Warn("Downloader wants to drop peer, but peerdrop-function is not set", "peer", p.id)
				break
			}
			// Header retrieval timed out, consider the peer bad and drop
			p.log.Debug("Header request timed out", "elapsed", ttl)
			headerTimeoutMeter.Mark(1)
			d.dropPeer(p.id)

			// Finish the sync gracefully instead of dumping the gathered data though
			for _, ch := range []chan bool{d.bodyWakeCh, d.receiptWakeCh} {
				select {
				case ch <- false:
				case <-d.cancelCh:
				}
			}
			select {
			case d.headerProcCh <- nil:
			case <-d.cancelCh:
			}
			return fmt.Errorf("%w: header request timed out", errBadPeer)
		}
	}
}

// fillHeaderSkeleton concurrently retrieves headers from all our available peers
// and maps them to the provided skeleton header chain.
//
// Any partial results from the beginning of the skeleton is (if possible) forwarded
// immediately to the header processor to keep the rest of the pipeline full even
// in the case of header stalls.
//
// The method returns the entire filled skeleton and also the number of headers
// already forwarded for processing.
func (d *Downloader) fillHeaderSkeleton(from uint64, skeleton []*types.Header) ([]*types.Header, int, error) {
	log.Debug("Filling up skeleton", "from", from)
	d.queue.ScheduleSkeleton(from, skeleton)

	var (
		deliver = func(packet dataPack) (int, error) {
			pack := packet.(*headerPack)
			return d.queue.DeliverHeaders(pack.peerID, pack.headers, d.headerProcCh)
		}
		expire  = func() map[string]int { return d.queue.ExpireHeaders(d.requestTTL()) }
		reserve = func(p *peerConnection, count int) (*fetchRequest, bool, bool) {
			return d.queue.ReserveHeaders(p, count), false, false
		}
		fetch    = func(p *peerConnection, req *fetchRequest) error { return p.FetchHeaders(req.From, MaxHeaderFetch) }
		capacity = func(p *peerConnection) int { return p.HeaderCapacity(d.requestRTT()) }
		setIdle  = func(p *peerConnection, accepted int, deliveryTime time.Time) {
			p.SetHeadersIdle(accepted, deliveryTime)
		}
	)
	err := d.fetchParts(d.headerCh, deliver, d.queue.headerContCh, expire,
		d.queue.PendingHeaders, d.queue.InFlightHeaders, reserve,
		nil, fetch, d.queue.CancelHeaders, capacity, d.peers.HeaderIdlePeers, setIdle, "headers")

	log.Debug("Skeleton fill terminated", "err", err)

	filled, proced := d.queue.RetrieveHeaders()
	return filled, proced, err
}

// fetchBodies iteratively downloads the scheduled block bodies, taking any
// available peers, reserving a chunk of blocks for each, waiting for delivery
// and also periodically checking for timeouts.
func (d *Downloader) fetchBodies(from uint64) error {
	log.Debug("Downloading block bodies", "origin", from)

	var (
		deliver = func(packet dataPack) (int, error) {
			pack := packet.(*bodyPack)
			return d.queue.DeliverBodies(pack.peerID, pack.transactions, pack.extraData)
		}
		expire   = func() map[string]int { return d.queue.ExpireBodies(d.requestTTL()) }
		fetch    = func(p *peerConnection, req *fetchRequest) error { return p.FetchBodies(req) }
		capacity = func(p *peerConnection) int { return p.BlockCapacity(d.requestRTT()) }
		setIdle  = func(p *peerConnection, accepted int, deliveryTime time.Time) { p.SetBodiesIdle(accepted, deliveryTime) }
	)
	err := d.fetchParts(d.bodyCh, deliver, d.bodyWakeCh, expire,
		d.queue.PendingBlocks, d.queue.InFlightBlocks, d.queue.ReserveBodies,
		d.bodyFetchHook, fetch, d.queue.CancelBodies, capacity, d.peers.BodyIdlePeers, setIdle, "bodies")

	log.Debug("Block body download terminated", "err", err)
	return err
}

// fetchReceipts iteratively downloads the scheduled block receipts, taking any
// available peers, reserving a chunk of receipts for each, waiting for delivery
// and also periodically checking for timeouts.
func (d *Downloader) fetchReceipts(from uint64) error {
	log.Debug("Downloading transaction receipts", "origin", from)

	var (
		deliver = func(packet dataPack) (int, error) {
			pack := packet.(*receiptPack)
			return d.queue.DeliverReceipts(pack.peerID, pack.receipts)
		}
		expire   = func() map[string]int { return d.queue.ExpireReceipts(d.requestTTL()) }
		fetch    = func(p *peerConnection, req *fetchRequest) error { return p.FetchReceipts(req) }
		capacity = func(p *peerConnection) int { return p.ReceiptCapacity(d.requestRTT()) }
		setIdle  = func(p *peerConnection, accepted int, deliveryTime time.Time) {
			p.SetReceiptsIdle(accepted, deliveryTime)
		}
	)
	err := d.fetchParts(d.receiptCh, deliver, d.receiptWakeCh, expire,
		d.queue.PendingReceipts, d.queue.InFlightReceipts, d.queue.ReserveReceipts,
		d.receiptFetchHook, fetch, d.queue.CancelReceipts, capacity, d.peers.ReceiptIdlePeers, setIdle, "receipts")

	log.Debug("Transaction receipt download terminated", "err", err)
	return err
}

// fetchParts iteratively downloads scheduled block parts, taking any available
// peers, reserving a chunk of fetch requests for each, waiting for delivery and
// also periodically checking for timeouts.
//
// As the scheduling/timeout logic mostly is the same for all downloaded data
// types, this method is used by each for data gathering and is instrumented with
// various callbacks to handle the slight differences between processing them.
//
// The instrumentation parameters:
//   - errCancel:   error type to return if the fetch operation is cancelled (mostly makes logging nicer)
//   - deliveryCh:  channel from which to retrieve downloaded data packets (merged from all concurrent peers)
//   - deliver:     processing callback to deliver data packets into type specific download queues (usually within `queue`)
//   - wakeCh:      notification channel for waking the fetcher when new tasks are available (or sync completed)
//   - expire:      task callback method to abort requests that took too long and return the faulty peers (traffic shaping)
//   - pending:     task callback for the number of requests still needing download (detect completion/non-completability)
//   - inFlight:    task callback for the number of in-progress requests (wait for all active downloads to finish)
//   - throttle:    task callback to check if the processing queue is full and activate throttling (bound memory use)
//   - reserve:     task callback to reserve new download tasks to a particular peer (also signals partial completions)
//   - fetchHook:   tester callback to notify of new tasks being initiated (allows testing the scheduling logic)
//   - fetch:       network callback to actually send a particular download request to a physical remote peer
//   - cancel:      task callback to abort an in-flight download request and allow rescheduling it (in case of lost peer)
//   - capacity:    network callback to retrieve the estimated type-specific bandwidth capacity of a peer (traffic shaping)
//   - idle:        network callback to retrieve the currently (type specific) idle peers that can be assigned tasks
//   - setIdle:     network callback to set a peer back to idle and update its estimated capacity (traffic shaping)
//   - kind:        textual label of the type being downloaded to display in log messages
func (d *Downloader) fetchParts(deliveryCh chan dataPack, deliver func(dataPack) (int, error), wakeCh chan bool,
	expire func() map[string]int, pending func() int, inFlight func() bool, reserve func(*peerConnection, int) (*fetchRequest, bool, bool),
	fetchHook func([]*types.Header), fetch func(*peerConnection, *fetchRequest) error, cancel func(*fetchRequest), capacity func(*peerConnection) int,
	idle func() ([]*peerConnection, int), setIdle func(*peerConnection, int, time.Time), kind string) error {

	// Create a ticker to detect expired retrieval tasks
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	update := make(chan struct{}, 1)

	// Prepare the queue and fetch block parts until the block header fetcher's done
	finished := false
	for {
		select {
		case <-d.cancelCh:
			return errCanceled

		case packet := <-deliveryCh:
			deliveryTime := time.Now()
			// If the peer was previously banned and failed to deliver its pack
			// in a reasonable time frame, ignore its message.
			if peer := d.peers.Peer(packet.PeerId()); peer != nil {
				// Deliver the received chunk of data and check chain validity
				accepted, err := deliver(packet)
				if errors.Is(err, errInvalidChain) {
					return err
				}
				// Unless a peer delivered something completely else than requested (usually
				// caused by a timed out request which came through in the end), set it to
				// idle. If the delivery's stale, the peer should have already been idled.
				if !errors.Is(err, errStaleDelivery) {
					setIdle(peer, accepted, deliveryTime)
				}
				// Issue a log to the user to see what's going on
				switch {
				case err == nil && packet.Items() == 0:
					peer.log.Trace("Requested data not delivered", "type", kind)
				case err == nil:
					peer.log.Trace("Delivered new batch of data", "type", kind, "count", packet.Stats())
				default:
					peer.log.Trace("Failed to deliver retrieved data", "type", kind, "err", err)
				}
			}
			// Blocks assembled, try to update the progress
			select {
			case update <- struct{}{}:
			default:
			}

		case cont := <-wakeCh:
			// The header fetcher sent a continuation flag, check if it's done
			if !cont {
				finished = true
			}
			// Headers arrive, try to update the progress
			select {
			case update <- struct{}{}:
			default:
			}

		case <-ticker.C:
			// Sanity check update the progress
			select {
			case update <- struct{}{}:
			default:
			}

		case <-update:
			// Short circuit if we lost all our peers
			if d.peers.Len() == 0 {
				return errNoPeers
			}
			// Check for fetch request timeouts and demote the responsible peers
			for pid, fails := range expire() {
				if peer := d.peers.Peer(pid); peer != nil {
					// If a lot of retrieval elements expired, we might have overestimated the remote peer or perhaps
					// ourselves. Only reset to minimal throughput but don't drop just yet. If even the minimal times
					// out that sync wise we need to get rid of the peer.
					//
					// The reason the minimum threshold is 2 is because the downloader tries to estimate the bandwidth
					// and latency of a peer separately, which requires pushing the measures capacity a bit and seeing
					// how response times reacts, to it always requests one more than the minimum (i.e. min 2).
					if fails > 2 {
						peer.log.Trace("Data delivery timed out", "type", kind)
						setIdle(peer, 0, time.Now())
					} else {
						peer.log.Debug("Stalling delivery, dropping", "type", kind)

						if d.dropPeer == nil {
							// The dropPeer method is nil when `--copydb` is used for a local copy.
							// Timeouts can occur if e.g. compaction hits at the wrong time, and can be ignored
							peer.log.Warn("Downloader wants to drop peer, but peerdrop-function is not set", "peer", pid)
						} else {
							d.dropPeer(pid)

							// If this peer was the master peer, abort sync immediately
							d.cancelLock.RLock()
							master := pid == d.cancelPeer
							d.cancelLock.RUnlock()

							if master {
								d.cancel()
								return errTimeout
							}
						}
					}
				}
			}
			// If there's nothing more to fetch, wait or terminate
			if pending() == 0 {
				if !inFlight() && finished {
					log.Debug("Data fetching completed", "type", kind)
					return nil
				}
				break
			}
			// Send a download request to all idle peers, until throttled
			progressed, throttled, running := false, false, inFlight()
			idles, total := idle()
			pendCount := pending()
			for _, peer := range idles {
				// Short circuit if throttling activated
				if throttled {
					break
				}
				// Short circuit if there is no more available task.
				if pendCount = pending(); pendCount == 0 {
					break
				}
				// Reserve a chunk of fetches for a peer. A nil can mean either that
				// no more headers are available, or that the peer is known not to
				// have them.
				request, progress, throttle := reserve(peer, capacity(peer))
				if progress {
					progressed = true
				}
				if throttle {
					throttled = true
					throttleCounter.Inc(1)
				}
				if request == nil {
					continue
				}
				if request.From > 0 {
					peer.log.Trace("Requesting new batch of data", "type", kind, "from", request.From)
				} else {
					peer.log.Trace("Requesting new batch of data", "type", kind, "count", len(request.Headers), "from", request.Headers[0].Number)
				}
				// Fetch the chunk and make sure any errors return the hashes to the queue
				if fetchHook != nil {
					fetchHook(request.Headers)
				}
				if err := fetch(peer, request); err != nil {
					// Although we could try and make an attempt to fix this, this error really
					// means that we've double allocated a fetch task to a peer. If that is the
					// case, the internal state of the downloader and the queue is very wrong so
					// better hard crash and note the error instead of silently accumulating into
					// a much bigger issue.
					panic(fmt.Sprintf("%v: %s fetch assignment failed", peer, kind))
				}
				running = true
			}
			// Make sure that we have peers available for fetching. If all peers have been tried
			// and all failed throw an error
			if !progressed && !throttled && !running && len(idles) == total && pendCount > 0 {
				return errPeersUnavailable
			}
		}
	}
}

// processHeaders takes batches of retrieved headers from an input channel and
// keeps processing and scheduling them into the header chain and downloader's
// queue until the stream ends or a failure occurs.
func (d *Downloader) processHeaders(origin uint64, pivot uint64, bn *big.Int) error {
	// Keep a count of uncertain headers to roll back
	var (
		rollback    uint64 // Zero means no rollback (fine as you can't unroll the genesis)
		rollbackErr error
		mode        = d.getMode()
	)
	defer func() {
		if rollback > 0 {
			lastHeader, lastFastBlock, lastBlock := d.lightchain.CurrentHeader().Number, common.Big0, common.Big0
			if mode != LightSync {
				lastFastBlock = d.blockchain.CurrentFastBlock().Number()
				lastBlock = d.blockchain.CurrentBlock().Number()
			}
			if err := d.lightchain.SetHead(rollback - 1); err != nil { // -1 to target the parent of the first uncertain block
				// We're already unwinding the stack, only print the error to make it more visible
				log.Error("Failed to roll back chain segment", "head", rollback-1, "err", err)
			}
			curFastBlock, curBlock := common.Big0, common.Big0
			if mode != LightSync {
				curFastBlock = d.blockchain.CurrentFastBlock().Number()
				curBlock = d.blockchain.CurrentBlock().Number()
			}
			log.Warn("Rolled back chain segment",
				"header", fmt.Sprintf("%d->%d", lastHeader, d.lightchain.CurrentHeader().Number),
				"fast", fmt.Sprintf("%d->%d", lastFastBlock, curFastBlock),
				"block", fmt.Sprintf("%d->%d", lastBlock, curBlock), "reason", rollbackErr)
		}
	}()

	// Wait for batches of headers to process
	gotHeaders := false

	for {
		select {
		case <-d.cancelCh:
			rollbackErr = errCanceled
			return errCanceled

		case headers := <-d.headerProcCh:
			// Terminate header processing if we synced up
			if len(headers) == 0 {
				// Notify everyone that headers are fully processed
				for _, ch := range []chan bool{d.bodyWakeCh, d.receiptWakeCh} {
					select {
					case ch <- false:
					case <-d.cancelCh:
					}
				}
				// If no headers were retrieved at all, the peer violated its TD promise that it had a
				// better chain compared to ours. The only exception is if its promised blocks were
				// already imported by other means (e.g. fetcher):
				//
				// R <remote peer>, L <local node>: Both at block 10
				// R: Mine block 11, and propagate it to L
				// L: Queue block 11 for import
				// L: Notice that R's head and TD increased compared to ours, start sync
				// L: Import of block 11 finishes
				// L: Sync begins, and finds common ancestor at 11
				// L: Request new headers up from 11 (R's TD was higher, it must have something)
				// R: Nothing to give
				if mode != LightSync {
					head := d.blockchain.CurrentBlock()
					if !gotHeaders && bn.Cmp(head.Number()) > 0 {
						return errStallingPeer
					}
				}
				// If fast or light syncing, ensure promised headers are indeed delivered. This is
				// needed to detect scenarios where an attacker feeds a bad pivot and then bails out
				// of delivering the post-pivot blocks that would flag the invalid content.
				//
				// This check cannot be executed "as is" for full imports, since blocks may still be
				// queued for processing when the header download completes. However, as long as the
				// peer gave us something useful, we're already happy/progressed (above check).
				if mode == SnapSync || mode == LightSync {
					head := d.lightchain.CurrentHeader()
					if bn.Cmp(head.Number) > 0 {
						return errStallingPeer
					}
				}
				// Disable any rollback and return
				rollback = 0
				return nil
			}
			// Otherwise split the chunk of headers into batches and process them
			gotHeaders = true
			for len(headers) > 0 {
				// Terminate if something failed in between processing chunks
				select {
				case <-d.cancelCh:
					rollbackErr = errCanceled
					return errCanceled
				default:
				}
				// Select the next chunk of headers to import
				limit := maxHeadersProcess
				if limit > len(headers) {
					limit = len(headers)
				}
				chunk := headers[:limit]
				// In case of header only syncing, validate the chunk immediately
				if mode == SnapSync || mode == LightSync {
					// If we're importing pure headers, verify based on their recentness
					frequency := fsHeaderCheckFrequency
					if chunk[len(chunk)-1].Number.Uint64()+uint64(fsHeaderForceVerify) > pivot {
						frequency = 1
					}
					if n, err := d.lightchain.InsertHeaderChain(chunk, frequency); err != nil {
						rollbackErr = err
						// If some headers were inserted, track them as uncertain
						if n > 0 && rollback == 0 {
							rollback = chunk[0].Number.Uint64()
						}
						log.Debug("Invalid header encountered", "number", chunk[n].Number, "hash", chunk[n].Hash(), "parent", chunk[n].ParentHash, "err", err)
						return fmt.Errorf("%w: %v", errInvalidChain, err)
					}
					// All verifications passed, track all headers within the alloted limits
					head := chunk[len(chunk)-1].Number.Uint64()
					if head-rollback > uint64(fsHeaderSafetyNet) {
						rollback = head - uint64(fsHeaderSafetyNet)
					}
				}
				// Unless we're doing light chains, schedule the headers for associated content retrieval
				if mode == FullSync || mode == SnapSync {
					// If we've reached the allowed number of pending headers, stall a bit
					for d.queue.PendingBlocks() >= maxQueuedHeaders || d.queue.PendingReceipts() >= maxQueuedHeaders {
						select {
						case <-d.cancelCh:
							rollbackErr = errCanceled
							return errCanceled
						case <-time.After(time.Second):
						}
					}
					// Otherwise insert the headers for content retrieval
					inserts := d.queue.Schedule(chunk, origin)
					if len(inserts) != len(chunk) {
						rollbackErr = fmt.Errorf("stale headers: len inserts %v len(chunk) %v", len(inserts), len(chunk))
						return fmt.Errorf("%w: stale headers", errBadPeer)
					}
				}
				headers = headers[limit:]
				origin += uint64(limit)
			}
			// Update the highest block number we know if a higher one is found.
			d.syncStatsLock.Lock()
			if d.syncStatsChainHeight < origin {
				d.syncStatsChainHeight = origin - 1
			}
			d.syncStatsLock.Unlock()

			// Signal the content downloaders of the availablility of new tasks
			for _, ch := range []chan bool{d.bodyWakeCh, d.receiptWakeCh} {
				select {
				case ch <- true:
				default:
				}
			}
		}
	}
}

// processFullSyncContent takes fetch results from the queue and imports them into the chain.
func (d *Downloader) processFullSyncContent() error {
	for {
		results := d.queue.Results(true)
		if len(results) == 0 {
			return nil
		}
		if d.chainInsertHook != nil {
			d.chainInsertHook(results)
		}
		if err := d.importBlockResults(results); err != nil {
			return err
		}
	}
}

func (d *Downloader) importBlockResults(results []*fetchResult) error {
	// Check for any early termination requests
	if len(results) == 0 {
		return nil
	}
	select {
	case <-d.quitCh:
		return errCancelContentProcessing
	default:
	}
	// Retrieve the a batch of results to import
	first, last := results[0].Header, results[len(results)-1].Header
	log.Debug("Inserting downloaded chain", "items", len(results),
		"firstnum", first.Number, "firsthash", first.Hash(),
		"lastnum", last.Number, "lasthash", last.Hash(),
	)
	blocks := make([]*types.Block, len(results))
	for i, result := range results {
		blocks[i] = types.NewBlockWithHeader(result.Header).WithBody(result.Transactions, result.ExtraData)
	}
	if index, err := d.blockchain.InsertChain(blocks); err != nil {
		if index < len(results) {
			log.Debug("Downloaded item processing failed", "number", results[index].Header.Number, "hash", results[index].Header.Hash(), "err", err)
		} else {
			// The InsertChain method in blockchain.go will sometimes return an out-of-bounds index,
			// when it needs to preprocess blocks to import a sidechain.
			// The importer will put together a new list of blocks to import, which is a superset
			// of the blocks delivered from the downloader, and the indexing will be off.
			log.Debug("Downloaded item processing failed on sidechain import", "index", index, "err", err)
		}
		return fmt.Errorf("%w: %v", errInvalidChain, err)
	}
	return nil
}

// processFastSyncContent takes fetch results from the queue and writes them to the
// database. It also controls the synchronisation of state nodes of the pivot block.
func (d *Downloader) processFastSyncContent(latest *types.Header, pivot uint64) error {
	// Start syncing state of the reported head block. This should get us most of
	// the state of the pivot block.
	sync := d.syncState(latest.Root)
	defer func() {
		// The `sync` object is replaced every time the pivot moves. We need to
		// defer close the very last active one, hence the lazy evaluation vs.
		// calling defer sync.Cancel() !!!
		sync.Cancel()
	}()

	closeOnErr := func(s *stateSync) {
		if err := s.Wait(); err != nil && err != errCancelStateFetch && err != errCanceled && err != snap.ErrCancelled {
			d.queue.Close() // wake up Results
		}
	}
	go closeOnErr(sync)
	// To cater for moving pivot points, track the pivot block and subsequently
	// accumulated download results separately.
	var (
		oldPivot *fetchResult   // Locked in pivot block, might change eventually
		oldTail  []*fetchResult // Downloaded content after the pivot
	)
	for {
		// Wait for the next batch of downloaded data to be available, and if the pivot
		// block became stale, move the goalpost
		results := d.queue.Results(oldPivot == nil) // Block if we're not monitoring pivot staleness
		if len(results) == 0 {
			// If pivot sync is done, stop
			if oldPivot == nil {
				return sync.Cancel()
			}
			// If sync failed, stop
			select {
			case <-d.cancelCh:
				sync.Cancel()
				return errCanceled
			default:
			}
		}
		if d.chainInsertHook != nil {
			d.chainInsertHook(results)
		}
		if oldPivot != nil {
			results = append(append([]*fetchResult{oldPivot}, oldTail...), results...)
		}

		P, beforeP, afterP := splitAroundPivot(pivot, results)
		if err := d.commitFastSyncData(beforeP, sync); err != nil {
			return err
		}
		if P != nil {
			// If new pivot block found, cancel old state retrieval and restart
			if oldPivot != P {
				sync.Cancel()
				sync = d.syncState(P.Header.Root)

				go closeOnErr(sync)
				oldPivot = P
			}
			// Wait for completion, occasionally checking for pivot staleness
			log.Debug("wait for pposStorageDoneCh")
			select {
			case <-d.pposStorageDoneCh:
				log.Debug("finish pposStorageDoneCh")
			case <-time.After(time.Second):
				oldTail = afterP
				continue
			}
			select {
			case <-sync.done:
				if sync.err != nil {
					return sync.err
				}
				if err := d.commitPivotBlock(P); err != nil {
					return err
				}
				oldPivot = nil

			case <-time.After(time.Second):
				oldTail = afterP
				continue
			}
		}
		// Fast sync done, pivot commit done, full import
		if err := d.importBlockResults(afterP); err != nil {
			return err
		}
	}
}

func splitAroundPivot(pivot uint64, results []*fetchResult) (p *fetchResult, before, after []*fetchResult) {
	if len(results) == 0 {
		return nil, nil, nil
	}
	if lastNum := results[len(results)-1].Header.Number.Uint64(); lastNum < pivot {
		// the pivot is somewhere in the future
		return nil, results, nil
	}
	// This can also be optimized, but only happens very seldom
	for _, result := range results {
		num := result.Header.Number.Uint64()
		switch {
		case num < pivot:
			before = append(before, result)
		case num == pivot:
			p = result
		default:
			after = append(after, result)
		}
	}
	return p, before, after
}

func (d *Downloader) commitFastSyncData(results []*fetchResult, stateSync *stateSync) error {
	// Check for any early termination requests
	if len(results) == 0 {
		return nil
	}
	select {
	case <-d.quitCh:
		return errCancelContentProcessing
	case <-stateSync.done:
		if err := stateSync.Wait(); err != nil {
			return err
		}
	default:
	}
	// Retrieve the a batch of results to import
	first, last := results[0].Header, results[len(results)-1].Header
	log.Debug("Inserting fast-sync blocks", "items", len(results),
		"firstnum", first.Number, "firsthash", first.Hash(),
		"lastnumn", last.Number, "lasthash", last.Hash(),
	)
	blocks := make([]*types.Block, len(results))
	receipts := make([]types.Receipts, len(results))
	for i, result := range results {
		blocks[i] = types.NewBlockWithHeader(result.Header).WithBody(result.Transactions, result.ExtraData)
		receipts[i] = result.Receipts
	}
	if index, err := d.blockchain.InsertReceiptChain(blocks, receipts, d.ancientLimit); err != nil {
		log.Debug("Downloaded item processing failed", "number", results[index].Header.Number, "hash", results[index].Header.Hash(), "err", err)
		return fmt.Errorf("%w: %v", errInvalidChain, err)
	}
	return nil
}

func (d *Downloader) commitPivotBlock(result *fetchResult) error {
	block := types.NewBlockWithHeader(result.Header).WithBody(result.Transactions, result.ExtraData)
	log.Debug("Committing fast sync pivot as new head", "number", block.Number(), "hash", block.Hash())

	// Commit the pivot block as the new head, will require full sync from here on
	if _, err := d.blockchain.InsertReceiptChain([]*types.Block{block}, []types.Receipts{result.Receipts}, d.ancientLimit); err != nil {
		return err
	}
	if err := d.blockchain.FastSyncCommitHead(block.Hash()); err != nil {
		return err
	}
	atomic.StoreInt32(&d.committed, 1)

	// If we had a bloom filter for the state sync, deallocate it now. Note, we only
	// deallocate internally, but keep the empty wrapper. This ensures that if we do
	// a rollback after committing the pivot and restarting fast sync, we don't end
	// up using a nil bloom. Empty bloom is fine, it just returns that it does not
	// have the info we need, so reach down to the database instead.
	if d.stateBloom != nil {
		d.stateBloom.Close()
	}
	return nil
}

// DeliverPposStorage injects a new batch of ppos storage received from a remote node.
func (d *Downloader) DeliverPposStorage(id string, kvs [][2][]byte, last bool, kvNum uint64) (err error) {
	return d.deliver(d.pposStorageCh, &pposStoragePack{id, kvs, last, kvNum}, pposStorageInMeter, pposStorageDropMeter)
}

// DeliverPposStorage injects a new batch of ppos storage received from a remote node.
func (d *Downloader) DeliverPposInfo(id string, latest, pivot *types.Header) (err error) {
	return d.deliver(d.pposInfoCh, &pposInfoPack{id, latest, pivot}, pposStorageInMeter, pposStorageDropMeter)
}

func (d *Downloader) DeliverOriginAndPivot(id string, headers []*types.Header) (err error) {
	return d.deliver(d.originAndPivotCh, &headerPack{id, headers}, headerInMeter, headerDropMeter)
}

// DeliverHeaders injects a new batch of block headers received from a remote
// node into the download schedule.
func (d *Downloader) DeliverHeaders(id string, headers []*types.Header) error {
	return d.deliver(d.headerCh, &headerPack{id, headers}, headerInMeter, headerDropMeter)
}

// DeliverBodies injects a new batch of block bodies received from a remote node.
func (d *Downloader) DeliverBodies(id string, transactions [][]*types.Transaction, extraData [][]byte) error {
	return d.deliver(d.bodyCh, &bodyPack{id, transactions, extraData}, bodyInMeter, bodyDropMeter)
}

// DeliverReceipts injects a new batch of receipts received from a remote node.
func (d *Downloader) DeliverReceipts(id string, receipts [][]*types.Receipt) error {
	return d.deliver(d.receiptCh, &receiptPack{id, receipts}, receiptInMeter, receiptDropMeter)
}

// DeliverNodeData injects a new batch of node state data received from a remote node.
func (d *Downloader) DeliverNodeData(id string, data [][]byte) error {
	return d.deliver(d.stateCh, &statePack{id, data}, stateInMeter, stateDropMeter)
}

// DeliverSnapPacket is invoked from a peer's message handler when it transmits a
// data packet for the local node to consume.
func (d *Downloader) DeliverSnapPacket(peer *snap.Peer, packet snap.Packet) error {
	switch packet := packet.(type) {
	case *snap.AccountRangePacket:
		hashes, accounts, err := packet.Unpack()
		if err != nil {
			return err
		}
		return d.SnapSyncer.OnAccounts(peer, packet.ID, hashes, accounts, packet.Proof)

	case *snap.StorageRangesPacket:
		hashset, slotset := packet.Unpack()
		return d.SnapSyncer.OnStorage(peer, packet.ID, hashset, slotset, packet.Proof)

	case *snap.ByteCodesPacket:
		return d.SnapSyncer.OnByteCodes(peer, packet.ID, packet.Codes)

	case *snap.TrieNodesPacket:
		return d.SnapSyncer.OnTrieNodes(peer, packet.ID, packet.Nodes)

	default:
		return fmt.Errorf("unexpected snap packet type: %T", packet)
	}
}

// deliver injects a new batch of data received from a remote node.
func (d *Downloader) deliver(destCh chan dataPack, packet dataPack, inMeter, dropMeter metrics.Meter) (err error) {
	// Update the delivery metrics for both good and failed deliveries
	inMeter.Mark(int64(packet.Items()))
	defer func() {
		if err != nil {
			dropMeter.Mark(int64(packet.Items()))
		}
	}()
	// Deliver or abort if the sync is canceled while queuing
	d.cancelLock.RLock()
	cancel := d.cancelCh
	d.cancelLock.RUnlock()
	if cancel == nil {
		return errNoSyncActive
	}
	select {
	case destCh <- packet:
		return nil
	case <-cancel:
		return errNoSyncActive
	}
}

// qosTuner is the quality of service tuning loop that occasionally gathers the
// peer latency statistics and updates the estimated request round trip time.
func (d *Downloader) qosTuner() {
	for {
		// Retrieve the current median RTT and integrate into the previoust target RTT
		rtt := time.Duration((1-qosTuningImpact)*float64(atomic.LoadUint64(&d.rttEstimate)) + qosTuningImpact*float64(d.peers.medianRTT()))
		atomic.StoreUint64(&d.rttEstimate, uint64(rtt))

		// A new RTT cycle passed, increase our confidence in the estimated RTT
		conf := atomic.LoadUint64(&d.rttConfidence)
		conf = conf + (1000000-conf)/2
		atomic.StoreUint64(&d.rttConfidence, conf)

		// Log the new QoS values and sleep until the next RTT
		log.Debug("Recalculated downloader QoS values", "rtt", rtt, "confidence", float64(conf)/1000000.0, "ttl", d.requestTTL())
		select {
		case <-d.quitCh:
			return
		case <-time.After(rtt):
		}
	}
}

// qosReduceConfidence is meant to be called when a new peer joins the downloader's
// peer set, needing to reduce the confidence we have in out QoS estimates.
func (d *Downloader) qosReduceConfidence() {
	// If we have a single peer, confidence is always 1
	peers := uint64(d.peers.Len())
	if peers == 0 {
		// Ensure peer connectivity races don't catch us off guard
		return
	}
	if peers == 1 {
		atomic.StoreUint64(&d.rttConfidence, 1000000)
		return
	}
	// If we have a ton of peers, don't drop confidence)
	if peers >= uint64(qosConfidenceCap) {
		return
	}
	// Otherwise drop the confidence factor
	conf := atomic.LoadUint64(&d.rttConfidence) * (peers - 1) / peers
	if float64(conf)/1000000 < rttMinConfidence {
		conf = uint64(rttMinConfidence * 1000000)
	}
	atomic.StoreUint64(&d.rttConfidence, conf)

	rtt := time.Duration(atomic.LoadUint64(&d.rttEstimate))
	log.Debug("Relaxed downloader QoS values", "rtt", rtt, "confidence", float64(conf)/1000000.0, "ttl", d.requestTTL())
}

// requestRTT returns the current target round trip time for a download request
// to complete in.
//
// Note, the returned RTT is .9 of the actually estimated RTT. The reason is that
// the downloader tries to adapt queries to the RTT, so multiple RTT values can
// be adapted to, but smaller ones are preferred (stabler download stream).
func (d *Downloader) requestRTT() time.Duration {
	return time.Duration(atomic.LoadUint64(&d.rttEstimate)) * 9 / 10
}

// requestTTL returns the current timeout allowance for a single download request
// to finish under.
func (d *Downloader) requestTTL() time.Duration {
	var (
		rtt  = time.Duration(atomic.LoadUint64(&d.rttEstimate))
		conf = float64(atomic.LoadUint64(&d.rttConfidence)) / 1000000.0
	)
	ttl := time.Duration(ttlScaling) * time.Duration(float64(rtt)/conf)
	if ttl > ttlLimit {
		ttl = ttlLimit
	}
	return ttl
}
