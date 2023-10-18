package bubble

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/rlp"
	"math/big"
)

var (
	AccListKeyPrefix     = []byte("AccList")  // the key prefix of the accounts list of the staking token
	AccAssetKeyPrefix    = []byte("AccAsset") // key prefix of the asset information of the pledge account
	TxHashKeyPrefix      = []byte("TxHash")
	TxHashListKeyPrefix  = []byte("TxHashList") // The key prefix of the transaction hash list
	BubBasicsKeyPrefix   = []byte("BubBasics")  // The key prefix of the bubble basics
	BubStatusKeyPrefix   = []byte("BubState")   // The key prefix of the bubble state
	BubbleSizeBasePrefix = []byte("BubbleSize")
)

// getBubBasicsKey bubble basics that press bubbleId key
func getBubBasicsKey(bubbleID *big.Int) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	return append(BubBasicsKeyPrefix, bid...)
}

// getBubStatusKey bubble state that press bubbleId key
func getBubStatusKey(bubbleID *big.Int) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	return append(BubStatusKeyPrefix, bid...)
}

func getSizedBubblePrefix(sizeCode uint8) []byte {
	size, err := rlp.EncodeToBytes(sizeCode)
	if nil != err {
		return nil
	}
	return append(BubbleSizeBasePrefix, size...)
}

func getSizedBubbleKey(sizeCode uint8, bubbleID *big.Int) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}

	return append(getSizedBubblePrefix(sizeCode), bid...)
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
