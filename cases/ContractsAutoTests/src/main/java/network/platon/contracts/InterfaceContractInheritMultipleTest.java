package network.platon.contracts;

import java.math.BigInteger;
import java.util.Arrays;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.GasProvider;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://github.com/PlatONnetwork/client-sdk-java/releases">platon-web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/PlatONnetwork/client-sdk-java/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 0.9.1.2-SNAPSHOT.
 */
public class InterfaceContractInheritMultipleTest extends Contract {
    private static final String BINARY = "608060405234801561001057600080fd5b50610118806100206000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c806399ecedf6146037578063cad0899b146080575b600080fd5b606a60048036036040811015604b57600080fd5b81019080803590602001909291908035906020019092919050505060c9565b6040518082815260200191505060405180910390f35b60b360048036036040811015609457600080fd5b81019080803590602001909291908035906020019092919050505060d6565b6040518082815260200191505060405180910390f35b6000818303905092915050565b600081830190509291505056fea265627a7a72315820a0c85ec21b0bd83ae664b6e98477e73353f5d0574d71f39986fe9b787bc5b04d64736f6c634300050d0032";

    public static final String FUNC_REDUCE = "reduce";

    public static final String FUNC_SUM = "sum";

    @Deprecated
    protected InterfaceContractInheritMultipleTest(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected InterfaceContractInheritMultipleTest(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected InterfaceContractInheritMultipleTest(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected InterfaceContractInheritMultipleTest(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<BigInteger> reduce(BigInteger c, BigInteger d) {
        final Function function = new Function(FUNC_REDUCE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(c), 
                new org.web3j.abi.datatypes.generated.Uint256(d)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<BigInteger> sum(BigInteger a, BigInteger b) {
        final Function function = new Function(FUNC_SUM, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(a), 
                new org.web3j.abi.datatypes.generated.Uint256(b)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public static RemoteCall<InterfaceContractInheritMultipleTest> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return deployRemoteCall(InterfaceContractInheritMultipleTest.class, web3j, credentials, contractGasProvider, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<InterfaceContractInheritMultipleTest> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(InterfaceContractInheritMultipleTest.class, web3j, credentials, gasPrice, gasLimit, BINARY, "");
    }

    public static RemoteCall<InterfaceContractInheritMultipleTest> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return deployRemoteCall(InterfaceContractInheritMultipleTest.class, web3j, transactionManager, contractGasProvider, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<InterfaceContractInheritMultipleTest> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(InterfaceContractInheritMultipleTest.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, "");
    }

    @Deprecated
    public static InterfaceContractInheritMultipleTest load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new InterfaceContractInheritMultipleTest(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static InterfaceContractInheritMultipleTest load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new InterfaceContractInheritMultipleTest(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static InterfaceContractInheritMultipleTest load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new InterfaceContractInheritMultipleTest(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static InterfaceContractInheritMultipleTest load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new InterfaceContractInheritMultipleTest(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
