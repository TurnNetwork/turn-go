// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package eth

import (
	"math/big"
	"time"

	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/eth/downloader"
	"github.com/PlatONnetwork/PlatON-Go/eth/gasprice"
	"github.com/PlatONnetwork/PlatON-Go/miner"
)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                  *core.Genesis       `toml:",omitempty"`
		CbftConfig               types.OptionsConfig `toml:",omitempty"`
		NetworkId                uint64
		SyncMode                 downloader.SyncMode
		NoPruning                bool
		LightServ                int  `toml:",omitempty"`
		LightPeers               int  `toml:",omitempty"`
		SkipBcVersionCheck       bool `toml:"-"`
		DatabaseHandles          int  `toml:"-"`
		DatabaseCache            int
		DatabaseFreezer          string
		TxLookupLimit            uint64 `toml:",omitempty"`
		TrieCache                int
		TrieTimeout              time.Duration
		TrieDBCache              int
		DBDisabledGC             bool
		DBGCInterval             uint64
		DBGCTimeout              time.Duration
		DBGCMpt                  bool
		DBGCBlock                int
		VMWasmType               string
		VmTimeoutDuration        uint64
		Miner                    miner.Config
		MiningLogAtDepth         uint
		TxChanSize               int
		ChainHeadChanSize        int
		ChainSideChanSize        int
		ResultQueueSize          int
		ResubmitAdjustChanSize   int
		MinRecommitInterval      time.Duration
		MaxRecommitInterval      time.Duration
		IntervalAdjustRatio      float64
		IntervalAdjustBias       float64
		StaleThreshold           uint64
		DefaultCommitRatio       float64
		BodyCacheLimit           int
		BlockCacheLimit          int
		MaxFutureBlocks          int
		BadBlockLimit            int
		TriesInMemory            int
		BlockChainVersion        int
		DefaultTxsCacheSize      int
		DefaultBroadcastInterval time.Duration
		TxPool                   core.TxPoolConfig
		GPO                      gasprice.Config
		DocRoot                  string `toml:"-"`
		Debug                    bool
		RPCGasCap                *big.Int `toml:",omitempty"`
		RPCTxFeeCap              float64  `toml:",omitempty"`
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.CbftConfig = c.CbftConfig
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.NoPruning = c.NoPruning
	enc.LightServ = c.LightServ
	enc.LightPeers = c.LightPeers
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.DatabaseFreezer = c.DatabaseFreezer
	enc.TxLookupLimit = c.TxLookupLimit
	enc.TrieCache = c.TrieCache
	enc.TrieTimeout = c.TrieTimeout
	enc.TrieDBCache = c.TrieDBCache
	enc.DBDisabledGC = c.DBDisabledGC
	enc.DBGCInterval = c.DBGCInterval
	enc.DBGCTimeout = c.DBGCTimeout
	enc.DBGCMpt = c.DBGCMpt
	enc.DBGCBlock = c.DBGCBlock
	enc.VMWasmType = c.VMWasmType
	enc.VmTimeoutDuration = c.VmTimeoutDuration
	enc.Miner = c.Miner
	enc.MiningLogAtDepth = c.MiningLogAtDepth
	enc.TxChanSize = c.TxChanSize
	enc.ChainHeadChanSize = c.ChainHeadChanSize
	enc.ChainSideChanSize = c.ChainSideChanSize
	enc.ResultQueueSize = c.ResultQueueSize
	enc.ResubmitAdjustChanSize = c.ResubmitAdjustChanSize
	enc.MinRecommitInterval = c.MinRecommitInterval
	enc.MaxRecommitInterval = c.MaxRecommitInterval
	enc.IntervalAdjustRatio = c.IntervalAdjustRatio
	enc.IntervalAdjustBias = c.IntervalAdjustBias
	enc.StaleThreshold = c.StaleThreshold
	enc.DefaultCommitRatio = c.DefaultCommitRatio
	enc.BodyCacheLimit = c.BodyCacheLimit
	enc.BlockCacheLimit = c.BlockCacheLimit
	enc.MaxFutureBlocks = c.MaxFutureBlocks
	enc.BadBlockLimit = c.BadBlockLimit
	enc.TriesInMemory = c.TriesInMemory
	enc.BlockChainVersion = c.BlockChainVersion
	enc.DefaultTxsCacheSize = c.DefaultTxsCacheSize
	enc.DefaultBroadcastInterval = c.DefaultBroadcastInterval
	enc.TxPool = c.TxPool
	enc.GPO = c.GPO
	enc.DocRoot = c.DocRoot
	enc.Debug = c.Debug
	enc.RPCGasCap = c.RPCGasCap
	enc.RPCTxFeeCap = c.RPCTxFeeCap
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                  *core.Genesis        `toml:",omitempty"`
		CbftConfig               *types.OptionsConfig `toml:",omitempty"`
		NetworkId                *uint64
		SyncMode                 *downloader.SyncMode
		NoPruning                *bool
		LightServ                *int  `toml:",omitempty"`
		LightPeers               *int  `toml:",omitempty"`
		SkipBcVersionCheck       *bool `toml:"-"`
		DatabaseHandles          *int  `toml:"-"`
		DatabaseCache            *int
		DatabaseFreezer          *string
		TxLookupLimit            *uint64 `toml:",omitempty"`
		TrieCache                *int
		TrieTimeout              *time.Duration
		TrieDBCache              *int
		DBDisabledGC             *bool
		DBGCInterval             *uint64
		DBGCTimeout              *time.Duration
		DBGCMpt                  *bool
		DBGCBlock                *int
		VMWasmType               *string
		VmTimeoutDuration        *uint64
		Miner                    *miner.Config
		MiningLogAtDepth         *uint
		TxChanSize               *int
		ChainHeadChanSize        *int
		ChainSideChanSize        *int
		ResultQueueSize          *int
		ResubmitAdjustChanSize   *int
		MinRecommitInterval      *time.Duration
		MaxRecommitInterval      *time.Duration
		IntervalAdjustRatio      *float64
		IntervalAdjustBias       *float64
		StaleThreshold           *uint64
		DefaultCommitRatio       *float64
		BodyCacheLimit           *int
		BlockCacheLimit          *int
		MaxFutureBlocks          *int
		BadBlockLimit            *int
		TriesInMemory            *int
		BlockChainVersion        *int
		DefaultTxsCacheSize      *int
		DefaultBroadcastInterval *time.Duration
		TxPool                   *core.TxPoolConfig
		GPO                      *gasprice.Config
		DocRoot                  *string `toml:"-"`
		Debug                    *bool
		RPCGasCap                *big.Int `toml:",omitempty"`
		RPCTxFeeCap              *float64 `toml:",omitempty"`
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.CbftConfig != nil {
		c.CbftConfig = *dec.CbftConfig
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.NoPruning != nil {
		c.NoPruning = *dec.NoPruning
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.DatabaseFreezer != nil {
		c.DatabaseFreezer = *dec.DatabaseFreezer
	}
	if dec.TxLookupLimit != nil {
		c.TxLookupLimit = *dec.TxLookupLimit
	}
	if dec.TrieCache != nil {
		c.TrieCache = *dec.TrieCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.TrieDBCache != nil {
		c.TrieDBCache = *dec.TrieDBCache
	}
	if dec.DBDisabledGC != nil {
		c.DBDisabledGC = *dec.DBDisabledGC
	}
	if dec.DBGCInterval != nil {
		c.DBGCInterval = *dec.DBGCInterval
	}
	if dec.DBGCTimeout != nil {
		c.DBGCTimeout = *dec.DBGCTimeout
	}
	if dec.DBGCMpt != nil {
		c.DBGCMpt = *dec.DBGCMpt
	}
	if dec.DBGCBlock != nil {
		c.DBGCBlock = *dec.DBGCBlock
	}
	if dec.VMWasmType != nil {
		c.VMWasmType = *dec.VMWasmType
	}
	if dec.VmTimeoutDuration != nil {
		c.VmTimeoutDuration = *dec.VmTimeoutDuration
	}
	if dec.Miner != nil {
		c.Miner = *dec.Miner
	}
	if dec.MiningLogAtDepth != nil {
		c.MiningLogAtDepth = *dec.MiningLogAtDepth
	}
	if dec.TxChanSize != nil {
		c.TxChanSize = *dec.TxChanSize
	}
	if dec.ChainHeadChanSize != nil {
		c.ChainHeadChanSize = *dec.ChainHeadChanSize
	}
	if dec.ChainSideChanSize != nil {
		c.ChainSideChanSize = *dec.ChainSideChanSize
	}
	if dec.ResultQueueSize != nil {
		c.ResultQueueSize = *dec.ResultQueueSize
	}
	if dec.ResubmitAdjustChanSize != nil {
		c.ResubmitAdjustChanSize = *dec.ResubmitAdjustChanSize
	}
	if dec.MinRecommitInterval != nil {
		c.MinRecommitInterval = *dec.MinRecommitInterval
	}
	if dec.MaxRecommitInterval != nil {
		c.MaxRecommitInterval = *dec.MaxRecommitInterval
	}
	if dec.IntervalAdjustRatio != nil {
		c.IntervalAdjustRatio = *dec.IntervalAdjustRatio
	}
	if dec.IntervalAdjustBias != nil {
		c.IntervalAdjustBias = *dec.IntervalAdjustBias
	}
	if dec.StaleThreshold != nil {
		c.StaleThreshold = *dec.StaleThreshold
	}
	if dec.DefaultCommitRatio != nil {
		c.DefaultCommitRatio = *dec.DefaultCommitRatio
	}
	if dec.BodyCacheLimit != nil {
		c.BodyCacheLimit = *dec.BodyCacheLimit
	}
	if dec.BlockCacheLimit != nil {
		c.BlockCacheLimit = *dec.BlockCacheLimit
	}
	if dec.MaxFutureBlocks != nil {
		c.MaxFutureBlocks = *dec.MaxFutureBlocks
	}
	if dec.BadBlockLimit != nil {
		c.BadBlockLimit = *dec.BadBlockLimit
	}
	if dec.TriesInMemory != nil {
		c.TriesInMemory = *dec.TriesInMemory
	}
	if dec.BlockChainVersion != nil {
		c.BlockChainVersion = *dec.BlockChainVersion
	}
	if dec.DefaultTxsCacheSize != nil {
		c.DefaultTxsCacheSize = *dec.DefaultTxsCacheSize
	}
	if dec.DefaultBroadcastInterval != nil {
		c.DefaultBroadcastInterval = *dec.DefaultBroadcastInterval
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.Debug != nil {
		c.Debug = *dec.Debug
	}
	if dec.RPCGasCap != nil {
		c.RPCGasCap = dec.RPCGasCap
	}
	if dec.RPCTxFeeCap != nil {
		c.RPCTxFeeCap = *dec.RPCTxFeeCap
	}
	return nil
}
