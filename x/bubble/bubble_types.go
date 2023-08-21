package bubble

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/rlp"
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
	ReleasedStatus   = 3
)

type CandidateQueue []*stakingL2.Candidate
type ValidatorQueue []*stakingL2.Candidate
type CommitteeQueue []*stakingL2.Candidate

type Bubble struct {
	BubbleId    *big.Int
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
	TokenAddr common.Address // ERC20 Token contract address
	Balance   *big.Int       // Token balance
}

type AccountAsset struct {
	Account      common.Address  // Account address
	NativeAmount *big.Int        // Native token balances
	TokenAssets  []AccTokenAsset // Token assets
}

type MintTokenTask struct {
	BubbleID *big.Int
	TxHash   common.Hash // The transaction hash of the staking Token transaction
	RPC      string      // Bubble The bubble sub-chain operates the node rpc
	AccAsset *AccountAsset
}

type SettlementInfo struct {
	AccAssets []AccountAsset // Keep asset information for all accounts
}

func (s SettlementInfo) Hash() (common.Hash, error) {
	enVal, err := rlp.EncodeToBytes(s)
	if err != nil {
		return common.ZeroHash, err
	}
	return crypto.Keccak256Hash(enVal), nil
}
