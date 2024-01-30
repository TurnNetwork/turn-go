package bubble

import (
	"errors"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type BubbleDB struct {
	db snapshotdb.DB
}

func NewBubbleDB() *BubbleDB {
	return &BubbleDB{
		db: snapshotdb.Instance(),
	}
}

func (bdb *BubbleDB) IteratorBubContract(blockHash common.Hash, ranges int) iterator.Iterator {
	return bdb.db.Ranking(blockHash, BubContractPrefix, ranges)
}

func (bdb *BubbleDB) GetBubContract(blockHash common.Hash, address *common.Address) (*common.Address, error) {
	data, err := bdb.db.Get(blockHash, getBubContractKey(address))
	if err != nil {
		return nil, err
	}

	var contract common.Address
	if data != nil {
		contract.SetBytes(data)
	} else {
		return nil, errors.New("contract address is nil")
	}

	return &contract, nil
}

func (bdb *BubbleDB) StoreBubContract(blockHash common.Hash, address *common.Address) error {
	return bdb.db.Put(blockHash, getBubContractKey(address), address.Bytes())
}

func (bdb *BubbleDB) DelBubContract(blockHash common.Hash, address *common.Address) error {
	return bdb.db.Del(blockHash, getBubContractKey(address))
}
