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
	"errors"
	"fmt"
	"math/big"
	"reflect"

	"github.com/bubblenet/bubble/accounts/abi"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/byteutil"
	"github.com/bubblenet/bubble/common/math"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/plugin"
)

const (
	TxBindTempPrivateKey       = 7200
	TxBehalfSignature          = 7201
	TxInvalidateTempPrivateKey = 7202
	TxAddLineOfCredit          = 7203
	TxGetLineOfCredit          = 7204
)

const (
	tempPrivateKeyPrefix = "TempPrivateKey"
	lineOfCreditPrefix   = "LineOfCredit"
)

var (
	ErrInvalidPeriod           = common.NewBizError(700000, "invalid period")
	ErrContractCaller          = common.NewBizError(700001, "invalid contract caller")
	ErrNoBindingTempPrivateKey = common.NewBizError(700002, "no binding temporary private key")
	ErrGetGameOperator         = common.NewBizError(700003, "Failed to get game operator")
	ErrGetLineOfCredit         = common.NewBizError(700004, "Failed to get line of credit")
	ErrCallGameContract        = common.NewBizError(700005, "Failed to call game contract")
)

func newCallContractError(errorInfo string) *common.BizError {
	info := "Failed to call game contract: " + errorInfo
	return common.NewBizError(700005, info)
}

func IsTxTxBehalfSignature(input []byte, to common.Address) bool {
	if len(input) == 0 {
		return false
	}

	if !bytes.Equal(to.Bytes(), vm.TempPrivateKeyContractAddr.Bytes()) {
		return false
	}

	var args [][]byte
	if err := rlp.Decode(bytes.NewReader(input), &args); nil != err {
		return false
	}

	fnCode := byteutil.BytesToUint16(args[0])

	return fnCode == TxBehalfSignature
}

func GetLineOfCredit(evm *EVM, workAddress, gameContractAddress common.Address) (lineOfCredit *big.Int, err error) {
	lineOfCreditDbValue := evm.StateDB.GetState(vm.TempPrivateKeyContractAddr, getLineOfCreditDBKey(workAddress, gameContractAddress))
	if nil == lineOfCreditDbValue || len(lineOfCreditDbValue) == 0 {
		contract := NewContract(AccountRef(workAddress), AccountRef(gameContractAddress), big.NewInt(0), uint64(math.MaxUint64/2))
		contract.SetCallCode(&gameContractAddress, evm.StateDB.GetCodeHash(gameContractAddress), evm.StateDB.GetCode(gameContractAddress))
		result, err := RunEvm(evm, contract, crypto.Keccak256([]byte("lineOfCredit()"))[:4])
		if err != nil {
			return big.NewInt(0), err
		}
		lineOfCredit = new(big.Int).SetBytes(result)
	} else {
		lineOfCredit = new(big.Int).SetBytes(lineOfCreditDbValue)
	}

	return
}

func SetLineOfCredit(evm *EVM, workAddress, gameContractAddress common.Address, lineOfCredit *big.Int) {
	evm.StateDB.SetState(vm.TempPrivateKeyContractAddr, getLineOfCreditDBKey(workAddress, gameContractAddress), lineOfCredit.Bytes())
}

func GetGameOperator(evm *EVM, workAddress, gameContractAddress common.Address) (operatorAddress common.Address, err error) {
	contract := NewContract(AccountRef(workAddress), AccountRef(gameContractAddress), big.NewInt(0), uint64(math.MaxUint64/2))
	contract.SetCallCode(&gameContractAddress, evm.StateDB.GetCodeHash(gameContractAddress), evm.StateDB.GetCode(gameContractAddress))
	result, err := RunEvm(evm, contract, crypto.Keccak256([]byte("issuer()"))[:4])
	if err != nil {
		return common.Address{}, err
	}

	return common.BytesToAddress(result), nil
}

func GetRatio(evm *EVM, workAddress, gameContractAddress common.Address) (ratio *big.Int, err error) {
	contract := NewContract(AccountRef(workAddress), AccountRef(gameContractAddress), big.NewInt(0), uint64(math.MaxUint64/2))
	contract.SetCallCode(&gameContractAddress, evm.StateDB.GetCodeHash(gameContractAddress), evm.StateDB.GetCode(gameContractAddress))
	result, err := RunEvm(evm, contract, crypto.Keccak256([]byte("ratio()"))[:4])
	if err != nil {
		return big.NewInt(0), err
	}
	return new(big.Int).SetBytes(result), nil
}

func GetBehalfSignatureParameterAddress(input []byte) (workAddress, gameContractAddress common.Address, err error) {
	var args [][]byte
	if err = rlp.Decode(bytes.NewReader(input), &args); nil != err {
		return
	}

	fnCode := byteutil.BytesToUint16(args[0])
	if fn, ok := (&TempPrivateKeyContract{}).FnSigns()[fnCode]; !ok {
		err = plugin.FuncNotExistErr
		return
	} else {

		//funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
		//fmt.Println("The FuncName is", funcName)

		// the func params type list
		paramList := reflect.TypeOf(fn)
		// the func params len
		paramNum := paramList.NumIn()

		if paramNum != len(args)-1 {
			err = plugin.FnParamsLenErr
			return
		}

		for i := 0; i < paramNum; i++ {
			//fmt.Println("byte:", args[i+1])

			targetType := paramList.In(i).String()
			inputByte := []reflect.Value{reflect.ValueOf(args[i+1])}

			if i == 0 {
				workAddress = (reflect.ValueOf(byteutil.Bytes2X_CMD[targetType]).Call(inputByte)[0]).Interface().(common.Address)
			}

			if i == 1 {
				gameContractAddress = (reflect.ValueOf(byteutil.Bytes2X_CMD[targetType]).Call(inputByte)[0]).Interface().(common.Address)
			}
			err = nil
		}
		return
	}
}

type TempPrivateKeyContract struct {
	Plugin   *plugin.TempPrivateKeyPlugin
	Contract *Contract
	Evm      *EVM
}

func (tpkc *TempPrivateKeyContract) RequiredGas(input []byte) uint64 {
	if checkInputEmpty(input) {
		return 0
	}
	return params.TempPrivateKeyGas
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
		TxAddLineOfCredit:          tpkc.addLineOfCredit,
		TxGetLineOfCredit:          tpkc.getLineOfCredit,
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

	return txResultHandlerWithRes(vm.TempPrivateKeyContractAddr, tpkc.Evm,
		"bindTempPrivateKey", "", TxBindTempPrivateKey, int(common.NoErr.Code), gameContractAddress, tempAddress, period), nil
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
	// check the temporary private key exists
	var (
		storeTempAddressBytes []byte
		storeTempAddress      common.Address
		state                 = tpkc.Evm.StateDB
		err                   error
	)

	dbValue := state.GetState(vm.TempPrivateKeyContractAddr, getTempPrivateKeyDBKey(workAddress, gameContractAddress))
	if nil == dbValue || len(dbValue) < common.AddressLength {
		log.Error("no binding temporary private key")
		err = ErrNoBindingTempPrivateKey
		goto resultHandle
	}

	// check parameter
	storeTempAddressBytes = dbValue[0:common.AddressLength]
	storeTempAddress = common.BytesToAddress(storeTempAddressBytes)
	if !bytes.Equal(tempAddress.Bytes(), storeTempAddress.Bytes()) {
		log.Error("invalid caller")
		err = ErrContractCaller
		goto resultHandle
	}

	state.SetState(vm.TempPrivateKeyContractAddr, getTempPrivateKeyDBKey(workAddress, gameContractAddress), []byte{})

resultHandle:

	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {

			return txResultHandler(vm.TempPrivateKeyContractAddr, tpkc.Evm, "invalidateTempPrivateKey",
				bizErr.Error(), TxInvalidateTempPrivateKey, bizErr)

		} else {
			log.Error("Failed to invalidateTempPrivateKey", "txHash", txHash,
				"blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	return txResultHandlerWithRes(vm.TempPrivateKeyContractAddr, tpkc.Evm,
		"invalidateTempPrivateKey", "", TxInvalidateTempPrivateKey, int(common.NoErr.Code), gameContractAddress, tempAddress), nil
}

// sign on behalf of workAddress
func (tpkc *TempPrivateKeyContract) behalfSignature(workAddress, gameContractAddress common.Address, periodArg []byte, input []byte) ([]byte, error) {
	txHash := tpkc.Evm.StateDB.TxHash()
	blockNumber := tpkc.Evm.Context.BlockNumber
	blockHash := tpkc.Evm.Context.BlockHash
	state := tpkc.Evm.StateDB

	// Calculating gas
	if !tpkc.Contract.UseGas(params.BehalfSignatureGas) {
		return nil, ErrOutOfGas
	}

	// get temporary private key information
	var (
		err                                      error
		tempAddress                              common.Address
		sender                                   = AccountRef(workAddress)
		vmRet, dbValue, tempAddressBytes, period []byte
	)
	dbValue = state.GetState(vm.TempPrivateKeyContractAddr, getTempPrivateKeyDBKey(workAddress, gameContractAddress))
	if nil == dbValue || len(dbValue) < common.AddressLength {
		log.Error("no binding temporary private key")
		err = ErrNoBindingTempPrivateKey
		goto resultHandle
	}
	tempAddressBytes = dbValue[0:common.AddressLength]
	tempAddress = common.BytesToAddress(tempAddressBytes)
	period = dbValue[common.AddressLength:]

	log.Debug("Call behalfSignature of TempPrivateKeyContract", "blockHash", blockHash, "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "workAddress", workAddress, "gameContractAddress", gameContractAddress, "tempAddress", tempAddress,
		"period", hex.EncodeToString(periodArg))

	// check period
	if nil != period && len(period) != 0 && !bytes.Equal(period, periodArg) {
		log.Error("invalid period")
		err = ErrInvalidPeriod
		goto resultHandle
	}

	// check from
	if !bytes.Equal(tempAddress.Bytes(), tpkc.Contract.CallerAddress.Bytes()) {
		log.Error("invalid caller")
		err = ErrContractCaller
		goto resultHandle
	}

	// Calculating gas
	if !tpkc.Contract.UseGas(params.CallStipend) {
		return nil, ErrOutOfGas
	}

	// run contract invoke
	vmRet, tpkc.Contract.Gas, err = tpkc.Evm.Call(sender, gameContractAddress, input, tpkc.Contract.Gas, big.NewInt(0))
	if err != nil {
		log.Error("Failed to call game contract", "gameContractAddress", gameContractAddress, "err", err)
		if errors.Is(err, ErrOutOfGas) {
			return nil, ErrOutOfGas
		}
		if errors.Is(err, ErrExecutionReverted) {
			reason, errUnpack := abi.UnpackRevert(vmRet)
			info := "execution reverted"
			if errUnpack == nil {
				info = fmt.Sprintf("execution reverted: %v", reason)
			}
			err = newCallContractError(info)
		} else {
			err = newCallContractError(err.Error())
		}

	}

resultHandle:

	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {

			return txResultHandler(vm.TempPrivateKeyContractAddr, tpkc.Evm, "behalfSignature",
				bizErr.Error(), TxBehalfSignature, bizErr)

		} else {
			log.Error("Failed to behalfSignature", "txHash", txHash,
				"blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	return txResultHandlerWithRes(vm.TempPrivateKeyContractAddr, tpkc.Evm,
		"behalfSignature", "", TxBehalfSignature, int(common.NoErr.Code), workAddress, gameContractAddress, tempAddress, periodArg, input, vmRet), nil
}

// add line of credit
func (tpkc *TempPrivateKeyContract) addLineOfCredit(gameContractAddress, workAddress common.Address, addValue *big.Int) ([]byte, error) {
	txHash := tpkc.Evm.StateDB.TxHash()
	blockNumber := tpkc.Evm.Context.BlockNumber
	operatorAddress := tpkc.Contract.CallerAddress
	blockHash := tpkc.Evm.Context.BlockHash
	log.Debug("Call addLineOfCredit of TempPrivateKeyContract", "blockHash", blockHash, "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "workAddress", operatorAddress, "gameContractAddress", gameContractAddress)

	// Calculating gas
	if !tpkc.Contract.UseGas(params.AddLineOfCreditGas) {
		return nil, ErrOutOfGas
	}

	var (
		err                     error
		lineOfCredit            *big.Int
		contractOperatorAddress common.Address
	)

	// Call handling logic
	// check operator
	contractOperatorAddress, err = GetGameOperator(tpkc.Evm, workAddress, gameContractAddress)
	if err != nil {
		log.Error("Failed to get game contract operator", "gameContractAddress", gameContractAddress, "err", err)
		err = ErrGetGameOperator
		goto resultHandle
	}
	if !bytes.Equal(operatorAddress.Bytes(), contractOperatorAddress.Bytes()) {
		log.Error("operatorAddress and contractOperatorAddress are not equal", "operatorAddress", operatorAddress, "contractOperatorAddress", contractOperatorAddress)
		err = ErrContractCaller
		goto resultHandle
	}

	// set new line of credit
	lineOfCredit, err = GetLineOfCredit(tpkc.Evm, workAddress, gameContractAddress)
	if err != nil {
		log.Error("Failed to get line of credit", "gameContractAddress", gameContractAddress, "err", err)
		err = ErrGetLineOfCredit
		goto resultHandle
	}

	lineOfCredit = lineOfCredit.Add(lineOfCredit, addValue)
	tpkc.Evm.StateDB.SetState(vm.TempPrivateKeyContractAddr, getLineOfCreditDBKey(workAddress, gameContractAddress), lineOfCredit.Bytes())

resultHandle:

	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {

			return txResultHandler(vm.TempPrivateKeyContractAddr, tpkc.Evm, "addLineOfCredit",
				bizErr.Error(), TxAddLineOfCredit, bizErr)

		} else {
			log.Error("Failed to addLineOfCredit", "txHash", txHash,
				"blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	return txResultHandlerWithRes(vm.TempPrivateKeyContractAddr, tpkc.Evm,
		"addLineOfCredit", "", TxAddLineOfCredit, int(common.NoErr.Code), workAddress, gameContractAddress, addValue, lineOfCredit), nil
}

// get line of credit
func (tpkc *TempPrivateKeyContract) getLineOfCredit(gameContractAddress common.Address) ([]byte, error) {

	txHash := tpkc.Evm.StateDB.TxHash()
	blockNumber := tpkc.Evm.Context.BlockNumber
	workAddress := tpkc.Contract.CallerAddress
	blockHash := tpkc.Evm.Context.BlockHash
	log.Debug("Call getLineOfCredit of TempPrivateKeyContract", "blockHash", blockHash, "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "workAddress", workAddress, "gameContractAddress", gameContractAddress)

	lineOfCredit, err := GetLineOfCredit(tpkc.Evm, workAddress, gameContractAddress)
	if err != nil {
		log.Error("Failed to get line of credit", "gameContractAddress", gameContractAddress, "err", err)
		err = ErrGetLineOfCredit
		bizErr, _ := err.(*common.BizError)
		return callResultHandler(tpkc.Evm, "getLineOfCredit",
			nil, bizErr), nil
	}

	return callResultHandler(tpkc.Evm, "getLineOfCredit",
		lineOfCredit.String(), nil), nil
}
