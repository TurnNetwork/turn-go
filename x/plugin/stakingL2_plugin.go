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
	"github.com/bubblenet/bubble/params"
	"math/big"
	"sync"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/ethdb"
	"github.com/bubblenet/bubble/event"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/x/gov"
	"github.com/bubblenet/bubble/x/stakingL2"
	"github.com/bubblenet/bubble/x/xcom"
	"github.com/bubblenet/bubble/x/xutil"
)

var (
	StakeThresholdL2 = new(big.Int).Mul(big.NewInt(params.BUB), big.NewInt(200))
)

type StakingL2Plugin struct {
	db            *stakingL2.StakingDB
	eventMux      *event.TypeMux
	chainReaderDB ethdb.KeyValueReader
	chainWriterDB ethdb.KeyValueWriter
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

func (sk *StakingL2Plugin) SetEventMux(eventMux *event.TypeMux) {
	sk.eventMux = eventMux
}

func (sk *StakingL2Plugin) SetChainDB(reader ethdb.KeyValueReader, writer ethdb.KeyValueWriter) {
	sk.chainReaderDB = reader
	sk.chainWriterDB = writer
}

func (sk *StakingL2Plugin) BeginBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	//blockNumber := header.Number.Uint64()
	//if xutil.IsEndOfConsensus(blockNumber) {
	//	// Store the list of consensus nodes for the next round in the DB in the last block of the consensus round.
	//	// Used to record historical consensus round node information.
	//	// 1. Simplify the consensus node information
	//	// 2. Calculate the identification ID
	//	// 3. Compute the hash of the simplified list of node information
	//	// 4. Replace the value of header.extra[0:32] with the Hash value.
	//	// 5. Form a list of identification IDs in the order of block generation and write them into the DB
	//	next, err := sk.getNextValList(blockHash, blockNumber, QueryStartNotIrr)
	//	if err != nil {
	//		log.Error("Failed to Query Next validators on StakingL2Plugin Begin When end of consensus",
	//			"blockNumber", blockNumber, "blockHash", blockHash.TerminalString(), "err", err)
	//		return err
	//	}
	//	historyValidatorList := make(stakingL2.HistoryValidatorList, len(next.Arr))
	//	historyValidatorIDList := make(stakingL2.HistoryValidatorIDList, len(next.Arr))
	//	for i := 0; i < len(next.Arr); i++ {
	//		hv := &stakingL2.HistoryValidator{
	//			NodeId:    next.Arr[i].NodeId,
	//			BlsPubKey: next.Arr[i].BlsPubKey,
	//		}
	//		id := hv.ID()
	//		historyValidatorList[i] = hv
	//		historyValidatorIDList[i] = id
	//		if err := sk.writeHistoryValidator(id, hv, blockHash, header, state); err != nil {
	//			return err
	//		}
	//	}
	//	if err := sk.writeHistoryValidatorIDList(historyValidatorIDList, next.Start, blockHash, header, state); err != nil {
	//		return err
	//	}
	//	listHash, err := historyValidatorList.Hash()
	//	if err != nil {
	//		log.Error("Failed to calculate Hash for consensus round node list", "blockNumber", blockNumber, "blockHash", blockHash.TerminalString(), "err", err)
	//		return err
	//	}
	//	// The outgoing block node writes to extra.
	//	// Non-outgoing block nodes validate extra.
	//	if xutil.IsWorker(header.Extra) {
	//		// The hash value will be signed by the node.
	//		// will also be counted in the block Hash.
	//		copy(header.Extra[:32], listHash.Bytes())
	//	} else {
	//		if !bytes.Equal(header.Extra[:32], listHash.Bytes()) {
	//			return errors.New("historical validator list Hash is not the same")
	//		}
	//	}
	//	log.Debug("Historical consensus node information written successfully", "blockNumber", blockNumber, "blockHash", blockHash.TerminalString(), "listHash", listHash.Hex())
	//}
	return nil
}

func (sk *StakingL2Plugin) EndBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {

	epoch := xutil.CalculateEpoch(header.Number.Uint64())

	if xutil.IsEndOfEpoch(header.Number.Uint64()) {

		// handle UnStaking Item
		err := sk.HandleUnCandidateItem(state, header.Number.Uint64(), blockHash, epoch)
		if nil != err {
			log.Error("Failed to call HandleUnCandidateItem on StakingL2Plugin EndBlock",
				"blockNumber", header.Number.Uint64(), "blockHash", blockHash.Hex(), "err", err)
			return err
		}

	}

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

func (sk *StakingL2Plugin) GetCandidateCompactInfo(blockHash common.Hash, blockNumber uint64, addr common.NodeAddress) (*stakingL2.CandidateHex, error) {
	can, err := sk.GetCandidateInfo(blockHash, addr)
	if nil != err {
		return nil, err
	}

	epoch := xutil.CalculateEpoch(blockNumber)
	lazyCalcL2StakeAmount(epoch, can.CandidateMutable)
	canHex := buildL2CanHex(can)
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
	can.ReleasedHes = amount
	can.StakingEpoch = uint32(xutil.CalculateEpoch(blockNumber.Uint64()))

	if err := sk.db.SetCandidateStore(blockHash, addr, can); nil != err {
		log.Error("Failed to CreateCandidate on StakingL2Plugin: Store Candidate info is failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
		return err
	}

	if can.IsOperator {
		if err := sk.db.SetOperatorStore(blockHash, addr, can); nil != err {
			log.Error("Failed to CreateCandidate on StakingL2Plugin: Store Operator info is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
			return err
		}
	} else {
		if err := sk.db.SetCommitteeStore(blockHash, addr, can); nil != err {
			log.Error("Failed to CreateCandidate on StakingL2Plugin: Store Committee info is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
			return err
		}
	}

	if err := sk.db.SetCanPowerStore(blockHash, addr, can); nil != err {
		log.Error("Failed to CreateCandidate on StakingL2Plugin: Store Candidate power is failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
		return err
	}

	// add the account staking Reference Count
	if err := sk.db.AddAccountStakeRc(blockHash, can.StakingAddress); nil != err {
		log.Error("Failed to CreateCandidate on StakingL2Plugin: Store Staking Account Reference Count (add) is failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "NodeID", can.NodeId.String(),
			"staking Account", can.StakingAddress.String(), "err", err)
		return err
	}

	return nil
}

func (sk *StakingL2Plugin) EditCandidate(blockHash common.Hash, blockNumber *big.Int, canAddr common.NodeAddress, can *stakingL2.Candidate) error {
	if err := sk.db.SetCanBaseStore(blockHash, canAddr, can.CandidateBase); nil != err {
		log.Error("Failed to EditCandidate on StakingL2Plugin: Store CandidateBase info is failed", "nodeId", can.NodeId.String(),
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "err", err)
		return err
	}
	if err := sk.db.SetCanMutableStore(blockHash, canAddr, can.CandidateMutable); nil != err {
		log.Error("Failed to EditCandidate on StakingL2Plugin: Store CandidateMutable info is failed", "nodeId", can.NodeId.String(),
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "err", err)
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
//	can.ReleasedHes = new(big.Int).Add(can.ReleasedHes, amount)
//
//	if err := sk.db.DelCanPowerStore(blockHash, can); nil != err {
//		log.Error("Failed to IncreaseStaking on StakingL2Plugin: Delete Candidate old power is failed",
//			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(),
//			"nodeId", can.NodeId.String(), "err", err)
//		return err
//	}
//
//	can.StakingEpoch = uint32(epoch)
//	can.AddShares(amount)
//
//	if err := sk.db.SetCanPowerStore(blockHash, canAddr, can); nil != err {
//		log.Error("Failed to IncreaseStaking on StakingL2Plugin: Store Candidate new power is failed",
//			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(),
//			"nodeId", can.NodeId.String(), "err", err)
//		return err
//	}
//
//	if err := sk.db.SetCanMutableStore(blockHash, canAddr, can.CandidateMutable); nil != err {
//		log.Error("Failed to IncreaseStaking on StakingL2Plugin: Store CandidateMutable info is failed",
//			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(),
//			"nodeId", can.NodeId.String(), "err", err)
//		return err
//	}
//
//	return nil
//}

func (sk *StakingL2Plugin) WithdrewStaking(state xcom.StateDB, blockHash common.Hash, blockNumber *big.Int,
	canAddr common.NodeAddress, can *stakingL2.Candidate) error {

	epoch := xutil.CalculateEpoch(blockNumber.Uint64())

	lazyCalcL2StakeAmount(epoch, can.CandidateMutable)

	if err := sk.db.DelCanPowerStore(blockHash, can); nil != err {
		log.Error("Failed to WithdrewStaking on StakingL2Plugin: Delete Candidate old power is failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
		return err
	}

	if err := sk.withdrewStakeAmount(state, blockHash, blockNumber.Uint64(), epoch, canAddr, can); nil != err {
		return err
	}

	can.StakingEpoch = uint32(epoch)

	if can.Released.Cmp(common.Big0) > 0 {
		if err := sk.db.SetCanMutableStore(blockHash, canAddr, can.CandidateMutable); nil != err {
			log.Error("Failed to WithdrewStaking on StakingL2Plugin: Store CandidateMutable info is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
			return err
		}
	} else {
		// 删除选举队列节点信息
		if can.IsOperator {
			if err := sk.db.DelOperatorStore(blockHash, canAddr); err != nil {
				log.Error("Failed to WithdrewStaking on StakingL2Plugin: Delete Operator info is failed",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
				return err
			}
		} else {
			if err := sk.db.DelCommitteeStore(blockHash, canAddr); err != nil {
				log.Error("Failed to WithdrewStaking on StakingL2Plugin: Delete Operator info is failed",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
				return err
			}
		}
		// 删除节点信息
		if err := sk.db.DelCandidateStore(blockHash, canAddr); nil != err {
			log.Error("Failed to WithdrewStaking on StakingL2Plugin: Delete Candidate info is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
			return err
		}
	}

	// sub the account staking Reference Count
	if err := sk.db.SubAccountStakeRc(blockHash, can.StakingAddress); nil != err {
		log.Error("Failed to WithdrewStaking on StakingL2Plugin: Store Staking Account Reference Count (sub) is failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(),
			"staking Account", can.StakingAddress.String(), "err", err)
		return err
	}

	return nil
}

func (sk *StakingL2Plugin) withdrewStakeAmount(state xcom.StateDB, blockHash common.Hash, blockNumber, epoch uint64,
	canAddr common.NodeAddress, can *stakingL2.Candidate) error {

	// Direct return of money during the hesitation period
	// Return according to the way of coming
	if can.ReleasedHes.Cmp(common.Big0) > 0 {
		state.AddBalance(can.StakingAddress, can.ReleasedHes)
		state.SubBalance(vm.StakingContractAddr, can.ReleasedHes)
		can.ReleasedHes = new(big.Int).SetInt64(0)
	}

	if can.Released.Cmp(common.Big0) > 0 {
		if err := sk.addUnStakeItem(state, blockNumber, blockHash, epoch, can.NodeId, canAddr, can.StakingBlockNum); nil != err {
			log.Error("Failed to WithdrewStaking on StakingL2Plugin: Add UnStakeItemStore failed",
				"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
			return err
		}
	}

	can.CleanShares()
	can.Status |= stakingL2.Invalided | stakingL2.Withdrew

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

		stakeItem, err := sk.db.GetUnStakeItemStore(blockHash, epoch, uint64(index))
		if nil != err {
			log.Error("Failed to HandleUnCandidateItem: Query the unStakeItem node addr is failed",
				"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err)
			return err
		}

		canAddr := stakeItem.NodeAddress

		//log.Debug("Call HandleUnCandidateItem: the candidate Addr",
		//	"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "addr", canAddr.Hex())

		if _, ok := filterAddr[canAddr]; ok {
			if err := sk.db.DelUnStakeItemStore(blockHash, epoch, uint64(index)); nil != err {
				log.Error("Failed to HandleUnCandidateItem: Delete already handle unstakeItem failed",
					"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err)
				return err
			}
			continue
		}

		can, err := sk.db.GetCandidateStore(blockHash, canAddr)
		if snapshotdb.NonDbNotFoundErr(err) {
			log.Error("Failed to HandleUnCandidateItem: Query candidate failed",
				"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "canAddr", canAddr.Hex(), "err", err)
			return err
		}

		// This should not be nil
		if snapshotdb.IsDbNotFoundErr(err) || can.IsEmpty() {

			if err := sk.db.DelUnStakeItemStore(blockHash, epoch, uint64(index)); nil != err {
				log.Error("Failed to HandleUnCandidateItem: Candidate is no exist, Delete unstakeItem failed",
					"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err)
				return err
			}

			continue
		}

		// if the item stakingBlockNum is not enough the stakingBlockNum of candidate info
		if stakeItem.StakingBlockNum != can.StakingBlockNum {

			log.Warn("Call HandleUnCandidateItem: the item stakingBlockNum no equal current candidate stakingBlockNum",
				"item stakingBlockNum", stakeItem.StakingBlockNum, "candidate stakingBlockNum", can.StakingBlockNum)

			if err := sk.db.DelUnStakeItemStore(blockHash, epoch, uint64(index)); nil != err {
				log.Error("Failed to HandleUnCandidateItem: The Item is invilad, cause the stakingBlockNum is less "+
					"than stakingBlockNum of curr candidate, Delete unstakeItem failed",
					"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err)
				return err
			}

			continue

		}

		// The state of the node needs to be restored
		if stakeItem.Recovery {
			// If the node is reported double-signed during the lock-up period，
			// Then you need to enter the double-signed lock-up period after the lock-up period expires and release the staking after the expiration
			// Otherwise, the state of the node is restored to the normal staking state
			if can.IsDuplicateSign() {

				// Because there is no need to release the staking when the zero-out block is locked, "SubAccountStakeRc" is not executed
				if err := sk.db.SubAccountStakeRc(blockHash, can.StakingAddress); nil != err {
					log.Error("Failed to HandleUnCandidateItem: Sub Account staking Reference Count is failed",
						"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
					return err
				}

				// Lock the node again and will release the staking
				if err := sk.addUnStakeItem(state, blockNumber, blockHash, epoch, can.NodeId, canAddr, can.StakingBlockNum); nil != err {
					log.Error("Failed to SlashCandidates on StakingL2Plugin: Add UnStakeItemStore failed",
						"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
					return err
				}
				can.CleanLowRatioStatus()
				if err := sk.db.SetCanMutableStore(blockHash, canAddr, can.CandidateMutable); nil != err {
					log.Error("Failed to HandleUnCandidateItem on StakingL2Plugin: Store CandidateMutable info is failed",
						"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
					return err
				}
				log.Debug("Call HandleUnCandidateItem: Node double sign", "blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(),
					"status", can.Status, "shares", can.Shares)
			} else {

				can.SetValided()
				if err := sk.db.SetCanPowerStore(blockHash, canAddr, can); nil != err {
					log.Error("Failed to HandleUnCandidateItem on StakingL2Plugin: Store Candidate power is failed",
						"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
					return err
				}
				if err := sk.db.SetCanMutableStore(blockHash, canAddr, can.CandidateMutable); nil != err {
					log.Error("Failed to HandleUnCandidateItem on StakingL2Plugin: Store CandidateMutable info is failed",
						"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(), "err", err)
					return err
				}
				log.Debug("Call HandleUnCandidateItem: Node state recovery", "blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", can.NodeId.String(),
					"status", can.Status, "shares", can.Shares)
			}
		} else {
			// Second handle balabala ...
			if err := sk.handleUnStake(state, blockNumber, blockHash, epoch, canAddr, can); nil != err {
				return err
			}
		}

		if err := sk.db.DelUnStakeItemStore(blockHash, epoch, uint64(index)); nil != err {
			log.Error("Failed to HandleUnCandidateItem: Delete unstakeItem failed",
				"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err)
			return err
		}

		filterAddr[canAddr] = struct{}{}
	}

	if err := sk.db.DelUnStakeCountStore(blockHash, epoch); nil != err {
		log.Error("Failed to HandleUnCandidateItem: Delete unstakeCount failed",
			"blockNUmber", blockNumber, "blockHash", blockHash.Hex(), "err", err)
		return err
	}

	return nil
}

func (sk *StakingL2Plugin) handleUnStake(state xcom.StateDB, blockNumber uint64, blockHash common.Hash, epoch uint64,
	addr common.NodeAddress, can *stakingL2.Candidate) error {

	log.Debug("Call handleUnStake", "blockNumber", blockNumber, "blockHash", blockHash.Hex(),
		"epoch", epoch, "nodeId", can.NodeId.String())

	lazyCalcL2StakeAmount(epoch, can.CandidateMutable)

	refundReleaseFn := func(balance *big.Int) *big.Int {
		if balance.Cmp(common.Big0) > 0 {
			state.AddBalance(can.StakingAddress, balance)
			state.SubBalance(vm.StakingContractAddr, balance)
			return new(big.Int).SetInt64(0)
		}
		return balance
	}

	can.ReleasedHes = refundReleaseFn(can.ReleasedHes)
	can.Released = refundReleaseFn(can.Released)

	if err := sk.db.DelCandidateStore(blockHash, addr); nil != err {
		log.Error("Failed to HandleUnCandidateItem: Delete candidate info failed",
			"blockNumber", blockNumber, "blockHash", blockHash.Hex(),
			"nodeId", can.NodeId.String(), "err", err)
		return err
	}

	return nil
}

func (sk *StakingL2Plugin) GetOperatorList(blockHash common.Hash, blockNumber uint64) (stakingL2.CandidateQueue, error) {

	epoch := xutil.CalculateEpoch(blockNumber)

	iter := sk.db.IteratorOperatorsStore(blockHash, 0)
	if err := iter.Error(); nil != err {
		return nil, err
	}
	defer iter.Release()

	candidates := make(stakingL2.CandidateQueue, 0)

	for iter.Valid(); iter.Next(); {

		addrSuffix := iter.Value()
		candidate, err := sk.db.GetCandidateStoreWithSuffix(blockHash, addrSuffix)
		if nil != err {
			return nil, err
		}

		lazyCalcL2StakeAmount(epoch, candidate.CandidateMutable)
		candidates = append(candidates, candidate)
	}

	return candidates, nil
}

func (sk *StakingL2Plugin) GetCandidateList(blockHash common.Hash, blockNumber uint64) (stakingL2.CandidateQueue, error) {

	epoch := xutil.CalculateEpoch(blockNumber)

	iter := sk.db.IteratorCandidatePowerByBlockHash(blockHash, 0)
	if err := iter.Error(); nil != err {
		return nil, err
	}
	defer iter.Release()

	candidates := make(stakingL2.CandidateQueue, 0)

	for iter.Valid(); iter.Next(); {

		addrSuffix := iter.Value()
		candidate, err := sk.db.GetCandidateStoreWithSuffix(blockHash, addrSuffix)
		if nil != err {
			return nil, err
		}

		lazyCalcL2StakeAmount(epoch, candidate.CandidateMutable)
		candidates = append(candidates, candidate)
	}

	return candidates, nil
}

func (sk *StakingL2Plugin) GetCanBaseList(blockHash common.Hash, blockNumber uint64) (stakingL2.CandidateBaseQueue, error) {

	iter := sk.db.IteratorCandidatePowerByBlockHash(blockHash, 0)
	if err := iter.Error(); nil != err {
		return nil, err
	}
	defer iter.Release()

	queue := make(stakingL2.CandidateBaseQueue, 0)

	for iter.Valid(); iter.Next(); {

		addrSuffix := iter.Value()
		can, err := sk.db.GetCanBaseStoreWithSuffix(blockHash, addrSuffix)
		if nil != err {
			return nil, err
		}
		queue = append(queue, can)
	}

	return queue, nil
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

	if can.ReleasedHes.Cmp(common.Big0) > 0 {
		can.Released = new(big.Int).Add(can.Released, can.ReleasedHes)
		can.ReleasedHes = new(big.Int).SetInt64(0)
	}

	log.Debug("lazyCalcL2StakeAmount end", "current epoch", epoch, "canMutable", can)

}

func (sk *StakingL2Plugin) addUnStakeItem(state xcom.StateDB, blockNumber uint64, blockHash common.Hash, epoch uint64,
	nodeId discover.NodeID, canAddr common.NodeAddress, stakingBlockNum uint64) error {

	endVoteNum, err := gov.GetMaxEndVotingBlock(nodeId, blockHash, state)
	if nil != err {
		return err
	}
	var refundEpoch, maxEndVoteEpoch, targetEpoch uint64
	if endVoteNum != 0 {
		maxEndVoteEpoch = xutil.CalculateEpoch(endVoteNum)
	}

	duration, err := gov.GovernUnStakeFreezeDuration(blockNumber, blockHash)
	if nil != err {
		return err
	}

	refundEpoch = xutil.CalculateEpoch(blockNumber) + duration

	if maxEndVoteEpoch <= refundEpoch {
		targetEpoch = refundEpoch
	} else {
		targetEpoch = maxEndVoteEpoch
	}

	log.Debug("Call addUnStakeItem, AddUnStakeItemStore start", "current blockNumber", blockNumber,
		"govenance max end vote blokNumber", endVoteNum, "unStakeFreeze Epoch", refundEpoch,
		"govenance max end vote epoch", maxEndVoteEpoch, "unstake item target Epoch", targetEpoch,
		"nodeId", nodeId.String())

	if err := sk.db.AddUnStakeItemStore(blockHash, targetEpoch, canAddr, stakingBlockNum, false); nil != err {
		return err
	}
	return nil
}

func (sk *StakingL2Plugin) addRecoveryUnStakeItem(blockNumber uint64, blockHash common.Hash, nodeId discover.NodeID,
	canAddr common.NodeAddress, stakingBlockNum uint64) error {

	duration, err := gov.GovernZeroProduceFreezeDuration(blockNumber, blockHash)
	if nil != err {
		return err
	}

	targetEpoch := xutil.CalculateEpoch(blockNumber) + duration

	log.Debug("Call addRecoveryUnStakeItem, AddUnStakeItemStore start", "current blockNumber", blockNumber,
		"duration", duration, "unstake item target Epoch", targetEpoch,
		"nodeId", nodeId.String())

	if err := sk.db.AddUnStakeItemStore(blockHash, targetEpoch, canAddr, stakingBlockNum, true); nil != err {
		return err
	}
	return nil
}

func (sk *StakingL2Plugin) CheckStakeThresholdL2(amount *big.Int) bool {
	return amount.Cmp(StakeThresholdL2) >= 0
}

func buildL2CanHex(can *stakingL2.Candidate) *stakingL2.CandidateHex {
	return &stakingL2.CandidateHex{
		NodeId:          can.NodeId,
		BlsPubKey:       can.BlsPubKey,
		StakingAddress:  can.StakingAddress,
		BenefitAddress:  can.BenefitAddress,
		StakingTxIndex:  can.StakingTxIndex,
		ProgramVersion:  can.ProgramVersion,
		Status:          can.Status,
		StakingEpoch:    can.StakingEpoch,
		StakingBlockNum: can.StakingBlockNum,
		Shares:          (*hexutil.Big)(can.Shares),
		Released:        (*hexutil.Big)(can.Released),
		ReleasedHes:     (*hexutil.Big)(can.ReleasedHes),
		Description:     can.Description,
		IsOperator:      can.IsOperator,
	}
}
