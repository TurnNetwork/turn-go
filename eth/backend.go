// Copyright 2014 The go-ethereum Authors
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

// Package eth implements the Ethereum protocol.
package eth

import (
	"bufio"
	"errors"
	"fmt"
	"math/big"
	"net"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/bubblenet/bubble/consensus/cbft/wal"

	"github.com/bubblenet/bubble/x/gov"

	"github.com/bubblenet/bubble/x/handler"

	"github.com/bubblenet/bubble/core/snapshotdb"

	"github.com/bubblenet/bubble/consensus/cbft/evidence"

	"github.com/bubblenet/bubble/accounts"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/consensus"
	"github.com/bubblenet/bubble/consensus/cbft"
	ctypes "github.com/bubblenet/bubble/consensus/cbft/types"
	"github.com/bubblenet/bubble/consensus/cbft/validator"
	"github.com/bubblenet/bubble/core"
	"github.com/bubblenet/bubble/core/bloombits"
	"github.com/bubblenet/bubble/core/rawdb"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/core/vm"
	"github.com/bubblenet/bubble/eth/downloader"
	"github.com/bubblenet/bubble/eth/filters"
	"github.com/bubblenet/bubble/eth/gasprice"
	"github.com/bubblenet/bubble/ethdb"
	"github.com/bubblenet/bubble/event"
	"github.com/bubblenet/bubble/internal/ethapi"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/miner"
	"github.com/bubblenet/bubble/node"
	"github.com/bubblenet/bubble/p2p"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/rpc"
	xplugin "github.com/bubblenet/bubble/x/plugin"
	"github.com/bubblenet/bubble/x/xcom"
)

// Ethereum implements the Ethereum full node service.
type Ethereum struct {
	config *Config

	// Handlers
	txPool          *core.TxPool
	blockchain      *core.BlockChain
	protocolManager *ProtocolManager

	// DB interfaces
	chainDb ethdb.Database // Block chain database

	eventMux       *event.TypeMux
	engine         consensus.Engine
	accountManager *accounts.Manager

	bloomRequests     chan chan *bloombits.Retrieval // Channel receiving bloom data retrieval requests
	bloomIndexer      *core.ChainIndexer             // Bloom indexer operating during block imports
	closeBloomHandler chan struct{}

	APIBackend *EthAPIBackend

	miner         *miner.Miner
	gasPrice      *big.Int
	networkID     uint64
	netRPCService *ethapi.PublicNetAPI

	p2pServer *p2p.Server

	lock sync.RWMutex // Protects the variadic fields (e.g. gas price and etherbase)
}

// New creates a new Ethereum object (including the
// initialisation of the common Ethereum object)
func New(stack *node.Node, config *Config) (*Ethereum, error) {
	// Ensure configuration values are compatible and sane
	if config.SyncMode == downloader.LightSync {
		return nil, errors.New("can't run bubble in light sync mode, use les.LightBubble")
	}
	if !config.SyncMode.IsValid() {
		return nil, fmt.Errorf("invalid sync mode %d", config.SyncMode)
	}
	if config.Miner.GasPrice == nil || config.Miner.GasPrice.Cmp(common.Big0) <= 0 {
		log.Warn("Sanitizing invalid miner gas price", "provided", config.Miner.GasPrice, "updated", DefaultConfig.Miner.GasPrice)
		config.Miner.GasPrice = new(big.Int).Set(DefaultConfig.Miner.GasPrice)
	}
	// Assemble the Ethereum object
	chainDb, err := stack.OpenDatabaseWithFreezer("chaindata", config.DatabaseCache, config.DatabaseHandles, config.DatabaseFreezer, "eth/db/chaindata/")
	if err != nil {
		return nil, err
	}
	snapshotdb.SetDBOptions(config.DatabaseCache, config.DatabaseHandles)

	snapshotBaseDB, err := snapshotdb.Open(stack.ResolvePath(snapshotdb.DBPath), config.DatabaseCache, config.DatabaseHandles, true)
	if err != nil {
		return nil, err
	}

	height := rawdb.ReadHeaderNumber(chainDb, rawdb.ReadHeadHeaderHash(chainDb))
	log.Debug("read header number from chain db", "height", height)
	if height != nil && *height > 0 {
		//when last  fast syncing fail,we will clean chaindb,wal,snapshotdb
		status, err := snapshotBaseDB.GetBaseDB([]byte(downloader.KeyFastSyncStatus))

		// systemError
		if err != nil && err != snapshotdb.ErrNotFound {
			if err := snapshotBaseDB.Close(); err != nil {
				return nil, err
			}
			return nil, err
		}
		//if find sync status,this means last syncing not finish,should clean all db to reinit
		//if not find sync status,no need init chain
		if err == nil {

			// Just commit the new block if there is no stored genesis block.
			stored := rawdb.ReadCanonicalHash(chainDb, 0)

			log.Info("last fast sync is fail,init  db", "status", common.BytesToUint32(status), "prichain", config.Genesis == nil)
			chainDb.Close()
			if err := snapshotBaseDB.Close(); err != nil {
				return nil, err
			}

			if config.DatabaseFreezer != "" {
				if err := os.RemoveAll(stack.Config().ResolveFreezerPath("chaindata", config.DatabaseFreezer)); err != nil {
					return nil, err
				}
			}

			if err := os.RemoveAll(stack.ResolvePath("chaindata")); err != nil {
				return nil, err
			}

			if err := os.RemoveAll(stack.ResolvePath(wal.WalDir(stack))); err != nil {
				return nil, err
			}

			if err := os.RemoveAll(stack.ResolvePath(snapshotdb.DBPath)); err != nil {
				return nil, err
			}

			chainDb, err = stack.OpenDatabaseWithFreezer("chaindata", config.DatabaseCache, config.DatabaseHandles, config.DatabaseFreezer, "eth/db/chaindata/")
			if err != nil {
				return nil, err
			}

			snapshotBaseDB, err = snapshotdb.Open(stack.ResolvePath(snapshotdb.DBPath), config.DatabaseCache, config.DatabaseHandles, true)
			if err != nil {
				return nil, err
			}

			//only private net  need InitGenesisAndSetEconomicConfig
			if stored != params.MainnetGenesisHash && config.Genesis == nil {
				// private net
				config.Genesis = new(core.Genesis)
				if err := config.Genesis.InitGenesisAndSetEconomicConfig(stack.GenesisPath()); err != nil {
					return nil, err
				}
			}
			log.Info("last fast sync is fail,init  db finish")
		}
	}

	chainConfig, opConfig, _, genesisErr := core.SetupGenesisBlock(chainDb, snapshotBaseDB, config.Genesis)
	if err := snapshotBaseDB.Close(); err != nil {
		return nil, err
	}

	if _, ok := genesisErr.(*params.ConfigCompatError); genesisErr != nil && !ok {
		return nil, genesisErr
	}
	// Configuring frps
	if nil != chainConfig && nil != chainConfig.Frps && nil != chainConfig.Cbft {
		nodeCfg := stack.Config()
		// Local node ID
		nodeId := discover.PubkeyID(&nodeCfg.P2P.PrivateKey.PublicKey)
		// 1.Assemble and generate the frpc profile
		// 1.1 Create frp configuration file and assemble frp server configuration information and p2p listening information of local node
		listenAddr := nodeCfg.P2P.ListenAddr
		dataDir := nodeCfg.DataDir
		err, file, writer, filePath := genFrpCfgFile(chainConfig.Frps, listenAddr, dataDir, nodeId)
		defer file.Close()
		if err != nil && nil != writer {
			fmt.Println("failed to generate config file:", err)
			return nil, err
		}
		// The frp configuration is initialized according to the set maximum number of peer connections
		maxCount := len(chainConfig.Cbft.InitialNodes)
		if maxCount > nodeCfg.P2P.MaxPeers {
			maxCount = nodeCfg.P2P.MaxPeers
		}
		for i := 0; i < maxCount; i++ {
			cbftNode := &chainConfig.Cbft.InitialNodes[i].Node
			// 1.2 Assemble visitor information for connecting to other peers, Only nodes with peer-to-peer p2p port 0 are processed
			if cbftNode.ID != nodeId && (0 == cbftNode.TCP || 0 == cbftNode.UDP) {
				if err := addFrpVisitor(cbftNode, writer); nil != err {
					return nil, err
				}
			}
		}
		// 1.3 Add the rpc proxy configuration
		if 0 != nodeCfg.HTTPPort && 0 != nodeCfg.ProxyRpcPort {
			if err := addFrpProxy(nodeCfg.HTTPPort, nodeCfg.ProxyRpcPort, "rpc_proxy", writer); nil != err {
				return nil, err
			}
		}

		// 2.Save the configuration file path to the node configuration
		stack.Server().FrpFilePath = filePath
	}
	if chainConfig.Cbft.Period == 0 || chainConfig.Cbft.Amount == 0 {
		chainConfig.Cbft.Period = config.CbftConfig.Period
		chainConfig.Cbft.Amount = config.CbftConfig.Amount
	}

	log.Info("Initialised chain configuration", "config", chainConfig)
	stack.SetP2pChainID(chainConfig.ChainID)

	eth := &Ethereum{
		config:            config,
		chainDb:           chainDb,
		eventMux:          stack.EventMux(),
		accountManager:    stack.AccountManager(),
		engine:            CreateConsensusEngine(stack, chainConfig, config.Miner.Noverify, chainDb, &config.CbftConfig, stack.EventMux()),
		closeBloomHandler: make(chan struct{}),
		networkID:         config.NetworkId,
		gasPrice:          config.Miner.GasPrice,
		bloomRequests:     make(chan chan *bloombits.Retrieval),
		bloomIndexer:      NewBloomIndexer(chainDb, params.BloomBitsBlocks, params.BloomConfirms),
		p2pServer:         stack.Server(),
	}

	bcVersion := rawdb.ReadDatabaseVersion(chainDb)

	var dbVer = "<nil>"
	if bcVersion != nil {
		dbVer = fmt.Sprintf("%d", *bcVersion)
	}
	log.Info("Initialising bubble protocol", "versions", ProtocolVersions, "network", config.NetworkId, "dbversion", dbVer)

	if !config.SkipBcVersionCheck {
		if bcVersion != nil && *bcVersion > core.BlockChainVersion {
			return nil, fmt.Errorf("database version is v%d, bubble %s only supports v%d", *bcVersion, params.VersionWithMeta, core.BlockChainVersion)
		} else if bcVersion == nil || *bcVersion < core.BlockChainVersion {
			log.Warn("Upgrade blockchain database version", "from", dbVer, "to", core.BlockChainVersion)
			rawdb.WriteDatabaseVersion(chainDb, core.BlockChainVersion)
		}
	}

	var (
		vmConfig = vm.Config{
			ConsoleOutput: config.Debug,
			WasmType:      vm.Str2WasmType(config.VMWasmType),
		}
		cacheConfig = &core.CacheConfig{Disabled: config.NoPruning, TrieDirtyLimit: config.TrieCache, TrieTimeLimit: config.TrieTimeout,
			BodyCacheLimit: config.BodyCacheLimit, BlockCacheLimit: config.BlockCacheLimit,
			MaxFutureBlocks: config.MaxFutureBlocks, BadBlockLimit: config.BadBlockLimit,
			TriesInMemory: config.TriesInMemory, TrieCleanLimit: config.TrieDBCache, Preimages: config.Preimages,
			TrieCleanJournal:   stack.ResolvePath(config.TrieCleanCacheJournal),
			TrieCleanRejournal: config.TrieCleanCacheRejournal,
			DBGCInterval:       config.DBGCInterval, DBGCTimeout: config.DBGCTimeout,
			DBGCMpt: config.DBGCMpt, DBGCBlock: config.DBGCBlock,
		}

		minningConfig = &core.MiningConfig{MiningLogAtDepth: config.MiningLogAtDepth, TxChanSize: config.TxChanSize,
			ChainHeadChanSize: config.ChainHeadChanSize, ChainSideChanSize: config.ChainSideChanSize,
			ResultQueueSize: config.ResultQueueSize, ResubmitAdjustChanSize: config.ResubmitAdjustChanSize,
			MinRecommitInterval: config.MinRecommitInterval, MaxRecommitInterval: config.MaxRecommitInterval,
			IntervalAdjustRatio: config.IntervalAdjustRatio, IntervalAdjustBias: config.IntervalAdjustBias,
			StaleThreshold: config.StaleThreshold, DefaultCommitRatio: config.DefaultCommitRatio,
		}
	)
	cacheConfig.DBDisabledGC.Set(config.DBDisabledGC)

	eth.blockchain, err = core.NewBlockChain(chainDb, cacheConfig, chainConfig, eth.engine, vmConfig, eth.shouldPreserve, &config.TxLookupLimit)
	if err != nil {
		return nil, err
	}
	snapshotdb.SetDBBlockChain(eth.blockchain)

	blockChainCache := core.NewBlockChainCache(eth.blockchain)

	// Rewind the chain in case of an incompatible config upgrade.
	if compat, ok := genesisErr.(*params.ConfigCompatError); ok {
		log.Warn("Rewinding chain to upgrade configuration", "err", compat)
		return nil, compat
		//eth.blockchain.SetHead(compat.RewindTo)
		//rawdb.WriteChainConfig(chainDb, genesisHash, chainConfig)
	}
	eth.bloomIndexer.Start(eth.blockchain)

	if config.TxPool.Journal != "" {
		config.TxPool.Journal = stack.ResolvePath(config.TxPool.Journal)
	}
	eth.txPool = core.NewTxPool(config.TxPool, chainConfig, core.NewTxPoolBlockChain(blockChainCache))

	core.SenderCacher.SetTxPool(eth.txPool)

	currentBlock := eth.blockchain.CurrentBlock()
	currentNumber := currentBlock.NumberU64()
	currentHash := currentBlock.Hash()
	gasCeil, err := gov.GovernMaxBlockGasLimit(currentNumber, currentHash)
	if nil != err {
		log.Error("Failed to query gasCeil from snapshotdb", "err", err)
		return nil, err
	}
	if config.Miner.GasFloor > uint64(gasCeil) {
		log.Error("The gasFloor must be less than gasCeil", "gasFloor", config.Miner.GasFloor, "gasCeil", gasCeil)
		return nil, fmt.Errorf("The gasFloor must be less than gasCeil, got: %d, expect range (0, %d]", config.Miner.GasFloor, gasCeil)
	}

	eth.miner = miner.New(eth, &config.Miner, eth.blockchain.Config(), minningConfig, eth.EventMux(), eth.engine,
		eth.isLocalBlock, blockChainCache, config.VmTimeoutDuration)

	reactor := core.NewBlockChainReactor(eth.EventMux(), eth.blockchain.Config().ChainID)
	node.GetCryptoHandler().SetPrivateKey(stack.Config().NodeKey())

	if engine, ok := eth.engine.(consensus.Bft); ok {
		var agency consensus.Agency
		core.NewExecutor(eth.blockchain.Config(), eth.blockchain, vmConfig, eth.txPool)
		// validatorMode:
		// - static (default)
		// - inner (via inner contract)eth/handler.go
		// - dpos

		log.Debug("Validator mode", "mode", chainConfig.Cbft.ValidatorMode)
		if chainConfig.Cbft.ValidatorMode == "" || chainConfig.Cbft.ValidatorMode == common.STATIC_VALIDATOR_MODE {
			agency = validator.NewStaticAgency(chainConfig.Cbft.InitialNodes)
			reactor.Start(common.STATIC_VALIDATOR_MODE)
		} else if chainConfig.Cbft.ValidatorMode == common.INNER_VALIDATOR_MODE {
			blocksPerNode := int(chainConfig.Cbft.Amount)
			offset := blocksPerNode * 2
			agency = validator.NewInnerAgency(chainConfig.Cbft.InitialNodes, eth.blockchain, blocksPerNode, offset)
			reactor.Start(common.INNER_VALIDATOR_MODE)
		} else if chainConfig.Cbft.ValidatorMode == common.DPOS_VALIDATOR_MODE {
			reactor.Start(common.DPOS_VALIDATOR_MODE)
			reactor.SetVRFhandler(handler.NewVrfHandler(eth.blockchain.Genesis().Nonce()))
			reactor.SetPluginEventMux()
			reactor.SetPrivateKey(stack.Config().NodeKey())
			if err := opConfig.SetSubOpPriKey(config.SubOpPriKey); err != nil {
				return nil, errors.New("failed to set the private key of child-chain operation address")
			}
			handlePlugin(reactor, chainDb, config.DBValidatorsHistory, opConfig, chainConfig.ChainID)
			agency = reactor

			//register Govern parameter verifiers
			gov.RegisterGovernParamVerifiers()
		}

		if err := recoverSnapshotDB(blockChainCache); err != nil {
			log.Error("recover SnapshotDB fail", "error", err)
			return nil, errors.New("Failed to recover SnapshotDB")
		}

		if err := engine.Start(eth.blockchain, blockChainCache, eth.txPool, agency); err != nil {
			log.Error("Init cbft consensus engine fail", "error", err)
			return nil, errors.New("Failed to init cbft consensus engine")
		}
	}

	// Permit the downloader to use the trie cache allowance during fast sync
	cacheLimit := cacheConfig.TrieCleanLimit + cacheConfig.TrieDirtyLimit
	if eth.protocolManager, err = NewProtocolManager(chainConfig, config.SyncMode, config.NetworkId, eth.eventMux, eth.txPool, eth.engine, eth.blockchain, chainDb, cacheLimit); err != nil {
		return nil, err
	}
	eth.APIBackend = &EthAPIBackend{stack.Config().ExtRPCEnabled(), eth, nil}
	gpoParams := config.GPO
	if gpoParams.Default == nil {
		gpoParams.Default = config.Miner.GasPrice
	}
	eth.APIBackend.gpo = gasprice.NewOracle(eth.APIBackend, gpoParams)
	// Start the RPC service
	eth.netRPCService = ethapi.NewPublicNetAPI(eth.p2pServer, eth.NetVersion())

	// Register the backend on the node
	stack.RegisterAPIs(eth.APIs())
	stack.RegisterProtocols(eth.Protocols())
	stack.RegisterLifecycle(eth)
	return eth, nil
}

// Generate the frp profile
func genFrpCfgFile(frps *params.FrpsConfig, listenAddr, dataDir string, nodeId discover.NodeID) (error, *os.File, *bufio.Writer, string) {
	// Create a new INI file
	frpDir := dataDir + "/bubble/frp/"
	if err := os.MkdirAll(frpDir, 0700); err != nil {
		log.Error(fmt.Sprintf("Failed to create directory: %v", err))
		return err, nil, nil, ""
	}
	fileName := "config.ini"
	filePath := common.AbsolutePath(frpDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		log.Error("failed to create Frp config file:", err)
		return err, file, nil, filePath
	}

	if nil == frps {
		log.Error("The frp server configuration is nil.")
		return errors.New("The frp server configuration is nil"), file, nil, filePath
	}

	// Create a writer
	writer := bufio.NewWriter(file)

	// Write the frp server configuration
	fmt.Fprintln(writer, "[common]")
	fmt.Fprintln(writer, "server_addr =", frps.ServerIP)
	fmt.Fprintln(writer, "server_port =", frps.ServerPort)
	if nil != frps.Auth {
		fmt.Fprintln(writer, "authentication_method =", frps.Auth.Method)
		fmt.Fprintln(writer, "authenticate_heartbeats =", frps.Auth.HeartBeats)
		fmt.Fprintln(writer, "authenticate_new_work_conns =", frps.Auth.NewWorkConns)
		fmt.Fprintln(writer, "token =", frps.Auth.Token)
	}
	fmt.Fprintln(writer, "")

	// Write the frp configuration that the node listens to
	startIndex := strings.Index(listenAddr, ":")
	if startIndex != -1 {
		listenAddr = listenAddr[startIndex+1:]
	}
	shortKey := nodeId.ShortString()
	nodeLabel := "[" + shortKey + "]"
	fmt.Fprintln(writer, nodeLabel)
	fmt.Fprintln(writer, "type = xtcp")
	fmt.Fprintln(writer, "sk =", shortKey)
	fmt.Fprintln(writer, "local_ip = 127.0.0.1")
	fmt.Fprintln(writer, "local_port =", listenAddr)
	fmt.Fprintln(writer, "use_encryption = false")
	fmt.Fprintln(writer, "use_compression = false")
	fmt.Fprintln(writer, "")

	// Flush the buffer and write the data to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Failure to flush the buffer and write data to the file:", err)
		return err, file, nil, filePath
	}

	return nil, file, writer, filePath
}

// Add the visitor entry to the frp configuration file
func addFrpVisitor(cbftNode *discover.Node, writer *bufio.Writer) error {
	if nil == cbftNode || nil == writer {
		return nil
	}
	// Create a TCP listener that listens on a random local port
	listener, err := net.Listen("tcp", "localhost:0")
	defer listener.Close()
	if err != nil {
		log.Error("Failed to get an unused port:", err)
		return err
	}

	// Get the local address of the listener
	address := listener.Addr().(*net.TCPAddr)
	// log.Debug("Available local address: %s:%d\n", address.IP.String(), address.Port)
	// Set the p2p port number
	cbftNode.UDP = uint16(address.Port)
	cbftNode.TCP = uint16(address.Port)

	visitorNode := cbftNode.ID.ShortString()
	visitorLabel := "[" + visitorNode + "_visitor]"
	fmt.Fprintln(writer, visitorLabel)
	fmt.Fprintln(writer, "type = xtcp")
	fmt.Fprintln(writer, "role = visitor")
	fmt.Fprintln(writer, "server_name =", visitorNode)
	fmt.Fprintln(writer, "sk =", visitorNode)
	fmt.Fprintln(writer, "bind_addr = 127.0.0.1")
	fmt.Fprintln(writer, "bind_port =", address.Port)
	//fmt.Fprintln(writer, "use_encryption = false")
	//fmt.Fprintln(writer, "use_compression = false")
	fmt.Fprintln(writer, "")
	err = writer.Flush()
	if err != nil {
		fmt.Println("Failure to flush the buffer and write data to the file:", err)
		return err
	}
	return nil
}

// Add the frpc agent configuration
func addFrpProxy(localPort, proxyPort int, proxyName string, writer *bufio.Writer) error {
	labelName := "[" + proxyName + "]"
	fmt.Fprintln(writer, labelName)
	fmt.Fprintln(writer, "type = tcp")
	fmt.Fprintln(writer, "local_ip = 127.0.0.1")
	fmt.Fprintln(writer, "local_port =", localPort)
	fmt.Fprintln(writer, "remote_port =", proxyPort)
	fmt.Fprintln(writer, "")
	if err := writer.Flush(); err != nil {
		fmt.Println("Failure to flush the buffer and write data to the file:", err)
		return err
	}
	return nil
}

func recoverSnapshotDB(blockChainCache *core.BlockChainCache) error {
	sdb := snapshotdb.Instance()
	ch := sdb.GetCurrent().GetHighest(false).Num.Uint64()
	blockChanHegiht := blockChainCache.CurrentHeader().Number.Uint64()
	if ch < blockChanHegiht {
		for i := ch + 1; i <= blockChanHegiht; i++ {
			block, parentBlock := blockChainCache.GetBlockByNumber(i), blockChainCache.GetBlockByNumber(i-1)
			log.Debug("snapshotdb recover block from blockchain", "num", block.Number(), "hash", block.Hash())
			if err := blockChainCache.Execute(block, parentBlock); err != nil {
				log.Error("snapshotdb recover block from blockchain  execute fail", "error", err)
				return err
			}
			if err := sdb.Commit(block.Hash()); err != nil {
				log.Error("snapshotdb recover block from blockchain  Commit fail", "error", err)
				return err
			}
		}
	}
	return nil
}

// CreateConsensusEngine creates the required type of consensus engine instance for an Ethereum service
func CreateConsensusEngine(stack *node.Node, chainConfig *params.ChainConfig, noverify bool, db ethdb.Database,
	cbftConfig *ctypes.OptionsConfig, eventMux *event.TypeMux) consensus.Engine {
	// If proof-of-authority is requested, set it up
	engine := cbft.New(chainConfig.Cbft, cbftConfig, eventMux, stack)
	if engine == nil {
		panic("create consensus engine fail")
	}
	return engine
}

// APIs return the collection of RPC services the ethereum package offers.
// NOTE, some of these services probably need to be moved to somewhere else.
func (s *Ethereum) APIs() []rpc.API {
	apis := ethapi.GetAPIs(s.APIBackend)

	// Append any APIs exposed explicitly by the consensus engine
	apis = append(apis, s.engine.APIs(s.BlockChain())...)

	// Append all the local APIs and return
	return append(apis, []rpc.API{
		{
			Namespace: "bub",
			Version:   "1.0",
			Service:   downloader.NewPublicDownloaderAPI(s.protocolManager.downloader, s.eventMux),
			Public:    true,
		}, {
			Namespace: "miner",
			Version:   "1.0",
			Service:   NewPrivateMinerAPI(s),
			Public:    false,
		}, {
			Namespace: "bub",
			Version:   "1.0",
			Service:   filters.NewPublicFilterAPI(s.APIBackend, false),
			Public:    true,
		}, {
			Namespace: "admin",
			Version:   "1.0",
			Service:   NewPrivateAdminAPI(s),
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(s),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(s),
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   xplugin.NewPublicDPOSAPI(),
		}, {
			Namespace: "net",
			Version:   "1.0",
			Service:   s.netRPCService,
			Public:    true,
		},
		{
			Namespace: "txgen",
			Version:   "1.0",
			Service:   NewTxGenAPI(s),
			Public:    true,
		},
	}...)
}

//func (s *Ethereum) ResetWithGenesisBlock(gb *types.Block) {
//	s.blockchain.ResetWithGenesisBlock(gb)
//}

// isLocalBlock checks whether the specified block is mined
// by local miner accounts.
//
// We regard two types of accounts as local miner account: etherbase
// and accounts specified via `txpool.locals` flag.
func (s *Ethereum) isLocalBlock(block *types.Block) bool {
	author, err := s.engine.Author(block.Header())
	if err != nil {
		log.Warn("Failed to retrieve block author", "number", block.NumberU64(), "hash", block.Hash(), "err", err)
		return false
	}
	// Check whether the given address is etherbase.
	s.lock.RLock()
	etherbase := common.Address{}
	s.lock.RUnlock()
	if author == etherbase {
		return true
	}
	// Check whether the given address is specified by `txpool.local`
	// CLI flag.
	for _, account := range s.config.TxPool.Locals {
		if account == author {
			return true
		}
	}
	return false
}

// shouldPreserve checks whether we should preserve the given block
// during the chain reorg depending on whether the author of block
// is a local account.
func (s *Ethereum) shouldPreserve(block *types.Block) bool {
	// The reason we need to disable the self-reorg preserving for clique
	// is it can be probable to introduce a deadlock.
	//
	// e.g. If there are 7 available signers
	//
	// r1   A
	// r2     B
	// r3       C
	// r4         D
	// r5   A      [X] F G
	// r6    [X]
	//
	// In the round5, the inturn signer E is offline, so the worst case
	// is A, F and G sign the block of round5 and reject the block of opponents
	// and in the round6, the last available signer B is offline, the whole
	// network is stuck.
	return s.isLocalBlock(block)
}

// start mining
func (s *Ethereum) StartMining() error {
	// If the miner was not running, initialize it
	if !s.IsMining() {
		// Propagate the initial price point to the transaction pool
		s.lock.RLock()
		price := s.gasPrice
		s.lock.RUnlock()
		s.txPool.SetGasPrice(price)

		// If mining is started, we can disable the transaction rejection mechanism
		// introduced to speed sync times.
		atomic.StoreUint32(&s.protocolManager.acceptTxs, 1)

		go s.miner.Start()
	}
	return nil
}

// StopMining terminates the miner, both at the consensus engine level as well as
// at the block creation level.
func (s *Ethereum) StopMining() {
	s.miner.Stop()
}

func (s *Ethereum) IsMining() bool      { return s.miner.Mining() }
func (s *Ethereum) Miner() *miner.Miner { return s.miner }

func (s *Ethereum) AccountManager() *accounts.Manager  { return s.accountManager }
func (s *Ethereum) BlockChain() *core.BlockChain       { return s.blockchain }
func (s *Ethereum) TxPool() *core.TxPool               { return s.txPool }
func (s *Ethereum) EventMux() *event.TypeMux           { return s.eventMux }
func (s *Ethereum) Engine() consensus.Engine           { return s.engine }
func (s *Ethereum) ChainDb() ethdb.Database            { return s.chainDb }
func (s *Ethereum) IsListening() bool                  { return true } // Always listening
func (s *Ethereum) EthVersion() int                    { return int(s.protocolManager.SubProtocols[0].Version) }
func (s *Ethereum) NetVersion() uint64                 { return s.networkID }
func (s *Ethereum) Downloader() *downloader.Downloader { return s.protocolManager.downloader }
func (s *Ethereum) BloomIndexer() *core.ChainIndexer   { return s.bloomIndexer }

// Protocols returns all the currently configured
// network protocols to start.
func (s *Ethereum) Protocols() []p2p.Protocol {
	protocols := make([]p2p.Protocol, 0)
	protocols = append(protocols, s.protocolManager.SubProtocols...)
	protocols = append(protocols, s.engine.Protocols()...)

	return protocols
}

// Start implements node.Lifecycle, starting all internal goroutines needed by the
// Ethereum protocol implementation.
func (s *Ethereum) Start() error {
	// Start the bloom bits servicing goroutines
	s.startBloomHandlers(params.BloomBitsBlocks)

	// Figure out a max peers count based on the server limits
	maxPeers := s.p2pServer.MaxPeers
	// Start the networking layer and the light server if requested
	s.protocolManager.Start(maxPeers)

	//log.Debug("node start", "srvr.Config.PrivateKey", srvr.Config.PrivateKey)
	if cbftEngine, ok := s.engine.(consensus.Bft); ok {
		if flag := cbftEngine.IsConsensusNode(); flag {
			for _, n := range s.blockchain.Config().Cbft.InitialNodes {
				// todo: Mock point.
				if !node.FakeNetEnable {
					s.p2pServer.AddConsensusPeer(discover.NewNode(n.Node.ID, n.Node.IP, n.Node.UDP, n.Node.TCP))
				}
			}
		}
		s.StartMining()
	}
	s.p2pServer.StartWatching(s.eventMux)

	return nil
}

// Stop implements node.Service, terminating all internal goroutines used by the
// Ethereum protocol.
func (s *Ethereum) Stop() error {
	s.protocolManager.Stop()

	// Then stop everything else.
	// Only the operations related to block execution are stopped here
	// and engine.Close cannot be called directly because it has a dependency on the following modules
	s.engine.Stop()
	s.bloomIndexer.Close()
	close(s.closeBloomHandler)
	s.txPool.Stop()
	s.miner.Stop()
	s.blockchain.Stop()
	s.engine.Close()
	core.GetReactorInstance().Close()
	s.chainDb.Close()
	s.eventMux.Stop()
	return nil
}

// RegisterPlugin one by one
func handlePlugin(reactor *core.BlockChainReactor, chainDB ethdb.Database, isValidatorsHistory bool, opConfig *params.OpConfig, chainId *big.Int) {
	xplugin.RewardMgrInstance().SetCurrentNodeID(reactor.NodeId)

	reactor.RegisterPlugin(xcom.SlashingRule, xplugin.SlashInstance())
	xplugin.SlashInstance().SetDecodeEvidenceFun(evidence.NewEvidence)
	reactor.RegisterPlugin(xcom.StakingRule, xplugin.StakingInstance())
	reactor.RegisterPlugin(xcom.RestrictingRule, xplugin.RestrictingInstance())
	reactor.RegisterPlugin(xcom.RewardRule, xplugin.RewardMgrInstance())
	xplugin.TokenInstance().SetOpConfig(opConfig)
	xplugin.TokenInstance().SetChainID(chainId)
	if reactor.NodeId == opConfig.SubChain.NodeId {
		// Set the sub-chain operation node identity
		xplugin.TokenInstance().SetSubOpIdentity(true)
	}
	reactor.RegisterPlugin(xcom.TokenRule, xplugin.TokenInstance())

	xplugin.GovPluginInstance().SetChainID(reactor.GetChainID())
	xplugin.GovPluginInstance().SetChainDB(chainDB)
	reactor.RegisterPlugin(xcom.GovernanceRule, xplugin.GovPluginInstance())

	xplugin.StakingInstance().SetChainDB(chainDB, chainDB)
	if isValidatorsHistory {
		xplugin.StakingInstance().EnableValidatorsHistory()
	}

	// set rule order
	reactor.SetBeginRule([]int{xcom.StakingRule, xcom.SlashingRule, xcom.CollectDeclareVersionRule, xcom.GovernanceRule, xcom.TokenRule})
	reactor.SetEndRule([]int{xcom.CollectDeclareVersionRule, xcom.RestrictingRule, xcom.RewardRule, xcom.GovernanceRule, xcom.StakingRule, xcom.TokenRule})

}
