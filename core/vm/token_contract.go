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
	"errors"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/plugin"
	"github.com/bubblenet/bubble/x/token"
	"math/big"
)

const (
	TmpERC20Addr = "0x0000000000000000000000000000000000000020" // ERC20 template address
	TxMintToken  = 6000
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
		TxMintToken: tkc.mintToken,

		// Get
	}
}
func (tkc *TokenContract) FnSigns() map[uint16]interface{} {
	fnSigns := tkc.FnSignsV1()
	return fnSigns
}

// 单个账户铸币（主链运营节点调用）
func (tkc *TokenContract) mintToken(accAsset token.AccountAsset) ([]byte, error) {
	from := tkc.Contract.CallerAddress
	if from != tkc.Plugin.MainOpAddr {
		return nil, errors.New("the transaction sender is not the main chain operator address")
	}

	txHash := tkc.Evm.StateDB.TxHash()
	blockNumber := tkc.Evm.Context.BlockNumber
	state := tkc.Evm.StateDB

	log.Debug("Call mintToken of TokenContract", "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "caller", from)

	// 计算gas
	if !tkc.Contract.UseGas(params.TokenGas) {
		return nil, ErrOutOfGas
	}

	if txHash == common.ZeroHash {
		return nil, nil
	}

	// 存储新用户信息
	accList := make([]common.Address, 1)
	accList[0] = accAsset.Account
	tokenAddrList := make([]common.Address, len(accAsset.TokenAssets))
	for i, tokenAsset := range accAsset.TokenAssets {
		tokenAddrList[i] = tokenAsset.TokenAddr
	}
	// 1.铸币原生Token
	// 1.1 从系统合约账户向account转账
	if accAsset.NativeAmount.Cmp(common.Big0) > 0 {
		state.AddBalance(accAsset.Account, accAsset.NativeAmount)
	}

	// 保存数据
	if err := tkc.Plugin.AddMintAccInfo(state, token.MintAccInfo{
		AccList:       accList,
		TokenAddrList: tokenAddrList,
	}); err != nil {
		return nil, err
	}

	// 2.铸币ERC20代币（默认精度为6）
	// 2.1 判断是否ERC20是否存在，不存在则需要部署
	for _, tokenAsset := range accAsset.TokenAssets {
		// ERC20地址
		erc20Addr := tokenAsset.TokenAddr
		// Token金额
		tokenAmount := tokenAsset.Balance
		code := tkc.Evm.StateDB.GetCode(erc20Addr)
		contract := tkc.Contract
		if len(code) == 0 {
			tmpErc20Addr := common.HexToAddress(TmpERC20Addr)
			tempCode := tkc.Evm.StateDB.GetCode(tmpErc20Addr)
			// 部署合约
			code = tempCode
			tkc.Evm.StateDB.SetCode(erc20Addr, code)
			// 初始化
		}

		// 开始铸ERC20 Token币
		// 2.修改调用为合约地址（表示合约的调用者，合约交易的发送者）
		contract.caller = AccountRef(vm.TokenContractAddr)
		contract.CallerAddress = vm.TokenContractAddr
		// 修改成ERC20合约地址
		contract.self = AccountRef(erc20Addr)
		contract.SetCallCode(&erc20Addr, tkc.Evm.StateDB.GetCodeHash(erc20Addr), code)
		for _, interpreter := range tkc.Evm.interpreters {
			if interpreter.CanRun(contract.Code) {
				if tkc.Evm.interpreter != interpreter {
					// Ensure that the interpreter pointer is set back
					// to its current value upon return.
					defer func(i Interpreter) {
						tkc.Evm.interpreter = i
					}(tkc.Evm.interpreter)
					tkc.Evm.interpreter = interpreter
				}
				input, err := encodeMintFuncCall(from, tokenAmount)
				if err != nil {
					log.Error("Failed to Mint ERC20 Token", "error", err)
					return nil, err
				}
				ret, err := interpreter.Run(contract, input, false)
				if err != nil {
					log.Error("Failed to Mint ERC20 Token", "ret", ret, "error", err)
					return ret, err
				}
			}
		}
	}

	return txResultHandlerWithRes(vm.TokenContractAddr, tkc.Evm, "",
		"", TxMintToken, int(common.NoErr.Code), accAsset.Account, accAsset.NativeAmount,
		accAsset.TokenAssets), nil
}
