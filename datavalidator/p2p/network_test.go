package p2p

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/bubblenet/bubble/core/rawdb"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/datavalidator/common"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/log/term"
	"github.com/mattn/go-colorable"
	"io"
	"net/http"
	"os"

	"github.com/bubblenet/bubble/datavalidator/db"
	"github.com/bubblenet/bubble/datavalidator/mock"
	"github.com/bubblenet/bubble/datavalidator/sync"
	wallet2 "github.com/bubblenet/bubble/datavalidator/wallet"
	"github.com/bubblenet/bubble/ethdb"
	"github.com/bubblenet/bubble/ethdb/memorydb"
	"math/big"
	"testing"
	"time"

	"github.com/bubblenet/bubble/datavalidator/types"
	p2p2 "github.com/bubblenet/bubble/p2p"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/stretchr/testify/require"
	_ "net/http/pprof"
)

func initPprof() {
	NewProfileHttpServer := func(addr string) {
		go func() {
			err := http.ListenAndServe(addr, nil)
			if err != nil {
				panic(err)
			}
		}()
	}
	NewProfileHttpServer("192.168.3.9:8080")
}

type DataValidator struct {
	Ctx     context.Context
	Cancel  context.CancelFunc
	Db      *db.DB
	Sync    *sync.Sync
	Network *Network
}

var skstr = []string{
	"84ce2fb332ae78d78bee8cbc0bf244d0dd6e8da97bafeb191d9765ba7fe9d39b",
	"b41ac21b6b185966cabf3673f0fce9a957afa0117987af247124cbeaeb94f194",
	"78f99961967c4a39e63e7b1d18f6b8be46608c66bcd28e1331223fd3821a0ca4",
}

var blsKeyStrs = []string{
	"f958b2708c0a6eae0ea5761edcf0257526a8bbe521cc099e32adbff14b049734",
	"86e71e0d2feb7cb2233aaf084beaca51a1dfa0107d4e8ae4417555786820de1a",
	"079d678e6e61d949c60d43237c5b62d9fb2b4e71747ddd804673358e5bbd253f",
	"6ad2849adea42f1a9f981a7caa3733024c08bcba3849cc7f7ec2fe808debaf5b",
}

func init() {
	initPprof()
	usecolor := term.IsTty(os.Stderr.Fd()) && os.Getenv("TERM") != "dumb"
	output := io.Writer(os.Stderr)
	if usecolor {
		output = colorable.NewColorableStderr()
	}
	ostream := log.StreamHandler(output, log.TerminalFormat(usecolor))
	glogger := log.NewGlogHandler(ostream)
	log.PrintOrigins(true)
	glogger.Verbosity(log.Lvl(log.LvlDebug))
	log.Root().SetHandler(glogger)
}
func TestBls(t *testing.T) {
	//for i := 0; i < 4; i++ {
	//	sk := bls.GenerateKey()
	//	fmt.Println(fmt.Sprintf("\"%s\",", hex.EncodeToString(sk.GetLittleEndian())))
	//}
	buf, _ := hex.DecodeString("86e71e0d2feb7cb2233aaf084beaca51a1dfa0107d4e8ae4417555786820de1a")
	var sec bls.SecretKey
	sec.SetLittleEndian(buf)
	fmt.Println(hex.EncodeToString(sec.GetLittleEndian()))
	fmt.Println(hex.EncodeToString(sec.GetPublicKey().Serialize()), common.BlsID(sec.GetPublicKey()))

}
func TestNetworkFlow(t *testing.T) {

	number := 3
	var sks []*ecdsa.PrivateKey
	var blsKeys []*bls.SecretKey
	for i := 0; i < number; i++ {
		sk, _ := crypto.HexToECDSA(skstr[i])
		sks = append(sks, sk)
		var key bls.SecretKey
		buf, _ := hex.DecodeString(blsKeyStrs[i])
		key.SetLittleEndian(buf)
		blsKeys = append(blsKeys, &key)
	}
	innercontract := mock.NewInnerContract(sks, blsKeys, map[uint64][]uint64{
		1: []uint64{0, 1, 2},
	})
	blockFilter := mock.NewBlockFilter([]uint64{1})

	mdb := memorydb.New()
	querydb := db.NewDataValidatorDB(rawdb.NewDatabase(mdb))
	querydb.StoreScanLog(0)
	vs := newMockDataValidator(blsKeys[0], innercontract, blockFilter, rawdb.NewDatabase(mdb), blockFilter)

	blockFilter.AddMessagePublished(1, 1)
	logs, err := blockFilter.RangeFilter(context.Background(), big.NewInt(0), big.NewInt(int64(blockFilter.BlockNumber())))
	require.Nil(t, err)
	owner := discover.PubkeyID(&sks[0].PublicKey)
	other := discover.PubkeyID(&sks[1].PublicKey)
	peer1, rw1, peer2, rw2 := p2p2.NewPeerByNodeID(owner, other, []p2p2.Protocol{vs.Network.Protocols()})

	fmt.Println("peer1", peer1.ID().TerminalString(), "peer2", peer2.ID().TerminalString())
	go vs.Network.Protocols().Run(peer2, rw2)
	queryDetail, err := querydb.GetUnSignChainIdNonce(1, 0)
	require.Nil(t, err)
	require.Nil(t, queryDetail)
	vs.Sync.HandleMessage(context.Background())
	queryDetail, err = querydb.GetUnSignChainIdNonce(1, 0)
	require.Nil(t, err)
	require.Equal(t, 0, len(queryDetail.Signatures))
	wallet := wallet2.FromSk(blsKeys[1])
	detail := logs[0]
	detail.Signatures = append(detail.Signatures, &types.Signature{
		Index:     1,
		Signature: wallet.Sign(logs[0].Hash()),
	})
	sendMsg(rw1, &SignMessageMsg{
		SignMessageData: *logs[0],
		Signature: &types.Signature{
			Index:     1,
			Signature: wallet.Sign(logs[0].Hash()),
		},
	})

	readMsg(t, rw1)
	queryDetail, err = querydb.GetUnSignChainIdNonce(detail.Log.ChainId, detail.Log.Nonce)
	require.Nil(t, err)
	require.Equal(t, 1, len(queryDetail.Signatures))
	wallet0 := wallet2.FromSk(blsKeys[0])
	wallet22 := wallet2.FromSk(blsKeys[2])
	sendMsg(rw1, &SignedObservation{
		ChainId: detail.Log.ChainId,
		ID:      detail.Hash(),
		Signatures: []*types.Signature{
			&types.Signature{
				Index:     0,
				Signature: wallet0.Sign(detail.Hash()),
			}, &types.Signature{
				Index:     1,
				Signature: wallet.Sign(detail.Hash()),
			}, &types.Signature{
				Index:     2,
				Signature: wallet22.Sign(detail.Hash()),
			},
		},
	})
	readMsg(t, rw1)
	//require.NotNil(t, msg)
	queryDetail, err = querydb.GetUnSignChainIdNonce(detail.Log.ChainId, detail.Log.Nonce)
	require.Nil(t, err)
	require.Nil(t, queryDetail)
	queryDetail, err = querydb.GetQuorumChainIdNonce(detail.Log.ChainId, detail.Log.Nonce)
	require.Nil(t, err)
	require.Equal(t, 3, len(queryDetail.Signatures))

	sendMsg(rw1, &SignedObservationRequest{
		ChainId: detail.Log.ChainId,
		ID:      detail.Hash(),
	})

	time.Sleep(500 * time.Millisecond)
	res := readMsg(t, rw1).(*SignedObservation)
	require.Equal(t, 3, len(res.Signatures))

	go vs.Network.HeartbeatLoop(context.Background(), time.Second)
	//time.Sleep(30 * time.Second)
	heart := readMsg(t, rw1).(*Heartbeat)
	require.NotNil(t, heart)
}

func readMsg(t *testing.T, rw p2p2.MsgReadWriter) interface{} {
	msg, err := rw.ReadMsg()
	require.Nil(t, err)
	switch msg.Code {
	case SignMessageType:
	case HeartbeatMessageType:
		var m Heartbeat
		require.Nil(t, msg.Decode(&m))
		return &m
	case SignedObservationRequestType:
	case SignedObservationType:
		var m SignedObservation
		require.Nil(t, msg.Decode(&m))
		return &m
	case SignedWithQuorumType:
		var m SignedWithQuorum
		require.Nil(t, msg.Decode(&m))
		return &m
	}
	return nil
}
func sendMsg(rw p2p2.MsgReadWriter, msg interface{}) {
	msgType := MessageType(msg)
	go p2p2.Send(rw, msgType, msg)
}
func newMockDataValidator(sk *bls.SecretKey, contract *mock.InnerContract, blockFilter *mock.BlockFilter, edb ethdb.Database, blockState types.BlockState) *DataValidator {
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

	network := NewNetwork(wallet, validatorContract, &mock.DataCheck{
		DB:            vdb,
		FilterMessage: messagePublished,
	}, blockState, nil)
	sets := sync.RefreshValidator()
	network.SetValidatorSet(sets)

	ctx, cancel := context.WithCancel(context.Background())
	return &DataValidator{
		Ctx:     ctx,
		Cancel:  cancel,
		Db:      vdb,
		Sync:    sync,
		Network: network,
	}
}
