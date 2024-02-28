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

package bubble

import (
	"fmt"
	_ "fmt"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/p2p/enode"
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
	Op1NodeId      = enode.IDv0{0x1}
	Op1ElectronRPC = "http://127.0.0.1:1789"
	Op1RPC         = "http://127.0.0.1:2789"
	Op1Addr        = common.HexToAddress("0x11111d46d924CC8437c806721496599FC3FFA268")
	Op1Amount      = big.NewInt(100000000000)

	Op2NodeId      = enode.IDv0{0x2}
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

	StakingTokenTxHash = []common.Hash{
		common.HexToHash("1"),
		common.HexToHash("2"),
		common.HexToHash("3"),
	}
	WithdrewTokenTxHash = []common.Hash{
		common.HexToHash("4"),
		common.HexToHash("5"),
		common.HexToHash("6"),
	}
	SettleBubbleTxHash = []common.Hash{
		common.HexToHash("7"),
		common.HexToHash("8"),
		common.HexToHash("9"),
	}
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

func verify_bubble_tx_hash(blockHash common.Hash, txHashList []common.Hash, txType TxType, t *testing.T) {
	bubDB := NewDB()
	if err := bubDB.StoreTxHashListToBub(blockHash, testBubbleId, txHashList, txType); err != nil {
		fmt.Errorf("failed to StoreTxHashListToBub, %v", err)
	}

	getTxHashes, err := bubDB.GetTxHashListByBub(blockHash, testBubbleId, txType)
	if getTxHashes == nil || err != nil {
		fmt.Errorf("failed to get %d tx hash list of bubble, %v", txType, err)
	}
	assert.True(t, nil == err)
	assert.True(t, nil != getTxHashes, "get transaction Hash list of bubble err")
	assert.True(t, len(txHashList) == len(*getTxHashes), "The length of fetched and stored transaction hash lists is not consistent")
	for i, txHash := range txHashList {
		assert.True(t, txHash == (*getTxHashes)[i], "Fetched and stored transaction hashes are inconsistent")
	}
}

func TestBubbleDB_StoreBubBasic_GetBubBasic(t *testing.T) {
	newBlockHash := newFirstBlock()
	// MOCK
	bubDB := NewDB()

	var opL1s []*Operator
	opL1 := Operator{
		NodeId:      Op1NodeId,
		ElectronRPC: Op1ElectronRPC,
		RPC:         Op1RPC,
		OpAddr:      Op1Addr, // The initial sub-chain operation address as the sender
		Balance:     Op1Amount,
	}
	opL1s = append(opL1s, &opL1)
	// sub-chain operators config
	var opL2s []*Operator
	opL2 := Operator{
		NodeId:      Op2NodeId,
		ElectronRPC: Op2ElectronRPC,
		RPC:         Op2RPC,
		OpAddr:      Op2Addr, // The initial sub-chain operation address as the sender
		Balance:     Op2Amount,
	}
	opL2s = append(opL2s, &opL2)
	basics := BasicsInfo{
		BubbleId:    testBubbleId,
		OperatorsL1: opL1s,
		OperatorsL2: opL2s,
		MicroNodes:  nil,
	}
	// store bubble basics
	if err := bubDB.StoreBasicsInfo(newBlockHash, testBubbleId, &basics); err != nil {
		fmt.Errorf("failed to StoreBasicsInfo, %v", err)
	}

	basic, err := bubDB.GetBasicsInfo(newBlockHash, testBubbleId)
	if basic == nil || err != nil {
		fmt.Errorf("failed to get bubble basic, %v", err)
	}
	assert.True(t, nil == err)
	assert.True(t, nil != basic)
	assert.Equal(t, basic.BubbleId, testBubbleId, "Query Bubble ID error")

	assert.Equal(t, basic.OperatorsL1[0].NodeId, Op1NodeId, "Query Bubble main-chain Operator NodeID error")
	assert.Equal(t, basic.OperatorsL1[0].ElectronRPC, Op1ElectronRPC, "Query Bubble main-chain Operator ElectronRPC error")
	assert.Equal(t, basic.OperatorsL1[0].RPC, Op1RPC, "Query Bubble main-chain Operator RPC error")
	assert.Equal(t, basic.OperatorsL1[0].OpAddr, Op1Addr, "Query Bubble main-chain Operator address error")
	assert.Equal(t, basic.OperatorsL1[0].Balance, Op1Amount, "Query Bubble main-chain Operator balance error")

	assert.Equal(t, basic.OperatorsL2[0].NodeId, Op2NodeId, "Query Bubble sub-chain Operator NodeID error")
	assert.Equal(t, basic.OperatorsL2[0].ElectronRPC, Op2ElectronRPC, "Query Bubble sub-chain Operator ElectronRPC error")
	assert.Equal(t, basic.OperatorsL2[0].RPC, Op2RPC, "Query Bubble sub-chain Operator RPC error")
	assert.Equal(t, basic.OperatorsL2[0].OpAddr, Op2Addr, "Query Bubble sub-chain Operator address error")
	assert.Equal(t, basic.OperatorsL2[0].Balance, Op2Amount, "Query Bubble sub-chain Operator balance error")

}

func TestBubbleDB_StoreStateInfo_GetStateInfo(t *testing.T) {
	newBlockHash := newFirstBlock()
	bubDB := NewDB()
	// store bubble state
	status := &StateInfo{
		BubbleId:        testBubbleId,
		State:           ActiveState,
		ContractCount:   0,
		CreateBlock:     0,
		PreReleaseBlock: 1,
		ReleaseBlock:    2,
	}
	if err := bubDB.StoreStateInfo(newBlockHash, testBubbleId, status); err != nil {
		fmt.Errorf("failed to StoreBasicsInfo, %v", err)
	}

	state, err := bubDB.GetStateInfo(newBlockHash, testBubbleId)
	if state == nil || err != nil {
		fmt.Errorf("failed to get bubble state, %v", err)
	}
	assert.True(t, nil == err)
	assert.True(t, nil != state)
	assert.Equal(t, state.State, ActiveState, "Query Bubble State error")
}

func TestBubbleDB_StoreAccListOfBub_GetAccListOfBub(t *testing.T) {
	newBlockHash := newFirstBlock()
	bubDB := NewDB()
	// store bubble account list
	if err := bubDB.StoreAccListOfBub(newBlockHash, testBubbleId, TestAccList); err != nil {
		fmt.Errorf("failed to StoreAccListOfBub, %v", err)
	}

	accList, err := bubDB.GetAccListOfBub(newBlockHash, testBubbleId)
	if accList == nil || err != nil {
		fmt.Errorf("failed to get account list of bubble, %v", err)
	}
	assert.True(t, nil == err)
	assert.True(t, nil != accList)
	assert.Equal(t, len(TestAccList), len(accList), "Query Bubble account list error")
	for i, addr := range TestAccList {
		assert.Equal(t, addr, accList[i], "Query Bubble account error")
	}
}

func TestBubbleDB_StoreAccAssetToBub_GetAccAssetToBub(t *testing.T) {
	newBlockHash := newFirstBlock()
	bubDB := NewDB()
	// store account Asset to bubble
	storeAccAsset := AccountAsset{
		Account:      TestAddr,
		NativeAmount: TestNativeAmount,
	}
	for _, tokenAddr := range TestERC20AddrList {
		tokenAsset := AccTokenAsset{
			TokenAddr: tokenAddr,
			Balance:   TestTokenAmount,
		}
		storeAccAsset.TokenAssets = append(storeAccAsset.TokenAssets, tokenAsset)
	}
	if err := bubDB.StoreAccAssetToBub(newBlockHash, testBubbleId, storeAccAsset); err != nil {
		fmt.Errorf("failed to StoreAccAssetToBub, %v", err)
	}

	accAsset, err := bubDB.GetAccAssetOfBub(newBlockHash, testBubbleId, TestAddr)
	if accAsset == nil || err != nil {
		fmt.Errorf("failed to get account Asset of bubble, %v", err)
	}
	assert.True(t, nil == err)
	assert.True(t, nil != accAsset)
	assert.Equal(t, accAsset.Account, TestAddr, "Query Bubble AccountAsset of account error")
	assert.Equal(t, accAsset.NativeAmount, TestNativeAmount, "Query Bubble AccountAsset of NativeAmount error")
	for i, accTokenAsset := range accAsset.TokenAssets {
		assert.Equal(t, accTokenAsset.TokenAddr, TestERC20AddrList[i], "Query Bubble AccountAsset of token address error")
		assert.Equal(t, accTokenAsset.Balance, TestTokenAmount, "Query Bubble AccountAsset of token balance error")
	}
}

func TestBubbleDB_StoreTxHashListToBub_GetTxHashListByBub(t *testing.T) {
	newBlockHash := newFirstBlock()
	// stakingToken transaction
	verify_bubble_tx_hash(newBlockHash, StakingTokenTxHash, StakingToken, t)
	// withdrewToken transaction
	verify_bubble_tx_hash(newBlockHash, WithdrewTokenTxHash, WithdrewToken, t)
	// settleBubble transaction
	verify_bubble_tx_hash(newBlockHash, SettleBubbleTxHash, SettleBubble, t)
}

func TestBubbleDB_StoreL2HashToL1Hash_GetL1HashByL2Hash(t *testing.T) {
	newBlockHash := newFirstBlock()
	bubDB := NewDB()

	if err := bubDB.StoreL2HashToL1Hash(newBlockHash, testBubbleId, L1TxHash, L2TxHash); err != nil {
		fmt.Errorf("failed to StoreL2HashToL1Hash, %v", err)
	}

	getL1TxHash, err := bubDB.GetL1HashByL2Hash(newBlockHash, testBubbleId, L2TxHash)
	if getL1TxHash == nil || err != nil {
		fmt.Errorf("failed to get L1 transaction hash by L2 transaction hash: %v", err)
	}
	assert.True(t, nil == err)
	assert.True(t, nil != getL1TxHash)
	assert.Equal(t, *getL1TxHash, L1TxHash, "failed to get L1 transaction hash by L2 transaction hash")
}
