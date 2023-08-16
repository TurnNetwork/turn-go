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
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/bubblenet/bubble/x/token"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/bubblenet/bubble/crypto/bls"

	"github.com/bubblenet/bubble/cmd/utils"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/restricting"
)

var (
	jsonFlag      = flag.String("json", "", "Json file path for contract parameters")
	funcTypesFlag = flag.String("funcTypes", "", "A collection of contract interface types")
)

// createStaking
type Dpos_1000 struct {
	Typ                uint16
	BenefitAddress     common.Address
	NodeId             discover.NodeID
	ExternalId         string
	NodeName           string
	Website            string
	Details            string
	Amount             *big.Int
	RewardPer          uint16
	ProgramVersion     uint32
	ProgramVersionSign common.VersionSign
	BlsPubKey          bls.PublicKeyHex
	BlsProof           bls.SchnorrProofHex
}

// editorCandidate
type Dpos_1001 struct {
	BenefitAddress common.Address
	NodeId         discover.NodeID
	RewardPer      uint16
	ExternalId     string
	NodeName       string
	Website        string
	Details        string
}

// increaseStaking
type Dpos_1002 struct {
	NodeId discover.NodeID
	Typ    uint16
	Amount *big.Int
}

// withdrewStaking
type Dpos_1003 struct {
	NodeId discover.NodeID
}

// delegate
type Dpos_1004 struct {
	Typ    uint16
	NodeId discover.NodeID
	Amount *big.Int
}

// withdrewDelegate
type Dpos_1005 struct {
	StakingBlockNum uint64
	NodeId          discover.NodeID
	Amount          *big.Int
}

// getRelatedListByDelAddr
type Dpos_1103 struct {
	Addr common.Address
}

// getDelegateInfo
type Dpos_1104 struct {
	StakingBlockNum uint64
	DelAddr         common.Address
	NodeId          discover.NodeID
}

// getCandidateInfo
type Dpos_1105 struct {
	NodeId discover.NodeID
}

// submitText
type Dpos_2000 struct {
	Verifier discover.NodeID
	PIPID    string
}

// submitVersion
type Dpos_2001 struct {
	Verifier        discover.NodeID
	PIPID           string
	NewVersion      uint32
	EndVotingRounds uint64
}

// submitParam
type Dpos_2002 struct {
	Verifier discover.NodeID
	PIPID    string
	Module   string
	Name     string
	NewValue string
}

// submitCancel
type Dpos_2005 struct {
	Verifier        discover.NodeID
	PIPID           string
	EndVotingRounds uint64
	TobeCanceled    common.Hash
}

// vote
type Dpos_2003 struct {
	Verifier       discover.NodeID
	ProposalID     common.Hash
	Option         uint8
	ProgramVersion uint32
	VersionSign    common.VersionSign
}

// declareVersion
type Dpos_2004 struct {
	Verifier       discover.NodeID
	ProgramVersion uint32
	VersionSign    common.VersionSign
}

// getProposal
type Dpos_2100 struct {
	ProposalID common.Hash
}

// getTallyResult
type Dpos_2101 struct {
	ProposalID common.Hash
}

// listProposal
type Dpos_2102 struct {
}

// getActiveVersion
type Dpos_2103 struct {
}

// getGovernParamValue
type Dpos_2104 struct {
	Module string
	Name   string
}

// getAccuVerifiersCount
type Dpos_2105 struct {
	ProposalID common.Hash
	BlockHash  common.Hash
}

// listGovernParam
type Dpos_2106 struct {
	Module string
}

// reportDuplicateSign
type Dpos_3000 struct {
	DupType uint8
	Data    string
}

// checkDuplicateSign
type Dpos_3001 struct {
	Etype       uint32
	Addr        common.Address
	BlockNumber uint64
}

// CreateRestrictingPlan
type Dpos_4000 struct {
	Account common.Address
	Plans   []restricting.RestrictingPlan
}

// GetRestrictingInfo
type Dpos_4100 struct {
	Account common.Address
}

// withdrawDelegateReward
type Dpos_5000 struct {
}

type Dpos_5100 struct {
	Addr    common.Address
	NodeIDs []discover.NodeID
}

// Dpos_6000 mintToken
type Dpos_6000 struct {
	//Account      common.Address       // 账户地址
	//NativeAmount *big.Int             // 原生代币余额
	//TokenAssets  []token.AccountAsset // Token资产
	AccAsset token.AccountAsset
}

type decDataConfig struct {
	P1000 Dpos_1000
	P1001 Dpos_1001
	P1002 Dpos_1002
	P1003 Dpos_1003
	P1004 Dpos_1004
	P1005 Dpos_1005
	P1103 Dpos_1103
	P1104 Dpos_1104
	P1105 Dpos_1105
	P2000 Dpos_2000
	P2001 Dpos_2001
	P2002 Dpos_2002
	P2005 Dpos_2005
	P2003 Dpos_2003
	P2004 Dpos_2004
	P2100 Dpos_2100
	P2101 Dpos_2101
	P2102 Dpos_2102
	P2103 Dpos_2103
	P2104 Dpos_2104
	P2105 Dpos_2105
	P2106 Dpos_2106
	P3000 Dpos_3000
	P3001 Dpos_3001
	P4000 Dpos_4000
	P4100 Dpos_4100
	P5100 Dpos_5100
	P6000 Dpos_6000
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
			rewardPer, _ := rlp.EncodeToBytes(cfg.P1000.RewardPer)
			programVersion, _ := rlp.EncodeToBytes(cfg.P1000.ProgramVersion)
			programVersionSign, _ := rlp.EncodeToBytes(cfg.P1000.ProgramVersionSign)
			blsPubKey, _ := rlp.EncodeToBytes(cfg.P1000.BlsPubKey)
			blsProof, _ := rlp.EncodeToBytes(cfg.P1000.BlsProof)

			params = append(params, typ)
			params = append(params, benefitAddress)
			params = append(params, nodeId)
			params = append(params, externalId)
			params = append(params, nodeName)
			params = append(params, website)
			params = append(params, details)
			params = append(params, amount)
			params = append(params, rewardPer)
			params = append(params, programVersion)
			params = append(params, programVersionSign)
			params = append(params, blsPubKey)
			params = append(params, blsProof)
		}
	case 1001:
		{
			benefitAddress, _ := rlp.EncodeToBytes(cfg.P1001.BenefitAddress.Bytes())
			nodeId, _ := rlp.EncodeToBytes(cfg.P1001.NodeId)
			rewardPer, _ := rlp.EncodeToBytes(cfg.P1001.RewardPer)
			externalId, _ := rlp.EncodeToBytes(cfg.P1001.ExternalId)
			nodeName, _ := rlp.EncodeToBytes(cfg.P1001.NodeName)
			website, _ := rlp.EncodeToBytes(cfg.P1001.Website)
			details, _ := rlp.EncodeToBytes(cfg.P1001.Details)

			params = append(params, benefitAddress)
			params = append(params, nodeId)
			params = append(params, rewardPer)
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
	case 1006:
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
	case 1106:

	case 1200:
	case 1201:
	case 1202:
	case 2000:
		{
			verifier, _ := rlp.EncodeToBytes(cfg.P2000.Verifier)
			pipID, _ := rlp.EncodeToBytes(cfg.P2000.PIPID)
			params = append(params, verifier)
			params = append(params, pipID)
		}
	case 2001:
		{
			verifier, _ := rlp.EncodeToBytes(cfg.P2001.Verifier)
			pipID, _ := rlp.EncodeToBytes(cfg.P2001.PIPID)
			newVersion, _ := rlp.EncodeToBytes(cfg.P2001.NewVersion)
			endVotingRounds, _ := rlp.EncodeToBytes(cfg.P2001.EndVotingRounds)
			params = append(params, verifier)
			params = append(params, pipID)
			params = append(params, newVersion)
			params = append(params, endVotingRounds)
		}
	case 2002:
		{
			verifier, _ := rlp.EncodeToBytes(cfg.P2002.Verifier)
			pipID, _ := rlp.EncodeToBytes(cfg.P2002.PIPID)
			module, _ := rlp.EncodeToBytes(cfg.P2002.Module)
			name, _ := rlp.EncodeToBytes(cfg.P2002.Name)
			newValue, _ := rlp.EncodeToBytes(cfg.P2002.NewValue)

			params = append(params, verifier)
			params = append(params, pipID)
			params = append(params, module)
			params = append(params, name)
			params = append(params, newValue)
		}
	case 2005:
		{
			verifier, _ := rlp.EncodeToBytes(cfg.P2005.Verifier)
			pipID, _ := rlp.EncodeToBytes(cfg.P2005.PIPID)
			endVotingRounds, _ := rlp.EncodeToBytes(cfg.P2005.EndVotingRounds)
			tobeCanceled, _ := rlp.EncodeToBytes(cfg.P2005.TobeCanceled)
			params = append(params, verifier)
			params = append(params, pipID)
			params = append(params, endVotingRounds)
			params = append(params, tobeCanceled)
		}
	case 2003:
		{
			verifier, _ := rlp.EncodeToBytes(cfg.P2003.Verifier)
			proposalID, _ := rlp.EncodeToBytes(cfg.P2003.ProposalID.Bytes())
			op, _ := rlp.EncodeToBytes(cfg.P2003.Option)
			programVersion, _ := rlp.EncodeToBytes(cfg.P2003.ProgramVersion)
			versionSign, _ := rlp.EncodeToBytes(cfg.P2003.VersionSign)
			params = append(params, verifier)
			params = append(params, proposalID)
			params = append(params, op)
			params = append(params, programVersion)
			params = append(params, versionSign)
		}
	case 2004:
		{
			verifier, _ := rlp.EncodeToBytes(cfg.P2004.Verifier)
			programVersion, _ := rlp.EncodeToBytes(cfg.P2004.ProgramVersion)
			versionSign, _ := rlp.EncodeToBytes(cfg.P2004.VersionSign)
			params = append(params, verifier)
			params = append(params, programVersion)
			params = append(params, versionSign)
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
		{
			module, _ := rlp.EncodeToBytes(cfg.P2104.Module)
			name, _ := rlp.EncodeToBytes(cfg.P2104.Name)
			params = append(params, module)
			params = append(params, name)
		}
	case 2105:
		{
			proposalID, _ := rlp.EncodeToBytes(cfg.P2105.ProposalID.Bytes())
			blockHash, _ := rlp.EncodeToBytes(cfg.P2105.BlockHash.Bytes())
			params = append(params, proposalID)
			params = append(params, blockHash)
		}
	case 2106:
		{
			module, _ := rlp.EncodeToBytes(cfg.P2106.Module)
			params = append(params, module)
		}
	case 3000:
		{
			dupType, _ := rlp.EncodeToBytes(cfg.P3000.DupType)
			data, _ := rlp.EncodeToBytes(cfg.P3000.Data)
			params = append(params, dupType)
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
	case 4100:
		{
			account, _ := rlp.EncodeToBytes(cfg.P4100.Account.Bytes())
			params = append(params, account)
		}
	case 5000:
	case 5100:
		{
			addr, _ := rlp.EncodeToBytes(cfg.P5100.Addr.Bytes())
			nodeIds, _ := rlp.EncodeToBytes(cfg.P5100.NodeIDs)
			params = append(params, addr)
			params = append(params, nodeIds)
		}
	case 6000:
		{
			//account, _ := rlp.EncodeToBytes(cfg.P6000.Account)
			//nativeAmount, _ := rlp.EncodeToBytes(cfg.P6000.NativeAmount)
			//tokenAssets, _ := rlp.EncodeToBytes(cfg.P6000.TokenAssets)
			//params = append(params, account)
			//params = append(params, nativeAmount)
			//params = append(params, tokenAssets)

			accAsset, _ := rlp.EncodeToBytes(cfg.P6000.AccAsset)
			params = append(params, accAsset)
		}
	case 6001:
		{
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
