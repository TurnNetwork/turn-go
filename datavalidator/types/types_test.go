package types

import (
	"encoding/json"
	"github.com/bubblenet/bubble/common"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMessageMarshal(t *testing.T) {
	m := &MessagePublished{
		Send:     common.Address{},
		ChainId:  2,
		Sequence: 3,
		Nonce:    4,
		Payload:  []byte{15, 2, 3},
	}
	q := &QuorumLog{
		MessagePublished: m,
		Signatures: []*Signature{
			&Signature{
				Index:     0,
				Signature: []byte{12, 23},
			},
		},
	}
	value, err := json.Marshal(&q)
	t.Log(string(value))
	require.Nil(t, err)
	m2 := new(QuorumLog)
	err = json.Unmarshal(value, &m2)
	require.Nil(t, err)
	require.Equal(t, len(m.Payload), len(m2.Payload))
}
