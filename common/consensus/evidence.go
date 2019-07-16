package consensus

import (
	"crypto/ecdsa"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

type EvidenceType int32
type Evidence interface {
	Verify(ecdsa.PublicKey) error
	Equal(Evidence) bool
	//return lowest number
	BlockNumber() uint64
	Hash() []byte
	Address() common.Address
	Validate() error
	Type() EvidenceType
}

type Evidences []Evidence

type EvidencePool interface {
	//Deserialization of evidence
	UnmarshalEvidence([]byte) (Evidence, error)
	//Get current evidences
	Evidences() []Evidence
	//Clear all evidences
	Clear()
	Close()
}
