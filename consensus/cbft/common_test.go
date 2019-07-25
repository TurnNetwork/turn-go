package cbft

import (
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/validator"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethdb"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/params"
)

var (
	chainConfig      = params.TestnetChainConfig
	testTxPoolConfig = core.DefaultTxPoolConfig
)

func CreateCBFT() *Cbft {
	priKey, _ := crypto.GenerateKey()

	sysConfig := &params.CbftConfig{
		Epoch:        1,
		Period:       10,
		Amount:       10,
		InitialNodes: []discover.Node{},
	}

	optConfig := &types.OptionsConfig{
		NodePriKey: priKey,
		NodeID:     discover.PubkeyID(&priKey.PublicKey),
	}

	ctx := node.NewServiceContext(&node.Config{DataDir: ""}, nil, new(event.TypeMux), nil)

	return New(sysConfig, optConfig, ctx.EventMux, ctx)
}

func CreateBackend(engine *Cbft, nodes []discover.Node) {
	var (
		db    = ethdb.NewMemDatabase()
		gspec = core.Genesis{
			Config: chainConfig,
			Alloc:  core.GenesisAlloc{},
		}
	)
	gspec.MustCommit(db)

	chain, _ := core.NewBlockChain(db, nil, gspec.Config, engine, vm.Config{}, nil)
	cache := core.NewBlockChainCache(chain)
	txpool := core.NewTxPool(testTxPoolConfig, chainConfig, cache)

	engine.Start(chain, cache, txpool, validator.NewStaticAgency(nodes))
}

//func TestMockNode(t *testing.T) {
//	cbft := CreateCBFT()
//	backend := CreateBackend(engine, validators.Nodes())
//}
