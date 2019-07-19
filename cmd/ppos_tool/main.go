// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/x/restricting"
)

var (
	jsonFlag      = flag.String("json", "", "Json file path for contract parameters")
	funcTypesFlag = flag.String("funcTypes", "", "A collection of contract interface types")
)

// createStaking
type Ppos_1000 struct {
	Typ            uint16
	BenefitAddress common.Address
	NodeId         discover.NodeID
	ExternalId     string
	NodeName       string
	Website        string
	Details        string
	Amount         *big.Int
	ProcessVersion uint32
}

// editorCandidate
type Ppos_1001 struct {
	BenefitAddress common.Address
	NodeId         discover.NodeID
	ExternalId     string
	NodeName       string
	Website        string
	Details        string
}

// increaseStaking
type Ppos_1002 struct {
	NodeId discover.NodeID
	Typ    uint16
	Amount *big.Int
}

// withdrewCandidate
type Ppos_1003 struct {
	NodeId discover.NodeID
}

// delegate
type Ppos_1004 struct {
	Typ    uint16
	NodeId discover.NodeID
	Amount *big.Int
}

// withdrewDelegate
type Ppos_1005 struct {
	StakingBlockNum uint64
	NodeId          discover.NodeID
	Amount          *big.Int
}

// getRelatedListByDelAddr
type Ppos_1103 struct {
	Addr common.Address
}

// getDelegateInfo
type Ppos_1104 struct {
	StakingBlockNum uint64
	DelAddr         common.Address
	NodeId          discover.NodeID
}

// getCandidateInfo
type Ppos_1105 struct {
	NodeId discover.NodeID
}

// submitText
type Ppos_2000 struct {
	Verifier       discover.NodeID
	GithubID       string
	Topic          string
	Desc           string
	Url            string
	EndVotingBlock uint64
}

// submitVersion
type Ppos_2001 struct {
	Verifier       discover.NodeID
	GithubID       string
	Topic          string
	Desc           string
	Url            string
	NewVersion     uint32
	EndVotingBlock uint64
	ActiveBlock    uint64
}

// submitParam
type Ppos_2002 struct {
	Verifier       discover.NodeID
	GithubID       string
	Topic          string
	Desc           string
	Url            string
	ParamName      string
	CurrentValue   interface{}
	NewValue       interface{}
	EndVotingBlock uint64
}

// vote
type Ppos_2003 struct {
	Verifier   discover.NodeID
	ProposalID common.Hash
	Op         uint8
}

//declareVersion
type Ppos_2004 struct {
	ActiveNode discover.NodeID
	Version    uint32
}

// getProposal
type Ppos_2100 struct {
	Verifier   discover.NodeID
	ProposalID common.Hash
}

// getTallyResult
type Ppos_2101 struct {
	ProposalID common.Hash
}

// ReportDuplicateSign
type Ppos_3000 struct {
	Data string
}

// CheckDuplicateSign
type Ppos_3001 struct {
	Etype       uint32
	Addr        common.Address
	BlockNumber uint64
}

// CreateRestrictingPlan
type Ppos_4000 struct {
	Account common.Address
	Plans   []restricting.RestrictingPlan
}

// GetRestrictingInfo
type Ppos_4001 struct {
	Account common.Address
}

type decDataConfig struct {
	P1000 Ppos_1000
	P1001 Ppos_1001
	P1002 Ppos_1002
	P1003 Ppos_1003
	P1004 Ppos_1004
	P1005 Ppos_1005
	P1103 Ppos_1103
	P1104 Ppos_1104
	P1105 Ppos_1105
	P2000 Ppos_2000
	P2001 Ppos_2001
	P2002 Ppos_2002
	P2003 Ppos_2003
	P2004 Ppos_2004
	P2100 Ppos_2100
	P2101 Ppos_2101
	P3000 Ppos_3000
	P3001 Ppos_3001
	P4000 Ppos_4000
	P4001 Ppos_4001
}

func parseConfigJson(configPath string, v *decDataConfig) error {
	if configPath == "" {
		panic(fmt.Errorf("parse config file error"))
	}

	file, err := os.Open(configPath)
	if err != nil {
		utils.Fatalf("Failed to read config file: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(v)
	if err != nil {
		panic(fmt.Errorf("parse config to json error,%s", err.Error()))
	}

	return nil
}

func getRlpData(funcType uint16, cfg *decDataConfig) string {
	rlpData := ""

	var params [][]byte
	params = make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(funcType)
	params = append(params, fnType)

	switch funcType {
	case 1000:
		{
			typ, _ := rlp.EncodeToBytes(cfg.P1000.Typ)
			benefitAddress, _ := rlp.EncodeToBytes(cfg.P1000.BenefitAddress.Bytes())
			nodeId, _ := rlp.EncodeToBytes(cfg.P1000.NodeId)
			externalId, _ := rlp.EncodeToBytes(cfg.P1000.ExternalId)
			nodeName, _ := rlp.EncodeToBytes(cfg.P1000.NodeName)
			website, _ := rlp.EncodeToBytes(cfg.P1000.Website)
			details, _ := rlp.EncodeToBytes(cfg.P1000.Details)
			amount, _ := rlp.EncodeToBytes(cfg.P1000.Amount)
			processVersion, _ := rlp.EncodeToBytes(cfg.P1000.ProcessVersion)
			params = append(params, typ)
			params = append(params, benefitAddress)
			params = append(params, nodeId)
			params = append(params, externalId)
			params = append(params, nodeName)
			params = append(params, website)
			params = append(params, details)
			params = append(params, amount)
			params = append(params, processVersion)
		}
	case 1001:
		{
			benefitAddress, _ := rlp.EncodeToBytes(cfg.P1001.BenefitAddress.Bytes())
			nodeId, _ := rlp.EncodeToBytes(cfg.P1001.NodeId)
			externalId, _ := rlp.EncodeToBytes(cfg.P1001.ExternalId)
			nodeName, _ := rlp.EncodeToBytes(cfg.P1001.NodeName)
			website, _ := rlp.EncodeToBytes(cfg.P1001.Website)
			details, _ := rlp.EncodeToBytes(cfg.P1001.Details)
			params = append(params, benefitAddress)
			params = append(params, nodeId)
			params = append(params, externalId)
			params = append(params, nodeName)
			params = append(params, website)
			params = append(params, details)
		}
	case 1002:
		{
			nodeId, _ := rlp.EncodeToBytes(cfg.P1002.NodeId)
			typ, _ := rlp.EncodeToBytes(cfg.P1002.Typ)
			amount, _ := rlp.EncodeToBytes(cfg.P1002.Amount)
			params = append(params, nodeId)
			params = append(params, typ)
			params = append(params, amount)
		}
	case 1003:
		{
			nodeId, _ := rlp.EncodeToBytes(cfg.P1003.NodeId)
			params = append(params, nodeId)
		}
	case 1004:
		{
			typ, _ := rlp.EncodeToBytes(cfg.P1004.Typ)
			nodeId, _ := rlp.EncodeToBytes(cfg.P1004.NodeId)
			amount, _ := rlp.EncodeToBytes(cfg.P1004.Amount)
			params = append(params, typ)
			params = append(params, nodeId)
			params = append(params, amount)
		}
	case 1005:
		{
			stakingBlockNum, _ := rlp.EncodeToBytes(cfg.P1005.StakingBlockNum)
			nodeId, _ := rlp.EncodeToBytes(cfg.P1005.NodeId)
			amount, _ := rlp.EncodeToBytes(cfg.P1005.Amount)
			params = append(params, stakingBlockNum)
			params = append(params, nodeId)
			params = append(params, amount)
		}
	case 1100:
	case 1101:
	case 1102:
	case 1103:
		{
			addr, _ := rlp.EncodeToBytes(cfg.P1103.Addr.Bytes())
			params = append(params, addr)
		}
	case 1104:
		{
			stakingBlockNum, _ := rlp.EncodeToBytes(cfg.P1104.StakingBlockNum)
			delAddr, _ := rlp.EncodeToBytes(cfg.P1104.DelAddr.Bytes())
			nodeId, _ := rlp.EncodeToBytes(cfg.P1104.NodeId)

			params = append(params, stakingBlockNum)
			params = append(params, delAddr)
			params = append(params, nodeId)
		}
	case 1105:
		{
			nodeId, _ := rlp.EncodeToBytes(cfg.P1105.NodeId)
			params = append(params, nodeId)
		}
	case 2000:
		{
			verifier, _ := rlp.EncodeToBytes(cfg.P2000.Verifier)
			githubID, _ := rlp.EncodeToBytes(cfg.P2000.GithubID)
			topic, _ := rlp.EncodeToBytes(cfg.P2000.Topic)
			desc, _ := rlp.EncodeToBytes(cfg.P2000.Desc)
			url, _ := rlp.EncodeToBytes(cfg.P2000.Url)
			endVotingBlock, _ := rlp.EncodeToBytes(cfg.P2000.EndVotingBlock)
			params = append(params, verifier)
			params = append(params, githubID)
			params = append(params, topic)
			params = append(params, desc)
			params = append(params, url)
			params = append(params, endVotingBlock)
		}
	case 2001:
		{
			verifier, _ := rlp.EncodeToBytes(cfg.P2001.Verifier)
			githubID, _ := rlp.EncodeToBytes(cfg.P2001.GithubID)
			topic, _ := rlp.EncodeToBytes(cfg.P2001.Topic)
			desc, _ := rlp.EncodeToBytes(cfg.P2001.Desc)
			url, _ := rlp.EncodeToBytes(cfg.P2001.Url)
			newVersion, _ := rlp.EncodeToBytes(cfg.P2001.NewVersion)
			endVotingBlock, _ := rlp.EncodeToBytes(cfg.P2001.EndVotingBlock)
			activeBlock, _ := rlp.EncodeToBytes(cfg.P2001.ActiveBlock)
			params = append(params, verifier)
			params = append(params, githubID)
			params = append(params, topic)
			params = append(params, desc)
			params = append(params, url)
			params = append(params, newVersion)
			params = append(params, endVotingBlock)
			params = append(params, activeBlock)
		}
	case 2002:
		{
			verifier, _ := rlp.EncodeToBytes(cfg.P2002.Verifier)
			githubID, _ := rlp.EncodeToBytes(cfg.P2002.GithubID)
			topic, _ := rlp.EncodeToBytes(cfg.P2002.Topic)
			desc, _ := rlp.EncodeToBytes(cfg.P2002.Desc)
			url, _ := rlp.EncodeToBytes(cfg.P2002.Url)
			paramName, _ := rlp.EncodeToBytes(cfg.P2002.ParamName)
			currentValue, _ := rlp.EncodeToBytes(cfg.P2002.CurrentValue)
			newValue, _ := rlp.EncodeToBytes(cfg.P2002.NewValue)
			/*
				t := reflect.TypeOf(cfg.P2002.CurrentValue)
				fmt.Println(t.Kind())
				fmt.Println(t.Name())*/

			endVotingBlock, _ := rlp.EncodeToBytes(cfg.P2002.EndVotingBlock)
			params = append(params, verifier)
			params = append(params, githubID)
			params = append(params, topic)
			params = append(params, desc)
			params = append(params, url)
			params = append(params, paramName)
			params = append(params, currentValue)
			params = append(params, newValue)
			params = append(params, endVotingBlock)
		}
	case 2003:
		{
			verifier, _ := rlp.EncodeToBytes(cfg.P2003.Verifier)
			proposalID, _ := rlp.EncodeToBytes(cfg.P2003.ProposalID.Bytes())
			op, _ := rlp.EncodeToBytes(cfg.P2003.Op)
			params = append(params, verifier)
			params = append(params, proposalID)
			params = append(params, op)
		}
	case 2004:
		{
			activeNode, _ := rlp.EncodeToBytes(cfg.P2004.ActiveNode)
			version, _ := rlp.EncodeToBytes(cfg.P2004.Version)
			params = append(params, activeNode)
			params = append(params, version)
		}
	case 2100:
		{
			proposalID, _ := rlp.EncodeToBytes(cfg.P2100.ProposalID.Bytes())
			params = append(params, proposalID)
		}
	case 2101:
		{
			proposalID, _ := rlp.EncodeToBytes(cfg.P2101.ProposalID.Bytes())
			params = append(params, proposalID)
		}
	case 2102:
	case 2103:
	case 2104:
	case 2105:
	case 3000:
		{
			data, _ := rlp.EncodeToBytes(cfg.P3000.Data)
			params = append(params, data)
		}
	case 3001:
		{
			etype, _ := rlp.EncodeToBytes(cfg.P3001.Etype)
			addr, _ := rlp.EncodeToBytes(cfg.P3001.Addr.Bytes())
			blockNumber, _ := rlp.EncodeToBytes(cfg.P3001.BlockNumber)
			params = append(params, etype)
			params = append(params, addr)
			params = append(params, blockNumber)
		}
	case 4000:
		{
			account, _ := rlp.EncodeToBytes(cfg.P4000.Account.Bytes())
			plans, _ := rlp.EncodeToBytes(cfg.P4000.Plans)
			params = append(params, account)
			params = append(params, plans)
		}
	case 4001:
		{
			account, _ := rlp.EncodeToBytes(cfg.P4001.Account.Bytes())
			params = append(params, account)
		}
	default:
		{
			//	panic(fmt.Errorf("funcType:%d is unknown!!!!", funcType))
			fmt.Printf("funcType:%d is unknown!!!!\n", funcType)
			return ""
		}
	}

	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, params)
	if err != nil {
		panic(fmt.Errorf("%d encode rlp data fail: %v", funcType, err))
	} else {
		rlpData = hexutil.Encode(buf.Bytes())
		fmt.Printf("funcType:%d rlp data = %s\n", funcType, rlpData)
	}

	return rlpData
}

func main() {
	// Parse and ensure all needed inputs are specified
	flag.Parse()

	if *jsonFlag == "" && *funcTypesFlag == "" {
		fmt.Printf("json path is null or funcid is null\n")
		os.Exit(-1)
	}

	cfg := decDataConfig{}
	parseConfigJson(*jsonFlag, &cfg)

	funcTypes := strings.Split(*funcTypesFlag, "|")

	for _, fnType := range funcTypes {
		funcType, e := strconv.Atoi(fnType)
		if e != nil {
			fmt.Printf("funcType is error\n")
			os.Exit(-1)
		}
		getRlpData(uint16(funcType), &cfg)
	}

}
