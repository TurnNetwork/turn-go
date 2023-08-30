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

const (
	MainChainSysAddr = "0x2000000000000000000000000000000000000002" // Main chain system contract address
)

var (
	tokenPluginOnce sync.Once
	tkp             *TokenPlugin
)

type TokenPlugin struct {
	IsSubOpNode bool // Whether it is a child chain operator node
	OpConfig    *params.OpConfig
	ChainID     *big.Int
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

// PostSettlementTask post the checkout task to the subscription event
func (tkp *TokenPlugin) PostSettlementTask(settleTask *token.SettleTask) error {
	if err := tkp.eventMux.Post(*settleTask); nil != err {
		log.Error("post settlementInfo failed", "err", err)
		return err
	}
	return nil
}

// Generate the rlp encoding of the main-chain settlement transaction
func genSettleTxRlpData(settleTask *token.SettleTask) []byte {
	var params [][]byte
	params = make([][]byte, 0)
	// Settlement function encoding
	settleFuncType := uint16(5)
	fnType, _ := rlp.EncodeToBytes(settleFuncType)
	params = append(params, fnType)

	txHash, _ := rlp.EncodeToBytes(settleTask.TxHash)
	bubId, _ := rlp.EncodeToBytes(settleTask.BubbleID)
	accAsset, _ := rlp.EncodeToBytes(settleTask.SettleInfo)
	params = append(params, txHash)
	params = append(params, bubId)
	params = append(params, accAsset)
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, params)
	if err != nil {
		// panic(fmt.Errorf("%d encode rlp data fail: %v", settleFuncType, err))
		return nil
	} else {
		rlpData := hexutil.Encode(buf.Bytes())
		fmt.Printf("funcType:%d rlp data = %s\n", settleFuncType, rlpData)
		return buf.Bytes()
	}
	return nil
}

//HandleSettleTask Handle settlement tasks
func (tkp *TokenPlugin) HandleSettleTask(settleTask *token.SettleTask) ([]byte, error) {
	if nil == settleTask {
		return nil, errors.New("the data in the settlement task is empty")
	}
	client, err := ethclient.Dial(tkp.OpConfig.MainChain.Rpc)
	if err != nil || client == nil {
		log.Error("failed connect operator node", "err", err)
		return nil, errors.New("failed connect operator node")
	}
	// Construct transaction parameters
	priKey := tkp.OpConfig.GetSubOpPriKey()
	// Invokes the main-chain system contract settlement interface
	toAddr := common.HexToAddress(MainChainSysAddr)
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
	data := genSettleTxRlpData(settleTask)
	if nil == data {
		log.Error("failed to generate the Code of the main chain settlement interface function", "err", err)
		return nil, err
	}
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

// SetChainID Set bubble's chainId
func (tkp *TokenPlugin) SetChainID(chainId *big.Int) {
	tkp.ChainID = chainId
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
	// Determine whether it is necessary to store minting account information
	if 0 == len(mintAccInfo.AccList) && 0 == len(mintAccInfo.TokenAddrList) {
		return nil
	}
	// Filter information about stored coins based on the address of the minting account
	var newAccInfo token.MintAccInfo
	for _, acc := range mintAccInfo.AccList {
		// Determine if the external account address is already stored
		if !token.ExistAccount(blockHash, acc) {
			newAccInfo.AccList = append(newAccInfo.AccList, acc)
			// Store the new minting account address
			token.StoreAccount(blockHash, acc)
		}
	}

	for _, tokenAddr := range mintAccInfo.TokenAddrList {
		// Determine if the contract account address is already stored
		if !token.ExistAccount(blockHash, tokenAddr) {
			newAccInfo.TokenAddrList = append(newAccInfo.TokenAddrList, tokenAddr)
			// Store the new contract address in the mintage account
			token.StoreAccount(blockHash, tokenAddr)
		}
	}

	// Whether there is a new mintage address to be saved
	if 0 < len(newAccInfo.AccList) || 0 < len(newAccInfo.TokenAddrList) {
		// Get the original minting account information
		oldMintAccInfo, err := tkp.GetMintAccInfo(blockHash)
		if err != nil {
			return err
		}

		var saveMintAccInfo token.MintAccInfo
		if nil != oldMintAccInfo && oldMintAccInfo.AccList != nil {
			saveMintAccInfo.AccList = oldMintAccInfo.AccList
		}
		if nil != oldMintAccInfo && oldMintAccInfo.TokenAddrList != nil {
			saveMintAccInfo.TokenAddrList = oldMintAccInfo.TokenAddrList
		}

		for _, acc := range newAccInfo.AccList {
			saveMintAccInfo.AccList = append(saveMintAccInfo.AccList, acc)
		}

		for _, tokenAddr := range newAccInfo.TokenAddrList {
			saveMintAccInfo.TokenAddrList = append(saveMintAccInfo.TokenAddrList, tokenAddr)
		}
		// Keep up-to-date minting account information
		return token.StoreMintInfo(blockHash, saveMintAccInfo)
	}
	return nil
}

// GetMintAccInfo Get a list of minting account information
func (tkp *TokenPlugin) GetMintAccInfo(blockHash common.Hash) (*token.MintAccInfo, error) {
	return token.GetMintAccInfo(blockHash)
}

// StoreL1HashToL2Hash The mapping relationship between the main-chain transaction hash and the sub-chain transaction hash is stored
// When the pledged token transaction of the main chain and the minting transaction of the child chain are mapped
func (tkp *TokenPlugin) StoreL1HashToL2Hash(blockHash common.Hash, L1TxHash common.Hash, L2TxHash common.Hash) error {
	return token.StoreL1HashToL2Hash(blockHash, L1TxHash, L2TxHash)
}

// GetL2HashByL1Hash The mintage transaction hash of the child chain is obtained according to the pledge token transaction hash of the main chain
func (tkp *TokenPlugin) GetL2HashByL1Hash(blockHash common.Hash, L2TxHash common.Hash) (*common.Hash, error) {
	return token.GetL2HashByL1Hash(blockHash, L2TxHash)
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
