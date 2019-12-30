// Copyright 2018-2019 The PlatON Network Authors
// This file is part of the PlatON-Go library.
//
// The PlatON-Go library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The PlatON-Go library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the PlatON-Go library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"testing"

	"github.com/PlatONnetwork/PlatON-Go/common/byteutil"
)

func TestMemorySha3(t *testing.T) {
	stack := newstack()
	stack.push(byteutil.BytesToBigInt([]byte{0x01}))
	stack.push(byteutil.BytesToBigInt([]byte{0x02}))
	r := memorySha3(stack)
	if r.Uint64() != 3 {
		t.Errorf("Expected: 3, got %d", r.Uint64())
	}
}

func TestMemoryCallDataCopy(t *testing.T) {
	stack := newstack()
	stack.push(byteutil.BytesToBigInt([]byte{0x01}))
	stack.push(byteutil.BytesToBigInt([]byte{0x02}))
	stack.push(byteutil.BytesToBigInt([]byte{0x03}))
	r := memoryCallDataCopy(stack)
	if r.Uint64() != 4 {
		t.Errorf("Expected: 4, got %d", r.Uint64())
	}
}

func TestMemoryReturnDataCopy(t *testing.T) {
	stack := newstack()
	stack.push(byteutil.BytesToBigInt([]byte{0x01}))
	stack.push(byteutil.BytesToBigInt([]byte{0x02}))
	stack.push(byteutil.BytesToBigInt([]byte{0x03}))
	r := memoryReturnDataCopy(stack)
	if r.Uint64() != 4 {
		t.Errorf("Expected: 4, got %d", r.Uint64())
	}
}
