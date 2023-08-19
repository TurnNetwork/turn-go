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

func (bdb *BubbleDB) GetAccListOfStakingTokenInBub(blockHash common.Hash, bubbleId uint32) ([]common.Address, error) {
	data, err := bdb.db.Get(blockHash, StakingTokenAccListKey(bubbleId))
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

// StoreAccListOfStakingToken Store the pledge token information into the snapshot db
func (bdb *BubbleDB) StoreAccListOfStakingToken(blockHash common.Hash, bubbleId uint32, accounts []common.Address) error {
	if data, err := rlp.EncodeToBytes(accounts); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, StakingTokenAccListKey(bubbleId), data)
	}
}

func (bdb *BubbleDB) GetAccAssetOfStakingInBub(blockHash common.Hash, bubbleId uint32, account common.Address) (*AccountAsset, error) {
	data, err := bdb.db.Get(blockHash, StakingTokenAccAssetKey(bubbleId, account))
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

func (bdb *BubbleDB) StoreAccAssetOfStakingInBub(blockHash common.Hash, bubbleId uint32, accAsset AccountAsset) error {
	if data, err := rlp.EncodeToBytes(accAsset); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, StakingTokenAccAssetKey(bubbleId, accAsset.Account), data)
	}
}

// StoreStakingTokenInfo Store the pledge token information into the snapshot db
func (bdb *BubbleDB) StoreStakingTokenInfo(blockHash common.Hash, bubbleId uint32, stakingAsset *AccountAsset) error {
	if data, err := rlp.EncodeToBytes(bubbleId); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, StakingTokenInBubKey(bubbleId), data)
	}
}
