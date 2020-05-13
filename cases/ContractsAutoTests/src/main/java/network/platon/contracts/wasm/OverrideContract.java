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
public class OverrideContract extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001400c60017f0060017f017f60000060027f7f0060037f7f7f0060037f7f7f017f60017f017e60047f7f7f7f0060027f7e006000017f60027f7f017f60027f7e017f025d0403656e760c706c61746f6e5f70616e6963000203656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000903656e7610706c61746f6e5f6765745f696e707574000003656e760d706c61746f6e5f72657475726e000303232202000b06020603000102050104020100020002000501010100070a030401030303080405017001040405030100020608017f0141d08a040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300040f5f5f66756e63735f6f6e5f65786974001406696e766f6b6500080909010041010b03050c170a802b220800100d101110160b0700419c0810130b2d002001427f7c2201420156044041000f0b2001a741016b0440200020002802002802001101000f0b4190ce000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010ba00302047f027e230041d0006b220024001004100122011012220210022000200136023c2000200236023820002000290338370308200041386a200041106a200041086a411c101822024100102002400240200041386a10092204500d0041800810072004510d0141850810072004520d0041012101200041386a200241011020200041386a10092104200042818080801037022c2000419808360228200041286a20041006210341002102200041003602404200210420004200370338200041386a4100100a200041cc006a41003602002000420037024420034180014f04402003ad2105034020042005845045044020044238862005420888842105200241016a2102200442088821040c010b0b024020024138490d002002210103402001450d01200241016a2102200141087621010c000b000b200241016a21010b20002802402001490440200041386a2001100a0b200041386a2003ad10252000280244200028024847044010000b2000280238200028023c1003200028024422010440200020013602480b200041386a100b0c010b10000b1014200041d0006a24000b9e0202057f017e2000101c0240024020001021450d002000280204450d0020002802002d000041c001490d010b10000b2000101b2204200028020422024b04401000200028020421020b200028020021050240024002402002044020052c00002200417f4a0d01027f200041ff0171220141bf014d04404100200041ff017141b801490d011a200141c97e6a0c010b4100200041ff017141f801490d001a200141897e6a0b41016a21010c010b4101210120050d00410021000c010b410021002002200449200120046a20024b720d0020022001490d01200120056a2103200220016b20042004417f461b22004109490d0110000c010b0b0340200004402000417f6a210020033100002006420886842106200341016a21030c010b0b20060b3401017f2000280208200149044020011012220220002802002000280204100e1a2000100b20002001360208200020023602000b0b080020002802001a0b050041e4000b3801017f419c08420037020041a408410036020041742100034020000440200041a8086a4100360200200041046a21000c010b0b410110150bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d0440200020012002100e1a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041d08a0436020c41c40a200028020c41076a417871220036020041c80a200036020041cc0a3f003602000b970101047f230041106b220124002001200036020c2000047f41cc0a200041086a2202411076220041cc0a2802006a220336020041c80a200241c80a28020022026a41076a417871220436020002400240200341107420044d044041cc0a200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104100e41086a0541000b200141106a24000b130020002d0000410171044020002802081a0b0b880101037f41a808410136020041ac082802002100034020000440034041b00841b0082802002201417f6a2202360200200141014845044041a8084100360200200020024102746a22004184016a280200200041046a28020011000041a808410136020041ac0828020021000c010b0b41b008412036020041ac08200028020022003602000c010b0b0b940101027f41a808410136020041ac08280200220145044041ac0841b40836020041b40821010b024041b0082802002202412046044041840210122201450d012001100f220141ac0828020036020041ac08200136020041b0084100360200410021020b41b008200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41a80841003602000b3801017f41b80a420037020041c00a410036020041742100034020000440200041c40a6a4100360200200041046a21000c010b0b410310150b070041b80a10130b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000101920012802044f0d002002410471044010000c010b200042003702000b02402002411071450d002000101920012802044d0d0020024104710440100020000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f2000101a2000101b6a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000101c41012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a411410181019200241306a24000b2101017f2001101b220220012802044b044010000b200020012001101a2002101d0bd60202077f017e230041206b220324002001280208220420024b0440200341186a2001101f20012003280218200328021c101e36020c200341106a2001101f410021042001027f410020032802102206450d001a410020032802142208200128020c2207490d001a200820072007417f461b210520060b360210200141146a2005360200200141003602080b200141106a210903400240200420024f0d002001280214450d00200341106a2001101f41002104027f410020032802102207450d001a410020032802142208200128020c2206490d001a200820066b2104200620076a0b21052001200436021420012005360210200341106a2009410020052004101e101d20012003290310220a3702102001200128020c200a422088a76a36020c2001200128020841016a22043602080c010b0b20032009290200220a3703082003200a37030020002003411410181a200341206a24000b980101037f200028020445044041000f0b2000101c200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b0f00200020011023200020013602040b2f01017f200028020820014904402001101220002802002000280204100e210220002001360208200020023602000b0b3f01027f2000280204220241016a220320002802084b047f20002003102320002802040520020b20002802006a20013a00002000200028020441016a3602040bbd0402067f037e02402001500440200041800110240c010b20014280015a044020012108034020082009845045044020094238862008420888842108200241016a2102200942088821090c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff017110242000200028020420046a1022200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff017110240b2000200028020420026a1022200028020420002802006a417f6a210203402001200a84500d02200220013c0000200a42388620014208888421012002417f6a2102200a420888210a0c000b000b20002001a741ff017110240b0340024020002802102202200028020c460d00200241786a2802004504401000200028021021020b200241786a22052005280200417f6a220336020020030d002000200536021041002103200028020422072002417c6a28020022066b2204210203402002044020024108762102200341016a21030c010b0b200020074101200341016a20044138491b22056a10222005200028020020066a22026a200220041010200441374d0440200028020020066a200441406a3a00000c020b200341084d0440200028020020066a200341776a3a0000200028020020066a20036a210203402004450d03200220043a0000200441087621042002417f6a21020c000b000510000c020b000b0b0b0b1a02004180080b0c696e69740067657441726561004198080b0102";

    private static String BINARY = BINARY_0;

    public static final String FUNC_GETAREA = "getArea";

    protected OverrideContract(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected OverrideContract(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Integer> getArea(Long input) {
        final WasmFunction function = new WasmFunction(FUNC_GETAREA, Arrays.asList(input), Integer.class);
        return executeRemoteCall(function, Integer.class);
    }

    public static RemoteCall<OverrideContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(OverrideContract.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<OverrideContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(OverrideContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static OverrideContract load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new OverrideContract(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static OverrideContract load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new OverrideContract(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
