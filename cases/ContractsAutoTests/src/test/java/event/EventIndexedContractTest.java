package event;

import beforetest.ContractPrepareTest;
import network.platon.autotest.junit.annotations.DataSource;
import network.platon.autotest.junit.enums.DataSourceType;
import network.platon.contracts.EventIndexedContract;
import network.platon.utils.DataChangeUtil;
import org.junit.Test;
import org.web3j.protocol.core.methods.response.TransactionReceipt;

import java.math.BigInteger;
import java.util.List;

/**
 * @title 事件索引测试
 * @description:
 * @author: albedo
 * @create: 2020/01/06
 */
public class EventIndexedContractTest extends ContractPrepareTest {
    @Test
    @DataSource(type = DataSourceType.EXCEL, file = "test.xls", sheetName = "testOneDimensionalArray",
            author = "albedo", showName = "event.EventIndexedContractTest-一维数组索引")
    public void testOneDimensionalArray() {
        try {
            prepare();
            EventIndexedContract eventTypeContract = EventIndexedContract.deploy(web3j, transactionManager, provider).send();
            String contractAddress = eventTypeContract.getContractAddress();
            String transactionHash = eventTypeContract.getTransactionReceipt().get().getTransactionHash();
            collector.logStepPass("EventIndexedContract issued successfully.contractAddress:" + contractAddress + ", hash:" + transactionHash);
            TransactionReceipt receipt = eventTypeContract.testOneDimensionalArray().send();
            List<EventIndexedContract.OneDimensionalArrayEventEventResponse> one = eventTypeContract.getOneDimensionalArrayEventEvents(receipt);
            byte[] data = one.get(0).array;
            String str=DataChangeUtil.bytesToHex(data);
            String except=one.get(0).log.getTopics().get(1);
            collector.assertEqual("0x"+str, except, "checkout one dimensional array index event");
        } catch (Exception e) {
            collector.logStepFail("EventIndexedContractTest testOneDimensionalArray failure,exception msg:" , e.getMessage());
            e.printStackTrace();
        }
    }

    @Test
    @DataSource(type = DataSourceType.EXCEL, file = "test.xls", sheetName = "testTwoDimensionalArray",
            author = "albedo", showName = "event.EventIndexedContractTest-二维数组索引")
    public void testTwoDimensionalArray() {
        try {
            prepare();
            EventIndexedContract eventCallContract = EventIndexedContract.deploy(web3j, transactionManager, provider).send();
            String contractAddress = eventCallContract.getContractAddress();
            String transactionHash = eventCallContract.getTransactionReceipt().get().getTransactionHash();
            collector.logStepPass("EventIndexedContract issued successfully.contractAddress:" + contractAddress + ", hash:" + transactionHash);
            TransactionReceipt receipt = eventCallContract.testTwoDimensionalArray().send();
            try {
                eventCallContract.getTwoDimensionalArrayEventEvents(receipt);
            } catch (UnsupportedOperationException e) {
                collector.assertEqual(e.getCause().getMessage(),"org.web3j.abi.datatypes.generated.StaticArray2<org.web3j.abi.datatypes.generated.Uint256>");
            }
        } catch (Exception e) {
            collector.logStepFail("EventIndexedContractTest testTwoDimensionalArray failure,exception msg:" , e.getMessage());
            e.printStackTrace();
        }
    }

    @Test
    @DataSource(type = DataSourceType.EXCEL, file = "test.xls", sheetName = "testStr",
            author = "albedo", showName = "event.EventIndexedContractTest-字符串索引")
    public void testStr() {
        try {
            prepare();
            EventIndexedContract eventCallContract = EventIndexedContract.deploy(web3j, transactionManager, provider).send();
            String contractAddress = eventCallContract.getContractAddress();
            String transactionHash = eventCallContract.getTransactionReceipt().get().getTransactionHash();
            collector.logStepPass("EventIndexedContract issued successfully.contractAddress:" + contractAddress + ", hash:" + transactionHash);
            TransactionReceipt receipt = eventCallContract.testStr().send();
            List<EventIndexedContract.StringEventEventResponse> str=eventCallContract.getStringEventEvents(receipt);
            byte[] s=str.get(0).str;
            String strIndexed=DataChangeUtil.bytesToHex(s);
            String except=str.get(0).log.getTopics().get(1);
            collector.assertEqual("0x"+strIndexed, except, "checkout string indexed event");
        } catch (Exception e) {
            collector.logStepFail("EventIndexedContractTest testStr failure,exception msg:" , e.getMessage());
            e.printStackTrace();
        }
    }
    @Test
    @DataSource(type = DataSourceType.EXCEL, file = "test.xls", sheetName = "testEnum",
            author = "albedo", showName = "event.EventIndexedContractTest-枚举类型索引")
    public void testEnum() {
        try {
            prepare();
            EventIndexedContract eventCallContract = EventIndexedContract.deploy(web3j, transactionManager, provider).send();
            String contractAddress = eventCallContract.getContractAddress();
            String transactionHash = eventCallContract.getTransactionReceipt().get().getTransactionHash();
            collector.logStepPass("EventIndexedContract issued successfully.contractAddress:" + contractAddress + ", hash:" + transactionHash);
            TransactionReceipt receipt = eventCallContract.testEnum().send();
            List<EventIndexedContract.EnumEventEventResponse> str=eventCallContract.getEnumEventEvents(receipt);
            BigInteger s=str.get(0).choices;
            String except=str.get(0).log.getTopics().get(1);
            collector.assertEqual(DataChangeUtil.subHexData(s.toString()), DataChangeUtil.subHexData(except), "checkout string indexed event");
        } catch (Exception e) {
            collector.logStepFail("EventIndexedContractTest testStr failure,exception msg:" , e.getMessage());
            e.printStackTrace();
        }
    }

    @Test
    @DataSource(type = DataSourceType.EXCEL, file = "test.xls", sheetName = "testComplex",
            author = "albedo", showName = "event.EventIndexedContractTest-复杂多索引")
    public void testComplex() {
        try {
            prepare();
            EventIndexedContract eventCallContract = EventIndexedContract.deploy(web3j, transactionManager, provider).send();
            String contractAddress = eventCallContract.getContractAddress();
            String transactionHash = eventCallContract.getTransactionReceipt().get().getTransactionHash();
            collector.logStepPass("EventIndexedContract issued successfully.contractAddress:" + contractAddress + ", hash:" + transactionHash);
            TransactionReceipt receipt = eventCallContract.testComplex().send();
            List<EventIndexedContract.ComplexIndexedEventEventResponse> str=eventCallContract.getComplexIndexedEventEvents(receipt);
            String strIndex=DataChangeUtil.bytesToHex(str.get(0).str);
            String arrayIndex=DataChangeUtil.bytesToHex(str.get(0).array);
            String choiceIndex=DataChangeUtil.bytesToHex(str.get(0).choice.toByteArray());
            String exceptStr=str.get(0).log.getTopics().get(3);
            String exceptArray=str.get(0).log.getTopics().get(1);
            String exceptId=str.get(0).log.getTopics().get(2);
            collector.assertEqual("0x"+strIndex, exceptStr, "checkout complex indexes event");
            collector.assertEqual("0x"+arrayIndex, exceptArray, "checkout complex indexes event");
            collector.assertEqual(DataChangeUtil.subHexData(choiceIndex), DataChangeUtil.subHexData(exceptId), "checkout complex indexes event");
        } catch (Exception e) {
            collector.logStepFail("EventIndexedContractTest testComplex failure,exception msg:" , e.getMessage());
            e.printStackTrace();
        }
    }

    @Test
    @DataSource(type = DataSourceType.EXCEL, file = "test.xls", sheetName = "testAnonymousIndexed",
            author = "albedo", showName = "event.EventIndexedContractTest-匿名事件索引数目")
    public void testAnonymousIndexed() {
        try {
            prepare();
            EventIndexedContract eventCallContract = EventIndexedContract.deploy(web3j, transactionManager, provider).send();
            String contractAddress = eventCallContract.getContractAddress();
            String transactionHash = eventCallContract.getTransactionReceipt().get().getTransactionHash();
            collector.logStepPass("EventIndexedContract issued successfully.contractAddress:" + contractAddress + ", hash:" + transactionHash);
            TransactionReceipt receipt = eventCallContract.testAnonymousIndexed().send();
            String u1Topic=receipt.getLogs().get(0).getTopics().get(0);
            String u2Topic=receipt.getLogs().get(0).getTopics().get(1);
            String u3Topic=receipt.getLogs().get(0).getTopics().get(2);
            String u4Topic=receipt.getLogs().get(0).getTopics().get(3);
            collector.assertEqual(DataChangeUtil.subHexData(u1Topic), DataChangeUtil.subHexData("1"), "checkout anonymous index num event");
            collector.assertEqual(DataChangeUtil.subHexData(u2Topic), DataChangeUtil.subHexData("2"), "checkout anonymous index num event");
            collector.assertEqual(DataChangeUtil.subHexData(u3Topic), DataChangeUtil.subHexData("3"), "checkout anonymous index num event");
            collector.assertEqual(DataChangeUtil.subHexData(u4Topic), DataChangeUtil.subHexData("4"), "checkout anonymous index num event");
        } catch (Exception e) {
            collector.logStepFail("EventIndexedContractTest testAnonymousIndexed failure,exception msg:" , e.getMessage());
            e.printStackTrace();
        }
    }

}
