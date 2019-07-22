package plugin_test

import (
	"bytes"
	"encoding/json"
	"math/big"
	"reflect"
	"testing"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/vm"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/x/plugin"
	"github.com/PlatONnetwork/PlatON-Go/x/restricting"
	"github.com/PlatONnetwork/PlatON-Go/x/xcom"
	"github.com/PlatONnetwork/PlatON-Go/x/xutil"
)

var (
	errParamPeriodInvalid = common.NewBizError("param epoch invalid")
	errBalanceNotEnough   = common.NewBizError("balance not enough to restrict")
	errAccountNotFound    = common.NewBizError("account is not found")
)

func showRestrictingAccountInfo(t *testing.T, state xcom.StateDB, account common.Address) {
	restrictingKey := restricting.GetRestrictingKey(account)
	bAccInfo := state.GetState(account, restrictingKey)

	if len(bAccInfo) == 0 {
		t.Logf("Restricting account not found, account: %v", account.String())
		return
	}

	var info restricting.RestrictingInfo
	if err := rlp.Decode(bytes.NewBuffer(bAccInfo), &info); err != nil {
		t.Fatalf("rlp decode info failed, info bytes: %+v", bAccInfo)
	}

	t.Log("actually balance of restrict account: ", info.Balance)
	t.Log("actually debt    of restrict account: ", info.Debt)
	t.Log("actually symbol  of restrict account: ", info.DebtSymbol)
	t.Log("actually list    of restrict account: ", info.ReleaseList)
}

func showReleaseEpoch(t *testing.T, state xcom.StateDB, epoch uint64) {
	releaseEpochKey := restricting.GetReleaseEpochKey(epoch)
	bAccNumbers := state.GetState(vm.RestrictingContractAddr, releaseEpochKey)

	if len(bAccNumbers) == 0 {
		t.Logf("release Epoch record not found, epoch: %d", epoch)
		return
	} else {
		t.Logf("actually account numbers of release epoch %d: %d", epoch, common.BytesToUint32(bAccNumbers))
	}

	for i := uint32(0); i < common.BytesToUint32(bAccNumbers); i++ {

		index := i + 1
		releaseAccountKey := restricting.GetReleaseAccountKey(epoch, index)
		bReleaseAcc := state.GetState(vm.RestrictingContractAddr, releaseAccountKey)

		if len(bReleaseAcc) == 0 {
			panic("system error, release account can't empty")
		} else {
			t.Logf("actually release accounts of epoch %d: %v", epoch, common.BytesToAddress(bReleaseAcc).String())
		}
	}
}

func showReleaseAmount(t *testing.T, state xcom.StateDB, account common.Address, epoch uint64) {
	releaseAmountKey := restricting.GetReleaseAmountKey(epoch, account)
	bAmount := state.GetState(account, releaseAmountKey)

	amount := new(big.Int)
	if len(bAmount) == 0 {
		t.Logf("record of restricting account amount not found, account: %v, epoch: %d", account.String(), epoch)
	} else {
		t.Logf("actually release amount of account [%s]: %v", account.String(), amount.SetBytes(bAmount))
	}
}

func TestRestrictingPlugin_EndBlock(t *testing.T) {

	// case1: blockChain not arrived settle block height
	{
		xcom.GetEc(xcom.DefaultDeveloperNet)
		stateDb := buildStateDB(t)

		buildDbRestrictingPlan(t, stateDb)
		head := types.Header{Number: big.NewInt(1)}

		err := plugin.RestrictingInstance().EndBlock(common.Hash{}, &head, stateDb)

		// show expected result
		t.Log("expected do nothing")

		if err != nil {
			t.Fatalf("The case1 of EndBlock failed. function returns error: %s", err.Error())
		} else {
			t.Logf("actually do nothing")
			t.Log("=====================")
			t.Log("case1 pass")
		}
	}

	// case2: blockChain arrived settle block height, restricting plan not exist
	{
		stateDb := buildStateDB(t)
		blockNumber := uint64(1) * xutil.CalcBlocksEachEpoch()

		head := types.Header{Number: big.NewInt(int64(blockNumber))}
		err := plugin.RestrictingInstance().EndBlock(common.Hash{}, &head, stateDb)

		// show expected result
		t.Logf("expected do nothing")

		if err != nil {
			t.Fatalf("The case2 of EndBlock failed. function returns error: %s", err.Error())
		} else {
			t.Logf("actually do nothing")
			t.Log("=====================")
			t.Log("case2 pass")
		}

	}

	// case3: blockChain arrived settle block height, restricting plan exist, debt symbol is false,
	// and debt symbol is false, debt more than release amount
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		blockNumber := uint64(1) * xutil.CalcBlocksEachEpoch()

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(1E18)
		info.Debt = big.NewInt(2E18)
		info.DebtSymbol = false
		info.ReleaseList = []uint64{1, 2}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		for _, epoch := range info.ReleaseList {
			// store epoch
			releaseEpochKey := restricting.GetReleaseEpochKey(epoch)
			stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

			// store release account
			releaseAccountKey := restricting.GetReleaseAccountKey(epoch, uint32(1))
			stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())

			// store release amount
			releaseAmountKey := restricting.GetReleaseAmountKey(epoch, restrictingAcc)
			amount := big.NewInt(int64(epoch) * 1000000000000000000)
			stateDb.SetState(restrictingAcc, releaseAmountKey, amount.Bytes())
		}

		stateDb.AddBalance(vm.RestrictingContractAddr, big.NewInt(1E18))

		// do EndBlock
		head := types.Header{Number: big.NewInt(int64(blockNumber))}
		err = plugin.RestrictingInstance().EndBlock(common.Hash{}, &head, stateDb)

		t.Log("=====================")
		t.Log("expected case3 of ReturnLockFunds success")
		t.Log("expected balance of restricting account:", big.NewInt(0))
		t.Log("expected balance of restricting contract contract:", big.NewInt(1E18))
		t.Log("expected balance of restrict account: ", big.NewInt(1E18))
		t.Log("expected debt    of restrict account: ", big.NewInt(1E18))
		t.Log("expected symbol  of restrict account: ", false)
		t.Log("expected list    of restrict account: ", []uint64{2})
		t.Log("=====================")
		t.Log("expected [release Epoch record not found, epoch: 1]")
		t.Log("expected [record of restricting account amount not found, account: 0x740cE31B3fAc20Dac379dB243021A51E80AaDd24, epoch: 1]")
		t.Log("expected account numbers of release epoch 2:", 1)
		t.Logf("expected release accounts of epoch 2: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(2E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case3 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case3 return success!")
			t.Log("actually balance of restricting account:", stateDb.GetBalance(restrictingAcc))
			t.Log("actually balance of restricting contract account:", stateDb.GetBalance(vm.RestrictingContractAddr))

			showRestrictingAccountInfo(t, stateDb, restrictingAcc)
			for _, epoch := range info.ReleaseList {
				showReleaseEpoch(t, stateDb, epoch)
				showReleaseAmount(t, stateDb, restrictingAcc, epoch)
			}
			t.Log("=====================")
			t.Log("case3 pass")
		}
	}

	// case4: blockChain arrived settle block height, restricting plan exist, debt symbol is false,
	// and total debt and restricting balance is more than release amount
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		blockNumber := uint64(1) * xutil.CalcBlocksEachEpoch()

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(2E18)
		info.Debt = big.NewInt(1E18)
		info.DebtSymbol = false
		info.ReleaseList = []uint64{1, 2}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		for _, epoch := range info.ReleaseList {
			// store epoch
			releaseEpochKey := restricting.GetReleaseEpochKey(epoch)
			stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

			// store release account
			releaseAccountKey := restricting.GetReleaseAccountKey(epoch, uint32(1))
			stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())
		}

		// store release amount
		releaseAmountKey := restricting.GetReleaseAmountKey(1, restrictingAcc)
		stateDb.SetState(restrictingAcc, releaseAmountKey, big.NewInt(2E18).Bytes())
		releaseAmountKey = restricting.GetReleaseAmountKey(2, restrictingAcc)
		stateDb.SetState(restrictingAcc, releaseAmountKey, big.NewInt(1E18).Bytes())

		stateDb.AddBalance(vm.RestrictingContractAddr, big.NewInt(2E18))

		// do EndBlock
		head := types.Header{Number: big.NewInt(int64(blockNumber))}
		err = plugin.RestrictingInstance().EndBlock(common.Hash{}, &head, stateDb)

		t.Log("=====================")
		t.Log("expected case4 of ReturnLockFunds success")
		t.Log("expected balance of restricting account:", big.NewInt(1E18))
		t.Log("expected balance of restricting contract contract:", big.NewInt(1E18))
		t.Log("expected balance of restrict account: ", big.NewInt(1E18))
		t.Log("expected debt    of restrict account: ", big.NewInt(0))
		t.Log("expected symbol  of restrict account: ", false)
		t.Log("expected list    of restrict account: ", []uint64{2})
		t.Log("=====================")
		t.Log("expected [release Epoch record not found, epoch: 1]")
		t.Log("expected [record of restricting account amount not found, account: 0x740cE31B3fAc20Dac379dB243021A51E80AaDd24, epoch: 1]")
		t.Log("expected account numbers of release epoch 2:", 1)
		t.Logf("expected release accounts of epoch 2: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(1E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case4 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case4 return success!")
			t.Log("actually balance of restricting account:", stateDb.GetBalance(restrictingAcc))
			t.Log("actually balance of restricting contract account:", stateDb.GetBalance(vm.RestrictingContractAddr))

			showRestrictingAccountInfo(t, stateDb, restrictingAcc)
			for _, epoch := range info.ReleaseList {
				showReleaseEpoch(t, stateDb, epoch)
				showReleaseAmount(t, stateDb, restrictingAcc, epoch)
			}
			t.Log("=====================")
			t.Log("case4 pass")
		}
	}

	// case5: blockChain arrived settle block height, restricting plan exist, debt symbol is false,
	// and total debt and restricting balance is less than release amount
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		blockNumber := uint64(1) * xutil.CalcBlocksEachEpoch()

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(2E18)
		info.Debt = big.NewInt(1E18)
		info.DebtSymbol = false
		info.ReleaseList = []uint64{1, 2}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		for _, epoch := range info.ReleaseList {
			// store epoch
			releaseEpochKey := restricting.GetReleaseEpochKey(epoch)
			stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

			// store release account
			releaseAccountKey := restricting.GetReleaseAccountKey(epoch, uint32(1))
			stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())
		}

		// store release amount
		releaseAmountKey := restricting.GetReleaseAmountKey(1, restrictingAcc)
		stateDb.SetState(restrictingAcc, releaseAmountKey, big.NewInt(4E18).Bytes())
		releaseAmountKey = restricting.GetReleaseAmountKey(2, restrictingAcc)
		stateDb.SetState(restrictingAcc, releaseAmountKey, big.NewInt(1E18).Bytes())

		stateDb.AddBalance(vm.RestrictingContractAddr, big.NewInt(2E18))

		// do EndBlock
		head := types.Header{Number: big.NewInt(int64(blockNumber))}
		err = plugin.RestrictingInstance().EndBlock(common.Hash{}, &head, stateDb)

		t.Log("=====================")
		t.Log("expected case5 of ReturnLockFunds success")
		t.Log("expected balance of restricting account:", big.NewInt(2E18))
		t.Log("expected balance of restricting contract contract:", big.NewInt(0))
		t.Log("expected balance of restrict account: ", big.NewInt(0))
		t.Log("expected debt    of restrict account: ", big.NewInt(1E18))
		t.Log("expected symbol  of restrict account: ", true)
		t.Log("expected list    of restrict account: ", []uint64{2})
		t.Log("=====================")
		t.Log("expected [release Epoch record not found, epoch: 1]")
		t.Log("expected [record of restricting account amount not found, account: 0x740cE31B3fAc20Dac379dB243021A51E80AaDd24, epoch: 1]")
		t.Log("expected account numbers of release epoch 2:", 1)
		t.Logf("expected release accounts of epoch 2: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(1E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case5 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case5 return success!")
			t.Log("actually balance of restricting account:", stateDb.GetBalance(restrictingAcc))
			t.Log("actually balance of restricting contract account:", stateDb.GetBalance(vm.RestrictingContractAddr))

			showRestrictingAccountInfo(t, stateDb, restrictingAcc)
			for _, epoch := range info.ReleaseList {
				showReleaseEpoch(t, stateDb, epoch)
				showReleaseAmount(t, stateDb, restrictingAcc, epoch)
			}
			t.Log("=====================")
			t.Log("case5 pass")
		}
	}

	// case6: blockChain arrived settle block height, restricting plan exist, debt symbol is true,
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		blockNumber := uint64(2) * xutil.CalcBlocksEachEpoch()
		plugin.SetLatestEpoch(stateDb, 1)

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(0)
		info.Debt = big.NewInt(1E18)
		info.DebtSymbol = true
		info.ReleaseList = []uint64{2, 3}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		for _, epoch := range info.ReleaseList {
			// store epoch
			releaseEpochKey := restricting.GetReleaseEpochKey(epoch)
			stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

			// store release account
			releaseAccountKey := restricting.GetReleaseAccountKey(epoch, uint32(1))
			stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())
		}

		// store release amount
		releaseAmountKey := restricting.GetReleaseAmountKey(2, restrictingAcc)
		stateDb.SetState(restrictingAcc, releaseAmountKey, big.NewInt(2E18).Bytes())
		releaseAmountKey = restricting.GetReleaseAmountKey(3, restrictingAcc)
		stateDb.SetState(restrictingAcc, releaseAmountKey, big.NewInt(1E18).Bytes())

		// do EndBlock
		head := types.Header{Number: big.NewInt(int64(blockNumber))}
		err = plugin.RestrictingInstance().EndBlock(common.Hash{}, &head, stateDb)

		t.Log("=====================")
		t.Log("expected case6 of ReturnLockFunds success")
		t.Log("expected balance of restricting account:", big.NewInt(0))
		t.Log("expected balance of restricting contract contract:", big.NewInt(0))
		t.Log("expected balance of restrict account: ", big.NewInt(0))
		t.Log("expected debt    of restrict account: ", big.NewInt(3E18))
		t.Log("expected symbol  of restrict account: ", true)
		t.Log("expected list    of restrict account: ", []uint64{3})
		t.Log("=====================")
		t.Log("expected [release Epoch record not found, epoch: 2]")
		t.Log("expected [record of restricting account amount not found, account: 0x740cE31B3fAc20Dac379dB243021A51E80AaDd24, epoch: 2]")
		t.Log("expected account numbers of release epoch 3:", 1)
		t.Logf("expected release accounts of epoch 3: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(1E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case6 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case6 return success!")
			t.Log("actually balance of restricting account:", stateDb.GetBalance(restrictingAcc))
			t.Log("actually balance of restricting contract account:", stateDb.GetBalance(vm.RestrictingContractAddr))

			showRestrictingAccountInfo(t, stateDb, restrictingAcc)
			for _, epoch := range info.ReleaseList {
				showReleaseEpoch(t, stateDb, epoch)
				showReleaseAmount(t, stateDb, restrictingAcc, epoch)
			}
			t.Log("=====================")
			t.Log("case6 pass")
		}
	}
}

func TestRestrictingPlugin_AddRestrictingRecord(t *testing.T) {

	var err error
	var plan restricting.RestrictingPlan

	// case1: release epoch is less than latest epoch
	{
		stateDb := buildStateDB(t)
		plugin.SetLatestEpoch(stateDb, 2)

		var plans = make([]restricting.RestrictingPlan, 1)
		plans[0].Epoch = 1
		plans[0].Amount = big.NewInt(int64(1E18))

		err = plugin.RestrictingInstance().AddRestrictingRecord(sender, addrArr[0], plans, stateDb)

		// show expected result
		t.Logf("expected error is [%s]", errParamPeriodInvalid)
		t.Logf("actually error is [%v]", err)

		if err != nil && err.Error() == errParamPeriodInvalid.Error() {
			t.Log("case1 of AddRestrictingRecord pass")
		} else {
			t.Error("case1 of AddRestrictingRecord failed.")
		}
		t.Log("=====================")
		t.Log("case1 pass")
	}

	// case2: balance of sender not enough
	{
		stateDb := buildStateDB(t)
		stateDb.AddBalance(sender, big.NewInt(1))

		var plans = make([]restricting.RestrictingPlan, 1)
		plans[0].Epoch = 1
		plans[0].Amount = big.NewInt(int64(1E18))

		err = plugin.RestrictingInstance().AddRestrictingRecord(sender, addrArr[0], plans, stateDb)

		// show expected result
		t.Logf("expected error is [%s]", errBalanceNotEnough)
		t.Logf("actually error is [%v]", err)

		if err != nil && err.Error() == errBalanceNotEnough.Error() {
			t.Log("case2 of AddRestrictingRecord pass")
		} else {
			t.Error("case2 of AddRestrictingRecord failed.")
		}
		t.Log("=====================")
		t.Log("case2 pass")
	}

	// case3: rlp decode failed
	{
		stateDb := buildStateDB(t)
		stateDb.AddBalance(sender, big.NewInt(1E18))
		restrictingAcc := addrArr[0]

		testData := "this is test data"
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, []byte(testData))

		var plans = make([]restricting.RestrictingPlan, 1)
		plans[0].Epoch = 1
		plans[0].Amount = big.NewInt(int64(1E18))

		err = plugin.RestrictingInstance().AddRestrictingRecord(sender, restrictingAcc, plans, stateDb)

		// show expected result
		t.Logf("expecetd error is [rlp: expected input list for restricting.RestrictingInfo]")
		t.Logf("actually error is [%v]", err)

		if err != nil {
			if _, ok := err.(*common.SysError); ok {
				t.Log("case3 of AddRestricting pass")
			} else {
				t.Error("case3 of AddRestrictingRecord failed.")
			}
		} else {
			t.Error("case3 of AddRestrictingRecord failed.")
		}

		t.Log("=====================")
		t.Log("case3 pass")
	}

	// case4: account is new user to restricting
	{
		stateDb := buildStateDB(t)

		// preset sender balance
		restrictingAmount := big.NewInt(int64(5E18))
		senderBalance := new(big.Int).Add(sender_balance, restrictingAmount)
		stateDb.AddBalance(sender, senderBalance)

		// build input plans for case1
		var plans = make([]restricting.RestrictingPlan, 5)
		for i := 0; i < 5; i++ {
			v := reflect.ValueOf(&plans[i]).Elem()

			epoch := i + 1
			amount := big.NewInt(int64(1E18))
			v.FieldByName("Epoch").SetUint(uint64(epoch))
			v.FieldByName("Amount").Set(reflect.ValueOf(amount))
		}

		// Deduct a portion of the money to contract in advance
		stateDb.SubBalance(sender, restrictingAmount)
		stateDb.AddBalance(vm.RestrictingContractAddr, restrictingAmount)

		err := plugin.RestrictingInstance().AddRestrictingRecord(sender, addrArr[0], plans, stateDb)

		// show expected result
		t.Log("=====================")
		t.Log("expected case4 of AddRestrictingRecord success")
		t.Log("expected balance of sender:", sender_balance)
		t.Log("expected balance of contract:", restrictingAmount)
		t.Log("expected balance of restrict account: ", big.NewInt(int64(5E18)))
		t.Log("expected debt    of restrict account: ", 0)
		t.Log("expected symbol  of restrict account: ", false)
		t.Log("expected list    of restrict account: ", []uint64{1, 2, 3, 4, 5})
		for i := 0; i < 5; i++ {
			epoch := i + 1
			t.Log("=====================")
			t.Logf("expected account numbers of release epoch %d: 1", epoch)
			t.Logf("expected release accounts of epoch %d: %v", epoch, addrArr[0].String())
			t.Logf("expected release amount of account [%s]: %v", addrArr[0].String(), big.NewInt(int64(1E18)))
		}
		t.Log("=====================")

		if err != nil {
			t.Errorf("case4 of AddRestrictingRecord failed. Actually returns error: %s", err.Error())

		} else {

			t.Log("=====================")
			t.Log("case4 return success!")
			t.Log("actually balance of sender:", stateDb.GetBalance(sender))
			t.Log("actually balance of contract:", stateDb.GetBalance(vm.RestrictingContractAddr))
			showRestrictingAccountInfo(t, stateDb, addrArr[0])
			for i := 0; i < 5; i++ {
				epoch := i + 1

				t.Log("=====================")
				showReleaseEpoch(t, stateDb, uint64(epoch))
				showReleaseAmount(t, stateDb, addrArr[0], uint64(epoch))
			}
			t.Log("=====================")
			t.Log("case4 pass")
		}
	}

	// case5: restricting account exist, but restricting epoch not intersect
	{
		stateDb := buildStateDB(t)

		// preset sender balance
		restrictingAmount := big.NewInt(int64(1E18))
		stateDb.AddBalance(sender, restrictingAmount)

		// build db info
		buildDbRestrictingPlan(t, stateDb)

		// build plans for case3
		var plans = make([]restricting.RestrictingPlan, 1)
		plan.Epoch = uint64(6)
		plan.Amount = restrictingAmount
		plans[0] = plan

		// Deduct a portion of the money to contract in advance
		stateDb.SubBalance(sender, restrictingAmount)
		stateDb.AddBalance(vm.RestrictingContractAddr, restrictingAmount)

		err := plugin.RestrictingInstance().AddRestrictingRecord(sender, addrArr[0], plans, stateDb)

		// show expected result
		t.Log("=====================")
		t.Log("expected case5 of AddRestrictingRecord success")
		t.Log("expected balance of sender:", sender_balance)
		t.Log("expected balance of contract:", big.NewInt(int64(6E18)))
		t.Log("expected balance of restrict account: ", big.NewInt(int64(6E18)))
		t.Log("expected debt    of restrict account: ", 0)
		t.Log("expected symbol  of restrict account: ", false)
		t.Log("expected list    of restrict account: ", []uint64{1, 2, 3, 4, 5, 6})
		for i := 0; i < 6; i++ {
			epoch := i + 1
			t.Log("=====================")
			t.Logf("expected account numbers of release epoch %d: 1", epoch)
			t.Logf("expect release accounts of epoch %d: %v", epoch, addrArr[0].String())
			t.Logf("expect release amount of account [%s]: %v", addrArr[0].String(), big.NewInt(int64(1E18)))
		}
		t.Log("=====================")

		if err != nil {
			t.Errorf("case5 of AddRestrictingRecord failed. Actually returns error: %s", err.Error())

		} else {

			t.Log("=====================")
			t.Log("case5 return success!")
			t.Log("actually balance of sender:", stateDb.GetBalance(sender))
			t.Log("actually balance of contract:", stateDb.GetBalance(vm.RestrictingContractAddr))

			showRestrictingAccountInfo(t, stateDb, addrArr[0])
			for i := 0; i < 6; i++ {
				epoch := i + 1

				t.Log("=====================")
				showReleaseEpoch(t, stateDb, uint64(epoch))
				showReleaseAmount(t, stateDb, addrArr[0], uint64(epoch))
			}
			t.Log("=====================")
			t.Log("case5 pass")
		}
	}

	// case6: restricting account exist, and restricting epoch intersect
	{
		stateDb := buildStateDB(t)

		// preset sender balance
		restrictingAmount := big.NewInt(int64(1E18))
		stateDb.AddBalance(sender, restrictingAmount)

		// build db info
		buildDbRestrictingPlan(t, stateDb)

		// build plans for case3
		var plans = make([]restricting.RestrictingPlan, 1)
		plan.Epoch = uint64(5)
		plan.Amount = restrictingAmount
		plans[0] = plan

		// Deduct a portion of the money to contract in advance
		stateDb.SubBalance(sender, restrictingAmount)
		stateDb.AddBalance(vm.RestrictingContractAddr, restrictingAmount)

		err := plugin.RestrictingInstance().AddRestrictingRecord(sender, addrArr[0], plans, stateDb)

		t.Log("=====================")
		t.Log("expected case6 of AddRestrictingRecord success")
		t.Log("expected balance of sender:", sender_balance)
		t.Log("expected balance of contract:", big.NewInt(int64(6E18)))
		t.Log("expected balance of restrict account: ", big.NewInt(int64(6E18)))
		t.Log("expected debt    of restrict account: ", 0)
		t.Log("expected symbol  of restrict account: ", false)
		t.Log("expected list    of restrict account: ", []uint64{1, 2, 3, 4, 5})
		for i := 0; i < 5; i++ {
			epoch := i + 1
			t.Log("=====================")
			t.Logf("expected account numbers of release epoch %d: 1", epoch)
			t.Logf("expect release accounts of epoch %d: %v", epoch, addrArr[0].String())
			if epoch == 5 {
				t.Logf("expect release amount of account [%s]: %v", addrArr[0].String(), big.NewInt(int64(2E18)))
			} else {
				t.Logf("expect release amount of account [%s]: %v", addrArr[0].String(), big.NewInt(int64(1E18)))
			}
		}
		t.Log("=====================")

		if err != nil {
			t.Errorf("case6 of AddRestrictingRecord failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case6 return success!")
			t.Log("actually balance of sender:", stateDb.GetBalance(sender))
			t.Log("actually balance of contract:", stateDb.GetBalance(vm.RestrictingContractAddr))
			showRestrictingAccountInfo(t, stateDb, addrArr[0])
			for i := 0; i < 5; i++ {
				epoch := i + 1
				t.Log("=====================")
				showReleaseEpoch(t, stateDb, uint64(epoch))
				showReleaseAmount(t, stateDb, addrArr[0], uint64(epoch))
			}
			t.Log("=====================")
			t.Log("case6 pass")
		}
	}
}

func TestRestrictingPlugin_PledgeLockFunds(t *testing.T) {

	var err error

	// case1: restricting account not exist
	{
		stateDb := buildStateDB(t)

		lockFunds := big.NewInt(int64(2E18))
		notFoundAccount := common.HexToAddress("0x11")

		err = plugin.RestrictingInstance().PledgeLockFunds(notFoundAccount, lockFunds, stateDb)

		// show expected result
		t.Logf("expected error is [%s]", errAccountNotFound)
		t.Logf("actually error is [%v]", err)

		if err != nil && err.Error() == errAccountNotFound.Error() {
			t.Log("case1 of PledgeLockFunds pass")
		} else {
			t.Error("case1 of PledgeLockFunds failed.")
		}
		t.Log("=====================")
		t.Log("case1 pass")
	}

	// case2: restricting account exist, but Balance not enough
	{
		stateDb := buildStateDB(t)

		// build data in stateDB for case2
		buildDbRestrictingPlan(t, stateDb)
		lockFunds := big.NewInt(int64(6E18))

		err = plugin.RestrictingInstance().PledgeLockFunds(addrArr[0], lockFunds, stateDb)

		// show expected result
		t.Logf("expected error is [%s]", errBalanceNotEnough)
		t.Logf("actually error is [%v]", err)

		if err != nil && err.Error() == errBalanceNotEnough.Error() {
			t.Log("case2 of PledgeLockFunds pass")
			showRestrictingAccountInfo(t, stateDb, addrArr[0])
		} else {
			t.Error("case2 of PledgeLockFunds failed.")
		}

		t.Log("=====================")
		t.Log("case2 pass")
	}

	// case3: rlp decode failed
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		lockFunds := big.NewInt(1)

		testData := "this is test data"
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, []byte(testData))

		err = plugin.RestrictingInstance().PledgeLockFunds(addrArr[0], lockFunds, stateDb)

		// show expected result
		t.Logf("expecetd error is [rlp: expected input list for restricting.RestrictingInfo]")
		t.Logf("actually error is [%v]", err)

		if err != nil {
			if _, ok := err.(*common.SysError); ok {
				t.Log("case3 of PledgeLockFunds pass")
			} else {
				t.Error("case3 of PledgeLockFunds failed.")
			}
		} else {
			t.Error("case3 of PledgeLockFunds failed.")
		}

		t.Log("=====================")
		t.Log("case3 pass")
	}

	// case4: restricting account exist, and Balance is enough
	{
		stateDb := buildStateDB(t)

		// build data in stateDB for case4
		buildDbRestrictingPlan(t, stateDb)

		lockFunds := big.NewInt(int64(2E18))

		err = plugin.RestrictingInstance().PledgeLockFunds(addrArr[0], lockFunds, stateDb)

		// show expected result
		t.Log("=====================")
		t.Log("expected case4 of PledgeLockFunds success")
		t.Log("expected balance of contract:", big.NewInt(int64(3E18)))
		t.Log("expected balance of restrict account: ", big.NewInt(int64(3E18)))
		t.Log("expected debt    of restrict account: ", 0)
		t.Log("expected symbol  of restrict account: ", false)
		t.Log("expected list    of restrict account: ", []uint64{1, 2, 3, 4, 5})
		for i := 0; i < 5; i++ {
			epoch := i + 1
			t.Log("=====================")
			t.Logf("expected account numbers of release epoch %d: 1", epoch)
			t.Logf("expected release accounts of epoch %d: %v", epoch, addrArr[0].String())
			t.Logf("expected release amount of account [%s]: %v", addrArr[0].String(), big.NewInt(int64(1E18)))
		}
		t.Log("=====================")

		if err != nil {
			t.Errorf("case4 of PledgeLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case4 return success!")
			t.Log("actually balance of contract:", stateDb.GetBalance(vm.RestrictingContractAddr))
			showRestrictingAccountInfo(t, stateDb, addrArr[0])
			for i := 0; i < 5; i++ {
				epoch := i + 1
				t.Log("=====================")
				showReleaseEpoch(t, stateDb, uint64(epoch))
				showReleaseAmount(t, stateDb, addrArr[0], uint64(epoch))
			}
			t.Log("=====================")
			t.Log("case4 pass")
		}

		t.Log("=====================")
		t.Log("case4 pass")
	}
}

func TestRestrictingPlugin_ReturnLockFunds(t *testing.T) {

	// case1: restricting account not exist
	{
		stateDb := buildStateDB(t)

		returnFunds := big.NewInt(int64(1E18))
		notFoundAccount := common.HexToAddress("0x11")

		err := plugin.RestrictingInstance().ReturnLockFunds(notFoundAccount, returnFunds, stateDb)

		// show expected result
		t.Logf("expected error is [%s]", errAccountNotFound)
		t.Logf("actually error is [%v]", err)

		if err != nil && err.Error() == errAccountNotFound.Error() {
			t.Log("case1 of ReturnLockFunds pass")
		} else {
			t.Error("case1 of ReturnLockFunds failed.")
		}

		t.Log("=====================")
		t.Log("case1 pass")
	}

	// case2: rlp decode failed
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		lockFunds := big.NewInt(1)

		testData := "this is test data"
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, []byte(testData))

		err := plugin.RestrictingInstance().ReturnLockFunds(addrArr[0], lockFunds, stateDb)

		// show expected result
		t.Logf("expecetd error is [rlp: expected input list for restricting.RestrictingInfo]")
		t.Logf("actually error is [%v]", err)

		if err != nil {
			if _, ok := err.(*common.SysError); ok {
				t.Log("case2 of PledgeLockFunds pass")
			} else {
				t.Error("case2 of PledgeLockFunds failed.")
			}
		} else {
			t.Error("case2 of PledgeLockFunds failed.")
		}

		t.Log("=====================")
		t.Log("case2 pass")
	}

	// case3: restricting account exist, debt symbol is false
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		returnFunds := big.NewInt(1E18)

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(1E18)
		info.Debt = big.NewInt(0)
		info.DebtSymbol = false
		info.ReleaseList = []uint64{5}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		// store epoch
		releaseEpochKey := restricting.GetReleaseEpochKey(uint64(5))
		stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

		// store release account
		releaseAccountKey := restricting.GetReleaseAccountKey(uint64(5), uint32(1))
		stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())

		// store release amount
		releaseAmountKey := restricting.GetReleaseAmountKey(uint64(5), restrictingAcc)
		amount := big.NewInt(2E18)
		stateDb.SetState(restrictingAcc, releaseAmountKey, amount.Bytes())

		stateDb.AddBalance(vm.StakingContractAddr, big.NewInt(1E18))
		stateDb.AddBalance(vm.RestrictingContractAddr, big.NewInt(1E18))

		// do ReturnLockFunds
		err = plugin.RestrictingInstance().ReturnLockFunds(addrArr[0], returnFunds, stateDb)

		// show expected result
		t.Log("=====================")
		t.Log("expected case3 of ReturnLockFunds success")
		t.Log("expected balance of restricting contract:", big.NewInt(int64(2E18)))
		t.Log("expected balance of staking contract:", big.NewInt(0))
		t.Log("expected balance of restrict account: ", big.NewInt(int64(2E18)))
		t.Log("expected debt    of restrict account: ", 0)
		t.Log("expected symbol  of restrict account: ", false)
		t.Log("expected list    of restrict account: ", []uint64{5})
		t.Log("=====================")
		t.Log("expected account numbers of release epoch 5: 1")
		t.Logf("expected release accounts of epoch 5: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(2E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case3 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case3 return success!")
			t.Log("actually balance of restricting contract:", stateDb.GetBalance(vm.RestrictingContractAddr))
			t.Log("actually balance of staking contract:", stateDb.GetBalance(vm.StakingContractAddr))
			showRestrictingAccountInfo(t, stateDb, restrictingAcc)

			showReleaseEpoch(t, stateDb, uint64(5))
			showReleaseAmount(t, stateDb, restrictingAcc, uint64(5))
			t.Log("=====================")
			t.Log("case3 pass")
		}
	}

	// case4: restricting account exist, and debt symbol is true, and amount is less than debt
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		returnFunds := big.NewInt(1E18)

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(0)
		info.Debt = big.NewInt(2E18)
		info.DebtSymbol = true
		info.ReleaseList = []uint64{5}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		// store epoch
		releaseEpochKey := restricting.GetReleaseEpochKey(uint64(5))
		stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

		// store release account
		releaseAccountKey := restricting.GetReleaseAccountKey(uint64(5), uint32(1))
		stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())

		// store release amount
		releaseAmountKey := restricting.GetReleaseAmountKey(uint64(5), restrictingAcc)
		amount := big.NewInt(3E18)
		stateDb.SetState(restrictingAcc, releaseAmountKey, amount.Bytes())

		stateDb.AddBalance(vm.StakingContractAddr, big.NewInt(1E18))

		// do ReturnLockFunds
		err = plugin.RestrictingInstance().ReturnLockFunds(addrArr[0], returnFunds, stateDb)

		// show expected result
		t.Log("=====================")
		t.Log("expected case4 of ReturnLockFunds success")
		t.Log("expected balance of restricting account:", big.NewInt(int64(1E18)))
		t.Log("expected balance of staking contract:", big.NewInt(0))
		t.Log("expected balance of restrict account: ", big.NewInt(0))
		t.Log("expected debt    of restrict account: ", big.NewInt(int64(1E18)))
		t.Log("expected symbol  of restrict account: ", true)
		t.Log("expected list    of restrict account: ", []uint64{5})
		t.Log("=====================")
		t.Log("expected account numbers of release epoch 5: 1")
		t.Logf("expected release accounts of epoch 5: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(3E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case4 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case4 return success!")
			t.Log("actually balance of restricting account:", stateDb.GetBalance(restrictingAcc))
			t.Log("actually balance of staking contract:", stateDb.GetBalance(vm.StakingContractAddr))
			showRestrictingAccountInfo(t, stateDb, restrictingAcc)

			showReleaseEpoch(t, stateDb, uint64(5))
			showReleaseAmount(t, stateDb, restrictingAcc, uint64(5))
			t.Log("=====================")
			t.Log("case4 pass")
		}
	}

	// case5: restricting account exist, and debt symbol is true, and amount is more than debt
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		returnFunds := big.NewInt(3E18)

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(0)
		info.Debt = big.NewInt(2E18)
		info.DebtSymbol = true
		info.ReleaseList = []uint64{5}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		// store epoch
		releaseEpochKey := restricting.GetReleaseEpochKey(uint64(5))
		stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

		// store release account
		releaseAccountKey := restricting.GetReleaseAccountKey(uint64(5), uint32(1))
		stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())

		// store release amount
		releaseAmountKey := restricting.GetReleaseAmountKey(uint64(5), restrictingAcc)
		amount := big.NewInt(3E18)
		stateDb.SetState(restrictingAcc, releaseAmountKey, amount.Bytes())

		stateDb.AddBalance(vm.StakingContractAddr, big.NewInt(3E18))

		// do ReturnLockFunds
		err = plugin.RestrictingInstance().ReturnLockFunds(addrArr[0], returnFunds, stateDb)

		// show expected result
		t.Log("=====================")
		t.Log("expected case5 of ReturnLockFunds success")
		t.Log("expected balance of restricting account:", big.NewInt(int64(2E18)))
		t.Log("expected balance of staking contract:", big.NewInt(0))
		t.Log("expected balance of restrict account: ", big.NewInt(1E18))
		t.Log("expected debt    of restrict account: ", big.NewInt(int64(0)))
		t.Log("expected symbol  of restrict account: ", false)
		t.Log("expected list    of restrict account: ", []uint64{5})
		t.Log("=====================")
		t.Log("expected account numbers of release epoch 5: 1")
		t.Logf("expected release accounts of epoch 5: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(3E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case5 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case5 return success!")
			t.Log("actually balance of restricting account:", stateDb.GetBalance(restrictingAcc))
			t.Log("actually balance of staking contract:", stateDb.GetBalance(vm.StakingContractAddr))
			showRestrictingAccountInfo(t, stateDb, restrictingAcc)

			showReleaseEpoch(t, stateDb, uint64(5))
			showReleaseAmount(t, stateDb, restrictingAcc, uint64(5))
			t.Log("=====================")
			t.Log("case5 pass")
		}
	}

	// case6: restricting account exist, and debt symbol is true, and amount is more than debt
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		returnFunds := big.NewInt(3E18)

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(0)
		info.Debt = big.NewInt(2E18)
		info.DebtSymbol = true
		info.ReleaseList = []uint64{5}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		// store epoch
		releaseEpochKey := restricting.GetReleaseEpochKey(uint64(5))
		stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

		// store release account
		releaseAccountKey := restricting.GetReleaseAccountKey(uint64(5), uint32(1))
		stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())

		// store release amount
		releaseAmountKey := restricting.GetReleaseAmountKey(uint64(5), restrictingAcc)
		amount := big.NewInt(3E18)
		stateDb.SetState(restrictingAcc, releaseAmountKey, amount.Bytes())

		stateDb.AddBalance(vm.StakingContractAddr, big.NewInt(3E18))

		// do ReturnLockFunds
		err = plugin.RestrictingInstance().ReturnLockFunds(addrArr[0], returnFunds, stateDb)

		// show expected result
		t.Log("=====================")
		t.Log("expected case5 of ReturnLockFunds success")
		t.Log("expected balance of restricting account:", big.NewInt(int64(2E18)))
		t.Log("expected balance of staking contract:", big.NewInt(0))
		t.Log("expected balance of restrict account: ", big.NewInt(1E18))
		t.Log("expected debt    of restrict account: ", big.NewInt(int64(0)))
		t.Log("expected symbol  of restrict account: ", false)
		t.Log("expected list    of restrict account: ", []uint64{5})
		t.Log("=====================")
		t.Log("expected account numbers of release epoch 5: 1")
		t.Logf("expected release accounts of epoch 5: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(3E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case5 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case5 return success!")
			t.Log("actually balance of restricting account:", stateDb.GetBalance(restrictingAcc))
			t.Log("actually balance of staking contract:", stateDb.GetBalance(vm.StakingContractAddr))
			showRestrictingAccountInfo(t, stateDb, restrictingAcc)

			showReleaseEpoch(t, stateDb, uint64(5))
			showReleaseAmount(t, stateDb, restrictingAcc, uint64(5))
			t.Log("=====================")
			t.Log("case5 pass")
		}

		t.Fatal("+++++++")
	}
}

func TestRestrictingPlugin_SlashingNotify(t *testing.T) {

	// case1: restricting account not exist
	{
		stateDb := buildStateDB(t)

		slashingFunds := big.NewInt(int64(1E18))
		notFoundAccount := common.HexToAddress("0x11")

		err := plugin.RestrictingInstance().SlashingNotify(notFoundAccount, slashingFunds, stateDb)

		// show expected result
		t.Logf("expected error is [%s]", errAccountNotFound)
		t.Logf("actually error is [%v]", err)

		if err != nil && err.Error() == errAccountNotFound.Error() {
			t.Log("case1 of SlashingNotify pass")
		} else {
			t.Error("case1 of SlashingNotify failed.")
		}

		t.Log("=====================")
		t.Log("case1 pass")
	}

	// case2: rlp decode failed
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		lockFunds := big.NewInt(1)

		testData := "this is test data"
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, []byte(testData))

		err := plugin.RestrictingInstance().SlashingNotify(addrArr[0], lockFunds, stateDb)

		// show expected result
		t.Logf("expecetd error is [rlp: expected input list for restricting.RestrictingInfo]")
		t.Logf("actually error is [%v]", err)

		if err != nil {
			if _, ok := err.(*common.SysError); ok {
				t.Log("case2 of SlashingNotify pass")
			} else {
				t.Error("case2 of SlashingNotify failed.")
			}
		} else {
			t.Error("case2 of SlashingNotify failed.")
		}

		t.Log("=====================")
		t.Log("case2 pass")
	}

	// case3: restricting account exist, and debt symbol is false
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		slashFunds := big.NewInt(1E18)

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(1E18)
		info.Debt = big.NewInt(2E18)
		info.DebtSymbol = false
		info.ReleaseList = []uint64{5}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		// store epoch
		releaseEpochKey := restricting.GetReleaseEpochKey(uint64(5))
		stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

		// store release account
		releaseAccountKey := restricting.GetReleaseAccountKey(uint64(5), uint32(1))
		stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())

		// store release amount
		releaseAmountKey := restricting.GetReleaseAmountKey(uint64(5), restrictingAcc)
		amount := big.NewInt(3E18)
		stateDb.SetState(restrictingAcc, releaseAmountKey, amount.Bytes())

		stateDb.AddBalance(vm.StakingContractAddr, big.NewInt(1E18))

		// do ReturnLockFunds
		err = plugin.RestrictingInstance().SlashingNotify(addrArr[0], slashFunds, stateDb)

		// show expected result
		t.Log("=====================")
		t.Log("expected case3 of ReturnLockFunds success")
		t.Log("expected balance of restricting account:", big.NewInt(0))
		t.Log("expected balance of restricting contract account:", big.NewInt(0))
		t.Log("expected balance of staking contract:", big.NewInt(1E18))
		t.Log("expected balance of slashing contract contract:", big.NewInt(0))
		t.Log("expected balance of restrict account: ", big.NewInt(1E18))
		t.Log("expected debt    of restrict account: ", big.NewInt(3E18))
		t.Log("expected symbol  of restrict account: ", false)
		t.Log("expected list    of restrict account: ", []uint64{5})
		t.Log("=====================")
		t.Log("expected account numbers of release epoch 5: 1")
		t.Logf("expected release accounts of epoch 5: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(3E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case3 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case3 return success!")
			t.Log("actually balance of restricting account:", stateDb.GetBalance(restrictingAcc))
			t.Log("expected balance of restricting contract account:", stateDb.GetBalance(vm.RestrictingContractAddr))
			t.Log("expected balance of staking contract:", stateDb.GetBalance(vm.StakingContractAddr))
			t.Log("expected balance of slashing contract contract:", stateDb.GetBalance(vm.SlashingContractAddr))
			showRestrictingAccountInfo(t, stateDb, restrictingAcc)
			showReleaseEpoch(t, stateDb, uint64(5))
			showReleaseAmount(t, stateDb, restrictingAcc, uint64(5))
			t.Log("=====================")
			t.Log("case3 pass")
		}
	}

	// case4: restricting account exist, and debt symbol is true, and amount is less than debt
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		slashFunds := big.NewInt(1E18)

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(1E18)
		info.Debt = big.NewInt(2E18)
		info.DebtSymbol = true
		info.ReleaseList = []uint64{5}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		// store epoch
		releaseEpochKey := restricting.GetReleaseEpochKey(uint64(5))
		stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

		// store release account
		releaseAccountKey := restricting.GetReleaseAccountKey(uint64(5), uint32(1))
		stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())

		// store release amount
		releaseAmountKey := restricting.GetReleaseAmountKey(uint64(5), restrictingAcc)
		amount := big.NewInt(3E18)
		stateDb.SetState(restrictingAcc, releaseAmountKey, amount.Bytes())

		stateDb.AddBalance(vm.StakingContractAddr, big.NewInt(1E18))

		// do ReturnLockFunds
		err = plugin.RestrictingInstance().SlashingNotify(addrArr[0], slashFunds, stateDb)

		// show expected result
		t.Log("=====================")
		t.Log("expected case4 of ReturnLockFunds success")
		t.Log("expected balance of restricting account:", big.NewInt(0))
		t.Log("expected balance of restricting contract account:", big.NewInt(0))
		t.Log("expected balance of staking contract:", big.NewInt(1E18))
		t.Log("expected balance of slashing contract contract:", big.NewInt(0))
		t.Log("expected balance of restrict account: ", big.NewInt(1E18))
		t.Log("expected debt    of restrict account: ", big.NewInt(1E18))
		t.Log("expected symbol  of restrict account: ", true)
		t.Log("expected list    of restrict account: ", []uint64{5})
		t.Log("=====================")
		t.Log("expected account numbers of release epoch 5: 1")
		t.Logf("expected release accounts of epoch 5: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(3E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case4 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case4 return success!")
			t.Log("actually balance of restricting account:", stateDb.GetBalance(restrictingAcc))
			t.Log("expected balance of restricting contract account:", stateDb.GetBalance(vm.RestrictingContractAddr))
			t.Log("expected balance of staking contract:", stateDb.GetBalance(vm.StakingContractAddr))
			t.Log("expected balance of slashing contract contract:", stateDb.GetBalance(vm.SlashingContractAddr))
			showRestrictingAccountInfo(t, stateDb, restrictingAcc)
			showReleaseEpoch(t, stateDb, uint64(5))
			showReleaseAmount(t, stateDb, restrictingAcc, uint64(5))
			t.Log("=====================")
			t.Log("case4 pass")
		}
	}

	// case5: restricting account exist, and debt symbol is true, and amount is more than debt
	{
		stateDb := buildStateDB(t)
		restrictingAcc := addrArr[0]
		slashFunds := big.NewInt(3E18)

		var info restricting.RestrictingInfo
		info.Balance = big.NewInt(1E18)
		info.Debt = big.NewInt(2E18)
		info.DebtSymbol = true
		info.ReleaseList = []uint64{5}

		bInfo, err := rlp.EncodeToBytes(info)
		if err != nil {
			t.Fatal("rlp encode test data failed")
		}

		// store restricting info
		restrictingKey := restricting.GetRestrictingKey(restrictingAcc)
		stateDb.SetState(restrictingAcc, restrictingKey, bInfo)

		// store epoch
		releaseEpochKey := restricting.GetReleaseEpochKey(uint64(5))
		stateDb.SetState(vm.RestrictingContractAddr, releaseEpochKey, common.Uint32ToBytes(1))

		// store release account
		releaseAccountKey := restricting.GetReleaseAccountKey(uint64(5), uint32(1))
		stateDb.SetState(vm.RestrictingContractAddr, releaseAccountKey, restrictingAcc.Bytes())

		// store release amount
		releaseAmountKey := restricting.GetReleaseAmountKey(uint64(5), restrictingAcc)
		amount := big.NewInt(3E18)
		stateDb.SetState(restrictingAcc, releaseAmountKey, amount.Bytes())

		stateDb.AddBalance(vm.StakingContractAddr, big.NewInt(1E18))

		// do ReturnLockFunds
		err = plugin.RestrictingInstance().SlashingNotify(addrArr[0], slashFunds, stateDb)

		// show expected result
		t.Log("=====================")
		t.Log("expected case5 of ReturnLockFunds success")
		t.Log("expected balance of restricting account:", big.NewInt(0))
		t.Log("expected balance of restricting contract account:", big.NewInt(0))
		t.Log("expected balance of staking contract:", big.NewInt(1E18))
		t.Log("expected balance of slashing contract contract:", big.NewInt(0))
		t.Log("expected balance of restrict account: ", big.NewInt(1E18))
		t.Log("expected debt    of restrict account: ", big.NewInt(1E18))
		t.Log("expected symbol  of restrict account: ", true)
		t.Log("expected list    of restrict account: ", []uint64{5})
		t.Log("=====================")
		t.Log("expected account numbers of release epoch 5: 1")
		t.Logf("expected release accounts of epoch 5: %s", restrictingAcc.String())
		t.Logf("expected release amount of account [%s]: %v", restrictingAcc.String(), big.NewInt(3E18))
		t.Log("=====================")

		if err != nil {
			t.Errorf("case5 of ReturnLockFunds failed. Actually returns error: %s", err.Error())
		} else {
			t.Log("=====================")
			t.Log("case5 return success!")
			t.Log("actually balance of restricting account:", stateDb.GetBalance(restrictingAcc))
			t.Log("expected balance of restricting contract account:", stateDb.GetBalance(vm.RestrictingContractAddr))
			t.Log("expected balance of staking contract:", stateDb.GetBalance(vm.StakingContractAddr))
			t.Log("expected balance of slashing contract contract:", stateDb.GetBalance(vm.SlashingContractAddr))
			showRestrictingAccountInfo(t, stateDb, restrictingAcc)
			showReleaseEpoch(t, stateDb, uint64(5))
			showReleaseAmount(t, stateDb, restrictingAcc, uint64(5))
			t.Log("=====================")
			t.Log("case5 pass")
		}
	}
}

func TestRestrictingPlugin_GetRestrictingInfo(t *testing.T) {

	// case1: restricting account not exist
	{
		stateDb := buildStateDB(t)

		notFoundAccount := common.HexToAddress("0x11")
		_, err := plugin.RestrictingInstance().GetRestrictingInfo(notFoundAccount, stateDb)

		// show expected result
		t.Logf("expected error is [%s]", errAccountNotFound)
		t.Logf("actually error is [%v]", err)

		if err != nil && err.Error() == errAccountNotFound.Error() {
			t.Log("case1 of GetRestrictingInfo pass")
		} else {
			t.Error("case1 of GetRestrictingInfo failed.")
		}
	}

	// case2: restricting account exist
	{
		stateDb := buildStateDB(t)

		buildDbRestrictingPlan(t, stateDb)

		result, err := plugin.RestrictingInstance().GetRestrictingInfo(addrArr[0], stateDb)

		t.Log("=====================")
		t.Log("expected case2 of GetRestrictingInfo success")
		t.Log("expected balance of restrict account: ", big.NewInt(int64(5E18)))
		t.Log("expected slash   of restrict account: ", big.NewInt(0))
		t.Log("expected debt    of restrict account: ", big.NewInt(0))
		t.Log("expected staking of restrict account: ", big.NewInt(0))
		for i := 0; i < 5; i++ {
			expectedBlocks := uint64(i+1) * xutil.CalcBlocksEachEpoch()
			t.Logf("expected release amount at blockNumber [%d] is: %v", expectedBlocks, big.NewInt(int64(1E18)))
		}
		t.Log("=====================")

		if err != nil {
			t.Errorf("case2 of GetRestrictingInfo failed. Actually returns error: %s", err.Error())
		} else {

			if len(result) == 0 {
				t.Log("case2 of GetRestrictingInfo failed. Actually result is empty")
			}

			var res restricting.Result
			if err = rlp.Decode(bytes.NewBuffer(result), &res); err != nil {
				t.Fatalf("failed to elp decode result, result: %s", result)
			}

			t.Log("actually balance of restrict account: ", res.Balance)
			t.Log("actually debt    of restrict account: ", res.Debt)
			t.Log("actually slash   of restrict account: ", res.Slash)
			t.Log("actually staking of restrict account: ", res.Staking)

			var infos []restricting.ReleaseAmountInfo
			if err = json.Unmarshal(res.Entry, &infos); err != nil {
				t.Fatalf("unmarshal release amout info failed, err:%s", err.Error())
			}

			for _, info := range infos {
				t.Logf("actually release amount at blockNumber [%d] is: %v", info.Height, info.Amount)
			}
		}
	}
}
