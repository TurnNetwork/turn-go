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
	"fmt"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/x/bubble"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/plugin"
)

const (
	TxCreateBubble     = 0001
	TxReleaseBubble    = 0002
	TxStakingToken     = 0003
	TxWithdrewToken    = 0004
	TxSettlementBubble = 0005
	CallGetBubbleInfo  = 1001
)

type BubbleContract struct {
	Plugin   *plugin.BubblePlugin
	Contract *Contract
	Evm      *EVM
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
		TxCreateBubble:     bc.createBubble,
		TxReleaseBubble:    bc.releaseBubble,
		TxStakingToken:     bc.stakingToken,
		TxWithdrewToken:    bc.withdrewToken,
		TxSettlementBubble: bc.settlementBubble,
		// Get
		CallGetBubbleInfo: bc.getBubbleInfo,
	}
}

func (bc *BubbleContract) CheckGasPrice(gasPrice *big.Int, fcode uint16) error {
	return nil
}

// createBubble create a Bubble chain using operator nodes and candidate nodes
func (bc *BubbleContract) createBubble(genesisData [][]byte) ([]byte, error) {

	from := bc.Contract.CallerAddress
	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber
	blockHash := bc.Evm.Context.BlockHash
	currentNonce := bc.Evm.StateDB.GetNonce(from)
	parentHash := bc.Evm.Context.ParentHash

	log.Debug("Call createBubble of bubbleContract", "blockNumber", blockNumber.Uint64(), "blockHash", blockHash.TerminalString(),
		"txHash", txHash.Hex(), "from", from.String())

	if !bc.Contract.UseGas(params.CreateBubbleGas) {
		return nil, ErrOutOfGas
	}

	bubbleID, err := bc.Plugin.CreateBubble(blockHash, blockNumber, from, currentNonce, parentHash)
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.BubbleContractAddr, bc.Evm, "createBubble", bizErr.Error(), TxCreateBubble, bizErr)
		} else {
			log.Error("Failed to createBubble", "txHash", txHash, "blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	// TODO: store genesisData and return the index for DApp reuse
	// TODO: store genesisData by a other interface

	return txResultHandlerWithRes(vm.BubbleContractAddr, bc.Evm, "", "", TxCreateBubble, int(common.NoErr.Code), bubbleID), nil
}

// releaseBubble release the node resources of a bubble chain and delete it`s information
func (bc *BubbleContract) releaseBubble(bubbleID uint32) ([]byte, error) {

	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber
	blockHash := bc.Evm.Context.BlockHash
	from := bc.Contract.CallerAddress

	log.Debug("Call releaseBubble of bubbleContract", "blockNumber", blockNumber.Uint64(),
		"blockHash", blockHash.TerminalString(), "txHash", txHash.Hex(), "from", from.String())

	if !bc.Contract.UseGas(params.ReleaseBubbleGas) {
		return nil, ErrOutOfGas
	}

	bub, err := bc.Plugin.GetBubbleInfo(blockHash, bubbleID)
	if snapshotdb.NonDbNotFoundErr(err) {
		log.Error("Failed to releaseBubble by GetBubbleInfo", "txHash", txHash,
			"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", bub.BubbleId, "err", err)
		return nil, err
	}
	if bub == nil {
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "releaseBubble",
			fmt.Sprintf("bubble %d is not exist", bub.BubbleId), TxReleaseBubble, bubble.ErrBubbleNotExist)
	}

	if from != bub.Creator {
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "releaseBubble",
			fmt.Sprintf("txSender: %s, bubble Creator: %s", from, bub.Creator), TxReleaseBubble, bubble.ErrSenderIsNotCreator)
	}

	// TODO: can release the bubble chain in the building state？
	if bub.State != bubble.PreReleaseStatus {
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "releaseBubble", fmt.Sprintf("bubble %d unable to release ", bub.BubbleId),
			TxReleaseBubble, bubble.ErrBubbleUnableRelease)
	}

	if err := bc.Plugin.ReleaseBubble(blockHash, blockNumber, bubbleID); err != nil {
		return nil, err
	}

	return txResultHandler(vm.BubbleContractAddr, bc.Evm, "", "", TxReleaseBubble, common.NoErr)
}

// getBubbleInfo return the bubble information by bubble ID
func (bc *BubbleContract) getBubbleInfo(bubbleID uint32) ([]byte, error) {
	blockHash := bc.Evm.Context.BlockHash

	bub, err := bc.Plugin.GetBubbleInfo(blockHash, bubbleID)
	if err != nil {
		return callResultHandler(bc.Evm, fmt.Sprintf("getBubbleInfo, bubbleID: %d", bubbleID), bub, bubble.ErrBubbleNotExist), err
	}

	return callResultHandler(bc.Evm, fmt.Sprintf("getBubbleInfo, bubbleID: %d", bubbleID), bub, nil), nil
}

// stakingToken The account pledges the token to the system contract
// Supports native and ERC20 tokens
// Specifies the bubbleID, and the pledged assets are minted in the specified bubble in the same currency and the same amount
func (bc *BubbleContract) stakingToken(bubbleID *big.Int, stakingAsset bubble.AccountAsset) ([]byte, error) {
	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber
	blockHash := bc.Evm.Context.BlockHash
	from := bc.Contract.CallerAddress
	state := bc.Evm.StateDB
	// Get Bubble Information
	bubInfo, err := bc.getBubbleInfo(uint32(bubbleID.Uint64()))
	if nil != err {
		return nil, err
	}
	if from != stakingAsset.Account {
		return nil, bubble.ErrStakingAccount
	}
	nativeAmount := stakingAsset.NativeAmount
	log.Debug("Call BubbleContract of stakingToken", "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "StakingTokenAddr", from, "amount", nativeAmount,
		"bubbleInfo", bubInfo)

	// Calculating gas
	if !bc.Contract.UseGas(params.StakingTokenGas) {
		return nil, ErrOutOfGas
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}
	// staking native tokens
	// get account balance
	origin := state.GetBalance(from)
	if origin.Cmp(nativeAmount) < 0 {
		log.Error("Failed to Staking Token: the account's balance is not Enough",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(),
			"stakeAddr", from, "balance", origin, "stakingAmount", nativeAmount)
		return nil, bubble.ErrAccountNoEnough
	}

	if stakingAsset.NativeAmount.Cmp(common.Big0) > 0 {
		// Deduct the account's native token
		state.SubBalance(from, nativeAmount)
		// The native token of the account is added to the system contract
		state.AddBalance(vm.BubbleContractAddr, nativeAmount)
	}

	// Staking ERC20 Token
	for _, tokenAsset := range stakingAsset.TokenAssets {
		// ERC20 Address
		erc20Addr := tokenAsset.TokenAddr
		// Token Amount
		tokenAmount := tokenAsset.Balance
		// Get code based on contract address
		// Check whether ERC20 exists, if not, the pledge fails
		code := bc.Evm.StateDB.GetCode(erc20Addr)
		if 0 == len(code) {
			return nil, bubble.ErrERC20NoExist
		}
		contract := bc.Contract
		// Change to ERC20 contract address
		contract.self = AccountRef(erc20Addr)
		contract.SetCallCode(&erc20Addr, bc.Evm.StateDB.GetCodeHash(erc20Addr), code)
		// Generate data for ERC20 transfer transactions
		input, err := encodeTransferFuncCall(vm.BubbleContractAddr, tokenAmount)
		if err != nil {
			log.Error("Failed to Staking ERC20 Token", "error", err)
			return nil, err
		}
		// Execute EVM
		_, err = RunEvm(bc.Evm, contract, input)
		if err != nil {
			log.Error("Failed to Staking ERC20 Token", "error", err)
			return nil, err
		}
	}

	// The assets pledged by the storage account
	if err := bc.Plugin.StoreAccStakingAsset(blockHash, uint32(bubbleID.Uint64()), &stakingAsset); nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.BubbleContractAddr, bc.Evm, "stakingToken", bizErr.Error(), TxStakingToken, bizErr)
		} else {
			log.Error("Failed to stakingToken", "txHash", txHash, "blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	return txResultHandlerWithRes(vm.BubbleContractAddr, bc.Evm, "",
		"", TxStakingToken, int(common.NoErr.Code), stakingAsset), nil
}

// withdrewToken Redeem account tokens, including native tokens and ERC20 tokens
func (bc *BubbleContract) withdrewToken(bubbleID uint32) ([]byte, error) {
	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber
	//from := bc.Contract.CallerAddress
	//state := bc.Evm.StateDB
	// Get Bubble Information
	bubInfo, err := bc.getBubbleInfo(bubbleID)
	if nil != err {
		return nil, err
	}
	log.Debug("Call BubbleContract of withdrewToken", "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "bubbleID", bubbleID, "bubInfo", bubInfo)

	// 计算gas
	if !bc.Contract.UseGas(params.BubbleGas) {
		return nil, ErrOutOfGas
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}
	return nil, nil
	//// 赎回原生Token
	//// 1.从DB中获取Bubble内对应账户可提取余额
	//// 2.从系统地址中转账amount给对应的账户
	//// origin := state.GetBalance(from)
	////if origin.Cmp(amount) < 0 {
	////	log.Error("Failed to Staking Token: the account's balance is not Enough",
	////		"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(),
	////		"stakeAddr", from, "balance", origin, "stakingAmount", amount)
	////	return nil, staking.ErrAccountVonNoEnough
	////}
	//state.SubBalance(vm.StakingContractAddr, amount)
	//state.AddBalance(from, amount)
	//// 更新账户原生代币质押记录为0
	//
	//// 赎回ERC20 Token
	//// 1.从DB中获取Bubble内对应账户可提取的ERC20信息
	//// 1.1 获取ERC20地址列表
	//// 2.根据合约地址获取code
	//code := bc.Evm.StateDB.GetCode(erc20Addr)
	//if len(code) != 0 {
	//	contract := bc.Contract
	//	// 2.修改调用为合约地址（表示合约的调用者，合约交易的发送者，相当于ERC20 Token转出方）
	//	contract.caller = AccountRef(vm.StakingContractAddr)
	//	contract.CallerAddress = vm.StakingContractAddr
	//	// 修改成ERC20合约地址
	//	contract.self = AccountRef(erc20Addr)
	//	contract.SetCallCode(&erc20Addr, bc.Evm.StateDB.GetCodeHash(erc20Addr), code)
	//	for _, interpreter := range bc.Evm.interpreters {
	//		if interpreter.CanRun(contract.Code) {
	//			if bc.Evm.interpreter != interpreter {
	//				// Ensure that the interpreter pointer is set back
	//				// to its current value upon return.
	//				defer func(i Interpreter) {
	//					bc.Evm.interpreter = i
	//				}(bc.Evm.interpreter)
	//				bc.Evm.interpreter = interpreter
	//			}
	//			input, err := encodeTransferFuncCall(from, transferERC20Amount)
	//			if err != nil {
	//				log.Error("Failed to Staking ERC20 Token", "error", err)
	//				return nil, err
	//			}
	//			ret, err := interpreter.Run(contract, input, false)
	//			if err != nil {
	//				log.Error("Failed to Staking ERC20 Token", "ret", ret, "error", err)
	//				return ret, err
	//			}
	//			// 保存ERC20代币质押记录
	//		}
	//	}
	//}
	//
	//return txResultHandlerWithRes(vm.BubbleContractAddr, bc.Evm, "",
	//	"", TxWithdrewToken, int(common.NoErr.Code), amount), nil
}

// 主链结算接口
func (bc *BubbleContract) settlementBubble(bubbleId *big.Int, settlementInfo bubble.SettlementInfo) ([]byte, error) {

	// 1.根据BubbleID获取Bubble信息

	// 2.根据Bubble信息中的子链运营节点信息中的运营节点地址判断是否有权限调用
	// from := bc.Contract.CallerAddress

	blockNumber := bc.Evm.Context.BlockNumber
	// state := bc.Evm.StateDB
	stHash, err := settlementInfo.Hash()
	if err != nil {
		log.Error("Failed to calculate Hash for SettlementInfo", "blockNumber", blockNumber, "blockHash", "err", err)
		return nil, err
	}
	log.Info("SettlementInfo hash:", stHash)
	txHash := bc.Evm.StateDB.TxHash()

	//log.Debug("Call withdrewDelegation of withdrewToken", "txHash", txHash.Hex(),
	//	"blockNumber", blockNumber.Uint64(), "StakingTokenAddr", from, "amount", amount)

	// 计算gas
	if !bc.Contract.UseGas(params.WithdrewDelegationGas) {
		return nil, ErrOutOfGas
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}

	// 1.铸币原生Token
	// 1.1 从系统合约账户向account转账
	//if mintAmount.Cmp(common.Big0) > 0 {
	//	state.SubBalance(vm.StakingContractAddr, mintAmount)
	//	state.AddBalance(receiveAccount, mintAmount)
	//}
	//
	//// 2.铸币ERC20代币（默认精度为6）
	//// 2.1 判断是否ERC20是否存在，不存在则需要部署
	//code := bc.Evm.StateDB.GetCode(erc20Addr)
	//contract := bc.Contract
	//if len(code) == 0 {
	//	tmpErc20Addr := common.HexToAddress("0x1111000000000000000000000000000000000001")
	//	tempCode := bc.Evm.StateDB.GetCode(tmpErc20Addr)
	//	// 部署合约
	//	code = tempCode
	//	bc.Evm.StateDB.SetCode(erc20Addr, code)
	//	// 初始化
	//}
	//// 开始铸ERC20 Token币
	//// 2.修改调用为合约地址（表示合约的调用者，合约交易的发送者）
	//contract.caller = AccountRef(vm.StakingContractAddr)
	//contract.CallerAddress = vm.StakingContractAddr
	//// 修改成ERC20合约地址
	//contract.self = AccountRef(erc20Addr)
	//contract.SetCallCode(&erc20Addr, bc.Evm.StateDB.GetCodeHash(erc20Addr), code)
	//for _, interpreter := range bc.Evm.interpreters {
	//	if interpreter.CanRun(contract.Code) {
	//		if bc.Evm.interpreter != interpreter {
	//			// Ensure that the interpreter pointer is set back
	//			// to its current value upon return.
	//			defer func(i Interpreter) {
	//				bc.Evm.interpreter = i
	//			}(bc.Evm.interpreter)
	//			bc.Evm.interpreter = interpreter
	//		}
	//		input, err := encodeMintFuncCall(from, mintERC20Amount)
	//		if err != nil {
	//			log.Error("Failed to Mint ERC20 Token", "error", err)
	//			return nil, err
	//		}
	//		ret, err := interpreter.Run(contract, input, false)
	//		if err != nil {
	//			log.Error("Failed to Mint ERC20 Token", "ret", ret, "error", err)
	//			return ret, err
	//		}
	//		// 保存ERC20代币质押记录
	//	}
	//}

	//return txResultHandlerWithRes(vm.BubbleContractAddr, bc.Evm, "",
	//	"", TxSettlementBubble, int(common.NoErr.Code), mintAmount), nil
	return nil, nil
}
