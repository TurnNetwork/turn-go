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
	"testing"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/rlp"
)

func TestIsTxTxBehalfSignature(t *testing.T) {
	gameContractAddress := common.HexToAddress("0x195667cDeFCad94C521BdfF0Bf85079761E0f8F3")
	tempAddress := common.HexToAddress("0x195667cDeFCad94C521BdfF0Bf85079761E0f8F3")
	period := []byte("Hello World")
	params := make([][]byte, 0)
	fnType, _ := rlp.EncodeToBytes(uint16(7200))
	gameContractAddressBytes, _ := rlp.EncodeToBytes(gameContractAddress)
	tempAddressBytes, _ := rlp.EncodeToBytes(tempAddress)
	periodBytes, _ := rlp.EncodeToBytes(period)
	params = append(params, fnType)
	params = append(params, gameContractAddressBytes)
	params = append(params, tempAddressBytes)
	params = append(params, periodBytes)
	buf := new(bytes.Buffer)
	rlp.Encode(buf, params)

	input := buf.Bytes()
	bTemp := IsTxTxBehalfSignature(input, vm.TempPrivateKeyContractAddr)
	if !bTemp {
		t.Fatal("test IsTxTxBehalfSignature fail")
	} else {
		t.Log("test IsTxTxBehalfSignature sucess")
	}
}
