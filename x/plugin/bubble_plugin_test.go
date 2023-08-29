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
	"github.com/bubblenet/bubble/event"
	"github.com/bubblenet/bubble/x/bubble"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	TestBubbleId = big.NewInt(1)
	TestTxHash   = common.HexToHash("0x12c171900f010b17e969702efa044d077e86808212c171900f010b17e969702e")
	Op2RPC       = "http://127.0.0.1:2789"
	Op1Addr      = common.HexToAddress("0x11111d46d924CC8437c806721496599FC3FFA268")
	TestAddr     = common.HexToAddress("0xc9E1C2B330Cf7e759F2493c5C754b34d98B07f93")
	TestAccList  = []common.Address{
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
	events := eventMux.Subscribe(bubble.MintTokenTask{})
	defer events.Unsubscribe()
	for {
		select {
		case mintTokenMsg := <-events.Chan():
			if mintTokenMsg == nil {
				t.Error("ev is nil, may be Server closing")
				continue
			}
			mintToken, ok := mintTokenMsg.Data.(bubble.MintTokenTask)
			if !ok {
				t.Error("failed to receive mintToken data conversion type")
				continue
			}
			t.Logf("P2P Received the MintToken Task is:%v", mintToken)
			// handle task
			hash, err := BubbleInstance().HandleMintTokenTask(&mintToken)
			if err != nil {
				t.Error("failed to process mintToken task")
				continue
			}
			t.Logf("the processing and MintToken task succeeded, tx hash:%v", common.BytesToHash(hash).Hex())
		}
	}
}

/**
Standard test cases
*/

func TestBubblePlugin_BeginBlock(t *testing.T) {
	// nothings in that
}

func TestBubblePlugin_EndBlock(t *testing.T) {
	// nothings in that
}

func TestBubblePlugin_Confirmed(t *testing.T) {
	// nothings in that
}

func TestBubblePlugin_PostMintTokenEvent(t *testing.T) {
	bp := BubbleInstance()
	eventMux := &event.TypeMux{}
	bp.SetEventMux(eventMux)

	go handleEvent(eventMux, t)
	accAsset := bubble.AccountAsset{
		Account:      TestAddr,
		NativeAmount: TestNativeAmount,
	}
	for _, tokenAddr := range TestERC20AddrList {
		tokenAsset := bubble.AccTokenAsset{
			TokenAddr: tokenAddr,
			Balance:   TestTokenAmount,
		}
		accAsset.TokenAssets = append(accAsset.TokenAssets, tokenAsset)
	}
	mintTokenTask := bubble.MintTokenTask{
		BubbleID: TestBubbleId,
		TxHash:   TestTxHash,
		RPC:      Op2RPC,
		OpAddr:   Op1Addr,
		AccAsset: &accAsset,
	}
	err := bp.PostMintTokenEvent(&mintTokenTask)
	assert.Nil(t, err, fmt.Sprintf("failed to post mintToken task event, err: %v", err))
}
