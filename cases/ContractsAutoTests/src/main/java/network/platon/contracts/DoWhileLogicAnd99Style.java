package network.platon.contracts;

import java.math.BigInteger;
import java.util.Arrays;
import java.util.Collections;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.GasProvider;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 0.7.5.0.
 */
public class DoWhileLogicAnd99Style extends Contract {
    private static final String BINARY = "608060405234801561001057600080fd5b506101d6806100206000396000f3fe608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806349de2f961461006757806361cc1193146100925780637cf7eab0146100cd578063c1633e2214610108575b600080fd5b34801561007357600080fd5b5061007c610133565b6040518082815260200191505060405180910390f35b34801561009e57600080fd5b506100cb600480360360208110156100b557600080fd5b810190808035906020019092919050505061013c565b005b3480156100d957600080fd5b50610106600480360360208110156100f057600080fd5b8101908080359060200190929190505050610176565b005b34801561011457600080fd5b5061011d6101a0565b6040518082815260200191505060405180910390f35b60008054905090565b6000600a8201905060006009830190505b6001830192508083111561016057610161565b5b818310151561014d5782600181905550505050565b60008090505b8181101561019c578060005401600081905550808060010191505061017c565b5050565b600060015490509056fea165627a7a72305820b84b4376997cc4269d165ab5382abc7acf8983b6815facc14a33270d416c05400029";

    public static final String FUNC_GETFORSUM = "getForSum";

    public static final String FUNC_DOWHILE = "dowhile";

    public static final String FUNC_FORSUM = "forsum";

    public static final String FUNC_GETDOWHILESUM = "getDoWhileSum";

    @Deprecated
    protected DoWhileLogicAnd99Style(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected DoWhileLogicAnd99Style(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected DoWhileLogicAnd99Style(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected DoWhileLogicAnd99Style(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<BigInteger> getForSum() {
        final Function function = new Function(FUNC_GETFORSUM, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<TransactionReceipt> dowhile(BigInteger x) {
        final Function function = new Function(
                FUNC_DOWHILE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(x)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> forsum(BigInteger x) {
        final Function function = new Function(
                FUNC_FORSUM, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.generated.Uint256(x)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<BigInteger> getDoWhileSum() {
        final Function function = new Function(FUNC_GETDOWHILESUM, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public static RemoteCall<DoWhileLogicAnd99Style> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return deployRemoteCall(DoWhileLogicAnd99Style.class, web3j, credentials, contractGasProvider, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<DoWhileLogicAnd99Style> deploy(Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(DoWhileLogicAnd99Style.class, web3j, credentials, gasPrice, gasLimit, BINARY, "");
    }

    public static RemoteCall<DoWhileLogicAnd99Style> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return deployRemoteCall(DoWhileLogicAnd99Style.class, web3j, transactionManager, contractGasProvider, BINARY, "");
    }

    @Deprecated
    public static RemoteCall<DoWhileLogicAnd99Style> deploy(Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return deployRemoteCall(DoWhileLogicAnd99Style.class, web3j, transactionManager, gasPrice, gasLimit, BINARY, "");
    }

    @Deprecated
    public static DoWhileLogicAnd99Style load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new DoWhileLogicAnd99Style(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static DoWhileLogicAnd99Style load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new DoWhileLogicAnd99Style(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static DoWhileLogicAnd99Style load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new DoWhileLogicAnd99Style(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static DoWhileLogicAnd99Style load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new DoWhileLogicAnd99Style(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
