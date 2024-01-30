package bubble

import (
	"github.com/bubblenet/bubble/common"
	"math/big"
)

// State is bubble chain status
type State uint8

const (
	//BuildingState   = 0
	ActiveState     State = 1
	PreReleaseState       = 2
	ReleasedState         = 3
)

type BubbleState struct {
	state State
}

type RemoteCallTask struct {
	BubbleID *big.Int
	TxHash   common.Hash // The transaction hash of the remoteDeployTask
	Caller   common.Address
	Contract common.Address
	Data     []byte
}
