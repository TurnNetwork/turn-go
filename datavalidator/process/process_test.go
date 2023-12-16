package process

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/bubblenet/bubble/core/rawdb"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/datavalidator/db"
	"github.com/bubblenet/bubble/datavalidator/mock"
	"github.com/bubblenet/bubble/datavalidator/p2p"
	"github.com/bubblenet/bubble/datavalidator/sync"
	"github.com/bubblenet/bubble/datavalidator/types"
	wallet2 "github.com/bubblenet/bubble/datavalidator/wallet"
	"github.com/bubblenet/bubble/ethdb"
	"github.com/bubblenet/bubble/ethdb/memorydb"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

var skstr = []string{
	"84ce2fb332ae78d78bee8cbc0bf244d0dd6e8da97bafeb191d9765ba7fe9d39b",
	"b41ac21b6b185966cabf3673f0fce9a957afa0117987af247124cbeaeb94f194",
	"78f99961967c4a39e63e7b1d18f6b8be46608c66bcd28e1331223fd3821a0ca4",
}
var blsKeyStrs = []string{
	"f958b2708c0a6eae0ea5761edcf0257526a8bbe521cc099e32adbff14b049734",
	"86e71e0d2feb7cb2233aaf084beaca51a1dfa0107d4e8ae4417555786820de1a",
	"079d678e6e61d949c60d43237c5b62d9fb2b4e71747ddd804673358e5bbd253f",
	"6ad2849adea42f1a9f981a7caa3733024c08bcba3849cc7f7ec2fe808debaf5b",
}

type DataValidator struct {
	Ctx     context.Context
	Cancel  context.CancelFunc
	Db      *db.DB
	Sync    *sync.Sync
	Network *p2p.Network
	Process *Process
}

func TestProcess(t *testing.T) {
	number := 3
	var sks []*ecdsa.PrivateKey
	var blsKeys []*bls.SecretKey
	for i := 0; i < number; i++ {
		sk, _ := crypto.HexToECDSA(skstr[i])
		sks = append(sks, sk)
		var key bls.SecretKey
		buf, _ := hex.DecodeString(blsKeyStrs[i])
		key.SetLittleEndian(buf)
		blsKeys = append(blsKeys, &key)
	}
	innercontract := mock.NewInnerContract(sks, blsKeys, map[uint64][]uint64{
		1: []uint64{0, 1, 2},
	})
	blockFilter := mock.NewBlockFilter([]uint64{1})

	mdb := memorydb.New()
	vs := newMockDataValidator(blsKeys[0], innercontract, blockFilter, rawdb.NewDatabase(mdb), blockFilter)

	blockFilter.AddMessagePublished(1, 1)
	logs, err := blockFilter.RangeFilter(context.Background(), big.NewInt(0), big.NewInt(int64(blockFilter.BlockNumber())))
	require.Nil(t, err)
	vs.Process.handleFreshLog(logs)
	unsign, err := vs.Db.GetUnSignLogRangeNonce(1, 0, 10)
	require.Nil(t, err)
	require.Equal(t, 1, len(unsign))
	require.Equal(t, 1, len(unsign[0].Signatures))
	require.Equal(t, 1, len(vs.Process.cache))

	vs.Db.SetQuorumLog(logs)
	unsign, err = vs.Db.GetUnSignLogRangeNonce(1, 0, 10)
	require.Nil(t, err)
	require.Equal(t, 0, len(unsign))
	signLogs, err := vs.Db.GetQuorumLogRangeNonce(1, 0, 10)
	require.Nil(t, err)
	require.Equal(t, 1, len(signLogs))
	vs.Process.handleCache(context.Background())
	require.Equal(t, 0, len(vs.Process.cache))

}

func newMockDataValidator(sk *bls.SecretKey, contract *mock.InnerContract, blockFilter *mock.BlockFilter, edb ethdb.Database, blockState types.BlockState) *DataValidator {
	vdb := db.NewDataValidatorDB(edb)
	var wallet wallet2.Wallet
	if sk != nil {
		wallet = wallet2.FromSk(sk)
	}
	validatorContract := contract
	childChainContract := contract
	messagePublished := blockFilter
	newLog := make(chan []*types.MessagePublishedDetail, 1)
	newValidator := make(chan types.ValidatorSets, 1)
	var owner *bls.PublicKey
	if wallet != nil {
		owner = wallet.PublicKey()
	}
	sync := sync.NewSync(owner, validatorContract, childChainContract, vdb, blockState, messagePublished, newLog, newValidator)

	network := p2p.NewNetwork(wallet, validatorContract, &mock.DataCheck{
		DB:            vdb,
		FilterMessage: messagePublished,
	}, blockState, nil)
	process := NewProcess(wallet, vdb, sync, newValidator, newLog, network, network)
	ctx, cancel := context.WithCancel(context.Background())
	return &DataValidator{
		Ctx:     ctx,
		Cancel:  cancel,
		Db:      vdb,
		Sync:    sync,
		Network: network,
		Process: process,
	}
}
