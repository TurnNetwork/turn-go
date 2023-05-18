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
	"github.com/bubblenet/bubble/p2p/discover"
	"gopkg.in/urfave/cli.v1"
)

var (
	RewardCmd = cli.Command{
		Name:  "reward",
		Usage: "use for reward",
		Subcommands: []cli.Command{
			getDelegateRewardCmd,
		},
	}
	getDelegateRewardCmd = cli.Command{
		Name:   "getDelegateReward",
		Usage:  "5100,query account not withdrawn commission rewards at each node,parameter:nodeList(can empty)",
		Action: getDelegateReward,
		Flags:  []cli.Flag{rpcUrlFlag, nodeList, jsonFlag},
	}
	nodeList = cli.StringSliceFlag{
		Name:  "nodeList",
		Usage: "node list,may empty",
	}
)

func getDelegateReward(c *cli.Context) error {
	nodeIDlist := c.StringSlice(nodeList.Name)
	idlist := make([]discover.NodeID, 0)
	for _, node := range nodeIDlist {
		nodeid, err := discover.HexID(node)
		if err != nil {
			return err
		}
		idlist = append(idlist, nodeid)
	}
	return query(c, 5100, idlist)
}
