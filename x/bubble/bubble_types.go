package bubble

import (
	"errors"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/common/math"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/stakingL2"
	"github.com/bubblenet/bubble/x/xcom"
)

// SizeConfig is bubble chain size
type SizeConfig struct {
	OperatorL1Size   uint
	OperatorL2Size   uint
	CommitteeSize    uint
	MinStakingAmount *big.Int
}

var (
	microConfig  = &SizeConfig{OperatorL1Size: 1, OperatorL2Size: 1, CommitteeSize: 3, MinStakingAmount: new(big.Int).Mul(big.NewInt(params.BUB), big.NewInt(5))}
	miniConfig   = &SizeConfig{OperatorL1Size: 1, OperatorL2Size: 1, CommitteeSize: 6, MinStakingAmount: new(big.Int).Mul(big.NewInt(params.BUB), big.NewInt(10))}
	mediumConfig = &SizeConfig{OperatorL1Size: 1, OperatorL2Size: 1, CommitteeSize: 12, MinStakingAmount: new(big.Int).Mul(big.NewInt(params.BUB), big.NewInt(20))}
	maxConfig    = &SizeConfig{OperatorL1Size: 1, OperatorL2Size: 1, CommitteeSize: 24, MinStakingAmount: new(big.Int).Mul(big.NewInt(params.BUB), big.NewInt(40))}

	sizeConfigs = map[uint8]*SizeConfig{
		1: microConfig,
		2: miniConfig,
		3: mediumConfig,
		4: maxConfig,
	}
)

func GetSizeConfig(sizeCode uint8) (*SizeConfig, error) {
	size := sizeConfigs[sizeCode]
	if size == nil {
		return nil, errors.New("unrecognized bubble size")
	}

	return size, nil
}

// Status bubble chain status
type State uint8

const (
	//BuildingStatus   = 0
	ActiveStatus     State = 1
	PreReleaseStatus       = 2
	ReleasedStatus         = 3
)

// BubTxType Define the Bubble's trade type
type BubTxType uint

const (
	CreateBubble  BubTxType = iota // createBubble transaction type
	ReleaseBubble                  // releaseBubble transaction type
	StakingToken                   // StakingToken transaction type
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
	Size        uint8
	OperatorsL1 []*Operator
	OperatorsL2 []*Operator
	MicroNodes  CandidateQueue
}

type BubStatus struct {
	BubbleId        *big.Int
	State           State
	ContractCount   int
	CreateBlock     uint64
	PreReleaseBlock uint64
	ReleaseBlock    uint64
}

type BubMutable struct {
	StakingTokenTxHashList  []common.Hash // List of stakingToken transaction hashes
	WithdrewTokenTxHashList []common.Hash // List of withdrewToken transaction hashes
	SettleBubbleTxHashList  []common.Hash // List of settleBubble transaction hashes
}

type Bubble struct {
	Basics *BubBasics // Bubble Basics
	*BubStatus
	*BubMutable
}

const (
	OperatorNode    = 1
	MaxNodeUseRatio = 0.75
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

type RemoteDeployTask struct {
	TxHash    common.Hash // The transaction hash of the remoteDeployTask
	BlockHash common.Hash
	BubbleID  *big.Int
	Address   common.Address
	Data      []byte
	RPC       string         // Bubble The bubble sub-chain operates the node rpc
	OpAddr    common.Address // Bubble The bubble main-chain operates address
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

// GenesisAlloc specifies the initial state that is part of the genesis block.
type GenesisAlloc map[common.Address]GenesisAccount

// GenesisAccount is an account in the state of the genesis block.
type GenesisAccount struct {
	Code       []byte                      `json:"code,omitempty"`
	Storage    map[common.Hash]common.Hash `json:"storage,omitempty"`
	Balance    *big.Int                    `json:"balance" gencodec:"required"`
	Nonce      uint64                      `json:"nonce,omitempty"`
	PrivateKey []byte                      `json:"secretKey,omitempty"` // for tests
}

type StakingConfig struct {
	xcom.StakingConfig
	xcom.StakingConfigExtend
}

type EconomicModel struct {
	xcom.EconomicModel
	Staking StakingConfig `json:"staking"`
}

// GenesisL2 specifies the header fields, state of a layer2 genesis block. It also defines hard
// fork switch-over blocks through the chain configuration.
type GenesisL2 struct {
	Config        *params.ChainConfig `json:"config"`
	OpConfig      *OpConfig           `json:"opConfig"`
	EconomicModel *EconomicModel      `json:"economicModel"`
	Nonce         hexutil.Bytes       `json:"nonce"`
	Timestamp     math.HexOrDecimal64 `json:"timestamp"`
	ExtraData     hexutil.Bytes       `json:"extraData"`
	GasLimit      math.HexOrDecimal64 `json:"gasLimit"   gencodec:"required"`
	Coinbase      common.Address      `json:"coinbase"`
	Alloc         GenesisAlloc        `json:"alloc"      gencodec:"required"`

	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
	Number     math.HexOrDecimal64 `json:"number"`
	GasUsed    math.HexOrDecimal64 `json:"gasUsed"`
	ParentHash common.Hash         `json:"parentHash"`
}
