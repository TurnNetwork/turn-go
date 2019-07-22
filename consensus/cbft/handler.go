package cbft

import (
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"

	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/router"

	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"

	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
)

const (

	// Protocol name of CBFT
	CbftProtocolName = "cbft"

	// Protocol version of CBFT
	CbftProtocolVersion = 1

	// CbftProtocolLength are the number of implemented message corresponding to cbft protocol versions.
	CbftProtocolLength = 20

	// Maximum threshold for the queue of messages waiting to be sent.
	sendQueueSize = 10240

	QCBnMonitorInterval     = 4 // Qc block synchronization detection interval
	LockedBnMonitorInterval = 4 // Locked block synchronization detection interval
	CommitBnMonitorInterval = 4 // Commit block synchronization detection interval

	//
	TypeForQCBn     = 1
	TypeForLockedBn = 2
	TypeForCommitBn = 3
)

// Responsible for processing the messages in the network.
type EngineManager struct {
	engine    *Cbft
	peers     *router.PeerSet
	sendQueue chan *types.MsgPackage
	quitSend  chan struct{}
}

// Create a new handler and do some initialization.
func NewEngineManger(engine *Cbft) *EngineManager {
	return &EngineManager{
		engine:    engine,
		peers:     router.NewPeerSet(),
		sendQueue: make(chan *types.MsgPackage, sendQueueSize),
		quitSend:  make(chan struct{}, 0),
	}
}

// Start the loop to send message.
func (h *EngineManager) Start() {
	// Launch goroutine loop release separately.
	go h.sendLoop()
	go h.synchronize()
}

// Close turns off the handler for sending messages.
func (h *EngineManager) Close() {
	close(h.quitSend)
}

// The loop reads data from the message queue and sends it.
// If the message specifies the peerId then sends it directionally,
// and if the message does not specify peerId then broadcasts the message.
func (h *EngineManager) sendLoop() {
	for {
		select {
		case m := <-h.sendQueue:
			// todo: Need to add to the processing judgment of wal

			if len(m.PeerID()) == 0 {
				h.broadcast(m)
			} else {
				h.sendMessage(m)
			}
		case <-h.quitSend:
			log.Error("Terminate sending message")
			break
		}
	}
}

// Broadcast forwards the message to the router for distribution.
func (h *EngineManager) broadcast(m *types.MsgPackage) {
	h.engine.Router().Gossip(m)
}

// Send message to a known peerId. Determine if the peerId has established
// a connection before sending.
func (h *EngineManager) sendMessage(m *types.MsgPackage) {
	h.engine.Router().SendMessage(m)
}

// Return the peer with the specified peerID.
func (h *EngineManager) GetPeer(peerID string) (*router.Peer, error) {
	if peerID == "" {
		return nil, fmt.Errorf("invalid peerID parameter - %v", peerID)
	}
	return h.peers.Get(peerID)
}

// Send imports messages into the send queue and send it according to the specified ID.
func (h *EngineManager) Send(peerID discover.NodeID, msg types.Message) {
	msgPkg := types.NewMsgPackage(peerID.TerminalString(), msg, types.FullMode)
	select {
	case h.sendQueue <- msgPkg:
		log.Debug("Send message to sendQueue", "msgHash", msg.MsgHash().TerminalString(), "BHash", msg.BHash().TerminalString())
	}
	// todo: Whether to consider the problem of blocking
}

// Broadcast imports messages into the send queue and send it according to broadcast.
//
// Note: The broadcast of this method defaults to FULL mode.
func (h *EngineManager) Broadcast(msg types.Message) {
	msgPkg := types.NewMsgPackage("", msg, types.FullMode)
	select {
	case h.sendQueue <- msgPkg:
		log.Debug("Broadcast message to sendQueue", "msgHash", msg.MsgHash().TerminalString(), "BHash", msg.BHash().TerminalString())
	}
}

// Broadcast imports messages into the send queue.
//
// Note: The broadcast of this method defaults to PartMode.
func (h *EngineManager) PartBroadcast(msg types.Message) {
	msgPkg := types.NewMsgPackage("", msg, types.PartMode)
	select {
	case h.sendQueue <- msgPkg:
		log.Debug("PartBroadcast message to sendQueue", "msgHash", msg.MsgHash().TerminalString(), "BHash", msg.BHash().TerminalString())
	}
}

// Protocols implemented the Protocols method and returned basic information about the CBFT protocol.
func (h *EngineManager) Protocols() []p2p.Protocol {
	// todo: version and ProtocolLengths need to confirm.
	return []p2p.Protocol{
		{
			Name:    CbftProtocolName,
			Version: CbftProtocolVersion,
			Length:  CbftProtocolLength,
			Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
				return h.handler(p, rw)
			},
			NodeInfo: func() interface{} {
				return h.NodeInfo()
			},
			PeerInfo: func(id discover.NodeID) interface{} {
				if p, err := h.peers.Get(fmt.Sprintf("%5x", id[:8])); err == nil {
					return p.Info()
				}
				return nil
			},
		},
	}
}

// Return all neighbor node lists.
func (h *EngineManager) Peers() ([]*router.Peer, error) {
	return h.peers.Peers(), nil
}

// Return a peer by id.
func (h *EngineManager) Get(id string) (*router.Peer, error) {
	return h.peers.Get(id)
}

// Remove the peer with the specified ID
func (h *EngineManager) Unregister(id string) error {
	return h.peers.Unregister(id)
}

// Representative node configuration information.
type NodeInfo struct {
	Config Config `json:"config"`
}

func (h *EngineManager) NodeInfo() *NodeInfo {
	cfg := h.engine.Config()
	return &NodeInfo{
		Config: *cfg,
	}
}

// After the node is successfully connected and the message belongs
// to the cbft protocol message, the method is called.
func (h *EngineManager) handler(p *p2p.Peer, rw p2p.MsgReadWriter) error {
	// Further confirm if the version number needs to be read from the configuration.
	peer := router.NewPeer(CbftProtocolVersion, p, newMeteredMsgWriter(rw))

	// execute handshake
	// todo:
	// 1.need qcBn/qcHash/lockedBn/lockedHash/commitBn/commitHash from cbft.
	var (
		qcBn       = 1
		qcHash     = common.Hash{}
		lockedBn   = 1
		lockedHash = common.Hash{}
		commitBn   = 1
		commitHash = common.Hash{}
	)
	p.Log().Debug("CBFT peer connected, do handshake", "name", peer.Name())

	// Build a new CbftStatusData object as a handshake parameter
	cbftStatus := &protocols.CbftStatusData{
		ProtocolVersion: CbftProtocolVersion,
		QCBn:            new(big.Int).SetUint64(uint64(qcBn)),
		QCBlock:         qcHash,
		LockBn:          new(big.Int).SetUint64(uint64(lockedBn)),
		LockBlock:       lockedHash,
		CmtBn:           new(big.Int).SetUint64(uint64(commitBn)),
		CmtBlock:        commitHash,
	}
	// do handshake
	if err := peer.Handshake(cbftStatus); err != nil {
		p.Log().Debug("CBFT handshake failed", "err", err)
		return err
	} else {
		p.Log().Debug("CBFT consensus handshake success", "msgHash", cbftStatus.MsgHash().TerminalString())
	}

	// The newly established node is registered to the neighbor node list.
	if err := h.peers.Register(peer); err != nil {
		p.Log().Error("Cbft peer registration failed", "err", err)
		return err
	}
	defer h.peers.Unregister(peer.PeerID())

	// start ping loop.
	go peer.Run()

	// main loop. handle incoming message.
	// Exit the loop and disconnect if the message
	// is processing abnormally.
	for {
		if err := h.handleMsg(peer); err != nil {
			p.Log().Error("CBFT message handling failed", "err", err)
			return err
		}
	}
}

// Main logic: Distribute according to message type and
// transfer message to CBFT layer
func (h *EngineManager) handleMsg(p *router.Peer) error {
	msg, err := p.ReadWriter().ReadMsg()
	if err != nil {
		p.Log().Error("read peer message error", "err", err)
		return err
	}

	// All messages cannot exceed the maximum specified by the agreement.
	if msg.Size > protocols.CbftProtocolMaxMsgSize {
		return types.ErrResp(types.ErrMsgTooLarge, "%v > %v", msg.Size, protocols.CbftProtocolMaxMsgSize)
	}
	defer msg.Discard()

	// Handle the message depending on msgType and it's content.
	switch {
	case msg.Code == protocols.CBFTStatusMsg:
		// CBFTStatusMsg belongs to the type of handshake message and will not appear here.
		return types.ErrResp(types.ErrExtraStatusMsg, "uncontrolled status message")

	case msg.Code == protocols.PrepareBlockMsg:
		var request protocols.PrepareBlock
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		p.MarkMessageHash((&request).MsgHash())
		request.Block.ReceivedAt = msg.ReceivedAt
		request.Block.ReceivedFrom = p
		// Message transfer to cbft message queue.
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.PrepareVoteMsg:
		var request protocols.PrepareVote
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		p.MarkMessageHash((&request).MsgHash())
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.ViewChangeMsg:
		var request protocols.ViewChange
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		p.MarkMessageHash((&request).MsgHash())
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.GetPrepareBlockMsg:
		var request protocols.GetPrepareBlock
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.GetQuorumCertMsg:
		var request protocols.GetQuorumCert
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.BlockQuorumCertMsg:
		var request protocols.BlockQuorumCert
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		p.MarkMessageHash((&request).MsgHash())
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.GetQCPrepareBlockMsg:
		var request protocols.GetQCPrepareBlock
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.QCPrepareBlockMsg:
		var request protocols.QCPrepareBlock
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.GetPrepareVoteMsg:
		var request protocols.GetPrepareVote
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.PrepareBlockHashMsg:
		var request protocols.PrepareBlockHash
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		p.MarkMessageHash((&request).MsgHash())
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.PrepareVotesMsg:
		var request protocols.PrepareVotes
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.QCBlockListMsg:
		var request protocols.QCBlockList
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		h.engine.ReceiveMessage(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.GetLatestStatusMsg:
		var request protocols.GetLatestStatus
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		h.engine.ReceiveSyncMsg(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.LatestStatusMsg:
		var request protocols.LatestStatus
		if err := msg.Decode(&request); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		h.engine.ReceiveSyncMsg(types.NewMessage(&request, p.PeerID()))
		return nil

	case msg.Code == protocols.PingMsg:
		var pingTime protocols.Ping
		if err := msg.Decode(&pingTime); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		// Directly respond to the response message after receiving the ping message.
		go p2p.SendItems(p.ReadWriter(), protocols.PongMsg, pingTime[0])
		p.Log().Trace("Respond to ping message done")

	case msg.Code == protocols.PongMsg:
		// Processed after receiving the pong message.
		curTime := time.Now().UnixNano()
		log.Debug("handle a eth Pong message", "curTime", curTime)
		var pingTime protocols.Pong
		if err := msg.Decode(&pingTime); err != nil {
			return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
		}
		for {
			// Return the first element of list l or nil if the list is empty.
			frontPing := p.PingList.Front()
			if frontPing == nil {
				log.Trace("end of p.PingList")
				break
			}
			log.Trace("Front element of p.PingList", "element", frontPing)
			if t, ok := p.PingList.Remove(frontPing).(string); ok {
				if t == pingTime[0] {
					tInt64, err := strconv.ParseInt(t, 10, 64)
					if err != nil {
						return types.ErrResp(types.ErrDecode, "%v: %v", msg, err)
					}
					log.Trace("calculate net latency", "sendPingTime", tInt64, "receivePongTime", curTime)
					latency := (curTime - tInt64) / 2 / 1000000
					h.engine.OnPong(p.Peer.ID(), latency)
					break
				}
			}
		}
		return nil

	default:
		return types.ErrResp(types.ErrInvalidMsgCode, "%v", msg.Code)
	}

	return nil
}

// Select a node with a height higher than the local node block from
// the neighbor node list, and then synchronize the block data of
// the height difference to the node.
//
// Note:
// 1. Synchronous blocks with inconsistent QC height.
// 2. Synchronous blocks with inconsistent locking block height.
// 3. Synchronous blocks with inconsistent commit block height.
func (h *EngineManager) synchronize() {
	log.Debug("~ Start synchronize in the handler")
	qcTicker := time.NewTicker(QCBnMonitorInterval * time.Second)
	lockedTicker := time.NewTicker(LockedBnMonitorInterval * time.Second)
	commitTicker := time.NewTicker(CommitBnMonitorInterval * time.Second)

	for {
		select {
		case <-qcTicker.C:
			qcBn := h.engine.HighestQCBlockBn()
			highPeers := h.peers.PeersWithHighestQCBn(qcBn)
			biggestPeer, biggestNumber := largerPeer(TypeForQCBn, highPeers, qcBn)
			if biggestPeer != nil {
				log.Debug("Synchronize for qc block send message", "localQCBn", qcBn, "remoteQCBn", biggestNumber, "remotePeerID", biggestPeer.PeerID())
				// todo: Build a message and then send a message
			}
		case <-lockedTicker.C:
			lockedBn := h.engine.HighestLockBlockBn()
			highPeers := h.peers.PeersWithHighestLockedBn(lockedBn)
			biggestPeer, biggestNumber := largerPeer(TypeForLockedBn, highPeers, lockedBn)
			if biggestPeer != nil {
				log.Debug("Synchronize for locked block send message", "localLockedBn", lockedBn, "remoteLockedBn", biggestNumber, "remotePeerID", biggestPeer.PeerID())
				// todo: Build a message and then send a message
			}
		case <-commitTicker.C:
			commitBn := h.engine.HighestCommitBlockBn()
			highPeers := h.peers.PeersWithHighestCommitBn(commitBn)
			biggestPeer, biggestNumber := largerPeer(TypeForCommitBn, highPeers, commitBn)
			if biggestPeer != nil {
				log.Debug("Synchronize for locked block send message", "localCommitBn", commitBn, "remoteCommitBn", biggestNumber, "remotePeerID", biggestPeer.PeerID())
				// todo: Build a message and then send a message
			}
		case <-h.quitSend:
			log.Error("synchronize quit")
			return
		}
	}
}

// Select a node from the list of nodes that is larger than the specified value.
//
// bType:
//  1 -> qcBlock, 2 -> lockedBlock, 3 -> CommitBlock
func largerPeer(bType uint64, peers []*router.Peer, number uint64) (*router.Peer, uint64) {
	if peers != nil && len(peers) != 0 {
		return nil, 0
	}
	largerNum := number
	largerIndex := -1
	for index, v := range peers {
		var pNumber uint64
		switch bType {
		case TypeForQCBn:
			pNumber = v.QCBn()
		case TypeForLockedBn:
			pNumber = v.LockedBn()
		case TypeForCommitBn:
			pNumber = v.CommitBn()
		default:
			return nil, 0
		}
		if pNumber > largerNum {
			largerNum, largerIndex = pNumber, index
		}
	}
	if largerIndex != -1 {
		return peers[largerIndex], largerNum
	}
	return nil, 0
}
