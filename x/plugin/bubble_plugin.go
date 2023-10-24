package plugin

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	gomath "math"
	"math/big"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/json"
	"github.com/bubblenet/bubble/common/math"
	"github.com/bubblenet/bubble/common/sort"
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
	"github.com/bubblenet/bubble/x/xutil"
	"golang.org/x/crypto/sha3"
)

const (
	SubChainSysAddr = "0x1000000000000000000000000000000000000020" // Sub chain system contract address
)

var (
	bubblePluginOnce sync.Once
	bubblePlugin     *BubblePlugin
)

const bubbleLife = 172800

type BubblePlugin struct {
	stkPlugin  *StakingPlugin
	stk2Plugin *StakingL2Plugin
	db         *bubble.BubbleDB
	NodeID     discover.NodeID // id of the local node
	eventMux   *event.TypeMux
	opPriKey   string // Main chain operation address private key
}

// BubbleInstance instance a global BubblePlugin
func BubbleInstance() *BubblePlugin {
	bubblePluginOnce.Do(func() {
		log.Info("Init bubble plugin ...")
		bubblePlugin = &BubblePlugin{
			stkPlugin:  StakingInstance(),
			stk2Plugin: StakingL2Instance(),
			db:         bubble.NewBubbleDB(),
		}
	})
	return bubblePlugin
}

func (bp *BubblePlugin) BeginBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	return nil
}

func (bp *BubblePlugin) EndBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {

	if xutil.IsEndOfEpoch(header.Number.Uint64()) {
		iter := bp.db.IteratorBubStatus(blockHash, 0)
		if err := iter.Error(); nil != err {
			return err
		}
		defer iter.Release()

		//bubs := make([]*bubble.Bubble, 0)

		for iter.Valid(); iter.Next(); {
			data := iter.Value()
			bubStatus := new(bubble.BubStatus)
			if err := rlp.DecodeBytes(data, bubStatus); err != nil {
				return err
			}

			blockNumber := header.Number.Uint64()
			if blockNumber < bubStatus.PreReleaseBlock {
				continue
			}

			if blockNumber < bubStatus.ReleaseBlock {
				bubStatus.State = bubble.PreReleaseStatus

				if bubStatus.ContractCount > 0 {
					continue
				}
			}

			err := bp.ReleaseBubble(blockHash, header.Number, bubStatus.BubbleId)
			if err != nil {
				log.Error("Failed to call ReleaseBubble on BubblePlugin EndBlock",
					"blockNumber", header.Number.Uint64(), "blockHash", blockHash.Hex(), "bubble", bubStatus.BubbleId, "err", err)
				return err
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

// GetBubbleInfo return the bubble information by bubble ID
func (bp *BubblePlugin) GetBubbleInfo(blockHash common.Hash, bubbleID *big.Int) (*bubble.Bubble, error) {
	// return bp.db.GetBubbleStore(blockHash, bubbleID)
	// get bubble basics
	basics, err := bp.GetBubBasics(blockHash, bubbleID)
	if nil == basics || err != nil {
		return nil, errors.New("failed to get bubble basics information")
	}

	// get bubble state
	state, err := bp.GetBubStatus(blockHash, bubbleID)
	if nil == state || err != nil {
		return nil, errors.New("failed to get bubble state")
	}
	//bub.State = *state
	// get bubble txHashList
	// get StakingToken Hash List
	stTxHashList, err := bp.GetTxHashListByBub(blockHash, bubbleID, bubble.StakingToken)
	if snapshotdb.NonDbNotFoundErr(err) {
		return nil, errors.New("failed to get bubble StakingToken transaction hash list")
	}
	//bub.StakingTokenTxHashList = stTxHashList

	// get WithdrewToken Hash List
	wdTxHashList, err := bp.GetTxHashListByBub(blockHash, bubbleID, bubble.WithdrewToken)
	if snapshotdb.NonDbNotFoundErr(err) {
		return nil, errors.New("failed to get bubble WithdrewToken transaction hash list")
	}
	//bub.WithdrewTokenTxHashList = wdTxHashList

	// get SettleBubble Hash List
	sbTxHashList, err := bp.GetTxHashListByBub(blockHash, bubbleID, bubble.SettleBubble)
	if snapshotdb.NonDbNotFoundErr(err) {
		return nil, errors.New("failed to get bubble SettleBubble transaction hash list")
	}
	//bub.SettleBubbleTxHashList = sbTxHashList

	bub := &bubble.Bubble{
		Basics:    basics,
		BubStatus: state,
		BubMutable: &bubble.BubMutable{
			StakingTokenTxHashList:  stTxHashList,
			WithdrewTokenTxHashList: wdTxHashList,
			SettleBubbleTxHashList:  sbTxHashList,
		},
	}

	return bub, nil
}

// GetBubBasics return the bubble basics by bubble ID
func (bp *BubblePlugin) GetBubBasics(blockHash common.Hash, bubbleID *big.Int) (*bubble.BubBasics, error) {
	bubBasic, err := bp.db.GetBubBasics(blockHash, bubbleID)
	if err != nil || nil == bubBasic {
		return nil, err
	}
	// The rpc of the bubble-chain operator node is obtained from the verifier
	if "" == bubBasic.OperatorsL2[0].RPC {
		sk := StakingL2Instance()
		nodeId := bubBasic.OperatorsL2[0].NodeId
		canAddr, err := xutil.NodeId2Addr(nodeId)
		if nil != err {
			return nil, err
		}
		base, err := sk.db.GetCanBaseStore(blockHash, canAddr)
		if nil != err {
			return nil, err
		}
		bubBasic.OperatorsL2[0].RPC = base.RPCURI

		// update bubble basic
		if err := bp.db.StoreBubBasics(blockHash, bubbleID, bubBasic); nil != err {
			return nil, err
		}
	}

	return bubBasic, nil
}

// GetBubStatus return the bubble state by bubble ID
func (bp *BubblePlugin) GetBubStatus(blockHash common.Hash, bubbleID *big.Int) (*bubble.BubStatus, error) {
	return bp.db.GetBubStatus(blockHash, bubbleID)
}

func (bp *BubblePlugin) CheckBubbleElements(blockHash common.Hash, sizeCode uint8) error {
	bubbleSize, err := bubble.GetBubbleSize(sizeCode)
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
func (bp *BubblePlugin) CreateBubble(blockHash common.Hash, blockNumber *big.Int, from common.Address, nonce uint64, parentNonce [][]byte, sizeCode uint8) (*bubble.Bubble, error) {
	bubbleSize, err := bubble.GetBubbleSize(sizeCode)
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
		operator := &bubble.Operator{
			NodeId: can.NodeId,
			RPC:    can.RPCURI,
			OpAddr: can.StakingAddress,
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

	basics := &bubble.BubBasics{
		BubbleId: bubbleID,

		OperatorsL1: OperatorsL1,
		OperatorsL2: OperatorsL2,
		MicroNodes:  microNodes,
	}

	preReleaseBlock := uint64(gomath.Ceil(float64(blockNumber.Uint64()+bubbleLife)/float64(xutil.EpochSize()))) * xutil.EpochSize()
	releaseBlock := uint64(gomath.Ceil(float64(blockNumber.Uint64()+bubbleLife*1.5)/float64(xutil.EpochSize()))) * xutil.EpochSize()
	status := &bubble.BubStatus{
		BubbleId:        bubbleID,
		State:           bubble.ActiveStatus,
		ContractCount:   0,
		CreateBlock:     blockNumber.Uint64(),
		PreReleaseBlock: preReleaseBlock,
		ReleaseBlock:    releaseBlock,
	}

	bub := &bubble.Bubble{
		Basics:    basics,
		BubStatus: status,
	}

	// store bubble basics
	if err := bp.db.StoreBubBasics(blockHash, basics.BubbleId, basics); err != nil {
		log.Error("Failed to CreateBubble on bubblePlugin: Store bubble basics failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleId", basics.BubbleId, "err", err)
		return nil, err
	}
	// store bubble state
	if err := bp.db.StoreBubStatus(blockHash, basics.BubbleId, status); err != nil {
		log.Error("Failed to CreateBubble on bubblePlugin: Store bubble state failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleId", basics.BubbleId, "err", err)
		return nil, err
	}

	if err := bp.db.StoreSizedBubbleID(blockHash, sizeCode, basics.BubbleId); err != nil {
		log.Error("Failed to CreateBubble on bubblePlugin: Store bubble sized info failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleId", basics.BubbleId, "err", err)
		return nil, err
	}

	return bub, nil
}

// generateBubbleID generate bubble ID use sha3 algorithm by bubble info
func (bp *BubblePlugin) generateBubbleID(creator common.Address, nonce *big.Int, committer bubble.CandidateQueue) (*big.Int, error) {
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
	vrfQueue, err := VRFQueueWrapper(operators, func(item interface{}) *VRFItem {
		w, _ := new(big.Int).SetString("1000000000000000000000000", 10)
		return &VRFItem{
			v: item,
			w: w,
		}
	})
	if err != nil {
		return nil, err
	}

	// VRF Elect
	log.Info("ElectOperatorL1 run VRF", "vrfQueue len", len(vrfQueue), "preNonces len", len(preNonces))
	electedVrfQueue, err := VRF(vrfQueue, operatorNumber, curNonce, preNonces[:len(vrfQueue)])
	if err != nil {
		return nil, err
	}

	// unwrap the VRF able queue
	electedOperators := make([]*bubble.Operator, 0)
	for _, item := range electedVrfQueue {
		if operator, ok := (item.v).(*bubble.Operator); ok {
			electedOperators = append(electedOperators, operator)
		} else {
			return nil, errors.New("type error")
		}
	}

	return electedOperators, nil
}

// ElectOperatorL2 Elect the Layer2 Operator nodes for the bubble chain by VRF
func (bp *BubblePlugin) ElectOperatorL2(blockHash common.Hash, operatorNumber uint, curNonce []byte, preNonces [][]byte) (bubble.CandidateQueue, error) {
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
	vrfQueue, err := VRFQueueWrapper(operatorQueue, func(item interface{}) *VRFItem {
		if candidate, ok := (item).(*stakingL2.Candidate); ok {
			return &VRFItem{
				v: item,
				w: candidate.Shares,
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// VRF Elect
	log.Info("ElectOperatorL2 run VRF", "vrfQueue len", len(vrfQueue), "preNonces len", len(preNonces))
	electedVrfQueue, err := VRF(vrfQueue, operatorNumber, curNonce, preNonces[:len(vrfQueue)])
	if err != nil {
		return nil, err
	}

	// unwrap the VRF able queue
	electedOperators := make(bubble.CandidateQueue, 0)
	for _, item := range electedVrfQueue {
		if Operator, ok := (item.v).(*stakingL2.Candidate); ok {
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
func (bp *BubblePlugin) ElectBubbleMicroNodes(blockHash common.Hash, committeeNumber uint, curNonce []byte, preNonces [][]byte) (bubble.CandidateQueue, error) {
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
	vrfQueue, err := VRFQueueWrapper(committeeQueue, func(item interface{}) *VRFItem {
		if candidate, ok := (item).(*stakingL2.Candidate); ok {
			return &VRFItem{
				v: item,
				w: candidate.Shares,
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// VRF Elect
	log.Info("ElectBubbleMicroNodes run VRF", "vrfQueue len", len(vrfQueue), "preNonces len", len(preNonces))
	electedVrfQueue, err := VRF(vrfQueue, committeeNumber, curNonce, preNonces[:len(vrfQueue)])
	if err != nil {
		return nil, err
	}

	// unwrap the VRF able queue
	committees := make(bubble.CandidateQueue, 0)
	for _, item := range electedVrfQueue {
		if candidate, ok := (item.v).(*stakingL2.Candidate); ok {
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

func (bp *BubblePlugin) GetSizedBubbleIDs(blockHash common.Hash, sizeCode uint8) ([]*big.Int, error) {
	iter := bp.db.IteratorSizedBubbleID(blockHash, sizeCode, 0)
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

func (bp *BubblePlugin) ElectBubble(blockHash common.Hash, nonce uint64, preNonces [][]byte, sizeCode uint8) (*big.Int, error) {
	bubIDs, err := bp.GetSizedBubbleIDs(blockHash, sizeCode)
	if err != nil {
		return nil, err
	}

	if len(preNonces) < len(bubIDs) {
		fitNonces := make([][]byte, len(bubIDs))
		fitNonces = append(fitNonces, preNonces...)
		preNonces = fitNonces
	}

	vrfQueue, err := VRFQueueWrapper(bubIDs, func(item interface{}) *VRFItem {
		w, _ := new(big.Int).SetString("1200000000000000000000000", 10)
		return &VRFItem{
			v: item,
			w: w,
		}
	})

	electedVrfQueue, err := VRF(vrfQueue, 1, common.Uint64ToBytes(nonce), preNonces)
	if err != nil {
		return nil, err
	}
	if len(electedVrfQueue) != 1 {
		return nil, errors.New("the number of bubbles elected is not correct")
	}

	// unwrap the VRF able queue
	vrfItem := electedVrfQueue[0]
	if bubID, ok := (vrfItem.v).(*big.Int); ok {
		return bubID, nil
	} else {
		return nil, errors.New("type error")
	}
}

// ReleaseBubble run the non-business logic to release the bubble
func (bp *BubblePlugin) ReleaseBubble(blockHash common.Hash, blockNumber *big.Int, bubbleID *big.Int) error {
	bub, err := bp.GetBubbleInfo(blockHash, bubbleID)
	if err != nil {
		return err
	}

	var committeeCount uint32
	// release the committeeL2 nodes to the DB
	for _, microNode := range bub.Basics.MicroNodes {
		addr, err := xutil.NodeId2Addr(microNode.NodeId)
		if err != nil {
			return err
		}
		// check whether the node has been withdrawn
		can, err := bp.stk2Plugin.db.GetCandidateStore(blockHash, addr)
		if can == nil || err == snapshotdb.ErrNotFound {
			break
		}

		if microNode.IsOperator {
			Operator, _ := bp.stk2Plugin.db.GetOperatorStore(blockHash, addr)
			if Operator != nil {
				log.Error("Failed to SetOperatorStore on ReleaseBubble: Operator info is exist",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", microNode.NodeId.String())
				return errors.New("failed to SetOperatorStore on ReleaseBubble: Operator info is exist")
			}

			if err := bp.stk2Plugin.db.SetOperatorStore(blockHash, addr, can); nil != err {
				log.Error("Failed to SetOperatorStore on ReleaseBubble: Store Operator info is failed",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", microNode.NodeId.String(), "err", err)
				return err
			}
		} else {
			committeeCount += 1
			Committee, _ := bp.stk2Plugin.db.GetCommitteeStore(blockHash, addr)
			if Committee != nil {
				log.Error("Failed to SetCommitteeStore on ReleaseBubble: Committee info is exist",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", microNode.NodeId.String())
				return errors.New("failed to SetCommitteeStore on ReleaseBubble: Committee info is exist")
			}

			if err := bp.stk2Plugin.db.SetCommitteeStore(blockHash, addr, can); nil != err {
				log.Error("Failed to SetCommitteeStore on ReleaseBubble: Store Committee info is failed",
					"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", microNode.NodeId.String(), "err", err)
				return err
			}
		}

	}

	status, err := bp.db.GetBubStatus(blockHash, bubbleID)
	if err != nil {
		log.Error("Failed to GetBubStatus on ReleaseBubble",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleID", bub.Basics.BubbleId.String(), "err", err)
		return err
	}
	if status.State <= bubble.ReleasedStatus {
		status.State = bubble.ReleasedStatus
	} else {
		log.Error("bubble is already released")
		return errors.New("bubble is already released")
	}
	if err := bp.db.StoreBubStatus(blockHash, bubbleID, status); err != nil {
		log.Error("Failed to StoreBubState on ReleaseBubble",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleID", bub.Basics.BubbleId.String(), "err", err)
		return err
	}

	if err := bp.db.DelSizedBubbleID(blockHash, bub.Basics.Size, bub.Basics.BubbleId); err != nil {
		log.Error("Failed to DelSizedBubbleID on ReleaseBubble",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleID", bub.Basics.BubbleId.String(), "err", err)
		return err
	}

	if err := bp.stk2Plugin.db.SubUsedCommitteeCount(blockHash, committeeCount); err != nil {
		log.Error("Failed to SubUsedCommitteeCount on ReleaseBubble",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleID", bub.Basics.BubbleId.String(), "err", err)
		return err
	}

	return nil
}

func (bp *BubblePlugin) GetByteCode(blockHash common.Hash, address common.Address) ([]byte, error) {
	return bp.db.GetContractByteCode(blockHash, address)
}

func (bp *BubblePlugin) StoreByteCode(blockHash common.Hash, address common.Address, byteCode []byte) error {
	return bp.db.StoreContractByteCode(blockHash, address, byteCode)
}

func (bp *BubblePlugin) GetBubContracts(blockHash common.Hash, bubbleID *big.Int) ([]*bubble.ContractInfo, error) {
	iter := bp.db.IteratorBubContract(blockHash, bubbleID, 0)
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
	return bp.db.GetBubContract(blockHash, bubbleID, address)
}

func (bp *BubblePlugin) StoreBubContract(blockHash common.Hash, bubbleID *big.Int, contractInfo *bubble.ContractInfo) error {
	return bp.db.StoreBubContract(blockHash, bubbleID, contractInfo)
}

func (bp *BubblePlugin) DelBubContract(blockHash common.Hash, bubbleID *big.Int, address common.Address) error {
	return bp.db.DelBubContract(blockHash, bubbleID, address)
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
func (bp *BubblePlugin) GetTxHashListByBub(blockHash common.Hash, bubbleID *big.Int, txType bubble.BubTxType) ([]common.Hash, error) {
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
func (bp *BubblePlugin) StoreTxHashToBub(blockHash common.Hash, bubbleID *big.Int, txHash common.Hash, txType bubble.BubTxType) error {
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
		log.Error("post mintToken task failed", "err", err)
		return err
	}

	return nil
}

// PostCreateBubbleEvent Send the create bubble event and wait for the task to be processed
func (bp *BubblePlugin) PostCreateBubbleEvent(task *bubble.CreateBubbleTask) error {
	log.Debug("PostCreateBubbleEvent", *task)
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post CreateBubble task failed", "err", err)
		return err
	}

	return nil
}

// PostReleaseBubbleEvent Send the release bubble event and wait for the task to be processed
func (bp *BubblePlugin) PostReleaseBubbleEvent(task *bubble.ReleaseBubbleTask) error {
	log.Debug("PostCreateBubbleEvent", *task)
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post ReleaseBubble task failed", "err", err)
		return err
	}

	return nil
}

// PostRemoteDeployEvent Send the remote deploy contract event and wait for the task to be processed
func (bp *BubblePlugin) PostRemoteDeployEvent(task *bubble.RemoteDeployTask) error {
	log.Debug("PostRemoteDeployEvent", *task)
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post RemoteDeployTask failed", "err", err)
		return err
	}

	return nil
}

// PostRemoteCallEvent Send the remote call contract function event and wait for the task to be processed
func (bp *BubblePlugin) PostRemoteCallEvent(task *bubble.RemoteCallTask) error {
	log.Debug("PostRemoteCallEvent", *task)
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post RemoteCallTask failed", "err", err)
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
		log.Error("failed connect operator node", "err", err)
		return nil, errors.New("failed connect operator node")
	}
	// Construct transaction parameters
	priKey := bp.opPriKey
	// Call the sub-chain system contract MintToken interface
	toAddr := common.HexToAddress(SubChainSysAddr)
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("Wrong private key", "err", err)
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
		log.Error("Failed to obtain the account nonce", "err", err)
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Failed to get gasPrice", "err", err)
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
		log.Error("Signing mintToken transaction failed", "err", err)
		return nil, err
	}

	// Sending transactions
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Failed to send mintToken transaction", "err", err)
		return nil, err
	}

	hash := signedTx.Hash()
	log.Debug("mintToken tx hash", hash.Hex())
	return hash.Bytes(), nil
}

func encodeRemoteDeploy(task *bubble.RemoteDeployTask) []byte {
	s := make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(uint16(6000))
	txHash, _ := rlp.EncodeToBytes(task.TxHash)
	address, _ := rlp.EncodeToBytes(task.Address)

	runtimeCode, err := BubbleInstance().db.GetContractByteCode(task.BlockHash, task.Address)
	if err != nil {
		return nil
	}
	bytecode, _ := rlp.EncodeToBytes(runtimeCode)

	data, _ := rlp.EncodeToBytes(task.Data)
	s = append(s, fnType)
	s = append(s, txHash)
	s = append(s, address)
	s = append(s, bytecode)
	s = append(s, data)

	buf := new(bytes.Buffer)
	err = rlp.Encode(buf, s)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func encodeRemoteCall(task *bubble.RemoteCallTask) []byte {
	s := make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(uint16(6000))
	txHash, _ := rlp.EncodeToBytes(task.TxHash)
	caller, _ := rlp.EncodeToBytes(task.Caller)
	Contract, _ := rlp.EncodeToBytes(task.Contract)
	data, _ := rlp.EncodeToBytes(task.Data)
	s = append(s, fnType)
	s = append(s, txHash)
	s = append(s, caller)
	s = append(s, Contract)
	s = append(s, data)

	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, s)
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
		log.Error("failed connect operator node", "err", err)
		return nil, errors.New("failed connect operator node")
	}
	// Construct transaction parameters
	priKey := bp.opPriKey
	// Call the sub-chain system contract RemoteDeploy interface
	toAddr := common.HexToAddress(SubChainSysAddr)
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("Wrong private key", "err", err)
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
		log.Error("Failed to obtain the account nonce", "err", err)
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Failed to get gasPrice", "err", err)
		return nil, err
	}
	value := big.NewInt(0)
	gasLimit := uint64(300000)
	data := encodeRemoteDeploy(task)
	if nil == data {
		return nil, errors.New("encode remoteDeploy transaction failed")
	}
	// Creating transaction objects
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, data)

	// The sender's private key is used to sign the transaction
	chainID, err := client.ChainID(context.Background())
	if err != nil || chainID != task.BubbleID {
		return nil, errors.New("chainID is wrong")
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("Signing remoteDeploy transaction failed", "err", err)
		return nil, err
	}

	// Sending transactions
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Failed to send remoteDeploy transaction", "err", err)
		return nil, err
	}

	hash := signedTx.Hash()
	log.Debug("remoteDeploy tx hash", hash.Hex())
	return hash.Bytes(), nil
}

// HandleRemoteCallTask Handle RemoteDeploy task
func (bp *BubblePlugin) HandleRemoteCallTask(task *bubble.RemoteCallTask) ([]byte, error) {
	if nil == task {
		return nil, errors.New("RemoteCallTask is empty")
	}
	client, err := ethclient.Dial(task.RPC)
	if err != nil || client == nil {
		log.Error("failed connect operator node", "err", err)
		return nil, errors.New("failed connect operator node")
	}
	// Construct transaction parameters
	priKey := bp.opPriKey
	// Call the sub-chain system contract RemoteDeploy interface
	toAddr := common.HexToAddress(SubChainSysAddr)
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("Wrong private key", "err", err)
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
		log.Error("Failed to obtain the account nonce", "err", err)
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Failed to get gasPrice", "err", err)
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
	if err != nil || chainID != task.BubbleID {
		return nil, errors.New("chainID is wrong")
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("Signing remoteCall transaction failed", "err", err)
		return nil, err
	}

	// Sending transactions
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("Failed to send remoteCall transaction", "err", err)
		return nil, err
	}

	hash := signedTx.Hash()
	log.Debug("remoteCall tx hash", hash.Hex())
	return hash.Bytes(), nil
}

func makeGenesisL2(bub *bubble.Bubble) *bubble.GenesisL2 {

	var initNodes []params.InitNode
	for _, node := range bub.Basics.MicroNodes {
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
			ChainID: bub.Basics.BubbleId,
			Frps:    params.DefaultFrpsCfg,
			Cbft: &params.CbftConfig{
				Period:        10000,
				Amount:        10,
				InitialNodes:  params.ConvertNodeUrl(initNodes),
				ValidatorMode: "dpos",
			},
			GenesisVersion: gov.L2Version,
		},
		OpConfig: &bubble.OpConfig{
			MainChain: bub.Basics.OperatorsL1[0],
			SubChain:  bub.Basics.OperatorsL2[0],
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
	rpcUrl := bub.Basics.OperatorsL1[0].RPC
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
	bub, err := bp.GetBubbleInfo(common.ZeroHash, task.BubbleID)
	if err != nil {
		log.Error("failed to get bubble info", "error", err.Error(), "bubbleId", task.BubbleID)
		return errors.New(fmt.Sprintf("failed to get bubble info: %s", err.Error()))
	}

	genesisL2 := makeGenesisL2(bub)
	genesisData, err := json.Marshal(genesisL2)
	if err != nil {
		log.Error("failed to marshal genesis", "error", err.Error())
		return errors.New(fmt.Sprintf("failed to marshal genesis: %s", err.Error()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(bub.Basics.MicroNodes))

	client := &http.Client{}
	data := fmt.Sprintf("{\"type\": %d, \"data\": %s}", bubble.CreateBubble, string(genesisData))

	// make request
	for _, microNode := range bub.Basics.MicroNodes {
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
	time.Sleep(5 * time.Second)
	bub, err := bp.GetBubbleInfo(common.ZeroHash, task.BubbleID)
	if err != nil {
		log.Error("failed to get bubble info", "error", err.Error())
		return errors.New(fmt.Sprintf("failed to get bubble info: %s", err.Error()))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(bub.Basics.MicroNodes))

	client := &http.Client{}
	data := fmt.Sprintf("{\"type\": %d, \"data\": %s}", bubble.ReleaseBubble, bub.Basics.BubbleId)

	// make request
	for _, microNode := range bub.Basics.MicroNodes {
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

// VRFItem is the element of the VRFQueue
type VRFItem struct {
	v interface{}
	x int64
	w *big.Int
}

// VRFQueue Wrap any slice to support VRF
type VRFQueue []*VRFItem

func (vq VRFQueue) Len() int {
	return len(vq)
}

func (vq VRFQueue) Less(i, j int) bool {
	return vq[i].x > vq[j].x // It's actually bigger
}

func (vq VRFQueue) Swap(i, j int) {
	vq[i], vq[j] = vq[j], vq[i]
}

// VRFQueueWrapper Wrap any slice to be VRFQueue
func VRFQueueWrapper(slice interface{}, wrapper func(interface{}) *VRFItem) (VRFQueue, error) {
	// convert slice to an interface queue, to Supports running wrapper
	s := reflect.ValueOf(slice)
	//fmt.Println(kind)
	queue := make([]interface{}, 0)
	if s.Kind() == reflect.Slice {
		for i := 0; i < s.Len(); i++ {
			queue = append(queue, s.Index(i).Interface())
		}
	} else {
		return nil, errors.New("the first parameter must be slice")
	}
	// wrap interface queue to an vrfQueue
	vrfQueue := make(VRFQueue, 0)
	for _, item := range queue {
		if vrfItem := wrapper(item); vrfItem == nil {
			return nil, errors.New("failed to convert the slice element to VRFItem")
		} else {
			vrfQueue = append(vrfQueue, vrfItem)
		}
	}

	return vrfQueue, nil
}

// VRF randomly pick number of elements from vrfQueue, it achieves randomness through the nonces
func VRF(vrfQueue VRFQueue, number uint, curNonce []byte, preNonces [][]byte) (VRFQueue, error) {
	// check params
	if len(curNonce) == 0 || len(preNonces) == 0 || len(vrfQueue) != len(preNonces) {
		log.Error("Failed to VRF", "vrfQueue Size", len(vrfQueue), "curNonceSize", len(curNonce), "preNoncesSize", len(preNonces))
		return nil, errors.New("vrf param is invalid")
	}

	totalWeights := new(big.Int)
	totalSqrtWeights := new(big.Int)
	for _, vrfer := range vrfQueue {
		totalWeights.Add(totalWeights, vrfer.w)
		totalSqrtWeights.Add(totalSqrtWeights, new(big.Int).Sqrt(vrfer.w))
	}

	var maxValue float64 = (1 << 256) - 1
	totalWeightsFloat, err := strconv.ParseFloat(totalWeights.Text(10), 64)
	if nil != err {
		return nil, err
	}
	totalSqrtWeightsFloat, err := strconv.ParseFloat(totalSqrtWeights.Text(10), 64)
	if nil != err {
		return nil, err
	}

	p := xcom.CalcP(totalSqrtWeightsFloat)
	shuffleSeed := new(big.Int).SetBytes(preNonces[0]).Int64()
	log.Debug("Call VRF parameter", "queueSize", len(vrfQueue), "p", p, "totalWeights", totalWeightsFloat, "totalSqrtWeightsFloat",
		totalSqrtWeightsFloat, "number", number, "shuffleSeed", shuffleSeed)

	rd := rand.New(rand.NewSource(shuffleSeed))
	rd.Shuffle(len(vrfQueue), func(i, j int) {
		vrfQueue[i], vrfQueue[j] = vrfQueue[j], vrfQueue[i]
	})

	for i, vrfer := range vrfQueue {
		resultStr := new(big.Int).Xor(new(big.Int).SetBytes(curNonce), new(big.Int).SetBytes(preNonces[i])).Text(10)
		xorValue, err := strconv.ParseFloat(resultStr, 64)
		if nil != err {
			return nil, err
		}

		xorP := xorValue / maxValue
		bd := math.NewBinomialDistribution(vrfer.w.Int64(), p)
		if x, err := bd.InverseCumulativeProbability(xorP); err != nil {
			return nil, err
		} else {
			vrfer.x = x
		}

		log.Debug("Call VRF finished", "index", i, "node", vrfer.v, "curNonce", hex.EncodeToString(curNonce), "preNonce",
			hex.EncodeToString(preNonces[i]), "xorValue", xorValue, "xorP", xorP, "weight", vrfer.w, "x", vrfer.x)
	}

	sort.Sort(vrfQueue)

	return vrfQueue[:number], nil
}

func sendTask(ctx context.Context, waitGroup *sync.WaitGroup, client *http.Client, url string, data string) {

	for i := 0; i < 10; i++ {
		// new request
		dataReader := strings.NewReader(data)
		req, err := http.NewRequest(http.MethodPost, url, dataReader)
		if err != nil {
			log.Error("new http request failed", "err", err)
		}
		req.WithContext(ctx)
		req.Header.Set("Content-Type", "application/json")

		// send request
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != 200 {
			log.Debug("send task to microNode failed", "retry", i, "error", err, "response", resp)
			time.Sleep(time.Duration(3) * time.Second)
			continue
		}
		log.Info("send task to microNode succeed", "ElectronURI", url, "req", req)
		resp.Body.Close()
		waitGroup.Done()
		break
	}
}
