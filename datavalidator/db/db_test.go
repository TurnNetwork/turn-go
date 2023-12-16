package db

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/datavalidator/types"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestDB(t *testing.T) {
	data := []*types.MessagePublishedDetail{
		&types.MessagePublishedDetail{
			BlockHash: common.BigToHash(big.NewInt(1)),
			TxHash:    common.BigToHash(big.NewInt(2)),
			Log: &types.MessagePublished{
				Send:     common.BigToAddress(big.NewInt(3)),
				ChainId:  1,
				Sequence: 1,
				Nonce:    1,
				Payload:  nil,
			},
			Signatures: []*types.Signature{
				&types.Signature{Index: 1, Signature: []byte{1, 2, 3}},
			},
		},
		&types.MessagePublishedDetail{
			BlockHash: common.BigToHash(big.NewInt(1)),
			TxHash:    common.BigToHash(big.NewInt(2)),
			Log: &types.MessagePublished{
				Send:     common.BigToAddress(big.NewInt(3)),
				ChainId:  1,
				Sequence: 2,
				Nonce:    2,
				Payload:  nil,
			},
			Signatures: []*types.Signature{
				&types.Signature{Index: 1, Signature: []byte{1, 2, 3}},
			},
		},
		&types.MessagePublishedDetail{
			BlockHash: common.BigToHash(big.NewInt(1)),
			TxHash:    common.BigToHash(big.NewInt(2)),
			Log: &types.MessagePublished{
				Send:     common.BigToAddress(big.NewInt(3)),
				ChainId:  1,
				Sequence: 3,
				Nonce:    13,
				Payload:  nil,
			},
			Signatures: nil,
		},
		&types.MessagePublishedDetail{
			BlockHash: common.BigToHash(big.NewInt(1)),
			TxHash:    common.BigToHash(big.NewInt(2)),
			Log: &types.MessagePublished{
				Send:     common.BigToAddress(big.NewInt(3)),
				ChainId:  1,
				Sequence: 4,
				Nonce:    23,
				Payload:  nil,
			},
			Signatures: []*types.Signature{
				&types.Signature{Index: 1, Signature: []byte{1, 2, 3}},
			},
		},
		&types.MessagePublishedDetail{
			BlockHash: common.BigToHash(big.NewInt(1)),
			TxHash:    common.BigToHash(big.NewInt(2)),
			Log: &types.MessagePublished{
				Send:     common.BigToAddress(big.NewInt(3)),
				ChainId:  2,
				Sequence: 5,
				Nonce:    2,
				Payload:  nil,
			},
			Signatures: []*types.Signature{
				&types.Signature{Index: 1, Signature: []byte{1, 2, 3}},
			},
		},
	}
	mdb := NewMemoryValidatorDB()
	mdb.SetMessagePublished(data)
	logs, err := mdb.GetAllMessagePublished()
	require.Nil(t, err)
	require.Equal(t, len(data), len(logs))
	logs, err = mdb.GetUnSignLogRangeNonce(1, 2, 10)
	require.Nil(t, err)
	for _, l := range logs {
		require.True(t, l.Log.Nonce >= 2)
	}
	logs, err = mdb.GetUnSignLogByTxHash((common.BigToHash(big.NewInt(2))))
	require.Nil(t, err)
	for _, l := range logs {
		require.Equal(t, common.BigToHash(big.NewInt(2)), l.TxHash)
	}
	mdb.SetQuorumLog(data)
	logs, err = mdb.GetQuorumLogRangeNonce(1, 2, 10)
	require.Nil(t, err)
	for _, l := range logs {
		require.True(t, l.Log.Nonce >= 2)
	}

	logs, err = mdb.GetQuorumLogByTxHash(common.BigToHash(big.NewInt(2)))
	require.Nil(t, err)
	for _, l := range logs {
		require.Equal(t, common.BigToHash(big.NewInt(2)), l.TxHash)
	}
}
