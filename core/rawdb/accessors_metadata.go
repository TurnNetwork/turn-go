// Copyright 2018 The go-ethereum Authors
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

package rawdb

import (
	"encoding/json"

	"github.com/bubblenet/bubble/ethdb"

	"github.com/bubblenet/bubble/x/xcom"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/rlp"
)

// ReadDatabaseVersion retrieves the version number of the database.
func ReadDatabaseVersion(db ethdb.KeyValueReader) *uint64 {
	var version uint64

	enc, _ := db.Get(databaseVerisionKey)
	if len(enc) == 0 {
		return nil
	}
	if err := rlp.DecodeBytes(enc, &version); err != nil {
		return nil
	}

	return &version
}

// WriteDatabaseVersion stores the version number of the database
func WriteDatabaseVersion(db ethdb.KeyValueWriter, version uint64) {
	enc, err := rlp.EncodeToBytes(version)
	if err != nil {
		log.Crit("Failed to encode database version", "err", err)
	}
	if err = db.Put(databaseVerisionKey, enc); err != nil {
		log.Crit("Failed to store the database version", "err", err)
	}
}

// ReadChainConfig retrieves the consensus settings based on the given genesis hash.
func ReadChainConfig(db ethdb.KeyValueReader, hash common.Hash) *params.ChainConfig {
	data, _ := db.Get(configKey(hash))
	if len(data) == 0 {
		return nil
	}
	var config params.ChainConfig
	if err := json.Unmarshal(data, &config); err != nil {
		log.Error("Invalid chain config JSON", "hash", hash, "err", err)
		return nil
	}
	return &config
}

// WriteChainConfig writes the chain config settings to the database.
func WriteChainConfig(db ethdb.KeyValueWriter, hash common.Hash, cfg *params.ChainConfig) {
	if cfg == nil {
		return
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		log.Crit("Failed to JSON encode chain config", "err", err)
	}
	if err := db.Put(configKey(hash), data); err != nil {
		log.Crit("Failed to store chain config", "err", err)
	}
}

// ReadMulSigner retrieves the Multi-signature verifier configuration information settings based on the given genesis hash.
func ReadMulSigner(db ethdb.KeyValueReader, hash common.Hash) *params.MulSigner {
	data, _ := db.Get(mulSignerKey(hash))
	if len(data) == 0 {
		return nil
	}
	var mulSigner params.MulSigner
	if err := json.Unmarshal(data, &mulSigner); err != nil {
		log.Error("Invalid Multi-signature verifier config JSON", "hash", hash, "err", err)
		return nil
	}
	return &mulSigner
}

// WriteMulSigner writes the Multi-signature verifier configuration information settings to the database.
func WriteMulSigner(db ethdb.KeyValueWriter, hash common.Hash, cfg *params.MulSigner) {
	if cfg == nil {
		return
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		log.Crit("Failed to JSON encode Multi-signature verifier config", "err", err)
	}
	if err := db.Put(mulSignerKey(hash), data); err != nil {
		log.Crit("Failed to store Multi-signature verifier config", "err", err)
	}
}

// ReadOperatorConfig retrieves the consensus settings based on the given genesis hash.
func ReadOperatorConfig(db ethdb.KeyValueReader, hash common.Hash) *params.OpConfig {
	data, _ := db.Get(opConfigKey(hash))
	if len(data) == 0 {
		return nil
	}
	var opConfig params.OpConfig
	if err := json.Unmarshal(data, &opConfig); err != nil {
		log.Error("Invalid Operator config JSON", "hash", hash, "err", err)
		return nil
	}
	return &opConfig
}

// WriteOperatorConfig writes the operator config settings to the database.
func WriteOperatorConfig(db ethdb.KeyValueWriter, hash common.Hash, cfg *params.OpConfig) {
	if cfg == nil {
		return
	}
	data, err := json.Marshal(cfg)
	if err != nil {
		log.Crit("Failed to JSON encode operator config", "err", err)
	}
	if err := db.Put(opConfigKey(hash), data); err != nil {
		log.Crit("Failed to store operator config", "err", err)
	}
}

// WriteEconomicModel writes the EconomicModel settings to the database.
func WriteEconomicModel(db ethdb.Writer, hash common.Hash, ec *xcom.EconomicModel) {
	if ec == nil {
		return
	}

	data, err := json.Marshal(ec)
	if err != nil {
		log.Crit("Failed to JSON encode EconomicModel config", "err", err)
	}
	if err := db.Put(economicModelKey(hash), data); err != nil {
		log.Crit("Failed to store EconomicModel", "err", err)
	}
}

// WriteEconomicModelExtend writes the EconomicModelExtend settings to the database.
func WriteEconomicModelExtend(db ethdb.Writer, hash common.Hash, ec *xcom.EconomicModelExtend) {
	if ec == nil {
		return
	}

	data, err := json.Marshal(ec)
	if err != nil {
		log.Crit("Failed to JSON encode EconomicModelExtend config", "err", err)
	}
	if err := db.Put(economicModelExtendKey(hash), data); err != nil {
		log.Crit("Failed to store EconomicModelExtend", "err", err)
	}
}

// ReadEconomicModel retrieves the EconomicModel settings based on the given genesis hash.
func ReadEconomicModel(db ethdb.Reader, hash common.Hash) *xcom.EconomicModel {
	data, _ := db.Get(economicModelKey(hash))
	if len(data) == 0 {
		return nil
	}

	var ec xcom.EconomicModel
	// reset the global ec
	if err := json.Unmarshal(data, &ec); err != nil {
		log.Error("Invalid EconomicModel JSON", "hash", hash, "err", err)
		return nil
	}
	return &ec
}

// ReadEconomicModelExtend retrieves the EconomicModelExtend settings based on the given genesis hash.
func ReadEconomicModelExtend(db ethdb.Reader, hash common.Hash) *xcom.EconomicModelExtend {
	data, _ := db.Get(economicModelExtendKey(hash))
	if len(data) == 0 {
		return nil
	}

	var ec xcom.EconomicModelExtend
	// reset the global ec
	if err := json.Unmarshal(data, &ec); err != nil {
		log.Error("Invalid EconomicModelExtend JSON", "hash", hash, "err", err)
		return nil
	}
	return &ec
}
