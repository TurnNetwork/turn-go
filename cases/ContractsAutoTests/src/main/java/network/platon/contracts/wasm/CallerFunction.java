package network.platon.contracts.wasm;

import java.math.BigInteger;
import java.util.Arrays;
import org.web3j.abi.WasmFunctionEncoder;
import org.web3j.abi.datatypes.WasmFunction;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
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
 * <p>Generated with platon-web3j version 0.13.1.1.
 */
public class CallerFunction extends WasmContract {
    private static String BINARY_0 = "0x0061736d01000000013d0b60027f7f0060017f0060017f017f60000060037f7f7f0060027f7f017f60047f7f7f7f006000017f60037f7f7f017f60047f7f7f7f017f60017f017e02c7010903656e760c706c61746f6e5f70616e6963000303656e760d706c61746f6e5f63616c6c6572000103656e760d726c705f6c6973745f73697a65000203656e760f706c61746f6e5f726c705f6c697374000403656e760e726c705f62797465735f73697a65000503656e7610706c61746f6e5f726c705f6279746573000403656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000703656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e000003201f030100000402080006040005000104030209000501020a01000500030302020405017001010105030100020608017f0141e088040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300090f5f5f66756e63735f6f6e5f65786974002406696e766f6b6500180ac1371f040010250bd00c01097f230041a0016b220124002001410036023820014200370330200141306a4114100b20012802301001200141186a200141306a100c2001200141306a100c200141c8006a4100360200200142003703400240034020034180086a2202410371044020034103460d02200341016a21030c010b0b200341fc076a21020340200241046a22022802002203417f73200341fffdfb776a7141808182847871450d000b0340200341ff0171450d01200241016a2d00002103200241016a21020c000b000b200141406b41800820024180086b100d4100210220014100360278200142003703704114100e20014114100f21054100210303402004411446044002402002044020012003410520026b74411f713a009001200141f0006a20014190016a10100b410021022001410036028801200142003703800120014180016a200128024420012d0040220341017620034101711b410174410172100b200141406b410172210503402002200128024420012d00402203410176200341017122031b22044f0d0120012802800120026a2001280248200520031b20026a2d000022034105763a0000200128028001200128024420012d0040220441017620044101711b6a20026a41016a2003411f713a0000200241016a21020c000b000b05200420056a2d0000200341087441801e71722103200241086a21020340200241054e0440200120032002417b6a220276411f713a009001200141f0006a20014190016a10100c010b0b200441016a21040c010b0b20012802800120046a41003a000020014190016a20014180016a2001280270200128027410112001280280012202044020012002360284010b20014190016a2001280294012001280290016b41066a100b20012802940120012802900122046b21034101210203402003044020042d000041002002411d764101716b41b3c5d1d0027141002002411c764101716b41dde788ea037141002002411b764101716b41fab384f5017141002002411a764101716b41ed9cc2b20271410020024119764101716b41b2afa9db0371200241057441e0ffffff037173737373737321022003417f6a2103200441016a21040c010b0b410021032001410036026820014200370360200141e0006a4106100b200241017321044119210203402002417b470440200128026020036a2004200276411f713a00002002417b6a2102200341016a21030c010b0b2001280290012202044020012002360294010b20014100360258200142003703500240200128027420012802706b2203450d0020012003100e2202360250200120023602542001200220036a3602582001280274200128027022046b22034101480d00200220042003100f1a2001200128025420036a3602540b20014190016a200141d0006a200128026020012802641011200128025022020440200120023602540b20004200370200200041086a41003602004100210203402002410c470440200020026a4100360200200241046a21020c010b0b2001280248200520012d0040220241017122031b210402402001280244200241017620031b220241016a410a4d0440200020024101743a0000200041016a21030c010b200241116a4170712205100e21032000200236020420002005410172360200200020033602080b2003200420021012200220036a41003a00002000413110130240200028020420002d00002205410176200541017122071b22032003200128029401220220012802900122046b6a2206200320064b1b220641106a417071417f6a410a2006410a4b1b220620002802002209417e71417f6a410a20071b2208460d00027f2006410a4604402009200520071b2107200041016a21054100210820002802080c010b4100200620084d200641016a100e22051b0d0120002d0000220741017104404101210820002802080c010b41012108200041016a0b2109200520092000280204200741fe017141017620074101711b41016a10122008044020002005360208200020033602042000200641016a4101723602000c010b200020034101743a00000b20042103034020022004470440200020032d00004184086a2c00001013200341016a21032002417f6a21020c010b0b2004044020012004360294010b200128026022020440200120023602640b200128027022020440200120023602740b200128023022020440200120023602340b200141a0016a24000bfa0101057f230041206b22022400024020002802042203200028020022046b22052001490440200028020820036b200120056b22044f04400340200341003a00002000200028020441016a22033602042004417f6a22040d000c030b000b2000200110142106200241186a200041086a3602002002410036021441002101200604402006100e21010b200220013602082002200120056a22033602102002200120066a3602142002200336020c0340200341003a00002002200228021041016a22033602102004417f6a22040d000b2000200241086a1015200241086a10160c010b200520014d0d002000200120046a3602040b200241206a24000b7d01037f20004200370000200041106a4100360000200041086a4200370000034020024114470440200020026a41003a0000200241016a21020c010b0b200128020021034100210203400240200241134d0440200220036a22042001280204470d010b0f0b200020026a20042d00003a0000200241016a21020c000b000b5a01027f02402002410a4d0440200020024101743a0000200041016a21030c010b200241106a4170712204100e21032000200236020420002004410172360200200020033602080b2003200120021012200220036a41003a00000b0b002000410120001b10190bfc0801067f03400240200020046a2105200120046a210320022004460d002003410371450d00200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220745044003402006411049450440200020046a2203200120046a2205290200370200200341086a200541086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2205200120046a2204290200370200200441086a2103200541086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002007417f6a220741024b0d00024002400240024002400240200741016b0e020102000b2005200120046a220328020022073a0000200541016a200341016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2203200120046a220541046a2802002202410874200741187672360200200341046a200541086a2802002207410874200241187672360200200341086a2005410c6a28020022024108742007411876723602002003410c6a200541106a2802002207410874200241187672360200200441106a2104200641706a21060c000b000b2005200120046a220328020022073a0000200541016a200341016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2203200120046a220541046a2802002202411074200741107672360200200341046a200541086a2802002207411074200241107672360200200341086a2005410c6a28020022024110742007411076723602002003410c6a200541106a2802002207411074200241107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022073a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2203200120046a220541046a2802002202411874200741087672360200200341046a200541086a2802002207411874200241087672360200200341086a2005410c6a28020022024118742007410876723602002003410c6a200541106a2802002207411874200241087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bb70101047f230041206b220224000240200028020422032000280208490440200320012d00003a00002000200028020441016a3602040c010b2000200320002802006b220441016a10142105200241186a200041086a3602004100210320024100360214200504402005100e21030b20022003360208200320046a220420012d00003a00002002200320056a3602142002200436020c2002200441016a3602102000200241086a1015200241086a10160b200241206a24000bc60301057f230041206b22042400200128020421050240200320026b22064101480d002006200128020820056b4c0440034020022003460d02200520022d00003a00002001200128020441016a2205360204200241016a21020c000b000b2001200520066a200128020022066b10142108200441186a200141086a36020020044100360214200520066b2106200804402008100e21070b200420073602082004200620076a22063602102004200720086a3602142004200636020c200441086a410472210703402002200346450440200620022d00003a00002004200428021041016a2206360210200241016a21020c010b0b200128020020052007101702402001280204220320056b220241004c0440200428021021020c010b2004200428021020052002100f20026a2202360210200128020421030b20012002360204200128020021022001200428020c3602002001280208210520012004280214360208200420033602102004200236020c2004200536021420042002360208200441086a1016200128020421050b20002005360204200141003602042000200128020036020020012802082102200141003602082000200236020820014100360200200441206a24000b100020020440200020012002100f1a0b0bf60101057f027f20002d00002202410171220345044020024101762104410a0c010b2000280204210420002802002202417e71417f6a0b210502400240024020042005460440027f2002410171044020002802080c010b200041016a0b2106416f2103200541e6ffffff074d0440410b20054101742202200541016a220320032002491b220241106a4170712002410b491b21030b2003100e220220062005101220002003410172360200200020023602080c010b2003450d01200028020821020b2000200441016a3602040c010b2000200441017441026a3a0000200041016a21020b200220046a220041003a0001200020013a00000b2e01017f2001200028020820002802006b2200410174220220022001491b41ffffffff07200041ffffffff03491b0b6701017f20002802002000280204200141046a1017200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b2b01027f200028020821012000280204210203402001200247044020002001417f6a22013602080c010b0b0b270020022002280200200120006b22016b2202360200200141014e0440200220002001100f1a0b0bdd0802087f017e230041e0006b220124001025100622001019220210070240200141086a20022000411c101a2204280208450440200441146a2802002100200428021021020c010b200141d0006a2004101b200420012802502001280254101c36020c200141386a2004101b410021002004027f410020012802382203450d001a4100200128023c2206200428020c2205490d001a200620052005417f461b210020030b2202360210200441146a2000360200200441003602080b200141386a200220004114101a2200101d024002402000280204450d002000101d0240200028020022032c0000220241004e044020020d010c020b200241807f460d00200241ff0171220541b7014d0440200028020441014d04401000200028020021030b20032d00010d010c020b200541bf014b0d012000280204200241ff017141ca7e6a22024d04401000200028020021030b200220036a2d0000450d010b2000280204450d0020032d000041c001490d010b10000b2000101e2206200028020422034b04401000200028020421030b20002802002107024002400240200304404100210520072c00002200417f4a0d01027f200041ff0171220541bf014d04404100200041ff017141b801490d011a200541c97e6a0c010b4100200041ff017141f801490d001a200541897e6a0b41016a21050c010b4101210520070d00410021000c010b41002100200520066a20034b0d0020032006490d004100210220032005490d01200520076a2102200320056b20062006417f461b22004109490d0110000c010b410021020b0340200004402000417f6a210020023100002008420886842108200241016a21020c010b0b02400240024002402008500d0041a508101f2008510440200410200c040b41aa08101f2008520d0020041020200141286a100a2001410036024020014200370338200141386a41001021200141cc006a410036020020014200370244410121020240200141d0006a200141286a1022220428020420012d00502200410176200041017122031b220041014d0440200041016b0d040c010b20004138490d02200041016a210203402000450d04200241016a2102200041087621000c000b000b2004280208200441016a20031b2c0000417f4c0d010c020b10000c020b200041016a21020b20012802402002490440200141386a200210210b200141386a200141d0006a200141286a10222200280208200041016a20012d0050220241017122041b22032000280204200241017620041b22001004200128023c22026a1023200320002002200128023822066a10050340024020012802482202200128024422034622050d00200241786a220028020022044504401000200028020021040b20002004417f6a220436020020040d0020012000360248200141386a2002417c6a2802002200200128023c20006b220210026a10232000200128023822066a22002002200010030c010b0b200545044010000b2006200128023c10082003450d00200120033602480b1024200141e0006a24000b9b0101047f230041106b220124002001200036020c2000047f41d008200041086a2202411076220041d0082802006a220336020041cc0841cc08280200220420026a41076a417871220236020002400240200341107420024d044041d008200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20042001410c6a4104100f41086a0541000b2100200141106a240020000b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000102620024f0d002003410471044010000c010b200042003702000b02402003411071450d002000102620024d0d0020034104710440100020000f0b200042003702000b20000b7201047f2001101e220220012802044b044010000b2001102721032000027f0240200128020022054504400c010b200220036a200128020422014b0d0020012003490d00410020012002490d011a200320056a2104200120036b20022002417f461b0c010b41000b360204200020043602000b2701017f230041206b22022400200241086a200020014114101a10262100200241206a240020000b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bff0201037f200028020445044041000f0b2000101d41012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b8b0101047f230041106b22012400024002402000280204450d0020002802002d000041c001490d00200141086a2000101b41012103200128020c2100034020000440200141002001280208220220022000101c22046a20024520002004497222021b3602084100200020046b20021b21002003417f6a21030c010b0b2003450d010b10000b200141106a24000b2f01017f200028020820014904402001101920002802002000280204100f210220002001360208200020023602000b0b4d01017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b200020012802082001280204100d20000b3601017f200028020820014904402001101920002802002000280204100f210220002001360208200020023602000b200020013602040b880101037f41bc08410136020041c0082802002100034020000440034041c40841c4082802002201417f6a2202360200200141014845044041bc084100360200200020024102746a22004184016a280200200041046a28020011010041bc08410136020041c00828020021000c010b0b41c408412036020041c008200028020022003602000c010b0b0b3501017f230041106b220041e0880436020c41c808200028020c41076a417871220036020041cc08200036020041d0083f003602000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f200010272000101e6a0520010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0b0b4201004180080b3b6c61780071707a7279397838676632747664773073336a6e35346b686365366d7561376c00696e6974006765745f706c61746f6e5f63616c6c6572";

    public static String BINARY = BINARY_0;

    public static final String FUNC_GET_PLATON_CALLER = "get_platon_caller";

    protected CallerFunction(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    protected CallerFunction(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }

    public RemoteCall<String> get_platon_caller() {
        final WasmFunction function = new WasmFunction(FUNC_GET_PLATON_CALLER, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static RemoteCall<CallerFunction> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CallerFunction.class, web3j, credentials, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<CallerFunction> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CallerFunction.class, web3j, transactionManager, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<CallerFunction> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CallerFunction.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public static RemoteCall<CallerFunction> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CallerFunction.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public static CallerFunction load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        return new CallerFunction(contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    public static CallerFunction load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        return new CallerFunction(contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }
}
