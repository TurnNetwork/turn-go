package process

import (
	"context"
	"crypto/ecdsa"
	"github.com/bubblenet/bubble/core/rawdb"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/datavalidator"
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
	"2488270d920dd57f0a2c7f69577b019c575480663356df4e94f90d5a1d6a72ec",
	"6b5ac12b43001b4f53f407878de9e0225bbfeb92541ed8a6e66af514bdac05ed",
	"1bf528b31769a7d1bfaf041ade586e2896a503e39e85bd0d463ef19bc45d3aa7",
	"4c2ea01368827230e8c86e27ad4f8a8f7925f32c49a70a98435e50b9e1411495",
}

func TestProcess(t *testing.T) {
	number := 3
	var sks []*ecdsa.PrivateKey
	var blsKeys []*bls.SecretKey
	for i := 0; i < number; i++ {
		sk, _ := crypto.HexToECDSA(skstr[i])
		sks = append(sks, sk)
		var key bls.SecretKey
		key.SetHexString(blsKeyStrs[i])
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

func newMockDataValidator(sk *bls.SecretKey, contract *mock.InnerContract, blockFilter *mock.BlockFilter, edb ethdb.Database, blockState types.BlockState) *datavalidator.DataValidator {
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
	return &datavalidator.DataValidator{
		Ctx:     ctx,
		Cancel:  cancel,
		Db:      vdb,
		Sync:    sync,
		Network: network,
		Process: process,
	}
}
