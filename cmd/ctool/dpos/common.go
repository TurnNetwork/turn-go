// Copyright 2021 The Bubble Network Authors
// This file is part of bubble.
//
// bubble is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// bubble is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with bubble. If not, see <http://www.gnu.org/licenses/>.

package dpos

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/urfave/cli.v1"

	bubble "github.com/bubblenet/bubble"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/ethclient"
	"github.com/bubblenet/bubble/rlp"
)

func CallDPOSContract(client *ethclient.Client, funcType uint16, params ...interface{}) ([]byte, error) {
	send, to := EncodeDPOS(funcType, params...)
	var msg bubble.CallMsg
	msg.Data = send
	msg.To = &to
	return client.CallContract(context.Background(), msg, nil)
}

// CallMsg contains parameters for contract calls.
type CallMsg struct {
	To   *common.Address // the destination contract (nil for contract creation)
	Data hexutil.Bytes   // input data, usually an ABI-encoded contract method invocation
}

func BuildDPOSContract(funcType uint16, params ...interface{}) ([]byte, error) {
	send, to := EncodeDPOS(funcType, params...)
	var msg CallMsg
	msg.Data = send
	msg.To = &to
	return json.Marshal(msg)
}

func EncodeDPOS(funcType uint16, params ...interface{}) ([]byte, common.Address) {
	par := buildParams(funcType, params...)
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, par)
	if err != nil {
		panic(fmt.Errorf("encode rlp data fail: %v", err))
	}
	return buf.Bytes(), funcTypeToContractAddress(funcType)
}

func buildParams(funcType uint16, params ...interface{}) [][]byte {
	var res [][]byte
	res = make([][]byte, 0)
	fnType, _ := rlp.EncodeToBytes(funcType)
	res = append(res, fnType)
	for _, param := range params {
		val, err := rlp.EncodeToBytes(param)
		if err != nil {
			panic(err)
		}
		res = append(res, val)
	}
	return res
}

func funcTypeToContractAddress(funcType uint16) common.Address {
	toadd := common.ZeroAddr
	switch {
	case 0 < funcType && funcType < 2000:
		toadd = vm.StakingContractAddr
	case funcType >= 2000 && funcType < 3000:
		toadd = vm.GovContractAddr
	case funcType >= 3000 && funcType < 4000:
		toadd = vm.SlashingContractAddr
	case funcType >= 4000 && funcType < 5000:
		toadd = vm.RestrictingContractAddr
	case funcType >= 5000 && funcType < 6000:
		toadd = vm.DelegateRewardPoolAddr
	}
	return toadd
}

func query(c *cli.Context, funcType uint16, params ...interface{}) error {
	url := c.String(rpcUrlFlag.Name)
	if url == "" {
		return errors.New("rpc url not set")
	}
	if c.Bool(jsonFlag.Name) {
		res, err := BuildDPOSContract(funcType, params...)
		if err != nil {
			return err
		}
		fmt.Println(string(res))
		return nil
	} else {
		client, err := ethclient.Dial(url)
		if err != nil {
			return err
		}
		res, err := CallDPOSContract(client, funcType, params...)
		if err != nil {
			return err
		}
		fmt.Println(string(res))
		return nil
	}
}
