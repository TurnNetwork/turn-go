package network.platon.contracts.wasm;

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
 * <p>Generated with platon-web3j version 0.9.1.2-SNAPSHOT.
 */
public class OOMException extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001510b60047f7f7f7f0060000060017f0060017f017f60027f7f006000017f60037f7f7f017f60097f7f7f7f7f7f7f7f7f017f600a7f7f7f7f7f7f7f7f7f7f017f60097f7f7f7e7f7e7f7f7f017f60017f017e025c0403656e760c706c61746f6e5f70616e6963000103656e760c706c61746f6e5f6465627567000403656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000503656e7610706c61746f6e5f6765745f696e707574000203262501040201030301010200040003070908020101020206030303030004040302030201010a010405017001050505030100020615037f0141b08b040b7f0041b08b040b7f0041a70b0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300040b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974001606696e766f6b650026090a010041010b040d0f18240acd41251c00100741a00a420037020041a80a410036020010154103101710280bf609010d7f4113210e417c21064112210a417d210d4111210c417e210b0340200020046a2102200120046a220341037145200441044672450440200220032d00003a0000200e417f6a210e200641016a2106200a417f6a210a200d41016a210d200c417f6a210c200b41016a210b200441016a21040c010b0b410420046b2105024002400240024020024103712207044020054120490d0320074101460d0120074102460d0220074103470d032002200120046a280200220a3a0000200041016a2107410320046b210b200421020340200b41134f0440200220076a2208200120026a220941046a2802002205411874200a41087672360200200841046a200941086a2802002203411874200541087672360200200841086a2009410c6a28020022054118742003410876723602002008410c6a200941106a280200220a411874200541087672360200200241106a2102200b41706a210b0c010b0b41032006416d2006416d4b1b200e6a4170716b20046b2105200120026a41016a2103200020026a41016a21020c030b200521030340200341104f0440200020046a2207200120046a2202290200370200200741086a200241086a290200370200200441106a2104200341706a21030c010b0b027f2005410871450440200120046a2103200020046a0c010b200020046a2202200120046a2200290200370200200041086a2103200241086a0b21042005410471044020042003280200360200200341046a2103200441046a21040b20054102710440200420032f00003b0000200341026a2103200441026a21040b2005410171450d03200420032d00003a00000f0b2002200120046a2205280200220a3a0000200241016a200541016a2f00003b0000200041036a2107410120046b2106200421020340200641114f0440200220076a2208200120026a220941046a2802002205410874200a41187672360200200841046a200941086a2802002203410874200541187672360200200841086a2009410c6a28020022054108742003411876723602002008410c6a200941106a280200220a410874200541187672360200200241106a2102200641706a21060c010b0b4101200b416f200b416f4b1b200c6a4170716b20046b2105200120026a41036a2103200020026a41036a21020c010b2002200120046a2205280200220c3a0000200241016a200541016a2d00003a0000200041026a2107410220046b2106200421020340200641124f0440200220076a2208200120026a220941046a2802002205411074200c41107672360200200841046a200941086a2802002203411074200541107672360200200841086a2009410c6a28020022054110742003411076723602002008410c6a200941106a280200220c411074200541107672360200200241106a2102200641706a21060c010b0b4102200d416e200d416e4b1b200a6a4170716b20046b2105200120026a41026a2103200020026a41026a21020b20054110710440200220032d00003a00002002200328000136000120022003290005370005200220032f000d3b000d200220032d000f3a000f200341106a2103200241106a21020b2005410871044020022003290000370000200341086a2103200241086a21020b2005410471044020022003280000360000200341046a2103200241046a21020b20054102710440200220032f00003b0000200341026a2103200241026a21020b2005410171450d00200220032d00003a00000b0bc70201027f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122016a22004100360200200041840220016b417c7122026a2201417c6a4100360200024020024109490d002000410036020820004100360204200141786a4100360200200141746a410036020020024119490d002000410036021820004100360214200041003602102000410036020c200141706a41003602002001416c6a4100360200200141686a4100360200200141646a41003602002002200041047141187222026b2101200020026a2100034020014120490d0120004200370300200041186a4200370300200041106a4200370300200041086a4200370300200041206a2100200141606a21010c000b000b0b3501017f230041106b220041b08b0436020c418408200028020c41076a41787122003602004180082000360200418c083f003602000b9c0101047f230041106b220224002002200036020c027f02400240024020000440418c08200041086a22014110762200418c082802006a2203360200418408200141840828020022016a41076a4178712204360200200341107420044d0d0120000d020c030b41000c030b418c08200341016a360200200041016a21000b200040000d0010000b20012002410c6a1005200141086a0b200241106a24000b7801027f20002101024003402001410371044020012d0000450d02200141016a21010c010b0b2001417c6a21010340200141046a22012802002202417f73200241fffdfb776a7141808182847871450d000b0340200241ff0171450d01200141016a2d00002102200141016a21010c000b000b200120006b0b0a0041900841013602000b0a0041900841003602000b3401017f23004190086b220124002001200036020c200141106a2000100e200141106a200141106a1009100120014190086a24000b140020022003490440200120026a20003a00000b0bda16030f7f027e037c41900b2105230041306b22082400200841900b36020c4101410220001b210a2008410f6a210f0340410020066b210702400240034020052d00002202450d0120024125470440200241187441187520002006418008200a1100002008200541016a220536020c2007417f6a2107200641016a21060c010b0b2008200541016a220536020c41002104034020052d000022034118744118752102200341204604402008200541016a220536020c200441087221040c010b200241234604402008200541016a220536020c200441107221040c010b2002412b4604402008200541016a220536020c200441047221040c010b2002412d4604402008200541016a220536020c200441027221040c010b200241304604402008200541016a220536020c200441017221040c010b0b0240200341506a41ff017141094d04402008410c6a1010210c200828020c21050c010b4100210c2003412a470d00200128020021022008200541016a220536020c2004410272200420024100481b210420022002411f7522036a200373210c200141046a21010b41002109024020052d0000412e470d002008200541016a220236020c200441800872210420052d0001220341506a41ff017141094d04402008410c6a10102109200828020c21050c010b2003412a460440200128020021022008200541026a220536020c20024100200241004a1b2109200141046a21010c010b200221050b02400240024020052c000041987f6a411f77220241094b0d000240024002400240200241016b0e09020304040400040400010b2008200541016a220536020c20044180027221040c030b2008200541016a220336020c20052d0001220241e800470d032008200541026a220536020c200441c0017221040c020b2008200541016a220536020c20044180047221040c010b2008200541016a220336020c20052d0001220241ec00470d022008200541026a220536020c20044180067221040b20052d000021020c030b2004418001722104200321050c020b2004418002722104200321050c010b41002000200641ff072006418008491b418008200a110000200841306a24000f0b024002400240024002400240024002400240024002400240027f02400240024002400240024020024118744118752203419e7f6a220b41164d0440200b41016b0e15031101020101110101010101110401010501110101110b20034125460d07200341c600460d01200341d800460d100b200320002006418008200a1100000c070b200941062004418008711b220941037441c00a6a2103200141076a417871220e2b030021134100210203402009410a492002411f4b72450440200841106a20026a41303a0000200341786a21032009417f6a2109200241016a21020c010b0b4400000000000000002013a12013201344000000000000000063220d1b22139944000000000000e041630d034180808080780c040b4101210220044102712203450d06200621070c070b200a20002006200128020041004110200941082004412172101121062008200541016a220536020c200141046a21010c0f0b20012802002207417f6a21020340200241016a22022d00000d000b200220076b2202200920022009491b2002200441800871220b410a761b21032004410271220d450d06200621020c070b2013aa0b2101027f410020132001b7a120032b03002215a2221444000000000000f0416320144400000000000000006671450d001a2014ab0b210b02402014200bb8a1221444000000000000e03f644101734504402015200b41016a220bb8654101730d01200141016a21014100210b0c010b201444000000000000e03f620d00200b45200b41017172200b6a210b0b410021032013440000c0ffffffdf41640d082009450d0603402002411f4d0440200841106a20026a200b200b410a6e2203410a6c6b4130723a0000200241016a21022009417f6a2109200b41094b2003210b0d010b0b03402002411f4b220320094572450440200841106a20026a41303a0000200241016a21022009417f6a21090c010b0b20030d07200841106a20026a412e3a0000200241016a21020c070b412520002006418008200a1100000b2008200541016a220536020c200641016a21060c0a0b0340200220066a417f6a21072002200c490440200241016a2102412020002007418008200a1100000c010b0b200241016a21020b20012c000020002007418008200a110000200741016a21062003450d0603402002200c4f0d07412020002006418008200a110000200641016a2106200241016a21020c000b000b410021040340200420066a2102200320046a220e200c490440412020002002418008200a110000200441016a21040c010b0b200e41016a21030b200141046a21010340024020072d00002206450d00200b04402009450d012009417f6a21090b200641187441187520002002418008200a110000200241016a2102200741016a21070c010b0b200d0440410021070340200220076a2106200320076a200c4f0d07412020002006418008200a110000200741016a21070c000b000b200221060c050b20132001b7a1221344000000000000e03f64410173450440200141016a21010c010b20012001201344000000000000e03f61716a21010b03402002411f4d0440200841106a20026a20012001410a6d2203410a6c6b41306a3a0000200241016a2102200141096a2003210141124b0d010b0b20044103712101034020014101472002411f4b722002200c4f72450440200841106a20026a41303a0000200241016a21020c010b0b20044101712103200441027121010240200c2004410c71410047200d726b20022002200c461b2202411f4b0d000240200d410173450440200841106a20026a412d3a00000c010b20044104714504402004410871450d02200841106a20026a41203a00000c010b200841106a20026a412b3a00000b200241016a21020b027f2001200372450440410021040340200420066a2203200220046a200c4f0d021a412020002003418008200a110000200441016a21040c000b000b20060b220320076a21060340200204402002200f6a2c000020002003418008200a110000200641016a21062002417f6a2102200341016a21030c010b0b2001450d0003402006200c4f0d01412020002003418008200a110000200641016a2106200341016a21030c000b000b200e41086a21012008200541016a220536020c200321060c030b41102103027f0240200241ff0171220741d800462202200741f80046724504400240200741ef00470440200741e200470d01410221030c030b410821030c020b2004416f712104410a21030b2004412072200420021b2204200741e40046200741e90046720d011a0b20044173710b2202417e7120022002418008711b2102027f02400240027f0240024002400240200741e900474100200741e400471b4504402002418004710d012002418002710d02200241c000710d042001280200220741107441107520072002418001711b0c050b2002418004710d022002418002710d05200241c000710d062001280200220741ffff037120072002418001711b0c070b200a20002006200141076a417871220129030022112011423f8722127c2012852011423f88a72003ad2009200c200210122106200141086a21010c080b200a20002006200128020022062006411f7522076a2007732006411f7620032009200c2002101121060c060b200a20002006200141076a417871220129030041002003ad2009200c200210122106200141086a21010c060b20012c00000b2107200a2000200620072007411f7522066a2006732007411f7620032009200c2002101121060c030b200a200020062001280200410020032009200c2002101121060c020b20012d00000b2107200a200020062007410020032009200c2002101121060b200141046a21010b2008200541016a220536020c0c000b000b0300010b4501037f20002802002101034020012d000041506a41ff017141094b4504402000200141016a220336020020012c00002002410a6c6a41506a2102200321010c010b0b20020ba10101057f230041206b2209240020082008416f7120031b210a0240200345044041002108200a418008710d010b200a41207141e1007341f6016a210c410021080340200820096a2003200320056e220d20056c6b220b4130200c200b41187441808080d000481b6a3a0000200841016a2208411f4b0d01200320054f200d21030d000b0b200020012002200920082004200520062007200a1013200941206a24000baa0102047f017e230041206b2209240020082008416f71200342005222081b210a0240200845044041002108200a418008710d010b200a41207141e1007341f6016a210c410021080340200820096a4130200c20032003200580220d20057e7da7220b41187441808080d000481b200b6a3a0000200841016a2208411f4b0d01200320055a200d21030d000b0b2000200120022009200820042005a720062007200a1013200941206a24000bf10401037f2009410271210b2004210a02400340200b0d01200a411f4b200a20074f724504402003200a6a41303a0000200a41016a210a0c010b0b200a21040b2009410371410147210c2004210a02400340200c0d01200a411f4b200a20084f724504402003200a6a41303a0000200a41016a210a0c010b0b200a21040b2009410171210c0240024002400240200941107145044020040d014100210420050d020c030b200445200941800871722004200747410020042008471b724504402004417e6a2004417f6a220420041b200420064110461b21040b0240024020064110460440200941207122062004411f4b720d01200320046a41f8003a0000200441016a21040c020b20064102472004411f4b720d01200320046a41e2003a0000200441016a21040c010b2006452004411f4b720d00200320046a41d8003a0000200441016a21040b2004411f4b0d00200320046a41303a0000200441016a21040b20082009410c714100472005726b200420042008461b2204411f4b0d022005450d010b200320046a412d3a0000200441016a21040c010b20094104714504402009410871450d01200320046a41203a0000200441016a21040c010b200320046a412b3a0000200441016a21040b2002210a0240200b200c720d00200421050340200520084f0d0141202001200a4180082000110000200541016a2105200a41016a210a0c000b000b2003417f6a2103034020040440200320046a2c00002001200a41800820001100002004417f6a2104200a41016a210a0c010b0b0240200b450d00410020026b210203402002200a6a20084f0d0141202001200a4180082000110000200a41016a210a0c000b000b200a0b130020002d0000410171044020002802081a0b0b2301017f03402000410c470440200041a00a6a4100360200200041046a21000c010b0b0b7601037f100a419408280200210003402000044003404198084198082802002201417f6a22023602002001410148450440200020024102746a22004184016a280200200041046a280200100b110200100a41940828020021000c010b0b4198084120360200419408200028020022003602000c010b0b0b900101027f100a4194082802002201450440419408419c08360200419c0821010b024041980828020022024120460440418402100822010440200110060b2001450d01200141940828020036020041940820013602004198084100360200410021020b419808200241016a360200200120024102746a22014184016a4100360200200141046a2000360200100b0f0b100b0b070041a00a10140b780020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000101a20012802044f0d002002410471450440200042003702000c010b10000b024002402002411071450d002000101a20012802044d0d0020024104710d01200042003702000b20000f0b100020000b290002402000280204044020002802002c0000417f4c0d0141010f0b41000f0b2000101b2000101c6a0b240002402000280204450d0020002802002c0000417f4c0d0041000f0b2000102141016a0b8a0301047f0240024020002802040440200010224101210220002802002c00002201417f4c0d010c020b41000f0b200141ff0171220241b7014d0440200241807f6a0f0b02400240200141ff0171220141bf014d04400240200041046a22042802002201200241c97e6a22034d047f100020042802000520010b4102490d0020002802002d00010d0010000b200341054f044010000b20002802002d000145044010000b410021024100210103402001200346450440200028020020016a41016a2d00002002410874722102200141016a21010c010b0b200241384f0d010c020b200141f7014d0440200241c07e6a0f0b0240200041046a22042802002201200241897e6a22034d047f100020042802000520010b4102490d0020002802002d00010d0010000b200341054f044010000b20002802002d000145044010000b410021024100210103402001200346450440200028020020016a41016a2d00002002410874722102200141016a21010c010b0b20024138490d010b200241ff7d490d010b100020020f0b20020b3902017f017e230041306b2201240020012000290200220237031020012002370308200141186a200141086a41141019101a200141306a24000b5e01027f2000027f027f2001280200220504404100200220036a200128020422014b2001200249720d011a410020012003490d021a200220056a2104200120026b20032003417f461b0c020b41000b210441000b360204200020043602000b2101017f2001101c220220012802044b044010000b200020012001101b2002101e0b900302097f017e230041406a220224002001280208220341004b0440200241386a2001101f200220022903383703182001200241186a101d36020c200241306a2001101f410021032001027f410020022802302205450d001a410020022802342207200128020c2204490d001a200720042004417f461b210820050b360210200141146a2008360200200141086a41003602000b200141106a2104200141146a21072001410c6a2105200141086a210803400240200341004f0d002007280200450d00200241306a2001101f41002103027f20022802302209044041002002280234220a20052802002206490d011a200a20066b2103200620096a0c010b41000b210620072003360200200420063602002002200336022c2002200636022820022002290328370310200241306a20044100200241106a101d101e20042002290330220b37020020052005280200200b422088a76a3602002008200828020041016a22033602000c010b0b20022004290200220b3703202002200b3703082000200241086a411410191a200241406b24000b4101017f02402000280204450d0020002802002d0000220041bf014d0440200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b4401017f200028020445044010000b0240200028020022012d0000418101470d00200041046a28020041014d047f100020002802000520010b2c00014100480d0010000b0b9f0101037f02402000280204044020001022200028020022022c000022014100480d0120014100470f0b41000f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200041046a28020041014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a200041046a280200200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b070041ac0a10140b3401027f230041106b2201240003402000418030470440200120003602002001100c200041016a21000c010b0b200141106a24000b950302077f017e230041406a22002400100410022202100822011003200020023602142000200136021020002000290310370308200041106a200041286a200041086a411c10191020200041106a102202400240200041106a1023450d002000280214450d0020002802102d000041c001490d010b10000b200041106a101c2204200028021422024b04401000200028021421020b2000280210210602400240027f027f024020020440410020062c00002201417f4a0d031a200141ff0171220341bf014b0d014100200141ff017141b801490d021a200341c97e6a0c020b410120060d021a0c030b4100200141ff017141f801490d001a200341897e6a0b41016a0b21032002200449200320046a20024b720d004100210120022003490d01200320066a2101200220036b20042004417f461b22054109490d0110000c010b410021010b0340200504402005417f6a210520013100002007420886842107200141016a21010c010b0b024002402007500d0041950b10272007510d01419a0b10272007520d0010250c010b10000b200041406b24000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b3801017f41ac0a420037020041b40a410036020041742100034020000440200041b80a6a4100360200200041046a21000c010b0b410410170b0b67010041c60a0b60f03f000000000000244000000000000059400000000000408f40000000000088c34000000000006af8400000000080842e4100000000d01263410000000084d797410000000065cdcd412564090a00696e6974006d656d6f72795f6c696d6974";

    public static String BINARY = BINARY_0;

    public static final String FUNC_MEMORY_LIMIT = "memory_limit";

    protected OOMException(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected OOMException(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<TransactionReceipt> memory_limit() {
        final WasmFunction function = new WasmFunction(FUNC_MEMORY_LIMIT, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> memory_limit(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_MEMORY_LIMIT, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static RemoteCall<OOMException> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(OOMException.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<OOMException> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(OOMException.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<OOMException> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(OOMException.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<OOMException> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(OOMException.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static OOMException load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new OOMException(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static OOMException load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new OOMException(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
