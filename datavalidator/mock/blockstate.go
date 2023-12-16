package mock

type MockBlockState struct {
	BlockNumberFn func() uint64
}

func (m MockBlockState) BlockNumber() uint64 {
	return m.BlockNumberFn()
}
