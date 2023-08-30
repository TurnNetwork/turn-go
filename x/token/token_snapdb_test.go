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
	"fmt"
	_ "fmt"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"

	"github.com/bubblenet/bubble/common"
)

var (
	testBubbleId     = big.NewInt(1)
	sndb             = snapshotdb.Instance()
	genesisBlockHash = common.HexToHash("")

	L1TxHash       = common.HexToHash("0x12c171900f010b17e969702efa044d077e86808212c171900f010b17e969702e")
	L2TxHash       = common.HexToHash("0x00000000000000000000000000000000000000886d5ba2d3dfb2e2f6a1814f22")
	Op1NodeId      = discover.NodeID{0x1}
	Op1ElectronRPC = "http://127.0.0.1:1789"
	Op1RPC         = "http://127.0.0.1:2789"
	Op1Addr        = common.HexToAddress("0x11111d46d924CC8437c806721496599FC3FFA268")
	Op1Amount      = big.NewInt(100000000000)

	Op2NodeId      = discover.NodeID{0x2}
	Op2ElectronRPC = "http://127.0.0.1:3789"
	Op2RPC         = "http://127.0.0.1:4789"
	Op2Addr        = common.HexToAddress("0x22222d46d924CC8437c806721496599FC3FFA268")
	Op2Amount      = big.NewInt(200000000000)

	TestAddr    = common.HexToAddress("0xc9E1C2B330Cf7e759F2493c5C754b34d98B07f93")
	TestAccList = []common.Address{
		common.HexToAddress("0xd87E10F8efd2C32f5e88b7C279953aEF6EE58902"),
		common.HexToAddress("0xeAEc60C738eeD9468e6AcCc1d403faCF1A670F6D"),
		common.HexToAddress("0x5c5994165265Ac31AAFE874a231f2C5d0eF29C3a"),
	}

	TestERC20AddrList = []common.Address{
		common.HexToAddress("0xe200000000000000000000000000000000000000"),
		common.HexToAddress("0xe200000000000000000000000000000000000001"),
		common.HexToAddress("0xe200000000000000000000000000000000000002"),
	}
	TestTokenAmount  = big.NewInt(100000000000)
	TestNativeAmount = big.NewInt(200000000000)
)

func newFirstBlock() common.Hash {
	header := types.Header{
		Number: big.NewInt(1),
	}
	newBlockHash := header.Hash()

	err := sndb.NewBlock(header.Number, genesisBlockHash, newBlockHash)
	if err != nil {
		fmt.Errorf("newBlock, %v", err)
	}
	return newBlockHash
}

func TestTokenDB_StoreAccount_ExistAccount(t *testing.T) {
	newBlockHash := newFirstBlock()
	// MOCK
	// store account
	if err := StoreAccount(newBlockHash, TestAddr); err != nil {
		fmt.Errorf("failed to StoreAccount, %v", err)
	}

	isExist := ExistAccount(newBlockHash, TestAddr)
	assert.True(t, isExist)
}

func TestTokenDB_StoreSettlementHash_GetSettlementHash(t *testing.T) {
	newBlockHash := newFirstBlock()
	// MOCK
	// store SettlementInfo
	// Assembly settlement information
	settleInfo := SettlementInfo{}
	for _, addr := range TestAccList {
		accAsset := AccountAsset{
			Account:      addr,
			NativeAmount: TestNativeAmount,
		}
		for _, tokenAddr := range TestERC20AddrList {
			tokenAsset := AccTokenAsset{
				TokenAddr: tokenAddr,
				Balance:   TestTokenAmount,
			}
			accAsset.TokenAssets = append(accAsset.TokenAssets, tokenAsset)
		}
		settleInfo.AccAssets = append(settleInfo.AccAssets, accAsset)
	}

	// Calculate the current account settlement Hash
	hash, err := settleInfo.Hash()
	assert.True(t, err == nil)
	if err := StoreSettlementHash(newBlockHash, hash); err != nil {
		fmt.Errorf("failed to StoreSettlementHash, %v", err)
	}

	settleHash, err := GetSettlementHash(newBlockHash)
	assert.True(t, err == nil)
	assert.True(t, settleHash != nil)
	assert.Equal(t, *settleHash, hash, "The stored and fetched settlement transaction hashes are inconsistent")
}

func TestTokenDB_StoreMintInfo_GetMintAccInfo(t *testing.T) {
	newBlockHash := newFirstBlock()
	// MOCK
	// store MintAccInfo
	mintAcc := MintAccInfo{}
	for _, acc := range TestAccList {
		mintAcc.AccList = append(mintAcc.AccList, acc)
	}
	for _, tokenAddr := range TestERC20AddrList {
		mintAcc.TokenAddrList = append(mintAcc.TokenAddrList, tokenAddr)
	}

	err := StoreMintInfo(newBlockHash, mintAcc)
	if err != nil {
		fmt.Errorf("failed to StoreMintInfo, %v", err)
	}
	assert.True(t, err == nil)

	mintAccInfo, err := GetMintAccInfo(newBlockHash)
	assert.True(t, err == nil)
	assert.True(t, mintAccInfo != nil)
	assert.Equal(t, len(mintAcc.AccList), len(TestAccList), "Store and retrieve the COINS account number")
	for i, acc := range mintAccInfo.AccList {
		assert.Equal(t, acc, TestAccList[i], "The account addresses of the mints stored and fetched are inconsistent")
	}
	assert.Equal(t, len(mintAcc.TokenAddrList), len(TestERC20AddrList), "Store and retrieve the COINS ERC20 Token account number")
	for i, tokenAddr := range mintAccInfo.TokenAddrList {
		assert.Equal(t, tokenAddr, TestERC20AddrList[i], "Storing and retrieving the COINS erc20 token address inconsistencies\n")
	}
}

func TestTokenDB_StoreL1HashToL2Hash_GetL2HashByL1Hash(t *testing.T) {
	newBlockHash := newFirstBlock()
	// MOCK
	err := StoreL1HashToL2Hash(newBlockHash, L1TxHash, L2TxHash)
	if err != nil {
		fmt.Errorf("failed to StoreL1HashToL2Hash, %v", err)
	}
	assert.True(t, err == nil)

	getL2TxHash, err := GetL2HashByL1Hash(newBlockHash, L1TxHash)
	assert.True(t, err == nil)
	assert.True(t, getL2TxHash != nil)
	assert.Equal(t, *getL2TxHash, L2TxHash, "The stored and fetched sub-chain transaction hashes are inconsistent")
}
