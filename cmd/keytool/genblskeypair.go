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

package main

import (
	"encoding/hex"
	"fmt"
	"gopkg.in/urfave/cli.v1"

	"github.com/bubblenet/bubble/crypto/bls"
)

type outputGenblskeypair struct {
	PrivateKey string
	PublicKey  string
}

var commandGenblskeypair = cli.Command{
	Name:      "genblskeypair",
	Usage:     "generate new bls private key pair",
	ArgsUsage: "[  ]",
	Description: `
Generate a new bls private key pair.
`,
	Flags: []cli.Flag{
		jsonFlag,
	},
	Action: func(ctx *cli.Context) error {
		err := bls.Init(int(bls.BLS12_381))
		if err != nil {
			return err
		}
		var privateKey bls.SecretKey
		privateKey.SetByCSPRNG()
		pubKey := privateKey.GetPublicKey()
		out := outputGenblskeypair{
			PrivateKey: hex.EncodeToString(privateKey.GetLittleEndian()),
			PublicKey:  hex.EncodeToString(pubKey.Serialize()),
		}
		if ctx.Bool(jsonFlag.Name) {
			mustPrintJSON(out)
		} else {
			fmt.Println("PrivateKey: ", out.PrivateKey)
			fmt.Println("PublicKey : ", out.PublicKey)
		}
		return nil
	},
}
