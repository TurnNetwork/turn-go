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

package token

import "github.com/bubblenet/bubble/common"

var (
	ErrStoreL1HashToL2Hash   = common.NewBizError(600000, "Failed to store the main-chain transaction hash to sub-chain hash")
	ErrGetL2TxHashByL1       = common.NewBizError(610000, "Failed to obtain the sub-chain transaction hash")
	ErrAddMintAccInfo        = common.NewBizError(600001, "Failed to add minting account information")
	ErrGetMintAccInfo        = common.NewBizError(610011, "Failed to obtain minting account information")
	ErrEncodeGetBalancesData = common.NewBizError(600002, "Failed to generate data for ERC20 batch get balance interface")
	ErrEVMExecERC20          = common.NewBizError(600003, "evm fails to execute erc20 contract transfer interface")
	ErrGetERC20Token         = common.NewBizError(600004, "failed to get Address ERC20 Token")
	ErrGetSettleInfoHash     = common.NewBizError(600005, "failed to get settlement information hash")
	ErrCalcSettleInfoHash    = common.NewBizError(600006, "failed to calculate settlement information hash")
	ErrNotNeedToSettle       = common.NewBizError(600007, "Settlement is not required if there are no new transactions in the relevant account")
	ErrNotMainOpAddr         = common.NewBizError(600008, "the transaction sender is not the main chain operator address")
	ErrEncodeMintData        = common.NewBizError(600011, "Failed to generate data for ERC20 mint interface")
)
