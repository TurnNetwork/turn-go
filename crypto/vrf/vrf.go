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

package vrf

import (
	"crypto/ecdsa"
	"errors"
)

var (
	NotSupportKey = errors.New("Unsupported key type")
)

func Prove(key *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	pk := ECP2OS(key.X, key.Y)
	sk := make([]byte, 32)
	blob := key.D.Bytes()
	copy(sk[32-len(blob):], blob)
	if pi, err := eCVRF_prove(pk, sk, data[:]); err != nil {
		return nil, err
	} else {
		return pi, nil
	}
}

func Verify(key *ecdsa.PublicKey, pi []byte, data []byte) (bool, error) {
	if res, err := eCVRF_verify(ECP2OS(key.X, key.Y), pi, data[:]); err != nil {
		return false, err
	} else {
		return res, nil
	}
}

func ProofToHash(pi []byte) []byte {
	return eCVRF_proof2hash(pi)
}
