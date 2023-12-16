package mock

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/datavalidator/types"
)

type MockValidator struct {
	ValidatorSetFn func() []*types.Validator
	IsValidatorFn  func(id string) bool
	GetValidatorFn func(addr common.Address) *types.Validator
}

func (m MockValidator) GetValidator(addr common.Address) *types.Validator {
	return m.GetValidatorFn(addr)
}

func (m MockValidator) ValidatorSet() []*types.Validator {
	return m.ValidatorSetFn()
}

func (m MockValidator) IsValidator(id string) bool {
	return m.IsValidatorFn((id))
}
