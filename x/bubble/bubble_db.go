package bubble

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/rlp"
)

type BubbleDB struct {
	db snapshotdb.DB
}

func NewBubbleDB() *BubbleDB {
	return &BubbleDB{
		db: snapshotdb.Instance(),
	}
}

func (bdb *BubbleDB) GetBubbleStore(blockHash common.Hash, bubbleID uint32) (*Bubble, error) {
	data, err := bdb.db.Get(blockHash, GetBubbleKey(bubbleID))
	if err != nil {
		return nil, err
	}

	var bubble Bubble
	if err := rlp.DecodeBytes(data, bubble); err != nil {
		return nil, err
	} else {
		return &bubble, nil
	}
}

func (bdb *BubbleDB) SetBubbleStore(blockHash common.Hash, bubble *Bubble) error {
	if data, err := rlp.EncodeToBytes(bubble); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, GetBubbleKey(bubble.BubbleId), data)
	}
}

func (bdb *BubbleDB) DelBubbleStore(blockHash common.Hash, bubbleID uint32) error {
	if err := bdb.db.Del(blockHash, GetBubbleKey(bubbleID)); err != nil {
		return err
	}
	return nil
}

func (bdb *BubbleDB) GetAccListOfBub(blockHash common.Hash, bubbleId uint32) ([]common.Address, error) {
	data, err := bdb.db.Get(blockHash, AccListByBubKey(bubbleId))
	if err != nil {
		return nil, err
	}

	var accList []common.Address
	if err := rlp.DecodeBytes(data, &accList); err != nil {
		return nil, err
	} else {
		return accList, nil
	}
}

// StoreAccListOfBub Store the staking tokens accounts into the snapshot db
func (bdb *BubbleDB) StoreAccListOfBub(blockHash common.Hash, bubbleId uint32, accounts []common.Address) error {
	if data, err := rlp.EncodeToBytes(accounts); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, AccListByBubKey(bubbleId), data)
	}
}

func (bdb *BubbleDB) GetAccAssetOfBub(blockHash common.Hash, bubbleId uint32, account common.Address) (*AccountAsset, error) {
	data, err := bdb.db.Get(blockHash, AccAssetByBubKey(bubbleId, account))
	if err != nil {
		return nil, err
	}

	var accAsset AccountAsset
	if err := rlp.DecodeBytes(data, &accAsset); err != nil {
		return nil, err
	} else {
		return &accAsset, nil
	}
}

func (bdb *BubbleDB) StoreAccAssetToBub(blockHash common.Hash, bubbleId uint32, accAsset AccountAsset) error {
	if data, err := rlp.EncodeToBytes(accAsset); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, AccAssetByBubKey(bubbleId, accAsset.Account), data)
	}
}
