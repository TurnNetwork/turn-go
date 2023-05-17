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

package plugin

import (
	"context"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/json"
	"github.com/bubblenet/bubble/core/snapshotdb"
)

// Provides an API interface to obtain data related to the economic model
type PublicDPOSAPI struct {
	snapshotDB snapshotdb.DB
}

func NewPublicDPOSAPI() *PublicDPOSAPI {
	return &PublicDPOSAPI{snapshotdb.Instance()}
}

// Get node list of zero-out blocks
func (p *PublicDPOSAPI) GetWaitSlashingNodeList() string {
	list, err := slash.getWaitSlashingNodeList(0, common.ZeroHash)
	if nil != err || len(list) == 0 {
		return ""
	}
	enVal, err := json.Marshal(list)
	if err != nil {
		return ""
	}
	return string(enVal)
}

func (p *PublicDPOSAPI) GetValidatorByBlockNumber(ctx context.Context, blockNumber uint64) string {
	list, err := stk.GetValidatorHistoryList(blockNumber)
	if nil != err || len(list) == 0 {
		return ""
	}
	enVal, err := json.Marshal(list)
	return string(enVal)
}
