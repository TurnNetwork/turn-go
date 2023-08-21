package bubble

import "github.com/bubblenet/bubble/common"

var (
	ErrBubbleNotExist         = common.NewBizError(322001, "The bubble is not exist")
	ErrSenderIsNotCreator     = common.NewBizError(322002, "transaction sender must be the bubble creator")
	ErrBubbleUnableRelease    = common.NewBizError(322003, "The bubble is unable to release")
	ErrAccountNoEnough        = common.NewBizError(330000, "The account balance is insufficient")
	ErrStakingAccount         = common.NewBizError(330001, "The sender of the pledged token transaction is not the person of the pledged account")
	ErrERC20NoExist           = common.NewBizError(330002, "erc20 contract address does not exist")
	ErrEVMOrContractEmpty     = common.NewBizError(330003, "Evm or Contract is nil")
	ErrNoExecutableVM         = common.NewBizError(330004, "there is no executable EVM or WASM interpreter")
	ErrCannotSettled          = common.NewBizError(330005, "the bubble has been released and cannot be settled")
	ErrIsNotSubChainOpAddr    = common.NewBizError(330006, "the transaction sender is not the main chain operator address")
	ErrSettleAccListIncLength = common.NewBizError(330007, "the length of the address participating in the settlement is incorrect")
	ErrSettleAccNoExist       = common.NewBizError(330008, "settlement account does not exist in the bubble")
	ErrStoreAccAssetToBub     = common.NewBizError(330009, "failed to Store the latest information about the staking assets of the account into bubble")
	ErrStoreL2HashToL1Hash    = common.NewBizError(330010, "failed to store sub-chain settlement transaction hash to main-chain settlement transaction hash")
	ErrBubbleIsNotRelease     = common.NewBizError(330011, "bubble cannot redeem the token until it is released")
	ErrEncodeTransferData     = common.NewBizError(330012, "Failed to generate data for ERC20 transfer interface")
	ErrEVMExecERC20           = common.NewBizError(330013, "evm fails to execute erc20 contract transfer interface")
	ErrStoreAccAsset          = common.NewBizError(330014, "failure to save account asset information to the specified bubble")
	ErrBubbleIsRelease        = common.NewBizError(330015, "bubble has released, Can not StakingToken token to the bubble")
)
