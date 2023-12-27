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

package vm

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bubblenet/bubble/accounts/abi"
	"github.com/bubblenet/bubble/x/token"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/bubble"
	"github.com/bubblenet/bubble/x/plugin"
)

const (
	TxRemoteDeployExecutor  = 8000
	TxRemoteRemoveExecutor  = 8002
	TxRemoteDestroyExecutor = 8001
	TxRemoteCall            = 8003
	TxRemoteCallExecutor    = 8004
)

type BubbleContract struct {
	Plugin      *plugin.BubblePlugin
	tokenPlugin *plugin.TokenPlugin
	Contract    *Contract
	Evm         *EVM
}

func (bc *BubbleContract) RequiredGas(input []byte) uint64 {
	if checkInputEmpty(input) {
		return 0
	}
	return params.BubbleGas
}

func (bc *BubbleContract) Run(input []byte) ([]byte, error) {
	if checkInputEmpty(input) {
		return nil, nil
	}
	return execBubbleContract(input, bc.FnSigns())
}

func (bc *BubbleContract) FnSigns() map[uint16]interface{} {
	return map[uint16]interface{}{
		// Set
		TxRemoteDeployExecutor:  bc.remoteDeployExecutor,
		TxRemoteDestroyExecutor: bc.remoteDestroyExecutor,
		TxRemoteRemoveExecutor:  bc.remoteRemoveExecutor,
		TxRemoteCall:            bc.remoteCall,
		TxRemoteCallExecutor:    bc.remoteCallExecutor,
	}
}

func (bc *BubbleContract) CheckGasPrice(gasPrice *big.Int, fcode uint16) error {
	return nil
}

// remoteDeployExecutor receive the remoteDeploy transaction from main chain and deploy the contract to the bubble chain
func (bc *BubbleContract) remoteDeployExecutor(remoteTxHash common.Hash, sender *common.Address, address *common.Address, bytecode []byte, data []byte) ([]byte, error) {
	from := bc.Contract.CallerAddress
	blockHash := bc.Evm.Context.BlockHash
	blockNumber := bc.Evm.Context.BlockNumber
	txHash := bc.Evm.StateDB.TxHash()
	//gas := bc.Contract.Gas

	log.Debug("Call remoteDeployExecutor of bubbleContract", "chainID", bc.tokenPlugin.ChainID, "txHash", txHash.Hex(), "blockNumber", blockNumber.Uint64(),
		"from", from, "address", address)

	if !bc.Contract.UseGas(params.RemoteDeployExecutorGas) {
		return nil, ErrOutOfGas
	}

	operator := bc.tokenPlugin.OpConfig.MainChain.OpAddr
	if from != operator {
		log.Error("the sender is not the operator of the layer1", "from", from, "operator", operator)
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteDeployExecutor", "the sender is not the operator of the layer1", TxRemoteRemoveExecutor,
			bubble.ErrSenderIsNotOperator)
	}

	if byteCode := bc.Evm.StateDB.GetCode(*address); len(byteCode) != 0 {
		log.Error("the contract is existed", "address", address)
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteDeployExecutor", "contract is existed", TxRemoteDeployExecutor, bubble.ErrContractIsExist)
	}

	isEstimateGas := txHash == common.ZeroHash

	// force set code to address
	bc.Evm.StateDB.SetCode(*address, bytecode)

	// call initialize function
	bc.Contract.caller = AccountRef(*sender)
	bc.Contract.CallerAddress = *sender
	bc.Contract.self = AccountRef(*address)
	bc.Contract.SetCallCode(address, bc.Evm.StateDB.GetCodeHash(*address), bytecode)
	if len(data) != 0 {
		_, err := RunEvm(bc.Evm, bc.Contract, data)
		if err != nil {
			errMsg := fmt.Sprintf("failed to call data when remoteDeployExecutor, error:%v", err.Error())
			log.Error(errMsg)
			if isEstimateGas {
				// The error returned by the action of deducting gas during the estimated gas process cannot be BizError,
				// otherwise the estimated process will be interrupted
				return nil, errors.New(errMsg)
			} else {
				return nil, bubble.ErrContractReturns
			}
		}
	}

	// discard!!!
	// switch to the evm runtime
	//caller := AccountRef(vm.BubbleContractAddr)
	//ch := &codeAndHash{code: append(bytecode, data...)}
	//_, _, _, err := bc.Evm.create(caller, ch, gas, big.NewInt(0), address)
	//if err != nil {
	//	log.Error("remote deploy contract returned an error", "chainID", bc.tokenPlugin.ChainID, "contract", address, "error", err)
	//	return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteDeployExecutor", "remote deploy contract returned an error", TxRemoteDeployExecutor,
	//		bubble.ErrContractReturns.Wrap(err.Error()))
	//}

	if err := bc.Plugin.StoreBubContract(blockHash, address); err != nil {
		return nil, err
	}

	if err := bc.tokenPlugin.StoreL1HashToL2Hash(blockHash, remoteTxHash, txHash); err != nil {
		return nil, token.ErrStoreL1HashToL2Hash
	}

	return txResultHandler(vm.BubbleContractAddr, bc.Evm, "", "", TxRemoteDeployExecutor, common.NoErr)
}

// remoteRemoveExecutor receive the remoteRemove transaction from main chain and remove the contract from the bubble chain
func (bc *BubbleContract) remoteRemoveExecutor(remoteTxHash common.Hash, address *common.Address) ([]byte, error) {
	from := bc.Contract.CallerAddress
	blockHash := bc.Evm.Context.BlockHash
	blockNumber := bc.Evm.Context.BlockNumber
	txHash := bc.Evm.StateDB.TxHash()

	log.Debug("Call remoteRemoveExecutor of bubbleContract", "chainID", bc.tokenPlugin.ChainID, "txHash", txHash.Hex(), "blockNumber", blockNumber.Uint64(),
		"from", from, "address", address)

	if !bc.Contract.UseGas(params.RemoteRemoveExecutor) {
		return nil, ErrOutOfGas
	}

	operator := bc.tokenPlugin.OpConfig.MainChain.OpAddr
	if from != operator {
		log.Error("the sender is not the operator of the layer1", "from", from, "operator", operator)
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteRemoveExecutor", "the sender is not the operator of the layer1", TxRemoteRemoveExecutor,
			bubble.ErrSenderIsNotOperator)
	}

	contract, err := bc.Plugin.GetBubContract(blockHash, address)
	if err != nil || contract == nil {
		log.Error("the contract is not exist", "address", contract)
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteRemoveExecutor", "the contract is not exist", TxRemoteRemoveExecutor,
			bubble.ErrContractNotExist)
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}

	if err := bc.Plugin.DelBubContract(blockHash, address); err != nil {
		log.Error("failed to delete contract info", "error", err.Error())
		return nil, err
	}

	if err := bc.tokenPlugin.StoreL1HashToL2Hash(blockHash, remoteTxHash, txHash); err != nil {
		return nil, token.ErrStoreL1HashToL2Hash
	}

	return txResultHandler(vm.BubbleContractAddr, bc.Evm, "", "", TxRemoteRemoveExecutor, common.NoErr)
}

// remoteDestroyExecutor receive the remoteDestroy transaction from main chain and destroy all contract from the bubble chain
// this function is only executed before release bubble
func (bc *BubbleContract) remoteDestroyExecutor(remoteBlockNumber *big.Int) ([]byte, error) {
	from := bc.Contract.CallerAddress
	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber
	blockHash := bc.Evm.Context.BlockHash

	log.Info("Call remoteDestroyExecutor of bubbleContract", "chainID", bc.tokenPlugin.ChainID, "txHash", txHash.Hex(), "blockNumber", blockNumber.Uint64(),
		"from", from, "remoteBlockNumber", remoteBlockNumber)

	if !bc.Contract.UseGas(params.RemoteDestroyExecutor) {
		return nil, ErrOutOfGas
	}

	operator := bc.tokenPlugin.OpConfig.MainChain.OpAddr
	if from != operator {
		log.Error("the sender is not the operator of the layer1", "from", from, "operator", operator)
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteDestroyExecutor", "the sender is not the operator of the layer1", TxRemoteDestroyExecutor,
			bubble.ErrSenderIsNotOperator)
	}

	contracts, err := bc.Plugin.GetBubContracts(blockHash)
	if err != nil || len(contracts) == 0 {
		log.Error("no contracts needs destroy")
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteDestroyExecutor", "no contracts needs destroy", TxRemoteDestroyExecutor,
			bubble.ErrContractNotExist)
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}

	for _, contract := range contracts {
		// contract code is empty
		code := bc.Evm.StateDB.GetCode(*contract)
		if len(code) == 0 {
			log.Info("the contract code is empty", "address", contract)
			return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteDestroyExecutor", "the contract code is empty", TxRemoteDestroyExecutor,
				bubble.ErrEmptyContractCode.Wrap(fmt.Sprintf("address: %s", contract.Hex())))
		}

		// switch to the evm runtime
		bc.Contract.caller = AccountRef(vm.BubbleContractAddr)
		bc.Contract.CallerAddress = vm.BubbleContractAddr
		bc.Contract.self = AccountRef(*contract)
		bc.Contract.SetCallCode(contract, bc.Evm.StateDB.GetCodeHash(*contract), code)
		// todo: call destroy and send destroy event once
		input, _ := hex.DecodeString("83197ef0")

		// destroy contract error
		if _, err = RunEvm(bc.Evm, bc.Contract, input); err != nil {
			log.Error("contract destroy returned an error", "error", err.Error())
			return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteDestroyExecutor", "contract destroy returned an error", TxRemoteRemoveExecutor,
				bubble.ErrContractReturns.Wrap(err.Error()))
		}

		if err := bc.Plugin.DelBubContract(blockHash, contract); err != nil {
			log.Error("failed to delete contract info", "error", err.Error())
			return nil, err
		}
	}

	return txResultHandler(vm.BubbleContractAddr, bc.Evm, "", "", TxRemoteDestroyExecutor, common.NoErr)
}

// remoteCall call the contract function on the main chain remotely
func (bc *BubbleContract) remoteCall(contract *common.Address, data []byte) ([]byte, error) {
	origin := bc.Evm.Origin
	from := bc.Contract.CallerAddress
	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber

	log.Debug("Call remoteCall of bubbleContract", "chainID", bc.tokenPlugin.ChainID, "txHash", txHash.Hex(), "blockNumber", blockNumber.Uint64(), "from", from,
		"contract", contract, "data", data)

	if !bc.Contract.UseGas(params.RemoteCallGas) {
		return nil, ErrOutOfGas
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}

	if bc.tokenPlugin.IsSubOpNode {
		task := &bubble.RemoteCallTask{
			TxHash:   bc.Evm.StateDB.TxHash(),
			Caller:   origin,
			BubbleID: bc.tokenPlugin.ChainID,
			Contract: *contract,
			Data:     data,
		}

		if err := bc.Plugin.PostRemoteCallTask(task); err != nil {
			log.Error("post remote call task failed", "error", err.Error())
			return nil, err
		}
	}

	return txResultHandler(vm.BubbleContractAddr, bc.Evm, "", "", TxRemoteCall, common.NoErr)
}

// remoteCallExecutor receive the remoteCall transaction from main chain and execute the function from the contract
func (bc *BubbleContract) remoteCallExecutor(remoteTxHash common.Hash, caller *common.Address, contract *common.Address, data []byte) ([]byte, error) {
	from := bc.Contract.CallerAddress
	blockHash := bc.Evm.Context.BlockHash
	blockNumber := bc.Evm.Context.BlockNumber
	txHash := bc.Evm.StateDB.TxHash()

	log.Debug("Call remoteCallExecutor of bubbleContract", "chainID", bc.tokenPlugin.ChainID, "txHash", txHash.Hex(), "blockNumber", blockNumber.Uint64(), "from", from,
		"caller", caller, "remoteTxHash", remoteTxHash, "contract", contract, "data", data)

	if !bc.Contract.UseGas(params.RemoteCallExecutorGas) {
		return nil, ErrOutOfGas
	}

	// sender only operator
	operator := bc.tokenPlugin.OpConfig.MainChain.OpAddr
	if from != operator {
		log.Error("the sender is not the operator of the layer1", "from", from, "operator", operator)
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteCallExecutor", "the sender is not the operator of the layer1", TxRemoteCallExecutor,
			bubble.ErrSenderIsNotOperator)
	}

	if code := bc.Evm.StateDB.GetCode(*contract); len(code) == 0 {
		log.Error("the contract is not exist", "address", contract)
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteCallExecutor", "the contract is not exist", TxRemoteCallExecutor, bubble.ErrEmptyContractCode)
	}

	// todo: whether to continue when estimating gas
	//if txHash == common.ZeroHash {
	//	return nil, nil
	//}

	// switch to the evm runtime
	bc.Contract.caller = AccountRef(*caller)
	bc.Contract.CallerAddress = *caller
	bc.Contract.self = AccountRef(*contract)
	bc.Contract.SetCallCode(contract, bc.Evm.StateDB.GetCodeHash(*contract), bc.Evm.StateDB.GetCode(*contract))

	vmRet, err := RunEvm(bc.Evm, bc.Contract, data)
	if errors.Is(err, ErrExecutionReverted) {
		reason, errUnpack := abi.UnpackRevert(vmRet)
		info := "execution reverted"
		if errUnpack == nil {
			info = fmt.Sprintf("execution reverted: %v", reason)
		}
		err = newCallContractError(info)
		log.Error("call contract error", "contract", contract, "error", err.Error())
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteCallExecutor", "the contract returned an error", TxRemoteCallExecutor,
			bubble.ErrContractReturns.Wrap(err.Error()))
	} else {
		log.Error("call contract error", "contract", contract, "error", err.Error())
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "remoteCallExecutor", "the contract returned an error", TxRemoteCallExecutor,
			bubble.ErrContractReturns.Wrap(err.Error()))
	}

	if err := bc.tokenPlugin.StoreL1HashToL2Hash(blockHash, remoteTxHash, txHash); err != nil {
		return nil, token.ErrStoreL1HashToL2Hash
	}

	return txResultHandler(vm.BubbleContractAddr, bc.Evm, "", "", TxRemoteCall, common.NoErr)

}
