package bubble

import (
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/rlp"
)

type DB struct {
	db snapshotdb.DB
}

func NewDB() *DB {
	return &DB{
		db: snapshotdb.Instance(),
	}
}

func (db *DB) IteratorBasicsInfo(blockHash common.Hash, ranges int) iterator.Iterator {
	return db.db.Ranking(blockHash, BasicsInfoKeyPrefix, ranges)
}

func (db *DB) GetBasicsInfo(blockHash common.Hash, bubbleId *big.Int) (*BasicsInfo, error) {
	data, err := db.db.Get(blockHash, getBasicsInfoKey(bubbleId))
	if err != nil {
		log.Error("failed to GetBasicsInfo", "error", err.Error())
		return nil, err
	}
	var basics BasicsInfo
	if err := rlp.DecodeBytes(data, &basics); err != nil {
		log.Error("failed to decode BasicsInfo", "error", err.Error())
		return nil, err
	} else {
		return &basics, nil
	}
}

func (db *DB) StoreBasicsInfo(blockHash common.Hash, bubbleId *big.Int, basics *BasicsInfo) error {
	if data, err := rlp.EncodeToBytes(basics); err != nil {
		return err
	} else {
		return db.db.Put(blockHash, getBasicsInfoKey(bubbleId), data)
	}
}

func (db *DB) DelBasicsInfo(blockHash common.Hash, bubbleID *big.Int) error {
	if err := db.db.Del(blockHash, getBasicsInfoKey(bubbleID)); err != nil {
		return err
	}
	return nil
}

/*
ValidatorInfo
*/

func (db *DB) GetValidatorInfo(blockHash common.Hash, bubbleId *big.Int) (*ValidatorInfo, error) {
	data, err := db.db.Get(blockHash, getValidatorInfoKey(bubbleId))
	if err != nil {
		log.Error("failed to GetValidatorInfo", "error", err.Error())
		return nil, err
	}
	var validator ValidatorInfo
	if err := rlp.DecodeBytes(data, &validator); err != nil {
		log.Error("failed to decode ValidatorInfo", "error", err.Error())
		return nil, err
	} else {
		return &validator, nil
	}
}

func (db *DB) StoreValidatorInfo(blockHash common.Hash, bubbleId *big.Int, validator *ValidatorInfo) error {
	if data, err := rlp.EncodeToBytes(validator); err != nil {
		return err
	} else {
		return db.db.Put(blockHash, getValidatorInfoKey(bubbleId), data)
	}
}

func (db *DB) DelValidatorInfo(blockHash common.Hash, bubbleID *big.Int) error {
	if err := db.db.Del(blockHash, getValidatorInfoKey(bubbleID)); err != nil {
		return err
	}
	return nil
}

func (db *DB) GetJoinBubble(blockHash common.Hash, nodeId discover.NodeID) (*big.Int, error) {
	data, err := db.db.Get(blockHash, getJoinBubbleKey(nodeId))
	if err != nil {
		return nil, err
	}
	bubbleId := new(big.Int).SetBytes(data)

	return bubbleId, nil
}

func (db *DB) StoreJoinBubble(blockHash common.Hash, nodeId discover.NodeID, bubbleId *big.Int) error {
	return db.db.Put(blockHash, getJoinBubbleKey(nodeId), bubbleId.Bytes())
}

func (db *DB) DelJoinBubble(blockHash common.Hash, nodeId discover.NodeID) error {
	return db.db.Del(blockHash, getJoinBubbleKey(nodeId))
}

func (db *DB) IteratorBubbleIdBySize(blockHash common.Hash, size Size, ranges int) iterator.Iterator {
	return db.db.Ranking(blockHash, getBubbleSizePrefix(size), ranges)

}

func (db *DB) StoreBubbleIdBySize(blockHash common.Hash, size Size, bubbleId *big.Int) error {
	if data, err := rlp.EncodeToBytes(bubbleId); err != nil {
		return err
	} else {
		return db.db.Put(blockHash, getSizedBubbleKey(size, bubbleId), data)
	}
}

func (db *DB) DelBubbleIdBySize(blockHash common.Hash, size Size, bubbleID *big.Int) error {
	if err := db.db.Del(blockHash, getSizedBubbleKey(size, bubbleID)); err != nil {
		return err
	}
	return nil
}

func (db *DB) GetByteCode(blockHash common.Hash, address common.Address) ([]byte, error) {
	return db.db.Get(blockHash, getByteCodeKey(address))
}

func (db *DB) StoreByteCode(blockHash common.Hash, address common.Address, byteCode []byte) error {
	return db.db.Put(blockHash, getByteCodeKey(address), byteCode)
}

func (db *DB) IteratorContractInfo(blockHash common.Hash, bubbleID *big.Int, ranges int) iterator.Iterator {
	return db.db.Ranking(blockHash, getBubContractKey(bubbleID), ranges)
}

func (db *DB) GetContractInfo(blockHash common.Hash, bubbleID *big.Int, address common.Address) (*ContractInfo, error) {
	data, err := db.db.Get(blockHash, getContractInfoKey(bubbleID, address))
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

func (db *DB) StoreContractInfo(blockHash common.Hash, bubbleID *big.Int, contractInfo *ContractInfo) error {
	if data, err := rlp.EncodeToBytes(contractInfo); err != nil {
		return err
	} else {
		return db.db.Put(blockHash, getContractInfoKey(bubbleID, contractInfo.Address), data)
	}
}

func (db *DB) DelContractInfo(blockHash common.Hash, bubbleID *big.Int, address common.Address) error {
	return db.db.Del(blockHash, getContractInfoKey(bubbleID, address))
}

func (db *DB) GetAccListOfBub(blockHash common.Hash, bubbleId *big.Int) ([]common.Address, error) {
	data, err := db.db.Get(blockHash, AccListByBubKey(bubbleId))
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
func (db *DB) StoreAccListOfBub(blockHash common.Hash, bubbleId *big.Int, accounts []common.Address) error {
	if data, err := rlp.EncodeToBytes(accounts); err != nil {
		return err
	} else {
		return db.db.Put(blockHash, AccListByBubKey(bubbleId), data)
	}
}

func (db *DB) GetAccAssetOfBub(blockHash common.Hash, bubbleId *big.Int, account common.Address) (*AccountAsset, error) {
	data, err := db.db.Get(blockHash, AccAssetByBubKey(bubbleId, account))
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

func (db *DB) StoreAccAssetToBub(blockHash common.Hash, bubbleId *big.Int, accAsset AccountAsset) error {
	if data, err := rlp.EncodeToBytes(accAsset); err != nil {
		return err
	} else {
		return db.db.Put(blockHash, AccAssetByBubKey(bubbleId, accAsset.Account), data)
	}
}

func (db *DB) StoreTxHashListToBub(blockHash common.Hash, bubbleID *big.Int, txHashList []common.Hash, txType TxType) error {
	if data, err := rlp.EncodeToBytes(txHashList); err != nil {
		return err
	} else {
		return db.db.Put(blockHash, TxHashListByBubKey(bubbleID, txType), data)
	}
}

func (db *DB) GetTxHashListByBub(blockHash common.Hash, bubbleID *big.Int, txType TxType) (*[]common.Hash, error) {
	data, err := db.db.Get(blockHash, TxHashListByBubKey(bubbleID, txType))
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

func (db *DB) StoreL2HashToL1Hash(blockHash common.Hash, bubbleID *big.Int, L1TxHash common.Hash, L2TxHash common.Hash) error {
	if data, err := rlp.EncodeToBytes(L1TxHash); err != nil {
		return err
	} else {
		return db.db.Put(blockHash, TxHashByBubKey(bubbleID, L2TxHash), data)
	}
}

func (db *DB) GetL1HashByL2Hash(blockHash common.Hash, bubbleID *big.Int, L2TxHash common.Hash) (*common.Hash, error) {
	data, err := db.db.Get(blockHash, TxHashByBubKey(bubbleID, L2TxHash))
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
