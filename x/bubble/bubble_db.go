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
