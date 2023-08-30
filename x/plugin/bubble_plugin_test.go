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
	"github.com/bubblenet/bubble/common/hexutil"
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

func TestVRF(t *testing.T) {
	queue := VRFQueue{
		{v: "a", w: big.NewInt(int64(4))},
		{v: "b", w: big.NewInt(int64(15))},
		{v: "c", w: big.NewInt(int64(20))},
		{v: "d", w: big.NewInt(int64(12))},
		{v: "e", w: big.NewInt(int64(31))},
		{v: "f", w: big.NewInt(int64(1))},
		{v: "g", w: big.NewInt(int64(7))},
		{v: "h", w: big.NewInt(int64(8))},
		{v: "i", w: big.NewInt(int64(6))},
		{v: "j", w: big.NewInt(int64(57))},
		//{w: big.NewInt(int64(5))},
		//{w: big.NewInt(int64(3))},
		//{w: big.NewInt(int64(4))},
		//{w: big.NewInt(int64(17))},
		//{w: big.NewInt(int64(46))},
		//{w: big.NewInt(int64(22))},
		//{w: big.NewInt(int64(16))},
		//{w: big.NewInt(int64(35))},
	}
	curNonce := hexutil.MustDecode("0x0299e70aa6b1c6028b61cc8c133180e88269ceb9e58bba90709c0e33dced26ca9b")
	preNoce := [][]byte{
		hexutil.MustDecode("0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23"),
		hexutil.MustDecode("0x03fa453d41b986f1ca7dfaccf2421c5ab05ccfc3f50ec69aef573881798434d74c"),
		hexutil.MustDecode("0x03a49430b191beb6416f220a451f6d7b87c014be6d1c9e6df4ceb016486921a8a4"),
		hexutil.MustDecode("0x0244b07649864495899091cf8646145f2ba8510af281c6157c820bc9ffed0f7fac"),
		hexutil.MustDecode("0x0327bae119cc3eeaa3eac5550c7714b5c4dc1941b6b23ef3d088271b39c0ab8fbf"),
		hexutil.MustDecode("0x038a49f3fb7635b7941f9897098a9db318f6a1630f9a68cad856bbd9ccdc124b1f"),
		hexutil.MustDecode("0x034bb00ebbc6ee84bdd3038e42c6135e77924082d77a7ea5379249f18d324cec78"),
		hexutil.MustDecode("0x03004e1515a56519794e0777e948d1b2866503cb6e0c5b2efdc4cc2504c140fa66"),
		hexutil.MustDecode("0x0306e54d8e29ffb5025e1d4f04f22ba873ea962ddc4a66dbe14b07f430cc8ee9f1"),
		hexutil.MustDecode("0x02498c4b7a4ae35e5d48756153030bb1af0709b4b62da8b38924302792c64c4482"),
	}
	for i := 0; i < 1; i++ {
		fmt.Println("=============================")
		if vrfed, err := VRF(queue, 10, curNonce, preNoce); err != nil {
			fmt.Printf("选举错误：%s \n", err.Error())
		} else {
			for _, v := range vrfed {
				fmt.Printf("VRF值： %s %d %d \n", v.v, v.x, v.w)
			}

		}
	}
}
