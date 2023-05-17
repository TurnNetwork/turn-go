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

package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
	"sort"

	"github.com/bubblenet/bubble/cmd/ctool/core"

	"github.com/bubblenet/bubble/cmd/ctool/dpos"
)

var (
	app *cli.App
)

func init() {
	app = cli.NewApp()

	// Initialize the CLI app
	app.Commands = []cli.Command{
		core.DeployCmd,
		core.InvokeCmd,
		core.SendTransactionCmd,
		core.SendRawTransactionCmd,
		core.GetTxReceiptCmd,
		core.StabilityCmd,
		core.StabPrepareCmd,
		core.AnalyzeStressTestCmd,
		dpos.GovCmd,
		dpos.SlashingCmd,
		dpos.StakingCmd,
		dpos.RestrictingCmd,
		dpos.RewardCmd,
	}

	app.Name = "ctool"
	app.Version = "1.0.0"

	sort.Sort(cli.CommandsByName(app.Commands))
	app.After = func(ctx *cli.Context) error {
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
