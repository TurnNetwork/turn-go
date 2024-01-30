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

package vm

import "github.com/bubblenet/bubble/common"

// the inner contract addr  table
var (
	TokenContractAddr          = common.HexToAddress("0x1000000000000000000000000000000000000020") // The Bubble Precompiled contract addr for token
	TempPrivateKeyContractAddr = common.HexToAddress("0x1000000000000000000000000000000000000021") // The Bubble Precompiled contract temporary private keycase cvm.TokenContractAddr:
	RestrictingContractAddr    = common.HexToAddress("0x1000000000000000000000000000000000000001") // The Bubble Precompiled contract addr for restricting
	StakingContractAddr        = common.HexToAddress("0x1000000000000000000000000000000000000002") // The Bubble Precompiled contract addr for staking
	RewardManagerPoolAddr      = common.HexToAddress("0x1000000000000000000000000000000000000003") // The Bubble Precompiled contract addr for reward
	SlashingContractAddr       = common.HexToAddress("0x1000000000000000000000000000000000000004") // The Bubble Precompiled contract addr for slashing
	GovContractAddr            = common.HexToAddress("0x1000000000000000000000000000000000000005") // The Bubble Precompiled contract addr for governance
	DelegateRewardPoolAddr     = common.HexToAddress("0x1000000000000000000000000000000000000006") // The Bubble Precompiled contract addr for delegate reward
	ValidatorInnerContractAddr = common.HexToAddress("0x2000000000000000000000000000000000000000") // The Bubble Precompiled contract addr for cbft inner
	BubbleContractAddr         = common.HexToAddress("0x2000000000000000000000000000000000000001") // The Bubble Precompiled contract addr for token
	VrfInnerContractAddr       = common.HexToAddress("0x3000000000000000000000000000000000000001") // The Bubble Precompiled contract addr for vrf inner
)

type PrecompiledContractCheck interface {
	IsBubblePrecompiledContract(address common.Address) bool
}

var PrecompiledContractCheckInstance PrecompiledContractCheck
