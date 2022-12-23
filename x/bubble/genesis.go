package bubble

import (
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/common/math"
	"github.com/bubblenet/bubble/p2p/enode"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/xcom"
)

// Operator Includes the operator's node ID, rpc url, operation address,
// and initial balance (to send the transaction in the child-chain to pay fees).
type Operator struct {
	NodeId      enode.IDv0     `json:"nodeId"` // Operator node id
	ElectronRPC string         `json:"ElectronRPC"`
	RPC         string         `json:"rpc"`     // Operation node RPC URL
	OpAddr      common.Address `json:"opAddr"`  // Address of operation
	Balance     *big.Int       `json:"balance"` // Operating address balance
}

// OpConfig is operator profiles, including main-chain operator
// and sub-chain operator profiles, interact between main-chain and sub-chain through the operator
type OpConfig struct {
	MainChain *Operator `json:"mainChain,omitempty"` // Main chain operator information configuration
	SubChain  *Operator `json:"subChain,omitempty"`  // Child chain operator information configuration
}

type StakingConfig struct {
	xcom.StakingConfig
	xcom.StakingConfigExtend
}

type EconomicModel struct {
	xcom.EconomicModel
	Staking StakingConfig `json:"staking"`
}

// GenesisAccount is an account in the state of the genesis block.
type GenesisAccount struct {
	Code       []byte                      `json:"code,omitempty"`
	Storage    map[common.Hash]common.Hash `json:"storage,omitempty"`
	Balance    *big.Int                    `json:"balance" gencodec:"required"`
	Nonce      uint64                      `json:"nonce,omitempty"`
	PrivateKey []byte                      `json:"secretKey,omitempty"` // for tests
}

// GenesisAlloc specifies the initial state that is part of the genesis block.
type GenesisAlloc map[common.Address]GenesisAccount

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
