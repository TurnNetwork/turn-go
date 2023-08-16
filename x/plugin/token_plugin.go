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
	"github.com/bubblenet/bubble/x/token"
	"sync"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/x/xcom"
)

var (
	tokenPluginOnce sync.Once
	tkp             *TokenPlugin
)

type TokenPlugin struct {
	MainOpAddr common.Address // Main chain operator address
}

func TokenPluginInstance() *TokenPlugin {
	tokenPluginOnce.Do(func() {
		log.Info("Init Token plugin ...")
		tkp = &TokenPlugin{}
	})
	return tkp
}

// SetMainOpAddr Set the main chain operator address
func (tkp *TokenPlugin) SetMainOpAddr(mainOpAddr common.Address) {
	tkp.MainOpAddr = mainOpAddr
}

// ExistAccount Add a list of minting account information
func (tkp *TokenPlugin) ExistAccount(state xcom.StateDB, mintAcc common.Address) bool {
	return false
}

// AddMintAccInfo Add a list of minting account information
func (tkp *TokenPlugin) AddMintAccInfo(state xcom.StateDB, mintAccInfo token.MintAccInfo) error {
	return token.SaveMintInfo(state, mintAccInfo)
}

// GetMintAccInfo Get a list of minting account information
func (tkp *TokenPlugin) GetMintAccInfo(state xcom.StateDB) (*token.MintAccInfo, error) {
	return token.GetMintAccInfo(state)
}

// BeginBlock implement BasePlugin
func (tkp *TokenPlugin) BeginBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	return nil
}

// EndBlock implement BasePlugin
func (tkp *TokenPlugin) EndBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	return nil
}

// Confirmed implement BasePlugin:does nothing
func (tkp *TokenPlugin) Confirmed(nodeId discover.NodeID, block *types.Block) error {
	return nil
}
