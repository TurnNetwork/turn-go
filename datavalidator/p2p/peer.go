package p2p

import (
	"fmt"

	"github.com/bubblenet/bubble/p2p"
	mapset "github.com/deckarep/golang-set"
)

type peer struct {
	id string

	*p2p.Peer
	rw         p2p.MsgReadWriter
	scanBlock  uint64
	knownQurom mapset.Set
}

func newPeer(p *p2p.Peer, rw p2p.MsgReadWriter) *peer {
	return &peer{
		id:   fmt.Sprintf("%x", p.ID().Bytes()[:8]),
		Peer: p,
		rw:   rw,
	}
}

type NodeInfo struct {
	ScanBlock uint64
}

func (p peer) Info() *NodeInfo {
	return &NodeInfo{
		ScanBlock: p.scanBlock,
	}
}

func (p *peer) close() {}
