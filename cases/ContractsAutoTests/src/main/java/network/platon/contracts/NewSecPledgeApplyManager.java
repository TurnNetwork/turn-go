package network.platon.contracts;

import java.util.Arrays;
import java.util.Collections;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.Utf8String;
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
 * <p>Please use the <a href="https://github.com/PlatONnetwork/client-sdk-java/releases">platon-web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/PlatONnetwork/client-sdk-java/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 0.13.1.1.
 */
public class NewSecPledgeApplyManager extends Contract {
    private static final String BINARY = "608060405234801561001057600080fd5b50611bf9806100206000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630766a725146100725780630bc5d1d61461015457806342d937671461023657806363e5f83214610318578063cc71764e146103fa575b600080fd5b34801561007e57600080fd5b506100d9600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506104dc565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101195780820151818401526020810190506100fe565b50505050905090810190601f1680156101465780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561016057600080fd5b506101bb600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506105ef565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101fb5780820151818401526020810190506101e0565b50505050905090810190601f1680156102285780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024257600080fd5b5061029d600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610702565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102dd5780820151818401526020810190506102c2565b50505050905090810190601f16801561030a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561032457600080fd5b5061037f600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610812565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103bf5780820151818401526020810190506103a4565b50505050905090810190601f1680156103ec5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040657600080fd5b50610461600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061093e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104a1578082015181840152602081019050610486565b50505050905090810190601f1680156104ce5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60606000826040518082805190602001908083835b60208310151561051657805182526020820191506020810190506020830392506104f1565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206015016002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105e35780601f106105b8576101008083540402835291602001916105e3565b820191906000526020600020905b8154815290600101906020018083116105c657829003601f168201915b50505050509050919050565b60606000826040518082805190602001908083835b6020831015156106295780518252602082019150602081019050602083039250610604565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020601a016002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106f65780601f106106cb576101008083540402835291602001916106f6565b820191906000526020600020905b8154815290600101906020018083116106d957829003601f168201915b50505050509050919050565b60606000826040518082805190602001908083835b60208310151561073c5780518252602082019150602081019050602083039250610717565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108065780601f106107db57610100808354040283529160200191610806565b820191906000526020600020905b8154815290600101906020018083116107e957829003601f168201915b50505050509050919050565b60606003826040518082805190602001908083835b60208310151561084c5780518252602082019150602081019050602083039250610827565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600081548110151561088d57fe5b90600052602060002090600402016003018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109325780601f1061090757610100808354040283529160200191610932565b820191906000526020600020905b81548152906001019060200180831161091557829003601f168201915b50505050509050919050565b6060610948611732565b610950611732565b6060600061095c61174c565b610964611807565b61096d88611442565b95506109ad6040805190810160405280600181526020017f2d00000000000000000000000000000000000000000000000000000000000000815250611442565b94506109c2858761147090919063ffffffff16565b6040519080825280602002602001820160405280156109f557816020015b60608152602001906001900390816109e05790505b509350600092505b8351831015610a4957610a21610a1c86886114e790919063ffffffff16565b611501565b8484815181101515610a2f57fe5b9060200190602002018190525082806001019350506109fd565b6001846000815181101515610a5a57fe5b906020019060200201519080600181540180825580915050906001820390600052602060002001600090919290919091509080519060200190610a9e929190611830565b5050836000815181101515610aaf57fe5b906020019060200201518260000181905250836001815181101515610ad057fe5b906020019060200201518260200181905250836002815181101515610af157fe5b906020019060200201518260400181905250836003815181101515610b1257fe5b906020019060200201518260600181905250836004815181101515610b3357fe5b906020019060200201518260800181905250836005815181101515610b5457fe5b906020019060200201518260a00181905250836006815181101515610b7557fe5b906020019060200201518260c00181905250836007815181101515610b9657fe5b906020019060200201518260e00181905250836008815181101515610bb757fe5b90602001906020020151826101000181905250836009815181101515610bd957fe5b9060200190602002015182610120018190525083600a815181101515610bfb57fe5b9060200190602002015182610140018190525083600b815181101515610c1d57fe5b9060200190602002015182610160018190525083600c815181101515610c3f57fe5b9060200190602002015182610180018190525083600d815181101515610c6157fe5b90602001906020020151826101a0018190525083600e815181101515610c8357fe5b90602001906020020151826101c0018190525083600f815181101515610ca557fe5b90602001906020020151826101e00181905250836010815181101515610cc757fe5b90602001906020020151826102000181905250836011815181101515610ce957fe5b90602001906020020151826102200181905250836012815181101515610d0b57fe5b90602001906020020151826102400181905250836013815181101515610d2d57fe5b90602001906020020151826102600181905250836014815181101515610d4f57fe5b90602001906020020151826102800181905250836008815181101515610d7157fe5b90602001906020020151816000018190525083600b815181101515610d9257fe5b90602001906020020151816020018190525083600c815181101515610db357fe5b90602001906020020151816040018190525083600d815181101515610dd457fe5b906020019060200201518160600181905250600281908060018154018082558091505090600182039060005260206000209060040201600090919290919091506000820151816000019080519060200190610e309291906118b0565b506020820151816001019080519060200190610e4d9291906118b0565b506040820151816002019080519060200190610e6a9291906118b0565b506060820151816003019080519060200190610e879291906118b0565b5050505060026003856000815181101515610e9e57fe5b906020019060200201516040518082805190602001908083835b602083101515610edd5780518252602082019150602081019050602083039250610eb8565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020908054610f1e929190611930565b50836000815181101515610f2e57fe5b90602001906020020151826102a0015160000181905250836001815181101515610f5457fe5b90602001906020020151826102a0015160200181905250836002815181101515610f7a57fe5b90602001906020020151826102a0015160400181905250836003815181101515610fa057fe5b90602001906020020151826102a0015160600181905250836004815181101515610fc657fe5b90602001906020020151826102a0015160800181905250836004815181101515610fec57fe5b90602001906020020151826102c001516000018190525083600581518110151561101257fe5b90602001906020020151826102c001516020018190525083600681518110151561103857fe5b90602001906020020151826102c001516040018190525081600083600001516040518082805190602001908083835b60208310151561108c5780518252602082019150602081019050602083039250611067565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060008201518160000190805190602001906110db9291906118b0565b5060208201518160010190805190602001906110f89291906118b0565b5060408201518160020190805190602001906111159291906118b0565b5060608201518160030190805190602001906111329291906118b0565b50608082015181600401908051906020019061114f9291906118b0565b5060a082015181600501908051906020019061116c9291906118b0565b5060c08201518160060190805190602001906111899291906118b0565b5060e08201518160070190805190602001906111a69291906118b0565b506101008201518160080190805190602001906111c49291906118b0565b506101208201518160090190805190602001906111e29291906118b0565b5061014082015181600a0190805190602001906112009291906118b0565b5061016082015181600b01908051906020019061121e9291906118b0565b5061018082015181600c01908051906020019061123c9291906118b0565b506101a082015181600d01908051906020019061125a9291906118b0565b506101c082015181600e0190805190602001906112789291906118b0565b506101e082015181600f0190805190602001906112969291906118b0565b506102008201518160100190805190602001906112b49291906118b0565b506102208201518160110190805190602001906112d29291906118b0565b506102408201518160120190805190602001906112f09291906118b0565b5061026082015181601301908051906020019061130e9291906118b0565b5061028082015181601401908051906020019061132c9291906118b0565b506102a08201518160150160008201518160000190805190602001906113539291906118b0565b5060208201518160010190805190602001906113709291906118b0565b50604082015181600201908051906020019061138d9291906118b0565b5060608201518160030190805190602001906113aa9291906118b0565b5060808201518160040190805190602001906113c79291906118b0565b5050506102c082015181601a0160008201518160000190805190602001906113f09291906118b0565b50602082015181600101908051906020019061140d9291906118b0565b50604082015181600201908051906020019061142a9291906118b0565b50505090505081600001519650505050505050919050565b61144a611732565b600060208301905060408051908101604052808451815260200182815250915050919050565b60008082600001516114948560000151866020015186600001518760200151611563565b0190505b8360000151846020015101811115156114e057818060010192505082600001516114d8856020015183038660000151038386600001518760200151611563565b019050611498565b5092915050565b6114ef611732565b6114fa838383611649565b5092915050565b606080600083600001516040519080825280601f01601f19166020018201604052801561153d5781602001602082028038833980820191505090505b50915060208201905061155981856020015186600001516116e7565b8192505050919050565b60008060008060008060008060008b97508c8b1115156116335760208b1115156115ed5760018b60200360080260020a03196001029550858a511694508a8d8d010393508588511692505b846000191683600019161415156115e55783881015156115d2578c8c019850611639565b87806001019850508588511692506115ae565b879850611639565b8a8a209150600096505b8a8d0387111515611632578a8820905080600019168260001916141561161f57879850611639565b60018801975086806001019750506115f7565b5b8c8c0198505b5050505050505050949350505050565b611651611732565b600061166f8560000151866020015186600001518760200151611563565b905084602001518360200181815250508460200151810383600001818152505084600001518560200151018114156116b15760008560000181815250506116dc565b8360000151836000015101856000018181510391508181525050836000015181018560200181815250505b829150509392505050565b60005b60208210151561170f57825184526020840193506020830192506020820391506116ea565b6001826020036101000a0390508019835116818551168181178652505050505050565b604080519081016040528060008152602001600081525090565b6103a0604051908101604052806060815260200160608152602001606081526020016060815260200160608152602001606081526020016060815260200160608152602001606081526020016060815260200160608152602001606081526020016060815260200160608152602001606081526020016060815260200160608152602001606081526020016060815260200160608152602001606081526020016117f4611a28565b8152602001611801611a58565b81525090565b608060405190810160405280606081526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061187157805160ff191683800117855561189f565b8280016001018555821561189f579182015b8281111561189e578251825591602001919060010190611883565b5b5090506118ac9190611a7a565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106118f157805160ff191683800117855561191f565b8280016001018555821561191f579182015b8281111561191e578251825591602001919060010190611903565b5b50905061192c9190611a7a565b5090565b828054828255906000526020600020906004028101928215611a175760005260206000209160040282015b82811115611a165782826000820181600001908054600181600116156101000203166002900461198c929190611a9f565b50600182018160010190805460018160011615610100020316600290046119b4929190611a9f565b50600282018160020190805460018160011615610100020316600290046119dc929190611a9f565b5060038201816003019080546001816001161561010002031660029004611a04929190611a9f565b5050509160040191906004019061195b565b5b509050611a249190611b26565b5090565b60a06040519081016040528060608152602001606081526020016060815260200160608152602001606081525090565b6060604051908101604052806060815260200160608152602001606081525090565b611a9c91905b80821115611a98576000816000905550600101611a80565b5090565b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611ad85780548555611b15565b82800160010185558215611b1557600052602060002091601f016020900482015b82811115611b14578254825591600101919060010190611af9565b5b509050611b229190611a7a565b5090565b611b8291905b80821115611b7e5760008082016000611b459190611b85565b600182016000611b559190611b85565b600282016000611b659190611b85565b600382016000611b759190611b85565b50600401611b2c565b5090565b90565b50805460018160011615610100020316600290046000825580601f10611bab5750611bca565b601f016020900490600052602060002090810190611bc99190611a7a565b5b505600a165627a7a723058207d526968f9f9b8eb50b68f7650aba977233da8f0b7f627fcb4315cc0b313724c0029";

    public static final String FUNC_SELECT_TRADEUSER_BYID = "select_tradeUser_byId";

    public static final String FUNC_SELECT_TRADEOPERATOR_BYTID = "select_tradeOperator_bytId";

    public static final String FUNC_SELECT_SECPLEDGEAPPLY_BYID = "select_SecPledgeApply_byId";

    public static final String FUNC_SELECT_PLEDGESECURITY_BYTID = "select_pledgeSecurity_bytId";

    public static final String FUNC_CREATEPLEDGEAPPLYCOMMON = "createPledgeApplyCommon";

    protected NewSecPledgeApplyManager(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    protected NewSecPledgeApplyManager(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }

    public RemoteCall<String> select_tradeUser_byId(String _id) {
        final Function function = new Function(FUNC_SELECT_TRADEUSER_BYID, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Utf8String(_id)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<String> select_tradeOperator_bytId(String _id) {
        final Function function = new Function(FUNC_SELECT_TRADEOPERATOR_BYTID, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Utf8String(_id)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<String> select_SecPledgeApply_byId(String _id) {
        final Function function = new Function(FUNC_SELECT_SECPLEDGEAPPLY_BYID, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Utf8String(_id)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<String> select_pledgeSecurity_bytId(String _id) {
        final Function function = new Function(FUNC_SELECT_PLEDGESECURITY_BYTID, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Utf8String(_id)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<TransactionReceipt> createPledgeApplyCommon(String secPledgeApplyJson) {
        final Function function = new Function(
                FUNC_CREATEPLEDGEAPPLYCOMMON, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Utf8String(secPledgeApplyJson)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public static RemoteCall<NewSecPledgeApplyManager> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        return deployRemoteCall(NewSecPledgeApplyManager.class, web3j, credentials, contractGasProvider, BINARY,  "", chainId);
    }

    public static RemoteCall<NewSecPledgeApplyManager> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        return deployRemoteCall(NewSecPledgeApplyManager.class, web3j, transactionManager, contractGasProvider, BINARY,  "", chainId);
    }

    public static NewSecPledgeApplyManager load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        return new NewSecPledgeApplyManager(contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    public static NewSecPledgeApplyManager load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        return new NewSecPledgeApplyManager(contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }
}
