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

package fetcher

import (
	"testing"
	"time"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/consensus/cbft/utils"

	"github.com/stretchr/testify/assert"
)

func TestLimitFetcher_AddTask(t *testing.T) {
	fetcher := NewLimitFetcher()

	msg1 := common.BytesToHash(utils.Rand32Bytes(32))
	msg2 := common.BytesToHash(utils.Rand32Bytes(32))
	msg3 := common.BytesToHash(utils.Rand32Bytes(32))
	assert.True(t, fetcher.AddTask(msg1))
	assert.True(t, fetcher.AddTask(msg2))
	assert.False(t, fetcher.AddTask(msg1))
	assert.True(t, fetcher.AddTask(msg3))
	assert.False(t, fetcher.AddTask(msg2))

	assert.Equal(t, 3, len(fetcher.fetching))

	time.Sleep(keepTimeout + 100*time.Millisecond)

	assert.Equal(t, 0, len(fetcher.fetching))

	fetcher.Stop()
}
