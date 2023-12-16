package types

import (
	"encoding/hex"
	"fmt"
	"github.com/bubblenet/bubble/common/json"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/crypto/bls"
	"sort"

	"golang.org/x/crypto/sha3"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/rlp"
)

type BlockState interface {
	BlockNumber() uint64
}

type Signature struct {
	Index     uint64
	Signature []byte
}
type signature struct {
	Index     uint64
	Signature string
}

func (m *Signature) MarshalJSON() (j []byte, err error) {
	msg := &signature{
		Index:     m.Index,
		Signature: hex.EncodeToString(m.Signature),
	}
	return json.Marshal(msg)
}

func (m *Signature) UnmarshalJSON(j []byte) (err error) {
	msg := new(signature)
	err = json.Unmarshal(j, &msg)
	if err != nil {
		return err
	}
	m.Index = msg.Index
	m.Signature, err = hex.DecodeString(msg.Signature)
	if err != nil {
		return err
	}
	return nil
}

type MessagePublishedDetail struct {
	BlockHash  common.Hash
	TxHash     common.Hash
	Log        *MessagePublished
	Signatures []*Signature
}

func (m MessagePublishedDetail) String() string {
	return fmt.Sprintf("{BlockHash:%s,TxHash:%s, Log:%s, signs:%d}", m.BlockHash.Hex(), m.TxHash.Hex(), m.Log.String(), len(m.Signatures))
}

func (m *MessagePublishedDetail) Hash() (h common.Hash) {
	return m.Log.Hash()
}

func (m *MessagePublishedDetail) Bytes() ([]byte, error) {
	return rlp.EncodeToBytes(m)
}

func (m *MessagePublishedDetail) AddSignature(sigs []*Signature) {
	signMap := make(map[uint64]struct{})
	for _, s := range m.Signatures {
		signMap[s.Index] = struct{}{}
	}
	for _, sig := range sigs {
		if _, ok := signMap[sig.Index]; !ok {
			m.Signatures = append(m.Signatures, sig)
		}
	}
	sort.Slice(m.Signatures, func(i, j int) bool {
		return m.Signatures[i].Index < m.Signatures[j].Index
	})
}

type PeerInfo struct {
	ID        string
	ScanBlock uint64
}
type MessagePublished struct {
	Send     common.Address
	ChainId  uint64
	Sequence uint64
	Nonce    uint64
	Payload  []byte
}

type messagePublished struct {
	Send     common.Address
	ChainId  uint64
	Sequence uint64
	Nonce    uint64
	Payload  string
}

func (m *MessagePublished) MarshalJSON() (j []byte, err error) {
	msg := &messagePublished{
		Send:     m.Send,
		ChainId:  m.ChainId,
		Sequence: m.Sequence,
		Nonce:    m.Nonce,
		Payload:  hex.EncodeToString(m.Payload),
	}
	return json.Marshal(msg)
}

func (m *MessagePublished) UnmarshalJSON(j []byte) (err error) {
	msg := new(messagePublished)
	err = json.Unmarshal(j, msg)
	if err != nil {
		return err
	}
	m.Send = msg.Send
	m.ChainId = msg.ChainId
	m.Sequence = msg.Sequence
	m.Nonce = msg.Nonce
	m.Payload, err = hex.DecodeString(msg.Payload)
	if err != nil {
		return err
	}
	return nil
}

func (m MessagePublished) String() string {
	return fmt.Sprintf("{Send:%s, ChainId:%d, Sequence:%d, Nonce:%d, Payload:%d}", m.Send.Hex(), m.ChainId, m.Sequence, m.Nonce, len(m.Payload))
}
func (m *MessagePublished) Hash() (h common.Hash) {
	sha := sha3.NewLegacyKeccak256().(crypto.KeccakState)
	rlp.Encode(sha, m)
	sha.Read(h[:])
	return h
}

func (m *MessagePublished) Bytes() ([]byte, error) {
	return rlp.EncodeToBytes(m)
}

type QuorumLog struct {
	Log        *MessagePublished
	Signatures []*Signature
}

func (q *QuorumLog) Bytes() ([]byte, error) {
	return rlp.EncodeToBytes(q)
}

type PeerGroup struct {
	Group     map[string]*Validator
	Threshold int
}

func NewPeerGroup(threshold int, group map[string]*Validator) *PeerGroup {
	return &PeerGroup{
		Group:     group,
		Threshold: threshold,
	}
}

func (p *PeerGroup) GetMember(index uint64) *Validator {
	for _, v := range p.Group {
		if v.Index == index {
			return v
		}
	}
	return nil
}

type ValidatorSets map[uint64]*PeerGroup

func NewValidatorSets() ValidatorSets {
	return make(ValidatorSets)
}

func (v *ValidatorSets) AddSet(chainId uint64, group *PeerGroup) {
	(*v)[chainId] = group
}

type Validator struct {
	Index       uint64
	Address     common.Address
	BlsPubKey   *bls.PublicKey
	P2pUrl      string
	BlockNumber uint64
}

type ChildChainValidator struct {
	Validators []*Validator
	Threshold  int
}
