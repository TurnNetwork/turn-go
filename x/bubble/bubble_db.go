package bubble

import (
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/log"
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

func (bdb *BubbleDB) IteratorBubBasics(blockHash common.Hash, ranges int) iterator.Iterator {
	return bdb.db.Ranking(blockHash, BubBasicsKeyPrefix, ranges)
}

func (bdb *BubbleDB) GetBubBasics(blockHash common.Hash, bubbleId *big.Int) (*BubBasics, error) {
	data, err := bdb.db.Get(blockHash, getBubBasicsKey(bubbleId))
	if err != nil {
		log.Error("failed to GetBubBasics", "error", err.Error())
		return nil, err
	}
	var basics BubBasics
	if err := rlp.DecodeBytes(data, &basics); err != nil {
		log.Error("failed to decode BubBasics", "error", err.Error())
		return nil, err
	} else {
		return &basics, nil
	}
}

func (bdb *BubbleDB) StoreBubBasics(blockHash common.Hash, bubbleId *big.Int, basics *BubBasics) error {
	if data, err := rlp.EncodeToBytes(basics); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, getBubBasicsKey(bubbleId), data)
	}
}

func (bdb *BubbleDB) DelBubBasics(blockHash common.Hash, bubbleID *big.Int) error {
	if err := bdb.db.Del(blockHash, getBubBasicsKey(bubbleID)); err != nil {
		return err
	}
	return nil
}

func (bdb *BubbleDB) IteratorBubStatus(blockHash common.Hash, ranges int) iterator.Iterator {
	return bdb.db.Ranking(blockHash, BubStatusKeyPrefix, ranges)
}

func (bdb *BubbleDB) GetBubStatus(blockHash common.Hash, bubbleId *big.Int) (*BubStatus, error) {
	data, err := bdb.db.Get(blockHash, getBubStatusKey(bubbleId))
	if err != nil {
		return nil, err
	}
	var bubStatus BubStatus
	if err := rlp.DecodeBytes(data, &bubStatus); err != nil {
		return nil, err
	} else {
		return &bubStatus, nil
	}
}

func (bdb *BubbleDB) StoreBubStatus(blockHash common.Hash, bubbleId *big.Int, bubStatus *BubStatus) error {
	if data, err := rlp.EncodeToBytes(bubStatus); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, getBubStatusKey(bubbleId), data)
	}
}

func (bdb *BubbleDB) IteratorSizedBubbleID(blockHash common.Hash, sizeCode uint8, ranges int) iterator.Iterator {
	return bdb.db.Ranking(blockHash, getSizedBubblePrefix(sizeCode), ranges)

}

func (bdb *BubbleDB) StoreSizedBubbleID(blockHash common.Hash, sizeCode uint8, bubbleId *big.Int) error {
	if data, err := rlp.EncodeToBytes(bubbleId); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, getSizedBubbleKey(sizeCode, bubbleId), data)
	}
}

func (bdb *BubbleDB) DelSizedBubbleID(blockHash common.Hash, sizeCode uint8, bubbleID *big.Int) error {
	if err := bdb.db.Del(blockHash, getSizedBubbleKey(sizeCode, bubbleID)); err != nil {
		return err
	}
	return nil
}

func (bdb *BubbleDB) GetContractByteCode(blockHash common.Hash, address common.Address) ([]byte, error) {
	return bdb.db.Get(blockHash, getByteCodeKey(address))
}

func (bdb *BubbleDB) StoreContractByteCode(blockHash common.Hash, address common.Address, byteCode []byte) error {
	return bdb.db.Put(blockHash, getByteCodeKey(address), byteCode)
}

func (bdb *BubbleDB) IteratorBubContract(blockHash common.Hash, bubbleID *big.Int, ranges int) iterator.Iterator {
	return bdb.db.Ranking(blockHash, getBubContractKey(bubbleID), ranges)
}

func (bdb *BubbleDB) GetBubContract(blockHash common.Hash, bubbleID *big.Int, address common.Address) (*ContractInfo, error) {
	data, err := bdb.db.Get(blockHash, getContractInfoKey(bubbleID, address))
	if err != nil {
		return nil, err
	}

	var contractInfo ContractInfo
	if err := rlp.DecodeBytes(data, &contractInfo); err != nil {
		return nil, err
	} else {
		return &contractInfo, nil
	}
}

func (bdb *BubbleDB) StoreBubContract(blockHash common.Hash, bubbleID *big.Int, contractInfo *ContractInfo) error {
	if data, err := rlp.EncodeToBytes(contractInfo); err != nil {
		return err
	} else {
		return bdb.db.Put(blockHash, getContractInfoKey(bubbleID, contractInfo.Address), data)
	}
}

func (bdb *BubbleDB) DelBubContract(blockHash common.Hash, bubbleID *big.Int, address common.Address) error {
	return bdb.db.Del(blockHash, getContractInfoKey(bubbleID, address))
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
