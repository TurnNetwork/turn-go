package chain

import "github.com/bubblenet/bubble/core"

type ChainBlockState struct {
	c *core.BlockChain
}

func NewChainBlockState(chain *core.BlockChain) *ChainBlockState {
	return &ChainBlockState{
		c: chain,
	}
}
func (c ChainBlockState) BlockNumber() uint64 {
	return c.c.CurrentBlock().NumberU64()
}
