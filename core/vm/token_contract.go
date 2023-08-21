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
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/plugin"
	"github.com/bubblenet/bubble/x/token"
	"math/big"
)

const (
	TmpERC20Addr          = "0x0000000000000000000000000000000000000020" // ERC20 template address
	TxMintToken           = 6000
	TxSettleBubble        = 6001
	CallGetL2HashByL1Hash = 6100
)

type TokenContract struct {
	Plugin   *plugin.TokenPlugin
	Contract *Contract
	Evm      *EVM
}

func (tkc *TokenContract) RequiredGas(input []byte) uint64 {
	if checkInputEmpty(input) {
		return 0
	}
	return params.TokenGas
}

func (tkc *TokenContract) Run(input []byte) ([]byte, error) {
	if checkInputEmpty(input) {
		return nil, nil
	}
	return execBubbleContract(input, tkc.FnSigns())
}

func (tkc *TokenContract) CheckGasPrice(gasPrice *big.Int, fcode uint16) error {
	return nil
}

func (tkc *TokenContract) FnSignsV1() map[uint16]interface{} {
	return map[uint16]interface{}{
		// Set
		TxMintToken:    tkc.mintToken,
		TxSettleBubble: tkc.settleBubble,

		// Get
		CallGetL2HashByL1Hash: tkc.getL2HashByL1Hash,
	}
}
func (tkc *TokenContract) FnSigns() map[uint16]interface{} {
	fnSigns := tkc.FnSignsV1()
	return fnSigns
}

// mintToken The operator node responsible for this bubble on the main chain calls this minting interface
// Minting the same currency and quantity of the corresponding account
// The mapping between the transaction hash of the pledged token on the main chain and the mintage transaction hash is recorded
func (tkc *TokenContract) mintToken(L1StakingTokenTxHash common.Hash, accAsset token.AccountAsset) ([]byte, error) {
	txHash := tkc.Evm.StateDB.TxHash()
	blockNumber := tkc.Evm.Context.BlockNumber

	// Call handling logic
	_, err := MintToken(tkc, L1StakingTokenTxHash, accAsset)
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.TokenContractAddr, tkc.Evm, "mintToken", bizErr.Error(), TxMintToken, bizErr)
		} else {
			log.Error("Failed to mintToken", "txHash", txHash, "blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}
	return txResultHandlerWithRes(vm.TokenContractAddr, tkc.Evm, "",
		"", TxMintToken, int(common.NoErr.Code), L1StakingTokenTxHash, accAsset), nil
}

// SettleBubble Logic functions that handle the settleBubble system's contract interface
func SettleBubble(tkc *TokenContract) (*token.SettlementInfo, error) {
	txHash := tkc.Evm.StateDB.TxHash()
	state := tkc.Evm.StateDB
	blockHash := tkc.Evm.Context.BlockHash

	// Calculating gas
	if !tkc.Contract.UseGas(params.TokenGas) {
		return nil, token.ErrOutOfGas
	}

	if txHash == common.ZeroHash {
		return nil, token.ErrZeroHash
	}

	// Get minting account information
	mintAccInfo, err := tkc.Plugin.GetMintAccInfo(blockHash)
	if err != nil || nil == mintAccInfo {
		return nil, token.ErrGetMintAccInfo
	}

	// Assembly settlement information
	var settlementInfo token.SettlementInfo
	// Get the account's assets
	// Get the account's native Token
	for i, acc := range mintAccInfo.AccList {
		balance := state.GetBalance(acc)
		// Assemble native Token settlement information
		var accAsset token.AccountAsset
		accAsset.Account = mintAccInfo.AccList[i]
		accAsset.NativeAmount = balance
		settlementInfo.AccAssets = append(settlementInfo.AccAssets, accAsset)
	}

	// Get the account's ERC20 Token
	for _, tokenAddr := range mintAccInfo.TokenAddrList {
		code := tkc.Evm.StateDB.GetCode(tokenAddr)
		if len(code) > 0 {
			// Contract deployed
			contract := tkc.Contract
			// Change to ERC20 contract address
			contract.self = AccountRef(tokenAddr)
			contract.SetCallCode(&tokenAddr, tkc.Evm.StateDB.GetCodeHash(tokenAddr), code)
			// Batch queries for erc20 token balances in the list of accounts
			input, err := encodeGetBalancesCall(mintAccInfo.AccList)
			if err != nil {
				log.Error("Failed to get Address ERC20 Token", "error", err)
				return nil, token.ErrEncodeGetBalancesData
			}
			// Execute EVM
			ret, err := RunEvm(tkc.Evm, contract, input)
			if err != nil {
				log.Error("Failed to Mint ERC20 Token", "error", err)
				return nil, token.ErrEVMExecERC20
			}
			// Parse byte array to uint256 array,
			// the first 32 bytes value indicates how many bytes to store the length of the array,
			// fixed as: 32
			resList := parseBytesToUint256Array(ret[32:])

			if len(resList) > 0 {
				// The value of the first element indicates the length of the returned array
				elemLen := resList[0].Uint64()
				if elemLen != uint64(len(mintAccInfo.AccList)) {
					log.Error("Failed to get Address ERC20 Token", "error",
						"The length of the number of accounts and the number of balances retrieved are inconsistent")
					return nil, token.ErrGetERC20Token
				}
				// Assemble the ERC20 Token settlement information
				for iAcc, balance := range resList[1:] {
					var accTokenAsset token.AccTokenAsset
					accTokenAsset.TokenAddr = tokenAddr
					accTokenAsset.Balance = balance
					settlementInfo.AccAssets[iAcc].TokenAssets = append(settlementInfo.AccAssets[iAcc].TokenAssets, accTokenAsset)
				}
			}
		}
	}
	// Get the hash of the last settlement transaction
	lastHash, err := token.GetSettlementHash(blockHash)
	if nil != err {
		return nil, token.ErrGetSettleInfoHash
	}
	// Calculate the current account settlement Hash
	hash, err := settlementInfo.Hash()
	if nil != err {
		return nil, token.ErrCalcSettleInfoHash
	}

	// Comparing
	if lastHash != nil && *lastHash == hash {
		// There are no related accounts within the bubble to generate new transactions, so there is no need for settlement
		return nil, token.ErrNotNeedToSettle
	} else {
		// settlement task
		// store current hash
		token.StoreSettlementHash(blockHash, hash)
		// Determine whether the current node is a sub-chain operator node
		if tkc.Plugin.IsSubOpNode {
			settleTask := token.SettleTask{TxHash: tkc.Evm.StateDB.TxHash(), SettleInfo: settlementInfo}
			// Send settlement task
			tkc.Plugin.PostSettlementTask(&settleTask)
		}
	}
	return &settlementInfo, nil
}

func MintToken(tkc *TokenContract, L1StakingTokenTxHash common.Hash, accAsset token.AccountAsset) ([]byte, error) {
	from := tkc.Contract.CallerAddress
	if from != tkc.Plugin.MainOpAddr {
		return nil, token.ErrNotMainOpAddr
	}

	txHash := tkc.Evm.StateDB.TxHash()
	blockNumber := tkc.Evm.Context.BlockNumber
	state := tkc.Evm.StateDB
	blockHash := tkc.Evm.Context.BlockHash

	log.Debug("Call mintToken of TokenContract", "blockHash", blockHash, "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "caller", from)

	// Calculating gas
	if !tkc.Contract.UseGas(params.MintTokenGas) {
		return nil, token.ErrOutOfGas
	}

	if txHash == common.ZeroHash {
		return nil, token.ErrZeroHash
	}

	// Store new user information
	accList := make([]common.Address, 1)
	accList[0] = accAsset.Account
	tokenAddrList := make([]common.Address, len(accAsset.TokenAssets))
	for i, tokenAsset := range accAsset.TokenAssets {
		tokenAddrList[i] = tokenAsset.TokenAddr
	}
	// Mint native tokens
	if accAsset.NativeAmount.Cmp(common.Big0) > 0 {
		// Transfer money from system contract account to Account
		state.AddBalance(accAsset.Account, accAsset.NativeAmount)
	}

	// Minting ERC20 tokens (default decimal is 6)
	for _, tokenAsset := range accAsset.TokenAssets {
		// ERC20 address
		erc20Addr := tokenAsset.TokenAddr
		// Token Amount
		tokenAmount := tokenAsset.Balance
		// contract code
		code := tkc.Evm.StateDB.GetCode(erc20Addr)
		contract := tkc.Contract
		// Determine whether ERC20 exists or not, if it does not exist, it needs to be deployed
		if len(code) == 0 {
			// Get the contract code from the ERC20 template
			tmpErc20Addr := common.HexToAddress(TmpERC20Addr)
			tempCode := tkc.Evm.StateDB.GetCode(tmpErc20Addr)
			// Deploy the ERC20 contract
			code = tempCode
			tkc.Evm.StateDB.SetCode(erc20Addr, code)
			// Initial combination information, such as decimal, etc
			// ToDo
		}

		// Start ERC20 Token mint
		// Change the call to the contract address (representing the caller of the contract, the sender of the contract transaction)
		contract.caller = AccountRef(vm.TokenContractAddr)
		contract.CallerAddress = vm.TokenContractAddr
		// Change to ERC20 contract address
		contract.self = AccountRef(erc20Addr)
		contract.SetCallCode(&erc20Addr, tkc.Evm.StateDB.GetCodeHash(erc20Addr), code)
		// Mint money to a designated account
		input, err := encodeMintFuncCall(accAsset.Account, tokenAmount)
		if err != nil {
			log.Error("Failed to Mint ERC20 Token", "error", err)
			return nil, token.ErrEncodeMintData
		}
		_, err = RunEvm(tkc.Evm, contract, input)
		if err != nil {
			log.Error("Failed to Mint ERC20 Token", "error", err)
			return nil, token.ErrEVMExecERC20
		}
	}

	// Add minting account information
	if err := tkc.Plugin.AddMintAccInfo(blockHash, token.MintAccInfo{
		AccList:       accList,
		TokenAddrList: tokenAddrList,
	}); err != nil {
		return nil, token.ErrAddMintAccInfo
	}

	// The mapping relationship between the sub-chain settlement transaction hash and the main chain settlement transaction hash is stored
	if err := tkc.Plugin.StoreL1HashToL2Hash(blockHash, tkc.Evm.StateDB.TxHash(), L1StakingTokenTxHash); nil != err {
		return nil, token.ErrStoreL1HashToL2Hash
	}
	return nil, nil
}

// settleBubble sub-chain settle transactions
func (tkc *TokenContract) settleBubble() ([]byte, error) {
	from := tkc.Contract.CallerAddress
	//if from != tkc.Plugin.MainOpAddr {
	//	return nil, errors.New("the transaction sender is not the main chain operator address")
	//}
	txHash := tkc.Evm.StateDB.TxHash()
	blockNumber := tkc.Evm.Context.BlockNumber
	// state := tkc.Evm.StateDB
	blockHash := tkc.Evm.Context.BlockHash
	log.Debug("Call mintToken of TokenContract", "blockHash", blockHash, "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "caller", from.Hex())

	// Call handling logic
	settlementInfo, err := SettleBubble(tkc)
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.TokenContractAddr, tkc.Evm, "settleBubble", bizErr.Error(), TxSettleBubble, bizErr)
		} else {
			log.Error("Failed to settleBubble", "txHash", txHash, "blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	// record log
	return txResultHandlerWithRes(vm.TokenContractAddr, tkc.Evm, "",
		"", TxSettleBubble, int(common.NoErr.Code), settlementInfo), nil
}

// getL2HashByL1Hash The sub-chain settlement transaction hash is obtained according to the settlement transaction hash of the main chain
func (tkc *TokenContract) getL2HashByL1Hash(L1TxHash common.Hash) ([]byte, error) {
	blockHash := tkc.Evm.Context.BlockHash

	txHash, err := tkc.Plugin.GetL2HashByL1Hash(blockHash, L1TxHash)
	if err != nil {
		return callResultHandler(tkc.Evm, "getL2HashByL1Hash, txHash", txHash, token.ErrGetL2TxHashByL1), err
	}

	return callResultHandler(tkc.Evm, "getL2HashByL1Hash, txHash", txHash, nil), nil
}
