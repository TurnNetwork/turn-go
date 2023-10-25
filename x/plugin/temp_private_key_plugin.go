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

package plugin

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/state"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/eth/gasprice"
)

type TempPrivateKeyPlugin struct {
	gpo          *gasprice.Oracle
	stateFunc    func(number uint64) (*state.StateDB, *types.Header, error)
	applyMessage func(ctx context.Context, from common.Address, to *common.Address, value *big.Int, gasPrice *big.Int, data []byte,
		state *state.StateDB, header *types.Header, timeout time.Duration) ([]byte, error)
}

var (
	tempPrivateKeyOnce sync.Once
	tpkp               *TempPrivateKeyPlugin
)

func TempPrivateKeyContractInstance() *TempPrivateKeyPlugin {
	tempPrivateKeyOnce.Do(func() {
		tpkp = &TempPrivateKeyPlugin{}
	})
	return tpkp
}

func (tpkp *TempPrivateKeyPlugin) SetGaspriceOracle(aPIBackendGpo *gasprice.Oracle) {
	tpkp.gpo = aPIBackendGpo
}

func (tpkp *TempPrivateKeyPlugin) SetStateFunc(stateFunc func(number uint64) (*state.StateDB, *types.Header, error)) {
	tpkp.stateFunc = stateFunc
}

func (tpkp *TempPrivateKeyPlugin) SetApplyMessage(applyMessage func(ctx context.Context, from common.Address, to *common.Address, value *big.Int, gasPrice *big.Int, data []byte,
	state *state.StateDB, header *types.Header, timeout time.Duration) ([]byte, error)) {
	tpkp.applyMessage = applyMessage
}

func (tpkp *TempPrivateKeyPlugin) SuggestPrice() (*big.Int, error) {
	return tpkp.gpo.SuggestPrice(context.Background())
}

func (tpkp *TempPrivateKeyPlugin) GetGameContractOperator(number uint64, workAddress, gameContractAddress common.Address, gasPrice *big.Int) (common.Address, error) {
	from := workAddress
	to := &gameContractAddress
	id := crypto.Keccak256([]byte("issuer()address"))[:4]
	result, err := tpkp.doCall(number, from, to, big.NewInt(0), gasPrice, id)
	if err != nil {
		return common.BigToAddress(big.NewInt(0)), err
	}

	return common.BytesToAddress(result), nil
}

func (tpkp *TempPrivateKeyPlugin) GetLineOfCredit(number uint64, workAddress, gameContractAddress common.Address, gasPrice *big.Int) (*big.Int, error) {
	from := workAddress
	to := &gameContractAddress
	id := crypto.Keccak256([]byte("lineOfCredit()uint256"))[:4]
	result, err := tpkp.doCall(number, from, to, big.NewInt(0), gasPrice, id)
	if err != nil {
		return big.NewInt(0), err
	}

	return big.NewInt(0).SetBytes(result), nil
}

func (tpkp *TempPrivateKeyPlugin) doCall(number uint64, from common.Address, to *common.Address, value *big.Int, gasPrice *big.Int, data []byte) ([]byte, error) {

	state, header, _ := tpkp.stateFunc(number)
	defer state.ClearParentReference()

	// Setup context so it may be cancelled the call has completed
	// or, in case of unmetered gas, setup a context with a timeout.
	var timeout time.Duration = 5 * time.Second
	ctx := context.Background()
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}
	// Make sure the context is cancelled when the call has completed
	// this makes sure resources are cleaned up.
	defer cancel()

	return tpkp.applyMessage(ctx, from, to, value, gasPrice, data, state, header, timeout)
}
