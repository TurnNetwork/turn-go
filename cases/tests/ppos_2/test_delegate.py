# -*- coding: utf-8 -*-

from tests.lib.utils import *
import pytest
from tests.lib.config import EconomicConfig


@pytest.mark.P1
def test_DI_001_009(client_new_node_obj):
    """
    :param client_new_node_obj:
    :param get_generate_account:
    :return:
    """
    address, pri_key = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                             10 ** 18 * 10000000)
    client_new_node_obj.staking.create_staking(0, address, address)
    address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                        10 ** 18 * 10000000)
    result = client_new_node_obj.delegate.delegate(0, address1)
    assert_code(result, 0)
    msg = client_new_node_obj.ppos.getCandidateInfo(client_new_node_obj.node.node_id)
    staking_blocknum = msg["Ret"]["StakingBlockNum"]
    msg = client_new_node_obj.ppos.getDelegateInfo(staking_blocknum, address1, client_new_node_obj.node.node_id)
    log.info(msg)
    assert client_new_node_obj.node.web3.toChecksumAddress(msg["Ret"]["Addr"]) == address1
    assert msg["Ret"]["NodeId"] == client_new_node_obj.node.node_id
    assert msg["Ret"]["ReleasedHes"] == client_new_node_obj.economic.delegate_limit


@pytest.mark.P1
def test_DI_002_003_004(client_new_node_obj_list):
    """
    :param client_new_node_obj_list:
    :return:
    """
    address, pri_key = client_new_node_obj_list[0].economic.account.generate_account(
        client_new_node_obj_list[0].node.web3,
        10 ** 18 * 10000000)
    client_new_node_obj_list[0].staking.create_staking(0, address, address, amount=1500000000000000000000000)

    address, pri_key = client_new_node_obj_list[1].economic.account.generate_account(
        client_new_node_obj_list[1].node.web3,
        10 ** 18 * 10000000)
    client_new_node_obj_list[1].staking.create_staking(1, address, address, amount=2000000000000000000000000)

    address, pri_key = client_new_node_obj_list[2].economic.account.generate_account(
        client_new_node_obj_list[2].node.web3,
        10 ** 18 * 10000000)
    client_new_node_obj_list[2].staking.create_staking(0, address, address, amount=2500000000000000000000000)

    client_new_node_obj_list[2].economic.wait_settlement_blocknum(client_new_node_obj_list[2].node)
    client_new_node_obj_list[2].economic.wait_consensus_blocknum(client_new_node_obj_list[2].node)

    nodeid_list2 = get_pledge_list(client_new_node_obj_list[2].ppos.getVerifierList)
    log.info(nodeid_list2)

    assert client_new_node_obj_list[0].node.node_id not in nodeid_list2

    address1, _ = client_new_node_obj_list[0].economic.account.generate_account(client_new_node_obj_list[0].node.web3,
                                                                                10 ** 18 * 10000000)
    # The candidate delegate
    result = client_new_node_obj_list[0].delegate.delegate(0, address1)
    assert_code(result, 0)

    assert client_new_node_obj_list[2].node.node_id in nodeid_list2
    address2, _ = client_new_node_obj_list[2].economic.account.generate_account(client_new_node_obj_list[2].node.web3,
                                                                                10 ** 18 * 10000000)
    # The verifier delegates
    result = client_new_node_obj_list[2].delegate.delegate(0, address2)
    assert_code(result, 0)
    nodeid_list3 = get_pledge_list(client_new_node_obj_list[2].ppos.getValidatorList)
    log.info(nodeid_list3)
    assert client_new_node_obj_list[2].node.node_id in nodeid_list3
    address3, _ = client_new_node_obj_list[2].economic.account.generate_account(client_new_node_obj_list[2].node.web3,
                                                                                10 ** 18 * 10000000)
    # Consensus verifier delegates
    result = client_new_node_obj_list[2].delegate.delegate(0, address3)
    assert_code(result, 0)


@pytest.mark.P3
def test_DI_005(client_consensus_obj):
    """
    The amount entrusted by the client is less than the threshold
    :param client_consensus_obj:
    :return:
    """
    address, _ = client_consensus_obj.economic.account.generate_account(client_consensus_obj.node.web3,
                                                                        10 ** 18 * 10000000)

    result = client_consensus_obj.delegate.delegate(0, address)
    log.info(result)
    assert_code(result, 301107)


@pytest.mark.P1
def test_DI_006(client_new_node_obj):
    """

    :param client_new_node_obj:
    :return:
    """
    address, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                       10 ** 18 * 10000000)

    client_new_node_obj.staking.create_staking(0, address, address)
    address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                        10 ** 18 * 10000000)
    delegate_limit = client_new_node_obj.economic.delegate_limit
    result = client_new_node_obj.delegate.delegate(0, address1, amount=delegate_limit - 1)
    log.info(result)
    assert_code(result, 301105)


@pytest.mark.P1
def test_DI_007(client_new_node_obj):
    """

    :param client_new_node_obj:
    :return:
    """
    address, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                       10 ** 18 * 10000000)

    client_new_node_obj.staking.create_staking(0, address, address)
    address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                        10 ** 18 * 10000000)

    fig = {"gas": 1}
    status = 0
    try:
        result = client_new_node_obj.delegate.delegate(0, address1, transaction_cfg=fig)
        log.info(result)
    except:
        status = 1
    assert status == 1


@pytest.mark.P1
def test_DI_008(client_new_node_obj):
    """
    :param client_new_node_obj:
    :return:
    """
    address, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                       10 ** 18 * 10000000)

    client_new_node_obj.staking.create_staking(0, address, address)
    address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                        10)
    status = 0
    try:
        result = client_new_node_obj.delegate.delegate(0, address1)
        log.info(result)
    except:
        status = 1
    assert status == 1


@pytest.mark.P3
def test_DI_010_020(client_new_node_obj):
    """
    Delegate to a candidate who doesn't exist
    :param client_new_node_obj:
    :return:
    """
    illegal_nodeID = "7ee3276fd6b9c7864eb896310b5393324b6db785a2528c00cc28ca8c" \
                     "3f86fc229a86f138b1f1c8e3a942204c03faeb40e3b22ab11b8983c35dc025de42865990"
    address1, pri_key = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                              10 ** 18 * 10000000)
    result = client_new_node_obj.delegate.delegate(0, address1, node_id=illegal_nodeID)
    log.info(result)


@pytest.mark.P1
@pytest.mark.parametrize('status', [0, 1, 2, 3])
def test_DI_011_012_013_014(client_new_node_obj, status):
    """
    0:A valid candidate whose commission is still in doubt
    1:The delegate is also a valid candidate at a lockup period
    2:A candidate whose mandate is voluntarily withdrawn but who is still in the freeze period
    3:A candidate whose mandate has been voluntarily withdrawn and whose freeze period has expired
    :param client_new_node_obj:
    :param status:
    :return:
    """

    address, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                       10 ** 18 * 10000000)

    client_new_node_obj.staking.create_staking(0, address, address)
    if status == 0:
        # A valid candidate whose commission is still in doubt
        address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                            10 ** 18 * 10000000)
        result = client_new_node_obj.delegate.delegate(0, address1)
        assert_code(result, 0)

    if status == 1:
        # The delegate is also a valid candidate at a lockup period
        address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                            10 ** 18 * 10000000)
        client_new_node_obj.economic.wait_settlement_blocknum(client_new_node_obj.node)
        result = client_new_node_obj.delegate.delegate(0, address1)
        assert_code(result, 0)

    if status == 2:
        address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                            10 ** 18 * 10000000)
        client_new_node_obj.economic.wait_settlement_blocknum(client_new_node_obj.node)
        result = client_new_node_obj.staking.withdrew_staking(address)
        assert_code(result, 0)
        result = client_new_node_obj.delegate.delegate(0, address1)
        assert_code(result, 301103)

    if status == 3:
        address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                            10 ** 18 * 10000000)
        client_new_node_obj.economic.wait_settlement_blocknum(client_new_node_obj.node)
        result = client_new_node_obj.staking.withdrew_staking(address)
        assert_code(result, 0)
        client_new_node_obj.economic.wait_settlement_blocknum(client_new_node_obj.node, number=2)
        result = client_new_node_obj.delegate.delegate(0, address1)
        log.info(result)
        assert_code(result, 301102)


@pytest.mark.P1
def test_DI_015_016(client_new_node_obj, greater_than_staking_amount, client_consensus_obj):
    """
    :param client_new_node_obj:
    :param greater_than_staking_amount:
    :param client_consensus_obj:
    :return:
    """
    address, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                       10 ** 18 * 10000000)
    address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                        10 ** 18 * 10000000)
    result = client_new_node_obj.staking.create_staking(0, address, address, amount=greater_than_staking_amount)
    assert_code(result, 0)
    log.info("Close one node")
    client_new_node_obj.node.stop()
    node = client_consensus_obj.node
    log.info("The next two periods")
    client_new_node_obj.economic.wait_settlement_blocknum(node, number=2)
    log.info("Restart the node")
    client_new_node_obj.node.start()
    result = client_new_node_obj.delegate.delegate(0, address1)
    log.info(result)
    assert_code(result, 301103)
    log.info("Next settlement period")
    client_new_node_obj.economic.wait_settlement_blocknum(node)
    result = client_new_node_obj.delegate.delegate(0, address1)
    log.info(result)
    assert_code(result, 301102)


@pytest.mark.P1
def test_DI_017(client_new_node_obj):
    """
    Use the pledge account as the entrustment
    :param client_new_node_obj:
    :return:
    """
    address, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                       10 ** 18 * 10000000)

    result = client_new_node_obj.staking.create_staking(0, address, address)
    assert_code(result, 0)
    result = client_new_node_obj.delegate.delegate(0, address)
    log.info(result)
    assert_code(result, 301106)


@pytest.mark.P1
def test_DI_019(client_new_node_obj):
    """
    The verification section receives the delegate, exits, becomes the verification node, and receives the delegate
    :param client_new_node_obj:
    :return:
    """
    address, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                       10 ** 18 * 10000000)
    result = client_new_node_obj.staking.create_staking(0, address, address)
    assert_code(result, 0)

    address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                        10 ** 18 * 10000000)
    result = client_new_node_obj.delegate.delegate(0, address1)
    log.info(result)
    # Exit the pledge
    result = client_new_node_obj.staking.withdrew_staking(address)
    assert_code(result, 0)
    # Repeat pledge
    result = client_new_node_obj.staking.create_staking(0, address, address)
    assert_code(result, 0)
    result = client_new_node_obj.delegate.delegate(0, address1)
    log.info(result)
    # Recheck wallet associations
    msg = client_new_node_obj.ppos.getRelatedListByDelAddr(address1)
    log.info(msg)
    print(len(msg["Ret"]))
    assert len(msg["Ret"]) == 2
    for i in msg["Ret"]:
        assert client_new_node_obj.node.web3.toChecksumAddress(i["Addr"]) == address1
        assert i["NodeId"] == client_new_node_obj.node.node_id


@pytest.mark.P3
def test_DI_021(client_new_node_obj, client_consensus_obj, greater_than_staking_amount):
    """

    :param client_new_node_obj:
    :return:
    """
    address, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                       10 ** 18 * 10000000)
    address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                        10 ** 18 * 10000000)
    result = client_new_node_obj.staking.create_staking(0, address, address, amount=greater_than_staking_amount)
    assert_code(result, 0)
    result = client_new_node_obj.delegate.delegate(0, address1)
    log.info(result)
    msg = client_new_node_obj.ppos.getCandidateInfo(client_new_node_obj.node.node_id)
    staking_blocknum = msg["Ret"]["StakingBlockNum"]
    log.info("Close one node")
    client_new_node_obj.node.stop()
    node = client_consensus_obj.node
    log.info("The next two periods")
    client_new_node_obj.economic.wait_settlement_blocknum(node, number=2)
    log.info("Restart the node")
    client_new_node_obj.node.start()
    msg = client_consensus_obj.ppos.getDelegateInfo(staking_blocknum, address1, client_new_node_obj.node.node_id)
    log.info(msg)


@pytest.mark.P2
@pytest.mark.parametrize('status', [0, 1, 2])
def test_DI_022_023_024(client_new_node_obj, status):
    """
    0:There is only the free amount of hesitation period when additional entrusting
    1:Only the free amount of the lockup period exists when the delegate is added
    2:The amount of both hesitation period and lockup period exists when additional entrustment is made
    :param client_new_node_obj:
    :param status:
    :return:
    """
    client_new_node_obj.economic.env.deploy_all()
    address, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                       10 ** 18 * 10000000)

    client_new_node_obj.staking.create_staking(0, address, address)
    address1, _ = client_new_node_obj.economic.account.generate_account(client_new_node_obj.node.web3,
                                                                        10 ** 18 * 10000000)
    result = client_new_node_obj.delegate.delegate(0, address1)
    log.info(result)
    msg = client_new_node_obj.ppos.getCandidateInfo(client_new_node_obj.node.node_id)
    staking_blocknum = msg["Ret"]["StakingBlockNum"]

    if status == 0:
        result = client_new_node_obj.delegate.delegate(0, address1)
        log.info(result)
        msg = client_new_node_obj.ppos.getDelegateInfo(staking_blocknum, address1, client_new_node_obj.node.node_id)
        log.info(msg)

    if status == 1:
        client_new_node_obj.economic.wait_settlement_blocknum(client_new_node_obj.node)
        result = client_new_node_obj.delegate.delegate(0, address1)
        log.info(result)
        msg = client_new_node_obj.ppos.getDelegateInfo(staking_blocknum, address1, client_new_node_obj.node.node_id)
        log.info(msg)

    if status == 2:
        client_new_node_obj.economic.wait_settlement_blocknum(client_new_node_obj.node)
        result = client_new_node_obj.delegate.delegate(0, address1)
        log.info(result)
        result = client_new_node_obj.delegate.delegate(0, address1)
        log.info(result)
        msg = client_new_node_obj.ppos.getDelegateInfo(staking_blocknum, address1, client_new_node_obj.node.node_id)
        log.info(msg)


