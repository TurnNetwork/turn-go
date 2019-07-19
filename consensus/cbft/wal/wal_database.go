package wal

import (
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type IWALDatabase interface {
	Put(key []byte, value []byte, wo *opt.WriteOptions) error
	Delete(key []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Close()
}

type WALDatabase struct {
	fn  string      // filename for reporting
	db  *leveldb.DB // LevelDB instance
	log log.Logger  // Contextual logger tracking the database path
}

func createWalDB(file string) (IWALDatabase, error) {
	if file == "" {
		return nil, errors.New("create waldb error,file is empty")
	}
	db, err := openDatabase(file)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func openDatabase(file string) (IWALDatabase, error) {
	db, err := newWALDatabase(file, 0, 0)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func newWALDatabase(file string, cache int, handles int) (*WALDatabase, error) {
	logger := log.New("Wal_database", file)

	// Ensure we have some minimal caching and file guarantees
	if cache < 16 {
		cache = 16
	}
	if handles < 16 {
		handles = 16
	}
	logger.Info("Allocated cache and file handles", "cache", cache, "handles", handles)

	// Open the db and recover any potential corruptions
	db, err := leveldb.OpenFile(file, &opt.Options{
		OpenFilesCacheCapacity: handles,
		BlockCacheCapacity:     cache / 2 * opt.MiB,
		WriteBuffer:            cache / 4 * opt.MiB, // Two of these are used internally
		Filter:                 filter.NewBloomFilter(10),
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(file, nil)
	}
	// (Re)check for errors and abort if opening of the db failed
	if err != nil {
		return nil, err
	}
	return &WALDatabase{
		fn:  file,
		db:  db,
		log: logger,
	}, nil
}

// Path returns the path to the database directory.
func (db *WALDatabase) Path() string {
	return db.fn
}

// Put puts the given key / value to the queue
func (db *WALDatabase) Put(key []byte, value []byte, wo *opt.WriteOptions) error {
	return db.db.Put(key, value, wo)
}

func (db *WALDatabase) Has(key []byte) (bool, error) {
	return db.db.Has(key, nil)
}

// Get returns the given key if it's present.
func (db *WALDatabase) Get(key []byte) ([]byte, error) {
	dat, err := db.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

// Delete deletes the key from the queue and database
func (db *WALDatabase) Delete(key []byte) error {
	return db.db.Delete(key, nil)
}

func (db *WALDatabase) Close() {
	// Stop the metrics collection to avoid internal database races
	err := db.db.Close()
	if err == nil {
		db.log.Info("Database closed")
	} else {
		db.log.Error("Failed to close database", "err", err)
	}
}
