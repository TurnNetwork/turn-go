package cbft

import (
	"fmt"
	"testing"
	"time"

	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/fetcher"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/network"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	types2 "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/utils"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/stretchr/testify/assert"
)

func init() {
	//log.Root().SetHandler(log.StdoutHandler)
	fetcher.SetArriveTimeout(10 * time.Second)
}

func TestFetch(t *testing.T) {
	pk, sk, cbftnodes := GenerateCbftNode(4)
	nodes := make([]*TestCBFT, 0)
	for i := 0; i < 4; i++ {
		node := MockNode(pk[i], sk[i], cbftnodes, 200000, 10)
		assert.Nil(t, node.Start())

		nodes = append(nodes, node)
		//fmt.Println(i, node.engine.config.Option.NodeID.TerminalString())
	}

	result := make(chan *types.Block, 1)

	var fetchBlock *types.Block
	qcBlocks := &protocols.QCBlockList{}
	parent := nodes[0].chain.Genesis()
	for i := 0; i < 3; i++ {
		block := NewBlock(parent.Hash(), parent.NumberU64()+1)
		assert.True(t, nodes[0].engine.state.HighestExecutedBlock().Hash() == block.ParentHash())
		nodes[0].engine.OnSeal(block, result, nil)
		fetchBlock = block
		qcBlocks.Blocks = append(qcBlocks.Blocks, block)
		_, qc := nodes[0].engine.blockTree.FindBlockAndQC(parent.Hash(), parent.NumberU64())
		select {
		case b := <-result:
			assert.NotNil(t, b)
			assert.Equal(t, uint32(i-1), nodes[0].engine.state.MaxQCIndex())
			for j := 1; j < 4; j++ {
				msg := &protocols.PrepareVote{
					Epoch:          nodes[0].engine.state.Epoch(),
					ViewNumber:     nodes[0].engine.state.ViewNumber(),
					BlockIndex:     uint32(i),
					BlockHash:      b.Hash(),
					BlockNumber:    b.NumberU64(),
					ValidatorIndex: uint32(j),
					ParentQC:       qc,
				}
				pb := nodes[0].engine.state.PrepareBlockByIndex(uint32(i))
				assert.NotNil(t, pb)
				execute := make(chan uint32, 1)
				timer := time.NewTimer(500 * time.Millisecond)
				nodes[j].engine.executeFinishHook = func(index uint32) {
					execute <- index
				}
				assert.Nil(t, nodes[j].engine.OnPrepareBlock("id", pb))

				select {
				case <-timer.C:
					t.Fatal("execute block timeout")
				case <-execute:

				}
				time.Sleep(50 * time.Millisecond)
				index, finish := nodes[j].engine.state.Executing()
				assert.True(t, index == uint32(i) && finish, fmt.Sprintf("i:%d,index:%d,finish:%v", i, index, finish))
				assert.Nil(t, nodes[j].engine.signMsgByBls(msg))
				assert.Nil(t, nodes[0].engine.OnPrepareVote("id", msg), fmt.Sprintf("number:%d", b.NumberU64()))
			}
			_, qc := nodes[0].engine.blockTree.FindBlockAndQC(block.Hash(), block.NumberU64())
			assert.NotNil(t, qc)
			qcBlocks.QC = append(qcBlocks.QC, qc)
			parent = b
		}
	}
	assert.Equal(t, uint64(3), nodes[0].engine.state.HighestQCBlock().NumberU64())
	assert.Equal(t, uint64(0), nodes[1].engine.state.HighestQCBlock().NumberU64())

	total := 3
	finish := make(chan struct{}, 1)
	nodes[1].engine.insertBlockQCHook = func(block *types.Block, qc *types2.QuorumCert) {
		total--
		if total == 0 {
			finish <- struct{}{}
		}
	}
	nodes[1].engine.fetchBlock("id", fetchBlock.Hash(), fetchBlock.NumberU64())
	nodes[1].engine.syncMsgCh <- &types2.MsgInfo{PeerID: "id", Msg: qcBlocks}
	select {
	case <-time.NewTimer(10000 * time.Millisecond).C:
		t.Fatal("fetch timeout")
	case <-finish:

	}
	assert.Equal(t, uint64(3), nodes[1].engine.state.HighestQCBlock().NumberU64())
}

func TestSyncBlock(t *testing.T) {
	nodes := Mock4NodePipe(false)
	result := make(chan *types.Block, 1)
	var fetchBlock *types.Block
	qcBlocks := &protocols.QCBlockList{}
	parent := nodes[0].chain.Genesis()
	for i := 0; i < 3; i++ {
		block := NewBlock(parent.Hash(), parent.NumberU64()+1)
		assert.True(t, nodes[0].engine.state.HighestExecutedBlock().Hash() == block.ParentHash())
		nodes[0].engine.OnSeal(block, result, nil)
		fetchBlock = block
		qcBlocks.Blocks = append(qcBlocks.Blocks, block)
		_, qc := nodes[0].engine.blockTree.FindBlockAndQC(parent.Hash(), parent.NumberU64())
		select {
		case b := <-result:
			assert.NotNil(t, b)
			assert.Equal(t, uint32(i-1), nodes[0].engine.state.MaxQCIndex())
			for j := 1; j < 4; j++ {
				msg := &protocols.PrepareVote{
					Epoch:          nodes[0].engine.state.Epoch(),
					ViewNumber:     nodes[0].engine.state.ViewNumber(),
					BlockIndex:     uint32(i),
					BlockHash:      b.Hash(),
					BlockNumber:    b.NumberU64(),
					ValidatorIndex: uint32(j),
					ParentQC:       qc,
				}
				pb := nodes[0].engine.state.PrepareBlockByIndex(uint32(i))
				assert.NotNil(t, pb)
				execute := make(chan uint32, 1)
				timer := time.NewTimer(500 * time.Millisecond)
				nodes[j].engine.executeFinishHook = func(index uint32) {
					execute <- index
				}
				assert.Nil(t, nodes[j].engine.OnPrepareBlock(nodes[0].engine.config.Option.NodeID.TerminalString(), pb))

				select {
				case <-timer.C:
					t.Fatal("execute block timeout")
				case <-execute:

				}
				index, finish := nodes[j].engine.state.Executing()
				assert.True(t, index == uint32(i) && finish, fmt.Sprintf("%d,%v", index, finish))
				assert.Nil(t, nodes[j].engine.signMsgByBls(msg))
				assert.Nil(t, nodes[0].engine.OnPrepareVote(nodes[j].engine.config.Option.NodeID.TerminalString(), msg), fmt.Sprintf("number:%d", b.NumberU64()))
			}
			_, qc := nodes[0].engine.blockTree.FindBlockAndQC(block.Hash(), block.NumberU64())
			assert.NotNil(t, qc)
			qcBlocks.QC = append(qcBlocks.QC, qc)
			parent = b
		}
	}
	assert.Equal(t, uint64(3), nodes[0].engine.state.HighestQCBlock().NumberU64())
	assert.Equal(t, uint64(0), nodes[1].engine.state.HighestQCBlock().NumberU64())
	assert.Equal(t, uint64(0), nodes[2].engine.state.HighestQCBlock().NumberU64())

	assert.Equal(t, 0, nodes[1].engine.fetcher.Len())
	assert.Equal(t, uint64(0), nodes[1].engine.state.HighestQCBlock().NumberU64())

	for i := 0; i < 4; i++ {
		nodes[i].engine.network.Testing()
	}

	//nodes[1].engine.insertBlockQCHook = func(block *types.Block, qc *types2.QuorumCert) {
	//	fmt.Println("block:", block.Hash().String(), "qc:", qc.BlockNumber)
	//}
	total := 4
	finish := make(chan struct{}, 1)
	nodes[1].engine.insertBlockQCHook = func(block *types.Block, qc *types2.QuorumCert) {
		total--
		if total == 0 {
			finish <- struct{}{}
		}
	}
	nodes[1].engine.fetchBlock(nodes[0].engine.config.Option.NodeID.TerminalString(), fetchBlock.Hash(), fetchBlock.NumberU64())

	select {
	case <-time.NewTimer(10000 * time.Millisecond).C:
		t.Fatal("fetch timeout")
	case <-finish:

	}
	//nodes[1].engine.syncMsgCh <- &types2.MsgInfo{PeerID: "id", Msg: qcBlocks}
	time.Sleep(1000 * time.Millisecond)
	assert.Equal(t, uint64(3), nodes[1].engine.state.HighestQCBlock().NumberU64())

}

func TestFetchBlockRules(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	done := make(chan struct{})
	total := 3
	hook := func(msg *types2.MsgPackage) {
		total--
		if total == 0 {
			done <- struct{}{}
		}
	}
	pk, sk, cbftnodes := GenerateCbftNode(1)
	node := MockNode(pk[0], sk[0], cbftnodes, 1000000, 10)
	node.Start()
	network.SetSendQueueHook(node.engine.network, hook)

	parent := common.Hash{}
	for i := 1; i < 3; i++ {
		block := NewBlock(common.Hash{}, uint64(i))
		parent = block.Hash()
		pb := &protocols.PrepareBlock{
			Epoch:         0,
			ViewNumber:    1,
			Block:         block,
			BlockIndex:    uint32(i - 1),
			ProposalIndex: 0,
		}
		node.engine.state.AddPrepareBlock(pb)
	}

	block := NewBlock(parent, uint64(5))

	pb := &protocols.PrepareBlock{
		Epoch:         0,
		ViewNumber:    1,
		Block:         block,
		BlockIndex:    uint32(5),
		ProposalIndex: 0,
	}

	node.engine.prepareBlockFetchRules("id", pb)
	select {
	case <-timer.C:

		t.Error("timeout")
	case <-done:
	}
}

func TestFetchVoteRule(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	done := make(chan struct{})
	total := 3
	hook := func(msg *types2.MsgPackage) {
		total--
		if total == 0 {
			done <- struct{}{}
		}
	}

	pk, sk, cbftnodes := GenerateCbftNode(1)
	node := MockNode(pk[0], sk[0], cbftnodes, 1000000, 10)
	node.Start()
	network.SetSendQueueHook(node.engine.network, hook)

	parent := common.Hash{}
	for i := 1; i < 3; i++ {
		block := NewBlock(common.Hash{}, uint64(i))
		parent = block.Hash()
		pb := &protocols.PrepareBlock{
			Epoch:         0,
			ViewNumber:    1,
			Block:         block,
			BlockIndex:    uint32(i - 1),
			ProposalIndex: 0,
		}
		node.engine.state.AddPrepareBlock(pb)
	}

	block := NewBlock(parent, uint64(5))

	pb := &protocols.PrepareVote{
		Epoch:       0,
		ViewNumber:  1,
		BlockHash:   block.Hash(),
		BlockNumber: block.NumberU64(),
		BlockIndex:  uint32(5),
	}

	node.engine.prepareVoteFetchRules("id", pb)

	select {
	case <-timer.C:

		t.Error("timeout")
	case <-done:
	}

}

func TestCbft_OnGetPrepareVote(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	done := make(chan struct{})
	hook := func(msg *types2.MsgPackage) {
		done <- struct{}{}

	}
	pk, sk, cbftnodes := GenerateCbftNode(1)
	node := MockNode(pk[0], sk[0], cbftnodes, 1000000, 10)
	node.Start()
	network.SetSendQueueHook(node.engine.network, hook)

	parent := common.Hash{}
	for i := 1; i < 3; i++ {
		block := NewBlock(common.Hash{}, uint64(i))
		parent = block.Hash()

		pb := &protocols.PrepareVote{
			Epoch:       0,
			ViewNumber:  1,
			BlockHash:   block.Hash(),
			BlockNumber: block.NumberU64(),
			BlockIndex:  uint32(i - 1),
		}
		for i := 0; i < 4; i++ {
			node.engine.state.AddPrepareVote(uint32(i), pb)
		}
	}

	assert.Nil(t, node.engine.OnGetPrepareVote("id", &protocols.GetPrepareVote{
		ViewNumber:  1,
		BlockHash:   parent,
		BlockNumber: 4,
		BlockIndex:  1,
		VoteBits:    utils.NewBitArray(4),
	}))
	select {
	case <-timer.C:

		t.Error("timeout")
	case <-done:
	}
}

func TestCbft_OnGetLatestStatus(t *testing.T) {
	engine, cNodes := buildSingleCbft()
	// use case.
	testCases := []struct {
		blockBn uint64
		reqBn   uint64
		reqType uint64
	}{
		{1, 1, network.TypeForQCBn},
		{1, 2, network.TypeForQCBn},
		{2, 1, network.TypeForQCBn},
		{1, 1, network.TypeForLockedBn},
		{1, 2, network.TypeForLockedBn},
		{2, 1, network.TypeForLockedBn},
		{1, 1, network.TypeForCommitBn},
		{1, 2, network.TypeForCommitBn},
		{2, 1, network.TypeForCommitBn},
	}
	peer, _ := engine.network.GetPeer(cNodes[0].TerminalString())
	for _, v := range testCases {
		message := &protocols.GetLatestStatus{
			BlockNumber: v.reqBn,
			LogicType:   v.reqType,
		}
		engine.state.SetHighestQCBlock(NewBlock(common.Hash{}, uint64(v.blockBn)))
		engine.state.SetHighestLockBlock(NewBlock(common.Hash{}, uint64(v.blockBn)))
		engine.state.SetHighestCommitBlock(NewBlock(common.Hash{}, uint64(v.blockBn)))
		err := engine.OnGetLatestStatus(peer.PeerID(), message)
		assert.Nil(t, err)
		if v.blockBn < v.reqBn {
			switch v.reqType {
			case network.TypeForQCBn:
				assert.Equal(t, v.blockBn, engine.state.HighestQCBlock().NumberU64())
				assert.Equal(t, v.reqBn, peer.QCBn())
			case network.TypeForLockedBn:
				assert.Equal(t, v.blockBn, engine.state.HighestLockBlock().NumberU64())
				assert.Equal(t, v.reqBn, peer.LockedBn())
			case network.TypeForCommitBn:
				assert.Equal(t, v.blockBn, engine.state.HighestCommitBlock().NumberU64())
				assert.Equal(t, v.reqBn, peer.CommitBn())
			}
		}
	}
}

func TestCbft_OnLatestStatus(t *testing.T) {
	engine, cNodes := buildSingleCbft()
	// use case.
	testCases := []struct {
		blockBn uint64
		rspBn   uint64
		rspType uint64
	}{
		{1, 1, network.TypeForQCBn},
		{1, 2, network.TypeForQCBn},
		{2, 1, network.TypeForQCBn},
		{1, 1, network.TypeForLockedBn},
		{1, 2, network.TypeForLockedBn},
		{2, 1, network.TypeForLockedBn},
		{1, 1, network.TypeForCommitBn},
		{1, 2, network.TypeForCommitBn},
		{2, 1, network.TypeForCommitBn},
	}
	peer, _ := engine.network.GetPeer(cNodes[0].TerminalString())
	for _, v := range testCases {
		message := &protocols.LatestStatus{
			BlockNumber: v.rspBn,
			LogicType:   v.rspType,
		}
		engine.state.SetHighestQCBlock(NewBlock(common.Hash{}, uint64(v.blockBn)))
		engine.state.SetHighestLockBlock(NewBlock(common.Hash{}, uint64(v.blockBn)))
		engine.state.SetHighestCommitBlock(NewBlock(common.Hash{}, uint64(v.blockBn)))
		err := engine.OnLatestStatus(peer.PeerID(), message)
		assert.Nil(t, err)
		if v.blockBn < v.rspBn {
			switch v.rspType {
			case network.TypeForQCBn:
				assert.Equal(t, v.blockBn, engine.state.HighestQCBlock().NumberU64())
				assert.Equal(t, v.rspBn, peer.QCBn())
			case network.TypeForLockedBn:
				assert.Equal(t, v.blockBn, engine.state.HighestLockBlock().NumberU64())
				assert.Equal(t, v.rspBn, peer.LockedBn())
			case network.TypeForCommitBn:
				assert.Equal(t, v.blockBn, engine.state.HighestCommitBlock().NumberU64())
				assert.Equal(t, v.rspBn, peer.CommitBn())
			}
		}
	}
}

func TestCbft_MissingViewChangeNodes(t *testing.T) {
	engine, cNodes := buildSingleCbft()
	nodes, message, err := engine.MissingViewChangeNodes()
	assert.Equal(t, len(cNodes), len(nodes))
	assert.Nil(t, err)
	assert.NotNil(t, message)
}

func buildSingleCbft() (*Cbft, []discover.NodeID) {
	// Init mock node.
	pk, sk, cbftnodes := GenerateCbftNode(1)
	node := MockNode(pk[0], sk[0], cbftnodes, 1000000, 10)
	node.Start()
	//node.engine.network.Close()
	// Add a node to the Handler.
	cNodes := network.RandomID()
	node.engine.consensusNodesMock = func() ([]discover.NodeID, error) {
		return cNodes, nil
	}
	network.FillEngineManager(cNodes, node.engine.network)
	return node.engine, cNodes
}

func TestCbft_OnPong(t *testing.T) {
	running := func(value int64) int64 {
		engine, cNodes := buildSingleCbft()
		for _, v := range cNodes {
			curTime := time.Now().UnixNano()
			tInt64 := curTime - value*1000000 // Suppose there is a 200 millisecond delay.
			latency := (curTime - tInt64) / 2 / 1000000
			engine.OnPong(v.TerminalString(), latency)
		}
		avg := engine.AvgLatency()
		return avg
	}

	// Simulation calls the OnPong method.
	testCases := []struct {
		value  int64
		expect int64
	}{
		{value: 200, expect: 100},
		{value: 300, expect: 150},
		{value: 400, expect: 200},
	}
	for i := 0; i < len(testCases); i++ {
		value := running(testCases[i].value)
		assert.Equal(t, testCases[i].expect, value)
	}
}
