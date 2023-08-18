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

func existAccount(blockHash common.Hash, account common.Address) bool {
	accBytes, _ := rlp.EncodeToBytes(account)
	_, err := get(blockHash, accBytes)
	if nil != err {
		return false
	}

	return true
}

func saveAccount(blockHash common.Hash, account common.Address) error {
	accBytes, _ := rlp.EncodeToBytes(account)
	return put(blockHash, accBytes, []byte{1})
}

func SaveSettlementHash(blockHash common.Hash, hash common.Hash) error {
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

func SaveMintInfo(blockHash common.Hash, mintAccInfo MintAccInfo) error {
	// 判断是否需要存储
	if 0 == len(mintAccInfo.AccList) && 0 == len(mintAccInfo.TokenAddrList) {
		return nil
	}
	// 过滤已存储的信息
	var newAccInfo MintAccInfo
	for _, acc := range mintAccInfo.AccList {
		// 判断地址是否已经存储
		if !existAccount(blockHash, acc) {
			newAccInfo.AccList = append(newAccInfo.AccList, acc)
			// 存储地址
			saveAccount(blockHash, acc)
		}
	}

	for _, tokenAddr := range mintAccInfo.TokenAddrList {
		// 判断地址是否已经存储
		if !existAccount(blockHash, tokenAddr) {
			newAccInfo.TokenAddrList = append(newAccInfo.TokenAddrList, tokenAddr)
			// 存储地址
			saveAccount(blockHash, tokenAddr)
		}
	}

	if 0 < len(newAccInfo.AccList) || 0 < len(newAccInfo.TokenAddrList) {
		oldMintAccInfo, err := GetMintAccInfo(blockHash)
		if err != nil {
			return err
		}

		var saveMintAccInfo MintAccInfo
		if nil != oldMintAccInfo && oldMintAccInfo.AccList != nil {
			saveMintAccInfo.AccList = oldMintAccInfo.AccList
		}
		if nil != oldMintAccInfo && oldMintAccInfo.TokenAddrList != nil {
			saveMintAccInfo.TokenAddrList = oldMintAccInfo.TokenAddrList
		}

		for _, acc := range newAccInfo.AccList {
			saveMintAccInfo.AccList = append(saveMintAccInfo.AccList, acc)
		}

		for _, tokenAddr := range newAccInfo.TokenAddrList {
			saveMintAccInfo.TokenAddrList = append(saveMintAccInfo.TokenAddrList, tokenAddr)
		}
		put(blockHash, KeyMintAccInfo(), saveMintAccInfo)
	}

	return nil
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
