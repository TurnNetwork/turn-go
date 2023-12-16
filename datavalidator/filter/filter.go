package filter

import (
	"context"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/datavalidator/types"
	"github.com/bubblenet/bubble/datavalidator/utils"
	"github.com/bubblenet/bubble/eth/filters"
)

type FilterFunc func(begin, end int64, addresses []common.Address, topics [][]common.Hash) *filters.Filter

type Filter struct {
	address       common.Address
	newFilterFunc FilterFunc
}

func NewFilter(address common.Address, filterFunc FilterFunc) *Filter {
	return &Filter{
		address:       address,
		newFilterFunc: filterFunc,
	}
}

func (f *Filter) RangeFilter(ctx context.Context, begin, to *big.Int) ([]*types.MessagePublishedDetail, error) {
	filter := f.newFilterFunc(begin.Int64(), to.Int64(), []common.Address{f.address}, [][]common.Hash{})
	logs, err := filter.Logs(ctx)
	if err != nil {
		return nil, err
	}
	var details []*types.MessagePublishedDetail
	for _, log := range logs {
		messageLog, err := utils.DecodeMessagePublishedLog(log)
		if err != nil {
			return nil, err
		}
		details = append(details, &types.MessagePublishedDetail{
			BlockHash:  log.BlockHash,
			TxHash:     log.TxHash,
			Log:        messageLog,
			Signatures: nil,
		})
	}
	return details, nil
}
