// Copyright 2021 The Bubble Network Authors
// This file is part of the bubble library.
//
// The bubble library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The bubble library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the bubble library. If not, see <http://www.gnu.org/licenses/>.

package vm

import (
	"fmt"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/x/bubble"
	"math/big"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/x/plugin"
)

const (
	TxCreateBubble    = 0001
	TxReleaseBubble   = 0002
	CallGetBubbleInfo = 1001
)

type BubbleContract struct {
	Plugin *plugin.BubblePlugin

	Contract *Contract
	Evm      *EVM
}

func (bc *BubbleContract) RequiredGas(input []byte) uint64 {
	if checkInputEmpty(input) {
		return 0
	}
	return params.BubbleGas
}

func (bc *BubbleContract) Run(input []byte) ([]byte, error) {
	if checkInputEmpty(input) {
		return nil, nil
	}
	return execBubbleContract(input, bc.FnSigns())
}

func (bc *BubbleContract) FnSigns() map[uint16]interface{} {
	return map[uint16]interface{}{
		// Set
		TxCreateBubble:  bc.createBubble,
		TxReleaseBubble: bc.releaseBubble,
		// Get
		CallGetBubbleInfo: bc.getBubbleInfo,
	}
}

func (bc *BubbleContract) CheckGasPrice(gasPrice *big.Int, fcode uint16) error {
	return nil
}

// createBubble create a Bubble chain using operator nodes and candidate nodes
func (bc *BubbleContract) createBubble(genesisData [][]byte) ([]byte, error) {

	from := bc.Contract.CallerAddress
	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber
	blockHash := bc.Evm.Context.BlockHash
	currentNonce := bc.Evm.StateDB.GetNonce(from)
	parentHash := bc.Evm.Context.ParentHash

	log.Debug("Call createBubble of bubbleContract", "blockNumber", blockNumber.Uint64(), "blockHash", blockHash.TerminalString(),
		"txHash", txHash.Hex(), "from", from.String())

	if !bc.Contract.UseGas(params.CreateBubbleGas) {
		return nil, ErrOutOfGas
	}

	bubbleID, err := bc.Plugin.CreateBubble(blockHash, blockNumber, from, currentNonce, parentHash)
	if nil != err {
		if bizErr, ok := err.(*common.BizError); ok {
			return txResultHandler(vm.BubbleContractAddr, bc.Evm, "createBubble", bizErr.Error(), TxCreateBubble, bizErr)
		} else {
			log.Error("Failed to createBubble", "txHash", txHash, "blockNumber", blockNumber, "err", err)
			return nil, err
		}
	}

	// TODO: store genesisData and return the index for DApp reuse
	// TODO: store genesisData by a other interface

	return txResultHandlerWithRes(vm.BubbleContractAddr, bc.Evm, "", "", TxCreateBubble, int(common.NoErr.Code), bubbleID), nil
}

// releaseBubble release the node resources of a bubble chain and delete it`s information
func (bc *BubbleContract) releaseBubble(bubbleID uint32) ([]byte, error) {

	txHash := bc.Evm.StateDB.TxHash()
	blockNumber := bc.Evm.Context.BlockNumber
	blockHash := bc.Evm.Context.BlockHash
	from := bc.Contract.CallerAddress

	log.Debug("Call releaseBubble of bubbleContract", "blockNumber", blockNumber.Uint64(),
		"blockHash", blockHash.TerminalString(), "txHash", txHash.Hex(), "from", from.String())

	if !bc.Contract.UseGas(params.ReleaseBubbleGas) {
		return nil, ErrOutOfGas
	}

	bub, err := bc.Plugin.GetBubbleInfo(blockHash, bubbleID)
	if snapshotdb.NonDbNotFoundErr(err) {
		log.Error("Failed to releaseBubble by GetBubbleInfo", "txHash", txHash,
			"blockNumber", blockNumber, "blockHash", blockHash.Hex(), "nodeId", bub.BubbleId, "err", err)
		return nil, err
	}
	if bub == nil {
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "releaseBubble",
			fmt.Sprintf("bubble %d is not exist", bub.BubbleId), TxReleaseBubble, bubble.ErrBubbleNotExist)
	}

	if from != bub.Creator {
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "releaseBubble",
			fmt.Sprintf("txSender: %s, bubble Creator: %s", from, bub.Creator), TxReleaseBubble, bubble.ErrSenderIsNotCreator)
	}

	// TODO: can release the bubble chain in the building state？
	if bub.State != bubble.PreReleaseStatus {
		return txResultHandler(vm.BubbleContractAddr, bc.Evm, "releaseBubble", fmt.Sprintf("bubble %d unable to release ", bub.BubbleId),
			TxReleaseBubble, bubble.ErrBubbleUnableRelease)
	}

	if err := bc.Plugin.ReleaseBubble(blockHash, blockNumber, bubbleID); err != nil {
		return nil, err
	}

	return txResultHandler(vm.BubbleContractAddr, bc.Evm, "", "", TxReleaseBubble, common.NoErr)
}

// getBubbleInfo return the bubble information by bubble ID
func (bc *BubbleContract) getBubbleInfo(bubbleID uint32) ([]byte, error) {
	blockHash := bc.Evm.Context.BlockHash

	bub, err := bc.Plugin.GetBubbleInfo(blockHash, bubbleID)
	if err != nil {
		return callResultHandler(bc.Evm, fmt.Sprintf("getBubbleInfo, bubbleID: %d", bubbleID), bub, bubble.ErrBubbleNotExist), nil
	}

	return callResultHandler(bc.Evm, fmt.Sprintf("getBubbleInfo, bubbleID: %d", bubbleID), bub, nil), nil
}
