package bubble

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/rlp"
	"math/big"
)

var (
	bubbleKeyPrefix     = []byte("Bubble")
	AccListKeyPrefix    = []byte("AccList")  // the key prefix of the accounts list of the staking token
	AccAssetKeyPrefix   = []byte("AccAsset") // key prefix of the asset information of the pledge account
	TxHashKeyPrefix     = []byte("TxHash")
	TxHashListKeyPrefix = []byte("TxHashList") // The key prefix of the transaction hash list
	BubBasicsKeyPrefix  = []byte("BubBasics")  // The key prefix of the bubble basics
	BubStateKeyPrefix   = []byte("BubState")   // The key prefix of the bubble state
)

func GetBubbleKey(bubbleID *big.Int) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	return append(bubbleKeyPrefix, bid...)
}

// BasicsByBubKey bubble basics that press bubbleId key
func BasicsByBubKey(bubbleID *big.Int) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	return append(BubBasicsKeyPrefix, bid...)
}

// StateByBubKey bubble state that press bubbleId key
func StateByBubKey(bubbleID *big.Int) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	return append(BubStateKeyPrefix, bid...)
}

// AccListByBubKey List of accounts that press bubble's key
func AccListByBubKey(bubbleID *big.Int) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	return append(AccListKeyPrefix, bid...)
}

// AccAssetByBubKey The key for the specified account inside the bubble
func AccAssetByBubKey(bubbleID *big.Int, account common.Address) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	key := append(AccAssetKeyPrefix, bid...)
	return append(key, account.Bytes()...)
}

// TxHashByBubKey The key for the specified TxHash inside the bubble
func TxHashByBubKey(bubbleID *big.Int, txHash common.Hash) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	key := append(TxHashKeyPrefix, bid...)
	return append(key, txHash.Bytes()...)
}

// TxHashListByBubKey Specifies the transaction type of bubble to generate the key of the transaction hash list
func TxHashListByBubKey(bubbleID *big.Int, txType BubTxType) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	txTypeRLP, err := rlp.EncodeToBytes(txType)
	if nil != err {
		return nil
	}
	key := append(TxHashListKeyPrefix, bid...)
	return append(key, txTypeRLP...)
}
