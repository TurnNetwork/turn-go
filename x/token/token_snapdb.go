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

package token

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/rlp"
)

func get(blockHash common.Hash, key []byte) ([]byte, error) {
	return snapshotdb.Instance().Get(blockHash, key)
}

func put(blockHash common.Hash, key []byte, value interface{}) error {
	bytes, err := rlp.EncodeToBytes(value)
	if err != nil {
		return err
	}
	return snapshotdb.Instance().Put(blockHash, key, bytes)
}

func del(blockHash common.Hash, key []byte) error {
	return snapshotdb.Instance().Del(blockHash, key)
}

func ExistAccount(blockHash common.Hash, account common.Address) bool {
	accBytes, _ := rlp.EncodeToBytes(account)
	_, err := get(blockHash, accBytes)
	if nil != err {
		return false
	}

	return true
}

func StoreAccount(blockHash common.Hash, account common.Address) error {
	accBytes, _ := rlp.EncodeToBytes(account)
	return put(blockHash, accBytes, []byte{1})
}

func StoreSettlementHash(blockHash common.Hash, hash common.Hash) error {
	return put(blockHash, KeyPrefixSettlementHash(), hash)
}

func GetSettlementHash(blockHash common.Hash) (*common.Hash, error) {
	hashBytes, err := get(blockHash, KeyPrefixSettlementHash())
	if snapshotdb.IsDbNotFoundErr(err) {
		return nil, nil
	}

	if nil == err && len(hashBytes) > 0 {
		var hash common.Hash
		if err = rlp.DecodeBytes(hashBytes, &hash); err != nil {
			return nil, err
		}
		return &hash, nil
	}
	return nil, err
}

func StoreMintInfo(blockHash common.Hash, mintAccInfo MintAccInfo) error {
	return put(blockHash, KeyMintAccInfo(), mintAccInfo)
}

func GetMintAccInfo(blockHash common.Hash) (*MintAccInfo, error) {
	mintAccInfoBytes, err := get(blockHash, keyPrefixMintAccInfo)
	if snapshotdb.IsDbNotFoundErr(err) {
		return nil, nil
	}

	if err == nil && len(mintAccInfoBytes) > 0 {
		var mintAccInfo MintAccInfo
		if err = rlp.DecodeBytes(mintAccInfoBytes, &mintAccInfo); err != nil {
			return nil, err
		}
		return &mintAccInfo, nil
	}
	return nil, err
}

func StoreL1HashToL2Hash(blockHash common.Hash, L1TxHash common.Hash, L2TxHash common.Hash) error {
	return put(blockHash, KeyPrefixTxHash(L1TxHash), L2TxHash)
}

func GetL2HashByL1Hash(blockHash common.Hash, L1TxHash common.Hash) (*common.Hash, error) {
	data, err := get(blockHash, KeyPrefixTxHash(L1TxHash))
	if err != nil {
		return nil, err
	}

	var L2TxHash common.Hash
	if err := rlp.DecodeBytes(data, &L2TxHash); err != nil {
		return nil, err
	} else {
		return &L2TxHash, nil
	}
}
