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
	"github.com/bubblenet/bubble/accounts/abi"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/x/plugin"
	"github.com/bubblenet/bubble/x/xcom"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

func execBubbleContract(input []byte, command map[uint16]interface{}) (ret []byte, err error) {
	// verify the tx data by contracts method
	_, fn, params, err := plugin.VerifyTxData(input, command)
	if nil != err {
		log.Error("Failed to verify contract tx before exec", "err", err)
		return xcom.NewResult(common.InvalidParameter, nil), err
	}

	// execute contracts method
	result := reflect.ValueOf(fn).Call(params)
	switch errtyp := result[1].Interface().(type) {
	case *common.BizError:
		log.Error("Failed to execute contract tx", "err", errtyp)
		return xcom.NewResult(errtyp, nil), errtyp
	case error:
		log.Error("Failed to execute contract tx", "err", errtyp)
		return xcom.NewResult(common.InternalError, nil), errtyp
	default:
	}
	return result[0].Bytes(), nil
}

func txResultHandler(contractAddr common.Address, evm *EVM, title, reason string, fncode int, errCode *common.BizError) ([]byte, error) {
	event := strconv.Itoa(fncode)
	receipt := strconv.Itoa(int(errCode.Code))
	blockNumber := evm.Context.BlockNumber.Uint64()
	if errCode.Code != 0 {
		txHash := evm.StateDB.TxHash()
		log.Error("Failed to "+title, "txHash", txHash.Hex(),
			"blockNumber", blockNumber, "receipt: ", receipt, "the reason", reason)
	}
	xcom.AddLogWithRes(evm.StateDB, blockNumber, contractAddr, event, receipt, nil)
	if errCode.Code == common.NoErr.Code {
		return []byte(receipt), nil
	}
	return []byte(receipt), errCode
}

func txResultHandlerWithRes(contractAddr common.Address, evm *EVM, title, reason string, fncode, errCode int, res ...interface{}) []byte {
	event := strconv.Itoa(fncode)
	receipt := strconv.Itoa(errCode)
	blockNumber := evm.Context.BlockNumber.Uint64()
	if errCode != 0 {
		txHash := evm.StateDB.TxHash()
		log.Error("Failed to "+title, "txHash", txHash.Hex(),
			"blockNumber", blockNumber, "receipt: ", receipt, "the reason", reason)
	}
	xcom.AddLogWithRes(evm.StateDB, blockNumber, contractAddr, event, receipt, res...)
	return []byte(receipt)
}

func callResultHandler(evm *EVM, title string, resultValue interface{}, err *common.BizError) []byte {
	txHash := evm.StateDB.TxHash()
	blockNumber := evm.Context.BlockNumber.Uint64()

	if nil != err {
		log.Error("Failed to "+title, "txHash", txHash.Hex(),
			"blockNumber", blockNumber, "the reason", err.Error())
		return xcom.NewResult(err, nil)
	}

	if IsBlank(resultValue) {
		return xcom.NewResult(common.NotFound, nil)
	}

	log.Debug("Call "+title+" finished", "blockNumber", blockNumber,
		"txHash", txHash, "result", resultValue)
	return xcom.NewResult(nil, resultValue)
}

func IsBlank(i interface{}) bool {
	defer func() {
		recover()
	}()

	typ := reflect.TypeOf(i)
	val := reflect.ValueOf(i)
	if typ == nil {
		return true
	} else {
		if typ.Kind() == reflect.Slice {
			return val.Len() == 0
		}
		if typ.Kind() == reflect.Map {
			return val.Len() == 0
		}
	}
	return val.IsNil()
}

func checkInputEmpty(input []byte) bool {
	if len(input) == 0 {
		return true
	} else {
		return false
	}
}

// encodeMintFuncCall: Generate function signatures for ERC20 contract minting transactions
func encodeMintFuncCall(to common.Address, amount *big.Int) ([]byte, error) {
	// Create a contract ABI resolver
	encodeABI, err := abi.JSON(strings.NewReader(`[
		{
			"inputs": [
				{
				  "internalType": "address",
				  "name": "to",
				  "type": "address"
				},
				{
				  "internalType": "uint256",
				  "name": "amount",
				  "type": "uint256"
				}
			],
			"name": "mint",
			"type": "function"
		}
	]`))
	if err != nil {
		return nil, err
	}

	functionName := "mint"
	// Encode function call data
	data, err := encodeABI.Pack(functionName, to, amount)
	if err != nil {
		return nil, err
	}

	// Convert the encoded data to a hexadecimal string
	// encodedData := common.Bytes2Hex(data)
	// fmt.Println(encodedData)
	return data, nil
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

// RunEvm Execute the EVM contract code
func RunEvm(evm *EVM, contract *Contract, input []byte) ([]byte, error) {
	if nil == evm || nil == contract {
		log.Error("Run Evm failed", "evm or contract is nil")
		return nil, errors.New("evm or contract is nil")
	}
	for _, interpreter := range evm.interpreters {
		if interpreter.CanRun(contract.Code) {
			// Determine and set the current virtual machine
			if evm.interpreter != interpreter {
				// Ensure that the interpreter pointer is set back
				// to its current value upon return.
				defer func(i Interpreter) {
					evm.interpreter = i
				}(evm.interpreter)
				evm.interpreter = interpreter
			}
			// Executing the virtual machine
			ret, err := interpreter.Run(contract, input, false)
			if err != nil {
				log.Error("Run Evm failed", "ret", ret, "error", err)
				// return ret, err
			}
			// Execution completes one of the EVM or WASM and returns
			return ret, err
		}
	}
	return nil, errors.New("there is no executable EVM or WASM interpreter")
}
