package xutil

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/x/xcom"
)

var SecondsPerYear = uint64(365 * 24 * 3600)

func NodeId2Addr(nodeId discover.NodeID) (common.Address, error) {
	if pk, err := nodeId.Pubkey(); nil != err {
		return common.ZeroAddr, err
	} else {
		return crypto.PubkeyToAddress(*pk), nil
	}
}

// The ProcessVersion: Major.Minor.Patch eg. 1.1.0
// Calculate the LargeVersion
// eg: 1.1.0 ==> 1.1
func CalcVersion(processVersion uint32) uint32 {
	return processVersion >> 8
}

func IsWorker(extra []byte) bool {
	return len(extra[32:]) >= common.ExtraSeal && bytes.Equal(extra[32:97], make([]byte, common.ExtraSeal))
}

func CheckStakeThreshold(stake *big.Int) bool {
	return stake.Cmp(xcom.StakeThreshold()) >= 0
}

func CheckDelegateThreshold(delegate *big.Int) bool {
	return delegate.Cmp(xcom.DelegateThreshold()) >= 0
}

// eg. 65536 => 1.0.0
func ProcessVersion2Str(processVersion uint32) string {
	major := processVersion << 8
	major = major >> 24

	minor := processVersion << 16
	minor = minor >> 24

	patch := processVersion << 24
	patch = patch >> 24

	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

// TODO: calculate common data of block height
// Number of blocks per consensus round
func ConsensusSize() uint64 {
	return xcom.BlocksWillCreate() * xcom.ConsValidatorNum()
}

// Each epoch (billing cycle) is a multiple of the consensus rounds
func EpochSize() uint64 {
	consensusSize := ConsensusSize()
	em := xcom.ExpectedMinutes()
	i := xcom.Interval()

	epochSize := em * 60 / (i * consensusSize) * consensusSize
	return epochSize
}

// epochs numbers each year
func EpochsPerYear() uint64 {
	epochSize := EpochSize()
	i := xcom.Interval()

	epochs := SecondsPerYear / (i * epochSize)
	return epochs
}

func CalculateBlocksEachEpoch() uint64 {
	return ConsensusSize() * EpochSize()
}

// calculate
func CalculateBlocksEachYear() uint64 {
	return EpochsPerYear() * EpochSize()
}

func IsElection(blockNumber uint64) bool {
	tmp := blockNumber + xcom.ElectionDistance()
	mod := tmp % ConsensusSize()
	return mod == 0
}

func IsSwitch(blockNumber uint64) bool {
	mod := blockNumber % ConsensusSize()
	return mod == 0
}

func IsSettlementPeriod(blockNumber uint64) bool {
	size := CalculateBlocksEachEpoch()
	mod := blockNumber % uint64(size)
	return mod == 0
}

func IsYearEnd(blockNumber uint64) bool {
	size := CalculateBlocksEachYear()
	return blockNumber > 0 && blockNumber%size == 0
}

// calculate the Epoch number by blockNumber
func CalculateEpoch(blockNumber uint64) uint64 {
	size := CalculateBlocksEachEpoch()

	var epoch uint64
	div := blockNumber / size
	mod := blockNumber % size

	switch {
	// first epoch
	case (div == 0 && mod == 0) || (div == 0 && mod > 0):
		epoch = 1
	case div > 0 && mod == 0:
		epoch = div
	case div > 0 && mod > 0:
		epoch = div + 1
	}

	return epoch
}

// calculate the Consensus number by blockNumber
func CalculateRound(blockNumber uint64) uint64 {
	size := ConsensusSize()

	var round uint64
	div := blockNumber / size
	mod := blockNumber % size
	switch {
	// first consensus round
	case (div == 0 && mod == 0) || (div == 0 && mod > 0):
		round = 1
	case div > 0 && mod == 0:
		round = div
	case div > 0 && mod > 0:
		round = div + 1
	}

	return round
}

// calculate the year by blockNumber.
// (V.0.1) If blockNumber eqs 0, year eqs 0 too, else rounded up the result of
// the blockNumber divided by the expected number of blocks per year
func CalculateYear(blockNumber uint64) uint64 {
	size := CalculateBlocksEachYear()

	div := blockNumber / uint64(size)
	mod := blockNumber % uint64(size)

	if mod == 0 {
		return div
	} else {
		return div + 1
	}
}

// TODO: calculate governed configure data for main net only
func MaxVotingDuration() uint64 {
	size := ConsensusSize()
	return uint64(14*24*60*60) / uint64(size) * uint64(size)
}

// TODO: calculate reward configure data for main net only
// SecondYearAllowance is 1.5% of GenesisIssuance
func SecondYearAllowance() *big.Int {
	issue := xcom.GenesisIssuance()
	allowance := new(big.Int).Mul(issue, big.NewInt(15))
	return allowance.Div(allowance, big.NewInt(100))
}

// SecondYearAllowance is 0.5% of GenesisIssuance
func ThirdYearAllowance() *big.Int {
	issue := xcom.GenesisIssuance()
	allowance := new(big.Int).Mul(issue, big.NewInt(5))
	return allowance.Div(allowance, big.NewInt(100))
}

// TODO: calculate restricting configure data for main net only
// GenesisRestrictingBalance is allowance at second year and the third year
func GenesisRestrictingBalance() *big.Int {
	return new(big.Int).Add(SecondYearAllowance(), ThirdYearAllowance())
}
