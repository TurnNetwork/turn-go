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

package cbft

import (
	"encoding/json"

	"github.com/bubblenet/bubble/consensus/cbft/state"
	"github.com/bubblenet/bubble/consensus/cbft/types"
	"github.com/bubblenet/bubble/crypto/bls"
)

type Status struct {
	Tree      *types.BlockTree `json:"blockTree"`
	State     *state.ViewState `json:"state"`
	Validator bool             `json:"validator"`
}

// API defines an exposed API function interface.
type API interface {
	Status() []byte
	Evidences() string
	GetPrepareQC(number uint64) *types.QuorumCert
	GetSchnorrNIZKProve() (*bls.SchnorrProof, error)
}

// PublicDebugConsensusAPI provides an API to access the Bubble blockchain.
// It offers only methods that operate on public data that
// is freely available to anyone.
type PublicDebugConsensusAPI struct {
	engine API
}

// NewDebugConsensusAPI creates a new Bubble blockchain API.
func NewDebugConsensusAPI(engine API) *PublicDebugConsensusAPI {
	return &PublicDebugConsensusAPI{engine: engine}
}

// ConsensusStatus returns the status data of the consensus engine.
func (s *PublicDebugConsensusAPI) ConsensusStatus() *Status {
	b := s.engine.Status()
	var status Status
	err := json.Unmarshal(b, &status)
	if err == nil {
		return &status
	}
	return nil
}

// GetPrepareQC returns the QC certificate corresponding to the blockNumber.
func (s *PublicDebugConsensusAPI) GetPrepareQC(number uint64) *types.QuorumCert {
	return s.engine.GetPrepareQC(number)
}

// PublicBubbleConsensusAPI provides an API to access the Bubble blockchain.
// It offers only methods that operate on public data that
// is freely available to anyone.
type PublicBubbleConsensusAPI struct {
	engine API
}

// NewPublicBubbleConsensusAPI creates a new Bubble blockchain API.
func NewPublicBubbleConsensusAPI(engine API) *PublicBubbleConsensusAPI {
	return &PublicBubbleConsensusAPI{engine: engine}
}

// Evidences returns the relevant data of the verification.
func (s *PublicBubbleConsensusAPI) Evidences() string {
	return s.engine.Evidences()
}

// PublicAdminConsensusAPI provides an API to access the Bubble blockchain.
// It offers only methods that operate on public data that
// is freely available to anyone.
type PublicAdminConsensusAPI struct {
	engine API
}

// NewPublicAdminConsensusAPI creates a new Bubble blockchain API.
func NewPublicAdminConsensusAPI(engine API) *PublicAdminConsensusAPI {
	return &PublicAdminConsensusAPI{engine: engine}
}

func (s *PublicAdminConsensusAPI) GetSchnorrNIZKProve() string {
	proof, err := s.engine.GetSchnorrNIZKProve()
	if nil != err {
		return err.Error()
	}
	proofByte, err := proof.MarshalText()
	if nil != err {
		return err.Error()
	}
	return string(proofByte)
}
