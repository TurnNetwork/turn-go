package bubble

import "github.com/bubblenet/bubble/common"

var (
	bubbleKeyPrefix = []byte("Bubble")
)

func GetBubbleKey(bubbleID uint32) []byte {
	bid := common.Uint32ToBytes(bubbleID)
	return append(bubbleKeyPrefix, bid...)
}
