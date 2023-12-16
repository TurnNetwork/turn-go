package db

import (
	"github.com/bubblenet/bubble/core/rawdb"
	"github.com/bubblenet/bubble/ethdb/memorydb"
)

func NewMemoryValidatorDB() *DB {
	return NewDataValidatorDB(rawdb.NewDatabase(memorydb.New()))
}
