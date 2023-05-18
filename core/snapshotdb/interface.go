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

package snapshotdb

import "github.com/bubblenet/bubble/common"

// Putter wraps the database write operation supported by both batches and regular databases.
type Putter interface {
	Put(hash common.Hash, key []byte, value []byte) error
}

// Deleter wraps the database delete operation supported by both batches and regular databases.
type Deleter interface {
	Delete(hash common.Hash, key []byte) error
}

type Writer interface {
	Putter
	Deleter
}

// Database wraps all database operations. All methods are safe for concurrent use.
type Database interface {
	Putter
	Deleter
	Get(hash common.Hash, key []byte) ([]byte, error)
	Has(hash common.Hash, key []byte) (bool, error)
}
