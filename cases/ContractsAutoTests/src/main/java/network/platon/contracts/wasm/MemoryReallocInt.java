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
public class MemoryReallocInt extends WasmContract {
    private static String BINARY_0 = "0x0061736d01000000013a0b60017f0060000060027f7f0060017f017f60037f7f7f017f60037f7f7f0060047f7f7f7f0060027f7e006000017f60027f7f017f60017f017e025d0403656e760c706c61746f6e5f70616e6963000103656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000803656e7610706c61746f6e5f6765745f696e707574000003656e760d706c61746f6e5f72657475726e000203212001000a01020001040305010303000100010004030303000609020203020202070405017001030305030100020608017f0141d08a040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300040f5f5f66756e63735f6f6e5f65786974001206696e766f6b6500070908010041010b0205150af62a200800100a100e10140b070041900810110b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b9a0502077f027e230041406a22002400100410012201100f220210022000200136022c2000200236022820002000290328370300200041286a200041086a2000411c1016101e200041286a101a02400240200041286a101f450d00200028022c450d0020002802282d000041c001490d010b10000b200041286a10192205200028022c22034b04401000200028022c21030b200028022821060240024002402003044020062c00002201417f4a0d01027f200141ff0171220241bf014d04404100200141ff017141b801490d011a200241c97e6a0c010b4100200141ff017141f801490d001a200241897e6a0b41016a21040c010b4101210420060d00410021010c010b410021012003200549200420056a20034b720d004100210220032004490d01200420066a2102200320046b20052005417f461b22014109490d0110000c010b410021020b0340200104402001417f6a210120023100002007420886842107200241016a21020c010b0b024002402007500d0041800810062007510d0141850810062007520d004104100f2201410a3602004108100f41e4003602002001101041e40036020041002101200041003602304200210720004200370328200041286a410010082000413c6a4100360200200042003702344101210242c8012108034020072008845045044020074238862008420888842108200141016a2101200742088821070c010b0b024020014138490d002001210203402002450d01200141016a2101200241087621020c000b000b200141016a210220002802302002490440200041286a200210080b200041286a42c80110232000280234200028023847044010000b2000280228200028022c1003200028023422010440200020013602380b200041286a10090c010b10000b1012200041406b24000b3401017f200028020820014904402001100f220220002802002000280204100b1a2000100920002001360208200020023602000b0b080020002802001a0b3801017f41900842003702004198084100360200417421000340200004402000419c086a4100360200200041046a21000c010b0b410110130bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d0440200020012002100b1a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041d08a0436020c41b80a200028020c41076a417871220036020041bc0a200036020041c00a3f003602000b970101047f230041106b220124002001200036020c2000047f41c00a200041086a2202411076220041c00a2802006a220336020041bc0a200241bc0a28020022026a41076a417871220436020002400240200341107420044d044041c00a200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104100b41086a0541000b200141106a24000b4f01017f230041106b22012400027f41002000450d001a2001410036020c2001410c6a200041786a4104100b1a2000200128020c41284f0d001a4128100f2000200128020c100b0b200141106a24000b130020002d0000410171044020002802081a0b0b880101037f419c08410136020041a0082802002100034020000440034041a40841a4082802002201417f6a22023602002001410148450440419c084100360200200020024102746a22004184016a280200200041046a280200110000419c08410136020041a00828020021000c010b0b41a408412036020041a008200028020022003602000c010b0b0b940101027f419c08410136020041a008280200220145044041a00841a80836020041a80821010b024041a40828020022024120460440418402100f2201450d012001100c220141a00828020036020041a008200136020041a4084100360200410021020b41a408200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b419c0841003602000b3801017f41ac0a420037020041b40a410036020041742100034020000440200041b80a6a4100360200200041046a21000c010b0b410210130b070041ac0a10110b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000101720012802044f0d002002410471044010000c010b200042003702000b02402002411071450d002000101720012802044d0d0020024104710440100020000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001018200010196a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000101a41012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a411410161017200241306a24000b2101017f20011019220220012802044b044010000b20002001200110182002101b0bd60202067f017e230041206b220224002001280208220341004b0440200241186a2001101d20012002280218200228021c101c36020c200241106a2001101d410021032001027f410020022802102204450d001a410020022802142206200128020c2205490d001a200620052005417f461b210720040b360210200141146a2007360200200141003602080b200141106a210503400240200341004f0d002001280214450d00200241106a2001101d41002103027f410020022802102206450d001a410020022802142207200128020c2204490d001a200720046b2103200420066a0b21042001200336021420012004360210200241106a2005410020042003101c101b2001200229031022083702102001200128020c2008422088a76a36020c2001200128020841016a22033602080c010b0b2002200529020022083703082002200837030020002002411410161a200241206a24000b980101037f200028020445044041000f0b2000101a200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b0f00200020011021200020013602040b2f01017f200028020820014904402001100f20002802002000280204100b210220002001360208200020023602000b0b3f01027f2000280204220241016a220320002802084b047f20002003102120002802040520020b20002802006a20013a00002000200028020441016a3602040bbd0402067f037e02402001500440200041800110220c010b20014280015a044020012108034020082009845045044020094238862008420888842108200241016a2102200942088821090c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff017110222000200028020420046a1020200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff017110220b2000200028020420026a1020200028020420002802006a417f6a210203402001200a84500d02200220013c0000200a42388620014208888421012002417f6a2102200a420888210a0c000b000b20002001a741ff017110220b0340024020002802102202200028020c460d00200241786a2802004504401000200028021021020b200241786a22052005280200417f6a220336020020030d002000200536021041002103200028020422072002417c6a28020022066b2204210203402002044020024108762102200341016a21030c010b0b200020074101200341016a20044138491b22056a10202005200028020020066a22026a20022004100d200441374d0440200028020020066a200441406a3a00000c020b200341084d0440200028020020066a200341776a3a0000200028020020066a20036a210203402004450d03200220043a0000200441087621042002417f6a21020c000b000510000c020b000b0b0b0b1601004180080b0f696e6974006765747265616c6c6f63";

    private static String BINARY = BINARY_0;

    public static final String FUNC_GETREALLOC = "getrealloc";

    protected MemoryReallocInt(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected MemoryReallocInt(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Integer> getrealloc() {
        final WasmFunction function = new WasmFunction(FUNC_GETREALLOC, Arrays.asList(), Integer.class);
        return executeRemoteCall(function, Integer.class);
    }

    public static RemoteCall<MemoryReallocInt> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(MemoryReallocInt.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<MemoryReallocInt> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(MemoryReallocInt.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static MemoryReallocInt load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new MemoryReallocInt(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static MemoryReallocInt load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new MemoryReallocInt(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
