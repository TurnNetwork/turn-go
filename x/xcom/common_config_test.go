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

package xcom

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/rlp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultEMConfig(t *testing.T) {
	t.Run("DefaultMainNet", func(t *testing.T) {
		if getDefaultEMConfig(DefaultMainNet) == nil {
			t.Error("DefaultMainNet can't be nil config")
		}
	})
	t.Run("DefaultTestNet", func(t *testing.T) {
		if getDefaultEMConfig(DefaultTestNet) == nil {
			t.Error("DefaultTestNet can't be nil config")
		}
	})
	t.Run("DefaultUnitTestNet", func(t *testing.T) {
		if getDefaultEMConfig(DefaultUnitTestNet) == nil {
			t.Error("DefaultUnitTestNet can't be nil config")
		}
	})
	if getDefaultEMConfig(10) != nil {
		t.Error("the chain config not support")
	}
}

func TestMainNetHash(t *testing.T) {
	tempEc := getDefaultEMConfig(DefaultMainNet)
	bytes, err := rlp.EncodeToBytes(tempEc)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, common.RlpHash(bytes).Hex() == MainNetECHash)
}
