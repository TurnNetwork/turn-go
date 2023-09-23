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
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/bubble"
	"github.com/bubblenet/bubble/x/plugin"
)

const (
	TxCreateBubble        = 8001
	TxReleaseBubble       = 8002
	TxStakingToken        = 8003
	TxWithdrewToken       = 8004
	TxSettleBubble        = 8005
	CallGetBubbleInfo     = 8100
	CallGetL1HashByL2Hash = 8101
	CallGetBubTxHashList  = 8102
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
func (bc *BubbleContract) createBubble() ([]byte, error) {

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

	if bizErr := bc.Plugin.CheckBubbleElements(blockHash); bizErr != nil {
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "createBubble", bizErr.Error(), TxCreateBubble, bizErr)
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}

	bub, err := bc.Plugin.CreateBubble(blockHash, blockNumber, from, currentNonce, parentHash)
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.BubbleContractAddr, bc.Evm, "createBubble", bizErr.Error(), TxCreateBubble, bizErr)
		} else {
			log.Error("Failed to createBubble", "txHash", txHash, "blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	// send create bubble event to the blockchain Mux if local node is operator
	task := &bubble.CreateBubbleTask{
		BubInfo:  bub,
		BubbleID: bub.Basics.BubbleId,
		TxHash:   txHash,
	}

	for _, operators := range bub.Basics.OperatorsL1 {
		if operators.NodeId == bc.Plugin.NodeID {
			if err := bc.Plugin.PostCreateBubbleEvent(task); err != nil {
				return nil, err
			}
		}
	}

	return txResultHandlerWithRes(vm.BubbleContractAddr, bc.Evm, "", "", TxCreateBubble, int(common.NoErr.Code), bub.Basics.BubbleId), nil
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
			"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "bubbleID", bubbleID, "err", "bubble is not exist")
		return nil, err
	}
	if bub == nil {
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "releaseBubble",
			fmt.Sprintf("bubble %d is not exist", bubbleID), TxReleaseBubble, bubble.ErrBubbleNotExist)
	}

	if from != bub.Basics.Creator {
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "releaseBubble",
			fmt.Sprintf("txSender: %s, bubble Creator: %s", from, bub.Basics.Creator), TxReleaseBubble, bubble.ErrSenderIsNotCreator)
	}

	if bub.State == bubble.ReleasedStatus {
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "releaseBubble", fmt.Sprintf("bubble %d is released ", bub.Basics.BubbleId),
			TxReleaseBubble, bubble.ErrBubbleUnableRelease)
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}

	if err := bc.Plugin.ReleaseBubble(blockHash, blockNumber, bubbleID); err != nil {
		return nil, err
	}

	// send release bubble event to the blockchain Mux if local node is operator
	task := &bubble.ReleaseBubbleTask{
		BubbleID: bub.Basics.BubbleId,
		TxHash:   txHash,
	}

	for _, operators := range bub.Basics.OperatorsL1 {
		if operators.NodeId == bc.Plugin.NodeID {
			if err := bc.Plugin.PostReleaseBubbleEvent(task); err != nil {
				return nil, err
			}
		}
	}

	return txResultHandler(vm.BubbleContractAddr, bc.Evm, "", "", TxReleaseBubble, common.NoErr)
}

// getBubbleInfo return the bubble information by bubble ID
func (bc *BubbleContract) getBubbleInfo(bubbleID *big.Int) ([]byte, error) {
	blockHash := bc.Evm.Context.BlockHash

	bub, err := bc.Plugin.GetBubbleInfo(blockHash, bubbleID)
	if err != nil {
		return callResultHandler(bc.Evm, fmt.Sprintf("getBubbleInfo, bubbleID: %d", bubbleID), bub,
			bubble.ErrBubbleNotExist.Wrap(err.Error())), nil
	}

	return callResultHandler(bc.Evm, fmt.Sprintf("getBubbleInfo, bubbleID: %d", bubbleID), bub, nil), nil
}

// getL1HashByL2Hash The settlement transaction hash of the main chain is obtained according to the sub-chain settlement transaction hash
func (bc *BubbleContract) getL1HashByL2Hash(bubbleID *big.Int, L2TxHash common.Hash) ([]byte, error) {
	blockHash := bc.Evm.Context.BlockHash

	txHash, err := bc.Plugin.GetL1HashByL2Hash(blockHash, bubbleID, L2TxHash)
	if err != nil {
		return callResultHandler(bc.Evm, fmt.Sprintf("getL1HashByL2Hash, bubbleID: %d", bubbleID),
			txHash, bubble.ErrGetL1HashByL2Hash.Wrap(err.Error())), nil
	}

	return callResultHandler(bc.Evm, fmt.Sprintf("getL1HashByL2Hash, bubbleID: %d", bubbleID), txHash, nil), nil
}

// getBubTxHashList Specify BubbleID and transaction type to get a list of bubble's transaction hashes
// Transaction types include: StakingToken, WithdrewToken, SettleBubble
func (bc *BubbleContract) getBubTxHashList(bubbleID *big.Int, txType bubble.BubTxType) ([]byte, error) {
	blockHash := bc.Evm.Context.BlockHash

	txHashList, err := bc.Plugin.GetTxHashListByBub(blockHash, bubbleID, txType)
	if err != nil {
		return callResultHandler(bc.Evm, fmt.Sprintf("getBubTxHashList, bubbleID: %d", bubbleID), txHashList,
			bubble.ErrGetTxHashListByBub.Wrap(err.Error())), nil
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
	// Call handling logic
	ret, err := StakingToken(bc, bubbleID, stakingAsset)
	// estimate gas
	if err == nil && ret == nil {
		return nil, nil
	}
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.BubbleContractAddr, bc.Evm, "stakingToken", bizErr.Error(), TxStakingToken, bizErr)
		} else {
			log.Error("Failed to stakingToken", "txHash", txHash, "blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	return txResultHandlerWithRes(vm.BubbleContractAddr, bc.Evm, "",
		"", TxStakingToken, int(common.NoErr.Code), bubbleID, stakingAsset), nil
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
	// Call handling logic
	accAsset, err := WithdrewToken(bc, bubbleID)
	// estimate gas
	if err == nil && accAsset == nil {
		return nil, nil
	}
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.BubbleContractAddr, bc.Evm, "withdrewToken", bizErr.Error(), TxWithdrewToken, bizErr)
		} else {
			log.Error("Failed to withdrewToken", "txHash", txHash, "blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	return txResultHandlerWithRes(vm.BubbleContractAddr, bc.Evm, "",
		"", TxWithdrewToken, int(common.NoErr.Code), bubbleID, accAsset), nil
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

	// Call handling logic
	ret, err := SettleBubble(bc, L2SettleTxHash, bubbleID, settlementInfo)
	// estimate gas
	if err == nil && ret == nil {
		return nil, nil
	}
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
		"", TxSettleBubble, int(common.NoErr.Code), L2SettleTxHash, bubbleID, settlementInfo), nil
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
	// Get Bubble Basics
	basics, err := bp.GetBubBasics(blockHash, bubbleID)
	if nil != err || nil == basics {
		return nil, bubble.ErrBubbleNotExist
	}

	// check bubble state
	bubState, err := bp.GetBubState(blockHash, bubbleID)
	if nil != err || nil == bubState {
		return nil, bubble.ErrBubbleNotExist
	}
	if *bubState == bubble.ReleasedStatus {
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

	// The transaction hash is empty when gas is estimated
	if bc.Evm.StateDB.TxHash() == common.ZeroHash {
		return nil, nil
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
	if bc.Plugin.NodeID == basics.OperatorsL1[0].NodeId {
		mintTokenTask := bubble.MintTokenTask{
			BubbleID: bubbleID,
			TxHash:   state.TxHash(),
			RPC:      basics.OperatorsL2[0].RPC,
			OpAddr:   basics.OperatorsL1[0].OpAddr,
			AccAsset: &stakingAsset,
		}

		if err := bp.PostMintTokenEvent(&mintTokenTask); err != nil {
			return nil, err
		}
	}

	return []byte{0x1}, nil
}

// WithdrewToken The processing logic of withdrewToken's trading interface
func WithdrewToken(bc *BubbleContract, bubbleID *big.Int) (*bubble.AccountAsset, error) {
	bp := bc.Plugin
	blockHash := bc.Evm.Context.BlockHash
	state := bc.Evm.StateDB

	// Get Bubble Basics
	basics, err := bp.GetBubBasics(blockHash, bubbleID)
	if nil != err || nil == basics {
		return nil, bubble.ErrBubbleNotExist
	}

	// check bubble state
	bubState, err := bp.GetBubState(blockHash, bubbleID)
	if nil != err || nil == bubState {
		return nil, bubble.ErrBubbleNotExist
	}
	// check bubble state
	if *bubState != bubble.ReleasedStatus {
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

	// The transaction hash is empty when gas is estimated
	if bc.Evm.StateDB.TxHash() == common.ZeroHash {
		return nil, nil
	}

	// Store the latest information about the staking assets of the account into bubble
	if err = bp.StoreAccAssetToBub(blockHash, bubbleID, &resetAsset); nil != err {
		return nil, bubble.ErrStoreAccAsset
	}
	// Store the withdrewToken transaction hash
	if err := bp.StoreTxHashToBub(blockHash, bubbleID, state.TxHash(), bubble.WithdrewToken); nil != err {
		return nil, err
	}
	return accAsset, nil
}

// SettleBubble The processing logic of settleBubble's trading interface
func SettleBubble(bc *BubbleContract, L2SettleTxHash common.Hash, bubbleID *big.Int, settlementInfo bubble.SettlementInfo) ([]byte, error) {
	bp := bc.Plugin
	blockHash := bc.Evm.Context.BlockHash
	from := bc.Contract.CallerAddress

	// Get Bubble Basics
	basics, err := bp.GetBubBasics(blockHash, bubbleID)
	if nil != err || nil == basics {
		return nil, bubble.ErrBubbleNotExist
	}

	// check bubble state
	bubState, err := bp.GetBubState(blockHash, bubbleID)
	if nil != err || nil == bubState {
		return nil, bubble.ErrBubbleNotExist
	}
	if *bubState == bubble.ReleasedStatus {
		return nil, bubble.ErrBubbleIsRelease
	}

	// Only the sub-chain operating address has the authority to submit settlement transactions
	if from != basics.OperatorsL2[0].OpAddr {
		return nil, bubble.ErrIsNotSubChainOpAddr
	}

	// Get the account address information
	accList, err := bp.GetAccListOfBub(blockHash, bubbleID)
	if len(accList) != len(settlementInfo.AccAssets) {
		return nil, bubble.ErrSettleAccListIncLength
	}

	var newAccAssets []bubble.AccountAsset
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

		newAccAssets = append(newAccAssets, newAccAsset)
	}

	// The transaction hash is empty when gas is estimated
	if bc.Evm.StateDB.TxHash() == common.ZeroHash {
		return nil, nil
	}

	// Store the latest information about the staking assets of the account into bubble
	for _, newAccAsset := range newAccAssets {
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
	return []byte{0x1}, nil
}
