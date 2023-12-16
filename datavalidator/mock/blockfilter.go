package mock

import (
	"context"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/datavalidator/types"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/rlp"
	"math/big"
	"sync"
)

type BlockFilter struct {
	sync.Mutex
	blockNumber uint64
	logs        map[common.Hash]*types.MessagePublishedDetail
	sequence    uint64
	nonceMap    map[uint64]uint64
}

func NewBlockFilter(chainId []uint64) *BlockFilter {
	nonceMap := make(map[uint64]uint64)
	for _, id := range chainId {
		nonceMap[id] = 0
	}
	return &BlockFilter{
		blockNumber: 0,
		sequence:    0,
		logs:        make(map[common.Hash]*types.MessagePublishedDetail),
		nonceMap:    nonceMap,
	}
}

func (a *BlockFilter) AddMessagePublished(chainId uint64, limit int) {
	a.Lock()
	defer a.Unlock()
	nonce := a.nonceMap[chainId]
	for i := 0; i < limit; i++ {
		message := &types.MessagePublishedDetail{
			BlockHash: common.BigToHash(big.NewInt(int64(a.blockNumber))),
			TxHash:    common.BigToHash(big.NewInt(int64(nonce))),
			Log: &types.MessagePublished{
				Send:     common.BigToAddress(big.NewInt(int64(chainId))),
				ChainId:  chainId,
				Sequence: a.sequence,
				Nonce:    nonce,
				Payload:  []byte{1, 2, 3},
			},
			Signatures: nil,
		}
		a.sequence++
		nonce++
		a.logs[message.Hash()] = message
		log.Debug("add message published success", "blocknumber", a.blockNumber, "nonce", nonce-1)
	}
	a.nonceMap[chainId] = nonce
	a.blockNumber += 1

}

func (a BlockFilter) BlockNumber() uint64 {
	a.Lock()
	defer a.Unlock()
	return a.blockNumber
}

func (a BlockFilter) RangeFilter(ctx context.Context, begin, to *big.Int) ([]*types.MessagePublishedDetail, error) {
	a.Lock()
	defer a.Unlock()
	hashes := make(map[common.Hash]struct{})
	for i := begin.Int64(); i < to.Int64(); i++ {
		hashes[common.BigToHash(big.NewInt(i))] = struct{}{}
	}
	var logs []*types.MessagePublishedDetail
	for _, log := range a.logs {
		if _, ok := hashes[log.BlockHash]; ok {
			buf, _ := rlp.EncodeToBytes(log)
			var l types.MessagePublishedDetail
			rlp.DecodeBytes(buf, &l)
			logs = append(logs, &l)
		}
	}
	return logs, nil
}

func (a BlockFilter) HasLogOnChain(txHash common.Hash, sequence uint64) *types.MessagePublishedDetail {
	a.Lock()
	defer a.Unlock()
	for _, log := range a.logs {
		if log.Log.Sequence == sequence {
			return log
		}
	}
	return nil
}
