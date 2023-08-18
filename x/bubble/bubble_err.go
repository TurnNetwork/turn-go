package bubble

import "github.com/bubblenet/bubble/common"

var (
	ErrBubbleNotExist      = common.NewBizError(322001, "The bubble is not exist")
	ErrSenderIsNotCreator  = common.NewBizError(322002, "transaction sender must be the bubble creator")
	ErrBubbleUnableRelease = common.NewBizError(322003, "The bubble is unable to release")
)
