package contracts

import (
	"github.com/bubblenet/bubble/datavalidator/types"
)

type InnerValidator struct {
}

func NewInnerValidator() *InnerValidator {
	return nil
}

func (i InnerValidator) ValidatorSet() []*types.Validator {
	//TODO implement me
	panic("implement me")
}

func (i InnerValidator) GetValidator() *types.Validator {
	return nil
}
func (i InnerValidator) GetValidators(chainId uint64) map[uint64]*types.Validator {
	//TODO implement me
	panic("implement me")
}
