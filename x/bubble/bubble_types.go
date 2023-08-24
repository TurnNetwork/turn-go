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

// BubState bubble chain status
type BubState uint8

const (
	//BuildingStatus   = 0
	ActiveStatus     BubState = 1
	PreReleaseStatus          = 2
	ReleasedStatus            = 3
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

// BubBasics bubble's underlying information type
// It stores the basic information when the bubble network is created,
// and the data will not be updated and changed during the operation of the bubble network
type BubBasics struct {
	BubbleId    *big.Int
	Creator     common.Address
	CreateBlock uint64
	OperatorsL1 []*Operator
	OperatorsL2 []*Operator
	MicroNodes  CandidateQueue
}

type Bubble struct {
	Basics                  *BubBasics // Bubble Basics
	State                   BubState
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
	NodeId      discover.NodeID `json:"nodeId"` // Operator node id
	ElectronRPC string          `json:"ElectronRPC"`
	RPC         string          `json:"rpc"`     // Operation node RPC URL
	OpAddr      common.Address  `json:"opAddr"`  // Address of operation
	Balance     *big.Int        `json:"balance"` // Operating address balance
}

// OpConfig is operator profiles, including main-chain operator
// and sub-chain operator profiles, interact between main-chain and sub-chain through the operator
type OpConfig struct {
	MainChain *Operator `json:"mainChain,omitempty"` // Main chain operator information configuration
	SubChain  *Operator `json:"subChain,omitempty"`  // Child chain operator information configuration
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
	TxHash   common.Hash    // The transaction hash of the staking Token transaction
	RPC      string         // Bubble The bubble sub-chain operates the node rpc
	OpAddr   common.Address // Bubble The bubble main-chain operates address
	AccAsset *AccountAsset
}

type CreateBubbleTask struct {
	BubbleID *big.Int
	TxHash   common.Hash // The transaction hash of the createBubbleTask
}

type ReleaseBubbleTask struct {
	BubbleID *big.Int
	TxHash   common.Hash // The transaction hash of the releaseBubbleTask
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
