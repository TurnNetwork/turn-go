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

func (tkp *TokenPlugin) SetEventMux(eventMux *event.TypeMux) {
	tkp.eventMux = eventMux
}

// AddSettlementTask Add the checkout task to the subscription event
func (tkp *TokenPlugin) AddSettlementTask(settlementInfo *token.SettlementInfo) {
	if err := tkp.eventMux.Post(*settlementInfo); nil != err {
		log.Error("post settlementInfo failed", "err", err)
	}
}

// 生成主链结算交易的rlp编码
func genSettleTxRlpData(settlementInfo *token.SettlementInfo) []byte {
	var params [][]byte
	params = make([][]byte, 0)
	// 结算函数编码
	settleFuncType := uint16(6002)
	fnType, _ := rlp.EncodeToBytes(settleFuncType)
	params = append(params, fnType)

	accAsset, _ := rlp.EncodeToBytes(settlementInfo)
	params = append(params, accAsset)
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, params)
	if err != nil {
		panic(fmt.Errorf("%d encode rlp data fail: %v", settleFuncType, err))
	} else {
		rlpData := hexutil.Encode(buf.Bytes())
		fmt.Printf("funcType:%d rlp data = %s\n", settleFuncType, rlpData)
		return buf.Bytes()
	}
	return nil
}

//HandleSettlementTask Handle settlement tasks
func (tkp *TokenPlugin) HandleSettlementTask(settlementInfo *token.SettlementInfo) ([]byte, error) {
	client, err := ethclient.Dial(tkp.OpConfig.MainChain.Rpc)
	if err != nil || client == nil {
		log.Error("failed connect operator node", "err", err)
		return nil, errors.New("failed connect operator node")
	}
	// Construct transaction parameters
	priKey := tkp.OpConfig.GetSubOpPriKey()
	// Invokes the main-chain system contract settlement interface
	toAddr := tkp.OpConfig.MainChain.SysAddr
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("Wrong private key", "err", err)
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("the private key to the public key failed")
		return nil, errors.New("the private key to the public key failed")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	// The pledge address of the operating node is taken as the address of the operating node
	// Check whether the transaction is signed by the sub-chain operating node
	if fromAddr != tkp.OpConfig.SubChain.OpAddr {
		log.Error("The settlement transaction sender is not the sub-chain operation address")
		return nil, errors.New("the settlement transaction sender is not the sub-chain operation address")
	}
	// get account nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		log.Error("Failed to obtain the account nonce", "err", err)
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Failed to get gasprice", "err", err)
		return nil, err
	}
	value := big.NewInt(0)
	gasLimit := uint64(300000)
	// Generate the Code of the main chain settlement interface function
	data := genSettleTxRlpData(settlementInfo)
	// Creating transaction objects
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, data)

	// The transaction is signed using the sender's private key
	chainID, err := client.ChainID(context.Background())
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("Signing settlement transaction failed", "err", err)
		return nil, err
	}

	// Send the settlement transaction to the main chain
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Failed to send settlement transaction", "err", err)
		return nil, err
	}

	hash := signedTx.Hash()
	log.Debug("settlement tx hash=========================================", hash.Hex())
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

// AddMintAccInfo Add a list of minting account information
func (tkp *TokenPlugin) AddMintAccInfo(blockHash common.Hash, mintAccInfo token.MintAccInfo) error {
	return token.SaveMintInfo(blockHash, mintAccInfo)
}

// GetMintAccInfo Get a list of minting account information
func (tkp *TokenPlugin) GetMintAccInfo(blockHash common.Hash) (*token.MintAccInfo, error) {
	return token.GetMintAccInfo(blockHash)
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
