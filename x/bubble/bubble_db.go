package bubble

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/rlp"
	"math/big"
)

type BubbleDB struct {
	db snapshotdb.DB
}

func NewBubbleDB() *BubbleDB {
	return &BubbleDB{
		db: snapshotdb.Instance(),
	}
}

func (bdb *BubbleDB) GetBubbleStore(blockHash common.Hash, bubbleID *big.Int) (*Bubble, error) {
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

func (bdb *BubbleDB) DelBubbleStore(blockHash common.Hash, bubbleID *big.Int) error {
	if err := bdb.db.Del(blockHash, GetBubbleKey(bubbleID)); err != nil {
		return err
	}
	return nil
}

func (bdb *BubbleDB) GetAccListOfBub(blockHash common.Hash, bubbleId *big.Int) ([]common.Address, error) {
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
func (bdb *BubbleDB) StoreAccListOfBub(blockHash common.Hash, bubbleId *big.Int, accounts []common.Address) error {
	if data, err := rlp.EncodeToBytes(accounts); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, AccListByBubKey(bubbleId), data)
	}
}

func (bdb *BubbleDB) GetAccAssetOfBub(blockHash common.Hash, bubbleId *big.Int, account common.Address) (*AccountAsset, error) {
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

func (bdb *BubbleDB) StoreAccAssetToBub(blockHash common.Hash, bubbleId *big.Int, accAsset AccountAsset) error {
	if data, err := rlp.EncodeToBytes(accAsset); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, AccAssetByBubKey(bubbleId, accAsset.Account), data)
	}
}

func (bdb *BubbleDB) StoreTxHashListToBub(blockHash common.Hash, bubbleID *big.Int, txHashList []common.Hash, txType BubTxType) error {
	if data, err := rlp.EncodeToBytes(txHashList); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, TxHashListByBubKey(bubbleID, txType), data)
	}
}

func (bdb *BubbleDB) GetTxHashListByBub(blockHash common.Hash, bubbleID *big.Int, txType BubTxType) (*[]common.Hash, error) {
	data, err := bdb.db.Get(blockHash, TxHashListByBubKey(bubbleID, txType))
	if err != nil {
		return nil, err
	}

	var txHashList []common.Hash
	if err := rlp.DecodeBytes(data, &txHashList); err != nil {
		return nil, err
	} else {
		return &txHashList, nil
	}
}

func (bdb *BubbleDB) StoreL2HashToL1Hash(blockHash common.Hash, bubbleID *big.Int, L1TxHash common.Hash, L2TxHash common.Hash) error {
	if data, err := rlp.EncodeToBytes(L1TxHash); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, TxHashByBubKey(bubbleID, L2TxHash), data)
	}
}

func (bdb *BubbleDB) GetL1HashByL2Hash(blockHash common.Hash, bubbleID *big.Int, L2TxHash common.Hash) (*common.Hash, error) {
	data, err := bdb.db.Get(blockHash, TxHashByBubKey(bubbleID, L2TxHash))
	if err != nil {
		return nil, err
	}

	var L1TxHash common.Hash
	if err := rlp.DecodeBytes(data, &L1TxHash); err != nil {
		return nil, err
	} else {
		return &L1TxHash, nil
	}
}
