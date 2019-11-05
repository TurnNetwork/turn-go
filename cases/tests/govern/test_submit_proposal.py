from common.log import log
from dacite import from_dict
from tests.lib import Genesis
import pytest
from tests.lib.utils import get_pledge_list, upload_platon, wait_block_number
from tests.lib.client import get_client_obj
import time, math

def test_VP_SU_001(submit_version):
    pip_obj = submit_version
    proposalinfo = pip_obj.get_effect_proposal_info_of_vote()
    log.info('获取升级提案信息为{}'.format(proposalinfo))
    endvotingblock_count = math.ceil(proposalinfo.get('SubmitBlock') / pip_obj.economic.consensus_size + 5
                                     ) * pip_obj.economic.consensus_size - 20
    log.info('计算投票截止块高为{},接口返回投票截止块高{}'.format(endvotingblock_count,
                                               proposalinfo.get('EndVotingBlock')))
    assert int(endvotingblock_count) == proposalinfo.get('EndVotingBlock')
    assert int(endvotingblock_count) + 21 == proposalinfo.get('ActiveBlock')

@pytest.mark.P0
def test_CP_SU_001_CP_UN_001(submit_cancel):
    pip_obj = submit_cancel
    proposalinfo = pip_obj.get_effect_proposal_info_of_vote(4)
    log.info('获取取消提案信息为{}'.format(proposalinfo))
    endvotingblock_count = math.ceil(proposalinfo.get('SubmitBlock') / pip_obj.economic.consensus_size + 4
                                     ) * pip_obj.economic.consensus_size - 20
    log.info('计算投票截止块高为{},接口返回投票截止块高{}'.format(endvotingblock_count,
                                               proposalinfo.get('EndVotingBlock')))
    assert int(endvotingblock_count) == proposalinfo.get('EndVotingBlock')
    pip_obj.submitCancel(pip_obj.node.node_id, str(time.time()), 1, proposalinfo.get('ProposalID'),
                         pip_obj.node.staking_address, transaction_cfg=pip_obj.cfg.transaction_cfg)

def test_VP_VE_001_to_VP_VE_004(no_version_proposal):
    pip_obj_tmp = no_version_proposal
    result = pip_obj_tmp.submitVersion(pip_obj_tmp.node.node_id, str(time.time()), pip_obj_tmp.cfg.version1, 1,
                                       pip_obj_tmp.node.staking_address, transaction_cfg=pip_obj_tmp.cfg.transaction_cfg)
    assert result.get("Code") == 302011

    result = pip_obj_tmp.submitVersion(pip_obj_tmp.node.node_id, str(time.time()), pip_obj_tmp.cfg.version2, 1,
                                       pip_obj_tmp.node.staking_address, transaction_cfg=pip_obj_tmp.cfg.transaction_cfg)
    assert result.get("Code") == 302011

    result = pip_obj_tmp.submitVersion(pip_obj_tmp.node.node_id, str(time.time()), pip_obj_tmp.cfg.version3, 1,
                                       pip_obj_tmp.node.staking_address, transaction_cfg=pip_obj_tmp.cfg.transaction_cfg)
    assert result.get("Code") == 302011

    result = pip_obj_tmp.submitVersion(pip_obj_tmp.node.node_id, str(time.time()), pip_obj_tmp.chain_version, 1,
                                   pip_obj_tmp.node.staking_address, transaction_cfg=pip_obj_tmp.cfg.transaction_cfg)
    assert result.get("Code") == 302011

def test_VP_WA_001(no_version_proposal):
    pip_obj_tmp = no_version_proposal
    address, _ = pip_obj_tmp.economic.account.generate_account(pip_obj_tmp.node.web3, 10**18 * 10000000)
    result = pip_obj_tmp.submitVersion(pip_obj_tmp.node.node_id, str(time.time()), pip_obj_tmp.cfg.version5, 1,
                                       address, transaction_cfg=pip_obj_tmp.cfg.transaction_cfg)
    log.info('发起升级提案结果为{}'.format(result))
    assert result.get('Code') == 302021

def test_TP_WA_001(client_verifier_obj):
    pip_obj = client_verifier_obj
    address, _ = pip_obj.economic.account.generate_account(pip_obj.node.web3, 10**18 * 10000000)
    result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 1,
                                       address, transaction_cfg=pip_obj.cfg.transaction_cfg)
    log.info('发起升级提案结果为{}'.format(result))
    assert result.get('Code') == 302021

def test_TP_UN_001(submit_text):
    pip_obj = submit_text
    result = pip_obj.submitText(pip_obj.node.node_id, str(time.time()), pip_obj.node.staking_address,
                                transaction_cfg=pip_obj.cfg.transaction_cfg)
    log.info('存在处于投票期文本提案，再次发起文本提案结果为{}'.format(result))
    assert result.get('Code') == 0

def test_VP_SU_001_VP_UN_001(submit_version):
    pip_obj = submit_version
    result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 1,
                                   pip_obj.node.staking_address, transaction_cfg=pip_obj.cfg.transaction_cfg)
    log.info('有处于投票期的升级提案，再次发起升级提案结果为{}'.format(result))
    assert result.get('Code') == 302012

def test_VP_UN_002_CP_ID_002(submit_version, client_list_obj):
    pip_obj = submit_version
    proposalinfo = pip_obj.get_effect_proposal_info_of_vote()
    log.info('获取升级提案信息为{}'.format(proposalinfo))
    verifier_list = get_pledge_list(pip_obj.node.ppos.getVerifierList)
    log.info('verifier_list:{}'.format(verifier_list))
    for nodeid in verifier_list:
        pip_obj_tmp = get_client_obj(nodeid, client_list_obj).pip
        log.info('节点id为{}'.format(pip_obj_tmp.node.node_id))
        log.info('替换节点{}版本'.format(pip_obj_tmp.node.node_id))
        upload_platon(pip_obj_tmp.node, pip_obj.cfg.PLATON_NEW_BIN)
        log.info('进行节点重启')
        pip_obj_tmp.node.restart()
        result = pip_obj_tmp.vote(pip_obj_tmp.node.node_id, proposalinfo.get("ProposalID"), pip_obj_tmp.cfg.vote_option_yeas,
                                  pip_obj_tmp.node.staking_address, transaction_cfg=pip_obj_tmp.cfg.transaction_cfg)
        log.info('节点{}升级提案投票结果为{}'.format(pip_obj_tmp.node.node_id, result))
        assert result.get('Code') == 0
    wait_block_number(pip_obj.node, proposalinfo.get('EndVotingBlock'))
    status = pip_obj.get_status_of_proposal(proposalinfo.get('ProposalID'))
    log.info('投票截止快高，升级提案状态为{}'.format(status))
    assert status == 4

    result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 1,
                                   pip_obj.node.staking_address,
                                   transaction_cfg=pip_obj.cfg.transaction_cfg)
    log.info('链上存在预生效的升级提案，再次发起升级提案结果为{}'.format(result))
    assert result.get('Code') == 302013

    result = pip_obj.submitCancel(pip_obj.node.node_id, str(time.time()), 1, proposalinfo.get('ProposalID'),
                                  pip_obj.node.staking_address, transaction_cfg=pip_obj.cfg.transaction_cfg)
    log.info('there is preactive version proposal, submit cancel proposal result: {}'.format(result))
    assert result.get('Code') == 302017

class TestEndVotingRounds():
    def test_VP_CR_001_VP_CR_002_VP_CR_007_TP_TE_002(self, pip_env, client_verifier_obj):
        '''
        投票周期配置成N个共识周期对应的秒数+1
        :param pip_env:
        :param pip_obj:
        :return:
        '''
        pip_obj = client_verifier_obj
        genesis = from_dict(data_class=Genesis, data=pip_env.genesis_config)
        genesis.EconomicModel.Gov.VersionProposalVote_DurationSeconds = 2 * pip_obj.cfg.consensus_block + 1
        genesis.EconomicModel.Gov.TextProposalVote_DurationSeconds = 5 * pip_obj.cfg.consensus_block + 1
        pip_env.set_genesis(genesis.to_dict())
        pip_env.deploy_all()
        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 3,
                                       pip_obj.node.staking_address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('投票共识轮数为3，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 302010

        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 0,
                                       pip_obj.node.staking_address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('投票共识轮数为0，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 302009

        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 2,
                                       pip_obj.node.staking_address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('投票共识轮数为2，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 0

        result = pip_obj.submitText(pip_obj.node.node_id, str(time.time()), pip_obj.node.staking_address,
                                    transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('发起文本提案结果为{}'.format(result))
        assert result.get('Code') == 0
        proposalinfo = pip_obj.get_effect_proposal_info_of_vote(pip_obj.cfg.text_proposal)
        log.info('获取文本提案信息{}'.format(proposalinfo))
        endvotingblock_count = math.ceil(proposalinfo.get('SubmitBlock') / pip_obj.economic.consensus_size + 5
                                         ) * pip_obj.economic.consensus_size - 20
        log.info('计算投票截止块高为{},接口返回投票截止块高{}'.format(endvotingblock_count,
                                                   proposalinfo.get('EndVotingBlock')))
        assert int(endvotingblock_count) == proposalinfo.get('EndVotingBlock')

    def test_VP_CR_003_VP_CR_004_VP_CR_007_TP_TE_003(self, pip_env, client_verifier_obj):
        pip_obj = client_verifier_obj.pip
        genesis = from_dict(data_class=Genesis, data=pip_env.genesis_config)
        genesis.EconomicModel.Gov.VersionProposalVote_DurationSeconds = 3 * pip_obj.cfg.consensus_block - 1
        genesis.EconomicModel.Gov.TextProposalVote_DurationSeconds = 5 * pip_obj.cfg.consensus_block - 1
        pip_env.set_genesis(genesis.to_dict())
        pip_env.deploy_all()
        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 3, pip_obj.node.staking_address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('投票共识轮数为3，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 302010

        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 0, pip_obj.node.staking_address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('投票共识轮数为0，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 302009

        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 2, pip_obj.node.staking_address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('投票共识轮数为2，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 0

        result = pip_obj.submitText(pip_obj.node.node_id, str(time.time()), pip_obj.node.staking_address,
                                    transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('发起文本提案结果为{}'.format(result))
        assert result.get('Code') == 0
        proposalinfo = pip_obj.get_effect_proposal_info_of_vote(pip_obj.cfg.text_proposal)
        log.info('获取文本提案信息{}'.format(proposalinfo))
        endvotingblock_count = math.ceil(proposalinfo.get('SubmitBlock') / pip_obj.economic.consensus_size + 4
                                         ) * pip_obj.economic.consensus_size - 20
        log.info('计算投票截止块高为{},接口返回投票截止块高{}'.format(endvotingblock_count,
                                                   proposalinfo.get('EndVotingBlock')))
        assert int(endvotingblock_count) == proposalinfo.get('EndVotingBlock')

    def test_VP_CR_005_VP_CR_006_TP_TE_001(self, pip_env, client_verifier_obj):
        pip_obj = client_verifier_obj.pip
        genesis = from_dict(data_class=Genesis, data=pip_env.genesis_config)
        genesis.EconomicModel.Gov.VersionProposalVote_DurationSeconds = 3 * pip_obj.economic.consensus_size
        genesis.EconomicModel.Gov.TextProposalVote_DurationSeconds = 5 * pip_obj.economic.consensus_size
        pip_env.set_genesis(genesis.to_dict())
        pip_env.deploy_all()
        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 4,
                                       pip_obj.node.staking_address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('投票共识轮数为3，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 302010

        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 0,
                                       pip_obj.node.staking_address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('投票共识轮数为0，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 302009

        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 3,
                                       pip_obj.node.staking_address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('投票共识轮数为2，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 0

        result = pip_obj.submitText(pip_obj.node.node_id, str(time.time()), pip_obj.node.staking_address,
                                    transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('发起文本提案结果为{}'.format(result))
        assert result.get('Code') == 0
        proposalinfo = pip_obj.get_effect_proposal_info_of_vote(pip_obj.cfg.text_proposal)
        log.info('获取文本提案信息{}'.format(proposalinfo))
        endvotingblock_count = math.ceil(proposalinfo.get('SubmitBlock') / pip_obj.economic.consensus_size + 5
                                         ) * pip_obj.economic.consensus_size - 20
        log.info('计算投票截止块高为{},接口返回投票截止块高{}'.format(endvotingblock_count,
                                                   proposalinfo.get('EndVotingBlock')))
        assert int(endvotingblock_count) == proposalinfo.get('EndVotingBlock')

        proosalinfo = pip_obj.get_effect_proposal_info_of_vote(1)
        log.info('text proposalinfo: {}'.format(proosalinfo))
        result = pip_obj.submitCancel(pip_obj.node.node_id, str(time.time()), 1, proosalinfo.get('ProposalID'),
                             pip_obj.node.staking_address, transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('submit cancel result: {}'.format(result))

class TestNoVerifierSubmitProposal():
    def test_VP_PR_002_TP_PR_002(self, client_new_node_obj):
        pip_obj = client_new_node_obj
        address, _ = pip_obj.economic.account.generate_account(pip_obj.node.web3, 10**18 * 10000000)
        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()), pip_obj.cfg.version5, 1, address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('新节点发起版本声明，结果为{}'.format(result))
        assert result.get('Code') == 302022

        result = pip_obj.submitText(pip_obj.node.node_id, str(time.time()), address,
                                       transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('新节点发起版本声明，结果为{}'.format(result))
        assert result.get('Code') == 302022

    def test_VP_PR_001_TP_PR_001(self, client_candidate_obj):
        pip_obj = client_candidate_obj
        result = pip_obj.submitVersion(pip_obj.node.node_id, str(time.time()),
                                                pip_obj.cfg.version5, 1, pip_obj.node.staking_address,
                                                transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('候选节点节点{}发起升级提案，结果为{}'.format(pip_obj.node.node_id, result))
        assert result.get('Code') == 302022

        result = pip_obj.submitText(pip_obj.node.node_id, str(time.time()), pip_obj.node.staking_address,
                                                transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('候选节点节点{}发起文本提案，结果为{}'.format(pip_obj.node.node_id, result))
        assert result.get('Code') == 302022

    def test_VP_PR_003_VP_PR_004_TP_PR_003_TP_PR_004(self, client_verifier_obj):
        address = client_verifier_obj.node.staking_address
        result = client_verifier_obj.staking.withdrew_staking(address)
        log.info('节点{}发起退质押结果为{}'.format(client_verifier_obj.node.node_id, result))
        assert result.get("Code") == 0
        log.info(client_verifier_obj.economic.account.find_pri_key(address))
        result = client_verifier_obj.pip.submitVersion(client_verifier_obj.node.node_id, str(time.time()),
                                                       client_verifier_obj.pip.cfg.version5, 1, address,
                                              transaction_cfg=client_verifier_obj.pip.cfg.transaction_cfg)
        log.info('节点退出中，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 302020

        result = client_verifier_obj.pip.submitText(client_verifier_obj.node.node_id, str(time.time()), address,
                                           transaction_cfg=client_verifier_obj.pip.cfg.transaction_cfg)
        log.info('节点退出中，发起文本提案结果为{}'.format(result))
        assert result.get('Code') == 302020

        client_verifier_obj.economic.wait_settlement_blocknum(client_verifier_obj.node,
                                                              number=client_verifier_obj.economic.unstaking_freeze_ratio)
        result = client_verifier_obj.pip.submitVersion(client_verifier_obj.node.node_id, str(time.time()),
                                                       client_verifier_obj.pip.cfg.version5, 1, address,
                                              transaction_cfg=client_verifier_obj.pip.cfg.transaction_cfg)
        log.info('节点已退出，发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 302022

        client_verifier_obj.economic.wait_settlement_blocknum(client_verifier_obj.node,
                                                              number=client_verifier_obj.economic.unstaking_freeze_ratio)
        result = client_verifier_obj.pip.submitText(client_verifier_obj.node.node_id, str(time.time()), address,
                                           transaction_cfg=client_verifier_obj.pip.cfg.transaction_cfg)
        log.info('节点已退出，发起文本提案结果为{}'.format(result))
        assert result.get('Code') == 302022

class TestSubmitCancel():
    @pytest.mark.P0
    def test_CP_WA_001(self, submit_version):
        pip_obj = submit_version
        address, _ = pip_obj.economic.account.generate_account(pip_obj.node.web3, 10**18 * 10000)
        proposalinfo = pip_obj.get_effect_proposal_info_of_vote()
        log.info('获取升级提案信息为{}'.format(proposalinfo))
        result = pip_obj.submitCancel(pip_obj.node.node_id, str(time.time()), 1, proposalinfo.get('ProposalID'), address,
                             transaction_cfg=pip_obj.cfg.transaction_cfg)
        assert result.get('Code') == 302021

    @pytest.mark.P0
    def test_CP_PR_001(self, new_node_has_proposal):
        pip_obj = new_node_has_proposal
        address, _ = pip_obj.economic.account.generate_account(pip_obj.node.web3, 10**18 * 10000)
        proposalinfo = pip_obj.get_effect_proposal_info_of_vote()
        log.info('升级提案信息为{}'.format(proposalinfo))
        result = pip_obj.submitCancel(pip_obj.node.node_id, str(time.time()), 1,
                                                    proposalinfo.get('ProposalID'),
                                               address, transaction_cfg=pip_obj.cfg.transaction_cfg)
        assert result.get('Code') == 302022

    @pytest.mark.P0
    def test_CP_PR_002(self, candidate_has_proposal):
        pip_obj = candidate_has_proposal
        proposalinfo = pip_obj.get_effect_proposal_info_of_vote()
        log.info('升级提案信息为{}'.format(proposalinfo))
        result = pip_obj.submitCancel(pip_obj.node.node_id, str(time.time()), 1,
                                                    proposalinfo.get('ProposalID'),
                                               pip_obj.node.staking_address, transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('候选人发起升级提案结果为{}'.format(result))
        assert result.get('Code') == 302022

    @pytest.mark.P2
    def test_CP_PR_003_CP_PR_004(self, submit_version, client_list_obj):
        pip_obj = submit_version
        proposalinfo = pip_obj.get_effect_proposal_info_of_vote()
        client_obj = get_client_obj(pip_obj.node.node_id, client_list_obj)
        address = client_obj.node.staking_address
        log.info('proposalinfo: {}'.format(proposalinfo))
        result = client_obj.staking.withdrew_staking(address)
        log.info('nodeid: {} withdrewstaking result: {}'.format(client_obj.node.node_id, result))
        assert result.get("Code") == 0
        result = client_obj.pip.submitCancel(client_obj.node.node_id, str(time.time()), 1,
                                                      proposalinfo.get('ProposalID'), address,
                                              transaction_cfg=client_obj.pip.cfg.transaction_cfg)
        log.info('node exiting，cancel proposal result: {}'.format(result))
        assert result.get('Code') == 302020

        client_obj.economic.wait_settlement_blocknum(client_obj.node,
                                                     number=client_obj.economic.unstaking_freeze_ratio)
        result = client_obj.pip.submitCancel(client_obj.node.node_id, str(time.time()), 1,
                                                      proposalinfo.get('ProposalID'), address,
                                              transaction_cfg=client_obj.pip.cfg.transaction_cfg)
        log.info('exited node，cancel proposal result: {}'.format(result))
        assert result.get('Code') == 302022

    def test_CP_CR_001(self, submit_version):
        pip_obj = submit_version
        proposalinfo = pip_obj.get_effect_proposal_info_of_vote()
        log.info('proposalinfo: {}'.format(proposalinfo))
        endvoting_rounds = (math.ceil(proposalinfo.get('EndVotingBlock')/pip_obj.economic.consensus_size) - math.ceil(
            pip_obj.node.block_number/pip_obj.economic.consensus_size)) / pip_obj.economic.consensus_size
        result = pip_obj.submitCancel(pip_obj.node.node_id, str(time.time()), endvoting_rounds,
                                                      proposalinfo.get('ProposalID'), pip_obj.node.staking_address,
                                              transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('endvoting_rounds:{}， cancel proposal result:{}'.format(endvoting_rounds, result))
        assert result.get('Code') == 302009

        result = pip_obj.submitCancel(pip_obj.node.node_id, str(time.time()), endvoting_rounds+1,
                                                      proposalinfo.get('ProposalID'), pip_obj.node.staking_address,
                                              transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('endvoting_rounds:{}， cancel proposal result:{}'.format(endvoting_rounds+1, result))
        assert result.get('Code') == 302009

    def test_CP_ID_001(self, no_version_proposal):
        pip_obj = no_version_proposal
        result = pip_obj.submitCancel(pip_obj.node.node_id, str(time.time()), 1,
                                                      '0x49b83cfc4b99462f7131d14d80c73b6657237753cd1e878e8d62dc2e9f574123',
                             pip_obj.node.staking_address, transaction_cfg=pip_obj.cfg.transaction_cfg)
        log.info('cancel proposal result: {}'.format(result))
        assert result.get('Code') == 302015