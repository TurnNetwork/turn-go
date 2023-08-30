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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/common/mock"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/plugin"
	"github.com/bubblenet/bubble/x/token"
	"github.com/bubblenet/bubble/x/xcom"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	testChainId          = big.NewInt(1)
	sndb                 = snapshotdb.Instance()
	genesisBlockHash     = common.HexToHash("")
	storeBlockHash       = common.HexToHash("")
	ERC20CodeTmp         = "608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063318f44481161008c57806370a082311161006657806370a08231146101a757806395d89b41146101d0578063a9059cbb146101d8578063dd62ed3e146101eb57600080fd5b8063318f44481461016f5780633bc807d01461017f57806340c10f191461019457600080fd5b806306b68323146100d457806306fdde03146100fd578063095ea7b31461011257806318160ddd1461013557806323b872dd14610147578063313ce5671461015a575b600080fd5b6100e76100e23660046107fd565b610224565b6040516100f49190610872565b60405180910390f35b6101056102fd565b6040516100f491906108b6565b610125610120366004610920565b61038f565b60405190151581526020016100f4565b6003545b6040519081526020016100f4565b61012561015536600461094a565b610433565b60025460405160ff90911681526020016100f4565b600254610100900460ff16610125565b61019261018d366004610986565b6105cf565b005b6101926101a2366004610920565b61062f565b6101396101b53660046109b0565b6001600160a01b031660009081526004602052604090205490565b61010561072c565b6101256101e6366004610920565b61073b565b6101396101f93660046109cb565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205490565b606060008267ffffffffffffffff811115610241576102416109fe565b60405190808252806020026020018201604052801561026a578160200160208202803683370190505b50905060005b838110156102f3576004600086868481811061028e5761028e610a14565b90506020020160208101906102a391906109b0565b6001600160a01b03166001600160a01b03168152602001908152602001600020548282815181106102d6576102d6610a14565b6020908102919091010152806102eb81610a40565b915050610270565b5090505b92915050565b60606000805461030c90610a59565b80601f016020809104026020016040519081016040528092919081815260200182805461033890610a59565b80156103855780601f1061035a57610100808354040283529160200191610385565b820191906000526020600020905b81548152906001019060200180831161036857829003601f168201915b5050505050905090565b3360008181526004602052604081205490919083908111156103cc5760405162461bcd60e51b81526004016103c390610a93565b60405180910390fd5b3360008181526005602090815260408083206001600160a01b038a1680855290835292819020889055518781529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3506001949350505050565b6001600160a01b0383166000908152600460205260408120548490839081111561046f5760405162461bcd60e51b81526004016103c390610a93565b6001600160a01b03861660009081526005602090815260408083203384529091529020548411156104e25760405162461bcd60e51b815260206004820152601d60248201527f54686520616d6f756e7420616c6c6f77656420746f206265207573656400000060448201526064016103c3565b6001600160a01b0386166000908152600460205260408120805486929061050a908490610aca565b90915550506001600160a01b03851660009081526004602052604081208054869290610537908490610add565b90915550506001600160a01b03861660009081526005602090815260408083203384529091528120805486929061056f908490610aca565b92505081905550846001600160a01b0316866001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef866040516105bb91815260200190565b60405180910390a350600195945050505050565b600254610100900460ff16156106155760405162461bcd60e51b815260206004820152600b60248201526a24ba1034b99024b734ba1760a91b60448201526064016103c3565b6002805461ffff191660ff90921691909117610100179055565b3360206001609c1b01146106a15760405162461bcd60e51b815260206004820152603360248201527f4e6f74206d6963726f2d6e6f64652073797374656d20636f6e7472616374206160448201527219191c995cdcc818d85b881b9bdd081b5a5b9d606a1b60648201526084016103c3565b6001600160a01b038216600090815260046020526040812080548392906106c9908490610add565b9250508190555080600360008282546106e29190610add565b90915550506040518181526001600160a01b038316906000907fab8530f87dc9b59234c4623bf917212bb2536d647574c8e7e5da92c2ede0c9f89060200160405180910390a35050565b60606001805461030c90610a59565b33600081815260046020526040812054909190839081111561076f5760405162461bcd60e51b81526004016103c390610a93565b336000908152600460205260408120805486929061078e908490610aca565b90915550506001600160a01b038516600090815260046020526040812080548692906107bb908490610add565b90915550506040518481526001600160a01b0386169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610420565b6000806020838503121561081057600080fd5b823567ffffffffffffffff8082111561082857600080fd5b818501915085601f83011261083c57600080fd5b81358181111561084b57600080fd5b8660208260051b850101111561086057600080fd5b60209290920196919550909350505050565b6020808252825182820181905260009190848201906040850190845b818110156108aa5783518352928401929184019160010161088e565b50909695505050505050565b600060208083528351808285015260005b818110156108e3578581018301518582016040015282016108c7565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461091b57600080fd5b919050565b6000806040838503121561093357600080fd5b61093c83610904565b946020939093013593505050565b60008060006060848603121561095f57600080fd5b61096884610904565b925061097660208501610904565b9150604084013590509250925092565b60006020828403121561099857600080fd5b813560ff811681146109a957600080fd5b9392505050565b6000602082840312156109c257600080fd5b6109a982610904565b600080604083850312156109de57600080fd5b6109e783610904565b91506109f560208401610904565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610a5257610a52610a2a565b5060010190565b600181811c90821680610a6d57607f821691505b602082108103610a8d57634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526018908201527f48617665206e6f7420656e6f7567682062616c616e63652e0000000000000000604082015260600190565b818103818111156102f7576102f7610a2a565b808201808211156102f7576102f7610a2a56fea2646970667358221220ab69fa229cb08ce4e2f3628266da3cfd1d5f7edb77e729acf151519dd3161dbb64736f6c63430008110033"
	L1StakingTokenTxHash = common.HexToHash("0x12c171900f010b17e969702efa044d077e86808212c171900f010b17e969702e")
	testERC20AddrList    = []common.Address{
		common.HexToAddress("0xe200000000000000000000000000000000000000"),
		common.HexToAddress("0xe200000000000000000000000000000000000001"),
		common.HexToAddress("0xe200000000000000000000000000000000000002"),
	}
	testTokenAmount    = big.NewInt(100000000000)
	testNativeAmount   = big.NewInt(200000000000)
	testStep           = big.NewInt(30000000000) // Use a multiple of settlement
	settleTokenAmount  = new(big.Int).Sub(testTokenAmount, testStep)
	settleNativeAmount = new(big.Int).Sub(testNativeAmount, testStep)
	testAddrList       = []common.Address{
		common.HexToAddress("0xfffff00000000000000000000000000000000000"),
		common.HexToAddress("0xfffff00000000000000000000000000000000001"),
	}
)

func runBubbleTx(bubContract *TokenContract, params [][]byte, title string, t *testing.T) {

	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, params)
	if err != nil {
		t.Errorf(title+" encode rlp data fail: %v", err)
	} else {
		t.Log(title+" data rlp: ", hexutil.Encode(buf.Bytes()))
	}

	res, err := bubContract.Run(buf.Bytes())
	assert.True(t, nil == err)
	var r uint32
	err = json.Unmarshal(res, &r)
	assert.True(t, nil == err)
	assert.Equal(t, common.OkCode, r)
}

func runBubbleCall(tkContract *TokenContract, params [][]byte, title string, verifyRet interface{}, t *testing.T) {
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, params)
	if err != nil {
		t.Errorf("%s encode rlp data fail: %v", title, err)
		return
	} else {
		t.Logf("%s data rlp: %s", title, hexutil.Encode(buf.Bytes()))
	}

	res, err := tkContract.Run(buf.Bytes())
	if nil != err {
		t.Errorf("err: %v", err)
	}

	assert.True(t, nil == err)
	var r xcom.Result
	err = json.Unmarshal(res, &r)
	assert.True(t, nil == err)
	assert.Equal(t, common.OkCode, r.Code)
	if nil != verifyRet {
		assert.Equal(t, verifyRet, r.Ret)
	}
	t.Logf("%s the result: %v\n", title, string(res))
}

func newOneBlock(parentHash common.Hash, blockNumber *uint64) common.Hash {
	header := types.Header{
		Number: big.NewInt(int64(*blockNumber)),
	}
	newBlockHash := header.Hash()

	err := sndb.NewBlock(header.Number, parentHash, newBlockHash)
	if err != nil {
		fmt.Errorf("newBlock, %v", err)
	}
	*blockNumber++
	lastBlockHash = newBlockHash
	return newBlockHash
}

// newTokenContract
func newTokenContract(caller common.Address, blockNumber *big.Int, blockHash common.Hash, chain *mock.Chain, index int) *TokenContract {
	blockContext := BlockContext{
		BlockNumber: blockNumber,
		BlockHash:   blockHash,
		Ctx:         context.Background(),
	}

	contract := &TokenContract{
		Plugin:   plugin.TokenInstance(),
		Contract: newContract(common.Big0, caller),
		Evm:      NewEVM(blockContext, TxContext{}, chain.SnapDB, chain.StateDB, params.TestChainConfig, Config{}),
	}
	// set ChainID
	contract.Plugin.ChainID = testChainId
	// set operator config
	var opConfig params.OpConfig
	opConfig.MainChain = new(params.OperatorInfo)
	opConfig.MainChain.OpAddr = sender
	contract.Plugin.SetOpConfig(&opConfig)

	chain.StateDB.Prepare(txHashArr[index], blockHash, index+1)
	return contract
}

func init_erc20_data(tkContract *TokenContract, block *uint64) {

	no := int64(*block)
	header := types.Header{
		Number: big.NewInt(no),
	}
	newBlockHash := header.Hash()
	// bubDB := bubble.NewBubbleDB()
	err := sndb.NewBlock(header.Number, lastBlockHash, newBlockHash)
	if err != nil {
		fmt.Errorf("newBlock, %v", err)
	}
	// MOCK
	// Deploy the erc20 contract template
	evm := tkContract.Evm
	if len(evm.StateDB.GetCode(common.HexToAddress(TmpERC20Addr))) == 0 {
		code := common.Hex2Bytes(ERC20CodeTmp)
		// deploy erc20 contract
		evm.StateDB.SetCode(common.HexToAddress(TmpERC20Addr), code)
	}

	lastBlockHash = newBlockHash
	lastBlockNumber = *block
	lastHeader = header
	storeBlockHash = newBlockHash
	*block++
}

func getBalances(evm *EVM, contract *Contract, accList []common.Address) (*[]token.AccountAsset, error) {
	// getBalances
	// Assembly settlement information
	var accAssets []token.AccountAsset
	for _, acc := range accList {
		accAssets = append(accAssets, token.AccountAsset{Account: acc, NativeAmount: evm.StateDB.GetBalance(acc)})
	}
	for _, tokenAddr := range testERC20AddrList {
		var tokenAssets []token.AccTokenAsset
		code := evm.StateDB.GetCode(tokenAddr)
		if len(code) > 0 {
			// Change to ERC20 contract address
			contract.self = AccountRef(tokenAddr)
			contract.SetCallCode(&tokenAddr, evm.StateDB.GetCodeHash(tokenAddr), code)
			// Batch queries for erc20 token balances in the list of accounts
			input, err := encodeGetBalancesCall(testAddrList)
			if err != nil {
				fmt.Errorf("failed to get Address ERC20 Token, error:%v", err)
				return nil, err
			}
			// Execute EVM
			ret, err := RunEvm(evm, contract, input)
			if err != nil {
				fmt.Errorf("failed to get Address ERC20 Token, error:%v", err)
				return nil, err
			}
			// Parse byte array to uint256 array,
			// the first 32 bytes value indicates how many bytes to store the length of the array,
			// fixed as: 32
			resList := parseBytesToUint256Array(ret[32:])

			if len(resList) > 0 {
				// The value of the first element indicates the length of the returned array
				elemLen := resList[0].Uint64()
				if elemLen != uint64(len(testAddrList)) {
					fmt.Errorf("failed to get Address ERC20 Token, error: %s",
						"The length of the number of accounts and the number of balances retrieved are inconsistent")
					return nil, nil
				}
				// Assemble the ERC20 Token settlement information
				for iAcc, balance := range resList[1:] {
					var accTokenAsset token.AccTokenAsset
					accTokenAsset.TokenAddr = tokenAddr
					accTokenAsset.Balance = balance
					tokenAssets = append(tokenAssets, accTokenAsset)
					accAssets[iAcc].TokenAssets = append(accAssets[iAcc].TokenAssets, accTokenAsset)
				}
			}
		}
	}
	return &accAssets, nil
}

// verify token count
func verify_token_amount(contract *TokenContract, nativeAmount *big.Int, tokenAmount *big.Int, t *testing.T) {
	// query balances
	accAssets, err := getBalances(contract.Evm, contract.Contract, testAddrList)
	if accAssets == nil || err != nil {
		panic(fmt.Errorf("failed to getBalances: %v", err.Error()))
	}
	assert.True(t, nil == err)
	assert.True(t, nil != accAssets)
	// Compare the number of native tokens and erc20 tokens of sender after initialization
	for _, accAsset := range *accAssets {
		result := nativeAmount.Cmp(accAsset.NativeAmount)
		assert.Equal(t, 0, result, "Native token count validation failed")
		for _, tokenAsset := range accAsset.TokenAssets {
			result := tokenAmount.Cmp(tokenAsset.Balance)
			assert.Equal(t, 0, result, "Erc20 token count verification failed")
		}
	}
}

// test mintToken interface
func mint_token(contract *TokenContract, mintAccount common.Address, t *testing.T) {

	var params [][]byte
	params = make([][]byte, 0)

	mintAsset := token.AccountAsset{
		Account:      mintAccount,
		NativeAmount: testNativeAmount,
	}
	for _, tokenAddr := range testERC20AddrList {
		tokenAsset := token.AccTokenAsset{
			TokenAddr: tokenAddr,
			Balance:   testTokenAmount,
		}
		mintAsset.TokenAssets = append(mintAsset.TokenAssets, tokenAsset)
	}

	fnType, _ := rlp.EncodeToBytes(uint16(TxMintToken))
	L1TxHash, _ := rlp.EncodeToBytes(L1StakingTokenTxHash)
	accAsset, _ := rlp.EncodeToBytes(mintAsset)

	params = append(params, fnType)
	params = append(params, L1TxHash)
	params = append(params, accAsset)

	runBubbleTx(contract, params, "mintToken", t)
}

// test settleBubble interface
func settle_bubble(contract *TokenContract, t *testing.T) {
	var params [][]byte
	params = make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(uint16(TxSettleBubble))
	params = append(params, fnType)
	runBubbleTx(contract, params, "settleBubble", t)
}

func decode_and_verify_mint_tx_receipt(tkContract *TokenContract, mintAccount common.Address, txIndex int, t *testing.T) {
	logs := tkContract.Evm.StateDB.GetLogs(txHashArr[txIndex])
	for _, log := range logs {
		// Deal only with switchable viewer is empty,
		// mintToken transaction logs will be written to erc20 COINS event information, need to filter
		if nil != (*log).Topics || (*log).Address != vm.TokenContractAddr {
			continue
		}
		data := (*log).Data
		// t.Logf("mintToken tx logs: %v", data)
		var L1StTxHash common.Hash
		var accAsset token.AccountAsset
		var m [][]byte
		if err := rlp.DecodeBytes(data, &m); err != nil {
			t.Error(err)
		}
		var code string
		err := rlp.DecodeBytes(m[0], &code)
		assert.True(t, nil == err)
		assert.True(t, code == "0")

		err = rlp.DecodeBytes(m[1], &L1StTxHash)
		assert.True(t, nil == err)
		assert.True(t, L1StTxHash == L1StakingTokenTxHash, "Error in obtaining the main chain pledge token transaction hash\n\n")
		// fmt.Printf("L1StakingTokenTxHash: %v\n", L1StTxHash)

		if err := rlp.DecodeBytes(m[2], &accAsset); err != nil {
			t.Error(err)
		}
		assert.True(t, mintAccount == accAsset.Account, "mintToken account error, and the test account is inconsistent")
		result := testNativeAmount.Cmp(accAsset.NativeAmount)
		assert.Equal(t, 0, result, "Incorrect mint native token count, inconsistent with test native token count")
		assert.True(t, len(testERC20AddrList) == len(accAsset.TokenAssets), "Inconsistent length of mint token accounts")
		for i, tokenAddr := range testERC20AddrList {
			tokenAsset := accAsset.TokenAssets[i]
			assert.True(t, tokenAddr == tokenAsset.TokenAddr, "Inconsistent erc20 token address")
			result := testTokenAmount.Cmp(tokenAsset.Balance)
			assert.Equal(t, 0, result, "Inconsistent erc20 token amount")
		}
		// fmt.Printf("MintToken AccoutAsset: %v\n", accAsset)
	}
}

func decode_and_verify_settle_tx_receipt(tkContract *TokenContract, txIndex int, t *testing.T) {
	logs := tkContract.Evm.StateDB.GetLogs(txHashArr[txIndex])
	for _, log := range logs {
		if nil != (*log).Topics || (*log).Address != vm.TokenContractAddr {
			continue
		}
		data := (*log).Data
		// t.Logf("settleBubble tx logs: %v", data)
		var settlementInfo token.SettlementInfo
		var m [][]byte
		if err := rlp.DecodeBytes(data, &m); err != nil {
			t.Error(err)
		}
		var code string
		if err := rlp.DecodeBytes(m[0], &code); err != nil {
			t.Error(err)
		}
		err := rlp.DecodeBytes(m[1], &settlementInfo)
		if err != nil {
			t.Error(err)
		}
		assert.True(t, nil == err)
		assert.True(t, len(testAddrList) == len(settlementInfo.AccAssets), "Inconsistent length of settlement accounts")
		for i, addr := range testAddrList {
			accAsset := settlementInfo.AccAssets[i]
			assert.True(t, addr == accAsset.Account, "Settlement account error, and the test account is inconsistent")
			result := testNativeAmount.Cmp(accAsset.NativeAmount)
			assert.Equal(t, 0, result, "Incorrect settlement native token count, inconsistent with test native token count")
			assert.True(t, len(testERC20AddrList) == len(accAsset.TokenAssets), "Inconsistent length of settlement token accounts")
			for j, tokenAddr := range testERC20AddrList {
				tokenAsset := accAsset.TokenAssets[j]
				assert.True(t, tokenAddr == tokenAsset.TokenAddr, "Inconsistent erc20 token address")
				result := testTokenAmount.Cmp(tokenAsset.Balance)
				assert.Equal(t, 0, result, "Inconsistent erc20 token amount")
			}
		}
		// t.Logf("settlementInfo: %v\n", settlementInfo)
	}
}

// test getL2HashByL1Hash interface
func getL2TxHashByL1TxHash(contract *TokenContract, VerifyL2TxHash common.Hash, t *testing.T) {

	var params [][]byte
	params = make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(uint16(CallGetL2HashByL1Hash))
	txHash, _ := rlp.EncodeToBytes(L1StakingTokenTxHash)

	params = append(params, fnType)
	params = append(params, txHash)
	runBubbleCall(contract, params, "getL1HashByL2Hash", VerifyL2TxHash.Hex(), t)
}

func TestTokenContract_mintToken(t *testing.T) {
	chain := newMockChain()
	defer chain.SnapDB.Clear()

	genesisBlockHash = chain.Genesis.Hash()
	t.Logf("genesisHash: %v", genesisBlockHash)
	// first block
	blkNumber := uint64(1)
	storeBlockHash = newOneBlock(genesisBlockHash, &blkNumber)

	tkContract := newTokenContract(sender, blockNumber, storeBlockHash, chain, 1)
	// init erc20 data
	init_erc20_data(tkContract, &blkNumber)

	// before mintToken: Determine the number of native and erc20 tokens in the minting account
	verify_token_amount(tkContract, big0, big0, t)

	index := 2
	// Cycle call mintToken system contracts interface
	for _, mintAccount := range testAddrList {
		tkContract.Contract = newContract(common.Big0, sender)
		chain.StateDB.Prepare(txHashArr[index], blockHash, index+1)
		mint_token(tkContract, mintAccount, t)
		// decode and verify mintToken transaction receipt
		decode_and_verify_mint_tx_receipt(tkContract, mintAccount, index, t)
		index++
	}

	// after mintToken: Determine the number of native and erc20 tokens in the minting account
	verify_token_amount(tkContract, testNativeAmount, testTokenAmount, t)
}

func TestTokenContract_settleBubble(t *testing.T) {
	chain := newMockChain()
	defer chain.SnapDB.Clear()

	genesisBlockHash = chain.Genesis.Hash()
	t.Logf("genesisHash: %v", genesisBlockHash)
	// first block
	blkNumber := uint64(1)
	storeBlockHash = newOneBlock(genesisBlockHash, &blkNumber)

	tkContract := newTokenContract(sender, blockNumber, storeBlockHash, chain, 1)
	// init erc20 data
	init_erc20_data(tkContract, &blkNumber)

	// settleBubble
	// Cycle call mintToken system contracts interface
	for _, mintAccount := range testAddrList {
		tkContract.Contract = newContract(common.Big0, sender)
		mint_token(tkContract, mintAccount, t)
	}
	index := 2
	t.Logf("settleBubble tx hash: %v", txHashArr[index])
	settle_bubble(tkContract, t)

	// decode and verify settleBubble transaction receipt
	decode_and_verify_settle_tx_receipt(tkContract, index, t)
}

func TestTokenContract_getL2HashByL1Hash(t *testing.T) {
	chain := newMockChain()
	defer chain.SnapDB.Clear()

	genesisBlockHash = chain.Genesis.Hash()
	t.Logf("genesisHash: %v", genesisBlockHash)
	// first block
	blkNumber := uint64(1)
	storeBlockHash = newOneBlock(genesisBlockHash, &blkNumber)

	tkContract := newTokenContract(sender, blockNumber, storeBlockHash, chain, 1)
	// init erc20 data
	init_erc20_data(tkContract, &blkNumber)

	index := 2
	chain.StateDB.Prepare(txHashArr[index], blockHash, index+1)
	tkContract.Contract = newContract(common.Big0, sender)
	mint_token(tkContract, testAddrList[0], t)

	getL2TxHashByL1TxHash(tkContract, tkContract.Evm.StateDB.TxHash(), t)
}
