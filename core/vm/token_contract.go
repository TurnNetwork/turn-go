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
	TxSettlement = 6001
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
		TxMintToken:  tkc.mintToken,
		TxSettlement: tkc.settlement,

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
		// 铸币给指定账户
		input, err := encodeMintFuncCall(accAsset.Account, tokenAmount)
		if err != nil {
			log.Error("Failed to Mint ERC20 Token", "error", err)
			return nil, err
		}
		_, err = RunEvm(tkc.Evm, contract, input)
		if err != nil {
			log.Error("Failed to Mint ERC20 Token", "error", err)
			return nil, err
		}
	}

	return txResultHandlerWithRes(vm.TokenContractAddr, tkc.Evm, "",
		"", TxMintToken, int(common.NoErr.Code), accAsset.Account, accAsset.NativeAmount,
		accAsset.TokenAssets), nil
}

// 单个账户铸币（主链运营节点调用）
func (tkc *TokenContract) settlement() ([]byte, error) {
	from := tkc.Contract.CallerAddress
	// 不做限制
	//if from != tkc.Plugin.MainOpAddr {
	//	return nil, errors.New("the transaction sender is not the main chain operator address")
	//}

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

	// 获取用户的信息
	mintAccInfo, err := tkc.Plugin.GetMintAccInfo(state)
	if err != nil || nil == mintAccInfo {
		return nil, err
	}

	// 组装结算信息
	var settlementInfo token.SettlementInfo
	// 获取用户的资产
	// 获取用户的原生Token
	for i, acc := range mintAccInfo.AccList {
		balance := state.GetBalance(acc)
		// 组装原生Token结算信息
		var accAsset token.AccountAsset
		accAsset.Account = mintAccInfo.AccList[i]
		accAsset.NativeAmount = balance
		settlementInfo.AccAssets = append(settlementInfo.AccAssets, accAsset)
	}

	// 获取用户的ERC20 Token
	for _, tokenAddr := range mintAccInfo.TokenAddrList {
		code := tkc.Evm.StateDB.GetCode(tokenAddr)
		if len(code) > 0 {
			// 合约已部署
			contract := tkc.Contract
			// 修改成ERC20合约地址
			contract.self = AccountRef(tokenAddr)
			contract.SetCallCode(&tokenAddr, tkc.Evm.StateDB.GetCodeHash(tokenAddr), code)
			input, err := encodeGetBalancesCall(mintAccInfo.AccList)
			if err != nil {
				log.Error("Failed to get Address ERC20 Token", "error", err)
				return nil, err
			}
			// 执行EVM
			ret, err := RunEvm(tkc.Evm, contract, input)
			if err != nil {
				log.Error("Failed to Mint ERC20 Token", "error", err)
				return nil, err
			}
			// 解析字节数组为 uint256 数组,前32个字节的值表示用多少个字节存储数组的长度，固定为：32
			balances := parseBytesToUint256Array(ret[32:])

			if len(balances) > 0 {
				// 第一个元素的值表示返回的数组长度
				elemLen := balances[0].Uint64()
				if elemLen != uint64(len(mintAccInfo.AccList)) {
					log.Error("Failed to get Address ERC20 Token", "error",
						"The length of the number of accounts and the number of balances retrieved are inconsistent")
					return ret, errors.New("failed to get Address ERC20 Token")
				}
				// 组装ERC20 Token结算信息
				for iAcc, balance := range balances[1:] {
					var accTokenAsset token.AccTokenAsset
					accTokenAsset.TokenAddr = tokenAddr
					accTokenAsset.Balance = balance
					settlementInfo.AccAssets[iAcc].TokenAssets = append(settlementInfo.AccAssets[iAcc].TokenAssets, accTokenAsset)
				}
			}
		}
	}
	// 获取最近一次的结算hash
	lastHash, err := token.GetSettlementHash(state)
	if nil != err {
		return nil, err
	}
	// 计算当前Hash
	hash, err := settlementInfo.Hash()
	if nil != err {
		return nil, err
	}
	// 比较hash
	if lastHash != nil && *lastHash == hash {
		// bubble网络内没有相关账户产生新的交易，不需要做结算
		return nil, errors.New("there are no related accounts in the bubble network to generate new transactions, " +
			"so there is no need for settlement")
	} else {
		// 需要结算
		// 保存hash(或在处理任务的时候保存)
		token.SaveSettlementHash(state, hash)
		// 判断当前节点是否是子链运营节点
		if tkc.Plugin.IsSubOpNode {

		}
	}
	// 记录Log日志
	return txResultHandlerWithRes(vm.TokenContractAddr, tkc.Evm, "",
		"", TxSettlement, int(common.NoErr.Code), settlementInfo), nil
}
