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

package core

import (
	"math/big"
	"testing"
	"time"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/cbfttypes"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/event"
	"github.com/bubblenet/bubble/trie"
)

func TestBlockChainReactorClose(t *testing.T) {
	t.Run("close after commit", func(t *testing.T) {
		eventmux := new(event.TypeMux)
		reacter := NewBlockChainReactor(eventmux, big.NewInt(100))
		reacter.Start(common.DPOS_VALIDATOR_MODE)
		var parenthash common.Hash
		cbftress := make(chan cbfttypes.CbftResult, 5)
		go func() {
			for i := 1; i < 11; i++ {
				header := new(types.Header)
				header.Number = big.NewInt(int64(i))
				header.Time = uint64(i)
				header.ParentHash = parenthash
				block := types.NewBlock(header, nil, nil, new(trie.Trie))
				snapshotdb.Instance().NewBlock(header.Number, header.ParentHash, block.Hash())
				parenthash = block.Hash()
				cbftress <- cbfttypes.CbftResult{Block: block}
			}
			close(cbftress)
		}()

		for value := range cbftress {
			eventmux.Post(value)
		}

		reacter.Close()

		time.Sleep(time.Second)
		snapshotdb.Instance().Clear()
	})
}
