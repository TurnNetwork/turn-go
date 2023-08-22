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
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/bubble"
	"github.com/bubblenet/bubble/x/plugin"
	"math/big"
)

const (
	TxCreateBubble        = 0001
	TxReleaseBubble       = 0002
	TxStakingToken        = 0003
	TxWithdrewToken       = 0004
	TxSettleBubble        = 0005
	CallGetBubbleInfo     = 0100
	CallGetL1HashByL2Hash = 0101
	CallGetBubTxHashList  = 102
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
		TxCreateBubble:  bc.createBubble,
		TxReleaseBubble: bc.releaseBubble,
		TxStakingToken:  bc.stakingToken,
		TxWithdrewToken: bc.withdrewToken,
		TxSettleBubble:  bc.settleBubble,
		// Get
		CallGetBubbleInfo:     bc.getBubbleInfo,
		CallGetL1HashByL2Hash: bc.getL1HashByL2Hash,
		CallGetBubTxHashList:  bc.getBubTxHashList,
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
func (bc *BubbleContract) releaseBubble(bubbleID *big.Int) ([]byte, error) {

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
func (bc *BubbleContract) getBubbleInfo(bubbleID *big.Int) ([]byte, error) {
	blockHash := bc.Evm.Context.BlockHash

	bub, err := bc.Plugin.GetBubbleInfo(blockHash, bubbleID)
	if err != nil {
		return callResultHandler(bc.Evm, fmt.Sprintf("getBubbleInfo, bubbleID: %d", bubbleID), bub, bubble.ErrBubbleNotExist), err
	}

	return callResultHandler(bc.Evm, fmt.Sprintf("getBubbleInfo, bubbleID: %d", bubbleID), bub, nil), nil
}

// getL1HashByL2Hash The settlement transaction hash of the main chain is obtained according to the sub-chain settlement transaction hash
func (bc *BubbleContract) getL1HashByL2Hash(bubbleID *big.Int, L2TxHash common.Hash) ([]byte, error) {
	blockHash := bc.Evm.Context.BlockHash

	txHash, err := bc.Plugin.GetL1HashByL2Hash(blockHash, bubbleID, L2TxHash)
	if err != nil {
		return callResultHandler(bc.Evm, fmt.Sprintf("getL1HashByL2Hash, bubbleID: %d", bubbleID), txHash, bubble.ErrGetL1HashByL2Hash), err
	}

	return callResultHandler(bc.Evm, fmt.Sprintf("getL1HashByL2Hash, bubbleID: %d", bubbleID), txHash, nil), nil
}

// getBubTxHashList Specify BubbleID and transaction type to get a list of bubble's transaction hashes
// Transaction types include: StakingToken, WithdrewToken, SettleBubble
func (bc *BubbleContract) getBubTxHashList(bubbleID *big.Int, txType bubble.BubTxType) ([]byte, error) {
	blockHash := bc.Evm.Context.BlockHash

	txHashList, err := bc.Plugin.GetTxHashListByBub(blockHash, bubbleID, txType)
	if err != nil {
		return callResultHandler(bc.Evm, fmt.Sprintf("getBubTxHashList, bubbleID: %d", bubbleID), txHashList, bubble.ErrGetTxHashListByBub), err
	}

	return callResultHandler(bc.Evm, fmt.Sprintf("getBubTxHashList, bubbleID: %d", bubbleID), txHashList, nil), nil
}

// stakingToken The account pledges the token to the system contract
// Supports native and ERC20 tokens
// Specifies the bubbleID, and the pledged assets are minted in the specified bubble in the same currency and the same amount
func (bc *BubbleContract) stakingToken(bubbleID *big.Int, stakingAsset bubble.AccountAsset) ([]byte, error) {
	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber
	from := bc.Contract.CallerAddress
	log.Debug("Call BubbleContract of stakingToken", "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "StakingTokenAddr", from)

	// Calculating gas
	if !bc.Contract.UseGas(params.StakingTokenGas) {
		return nil, ErrOutOfGas
	}
	if txHash == common.ZeroHash {
		return nil, nil
	}

	// Call handling logic
	_, err := StakingToken(bc, bubbleID, stakingAsset)
	if nil != err {
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
func (bc *BubbleContract) withdrewToken(bubbleID *big.Int) ([]byte, error) {
	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber
	// Check the bubble status.Only when the bubble state is release can the account redeem the pledged token
	log.Debug("Call BubbleContract of withdrewToken", "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "bubbleID", bubbleID)

	if !bc.Contract.UseGas(params.WithdrewTokenGas) {
		return nil, ErrOutOfGas
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}
	// Call handling logic
	accAsset, err := WithdrewToken(bc, bubbleID)
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.BubbleContractAddr, bc.Evm, "withdrewToken", bizErr.Error(), TxWithdrewToken, bizErr)
		} else {
			log.Error("Failed to withdrewToken", "txHash", txHash, "blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	return txResultHandlerWithRes(vm.BubbleContractAddr, bc.Evm, "",
		"", TxWithdrewToken, int(common.NoErr.Code), accAsset), nil
}

// settleBubble Count the account assets in the bubble and record them
// The mapping relationship between the sub-chain settlement transaction hash and the main chain settlement transaction hash is stored,
// The transaction receipt can be queried through the returned main-chain settlement transaction hash,
// and the actual settlement information can be obtained by parsing the log in the transaction receipt
func (bc *BubbleContract) settleBubble(L2SettleTxHash common.Hash, bubbleID *big.Int, settlementInfo bubble.SettlementInfo) ([]byte, error) {
	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber
	//from := bc.Contract.CallerAddress
	//blockHash := bc.Evm.Context.BlockHash
	//log.Debug("Call mintToken of TokenContract", "blockHash", blockHash, "txHash", txHash.Hex(),
	//	"blockNumber", blockNumber.Uint64(), "caller", from.Hex())

	// Calculating gas
	if !bc.Contract.UseGas(params.SettleBubbleGas) {
		return nil, ErrOutOfGas
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}

	// Call handling logic
	_, err := SettleBubble(bc, L2SettleTxHash, bubbleID, settlementInfo)
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.BubbleContractAddr, bc.Evm, "settleBubble", bizErr.Error(), TxSettleBubble, bizErr)
		} else {
			log.Error("Failed to settleBubble", "txHash", txHash, "blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}
	// log record
	return txResultHandlerWithRes(vm.BubbleContractAddr, bc.Evm, "",
		"", TxSettleBubble, int(common.NoErr.Code), L2SettleTxHash, settlementInfo), nil
}

// StakingToken The processing logic of stakingToken's trading interface
func StakingToken(bc *BubbleContract, bubbleID *big.Int, stakingAsset bubble.AccountAsset) ([]byte, error) {
	blockHash := bc.Evm.Context.BlockHash
	state := bc.Evm.StateDB
	from := bc.Contract.CallerAddress
	blockNumber := bc.Evm.Context.BlockNumber
	bp := bc.Plugin
	if from != stakingAsset.Account {
		return nil, bubble.ErrStakingAccount
	}
	// Get Bubble Information
	bubInfo, err := bp.GetBubbleInfo(blockHash, bubbleID)
	if nil != err || nil == bubInfo {
		return nil, bubble.ErrBubbleNotExist
	}

	// check bubble state
	if bubInfo.State == bubble.ReleasedStatus {
		return nil, bubble.ErrBubbleIsRelease
	}

	// staking native tokens
	// get account balance
	nativeAmount := stakingAsset.NativeAmount
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
			return nil, bubble.ErrEncodeTransferData
		}
		// Execute EVM
		_, err = RunEvm(bc.Evm, contract, input)
		if err != nil {
			log.Error("Failed to Staking ERC20 Token", "error", err)
			return nil, bubble.ErrEVMExecERC20
		}
	}

	// The assets staking by the storage account
	if err := bp.AddAccAssetToBub(blockHash, bubbleID, &stakingAsset); nil != err {
		return nil, err
	}

	// Store the stakingToken transaction hash
	if err := bp.StoreTxHashToBub(blockHash, bubbleID, state.TxHash(), bubble.StakingToken); nil != err {
		return nil, err
	}
	// Send the corresponding minting task
	// Only bubble's main-chain operator node needs to handle this task
	// if bc.Plugin.NodeID == bubInfo.OperatorsL1[0].NodeId
	{
		var mintTokenTask bubble.MintTokenTask
		mintTokenTask.BubbleID = bubbleID
		mintTokenTask.TxHash = state.TxHash()
		mintTokenTask.AccAsset = &stakingAsset
		if err := bp.PostMintTokenEvent(&mintTokenTask); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// WithdrewToken The processing logic of withdrewToken's trading interface
func WithdrewToken(bc *BubbleContract, bubbleID *big.Int) (*bubble.AccountAsset, error) {
	bp := bc.Plugin
	blockHash := bc.Evm.Context.BlockHash
	state := bc.Evm.StateDB
	// Get Bubble Information
	bubInfo, err := bp.GetBubbleInfo(blockHash, bubbleID)
	if nil != err || nil == bubInfo {
		return nil, err
	}
	// check bubble state
	if bubInfo.State != bubble.ReleasedStatus {
		return nil, bubble.ErrBubbleIsNotRelease
	}

	from := bc.Contract.CallerAddress
	// Obtain the staking assets of the account
	accAsset, err := bp.GetAccAssetOfBub(blockHash, bubbleID, from)
	if nil != err || nil == accAsset {
		return nil, err
	}
	var resetAsset bubble.AccountAsset
	resetAsset.Account = from
	// withdrew native tokens
	if accAsset.NativeAmount.Cmp(common.Big0) > 0 {
		// Transfer money from the system address to the corresponding account
		state.SubBalance(vm.BubbleContractAddr, accAsset.NativeAmount)
		state.AddBalance(from, accAsset.NativeAmount)
		resetAsset.NativeAmount = common.Big0
	}

	// withdrew ERC20 tokens
	for _, tokenAsset := range accAsset.TokenAssets {
		// ERC20 Address
		erc20Addr := tokenAsset.TokenAddr
		// Token Amount
		tokenAmount := tokenAsset.Balance
		// Get code based on contract address
		// Check whether ERC20 exists, if not, the failure fails
		code := bc.Evm.StateDB.GetCode(erc20Addr)
		if 0 == len(code) {
			return nil, bubble.ErrERC20NoExist
		}
		contract := bc.Contract

		// Change to ERC20 contract address
		contract.self = AccountRef(erc20Addr)
		// Change the call to the contract address (represents the caller of the contract,
		// the sender of the contract transaction, equivalent to the ERC20 Token forwarder)
		contract.caller = AccountRef(vm.BubbleContractAddr)
		contract.CallerAddress = vm.BubbleContractAddr
		contract.SetCallCode(&erc20Addr, bc.Evm.StateDB.GetCodeHash(erc20Addr), code)
		// Generate data for ERC20 transfer transactions, Transfer from system contract to withdrew account
		input, err := encodeTransferFuncCall(from, tokenAmount)
		if err != nil {
			log.Error("Failed to Withdrew ERC20 Token", "error", err)
			return nil, bubble.ErrEncodeTransferData
		}
		// Execute EVM
		_, err = RunEvm(bc.Evm, contract, input)
		if err != nil {
			log.Error("Failed to Withdrew ERC20 Token", "error", err)
			return nil, bubble.ErrEVMExecERC20
		}
		resetAsset.TokenAssets = append(resetAsset.TokenAssets, bubble.AccTokenAsset{TokenAddr: erc20Addr, Balance: common.Big0})
	}
	// Store the latest information about the staking assets of the account into bubble
	if err = bp.StoreAccAssetToBub(blockHash, bubbleID, &resetAsset); nil != err {
		return nil, bubble.ErrStoreAccAsset
	}
	// Store the withdrewToken transaction hash
	if err := bp.StoreTxHashToBub(blockHash, bubbleID, state.TxHash(), bubble.WithdrewToken); nil != err {
		return nil, err
	}
	return &resetAsset, nil
}

// SettleBubble The processing logic of settleBubble's trading interface
func SettleBubble(bc *BubbleContract, L2SettleTxHash common.Hash, bubbleID *big.Int, settlementInfo bubble.SettlementInfo) ([]byte, error) {
	bp := bc.Plugin
	blockHash := bc.Evm.Context.BlockHash
	// Get Bubble Information
	bubInfo, err := bp.GetBubbleInfo(blockHash, bubbleID)
	if nil != err || nil == bubInfo {
		return nil, bubble.ErrBubbleNotExist
	}
	// check bubble state
	if bubInfo.State == bubble.ReleasedStatus {
		return nil, bubble.ErrCannotSettled
	}

	// Only the child-chain operating address has the authority to submit settlement transactions
	//if from != bubInfo.subChain.opAddr {
	//	return nil, ErrIsNotSubChainOpAddr
	//}

	// Get the account address information
	accList, err := bp.GetAccListOfBub(blockHash, bubbleID)
	if len(accList) != len(settlementInfo.AccAssets) {
		return nil, bubble.ErrSettleAccListIncLength
	}

	for _, accAsset := range settlementInfo.AccAssets {
		account := accAsset.Account
		// Query account assets
		localAsset, err := bp.GetAccAssetOfBub(blockHash, bubbleID, account)
		if nil != err || nil == localAsset {
			return nil, bubble.ErrSettleAccNoExist
		}
		var newAccAsset bubble.AccountAsset
		newAccAsset.Account = account
		// modify the amount of Native Token
		newAccAsset.NativeAmount = accAsset.NativeAmount
		// modify the ERC20 tokens
		for _, tokenAsset := range accAsset.TokenAssets {
			tokenAddr := tokenAsset.TokenAddr
			amount := tokenAsset.Balance
			newAccAsset.TokenAssets = append(newAccAsset.TokenAssets, bubble.AccTokenAsset{TokenAddr: tokenAddr, Balance: amount})
		}

		// Store the latest information about the staking assets of the account into bubble
		if err = bp.StoreAccAssetToBub(blockHash, bubbleID, &newAccAsset); nil != err {
			return nil, bubble.ErrStoreAccAssetToBub
		}
	}

	// The mapping relationship between the sub-chain settlement transaction hash and the main chain settlement transaction hash is stored
	if err = bp.StoreL2HashToL1Hash(blockHash, bubbleID, bc.Evm.StateDB.TxHash(), L2SettleTxHash); nil != err {
		return nil, bubble.ErrStoreL2HashToL1Hash
	}

	// Store the settleBubble transaction hash
	if err := bp.StoreTxHashToBub(blockHash, bubbleID, bc.Evm.StateDB.TxHash(), bubble.SettleBubble); nil != err {
		return nil, err
	}
	return nil, nil
}
