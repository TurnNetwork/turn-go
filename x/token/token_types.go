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
	TokenAddr common.Address // ERC20 Token contract address
	Balance   *big.Int       // Token balance
}
type AccountAsset struct {
	Account      common.Address  // Account address
	NativeAmount *big.Int        // Native token balances
	TokenAssets  []AccTokenAsset // Token assets
}

// MintAccInfo Minting account information
type MintAccInfo struct {
	AccList       []common.Address // List of minting addresses
	TokenAddrList []common.Address // List of ERC20 Token contract addresses
}

type SettlementInfo struct {
	AccAssets []AccountAsset // Asset information for all accounts
}

type SettleTask struct {
	TxHash     common.Hash    // The transaction hash of the staking Token transaction
	BubbleID   *big.Int       // bubbleID for settlement
	SettleInfo SettlementInfo // Asset information for all accounts
}

func (s SettlementInfo) Hash() (common.Hash, error) {
	enVal, err := rlp.EncodeToBytes(s)
	if err != nil {
		return common.ZeroHash, err
	}
	return crypto.Keccak256Hash(enVal), nil
}
