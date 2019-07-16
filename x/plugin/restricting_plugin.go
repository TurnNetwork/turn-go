package plugin

import (
	"bytes"
	"encoding/json"
	"math/big"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/vm"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/x/restricting"
	"github.com/PlatONnetwork/PlatON-Go/x/xcom"
)

var (
	errParamPeriodInvalid = common.NewBizError("param epoch invalid")
	errBalanceNotEnough   = common.NewBizError("balance not enough to restrict")
	errAccountNotFound    = common.NewBizError("account is not found")
)

type RestrictingPlugin struct {
}

var rt *RestrictingPlugin = nil

func RestrictingInstance() *RestrictingPlugin {
	if rt == nil {
		rt = &RestrictingPlugin{}
	}
	return rt
}

/*func ClearRestricting() error {
	if nil == rt {
		return common.NewSysError("the RestrictingPlugin already be nil")
	}
	rt = nil
	return nil
}*/

// BeginBlock does something like check input params before execute transactions,
// in RestrictingPlugin it does nothing.
func (rp *RestrictingPlugin) BeginBlock(blockHash common.Hash, head *types.Header, state xcom.StateDB) error {
	return nil
}

// EndBlock invoke releaseRestricting
func (rp *RestrictingPlugin) EndBlock(blockHash common.Hash, head *types.Header, state xcom.StateDB) error {
	epoch := getLatestEpoch(state)

	expect := epoch + 1
	expectBlock := getBlockNumberByEpoch(expect)

	if expectBlock != head.Number.Uint64() {
		return nil
	}

	log.Info("begin to release restricting", "curr", head.Number, "epoch", expectBlock)
	return rp.releaseRestricting(expect, state)
}

// Confirmed is empty function
func (rp *RestrictingPlugin) Confirmed(block *types.Block) error {
	return nil
}

// AddRestrictingRecord stores four K-V record in StateDB:
// RestrictingInfo: the account info to be released
// ReleaseEpoch:   the number of accounts to be released on the epoch corresponding to the target block height
// ReleaseAccount: the account on the index on the target epoch
// ReleaseAmount: the amount of the account to be released on the target epoch
func (rp *RestrictingPlugin) AddRestrictingRecord(sender common.Address, account common.Address, plans []restricting.RestrictingPlan,
	state xcom.StateDB) error {

	// pre-check
	latest := getLatestEpoch(state)

	totalAmount := new(big.Int) // total restricting amount
	for i := 0; i < len(plans); i++ {
		epoch := plans[i].Epoch
		amount := plans[i].Amount

		if epoch < latest {
			log.Error("param epoch invalid", "epoch", epoch, "latest", latest)
			return errParamPeriodInvalid
		}
		totalAmount = totalAmount.Add(totalAmount, amount)
	}

	if state.GetBalance(sender).Cmp(totalAmount) == -1 {
		log.Error("sender's Balance not enough", "total", totalAmount)
		return errBalanceNotEnough
	}

	// TODO
	var (
		err        error
		epochList  []uint64
		index      uint32
		info       restricting.RestrictingInfo
		accNumbers uint32
	)

	restrictingKey := restricting.GetRestrictingKey(account)
	bAccInfo := state.GetState(account, restrictingKey)

	var newInfo1 restricting.RestrictingInfo

	_ = rlp.Decode(bytes.NewBuffer(bAccInfo), &newInfo1)
	// fmt.Println(bAccInfo)
	// fmt.Println(newInfo1)

	if len(bAccInfo) == 0 {
		log.Debug("restricting record not exist", "account", account.Bytes())

		for i := 0; i < len(plans); i++ {
			epoch := plans[i].Epoch
			amount := plans[i].Amount

			// step1: get account numbers at target epoch
			releaseEpochKey := restricting.GetReleaseEpochKey(epoch)
			bAccNumbers := state.GetState(vm.RestrictingContractAddr, releaseEpochKey)

			if len(bAccNumbers) == 0 {
				accNumbers = uint32(1)
			} else {
				accNumbers = common.BytesToUint32(bAccNumbers) + 1
			}
			index = accNumbers

			// step2: save account numbers at target epoch
			state.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(accNumbers))

			// step3: save account at target index
			releaseAccountKey := restricting.GetReleaseAccountKey(epoch, index)
			state.SetState(vm.RestrictingContractAddr, releaseAccountKey, account.Bytes())

			// step4: save restricting amount at target epoch
			releaseAmountKey := restricting.GetReleaseAmountKey(epoch, account)

			state.SetState(account, releaseAmountKey, amount.Bytes())

			epochList = append(epochList, epoch)
		}

		info.Balance = totalAmount
		info.Debt = big.NewInt(0)
		info.DebtSymbol = false
		info.ReleaseList = epochList

	} else {
		log.Debug("restricting record exist", "account", account.Bytes())

		if err = rlp.Decode(bytes.NewReader(bAccInfo), &info); err != nil {
			log.Error("failed to rlp decode the restricting account", "err", err.Error())
			return common.NewSysError(err.Error())
		}

		for i := 0; i < len(plans); i++ {
			epoch := plans[i].Epoch
			amount := plans[i].Amount

			// step1: get restricting amount at target epoch
			releaseAmountKey := restricting.GetReleaseAmountKey(epoch, account)
			bAmount := state.GetState(account, releaseAmountKey)

			if len(bAmount) == 0 {
				log.Trace("release record not exist on curr epoch ", "account", account, "epoch", epoch)

				releaseEpochKey := restricting.GetReleaseEpochKey(epoch)
				bAccNumbers := state.GetState(vm.RestrictingContractAddr, releaseEpochKey)

				if len(bAccNumbers) == 0 {
					accNumbers = uint32(1)
				} else {
					accNumbers = common.BytesToUint32(bAccNumbers) + 1
				}
				index = accNumbers

				// step2: save account numbers at target epoch
				state.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(accNumbers))

				// step3: save account at target index
				releaseAccountKey := restricting.GetReleaseAccountKey(epoch, index)
				state.SetState(vm.RestrictingContractAddr, releaseAccountKey, account.Bytes())

				info.ReleaseList = append(info.ReleaseList, epoch)

			} else {
				log.Trace("release record exist at curr epoch", "account", account, "epoch", epoch)

				origAmount := new(big.Int)
				origAmount = origAmount.SetBytes(bAmount)
				amount = amount.Add(amount, origAmount)
			}

			// step4: save restricting amount at target epoch
			state.SetState(account, releaseAmountKey, amount.Bytes())
		}

		info.Balance = info.Balance.Add(info.Balance, totalAmount)
	}

	// step5: save restricting account info at target epoch
	bAccInfo, err = rlp.EncodeToBytes(info)
	if err != nil {
		log.Error("failed to rlp encode restricting info", "account", account, "error", err)
		return common.NewSysError(err.Error())
	}

	state.SetState(account, restrictingKey, bAccInfo)

	return nil
}

// PledgeLockFunds transfer the money from the restricting contract account to the staking contract account
func (rp *RestrictingPlugin) PledgeLockFunds(account common.Address, amount *big.Int, state xcom.StateDB) error {

	restrictingKey := restricting.GetRestrictingKey(account)
	bAccInfo := state.GetState(account, restrictingKey)

	if len(bAccInfo) == 0 {
		log.Error("record not found in PledgeLockFunds", "account", account, "funds", amount)
		return errAccountNotFound
	}

	var (
		err  error
		info restricting.RestrictingInfo
	)

	if err := rlp.Decode(bytes.NewReader(bAccInfo), &info); err != nil {
		log.Error("failed to rlp decode the restricting account", "info", bAccInfo, "error", err.Error())
		return common.NewSysError(err.Error())
	}

	if info.Balance.Cmp(amount) == -1 {
		log.Error("Balance of restricting account not enough", "balance", info.Balance, "funds", amount)
		return errBalanceNotEnough
	}

	// sub Balance
	info.Balance = info.Balance.Sub(info.Balance, amount)

	// save restricting account info
	if bAccInfo, err = rlp.EncodeToBytes(info); err != nil {
		log.Error("failed to rlp encode the restricting account", "account", account, "error", err)
		return common.NewSysError(err.Error())
	}
	state.SetState(account, restrictingKey, bAccInfo)

	state.SubBalance(vm.RestrictingContractAddr, amount)
	state.AddBalance(vm.StakingContractAddr, amount)

	return nil
}

// ReturnLockFunds transfer the money from the staking contract account to the restricting contract account
func (rp *RestrictingPlugin) ReturnLockFunds(account common.Address, amount *big.Int, state xcom.StateDB) error {

	restrictingKey := restricting.GetRestrictingKey(account)
	bAccInfo := state.GetState(account, restrictingKey)

	if len(bAccInfo) == 0 {
		log.Error("record not found in ReturnLockFunds", "account", account, "funds", amount)
		return errAccountNotFound
	}

	var (
		err   error
		info  restricting.RestrictingInfo
		repay = new(big.Int) // repay the money owed in the past
		left  = new(big.Int) // money left after the repayment
	)

	if err = rlp.Decode(bytes.NewReader(bAccInfo), &info); err != nil {
		log.Error("failed to rlp encode the restricting account", "error", err.Error())
		return common.NewSysError(err.Error())
	}

	if info.DebtSymbol {
		log.Trace("Balance was owed to release in the past", "account", account, "Debt", info.Debt, "funds", amount)

		if amount.Cmp(info.Debt) == -1 {
			// the money returned back is not enough to repay the money owed to release
			repay = amount
			info.Debt = info.Debt.Sub(info.Debt, amount)

		} else {
			// the money returned back is more than the money owed to release
			repay = info.Debt

			left = left.Sub(amount, info.Debt)
			if left.Cmp(big.NewInt(0)) == 1 {
				info.Balance = info.Balance.Add(info.Balance, left)
			}

			info.Debt = big.NewInt(0)
			info.DebtSymbol = false
		}

	} else {
		log.Trace("directly add Balance while symbol is false", "account", account, "Debt", info.Debt)

		repay = big.NewInt(0)
		left = amount
		info.Balance = info.Balance.Add(info.Balance, left)
	}

	// save restricting account info
	if bAccInfo, err = rlp.EncodeToBytes(info); err != nil {
		log.Error("failed to rlp encode the restricting account", "account", account, "error", err)
		return common.NewSysError(err.Error())
	}
	state.SetState(account, restrictingKey, bAccInfo)

	state.SubBalance(vm.StakingContractAddr, amount)
	if repay.Cmp(big.NewInt(0)) == 1 {
		state.AddBalance(account, repay)
	}
	state.AddBalance(vm.RestrictingContractAddr, left)

	return nil
}

// SlashingNotify modify Debt of restricting account
func (rp *RestrictingPlugin) SlashingNotify(account common.Address, amount *big.Int, state xcom.StateDB) error {

	restrictingKey := restricting.GetRestrictingKey(account)
	bAccInfo := state.GetState(account, restrictingKey)

	if len(bAccInfo) == 0 {
		log.Error("record not found in SlashingNotify", "account", account, "funds", amount)
		return errAccountNotFound
	}

	var (
		err  error
		info restricting.RestrictingInfo
	)

	if err = rlp.Decode(bytes.NewReader(bAccInfo), &info); err != nil {
		log.Error("failed to rlp decode restricting account", "error", err.Error(), "info", bAccInfo)
		return common.NewSysError(err.Error())
	}

	if info.DebtSymbol {
		log.Trace("Balance was owed to release in the past", "account", account, "Debt", info.Debt, "funds", amount)

		if amount.Cmp(info.Debt) < 0 {
			info.Debt = info.Debt.Sub(info.Debt, amount)

		} else {
			info.Debt = info.Debt.Sub(amount, info.Debt)
			info.DebtSymbol = false
		}

	} else {
		info.Debt = info.Debt.Add(info.Debt, amount)
	}

	// save restricting account info
	if bAccInfo, err = rlp.EncodeToBytes(info); err != nil {
		log.Error("failed to encode restricting account", "account", account, "error", err)
		return common.NewSysError(err.Error())
	}
	state.SetState(account, restrictingKey, bAccInfo)

	return nil
}

// releaseRestricting will release restricting plans on target epoch
func (rp *RestrictingPlugin) releaseRestricting(epoch uint64, state xcom.StateDB) error {

	releaseEpochKey := restricting.GetReleaseEpochKey(epoch)
	bAccNumbers := state.GetState(vm.RestrictingContractAddr, releaseEpochKey)

	if len(bAccNumbers) == 0 {
		log.Debug("there is no release record on curr epoch", "epoch", epoch)
		return nil
	}
	numbers := common.BytesToUint32(bAccNumbers)
	log.Debug("many restricting records need release", "epoch", epoch, "records", numbers)

	// TODO
	var (
		info    restricting.RestrictingInfo
		release = new(big.Int) // amount need released
	)

	for index := numbers; index > 0; index-- {

		releaseAccountKey := restricting.GetReleaseAccountKey(epoch, index)
		bAccount := state.GetState(vm.RestrictingContractAddr, releaseAccountKey)
		account := common.BytesToAddress(bAccount)

		log.Trace("begin to release record", "index", index, "account", account)

		restrictingKey := restricting.GetRestrictingKey(account)
		bAccInfo := state.GetState(account, restrictingKey)

		if err := rlp.Decode(bytes.NewReader(bAccInfo), &info); err != nil {
			log.Error("failed to rlp decode restricting account", "error", err.Error(), "info", bAccInfo)
			return common.NewSysError(err.Error())
		}

		releaseAmountKey := restricting.GetReleaseAmountKey(epoch, account)
		bRelease := state.GetState(account, releaseAmountKey)
		release = release.SetBytes(bRelease)

		if info.DebtSymbol {
			log.Debug("Balance is owed to release in the past", "account", account, "Debt", info.Debt, "symbol", info.DebtSymbol)
			info.Debt = info.Debt.Add(info.Debt, release)

		} else {
			temp := new(big.Int)
			if release.Cmp(info.Debt) <= 0 {
				info.Debt = info.Debt.Sub(info.Debt, release)

			} else if release.Cmp(temp.Add(info.Debt, info.Balance)) <= 0 {
				release = release.Sub(release, info.Debt)
				info.Balance = info.Balance.Sub(info.Balance, release)
				info.Debt = big.NewInt(0)

				log.Trace("show balance", "balance", info.Balance)

				state.SubBalance(vm.RestrictingContractAddr, release)
				state.AddBalance(account, release)

			} else {
				temp := info.Balance

				release = release.Sub(release, info.Balance)
				info.Balance = big.NewInt(0)
				info.Debt = info.Debt.Sub(release, info.Debt)
				info.DebtSymbol = true

				state.SubBalance(vm.RestrictingContractAddr, temp)
				state.AddBalance(account, temp)
			}
		}

		// delete ReleaseAmount
		state.SetState(account, releaseAmountKey, []byte{})

		// delete ReleaseAccount
		state.SetState(vm.RestrictingContractAddr, releaseAccountKey, []byte{})

		// delete epoch in ReleaseList
		// In general, the first epoch is released first.
		// info.ReleaseList = info.ReleaseList[1:]
		for i, target := range info.ReleaseList {
			if target == epoch {
				info.ReleaseList = append(info.ReleaseList[:i], info.ReleaseList[i+1:]...)
				break
			}
		}

		// restore restricting info
		if bNewInfo, err := rlp.EncodeToBytes(info); err != nil {
			log.Error("failed to rlp encode new info while release", "account", account, "info", info)
			return common.NewSysError(err.Error())
		} else {
			state.SetState(account, restrictingKey, bNewInfo)
		}
	}

	// delete ReleaseEpoch
	state.SetState(vm.RestrictingContractAddr, releaseEpochKey, []byte{})

	return nil
}

func (rp *RestrictingPlugin) GetRestrictingInfo(account common.Address, state xcom.StateDB) ([]byte, error) {

	restrictingKey := restricting.GetRestrictingKey(account)
	bAccInfo := state.GetState(account, restrictingKey)

	if len(bAccInfo) == 0 {
		log.Error("record not found in GetRestrictingInfo", "account", account)
		return []byte{}, errAccountNotFound
	}

	var (
		bAmount          []byte
		info             restricting.RestrictingInfo
		plan             restricting.ReleaseAmountInfo
		plans            []restricting.ReleaseAmountInfo
		releaseAmountKey []byte
		result           restricting.Result
	)

	if err := rlp.Decode(bytes.NewReader(bAccInfo), &info); err != nil {
		log.Error("failed to rlp encode the restricting account", "error", err.Error(), "info", bAccInfo)
		return []byte{}, common.NewSysError(err.Error())
	}

	var amount = new(big.Int)
	for i := 0; i < len(info.ReleaseList); i++ {
		epoch := info.ReleaseList[i]
		releaseAmountKey = restricting.GetReleaseAmountKey(epoch, account)
		bAmount = state.GetState(account, releaseAmountKey)

		plan.Height = getBlockNumberByEpoch(epoch)
		plan.Amount = amount.SetBytes(bAmount)
		plans = append(plans, plan)
	}

	bPlans, err := json.Marshal(plans)
	if err != nil {
		log.Error("failed to Marshal restricting result")
		return []byte{}, err
	}

	result.Balance = info.Balance
	result.Debt = info.Debt
	result.Slash = big.NewInt(0)
	result.Staking = big.NewInt(0)
	result.Entry = bPlans

	log.Trace("get restricting result", "account", account, "result", result)

	return rlp.EncodeToBytes(result)
}

// state DB operation
func SetLatestEpoch(stateDb xcom.StateDB, epoch uint64) {
	key := restricting.GetLatestEpochKey()
	stateDb.SetState(vm.RestrictingContractAddr, key, common.Uint64ToBytes(epoch))
}

func getLatestEpoch(stateDb xcom.StateDB) uint64 {
	key := restricting.GetLatestEpochKey()
	bEpoch := stateDb.GetState(vm.RestrictingContractAddr, key)

	if len(bEpoch) == 0 {
		return 0
	} else {
		return common.BytesToUint64(bEpoch)
	}
}

func getBlockNumberByEpoch(epoch uint64) uint64 {
	return epoch * xcom.ConsensusSize() * xcom.EpochSize()
}
