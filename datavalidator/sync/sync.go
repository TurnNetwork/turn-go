package sync

import (
	"context"
	"github.com/bubblenet/bubble/crypto/bls"
	common2 "github.com/bubblenet/bubble/datavalidator/common"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/metrics"
	"github.com/bubblenet/bubble/p2p/discover"
	"math/big"
	"sync"
	"time"

	"github.com/bubblenet/bubble/datavalidator/utils"

	"github.com/bubblenet/bubble/datavalidator/types"
)

const (
	maxScanBlock = uint64(100)
	tickInterval = time.Second
)

var (
	filterLogsMeter = metrics.NewRegisteredMeter("datavalidator/meter/sync/filterlog", nil)
)

type SyncDB interface {
	GetScanLog() (uint64, error)
	StoreScanLog(block uint64) error
	SetMessagePublished(logs []*types.MessagePublishedDetail) error
}

type Sync struct {
	mutex              sync.Mutex
	owner              *bls.PublicKey
	validatorContract  types.ValidatorContract
	childChainContract types.ChildChainContract
	db                 SyncDB
	blockState         types.BlockState
	filter             types.FilterMessage
	newLog             chan<- []*types.MessagePublishedDetail
	newValidator       chan<- types.ValidatorSets
	chainIds           map[uint64]struct{}
}

func NewSync(owner *bls.PublicKey, validatorContract types.ValidatorContract, childChainContract types.ChildChainContract, db SyncDB, blockState types.BlockState, filter types.FilterMessage, newLog chan<- []*types.MessagePublishedDetail, newValidator chan<- types.ValidatorSets) *Sync {
	if v := validatorContract.GetValidator(owner); v != nil {
		number, _ := db.GetScanLog()
		if v.BlockNumber >= number {
			db.StoreScanLog(v.BlockNumber)
		}
	}
	return &Sync{
		owner:              owner,
		validatorContract:  validatorContract,
		childChainContract: childChainContract,
		db:                 db,
		blockState:         blockState,
		filter:             filter,
		newLog:             newLog,
		newValidator:       newValidator,
		chainIds:           make(map[uint64]struct{}),
	}
}

func (s *Sync) Run(ctx context.Context) {
	utils.Ticker(ctx, func(ctx context.Context) {
		if s.owner != nil {
			s.HandleMessage(ctx)
		}
	}, tickInterval)
}
func (s *Sync) Owner() *types.Validator {
	set := s.validatorContract.ValidatorSet()
	for _, v := range set {
		if v.BlsPubKey.IsEqual(s.owner) {
			return v
		}
	}
	return nil
}
func (s *Sync) RefreshValidator() types.ValidatorSets {
	if s.owner == nil {
		return nil
	}
	ids := s.childChainContract.GetBubbleId()
	sets := types.NewValidatorSets()
	chainIds := make(map[uint64]struct{})
	for _, chainId := range ids {
		vs := s.childChainContract.GetBubbleValidator(chainId)
		members := make(map[string]*types.Validator)
		flag := false
		for _, v := range vs.Validators {
			node, err := discover.ParseNode(v.P2pUrl)
			if err != nil {
				continue
			}
			members[common2.PeerID(node.ID.Bytes())] = v
			if v.BlsPubKey.IsEqual(s.owner) {
				flag = true
			}
		}
		if flag {
			sets.AddSet(chainId, types.NewPeerGroup(vs.Threshold, members))
			chainIds[chainId] = struct{}{}
		}
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.chainIds = chainIds
	return sets
}
func (s *Sync) HandleMessage(ctx context.Context) error {
	logs, err := s.scanMessage(ctx)
	if err != nil {
		return err
	}
	s.newLog <- logs
	return err
}
func (s *Sync) scanMessage(ctx context.Context) ([]*types.MessagePublishedDetail, error) {
	start, err := s.db.GetScanLog()
	if err != nil {
		return nil, err
	}
	log.Debug("get scan log", "start", start)
	end := s.blockState.BlockNumber()
	if end == start {
		return nil, nil
	} else if end > start+maxScanBlock {
		end = start + maxScanBlock
	}
	log.Debug("prepare scan log", "start", start, "end", end)
	logs, err := s.filter.RangeFilter(ctx, new(big.Int).SetUint64(start), new(big.Int).SetUint64(end))
	if err != nil {
		return nil, err
	}
	log.Debug("filter logs", "logs", len(logs))

	s.newValidator <- s.RefreshValidator()
	log.Debug("refresh validator success")
	var need []*types.MessagePublishedDetail
	for _, log := range logs {
		s.mutex.Lock()
		if _, ok := s.chainIds[log.Log.ChainId]; ok {
			need = append(need, log)
		}
		s.mutex.Unlock()
	}
	if need != nil {
		if err := s.db.SetMessagePublished(need); err != nil {
			return nil, err
		}
	}

	log.Debug("store scan log", "start", start, "end", end, "scan log", len(need))
	s.db.StoreScanLog(end)
	filterLogsMeter.Mark(int64(len(need)))
	return need, nil
}
