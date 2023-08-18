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

import "github.com/bubblenet/bubble/common"

var (
	ErrWrongBlsPubKey          = common.NewBizError(321000, "Invalid BLS public key length")
	ErrWrongBlsPubKeyProof     = common.NewBizError(321001, "The BLS proof is incorrect")
	ErrDescriptionLen          = common.NewBizError(321002, "The Description length is incorrect")
	ErrWrongProgramVersionSign = common.NewBizError(321003, "The program version signature is invalid")
	ErrProgramVersionTooLow    = common.NewBizError(321004, "The program version is too low")
	ErrNoSameStakingAddr       = common.NewBizError(321005, "The address must be the same as the one initiated staking")
	ErrStakeVonTooLow          = common.NewBizError(321100, "Staking deposit is insufficient")
	ErrCanAlreadyExist         = common.NewBizError(321101, "The candidate already existed")
	ErrCanNoExist              = common.NewBizError(321102, "The candidate does not exist")
	ErrCanStatusInvalid        = common.NewBizError(321103, "This candidate status is expired")
	ErrIncreaseStakeVonTooLow  = common.NewBizError(321104, "Increased stake is insufficient")
	ErrAccountVonNoEnough      = common.NewBizError(321111, "The account balance is insufficient")
	//ErrGetCandidateList        = common.NewBizError(321202, "Retreiving candidate list failed")
	ErrQueryCandidateInfo = common.NewBizError(321201, "Query candidate info failed")
	ErrNodeID2Addr        = common.NewBizError(321202, "Failed to convert Node ID to address")
)
