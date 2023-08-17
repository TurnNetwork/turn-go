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
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/ethclient"
	"github.com/bubblenet/bubble/event"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/token"
	"github.com/bubblenet/bubble/x/xcom"
	"math/big"
	"sync"
)

var (
	tokenPluginOnce sync.Once
	tkp             *TokenPlugin
)

type TokenPlugin struct {
	MainOpAddr  common.Address // Main chain operator address
	IsSubOpNode bool           // Whether it is a child chain operator node
	subOpPriKey string         // The child chain operates the private key of the node, which is used to sign and send the transactions of the main chain
	OpConfig    *params.OpConfig
	eventMux    *event.TypeMux
}

func TokenInstance() *TokenPlugin {
	tokenPluginOnce.Do(func() {
		log.Info("Init Token plugin ...")
		tkp = &TokenPlugin{}
	})
	return tkp
}

func (tkp *TokenPlugin) SetSubOpPriKey(subOpPriKey string) {
	tkp.subOpPriKey = subOpPriKey
}

func (tkp *TokenPlugin) SetEventMux(eventMux *event.TypeMux) {
	tkp.eventMux = eventMux
}

// AddSettlementTask Add the checkout task to the subscription event
func (tkp *TokenPlugin) AddSettlementTask(settlementInfo *token.SettlementInfo) {
	if err := tkp.eventMux.Post(*settlementInfo); nil != err {
		log.Error("post settlementInfo failed", "err", err)
	}
}

func genSettleTxRlpData(funcType uint16, settlementInfo *token.SettlementInfo) []byte {
	var params [][]byte
	params = make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(funcType)
	params = append(params, fnType)

	accAsset, _ := rlp.EncodeToBytes(settlementInfo)
	params = append(params, accAsset)
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, params)
	if err != nil {
		panic(fmt.Errorf("%d encode rlp data fail: %v", funcType, err))
	} else {
		rlpData := hexutil.Encode(buf.Bytes())
		fmt.Printf("funcType:%d rlp data = %s\n", funcType, rlpData)
		return buf.Bytes()
	}
	return nil
}

//HandleSettlementTask Handle settlement tasks
func (tkp *TokenPlugin) HandleSettlementTask(settlementInfo *token.SettlementInfo) ([]byte, error) {
	// time.Sleep(20 * time.Second)
	client, err := ethclient.Dial(tkp.OpConfig.MainChain.Rpc)
	if err != nil || client == nil {
		log.Error("failed connect operator node", "err", err)
	}
	// 发送交易
	// 构建交易参数
	//priKey := tkp.subOpPriKey
	priKey := "51b50bc613d2479f1c4bf1447df03c5d64308734567ef4532a0ca5457660c6b7"
	// 调用主链系统合约结算接口
	toAddr := tkp.OpConfig.MainChain.SysAddr
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("failed connect operator node", "err", err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("failed connect operator node", "err", "Could not get public key")
		return nil, errors.New("could not get public key")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 以运营节点的质押地址作为运营节点地址，判断是否是子链运营节点签名交易
	if fromAddr != tkp.OpConfig.SubChain.OpAddr {
		log.Error("failed connect operator node", "err", err)
		return nil, err
	}
	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		log.Error("failed connect operator node", "err", err)
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("failed connect operator node", "err", err)
		return nil, err
	}
	value := big.NewInt(0)
	gasLimit := uint64(300000)
	// 组装结算接口的data
	// data := []byte("")
	// 主链结算接口函数Code
	settleFuncType := 6002
	toAddr = vm.TokenContractAddr
	data := genSettleTxRlpData(uint16(settleFuncType), settlementInfo)

	// 创建交易对象
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, data)

	// 使用发送方的私钥进行交易签名
	chainID, err := client.ChainID(context.Background())
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("failed connect operator node", "err", err)
		return nil, err
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("failed connect operator node", "err", err)
		return nil, err
	}

	hash := signedTx.Hash()
	log.Error("hash=========================================", hash.Hex())
	return hash.Bytes(), nil
}

// SetMainOpAddr Set the main chain operator address
func (tkp *TokenPlugin) SetMainOpAddr(mainOpAddr common.Address) {
	tkp.MainOpAddr = mainOpAddr
}

// SetOpConfig Set the main chain operator address
func (tkp *TokenPlugin) SetOpConfig(opConfig *params.OpConfig) {
	tkp.OpConfig = opConfig
}

// SetSubOpIdentity Set the sub-chain operation node identity
func (tkp *TokenPlugin) SetSubOpIdentity(isSubOpNode bool) {
	tkp.IsSubOpNode = isSubOpNode
}

// ExistAccount Add a list of minting account information
func (tkp *TokenPlugin) ExistAccount(state xcom.StateDB, mintAcc common.Address) bool {
	return false
}

// AddMintAccInfo Add a list of minting account information
func (tkp *TokenPlugin) AddMintAccInfo(state xcom.StateDB, mintAccInfo token.MintAccInfo) error {
	return token.SaveMintInfo(state, mintAccInfo)
}

// GetMintAccInfo Get a list of minting account information
func (tkp *TokenPlugin) GetMintAccInfo(state xcom.StateDB) (*token.MintAccInfo, error) {
	return token.GetMintAccInfo(state)
}

// BeginBlock implement BasePlugin
func (tkp *TokenPlugin) BeginBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	return nil
}

// EndBlock implement BasePlugin
func (tkp *TokenPlugin) EndBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	return nil
}

// Confirmed implement BasePlugin:does nothing
func (tkp *TokenPlugin) Confirmed(nodeId discover.NodeID, block *types.Block) error {
	return nil
}
