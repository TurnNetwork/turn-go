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
 * <p>Generated with platon-web3j version 0.9.1.2-SNAPSHOT.
 */
public class MemoryFunction_1 extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001360a60000060017f0060027f7f0060017f017f60027f7f017f60037f7f7f006000017f60037f7f7f017f60047f7f7f7f017f60017f017e025d0403656e760c706c61746f6e5f70616e6963000003656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000603656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e0002031d1c000001010104030405000802030103090204020302020705000000030405017001030305030100020608017f0141d08a040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300040f5f5f66756e63735f6f6e5f65786974001c06696e766f6b65000d0908010041010b0206060adc2b1c08001005101d101e0b3801017f419808420037020041a008410036020041742100034020000440200041a4086a4100360200200041046a21000c010b0b410110070b0300010b970101027f41a408410136020041a808280200220145044041a80841b00836020041b00821010b024041ac0828020022024120460440418402100a2201450d012001418402100b220141a80828020036020041a808200136020041ac084100360200410021020b41ac08200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41a40841003602000b950201087f230041106b220424000240024020044180081009220128020420042d0000220241017620024101711b41016a100a200128020420042d0000220241017620024101711b41016a100b22022001280208200141016a20042d00004101711b220173410371044020022103200121050c010b20022106034020014103710440200620012d000022033a00002003450d03200641016a2106200141016a21010c010b0b0340200620076a2103200120076a22052802002208417f73200841fffdfb776a71418081828478710d0120032008360200200741046a21070c000b000b0340200320052d000022013a00002001450d01200341016a2103200541016a21050c000b000b2000200210091a200441106a24000b910101027f20004200370200200041086a410036020020012102024003402002410371044020022d0000450d02200241016a21020c010b0b2002417c6a21020340200241046a22022802002203417f73200341fffdfb776a7141808182847871450d000b0340200341ff0171450d01200241016a2d00002103200241016a21020c000b000b20002001200220016b100c20000b970101047f230041106b220124002001200036020c2000047f41c80a200041086a2202411076220041c80a2802006a220336020041c40a200241c40a28020022026a41076a417871220436020002400240200341107420044d044041c80a200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104101a41086a0541000b200141106a24000be10201027f02402001450d00200041003a0000200020016a2202417f6a41003a000020014103490d00200041003a0002200041003a00012002417d6a41003a00002002417e6a41003a000020014107490d00200041003a00032002417c6a41003a000020014109490d002000410020006b41037122036a220241003602002002200120036b417c7122036a2201417c6a410036020020034109490d002002410036020820024100360204200141786a4100360200200141746a410036020020034119490d002002410036021820024100360214200241003602102002410036020c200141706a41003602002001416c6a4100360200200141686a4100360200200141646a41003602002003200241047141187222036b2101200220036a2102034020014120490d0120024200370300200241186a4200370300200241106a4200370300200241086a4200370300200241206a2102200141606a21010c000b000b20000b6601027f027f02402002410b4f0440200241106a4170712204410120041b100a21032000200236020420002004410172360200200020033602080c010b200020024101743a0000200041016a22032002450d011a0b200320012002101a0b20026a41003a00000b8e0b02077f017e230041e0006b22012400100410012200100a220210020240200141086a20022000411c100e2204280208450440200441146a2802002103200428021021020c010b200141386a2004100f2004200141c8006a2001280238200128023c4114100e101036020c200141c8006a2004100f2004027f410020012802482200450d001a4100200128024c2206200428020c2205490d001a200620052005417f461b210320000b2202360210200441146a2003360200200441003602080b200141c8006a200220034114100e22001011024002402000280204450d00200010110240200028020022022c0000220341004e044020030d010c020b200341807f460d00200341ff0171220441b7014d0440200028020441014d04401000200028020021020b20022d00010d010c020b200441bf014b0d012000280204200341ff017141ca7e6a22034d04401000200028020021020b200220036a2d0000450d010b2000280204450d0020022d000041c001490d010b10000b200010122205200028020422024b04401000200028020421020b20002802002106024002400240200204404100210420062c00002200417f4a0d01027f200041ff0171220441bf014d04404100200041ff017141b801490d011a200441c97e6a0c010b4100200041ff017141f801490d001a200441897e6a0b41016a21040c010b4101210420060d00410021000c010b410021002002200549200420056a20024b720d004100210320022004490d01200420066a2103200220046b20052005417f461b22004109490d0110000c010b410021030b0340200004402000417f6a210020033100002007420886842107200341016a21030c010b0b02400240024002402007500d0041890810132007510d03418e0810132007520d00200141286a10082001410036025020014200370348200141c8006a41001014200141dc006a410036020020014200370254410121030240200141386a200141286a1015220228020420012d00382200410176200041017122041b220041014d0440200041016b0d040c010b20004138490d02200041016a210303402000450d04200341016a2103200041087621000c000b000b2002280208200241016a20041b2c0000417f4c0d010c020b10000c020b200041016a21030b20012802502003490440200141c8006a200310140b41012100200141386a200141286a10152202280208200241016a20012d0038220341017122051b210402400240024002402002280204200341017620051b2202410146044020042c000022024100480d01200141c8006a200241ff017110160c040b200241374b0d01200221000b200141c8006a20004180017341ff017110160c010b20021017220041b7016a22034180024e044010000b200141c8006a200341ff01711016200141c8006a200128024c20006a1018200128024c20012802486a417f6a210320022100037f2000047f200320003a0000200041087621002003417f6a21030c010520020b0b21000b200141c8006a20001019200128024c220220012802486a20042000101a1a2001200020026a36024c0b034002402001280258220320012802544622040d00200341786a220028020022024504401000200028020021020b20002002417f6a220236020020020d0020012000360258200141c8006a4101200128024c22022003417c6a28020022036b22001017220441016a20004138491b220520026a101820052003200128024822056a22026a20022000101b200041374d04402002200041406a3a00000c020b200441f7016a220641ff014d0440200220063a00002005200320046a6a210303402000450d03200320003a0000200041087621002003417f6a21030c000b000510000c020b000b0b200445044010000b2001280248200128024c100320012802542200450d00200120003602580b101c200141e0006a24000b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000101020024f0d002003410471044010000c010b200042003702000b02402003411071450d002000101020024d0d0020034104710440100020000f0b200042003702000b20000b7101047f20011012220220012802044b044010000b2001101f21032000027f0240200128020022054504400c010b200220036a200128020422014b2001200349720d00410020012002490d011a200320056a2104200120036b20022002417f461b0c010b41000b360204200020043602000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f2000101f200010126a0541010b0b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bff0201037f200028020445044041000f0b2000101141012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b2f01017f200028020820014904402001100a20002802002000280204101a210220002001360208200020023602000b0b4d01017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b200020012802082001280204100c20000b2500200041011019200028020020002802046a20013a00002000200028020441016a3602040b1e01017f03402000044020004108762100200141016a21010c010b0b20010b0f00200020011014200020013602040b1b00200028020420016a220120002802084b04402000200110140b0bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d0440200020012002101a1a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b880101037f41a408410136020041a8082802002100034020000440034041ac0841ac082802002201417f6a2202360200200141014845044041a4084100360200200020024102746a22004184016a280200200041046a28020011010041a408410136020041a80828020021000c010b0b41ac08412036020041a808200028020022003602000c010b0b0b3501017f230041106b220041d08a0436020c41c00a200028020c41076a417871220036020041c40a200036020041c80a3f003602000b3801017f41b40a420037020041bc0a410036020041742100034020000440200041c00a6a4100360200200041046a21000c010b0b410210070b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0b0b1e01004180080b175761736d5465737400696e6974006765746d616c6c6f63";

    public static String BINARY = BINARY_0;

    public static final String FUNC_GETMALLOC = "getmalloc";

    protected MemoryFunction_1(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected MemoryFunction_1(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<String> getmalloc() {
        final WasmFunction function = new WasmFunction(FUNC_GETMALLOC, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static RemoteCall<MemoryFunction_1> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(MemoryFunction_1.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<MemoryFunction_1> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(MemoryFunction_1.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<MemoryFunction_1> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(MemoryFunction_1.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<MemoryFunction_1> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(MemoryFunction_1.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static MemoryFunction_1 load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new MemoryFunction_1(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static MemoryFunction_1 load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new MemoryFunction_1(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
