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

package plugin

import (
	"fmt"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/event"
	"github.com/bubblenet/bubble/x/token"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	genesisBlockHash = common.HexToHash("")
	TestBubbleId     = big.NewInt(1)
	TestTxHash       = common.HexToHash("0x12c171900f010b17e969702efa044d077e86808212c171900f010b17e969702e")
	Op2RPC           = "http://127.0.0.1:2789"
	Op1Addr          = common.HexToAddress("0x11111d46d924CC8437c806721496599FC3FFA268")
	TestAddr         = common.HexToAddress("0xc9E1C2B330Cf7e759F2493c5C754b34d98B07f93")
	TestAccList      = []common.Address{
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

func handleEvent(eventMux *event.TypeMux, t *testing.T) {
	events := eventMux.Subscribe(token.SettleTask{})
	defer events.Unsubscribe()
	for {
		select {
		case settleMsg := <-events.Chan():
			if settleMsg == nil {
				t.Error("ev is nil, may be Server closing")
				continue
			}
			settleData, ok := settleMsg.Data.(token.SettleTask)
			if !ok {
				t.Error("failed to receive settleBubble data conversion type")
				continue
			}
			t.Logf("P2P Received the settleBubble Task is:%v", settleData)
			// handle task
			hash, err := TokenInstance().HandleSettleTask(&settleData)
			if err != nil {
				t.Error("failed to process settleBubble task")
				continue
			}
			t.Logf("the processing and settleBubble task succeeded, tx hash:%v", common.BytesToHash(hash).Hex())
		}
	}
}

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

/**
Standard test cases
*/

func TestTokenPlugin_BeginBlock(t *testing.T) {
	// nothings in that
}

func TestTokenPlugin_EndBlock(t *testing.T) {
	// nothings in that
}

func TestTokenPlugin_Confirmed(t *testing.T) {
	// nothings in that
}

func TestTokenPlugin_PostSettlementTask(t *testing.T) {
	bp := TokenInstance()
	eventMux := &event.TypeMux{}
	bp.SetEventMux(eventMux)

	go handleEvent(eventMux, t)
	accAsset := token.AccountAsset{
		Account:      TestAddr,
		NativeAmount: TestNativeAmount,
	}
	for _, tokenAddr := range TestERC20AddrList {
		tokenAsset := token.AccTokenAsset{
			TokenAddr: tokenAddr,
			Balance:   TestTokenAmount,
		}
		accAsset.TokenAssets = append(accAsset.TokenAssets, tokenAsset)
	}
	mintTokenTask := token.SettleTask{
		BubbleID: TestBubbleId,
		TxHash:   TestTxHash,
	}
	err := bp.PostSettlementTask(&mintTokenTask)
	assert.Nil(t, err, fmt.Sprintf("failed to post Settlement task event, err: %v", err))
}

func TestTokenPlugin_AddMintAccInfo_GetMintAccInfo(t *testing.T) {
	newBlockHash := newFirstBlock()
	// MOCK
	// store MintAccInfo
	mintAcc := token.MintAccInfo{}
	for _, acc := range TestAccList {
		mintAcc.AccList = append(mintAcc.AccList, acc)
	}
	for _, tokenAddr := range TestERC20AddrList {
		mintAcc.TokenAddrList = append(mintAcc.TokenAddrList, tokenAddr)
	}
	bp := TokenInstance()
	// Number of additions
	addTimes := 2
	for i := 0; i < addTimes; i++ {
		err := bp.AddMintAccInfo(newBlockHash, mintAcc)
		if err != nil {
			fmt.Errorf("failed to AddMintAccInfo, %v", err)
		}
		assert.True(t, err == nil)
	}

	mintAccInfo, err := bp.GetMintAccInfo(newBlockHash)
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
