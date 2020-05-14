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
public class CallerFunction extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001350a60027f7f0060017f0060017f017f60000060037f7f7f0060027f7f017f60047f7f7f7f006000017f60037f7f7f017f60017f017e02710503656e760c706c61746f6e5f70616e6963000303656e760d706c61746f6e5f63616c6c6572000103656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000703656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e00000334330301010000000104010500010903000104070802040302020504040100000603010301080202020106050000020200000000000405017001030305030100020608017f0141f08a040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300050f5f5f66756e63735f6f6e5f65786974002406696e766f6b6500120908010041010b0206270a8245331f0041bc08420037020041c408410036020041bc08100d41011025101a10260b070041bc0810200b9a0901067f230041a0016b220124002001410036023820014200370330200141306a4114100820012802301001200141186a200141306a10092001200141306a1009200141c8006a410036020020014200370340200141406b41a1081016101e200141003602782001420037037020014114101c2205360290012001200541146a22063602980120052001411410171a200120063602940103402004200620056b4f044002402003044020012002410520036b74411f713a008001200141f0006a20014180016a100a0b20014190016a100b410021032001410036028801200142003703800120014180016a200128024420012d0040220241017620024101711b4101744101721008200141406b410172210503402003200128024420012d00402202410176200241017122021b22044f0d0120012802800120036a2001280248200520021b20036a2d000022024105763a0000200128028001200128024420012d0040220441017620044101711b6a20036a41016a2002411f713a0000200341016a21030c000b000b05200420056a2d0000200241087441801e71722102200341086a21030340200341054e0440200120022003417b6a220376411f713a008001200141f0006a20014180016a100a0c010b0b200441016a2104200128029001210520012802940121060c010b0b20012802800120046a41003a000020014190016a20014180016a200141f0006a100c20014180016a100b20014190016a2001280294012001280290016b41066a100820012802940120012802900122046b21024101210303402002044020042d000041002003411d764101716b41b3c5d1d0027141002003411c764101716b41dde788ea037141002003411b764101716b41fab384f5017141002003411a764101716b41ed9cc2b20271410020034119764101716b41b2afa9db0371200341057441e0ffffff037173737373737321032002417f6a2102200441016a21040c010b0b410021022001410036026820014200370360200141e0006a41061008200341017321044119210303402003417b470440200128026020026a2004200376411f713a00002003417b6a2103200241016a21020c010b0b20014190016a100b41002103200141003602582001420037035002400240200128027420012802706b2202450d002002417f4c0d0120012002101c2204360250200120043602542001200220046a3602582001280274200128027022066b22024101480d0020042006200210171a2001200128025420026a3602540b20014190016a200141d0006a200141e0006a100c200141d0006a100b200041086a4100360200200042003702002000100d20002001280248200520012d0040220241017122041b2001280244200241017620041b2202200241016a10232000413110222000200128029401220420012802900122026b200028020420002d0000220541017620054101711b6a102103402003200420026b4904402000200220036a2d00004180086a2c00001022200341016a2103200128029001210220012802940121040c010b0b20014190016a100b200141e0006a100b200141f0006a100b200141406b1020200141306a100b200141a0016a24000f0b000b870201047f230041206b22022400024020002802042203200028020022056b22042001490440200028020820036b200120046b22044f04400340200341003a00002000200028020441016a22033602042004417f6a22040d000c030b000b20002001100e2105200241186a200041086a36020020024100360214200028020420002802006b210341002101200504402005101c21010b200220013602082002200120036a22033602102002200120056a3602142002200336020c0340200341003a00002002200228021041016a22033602102004417f6a22040d000b2000200241086a100f200241086a10100c010b200420014d0d002000200120056a3602040b200241206a24000b6201037f034020024114470440200020026a41003a0000200241016a21020c010b0b200128020021034100210203400240200241134d0440200220036a22042001280204470d010b0f0b200020026a20042d00003a0000200241016a21020c000b000bc20101047f230041206b220224000240200028020422032000280208490440200320012d00003a00002000200028020441016a3602040c010b2000200320002802006b41016a100e2105200241186a200041086a3602004100210320024100360214200028020420002802006b2104200504402005101c21030b20022003360208200320046a220420012d00003a00002002200320056a3602142002200436020c2002200441016a3602102000200241086a100f200241086a10100b200241206a24000b1501017f200028020022010440200020013602040b0bda0301067f230041206b2203240020012802042105024020022802042208200228020022026b22044101480d002004200128020820056b4c0440034020022008460d02200520022d00003a00002001200128020441016a2205360204200241016a21020c000b000b2001200420056a20012802006b100e2107200341186a200141086a3602004100210420034100360214200520012802006b2106200704402007101c21040b200320043602082003200420066a22063602102003200420076a3602142003200636020c200341086a410472210403402002200846450440200620022d00003a00002003200328021041016a2206360210200241016a21020c010b0b200128020020052004101502402001280204220420056b220241004c0440200328021021020c010b200328021022042005200210171a2003200220046a2202360210200128020421040b20012002360204200128020021022001200328020c3602002001280208210520012003280214360208200320043602102003200236020c2003200536021420032002360208200341086a1010200128020421050b20002005360204200141003602042000200128020036020020012802082102200141003602082000200236020820014100360200200341206a24000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0b3701017f2001417f4c0440000b2001200028020820002802006b2200410174220220022001491b41ffffffff07200041ffffffff03491b0b6701017f20002802002000280204200141046a1015200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b3101027f200028020821012000280204210203402001200247044020002001417f6a22013602080c010b0b20002802001a0b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010bf90502077f017e230041f0006b22002400100510022201101b22021003200020013602442000200236024020002000290340370308200041406b200041106a200041086a411c10281030200041406b102c02400240200041406b1031450d002000280244450d0020002802402d000041c001490d010b10000b200041406b102b2205200028024422034b04401000200028024421030b200028024021060240024002402003044020062c00002201417f4a0d01027f200141ff0171220241bf014d04404100200141ff017141b801490d011a200241c97e6a0c010b4100200141ff017141f801490d001a200241897e6a0b41016a21040c010b4101210420060d00410021010c010b410021012003200549200420056a20034b720d004100210220032004490d01200420066a2102200320046b20052005417f461b22014109490d0110000c010b410021020b0340200104402001417f6a210120023100002007420886842107200241016a21020c010b0b02400240024002402007500d0041a50810112007510d0341aa0810112007520d00200041306a10072000410036024820004200370340200041406b41001013200041d4006a41003602002000420037024c410121020240200041d8006a200041306a101d220328020420002d00582201410176200141017122041b220141014d0440200141016b0d040c010b20014138490d02200141016a210203402001450d04200241016a2102200141087621010c000b000b2003280208200341016a20041b2c0000417f4c0d010c020b10000c020b200141016a21020b2003102020002802482002490440200041406b200210130b2000200041d8006a200041306a101d2201280208200141016a20002d0058220241017122031b36026820002001280204200241017620031b36026c20002000290368370300200041406b2000103720011020200028024c200028025047044010000b200028024020002802441004200028024c22010440200020013602500b200041406b1014200041306a10200b1024200041f0006a24000b3401017f200028020820014904402001101b22022000280200200028020410171a2000101420002001360208200020023602000b0b080020002802001a0b270020022002280200200120006b22016b2202360200200141014e044020022000200110171a0b0b7a01027f41a1082100024003402000410371044020002d0000450d02200041016a21000c010b0b2000417c6a21000340200041046a22002802002201417f73200141fffdfb776a7141808182847871450d000b0340200141ff0171450d01200041016a2d00002101200041016a21000c000b000b200041a1086b0bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210171a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041f08a0436020c41e40a200028020c41076a417871220036020041e80a200036020041ec0a3f003602000b970101047f230041106b220124002001200036020c2000047f41ec0a200041086a2202411076220041ec0a2802006a220336020041e80a200241e80a28020022026a41076a417871220436020002400240200341107420044d044041ec0a200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104101741086a0541000b200141106a24000b0b002000410120001b101b0b4d01017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b200020012802082001280204101e20000b5a01027f02402002410a4d0440200020024101743a0000200041016a21030c010b200241106a4170712204101c21032000200236020420002004410172360200200020033602080b200320012002101f200220036a41003a00000b10002002044020002001200210171a0b0b130020002d0000410171044020002802081a0b0bf80101057f0240027f20002d0000220241017104402000280204210320002802002202417e71417f6a0c010b20024101762103410a0b220420032001200320014b1b220141106a417071417f6a410a2001410a4b1b2201460d00027f2001410a460440200041016a210420002802080c010b4100200120044d200141016a101c22041b0d0120002d0000220241017104404101210520002802080c010b41012105200041016a0b210620042006027f2002410171044020002802040c010b200241fe01714101760b41016a101f2005044020002004360208200020033602042000200141016a4101723602000f0b200020034101743a00000b0bf40101057f024002400240027f20002d00002202410171220345044020024101762104410a0c010b2000280204210420002802002202417e71417f6a0b22052004460440027f2002410171044020002802080c010b200041016a0b2106416f2103200541e6ffffff074d0440410b20054101742202200541016a220320032002491b220241106a4170712002410b491b21030b2003101c220220062005101f20002002360208200020034101723602000c010b2003450d01200028020821020b2000200441016a3602040c010b2000200441017441026a3a0000200041016a21020b200220046a220041003a0001200020013a00000b5a01017f02402003410a4d0440200020024101743a0000200041016a21030c010b200341106a4170712204101c21032000200236020420002004410172360200200020033602080b200320012002101f200220036a41003a00000b880101037f41c808410136020041cc082802002100034020000440034041d00841d0082802002201417f6a2202360200200141014845044041c8084100360200200020024102746a22004184016a280200200041046a28020011010041c808410136020041cc0828020021000c010b0b41d008412036020041cc08200028020022003602000c010b0b0b940101027f41c808410136020041cc08280200220145044041cc0841d40836020041d40821010b024041d00828020022024120460440418402101b2201450d0120011018220141cc0828020036020041cc08200136020041d0084100360200410021020b41d008200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41c80841003602000b3801017f41d80a420037020041e00a410036020041742100034020000440200041e40a6a4100360200200041046a21000c010b0b410210250b070041d80a10200b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000102920012802044f0d002002410471044010000c010b200042003702000b02402002411071450d002000102920012802044d0d0020024104710440100020000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f2000102a2000102b6a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000102c41012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a411410281029200241306a24000b2101017f2001102b220220012802044b044010000b200020012001102a2002102d0bd60202067f017e230041206b220224002001280208220341004b0440200241186a2001102f20012002280218200228021c102e36020c200241106a2001102f410021032001027f410020022802102204450d001a410020022802142206200128020c2205490d001a200620052005417f461b210720040b360210200141146a2007360200200141003602080b200141106a210503400240200341004f0d002001280214450d00200241106a2001102f41002103027f410020022802102206450d001a410020022802142207200128020c2204490d001a200720046b2103200420066a0b21042001200336021420012004360210200241106a2005410020042003102e102d2001200229031022083702102001200128020c2008422088a76a36020c2001200128020841016a22033602080c010b0b2002200529020022083703082002200837030020002002411410281a200241206a24000b980101037f200028020445044041000f0b2000102c200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b0f00200020011034200020013602040b2f01017f200028020820014904402001101b200028020020002802041017210220002001360208200020023602000b0b1b00200028020420016a220120002802084b04402000200110340b0b2500200041011035200028020020002802046a20013a00002000200028020441016a3602040bd40301047f2001280200210441012102024002400240024020012802042201410146044020042c000022014100480d012000200141ff017110360c040b200141374b0d01200121020b200020024180017341ff017110360c010b20011032220241b7016a22034180024e044010000b2000200341ff017110362000200028020420026a1033200028020420002802006a417f6a210320012102037f2002047f200320023a0000200241087621022003417f6a21030c010520010b0b21020b200020021035200028020020002802046a2004200210171a2000200028020420026a3602040b0340024020002802102201200028020c460d00200141786a2802004504401000200028021021010b200141786a22022002280200417f6a220336020020030d002000200236021020004101200028020422042001417c6a28020022026b22011032220341016a20014138491b220520046a1033200220002802006a220420056a200420011019200141374d0440200028020020026a200141406a3a00000c020b200341f7016a220441ff014d0440200028020020026a20043a00002000280200200220036a6a210203402001450d03200220013a0000200141087621012002417f6a21020c000b000510000c020b000b0b0b0b4201004180080b3b71707a7279397838676632747664773073336a6e35346b686365366d7561376c006c617800696e6974006765745f706c61746f6e5f63616c6c6572";

    public static String BINARY = BINARY_0;

    public static final String FUNC_GET_PLATON_CALLER = "get_platon_caller";

    protected CallerFunction(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected CallerFunction(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<String> get_platon_caller() {
        final WasmFunction function = new WasmFunction(FUNC_GET_PLATON_CALLER, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static RemoteCall<CallerFunction> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CallerFunction.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<CallerFunction> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CallerFunction.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<CallerFunction> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CallerFunction.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<CallerFunction> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CallerFunction.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static CallerFunction load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new CallerFunction(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static CallerFunction load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new CallerFunction(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
