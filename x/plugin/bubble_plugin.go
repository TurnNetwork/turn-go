package plugin

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/ethclient"
	"github.com/bubblenet/bubble/event"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/x/bubble"
	"github.com/bubblenet/bubble/x/xcom"
	"math/big"
	"sync"
)

const (
	remoteBubbleContract = "0x2000000000000000000000000000000000000001"
)

var (
	bubblePluginOnce sync.Once
	bp               *BubblePlugin
)

type BubblePlugin struct {
	db       *bubble.BubbleDB
	ChainID  *big.Int
	eventMux *event.TypeMux
}

func BubbleInstance() *BubblePlugin {
	bubblePluginOnce.Do(func() {
		log.Info("Init bubble plugin ...")
		bp = &BubblePlugin{
			db: bubble.NewBubbleDB(),
		}
	})
	return bp
}

func (bp *BubblePlugin) BeginBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	return nil
}

func (bp *BubblePlugin) EndBlock(blockHash common.Hash, header *types.Header, state xcom.StateDB) error {
	return nil
}

func (bp *BubblePlugin) Confirmed(nodeId discover.NodeID, block *types.Block) error {
	return nil
}

func (bp *BubblePlugin) SetEventMux(eventMux *event.TypeMux) {
	bp.eventMux = eventMux
}

func (bp *BubblePlugin) GetBubContracts(blockHash common.Hash) ([]*common.Address, error) {
	iter := bp.db.IteratorBubContract(blockHash, 0)
	if err := iter.Error(); nil != err {
		return nil, err
	}
	defer iter.Release()

	queue := make([]*common.Address, 0)
	for iter.Valid(); iter.Next(); {
		data := iter.Value()
		address := new(common.Address)
		address.SetBytes(data)
		queue = append(queue, address)
	}

	return queue, nil
}

func (bp *BubblePlugin) GetBubContract(blockHash common.Hash, address *common.Address) (*common.Address, error) {
	addr, err := bp.db.GetBubContract(blockHash, address)
	if err != nil {
		return nil, err
	}

	return addr, nil
}

func (bp *BubblePlugin) DelBubContract(blockHash common.Hash, address *common.Address) error {
	err := bp.db.DelBubContract(blockHash, address)
	if err != nil {
		return err
	}

	return nil
}

func (bp *BubblePlugin) PostRemoteCallTask(task *bubble.RemoteCallTask) error {
	if err := bp.eventMux.Post(*task); nil != err {
		log.Error("post remoteCallTask failed", "err", err)
		return err
	}

	return nil
}

// HandleRemoteCallTask Handle RemoteCall task
func (bp *BubblePlugin) HandleRemoteCallTask(task *bubble.RemoteCallTask) ([]byte, error) {
	if nil == task {
		return nil, errors.New("RemoteCallTask is empty")
	}

	// get token plugin
	tp := TokenInstance()

	client, err := ethclient.Dial(tp.OpConfig.MainChain.Rpc)
	if err != nil || client == nil {
		log.Error("failed connect operator node", "err", err)
		return nil, errors.New("failed connect operator node")
	}

	// Construct transaction parameters
	priKey := tp.OpConfig.GetSubOpPriKey()
	toAddr := common.HexToAddress(remoteBubbleContract)
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		log.Error("wrong private key", "err", err)
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("the private key to the public key failed")
		return nil, errors.New("the private key to the public key failed")
	}
	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	if fromAddr != tp.OpConfig.SubChain.OpAddr {
		log.Error("the key is not the layer2 operator")
		return nil, errors.New("the key is not the layer2 operator")
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		log.Error("failed to obtain the account nonce", "err", err)
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil || chainID != task.BubbleID {
		return nil, errors.New("chainID is wrong")
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Failed to get gasPrice", "err", err)
		return nil, err
	}
	value := big.NewInt(0)
	gasLimit := uint64(300000)

	// encode transaction data
	data := encodeRemoteCall(task)
	if data == nil {
		return nil, errors.New("encode remoteCall transaction failed")
	}

	// Creating transaction objects
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Error("sign remoteCall transaction failed", "err", err)
		return nil, err
	}

	// Sending transactions
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("failed to send remoteCall transaction", "err", err)
		return nil, err
	}
	hash := signedTx.Hash()
	log.Debug("send remoteCall transaction succeed", "hash", hash.Hex())
	return hash.Bytes(), nil
}

func encodeRemoteCall(task *bubble.RemoteCallTask) []byte {
	queue := make([][]byte, 0)

	fnType, _ := rlp.EncodeToBytes(uint16(8007))
	bubbleID, _ := rlp.EncodeToBytes(task.BubbleID)
	txHash, _ := rlp.EncodeToBytes(task.TxHash)
	caller, _ := rlp.EncodeToBytes(task.Caller)
	Contract, _ := rlp.EncodeToBytes(task.Contract)
	data, _ := rlp.EncodeToBytes(task.Data)
	queue = append(queue, fnType)
	queue = append(queue, bubbleID)
	queue = append(queue, txHash)
	queue = append(queue, caller)
	queue = append(queue, Contract)
	queue = append(queue, data)

	buf := new(bytes.Buffer)
	err := rlp.Encode(buf, queue)
	if err != nil {
		return nil
	}
	return buf.Bytes()
}
