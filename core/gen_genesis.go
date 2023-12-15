package core

import (
	"encoding/json"
	"errors"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/hexutil"
	"github.com/bubblenet/bubble/common/math"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/xcom"
)

var _ = (*genesisSpecMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (g Genesis) MarshalJSON() ([]byte, error) {
	type Genesis struct {
		Config        *params.ChainConfig               `json:"config"`
		MulSigner     *params.MulSigner                 `json:"mulSigner"`
		OpConfig      *params.OpConfig                  `json:"opConfig"`
		EconomicModel *xcom.EconomicModel               `json:"economicModel"`
		Nonce         hexutil.Bytes                     `json:"nonce"`
		Timestamp     math.HexOrDecimal64               `json:"timestamp"`
		ExtraData     hexutil.Bytes                     `json:"extraData"`
		GasLimit      math.HexOrDecimal64               `json:"gasLimit"   gencodec:"required"`
		Coinbase      common.Address                    `json:"coinbase"`
		Alloc         map[common.Address]GenesisAccount `json:"alloc"      gencodec:"required"`
		Number        math.HexOrDecimal64               `json:"number"`
		GasUsed       math.HexOrDecimal64               `json:"gasUsed"`
		ParentHash    common.Hash                       `json:"parentHash"`
	}
	var enc Genesis
	enc.Config = g.Config
	enc.MulSigner = g.MulSigner
	enc.OpConfig = g.OpConfig
	enc.EconomicModel = g.EconomicModel
	enc.Nonce = g.Nonce
	enc.Timestamp = math.HexOrDecimal64(g.Timestamp)
	enc.ExtraData = g.ExtraData
	enc.GasLimit = math.HexOrDecimal64(g.GasLimit)
	enc.Coinbase = g.Coinbase
	enc.Alloc = g.Alloc
	enc.Number = math.HexOrDecimal64(g.Number)
	enc.GasUsed = math.HexOrDecimal64(g.GasUsed)
	enc.ParentHash = g.ParentHash
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (g *Genesis) UnmarshalJSON(input []byte) error {
	type Genesis struct {
		Config        *params.ChainConfig               `json:"config"`
		MulSigner     *params.MulSigner                 `json:"mulSigner"`
		OpConfig      *params.OpConfig                  `json:"opConfig"`
		EconomicModel *xcom.EconomicModel               `json:"economicModel"`
		Nonce         *hexutil.Bytes                    `json:"nonce"`
		Timestamp     *math.HexOrDecimal64              `json:"timestamp"`
		ExtraData     *hexutil.Bytes                    `json:"extraData"`
		GasLimit      *math.HexOrDecimal64              `json:"gasLimit"   gencodec:"required"`
		Coinbase      *common.Address                   `json:"coinbase"`
		Alloc         map[common.Address]GenesisAccount `json:"alloc"      gencodec:"required"`
		Number        *math.HexOrDecimal64              `json:"number"`
		GasUsed       *math.HexOrDecimal64              `json:"gasUsed"`
		ParentHash    *common.Hash                      `json:"parentHash"`
	}
	var dec Genesis
	dec.EconomicModel = g.EconomicModel
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Config != nil {
		g.Config = dec.Config
	}
	if dec.EconomicModel != nil {
		g.EconomicModel = dec.EconomicModel
	}

	if dec.MulSigner != nil {
		g.MulSigner = dec.MulSigner
	}

	if dec.OpConfig != nil {
		g.OpConfig = dec.OpConfig
	}

	if dec.Nonce != nil {
		g.Nonce = *dec.Nonce
	}
	if dec.Timestamp != nil {
		g.Timestamp = uint64(*dec.Timestamp)
	}
	if dec.ExtraData != nil {
		g.ExtraData = *dec.ExtraData
	}
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Genesis")
	}
	g.GasLimit = uint64(*dec.GasLimit)
	if dec.Coinbase != nil {
		g.Coinbase = *dec.Coinbase
	}
	if dec.Alloc == nil {
		return errors.New("missing required field 'alloc' for Genesis")
	} else {
		g.Alloc = dec.Alloc
	}
	if dec.Number != nil {
		g.Number = uint64(*dec.Number)
	}
	if dec.GasUsed != nil {
		g.GasUsed = uint64(*dec.GasUsed)
	}
	if dec.ParentHash != nil {
		g.ParentHash = *dec.ParentHash
	}
	return nil
}
