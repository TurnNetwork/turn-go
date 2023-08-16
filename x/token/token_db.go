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
	"encoding/json"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/x/xcom"
)

func ExistAccount(state xcom.StateDB, account common.Address) bool {
	accBytes, _ := json.Marshal(account)
	if len(state.GetState(vm.TokenContractAddr, accBytes)) > 0 {
		return true
	}
	return false
}

func SaveAccount(state xcom.StateDB, account common.Address) error {
	accBytes, _ := json.Marshal(account)
	state.SetState(vm.TokenContractAddr, accBytes, []byte{1})
	return nil
}

func SaveSettlementHash(state xcom.StateDB, hash common.Hash) error {
	hashBytes, _ := json.Marshal(hash)
	state.SetState(vm.TokenContractAddr, KeyPrefixSettlementHash(), hashBytes)
	return nil
}

func GetSettlementHash(state xcom.StateDB) (*common.Hash, error) {
	var hash common.Hash
	hashBytes := state.GetState(vm.TokenContractAddr, KeyPrefixSettlementHash())
	if len(hashBytes) > 0 {
		if err := json.Unmarshal(hashBytes, &hash); err != nil {
			return nil, err
		}
		return &hash, nil
	} else {
		return nil, nil
	}
}

func SaveMintInfo(state xcom.StateDB, mintAccInfo MintAccInfo) error {
	// 判断是否需要存储
	if 0 == len(mintAccInfo.AccList) && 0 == len(mintAccInfo.TokenAddrList) {
		return nil
	}
	// 过滤已存储的信息
	var newAccInfo MintAccInfo
	for _, acc := range mintAccInfo.AccList {
		// 判断地址是否已经存储
		if !ExistAccount(state, acc) {
			newAccInfo.AccList = append(newAccInfo.AccList, acc)
			// 存储地址
			SaveAccount(state, acc)
		}
	}

	for _, tokenAddr := range mintAccInfo.TokenAddrList {
		// 判断地址是否已经存储
		if !ExistAccount(state, tokenAddr) {
			newAccInfo.TokenAddrList = append(newAccInfo.TokenAddrList, tokenAddr)
			// 存储地址
			SaveAccount(state, tokenAddr)
		}
	}

	if 0 < len(newAccInfo.AccList) || 0 < len(newAccInfo.TokenAddrList) {
		oldMintAccInfo, err := GetMintAccInfo(state)
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
		mintAccInfoBytes, _ := json.Marshal(saveMintAccInfo)
		state.SetState(vm.TokenContractAddr, KeyMintAccInfo(), mintAccInfoBytes)
	}

	return nil
}

func GetMintAccInfo(state xcom.StateDB) (*MintAccInfo, error) {
	var mintAccInfo MintAccInfo
	mintAccInfoBytes := state.GetState(vm.TokenContractAddr, KeyMintAccInfo())
	if len(mintAccInfoBytes) > 0 {
		if err := json.Unmarshal(mintAccInfoBytes, &mintAccInfo); err != nil {
			return nil, err
		}
		return &mintAccInfo, nil
	} else {
		return nil, nil
	}
}
