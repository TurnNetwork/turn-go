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
	_ "fmt"
	"github.com/bubblenet/bubble/accounts/abi"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/bubble"
	"github.com/bubblenet/bubble/x/xcom"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bubblenet/bubble/common/mock"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/plugin"
)

var (
	testBubbleId     = big.NewInt(1)
	sndb             = snapshotdb.Instance()
	genesisBlockHash = common.HexToHash("")
	storeBlockHash   = common.HexToHash("")
	BlockHashList    = []common.Hash{
		common.HexToHash(""), // 0
		common.HexToHash("9d4fb5346abcf593ad80a0d3d5a371b22c962418ad34189d5b1b39065668d663"), //1
	}
	// ERC20 config
	erc20Code         = "608060405234801561001057600080fd5b50600436106100a95760003560e01c8063313ce56711610071578063313ce567146101345780633c333cea1461014957806370a082311461015e57806395d89b4114610187578063a9059cbb1461018f578063dd62ed3e146101a257600080fd5b806306b68323146100ae57806306fdde03146100d7578063095ea7b3146100ec57806318160ddd1461010f57806323b872dd14610121575b600080fd5b6100c16100bc366004610724565b6101db565b6040516100ce9190610799565b60405180910390f35b6100df6102b4565b6040516100ce91906107dd565b6100ff6100fa366004610847565b610346565b60405190151581526020016100ce565b6003545b6040519081526020016100ce565b6100ff61012f366004610871565b6103ea565b60025460405160ff90911681526020016100ce565b61015c6101573660046108c3565b610586565b005b61011361016c366004610994565b6001600160a01b031660009081526004602052604090205490565b6100df610653565b6100ff61019d366004610847565b610662565b6101136101b03660046109b6565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205490565b606060008267ffffffffffffffff8111156101f8576101f86108ad565b604051908082528060200260200182016040528015610221578160200160208202803683370190505b50905060005b838110156102aa5760046000868684818110610245576102456109e9565b905060200201602081019061025a9190610994565b6001600160a01b03166001600160a01b031681526020019081526020016000205482828151811061028d5761028d6109e9565b6020908102919091010152806102a281610a15565b915050610227565b5090505b92915050565b6060600080546102c390610a2e565b80601f01602080910402602001604051908101604052809291908181526020018280546102ef90610a2e565b801561033c5780601f106103115761010080835404028352916020019161033c565b820191906000526020600020905b81548152906001019060200180831161031f57829003601f168201915b5050505050905090565b3360008181526004602052604081205490919083908111156103835760405162461bcd60e51b815260040161037a90610a68565b60405180910390fd5b3360008181526005602090815260408083206001600160a01b038a1680855290835292819020889055518781529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3506001949350505050565b6001600160a01b038316600090815260046020526040812054849083908111156104265760405162461bcd60e51b815260040161037a90610a68565b6001600160a01b03861660009081526005602090815260408083203384529091529020548411156104995760405162461bcd60e51b815260206004820152601d60248201527f54686520616d6f756e7420616c6c6f77656420746f2062652075736564000000604482015260640161037a565b6001600160a01b038616600090815260046020526040812080548692906104c1908490610a9f565b90915550506001600160a01b038516600090815260046020526040812080548692906104ee908490610ab2565b90915550506001600160a01b038616600090815260056020908152604080832033845290915281208054869290610526908490610a9f565b92505081905550846001600160a01b0316866001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8660405161057291815260200190565b60405180910390a350600195945050505050565b60005b815181101561064e5782600460008484815181106105a9576105a96109e9565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020819055508181815181106105e7576105e76109e9565b60200260200101516001600160a01b031660006001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405161063491815260200190565b60405180910390a38061064681610a15565b915050610589565b505050565b6060600180546102c390610a2e565b3360008181526004602052604081205490919083908111156106965760405162461bcd60e51b815260040161037a90610a68565b33600090815260046020526040812080548692906106b5908490610a9f565b90915550506001600160a01b038516600090815260046020526040812080548692906106e2908490610ab2565b90915550506040518481526001600160a01b0386169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016103d7565b6000806020838503121561073757600080fd5b823567ffffffffffffffff8082111561074f57600080fd5b818501915085601f83011261076357600080fd5b81358181111561077257600080fd5b8660208260051b850101111561078757600080fd5b60209290920196919550909350505050565b6020808252825182820181905260009190848201906040850190845b818110156107d1578351835292840192918401916001016107b5565b50909695505050505050565b600060208083528351808285015260005b8181101561080a578581018301518582016040015282016107ee565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461084257600080fd5b919050565b6000806040838503121561085a57600080fd5b6108638361082b565b946020939093013593505050565b60008060006060848603121561088657600080fd5b61088f8461082b565b925061089d6020850161082b565b9150604084013590509250925092565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156108d657600080fd5b8235915060208084013567ffffffffffffffff808211156108f657600080fd5b818601915086601f83011261090a57600080fd5b81358181111561091c5761091c6108ad565b8060051b604051601f19603f83011681018181108582111715610941576109416108ad565b60405291825284820192508381018501918983111561095f57600080fd5b938501935b82851015610984576109758561082b565b84529385019392850192610964565b8096505050505050509250929050565b6000602082840312156109a657600080fd5b6109af8261082b565b9392505050565b600080604083850312156109c957600080fd5b6109d28361082b565b91506109e06020840161082b565b90509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610a2757610a276109ff565b5060010190565b600181811c90821680610a4257607f821691505b602082108103610a6257634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526018908201527f48617665206e6f7420656e6f7567682062616c616e63652e0000000000000000604082015260600190565b818103818111156102ae576102ae6109ff565b808201808211156102ae576102ae6109ff56fea26469706673582212208ae36bdab54f064c9dab7b22f6755ad349d754b18cfb8838a9c1809cb45dd02164736f6c63430008110033"
	testERC20AddrList = []common.Address{
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
		sender, // sender
		delegateSender,
	}
	L2SettleTxHash = common.HexToHash("0x12c171900f010b17e969702efa044d077e86808212c171900f010b17e969702e")
	currentNodeId  = discover.NodeID{0x1}
)

func runBubbleTx(bubContract *BubbleContract, params [][]byte, title string, t *testing.T) {

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

func runBubbleCall(bubContract *BubbleContract, params [][]byte, title string, verifyRet interface{}, t *testing.T) {
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, params)
	if err != nil {
		t.Errorf("%s encode rlp data fail: %v", title, err)
		return
	} else {
		t.Logf("%s data rlp: %s", title, hexutil.Encode(buf.Bytes()))
	}

	res, err := bubContract.Run(buf.Bytes())
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

// encodeGetBalancesCall: Generate function signatures for the ERC20 batch get account balance interface
func encodeGetBalancesCall(addrList []common.Address) ([]byte, error) {
	// Create a contract ABI resolver
	encodeABI, err := abi.JSON(strings.NewReader(`[
		{
			"inputs": [
				{
					"internalType": "address[]",
					"name": "_addrList",
					"type": "address[]"
				}
			],
			"outputs": [
				{
					"internalType": "uint256[]",
					"name": "",
					"type": "uint256[]"
				}
			],
			"name": "balanceOf",
			"type": "function"
		}
	]`))
	if err != nil {
		return nil, err
	}

	functionName := "balanceOf"
	// Encode function call data
	data, err := encodeABI.Pack(functionName, addrList)
	if err != nil {
		return nil, err
	}

	// Convert the encoded data to a hexadecimal string
	// encodedData := common.Bytes2Hex(data)
	// fmt.Println(encodedData)
	return data, nil
}

// encodeERC20InitFunc: Generate function signatures for ERC20 contract init transactions
func encodeERC20InitFunc(amount *big.Int, addrList []common.Address) ([]byte, error) {
	// Create a contract ABI resolver
	encodeABI, err := abi.JSON(strings.NewReader(`[
		{
			"inputs": [
				{
					"internalType": "uint256",
					"name": "_amount",
					"type": "uint256"
				},
				{
					"internalType": "address[]",
					"name": "addrList",
					"type": "address[]"
				}
			],
			"name": "init",
			"type": "function"
		}
	]`))
	if err != nil {
		return nil, err
	}

	functionName := "init"
	// Encode function call data
	data, err := encodeABI.Pack(functionName, amount, addrList)
	if err != nil {
		return nil, err
	}

	// Convert the encoded data to a hexadecimal string
	// encodedData := common.Bytes2Hex(data)
	// fmt.Println(encodedData)
	return data, nil
}

func getBalances(evm *EVM, contract *Contract, accList []common.Address) (*[]bubble.AccountAsset, error) {
	// getBalances
	// Assembly settlement information
	var accAssets []bubble.AccountAsset
	for _, acc := range accList {
		accAssets = append(accAssets, bubble.AccountAsset{Account: acc, NativeAmount: evm.StateDB.GetBalance(acc)})
	}
	for _, tokenAddr := range testERC20AddrList {
		var tokenAssets []bubble.AccTokenAsset
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
					var accTokenAsset bubble.AccTokenAsset
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

// Parse the byte array into a uint256 array
func parseBytesToUint256Array(bytes []byte) []*big.Int {
	const Uint256Size = 32 // The size of uint256 is 32 bytes
	var uint256Array []*big.Int

	for i := 0; i < len(bytes); i += Uint256Size {
		end := i + Uint256Size
		if end > len(bytes) {
			end = len(bytes)
		}
		slice := bytes[i:end]

		uint256 := new(big.Int).SetBytes(slice)
		uint256Array = append(uint256Array, uint256)
	}

	return uint256Array
}

func init_erc20_data(bubContract *BubbleContract, block uint64) {

	no := int64(block)
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
	// deploy and init erc20 contract
	contract := bubContract.Contract
	evm := bubContract.Evm
	for _, erc20Addr := range testERC20AddrList {
		code := common.Hex2Bytes(erc20Code)
		// deploy erc20 contract
		evm.StateDB.SetCode(erc20Addr, code)
		// init
		contract.self = AccountRef(erc20Addr)
		contract.SetCallCode(&erc20Addr, evm.StateDB.GetCodeHash(erc20Addr), code)
		// encode init function call
		input, err := encodeERC20InitFunc(testTokenAmount, testAddrList)
		if err != nil {
			fmt.Errorf("encodeERC20InitFunc, %v", err)
			return
		}
		_, err = RunEvm(evm, contract, input)
		if err != nil {
			fmt.Errorf("faild to init erc20, %v", err)
			return
		}
	}

	lastBlockHash = newBlockHash
	lastBlockNumber = block
	lastHeader = header
}

// build_bubble_data Building bubble data
func build_bubble_data(block uint64, parentHash common.Hash) {
	no := int64(block)
	header := types.Header{
		Number: big.NewInt(no),
	}
	newBlockHash := header.Hash()
	bubDB := bubble.NewBubbleDB()
	err := sndb.NewBlock(header.Number, parentHash, newBlockHash)
	if err != nil {
		fmt.Errorf("newBlock, %v", err)
	}
	// MOCK
	// store bubble basic
	// main-chain operators config
	var opL1s []*bubble.Operator
	opL1 := bubble.Operator{}
	opL1s = append(opL1s, &opL1)
	// sub-chain operators config
	var opL2s []*bubble.Operator
	opL2 := bubble.Operator{
		OpAddr: sender, // The initial sub-chain operation address as the sender
	}
	opL2s = append(opL2s, &opL2)
	basics := bubble.BubBasics{
		BubbleId:    testBubbleId,
		Creator:     sender,
		CreateBlock: block,
		OperatorsL1: opL1s,
		OperatorsL2: opL2s,
		MicroNodes:  nil,
	}
	// store bubble basics
	if err := bubDB.StoreBubBasics(newBlockHash, testBubbleId, &basics); err != nil {
		fmt.Errorf("failed to StoreBubBasics, %v", err)
	}

	// store bubble state
	if err := bubDB.StoreBubState(newBlockHash, testBubbleId, bubble.ActiveStatus); err != nil {
		fmt.Errorf("failed to StoreBubBasics, %v", err)
	}
	//basic, err := bubDB.GetBubBasics(newBlockHash, testBubbleId)
	//if basic == nil || err != nil {
	//	fmt.Println("basic, %", basic)
	//}
	lastBlockHash = newBlockHash
	lastBlockNumber = block
	lastHeader = header
}

// newBubbleContract
func newBubbleContract(caller common.Address, blockNumber *big.Int, blockHash common.Hash, chain *mock.Chain, index int) *BubbleContract {
	blockContext := BlockContext{
		BlockNumber: blockNumber,
		BlockHash:   blockHash,
		Ctx:         context.Background(),
	}

	contract := &BubbleContract{
		Plugin:   plugin.BubbleInstance(),
		Contract: newContract(common.Big0, caller),
		// Evm:      newEvm(blockNumber, blockHash, chain),
		Evm: NewEVM(blockContext, TxContext{}, chain.SnapDB, chain.StateDB, params.TestChainConfig, Config{}),
	}
	// set current nodeId
	contract.Plugin.NodeID = currentNodeId
	chain.StateDB.Prepare(txHashArr[index], blockHash, index+1)
	return contract
}

// test createBubble interface
func create_bubble(contract *BubbleContract, t *testing.T) {
	var params [][]byte
	params = make([][]byte, 0)
	fnType, _ := rlp.EncodeToBytes(uint16(1))

	params = append(params, fnType)
	runBubbleTx(contract, params, "createBubble", t)
}

// test stakingToken interface
func staking_token(contract *BubbleContract, t *testing.T) {

	var params [][]byte
	params = make([][]byte, 0)
	bubbleId := testBubbleId

	stakingAsset := bubble.AccountAsset{
		Account:      contract.Contract.CallerAddress,
		NativeAmount: testNativeAmount,
	}
	for _, tokenAddr := range testERC20AddrList {
		tokenAsset := bubble.AccTokenAsset{
			TokenAddr: tokenAddr,
			Balance:   testTokenAmount,
		}
		stakingAsset.TokenAssets = append(stakingAsset.TokenAssets, tokenAsset)
	}

	fnType, _ := rlp.EncodeToBytes(uint16(3))
	bubId, _ := rlp.EncodeToBytes(bubbleId)
	accAsset, _ := rlp.EncodeToBytes(stakingAsset)

	params = append(params, fnType)
	params = append(params, bubId)
	params = append(params, accAsset)

	runBubbleTx(contract, params, "stakingToken", t)
}

// test withdrewToken interface
func withdrew_token(contract *BubbleContract, t *testing.T) {

	var params [][]byte
	params = make([][]byte, 0)
	bubbleId := big.NewInt(1)

	fnType, _ := rlp.EncodeToBytes(uint16(4))
	bubId, _ := rlp.EncodeToBytes(bubbleId)

	params = append(params, fnType)
	params = append(params, bubId)
	runBubbleTx(contract, params, "withdrewToken", t)
}

// test settleBubble interface
func settle_bubble(contract *BubbleContract, t *testing.T) {
	var params [][]byte
	params = make([][]byte, 0)
	bubbleId := testBubbleId
	settleInfo := bubble.SettlementInfo{}
	for _, addr := range testAddrList {
		accAsset := bubble.AccountAsset{
			Account:      addr,
			NativeAmount: settleNativeAmount,
		}
		for _, tokenAddr := range testERC20AddrList {
			tokenAsset := bubble.AccTokenAsset{
				TokenAddr: tokenAddr,
				Balance:   settleTokenAmount,
			}
			accAsset.TokenAssets = append(accAsset.TokenAssets, tokenAsset)
		}
		settleInfo.AccAssets = append(settleInfo.AccAssets, accAsset)
	}

	fnType, _ := rlp.EncodeToBytes(uint16(5))
	txHash, _ := rlp.EncodeToBytes(L2SettleTxHash)
	bubId, _ := rlp.EncodeToBytes(bubbleId)
	settle, _ := rlp.EncodeToBytes(settleInfo)

	params = append(params, fnType)
	params = append(params, txHash)
	params = append(params, bubId)
	params = append(params, settle)
	runBubbleTx(contract, params, "settleBubble", t)
}

// test getBubbleInfo interface
func get_bubble_info(contract *BubbleContract, t *testing.T) {

	var params [][]byte
	params = make([][]byte, 0)
	bubbleId := big.NewInt(1)

	fnType, _ := rlp.EncodeToBytes(uint16(100))
	bubId, _ := rlp.EncodeToBytes(bubbleId)

	params = append(params, fnType)
	params = append(params, bubId)
	runBubbleCall(contract, params, "getBubbleInfo", nil, t)
}

// test getL1HashByL2Hash interface
func getL1TxHashByL2TxHash(contract *BubbleContract, VerifyL1TxHash common.Hash, t *testing.T) {

	var params [][]byte
	params = make([][]byte, 0)
	bubbleId := testBubbleId

	fnType, _ := rlp.EncodeToBytes(uint16(101))
	bubId, _ := rlp.EncodeToBytes(bubbleId)
	txHash, _ := rlp.EncodeToBytes(L2SettleTxHash)

	params = append(params, fnType)
	params = append(params, bubId)
	params = append(params, txHash)
	runBubbleCall(contract, params, "getL1HashByL2Hash", VerifyL1TxHash.Hex(), t)
}

// test getBubTxHashList interface
func getBubbleTxHashList(contract *BubbleContract, txType bubble.BubTxType, t *testing.T) {

	var params [][]byte
	params = make([][]byte, 0)
	bubbleId := testBubbleId

	fnType, _ := rlp.EncodeToBytes(uint16(102))
	bubId, _ := rlp.EncodeToBytes(bubbleId)
	bubTxType, _ := rlp.EncodeToBytes(txType)

	params = append(params, fnType)
	params = append(params, bubId)
	params = append(params, bubTxType)
	runBubbleCall(contract, params, "getBubTxHashList", nil, t)
}

// verify token count
func verify_token_amount(contract *BubbleContract, NativeAmount *big.Int, tokenAmount *big.Int, t *testing.T) {
	// query balances
	accAssets, err := getBalances(contract.Evm, contract.Contract, testAddrList)
	if accAssets == nil || err != nil {
		panic(fmt.Errorf("failed to getBalances: %v", err.Error()))
	}
	assert.True(t, nil == err)
	assert.True(t, nil != accAssets)
	// Compare the number of native tokens and erc20 tokens of sender after initialization
	for _, accAsset := range *accAssets {
		result := NativeAmount.Cmp(accAsset.NativeAmount)
		assert.Equal(t, 0, result, "Native token count validation failed")
		for _, tokenAsset := range accAsset.TokenAssets {
			result := tokenAmount.Cmp(tokenAsset.Balance)
			assert.Equal(t, 0, result, "Erc20 token count verification failed")
		}
	}
}

// decode_and_verify_stakingToken_tx_receipt Parse and verify the logs information in the stakingToken transaction receipt
func decode_and_verify_stakingToken_tx_receipt(bubContract *BubbleContract, stakingAccount common.Address, txIndex int, t *testing.T) {
	logs := bubContract.Evm.StateDB.GetLogs(txHashArr[txIndex])
	for _, log := range logs {
		// Deal only with switchable viewer is empty,
		// stakingToken transaction logs will be written to erc20 COINS event information, need to filter
		if nil != (*log).Topics || (*log).Address != vm.BubbleContractAddr {
			continue
		}
		data := (*log).Data
		// t.Logf("stakingToken tx logs: %v", data)
		var bubbleId *big.Int
		var stakingAsset bubble.AccountAsset
		var m [][]byte
		if err := rlp.DecodeBytes(data, &m); err != nil {
			t.Error(err)
		}
		var code string
		err := rlp.DecodeBytes(m[0], &code)
		assert.True(t, nil == err)
		assert.True(t, code == "0")

		err = rlp.DecodeBytes(m[1], &bubbleId)
		assert.True(t, nil == err)
		result := bubbleId.Cmp(testBubbleId)
		assert.Equal(t, 0, result, "BubbleID error obtained")

		if err := rlp.DecodeBytes(m[2], &stakingAsset); err != nil {
			t.Error(err)
		}
		assert.True(t, stakingAccount == stakingAsset.Account, "stakingToken account error, and the test account is inconsistent")
		result = testNativeAmount.Cmp(stakingAsset.NativeAmount)
		assert.Equal(t, 0, result, "Incorrect staking native token count, inconsistent with test native token count")
		assert.True(t, len(testERC20AddrList) == len(stakingAsset.TokenAssets), "Inconsistent length of staking token accounts")
		for i, tokenAddr := range testERC20AddrList {
			tokenAsset := stakingAsset.TokenAssets[i]
			assert.True(t, tokenAddr == tokenAsset.TokenAddr, "Inconsistent erc20 token address")
			result := testTokenAmount.Cmp(tokenAsset.Balance)
			assert.Equal(t, 0, result, "Inconsistent erc20 token amount")
		}
		// fmt.Printf("staking AccoutAsset: %v\n", stakingAsset)
	}
}

// decode_and_verify_withdrewToken_tx_receipt Parse and verify the logs information in the withdrewToken transaction receipt
func decode_and_verify_withdrewToken_tx_receipt(bubContract *BubbleContract, caller common.Address, txIndex int, t *testing.T) {
	logs := bubContract.Evm.StateDB.GetLogs(txHashArr[txIndex])
	for _, log := range logs {
		// Deal only with switchable viewer is empty,
		// withdrewToken transaction logs will be written to erc20 COINS event information, need to filter
		if nil != (*log).Topics || (*log).Address != vm.BubbleContractAddr {
			continue
		}
		data := (*log).Data
		// t.Logf("withdrewToken tx logs: %v", data)
		var bubbleId *big.Int
		var accAsset bubble.AccountAsset
		var m [][]byte
		if err := rlp.DecodeBytes(data, &m); err != nil {
			t.Error(err)
		}
		var code string
		err := rlp.DecodeBytes(m[0], &code)
		assert.True(t, nil == err)
		assert.True(t, code == "0")

		err = rlp.DecodeBytes(m[1], &bubbleId)
		assert.True(t, nil == err)
		result := bubbleId.Cmp(testBubbleId)
		assert.Equal(t, 0, result, "BubbleID error obtained")

		if err := rlp.DecodeBytes(m[2], &accAsset); err != nil {
			t.Error(err)
		}
		assert.True(t, caller == accAsset.Account, "withdrewToken account error, and the test account is inconsistent")
		result = testNativeAmount.Cmp(accAsset.NativeAmount)
		assert.Equal(t, 0, result, "Incorrect withdrew native token count, inconsistent with test native token count")
		assert.True(t, len(testERC20AddrList) == len(accAsset.TokenAssets), "Inconsistent length of withdrew token accounts")
		for i, tokenAddr := range testERC20AddrList {
			tokenAsset := accAsset.TokenAssets[i]
			assert.True(t, tokenAddr == tokenAsset.TokenAddr, "Inconsistent erc20 token address")
			result := testTokenAmount.Cmp(tokenAsset.Balance)
			assert.Equal(t, 0, result, "Inconsistent erc20 token amount")
		}
		// fmt.Printf("withdrew AccoutAsset: %v\n", accAsset)
	}
}

// decode_and_verify_settleBubble_tx_receipt Parse and verify the logs information in the settleBubble transaction receipt
func decode_and_verify_settleBubble_tx_receipt(bubContract *BubbleContract, txIndex int, t *testing.T) {
	logs := bubContract.Evm.StateDB.GetLogs(txHashArr[txIndex])
	for _, log := range logs {
		// Deal only with switchable viewer is empty,
		// settleBubble transaction logs will be written to erc20 COINS event information, need to filter
		if nil != (*log).Topics || (*log).Address != vm.BubbleContractAddr {
			continue
		}
		data := (*log).Data
		// t.Logf("settleBubble tx logs: %v", data)
		var m [][]byte
		if err := rlp.DecodeBytes(data, &m); err != nil {
			t.Error(err)
		}
		var code string
		err := rlp.DecodeBytes(m[0], &code)
		assert.True(t, nil == err)
		assert.True(t, code == "0")
		// parse L2TxHash
		var L2TxHash common.Hash
		err = rlp.DecodeBytes(m[1], &L2TxHash)
		assert.True(t, nil == err)
		assert.True(t, L2TxHash == L2SettleTxHash, "Error in obtaining the sub-chain settleBubble transaction hash")

		// parse bubbleID
		var bubbleId *big.Int
		err = rlp.DecodeBytes(m[2], &bubbleId)
		assert.True(t, nil == err)
		result := bubbleId.Cmp(testBubbleId)
		assert.Equal(t, 0, result, "BubbleID error obtained")

		// parse settlementInfo
		var settleInfo bubble.SettlementInfo
		if err := rlp.DecodeBytes(m[3], &settleInfo); err != nil {
			t.Error(err)
		}
		assert.True(t, nil == err)
		assert.True(t, len(testAddrList) == len(settleInfo.AccAssets), "Settlement account number error")

		for i, accAsset := range settleInfo.AccAssets {
			account := testAddrList[i]
			assert.True(t, account == accAsset.Account, "settleBubble account error, and the test account is inconsistent")
			result = settleNativeAmount.Cmp(accAsset.NativeAmount)
			assert.Equal(t, 0, result, "Incorrect settleBubble native token count, inconsistent with test native token count")
			assert.True(t, len(testERC20AddrList) == len(accAsset.TokenAssets), "Inconsistent length of settleBubble token accounts")
			for i, tokenAddr := range testERC20AddrList {
				tokenAsset := accAsset.TokenAssets[i]
				assert.True(t, tokenAddr == tokenAsset.TokenAddr, "Inconsistent erc20 token address")
				result := settleTokenAmount.Cmp(tokenAsset.Balance)
				assert.Equal(t, 0, result, "Inconsistent erc20 token amount")
			}
		}

		// fmt.Printf("withdrew AccoutAsset: %v\n", accAsset)
	}
}

/**
Standard test cases
*/
func TestBubbleContract_createBubble(t *testing.T) {
	//chain := newMockChain()
	//defer chain.SnapDB.Clear()
	//newPlugins()
	//
	//defer func() {
	//	sndb.Clear()
	//}()
	//handler.NewVrfHandler(hexutil.MustDecode("0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23"))
	//
	//if err := chain.SnapDB.NewBlock(blockNumber, chain.Genesis.Hash(), blockHash); nil != err {
	//	t.Error("newBlock err", err)
	//}
	//
	//chain.StateDB.Prepare(txHashArr[0], blockHash3, 0)
	//contract := newBubbleContract(sender, blockNumber, blockHash, chain, 1)
	//create_bubble(contract, t)
}

func TestBubbleContract_stakingToken(t *testing.T) {
	chain := newMockChain()
	defer chain.SnapDB.Clear()

	genesisBlockHash = chain.Genesis.Hash()
	t.Logf("genesisHash: %v", genesisBlockHash)

	// first block
	build_bubble_data(1, genesisBlockHash)
	storeBlockHash = lastBlockHash
	contract := newBubbleContract(sender, blockNumber, storeBlockHash, chain, 1)
	// init erc20 data
	init_erc20_data(contract, 2)
	sBalance, _ := new(big.Int).SetString(senderBalance, 10)
	// Validation after initialization sender native tokens and erc20 token number
	verify_token_amount(contract, sBalance, testTokenAmount, t)

	// Cycle call stakingToken system contracts interface
	index := 2
	for _, caller := range testAddrList {
		contract.Contract = newContract(common.Big0, caller)
		chain.StateDB.Prepare(txHashArr[index], blockHash, index+1)
		staking_token(contract, t)
		decode_and_verify_stakingToken_tx_receipt(contract, caller, index, t)
		index++
	}

	// The sender's number of native tokens and erc20 tokens after verifying the pledged token
	verify_token_amount(contract, new(big.Int).Sub(sBalance, testNativeAmount), big0, t)
}

func TestBubbleContract_withdrewToken(t *testing.T) {
	chain := newMockChain()
	defer chain.SnapDB.Clear()
	genesisBlockHash = chain.Genesis.Hash()
	t.Logf("genesisHash: %v", genesisBlockHash)

	// first block
	build_bubble_data(1, genesisBlockHash)
	storeBlockHash = lastBlockHash
	contract := newBubbleContract(sender, blockNumber, storeBlockHash, chain, 1)
	// init erc20 data
	init_erc20_data(contract, 2)
	// Cycle call stakingToken system contracts interface
	for _, caller := range testAddrList {
		contract.Contract = newContract(common.Big0, caller)
		staking_token(contract, t)
	}

	sBalance, _ := new(big.Int).SetString(senderBalance, 10)
	// The sender's number of native tokens and erc20 tokens after verifying the pledged token
	verify_token_amount(contract, new(big.Int).Sub(sBalance, testNativeAmount), big0, t)
	// modify state: The simulation bubble has been released
	// store bubble state
	bubDB := bubble.NewBubbleDB()
	if err := bubDB.StoreBubState(storeBlockHash, testBubbleId, bubble.ReleasedStatus); err != nil {
		fmt.Errorf("failed to StoreBubBasics, %v", err)
	}
	// call withdrewToken
	index := 2
	for _, caller := range testAddrList {
		contract.Contract = newContract(common.Big0, caller)
		chain.StateDB.Prepare(txHashArr[index], blockHash, index+1)
		withdrew_token(contract, t)
		decode_and_verify_withdrewToken_tx_receipt(contract, caller, index, t)
		index++
	}
	// Validation of redemption after primary tokens and erc20 token number
	verify_token_amount(contract, sBalance, testTokenAmount, t)
}

func TestBubbleContract_settleBubble(t *testing.T) {
	chain := newMockChain()
	defer chain.SnapDB.Clear()
	genesisBlockHash = chain.Genesis.Hash()
	t.Logf("genesisHash: %v", genesisBlockHash)

	// first block
	build_bubble_data(1, genesisBlockHash)
	storeBlockHash = lastBlockHash
	contract := newBubbleContract(sender, blockNumber, storeBlockHash, chain, 1)
	// init erc20 data
	init_erc20_data(contract, 2)
	// Cycle call stakingToken system contracts interface
	for _, caller := range testAddrList {
		contract.Contract = newContract(common.Big0, caller)
		staking_token(contract, t)
	}

	// call settleBubble：Only the sub-chain operating address can send settlement transactions
	index := 2
	contract.Contract = newContract(common.Big0, sender)
	chain.StateDB.Prepare(txHashArr[index], blockHash, index+1)
	settle_bubble(contract, t)
	decode_and_verify_settleBubble_tx_receipt(contract, index, t)
	index++

	// call withdrewToken：It can be redeemed only when the bubble state is released
	// store bubble state
	bubDB := bubble.NewBubbleDB()
	if err := bubDB.StoreBubState(storeBlockHash, testBubbleId, bubble.ReleasedStatus); err != nil {
		fmt.Errorf("failed to StoreBubBasics, %v", err)
	}
	for _, caller := range testAddrList {
		contract.Contract = newContract(common.Big0, caller)
		withdrew_token(contract, t)
	}
	// Validation of redemption after primary tokens and erc20 token number
	sBalance, _ := new(big.Int).SetString(senderBalance, 10)
	verify_token_amount(contract, new(big.Int).Sub(sBalance, testStep), settleTokenAmount, t)
}

func TestBubbleContract_getBubbleInfo(t *testing.T) {
	chain := newMockChain()
	defer chain.SnapDB.Clear()
	genesisBlockHash = chain.Genesis.Hash()
	t.Logf("genesisHash: %v", genesisBlockHash)

	// first block
	build_bubble_data(1, genesisBlockHash)
	storeBlockHash = lastBlockHash
	contract := newBubbleContract(sender, blockNumber, storeBlockHash, chain, 1)

	get_bubble_info(contract, t)
}

func TestBubbleContract_getL1HashByL2Hash(t *testing.T) {
	chain := newMockChain()
	defer chain.SnapDB.Clear()
	genesisBlockHash = chain.Genesis.Hash()
	t.Logf("genesisHash: %v", genesisBlockHash)

	// first block
	build_bubble_data(1, genesisBlockHash)
	storeBlockHash = lastBlockHash
	contract := newBubbleContract(sender, blockNumber, storeBlockHash, chain, 1)
	// init erc20 data
	init_erc20_data(contract, 2)
	// Cycle call stakingToken system contracts interface
	for _, caller := range testAddrList {
		contract.Contract = newContract(common.Big0, caller)
		staking_token(contract, t)
	}

	contract.Contract = newContract(common.Big0, sender)
	settle_bubble(contract, t)
	t.Logf("settle_bubble transaction tx hash: %v", contract.Evm.StateDB.TxHash())
	getL1TxHashByL2TxHash(contract, contract.Evm.StateDB.TxHash(), t)
}

func TestBubbleContract_getBubTxHashList(t *testing.T) {
	chain := newMockChain()
	defer chain.SnapDB.Clear()
	genesisBlockHash = chain.Genesis.Hash()
	t.Logf("genesisHash: %v", genesisBlockHash)

	// first block
	build_bubble_data(1, genesisBlockHash)
	storeBlockHash = lastBlockHash
	contract := newBubbleContract(sender, blockNumber, storeBlockHash, chain, 1)
	// init erc20 data
	init_erc20_data(contract, 2)
	// Cycle call stakingToken system contracts interface
	for _, caller := range testAddrList {
		contract.Contract = newContract(common.Big0, caller)
		staking_token(contract, t)
	}
	// query stakingToken txs
	getBubbleTxHashList(contract, bubble.StakingToken, t)

	contract.Contract = newContract(common.Big0, sender)
	settle_bubble(contract, t)
	// query settleBubble txs
	getBubbleTxHashList(contract, bubble.SettleBubble, t)

	bubDB := bubble.NewBubbleDB()
	if err := bubDB.StoreBubState(storeBlockHash, testBubbleId, bubble.ReleasedStatus); err != nil {
		fmt.Errorf("failed to StoreBubBasics, %v", err)
	}
	for _, caller := range testAddrList {
		contract.Contract = newContract(common.Big0, caller)
		withdrew_token(contract, t)
	}
	// query withdrewToken txs
	getBubbleTxHashList(contract, bubble.WithdrewToken, t)
}
