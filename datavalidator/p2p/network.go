package p2p

import (
	"context"
	"errors"
	"github.com/bubblenet/bubble/datavalidator/wallet"
	"github.com/bubblenet/bubble/metrics"
	"github.com/bubblenet/bubble/rlp"
	"math"
	"sync"
	"time"

	"github.com/bubblenet/bubble/common"
	datacommon "github.com/bubblenet/bubble/datavalidator/common"
	"github.com/bubblenet/bubble/datavalidator/types"

	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/p2p"
	"github.com/bubblenet/bubble/p2p/discover"
)

const (
	Version = 1
)

var (
	p2pmsg = metrics.NewRegisteredMeter("datavalidator/meter/p2p/p2pmsg", nil)
)

type Validators interface {
	GetValidators(chainId uint64) map[uint64]*types.Validator
}

type DataCheck interface {
	GetScanLog() (uint64, error)
	SetQuorumLog(detail []*types.MessagePublishedDetail) error
	UpdateMessagePublished(log *types.MessagePublishedDetail) error
	GetMessagePublished(id common.Hash) (*types.MessagePublishedDetail, error)
	//HasLogOnChain(txHash common.Hash, sequence uint64) *types.MessagePublishedDetail
}

type Network struct {
	mutex      sync.Mutex
	dbMutex    sync.Mutex
	log        log.Logger
	wallet     wallet.Wallet
	sets       types.ValidatorSets
	validators Validators
	dataCheck  DataCheck
	blockState types.BlockState
	peerSet    *peerSet
	p2pServer  types.P2PServer
}

func NewNetwork(wallet wallet.Wallet, validators Validators, dataCheck DataCheck, blockState types.BlockState, p2pServer types.P2PServer) *Network {
	network := &Network{
		log:        log.New(),
		wallet:     wallet,
		validators: validators,
		dataCheck:  dataCheck,
		blockState: blockState,
		peerSet:    newPeerSet(),
		sets:       types.NewValidatorSets(),
		p2pServer:  p2pServer,
	}
	if wallet != nil {
		network.log = log.New("self", datacommon.BlsID(wallet.PublicKey()))
	}
	return network
}

func (n *Network) PeersInfo() []*types.PeerInfo {
	var peersInfo []*types.PeerInfo
	peers := n.peerSet.Peers()
	for _, p := range peers {
		peersInfo = append(peersInfo, &types.PeerInfo{
			ID:        p.id,
			ScanBlock: p.scanBlock,
		})
	}
	return peersInfo
}
func (n *Network) SetP2PServer(server types.P2PServer) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.p2pServer = server
}
func (n *Network) SetValidatorSet(sets types.ValidatorSets) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.sets = sets
	n.TryConnect()
}
func (n *Network) TryConnect() {
	owner := n.wallet.PublicKey()
	if n.p2pServer == nil {
		return
	}
	for _, group := range n.sets {
		for id, v := range group.Group {
			if n.peerSet.Peer(id) == nil {
				node, err := discover.ParseNode(v.P2pUrl)
				if err != nil || v.BlsPubKey.IsEqual(owner) {
					continue
				}
				n.log.Debug("add p2p node", "peer", v.P2pUrl)
				go n.p2pServer.AddPeer(node)
			}
		}
	}
}
func (n *Network) Run(ctx context.Context) {
	if n.wallet != nil {
		go n.HeartbeatLoop(ctx, 30*time.Second)
	}
}
func (n *Network) HeartbeatLoop(ctx context.Context, duration time.Duration) {
	tick := time.NewTicker(duration)
	bootTimestamp := uint64(time.Now().UnixNano())
	counter := uint64(0)
	for {
		counter += 1
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			scanBlock, _ := n.dataCheck.GetScanLog()
			for _, v := range n.peerSet.Peers() {
				n.log.Debug("send heartbeat", "peer", v.id)
				n.Send(v.id, &Heartbeat{
					Counter:       counter,
					Timestamp:     uint64(time.Now().UnixNano()),
					Version:       "1",
					BlsPub:        datacommon.BlsID(n.wallet.PublicKey()),
					BootTimestamp: bootTimestamp,
					ScanBlock:     scanBlock,
				})
			}
		}
	}

}
func (n *Network) Protocols() p2p.Protocol {
	return p2p.Protocol{
		Name:    "data_validator",
		Version: Version,
		Length:  10,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			return n.handle(p, rw)
		},
		//NodeInfo: func() interface{} {
		//	return nil
		//},
		//PeerInfo: func(id discover.NodeID) interface{} {
		//	if p := n.peerSet.Peer(fmt.Sprintf("%x", id[:8])); p != nil {
		//		return p.Info()
		//	}
		//	return nil
		//},
	}
}
func (n *Network) handle(p *p2p.Peer, rw p2p.MsgReadWriter) error {
	//if !n.validators.IsValidator(fmt.Sprintf("%x", p.ID().Bytes()[:8])) {
	//	return nil
	//}
	//TODO handshake
	peer := newPeer(p, rw)
	n.peerSet.Register(peer)
	defer n.peerSet.Unregister(peer.id)
	log.Debug("receive new peer", "peer", p.ID().TerminalString(), "len:", n.peerSet.Len())

	for {
		if err := n.handleMsg(peer); err != nil {
			n.log.Warn("handle msg", "err", err)
			p.Log().Error("datavalidator message handling failed", "err", err)
			return err
		}
	}
}
func (n *Network) handleMsg(p *peer) error {
	msg, err := p.rw.ReadMsg()
	n.log.Debug("read message", "id", p.id)
	if err != nil {
		p.Log().Error("read peer message error", "err", err)
		return err
	}
	p2pmsg.Mark(1)
	if n.wallet == nil {
		return nil
	}
	// if msg.Size > ProtocolMaxMsgSize {
	// 	return errResp(ErrMsgTooLarge, "%v > %v", msg.Size, ProtocolMaxMsgSize)
	// }
	defer msg.Discard()

	switch {
	case msg.Code == SignMessageType:
		var sign SignMessageMsg
		if err := msg.Decode(&sign); err != nil {
			return err
		}
		return n.handleSignMessageType(p, &sign)
	case msg.Code == HeartbeatMessageType:
		var so Heartbeat
		if err := msg.Decode(&so); err != nil {
			return err
		}
		return n.handleHeartbeatMessageType(p, &so)
	case msg.Code == SignedObservationRequestType:
		var so SignedObservationRequest
		if err := msg.Decode(&so); err != nil {
			return err
		}
		return n.handleSignedObservationRequestType(p, &so)
	case msg.Code == SignedObservationType:
		var so SignedObservation
		if err := msg.Decode(&so); err != nil {
			return err
		}
		return n.handleSignedObservationType(p, &so)
	case msg.Code == SignedWithQuorumType:
		var so SignedWithQuorum
		if err := msg.Decode(&so); err != nil {
			return err
		}
		return n.handleSignedWithQuorumType(p, &so)
	}
	return nil
}

//func (n *Network) handleHandshakeType(p *peer, rw p2p.MsgReadWriter) error {
//	scanBlock, _ := n.dataCheck.GetScanLog()
//	handshakeMsg := &Handshake{
//		Version:     Version,
//		BlockNumber: n.blockState.BlockNumber(),
//		ScanBlock:   scanBlock,
//	}
//	werr := make(chan error, 1)
//	go func() {
//		werr <- p2p.Send(p.rw, HandshakeType, handshakeMsg)
//	}()
//	var msg *Handshake
//	var err error
//	if msg, err = n.readHandshake(p.rw); err != nil {
//		<-werr
//		return err
//	}
//	if err := <-werr; err != nil {
//		return fmt.Errorf("write error: %v", err)
//	}
//	log.Info("handshake success", "peer", p.id, "blockNumber", msg.BlockNumber, "scan", msg.ScanBlock, "version", msg.Version)
//	return nil
//}

//func (n *Network) readHandshake(rw p2p.MsgReader) (*Handshake, error) {
//	msg, err := rw.ReadMsg()
//	if err != nil {
//		return nil, err
//	}
//	if msg.Code != HandshakeType {
//		return nil, errors.New("invalid message type")
//	}
//
//	var hmsg Handshake
//	err = msg.Decode(&hmsg)
//	if err != nil {
//		return nil, err
//	}
//	return &hmsg, nil
//}

func (n *Network) handleHeartbeatMessageType(p *peer, msg *Heartbeat) error {
	n.log.Debug("handle heartbeat message", "peer", p.id)
	p.scanBlock = msg.ScanBlock
	return nil
}

func (n *Network) handleSignMessageType(p *peer, msg *SignMessageMsg) error {
	log := n.log.New("peer", p.id)
	log.Debug("handle sign message", "msg", msg.SignMessageData.String())
	id := msg.SignMessageData.Hash()

	validator := n.validators.GetValidators(msg.SignMessageData.Log.ChainId)
	if validator == nil {
		log.Warn("invalid chainId in validators")
		return errors.New("invalid chainId in validators")
	}
	v := validator[msg.Signature.Index]
	if v == nil {
		log.Warn("not found validator in validator group")
		return errors.New("not found validator in validator group")
	}
	success := datacommon.VerifySignature(id.Bytes(), msg.Signature, v.BlsPubKey)
	if success == false {
		log.Warn("verify sign message failed", "peer", p.String())
		return errors.New("invalid signature")
	}
	var responseMsg interface{}
	var group *types.PeerGroup
	err := func() error {
		n.dbMutex.Lock()
		defer n.dbMutex.Unlock()
		detail, err := n.dataCheck.GetMessagePublished(id)
		if err != nil {
			return err
		}
		log.Debug("get message detail", "id", id)

		if detail != nil {
			group = n.PeerGroup(detail.Log.ChainId)
			if group == nil {
				return errors.New("group is invalid")
			}
			if err := VerifySignatures(detail.Hash().Bytes(), msg.SignMessageData.Signatures, group); err != nil {
				return err
			}
			log.Debug("verify sign message success", "id", id, "current sign", len(detail.Signatures), "msg sign", len(msg.SignMessageData.Signatures))
			//print signature
			//for _, d := range detail.Signatures {
			//	log.Debug("detail signature", "index", d.Index)
			//}
			//for _, d := range msg.SignMessageData.Signatures {
			//	log.Debug("msg signature", "index", d.Index)
			//
			//}
			//
			detail.AddSignature(msg.SignMessageData.Signatures)
			if len(detail.Signatures) >= group.Threshold {
				n.log.Debug("sign message had quorum", "id", id)
				n.dataCheck.SetQuorumLog([]*types.MessagePublishedDetail{detail})
				//broadcast quorum message
				responseMsg = &SignedWithQuorum{
					ChainId:    detail.Log.ChainId,
					ID:         detail.Hash(),
					Signatures: detail.Signatures,
				}

			} else {
				n.log.Debug("update sign message", "id", id, "signs", len(detail.Signatures))
				n.dataCheck.UpdateMessagePublished(detail)
				responseMsg = &SignedObservation{
					ChainId:    detail.Log.ChainId,
					ID:         id,
					Signatures: detail.Signatures,
				}
			}
		} else {
			//detail = n.dataCheck.HasLogOnChain(msg.SignMessageData.TxHash, msg.SignMessageData.Log.Sequence)
			//if detail != nil {
			//	sig := n.wallet.Sign(msg.SignMessageData.Hash())
			//	p2p.Send(p.rw, SignedObservationType, &SignedObservation{
			//		ChainId:    msg.SignMessageData.Log.ChainId,
			//		ID:         id,
			//		Signatures: &types.Signature{
			//			Index: n.
			//		},
			//	})
			//}
		}
		return nil
	}()
	if err != nil {
		return err
	}
	n.SendGroup(responseMsg, group)

	return nil
}

func (n *Network) handleSignedObservationRequestType(p *peer, msg *SignedObservationRequest) error {
	log := n.log.New("peer", p.id)
	log.Debug("handle sign observation request", "msgid", msg.ID, "chainid", msg.ChainId)
	detail, err := n.dataCheck.GetMessagePublished(msg.ID)
	if err != nil {
		log.Warn("get message log failed", "error", err)
		return err
	}
	if detail != nil {
		log.Debug("local database had log", "msgid", msg.ID, "chainid", msg.ChainId)
		go p2p.Send(p.rw, SignedObservationType, &SignedObservation{
			ChainId:    detail.Log.ChainId,
			ID:         msg.ID,
			Signatures: detail.Signatures,
		})
	}

	return nil
}
func (n *Network) handleSignedObservationType(p *peer, msg *SignedObservation) error {
	log := n.log.New("peer", p.id)
	log.Debug("handle sign observation", "peer", p.id)
	var responseMsg interface{}
	var group *types.PeerGroup
	err := func() error {
		n.dbMutex.Lock()
		defer n.dbMutex.Unlock()
		detail, err := n.dataCheck.GetMessagePublished(msg.ID)
		if err != nil {
			return err
		}

		if detail != nil {
			group = n.PeerGroup(detail.Log.ChainId)
			if group == nil {
				return errors.New("group is invalid")
			}
			if err := VerifySignatures(msg.ID.Bytes(), msg.Signatures, group); err != nil {
				return err
			}
			detail.AddSignature(msg.Signatures)
			if len(detail.Signatures) >= group.Threshold {
				n.dataCheck.SetQuorumLog([]*types.MessagePublishedDetail{detail})
				//broadcast quorum message
				responseMsg = &SignedWithQuorum{
					ChainId:    msg.ChainId,
					ID:         msg.ID,
					Signatures: detail.Signatures,
				}

			} else {
				n.dataCheck.UpdateMessagePublished(detail)
			}
		}
		return nil
	}()
	if err != nil {
		return err
	}
	n.SendGroup(responseMsg, group)
	return nil
}
func (n *Network) PeerGroup(chainId uint64) *types.PeerGroup {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if n.sets == nil {
		return nil
	}
	group := n.sets[chainId]
	return group
}
func (n *Network) handleSignedWithQuorumType(p *peer, msg *SignedWithQuorum) error {
	log := n.log.New("peer", p.id)
	log.Debug("handle sign quorum", "peer", p.id)
	detail, err := n.dataCheck.GetMessagePublished(msg.ID)
	if err != nil {
		return err
	}
	if detail != nil {
		group := n.PeerGroup(detail.Log.ChainId)
		if group == nil {
			return errors.New("group is invalid")
		}
		if len(msg.Signatures) < group.Threshold {
			return errors.New("invalid quorum message, msg sign is non-enough")
		}
		if err := VerifySignatures(msg.ID.Bytes(), msg.Signatures, group); err != nil {
			return err
		}
		detail.AddSignature(msg.Signatures)
		n.dataCheck.SetQuorumLog([]*types.MessagePublishedDetail{detail})
	}
	return nil
}
func (n *Network) SendGroup(msg interface{}, group *types.PeerGroup) {

	if msg != nil && group != nil {
		if size, _, _ := rlp.EncodeToReader(msg); size == 1 && MessageType(msg) == 4 {
			panic(msg)
		}
		go func() {
			for peer, _ := range group.Group {
				n.Send(peer, msg)
			}
		}()
	}
}
func (n *Network) Send(id string, msg interface{}) error {
	if msg != nil {
		msgType := MessageType(msg)
		if msgType < math.MaxUint64 {
			peer := n.peerSet.Peer(id)
			if peer != nil {
				n.log.Debug("send message", "peer", id)

				p2p.Send(peer.rw, msgType, msg)
			}
		}
	}
	return nil
}

func VerifySignatures(id []byte, signatures []*types.Signature, group *types.PeerGroup) error {
	for _, sig := range signatures {
		member := group.GetMember(sig.Index)
		if member == nil {
			return errors.New("member is invalid")
		}
		if !datacommon.VerifySignature(id, sig, member.BlsPubKey) {
			return errors.New("verify signature failed")
		}
	}
	return nil
}
