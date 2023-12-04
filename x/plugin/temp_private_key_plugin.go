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

	"github.com/bubblenet/bubble/eth/gasprice"
)

type TempPrivateKeyPlugin struct {
	gpo *gasprice.Oracle
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

func (tpkp *TempPrivateKeyPlugin) SuggestPrice() (*big.Int, error) {
	return tpkp.gpo.SuggestPrice(context.Background())
}
