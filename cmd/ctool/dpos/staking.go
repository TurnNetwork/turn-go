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
	"errors"
	"gopkg.in/urfave/cli.v1"

	"github.com/bubblenet/bubble/common"

	"github.com/bubblenet/bubble/p2p/enode"
)

var (
	StakingCmd = cli.Command{
		Name:  "staking",
		Usage: "use for staking",
		Subcommands: []cli.Command{
			GetVerifierListCmd,
			getValidatorListCmd,
			getCandidateListCmd,
			getRelatedListByDelAddrCmd,
			getDelegateInfoCmd,
			getCandidateInfoCmd,
			getDelegationLockCmd,
			getPackageRewardCmd,
			getStakingRewardCmd,
			getAvgPackTimeCmd,
		},
	}
	GetVerifierListCmd = cli.Command{
		Name:   "getVerifierList",
		Usage:  "1100,query the validator queue of the current settlement epoch",
		Action: getVerifierList,
		Flags:  []cli.Flag{rpcUrlFlag, jsonFlag},
	}
	getValidatorListCmd = cli.Command{
		Name:   "getValidatorList",
		Usage:  "1101,query the list of validators in the current consensus round",
		Action: getValidatorList,
		Flags:  []cli.Flag{rpcUrlFlag, jsonFlag},
	}
	getCandidateListCmd = cli.Command{
		Name:   "getCandidateList",
		Usage:  "1102,Query the list of all real-time candidates",
		Action: getCandidateList,
		Flags:  []cli.Flag{rpcUrlFlag, jsonFlag},
	}
	getRelatedListByDelAddrCmd = cli.Command{
		Name:   "getRelatedListByDelAddr",
		Usage:  "1103,Query the NodeID and staking Id of the node entrusted by the current account address,parameter:add",
		Action: getRelatedListByDelAddr,
		Flags:  []cli.Flag{rpcUrlFlag, addFlag, jsonFlag},
	}
	getDelegateInfoCmd = cli.Command{
		Name:   "getDelegateInfo",
		Usage:  "1104,Query the delegation information of the current single node,parameter:stakingBlock,address,nodeid",
		Action: getDelegateInfo,
		Flags:  []cli.Flag{rpcUrlFlag, stakingBlockNumFlag, addFlag, nodeIdFlag, jsonFlag},
	}
	getCandidateInfoCmd = cli.Command{
		Name:   "getCandidateInfo",
		Usage:  "1105,Query the staking information of the current node,parameter:nodeid",
		Action: getCandidateInfo,
		Flags:  []cli.Flag{rpcUrlFlag, nodeIdFlag, jsonFlag},
	}
	getDelegationLockCmd = cli.Command{
		Name:   "getDelegationLock",
		Usage:  "1106,Query the delegation lock information of the current account,parameter:address",
		Action: getDelegationLock,
		Flags:  []cli.Flag{rpcUrlFlag, addFlag, jsonFlag},
	}
	getPackageRewardCmd = cli.Command{
		Name:   "getPackageReward",
		Usage:  "1200,query the block reward of the current settlement epoch",
		Action: getPackageReward,
		Flags:  []cli.Flag{rpcUrlFlag, jsonFlag},
	}
	getStakingRewardCmd = cli.Command{
		Name:   "getStakingReward",
		Usage:  "1201,query the staking reward of the current settlement epoch",
		Action: getStakingReward,
		Flags:  []cli.Flag{rpcUrlFlag, jsonFlag},
	}
	getAvgPackTimeCmd = cli.Command{
		Name:   "getAvgPackTime",
		Usage:  "1202,average time to query packaged blocks",
		Action: getAvgPackTime,
		Flags:  []cli.Flag{rpcUrlFlag, jsonFlag},
	}
	addFlag = cli.StringFlag{
		Name:  "address",
		Usage: "account address",
	}
	stakingBlockNumFlag = cli.Uint64Flag{
		Name:  "stakingBlock",
		Usage: "block height when staking is initiated",
	}
	nodeIdFlag = cli.StringFlag{
		Name:  "nodeid",
		Usage: "node id",
	}
)

func getVerifierList(c *cli.Context) error {
	return query(c, 1100)
}

func getValidatorList(c *cli.Context) error {
	return query(c, 1101)
}

func getCandidateList(c *cli.Context) error {
	return query(c, 1102)
}

func getRelatedListByDelAddr(c *cli.Context) error {
	addstring := c.String(addFlag.Name)
	if addstring == "" {
		return errors.New("The Del's account address is not set")
	}
	add, err := common.StringToAddress(addstring)
	if err != nil {
		return err
	}
	return query(c, 1103, add)
}

func getDelegateInfo(c *cli.Context) error {
	addstring := c.String(addFlag.Name)
	if addstring == "" {
		return errors.New("The Del's account address is not set")
	}
	add, err := common.StringToAddress(addstring)
	if err != nil {
		return err
	}
	nodeIDstring := c.String(nodeIdFlag.Name)
	if nodeIDstring == "" {
		return errors.New("The verifier's node ID is not set")
	}
	nodeid, err := enode.HexIDv0(nodeIDstring)
	if err != nil {
		return err
	}
	stakingBlockNum := c.Uint64(stakingBlockNumFlag.Name)
	return query(c, 1104, stakingBlockNum, add, nodeid)
}

func getCandidateInfo(c *cli.Context) error {
	nodeIDstring := c.String(nodeIdFlag.Name)
	if nodeIDstring == "" {
		return errors.New("The verifier's node ID is not set")
	}
	nodeid, err := enode.HexIDv0(nodeIDstring)
	if err != nil {
		return err
	}
	return query(c, 1105, nodeid)
}

func getDelegationLock(c *cli.Context) error {
	addstring := c.String(addFlag.Name)
	if addstring == "" {
		return errors.New("The Del's account address is not set")
	}
	add, err := common.StringToAddress(addstring)
	if err != nil {
		return err
	}
	return query(c, 1106, add)
}

func getPackageReward(c *cli.Context) error {
	return query(c, 1200)
}

func getStakingReward(c *cli.Context) error {
	return query(c, 1201)
}

func getAvgPackTime(c *cli.Context) error {
	return query(c, 1202)
}
