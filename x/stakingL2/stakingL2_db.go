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

package stakingL2

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/rlp"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type StakingDB struct {
	db snapshotdb.DB
}

func NewStakingDB() *StakingDB {
	return &StakingDB{
		db: snapshotdb.Instance(),
	}
}

func (db *StakingDB) GetDB() snapshotdb.DB {
	return db.db
}

func (db *StakingDB) get(blockHash common.Hash, key []byte) ([]byte, error) {
	return db.db.Get(blockHash, key)
}

func (db *StakingDB) put(blockHash common.Hash, key, value []byte) error {
	return db.db.Put(blockHash, key, value)
}

func (db *StakingDB) del(blockHash common.Hash, key []byte) error {
	return db.db.Del(blockHash, key)
}

func (db *StakingDB) ranking(blockHash common.Hash, prefix []byte, ranges int) iterator.Iterator {
	return db.db.Ranking(blockHash, prefix, ranges)
}

// about Candidate ...

func (db *StakingDB) GetCandidateStore(blockHash common.Hash, addr common.NodeAddress) (*Candidate, error) {
	base, err := db.GetCanBaseStore(blockHash, addr)
	if nil != err {
		return nil, err
	}
	mutable, err := db.GetCanMutableStore(blockHash, addr)
	if nil != err {
		return nil, err
	}

	can := &Candidate{}
	can.CandidateBase = base
	can.CandidateMutable = mutable
	return can, nil
}

func (db *StakingDB) GetCandidateStoreWithSuffix(blockHash common.Hash, suffix []byte) (*Candidate, error) {
	base, err := db.GetCanBaseStoreWithSuffix(blockHash, suffix)
	if nil != err {
		return nil, err
	}
	mutable, err := db.GetCanMutableStoreWithSuffix(blockHash, suffix)
	if nil != err {
		return nil, err
	}

	can := &Candidate{}
	can.CandidateBase = base
	can.CandidateMutable = mutable
	return can, nil
}

func (db *StakingDB) SetCandidateStore(blockHash common.Hash, addr common.NodeAddress, can *Candidate) error {

	if err := db.SetCanBaseStore(blockHash, addr, can.CandidateBase); nil != err {
		return err
	}
	if err := db.SetCanMutableStore(blockHash, addr, can.CandidateMutable); nil != err {
		return err
	}
	return nil
}

func (db *StakingDB) DelCandidateStore(blockHash common.Hash, addr common.NodeAddress) error {
	if err := db.DelCanBaseStore(blockHash, addr); nil != err {
		return err
	}
	if err := db.DelCanMutableStore(blockHash, addr); nil != err {
		return err
	}
	return nil
}

// about canbase ...

func (db *StakingDB) GetCanBaseStore(blockHash common.Hash, addr common.NodeAddress) (*CandidateBase, error) {

	key := CanBaseKeyByAddr(addr)

	canByte, err := db.get(blockHash, key)

	if nil != err {
		return nil, err
	}

	var can CandidateBase
	if err := rlp.DecodeBytes(canByte, &can); nil != err {
		return nil, err
	}

	return &can, nil
}

func (db *StakingDB) GetCanBaseStoreWithSuffix(blockHash common.Hash, suffix []byte) (*CandidateBase, error) {
	key := CanBaseKeyBySuffix(suffix)

	canByte, err := db.get(blockHash, key)

	if nil != err {
		return nil, err
	}
	var can CandidateBase

	if err := rlp.DecodeBytes(canByte, &can); nil != err {
		return nil, err
	}
	return &can, nil
}

func (db *StakingDB) SetCanBaseStore(blockHash common.Hash, addr common.NodeAddress, can *CandidateBase) error {

	key := CanBaseKeyByAddr(addr)

	if val, err := rlp.EncodeToBytes(can); nil != err {
		return err
	} else {

		return db.put(blockHash, key, val)
	}
}

func (db *StakingDB) DelCanBaseStore(blockHash common.Hash, addr common.NodeAddress) error {
	key := CanBaseKeyByAddr(addr)
	return db.del(blockHash, key)
}

// about canmutable ...

func (db *StakingDB) GetCanMutableStore(blockHash common.Hash, addr common.NodeAddress) (*CandidateMutable, error) {

	key := CanMutableKeyByAddr(addr)

	canByte, err := db.get(blockHash, key)

	if nil != err {
		return nil, err
	}

	var can CandidateMutable
	if err := rlp.DecodeBytes(canByte, &can); nil != err {
		return nil, err
	}

	return &can, nil
}

func (db *StakingDB) GetCanMutableStoreWithSuffix(blockHash common.Hash, suffix []byte) (*CandidateMutable, error) {
	key := CanMutableKeyBySuffix(suffix)

	canByte, err := db.get(blockHash, key)

	if nil != err {
		return nil, err
	}
	var can CandidateMutable

	if err := rlp.DecodeBytes(canByte, &can); nil != err {
		return nil, err
	}
	return &can, nil
}

func (db *StakingDB) SetCanMutableStore(blockHash common.Hash, addr common.NodeAddress, can *CandidateMutable) error {

	key := CanMutableKeyByAddr(addr)

	if val, err := rlp.EncodeToBytes(can); nil != err {
		return err
	} else {

		return db.put(blockHash, key, val)
	}
}

func (db *StakingDB) DelCanMutableStore(blockHash common.Hash, addr common.NodeAddress) error {
	key := CanMutableKeyByAddr(addr)
	return db.del(blockHash, key)
}

func (db *StakingDB) GetOperatorStore(blockHash common.Hash, addr common.NodeAddress) (*Candidate, error) {
	data, err := db.db.Get(blockHash, OperatorKeyByAddr(addr))
	if nil != err {
		return nil, err
	}

	var can Candidate
	if err := rlp.DecodeBytes(data, &can); nil != err {
		return nil, err
	}

	return &can, nil
}

func (db *StakingDB) IteratorOperatorsStore(blockHash common.Hash, ranges int) iterator.Iterator {
	return db.ranking(blockHash, OperatorKeyPrefix, ranges)
}

func (db *StakingDB) SetOperatorStore(blockHash common.Hash, addr common.NodeAddress, can *Candidate) error {

	if data, err := rlp.EncodeToBytes(can); nil != err {
		return err
	} else {
		return db.put(blockHash, OperatorKeyByAddr(addr), data)
	}
}

func (db *StakingDB) DelOperatorStore(blockHash common.Hash, addr common.NodeAddress) error {
	return db.del(blockHash, OperatorKeyByAddr(addr))
}

func (db *StakingDB) GetCommitteeStore(blockHash common.Hash, addr common.NodeAddress) (*Candidate, error) {
	data, err := db.db.Get(blockHash, CommitteeKeyByAddr(addr))
	if nil != err {
		return nil, err
	}

	var can Candidate
	if err := rlp.DecodeBytes(data, &can); nil != err {
		return nil, err
	}

	return &can, nil
}

func (db *StakingDB) IteratorCommitteeStore(blockHash common.Hash, ranges int) iterator.Iterator {
	return db.ranking(blockHash, CommitteeKeyPrefix, ranges)
}

func (db *StakingDB) SetCommitteeStore(blockHash common.Hash, addr common.NodeAddress, can *Candidate) error {

	if data, err := rlp.EncodeToBytes(can); nil != err {
		return err
	} else {
		return db.put(blockHash, CommitteeKeyByAddr(addr), data)
	}
}

func (db *StakingDB) DelCommitteeStore(blockHash common.Hash, addr common.NodeAddress) error {
	return db.del(blockHash, CommitteeKeyByAddr(addr))
}

func (db *StakingDB) GetCommitteeCount(blockHash common.Hash) (uint32, error) {
	count, err := db.db.Get(blockHash, CommitteeCountKey)
	if err == snapshotdb.ErrNotFound {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}

	return common.BytesToUint32(count), nil
}

func (db *StakingDB) SetCommitteeCount(blockHash common.Hash, count uint32) error {
	return db.db.Put(blockHash, CommitteeCountKey, common.Uint32ToBytes(count))
}

func (db *StakingDB) AddCommitteeCount(blockHash common.Hash, number uint32) error {
	count, err := db.GetCommitteeCount(blockHash)
	if err != nil {
		return err
	}
	if err := db.SetCommitteeCount(blockHash, count+number); err != nil {
		return err
	}

	return nil
}

func (db *StakingDB) SubCommitteeCount(blockHash common.Hash, number uint32) error {
	count, err := db.GetCommitteeCount(blockHash)
	if err != nil {
		return err
	}
	if err := db.SetCommitteeCount(blockHash, count-number); err != nil {
		return err
	}

	return nil
}

func (db *StakingDB) GetUsedCommitteeCount(blockHash common.Hash) (uint32, error) {
	count, err := db.db.Get(blockHash, UsedCommitteeCountKey)
	if err == snapshotdb.ErrNotFound {
		return 0, nil
	}
	if nil != err {
		return 0, err
	}

	return common.BytesToUint32(count), nil
}

func (db *StakingDB) SetUsedCommitteeCount(blockHash common.Hash, count uint32) error {
	return db.db.Put(blockHash, CommitteeCountKey, common.Uint32ToBytes(count))
}

func (db *StakingDB) AddUsedCommitteeCount(blockHash common.Hash, number uint32) error {
	count, err := db.GetUsedCommitteeCount(blockHash)
	if err != nil {
		return err
	}
	if err := db.SetUsedCommitteeCount(blockHash, count+number); err != nil {
		return err
	}

	return nil
}

func (db *StakingDB) SubUsedCommitteeCount(blockHash common.Hash, number uint32) error {
	count, err := db.GetUsedCommitteeCount(blockHash)
	if err != nil {
		return err
	}
	if err := db.SetUsedCommitteeCount(blockHash, count-number); err != nil {
		return err
	}

	return nil
}

// about UnStakeRecord ...

func (db *StakingDB) AddUnStakeRecordStore(blockHash common.Hash, epoch uint64, canAddr common.NodeAddress, stakeBlockNumber uint64) error {

	count_key := GetUnStakeCountKey(epoch)

	val, err := db.get(blockHash, count_key)
	var v uint64
	switch {
	case snapshotdb.NonDbNotFoundErr(err):
		return err
	case nil == err && len(val) != 0:
		v = common.BytesToUint64(val)
	}

	v++

	if err := db.put(blockHash, count_key, common.Uint64ToBytes(v)); nil != err {
		return err
	}
	item_key := GetUnStakeRecordKey(epoch, v)

	unStakeRecord := &UnStakeRecord{
		NodeAddress:     canAddr,
		StakingBlockNum: stakeBlockNumber,
	}

	item, err := rlp.EncodeToBytes(unStakeRecord)
	if nil != err {
		return err
	}

	return db.put(blockHash, item_key, item)
}

func (db *StakingDB) GetUnStakeCountStore(blockHash common.Hash, epoch uint64) (uint64, error) {
	count_key := GetUnStakeCountKey(epoch)

	val, err := db.get(blockHash, count_key)
	if nil != err {
		return 0, err
	}
	return common.BytesToUint64(val), nil
}

func (db *StakingDB) GetUnStakeRecordStore(blockHash common.Hash, epoch, index uint64) (*UnStakeRecord, error) {
	item_key := GetUnStakeRecordKey(epoch, index)
	itemByte, err := db.get(blockHash, item_key)
	if nil != err {
		return nil, err
	}

	var unStakeRecord UnStakeRecord
	if err := rlp.DecodeBytes(itemByte, &unStakeRecord); nil != err {
		return nil, err
	}
	return &unStakeRecord, nil
}

func (db *StakingDB) DelUnStakeCountStore(blockHash common.Hash, epoch uint64) error {
	count_key := GetUnStakeCountKey(epoch)

	return db.del(blockHash, count_key)
}

func (db *StakingDB) DelUnStakeRecordStore(blockHash common.Hash, epoch, index uint64) error {
	item_key := GetUnStakeRecordKey(epoch, index)

	return db.del(blockHash, item_key)
}

func (db *StakingDB) IteratorCandidateBase(blockHash common.Hash, ranges int) iterator.Iterator {
	return db.ranking(blockHash, CanBaseKeyPrefix, ranges)
}

// about account staking reference count ...

func (db *StakingDB) AddAccountStakeRc(blockHash common.Hash, addr common.Address) error {
	key := GetAccountStakeRcKey(addr)
	val, err := db.get(blockHash, key)
	var v uint64
	switch {
	case snapshotdb.NonDbNotFoundErr(err):
		return err
	case nil == err && len(val) != 0:
		v = common.BytesToUint64(val)
	}

	v++

	return db.put(blockHash, key, common.Uint64ToBytes(v))
}

func (db *StakingDB) SubAccountStakeRc(blockHash common.Hash, addr common.Address) error {
	key := GetAccountStakeRcKey(addr)
	val, err := db.get(blockHash, key)
	var v uint64
	switch {
	case snapshotdb.NonDbNotFoundErr(err):
		return err
	case nil == err && len(val) != 0:
		v = common.BytesToUint64(val)
	}

	// Prevent large numbers from being directly called after the uint64 overflow
	if v == 0 {
		return nil
	}

	v--

	if v == 0 {

		return db.del(blockHash, key)
	} else {

		return db.put(blockHash, key, common.Uint64ToBytes(v))
	}
}