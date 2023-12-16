package mock

import "github.com/bubblenet/bubble/datavalidator/types"

type MockChildChain struct {
	GetBubbleIdFn        func() []uint64
	GetBubbleValidatorFn func(chainId uint64) *types.ChildChainValidator
	GetValidatorsFn      func(chainId uint64) map[uint64]*types.Validator
}

func (m MockChildChain) GetBubbleId() []uint64 {
	return m.GetBubbleIdFn()
}

func (m MockChildChain) GetBubbleValidator(chainId uint64) *types.ChildChainValidator {
	return m.GetBubbleValidatorFn(chainId)
}
func (m MockChildChain) GetValidators(chainId uint64) map[uint64]*types.Validator {
	return m.GetValidatorsFn(chainId)
}
