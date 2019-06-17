package core

import (
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/x/common"
	"github.com/PlatONnetwork/PlatON-Go/x/core/staking"
)

type BlockChainReactor struct {


	eventMux      	*event.TypeMux
	bftResultSub 	*event.TypeMuxSubscription


	// xxPlugin container
	basePluginMap  	map[int]BasePlugin
	// Order rules for xxPlugins called in BeginBlocker
	beginRule		[]int
	// Order rules for xxPlugins called in EndBlocker
	endRule 		[]int
}


var bcr *BlockChainReactor


func New (mux *event.TypeMux) *BlockChainReactor {
	if nil == bcr {
		bcr = &BlockChainReactor{
			eventMux: 		mux,
			basePluginMap: 	make(map[int]BasePlugin, 0),
			//beginRule:		make([]string, 0),
			//endRule: 		make([]string, 0),
		}
		// Subscribe events for confirmed blocks
		bcr.bftResultSub = bcr.eventMux.Subscribe(cbfttypes.CbftResult{})

		// start the loop rutine
		go bcr.loop()
	}
	return bcr
}

// Getting the global bcr single instance
func GetInstance () *BlockChainReactor {
	return bcr
}


func (brc *BlockChainReactor) loop () {

	for {
		select {
		case obj := <-bcr.bftResultSub.Chan():
			if obj == nil {
				log.Error("BlockChainReactor receive nil bftResultEvent maybe channel is closed")
				continue
			}
			cbftResult, ok := obj.Data.(cbfttypes.CbftResult)
			if !ok {
				log.Error("receive bft result type error")
				continue
			}
			block := cbftResult.Block
			// Short circuit when receiving empty result.
			if block == nil {
				log.Error("Cbft result error, block is nil")
				continue
			}

			/**
			TODO flush the seed and the package ratio
			 */

			if plugin, ok := brc.basePluginMap[common.StakingRule]; ok {
				if staking, ok := plugin.(*core.StakingPlugin); ok {
					if err := staking.Confirmed(block); nil != err {
						log.Error("Failed to call Staking Confirmed", "blockNumber", block.Number(), "blockHash", block.Hash().Hex(), "err", err.Error())
					}
				}

			}

			/*// TODO Slashing
			if plugin, ok := brc.basePluginMap[common.StakingRule]; ok {
				if slashing, ok := plugin.(*core.StakingPlugin); ok {
					if err := slashing.Confirmed(block); nil != err {
						log.Error("Failed to call Staking Confirmed", "blockNumber", block.Number(), "blockHash", block.Hash().Hex(), "err", err.Error())
					}
				}

			}*/


		default:
				return

		}
	}

}


func (bcr *BlockChainReactor) RegisterPlugin (pluginRule int, plugin BasePlugin) {
	bcr.basePluginMap[pluginRule] = plugin
}
func (bcr *BlockChainReactor) SetBeginRule(rule []int) {
	bcr.beginRule = rule
}
func (bcr *BlockChainReactor) SetEndRule(rule []int) {
	bcr.endRule = rule
}


// Called before every block has not executed all txs
func (bcr *BlockChainReactor) BeginBlocker (header *types.Header, state *state.StateDB) (bool, error) {

	for _, pluginName := range bcr.beginRule {
		if plugin, ok := bcr.basePluginMap[pluginName]; ok {
			if flag, err := plugin.BeginBlock(header, state); nil != err {
				return flag, err
			}
		}
	}
	return false, nil
}

// Called after every block had executed all txs
func (bcr *BlockChainReactor) EndBlocker (header *types.Header, state *state.StateDB) (bool, error) {

	for _, pluginName := range bcr.endRule {
		if plugin, ok := bcr.basePluginMap[pluginName]; ok {
			if flag, err := plugin.EndBlock(header, state); nil != err {
				return flag, err
			}
		}
	}
	return false, nil
}


