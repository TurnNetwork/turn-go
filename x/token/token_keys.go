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
	"github.com/bubblenet/bubble/rlp"
)

var (
	keyPrefixMintAccInfo    = []byte("MintAccInfo")
	keyPrefixSettlementHash = []byte("SettlementHash")
	keyPrefixTxHash         = []byte("TxHash")
)

func KeyMintAccInfo() []byte {
	return keyPrefixMintAccInfo
}

func KeyPrefixTxHash(txHash common.Hash) []byte {
	hash, err := rlp.EncodeToBytes(txHash)
	if nil != err {
		return nil
	}
	return append(keyPrefixTxHash, hash...)
}

func KeyPrefixSettlementHash() []byte {
	return keyPrefixSettlementHash
}
