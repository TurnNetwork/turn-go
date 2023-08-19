package bubble

import "github.com/bubblenet/bubble/common"

var (
	ErrBubbleNotExist      = common.NewBizError(322001, "The bubble is not exist")
	ErrSenderIsNotCreator  = common.NewBizError(322002, "transaction sender must be the bubble creator")
	ErrBubbleUnableRelease = common.NewBizError(322003, "The bubble is unable to release")
	ErrAccountNoEnough     = common.NewBizError(330000, "The account balance is insufficient")
	ErrStakingAccount      = common.NewBizError(330001, "The sender of the pledged token transaction is not the person of the pledged account")
	ErrERC20NoExist        = common.NewBizError(330002, "erc20 contract address does not exist")
	ErrEVMOrContractEmpty  = common.NewBizError(330003, "Evm or Contract is nil")
	ErrNoExecutableVM      = common.NewBizError(330004, "there is no executable EVM or WASM interpreter")
)
