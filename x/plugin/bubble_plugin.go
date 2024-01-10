package plugin

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/json"
	"github.com/bubblenet/bubble/common/math"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core/snapshotdb"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/ethclient"
	"github.com/bubblenet/bubble/event"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/params"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/bubble"
	"github.com/bubblenet/bubble/x/gov"
	"github.com/bubblenet/bubble/x/stakingL2"
	"github.com/bubblenet/bubble/x/xcom"
	"github.com/bubblenet/bubble/x/xcom/vrf"
	"github.com/bubblenet/bubble/x/xutil"
	"golang.org/x/crypto/sha3"
	gomath "math"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	SubChainSysAddr  = "0x1000000000000000000000000000000000000020" // Sub chain system contract address
	remoteBubbleAddr = "0x2000000000000000000000000000000000000001"
	bubbleLife       = 172800
)

var (
	bubblePluginOnce sync.Once
	bubblePlugin     *BubblePlugin
	preReleaseLife   uint64
	releaseLife      uint64
)

type BubblePlugin struct {
	stkPlugin  *StakingPlugin
	stk2Plugin *StakingL2Plugin
	db         *bubble.DB
	NodeID     discover.NodeID // id of the local node
	eventMux   *event.TypeMux
	opPriKey   string // Main chain operation address private key
}

// BubbleInstance instance a global BubblePlugin
func BubbleInstance() *BubblePlugin {
	bubblePluginOnce.Do(func() {
		log.Info("Init bubble plugin ...")
		preReleaseLife = uint64(gomath.Ceil(float64(bubbleLife)/float64(xutil.CalcBlocksEachEpoch()))) * xutil.CalcBlocksEachEpoch()
		releaseLife = uint64(gomath.Ceil(float64(bubbleLife)*0.5/float64(xutil.CalcBlocksEachEpoch()))) * xutil.CalcBlocksEachEpoch()
		bubblePlugin = &BubblePlugin{
			stkPlugin:  StakingInstance(),
			stk2Plugin: StakingL2Instance(),
			db:         bubble.NewDB(),
		}
	})
	return bubblePlugin
}

func (bp *BubblePlugin) BeginBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	return nil
}

func (bp *BubblePlugin) EndBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	curBlock := header.Number.Uint64()
	preBlock := curBlock + 20

	bubs, err := bp.GetBubbles(blockHash)
	if err != nil {
		return err
	}

	if xutil.IsEndOfEpoch(curBlock) {
		for _, bub := range bubs {

			if curBlock < bub.PreReleaseBlock {
				continue
			}
			// preRelease bubble
			if curBlock < bub.ReleaseBlock {
				if bub.State < bubble.PreReleaseState {
					bub.State = bubble.PreReleaseState
					if err := bp.db.StoreBasicsInfo(blockHash, bub.BubbleId, bub); err != nil {
						log.Error("Failed to store bubble on BubblePlugin EndBlock",
							"blockNumber", curBlock, "blockHash", blockHash.Hex(), "bubble", bub.BubbleId, "err", err.Error())
						return err
					}
				}
				if bub.ContractCount > 0 {
					continue
				}
			}
			if bub.State == bubble.ReleasedState {
				continue
			}
			// prerelease and not contract OR current block is release block
			log.Debug("start to ReleaseBubble", "curBlock", curBlock, "bubbleId", bub.BubbleId, "bubbleState", bub.State, "releaseBlock", bub.ReleaseBlock,
				"contractCount", bub.ContractCount)
			err := bp.ReleaseBubble(blockHash, header.Number, bub.BubbleId)
			if err != nil {
				log.Error("Failed to release bubble on BubblePlugin EndBlock",
					"blockNumber", curBlock, "blockHash", blockHash.Hex(), "bubble", bub.BubbleId, "err", err.Error())
				return err
			}
		}
	}

	// destroy and clean bubble
	if xutil.IsEndOfEpoch(preBlock) {
		for _, bub := range bubs {
			log.Debug("prepare to RemoteDestroy", "bubbleID", bub.BubbleId, "State", bub.State, "curBlock", curBlock, "preBlock", preBlock,
				"PreReleaseBlock", bub.PreReleaseBlock, "releaseBlock", bub.ReleaseBlock)
			if bub.State == bubble.PreReleaseState && preBlock == bub.ReleaseBlock {
				if err := bp.DestroyBubble(blockHash, curBlock, bub.BubbleId); err != nil {
					log.Error("Failed to destroy bubble on BubblePlugin EndBlock",
						"blockNumber", curBlock, "blockHash", blockHash.Hex(), "bubble", bub.BubbleId, "err", err.Error())
					return err
				}
			}
		}
	}

	return nil
}

func (bp *BubblePlugin) Confirmed(nodeId discover.NodeID, block *types.Block) error {
	return nil
}

func (bp *BubblePlugin) SetCurrentNodeID(nodeId discover.NodeID) {
	bp.NodeID = nodeId
}

func (bp *BubblePlugin) SetEventMux(eventMux *event.TypeMux) {
	bp.eventMux = eventMux
}

func (bp *BubblePlugin) SetOpPriKey(opPriKey string) error {
	if "0x" == opPriKey[0:2] || "0X" == opPriKey[0:2] {
		opPriKey = opPriKey[2:]
	}
	if 64 != len(opPriKey) {
		return errors.New("the private key is of the wrong size")
	}
	bp.opPriKey = opPriKey
	return nil
}

func (bp *BubblePlugin) GetNodeUseRatio(blockHash common.Hash) (float32, error) {
	used, err := bp.stk2Plugin.db.GetUsedCommitteeCount(blockHash)
	if err != nil {
		return 0, err
	}

	total, err := bp.stk2Plugin.db.GetCommitteeCount(blockHash)
	if err != nil {
		return 0, err
	}

	return float32(used) / float32(total), nil
}

// GetBubbles return the all bubble information
func (bp *BubblePlugin) GetBubbles(blockHash common.Hash) ([]*bubble.BasicsInfo, error) {
	iter := bp.db.IteratorBasicsInfo(blockHash, 0)
	if err := iter.Error(); nil != err {
		return nil, err
	}
	defer iter.Release()

	var infos []*bubble.BasicsInfo
	for iter.Valid(); iter.Next(); {
		data := iter.Value()
		stateInfo := new(bubble.BasicsInfo)
		if err := rlp.DecodeBytes(data, stateInfo); err != nil {
			return nil, err
		}
		infos = append(infos, stateInfo)
	}

	return infos, nil
}

// GetBubbleInfo return the bubble information by bubble ID
func (bp *BubblePlugin) GetBubbleInfo(blockHash common.Hash, bubbleID *big.Int) (*bubble.BasicsInfo, error) {
	return bp.db.GetBasicsInfo(blockHash, bubbleID)
}

func (bp *BubblePlugin) SetBubbleInfo(blockHash common.Hash, bubble *bubble.BasicsInfo) error {
	return bp.db.StoreBasicsInfo(blockHash, bubble.BubbleId, bubble)
}

// GetValidatorInfo return the bubble Validators by bubble ID
func (bp *BubblePlugin) GetValidatorInfo(blockHash common.Hash, bubbleID *big.Int) (*bubble.ValidatorInfo, error) {
	validator, err := bp.db.GetValidatorInfo(blockHash, bubbleID)
	if err != nil || validator == nil {
		return nil, err
	}

	// update operators rpc url
	nodeId := validator.OperatorsL2[0].NodeId
	canAddr, err := xutil.NodeId2Addr(nodeId)
	if nil != err {
		return nil, err
	}
	base, err := bp.stk2Plugin.db.GetCanBaseStore(blockHash, canAddr)
	if err != nil {
		return nil, err
	}
	// When the rpc url in the micro-node information is not empty, it will be updated
	if "" != base.RPCURI {
		validator.OperatorsL2[0].RPC = base.RPCURI
	}
	// update nodes rpc url
	for _, node := range validator.MicroNodes {
		addr, err := xutil.NodeId2Addr(node.NodeId)
		if nil != err {
			return nil, err
		}
		base, err := bp.stk2Plugin.db.GetCanBaseStore(blockHash, addr)
		if err != nil {
			return nil, err
		}
		// When the rpc url in the micro-node information is not empty, it will be updated
		if base.RPCURI != "" {
			node.RPCURI = base.RPCURI
		}
	}

	return validator, nil
}

func (bp *BubblePlugin) CheckBubbleElements(blockHash common.Hash, size bubble.Size) error {
	bubbleSize, err := bubble.GetConfig(size)
	if err != nil {
		return err
	}
	// check L1 operators
	if operators, err := bp.stkPlugin.db.GetOperatorArrStore(blockHash); err != nil || len(operators) < int(bubbleSize.OperatorL1Size) {
		return bubble.ErrOperatorL1IsInsufficient
	}
	// check L2 operators
	if operators, err := bp.stk2Plugin.GetOperatorList(blockHash); err != nil || len(operators) < int(bubbleSize.OperatorL2Size) {
		return bubble.ErrOperatorL2IsInsufficient
	}
	// check L2 committees
	if committees, err := bp.stk2Plugin.GetCommitteeList(blockHash); err != nil || len(committees) < int(bubbleSize.CommitteeSize) {
		return bubble.ErrMicroNodeIsInsufficient
	}

	return nil
}

// CreateBubble run the non-business logic to create bubble
func (bp *BubblePlugin) CreateBubble(blockHash common.Hash, blockNumber *big.Int, txHash common.Hash, from common.Address, nonce uint64, parentNonce [][]byte, size bubble.Size) (*bubble.BasicsInfo, error) {
	bubbleSize, err := bubble.GetConfig(size)
	if err != nil {
		return nil, err
	}

	// elect the operatorsL1 by VRF
	OperatorsL1, err := bp.ElectOperatorL1(blockHash, bubbleSize.OperatorL1Size, common.Uint64ToBytes(nonce), parentNonce)
	if err != nil {
		return nil, err
	}

	// elect the operatorsL2 by VRF
	candidateL2, err := bp.ElectOperatorL2(blockHash, bubbleSize.OperatorL2Size, common.Uint64ToBytes(nonce), parentNonce)
	if err != nil {
		return nil, err
	}
	var OperatorsL2 []*bubble.Operator
	for _, can := range candidateL2 {
		// RPCURI is updating
		canAddr, err := xutil.NodeId2Addr(can.NodeId)
		if err != nil {
			return nil, err
		}
		canBase, err := bp.stk2Plugin.GetCanBase(blockHash, canAddr)

		operator := &bubble.Operator{
			NodeId: canBase.NodeId,
			RPC:    canBase.RPCURI,
			OpAddr: canBase.StakingAddress,
		}
		OperatorsL2 = append(OperatorsL2, operator)
	}

	// elect the microNodesL2 by VRF
	microNodes, err := bp.ElectBubbleMicroNodes(blockHash, bubbleSize.CommitteeSize, common.Uint64ToBytes(nonce), parentNonce)
	if err != nil {
		return nil, err
	}

	if err := bp.stk2Plugin.db.AddUsedCommitteeCount(blockHash, uint32(len(microNodes))); err != nil {
		return nil, err
	}

	microNodes = append(microNodes, candidateL2...)

	// build bubble infos
	bubbleID, err := bp.generateBubbleID(from, big.NewInt(int64(nonce)), microNodes)
	if err != nil {
		return nil, err
	}
	if data, _ := bp.GetBubbleInfo(blockHash, bubbleID); data != nil {
		return nil, errors.New(fmt.Sprintf("bubble %d already exist", bubbleID))
	}

	baseBlock := uint64(gomath.Ceil(float64(blockNumber.Uint64())/float64(xutil.CalcBlocksEachEpoch()))) * xutil.CalcBlocksEachEpoch()

	bub := &bubble.BasicsInfo{
		BubbleId:        bubbleID,
		Size:            size,
		State:           bubble.ActiveState,
		CreateBlock:     blockNumber.Uint64(),
		PreReleaseBlock: baseBlock + preReleaseLife,
		ReleaseBlock:    baseBlock + preReleaseLife + releaseLife,
		ContractCount:   0,
	}

	validator := &bubble.ValidatorInfo{
		BubbleId:    bubbleID,
		OperatorsL1: OperatorsL1,
		OperatorsL2: OperatorsL2,
		MicroNodes:  microNodes,
	}

	// store bubble basics
	if err := bp.db.StoreBasicsInfo(blockHash, bub.BubbleId, bub); err != nil {
		log.Error("Failed to CreateBubble on bubblePlugin: Store bubble basics failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleId", bub.BubbleId, "err", err.Error())
		return nil, err
	}

	// store bubble state
	if err := bp.db.StoreValidatorInfo(blockHash, validator.BubbleId, validator); err != nil {
		log.Error("Failed to CreateBubble on bubblePlugin: Store bubble validator failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleId", validator.BubbleId, "err", err.Error())
		return nil, err
	}

	if err := bp.db.StoreBubbleIdBySize(blockHash, size, bub.BubbleId); err != nil {
		log.Error("Failed to CreateBubble on bubblePlugin: Store bubble sized info failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleId", bub.BubbleId, "err", err.Error())
		return nil, err
	}

	// send create bubble event to the blockchain Mux if local node is operator
	task := &bubble.CreateBubbleTask{
		BubbleID: bub.BubbleId,
		TxHash:   txHash,
	}

	for _, operators := range validator.OperatorsL1 {
		if operators.NodeId == bp.NodeID {
			if err := bp.PostCreateBubbleEvent(task); err != nil {
				return nil, err
			}
		}
	}

	return bub, nil
}

// generateBubbleID generate bubble ID use sha3 algorithm by bubble info
func (bp *BubblePlugin) generateBubbleID(creator common.Address, nonce *big.Int, committer bubble.ValidatorQueue) (*big.Int, error) {
	committerData, err := rlp.EncodeToBytes(committer)
	if err != nil {
		return nil, err
	}
	data := bytes.Join([][]byte{creator.Bytes(), nonce.Bytes(), committerData}, []byte(""))
	hash := sha3.Sum256(data)

	return new(big.Int).SetBytes(hash[:4]), nil
}

// ElectOperatorL1 Elect the Layer1 Operator nodes for the bubble chain by VRF
func (bp *BubblePlugin) ElectOperatorL1(blockHash common.Hash, operatorNumber uint, curNonce []byte, preNonces [][]byte) ([]*bubble.Operator, error) {
	operators, err := bp.stkPlugin.db.GetOperatorArrStore(blockHash)
	if err != nil {
		return nil, err
	}
	if len(operators) < int(operatorNumber) || len(operators) == 0 {
		return nil, bubble.ErrOperatorL1IsInsufficient
	}

	// fill empty nonce preNonces to if nonce insufficient
	if len(preNonces) < len(operators) {
		newNonces := make([][]byte, len(operators))
		newNonces = append(newNonces, preNonces...)
		preNonces = newNonces
	}

	// wrap the operators to the VRF able queue
	vrfQueue, err := vrf.VRFQueueWrapper(operators, func(item interface{}) *vrf.VRFItem {
		w, _ := new(big.Int).SetString("1000000000000000000000000", 10)
		return &vrf.VRFItem{
			V: item,
			W: w,
		}
	})
	if err != nil {
		return nil, err
	}

	// VRF Elect
	log.Info("ElectOperatorL1 run VRF", "vrfQueue len", len(vrfQueue), "preNonces len", len(preNonces))
	electedVrfQueue, err := vrf.VRF(vrfQueue, operatorNumber, curNonce, preNonces[:len(vrfQueue)])
	if err != nil {
		return nil, err
	}

	// unwrap the VRF able queue
	electedOperators := make([]*bubble.Operator, 0)
	for _, item := range electedVrfQueue {
		if operator, ok := (item.V).(*bubble.Operator); ok {
			electedOperators = append(electedOperators, operator)
		} else {
			return nil, errors.New("type error")
		}
	}

	return electedOperators, nil
}

// ElectOperatorL2 Elect the Layer2 Operator nodes for the bubble chain by VRF
func (bp *BubblePlugin) ElectOperatorL2(blockHash common.Hash, operatorNumber uint, curNonce []byte, preNonces [][]byte) (bubble.ValidatorQueue, error) {
	operatorQueue, err := bp.stk2Plugin.GetOperatorList(blockHash)
	if err != nil {
		return nil, err
	}
	if len(operatorQueue) < int(operatorNumber) || len(operatorQueue) == 0 {
		return nil, bubble.ErrOperatorL2IsInsufficient
	}

	// fill empty nonce preNonces to if nonce insufficient
	if len(preNonces) < len(operatorQueue) {
		newNonces := make([][]byte, len(operatorQueue))
		newNonces = append(newNonces, preNonces...)
		preNonces = newNonces
	}

	// wrap the operators to the VRF able queue
	vrfQueue, err := vrf.VRFQueueWrapper(operatorQueue, func(item interface{}) *vrf.VRFItem {
		if candidate, ok := (item).(*stakingL2.Candidate); ok {
			return &vrf.VRFItem{
				V: item,
				W: candidate.Shares,
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// VRF Elect
	log.Info("ElectOperatorL2 run VRF", "vrfQueue len", len(vrfQueue), "preNonces len", len(preNonces))
	electedVrfQueue, err := vrf.VRF(vrfQueue, operatorNumber, curNonce, preNonces[:len(vrfQueue)])
	if err != nil {
		return nil, err
	}

	// unwrap the VRF able queue
	electedOperators := make(bubble.ValidatorQueue, 0)
	for _, item := range electedVrfQueue {
		if Operator, ok := (item.V).(*stakingL2.Candidate); ok {
			electedOperators = append(electedOperators, Operator)
		} else {
			return nil, errors.New("type error")
		}
	}

	// delete the elected operators from the operator list
	for _, operator := range electedOperators {
		addr, err := xutil.NodeId2Addr(operator.NodeId)
		if err != nil {
			return nil, err
		}

		if err := bp.stk2Plugin.db.DelOperatorStore(blockHash, addr); err != nil {
			return nil, err
		}
	}

	return electedOperators, nil
}

// ElectBubbleMicroNodes Elect the Committee nodes for the bubble chain by VRF
func (bp *BubblePlugin) ElectBubbleMicroNodes(blockHash common.Hash, committeeNumber uint, curNonce []byte, preNonces [][]byte) (bubble.ValidatorQueue, error) {
	committeeQueue, err := bp.stk2Plugin.GetCommitteeList(blockHash)
	if err != nil {
		return nil, err
	}
	if len(committeeQueue) < int(committeeNumber) || len(committeeQueue) == 0 {
		return nil, bubble.ErrMicroNodeIsInsufficient
	}

	// fill empty nonce preNonces to if nonce insufficient
	if len(preNonces) < len(committeeQueue) {
		newNonces := make([][]byte, len(committeeQueue))
		newNonces = append(newNonces, preNonces...)
		preNonces = newNonces
	}

	// wrap the candidates to the VRF able queue
	vrfQueue, err := vrf.VRFQueueWrapper(committeeQueue, func(item interface{}) *vrf.VRFItem {
		if candidate, ok := (item).(*stakingL2.Candidate); ok {
			return &vrf.VRFItem{
				V: item,
				W: candidate.Shares,
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// VRF Elect
	log.Info("ElectBubbleMicroNodes run VRF", "vrfQueue len", len(vrfQueue), "preNonces len", len(preNonces))
	electedVrfQueue, err := vrf.VRF(vrfQueue, committeeNumber, curNonce, preNonces[:len(vrfQueue)])
	if err != nil {
		return nil, err
	}

	// unwrap the VRF able queue
	committees := make(bubble.ValidatorQueue, 0)
	for _, item := range electedVrfQueue {
		if candidate, ok := (item.V).(*stakingL2.Candidate); ok {
			committees = append(committees, candidate)
		} else {
			return nil, errors.New("type error")
		}
	}

	// delete the elected candidates from the candidate list
	for _, committee := range committees {
		addr, err := xutil.NodeId2Addr(committee.NodeId)
		if err != nil {
			return nil, err
		}
		if err := bp.stk2Plugin.db.DelCommitteeStore(blockHash, addr); err != nil {
			return nil, err
		}
	}

	return committees, nil
}

func (bp *BubblePlugin) GetSizedBubbleIDs(blockHash common.Hash, size bubble.Size) ([]*big.Int, error) {
	iter := bp.db.IteratorBubbleIdBySize(blockHash, size, 0)
	if err := iter.Error(); nil != err {
		return nil, err
	}
	defer iter.Release()

	queue := make([]*big.Int, 0)
	for iter.Valid(); iter.Next(); {
		data := iter.Value()
		bubID := new(big.Int)
		if err := rlp.DecodeBytes(data, bubID); err != nil {
			return nil, err
		}
		queue = append(queue, bubID)
	}

	return queue, nil
}

func (bp *BubblePlugin) ElectBubble(blockHash common.Hash, nonce uint64, preNonces [][]byte, size bubble.Size) (*big.Int, error) {
	bubIDs, err := bp.GetSizedBubbleIDs(blockHash, size)
	if err != nil {
		return nil, err
	}

	if len(preNonces) < len(bubIDs) {
		fitNonces := make([][]byte, len(bubIDs))
		fitNonces = append(fitNonces, preNonces...)
		preNonces = fitNonces
	}

	vrfQueue, err := vrf.VRFQueueWrapper(bubIDs, func(item interface{}) *vrf.VRFItem {
		w, _ := new(big.Int).SetString("1200000000000000000000000", 10)
		return &vrf.VRFItem{
			V: item,
			W: w,
		}
	})

	electedVrfQueue, err := vrf.VRF(vrfQueue, 1, common.Uint64ToBytes(nonce), preNonces[:len(vrfQueue)])
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if len(electedVrfQueue) != 1 {
		return nil, errors.New("the number of bubbles elected is not correct")
	}

	// unwrap the VRF able queue
	vrfItem := electedVrfQueue[0]
	if bubID, ok := (vrfItem.V).(*big.Int); ok {
		return bubID, nil
	} else {
		return nil, errors.New("type error")
	}
}

func (bp *BubblePlugin) DestroyBubble(blockHash common.Hash, blockNumber uint64, bubbleID *big.Int) error {
	val, err := bp.GetValidatorInfo(blockHash, bubbleID)
	if err != nil {
		return err
	}

	task := &bubble.RemoteDestroyTask{
		BlockNumber: blockNumber,
		BubbleID:    bubbleID,
		RPC:         val.OperatorsL2[0].RPC,
		OpAddr:      val.OperatorsL1[0].OpAddr,
	}

	for _, operators := range val.OperatorsL1 {
		if operators.NodeId == bp.NodeID {
			if err := bp.PostRemoteDestroyEvent(task); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReleaseBubble run the non-business logic to release the bubble
func (bp *BubblePlugin) ReleaseBubble(blockHash common.Hash, blockNumber *big.Int, bubbleID *big.Int) error {
	bub, err := bp.GetBubbleInfo(blockHash, bubbleID)
	if err != nil {
		log.Error("Failed to GetBubbleInfo on ReleaseBubble", "blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleID",
			bub.BubbleId.String(), "err", err.Error())
		return err
	}

	if bub.State <= bubble.ReleasedState {
		bub.State = bubble.ReleasedState
	} else {
		log.Error("bubble is already released")
		return errors.New("bubble is already released")
	}

	// release the committee nodes to the DB
	var committeeCount uint32
	val, err := bp.GetValidatorInfo(blockHash, bubbleID)
	if err != nil {
		log.Error("Failed to GetValidatorInfo on ReleaseBubble", "blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleID",
			bub.BubbleId.String(), "err", err.Error())
		return err
	}

	for _, microNode := range val.MicroNodes {
		addr, err := xutil.NodeId2Addr(microNode.NodeId)
		if err != nil {
			return err
		}
		// check whether the node has been withdrawn
		can, err := bp.stk2Plugin.db.GetCandidateStore(blockHash, addr)
		if can == nil || err != nil {
			break
		}

		if microNode.IsOperator {
			// release operator node
			Operator, _ := bp.stk2Plugin.db.GetOperatorStore(blockHash, addr)
			if Operator != nil {
				log.Error("Failed to SetOperatorStore on ReleaseBubble: Operator info is exist",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", microNode.NodeId.String())
				return errors.New("failed to SetOperatorStore on ReleaseBubble: Operator info is exist")
			}

			if err := bp.stk2Plugin.db.SetOperatorStore(blockHash, addr, can); nil != err {
				log.Error("Failed to SetOperatorStore on ReleaseBubble: Store Operator info is failed",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", microNode.NodeId.String(), "err", err.Error())
				return err
			}
		} else {
			// release committee node
			committeeCount += 1
			Committee, _ := bp.stk2Plugin.db.GetCommitteeStore(blockHash, addr)
			if Committee != nil {
				log.Error("Failed to SetCommitteeStore on ReleaseBubble: Committee info is exist",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", microNode.NodeId.String())
				return errors.New("failed to SetCommitteeStore on ReleaseBubble: Committee info is exist")
			}

			if err := bp.stk2Plugin.db.SetCommitteeStore(blockHash, addr, can); nil != err {
				log.Error("Failed to SetCommitteeStore on ReleaseBubble: Store Committee info is failed",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", microNode.NodeId.String(), "err", err.Error())
				return err
			}
		}
	}

	if err := bp.db.StoreBasicsInfo(blockHash, bubbleID, bub); err != nil {
		log.Error("Failed to StoreBasicsInfo on ReleaseBubble",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleID", bub.BubbleId.String(), "err", err.Error())
		return err
	}

	if err := bp.stk2Plugin.db.SubUsedCommitteeCount(blockHash, committeeCount); err != nil {
		log.Error("Failed to SubUsedCommitteeCount on ReleaseBubble",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleID", bub.BubbleId.String(), "err", err.Error())
		return err
	}

	if err := bp.db.DelBubbleIdBySize(blockHash, bub.Size, bub.BubbleId); err != nil {
		log.Error("Failed to DelBubbleIdBySize on ReleaseBubble",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleID", bub.BubbleId.String(), "err", err.Error())
		return err
	}

	// send release bubble event to the blockchain Mux if local node is operator
	task := &bubble.ReleaseBubbleTask{
		BubbleID: val.BubbleId,
	}

	for _, operators := range val.OperatorsL1 {
		if operators.NodeId == bp.NodeID {
			if err := bp.PostReleaseBubbleEvent(task); err != nil {
				return err
			}
		}
	}

	return nil
}

func (bp *BubblePlugin) GetByteCode(blockHash common.Hash, address common.Address) ([]byte, error) {
	return bp.db.GetByteCode(blockHash, address)
}

func (bp *BubblePlugin) StoreByteCode(blockHash common.Hash, address common.Address, byteCode []byte) error {
	return bp.db.StoreByteCode(blockHash, address, byteCode)
}

func (bp *BubblePlugin) GetBubContracts(blockHash common.Hash, bubbleID *big.Int) ([]*bubble.ContractInfo, error) {
	iter := bp.db.IteratorContractInfo(blockHash, bubbleID, 0)
	if err := iter.Error(); nil != err {
		return nil, err
	}
	defer iter.Release()

	queue := make([]*bubble.ContractInfo, 0)
	for iter.Valid(); iter.Next(); {
		data := iter.Value()
		contractInfo := new(bubble.ContractInfo)
		err := rlp.DecodeBytes(data, contractInfo)
		if err != nil {
			return nil, err
		}

		queue = append(queue, contractInfo)
	}

	return queue, nil
}

func (bp *BubblePlugin) GetBubContract(blockHash common.Hash, bubbleID *big.Int, address common.Address) (*bubble.ContractInfo, error) {
	return bp.db.GetContractInfo(blockHash, bubbleID, address)
}

func (bp *BubblePlugin) StoreBubContract(blockHash common.Hash, bubbleID *big.Int, contractInfo *bubble.ContractInfo) error {
	return bp.db.StoreContractInfo(blockHash, bubbleID, contractInfo)
}

func (bp *BubblePlugin) DelBubContract(blockHash common.Hash, bubbleID *big.Int, address common.Address) error {
	return bp.db.DelContractInfo(blockHash, bubbleID, address)
}

// GetAccListOfBub Get the list of accounts inside bubble
// An account is activated within a bubble by staking tokens with a specified bubbleId
func (bp *BubblePlugin) GetAccListOfBub(blockHash common.Hash, bubbleId *big.Int) ([]common.Address, error) {
	return bp.db.GetAccListOfBub(blockHash, bubbleId)
}

// AddAccToBub Add the account address of the staking tokens to bubble
func (bp *BubblePlugin) AddAccToBub(blockHash common.Hash, bubbleId *big.Int, account common.Address) error {
	accList, err := bp.GetAccListOfBub(blockHash, bubbleId)
	if snapshotdb.NonDbNotFoundErr(err) {
		return err
	}
	// Store it in a list of accounts in bubble
	accList = append(accList, account)
	return bp.db.StoreAccListOfBub(blockHash, bubbleId, accList)
}

// GetAccAssetOfBub Get the assets staking by the account within the specified bubble
func (bp *BubblePlugin) GetAccAssetOfBub(blockHash common.Hash, bubbleId *big.Int, account common.Address) (*bubble.AccountAsset, error) {
	return bp.db.GetAccAssetOfBub(blockHash, bubbleId, account)
}

// StoreAccAssetToBub Store the information of the staking assets of the account into bubble
func (bp *BubblePlugin) StoreAccAssetToBub(blockHash common.Hash, bubbleId *big.Int, stakingAsset *bubble.AccountAsset) error {
	if nil == stakingAsset {
		return errors.New("null pointer")
	}
	return bp.db.StoreAccAssetToBub(blockHash, bubbleId, *stakingAsset)
}

// GetL1HashByL2Hash The transaction hash of the main chain is queried according to the transaction hash of the child chain
func (bp *BubblePlugin) GetL1HashByL2Hash(blockHash common.Hash, bubbleID *big.Int, L2TxHash common.Hash) (*common.Hash, error) {
	return bp.db.GetL1HashByL2Hash(blockHash, bubbleID, L2TxHash)
}

// StoreL2HashToL1Hash The mapping relationship between the sub-chain transaction hash and the main chain transaction hash is stored
func (bp *BubblePlugin) StoreL2HashToL1Hash(blockHash common.Hash, bubbleID *big.Int, L1TxHash common.Hash, L2TxHash common.Hash) error {
	return bp.db.StoreL2HashToL1Hash(blockHash, bubbleID, L1TxHash, L2TxHash)
}

// GetTxHashListByBub The mapping relationship between the sub-chain transaction hash and the main chain transaction hash is stored
func (bp *BubblePlugin) GetTxHashListByBub(blockHash common.Hash, bubbleID *big.Int, txType bubble.TxType) ([]common.Hash, error) {
	txHashList, err := bp.db.GetTxHashListByBub(blockHash, bubbleID, txType)
	if snapshotdb.NonDbNotFoundErr(err) {
		return nil, err
	}
	if nil == txHashList {
		return []common.Hash{}, err
	}
	return *txHashList, err
}

// StoreTxHashToBub The mapping relationship between the sub-chain transaction hash and the main chain transaction hash is stored
func (bp *BubblePlugin) StoreTxHashToBub(blockHash common.Hash, bubbleID *big.Int, txHash common.Hash, txType bubble.TxType) error {
	// get hash list
	txHashList, err := bp.GetTxHashListByBub(blockHash, bubbleID, txType)
	if snapshotdb.NonDbNotFoundErr(err) {
		return err
	}
	// add new tx hash
	txHashList = append(txHashList, txHash)
	return bp.db.StoreTxHashListToBub(blockHash, bubbleID, txHashList, txType)
}

// AddAccAssetToBub Add account staking assets to bubble
func (bp *BubblePlugin) AddAccAssetToBub(blockHash common.Hash, bubbleId *big.Int, stakingAsset *bubble.AccountAsset) error {
	if nil == stakingAsset {
		return errors.New("the staking tokens information is empty")
	}
	// Check if a bubble exists. You cannot pledge assets to a bubble that does not exist
	bubInfo, err := bp.GetBubbleInfo(blockHash, bubbleId)
	if nil != err || nil == bubInfo {
		return bubble.ErrBubbleNotExist
	}

	// Determine whether the account has a history of pledging tokens within the bubble
	accAsset, err := bp.GetAccAssetOfBub(blockHash, bubbleId, stakingAsset.Account)
	if snapshotdb.NonDbNotFoundErr(err) {
		return err
	}
	// New staking account
	if nil == accAsset {
		// Store the staking tokens account in bubble
		if err := bp.AddAccToBub(blockHash, bubbleId, stakingAsset.Account); nil != err {
			return err
		}
		// Store new account assets
		if err = bp.db.StoreAccAssetToBub(blockHash, bubbleId, *stakingAsset); nil != err {
			return err
		}
	} else {
		// Update account assets
		// Update native token
		accAsset.NativeAmount = new(big.Int).Add(accAsset.NativeAmount, stakingAsset.NativeAmount)
		// Update ERC20(add new asset to old asset)
		for _, newAsset := range stakingAsset.TokenAssets {
			tokenAddr := newAsset.TokenAddr
			amount := newAsset.Balance
			isFind := false
			for i, oldAsset := range accAsset.TokenAssets {
				// The currency already exists
				if tokenAddr == oldAsset.TokenAddr {
					accAsset.TokenAssets[i].Balance = new(big.Int).Add(accAsset.TokenAssets[i].Balance, amount)
					isFind = true
					break
				}
			}
			// New currency species： New ERC20 Token
			if !isFind {
				accAsset.TokenAssets = append(accAsset.TokenAssets, bubble.AccTokenAsset{TokenAddr: tokenAddr, Balance: amount})
			}
		}
		// Update account asset information
		if err = bp.db.StoreAccAssetToBub(blockHash, bubbleId, *accAsset); nil != err {
			return err
		}
	}
	return nil
}

// PostMintTokenEvent Send the coin mint event and wait for the coin mint task to be processed
func (bp *BubblePlugin) PostMintTokenEvent(mintTokenTask *bubble.MintTokenTask) error {
	if err := bp.eventMux.Post(*mintTokenTask); nil != err {
		log.Error("post mintToken task failed", "err", err.Error())
		return err
	}

	return nil
}

// PostCreateBubbleEvent Send the create bubble event and wait for the task to be processed
func (bp *BubblePlugin) PostCreateBubbleEvent(task *bubble.CreateBubbleTask) error {
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post CreateBubble task failed", "err", err.Error())
		return err
	}

	return nil
}

// PostReleaseBubbleEvent Send the release bubble event and wait for the task to be processed
func (bp *BubblePlugin) PostReleaseBubbleEvent(task *bubble.ReleaseBubbleTask) error {
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post ReleaseBubble task failed", "err", err.Error())
		return err
	}

	return nil
}

// PostRemoteDestroyEvent Send the release bubble event and wait for the task to be processed
func (bp *BubblePlugin) PostRemoteDestroyEvent(task *bubble.RemoteDestroyTask) error {
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post remoteDestroy task failed", "err", err.Error())
		return err
	}

	return nil
}

// PostRemoteDeployEvent Send the remote deploy contract event and wait for the task to be processed
func (bp *BubblePlugin) PostRemoteDeployEvent(task *bubble.RemoteDeployTask) error {
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post RemoteDeployTask failed", "err", err.Error())
		return err
	}

	return nil
}

// PostRemoteRemoveEvent Send the remote remove event and wait for the task to be processed
func (bp *BubblePlugin) PostRemoteRemoveEvent(task *bubble.RemoteRemoveTask) error {
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post RemoteRemoveTask failed", "err", err.Error())
		return err
	}

	return nil
}

// PostRemoteCallEvent Send the remote call contract function event and wait for the task to be processed
func (bp *BubblePlugin) PostRemoteCallEvent(task *bubble.RemoteCallTask) error {
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post RemoteCallTask failed", "err", err.Error())
		return err
	}

	return nil
}

// Generate the rlp encoding of the sub-chain minting transaction
func genMintTokenRlpData(mintToken bubble.MintTokenTask) []byte {
	var params [][]byte
	params = make([][]byte, 0)
	// sub-chain mintToken function coding
	mintTokenType := uint16(6000)
	fnType, _ := rlp.EncodeToBytes(mintTokenType)
	params = append(params, fnType)

	txHash, _ := rlp.EncodeToBytes(mintToken.TxHash)
	accAsset, _ := rlp.EncodeToBytes(mintToken.AccAsset)
	params = append(params, txHash)
	params = append(params, accAsset)
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, params)
	if err != nil {
		return nil
	}
	// rlpData := hexutil.Encode(buf.Bytes())
	// fmt.Printf("funcType:%d rlp data = %s\n", mintTokenType, rlpData)
	return buf.Bytes()
}

// HandleMintTokenTask Handle MintToken task
func (bp *BubblePlugin) HandleMintTokenTask(mintToken *bubble.MintTokenTask) ([]byte, error) {
	log.Info("failed connect operator node", mintToken)
	if nil == mintToken || nil == mintToken.AccAsset {
		return nil, errors.New("mintToken task information is empty")
	}
	client, err := ethclient.Dial(mintToken.RPC)
	if err != nil || client == nil {
		log.Error("failed connect operator node", "err", err.Error())
		return nil, errors.New("failed connect operator node")
	}
	// Construct transaction parameters
	priKey := bp.opPriKey
	// Call the sub-chain system contract MintToken interface
	toAddr := common.HexToAddress(SubChainSysAddr)
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("Wrong private key", "err", err.Error())
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("the private key to the public key failed")
		return nil, errors.New("the private key to the public key failed")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	// The staking address of the operation node is taken as the operation node address
	// It is determined whether the main chain operation node signs the transaction
	if fromAddr != mintToken.OpAddr {
		log.Error("The mintToken transaction sender is not the main-chain operation address")
		return nil, errors.New("the mintToken transaction sender is not the main-chain operation address")
	}
	// get account nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		log.Error("Failed to obtain the account nonce", "err", err.Error())
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Failed to get gasPrice", "err", err.Error())
		return nil, err
	}
	value := big.NewInt(0)
	gasLimit := uint64(300000)
	// Assemble the data of the minting interface
	data := genMintTokenRlpData(*mintToken)
	if nil == data {
		return nil, errors.New("genMintTokenRlpData failed")
	}
	// Creating transaction objects
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, data)

	// The sender's private key is used to sign the transaction
	chainID, err := client.ChainID(context.Background())
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("Signing mintToken transaction failed", "err", err.Error())
		return nil, err
	}

	// Sending transactions
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Failed to send mintToken transaction", "err", err.Error())
		return nil, err
	}

	hash := signedTx.Hash()
	log.Debug("mintToken tx hash", hash.Hex())
	return hash.Bytes(), nil
}

func encodeRemoteDeploy(task *bubble.RemoteDeployTask, code []byte) []byte {
	queue := make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(uint16(8000))
	txHash, _ := rlp.EncodeToBytes(task.TxHash)
	sender, _ := rlp.EncodeToBytes(task.Caller)
	address, _ := rlp.EncodeToBytes(task.Address)
	byteCode, _ := rlp.EncodeToBytes(code)
	data, _ := rlp.EncodeToBytes(task.Data)
	queue = append(queue, fnType)
	queue = append(queue, txHash)
	queue = append(queue, sender)
	queue = append(queue, address)
	queue = append(queue, byteCode)
	queue = append(queue, data)

	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, queue)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func encodeRemoteCall(task *bubble.RemoteCallTask) []byte {
	queue := make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(uint16(8004))
	txHash, _ := rlp.EncodeToBytes(task.TxHash)
	caller, _ := rlp.EncodeToBytes(task.Caller)
	Contract, _ := rlp.EncodeToBytes(task.Contract)
	Data, _ := rlp.EncodeToBytes(task.Data)
	queue = append(queue, fnType)
	queue = append(queue, txHash)
	queue = append(queue, caller)
	queue = append(queue, Contract)
	queue = append(queue, Data)
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, queue)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func encodeRemoteRemove(task *bubble.RemoteRemoveTask) []byte {
	queue := make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(uint16(8002))
	txHash, _ := rlp.EncodeToBytes(task.TxHash)
	Contract, _ := rlp.EncodeToBytes(task.Contract)
	queue = append(queue, fnType)
	queue = append(queue, txHash)
	queue = append(queue, Contract)
	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, queue)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func encodeRemoteDestroy(task *bubble.RemoteDestroyTask) []byte {
	queue := make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(uint16(8001))
	blockNumber, _ := rlp.EncodeToBytes(task.BlockNumber)
	queue = append(queue, fnType)
	queue = append(queue, blockNumber)

	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, queue)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

// HandleRemoteDeployTask Handle RemoteDeploy task
func (bp *BubblePlugin) HandleRemoteDeployTask(task *bubble.RemoteDeployTask) ([]byte, error) {
	if nil == task {
		return nil, errors.New("RemoteDeployTask is empty")
	}
	client, err := ethclient.Dial(task.RPC)
	if err != nil || client == nil {
		log.Error("failed connect operator node", "err", err.Error(), "rpc", task.RPC)
		return nil, errors.New("failed connect operator node")
	}
	// Construct transaction parameters
	priKey := bp.opPriKey
	// Call the sub-chain system contract RemoteDeploy interface
	toAddr := common.HexToAddress(remoteBubbleAddr)
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("Wrong private key", "err", err.Error())
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("the private key to the public key failed")
		return nil, errors.New("the private key to the public key failed")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	// The staking address of the operation node is taken as the operation node address
	// It is determined whether the main chain operation node signs the transaction
	if fromAddr != task.OpAddr {
		log.Error("The transaction sender is not the main-chain operation address")
		return nil, errors.New("the transaction sender is not the main-chain operation address")
	}
	// get account nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		log.Error("Failed to obtain the account nonce", "err", err.Error())
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Failed to get gasPrice", "err", err.Error())
		return nil, err
	}
	value := big.NewInt(0)
	gasLimit := uint64(300000)

	time.Sleep(3 * time.Second)
	byteCode, err := BubbleInstance().GetByteCode(task.BlockHash, task.Address)
	if err != nil {
		return nil, err
	}
	data := encodeRemoteDeploy(task, byteCode)
	if nil == data {
		return nil, errors.New("encode remoteDeploy transaction failed")
	}
	// Creating transaction objects
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, data)

	// The sender's private key is used to sign the transaction
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get chainId error: %s", err.Error()))
	}

	if chainID.Cmp(task.BubbleID) != 0 {
		return nil, errors.New(fmt.Sprintf("chainID is wrong, expect %d, actual %d", task.BubbleID.Uint64(), chainID.Uint64()))
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("Signing remoteDeploy transaction failed", "err", err.Error())
		return nil, err
	}

	// Sending transactions
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Failed to send remoteDeploy transaction", "err", err.Error())
		return nil, err
	}

	hash := signedTx.Hash()
	log.Debug("remoteDeploy tx hash", hash.Hex())
	return hash.Bytes(), nil
}

// HandleRemoteCallTask Handle RemoteCall task
func (bp *BubblePlugin) HandleRemoteCallTask(task *bubble.RemoteCallTask) ([]byte, error) {
	if nil == task {
		return nil, errors.New("RemoteCallTask is empty")
	}
	client, err := ethclient.Dial(task.RPC)
	if err != nil || client == nil {
		log.Error("failed connect operator node", "err", err.Error())
		return nil, errors.New("failed connect operator node")
	}
	// Construct transaction parameters
	priKey := bp.opPriKey
	// Call the sub-chain system contract RemoteDeploy interface
	toAddr := common.HexToAddress(remoteBubbleAddr)
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("Wrong private key", "err", err.Error())
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("the private key to the public key failed")
		return nil, errors.New("the private key to the public key failed")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	// The staking address of the operation node is taken as the operation node address
	// It is determined whether the main chain operation node signs the transaction
	if fromAddr != task.OpAddr {
		log.Error("The transaction sender is not the main-chain operation address")
		return nil, errors.New("the transaction sender is not the main-chain operation address")
	}
	// get account nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		log.Error("Failed to obtain the account nonce", "err", err.Error())
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Failed to get gasPrice", "err", err.Error())
		return nil, err
	}
	value := big.NewInt(0)
	gasLimit := uint64(300000)
	data := encodeRemoteCall(task)
	if nil == data {
		return nil, errors.New("encode remoteCall transaction failed")
	}
	// Creating transaction objects
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, data)

	// The sender's private key is used to sign the transaction
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get chainId error: %s", err.Error()))
	}

	if chainID.Cmp(task.BubbleID) != 0 {
		return nil, errors.New(fmt.Sprintf("chainID is wrong, expect %d, actual %d", task.BubbleID.Uint64(), chainID.Uint64()))
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("Signing remoteCall transaction failed", "err", err.Error())
		return nil, err
	}

	// Sending transactions
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Failed to send remoteCall transaction", "err", err.Error())
		return nil, err
	}

	hash := signedTx.Hash()
	log.Debug("remoteCall tx hash", hash.Hex())
	return hash.Bytes(), nil
}

// HandleRemoteRemoveTask Handle RemoteRemove task
func (bp *BubblePlugin) HandleRemoteRemoveTask(task *bubble.RemoteRemoveTask) ([]byte, error) {
	if nil == task {
		return nil, errors.New("RemoteRemoveTask is empty")
	}
	client, err := ethclient.Dial(task.RPC)
	if err != nil || client == nil {
		log.Error("failed connect operator node", "err", err.Error())
		return nil, errors.New("failed connect operator node")
	}
	// Construct transaction parameters
	priKey := bp.opPriKey
	// Call the sub-chain system contract RemoteRemove interface
	toAddr := common.HexToAddress(remoteBubbleAddr)
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("Wrong private key", "err", err.Error())
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("the private key to the public key failed")
		return nil, errors.New("the private key to the public key failed")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	// The staking address of the operation node is taken as the operation node address
	// It is determined whether the main chain operation node signs the transaction
	if fromAddr != task.OpAddr {
		log.Error("The transaction sender is not the main-chain operation address")
		return nil, errors.New("the transaction sender is not the main-chain operation address")
	}
	// get account nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		log.Error("Failed to obtain the account nonce", "err", err.Error())
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Failed to get gasPrice", "err", err.Error())
		return nil, err
	}
	value := big.NewInt(0)
	gasLimit := uint64(300000)
	data := encodeRemoteRemove(task)
	if nil == data {
		return nil, errors.New("encode remoteCall transaction failed")
	}
	// Creating transaction objects
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, data)

	// The sender's private key is used to sign the transaction
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get chainId error: %s", err.Error()))
	}

	if chainID.Cmp(task.BubbleID) != 0 {
		return nil, errors.New(fmt.Sprintf("chainID is wrong, expect %d, actual %d", task.BubbleID.Uint64(), chainID.Uint64()))
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("Signing remoteCall transaction failed", "err", err.Error())
		return nil, err
	}

	// Sending transactions
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Failed to send remoteCall transaction", "err", err.Error())
		return nil, err
	}

	hash := signedTx.Hash()
	log.Debug("remoteCall tx hash", hash.Hex())
	return hash.Bytes(), nil
}

func makeGenesisL2(val *bubble.ValidatorInfo) *bubble.GenesisL2 {

	var initNodes []params.InitNode
	for _, node := range val.MicroNodes {
		initNode := params.InitNode{
			Enode:     node.P2PURI,
			BlsPubkey: node.BlsPubKey.String(),
		}
		initNodes = append(initNodes, initNode)
	}
	generalAddr := common.HexToAddress("0x51625d7FFda8B38a6987EAa99aeA3269923237a3")
	generalBalance, _ := new(big.Int).SetString("9727638019000000000000000000", 10)
	rewardMgrPoolIssue, _ := new(big.Int).SetString("200000000000000000000000000", 10)

	ec := xcom.GetEc(xcom.DefaultMainNet)
	em := &bubble.EconomicModel{
		EconomicModel: *ec,
		Staking: bubble.StakingConfig{
			StakingConfig:       xcom.GetEc(xcom.DefaultMainNet).Staking,
			StakingConfigExtend: xcom.GetEce().Staking,
		},
	}

	genesisL2 := &bubble.GenesisL2{
		Config: &params.ChainConfig{
			ChainID: val.BubbleId,
			Frps:    params.GlobalFrpsCfg,
			Cbft: &params.CbftConfig{
				Period:        10000,
				Amount:        10,
				InitialNodes:  params.ConvertNodeUrl(initNodes),
				ValidatorMode: "dpos",
			},
			GenesisVersion: gov.L2Version,
		},
		OpConfig: &bubble.OpConfig{
			MainChain: val.OperatorsL1,
			SubChain:  val.OperatorsL2,
		},
		EconomicModel: em,
		Nonce:         []byte{},
		Timestamp:     params.MainNetGenesisTimestamp,
		ExtraData:     []byte{},
		GasLimit:      math.HexOrDecimal64(params.GenesisGasLimit),
		Coinbase:      common.ZeroAddr,
		Alloc: bubble.GenesisAlloc{
			vm.RewardManagerPoolAddr: {Balance: rewardMgrPoolIssue},
			generalAddr:              {Balance: generalBalance},
		},
		Number:     0,
		GasUsed:    0,
		ParentHash: common.ZeroHash,
	}
	// The ip address of the frps is updated to be the ip address of the operating node of L1
	rpcUrl := val.OperatorsL1[0].RPC
	index := strings.Index(rpcUrl, "//")
	if index != -1 {
		rpcUrl = rpcUrl[index+2:]
	}
	index = strings.Index(rpcUrl, ":")
	if index != -1 {
		rpcUrl = rpcUrl[0:index]
	}
	genesisL2.Config.Frps.ServerIP = rpcUrl
	return genesisL2
}

// HandleCreateBubbleTask Handle create bubble task
func (bp *BubblePlugin) HandleCreateBubbleTask(task *bubble.CreateBubbleTask) error {
	if task == nil {
		log.Error("CreateBubbleTask is nil")
		return errors.New("CreateBubbleTask is nil")
	}

	// wait for blocks to be written to the db
	// TODO：if not, panic by BLS segmentation violation
	time.Sleep(3 * time.Second)
	val, err := bp.GetValidatorInfo(common.ZeroHash, task.BubbleID)
	if err != nil {
		log.Error("failed to get bubble info", "error", err.Error(), "bubbleId", task.BubbleID)
		return errors.New(fmt.Sprintf("failed to get bubble info: %s", err.Error()))
	}

	genesisL2 := makeGenesisL2(val)
	genesisData, err := json.Marshal(genesisL2)
	if err != nil {
		log.Error("failed to marshal genesis", "error", err.Error())
		return errors.New(fmt.Sprintf("failed to marshal genesis: %s", err.Error()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(val.MicroNodes))

	client := &http.Client{}
	data := fmt.Sprintf("{\"type\": %d, \"data\": %s}", bubble.CreateBubble, string(genesisData))

	// make request
	for _, microNode := range val.MicroNodes {
		// send and retry CreateBubbleTask
		log.Debug("prepare to send CreateBubbleTask", "ElectronURI", microNode.ElectronURI, "data", data)
		go sendTask(ctx, waitGroup, client, microNode.ElectronURI, data)
	}

	// wait task done
	go func() {
		waitGroup.Wait()
		cancel()
	}()
	<-ctx.Done()
	if ctx.Err().Error() == "context deadline exceeded" {
		return errors.New("task timeout")
	}

	return nil
}

// HandleReleaseBubbleTask Handle release bubble task
func (bp *BubblePlugin) HandleReleaseBubbleTask(task *bubble.ReleaseBubbleTask) error {
	if task == nil {
		return errors.New("releaseBubbleTask is nil")
	}

	// wait for blocks to be written to the db
	// TODO：if not, panic by BLS segmentation violation
	time.Sleep(3 * time.Second)
	val, err := bp.GetValidatorInfo(common.ZeroHash, task.BubbleID)
	if err != nil {
		log.Error("failed to get bubble info", "error", err.Error())
		return errors.New(fmt.Sprintf("failed to get bubble info: %s", err.Error()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(val.MicroNodes))

	client := &http.Client{}
	data := fmt.Sprintf("{\"type\": %d, \"data\": %s}", bubble.ReleaseBubble, val.BubbleId)

	// make request
	for _, microNode := range val.MicroNodes {
		// send and retry ReleaseBubbleTask
		log.Debug("prepare to send ReleaseBubbleTask", "ElectronURI", microNode.ElectronURI, "data", data)
		go sendTask(ctx, waitGroup, client, microNode.ElectronURI, data)
	}

	// wait task done
	go func() {
		waitGroup.Wait()
		cancel()
	}()
	<-ctx.Done()
	if ctx.Err().Error() == "context deadline exceeded" {
		return errors.New("task timeout")
	}

	return nil
}

// HandleRemoteDestroyTask Handle RemoteDestroy task
func (bp *BubblePlugin) HandleRemoteDestroyTask(task *bubble.RemoteDestroyTask) ([]byte, error) {
	if nil == task {
		return nil, errors.New("RemoteDestroyTask is empty")
	}
	client, err := ethclient.Dial(task.RPC)
	if err != nil || client == nil {
		log.Error("failed connect operator node", "err", err.Error())
		return nil, errors.New("failed connect operator node")
	}
	// Construct transaction parameters
	priKey := bp.opPriKey
	// Call the sub-chain system contract RemoteDeploy interface
	toAddr := common.HexToAddress(remoteBubbleAddr)
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("Wrong private key", "err", err.Error())
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("the private key to the public key failed")
		return nil, errors.New("the private key to the public key failed")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	if fromAddr != task.OpAddr {
		log.Error("The transaction sender is not the main-chain operation address")
		return nil, errors.New("the transaction sender is not the main-chain operation address")
	}
	// get account nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		log.Error("Failed to obtain the account nonce", "err", err.Error())
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Failed to get gasPrice", "err", err.Error())
		return nil, err
	}
	value := big.NewInt(0)
	gasLimit := uint64(300000)
	data := encodeRemoteDestroy(task)
	if nil == data {
		return nil, errors.New("encode remoteDeploy transaction failed")
	}
	// Creating transaction objects
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, data)

	// The sender's private key is used to sign the transaction
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("get chainId error: %s", err.Error()))
	}

	if chainID.Cmp(task.BubbleID) != 0 {
		return nil, errors.New(fmt.Sprintf("chainID is wrong, expect %d, actual %d", task.BubbleID.Uint64(), chainID.Uint64()))
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("Signing remoteDeploy transaction failed", "err", err.Error())
		return nil, err
	}

	// Sending transactions
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Failed to send remoteDeploy transaction", "err", err.Error())
		return nil, err
	}

	hash := signedTx.Hash()
	log.Debug("remoteDeploy tx hash", hash.Hex())
	return hash.Bytes(), nil
}

func sendTask(ctx context.Context, waitGroup *sync.WaitGroup, client *http.Client, url string, data string) {

	for i := 0; i < 10; i++ {
		// new request
		dataReader := strings.NewReader(data)
		req, err := http.NewRequest(http.MethodPost, url, dataReader)
		if err != nil {
			log.Error("new http request failed", "err", err.Error())
		}
		req.WithContext(ctx)
		req.Header.Set("Content-Type", "application/json")

		// send request
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != 200 {
			log.Debug("send task to microNode failed", "retry", i, "error", err.Error(), "response", resp)
			time.Sleep(time.Duration(3) * time.Second)
			continue
		}
		log.Info("send task to microNode succeed", "ElectronURI", url, "req", req)
		resp.Body.Close()
		waitGroup.Done()
		break
	}
}
