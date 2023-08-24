package plugin

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/common/math"
	"github.com/bubblenet/bubble/common/sort"
	"github.com/bubblenet/bubble/common/vm"
	"github.com/bubblenet/bubble/core"
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
	"github.com/bubblenet/bubble/x/handler"
	"github.com/bubblenet/bubble/x/stakingL2"
	"github.com/bubblenet/bubble/x/xcom"
	"github.com/bubblenet/bubble/x/xutil"
	"golang.org/x/crypto/sha3"
	gomath "math"
	"math/big"
	"math/rand"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"reflect"
	"strconv"
	"sync"
)

const (
	SubChainSysAddr = "0x1000000000000000000000000000000000000020" // Sub chain system contract address
)

var (
	bubblePluginOnce sync.Once
	bubblePlugin     *BubblePlugin
)

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
			db: bubble.NewBubbleDB(),
		}
	})
	return bubblePlugin
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

// GetBubbleInfo return the bubble information by bubble ID
func (bp *BubblePlugin) GetBubbleInfo(blockHash common.Hash, bubbleID *big.Int) (*bubble.Bubble, error) {
	// return bp.db.GetBubbleStore(blockHash, bubbleID)
	// get bubble basics
	basics, err := bp.GetBubBasics(blockHash, bubbleID)
	if nil == basics || err != nil {
		return nil, errors.New("failed to get bubble basics information")
	}
	var bub bubble.Bubble
	bub.Basics = basics
	// get bubble state
	state, err := bp.GetBubState(blockHash, bubbleID)
	if nil == state || err != nil {
		return nil, errors.New("failed to get bubble state")
	}
	bub.State = *state
	// get bubble txHashList
	// get StakingToken Hash List
	stTxHashList, err := bp.GetTxHashListByBub(blockHash, bubbleID, bubble.StakingToken)
	if snapshotdb.NonDbNotFoundErr(err) {
		return nil, errors.New("failed to get bubble StakingToken transaction hash list")
	}
	bub.StakingTokenTxHashList = stTxHashList

	// get WithdrewToken Hash List
	wdTxHashList, err := bp.GetTxHashListByBub(blockHash, bubbleID, bubble.WithdrewToken)
	if snapshotdb.NonDbNotFoundErr(err) {
		return nil, errors.New("failed to get bubble WithdrewToken transaction hash list")
	}
	bub.WithdrewTokenTxHashList = wdTxHashList

	// get SettleBubble Hash List
	sbTxHashList, err := bp.GetTxHashListByBub(blockHash, bubbleID, bubble.SettleBubble)
	if snapshotdb.NonDbNotFoundErr(err) {
		return nil, errors.New("failed to get bubble SettleBubble transaction hash list")
	}
	bub.SettleBubbleTxHashList = sbTxHashList

	return &bub, nil
}

// GetBubBasics return the bubble basics by bubble ID
func (bp *BubblePlugin) GetBubBasics(blockHash common.Hash, bubbleID *big.Int) (*bubble.BubBasics, error) {
	return bp.db.GetBubBasics(blockHash, bubbleID)
}

// GetBubState return the bubble state by bubble ID
func (bp *BubblePlugin) GetBubState(blockHash common.Hash, bubbleID *big.Int) (*bubble.BubState, error) {
	return bp.db.GetBubState(blockHash, bubbleID)
}

// CreateBubble run the non-business logic to create bubble
func (bp *BubblePlugin) CreateBubble(blockHash common.Hash, blockNumber *big.Int, from common.Address, nonce uint64, parentHash common.Hash) (*bubble.Bubble, error) {

	// get the nonces of the historical block
	preNonces, err := handler.GetVrfHandlerInstance().Load(parentHash)
	if err != nil {
		return nil, err
	}
	maxLen := int(gomath.Max(gomath.Max(bubble.OperatorL1Size, bubble.OperatorL2Size), bubble.CommitteeSize))
	if len(preNonces) < maxLen {
		newNonces := make([][]byte, maxLen)
		newNonces = append(newNonces, preNonces...)
		preNonces = newNonces
	}

	// elect the operatorsL1 by VRF
	OperatorsL1, err := bp.ElectOperatorL1(blockHash, bubble.OperatorL1Size, common.Uint64ToBytes(nonce), preNonces)
	if err != nil {
		return nil, err
	}

	// elect the operatorsL2 by VRF
	candidateL2, err := bp.ElectCandidateL2(blockHash, bubble.OperatorL2Size, blockNumber, common.Uint64ToBytes(nonce), preNonces)
	if err != nil {
		return nil, err
	}
	var OperatorsL2 []*bubble.Operator
	for _, can := range candidateL2 {
		operator := &bubble.Operator{
			NodeId: can.NodeId,
			RPC:    can.ElectronURI,
			OpAddr: can.StakingAddress,
		}
		OperatorsL2 = append(OperatorsL2, operator)
	}

	// // elect the microNodesL2 by VRF
	microNodes, err := bp.ElectBubbleMicroNodes(blockHash, blockNumber, bubble.CommitteeSize, common.Uint64ToBytes(nonce), preNonces)
	if err != nil {
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
		BubbleId:    bubbleID,
		Creator:     from,
		CreateBlock: blockNumber.Uint64(),
		OperatorsL1: OperatorsL1,
		OperatorsL2: OperatorsL2,
		MicroNodes:  microNodes,
	}

	bub := &bubble.Bubble{
		Basics: basics,
		State:  bubble.ActiveStatus,
	}

	// store bubble basics
	if err := bp.db.StoreBubBasics(blockHash, basics.BubbleId, basics); err != nil {
		log.Error("Failed to CreateBubble on bubblePlugin: Store bubble basics failed",
			"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "bubbleId", basics.BubbleId, "err", err)
		return nil, err
	}
	// store bubble state
	if err := bp.db.StoreBubState(blockHash, basics.BubbleId, bub.State); err != nil {
		log.Error("Failed to CreateBubble on bubblePlugin: Store bubble state failed",
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
	if len(operators) == 0 {
		return nil, bubble.ErrOperatorL1IsInsufficient
	}
	// wrap the operators to the VRF able queue
	vrfQueue, err := VRFQueueWrapper(operators, func(item interface{}) *VRFItem {
		return &VRFItem{
			v: item,
			w: new(big.Int).SetInt64(10000),
		}
	})
	if err != nil {
		return nil, err
	}
	// VRF Elect
	electedVrfQueue, err := VRF(vrfQueue, operatorNumber, curNonce, preNonces[:operatorNumber])
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

// ElectCandidateL2 Elect the Layer2 Operator nodes for the bubble chain by VRF
func (bp *BubblePlugin) ElectCandidateL2(blockHash common.Hash, operatorNumber uint, blockNumber *big.Int, curNonce []byte, preNonces [][]byte) (bubble.CandidateQueue, error) {
	operatorQueue, err := bp.stk2Plugin.GetOperatorList(blockHash, blockNumber.Uint64())
	if err != nil {
		return nil, err
	}
	if len(operatorQueue) == 0 {
		return nil, bubble.ErrOperatorL2IsInsufficient
	}
	// wrap the operators to the VRF able queue
	vrfQueue, err := VRFQueueWrapper(operatorQueue, func(item interface{}) *VRFItem {
		if candidate, ok := (item).(stakingL2.Candidate); ok {
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
	electedVrfQueue, err := VRF(vrfQueue, operatorNumber, curNonce, preNonces[:operatorNumber])
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
func (bp *BubblePlugin) ElectBubbleMicroNodes(blockHash common.Hash, blockNumber *big.Int, committeeNumber uint, curNonce []byte, preNonces [][]byte) (bubble.CandidateQueue, error) {
	candidateQueue, err := bp.stk2Plugin.GetCandidateList(blockHash, blockNumber.Uint64())
	if err != nil {
		return nil, err
	}
	if len(candidateQueue) == 0 {
		return nil, bubble.ErrMicroNodeIsInsufficient
	}
	// wrap the candidates to the VRF able queue
	vrfQueue, err := VRFQueueWrapper(candidateQueue, func(item interface{}) *VRFItem {
		if candidate, ok := (item).(stakingL2.Candidate); ok {
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
	electedVrfQueue, err := VRF(vrfQueue, committeeNumber, curNonce, preNonces[:committeeNumber])
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
		if err := bp.stk2Plugin.db.DelCandidateStore(blockHash, addr); err != nil {
			return nil, err
		}
	}

	return committees, nil
}

// ReleaseBubble run the non-business logic to release the bubble
func (bp *BubblePlugin) ReleaseBubble(blockHash common.Hash, blockNumber *big.Int, bubbleID *big.Int) error {
	bub, err := bp.GetBubbleInfo(blockHash, bubbleID)
	if err != nil {
		return err
	}
	// release the operator nodes of the L2 to the database
	for _, operator := range bub.Basics.OperatorsL2 {
		addr, err := xutil.NodeId2Addr(operator.NodeId)
		if err != nil {
			return err
		}
		// check whether the node has been withdrawn
		can, err := bp.stk2Plugin.db.GetCandidateStore(blockHash, addr)
		if can != nil || err == snapshotdb.ErrNotFound {
			break
		}
		if err := bp.stk2Plugin.db.SetOperatorStore(blockHash, addr, can); nil != err {
			log.Error("Failed to SetOperatorStore on ReleaseBubble: Store Operator info is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", operator.NodeId.String(), "err", err)
			return err
		}
	}
	// release the candidate nodes of the L2 to the database
	for _, candidate := range bub.Basics.MicroNodes {
		addr, err := xutil.NodeId2Addr(candidate.NodeId)
		if err != nil {
			return err
		}
		// check whether the node has been withdrawn
		can, err := bp.stk2Plugin.db.GetCandidateStore(blockHash, addr)
		if can != nil || err == snapshotdb.ErrNotFound {
			break
		}
		if err := bp.stk2Plugin.db.SetCommitteeStore(blockHash, addr, candidate); nil != err {
			log.Error("Failed to SetCandidateStore on ReleaseBubble: Store Candidate info is failed",
				"blockNumber", blockNumber.Uint64(), "blockHash", blockHash.Hex(), "nodeId", candidate.NodeId.String(), "err", err)
			return err
		}
	}

	if err := bp.db.DelBubbleStore(blockHash, bubbleID); err != nil {
		return err
	}

	return nil
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
	if err := bp.eventMux.Post(task); nil != err {
		log.Error("post CreateBubble task failed", "err", err)
		return err
	}

	return nil
}

// PostReleaseBubbleEvent Send the release bubble event and wait for the task to be processed
func (bp *BubblePlugin) PostReleaseBubbleEvent(task *bubble.ReleaseBubbleTask) error {
	if err := bp.eventMux.Post(task); nil != err {
		log.Error("post ReleaseBubble task failed", "err", err)
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
	log.Debug("mintToken tx hash=========================================", hash.Hex())
	return hash.Bytes(), nil
}

func makeGenesisL2(bub *bubble.Bubble) *core.GenesisL2 {

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

	genesisL2 := &core.GenesisL2{
		Config: &params.ChainConfig{
			ChainID: bub.Basics.BubbleId,
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
		EconomicModel: xcom.GetEc(xcom.DefaultMainNet),
		Nonce:         []byte{},
		Timestamp:     params.MainNetGenesisTimestamp,
		ExtraData:     []byte{},
		GasLimit:      math.HexOrDecimal64(params.GenesisGasLimit),
		Coinbase:      common.ZeroAddr,
		Alloc: core.GenesisAlloc{
			vm.RewardManagerPoolAddr: {Balance: rewardMgrPoolIssue},
			generalAddr:              {Balance: generalBalance},
		},
		Number:     0,
		GasUsed:    0,
		ParentHash: common.ZeroHash,
	}

	return genesisL2
}

// HandleCreateBubbleTask Handle create bubble task
func (bp *BubblePlugin) HandleCreateBubbleTask(task *bubble.CreateBubbleTask) error {
	if task == nil {
		return errors.New("CreateBubbleTask is empty")
	}

	bub, err := bp.GetBubbleInfo(common.ZeroHash, task.BubbleID)
	if err != nil {
		log.Error(fmt.Sprintf("failed to get bubble info: %s", err.Error()))
		return errors.New(fmt.Sprintf("failed to get bubble info: %s", err.Error()))
	}
	genesisL2 := makeGenesisL2(bub)
	args, err := genesisL2.MarshalJSON()
	if err != nil {
		log.Warn(fmt.Sprintf("failed to marshal genesis: %s", err.Error()))
		return errors.New(fmt.Sprintf("failed to marshal genesis: %s", err.Error()))
	}

	for _, operator := range bub.Basics.OperatorsL2 {
		conn, err := net.Dial("http", operator.RPC)
		if err != nil {
			log.Warn(fmt.Sprintf("failed to dial rpc: %s", err.Error()))
			return errors.New(fmt.Sprintf("failed to dial rpc: %s", err.Error()))
		}
		client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

		// TODO: asynchronous process and check whether the network is built
		var result string
		err = client.Call("", args, result)
		if err != nil {
			log.Warn(fmt.Sprintf("failed to call remote function: %s", err.Error()))
			return errors.New(fmt.Sprintf("failed to call remote function: %s", err.Error()))
		}
		if result != "" {
			log.Warn(fmt.Sprintf("call remote function result error: %s", result))
			return errors.New(fmt.Sprintf("call remote function result error: %s", result))
		}
	}
	return nil
}

// HandleReleaseBubbleTask Handle release bubble task
func (bp *BubblePlugin) HandleReleaseBubbleTask(task *bubble.ReleaseBubbleTask) error {
	if task == nil {
		return errors.New("ReleaseBubbleTask is empty")
	}

	bub, err := bp.GetBubbleInfo(common.ZeroHash, task.BubbleID)
	if err != nil {
		log.Error(fmt.Sprintf("failed to get bubble info: %s", err.Error()))
		return errors.New(fmt.Sprintf("failed to get bubble info: %s", err.Error()))
	}

	for _, operator := range bub.Basics.OperatorsL2 {
		conn, err := net.Dial("http", operator.RPC)
		if err != nil {
			log.Warn(fmt.Sprintf("failed to dial rpc: %s", err.Error()))
			return errors.New(fmt.Sprintf("failed to dial rpc: %s", err.Error()))
		}
		client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
		args := bub.Basics.BubbleId

		// TODO: asynchronous process and check whether the network is built
		var reply string
		err = client.Call("", args, reply)
		if err != nil {
			log.Warn(fmt.Sprintf("failed to call remote function: %s", err.Error()))
			return errors.New(fmt.Sprintf("failed to call remote function: %s", err.Error()))
		}
		if reply != "" {
			log.Warn(fmt.Sprintf("call remote function result error: %s", reply))
			return errors.New(fmt.Sprintf("call remote function result error: %s", reply))
		}
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
	queue := make([]interface{}, 0)
	if s := reflect.ValueOf(slice); s.Kind() == reflect.Slice {
		for i := 0; i < s.Len(); i++ {
			queue[i] = s.Index(i).Interface()
		}
	} else {
		return nil, errors.New("the first parameter must be slice")
	}
	// wrap interface queue to an vrfQueue
	vrfQueue := make(VRFQueue, 0)
	for item := range queue {
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
		xorValue := float64(new(big.Int).Xor(new(big.Int).SetBytes(curNonce), new(big.Int).SetBytes(preNonces[i])).Int64())
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
