package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Uint64;
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
 * <p>Generated with platon-web3j version 0.9.1.2-SNAPSHOT.
 */
public class InnerFunction extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001400c60017f017f60017f0060027f7f0060000060017f017e60037f7e7e006000017e60037f7f7f006000017f60027f7f017f60037f7f7f017f60047f7f7f7f017f02bc010803656e760c706c61746f6e5f70616e6963000303656e7610706c61746f6e5f6761735f7072696365000003656e7610706c61746f6e5f74696d657374616d70000603656e7613706c61746f6e5f626c6f636b5f6e756d626572000603656e7610706c61746f6e5f6761735f6c696d6974000603656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000803656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e000203282703000101010007050504040403000b02000100040009000102020201020a0103020007030300050405017001060605030100020608017f0141f08a040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300080f5f5f66756e63735f6f6e5f65786974002706696e766f6b650014090b010041010b050a1112130a0a863927100041bc0810091a4101100b102b102c0b3601017f20004200370200200041086a410036020003402001410c46450440200020016a4100360200200141046a21010c010b0b20000b0300010b940101027f41c808410136020041cc08280200220145044041cc0841d40836020041d40821010b024041d0082802002202412046044041840210152201450d0120011029220141cc0828020036020041cc08200136020041d0084100360200410021020b41d008200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41c80841003602000bae0502097f047e230041306b22052400200541206a2102200541206a10012101034020010440200b420886200a42388884210b2001417f6a21012002310000200a42088684210a200241016a21020c010b0b0240027f2000100922032d00002200410171450440200041017622004128200041284b1b41106a41f001712102410a0c010b200328020422004128200041284b1b41106a41707122022003280200417e712201460d012001417f6a0b210141002002417f6a20014d2002100d22011b0d0020012003280208200341016a20032d0000220641017122041b2003280204200641017620041b41016a100e2003200136020820032000360204200320024101723602000b200341016a2100200541186a21070340200541106a200a200b100f20052005290310220c2007290300220d10102005290300200a7ca74180086a2d00002108024002400240027f20032d00002201410171220445044020014101762102410a0c010b2003280204210220032802002201417e71417f6a0b220620024604402003280208200020014101711b2109416f2104200641e6ffffff074d0440410b20064101742201200641016a220420042001491b220141106a4170712001410b491b21040b2004100d220120092006100e20032004410172360200200320013602080c010b2004450d01200328020821010b2003200241016a3602040c010b2003200241017441026a3a0000200021010b200120026a220141003a0001200120083a0000200a420956200b420052200b50200c210a200d210b1b0d000b0240200328020420032d00002201410176200141017122011b2202450d0020022003280208200020011b22016a417f6a21020340200120024f0d0120012d00002100200120022d00003a0000200220003a00002002417f6a2102200141016a21010c000b000b200541306a24000b0b002000410120001b10150b10002002044020002001200210251a0b0b3701017f230041106b22032400200320012002102e200329030021012000200341086a29030037030820002001370300200341106a24000b7701017e20002001427f7e200242767e7c2001422088220242ffffffff0f7e7c200242f6ffffff0f7e200142ffffffff0f83220142f6ffffff0f7e22024220887c22034220887c200142ffffffff0f7e200342ffffffff0f837c22014220887c3703082000200242ffffffff0f832001422086843703000b040010030b040010040b040010020bd30902077f017e23004180016b220424001008100522001015220210060240200441086a20022000411c10162201280208450440200141146a2802002103200128021021020c010b200441386a200110172001200441e0006a2004280238200428023c41141016101836020c200441e0006a200110172001027f410020042802602200450d001a410020042802642206200128020c2205490d001a200620052005417f461b210320000b2202360210200141146a2003360200200141003602080b200441e0006a200220034114101622001019024002402000280204450d00200010190240200028020022022c0000220341004e044020030d010c020b200341807f460d00200341ff0171220141b7014d0440200028020441014d04401000200028020021020b20022d00010d010c020b200141bf014b0d012000280204200341ff017141ca7e6a22034d04401000200028020021020b200220036a2d0000450d010b2000280204450d0020022d000041c001490d010b10000b2000101a2205200028020422024b04401000200028020421020b20002802002106024002400240200204404100210120062c00002200417f4a0d01027f200041ff0171220141bf014d04404100200041ff017141b801490d011a200141c97e6a0c010b4100200041ff017141f801490d001a200141897e6a0b41016a21010c010b4101210120060d00410021000c010b410021002002200549200120056a20024b720d004100210320022001490d01200120066a2103200220016b20052005417f461b22004109490d0110000c010b410021030b0340200004402000417f6a210020033100002007420886842107200341016a21030c010b0b02400240024002402007500d00418b08101b2007510d03419008101b2007510440200441286a100c200441386a101c2101200441f8006a4100360200200441f0006a4200370300200441e8006a42003703002004420037036041012100024002400240200441d0006a200441286a101d220328020420042d00502202410176200241017122051b220241014d0440200241016b0d032003280208200341016a20051b2c0000417f4c0d010c030b200241374b0d010b200241016a21000c010b2002101e20026a41016a21000b20042000360260200441e0006a410472101f20012000102041012100200441e0006a200441286a101d2202280208200241016a20042d0060220341017122061b2105024002402002280204200341017620061b2202410146044020052c000022024100480d012001200241ff017110210c060b200241374b0d01200221000b200120004180017341ff017110210c030b2002101e220041b7016a22034180024e044010000b2001200341ff017110212001200128020420006a1022200128020420012802006a417f6a210320022100034020000440200320003a0000200041087621002003417f6a21030c0105200221000c040b000b000b419a08101b2007510440410210230c040b41a708101b2007510440410310230c040b41b108101b2007520d00410410230c030b10000c020b200120001024200128020020012802046a2005200010251a2001200128020420006a3602040b20011026200128020c200141106a28020047044010000b200128020020012802041007200128020c2200450d00200120003602100b102720044180016a24000b970101047f230041106b220124002001200036020c2000047f41ec0a200041086a2202411076220041ec0a2802006a220336020041e80a200241e80a28020022026a41076a417871220436020002400240200341107420044d044041ec0a200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104102541086a0541000b200141106a24000b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000101820024f0d002003410471044010000c010b200042003702000b02402003411071450d002000101820024d0d0020034104710440100020000f0b200042003702000b20000b7101047f2001101a220220012802044b044010000b2001102d21032000027f0240200128020022054504400c010b200220036a200128020422014b2001200349720d00410020012002490d011a200320056a2104200120036b20022002417f461b0c010b41000b360204200020043602000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f2000102d2000101a6a0541010b0b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bff0201037f200028020445044041000f0b2000101941012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b29002000410036020820004200370200200041001028200041146a41003602002000420037020c20000ba10101037f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b20012802082103024020012802042201410a4d0440200020014101743a0000200041016a21020c010b200141106a4170712204100d21022000200136020420002004410172360200200020023602080b200220032001100e200120026a41003a000020000b1e01017f03402000044020004108762100200141016a21010c010b0b20010b860201067f200028020422032000280210220241087641fcffff07716a2101027f200320002802082204460440200041146a210541000c010b2003200028021420026a220541087641fcffff07716a280200200541ff07714102746a2106200041146a21052001280200200241ff07714102746a0b21020340024020022006460440200541003602000340200420036b41027522014103490d022000200341046a22033602040c000b000b200241046a220220012802006b418020470d0120012802042102200141046a21010c010b0b2001417f6a220141014d04402000418004418008200141016b1b3602100b03402003200447044020002004417c6a22043602080c010b0b0b1300200028020820014904402000200110280b0b2500200041011024200028020020002802046a20013a00002000200028020441016a3602040b0f00200020011028200020013602040ba80402047f037e230041406a2203240041012102200320001104002106200341086a101c210141002100200341386a4100360200200341306a4200370300200341286a42003703002003420037032020064280015a044020062107034020052007845045044020054238862007420888842107200041016a2100200542088821050c010b0b200041384f047f2000101e20006a0520000b41016a21020b20032002360220200341206a410472101f20012002102002402006500440200141800110210c010b20064280015a0440420021054100210020062107034020052007845045044020054238862007420888842107200041016a2100200542088821050c010b0b0240200041384f04402000210203402002044020024108762102200441016a21040c010b0b200441c9004f044010000b2001200441b77f6a41ff017110212001200128020420046a1022200128020420012802006a417f6a21042000210203402002450d02200420023a0000200241087621022004417f6a21040c000b000b200120004180017341ff017110210b2001200128020420006a1022200128020420012802006a417f6a21004200210503402005200684500d02200020063c0000200542388620064208888421062000417f6a2100200542088821050c000b000b20012006a741ff017110210b20011026200128020c200141106a28020047044010000b200128020020012802041007200128020c22000440200120003602100b200341406b24000b1b00200028020420016a220120002802084b04402000200110280b0bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bf80101057f0340024020002802102201200028020c460d00200141786a28020041014904401000200028021021010b200141786a2202200228020041016b220436020020040d002000200236021020004101200028020422032001417c6a28020022026b2201101e220441016a20014138491b220520036a1022200220002802006a220320056a20032001102a0240200141374d0440200028020020026a200141406a3a00000c010b200441f7016a220341ff014d0440200028020020026a20033a00002000280200200220046a6a210203402001450d02200220013a0000200141087621012002417f6a21020c000b000b10000b0c010b0b0b880101037f41c808410136020041cc082802002100034020000440034041d00841d0082802002201417f6a2202360200200141014845044041c8084100360200200020024102746a22004184016a280200200041046a28020011010041c808410136020041cc0828020021000c010b0b41d008412036020041cc08200028020022003602000c010b0b0b2f01017f2000280208200149044020011015200028020020002802041025210220002001360208200020023602000b0bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210251a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041f08a0436020c41e40a200028020c41076a417871220036020041e80a200036020041ec0a3f003602000b3801017f41d80a420037020041e00a410036020041742100034020000440200041e40a6a4100360200200041046a21000c010b0b4105100b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bbe0202027f047e2000027e2002500440420021022001420a800c010b0240027e200141fd00200279a76b220341c000460d001a2003413f4d0440200241c00020036bad22078620012003ad2206888421052002200688210820012007862107420021060c020b200241800120036bad2206862001200341406aad22058884210720022005882105200120068621060c010b21072002210541c00021030b03402003044020084201862005423f8884220220054201862007423f88842201427f852205420a7c200554ad2002427f8542007c7c423f8722024200837d20012002420a83220554ad7d2108200120057d210520074201862006423f888421072003417f6a21032004ad20064201868421062002a741017121040c010b0b20074201862006423f888421022004ad2006420186427e83840b370300200020023703080b0b4101004180080b3a3031323334353637383900696e6974006761735f707269636500626c6f636b5f6e756d626572006761735f6c696d69740074696d657374616d70";

    public static String BINARY = BINARY_0;

    public static final String FUNC_BLOCK_NUMBER = "block_number";

    public static final String FUNC_GAS_PRICE = "gas_price";

    public static final String FUNC_GAS_LIMIT = "gas_limit";

    public static final String FUNC_TIMESTAMP = "timestamp";

    protected InnerFunction(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected InnerFunction(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Uint64> block_number() {
        final WasmFunction function = new WasmFunction(FUNC_BLOCK_NUMBER, Arrays.asList(), Uint64.class);
        return executeRemoteCall(function, Uint64.class);
    }

    public RemoteCall<String> gas_price() {
        final WasmFunction function = new WasmFunction(FUNC_GAS_PRICE, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static RemoteCall<InnerFunction> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(InnerFunction.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<InnerFunction> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(InnerFunction.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<InnerFunction> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(InnerFunction.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<InnerFunction> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(InnerFunction.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<Uint64> gas_limit() {
        final WasmFunction function = new WasmFunction(FUNC_GAS_LIMIT, Arrays.asList(), Uint64.class);
        return executeRemoteCall(function, Uint64.class);
    }

    public RemoteCall<Uint64> timestamp() {
        final WasmFunction function = new WasmFunction(FUNC_TIMESTAMP, Arrays.asList(), Uint64.class);
        return executeRemoteCall(function, Uint64.class);
    }

    public static InnerFunction load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new InnerFunction(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static InnerFunction load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new InnerFunction(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
