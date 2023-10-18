// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"fmt"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/p2p/discover"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0xbd3e148a58392bc546aee5e69aa751c8e572e2cd2c82d06d9418a686b82252b9")
	TestnetGenesisHash = common.HexToHash("0x47c7f3dea852b28c45d5fee6c950321486780065745179c7a296934e493f18e1")
)

var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{
	MainnetGenesisHash: MainnetTrustedCheckpoint,
	TestnetGenesisHash: TestnetTrustedCheckpoint,
}

var (
	initialMainNetConsensusNodes = []InitNode{
		{
			"enode://c6a19b488e56c2cb0fdcdf245809cdffcb6af843c70b70790e58a6f161837c5f691bebf0b2c41a30fe8ab2175070106d28bd4b06b234fe40eb7e4d3eec14ae44@mf1.bubble.org:16789",
			"edfb6f219bf7f044fd15b11a89e0781777d02d36aee4c8a656994c523945b57f5eb9e3d3f2705671859bf242077225004782a1edafdb3d318ce0de3e50a7dbe2246e857c94d0c3534d8e9105e093dc49ff74b3e316427b04ca260dcc9d7c1d99",
			"",
		},
		{
			"enode://df3216d519f40905c005791486eb75e574d42832af93f639d75f44c1fea306723b2b31e898d9dc1fab055579014402c15426105cbcf53660edc305269d906200@mf2.bubble.org:16789",
			"8d808171d8a77aa68130c23fce38f11271adf82911d0082d103cd69a8f7ab60b4e9c4b57f2f4a08fe33f11d1e79e8e13901eab37337ab7ef5f707af6accb0a4b9c8863c4954ef63b7431ba9fcb583cc184525d5bac25bf963bf7695b22c9c509",
			"",
		},
		{
			"enode://c5fe718e51ab5a3a03264538989ba75c6e56b69668b8d5ac3484c1c29e5756ed75aa5ea0ae7b1401387b9fdb71719764638524786e9ce38bfbfc0cb849ab519a@mf3.bubble.org:16789",
			"c1c4af29449a81c8d08c237d61263d9ed7ca1e075b78c4a9ef0c6d39dcaab4f6dd9ab13d5a858a17fdb354c0e2c1121719bf4e7098afce1f42478ff61430dc368ada9a68fcf19ceefcfb1bb05c76cfd667af589cd8c550eaca353587768f1903",
			"",
		},
		{
			"enode://0b949b48ad6f5be62fee3b8e83dd59b43c3cac79bed861ce00d00d4102960171f6f108dea61981153c8d36257a101fa9bca6d42016a3b584319918edfb157e10@mf4.bubble.org:16789",
			"02b99c2daa814f838b00b47c4bbeb7074c1b7d99c7a97008cb920f8c5ad30bbf8c7bd631a35ed25000c98b7e35e76700c3c262ebc39dc9889354d8b983f25b5397044a8b0bad5bc7e5734a49712661eba113f60085ca3ddb7f2a32fd0498f28d",
			"",
		},
		{
			"enode://42443557406e64bea58a15fc12b9f78a723cccec652a3cc7edc9b50a74198ff0438d1a107f140681f03a23c7e93c0125e809909a4280298a3ae5947843116786@mf5.bubble.org:16789",
			"cbbee10412141c4784afb4c7dcc749c01b9683372b599e63f544c829477c56fccb79f6ff8059c2b0a913907b76300805259bd20381608f824a47b796cb6454501d09473516fe438daf1306578b790c461558b0359d6cfebb623dbf400b8fcd90",
			"",
		},
		{
			"enode://b9187f73960d062f3dd96099abab78178a03f772d628c18bc654866761dadeac8956e77a72115476dc8bbcc8c58d4c492e9db003ba4b36630926ea4fb93f1dc5@mf6.bubble.org:16789",
			"43fa5c17c95c7ec9524e8f3ef19e25507dcc26b3b5c81fd44d169e8049eae6741e68f07d206704d3d02461a05d5d5a1832e91c3ec73768f4088a944b953eb118c682230e17f47aaa4a5685a3d103adf9a1d55dba66822cb72c22e51fdce1178d",
			"",
		},
		{
			"enode://b7913cc6f6e47d5f6fe04e1b33521bec788fc4810bf5b9e232a499f4dfd7117974912645f6385e155cb3bab4eb4ef1e51502d302e101dfb4cee409ce702ba657@mf7.bubble.org:16789",
			"490979dd7cf894a729438bff0d14fb514bcd96cd2a89645162e02860916781da78dfdc8a2c24aa0ae233ce7d32d53214c82ce04fd2dabaf60aa819b9d6400038d18e82ad69902ab257aa350c366b1387e00794b2b2b999f778ca89d85591018d",
			"",
		},
	}

	initialTestnetConsensusNodes = []InitNode{
		{
			"enode://b7f1f7757a900cce7ce4caf8663ecf871205763ac201c65f9551d5b841731a9cd9550bc05f3a16fbc2ef589c9faeef74d4500b60d76047939e2ba7fa4a5915aa@127.0.0.1:16789",
			"f1735bac863706b49809a4e635fe0c2e224aef5ad549f18ba3f2f6b61c0c9d0005f12d497a301ba26a8aaf009c90e4198301875002984c5cd9bd614cd2fbcb81c57f6355a8400d56c20804edfb34782c1f2eadda82c8b226aa4a71bfa60be8c",
			"",
		},
	}

	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:     big.NewInt(2500),
		EmptyBlock:  "on",
		EIP155Block: big.NewInt(1),
		Cbft: &CbftConfig{
			InitialNodes:  ConvertNodeUrl(initialMainNetConsensusNodes),
			Amount:        10,
			ValidatorMode: "dpos",
			Period:        20000,
		},
		GenesisVersion: GenesisVersion,
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "mainnet",
		SectionIndex: 193,
		SectionHead:  common.HexToHash("0xc2d574295ecedc4d58530ae24c31a5a98be7d2b3327fba0dd0f4ed3913828a55"),
		CHTRoot:      common.HexToHash("0x5d1027dfae688c77376e842679ceada87fd94738feb9b32ef165473bfbbb317b"),
		BloomRoot:    common.HexToHash("0xd38be1a06aabd568e10957fee4fcc523bc64996bcf31bae3f55f86e0a583919f"),
	}

	// TestnetChainConfig is the chain parameters to run a node on the test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:     big.NewInt(104),
		EmptyBlock:  "on",
		EIP155Block: big.NewInt(1),
		Cbft: &CbftConfig{
			InitialNodes:  ConvertNodeUrl(initialTestnetConsensusNodes),
			Amount:        10,
			ValidatorMode: "dpos",
			Period:        20000,
		},
		GenesisVersion: GenesisVersion,
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the test network.
	TestnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "testnet",
		SectionIndex: 123,
		SectionHead:  common.HexToHash("0xa372a53decb68ce453da12bea1c8ee7b568b276aa2aab94d9060aa7c81fc3dee"),
		CHTRoot:      common.HexToHash("0x6b02e7fada79cd2a80d4b3623df9c44384d6647fc127462e1c188ccd09ece87b"),
		BloomRoot:    common.HexToHash("0xf2d27490914968279d6377d42868928632573e823b5d1d4a944cba6009e16259"),
	}

	GrapeChainConfig = &ChainConfig{
		ChainID:     big.NewInt(304),
		EmptyBlock:  "on",
		EIP155Block: big.NewInt(3),
		Cbft: &CbftConfig{
			Period: 3,
		},
		GenesisVersion: GenesisVersion,
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), "", big.NewInt(0), big.NewInt(0), nil, nil, nil, GenesisVersion}

	TestChainConfig = &ChainConfig{big.NewInt(1), "", big.NewInt(0), big.NewInt(0), nil, new(CbftConfig), nil, GenesisVersion}

	// DefaultFrpsCfg Default frps configuration
	DefaultFrpsCfg = &FrpsConfig{"0.0.0.0", 7000, &AuthConfig{"token", true, true, "12345678_"}}
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	Name         string      `json:"-"`
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID     *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection
	EmptyBlock  string   `json:"emptyBlock"`
	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EWASMBlock  *big.Int `json:"ewasmBlock,omitempty"`  // EWASM switch block (nil = no fork, 0 = already activated)
	// Various consensus engines
	Clique         *CliqueConfig `json:"clique,omitempty"`
	Cbft           *CbftConfig   `json:"cbft,omitempty"`
	Frps           *FrpsConfig   `json:"frps,omitempty"`
	GenesisVersion uint32        `json:"genesisVersion"`
}

type CbftNode struct {
	Node      discover.Node `json:"node"`
	BlsPubKey bls.PublicKey `json:"blsPubKey"`
	RPC       string
}

type InitNode struct {
	Enode     string
	BlsPubkey string
	RPC       string
}

// AuthConfig frp server authenticate struct:Enable token authentication
type AuthConfig struct {
	Method       string // authentication method
	HeartBeats   bool   // authenticate heartbeats
	NewWorkConns bool   // authenticate_new_work_conns
	Token        string // token authentication
}

// FrpsConfig frp server configuration structure
type FrpsConfig struct {
	ServerIP   string      // frp server ip
	ServerPort int         // frp server port
	Auth       *AuthConfig // Enable authentication mode
}

type CbftConfig struct {
	Period        uint64     `json:"period,omitempty"`        // Number of seconds between blocks to enforce
	Amount        uint32     `json:"amount,omitempty"`        //The maximum number of blocks generated per cycle
	InitialNodes  []CbftNode `json:"initialNodes,omitempty"`  //Genesis consensus node
	ValidatorMode string     `json:"validatorMode,omitempty"` //Validator mode for easy testing
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Clique != nil:
		engine = c.Clique
	case c.Cbft != nil:
		engine = c.Cbft
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v EIP155: %v Engine: %v}",
		c.ChainID,
		c.EIP155Block,
		engine,
	)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	//	return isForked(c.EIP155Block, num)
	return true
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	return GasTableConstantinople
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

func ConvertNodeUrl(initialNodes []InitNode) []CbftNode {
	bls.Init(bls.BLS12_381)
	NodeList := make([]CbftNode, 0, len(initialNodes))
	for _, n := range initialNodes {

		cbftNode := new(CbftNode)

		if node, err := discover.ParseNode(n.Enode); nil == err {
			cbftNode.Node = *node
		}

		if n.BlsPubkey != "" {
			var blsPk bls.PublicKey
			if err := blsPk.UnmarshalText([]byte(n.BlsPubkey)); nil == err {
				cbftNode.BlsPubKey = blsPk
			}
		}

		cbftNode.RPC = n.RPC
		NodeList = append(NodeList, *cbftNode)
	}
	return NodeList
}
