package bubble

import "github.com/bubblenet/bubble/common"

var (
	ErrBubbleNotExist           = common.NewBizError(322001, "The bubble is not exist")
	ErrSenderIsNotCreator       = common.NewBizError(322002, "transaction sender must be the contract creator")
	ErrBubbleUnableRelease      = common.NewBizError(322003, "The bubble is unable to release")
	ErrOperatorL1IsInsufficient = common.NewBizError(322004, "The operator of layer 1 is insufficient")
	ErrOperatorL2IsInsufficient = common.NewBizError(322005, "The operator of layer 2 is insufficient")
	ErrMicroNodeIsInsufficient  = common.NewBizError(322006, "The micro node is insufficient")
	ErrEmptyContractCode        = common.NewBizError(322007, "The contract code is empty or abnormal")
	ErrContractReturns          = common.NewBizError(322008, "The contract returns an error")
	ErrBubbleIsPreRelease       = common.NewBizError(322009, "the bubble is ready to release")
	ErrContractIsExist          = common.NewBizError(322010, "the contract is exist")
	ErrContractNotExist         = common.NewBizError(322010, "the contract is not exist")
	ErrSenderIsNotOperator      = common.NewBizError(322011, "transaction sender must be the bubble operator")
)
