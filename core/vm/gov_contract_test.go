package vm

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"testing"

	//"github.com/PlatONnetwork/PlatON-Go/log"

	"github.com/PlatONnetwork/PlatON-Go/x/xutil"

	"github.com/PlatONnetwork/PlatON-Go/node"

	"github.com/stretchr/testify/assert"

	"github.com/PlatONnetwork/PlatON-Go/common/mock"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/x/gov"

	"github.com/PlatONnetwork/PlatON-Go/common"
	commonvm "github.com/PlatONnetwork/PlatON-Go/common/vm"
	"github.com/PlatONnetwork/PlatON-Go/x/plugin"
	"github.com/PlatONnetwork/PlatON-Go/x/xcom"
)

var (
	govPlugin   *plugin.GovPlugin
	gc          *GovContract
	versionSign common.VersionSign
	chandler    *node.CryptoHandler

	paramModule       = gov.ModuleStaking
	paramName         = gov.KeyMaxValidators
	defaultProposalID = txHashArr[1]
)

func commit_sndb(chain *mock.Chain) {
	/*
		//Flush() signs a Hash to the current block which has no hash yet. Flush() do not write the data to database.
		//in this file, all blocks in each test case has a hash already, so, do not call Flush()
				if err := chain.SnapDB.Flush(chain.CurrentHeader().Hash(), chain.CurrentHeader().Number); err != nil {
					fmt.Println("commit_sndb error:", err)
				}
	*/
	if err := chain.SnapDB.Commit(chain.CurrentHeader().Hash()); err != nil {
		fmt.Println("commit_sndb error:", err)
	}
}

func prepair_sndb(chain *mock.Chain, txHash common.Hash) {
	if txHash == common.ZeroHash {
		chain.AddBlock()
	} else {
		chain.AddBlockWithTxHash(txHash)
	}

	//fmt.Println("prepair_sndb::::::", chain.CurrentHeader().ParentHash.Hex())
	if err := chain.SnapDB.NewBlock(chain.CurrentHeader().Number, chain.CurrentHeader().ParentHash, chain.CurrentHeader().Hash()); err != nil {
		fmt.Println("prepair_sndb error:", err)
	}

	//prepare gc to run contract
	gc.Evm = newEvm(chain.CurrentHeader().Number, chain.CurrentHeader().Hash(), chain.StateDB)
}

func skip_emptyBlock(chain *mock.Chain, blockNumber uint64) {
	cnt := blockNumber - chain.CurrentHeader().Number.Uint64()
	for i := uint64(0); i < cnt; i++ {
		prepair_sndb(chain, common.ZeroHash)
		commit_sndb(chain)
	}
}

func init() {
	chandler = node.GetCryptoHandler()
	chandler.SetPrivateKey(priKeyArr[0])
	versionSign.SetBytes(chandler.MustSign(promoteVersion))
}

func buildSubmitTextInput() []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2000))) // func type code
	input = append(input, common.MustRlpEncode(nodeIdArr[1])) // param 1 ...
	input = append(input, common.MustRlpEncode("textUrl"))

	return common.MustRlpEncode(input)
}
func buildSubmitText(nodeID discover.NodeID, pipID string) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2000))) // func type code
	input = append(input, common.MustRlpEncode(nodeID))       // param 1 ...
	input = append(input, common.MustRlpEncode(pipID))

	return common.MustRlpEncode(input)
}

func buildSubmitParam(nodeID discover.NodeID, pipID string, module, name, newValue string) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2002))) // func type code
	input = append(input, common.MustRlpEncode(nodeID))       // param 1 ...
	input = append(input, common.MustRlpEncode(pipID))
	input = append(input, common.MustRlpEncode(module))
	input = append(input, common.MustRlpEncode(name))
	input = append(input, common.MustRlpEncode(newValue))

	return common.MustRlpEncode(input)
}

func buildSubmitVersionInput() []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2001))) // func type code
	input = append(input, common.MustRlpEncode(nodeIdArr[0])) // param 1 ...
	input = append(input, common.MustRlpEncode("verionPIPID"))
	input = append(input, common.MustRlpEncode(promoteVersion)) //new version : 1.1.1
	input = append(input, common.MustRlpEncode(xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())))

	return common.MustRlpEncode(input)
}

func buildSubmitVersion(nodeID discover.NodeID, pipID string, newVersion uint32, endVotingRounds uint64) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2001))) // func type code
	input = append(input, common.MustRlpEncode(nodeID))       // param 1 ...
	input = append(input, common.MustRlpEncode(pipID))
	input = append(input, common.MustRlpEncode(newVersion)) //new version : 1.1.1
	input = append(input, common.MustRlpEncode(endVotingRounds))

	return common.MustRlpEncode(input)
}

func buildSubmitCancelInput() []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2005))) // func type code
	input = append(input, common.MustRlpEncode(nodeIdArr[0])) // param 1 ..
	input = append(input, common.MustRlpEncode("cancelPIPID"))
	input = append(input, common.MustRlpEncode(xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())-1))
	input = append(input, common.MustRlpEncode(defaultProposalID))
	return common.MustRlpEncode(input)
}

func buildSubmitCancel(nodeID discover.NodeID, pipID string, endVotingRounds uint64, tobeCanceledProposalID common.Hash) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2005))) // func type code
	input = append(input, common.MustRlpEncode(nodeID))       // param 1 ..
	input = append(input, common.MustRlpEncode(pipID))
	input = append(input, common.MustRlpEncode(endVotingRounds))
	input = append(input, common.MustRlpEncode(tobeCanceledProposalID))
	return common.MustRlpEncode(input)
}

func buildVoteInput(nodeIdx int, proposalID common.Hash) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2003)))       // func type code
	input = append(input, common.MustRlpEncode(nodeIdArr[nodeIdx])) // param 1 ...
	input = append(input, common.MustRlpEncode(proposalID))
	input = append(input, common.MustRlpEncode(uint8(1)))
	input = append(input, common.MustRlpEncode(promoteVersion))
	input = append(input, common.MustRlpEncode(versionSign))

	return common.MustRlpEncode(input)
}

func buildVote(nodeIdx int, proposalID common.Hash, option gov.VoteOption, programVersion uint32, sign common.VersionSign) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2003)))       // func type code
	input = append(input, common.MustRlpEncode(nodeIdArr[nodeIdx])) // param 1 ...
	input = append(input, common.MustRlpEncode(proposalID))
	input = append(input, common.MustRlpEncode(option))
	input = append(input, common.MustRlpEncode(programVersion))
	input = append(input, common.MustRlpEncode(sign))

	return common.MustRlpEncode(input)
}

func buildDeclareInput() []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2004))) // func type code
	input = append(input, common.MustRlpEncode(nodeIdArr[0])) // param 1 ...
	input = append(input, common.MustRlpEncode(promoteVersion))
	input = append(input, common.MustRlpEncode(versionSign))
	return common.MustRlpEncode(input)
}

func buildDeclare(nodeID discover.NodeID, declaredVersion uint32, sign common.VersionSign) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2004))) // func type code
	input = append(input, common.MustRlpEncode(nodeID))       // param 1 ...
	input = append(input, common.MustRlpEncode(declaredVersion))
	input = append(input, common.MustRlpEncode(sign))
	return common.MustRlpEncode(input)
}

func buildGetProposalInput(proposalID common.Hash) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2100))) // func type code
	input = append(input, common.MustRlpEncode(proposalID))   // param 1 ...

	return common.MustRlpEncode(input)
}

func buildGetTallyResultInput(idx int) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2101))) // func type code
	input = append(input, common.MustRlpEncode(txHashArr[0])) // param 1 ...

	return common.MustRlpEncode(input)
}

func buildListProposalInput() []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2102))) // func type code
	return common.MustRlpEncode(input)
}

func buildGetActiveVersionInput() []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2103))) // func type code
	return common.MustRlpEncode(input)
}

func buildGetProgramVersionInput() []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2104))) // func type code
	return common.MustRlpEncode(input)
}

func buildGetAccuVerifiersCountInput(proposalID, blockHash common.Hash) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2105))) // func type code
	input = append(input, common.MustRlpEncode(proposalID))
	input = append(input, common.MustRlpEncode(blockHash))
	return common.MustRlpEncode(input)
}

func buildListGovernParam(module string) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2106))) // func type code
	input = append(input, common.MustRlpEncode(module))
	return common.MustRlpEncode(input)
}

func buildGetGovernParamValueInput(module, name string) []byte {
	var input [][]byte
	input = make([][]byte, 0)
	input = append(input, common.MustRlpEncode(uint16(2104)))
	input = append(input, common.MustRlpEncode(module))
	input = append(input, common.MustRlpEncode(name))
	return common.MustRlpEncode(input)
}

func setup(t *testing.T) *mock.Chain {
	t.Log("setup()......")
	//to turn on log's debug level
	//log.Root().SetHandler(log.CallerFileHandler(log.LvlFilterHandler(log.Lvl(6), log.StreamHandler(os.Stderr, log.TerminalFormat(true)))))

	precompiledContract := PlatONPrecompiledContracts[commonvm.GovContractAddr]
	gc, _ = precompiledContract.(*GovContract)
	//default sender of tx, this could be changed in different test case if necessary
	gc.Contract = newContract(common.Big0, sender)

	chain, _ := newChain()
	newPlugins()
	govPlugin = plugin.GovPluginInstance()
	gc.Plugin = govPlugin
	build_staking_data_new(chain)

	if err := gov.InitGenesisGovernParam(chain.SnapDB); err != nil {
		t.Error("error", err)
	}
	gov.RegisterGovernParamVerifiers()

	commit_sndb(chain)

	//the contract will retrieve this txHash as ProposalID
	prepair_sndb(chain, defaultProposalID)
	return chain
}

func clear(chain *mock.Chain, t *testing.T) {
	t.Log("tear down()......")
	if err := chain.SnapDB.Clear(); err != nil {
		t.Error("clear chain.SnapDB error", err)
	}

}

func TestGovContract_SubmitText(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(false, gc, buildSubmitTextInput(), t)
}

func TestGovContract_GetTextProposal(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	//submit a proposal and get it, this tx hash in gc is = txHashArr[1]
	runGovContract(false, gc, buildSubmitTextInput(), t)

	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])

	// get the Proposal by txHashArr[1]
	runGovContract(true, gc, buildGetProposalInput(defaultProposalID), t)
}

func TestGovContract_SubmitText_Sender_wrong(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	gc.Contract.CallerAddress = anotherSender

	runGovContract(false, gc, buildSubmitTextInput(), t, gov.TxSenderDifferFromStaking)
}

func TestGovContract_SubmitText_PIPID_empty(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	runGovContract(false, gc, buildSubmitText(nodeIdArr[1], ""), t, gov.PIPIDEmpty)
}

func TestGovContract_SubmitText_ProposalID_exist(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	runGovContract(false, gc, buildSubmitText(nodeIdArr[1], "pipid1"), t)

	runGovContract(false, gc, buildSubmitText(nodeIdArr[1], "pipid33"), t, gov.ProposalIDExist)
}

func TestGovContract_SubmitText_PIPID_exist(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	runGovContract(false, gc, buildSubmitText(nodeIdArr[1], "pipid1"), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	runGovContract(false, gc, buildSubmitText(nodeIdArr[1], "pipid1"), t, gov.PIPIDExist)
}

func TestGovContract_SubmitText_Proposal_Empty(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	runGovContract(false, gc, buildSubmitText(discover.ZeroNodeID, "pipid1"), t, gov.ProposerEmpty)
}

func TestGovContract_ListGovernParam(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(true, gc, buildListGovernParam("Staking"), t)
}

func TestGovContract_ListGovernParam_all(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(true, gc, buildListGovernParam(""), t)
}

func TestGovContract_GetGovernParamValue(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(true, gc, buildGetGovernParamValueInput("Staking", "StakeThreshold"), t)
}

func TestGovContract_GetGovernParamValue_NotFound(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(true, gc, buildGetGovernParamValueInput("Staking", "StakeThreshold_Err"), t, gov.UnsupportedGovernParam)
}

func TestGovContract_SubmitParam(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	runGovContract(false, gc, buildSubmitParam(nodeIdArr[1], "pipid3", paramModule, paramName, "30"), t)

	p, err := gov.GetProposal(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Fatal("find proposal error", "err", err)
	} else {
		if p == nil {
			t.Fatal("not find proposal error")
		} else {
			pp := p.(*gov.ParamProposal)
			assert.Equal(t, "30", pp.NewValue)
		}
	}
}

func TestGovContract_SubmitParam_thenSubmitParamFailed(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	runGovContract(false, gc, buildSubmitParam(nodeIdArr[1], "pipid3", paramModule, paramName, "30"), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	runGovContract(false, gc, buildSubmitParam(nodeIdArr[2], "pipid4", paramModule, paramName, "35"), t, gov.VotingParamProposalExist)
}

func TestGovContract_SubmitParam_thenSubmitVersionFailed(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	runGovContract(false, gc, buildSubmitParam(nodeIdArr[1], "pipid3", paramModule, paramName, "30"), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	runGovContract(false, gc, buildSubmitVersionInput(), t, gov.VotingParamProposalExist)
}

func TestGovContract_SubmitParam_GetAccuVerifiers(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	value, err := gov.GetGovernParamValue(paramModule, paramName, chain.CurrentHeader().Number.Uint64(), chain.CurrentHeader().Hash())
	if err != nil {
		t.Errorf("%s", err)
	} else {
		assert.Equal(t, "25", value)
	}

	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitParam(nodeIdArr[1], "pipid3", paramModule, paramName, "30"), t)
	//runGovContract(false, gc, buildSubmitTextInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	allVote(chain, t, txHashArr[1])
	commit_sndb(chain)

	runGovContract(true, gc, buildGetAccuVerifiersCountInput(defaultProposalID, chain.CurrentHeader().Hash()), t)

}

func TestGovContract_SubmitParam_Pass(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	value, err := gov.GetGovernParamValue(paramModule, paramName, chain.CurrentHeader().Number.Uint64(), chain.CurrentHeader().Hash())
	if err != nil {
		t.Errorf("%s", err)
	} else {
		assert.Equal(t, "25", value)
	}

	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitParam(nodeIdArr[1], "pipid3", paramModule, paramName, "30"), t)
	//runGovContract(false, gc, buildSubmitTextInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	allVote(chain, t, txHashArr[1])
	commit_sndb(chain)

	runGovContract(true, gc, buildGetAccuVerifiersCountInput(defaultProposalID, chain.CurrentHeader().Hash()), t)

	p, err := gov.GetProposal(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Fatal("find proposal error", "err", err)
	}

	//skip empty block
	skip_emptyBlock(chain, p.GetEndVotingBlock()-1)

	// build_staking_data_more will build a new block base on chain.SnapDB.Current
	build_staking_data_more(chain)
	endBlock(chain, t)
	commit_sndb(chain)

	//at the end of voting block, the status=pass;
	result, err := gov.GetTallyResult(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Errorf("%s", err)
	}
	if result == nil {
		t.Fatal("cannot find the tally result")
	} else if result.Status == gov.Pass {
		t.Log("the result status", result.Status, "yeas", result.Yeas, "accuVerifiers", result.AccuVerifiers)
	} else {
		t.Fatal("tallyResult", "status", result.Status, "yeas", result.Yeas, "accuVerifiers", result.AccuVerifiers)
	}

	//from the next to voting block, the parameter value will be the new value
	skip_emptyBlock(chain, p.GetEndVotingBlock()+1)
	value, err = gov.GetGovernParamValue(paramModule, paramName, chain.CurrentHeader().Number.Uint64(), chain.CurrentHeader().Hash())
	if err != nil {
		t.Errorf("%s", err)
	} else {
		assert.Equal(t, "30", value)
	}
}

func TestGovContract_SubmitVersion(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	runGovContract(false, gc, buildSubmitVersionInput(), t)
}

func TestGovContract_SubmitVersion_thenSubmitParamFailed(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	runGovContract(false, gc, buildSubmitParam(nodeIdArr[1], "pipid3", paramModule, paramName, "30"), t, gov.VotingVersionProposalExist)
}

func TestGovContract_GetVersionProposal(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and get it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	runGovContract(true, gc, buildGetProposalInput(defaultProposalID), t)
}

func TestGovContract_SubmitVersion_AnotherVoting(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	//submit a proposal
	runGovContract(false, gc, buildSubmitVersion(nodeIdArr[1], "versionPIPID", promoteVersion, xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	//submit a proposal
	runGovContract(false, gc, buildSubmitVersion(nodeIdArr[2], "versionPIPID2", promoteVersion, xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())), t, gov.VotingVersionProposalExist)
}

func TestGovContract_SubmitVersion_AnotherPreActive(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and vote for it. the proposalID = txHashArr[1]
	runGovContract(false, gc, buildSubmitVersionInput(), t)

	commit_sndb(chain)

	build_staking_data_more(chain)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	allVote(chain, t, defaultProposalID)
	commit_sndb(chain)

	pTemp, err := gov.GetProposal(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Fatal("find proposal error", "err", err)
	}
	p := pTemp.(*gov.VersionProposal)

	//skip empty blocks
	skip_emptyBlock(chain, p.GetEndVotingBlock()-1)

	// build_staking_data_more will build a new block base on chain.SnapDB.Current
	build_staking_data_more(chain)
	endBlock(chain, t)
	commit_sndb(chain)

	result, err := gov.GetTallyResult(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Errorf("%s", err)
	}
	if result == nil {
		t.Fatal("cannot find the tally result")
	} else if result.Status == gov.PreActive {
		t.Log("the result status", result.Status, "yeas", result.Yeas, "accuVerifiers", result.AccuVerifiers)
	} else {
		t.Fatal("tallyResult", "status", result.Status, "yeas", result.Yeas, "accuVerifiers", result.AccuVerifiers)
	}

	//skip empty blocks, this version proposal is pre-active
	skip_emptyBlock(chain, p.GetActiveBlock()-1)
	//submit another version proposal
	runGovContract(false, gc, buildSubmitVersion(nodeIdArr[2], "versionPIPID2", promoteVersion, xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())), t, gov.PreActiveVersionProposalExist)
}

func TestGovContract_SubmitVersion_NewVersionError(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(false, gc, buildSubmitVersion(nodeIdArr[1], "versionPIPID", uint32(32), xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())), t, gov.NewVersionError)
}

func TestGovContract_SubmitVersion_EndVotingRoundsTooSmall(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(false, gc, buildSubmitVersion(nodeIdArr[1], "versionPIPID", promoteVersion, 0), t, gov.EndVotingRoundsTooSmall)
}

func TestGovContract_SubmitVersion_EndVotingRoundsTooLarge(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	//the default rounds is 6 for developer test net
	runGovContract(false, gc, buildSubmitVersion(nodeIdArr[1], "versionPIPID", promoteVersion, xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())+1), t, gov.EndVotingRoundsTooLarge)
}

func TestGovContract_Float(t *testing.T) {
	t.Log(int(math.Ceil(0.667 * 1000)))
	t.Log(int(math.Floor(0.5 * 1000)))
}

func TestGovContract_DeclareVersion_VotingStage_NotVoted_DeclareActiveVersion(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and get it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	chandler := node.GetCryptoHandler()
	chandler.SetPrivateKey(priKeyArr[0])

	var sign common.VersionSign
	sign.SetBytes(chandler.MustSign(initProgramVersion))

	//fmt.Println("hash:::", chain.CurrentHeader().Hash().Hex())

	runGovContract(false, gc, buildDeclare(nodeIdArr[0], initProgramVersion, sign), t)

	//if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
	//	t.Error("cannot list ActiveNode")
	//} else if len(nodeList) == 0 {
	//	t.Log("in this case, Gov will notify Staking immediately, so, there's no active node list")
	//} else {
	//	t.Fatal("cannot list ActiveNode")
	//}
}

func TestGovContract_DeclareVersion_VotingStage_NotVoted_DeclareNewVersion(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and get it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	chandler := node.GetCryptoHandler()
	chandler.SetPrivateKey(priKeyArr[0])
	runGovContract(false, gc, buildDeclareInput(), t)

	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
		t.Error("cannot list ActiveNode")
	} else if len(nodeList) == 1 {
		t.Log("in this case, Gov will save the declared node, and notify Staking if the proposal is passed later")
	} else {
		t.Fatal("cannot list ActiveNode")
	}
}

func TestGovContract_DeclareVersion_VotingStage_NotVoted_DeclareOtherVersion_Error(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and get it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)

	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])

	chandler := node.GetCryptoHandler()
	chandler.SetPrivateKey(priKeyArr[0])

	otherVersion := uint32(1<<16 | 3<<8 | 0)
	var sign common.VersionSign
	sign.SetBytes(chandler.MustSign(otherVersion))

	runGovContract(false, gc, buildDeclare(nodeIdArr[0], otherVersion, sign), t, gov.DeclareVersionError)

}

func TestGovContract_DeclareVersion_VotingStage_Voted_DeclareNewVersion(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and get it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	chandler := node.GetCryptoHandler()
	chandler.SetPrivateKey(priKeyArr[0])

	//vote new version
	runGovContract(false, gc, buildVoteInput(0, defaultProposalID), t)

	//declare new version
	runGovContract(false, gc, buildDeclare(nodeIdArr[0], promoteVersion, versionSign), t)

	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
		t.Error("cannot list ActiveNode")
	} else if len(nodeList) == 1 {
		t.Log("voted, Gov will save the declared node, and notify Staking if the proposal is passed later")
	} else {
		t.Fatal("cannot list ActiveNode")
	}

}

func TestGovContract_DeclareVersion_VotingStage_Voted_DeclareOldVersion_ERROR(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and get it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])

	chandler := node.GetCryptoHandler()
	chandler.SetPrivateKey(priKeyArr[0])

	//vote new version
	runGovContract(false, gc, buildVoteInput(0, defaultProposalID), t)

	var sign common.VersionSign
	sign.SetBytes(chandler.MustSign(initProgramVersion))

	//vote new version, but declare old version
	runGovContract(false, gc, buildDeclare(nodeIdArr[0], initProgramVersion, sign), t, gov.DeclareVersionError)

	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
		t.Error("cannot list ActiveNode")
	} else if len(nodeList) == 1 {
		t.Log("voted, Gov will save the declared node, and notify Staking if the proposal is passed later")
	} else {
		t.Fatal("cannot list ActiveNode")
	}
}

func TestGovContract_DeclareVersion_VotingStage_Voted_DeclareOtherVersion_ERROR(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and get it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])

	chandler := node.GetCryptoHandler()
	chandler.SetPrivateKey(priKeyArr[0])
	//vote new version
	runGovContract(false, gc, buildVoteInput(0, defaultProposalID), t)

	otherVersion := uint32(1<<16 | 3<<8 | 0)
	var sign common.VersionSign
	sign.SetBytes(chandler.MustSign(otherVersion))

	//vote new version, but declare other version
	runGovContract(false, gc, buildDeclare(nodeIdArr[0], otherVersion, sign), t, gov.DeclareVersionError)

	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
		t.Error("cannot list ActiveNode")
	} else if len(nodeList) == 1 {
		t.Log("voted, Gov will save the declared node, and notify Staking if the proposal is passed later")
	} else {
		t.Fatal("cannot list ActiveNode")
	}
}

func TestGovContract_SubmitCancel(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	runGovContract(false, gc, buildSubmitCancelInput(), t)

}

func TestGovContract_SubmitCancel_AnotherVoting(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	//submit a proposal
	runGovContract(false, gc, buildSubmitVersion(nodeIdArr[0], "versionPIPID", promoteVersion, xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	runGovContract(false, gc, buildSubmitCancel(nodeIdArr[1], "cancelPIPID", xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())-1, defaultProposalID), t)

	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[3])
	runGovContract(false, gc, buildSubmitCancel(nodeIdArr[2], "cancelPIPIDAnother", xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())-1, defaultProposalID), t, gov.VotingCancelProposalExist)
}

func TestGovContract_SubmitCancel_EndVotingRounds_TooLarge(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	runGovContract(false, gc, buildSubmitCancel(nodeIdArr[0], "cancelPIPID", xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds()), defaultProposalID), t, gov.EndVotingRoundsTooLarge)
}

func TestGovContract_SubmitCancel_EndVotingRounds_TobeCanceledNotExist(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(false, gc, buildSubmitVersionInput(), t)

	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	//the version proposal's endVotingRounds=5
	runGovContract(false, gc, buildSubmitCancel(nodeIdArr[0], "cancelPIPID", xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())-1, txHashArr[3]), t, gov.TobeCanceledProposalNotFound)
}

func TestGovContract_SubmitCancel_EndVotingRounds_TobeCanceledNotVersionProposal(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//txHash = txHashArr[2] is a text proposal
	runGovContract(false, gc, buildSubmitTextInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	//try to cancel a text proposal
	runGovContract(false, gc, buildSubmitCancel(nodeIdArr[0], "cancelPIPID", xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())-1, defaultProposalID), t, gov.TobeCanceledProposalTypeError)
}

func TestGovContract_SubmitCancel_EndVotingRounds_TobeCanceledNotAtVotingStage(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	//move the proposal ID from voting-list to end-list
	err := gov.MoveVotingProposalIDToEnd(defaultProposalID, chain.CurrentHeader().Hash())
	if err != nil {
		t.Fatal("err", err)
	}
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[3])
	//try to cancel a closed version proposal
	runGovContract(false, gc, buildSubmitCancel(nodeIdArr[0], "cancelPIPID", xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())-1, defaultProposalID), t, gov.TobeCanceledProposalNotAtVoting)
}

func TestGovContract_GetCancelProposal(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	//submit a proposal and get it.
	runGovContract(false, gc, buildSubmitCancel(nodeIdArr[0], "cancelPIPID", xutil.CalcConsensusRounds(xcom.VersionProposalVote_DurationSeconds())-1, defaultProposalID), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[3])
	runGovContract(true, gc, buildGetProposalInput(txHashArr[2]), t)
}

func TestGovContract_Vote_VersionProposal(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])

	runGovContract(false, gc, buildVoteInput(0, defaultProposalID), t)

	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
		t.Error("cannot list ActiveNode", "err", err)
	} else if len(nodeList) == 1 {
		t.Log("voted, Gov will save the declared node, and notify Staking if the proposal is passed later")
	} else {
		t.Fatal("cannot list ActiveNode")
	}
}
func TestGovContract_Vote_Duplicated(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	runGovContract(false, gc, buildVoteInput(0, defaultProposalID), t)
	runGovContract(false, gc, buildVoteInput(0, defaultProposalID), t, gov.VoteDuplicated)
	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
		t.Error("cannot list ActiveNode")
	} else if len(nodeList) == 1 {
		t.Log("voted duplicated, Gov will count this node once in active node list")
	} else {
		t.Fatal("cannot list ActiveNode")
	}
}

func TestGovContract_Vote_OptionError(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	// vote option = 0, it's wrong
	runGovContract(false, gc, buildVote(0, defaultProposalID, 0, promoteVersion, versionSign), t, gov.VoteOptionError)

	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
		t.Error("cannot list ActiveNode")
	} else if len(nodeList) == 0 {
		t.Log("option error, this node will not be added to active list")
	} else {
		t.Fatal("cannot list ActiveNode")
	}
}

func TestGovContract_Vote_ProposalNotExist(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	chandler := node.GetCryptoHandler()
	chandler.SetPrivateKey(priKeyArr[0])

	var sign common.VersionSign
	sign.SetBytes(chandler.MustSign(initProgramVersion))

	//verify vote new version, but it has not upgraded
	// txIdx=4, not a proposalID
	prepair_sndb(chain, txHashArr[2])
	runGovContract(false, gc, buildVote(0, txHashArr[4], gov.Yes, initProgramVersion, sign), t, gov.ProposalNotFound)

	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
		t.Error("list ActiveNode error", "err", err)
	} else if len(nodeList) == 0 {
		t.Log("proposal not found, this node will not be added to active list")
	} else {
		t.Fatal("cannot list ActiveNode")
	}
}

func TestGovContract_Vote_TextProposalPassed(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitTextInput(), t)
	commit_sndb(chain)

	build_staking_data_more(chain)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	allVote(chain, t, defaultProposalID)
	commit_sndb(chain)

	p, err := gov.GetProposal(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Fatal("find proposal error", "err", err)
	}

	//skip empty block
	skip_emptyBlock(chain, p.GetEndVotingBlock()-1)

	// build_staking_data_more will build a new block base on chain.SnapDB.Current
	build_staking_data_more(chain)
	endBlock(chain, t)
	commit_sndb(chain)

	// vote option = 0, it's wrong
	runGovContract(false, gc, buildVote(0, defaultProposalID, gov.No, promoteVersion, versionSign), t, gov.ProposalNotAtVoting)

	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
		t.Error("cannot list ActiveNode")
	} else if len(nodeList) == 0 {
		t.Log("option error, this node will not be added to active list")
	} else {
		t.Fatal("cannot list ActiveNode")
	}

}

func TestGovContract_Vote_VerifierNotUpgraded(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	//the proposalID will be txHashArr[1]
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	chandler := node.GetCryptoHandler()
	chandler.SetPrivateKey(priKeyArr[0])

	var sign common.VersionSign
	sign.SetBytes(chandler.MustSign(initProgramVersion))

	prepair_sndb(chain, txHashArr[2])
	//verify vote new version, but it has not upgraded
	//txIdx should figure out the proposalID
	runGovContract(false, gc, buildVote(0, defaultProposalID, gov.Yes, initProgramVersion, sign), t, gov.VerifierNotUpgraded)
	commit_sndb(chain)

	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), txHashArr[1]); err != nil {
		t.Error("list ActiveNode error", "err", err)
	} else if len(nodeList) == 0 {
		t.Log("verifier not upgraded, this node will not be added to active list")
	} else {
		t.Fatal("cannot list ActiveNode")
	}
}

func TestGovContract_Vote_ProgramVersionError(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and vote for it. proposalID= txHashArr[1]
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	chandler := node.GetCryptoHandler()
	chandler.SetPrivateKey(priKeyArr[0])

	otherVersion := uint32(1<<16 | 3<<8 | 0)
	var sign common.VersionSign
	sign.SetBytes(chandler.MustSign(otherVersion))

	prepair_sndb(chain, txHashArr[2])
	//verify vote new version, but it has not upgraded
	runGovContract(false, gc, buildVote(0, defaultProposalID, gov.Yes, otherVersion, sign), t, gov.VerifierNotUpgraded)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[3])
	if nodeList, err := gov.GetActiveNodeList(chain.CurrentHeader().Hash(), defaultProposalID); err != nil {
		t.Error("list ActiveNode error", "err", err)
	} else if len(nodeList) == 0 {
		t.Log("verifier program version error, this node will not be added to active list")
	} else {
		t.Fatal("cannot list ActiveNode")
	}
}

func TestGovContract_AllNodeVoteVersionProposal(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and vote for it. proposalID= txHashArr[1]
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	chandler := node.GetCryptoHandler()
	prepair_sndb(chain, txHashArr[2])
	for i := 0; i < 3; i++ {
		chandler.SetPrivateKey(priKeyArr[i])
		var sign common.VersionSign
		sign.SetBytes(chandler.MustSign(promoteVersion))
		//verify vote new version, but it has not upgraded
		runGovContract(false, gc, buildVote(i, defaultProposalID, gov.Yes, promoteVersion, sign), t)
	}
}

func TestGovContract_TextProposal_pass(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitTextInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	allVote(chain, t, txHashArr[1])
	commit_sndb(chain)

	p, err := gov.GetProposal(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Fatal("find proposal error", "err", err)
	}

	//skip empty block
	skip_emptyBlock(chain, p.GetEndVotingBlock()-1)

	// build_staking_data_more will build a new block base on chain.SnapDB.Current
	build_staking_data_more(chain)
	endBlock(chain, t)
	commit_sndb(chain)

	result, err := gov.GetTallyResult(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Errorf("%s", err)
	}
	if result == nil {
		t.Fatal("cannot find the tally result")
	} else if result.Status == gov.Pass {
		t.Log("the result status", result.Status, "yeas", result.Yeas, "accuVerifiers", result.AccuVerifiers)
	} else {
		t.Fatal("tallyResult", "status", result.Status, "yeas", result.Yeas, "accuVerifiers", result.AccuVerifiers)
	}
}

func TestGovContract_VersionProposal_Active(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	//submit a proposal and vote for it. proposalID= txHashArr[1]
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	allVote(chain, t, defaultProposalID)
	commit_sndb(chain)

	pTemp, err := gov.GetProposal(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Fatal("find proposal error", "err", err)
	}
	p := pTemp.(*gov.VersionProposal)

	//skip empty block
	skip_emptyBlock(chain, p.GetEndVotingBlock()-1)

	// build_staking_data_more will build a new block base on chain.SnapDB.Current
	build_staking_data_more(chain)
	endBlock(chain, t)

	commit_sndb(chain)

	result, err := gov.GetTallyResult(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Errorf("%s", err)
	}
	if result == nil {
		t.Fatal("cannot find the tally result")
	} else if result.Status == gov.PreActive {
		t.Log("the result status", result.Status, "yeas", result.Yeas, "accuVerifiers", result.AccuVerifiers)
	} else {
		t.Fatal("tallyResult", "status", result.Status, "yeas", result.Yeas, "accuVerifiers", result.AccuVerifiers)
	}

	//skip empty block
	skip_emptyBlock(chain, p.GetActiveBlock()-1)

	// build_staking_data_more will build a new block base on chain.SnapDB.Current
	build_staking_data_more(chain)
	beginBlock(chain, t)
	commit_sndb(chain)

	result, err = gov.GetTallyResult(defaultProposalID, chain.StateDB)
	if err != nil {
		t.Errorf("%s", err)
	}
	if result == nil {
		t.Fatal("cannot find the tally result")
	} else if result.Status == gov.Active {
		t.Log("the result status", result.Status, "yeas", result.Yeas, "accuVerifiers", result.AccuVerifiers)
	} else {
		t.Fatal("tallyResult", "status", result.Status, "yeas", result.Yeas, "accuVerifiers", result.AccuVerifiers)
	}
}

func TestGovContract_ListProposal(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	runGovContract(false, gc, buildSubmitTextInput(), t)
	commit_sndb(chain)

	runGovContract(true, gc, buildListProposalInput(), t)

}

func TestGovContract_GetActiveVersion(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	runGovContract(true, gc, buildGetActiveVersionInput(), t)
}

func TestGovContract_getAccuVerifiersCount(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)
	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	//get accu verifiers
	runGovContract(true, gc, buildGetAccuVerifiersCountInput(txHashArr[1], chain.CurrentHeader().Hash()), t)
}

func TestGovContract_getAccuVerifiersCount_wrongProposalID(t *testing.T) {
	chain := setup(t)
	defer clear(chain, t)

	//submit a proposal and vote for it.
	runGovContract(false, gc, buildSubmitVersionInput(), t)
	commit_sndb(chain)

	prepair_sndb(chain, txHashArr[2])
	////get accu verifiers
	runGovContract(true, gc, buildGetAccuVerifiersCountInput(txHashArr[2], chain.CurrentHeader().Hash()), t, gov.ProposalNotFound)
}

func runGovContract(callType bool, contract *GovContract, buf []byte, t *testing.T, expectedErrors ...error) {
	res, err := contract.Run(buf)
	assert.True(t, nil == err)

	var result xcom.Result
	if callType {
		err = json.Unmarshal(res, &result)
		assert.True(t, nil == err)
	} else {
		var retCode uint32
		err = json.Unmarshal(res, &retCode)
		assert.True(t, nil == err)
		result.Code = retCode
	}

	if expectedErrors != nil {
		assert.NotEqual(t, common.OkCode, result.Code)
		var expected = false
		for _, expectedError := range expectedErrors {
			expectedCode, expectedMsg := common.DecodeError(expectedError)
			expected = expected || result.Code == expectedCode || strings.Contains(result.Ret, expectedMsg)
		}
		assert.True(t, expected)
		t.Log("the expected errCode:", result.Code, "errMsg:", expectedErrors)
	} else {
		assert.Equal(t, common.OkCode, result.Code)
		t.Log("the expected resultCode:", result.Code)
	}
}

func Test_ResetVoteOption(t *testing.T) {
	v := gov.VoteInfo{}
	v.ProposalID = common.ZeroHash
	v.VoteNodeID = discover.NodeID{}
	v.VoteOption = gov.Abstention
	t.Log(v)

	v.VoteOption = gov.Yes
	t.Log(v)
}

func allVote(chain *mock.Chain, t *testing.T, pid common.Hash) {
	//for _, nodeID := range nodeIdArr {
	currentValidatorList, _ := plugin.StakingInstance().ListCurrentValidatorID(chain.CurrentHeader().Hash(), chain.CurrentHeader().Number.Uint64())
	voteCount := len(currentValidatorList)
	chandler := node.GetCryptoHandler()
	//log.Root().SetHandler(log.CallerFileHandler(log.LvlFilterHandler(log.Lvl(6), log.StreamHandler(os.Stderr, log.TerminalFormat(true)))))
	for i := 0; i < voteCount; i++ {
		vote := gov.VoteInfo{
			ProposalID: pid,
			VoteNodeID: nodeIdArr[i],
			VoteOption: gov.Yes,
		}

		chandler.SetPrivateKey(priKeyArr[i])
		versionSign := common.VersionSign{}
		versionSign.SetBytes(chandler.MustSign(promoteVersion))

		err := gov.Vote(sender, vote, chain.CurrentHeader().Hash(), 1, promoteVersion, versionSign, plugin.StakingInstance(), chain.StateDB)
		if err != nil {
			t.Fatalf("vote err: %s.", err)
		}
	}
}

func beginBlock(chain *mock.Chain, t *testing.T) {
	err := govPlugin.BeginBlock(chain.CurrentHeader().Hash(), chain.CurrentHeader(), chain.StateDB)
	if err != nil {
		t.Fatalf("begin block err... %s", err)
	}
}

func endBlock(chain *mock.Chain, t *testing.T) {
	err := govPlugin.EndBlock(chain.CurrentHeader().Hash(), chain.CurrentHeader(), chain.StateDB)
	if err != nil {
		t.Fatalf("end block err... %s", err)
	}
}
