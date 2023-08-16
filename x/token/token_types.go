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
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/rlp"
	"math/big"
)

type AccTokenAsset struct {
	TokenAddr common.Address // ERC20 Token合约地址
	Balance   *big.Int       // Token余额
}
type AccountAsset struct {
	Account      common.Address  // 账户地址
	NativeAmount *big.Int        // 原生代币余额
	TokenAssets  []AccTokenAsset // Token资产
}

// MintAccInfo 铸币账户信息
type MintAccInfo struct {
	AccList       []common.Address // 铸币地址列表
	TokenAddrList []common.Address // ERC20 Token合约地址列表
}

// Hash Calculating the hash of the AccountAsset.
func (acc AccountAsset) Hash() (common.Hash, error) {
	enVal, err := rlp.EncodeToBytes(acc)
	if err != nil {
		return common.ZeroHash, err
	}
	return crypto.Keccak256Hash(enVal), nil
}

type SettlementInfo struct {
	AccAssets []AccountAsset // 所有账户的资产信息
}

func (s SettlementInfo) Hash() (common.Hash, error) {
	enVal, err := rlp.EncodeToBytes(s)
	if err != nil {
		return common.ZeroHash, err
	}
	return crypto.Keccak256Hash(enVal), nil
}
