package contracts

import "github.com/bubblenet/bubble/datavalidator/types"

type InnerChildChainContract struct {
}

func NewInnerChildChainContract() *InnerChildChainContract {
	return nil
}

func (i InnerChildChainContract) GetBubbleId() []uint64 {
	//TODO implement me
	panic("implement me")
}

func (i InnerChildChainContract) GetBubbleValidator(chainId uint64) *types.ChildChainValidator {
	//TODO implement me
	panic("implement me")
}
