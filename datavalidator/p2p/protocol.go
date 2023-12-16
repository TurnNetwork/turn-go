package p2p

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/datavalidator/types"
	"math"
)

const (
	SignMessageType              = 0x00
	HeartbeatMessageType         = 0x01
	SignedObservationRequestType = 0x02
	SignedObservationType        = 0x03
	SignedWithQuorumType         = 0x04
)

func MessageType(msg interface{}) uint64 {
	switch msg.(type) {
	case *SignMessageMsg:
		return SignMessageType
	case *Heartbeat:
		return HeartbeatMessageType
	case *SignedObservationRequest:
		return SignedObservationRequestType
	case *SignedObservation:
		return SignedObservationType
	case *SignedWithQuorum:
		return SignedWithQuorumType
	}
	return math.MaxUint64
}

type SignMessageMsg struct {
	SignMessageData types.MessagePublishedDetail
	Signature       *types.Signature
}

type Heartbeat struct {
	Counter   uint64
	Timestamp uint64
	Version   string
	BlsPub    string
	//ValidatorAddr common.Address
	BootTimestamp uint64
	ScanBlock     uint64
}

type SignedObservationRequest struct {
	ChainId uint64
	ID      common.Hash
}

type SignedObservation struct {
	ChainId    uint64
	ID         common.Hash
	Signatures []*types.Signature
}

type SignedWithQuorum struct {
	ChainId    uint64
	ID         common.Hash
	Signatures []*types.Signature
}
