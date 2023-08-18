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

package stakingL2

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/crypto/bls"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/p2p/discover"
)

const (
	Invalided     CandidateStatus = 1 << iota // 0001: The current candidate withdraws from the staking qualification (Active OR Passive)
	LowRatio                                  // 0010: The candidate was low package ratio AND no delete
	NotEnough                                 // 0100: The current candidate's von does not meet the minimum staking threshold
	DuplicateSign                             // 1000: The Duplicate package or Duplicate sign
	LowRatioDel                               // 0001,0000: The lowRatio AND must delete
	Withdrew                                  // 0010,0000: The Active withdrew
	InBubble
	Valided  = 0       // 0000: The current candidate is in force
	NotExist = 1 << 31 // 1000,xxxx,... : The candidate is not exist
)

const (
	OperatorNode = 1
)

type Operator struct {
	NodeId discover.NodeID
	RPC    string
}

func (op Operator) string() string {
	return fmt.Sprintf("%s@%s:%s", op.NodeId, op.RPC)
}

type CandidateStatus uint32

func (status CandidateStatus) IsValid() bool {
	return !status.IsInvalid()
}

func (status CandidateStatus) IsInvalid() bool {
	return status&Invalided == Invalided
}

func (status CandidateStatus) IsPureInvalid() bool {
	return status&Invalided == status|Invalided
}

func (status CandidateStatus) IsLowRatio() bool {
	return status&LowRatio == LowRatio
}

func (status CandidateStatus) IsPureLowRatio() bool {
	return status&LowRatio == status|LowRatio
}

func (status CandidateStatus) IsNotEnough() bool {
	return status&NotEnough == NotEnough
}

func (status CandidateStatus) IsPureNotEnough() bool {
	return status&NotEnough == status|NotEnough
}

func (status CandidateStatus) IsInvalidLowRatio() bool {
	return status&(Invalided|LowRatio) == (Invalided | LowRatio)
}

func (status CandidateStatus) IsInvalidNotEnough() bool {
	return status&(Invalided|NotEnough) == (Invalided | NotEnough)
}

func (status CandidateStatus) IsInvalidLowRatioNotEnough() bool {
	return status&(Invalided|LowRatio|NotEnough) == (Invalided | LowRatio | NotEnough)
}

func (status CandidateStatus) IsLowRatioNotEnough() bool {
	return status&(LowRatio|NotEnough) == (LowRatio | NotEnough)
}

func (status CandidateStatus) IsDuplicateSign() bool {
	return status&DuplicateSign == DuplicateSign
}

func (status CandidateStatus) IsInvalidDuplicateSign() bool {
	return status&(DuplicateSign|Invalided) == (DuplicateSign | Invalided)
}

func (status CandidateStatus) IsLowRatioDel() bool {
	return status&LowRatioDel == LowRatioDel
}

func (status CandidateStatus) IsPureLowRatioDel() bool {
	return status&LowRatioDel == status|LowRatioDel
}

func (status CandidateStatus) IsInvalidLowRatioDel() bool {
	return status&(Invalided|LowRatioDel) == (Invalided | LowRatioDel)
}

func (status CandidateStatus) IsWithdrew() bool {
	return status&Withdrew == Withdrew
}

func (status CandidateStatus) IsPureWithdrew() bool {
	return status&Withdrew == status|Withdrew
}

func (status CandidateStatus) IsInvalidWithdrew() bool {
	return status&(Invalided|Withdrew) == (Invalided | Withdrew)
}

// The Candidate info
type Candidate struct {
	*CandidateBase
	*CandidateMutable
}

func (can *Candidate) String() string {
	return fmt.Sprintf(`{"NodeId": "%s","BlsPubKey": "%s","StakingAddress": "%s","BenefitAddress": "%s",
						"StakingTxIndex": %d,"ProgramVersion": %d,"Status": %d,"StakingEpoch": %d,"StakingBlockNum": %d,
						"Shares": %d,"Released": %d,"ReleasedHes": %d,"ExternalId": "%s","NodeName": "%s","Website": "%s","Details": "%s"}`,
		fmt.Sprintf("%x", can.NodeId.Bytes()),
		fmt.Sprintf("%x", can.BlsPubKey.Bytes()),
		fmt.Sprintf("%x", can.StakingAddress.Bytes()),
		fmt.Sprintf("%x", can.BenefitAddress.Bytes()),
		can.StakingTxIndex,
		can.ProgramVersion,
		can.Status,
		can.StakingEpoch,
		can.StakingBlockNum,
		can.Shares,
		can.Released,
		can.ReleasedHes,
		//can.RestrictingPlan,
		//can.RestrictingPlanHes,
		can.ExternalId,
		can.NodeName,
		can.Website,
		can.Details)
}

func (can *Candidate) IsNotEmpty() bool {
	return !can.IsEmpty()
}

func (can *Candidate) IsEmpty() bool {
	return nil == can
}

type CandidateBase struct {
	NodeId discover.NodeID
	// bls public key
	BlsPubKey bls.PublicKeyHex
	// The account used to initiate the staking
	StakingAddress common.Address
	// The account receive the block rewards and the staking rewards
	BenefitAddress common.Address
	// The tx index at the time of staking
	StakingTxIndex uint32
	// The version of the node program
	// (Store Large Verson: the 2.1.x large version is 2.1.0)
	ProgramVersion uint32
	// Block height at the time of staking
	StakingBlockNum uint64
	// is Operation node
	IsOperator bool
	// Node desc
	Description
}

func (can *CandidateBase) String() string {
	return fmt.Sprintf(`{"NodeId": "%s","BlsPubKey": "%s","StakingAddress": "%s","BenefitAddress": "%s","StakingTxIndex": %d,"ProgramVersion": %d,"StakingBlockNum": %d,"ExternalId": "%s","NodeName": "%s","Website": "%s","Details": "%s"}`,
		fmt.Sprintf("%x", can.NodeId.Bytes()),
		fmt.Sprintf("%x", can.BlsPubKey.Bytes()),
		fmt.Sprintf("%x", can.StakingAddress.Bytes()),
		fmt.Sprintf("%x", can.BenefitAddress.Bytes()),
		can.StakingTxIndex,
		can.ProgramVersion,
		can.StakingBlockNum,
		can.ExternalId,
		can.NodeName,
		can.Website,
		can.Details)
}

func (can *CandidateBase) IsNotEmpty() bool {
	return !can.IsEmpty()
}

func (can *CandidateBase) IsEmpty() bool {
	return nil == can
}

type CandidateMutable struct {
	// The candidate status
	// Reference `THE CANDIDATE  STATUS`
	Status CandidateStatus
	// The epoch number at staking or edit
	StakingEpoch uint32
	// All vons of staking
	Shares *big.Int
	// The staking von  is circulating for effective epoch (in effect)
	Released *big.Int
	// The staking von  is circulating for hesitant epoch (in hesitation)
	ReleasedHes *big.Int
	// The staking von  is RestrictingPlan for effective epoch (in effect)
	//RestrictingPlan *big.Int
	// The staking von  is RestrictingPlan for hesitant epoch (in hesitation)
	//RestrictingPlanHes *big.Int
	// Internet accessible RPC link
	RPC string
}

func (can *CandidateMutable) String() string {
	return fmt.Sprintf(`{"Status": %d,"StakingEpoch": %d,"Shares": %d,"Released": %d,"ReleasedHes": %d}`,
		can.Status,
		can.StakingEpoch,
		can.Shares,
		can.Released,
		can.ReleasedHes)
	//can.RestrictingPlan,
	//can.RestrictingPlanHes
}

func (can *CandidateMutable) SetValided() {
	can.Status = Valided
}

func (can *CandidateMutable) SetStatus(status CandidateStatus) {
	can.Status = status
}

func (can *CandidateMutable) AppendStatus(status CandidateStatus) {
	can.Status |= status
}

func (can *CandidateMutable) CleanLowRatioStatus() {
	can.Status &^= LowRatio
}

func (can *CandidateMutable) CleanShares() {
	can.Shares = new(big.Int).SetInt64(0)
}

func (can *CandidateMutable) AddShares(amount *big.Int) {
	can.Shares = new(big.Int).Add(can.Shares, amount)
}

func (can *CandidateMutable) SubShares(amount *big.Int) {
	can.Shares = new(big.Int).Sub(can.Shares, amount)
}

func (can *CandidateMutable) IsNotEmpty() bool {
	return !can.IsEmpty()
}

func (can *CandidateMutable) IsEmpty() bool {
	return nil == can
}

func (can *CandidateMutable) IsValid() bool {
	return can.Status.IsValid()
}

func (can *CandidateMutable) IsInvalid() bool {
	return can.Status.IsInvalid()
}

func (can *CandidateMutable) IsPureInvalid() bool {
	return can.Status.IsPureInvalid()
}

func (can *CandidateMutable) IsLowRatio() bool {
	return can.Status.IsLowRatio()
}

func (can *CandidateMutable) IsPureLowRatio() bool {
	return can.Status.IsPureLowRatio()
}

func (can *CandidateMutable) IsNotEnough() bool {
	return can.Status.IsNotEnough()
}

func (can *CandidateMutable) IsPureNotEnough() bool {
	return can.Status.IsPureNotEnough()
}

func (can *CandidateMutable) IsInvalidLowRatio() bool {
	return can.Status.IsInvalidLowRatio()
}

func (can *CandidateMutable) IsInvalidNotEnough() bool {
	return can.Status.IsInvalidNotEnough()
}

func (can *CandidateMutable) IsInvalidLowRatioNotEnough() bool {
	return can.Status.IsInvalidLowRatioNotEnough()
}

func (can *CandidateMutable) IsLowRatioNotEnough() bool {
	return can.Status.IsLowRatioNotEnough()
}

func (can *CandidateMutable) IsDuplicateSign() bool {
	return can.Status.IsDuplicateSign()
}

func (can *CandidateMutable) IsInvalidDuplicateSign() bool {
	return can.Status.IsInvalidDuplicateSign()
}

func (can *CandidateMutable) IsLowRatioDel() bool {
	return can.Status.IsLowRatioDel()
}

func (can *CandidateMutable) IsPureLowRatioDel() bool {
	return can.Status.IsPureLowRatioDel()
}

func (can *CandidateMutable) IsInvalidLowRatioDel() bool {
	return can.Status.IsInvalidLowRatioDel()
}

func (can *CandidateMutable) IsWithdrew() bool {
	return can.Status.IsWithdrew()
}

func (can *CandidateMutable) IsPureWithdrew() bool {
	return can.Status.IsPureWithdrew()
}

func (can *CandidateMutable) IsInvalidWithdrew() bool {
	return can.Status.IsInvalidWithdrew()
}

// Display amount field using 0x hex
type CandidateHex struct {
	NodeId          discover.NodeID
	BlsPubKey       bls.PublicKeyHex
	StakingAddress  common.Address
	BenefitAddress  common.Address
	StakingTxIndex  uint32
	ProgramVersion  uint32
	Status          CandidateStatus
	StakingEpoch    uint32
	StakingBlockNum uint64
	Shares          *hexutil.Big
	Released        *hexutil.Big
	ReleasedHes     *hexutil.Big
	//RestrictingPlan      *hexutil.Big
	//RestrictingPlanHes   *hexutil.Big
	Description
}

func (can *CandidateHex) String() string {
	return fmt.Sprintf(`{"NodeId": "%s","BlsPubKey": "%s","StakingAddress": "%s","BenefitAddress": "%s",
						"StakingTxIndex": %d,"ProgramVersion": %d,"Status": %d,"StakingEpoch": %d,"StakingBlockNum": %d,
						"Shares": "%s","Released": "%s","ReleasedHes": "%s","ExternalId": "%s","NodeName": "%s","Website": "%s","Details": "%s"}`,
		fmt.Sprintf("%x", can.NodeId.Bytes()),
		fmt.Sprintf("%x", can.BlsPubKey.Bytes()),
		fmt.Sprintf("%x", can.StakingAddress.Bytes()),
		fmt.Sprintf("%x", can.BenefitAddress.Bytes()),
		can.StakingTxIndex,
		can.ProgramVersion,
		can.Status,
		can.StakingEpoch,
		can.StakingBlockNum,
		can.Shares,
		can.Released,
		can.ReleasedHes,
		//can.RestrictingPlan,
		//can.RestrictingPlanHes,
		can.ExternalId,
		can.NodeName,
		can.Website,
		can.Details)
}

func (can *CandidateHex) IsNotEmpty() bool {
	return !can.IsEmpty()
}

func (can *CandidateHex) IsEmpty() bool {
	return nil == can
}

//// EncodeRLP implements rlp.Encoder
//func (c *Candidate) EncodeRLP(w io.Writer) error {
//	return rlp.Encode(w, &c)
//}
//
//
//// DecodeRLP implements rlp.Decoder
//func (c *Candidate) DecodeRLP(s *rlp.Stream) error {
//	if err := s.Decode(&c); err != nil {
//		return err
//	}
//	return nil
//}

const (
	MaxExternalIdLen = 70
	MaxNodeNameLen   = 30
	MaxWebsiteLen    = 140
	MaxDetailsLen    = 280
)

type Description struct {
	// External Id for the third party to pull the node description (with length limit)
	ExternalId string
	// The Candidate Node's Name  (with a length limit)
	NodeName string
	// The third-party home page of the node (with a length limit)
	Website string
	RPC     string
	// Description of the node (with a length limit)
	Details string
}

func (desc *Description) CheckLength() error {

	if len(desc.ExternalId) > MaxExternalIdLen {
		return fmt.Errorf("ExternalId overflow, got len is: %d, max len is: %d", len(desc.ExternalId), MaxExternalIdLen)
	}
	if len(desc.NodeName) > MaxNodeNameLen {
		return fmt.Errorf("NodeName overflow, got len is: %d, max len is: %d", len(desc.NodeName), MaxNodeNameLen)
	}
	if len(desc.Website) > MaxWebsiteLen {
		return fmt.Errorf("Website overflow, got len is: %d, max len is: %d", len(desc.Website), MaxWebsiteLen)
	}
	if len(desc.Details) > MaxDetailsLen {
		return fmt.Errorf("Details overflow, got len is: %d, max len is: %d", len(desc.Details), MaxDetailsLen)
	}
	return nil
}

type CandidateQueue []*Candidate

func (queue CandidateQueue) String() string {
	arr := make([]string, len(queue))
	for i, c := range queue {
		arr[i] = c.String()
	}
	return "[" + strings.Join(arr, ",") + "]"
}

type CandidateHexQueue []*CandidateHex

func (queue CandidateHexQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}

func (queue CandidateHexQueue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue CandidateHexQueue) String() string {
	arr := make([]string, len(queue))
	for i, c := range queue {
		arr[i] = c.String()
	}
	return "[" + strings.Join(arr, ",") + "]"
}

type CandidateBaseQueue []*CandidateBase

func (queue CandidateBaseQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}

func (queue CandidateBaseQueue) IsEmpty() bool {
	return len(queue) == 0
}

type CandidateMap map[discover.NodeID]*Candidate

type UnStakeItem struct {
	// this is the nodeAddress
	NodeAddress     common.NodeAddress
	StakingBlockNum uint64
	// Return to normal staking state
	Recovery bool
}
