package tests

import (
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/datavalidator/types"
	"github.com/bubblenet/bubble/datavalidator/wallet"
	p2p2 "github.com/bubblenet/bubble/p2p"
	"github.com/bubblenet/bubble/p2p/discover"
	"testing"
)

type Node struct {
	v      *types.Validator
	sk     wallet.Wallet
	nodeId discover.NodeID
	peer1  *p2p2.Peer
	rw1    p2p2.MsgReadWriter
	peer2  *p2p2.Peer
	rw2    p2p2.MsgReadWriter
}

func MakeNode(id discover.NodeID, index uint64, protocols []p2p2.Protocol) *Node {
	blsKey := bls.GenerateKey()
	sk, _ := crypto.GenerateKey()
	nodeId := discover.PubkeyID(&sk.PublicKey)
	peer1, rw1, peer2, rw2 := p2p2.NewPeerByNodeID(nodeId, id, protocols)
	return &Node{
		v: &types.Validator{
			Index:       index,
			Address:     crypto.PubkeyToAddress(sk.PublicKey),
			P2pUrl:      nodeId.String(),
			BlockNumber: 0,
		},
		sk:     wallet.FromSk(blsKey),
		nodeId: nodeId,
		peer1:  peer1,
		rw1:    rw1,
		peer2:  peer2,
		rw2:    rw2,
	}
}

func TestFlow(t *testing.T) {
	//details := []*types.MessagePublishedDetail{
	//	&types.MessagePublishedDetail{
	//		BlockHash: common.BigToHash(big.NewInt(1)),
	//		TxHash:    common.BigToHash(big.NewInt(2)),
	//		Log: &types.MessagePublished{
	//			Send:     common.BigToAddress(big.NewInt(11)),
	//			ChainId:  102,
	//			Sequence: 1,
	//			Nonce:    1,
	//			Payload:  []byte{1, 2, 3},
	//		},
	//		Signatures: []*types.Signature{},
	//	},
	//	&types.MessagePublishedDetail{
	//		BlockHash: common.BigToHash(big.NewInt(3)),
	//		TxHash:    common.BigToHash(big.NewInt(4)),
	//		Log: &types.MessagePublished{
	//			Send:     common.BigToAddress(big.NewInt(11)),
	//			ChainId:  102,
	//			Sequence: 2,
	//			Nonce:    2,
	//			Payload:  []byte{4, 5, 6},
	//		},
	//	},
	//}
	//
	//sk, err := crypto.GenerateKey()
	//require.Nil(t, err)
	//owner := crypto.PubkeyToAddress(sk.PublicKey)
	//filterMessage := &mock.MockFilterMessage{
	//	RangeFilterFn: func(ctx context.Context, begin, to *big.Int) ([]*types.MessagePublishedDetail, error) {
	//		if to.Uint64() > 10 {
	//			return nil, nil
	//		}
	//		return []*types.MessagePublishedDetail{details[0]}, nil
	//	},
	//	HasLogOnChainFn: func(txHash common.Hash, sequence uint64) *types.MessagePublishedDetail {
	//		if txHash == details[1].TxHash && sequence == details[1].Log.Sequence {
	//			return details[1]
	//		}
	//		return nil
	//	},
	//}
	//mdb := db.NewMemoryValidatorDB()
	//validatorContract := &mock.MockValidator{
	//	GetValidatorFn: func(addr common.Address) *types.Validator {
	//		return nil
	//	},
	//	ValidatorSetFn: func() []*types.Validator {
	//		return nil
	//	},
	//	IsValidatorFn: func(id string) bool {
	//		return false
	//	},
	//}
	//childChainContract := &mock.MockChildChain{
	//	GetBubbleIdFn: func() []uint64 {
	//		return []uint64{102}
	//	},
	//	GetValidatorsFn: func(chainId uint64) map[uint64]*types.Validator {
	//		return nil
	//	},
	//	GetBubbleValidatorFn: func(chainId uint64) *types.ChildChainValidator {
	//		if chainId == 102 {
	//			return &types.ChildChainValidator{
	//				Validators: validatorContract.ValidatorSet(),
	//				Threshold:  len(validatorContract.ValidatorSet()),
	//			}
	//		}
	//		return nil
	//	},
	//}
	//blockState := mock.MockBlockState{
	//	BlockNumberFn: func() uint64 {
	//		return 10
	//	},
	//}
	//newLog := make(chan []*types.MessagePublishedDetail, 1)
	//newValidator := make(chan types.ValidatorSets, 1)
	//sync := sync.NewSync(owner, validatorContract, childChainContract, mdb, blockState, filterMessage, newLog, newValidator)
	//network := p2p.NewNetwork(wallet.FromSk(sk), childChainContract, &mock.DataCheck{
	//	DB:            mdb,
	//	FilterMessage: filterMessage,
	//}, blockState, nil)
	//
	//ownNodeId := discover.PubkeyID(&sk.PublicKey)
	//ownvalidator := &types.Validator{
	//	Index:       1,
	//	Address:     owner,
	//	P2pUrl:      ownNodeId.String(),
	//	BlockNumber: 0,
	//}
	//node2 := MakeNode(ownNodeId, 2, []p2p2.Protocol{network.Protocols()})
	//
	//childChainContract.GetValidatorsFn = func(chainId uint64) map[uint64]*types.Validator {
	//	if chainId == 101 {
	//		return map[uint64]*types.Validator{
	//			ownvalidator.Index: ownvalidator,
	//			node2.v.Index:      node2.v,
	//		}
	//	}
	//	return nil
	//}
	//
	//validatorContract.ValidatorSetFn = func() []*types.Validator {
	//	return []*types.Validator{ownvalidator, node2.v}
	//}
	//
	//validatorContract.IsValidatorFn = func(id string) bool {
	//	return true
	//}
	//mdb.StoreScanLog(0)
	//_ = func() {
	//	sync.RefreshValidator()
	//	sync.HandleMessage(context.Background())
	//	res := <-newLog
	//	require.Equal(t, details[0].BlockHash, res[0].BlockHash)
	//	fmt.Println(res[0].BlockHash.Hex())
	//}
	//go network.Protocols().Run(node2.peer2, node2.rw2)
	//process := process.NewProcess(wallet.FromSk(sk), mdb, sync, newValidator, newLog, network, network)
	//process.Run(context.Background())
	//
	////msg, err := node2.rw1.ReadMsg()
	////require.Nil(t, err)
	////require.NotNil(t, msg)
	//var signMsg p2p.SignMessageMsg
	//msg.Decode(&signMsg)
	//require.NotNil(t, signMsg.SignMessageData.Signatures)
	//
	//sign := node2.sk.Sign(signMsg.SignMessageData.Hash())
	//go p2p2.Send(node2.rw1, p2p.SignedObservationType, &p2p.SignedObservation{
	//	ChainId: signMsg.SignMessageData.Log.ChainId,
	//	ID:      signMsg.SignMessageData.Log.Hash(),
	//	Signatures: []*types.Signature{&types.Signature{
	//		Index:     node2.v.Index,
	//		Signature: sign,
	//	}},
	//})
	//msg, err = node2.rw1.ReadMsg()
	//require.Nil(t, err)
	//require.NotNil(t, msg)
	//require.True(t, p2p.SignedWithQuorumType == msg.Code)
	//fmt.Println("msg code", msg.Code, "msg.size", msg.Size)
	//go p2p2.Send(node2.rw1, p2p.SignedObservationRequestType, &p2p.SignedObservationRequest{
	//	ChainId: signMsg.SignMessageData.Log.ChainId,
	//	ID:      signMsg.SignMessageData.Log.Hash(),
	//})
	//msg, err = node2.rw1.ReadMsg()
	//require.Nil(t, err)
	//require.NotNil(t, msg)
	//require.True(t, p2p.SignedObservationType == msg.Code)
	//fmt.Println("msg code", msg.Code, "msg.size", msg.Size)
	//
	//ctx := context.WithoutCancel(context.Background())
	//go network.HeartbeatLoop(ctx, time.Second)
	//msg, err = node2.rw1.ReadMsg()
	//require.Nil(t, err)
	//require.NotNil(t, msg)
	//require.True(t, p2p.HeartbeatMessageType == msg.Code)
	//fmt.Println("msg code", msg.Code, "msg.size", msg.Size)

}
