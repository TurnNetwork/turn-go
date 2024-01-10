package bubble

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/rlp"
	"math/big"
)

var (
	AccListKeyPrefix       = []byte("AccList")  // the key prefix of the accounts list of the staking token
	AccAssetKeyPrefix      = []byte("AccAsset") // key prefix of the asset information of the pledge account
	TxHashKeyPrefix        = []byte("TxHash")
	TxHashListKeyPrefix    = []byte("TxHashList")       // The key prefix of the transaction hash list
	BasicsInfoKeyPrefix    = []byte("BubBasicsInfo")    // The key prefix of the bubble basics
	ValidatorInfoKeyPrefix = []byte("BubValidatorInfo") // The key prefix of the bubble basics
	SizeRootPrefix         = []byte("BubSize")
	ByteCodePrefix         = []byte("BubBytecode")
	ContractInfoPrefix     = []byte("BubContractInfo")
)

// getBasicsInfoKey bubble basics that press bubbleId key
func getBasicsInfoKey(bubbleID *big.Int) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	return append(BasicsInfoKeyPrefix, bid...)
}

// getValidatorInfoKey bubble basics that press bubbleId key
func getValidatorInfoKey(bubbleID *big.Int) []byte {
	bid, err := rlp.EncodeToBytes(bubbleID)
	if nil != err {
		return nil
	}
	return append(ValidatorInfoKeyPrefix, bid...)
}

func getBubbleSizePrefix(size Size) []byte {
	data, err := rlp.EncodeToBytes(size)
	if nil != err {
		return nil
	}
	return append(SizeRootPrefix, data...)
}

func getSizedBubbleKey(size Size, bubbleID *big.Int) []byte {
	return append(getBubbleSizePrefix(size), bubbleID.Bytes()...)
}

func getBubContractKey(bubbleID *big.Int) []byte {
	return append(ContractInfoPrefix, bubbleID.Bytes()...)
}

func getContractInfoKey(bubbleID *big.Int, address common.Address) []byte {
	return append(getBubContractKey(bubbleID), address.Bytes()...)
}

func getByteCodeKey(address common.Address) []byte {
	return append(ByteCodePrefix, address.Bytes()...)
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
func TxHashListByBubKey(bubbleID *big.Int, txType TxType) []byte {
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
