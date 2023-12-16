package datavalidator

import (
	"context"
	"github.com/bubblenet/bubble/crypto/bls"
	_ "github.com/bubblenet/bubble/datavalidator/contracts"
	"github.com/bubblenet/bubble/datavalidator/db"
	"github.com/bubblenet/bubble/datavalidator/mock"
	"github.com/bubblenet/bubble/datavalidator/p2p"
	"github.com/bubblenet/bubble/datavalidator/process"
	rpc2 "github.com/bubblenet/bubble/datavalidator/rpc"
	"github.com/bubblenet/bubble/datavalidator/sync"
	"github.com/bubblenet/bubble/datavalidator/types"
	wallet2 "github.com/bubblenet/bubble/datavalidator/wallet"
	"github.com/bubblenet/bubble/log"
	p2p2 "github.com/bubblenet/bubble/p2p"
	"github.com/bubblenet/bubble/rpc"
)

type DataValidator struct {
	Ctx     context.Context
	Cancel  context.CancelFunc
	Db      *db.DB
	Sync    *sync.Sync
	Network *p2p.Network
	Process *process.Process
}

func NewDataValidatorMock(sk *bls.SecretKey, dbPath string, blockState types.BlockState, server types.P2PServer) (*DataValidator, error) {
	vdb, err := db.NewLevelDbDataValidatorDB(dbPath)
	if err != nil {
		return nil, err
	}
	var wallet wallet2.Wallet
	if sk != nil {
		wallet = wallet2.FromSk(sk)
		log.Debug("init wallet success", "addr", wallet.PublicKey().GetHexString())
	}

	vdb.StoreScanLog(0)
	validatorContract := mock.SystemInnerContract
	childChainContract := mock.SystemInnerContract
	messagePublished := mock.SystemBlockFilter
	// validatorContract := contracts.NewInnerValidator()
	// childChainContract := contracts.NewInnerChildChainContract()
	// messagePublished := contracts.NewInnerMessagePublished()
	newLog := make(chan []*types.MessagePublishedDetail, 1)
	newValidator := make(chan types.ValidatorSets, 1)
	var owner *bls.PublicKey
	if wallet != nil {
		owner = wallet.PublicKey()
	}
	sync := sync.NewSync(owner, validatorContract, childChainContract, vdb, messagePublished, messagePublished, newLog, newValidator)

	network := p2p.NewNetwork(wallet, validatorContract, &mock.DataCheck{
		DB:            vdb,
		FilterMessage: messagePublished,
	}, blockState, server)
	process := process.NewProcess(wallet, vdb, sync, newValidator, newLog, network, network)
	ctx, cancel := context.WithCancel(context.Background())
	return &DataValidator{
		Ctx:     ctx,
		Cancel:  cancel,
		Db:      vdb,
		Sync:    sync,
		Network: network,
		Process: process,
	}, nil
}

func NewDataValidator(sk *bls.SecretKey, dbPath string, blockState types.BlockState, server types.P2PServer) (*DataValidator, error) {
	return NewDataValidatorMock(sk, dbPath, blockState, server)
	vdb, err := db.NewLevelDbDataValidatorDB(dbPath)
	if err != nil {
		return nil, err
	}
	var wallet wallet2.Wallet
	if sk != nil {
		wallet = wallet2.FromSk(sk)
	}
	innercontract := mock.NewInnerContract(nil, nil, map[uint64][]uint64{})
	blockFilter := mock.NewBlockFilter([]uint64{1})
	validatorContract := innercontract
	childChainContract := innercontract
	messagePublished := blockFilter
	// validatorContract := contracts.NewInnerValidator()
	// childChainContract := contracts.NewInnerChildChainContract()
	// messagePublished := contracts.NewInnerMessagePublished()
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
	}, blockState, server)
	process := process.NewProcess(wallet, vdb, sync, newValidator, newLog, network, network)
	ctx, cancel := context.WithCancel(context.Background())
	return &DataValidator{
		Ctx:     ctx,
		Cancel:  cancel,
		Db:      vdb,
		Sync:    sync,
		Network: network,
		Process: process,
	}, nil
}

func (d *DataValidator) Api() []rpc.API {
	return []rpc.API{rpc2.NewDataValidatorRpc(d.Db, d.Network).Api()}
}

func (d *DataValidator) P2P() []p2p2.Protocol {
	return []p2p2.Protocol{d.Network.Protocols()}
}

func (d *DataValidator) Start() error {
	d.Network.Run(d.Ctx)
	d.Process.Run(d.Ctx)
	d.Sync.Run(d.Ctx)

	return nil
}

func (d *DataValidator) Stop() error {
	d.Cancel()
	return nil
}
