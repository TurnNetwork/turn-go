package bubble

import "github.com/bubblenet/bubble/common"

var (
	bubbleKeyPrefix               = []byte("Bubble")
	stakingTokenAccListKeyPrefix  = []byte("StakingTokenAccList") // the key prefix of the accounts list of the staking token
	stakingTokenKeyPrefix         = []byte("StakingToken")
	stakingTokenAccAssetKeyPrefix = []byte("StakingTokenAccAsset") // key prefix of the asset information of the pledge account
)

func GetBubbleKey(bubbleID uint32) []byte {
	bid := common.Uint32ToBytes(bubbleID)
	return append(bubbleKeyPrefix, bid...)
}

// StakingTokenInBubKey Pledge the token to the key in the specified bubble
func StakingTokenInBubKey(bubbleID uint32) []byte {
	bid := common.Uint32ToBytes(bubbleID)
	return append(stakingTokenKeyPrefix, bid...)
}

// StakingTokenAccListKey key that generates the list of accounts for the pledged tokens inside the bubble
func StakingTokenAccListKey(bubbleID uint32) []byte {
	bid := common.Uint32ToBytes(bubbleID)
	return append(stakingTokenAccListKeyPrefix, bid...)
}

// StakingTokenAccAssetKey Generate the keys of the assets secured by the accounts in the bubble
func StakingTokenAccAssetKey(bubbleID uint32, account common.Address) []byte {
	bid := common.Uint32ToBytes(bubbleID)
	key := append(stakingTokenAccAssetKeyPrefix, bid...)
	return append(key, account.Bytes()...)
}
