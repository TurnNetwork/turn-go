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

import (
	"fmt"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/token"
	"github.com/status-im/keycard-go/hexutils"
	"testing"
)

// 解析Settlement交易日志
func TestDecodeSettlementTxLog(t *testing.T) {
	var settlementInfo token.SettlementInfo
	var m [][]byte
	data := "f86030b85df85bf859f857946a311b9d42ea0cb4f62760383c0eff06ac68f1f78ba56fa60cd461dbb4580000f5d9948a16806861ca61ef5d3bb99498b94d3367a248ba83989680da9412c171900f010b17e969702efa044d077e8680828401312d00"
	if err := rlp.DecodeBytes(hexutils.HexToBytes(data), &m); err != nil {
		t.Error(err)
	}
	var code string
	if err := rlp.DecodeBytes(m[0], &code); err != nil {
		t.Error(err)
	}
	if err := rlp.DecodeBytes(m[1], &settlementInfo); err != nil {
		t.Error(err)
	}

	fmt.Printf("settlementInfo: %v\n", settlementInfo)
	return
}
