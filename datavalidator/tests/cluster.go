package tests

import (
	"context"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/datavalidator"
	"github.com/bubblenet/bubble/datavalidator/db"
	"github.com/bubblenet/bubble/datavalidator/mock"
	"github.com/bubblenet/bubble/datavalidator/p2p"
	"github.com/bubblenet/bubble/datavalidator/process"
	"github.com/bubblenet/bubble/datavalidator/sync"
	"github.com/bubblenet/bubble/datavalidator/types"
	wallet2 "github.com/bubblenet/bubble/datavalidator/wallet"
	"github.com/bubblenet/bubble/ethdb"
)

func NewDataValidator(sk *bls.SecretKey, contract *mock.InnerContract, blockFilter *mock.BlockFilter, edb ethdb.Database, blockState types.BlockState) *datavalidator.DataValidator {
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
	}, blockState, &mock.P2pServer{})
	process := process.NewProcess(wallet, vdb, sync, newValidator, newLog, network, network)
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
