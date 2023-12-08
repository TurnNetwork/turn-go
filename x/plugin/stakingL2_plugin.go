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

package plugin

import (
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/rlp"
	"math/big"
	"sync"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/x/stakingL2"
	"github.com/bubblenet/bubble/x/xcom"
	"github.com/bubblenet/bubble/x/xutil"
)

var (
	StakeThresholdL2 = new(big.Int).Mul(big.NewInt(params.BUB), big.NewInt(200))
)

type StakingL2Plugin struct {
	db *stakingL2.StakingDB
}

var (
	stkL2Oncer sync.Once
	stkL2      *StakingL2Plugin
)

// StakingL2Instance return the StakingL2Plugin object by singleton pattern
func StakingL2Instance() *StakingL2Plugin {
	stkL2Oncer.Do(func() {
		log.Info("Init StakingL2 plugin ...")
		stkL2 = &StakingL2Plugin{
			db: stakingL2.NewStakingDB(),
		}
	})
	return stkL2
}

func (sk *StakingL2Plugin) BeginBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	return nil
}

func (sk *StakingL2Plugin) EndBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	epoch := xutil.CalculateEpoch(header.Number.Uint64())

	if xutil.IsEndOfEpoch(header.Number.Uint64()) {
		err := sk.HandleUnCandidateItem(state, header.Number.Uint64(), blockHash, epoch)
		if nil != err {
			log.Error("Failed to call HandleUnCandidateItem on StakingL2Plugin EndBlock",
				"blockNumber", header.Number.Uint64(), "blockHash", blockHash.Hex(), "err", err.Error())
			return err
		}
	}

	return nil
}

func (sk *StakingL2Plugin) Confirmed(nodeId discover.NodeID, block *types.Block) error {
	return nil
}

func (sk *StakingL2Plugin) GetCandidateInfo(blockHash common.Hash, addr common.NodeAddress) (*stakingL2.Candidate, error) {
	return sk.db.GetCandidateStore(blockHash, addr)
}

func (sk *StakingL2Plugin) GetCanBase(blockHash common.Hash, addr common.NodeAddress) (*stakingL2.CandidateBase, error) {
	return sk.db.GetCanBaseStore(blockHash, addr)
}

func (sk *StakingL2Plugin) GetCanMutable(blockHash common.Hash, addr common.NodeAddress) (*stakingL2.CandidateMutable, error) {
	return sk.db.GetCanMutableStore(blockHash, addr)
}

func (sk *StakingL2Plugin) GetFacadeCandidateInfo(blockHash common.Hash, blockNumber uint64, addr common.NodeAddress) (*stakingL2.MarshalAbleCandidate, error) {
	can, err := sk.GetCandidateInfo(blockHash, addr)
	if nil != err {
		return nil, err
	}

	epoch := xutil.CalculateEpoch(blockNumber)
	lazyCalcL2StakeAmount(epoch, can.CandidateMutable)
	canHex := buildMarshalAbleCandidate(can)
	return canHex, nil
}

func (sk *StakingL2Plugin) CreateCandidate(state xcom.StateDB, blockHash common.Hash, blockNumber, amount *big.Int,
	addr common.NodeAddress, can *stakingL2.Candidate) error {

	origin := state.GetBalance(can.StakingAddress)
	if origin.Cmp(amount) < 0 {
		log.Error("Failed to CreateCandidate on StakingL2Plugin: the account free von is not Enough",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(),
			"stakeAddr", can.StakingAddress, "originVon", origin, "stakingVon", amount)
		return stakingL2.ErrAccountVonNoEnough
	}
	state.SubBalance(can.StakingAddress, amount)
	state.AddBalance(vm.StakingContractAddr, amount)
	can.PendingShares = amount
	can.StakingEpoch = uint32(xutil.CalculateEpoch(blockNumber.Uint64()))

	if err := sk.db.SetCandidateStore(blockHash, addr, can); nil != err {
		log.Error("Failed to CreateCandidate on StakingL2Plugin: Store Candidate info is failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err.Error())
		return err
	}

	if can.IsOperator {
		if err := sk.db.SetOperatorStore(blockHash, addr, can); nil != err {
			log.Error("Failed to CreateCandidate on StakingL2Plugin: Store Operator info is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err.Error())
			return err
		}
	} else {
		if err := sk.db.SetCommitteeStore(blockHash, addr, can); nil != err {
			log.Error("Failed to CreateCandidate on StakingL2Plugin: Store Committee info is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err.Error())
			return err
		}

		if err := sk.db.AddCommitteeCount(blockHash, 1); err != nil {
			log.Error("Failed to CreateCandidate on StakingL2Plugin: add Committee count is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err.Error())
			return err
		}
	}

	// add the account staking Reference Count
	if err := sk.db.AddAccountStakeRc(blockHash, can.StakingAddress); nil != err {
		log.Error("Failed to CreateCandidate on StakingL2Plugin: Store Staking Account Reference Count (add) is failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "NodeID", can.NodeId.String(),
			"staking Account", can.StakingAddress.String(), "err", err.Error())
		return err
	}

	return nil
}

func (sk *StakingL2Plugin) EditCandidate(blockHash common.Hash, blockNumber *big.Int, canAddr common.NodeAddress, can *stakingL2.Candidate) error {
	if err := sk.db.SetCanBaseStore(blockHash, canAddr, can.CandidateBase); nil != err {
		log.Error("Failed to EditCandidate on StakingL2Plugin: Store CandidateBase info is failed", "nodeId", can.NodeId.String(),
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "err", err.Error())
		return err
	}
	if err := sk.db.SetCanMutableStore(blockHash, canAddr, can.CandidateMutable); nil != err {
		log.Error("Failed to EditCandidate on StakingL2Plugin: Store CandidateMutable info is failed", "nodeId", can.NodeId.String(),
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "err", err.Error())
		return err
	}
	return nil
}

//func (sk *StakingL2Plugin) IncreaseStaking(state xcom.StateDB, blockHash common.Hash, blockNumber,
//	amount *big.Int, canAddr common.NodeAddress, can *stakingL2.Candidate) error {
//
//	epoch := xutil.CalculateEpoch(blockNumber.Uint64())
//
//	lazyCalcL2StakeAmount(epoch, can.CandidateMutable)
//
//	origin := state.GetBalance(can.StakingAddress)
//	if origin.Cmp(amount) < 0 {
//		log.Error("Failed to IncreaseStaking on StakingL2Plugin: the account free von is not Enough",
//			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(),
//			"nodeId", can.NodeId.String(), "account", can.StakingAddress,
//			"originVon", origin, "stakingVon", amount)
//		return stakingL2.ErrAccountVonNoEnough
//	}
//	state.SubBalance(can.StakingAddress, amount)
//	state.AddBalance(vm.StakingContractAddr, amount)
//	can.PendingShares = new(big.Int).Add(can.PendingShares, amount)
//
//	if err := sk.db.DelCanPowerStore(blockHash, can); nil != err {
//		log.Error("Failed to IncreaseStaking on StakingL2Plugin: Delete Candidate old power is failed",
//			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(),
//			"nodeId", can.NodeId.String(), "err", err.Error())
//		return err
//	}
//
//	can.StakingEpoch = uint32(epoch)
//	can.AddShares(amount)
//
//	if err := sk.db.SetCanPowerStore(blockHash, canAddr, can); nil != err {
//		log.Error("Failed to IncreaseStaking on StakingL2Plugin: Store Candidate new power is failed",
//			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(),
//			"nodeId", can.NodeId.String(), "err", err.Error())
//		return err
//	}
//
//	if err := sk.db.SetCanMutableStore(blockHash, canAddr, can.CandidateMutable); nil != err {
//		log.Error("Failed to IncreaseStaking on StakingL2Plugin: Store CandidateMutable info is failed",
//			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(),
//			"nodeId", can.NodeId.String(), "err", err.Error())
//		return err
//	}
//
//	return nil
//}

func (sk *StakingL2Plugin) WithdrewStaking(state xcom.StateDB, blockHash common.Hash, blockNumber *big.Int,
	canAddr common.NodeAddress, can *stakingL2.Candidate) error {

	epoch := xutil.CalculateEpoch(blockNumber.Uint64())

	lazyCalcL2StakeAmount(epoch, can.CandidateMutable)

	if err := sk.withdrewStakeAmount(state, blockHash, blockNumber.Uint64(), epoch, canAddr, can); nil != err {
		return err
	}

	can.StakingEpoch = uint32(epoch)

	if can.LockedShares.Cmp(common.Big0) > 0 {
		if err := sk.db.SetCanMutableStore(blockHash, canAddr, can.CandidateMutable); nil != err {
			log.Error("Failed to WithdrewStaking on StakingL2Plugin: Store CandidateMutable info is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err.Error())
			return err
		}
	} else {
		// delete node store
		if can.IsOperator {
			if err := sk.db.DelOperatorStore(blockHash, canAddr); err != nil {
				log.Error("Failed to WithdrewStaking on StakingL2Plugin: Delete Operator info is failed",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err.Error())
				return err
			}
		} else {
			if err := sk.db.DelCommitteeStore(blockHash, canAddr); err != nil {
				log.Error("Failed to WithdrewStaking on StakingL2Plugin: Delete Operator info is failed",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err.Error())
				return err
			}
		}
		if err := sk.db.DelCandidateStore(blockHash, canAddr); nil != err {
			log.Error("Failed to WithdrewStaking on StakingL2Plugin: Delete Candidate info is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err.Error())
			return err
		}
	}

	// sub the account staking Reference Count
	if err := sk.db.SubAccountStakeRc(blockHash, can.StakingAddress); nil != err {
		log.Error("Failed to WithdrewStaking on StakingL2Plugin: Store Staking Account Reference Count (sub) is failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(),
			"staking Account", can.StakingAddress.String(), "err", err.Error())
		return err
	}

	return nil
}

func (sk *StakingL2Plugin) withdrewStakeAmount(state xcom.StateDB, blockHash common.Hash, blockNumber, epoch uint64,
	canAddr common.NodeAddress, can *stakingL2.Candidate) error {

	// Direct return of money during the hesitation period
	// Return according to the way of coming
	if can.PendingShares.Cmp(common.Big0) > 0 {
		state.AddBalance(can.StakingAddress, can.PendingShares)
		state.SubBalance(vm.StakingContractAddr, can.PendingShares)
		can.PendingShares = new(big.Int).SetInt64(0)
	}

	if can.LockedShares.Cmp(common.Big0) > 0 {
		if err := sk.addUnStakeRecord(state, blockNumber, blockHash, epoch, can.NodeId, canAddr, can.StakingBlockNum); nil != err {
			log.Error("Failed to WithdrewStaking on StakingL2Plugin: Add UnStakeRecordStore failed",
				"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err.Error())
			return err
		}
	}

	can.CleanShares()
	can.Status |= stakingL2.Invalid

	return nil
}

func (sk *StakingL2Plugin) HandleUnCandidateItem(state xcom.StateDB, blockNumber uint64, blockHash common.Hash, epoch uint64) error {

	unStakeCount, err := sk.db.GetUnStakeCountStore(blockHash, epoch)
	switch {
	case snapshotdb.NonDbNotFoundErr(err):
		return err
	case snapshotdb.IsDbNotFoundErr(err):
		unStakeCount = 0
	}

	if unStakeCount == 0 {
		return nil
	}

	filterAddr := make(map[common.NodeAddress]struct{})

	for index := 1; index <= int(unStakeCount); index++ {

		stakeItem, err := sk.db.GetUnStakeRecordStore(blockHash, epoch, uint64(index))
		if nil != err {
			log.Error("Failed to HandleUnCandidateItem: Query the unStakeRecord node addr is failed",
				"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err.Error())
			return err
		}

		canAddr := stakeItem.NodeAddress

		if _, ok := filterAddr[canAddr]; ok {
			if err := sk.db.DelUnStakeRecordStore(blockHash, epoch, uint64(index)); nil != err {
				log.Error("Failed to HandleUnCandidateItem: Delete already handle unStakeRecord failed",
					"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err.Error())
				return err
			}
			continue
		}

		can, err := sk.db.GetCandidateStore(blockHash, canAddr)
		if snapshotdb.NonDbNotFoundErr(err) {
			log.Error("Failed to HandleUnCandidateItem: Query candidate failed",
				"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "canAddr", canAddr.Hex(), "err", err.Error())
			return err
		}

		// This should not be nil
		if snapshotdb.IsDbNotFoundErr(err) || can.IsEmpty() {

			if err := sk.db.DelUnStakeRecordStore(blockHash, epoch, uint64(index)); nil != err {
				log.Error("Failed to HandleUnCandidateItem: Candidate is no exist, Delete unStakeRecord failed",
					"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err.Error())
				return err
			}

			continue
		}

		if stakeItem.StakingBlockNum != can.StakingBlockNum {

			log.Warn("Call HandleUnCandidateItem: the item stakingBlockNum no equal current candidate stakingBlockNum",
				"item stakingBlockNum", stakeItem.StakingBlockNum, "candidate stakingBlockNum", can.StakingBlockNum)

			if err := sk.db.DelUnStakeRecordStore(blockHash, epoch, uint64(index)); nil != err {
				log.Error("Failed to HandleUnCandidateItem: The Item is invilad, cause the stakingBlockNum is less "+
					"than stakingBlockNum of curr candidate, Delete unStakeRecord failed",
					"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err.Error())
				return err
			}

			continue
		}

		if err := sk.handleUnStake(state, blockNumber, blockHash, epoch, canAddr, can); nil != err {
			return err
		}

		if err := sk.db.DelUnStakeRecordStore(blockHash, epoch, uint64(index)); nil != err {
			log.Error("Failed to HandleUnCandidateItem: Delete unStakeRecord failed",
				"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err.Error())
			return err
		}

		filterAddr[canAddr] = struct{}{}
	}

	if err := sk.db.DelUnStakeCountStore(blockHash, epoch); nil != err {
		log.Error("Failed to HandleUnCandidateItem: Delete unstakeCount failed",
			"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err.Error())
		return err
	}

	return nil
}

func (sk *StakingL2Plugin) handleUnStake(state xcom.StateDB, blockNumber uint64, blockHash common.Hash, epoch uint64,
	addr common.NodeAddress, can *stakingL2.Candidate) error {

	log.Debug("Call handleUnStake", "blockNumber", blockNumber, "blockHash", blockHash.Hex(),
		"epoch", epoch, "nodeId", can.NodeId.String())

	lazyCalcL2StakeAmount(epoch, can.CandidateMutable)

	refundSharesFn := func(balance *big.Int) *big.Int {
		if balance.Cmp(common.Big0) > 0 {
			state.AddBalance(can.StakingAddress, balance)
			state.SubBalance(vm.StakingContractAddr, balance)
			return new(big.Int).SetInt64(0)
		}
		return balance
	}

	can.PendingShares = refundSharesFn(can.PendingShares)
	can.LockedShares = refundSharesFn(can.LockedShares)

	if err := sk.db.DelCandidateStore(blockHash, addr); nil != err {
		log.Error("Failed to HandleUnCandidateItem: Delete candidate info failed",
			"blockNumber", blockNumber, "blockHash", blockHash.Hex(),
			"nodeId", can.NodeId.String(), "err", err.Error())
		return err
	}

	if can.IsOperator == false {
		if err := sk.db.SubCommitteeCount(blockHash, 1); err != nil {
			log.Error("Failed to HandleUnCandidateItem: Subtraction unStakeRecord failed",
				"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err.Error())
			return err
		}
	}

	return nil
}

func (sk *StakingL2Plugin) GetOperatorList(blockHash common.Hash) (stakingL2.CandidateQueue, error) {

	iter := sk.db.IteratorOperatorsStore(blockHash, 0)
	if err := iter.Error(); nil != err {
		return nil, err
	}
	defer iter.Release()

	candidates := make(stakingL2.CandidateQueue, 0)

	for iter.Valid(); iter.Next(); {
		data := iter.Value()
		candidate := new(stakingL2.Candidate)
		if err := rlp.DecodeBytes(data, candidate); err != nil {
			return nil, err
		}
		candidates = append(candidates, candidate)
	}

	return candidates, nil
}

func (sk *StakingL2Plugin) GetCommitteeList(blockHash common.Hash) (stakingL2.CandidateQueue, error) {

	iter := sk.db.IteratorCommitteeStore(blockHash, 0)
	if err := iter.Error(); nil != err {
		return nil, err
	}
	defer iter.Release()

	candidates := make(stakingL2.CandidateQueue, 0)

	for iter.Valid(); iter.Next(); {
		data := iter.Value()
		candidate := new(stakingL2.Candidate)
		if err := rlp.DecodeBytes(data, candidate); err != nil {
			return nil, err
		}
		candidates = append(candidates, candidate)
	}

	return candidates, nil
}

func (sk *StakingL2Plugin) GetCandidateList(blockHash common.Hash, blockNumber uint64) (stakingL2.CandidateQueue, error) {

	epoch := xutil.CalculateEpoch(blockNumber)

	iter := sk.db.IteratorCandidateBase(blockHash, 0)
	if err := iter.Error(); nil != err {
		return nil, err
	}
	defer iter.Release()

	candidates := make(stakingL2.CandidateQueue, 0)

	for iter.Valid(); iter.Next(); {

		data := iter.Value()
		canBase := new(stakingL2.CandidateBase)
		if err := rlp.DecodeBytes(data, canBase); err != nil {
			return nil, err
		}
		nodeAddress, err := xutil.NodeId2Addr(canBase.NodeId)
		if err != nil {
			return nil, err
		}
		canMutable, err := sk.db.GetCanMutableStore(blockHash, nodeAddress)
		if nil != err {
			return nil, err
		}

		candidate := &stakingL2.Candidate{
			CandidateBase:    canBase,
			CandidateMutable: canMutable,
		}

		lazyCalcL2StakeAmount(epoch, candidate.CandidateMutable)
		candidates = append(candidates, candidate)
	}

	return candidates, nil
}

func lazyCalcL2StakeAmount(epoch uint64, can *stakingL2.CandidateMutable) {
	if can.IsEmpty() {
		return
	}

	changeAmountEpoch := can.StakingEpoch

	sub := epoch - uint64(changeAmountEpoch)

	log.Debug("lazyCalcL2StakeAmount before", "current epoch", epoch, "canMutable", can)

	// If it is during the same hesitation period, short circuit
	if sub < xcom.HesitateRatio() {
		return
	}

	if can.PendingShares.Cmp(common.Big0) > 0 {
		can.LockedShares = new(big.Int).Add(can.LockedShares, can.PendingShares)
		can.PendingShares = new(big.Int).SetInt64(0)
	}

	log.Debug("lazyCalcL2StakeAmount end", "current epoch", epoch, "canMutable", can)

}

func (sk *StakingL2Plugin) addUnStakeRecord(state xcom.StateDB, blockNumber uint64, blockHash common.Hash, epoch uint64,
	nodeId discover.NodeID, canAddr common.NodeAddress, stakingBlockNum uint64) error {

	refundEpoch := xutil.CalculateEpoch(blockNumber) + 0
	log.Debug("Call addUnStakeRecord, AddUnStakeRecordStore start", "current blockNumber", blockNumber,
		"unStakeFreeze Epoch", 0, "unStake record refund Epoch", refundEpoch, "nodeId", nodeId.String())

	if err := sk.db.AddUnStakeRecordStore(blockHash, refundEpoch, canAddr, stakingBlockNum); nil != err {
		return err
	}
	return nil
}

func (sk *StakingL2Plugin) CheckStakeThresholdL2(amount *big.Int) bool {
	return amount.Cmp(StakeThresholdL2) >= 0
}

func buildMarshalAbleCandidate(can *stakingL2.Candidate) *stakingL2.MarshalAbleCandidate {
	return &stakingL2.MarshalAbleCandidate{
		NodeId:      can.NodeId,
		Name:        can.Name,
		Status:      can.Status,
		Version:     can.Version,
		ElectronURI: can.ElectronURI,
		RPCURI:      can.RPCURI,
		P2PURI:      can.P2PURI,
		IsOperator:  can.IsOperator,

		StakingAddress:  can.StakingAddress,
		BenefitAddress:  can.BenefitAddress,
		StakingEpoch:    can.StakingEpoch,
		StakingBlockNum: can.StakingBlockNum,
		StakingTxIndex:  can.StakingTxIndex,

		Shares:        (*hexutil.Big)(can.Shares),
		LockedShares:  (*hexutil.Big)(can.LockedShares),
		PendingShares: (*hexutil.Big)(can.PendingShares),

		BlsPubKey: can.BlsPubKey,
		Detail:    can.Detail,
	}
}
