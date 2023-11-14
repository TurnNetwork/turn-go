package bubble

import (
	"errors"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/rlp"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/stakingL2"
)

// Size is the bubble Size
type Size uint8

// Config contains configuration parameters for bubble
type Config struct {
	OperatorL1Size   uint
	OperatorL2Size   uint
	CommitteeSize    uint
	MinStakingAmount *big.Int
}

var (
	microBubbleConfig  = &Config{OperatorL1Size: 1, OperatorL2Size: 1, CommitteeSize: 3, MinStakingAmount: new(big.Int).Mul(big.NewInt(params.BUB), big.NewInt(5))}
	smallBubbleConfig  = &Config{OperatorL1Size: 1, OperatorL2Size: 1, CommitteeSize: 6, MinStakingAmount: new(big.Int).Mul(big.NewInt(params.BUB), big.NewInt(10))}
	mediumBubbleConfig = &Config{OperatorL1Size: 1, OperatorL2Size: 1, CommitteeSize: 12, MinStakingAmount: new(big.Int).Mul(big.NewInt(params.BUB), big.NewInt(20))}
	largeBubbleConfig  = &Config{OperatorL1Size: 1, OperatorL2Size: 1, CommitteeSize: 24, MinStakingAmount: new(big.Int).Mul(big.NewInt(params.BUB), big.NewInt(40))}
)

var configsTable = map[Size]*Config{
	1: microBubbleConfig,
	2: smallBubbleConfig,
	3: mediumBubbleConfig,
	4: largeBubbleConfig,
}

func GetConfig(size Size) (*Config, error) {
	config := configsTable[size]
	if config == nil {
		return nil, errors.New("unrecognized bubble Size")
	}

	return config, nil
}

// State is the bubble state
type State uint8

const (
	BuildingState State = iota // BuildingStatus is not in use
	ActiveState
	PreReleaseState
	ReleasedState
)

// TxType Define the Bubble's trade type
type TxType uint

const (
	CreateBubble  TxType = iota // createBubble transaction type
	ReleaseBubble               // releaseBubble transaction type
	StakingToken                // StakingToken transaction type
	WithdrewToken               // WithdrewToken transaction type
	SettleBubble                // SettleBubble transaction type
)

type ValidatorQueue []*stakingL2.Candidate
type CommitteeQueue []*stakingL2.Candidate

// BasicsInfo stores the basic information when the bubble network is created,
// and the data will not be chang during the lifetime of the bubble network
type BasicsInfo struct {
	BubbleId    *big.Int
	Size        Size
	OperatorsL1 []*Operator
	OperatorsL2 []*Operator
	MicroNodes  ValidatorQueue
}

type StateInfo struct {
	BubbleId        *big.Int
	State           State
	CreateBlock     uint64
	PreReleaseBlock uint64
	ReleaseBlock    uint64
	ContractCount   uint
}

type TransactionInfo struct {
	StakingTokenTxHashList  []common.Hash // List of stakingToken transaction hashes
	WithdrewTokenTxHashList []common.Hash // List of withdrewToken transaction hashes
	SettleBubbleTxHashList  []common.Hash // List of settleBubble transaction hashes
}

type Bubble struct {
	*BasicsInfo // Bubble Basics
	*StateInfo
	*TransactionInfo
}

const (
	OperatorNode    = 1
	MaxNodeUseRatio = 0.75
)

type AccTokenAsset struct {
	TokenAddr common.Address // ERC20 Token contract address
	Balance   *big.Int       // Token balance
}

type AccountAsset struct {
	Account      common.Address  // Account address
	NativeAmount *big.Int        // Native token balances
	TokenAssets  []AccTokenAsset // Token assets
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

type ContractInfo struct {
	Creator common.Address
	Address common.Address
	Amount  *big.Int
}

type MintTokenTask struct {
	BubbleID *big.Int
	TxHash   common.Hash    // The transaction hash of the staking Token transaction
	RPC      string         // The bubble sub-chain operates the node rpc
	OpAddr   common.Address // Bubble The bubble main-chain operates address
	AccAsset *AccountAsset
}

type CreateBubbleTask struct {
	TxHash   common.Hash // The transaction hash of the createBubbleTask
	BubbleID *big.Int
}

type ReleaseBubbleTask struct {
	BubbleID *big.Int
}

type RemoteDeployTask struct {
	TxHash    common.Hash // The transaction hash of the remoteDeployTask
	BlockHash common.Hash
	BubbleID  *big.Int
	Address   common.Address
	Data      []byte
	RPC       string         // Bubble The bubble sub-chain operates the node rpc
	OpAddr    common.Address // Bubble The bubble main-chain operates address
}

type RemoteDestroyTask struct {
	BubbleID *big.Int
	//TxHash   common.Hash    // The transaction hash of the RemoteDestroyTask
	RPC    string         // Bubble The bubble sub-chain operates the node rpc
	OpAddr common.Address // Bubble The bubble main-chain operates address
}

type RemoteCallTask struct {
	TxHash   common.Hash // The transaction hash of the remoteDeployTask
	Caller   common.Address
	BubbleID *big.Int
	Contract common.Address
	Data     []byte
	RPC      string         // Bubble The bubble sub-chain operates the node rpc
	OpAddr   common.Address // Bubble The bubble main-chain operates address
}
