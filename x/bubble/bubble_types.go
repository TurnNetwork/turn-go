package bubble

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/rlp"
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

// BubTxType Define the Bubble's trade type
type BubTxType uint

const (
	StakingToken  BubTxType = iota // StakingToken transaction type
	WithdrewToken                  // WithdrewToken transaction type
	SettleBubble                   // SettleBubble transaction type
)

type CandidateQueue []*stakingL2.Candidate
type ValidatorQueue []*stakingL2.Candidate
type CommitteeQueue []*stakingL2.Candidate

type Bubble struct {
	BubbleId                *big.Int
	Creator                 common.Address
	CreateBlock             uint64
	State                   int // unused
	Member                  SettlementInfo
	OperatorsL1             []*Operator
	OperatorsL2             []*Operator
	MicroNodes              CandidateQueue
	StakingTokenTxHashList  []common.Hash // List of stakingToken transaction hashes
	WithdrewTokenTxHashList []common.Hash // List of withdrewToken transaction hashes
	SettleBubbleTxHashList  []common.Hash // List of settleBubble transaction hashes
}

const (
	OperatorNode = 1
)

// Operator Includes the operator's node ID, rpc url, operation address,
// and initial balance (to send the transaction in the child-chain to pay fees).
type Operator struct {
	NodeId  discover.NodeID `json:"nodeId"`  // Operator node id
	RPC     string          `json:"rpc"`     // Operation node RPC URL
	OpAddr  common.Address  `json:"opAddr"`  // Address of operation
	Balance *big.Int        `json:"balance"` // Operating address balance
}

// OptConfig is operator profiles, including main-chain operator
// and sub-chain operator profiles, interact between main-chain and sub-chain through the operator
type OptConfig struct {
	// Private key of sub-chain operation address (pledged address of operation node)
	subOpPriKey string
	MainChain   *Operator `json:"mainChain,omitempty"` // Main chain operator information configuration
	SubChain    *Operator `json:"subChain,omitempty"`  // Child chain operator information configuration
}

type CreateBubbleTask struct {
	BubbleID *big.Int
	RPCs     string // Bubble The bubble sub-chain operates the node rpc
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
