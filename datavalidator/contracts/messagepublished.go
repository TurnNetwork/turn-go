package contracts

import (
	"context"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/datavalidator/types"
	"math/big"
)

type InnerMessagePublished struct {
}

func NewInnerMessagePublished() *InnerMessagePublished {
	return nil
}

func (i InnerMessagePublished) RangeFilter(ctx context.Context, begin, to *big.Int) ([]*types.MessagePublishedDetail, error) {
	//TODO implement me
	panic("implement me")
}

func (i InnerMessagePublished) HasLogOnChain(txHash common.Hash, sequence uint64) *types.MessagePublishedDetail {
	//TODO implement me
	panic("implement me")
}
