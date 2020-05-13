package network.platon.contracts.wasm;

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
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 0.7.5.3-SNAPSHOT.
 */
public class MemoryFunction_2 extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001350a60027f7f0060017f0060000060017f017f60027f7f017f60037f7f7f0060037f7f7f017f60047f7f7f7f006000017f60017f017e025d0403656e760c706c61746f6e5f70616e6963000203656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000803656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e0000032928020104010902000102030604050203040405010002010201060303030107040000030300000000000405017001030305030100020608017f0141d08a040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300040f5f5f66756e63735f6f6e5f65786974001806696e766f6b6500090908010041010b02051b0ab830280800100c1011101a0b0700419c0810160b1f0020004200370200200041086a4100360200200020012001100d101520000b8e0101047f230041106b2201240020014180081006220328020420012d0000220241017620024101711b41016a1012200328020420012d0000220241017620024101711b41016a10132202200328020420012d0000220441017620044101711b41016a100f1a20022003280208200341016a20012d00004101711b10172000200210061a20031016200141106a24000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010bf90502077f017e230041f0006b22002400100410012201101222021002200020013602442000200236024020002000290340370308200041406b200041106a200041086a411c101c1024200041406b102002400240200041406b1025450d002000280244450d0020002802402d000041c001490d010b10000b200041406b101f2205200028024422034b04401000200028024421030b200028024021060240024002402003044020062c00002201417f4a0d01027f200141ff0171220241bf014d04404100200141ff017141b801490d011a200241c97e6a0c010b4100200141ff017141f801490d001a200241897e6a0b41016a21040c010b4101210420060d00410021010c010b410021012003200549200420056a20034b720d004100210220032004490d01200420066a2102200320046b20052005417f461b22014109490d0110000c010b410021020b0340200104402001417f6a210120023100002007420886842107200241016a21020c010b0b02400240024002402007500d00418a0810082007510d03418f0810082007520d00200041306a10072000410036024820004200370340200041406b4100100a200041d4006a41003602002000420037024c410121020240200041d8006a200041306a1014220328020420002d00582201410176200141017122041b220141014d0440200141016b0d040c010b20014138490d02200141016a210203402001450d04200241016a2102200141087621010c000b000b2003280208200341016a20041b2c0000417f4c0d010c020b10000c020b200141016a21020b2003101620002802482002490440200041406b2002100a0b2000200041d8006a200041306a10142201280208200141016a20002d0058220241017122031b36026820002001280204200241017620031b36026c20002000290368370300200041406b2000102b20011016200028024c200028025047044010000b200028024020002802441003200028024c22010440200020013602500b200041406b100b200041306a10160b1018200041f0006a24000b3401017f2000280208200149044020011012220220002802002000280204100e1a2000100b20002001360208200020023602000b0b080020002802001a0b3801017f419c08420037020041a408410036020041742100034020000440200041a8086a4100360200200041046a21000c010b0b410110190b7801027f20002101024003402001410371044020012d0000450d02200141016a21010c010b0b2001417c6a21010340200141046a22012802002202417f73200241fffdfb776a7141808182847871450d000b0340200241ff0171450d01200141016a2d00002102200141016a21010c000b000b200120006b0bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000be10201027f02402001450d00200041003a0000200020016a2202417f6a41003a000020014103490d00200041003a0002200041003a00012002417d6a41003a00002002417e6a41003a000020014107490d00200041003a00032002417c6a41003a000020014109490d002000410020006b41037122036a220241003602002002200120036b417c7122036a2201417c6a410036020020034109490d002002410036020820024100360204200141786a4100360200200141746a410036020020034119490d002002410036021820024100360214200241003602102002410036020c200141706a41003602002001416c6a4100360200200141686a4100360200200141646a41003602002003200241047141187222036b2101200220036a2102034020014120490d0120024200370300200241186a4200370300200241106a4200370300200241086a4200370300200241206a2102200141606a21010c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d0440200020012002100e1a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041d08a0436020c41c40a200028020c41076a417871220036020041c80a200036020041cc0a3f003602000b970101047f230041106b220124002001200036020c2000047f41cc0a200041086a2202411076220041cc0a2802006a220336020041c80a200241c80a28020022026a41076a417871220436020002400240200341107420044d044041cc0a200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104100e41086a0541000b200141106a24000b5301017f230041106b22022400027f4100200045200145720d001a2002410036020c2002410c6a200041786a4104100e1a2000200228020c20014f0d001a200110122000200228020c100e0b200241106a24000b4d01017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b200020012802082001280204101520000b6601027f027f02402002410b4f0440200241106a4170712204410120041b101221032000200236020420002004410172360200200020033602080c010b200020024101743a0000200041016a22032002450d011a0b200320012002100e0b20026a41003a00000b130020002d0000410171044020002802081a0b0bac0101047f0240024020002001734103710440200121020c010b20002103034020014103710440200320012d000022003a00002000450d03200341016a2103200141016a21010c010b0b0340200320046a2100200120046a22022802002205417f73200541fffdfb776a71418081828478710d0120002005360200200441046a21040c000b000b0340200020022d000022013a00002001450d01200041016a2100200241016a21020c000b000b0b880101037f41a808410136020041ac082802002100034020000440034041b00841b0082802002201417f6a2202360200200141014845044041a8084100360200200020024102746a22004184016a280200200041046a28020011010041a808410136020041ac0828020021000c010b0b41b008412036020041ac08200028020022003602000c010b0b0b970101027f41a808410136020041ac08280200220145044041ac0841b40836020041b40821010b024041b0082802002202412046044041840210122201450d012001418402100f220141ac0828020036020041ac08200136020041b0084100360200410021020b41b008200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41a80841003602000b3801017f41b80a420037020041c00a410036020041742100034020000440200041c40a6a4100360200200041046a21000c010b0b410210190b070041b80a10160b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000101d20012802044f0d002002410471044010000c010b200042003702000b02402002411071450d002000101d20012802044d0d0020024104710440100020000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f2000101e2000101f6a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000102041012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a4114101c101d200241306a24000b2101017f2001101f220220012802044b044010000b200020012001101e200210210bd60202067f017e230041206b220224002001280208220341004b0440200241186a2001102320012002280218200228021c102236020c200241106a20011023410021032001027f410020022802102204450d001a410020022802142206200128020c2205490d001a200620052005417f461b210720040b360210200141146a2007360200200141003602080b200141106a210503400240200341004f0d002001280214450d00200241106a2001102341002103027f410020022802102206450d001a410020022802142207200128020c2204490d001a200720046b2103200420066a0b21042001200336021420012004360210200241106a2005410020042003102210212001200229031022083702102001200128020c2008422088a76a36020c2001200128020841016a22033602080c010b0b20022005290200220837030820022008370300200020024114101c1a200241206a24000b980101037f200028020445044041000f0b20001020200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b0f00200020011028200020013602040b2f01017f200028020820014904402001101220002802002000280204100e210220002001360208200020023602000b0b1b00200028020420016a220120002802084b04402000200110280b0b2500200041011029200028020020002802046a20013a00002000200028020441016a3602040bd40301047f2001280200210441012102024002400240024020012802042201410146044020042c000022014100480d012000200141ff0171102a0c040b200141374b0d01200121020b200020024180017341ff0171102a0c010b20011026220241b7016a22034180024e044010000b2000200341ff0171102a2000200028020420026a1027200028020420002802006a417f6a210320012102037f2002047f200320023a0000200241087621022003417f6a21030c010520010b0b21020b200020021029200028020020002802046a20042002100e1a2000200028020420026a3602040b0340024020002802102201200028020c460d00200141786a2802004504401000200028021021010b200141786a22022002280200417f6a220336020020030d002000200236021020004101200028020422042001417c6a28020022026b22011026220341016a20014138491b220520046a1027200220002802006a220420056a200420011010200141374d0440200028020020026a200141406a3a00000c020b200341f7016a220441ff014d0440200028020020026a20043a00002000280200200220036a6a210203402001450d03200220013a0000200141087621012002417f6a21020c000b000510000c020b000b0b0b0b2001004180080b195761736d546573743200696e6974006765747265616c6c6f63";

    private static String BINARY = BINARY_0;

    public static final String FUNC_GETREALLOC = "getrealloc";

    protected MemoryFunction_2(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected MemoryFunction_2(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<String> getrealloc() {
        final WasmFunction function = new WasmFunction(FUNC_GETREALLOC, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static RemoteCall<MemoryFunction_2> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(MemoryFunction_2.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<MemoryFunction_2> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(MemoryFunction_2.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static MemoryFunction_2 load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new MemoryFunction_2(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static MemoryFunction_2 load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new MemoryFunction_2(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
