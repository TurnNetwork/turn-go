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
	"github.com/holiman/uint256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallGas(t *testing.T) {

	_, err := callGas(100, 2, uint256.NewInt().SetUint64(10))
	assert.Nil(t, err)

	_, err = callGas(100, 2, uint256.NewInt().SetUint64(1000000000000000))
	assert.Nil(t, err)
}
