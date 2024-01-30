// Copyright 2021 The Bubble Network Authors
// This file is part of the bubble library.
//
// The bubble library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The bubble library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the bubble library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/bubblenet/bubble/x/bubble"

	"github.com/bubblenet/bubble/x/token"

	"github.com/bubblenet/bubble/common"
	cvm "github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core/cbfttypes"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/core/state"
	"github.com/bubblenet/bubble/core/vm"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/x/handler"
	"github.com/bubblenet/bubble/x/staking"
	"github.com/bubblenet/bubble/x/xutil"

	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/event"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/x/plugin"
	"github.com/bubblenet/bubble/x/xcom"
)

type BlockChainReactor struct {
	vh            *handler.VrfHandler
	eventMux      *event.TypeMux
	bftResultSub  *event.TypeMuxSubscription
	settleTaskSub *event.TypeMuxSubscription // Settlement task subscription
	remoteCallSub *event.TypeMuxSubscription
	basePluginMap map[int]plugin.BasePlugin // xxPlugin container
	beginRule     []int                     // Order rules for xxPlugins called in BeginBlocker
	endRule       []int                     // Order rules for xxPlugins called in EndBlocker
	validatorMode string                    // mode: static, inner, dpos
	NodeId        discover.NodeID           // The nodeId of current node
	exitCh        chan chan struct{}        // Used to receive an exit signal
	exitOnce      sync.Once
	chainID       *big.Int
}

var (
	bcrOnce sync.Once
	bcr     *BlockChainReactor
)

func NewBlockChainReactor(mux *event.TypeMux, chainId *big.Int) *BlockChainReactor {
	bcrOnce.Do(func() {
		log.Info("Init BlockChainReactor ...")
		bcr = &BlockChainReactor{
			eventMux:      mux,
			basePluginMap: make(map[int]plugin.BasePlugin, 0),
			exitCh:        make(chan chan struct{}),
			chainID:       chainId,
		}
	})
	return bcr
}

func (bcr *BlockChainReactor) Start(mode string) {
	bcr.setValidatorMode(mode)
	if mode == common.DPOS_VALIDATOR_MODE {
		// Subscribe events for confirmed blocks
		bcr.bftResultSub = bcr.eventMux.Subscribe(cbfttypes.CbftResult{})
		bcr.settleTaskSub = bcr.eventMux.Subscribe(token.SettleTask{})
		bcr.remoteCallSub = bcr.eventMux.Subscribe(bubble.RemoteCallTask{})
		// start the loop rutine
		go bcr.loop()
		go bcr.handleTask()
	}
}

func (bcr *BlockChainReactor) Close() {
	if bcr.validatorMode == common.DPOS_VALIDATOR_MODE {
		bcr.exitOnce.Do(func() {
			exitDone := make(chan struct{})
			bcr.exitCh <- exitDone
			<-exitDone
			close(exitDone)
		})
	}
	log.Info("blockchain_reactor closed")
}

func (bcr *BlockChainReactor) GetChainID() *big.Int {
	return bcr.chainID
}

// Getting the global bcr single instance
func GetReactorInstance() *BlockChainReactor {
	return bcr
}

func (bcr *BlockChainReactor) loop() {

	for {
		select {
		case obj := <-bcr.bftResultSub.Chan():
			if obj == nil {
				//log.Error("blockchain_reactor receive nil bftResultEvent maybe channel is closed")
				continue
			}
			cbftResult, ok := obj.Data.(cbfttypes.CbftResult)
			if !ok {
				log.Error("blockchain_reactor receive bft result type error")
				continue
			}
			bcr.commit(cbftResult.Block)
		// stop this routine
		case done := <-bcr.exitCh:
			close(bcr.exitCh)
			log.Info("blockChain reactor loop exit")
			done <- struct{}{}
			return
		}
	}
}

func (bcr *BlockChainReactor) handleTask() {

	for {
		select {
		case settleInfo := <-bcr.settleTaskSub.Chan():
			if settleInfo == nil {
				continue
			}
			settleData, ok := settleInfo.Data.(token.SettleTask)
			if !ok {
				log.Error("blockchain_reactor failed to receive settlement data conversion type")
				continue
			}
			// handle task
			hash, err := plugin.TokenInstance().HandleSettleTask(&settleData)
			if err != nil {
				log.Error("blockchain_reactor failed to process settlement task")
				continue
			}
			log.Info("The processing and settlement task succeeded, tx hash:", common.BytesToHash(hash).Hex())
		case msg := <-bcr.remoteCallSub.Chan():
			if msg == nil {
				continue
			}
			task, ok := msg.Data.(bubble.RemoteCallTask)
			if !ok {
				log.Error("blockchain_reactor failed to receive remoteCall task", "msg", msg.Data)
				continue
			}
			// handle task
			hash, err := plugin.BubbleInstance().HandleRemoteCallTask(&task)
			if err != nil {
				log.Error("blockchain_reactor failed to process RemoteCall task", "err", err.Error())
				continue
			}
			log.Info("The processing and RemoteCall task succeeded", "bubbleID", task.BubbleID, "Contract", task.Contract,
				"txHash", task.TxHash, "remoteTxHash", common.BytesToHash(hash).Hex())
		}
	}
}

func (bcr *BlockChainReactor) commit(block *types.Block) error {
	if block == nil {
		log.Error("blockchain_reactor receive Cbft result error, block is nil")
		return nil
	}
	/**
	notify P2P module the nodeId of the next round validator
	*/
	if plugin, ok := bcr.basePluginMap[xcom.StakingRule]; ok {
		if err := plugin.Confirmed(bcr.NodeId, block); nil != err {
			log.Error("Failed to call Staking Confirmed", "blockNumber", block.Number(), "blockHash", block.Hash().Hex(), "err", err.Error())
		}

	}

	log.Info("Call snapshotdb commit on blockchain_reactor", "blockNumber", block.Number(), "blockHash", block.Hash())
	if err := snapshotdb.Instance().Commit(block.Hash()); nil != err {
		log.Error("Failed to call snapshotdb commit on blockchain_reactor", "blockNumber", block.Number(), "blockHash", block.Hash(), "err", err)
		return err
	}
	return nil
}

func (bcr *BlockChainReactor) OnCommit(block *types.Block) error {
	if bcr.validatorMode == common.DPOS_VALIDATOR_MODE {
		return bcr.commit(block)
	}
	return nil
}

func (bcr *BlockChainReactor) RegisterPlugin(pluginRule int, plugin plugin.BasePlugin) {
	bcr.basePluginMap[pluginRule] = plugin
}

func (bcr *BlockChainReactor) SetPluginEventMux() {
	plugin.StakingInstance().SetEventMux(bcr.eventMux)
	plugin.TokenInstance().SetEventMux(bcr.eventMux)
	plugin.BubbleInstance().SetEventMux(bcr.eventMux)
}

func (bcr *BlockChainReactor) setValidatorMode(mode string) {
	bcr.validatorMode = mode
}

func (bcr *BlockChainReactor) SetVRFhandler(vher *handler.VrfHandler) {
	bcr.vh = vher
}

func (bcr *BlockChainReactor) SetPrivateKey(privateKey *ecdsa.PrivateKey) {
	if bcr.validatorMode == common.DPOS_VALIDATOR_MODE && nil != privateKey {
		if nil != bcr.vh {
			bcr.vh.SetPrivateKey(privateKey)
		}
		plugin.SlashInstance().SetPrivateKey(privateKey)
		bcr.NodeId = discover.PubkeyID(&privateKey.PublicKey)
	}
}

func (bcr *BlockChainReactor) SetBeginRule(rule []int) {
	bcr.beginRule = rule
}
func (bcr *BlockChainReactor) SetEndRule(rule []int) {
	bcr.endRule = rule
}

func (bcr *BlockChainReactor) SetWorkerCoinBase(header *types.Header, nodeId discover.NodeID) {

	/**
	this things about dpos
	*/
	if bcr.validatorMode != common.DPOS_VALIDATOR_MODE {
		return
	}

	nodeIdAddr, err := xutil.NodeId2Addr(nodeId)
	if nil != err {
		log.Error("Failed to SetWorkerCoinBase: parse current nodeId is failed", "err", err)
		panic(fmt.Sprintf("parse current nodeId is failed: %s", err.Error()))
	}

	if plu, ok := bcr.basePluginMap[xcom.StakingRule]; ok {
		stake := plu.(*plugin.StakingPlugin)
		can, err := stake.GetCandidateInfo(common.ZeroHash, nodeIdAddr)
		if nil != err {
			log.Error("Failed to SetWorkerCoinBase: Query candidate info is failed", "blockNumber", header.Number,
				"nodeId", nodeId.String(), "nodeIdAddr", nodeIdAddr.Hex(), "err", err)
			return
		}
		header.Coinbase = can.BenefitAddress
		log.Info("SetWorkerCoinBase Successfully", "blockNumber", header.Number,
			"nodeId", nodeId.String(), "nodeIdAddr", nodeIdAddr.Hex(), "coinbase", header.Coinbase.String())
	}

}

// Called before every block has not executed all txs
func (bcr *BlockChainReactor) BeginBlocker(header *types.Header, state xcom.StateDB) error {

	/**
	this things about dpos
	*/
	if bcr.validatorMode != common.DPOS_VALIDATOR_MODE {
		return nil
	}

	blockHash := common.ZeroHash

	// store the sign in  header.Extra[32:97]
	if xutil.IsWorker(header.Extra) {
		// Generate vrf proof
		if value, err := bcr.vh.GenerateNonce(header.Number, header.ParentHash); nil != err {
			return err
		} else {
			header.Nonce = types.EncodeNonce(value)
		}
	} else {
		blockHash = header.CacheHash()
		// Verify vrf proof
		pk := header.CachePublicKey()
		if pk == nil {
			return errors.New("failed to get the public key of the block producer")
		}
		if err := bcr.vh.VerifyVrf(pk, header.Number, header.ParentHash, blockHash, header.Nonce.Bytes()); nil != err {
			return err
		}
	}

	log.Debug("Call snapshotDB newBlock on blockchain_reactor", "blockNumber", header.Number.Uint64(),
		"hash", blockHash, "parentHash", header.ParentHash)
	if err := snapshotdb.Instance().NewBlock(header.Number, header.ParentHash, blockHash); nil != err {
		log.Error("Failed to call snapshotDB newBlock on blockchain_reactor", "blockNumber",
			header.Number.Uint64(), "hash", hex.EncodeToString(blockHash.Bytes()), "parentHash",
			hex.EncodeToString(header.ParentHash.Bytes()), "err", err)
		return err
	}

	for _, pluginRule := range bcr.beginRule {
		if plugin, ok := bcr.basePluginMap[pluginRule]; ok {
			if err := plugin.BeginBlock(blockHash, header, state); nil != err {
				return err
			}
		}
	}

	// This must not be deleted
	root := state.IntermediateRoot(true)
	log.Debug("BeginBlock StateDB root, end", "blockHash", header.Hash(), "blockNumber",
		header.Number.Uint64(), "root", root, "pointer", fmt.Sprintf("%p", state))

	return nil
}

// Called after every block had executed all txs
func (bcr *BlockChainReactor) EndBlocker(header *types.Header, state xcom.StateDB) error {

	/**
	this things about dpos
	*/
	if bcr.validatorMode != common.DPOS_VALIDATOR_MODE {
		return nil
	}

	blockHash := common.ZeroHash

	if !xutil.IsWorker(header.Extra) {
		blockHash = header.CacheHash()
	}

	// Store the previous vrf random number
	if err := bcr.vh.Storage(header.Number, header.ParentHash, blockHash, header.Nonce.Bytes()); nil != err {
		log.Error("blockchain_reactor Storage proof failed", "blockNumber", header.Number.Uint64(),
			"blockHash", hex.EncodeToString(blockHash.Bytes()), "err", err)
		return err
	}

	for _, pluginRule := range bcr.endRule {
		if plugin, ok := bcr.basePluginMap[pluginRule]; ok {
			if err := plugin.EndBlock(blockHash, header, state); nil != err {
				return err
			}
		}
	}

	// storage the dpos k-v Hash
	dposHash := snapshotdb.Instance().GetLastKVHash(blockHash)

	if len(dposHash) != 0 && !bytes.Equal(dposHash, make([]byte, len(dposHash))) {
		// store hash about dpos
		state.SetState(cvm.StakingContractAddr, staking.GetDPOSHASHKey(), dposHash)
		log.Debug("Store dpos hash", "blockHash", blockHash, "blockNumber", header.Number.Uint64(),
			"dposHash", hex.EncodeToString(dposHash))
	}

	// This must not be deleted
	root := state.IntermediateRoot(true)
	log.Debug("EndBlock StateDB root, end", "blockHash", blockHash, "blockNumber",
		header.Number.Uint64(), "root", root, "pointer", fmt.Sprintf("%p", state))

	return nil
}

func (bcr *BlockChainReactor) VerifyTx(tx *types.Transaction, to common.Address) error {

	if !vm.IsBubblePrecompiledContract(to) {
		return nil
	}

	input := tx.Data()
	if len(input) == 0 {
		return nil
	}

	var contract vm.BubblePrecompiledContract
	switch to {
	case cvm.TokenContractAddr:
		c := vm.BubblePrecompiledContracts[cvm.TokenContractAddr]
		contract = c.(vm.BubblePrecompiledContract)
	case cvm.BubbleContractAddr:
		c := vm.BubblePrecompiledContracts[cvm.BubbleContractAddr]
		contract = c.(vm.BubblePrecompiledContract)
	case cvm.TempPrivateKeyContractAddr:
		c := vm.BubblePrecompiledContracts[cvm.TempPrivateKeyContractAddr]
		contract = c.(vm.BubblePrecompiledContract)
	case cvm.StakingContractAddr:
		c := vm.BubblePrecompiledContracts[cvm.StakingContractAddr]
		contract = c.(vm.BubblePrecompiledContract)
	case cvm.RestrictingContractAddr:
		c := vm.BubblePrecompiledContracts[cvm.RestrictingContractAddr]
		contract = c.(vm.BubblePrecompiledContract)
	case cvm.GovContractAddr:
		c := vm.BubblePrecompiledContracts[cvm.GovContractAddr]
		contract = c.(vm.BubblePrecompiledContract)
	case cvm.SlashingContractAddr:
		c := vm.BubblePrecompiledContracts[cvm.SlashingContractAddr]
		contract = c.(vm.BubblePrecompiledContract)
	default:
		// pass if the contract is validatorInnerContract
		return nil
	}
	// verify the dpos contract tx.data
	if contract != nil {
		if fcode, _, _, err := plugin.VerifyTxData(input, contract.FnSigns()); nil != err {
			return err
		} else {
			return contract.CheckGasPrice(tx.GasPrice(), fcode)
		}
	} else {
		log.Warn("Cannot find an appropriate BubblePrecompiledContract!")
		return nil
	}
}

func (bcr *BlockChainReactor) Sign(msg interface{}) error {
	return nil
}

func (bcr *BlockChainReactor) VerifySign(msg interface{}) error {
	return nil
}

func (bcr *BlockChainReactor) VerifyHeader(header *types.Header, stateDB *state.StateDB) error {
	return nil
}

func (bcr *BlockChainReactor) GetLastNumber(blockNumber uint64) uint64 {
	return plugin.StakingInstance().GetLastNumber(blockNumber)
}

func (bcr *BlockChainReactor) GetValidator(blockNumber uint64) (*cbfttypes.Validators, error) {
	return plugin.StakingInstance().GetValidator(blockNumber)
}

func (bcr *BlockChainReactor) IsCandidateNode(nodeID discover.NodeID) bool {
	return plugin.StakingInstance().IsCandidateNode(nodeID)
}

func (bcr *BlockChainReactor) Flush(header *types.Header) error {
	log.Debug("Call snapshotdb flush on blockchain_reactor", "blockNumber", header.Number.Uint64(), "hash", header.Hash())
	if err := snapshotdb.Instance().Flush(header.Hash(), header.Number); nil != err {
		log.Error("Failed to call snapshotdb flush on blockchain_reactor", "blockNumber", header.Number.Uint64(), "hash", header.Hash(), "err", err)
		return err
	}
	return nil
}
