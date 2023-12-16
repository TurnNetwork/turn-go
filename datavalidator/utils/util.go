package utils

import (
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/types"
	dtypes "github.com/bubblenet/bubble/datavalidator/types"
	"github.com/bubblenet/bubble/rlp"
)

func DecodeMessagePublishedLog(log *types.Log) (*dtypes.MessagePublished, error) {
	var args [][]byte
	err := rlp.DecodeBytes(log.Data, &args)
	if err != nil {
		return nil, err
	}
	var (
		send     common.Address
		chainId  uint64
		sequence uint64
		nonce    uint64
		payload  []byte
	)
	rlp.DecodeBytes(args[0], &send)
	rlp.DecodeBytes(args[1], &chainId)
	rlp.DecodeBytes(args[2], &sequence)
	rlp.DecodeBytes(args[3], &nonce)
	rlp.DecodeBytes(args[4], &payload)
	return &dtypes.MessagePublished{
		Send:     send,
		ChainId:  chainId,
		Sequence: sequence,
		Nonce:    nonce,
		Payload:  payload,
	}, nil
}

func EncodeMessagePublishedData(send common.Address, chainId uint64, sequence uint64, nonce uint64, payload []byte) ([]byte, error) {
	var args [][]byte
	encode := func(v interface{}) []byte {
		buf, _ := rlp.EncodeToBytes(v)
		return buf
	}
	args = append(args, encode(send))
	args = append(args, encode(chainId))
	args = append(args, encode(sequence))
	args = append(args, encode(nonce))
	args = append(args, encode(payload))

	return encode(args), nil
}
