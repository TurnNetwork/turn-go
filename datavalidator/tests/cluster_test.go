package tests

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/bubblenet/bubble/core/rawdb"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/datavalidator"
	"github.com/bubblenet/bubble/datavalidator/mock"
	"github.com/bubblenet/bubble/ethdb/memorydb"
	"github.com/bubblenet/bubble/log"
	"github.com/bubblenet/bubble/log/term"
	"github.com/bubblenet/bubble/p2p"
	p2p2 "github.com/bubblenet/bubble/p2p"
	"github.com/bubblenet/bubble/p2p/discover"
	"github.com/mattn/go-colorable"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"testing"
	"time"
)

type P2PPeer struct {
	peer1 *p2p.Peer
	rw1   p2p.MsgReadWriter
	peer2 *p2p.Peer
	rw2   p2p.MsgReadWriter
}
type MockValidator struct {
	idValidator bool
	db          *memorydb.Database
	peers       map[string]*P2PPeer
	dv          *datavalidator.DataValidator
}

var skstr = []string{
	"84ce2fb332ae78d78bee8cbc0bf244d0dd6e8da97bafeb191d9765ba7fe9d39b",
	"b41ac21b6b185966cabf3673f0fce9a957afa0117987af247124cbeaeb94f194",
	"78f99961967c4a39e63e7b1d18f6b8be46608c66bcd28e1331223fd3821a0ca4",
}
var blsKeyStrs = []string{
	"2488270d920dd57f0a2c7f69577b019c575480663356df4e94f90d5a1d6a72ec",
	"6b5ac12b43001b4f53f407878de9e0225bbfeb92541ed8a6e66af514bdac05ed",
	"1bf528b31769a7d1bfaf041ade586e2896a503e39e85bd0d463ef19bc45d3aa7",
	"4c2ea01368827230e8c86e27ad4f8a8f7925f32c49a70a98435e50b9e1411495",
}

func init() {
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
func initPprof() {
	NewProfileHttpServer := func(addr string) {
		go func() {
			err := http.ListenAndServe(addr, nil)
			if err != nil {
				panic(err)
			}
		}()
	}
	NewProfileHttpServer(":8080")
}
func TestCluster(t *testing.T) {

	number := 3
	var sks []*ecdsa.PrivateKey
	var blsKeys []*bls.SecretKey

	for i := 0; i < number; i++ {
		sk, _ := crypto.HexToECDSA(skstr[i])
		fmt.Println("sk:", discover.PubkeyID(&sk.PublicKey).TerminalString())
		sks = append(sks, sk)
		var key bls.SecretKey
		key.SetHexString(blsKeyStrs[i])
		blsKeys = append(blsKeys, &key)
	}
	innercontract := mock.NewInnerContract(sks, blsKeys, map[uint64][]uint64{
		1: []uint64{0, 1, 2},
	})
	blockFilter := mock.NewBlockFilter([]uint64{1})

	validators := make(map[uint64]*MockValidator)

	//var vs []*datavalidator.DataValidator
	for i, _ := range sks {
		mdb := memorydb.New()
		vs := NewDataValidator(blsKeys[i], innercontract, blockFilter, rawdb.NewDatabase(mdb), blockFilter)
		validators[uint64(i)] = &MockValidator{
			idValidator: true,
			db:          mdb,
			peers:       make(map[string]*P2PPeer),
			dv:          vs,
		}
	}
	go func() {
		for i := 0; i < 1; i++ {
			blockFilter.AddMessagePublished(1, 1)
		}
		time.Sleep(5000 * time.Second)

	}()
	for i := 0; i < len(sks)-1; i++ {
		owner := discover.PubkeyID(&sks[i].PublicKey)
		fmt.Println("owner", owner.TerminalString())
		for j := i + 1; j < len(sks); j++ {
			other := discover.PubkeyID(&sks[j].PublicKey)
			//fmt.Println("other", other.TerminalString())
			peer1, rw1, peer2, rw2 := p2p2.NewPeerByNodeID(owner, other, validators[uint64(i)].dv.P2P())
			fmt.Println("peer1", peer1.ID().TerminalString(), "peer2", peer2.ID().TerminalString())
			validators[uint64(i)].peers[other.String()] = &P2PPeer{
				peer1: peer1,
				rw1:   rw1,
				peer2: peer2,
				rw2:   rw2,
			}
			validators[uint64(j)].peers[owner.String()] = &P2PPeer{
				peer1: peer2,
				rw1:   rw2,
				peer2: peer1,
				rw2:   rw1,
			}
			go validators[uint64(i)].dv.Network.Protocols().Run(peer2, rw2)
			go validators[uint64(j)].dv.Network.Protocols().Run(peer1, rw1)
		}
	}
	t.Log("init success")
	go func() {
		for _, d := range validators {
			d.dv.Db.StoreScanLog(0)
			d.dv.Start()
		}
	}()

	for {
		time.Sleep(1 * time.Second)

		logs, err := validators[0].dv.Db.GetQuorumLogRangeNonce(1, 0, 11)
		logs2, err2 := validators[1].dv.Db.GetQuorumLogRangeNonce(1, 0, 11)
		logs3, _ := validators[2].dv.Db.GetQuorumLogRangeNonce(1, 0, 11)

		require.Nil(t, err)
		require.Nil(t, err2)
		t.Log("logs", len(logs))
		if len(logs) == 1 && len(logs2) == 1 && len(logs3) == 1 {
			return
		}
	}
}
