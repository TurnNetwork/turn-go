package bubble

import (
	"github.com/bubblenet/bubble/common"
)

var (
	BubContractPrefix = []byte("BubbleContract")
)

func getBubContractKey(address *common.Address) []byte {
	return append(BubContractPrefix, address.Bytes()...)
}
