// Copyright 2015 The go-ethereum Authors
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

package core

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/consensus"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/core/state"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/core/vm"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/rlp"
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for block rewards
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(config *params.ChainConfig, bc *BlockChain, engine consensus.Engine) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
		engine: engine,
	}
}

// Process processes the state changes according to the Ethereum rules by running
// the transaction messages using the statedb and applying any rewards to
// the processor (coinbase).
//
// Process returns the receipts and logs accumulated during the process and
// returns the amount of gas that was used in the process. If any of the
// transactions failed to execute due to insufficient gas it will return an error.
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, cfg vm.Config) (types.Receipts, []*types.Log, uint64, error) {
	var (
		receipts types.Receipts
		usedGas  = new(uint64)
		header   = block.Header()
		allLogs  []*types.Log
		gp       = new(GasPool).AddGas(block.GasLimit())
	)
	blockContext := NewEVMBlockContext(header, p.bc)
	vmenv := vm.NewEVM(blockContext, vm.TxContext{}, snapshotdb.Instance(), statedb, p.config, cfg)
	if bcr != nil {
		// BeginBlocker()
		if err := bcr.BeginBlocker(header, statedb); nil != err {
			log.Error("Failed to call BeginBlocker on StateProcessor", "blockNumber", block.Number(),
				"blockHash", block.Hash(), "err", err)
			return nil, nil, 0, err
		}
	}

	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		msg, err := tx.AsMessage(types.MakeSigner(p.config))
		if err != nil {
			return nil, nil, 0, err
		}
		statedb.Prepare(tx.Hash(), block.Hash(), i)
		//preUsedGas := uint64(0)

		receipt, err := applyTransaction(msg, p.config, p.bc, gp, statedb, header, tx, usedGas, vmenv)
		if err != nil {
			log.Error("Failed to execute tx on StateProcessor", "blockNumber", block.Number(),
				"blockHash", block.Hash().TerminalString(), "txHash", tx.Hash().String(), "err", err)
			return nil, nil, 0, fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash().Hex(), err)
		}
		//log.Debug("tx process success", "txHash", tx.Hash().Hex(), "txTo", tx.To().Hex(), "dataLength", len(tx.Data()), "toCodeSize", statedb.GetCodeSize(*tx.To()), "txUsedGas", *usedGas-preUsedGas)
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}

	if bcr != nil {
		// EndBlocker()
		if err := bcr.EndBlocker(header, statedb); nil != err {
			log.Error("Failed to call EndBlocker on StateProcessor", "blockNumber", block.Number(),
				"blockHash", block.Hash().TerminalString(), "err", err)
			return nil, nil, 0, err
		}
	}

	// Finalize the block, applying any consensus engine specific extras (e.g. block rewards)
	p.engine.Finalize(p.bc, header, statedb, block.Transactions(), receipts)

	return receipts, allLogs, *usedGas, nil
}

func applyTransaction(msg types.Message, config *params.ChainConfig, bc ChainContext, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *uint64, evm *vm.EVM) (*types.Receipt, error) {
	// Create a new context to be used in the EVM environment
	txContext := NewEVMTxContext(msg)
	// Add addresses to access list if applicable
	log.Trace("execute tx start", "blockNumber", header.Number, "txHash", tx.Hash().String())

	// Update the evm with the new transaction context.
	evm.Reset(txContext, statedb)
	// Apply the transaction to the current state (included in the env)
	result, err := ApplyMessage(evm, msg, gp)
	if err != nil {
		return nil, err
	}
	// Update the state with pending changes
	statedb.Finalise(true)

	var root []byte
	*usedGas += result.UsedGas

	// Create a new receipt for the transaction, storing the intermediate root and gas used by the tx
	// based on the eip phase, we're passing whether the root touch-delete accounts.
	receipt := types.NewReceipt(root, result.Failed(), *usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = result.UsedGas
	// if the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(evm.TxContext.Origin, tx.Nonce())
	}
	// Set the receipt logs
	if result.Failed() {
		if bizError, ok := result.Err.(*common.BizError); ok {
			buf := new(bytes.Buffer)
			res := strconv.Itoa(int(bizError.Code))
			if err := rlp.Encode(buf, [][]byte{[]byte(res)}); nil != err {
				log.Error("Cannot RlpEncode the log data", "data", bizError.Code, "err", err)
				return nil, err
			}
			receipt.Logs = []*types.Log{
				&types.Log{
					Address:     *msg.To(),
					Topics:      nil,
					Data:        buf.Bytes(),
					BlockNumber: header.Number.Uint64(),
				},
			}
		} else {
			receipt.Logs = statedb.GetLogs(tx.Hash())
		}
	} else {
		receipt.Logs = statedb.GetLogs(tx.Hash())
	}
	//create a bloom for filtering
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	receipt.BlockHash = statedb.BlockHash()
	receipt.BlockNumber = header.Number
	receipt.TransactionIndex = uint(statedb.TxIndex())
	return receipt, err
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(config *params.ChainConfig, bc ChainContext, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *uint64, cfg vm.Config) (*types.Receipt, error) {
	msg, err := tx.AsMessage(types.MakeSigner(config))
	if err != nil {
		return nil, err
	}
	// Create a new context to be used in the EVM environment
	blockContext := NewEVMBlockContext(header, bc)
	vmenv := vm.NewEVM(blockContext, vm.TxContext{}, snapshotdb.Instance(), statedb, config, cfg)
	return applyTransaction(msg, config, bc, gp, statedb, header, tx, usedGas, vmenv)
}
