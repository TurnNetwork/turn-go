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

package common

import (
	"bytes"
	"golang.org/x/crypto/sha3"

	"github.com/bubblenet/bubble/rlp"
)

func MustRlpEncode(val interface{}) []byte {
	bytes, err := rlp.EncodeToBytes(val)
	if err != nil {
		panic(err)
	}
	return bytes
}

func GenerateKVHash(k, v []byte, oldHash Hash) Hash {
	var buf bytes.Buffer
	buf.Write(k)
	buf.Write(v)
	buf.Write(oldHash.Bytes())
	return RlpHash(buf.Bytes())
}

func RlpHash(x interface{}) (h Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}
