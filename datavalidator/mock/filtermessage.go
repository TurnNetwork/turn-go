package mock

import (
	"context"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/datavalidator/types"
)

type MockFilterMessage struct {
	RangeFilterFn   func(ctx context.Context, begin, to *big.Int) ([]*types.MessagePublishedDetail, error)
	HasLogOnChainFn func(txHash common.Hash, sequence uint64) *types.MessagePublishedDetail
}

func (m MockFilterMessage) RangeFilter(ctx context.Context, begin, to *big.Int) ([]*types.MessagePublishedDetail, error) {
	return m.RangeFilterFn(ctx, begin, to)
}

func (m MockFilterMessage) HasLogOnChain(txHash common.Hash, sequence uint64) *types.MessagePublishedDetail {
	return m.HasLogOnChainFn(txHash, sequence)
}
