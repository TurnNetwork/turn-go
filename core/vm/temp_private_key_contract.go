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
	"encoding/hex"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/plugin"
)

const (
	TxBindTempPrivateKey       = 7000
	TxBehalfSignature          = 7001
	TxInvalidateTempPrivateKey = 7002
)

const (
	tempPrivateKeyPrefix = "TempPrivateKey"
	lineOfCreditPrefix   = "LineOfCredit"
)

var (
	ErrInvalidPeriod      = common.NewBizError(700000, "invalid period")
	ErrGetCurrentGasPrice = common.NewBizError(700001, "Failed to get current gas price")
	ErrGetGameOperator    = common.NewBizError(700002, "Failed to get game operator")
	ErrGetLineOfCredit    = common.NewBizError(700003, "Failed to get line of credit")
	ErrCallGameContract   = common.NewBizError(700004, "Failed to call game contract")
)

type TempPrivateKeyContract struct {
	Plugin   *plugin.TempPrivateKeyPlugin
	Contract *Contract
	Evm      *EVM
}

func (tpkc *TempPrivateKeyContract) RequiredGas(input []byte) uint64 {
	if checkInputEmpty(input) {
		return 0
	}
	return params.TokenGas
}

func (tpkc *TempPrivateKeyContract) Run(input []byte) ([]byte, error) {
	if checkInputEmpty(input) {
		return nil, nil
	}
	return execBubbleContract(input, tpkc.FnSigns())
}

func (tpkc *TempPrivateKeyContract) CheckGasPrice(gasPrice *big.Int, fcode uint16) error {
	return nil
}

func (tpkc *TempPrivateKeyContract) FnSigns() map[uint16]interface{} {
	fnSigns := tpkc.FnSignsV1()
	return fnSigns
}

func (tpkc *TempPrivateKeyContract) FnSignsV1() map[uint16]interface{} {
	return map[uint16]interface{}{
		TxBindTempPrivateKey:       tpkc.bindTempPrivateKey,
		TxBehalfSignature:          tpkc.behalfSignature,
		TxInvalidateTempPrivateKey: tpkc.invalidateTempPrivateKey,
	}
}

func getTempPrivateKeyDBKey(workAddress, gameContractAddress common.Address) []byte {
	key := append(workAddress.Bytes(), gameContractAddress.Bytes()...)
	prefix := []byte(tempPrivateKeyPrefix)
	return append(prefix, key...)
}

func getTempPrivateKeyDBValue(tempAddress common.Address, period []byte) []byte {
	return append(tempAddress.Bytes(), period...)
}

func getLineOfCreditDBKey(workAddress, gameContractAddress common.Address) []byte {
	key := append(workAddress.Bytes(), gameContractAddress.Bytes()...)
	prefix := []byte(lineOfCreditPrefix)
	return append(prefix, key...)
}

// Set temporary private key
func (tpkc *TempPrivateKeyContract) bindTempPrivateKey(gameContractAddress, tempAddress common.Address, period []byte) ([]byte, error) {
	txHash := tpkc.Evm.StateDB.TxHash()
	blockNumber := tpkc.Evm.Context.BlockNumber
	workAddress := tpkc.Contract.CallerAddress
	blockHash := tpkc.Evm.Context.BlockHash
	log.Debug("Call bindTempPrivateKey of TempPrivateKeyContract", "blockHash", blockHash, "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "workAddress", workAddress, "gameContractAddress", gameContractAddress, "tempAddress", tempAddress,
		"period", hex.EncodeToString(period))

	// Calculating gas
	if !tpkc.Contract.UseGas(params.BindTempPrivateKeyGas) {
		return nil, ErrOutOfGas
	}

	// Call handling logic
	state := tpkc.Evm.StateDB
	state.SetState(vm.TempPrivateKeyContractAddr, getTempPrivateKeyDBKey(workAddress, gameContractAddress), getTempPrivateKeyDBValue(tempAddress, period))

	return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "bindTempPrivateKey", "", TxBindTempPrivateKey, nil)
}

// invalidate temporary private key
func (tpkc *TempPrivateKeyContract) invalidateTempPrivateKey(gameContractAddress, tempAddress common.Address) ([]byte, error) {
	txHash := tpkc.Evm.StateDB.TxHash()
	blockNumber := tpkc.Evm.Context.BlockNumber
	workAddress := tpkc.Contract.CallerAddress
	blockHash := tpkc.Evm.Context.BlockHash
	log.Debug("Call invalidateTempPrivateKey of TempPrivateKeyContract", "blockHash", blockHash, "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "workAddress", workAddress, "gameContractAddress", gameContractAddress, "tempAddress", tempAddress)

	// Calculating gas
	if !tpkc.Contract.UseGas(params.InvalidateTempPrivateKeyGas) {
		return nil, ErrOutOfGas
	}

	// Call handling logic
	state := tpkc.Evm.StateDB
	state.SetState(vm.TempPrivateKeyContractAddr, getTempPrivateKeyDBKey(workAddress, gameContractAddress), []byte{})
	return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "invalidateTempPrivateKey", "", TxInvalidateTempPrivateKey, nil)
}

func (tpkc *TempPrivateKeyContract) defaultGasPrice() *big.Int {
	return big.NewInt(50000000000)
}

// sign on behalf of workAddress
func (tpkc *TempPrivateKeyContract) behalfSignature(workAddress, gameContractAddress common.Address, periodArg []byte, input []byte) ([]byte, error) {
	txHash := tpkc.Evm.StateDB.TxHash()
	blockNumber := tpkc.Evm.Context.BlockNumber
	blockHash := tpkc.Evm.Context.BlockHash
	state := tpkc.Evm.StateDB

	// get temporary private key information
	dbValue := state.GetState(vm.GovContractAddr, getTempPrivateKeyDBKey(workAddress, gameContractAddress))
	tempAddressBytes := dbValue[0:common.AddressLength]
	tempAddress := common.BytesToAddress(tempAddressBytes)
	period := dbValue[common.AddressLength:]

	log.Debug("Call behalfSignature of TempPrivateKeyContract", "blockHash", blockHash, "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "workAddress", workAddress, "gameContractAddress", gameContractAddress, "tempAddress", tempAddress)

	// get gas price
	gasprice, err := tpkc.Plugin.SuggestPrice()
	if err != nil {
		// gasprice = tpkc.defaultGasPrice()
		if !tpkc.Contract.UseGas(params.BehalfSignatureGas) {
			return nil, ErrOutOfGas
		}
		log.Error("Failed to behal fSignatur", "error", err)
		return nil, ErrGetCurrentGasPrice
	}

	// get game contract operator
	operatorAddress, err := tpkc.Plugin.GetGameContractOperator(tpkc.Evm.Context.BlockNumber.Uint64(), workAddress, gameContractAddress, gasprice)
	if err != nil {
		if !tpkc.Contract.UseGas(params.BehalfSignatureGas) {
			return nil, ErrOutOfGas
		}
		log.Error("Failed to behal fSignatur", "error", err)
		return nil, ErrGetGameOperator
	}

	// Call handling logic
	if bytes.Compare(period, periodArg) != 0 {
		log.Error("invalid period")
		err = ErrInvalidPeriod
	}

	// get line of credit
	lineOfCredit := big.NewInt(0)
	lineOfCreditDbValue := state.GetState(vm.TempPrivateKeyContractAddr, getLineOfCreditDBKey(workAddress, gameContractAddress))
	if nil == lineOfCreditDbValue || len(lineOfCreditDbValue) == 0 {
		lineOfCredit, err = tpkc.Plugin.GetLineOfCredit(tpkc.Evm.Context.BlockNumber.Uint64(), workAddress, gameContractAddress, gasprice)
		if err != nil {
			err = ErrGetLineOfCredit
		}
	} else {
		lineOfCredit.SetBytes(lineOfCreditDbValue)
	}

	workAddressBalance := state.GetBalance(workAddress)
	operatorAddressBalance := state.GetBalance(operatorAddress)

	realLineOfCredit := big.NewInt(0).SetBytes(lineOfCredit.Bytes())
	if lineOfCredit.Cmp(operatorAddressBalance) > 0 {
		realLineOfCredit = operatorAddressBalance
	}

	maxUsedValue := workAddressBalance.Add(workAddressBalance, realLineOfCredit)
	gasLimit := maxUsedValue.Div(maxUsedValue, gasprice).Uint64()

	// run contract invoke
	sender := AccountRef(workAddress)
	_, leftOverGas, err := tpkc.Evm.Call(sender, gameContractAddress, input, gasLimit, big.NewInt(0))
	if err != nil {
		err = ErrCallGameContract
	}
	usedGas := gasLimit - leftOverGas
	usedValue := new(big.Int).Mul(new(big.Int).SetUint64(usedGas), gasprice)
	if usedValue.Cmp(realLineOfCredit) > 0 {
		state.SubBalance(operatorAddress, realLineOfCredit)
		lineOfCredit = lineOfCredit.Sub(lineOfCredit, realLineOfCredit)
		state.SubBalance(workAddress, usedValue.Sub(usedValue, realLineOfCredit))
	} else {
		state.SubBalance(operatorAddress, usedValue)
		lineOfCredit = lineOfCredit.Sub(lineOfCredit, usedValue)
	}

	state.SetState(vm.TempPrivateKeyContractAddr, getLineOfCreditDBKey(workAddress, gameContractAddress), lineOfCredit.Bytes())

	return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "behalfSignature", "", TxBehalfSignature, nil)
}
