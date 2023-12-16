package types

import (
	"context"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/p2p/discover"
	"math/big"
)

type BlockState interface {
	BlockNumber() uint64
}
type P2PServer interface {
	AddPeer(node *discover.Node)
	RemovePeer(node *discover.Node)
}
type FilterMessage interface {
	RangeFilter(ctx context.Context, begin, to *big.Int) ([]*MessagePublishedDetail, error)
}

type ValidatorContract interface {
	ValidatorSet() []*Validator
	GetValidator(addr *bls.PublicKey) *Validator
}

type ChildChainContract interface {
	GetBubbleId() []uint64
	GetBubbleValidator(chainId uint64) *ChildChainValidator
	GetValidators(chainId uint64) map[uint64]*Validator
}
