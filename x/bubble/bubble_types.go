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

// BubbleSize is bubble chain size
type BubbleSize struct {
	OperatorL1Size uint
	OperatorL2Size uint
	CommitteeSize  uint
}

var (
	microBubble  = &BubbleSize{1, 1, 3}
	miniBubble   = &BubbleSize{1, 1, 6}
	mediumBubble = &BubbleSize{1, 1, 12}
	maxBubble    = &BubbleSize{1, 1, 24}

	sizeInfo = map[uint8]*BubbleSize{
		1: microBubble,
		2: miniBubble,
		3: mediumBubble,
		4: maxBubble,
	}
)

func GetBubbleSize(sizeCode uint8) (*BubbleSize, error) {
	size := sizeInfo[sizeCode]
	if size == nil {
		return nil, errors.New("unrecognized bubble size")
	}

	return size, nil
}

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
