package wasm.data_type;

import network.platon.autotest.junit.annotations.DataSource;
import network.platon.autotest.junit.enums.DataSourceType;
import network.platon.contracts.wasm.ReferenceDataTypeMapContract;
import network.platon.contracts.wasm.ReferenceDataTypeMapTestContract;
import network.platon.contracts.wasm.StorageType_map_string;
import org.junit.Before;
import org.junit.Test;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import wasm.beforetest.WASMContractPrepareTest;

/**
 * @title 测试引用类型Map集合
 * @description:
 * @author: qudong
 * @create: 2020/02/07
 */
public class ReferenceDataTypeMapTest2 extends WASMContractPrepareTest {

    private String key;
    private String value;

    private String keyPerson;
    private String personName;
    private String personAge;

    @Before
    public void before() {
        key = driverService.param.get("key");
        value = driverService.param.get("value");
        keyPerson = driverService.param.get("keyPerson");
        personName = driverService.param.get("personName");
        personAge = driverService.param.get("personAge");
    }


    @Test
    @DataSource(type = DataSourceType.EXCEL, file = "test.xls", sheetName = "Sheet1",
            author = "qudong", showName = "wasm.referenceDataTypeMapTest验证Map集合",sourcePrefix = "wasm")
    public void testReferenceDataTypeMap() {

         //部署合约
        ReferenceDataTypeMapTestContract referenceDataTypeMapTestContract = null;
        try {
            prepare();
            referenceDataTypeMapTestContract = ReferenceDataTypeMapTestContract.deploy(web3j, transactionManager, provider).send();
            String contractAddress = referenceDataTypeMapTestContract.getContractAddress();
            TransactionReceipt tx = referenceDataTypeMapTestContract.getTransactionReceipt().get();
            collector.logStepPass("referenceDataTypeMapContract issued successfully.contractAddress:" + contractAddress
                                  + ", hash:" + tx.getTransactionHash());
            collector.logStepPass("deployFinishCurrentBlockNumber:" + tx.getBlockNumber());
        } catch (Exception e) {
            collector.logStepFail("referenceDataTypeMapContract deploy fail.", e.toString());
            e.printStackTrace();
        }
        //调用合约方法
        try {
            //1、验证：map中的key与value可以是任意类型
            TransactionReceipt  transactionReceipt = referenceDataTypeMapTestContract.setMapKeyType().send();
            collector.logStepPass("referenceDataTypeMapContract 【验证map中的key与value可以是任意类型】 successfully hash:" + transactionReceipt.getTransactionHash());
            //2、验证：string类型map容器赋值
            TransactionReceipt  transactionReceipt1 = referenceDataTypeMapTestContract.addMapString(key,value).send();
            collector.logStepPass("referenceDataTypeMapContract 【验证string类型map容器赋值】 执行storageType_map_string() successfully hash:" + transactionReceipt1.getTransactionHash());
            //3、验证：map容器数量
            Long actualValueSize = referenceDataTypeMapTestContract.getMapStringSize().send();
            collector.logStepPass("referenceDataTypeMapContract 【验证map容器数量】 执行getMapStringSize() actualValueSize:" + actualValueSize);
            collector.assertEqual(actualValueSize,Long.parseLong("1"), "checkout  execute success.");
            //3、验证：map容器根据key获取值
            //Byte expectMapLength = 3;
            String actualValue = referenceDataTypeMapTestContract.getMapValueByString(key).send();
            collector.logStepPass("referenceDataTypeMapContract 【验证map容器根据key获取值】 执行getMapValueByString() actualValue:" + actualValue);
            collector.assertEqual(actualValue,value, "checkout  execute success.");
            //4、验证：person类型map容器赋值
            ReferenceDataTypeMapTestContract.Person person = new ReferenceDataTypeMapTestContract.Person();
            person.name = personName;
            person.age  = Long.parseLong(personAge);
            TransactionReceipt  transactionReceipt2 = referenceDataTypeMapTestContract.addMapByPerson(Byte.valueOf(keyPerson),person).send();
            collector.logStepPass("referenceDataTypeMapContract 【验证person类型map容器赋值】 执行setMapByPerson() successfully hash:" + transactionReceipt2.getTransactionHash());
            //4、验证：map容器数量
            Long actualPersonSize = referenceDataTypeMapTestContract.getMapByPersonSize().send();
            collector.logStepPass("referenceDataTypeMapContract 【验证map容器数量】 执行getMapByPersonSize() actualPersonSize:" + actualPersonSize);
            collector.assertEqual(actualPersonSize,Long.parseLong("1"), "checkout  execute success.");
            //5、验证：map容器根据key获取值Person
            String actualValueName = referenceDataTypeMapTestContract.getMapByPerson(Byte.valueOf(keyPerson)).send();
            collector.logStepPass("referenceDataTypeMapContract 【验证map容器根据key获取值】 执行getMapByPerson() actualValueName:" + actualValueName);
            //collector.assertEqual(actualValuePerson.name,personName, "checkout  execute success.");

        } catch (Exception e) {
            collector.logStepFail("referenceDataTypeMapContract Calling Method fail.", e.toString());
            e.printStackTrace();
        }

    }
}
