package bubble

import "github.com/bubblenet/bubble/common"

var (
	bubbleKeyPrefix   = []byte("Bubble")
	AccListKeyPrefix  = []byte("AccList")  // the key prefix of the accounts list of the staking token
	AccAssetKeyPrefix = []byte("AccAsset") // key prefix of the asset information of the pledge account
)

func GetBubbleKey(bubbleID uint32) []byte {
	bid := common.Uint32ToBytes(bubbleID)
	return append(bubbleKeyPrefix, bid...)
}

// AccListByBubKey List of accounts that press bubble's key
func AccListByBubKey(bubbleID uint32) []byte {
	bid := common.Uint32ToBytes(bubbleID)
	return append(AccListKeyPrefix, bid...)
}

// AccAssetByBubKey The key for the specified account inside the bubble
func AccAssetByBubKey(bubbleID uint32, account common.Address) []byte {
	bid := common.Uint32ToBytes(bubbleID)
	key := append(AccAssetKeyPrefix, bid...)
	return append(key, account.Bytes()...)
}
