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
	"math/big"

	"github.com/bubblenet/bubble/x/xutil"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/math"
	"github.com/bubblenet/bubble/p2p/discover"
)

var (
	CanBaseKeyPrefix     = []byte("CanBaseL2")
	CanMutableKeyPrefix  = []byte("CanMutL2")
	OperatorKeyPrefix    = []byte("OperatorL2")
	CommitteeKeyPrefix   = []byte("CommitteeL2")
	CanPowerKeyPrefix    = []byte("PowerL2")
	UnStakeCountKey      = []byte("UnStakeCountL2")
	UnStakeItemKey       = []byte("UnStakeItemL2")
	AccountStakeRcPrefix = []byte("AccStakeRcL2")

	b104Len = len(math.MaxBig104.Bytes())
)

func CanBaseKeyByAddr(addr common.NodeAddress) []byte {
	return append(CanBaseKeyPrefix, addr.Bytes()...)
}
func CanBaseKeyBySuffix(addr []byte) []byte {
	return append(CanBaseKeyPrefix, addr...)
}

func CanMutableKeyByAddr(addr common.NodeAddress) []byte {
	return append(CanMutableKeyPrefix, addr.Bytes()...)
}

func CanMutableKeyBySuffix(addr []byte) []byte {
	return append(CanMutableKeyPrefix, addr...)
}

func OperatorKeyByAddr(addr common.NodeAddress) []byte {
	return append(OperatorKeyPrefix, addr.Bytes()...)
}

func CommitteeKeyByAddr(addr common.NodeAddress) []byte {
	return append(CommitteeKeyPrefix, addr.Bytes()...)
}

func TallyPowerKey(programVersion uint32, shares *big.Int, stakeBlockNum uint64, stakeTxIndex uint32, nodeID discover.NodeID) []byte {

	// Only sort Major and Minor
	// eg. 1.1.x => 1.1.0
	subVersion := math.MaxInt32 - xutil.CalcVersion(programVersion)
	sortVersion := common.Uint32ToBytes(subVersion)

	priority := new(big.Int).Sub(math.MaxBig104, shares)
	zeros := make([]byte, b104Len)
	prio := append(zeros, priority.Bytes()...)

	id := nodeID.Bytes()

	num := common.Uint64ToBytes(stakeBlockNum)
	txIndex := common.Uint32ToBytes(stakeTxIndex)

	// some index of pivots
	indexPre := len(CanPowerKeyPrefix)
	indexVersion := indexPre + len(sortVersion)
	indexPrio := indexVersion + len(prio)
	indexNum := indexPrio + len(num)
	indexTxIndex := indexNum + len(txIndex)
	size := indexTxIndex + len(id)

	// construct key
	key := make([]byte, size)
	copy(key[:len(CanPowerKeyPrefix)], CanPowerKeyPrefix)
	copy(key[indexPre:indexVersion], sortVersion)
	copy(key[indexVersion:indexPrio], prio)
	copy(key[indexPrio:indexNum], num)
	copy(key[indexNum:indexTxIndex], txIndex)
	copy(key[indexTxIndex:], id)
	return key
}

func GetUnStakeCountKey(epoch uint64) []byte {
	return append(UnStakeCountKey, common.Uint64ToBytes(epoch)...)
}

func GetUnStakeItemKey(epoch, index uint64) []byte {

	epochByte := common.Uint64ToBytes(epoch)
	indexByte := common.Uint64ToBytes(index)

	markPre := len(UnStakeItemKey)
	markEpoch := markPre + len(epochByte)
	size := markEpoch + len(indexByte)

	key := make([]byte, size)
	copy(key[:markPre], UnStakeItemKey)
	copy(key[markPre:markEpoch], epochByte)
	copy(key[markEpoch:], indexByte)

	return key
}

func GetAccountStakeRcKey(addr common.Address) []byte {
	return append(AccountStakeRcPrefix, addr.Bytes()...)
}
