package bubble

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/x/staking"
	"github.com/bubblenet/bubble/x/stakingL2"
	"math/big"
)

// bubble chain size
const (
	CommitteeSize  = 3
	OperatorL1Size = 1
	OperatorL2Size = 1
)

// bubble chain status
const (
	//BuildingStatus   = 0
	ActiveStatus     = 1
	PreReleaseStatus = 2
	//ReleasedStatus   = 3
)

type CandidateQueue []*stakingL2.Candidate
type ValidatorQueue []*stakingL2.Candidate
type CommitteeQueue []*stakingL2.Candidate

type Bubble struct {
	BubbleId    uint32
	Creator     common.Address
	State       int // unused
	InitBlock   uint64
	SettleBlock uint64 // unused
	Member      SettlementInfo
	OperatorsL1 []*staking.Operator
	OperatorsL2 []*stakingL2.Candidate
	Committees  CandidateQueue
}

type AccTokenAsset struct {
	TokenAddr common.Address // ERC20 Token合约地址
	Balance   *big.Int       // Token余额
}

type AccountAsset struct {
	Account      common.Address  // 账户地址
	NativeAmount *big.Int        // 原生代币余额
	TokenAssets  []AccTokenAsset // Token资产
}

type SettlementInfo struct {
	AccAssets []AccountAsset // 所有账户的资产信息
}
