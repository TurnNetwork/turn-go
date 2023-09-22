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
	"github.com/bubblenet/bubble/common/hexutil"
	"math/big"
	"strings"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/p2p/discover"
)

const (
	Invalid CandidateStatus = 1 << iota // 0001:  candidate is invalid caused by withdraw
	Frozen                              // 0010: candidate is frozen caused by duplicate signature
	Valid   = 0                         // 0000: candidate is valid
)

// CandidateStatus is the candidate status
type CandidateStatus uint32

func (status CandidateStatus) IsValid() bool {
	return status == Valid
}

func (status CandidateStatus) IsInvalid() bool {
	return status&Invalid == Invalid
}

func (status CandidateStatus) IsInvalidOnly() bool {
	return status&Invalid == status|Invalid
}

func (status CandidateStatus) IsFrozen() bool {
	return status&Frozen == Frozen
}

func (status CandidateStatus) IsFrozenOnly() bool {
	return status&Frozen == status|Frozen
}

// The Candidate info
type Candidate struct {
	*CandidateBase
	*CandidateMutable
}

func (can *Candidate) String() string {
	return fmt.Sprintf(`{"NodeId": "%s","Name": "%s","Status": %d,"Version": %d,"ElectronURI": "%s","P2PURI": "%s",
							"IsOperator": "%t","StakingAddress": "%s","BenefitAddress": "%s","StakingEpoch": %d,"StakingBlockNum": %d,
							"StakingTxIndex": %d,"Shares": %d,"PendingShares": %d,"LockedShares": %d,"BlsPubKey": "%s","Detail": "%s"}`,
		fmt.Sprintf("%x", can.NodeId.Bytes()),
		can.Name,
		can.Status,
		can.Version,
		can.ElectronURI,
		can.P2PURI,
		can.IsOperator,
		fmt.Sprintf("%x", can.StakingAddress.Bytes()),
		fmt.Sprintf("%x", can.BenefitAddress.Bytes()),
		can.StakingEpoch,
		can.StakingBlockNum,
		can.StakingTxIndex,
		can.Shares,
		can.PendingShares,
		can.LockedShares,
		fmt.Sprintf("%x", can.BlsPubKey.Bytes()),
		can.Detail,
	)
}

func (can *Candidate) IsEmpty() bool {
	return nil == can
}

type CandidateBase struct {
	NodeId discover.NodeID
	Name   string
	// The micro node version represent by uint32, only store large version (2.1.x == 2.1.0)
	Version uint32
	// The http URL to receive the genesis messages
	ElectronURI string
	// The RPC URI to receive the staking token txs
	RPCURI string
	// The P2P URI used to link peers
	P2PURI string
	// is operation node
	IsOperator bool

	// The staking transaction sender
	StakingAddress common.Address
	// The account to receive the block rewards and the staking rewards
	BenefitAddress common.Address
	// The staking block number
	StakingBlockNum uint64
	// The transaction index on the staking block
	StakingTxIndex uint32

	// BLS public key
	BlsPubKey bls.PublicKeyHex
	Detail    string
}

func (can *CandidateBase) String() string {
	return fmt.Sprintf(`{"NodeId": "%s","Name": "%s","Version": %d,"ElectronURI": "%s","P2PURI": "%s","IsOperator":"%t",
							StakingAddress": "%s","BenefitAddress": "%s","StakingBlockNum": %d,"StakingTxIndex": %d,"BlsPubKey": "%s","Detail": "%s"}`,
		fmt.Sprintf("%x", can.NodeId.Bytes()),
		can.Name,
		can.Version,
		can.ElectronURI,
		can.P2PURI,
		can.IsOperator,
		fmt.Sprintf("%x", can.StakingAddress.Bytes()),
		fmt.Sprintf("%x", can.BenefitAddress.Bytes()),
		can.StakingBlockNum,
		can.StakingTxIndex,
		fmt.Sprintf("%x", can.BlsPubKey.Bytes()),
		can.Detail,
	)
}

func (can *CandidateBase) IsEmpty() bool {
	return nil == can
}

func (can *CandidateBase) CheckDescription() error {
	if len(can.Name) > 30 {
		return fmt.Errorf("node name overlength, got: %d, expect: %d", len(can.Name), 30)
	}
	if len(can.Detail) > 280 {
		return fmt.Errorf("details overlength, got: %d, expect: %d", len(can.Detail), 280)
	}
	if len(can.ElectronURI) > 200 {
		return fmt.Errorf("electron URI overlength, got: %d, expect: %d", len(can.ElectronURI), 200)
	}
	if len(can.P2PURI) > 200 {
		return fmt.Errorf("P2P URI overlength, got: %d, expect: %d", len(can.P2PURI), 200)
	}
	return nil
}

type CandidateMutable struct {
	// The candidate status
	Status CandidateStatus
	// The epoch number of the block that staking or edit staking
	StakingEpoch uint32
	// The total amount of staking von
	Shares *big.Int
	// The total amount of pending staking von
	PendingShares *big.Int
	// The total amount of locked staking von
	LockedShares *big.Int
}

func (can *CandidateMutable) String() string {
	return fmt.Sprintf(`{"Status": %d,"StakingEpoch": %d,"Shares": %d,"PendingShares": %d,"LockedShares": %d,}`,
		can.Status,
		can.StakingEpoch,
		can.Shares,
		can.PendingShares,
		can.LockedShares)
}

func (can *CandidateMutable) SetStatus(status CandidateStatus) {
	can.Status = status
}

func (can *CandidateMutable) AppendStatus(status CandidateStatus) {
	can.Status |= status
}

func (can *CandidateMutable) CleanShares() {
	can.Shares = new(big.Int).SetInt64(0)
}

func (can *CandidateMutable) AddShares(amount *big.Int) {
	can.Shares = new(big.Int).Add((*big.Int)(can.Shares), amount)
}

func (can *CandidateMutable) SubShares(amount *big.Int) {
	can.Shares = new(big.Int).Sub((*big.Int)(can.Shares), amount)
}

func (can *CandidateMutable) IsEmpty() bool {
	return nil == can
}

type MarshalAbleCandidate struct {
	NodeId      discover.NodeID
	Name        string
	Status      CandidateStatus
	Version     uint32
	ElectronURI string
	RPCURI      string
	P2PURI      string
	IsOperator  bool

	StakingAddress  common.Address
	BenefitAddress  common.Address
	StakingEpoch    uint32
	StakingBlockNum uint64
	StakingTxIndex  uint32
	Shares          *hexutil.Big
	LockedShares    *hexutil.Big
	PendingShares   *hexutil.Big

	BlsPubKey bls.PublicKeyHex
	Detail    string
}

func (can *MarshalAbleCandidate) String() string {
	return fmt.Sprintf(`{"NodeId": "%s","Name": "%s","Status": %d,"Version": %d,"ElectronURI": "%s","P2PURI": "%s",
							"IsOperator": "%t","StakingAddress": "%s","BenefitAddress": "%s","StakingEpoch": %d,"StakingBlockNum": %d,
							"StakingTxIndex": %d,"Shares": %d,"PendingShares": %d,"LockedShares": %d,"BlsPubKey": "%s","Detail": "%s"}`,
		fmt.Sprintf("%x", can.NodeId.Bytes()),
		can.Name,
		can.Status,
		can.Version,
		can.ElectronURI,
		can.P2PURI,
		can.IsOperator,
		fmt.Sprintf("%x", can.StakingAddress.Bytes()),
		fmt.Sprintf("%x", can.BenefitAddress.Bytes()),
		can.StakingEpoch,
		can.StakingBlockNum,
		can.StakingTxIndex,
		can.Shares,
		can.PendingShares,
		can.LockedShares,
		fmt.Sprintf("%x", can.BlsPubKey.Bytes()),
		can.Detail,
	)
}

func (can *MarshalAbleCandidate) IsEmpty() bool {
	return nil == can
}

type CandidateQueue []*Candidate

func (queue CandidateQueue) String() string {
	arr := make([]string, len(queue))
	for i, c := range queue {
		arr[i] = c.String()
	}
	return "[" + strings.Join(arr, ",") + "]"
}

type UnStakeRecord struct {
	NodeAddress     common.NodeAddress
	StakingBlockNum uint64
}
