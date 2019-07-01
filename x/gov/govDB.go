package gov

import (
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/x/xcom"
	"sync"
)

var (
	ValueDelimiter = []byte(":")
)

type VoteValue struct {
	Voter  discover.NodeID
	Option VoteOption
}

var dbOnce sync.Once
var govDB *GovDB

type GovDB struct {
	govdbErr error
	snapdb   GovSnapshotDB
}

func GovDBInstance() *GovDB {
	dbOnce.Do(func() {
		govDB = &GovDB{snapdb: GovSnapshotDB{}}
	})
	return govDB
}

func tobytes(data interface{}) []byte {
	if bytes, err := json.Marshal(data); err != nil {
		return bytes
	} else {
		log.Error("govdb, marshal value to bytes error..")
		panic(err)
	}
}

// 保存提案记录，value编码规则:
//  value 为[]byte，其中byte[byte.len -2] 为type,byte[0:byte.len-2]为proposal
func (self *GovDB) SetProposal(proposal Proposal, state xcom.StateDB) error {

	bytes, e := json.Marshal(proposal)
	if e != nil {
		return e
	}

	value := append(bytes, byte(proposal.GetProposalType()))
	state.SetState(vm.GovContractAddr, KeyProposal(proposal.GetProposalID()), value)

	return nil
}

func (self *GovDB) setError(err error) {
	if err != nil {
		self.govdbErr = err
		panic(err)
	}
}

// 查询提案记录，获取value后，解码
func (self *GovDB) GetProposal(proposalID common.Hash, state xcom.StateDB) (Proposal, error) {
	value := state.GetState(vm.GovContractAddr, KeyProposal(proposalID))
	if len(value) == 0 {
		return nil, fmt.Errorf("no value found!")
	}
	var p Proposal
	pData := value[0 : len(value)-1]
	pType := value[len(value)-1]
	if pType == byte(Text) {
		var proposal TextProposal
		if e := json.Unmarshal(pData, &proposal); e != nil {
			return nil, e
		}
		p = proposal
	} else if pType == byte(Version) {
		var proposal VersionProposal
		//proposal = VersionProposal{TextProposal{},0,common.Big0}
		if e := json.Unmarshal(pData, &proposal); e != nil {
			return nil, e
		}
		p = proposal
	} else {
		return nil, fmt.Errorf("incorrect propsal type:%b!", pType)
	}

	return p, nil
}

// 从snapdb查询各个列表id,然后从逐条从statedb查询
func (self *GovDB) GetProposalList(blockHash common.Hash, state xcom.StateDB) ([]Proposal, error) {
	proposalIds, err := self.snapdb.getAllProposalIDList(blockHash)
	if err != nil {
		return nil, err
	}
	var proposls []Proposal
	for _, hash := range proposalIds {
		proposal, _ := self.GetProposal(hash, state)
		if proposal != nil {
			proposls = append(proposls, proposal)
		}
	}
	return proposls, nil
}

//保存投票记录
func (self *GovDB) SetVote(proposalID common.Hash, voter discover.NodeID, option VoteOption, state xcom.StateDB) bool {
	voteList := self.ListVote(proposalID, state)
	voteList = append(voteList, VoteValue{voter, option})

	voteListBytes, _ := json.Marshal(voteList)

	state.SetState(vm.GovContractAddr, KeyVote(proposalID), voteListBytes)
	return true
}

// 查询投票记录
func (self *GovDB) ListVote(proposalID common.Hash, state xcom.StateDB) []VoteValue {
	voteListBytes := state.GetState(vm.GovContractAddr, KeyVote(proposalID))

	var voteList []VoteValue
	if err := json.Unmarshal(voteListBytes, &voteList); err != nil {
		return nil
	}
	return voteList
}

func (self *GovDB) ListVotedVerifier(proposalID common.Hash, state xcom.StateDB) ([]discover.NodeID, error) {
	return nil, nil
}

// 保存投票结果
func (self *GovDB) SetTallyResult(tallyResult TallyResult, state xcom.StateDB) bool {
	value, _ := json.Marshal(tallyResult)
	state.SetState(vm.GovContractAddr, KeyTallyResult(tallyResult.ProposalID), value)
	return true
}

// 查询投票结果
func (self *GovDB) GetTallyResult(proposalID common.Hash, state xcom.StateDB) (*TallyResult, error) {
	value := state.GetState(vm.GovContractAddr, KeyTallyResult(proposalID))

	var tallyResult TallyResult
	if err := json.Unmarshal(value, &tallyResult); err != nil {
		return nil, err
	}

	return &tallyResult, nil
}

// 保存生效版本记录
func (self *GovDB) SetPreActiveVersion(preActiveVersion uint32, state xcom.StateDB) bool {
	state.SetState(vm.GovContractAddr, KeyPreActiveVersion(), common.Uint32ToBytes(preActiveVersion))
	return true
}

// 查询生效版本记录
func (self *GovDB) GetPreActiveVersion(state xcom.StateDB) uint32 {
	value := state.GetState(vm.GovContractAddr, KeyPreActiveVersion())
	return common.BytesToUint32(value)
}

// 保存生效版本记录
func (self *GovDB) SetActiveVersion(activeVersion uint32, state xcom.StateDB) bool {
	state.SetState(vm.GovContractAddr, KeyActiveVersion(), common.Uint32ToBytes(activeVersion))
	return true
}

// 查询生效版本记录
func (self *GovDB) GetActiveVersion(state xcom.StateDB) uint32 {
	value := state.GetState(vm.GovContractAddr, KeyActiveVersion())
	return common.BytesToUint32(value)
}

// 查询正在投票的提案
func (self *GovDB) ListVotingProposal(blockHash common.Hash, state xcom.StateDB) []common.Hash {
	value, err := govDB.snapdb.getVotingIDList(blockHash)
	if err != nil {
		log.Error("List voting proposal ID error")
		return nil
	}
	return value
}

// 获取投票结束的提案
func (self *GovDB) ListEndProposalID(blockHash common.Hash, state xcom.StateDB) []common.Hash {
	value, err := govDB.snapdb.getEndIDList(blockHash)
	if err != nil {
		log.Error("List end proposal ID error")
		return nil
	}

	return value
}

// 查询预生效的升级提案
func (self *GovDB) GetPreActiveProposalID(blockHash common.Hash, state xcom.StateDB) common.Hash {
	value, err := govDB.snapdb.getPreActiveIDList(blockHash)
	if err != nil {
		log.Error("Get pre-active proposal ID error")
		return common.Hash{}
	}
	return value[0]
}

// 把新增提案的ID增加到正在投票的提案队列中
func (self *GovDB) AddVotingProposalID(blockHash common.Hash, proposalID common.Hash, state xcom.StateDB) bool {
	if err := govDB.snapdb.addProposalByKey(blockHash, KeyVotingProposals(), proposalID); err != nil {
		log.Error("add voting proposal to snapshot db error:%s", err)
		return false
	}

	return true
}

// 把提案的ID从正在投票的提案队列中移动到预激活中
func (self *GovDB) MoveVotingProposalIDToPreActive(blockHash common.Hash, proposalID common.Hash, state xcom.StateDB) bool {

	voting, _ := self.snapdb.getVotingIDList(blockHash)
	voting = remove(voting, proposalID)

	pre, _ := self.snapdb.getPreActiveIDList(blockHash)
	pre = append(pre, proposalID)

	//重新写入
	self.snapdb.addProposalByKey(blockHash, KeyVotingProposals(), proposalID)
	self.snapdb.addProposalByKey(blockHash, KeyPreActiveProposals(), proposalID)

	return true
}

func remove(list []common.Hash, item common.Hash) []common.Hash {
	for i, id := range list {
		if id == item {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list
}

// 把提案的ID从正在投票的提案队列中移动到投票结束的提案队列中
func (self *GovDB) MoveVotingProposalIDToEnd(blockHash common.Hash, proposalID common.Hash, state xcom.StateDB) bool {

	voting, _ := self.snapdb.getVotingIDList(blockHash)
	voting = remove(voting, proposalID)

	end, _ := self.snapdb.getEndIDList(blockHash)
	end = append(end, proposalID)

	//重新写入
	self.snapdb.addProposalByKey(blockHash, KeyVotingProposals(), proposalID)
	self.snapdb.addProposalByKey(blockHash, KeyEndProposals(), proposalID)

	return true
}

// 把提案的ID从预激活的提案队列中移动到投票结束的提案队列中
func (self *GovDB) MovePreActiveProposalIDToEnd(blockHash common.Hash, proposalID common.Hash, state xcom.StateDB) bool {

	pre, _ := self.snapdb.getPreActiveIDList(blockHash)
	pre = remove(pre, proposalID)

	end, _ := self.snapdb.getEndIDList(blockHash)
	end = append(end, proposalID)

	//重新写入
	self.snapdb.addProposalByKey(blockHash, KeyPreActiveProposals(), proposalID)
	self.snapdb.addProposalByKey(blockHash, KeyEndProposals(), proposalID)

	return true
}

// 增加升级提案投票期间版本声明的验证人/候选人记录
func (self *GovDB) AddActiveNode(blockHash common.Hash, proposalID common.Hash, nodeID discover.NodeID) bool {
	if err := self.snapdb.addActiveNode(blockHash, nodeID, proposalID); err != nil {
		log.Error("add declared node to snapshot db error,", err)
		return false
	}
	return true
}

// 获取升级提案投票期间版本升声明的节点列表
func (self *GovDB) GetActiveNodeList(blockHash common.Hash, proposalID common.Hash) []discover.NodeID {
	nodes, err := self.snapdb.getActiveNodeList(blockHash, proposalID)
	if err != nil {
		log.Error("get declared node list from snapshot db error,", err)
		return nil
	}
	return nodes
}

// 升级后，清除做过版本声明的节点
func (self *GovDB) ClearActiveNodes(blockHash common.Hash, proposalID common.Hash) bool {
	err := self.snapdb.deleteActiveNodeList(blockHash, proposalID)
	if err != nil {
		log.Error("delete declared node list from snapshot db error,", err)
		return false
	}
	return true
}

// 增加已投票验证人记录
func (self *GovDB) AddVotedVerifier(blockHash common.Hash, proposalID common.Hash, voter discover.NodeID) bool {
	if err := self.snapdb.addVotedVerifier(blockHash, voter, proposalID); err != nil {
		log.Error("add voted node to snapshot db error,", err)
		return false
	}
	return true
}

// 累计在结算周期内可投票的所有验证人
func (self *GovDB) AccuVerifiers(blockHash common.Hash, proposalID common.Hash, verifierList []discover.NodeID) bool {
	if err := self.snapdb.addTotalVerifiers(blockHash, proposalID, verifierList); err != nil {
		log.Error("add total verifier to  snapshot db error,", err)
		return false
	}
	return true
}

// 获取所有可投票验证人总数
func (self *GovDB) AccuVerifiersLength(blockHash common.Hash, proposalID common.Hash) int {
	if l, err := self.snapdb.getAccuVerifiersLength(blockHash, proposalID); err != nil {
		log.Error("add total verifier to  snapshot db error,", err)
		return 0
	} else {
		return l
	}
}
