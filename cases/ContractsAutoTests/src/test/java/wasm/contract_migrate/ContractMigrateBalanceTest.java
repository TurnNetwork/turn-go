package wasm.contract_migrate;

import network.platon.autotest.junit.annotations.DataSource;
import network.platon.autotest.junit.enums.DataSourceType;
import network.platon.autotest.utils.FileUtil;
import network.platon.contracts.wasm.ContractMigrate_v1;
import network.platon.utils.RlpUtil;
import org.junit.Test;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.rlp.RlpEncoder;
import org.web3j.rlp.RlpString;
import org.web3j.tx.Transfer;
import org.web3j.utils.Convert;
import org.web3j.protocol.core.DefaultBlockParameterName;

import wasm.beforetest.WASMContractPrepareTest;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.math.BigDecimal;
import java.math.BigInteger;
import java.nio.file.Paths;
import java.util.ArrayList;

/**
 * @title contract migrate
 * @description:
 * @author: yuanwenjun
 * @create: 2020/02/12
 */
public class ContractMigrateBalanceTest extends WASMContractPrepareTest {

    //the file name of migrate contract
    private String wasmFile = "ContractMigrate_v1.bin";

    @Test
    @DataSource(type = DataSourceType.EXCEL, file = "test.xls", sheetName = "Sheet1",
            author = "yuanwenjun", showName = "wasm.contract_migrate",sourcePrefix = "wasm")
    public void testMigrateContractBalance() {

        Byte[] init_arg = null;
        Long transfer_value = 100000L;
        Long gas_value = 200000L;
        BigInteger origin_contract_value = BigInteger.valueOf(10000);

        try {
            prepare();
            ContractMigrate_v1 contractMigratev1 = ContractMigrate_v1.deploy(web3j, transactionManager, provider).send();
            String contractAddress = contractMigratev1.getContractAddress();
            String transactionHash = contractMigratev1.getTransactionReceipt().get().getTransactionHash();
            collector.logStepPass("ContractMigrateV1 issued successfully.contractAddress:" + contractAddress + ", hash:" + transactionHash);
            
            Transfer transfer = new Transfer(web3j, transactionManager);
            transfer.sendFunds(contractAddress, new BigDecimal(origin_contract_value), Convert.Unit.VON).send();
            BigInteger originBalance = web3j.platonGetBalance(contractAddress, DefaultBlockParameterName.LATEST).send().getBalance();
            collector.logStepPass("origin contract balance is: " + originBalance);

            String filePath = FileUtil.pathOptimization(Paths.get("src", "test", "resources", "contracts", "wasm", "contract_migrate").toUri().getPath());
            init_arg = RlpUtil.loadInitArg(filePath+File.separator+wasmFile,new ArrayList<String>());
            TransactionReceipt transactionReceipt = contractMigratev1.migrate_contract(init_arg,transfer_value,gas_value).send();
            collector.logStepPass("Contract Migrate V1  successfully hash:" + transactionReceipt.getTransactionHash());
            
            BigInteger originAfterMigrateBalance = web3j.platonGetBalance(contractAddress, DefaultBlockParameterName.LATEST).send().getBalance();
            collector.logStepPass("After migrate, origin contract balance is: " + originAfterMigrateBalance);
            collector.assertEqual(originAfterMigrateBalance, BigInteger.valueOf(0), "checkout origin contract balance");
            
            String newContractAddress = contractMigratev1.getTransferEvents(transactionReceipt).get(0).arg1;
            collector.logStepPass("new Contract Address is:"+newContractAddress);
            BigInteger newMigrateBalance = web3j.platonGetBalance(newContractAddress, DefaultBlockParameterName.LATEST).send().getBalance();
            collector.logStepPass("new contract balance is: " + newMigrateBalance);
            collector.assertEqual(newMigrateBalance, origin_contract_value.add(BigInteger.valueOf(transfer_value)), "checkout new contract balance");

        } catch (Exception e) {
            collector.logStepFail("ContractDistoryTest failure,exception msg:" , e.getMessage());
            e.printStackTrace();
        }
    }
}
