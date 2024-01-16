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
	"net/url"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/gov"
	"github.com/bubblenet/bubble/x/plugin"
	stakingL2 "github.com/bubblenet/bubble/x/stakingL2"
	"github.com/bubblenet/bubble/x/xutil"
)

const (
	TxCreateStakingL2     = 7000
	TxEditorCandidateL2   = 7001
	TxIncreaseStakingL2   = 7002
	TxWithdrewCandidateL2 = 7003
	QueryCandidateListL2  = 7102
	QueryCandidateInfoL2  = 7103
	GetPackageRewardL2    = 7200
	GetStakingRewardL2    = 7201
)

type StakingL2Contract struct {
	Plugin   *plugin.StakingL2Plugin
	Contract *Contract
	Evm      *EVM
}

func (stk *StakingL2Contract) RequiredGas(input []byte) uint64 {
	if checkInputEmpty(input) {
		return 0
	}
	return params.StakingL2Gas
}

func (stk *StakingL2Contract) Run(input []byte) ([]byte, error) {
	if checkInputEmpty(input) {
		return nil, nil
	}
	return execBubbleContract(input, stk.FnSigns())
}

func (stk *StakingL2Contract) CheckGasPrice(gasPrice *big.Int, fcode uint16) error {
	// TODO: coding
	return nil
}

func (stk *StakingL2Contract) FnSigns() map[uint16]interface{} {
	return map[uint16]interface{}{
		// Set
		TxCreateStakingL2:   stk.createStaking,
		TxEditorCandidateL2: stk.editCandidate,
		//TxIncreaseStakingL2:   stk.increaseStaking,
		TxWithdrewCandidateL2: stk.withdrewStaking,
		// Get
		QueryCandidateListL2: stk.getCandidateList,
		QueryCandidateInfoL2: stk.getCandidateInfo,
		// Reward
		GetPackageReward: stk.GetPackageReward,
		GetStakingReward: stk.GetStakingReward,
	}
}

func (stk *StakingL2Contract) createStaking(nodeId discover.NodeID, amount *big.Int, benefitAddress common.Address, name, detail,
	electronURI, rpcURI, p2pURI string, programVersion uint32, blsPubKey bls.PublicKeyHex, isOperator bool) ([]byte, error) {

	txHash := stk.Evm.StateDB.TxHash()
	txIndex := stk.Evm.StateDB.TxIdx()
	blockNumber := stk.Evm.Context.BlockNumber
	blockHash := stk.Evm.Context.BlockHash
	from := stk.Contract.CallerAddress
	state := stk.Evm.StateDB

	log.Debug("Call createStaking of StakingL2Contract", "txHash", txHash.Hex(), "blockNumber", blockNumber.Uint64(),
		"blockHash", blockHash.Hex(), "benefitAddress", benefitAddress.String(), "nodeId", nodeId.String(), "name", name, "detail", detail,
		"amount", amount, "programVersion", programVersion, "from", from, "blsPubKey", blsPubKey)
	if !stk.Contract.UseGas(params.CreateStakeL2Gas) {
		return nil, ErrOutOfGas
	}

	// parse bls publickey
	_, err := blsPubKey.ParseBlsPubKey()
	if nil != err {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "createStaking", fmt.Sprintf("failed to parse blspubkey: %s", err.Error()),
			TxCreateStakingL2, stakingL2.ErrWrongBlsPubKey)
	}

	if !stk.Plugin.CheckStakeThresholdL2(amount) {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "createStaking", fmt.Sprintf("staking threshold: %d, deposit: %d", plugin.StakeThresholdL2, amount),
			TxCreateStakingL2, stakingL2.ErrStakeVonTooLow)
	}

	// Query current active version
	if programVersion < gov.L2Version {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "createStaking",
			fmt.Sprintf("input Version: %s, current valid Version: %s", xutil.ProgramVersion2Str(programVersion), xutil.ProgramVersion2Str(gov.L2Version)),
			TxCreateStakingL2, stakingL2.ErrProgramVersionTooLow)
	} else {
		//If the node version is higher than the current governance version, temporarily use the governance version,  wait for the version to pass the governance proposal, and then replace it
		programVersion = gov.L2Version
	}

	// check whether the candidate exists
	canAddr, err := xutil.NodeId2Addr(nodeId)
	if nil != err {
		log.Error("Failed to createStaking by parse nodeId", "txHash", txHash, "blockNumber", blockNumber, "blockHash", blockHash.Hex(),
			"nodeId", nodeId.String(), "err", err.Error())
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "createStaking", fmt.Sprintf("nodeid %s to address fail: %s",
			nodeId.String(), err.Error()), TxCreateStakingL2, stakingL2.ErrNodeID2Addr)
	}
	canOld, err := stk.Plugin.GetCandidateInfo(blockHash, canAddr)
	if snapshotdb.NonDbNotFoundErr(err) {
		log.Error("Failed to createStaking by GetCandidateInfo", "txHash", txHash, "blockNumber", blockNumber, "err", err.Error())
		return nil, err
	}
	if !canOld.IsEmpty() {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "createStaking", "can is not nil", TxCreateStakingL2, stakingL2.ErrCanAlreadyExist)
	}

	// init candidate info
	canBase := &stakingL2.CandidateBase{
		NodeId:          nodeId,
		Name:            name,
		Version:         programVersion,
		ElectronURI:     electronURI,
		RPCURI:          rpcURI,
		P2PURI:          p2pURI,
		IsOperator:      isOperator,
		StakingAddress:  from,
		BenefitAddress:  benefitAddress,
		StakingBlockNum: blockNumber.Uint64(),
		StakingTxIndex:  txIndex,
		BlsPubKey:       blsPubKey,
		Detail:          detail,
	}
	canMutable := &stakingL2.CandidateMutable{
		Shares:        amount,
		PendingShares: new(big.Int).SetInt64(0),
		LockedShares:  new(big.Int).SetInt64(0),
	}
	can := &stakingL2.Candidate{}
	can.CandidateBase = canBase
	can.CandidateMutable = canMutable

	// check Description length
	if err := can.CheckDescription(); nil != err {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "createStaking", stakingL2.ErrDescriptionLen.Msg+":"+err.Error(),
			TxCreateStakingL2, stakingL2.ErrDescriptionLen)
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}

	// run non-business logic to create candidate
	err = stk.Plugin.CreateCandidate(state, blockHash, blockNumber, amount, canAddr, can)
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "createStaking", bizErr.Error(), TxCreateStakingL2, bizErr)
		} else {
			log.Error("Failed to createStaking by CreateCandidate", "txHash", txHash, "blockNumber", blockNumber, "err", err.Error())
			return nil, err
		}
	}

	return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "", "", TxCreateStakingL2, common.NoErr)
}

func (stk *StakingL2Contract) editCandidate(nodeId discover.NodeID, benefitAddress *common.Address, name, detail, rpcURI *string) ([]byte, error) {

	txHash := stk.Evm.StateDB.TxHash()
	blockNumber := stk.Evm.Context.BlockNumber
	blockHash := stk.Evm.Context.BlockHash
	from := stk.Contract.CallerAddress

	//log.Debug("Call editCandidate of StakingL2Contract", "txHash", txHash.Hex(),
	//	"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", nodeId.String(), "from", from, "benefitAddress", benefitAddress,
	//	"name", *name, "detail", *detail, "rpcURI", *rpcURI)

	if !stk.Contract.UseGas(params.EditCandidateL2Gas) {
		return nil, ErrOutOfGas
	}

	canAddr, err := xutil.NodeId2Addr(nodeId)
	if nil != err {
		log.Error("Failed to editCandidate by parse nodeId", "txHash", txHash,
			"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", nodeId.String(), "err", err.Error())
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "createStaking",
			fmt.Sprintf("nodeid %s to address fail: %s",
				nodeId.String(), err.Error()),
			TxCreateStakingL2, stakingL2.ErrNodeID2Addr)
	}

	canOld, err := stk.Plugin.GetCandidateInfo(blockHash, canAddr)
	if snapshotdb.NonDbNotFoundErr(err) {
		log.Error("Failed to editCandidate by GetCandidateInfo", "txHash", txHash,
			"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "err", err.Error())
		return nil, err
	}

	if canOld.IsEmpty() {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "editCandidate",
			"can is nil", TxEditorCandidateL2, stakingL2.ErrCanNoExist)
	}
	if canOld.Status.IsInvalid() {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "editCandidate",
			fmt.Sprintf("can status is: %d", canOld.Status),
			TxEditorCandidateL2, stakingL2.ErrCanStatusInvalid)
	}

	if from != canOld.StakingAddress {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "editCandidate",
			fmt.Sprintf("contract sender: %s, can stake addr: %s", from, canOld.StakingAddress),
			TxEditorCandidateL2, stakingL2.ErrNoSameStakingAddr)
	}

	if benefitAddress != nil && canOld.BenefitAddress != vm.RewardManagerPoolAddr {
		canOld.BenefitAddress = *benefitAddress
	}
	if name != nil {
		canOld.Name = *name
	}
	if detail != nil {
		canOld.Detail = *detail
	}

	if rpcURI != nil {
		_, err := url.ParseRequestURI(*rpcURI)
		if err != nil {
			return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "editCandidate", "RLP url is incorrect", TxEditorCandidateL2, stakingL2.ErrRlpUrl)
		}
		canOld.RPCURI = *rpcURI
	}

	if err := canOld.CheckDescription(); nil != err {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "editCandidate",
			stakingL2.ErrDescriptionLen.Msg+":"+err.Error(),
			TxEditorCandidateL2, stakingL2.ErrDescriptionLen)
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}
	err = stk.Plugin.EditCandidate(blockHash, blockNumber, canAddr, canOld)
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "editCandidate",
				bizErr.Error(), TxEditorCandidateL2, bizErr)
		} else {
			log.Error("Failed to editCandidate by EditCandidate", "txHash", txHash,
				"blockNumber", blockNumber, "err", err.Error())
			return nil, err
		}
	}

	return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "", "", TxEditorCandidateL2, common.NoErr)
}

//func (stk *StakingL2Contract) increaseStaking(nodeId discover.NodeID, amount *big.Int) ([]byte, error) {
//
//	txHash := stk.Evm.StateDB.TxHash()
//	blockNumber := stk.Evm.Context.BlockNumber
//	blockHash := stk.Evm.Context.BlockHash
//	from := stk.Contract.CallerAddress
//	state := stk.Evm.StateDB
//
//	log.Debug("Call increaseStaking of StakingL2Contract", "txHash", txHash.Hex(),
//		"blockNumber", blockNumber.Uint64(), "nodeId", nodeId.String(),
//		"amount", amount, "from", from)
//
//	if !stk.Contract.UseGas(params.IncStakeL2Gas) {
//		return nil, ErrOutOfGas
//	}
//
//	if ok, threshold := plugin.CheckOperatingThreshold(blockNumber.Uint64(), blockHash, amount); !ok {
//		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "increaseStaking",
//			fmt.Sprintf("increase staking threshold: %d, deposit: %d", threshold, amount),
//			TxIncreaseStakingL2, stakingL2.ErrIncreaseStakeVonTooLow)
//	}
//
//	canAddr, err := xutil.NodeId2Addr(nodeId)
//	if nil != err {
//		log.Error("Failed to increaseStaking by parse nodeId", "txHash", txHash,
//			"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", nodeId.String(), "err", err.Error())
//		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "createStaking",
//			fmt.Sprintf("nodeid %s to address fail: %s",
//				nodeId.String(), err.Error()),
//			TxCreateStakingL2, stakingL2.ErrNodeID2Addr)
//	}
//
//	canOld, err := stk.Plugin.GetCandidateInfo(blockHash, canAddr)
//	if snapshotdb.NonDbNotFoundErr(err) {
//		log.Error("Failed to increaseStaking by GetCandidateInfo", "txHash", txHash,
//			"blockNumber", blockNumber, "err", err.Error())
//		return nil, err
//	}
//
//	if canOld.IsEmpty() {
//		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "increaseStaking",
//			"can is nil", TxIncreaseStakingL2, stakingL2.ErrCanNoExist)
//	}
//
//	if canOld.IsInvalid() {
//		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "increaseStaking",
//			fmt.Sprintf("can status is: %d", canOld.Status),
//			TxIncreaseStakingL2, stakingL2.ErrCanStatusInvalid)
//	}
//
//	if from != canOld.StakingAddress {
//		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "increaseStaking",
//			fmt.Sprintf("contract sender: %s, can stake addr: %s", from, canOld.StakingAddress),
//			TxIncreaseStakingL2, stakingL2.ErrNoSameStakingAddr)
//	}
//	if txHash == common.ZeroHash {
//		return nil, nil
//	}
//
//	err = stk.Plugin.IncreaseStaking(state, blockHash, blockNumber, amount, canAddr, canOld)
//
//	if nil != err {
//		if bizErr, ok := err.(*common.BizError); ok {
//			return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "increaseStaking",
//				bizErr.Error(), TxIncreaseStakingL2, bizErr)
//
//		} else {
//			log.Error("Failed to increaseStaking by EditCandidate", "txHash", txHash,
//				"blockNumber", blockNumber, "err", err.Error())
//			return nil, err
//		}
//
//	}
//	return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "",
//		"", TxIncreaseStakingL2, common.NoErr)
//}

func (stk *StakingL2Contract) withdrewStaking(nodeId discover.NodeID) ([]byte, error) {

	txHash := stk.Evm.StateDB.TxHash()
	blockNumber := stk.Evm.Context.BlockNumber
	blockHash := stk.Evm.Context.BlockHash
	from := stk.Contract.CallerAddress
	state := stk.Evm.StateDB

	log.Debug("Call withdrewStaking of StakingL2Contract", "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "nodeId", nodeId.String(), "from", from)

	if !stk.Contract.UseGas(params.WithdrewStakeL2Gas) {
		return nil, ErrOutOfGas
	}

	canAddr, err := xutil.NodeId2Addr(nodeId)
	if nil != err {
		log.Error("Failed to withdrewStaking by parse nodeId", "txHash", txHash,
			"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", nodeId.String(), "err", err.Error())
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "createStaking",
			fmt.Sprintf("nodeid %s to address fail: %s",
				nodeId.String(), err.Error()),
			TxCreateStakingL2, stakingL2.ErrNodeID2Addr)
	}

	canOld, err := stk.Plugin.GetCandidateInfo(blockHash, canAddr)
	if snapshotdb.NonDbNotFoundErr(err) {
		log.Error("Failed to withdrewStaking by GetCandidateInfo", "txHash", txHash,
			"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", nodeId.String(), "err", err.Error())
		return nil, err
	}

	if canOld.IsEmpty() {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "withdrewStaking",
			"can is nil", TxWithdrewCandidateL2, stakingL2.ErrCanNoExist)
	}

	if canOld.Status.IsInvalid() {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "withdrewStaking",
			fmt.Sprintf("can status is: %d", canOld.Status),
			TxWithdrewCandidateL2, stakingL2.ErrCanStatusInvalid)
	}

	if from != canOld.StakingAddress {
		return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "withdrewStaking",
			fmt.Sprintf("contract sender: %s, can stake addr: %s", from, canOld.StakingAddress),
			TxWithdrewCandidateL2, stakingL2.ErrNoSameStakingAddr)
	}
	if txHash == common.ZeroHash {
		return nil, nil
	}
	err = stk.Plugin.WithdrewStaking(state, blockHash, blockNumber, canAddr, canOld)
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "withdrewStaking",
				bizErr.Error(), TxWithdrewCandidateL2, bizErr)
		} else {
			log.Error("Failed to withdrewStaking by WithdrewStaking", "txHash", txHash,
				"blockNumber", blockNumber, "err", err.Error())
			return nil, err
		}

	}

	return txResultHandler(vm.StakingL2ContractAddr, stk.Evm, "",
		"", TxWithdrewCandidateL2, common.NoErr)
}

func (stk *StakingL2Contract) getCandidateList() ([]byte, error) {

	blockNumber := stk.Evm.Context.BlockNumber
	blockHash := stk.Evm.Context.BlockHash

	arr, err := stk.Plugin.GetCandidateList(blockHash, blockNumber.Uint64())
	if snapshotdb.NonDbNotFoundErr(err) {
		return callResultHandler(stk.Evm, "getCandidateList", arr, stakingL2.ErrGetCandidateList.Wrap(err.Error())), nil
	}
	if snapshotdb.IsDbNotFoundErr(err) || len(arr) == 0 {
		return callResultHandler(stk.Evm, "getCandidateList", arr, stakingL2.ErrGetCandidateList.Wrap("CandidateList info is not found")), nil
	}
	return callResultHandler(stk.Evm, "getCandidateList", arr, nil), nil
}

func (stk *StakingL2Contract) getCandidateInfo(nodeId discover.NodeID) ([]byte, error) {
	blockNumber := stk.Evm.Context.BlockNumber
	blockHash := stk.Evm.Context.BlockHash

	canAddr, err := xutil.NodeId2Addr(nodeId)
	if nil != err {
		return callResultHandler(stk.Evm, fmt.Sprintf("getCandidateInfo, nodeId: %s",
			nodeId), nil, stakingL2.ErrQueryCandidateInfo.Wrap(err.Error())), nil
	}
	can, err := stk.Plugin.GetFacadeCandidateInfo(blockHash, blockNumber.Uint64(), canAddr)
	if snapshotdb.NonDbNotFoundErr(err) {
		return callResultHandler(stk.Evm, fmt.Sprintf("getCandidateInfo, nodeId: %s",
			nodeId), can, stakingL2.ErrQueryCandidateInfo.Wrap(err.Error())), nil
	}
	if snapshotdb.IsDbNotFoundErr(err) || can.IsEmpty() {
		return callResultHandler(stk.Evm, fmt.Sprintf("getCandidateInfo, nodeId: %s",
			nodeId), can, stakingL2.ErrQueryCandidateInfo.Wrap("Candidate info is not found")), nil
	}

	return callResultHandler(stk.Evm, fmt.Sprintf("getCandidateInfo, nodeId: %s",
		nodeId), can, nil), nil
}

func (stk *StakingL2Contract) GetPackageReward() ([]byte, error) {
	packageReward, err := plugin.LoadNewBlockReward(common.ZeroHash, stk.Evm.SnapshotDB)
	if nil != err {
		return callResultHandler(stk.Evm, "GetPackageReward", nil, common.NotFound.Wrap(err.Error())), nil
	}

	return callResultHandler(stk.Evm, "GetPackageReward", (*hexutil.Big)(packageReward), nil), nil
}

func (stk *StakingL2Contract) GetStakingReward() ([]byte, error) {
	stakingReward, err := plugin.LoadStakingReward(common.ZeroHash, stk.Evm.SnapshotDB)
	if nil != err {
		return callResultHandler(stk.Evm, "GetStakingReward", nil, common.NotFound.Wrap(err.Error())), nil
	}

	return callResultHandler(stk.Evm, "GetStakingReward", (*hexutil.Big)(stakingReward), nil), nil
}
