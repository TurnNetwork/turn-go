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

package vm

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/ethclient"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/token"
	"github.com/status-im/keycard-go/hexutils"
	"log"
	"math/big"
	"testing"
)

// 解析Settlement交易日志
func TestDecodeSettlementTxLog(t *testing.T) {
	var settlementInfo token.SettlementInfo
	var m [][]byte
	data := "f86030b85df85bf859f857946a311b9d42ea0cb4f62760383c0eff06ac68f1f78ba56fa60cd461dbb4580000f5d9948a16806861ca61ef5d3bb99498b94d3367a248ba83989680da9412c171900f010b17e969702efa044d077e8680828401312d00"
	if err := rlp.DecodeBytes(hexutils.HexToBytes(data), &m); err != nil {
		t.Error(err)
	}
	var code string
	if err := rlp.DecodeBytes(m[0], &code); err != nil {
		t.Error(err)
	}
	if err := rlp.DecodeBytes(m[1], &settlementInfo); err != nil {
		t.Error(err)
	}

	fmt.Printf("settlementInfo: %v\n", settlementInfo)
	return
}

// rpc调用
func TestRpc(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:1789")
	if err != nil {
		t.Error("error", err)
	}
	if client != nil {
		// 发送交易
		// 构建交易参数
		priKey := "000000000000cef6621103622f27a31d65c0856a0a66ba2fd03e4663161f1c5b"
		toAddr := "0x6A311b9D42Ea0Cb4F62760383C0EfF06Ac68F1f7"
		privateKey, err := crypto.HexToECDSA(priKey)
		if err != nil {
			t.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			t.Fatal("无法获取公钥")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			t.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		toAddress := common.HexToAddress(toAddr)
		value := big.NewInt(1000000000000000000) // 发送 1 ETH
		gasLimit := uint64(21000)
		data := []byte("")

		// 创建交易对象
		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

		// 使用发送方的私钥进行交易签名
		chainID, err := client.ChainID(context.Background())
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			t.Fatal(err)
		}

		// 发送交易
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			t.Fatal(err)
		}

		hash := signedTx.Hash().Hex()
		fmt.Printf("hash:%s\n", hash)
	}
}
