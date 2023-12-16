package process

import (
	"context"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/datavalidator/db"
	"github.com/bubblenet/bubble/datavalidator/p2p"
	"github.com/bubblenet/bubble/datavalidator/sync"
	"github.com/bubblenet/bubble/datavalidator/types"
	"github.com/bubblenet/bubble/datavalidator/utils"
	wallet2 "github.com/bubblenet/bubble/datavalidator/wallet"
	logger "github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/metrics"
	sync2 "sync"
	"time"
)

var (
	signMsgMeter = metrics.NewRegisteredMeter("datavalidator/meter/process/signmsg", nil)
)

type P2PSender interface {
	Send(string, interface{}) error
}

type UpdateValidatorSet interface {
	SetValidatorSet(sets types.ValidatorSets)
}

type Process struct {
	me         *types.Validator
	wallet     wallet2.Wallet
	db         *db.DB
	sets       types.ValidatorSets
	sync       *sync.Sync
	vsChan     <-chan types.ValidatorSets
	logChan    <-chan []*types.MessagePublishedDetail
	p2pSender  P2PSender
	updateVs   UpdateValidatorSet
	cacheMutex sync2.Mutex
	cache      map[common.Hash]*types.MessagePublishedDetail
}

func NewProcess(wallet wallet2.Wallet, db *db.DB, sync *sync.Sync, vsChan <-chan types.ValidatorSets, logChan <-chan []*types.MessagePublishedDetail, sender P2PSender, updateVs UpdateValidatorSet) *Process {
	p := &Process{
		wallet:    wallet,
		db:        db,
		sets:      types.NewValidatorSets(),
		sync:      sync,
		vsChan:    vsChan,
		logChan:   logChan,
		p2pSender: sender,
		updateVs:  updateVs,
		cache:     make(map[common.Hash]*types.MessagePublishedDetail),
	}
	p.me = sync.Owner()
	p.sets = sync.RefreshValidator()
	p.updateVs.SetValidatorSet(p.sets)
	p.LoadLog()
	return p
}

func (p *Process) LoadLog() {
	if p.wallet == nil {
		return
	}
	logs, _ := p.db.GetAllUnSignMessagePublished()
	for _, log := range logs {
		hash := log.Hash()
		if _, ok := p.cache[hash]; !ok {
			if len(log.Signatures) == 0 {
				p.signLog(log)
			}
			p.cache[hash] = log
		}
	}
}

func (p *Process) Run(ctx context.Context) {
	go p.handleChan(ctx)
	utils.Ticker(ctx, p.handleCache, time.Second*10)
}

func (p *Process) handleChan(ctx context.Context) {
	for {
		select {
		case s := <-p.vsChan:
			p.sets = s
			p.updateVs.SetValidatorSet(s)
		case log := <-p.logChan:
			p.handleFreshLog(log)
		case <-ctx.Done():
			return
		}
	}
}
func (p *Process) handleFreshLog(logs []*types.MessagePublishedDetail) {
	logger.Debug("handle new log", "logs", len(logs))
	if p.wallet == nil {
		return
	}
	for _, log := range logs {
		sig := p.wallet.Sign(log.Log.Hash())
		signMsgMeter.Mark(1)
		log.Signatures = append(log.Signatures, &types.Signature{Index: p.me.Index, Signature: sig})
		logger.Debug("signature success", "detail", log.String())
		p.db.SetMessagePublished([]*types.MessagePublishedDetail{&types.MessagePublishedDetail{
			BlockHash:  log.BlockHash,
			TxHash:     log.TxHash,
			Log:        log.Log,
			Signatures: []*types.Signature{&types.Signature{Index: p.me.Index, Signature: sig}},
		}})
		p.cacheMutex.Lock()
		p.cache[log.Log.Hash()] = log
		p.cacheMutex.Unlock()
		p.sendGroup(log.Log.ChainId, &p2p.SignMessageMsg{
			SignMessageData: *log,
			Signature: &types.Signature{
				Index:     p.me.Index,
				Signature: sig,
			},
		})
	}
}
func (p *Process) handleCache(ctx context.Context) {
	p.cacheMutex.Lock()
	defer p.cacheMutex.Unlock()
	for k, v := range p.cache {
		log := logger.New("logid", v.Log.Hash())
		if p.sets[v.Log.ChainId] == nil {
			log.Debug("Remove log cache", "cause", "chain id unexist")
			delete(p.cache, k)
		}
		detail, err := p.db.GetQuorumChainIdNonce(v.Log.ChainId, v.Log.Nonce)
		if err != nil {
			return
		}
		if detail != nil {
			log.Debug("Remove log cache", "cause", "had quorum")
			delete(p.cache, k)
		} else {
			log.Debug("observer request signs")
			p.sendGroup(v.Log.ChainId, &p2p.SignedObservationRequest{
				ChainId: v.Log.ChainId,
				ID:      k,
			})
		}
	}
}
func (p *Process) signLog(log *types.MessagePublishedDetail) {
	sig := p.wallet.Sign(log.Log.Hash())
	signMsgMeter.Mark(1)
	log.Signatures = append(log.Signatures, &types.Signature{Index: p.me.Index, Signature: sig})
	logger.Debug("signature success", "detail", log.String())
	p.db.SetMessagePublished([]*types.MessagePublishedDetail{&types.MessagePublishedDetail{
		BlockHash:  log.BlockHash,
		TxHash:     log.TxHash,
		Log:        log.Log,
		Signatures: []*types.Signature{&types.Signature{Index: p.me.Index, Signature: sig}},
	}})
}
func (p *Process) sendGroup(chainId uint64, msg interface{}) error {
	group := p.sets[chainId]
	if group != nil {
		for k, _ := range group.Group {
			p.p2pSender.Send(k, msg)
		}
	}
	return nil
}
