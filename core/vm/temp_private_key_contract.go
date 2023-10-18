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
	"encoding/hex"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/plugin"
)

const (
	TxBindTempPrivateKey       = 6000
	TxChangeTempPrivateKey     = 6001
	TxBehalfSignature          = 6002
	TxInvalidateTempPrivateKey = 6003
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
		TxChangeTempPrivateKey:     tpkc.changeTempPrivateKey,
		TxInvalidateTempPrivateKey: tpkc.invalidateTempPrivateKey,
		TxBehalfSignature:          tpkc.behalfSignature,
	}
}

func getDBKey(workAddress, gameContractAddress common.Address) []byte {
	return append(workAddress.Bytes(), gameContractAddress.Bytes()...)
}

func getDBValue(tempAddress common.Address, period []byte) []byte {
	return append(tempAddress.Bytes(), period...)
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
	state.SetState(vm.TempPrivateKeyContractAddr, getDBKey(workAddress, gameContractAddress), getDBValue(tempAddress, period))

	// // estimate gas
	// if err == nil && ret == nil {
	// 	return nil, nil
	// }
	// if nil != err {
	// 	if bizErr, ok := err.(*common.BizError); ok {
	// 		return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "bindTempPrivateKey", bizErr.Error(), TxBindTempPrivateKey, bizErr)
	// 	} else {
	// 		log.Error("Failed to bindTempPrivateKey", "txHash", txHash, "blockNumber", blockNumber, "err", err)
	// 		return nil, err
	// 	}
	// }

	return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "bindTempPrivateKey", "", TxBindTempPrivateKey, nil)
}

// change temporary private key
func (tpkc *TempPrivateKeyContract) changeTempPrivateKey(gameContractAddress, tempAddress common.Address, period []byte) ([]byte, error) {
	txHash := tpkc.Evm.StateDB.TxHash()
	blockNumber := tpkc.Evm.Context.BlockNumber
	workAddress := tpkc.Contract.CallerAddress
	blockHash := tpkc.Evm.Context.BlockHash
	log.Debug("Call changeTempPrivateKey of TempPrivateKeyContract", "blockHash", blockHash, "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "workAddress", workAddress, "gameContractAddress", gameContractAddress, "tempAddress", tempAddress,
		"period", hex.EncodeToString(period))
	// Calculating gas
	if !tpkc.Contract.UseGas(params.ChangeTempPrivateKeyGas) {
		return nil, ErrOutOfGas
	}

	// Call handling logic
	state := tpkc.Evm.StateDB
	state.SetState(vm.TempPrivateKeyContractAddr, getDBKey(workAddress, gameContractAddress), getDBValue(tempAddress, period))

	// // estimate gas
	// if err == nil && ret == nil {
	// 	return nil, nil
	// }
	// if nil != err {
	// 	if bizErr, ok := err.(*common.BizError); ok {
	// 		return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "changeTempPrivateKey", bizErr.Error(), TxChangeTempPrivateKey, bizErr)
	// 	} else {
	// 		log.Error("Failed to changeTempPrivateKey", "txHash", txHash, "blockNumber", blockNumber, "err", err)
	// 		return nil, err
	// 	}
	// }

	return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "changeTempPrivateKey", "", TxChangeTempPrivateKey, nil)
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
	state.SetState(vm.TempPrivateKeyContractAddr, getDBKey(workAddress, gameContractAddress), []byte{})

	// // estimate gas
	// if err == nil && ret == nil {
	// 	return nil, nil
	// }
	// if nil != err {
	// 	if bizErr, ok := err.(*common.BizError); ok {
	// 		return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "invalidateTempPrivateKey", bizErr.Error(), TxInvalidateTempPrivateKey, bizErr)
	// 	} else {
	// 		log.Error("Failed to invalidateTempPrivateKey", "txHash", txHash, "blockNumber", blockNumber, "err", err)
	// 		return nil, err
	// 	}
	// }

	return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "invalidateTempPrivateKey", "", TxInvalidateTempPrivateKey, nil)
}

// sign on behalf of workAddress

func (tpkc *TempPrivateKeyContract) defaultGasPrice() (*big.Int, error) {
	return big.NewInt(int64(0)), nil
}

func (tpkc *TempPrivateKeyContract) behalfSignature(workAddress, gameContractAddress, gamePublisherAddress common.Address, periodArg []byte, callData []byte, signedData []byte) ([]byte, error) {
	txHash := tpkc.Evm.StateDB.TxHash()
	blockNumber := tpkc.Evm.Context.BlockNumber
	blockHash := tpkc.Evm.Context.BlockHash
	state := tpkc.Evm.StateDB
	dbValue := state.GetState(vm.GovContractAddr, getDBKey(workAddress, gameContractAddress))
	tempAddressBytes := dbValue[0:common.AddressLength]
	tempAddress := common.BytesToAddress(tempAddressBytes)
	period := dbValue[common.AddressLength:]
	log.Debug("Call behalfSignature of TempPrivateKeyContract", "blockHash", blockHash, "txHash", txHash.Hex(),
		"blockNumber", blockNumber.Uint64(), "workAddress", workAddress, "gameContractAddress", gameContractAddress, "tempAddress", tempAddress)
	// Calculating gas
	if !tpkc.Contract.UseGas(params.InvalidateTempPrivateKeyGas) {
		return nil, ErrOutOfGas
	}

	// Call handling logic
	periodArg = period
	// check temp address
	// get gas price

	// estimate gas
	// if err == nil && ret == nil {
	// 	return nil, nil
	// }
	// if nil != err {
	// 	if bizErr, ok := err.(*common.BizError); ok {
	// 		return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "behalfSignature", bizErr.Error(), TxBehalfSignature, bizErr)
	// 	} else {
	// 		log.Error("Failed to behalfSignature", "txHash", txHash, "blockNumber", blockNumber, "err", err)
	// 		return nil, err
	// 	}
	// }

	return txResultHandler(vm.TokenContractAddr, tpkc.Evm, "behalfSignature", "", TxBehalfSignature, nil)
}
