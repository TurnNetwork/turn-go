package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Uint64;
import com.platon.rlp.datatypes.WasmAddress;
import java.math.BigInteger;
import java.util.Arrays;
import org.web3j.abi.WasmFunctionEncoder;
import org.web3j.abi.datatypes.WasmFunction;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.WasmContract;
import org.web3j.tx.gas.GasProvider;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://github.com/PlatONnetwork/client-sdk-java/releases">platon-web3j command line tools</a>,
 * or the org.web3j.codegen.WasmFunctionWrapperGenerator in the 
 * <a href="https://github.com/PlatONnetwork/client-sdk-java/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with platon-web3j version 0.13.1.4.
 */
public class SimpleContract extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001530f60017f017f60027f7f0060017f0060000060027f7f017f60037f7f7f0060047f7f7f7f0060037f7f7f017f60047f7f7f7f017f60017f017e60027f7e0060037e7e7f006000017f60017e017f60027e7e017f02a9020d03656e760c706c61746f6e5f70616e6963000303656e760d726c705f6c6973745f73697a65000003656e760f706c61746f6e5f726c705f6c697374000503656e760e726c705f62797465735f73697a65000403656e7610706c61746f6e5f726c705f6279746573000503656e760d726c705f753132385f73697a65000e03656e760f706c61746f6e5f726c705f75313238000b03656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000c03656e7610706c61746f6e5f6765745f696e707574000203656e7617706c61746f6e5f6765745f73746174655f6c656e677468000403656e7610706c61746f6e5f6765745f7374617465000803656e7610706c61746f6e5f7365745f7374617465000603656e760d706c61746f6e5f72657475726e00010327260303000705090902000200000d010a01000403080201010001010401020102070000030000060405017001010105030100020608017f0141c088040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f7273000d0f5f5f66756e63735f6f6e5f65786974001f06696e766f6b65000e0a9b37260400102f0be70502057f027e23004190026b22002400102f10072201100f22021008200041c8006a200020022001101022014100101102400240200041c8006a10122205500d004180081013200551044020011014200041c8006a101510160c020b4185081013200551044020011017410247044010000b200041c8006a200141011011200041c8006a10122105200041c8006a101522012005370310200110160c020b4189081013200551044020011014200041c8006a101522022903102105200041c8016a1018220120051019101a20012005101b200128020c200141106a28020047044010000b20012802002001280204100c200128020c22030440200120033602100b200210160c020b418d0810132005510440200041d8016a22024100360200200041d0016a22034200370300200042003703c80120011017410247044010000b200041c8006a200141011011200041c8006a200041c8016a101c200041c8006a10152101200041286a20022802002202360200200041206a20032903002205370300200020002903c8012206370318200141386a2006370300200141406b2005370300200141c8006a2002360200200110160c020b41990810132005520d0020011014200041a8016a2201200041c8006a1015220241c8006a280000360200200041a0016a2203200241406b2900003703002000200241386a29000037039801200041c8016a1018220420004198016a101d101a200041c0016a20012802002201360200200041b8016a20032903002205370300200020002903980122063703b001200041e8016a2005370300200041f0016a2001360200200041206a2005370300200041286a2001360200200020063703e0012000200637031820004188026a200136020020004180026a2005370300200020063703f8012004200041f8016a101e220128020c200141106a28020047044010000b20012802002001280204100c200128020c22030440200120033602100b200210160c010b10000b101f20004190026a24000b9b0101047f230041106b220124002001200036020c2000047f41bc08200041086a2202411076220041bc082802006a220336020041b80841b808280200220420026a41076a417871220236020002400240200341107420024d044041bc08200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20042001410c6a4104102c41086a0541000b2100200141106a240020000b0c00200020012002411c10200bc90202067f017e230041106b220324002001280208220520024b0440200341086a2001102620012003280208200328020c102736020c200320011026410021052001027f410020032802002207450d001a410020032802042208200128020c2206490d001a200820062006417f461b210420070b360210200141146a2004360200200141003602080b200141106a210603402001280214210402402005200249044020040d01410021040b200020062802002004411410201a200341106a24000f0b20032001102641002104027f410020032802002205450d001a410020032802042208200128020c2207490d001a200820076b2104200520076a0b2105200120043602142001200536021020032006410020052004102710322001200329030022093702102001200128020c2009422088a76a36020c2001200128020841016a22053602080c000b000b870202047f017e230041106b2203240020001021024002402000280204450d00200010210240200028020022022c0000220141004e044020010d010c020b200141807f460d00200141ff0171220441b7014d0440200028020441014d04401000200028020021020b20022d00010d010c020b200441bf014b0d012000280204200141ff017141ca7e6a22014d04401000200028020021020b200120026a2d0000450d010b2000280204450d0020022d000041c001490d010b10000b200341086a20001022200328020c220041094f044010000b200328020821010340200004402000417f6a210020013100002005420886842105200141016a21010c010b0b200341106a240020050b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b0e0020001017410147044010000b0b960401087f230041406a22012400200042d1f0fad48ae09ad34537030820004200370300200141286a101822022000290308101b200228020c200241106a28020047044010000b02402002280200220720022802042206100922054504400c010b2001410036022020014200370318200141186a200510232007200620012802182204200128021c220820046b100a417f47044020002001200441016a20082004417f736a10101012370310200521030b2004450d002001200436021c0b200228020c22040440200220043602100b2003450440200020002903003703100b2000420037001841002104200041286a4100360000200041206a4200370000200041386a22054200370000200041306a220342dda4afedfcbbf09c977f370300200041406b4200370000200041c8006a4100360000200141286a101822022003290300101b200228020c200241106a28020047044010000b0240200228020022062002280204220810092207450d002001410036022020014200370318200141186a200710232006200820012802182203200128021c220620036b100a417f4704402001200341016a20062003417f736a10102005101c200721040b2003450d002001200336021c0b200228020c22030440200220033602100b20044504402005200041186a2202290300370300200541106a200241106a280200360200200541086a200241086a2903003703000b200141406b240020000b9e0502087f027e230041a0016b2203240020034188016a10182202200041306a22012903001019101a20022001290300101b200041386a2101200228020c200241106a28020047044010000b2002280204210720022802002108200341f0006a101821052001101d21062005200341186a1024220410252005200620042802046a20042802006b101a200341106a200141106a2800002206360200200341086a200141086a290000220937030020032001290000220a370300200341306a2009370300200341386a2006360200200341c8006a2009370300200341d0006a20063602002003200a3703282003200a370340200341e8006a2006360200200341e0006a20093703002003200a37035802402005200341d8006a101e220128020c200141106a280200460440200141046a2105200128020021060c010b200141046a2105100020012802002106200128020c2001280210460d0010000b2008200720062005280200100b200428020022050440200420053602040b200128020c22040440200120043602100b200228020c22010440200220013602100b20034188016a1018220120002903081019101a20012000290308101b200128020c200141106a28020047044010000b2001280204210520012802002106200341f0006a101821022000290310101921072002200341d8006a1024220410252002200720042802046a20042802006b101a20022000290310101b0240200228020c200241106a280200460440200241046a2107200228020021080c010b200241046a2107100020022802002108200228020c2002280210460d0010000b2006200520082007280200100b200428020022050440200420053602040b200228020c22040440200220043602100b200128020c22020440200120023602100b200341a0016a24000b800101047f230041106b2201240002402000280204450d0020002802002d000041c001490d00200141086a20001026200128020c210003402000450d01200141002001280208220320032000102722046a20034520002004497222031b3602084100200020046b20031b2100200241016a21020c000b000b200141106a240020020b29002000410036020820004200370200200041001028200041146a41003602002000420037020c20000bb00102037f017e230041206b22012400200141186a4100360200200141106a4200370300200141086a4200370300200142003703004101210320004280015a0440034020002004845045044020044238862000420888842100200241016a2102200442088821040c010b0b024020024138490d002002210303402003450d01200241016a2102200341087621030c000b000b200241016a21030b2001200336020020014104721029200141206a240020030b1300200028020820014904402000200110280b0b2801017f2000420020011005200028020422026a102a42002001200220002802006a10062000102b0bcf0102037f027e230041406a2202240020001021200241386a20001022200228023c2103024002402000280204450d00200341144b0d0020002802002d000041c001490d010b10000b200241306a22004100360200200241286a220442003703002002420037032020022003411420034114491b22036b41346a20022802382003102c1a200241186a20002802002200360200200241106a20042903002205370300200220022903202206370308200141106a2000360000200141086a200537000020012006370000200241406b24000b850202037f027e23004180016b22012400200141306a4100360200200141286a4200370300200141206a4200370300200141086a200041086a2900002204370300200141106a200041106a280000220236020020014200370318200120002900002205370300200141406b2004370300200141c8006a2002360200200141d8006a2004370300200141e0006a20023602002001200537033820012005370350200141f8006a2002360200200141f0006a200437030020012005370368410121020240034020034114460d01200141e8006a20036a2100200341016a210320002d0000450d000b411521020b20012002360218200141186a410472102920014180016a240020020b2a01017f2000200141141003200028020422026a102a20014114200220002802006a10042000102b20000b880101037f41a808410136020041ac082802002100034020000440034041b00841b0082802002201417f6a2202360200200141014845044041a8084100360200200020024102746a22004184016a280200200041046a28020011020041a808410136020041ac0828020021000c010b0b41b008412036020041ac08200028020022003602000c010b0b0b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000103020024f0d002003410471044010000c010b200042003702000b02402003411071450d002000103020024d0d0020034104710440100020000f0b200042003702000b20000b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bd50101047f2001102d2204200128020422024b04401000200128020421020b200128020021052000027f02400240200204404100210120052c00002203417f4a0d01027f200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a21010c010b4101210120050d000c010b41002103200120046a20024b0d0020022001490d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000bfd0101067f024020002802042202200028020022046b220520014904402000280208220720026b200120056b22034f04400340200241003a00002000200028020441016a22023602042003417f6a22030d000c030b000b2001200720046b2202410174220420042001491b41ffffffff07200241ffffffff03491b220104402001102e21060b200520066a220521020340200241003a0000200241016a21022003417f6a22030d000b200120066a210420052000280204200028020022066b22016b2103200141014e0440200320062001102c1a0b2000200436020820002002360204200020033602000f0b200520014d0d002000200120046a3602040b0b3a01017f200041003602082000420037020020004101102e2201360200200141fe013a00002000200141016a22013602082000200136020420000b6101037f200028020c200041106a28020047044010000b200028020422022001280204200128020022036b22016a220420002802084b047f20002004102820002802040520020b20002802006a20032001102c1a2000200028020420016a3602040b2101017f2001102d220220012802044b044010000b2000200120011031200210320b2701017f230041206b22022400200241086a200020014114102010302100200241206a240020000b2f01017f200028020820014904402001100f20002802002000280204102c210220002001360208200020023602000b0b860201067f200028020422032000280210220141087641fcffff07716a2102027f200320002802082204460440200041146a210641000c010b2003200028021420016a220541087641fcffff07716a280200200541ff07714102746a2105200041146a21062002280200200141ff07714102746a0b21010340024020012005460440200641003602000340200420036b41027522014103490d022000200341046a22033602040c000b000b200141046a220120022802006b418020470d0120022802042101200241046a21020c010b0b2001417f6a220241014d04402000418004418008200241016b1b3602100b03402003200447044020002004417c6a22043602080c010b0b0b3601017f200028020820014904402001100f20002802002000280204102c210220002001360208200020023602000b200020013602040b7a01037f0340024020002802102201200028020c460d00200141786a2802004504401000200028021021010b200141786a22022002280200417f6a220336020020030d002000200236021020002001417c6a2802002201200028020420016b220210016a102a200120002802006a22012002200110020c010b0b0bfc0801067f03400240200020046a2105200120046a210320022004460d002003410371450d00200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220745044003402006411049450440200020046a2203200120046a2205290200370200200341086a200541086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2205200120046a2204290200370200200441086a2103200541086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002007417f6a220741024b0d00024002400240024002400240200741016b0e020102000b2005200120046a220328020022073a0000200541016a200341016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2203200120046a220541046a2802002202410874200741187672360200200341046a200541086a2802002207410874200241187672360200200341086a2005410c6a28020022024108742007411876723602002003410c6a200541106a2802002207410874200241187672360200200441106a2104200641706a21060c000b000b2005200120046a220328020022073a0000200541016a200341016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2203200120046a220541046a2802002202411074200741107672360200200341046a200541086a2802002207411074200241107672360200200341086a2005410c6a28020022024110742007411076723602002003410c6a200541106a2802002207411074200241107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022073a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2203200120046a220541046a2802002202411874200741087672360200200341046a200541086a2802002207411874200241087672360200200341086a2005410c6a28020022024118742007410876723602002003410c6a200541106a2802002207411874200241087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bff0201037f200028020445044041000f0b2000102141012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b0b002000410120001b100f0b3501017f230041106b220041c0880436020c41b408200028020c41076a417871220036020041b808200036020041bc083f003602000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f200010312000102d6a0520010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0b5b01027f2000027f0240200128020022054504400c010b200220036a200128020422014b0d0020012002490d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b0b2b01004180080b24696e69740073657400676574007365745f61646472657373006765745f61646472657373";

    public static String BINARY = BINARY_0;

    public static final String FUNC_SET = "set";

    public static final String FUNC_GET = "get";

    public static final String FUNC_SET_ADDRESS = "set_address";

    public static final String FUNC_GET_ADDRESS = "get_address";

    public SimpleContract(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    public SimpleContract(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }

    public static RemoteCall<SimpleContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(SimpleContract.class, web3j, credentials, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<SimpleContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(SimpleContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<SimpleContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(SimpleContract.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public static RemoteCall<SimpleContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(SimpleContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public RemoteCall<TransactionReceipt> set(Uint64 input) {
        final WasmFunction function = new WasmFunction(FUNC_SET, Arrays.asList(input), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> set(Uint64 input, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SET, Arrays.asList(input), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<Uint64> get() {
        final WasmFunction function = new WasmFunction(FUNC_GET, Arrays.asList(), Uint64.class);
        return executeRemoteCall(function, Uint64.class);
    }

    public RemoteCall<TransactionReceipt> set_address(WasmAddress addr) {
        final WasmFunction function = new WasmFunction(FUNC_SET_ADDRESS, Arrays.asList(addr), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> set_address(WasmAddress addr, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SET_ADDRESS, Arrays.asList(addr), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<WasmAddress> get_address() {
        final WasmFunction function = new WasmFunction(FUNC_GET_ADDRESS, Arrays.asList(), WasmAddress.class);
        return executeRemoteCall(function, WasmAddress.class);
    }

    public static SimpleContract load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        return new SimpleContract(contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    public static SimpleContract load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        return new SimpleContract(contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }
}
