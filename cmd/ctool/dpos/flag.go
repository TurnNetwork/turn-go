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

import "gopkg.in/urfave/cli.v1"

var (
	rpcUrlFlag = cli.StringFlag{
		Name:  "rpcurl",
		Usage: "the rpc url",
	}

	jsonFlag = cli.BoolFlag{
		Name:  "json",
		Usage: "print raw transaction",
	}
)
