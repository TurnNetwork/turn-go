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
public class ContractDistoryWithPermissionCheck extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001570f60027f7f0060017f0060017f017f60000060037f7f7f0060037f7f7f017f60047f7f7f7f0060027f7e0060027f7f017f6000017f60077f7f7f7f7f7f7f0060037f7e7f0060047f7f7f7f017f60017e017f60017f017e02d2010903656e760c706c61746f6e5f70616e6963000303656e760d706c61746f6e5f6f726967696e000103656e760e706c61746f6e5f64657374726f79000203656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000903656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e000003656e7617706c61746f6e5f6765745f73746174655f6c656e677468000803656e7610706c61746f6e5f6765745f7374617465000c03656e7610706c61746f6e5f7365745f73746174650006034d4c03050503020209050503030804040100040a040806000102030101050202020206000402010204000002000000000400070d0b0101010000020305000e0201020701000701000200010001020405017001030305030100020615037f0141808b040b7f0041808b040b7f0041fd0a0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300090b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974002106696e766f6b6500420908010041010b02233c0aba524c1800100c41a40a10201a4101102241b00a10201a410210220ba20a010d7f2002410f6a210f410020026b21072002410e6a210a410120026b210e2002410d6a210d410220026b210c0340200020056a2103200120056a220441037145200220054672450440200320042d00003a0000200f417f6a210f200741016a2107200a417f6a210a200e41016a210e200d417f6a210d200c41016a210c200541016a21050c010b0b200220056b210602400240024002402003410371220b044020064120490d03200b4101460d01200b4102460d02200b4103470d032003200120056a280200220a3a0000200041016a210b200220056b417f6a210c200521030340200c4113494504402003200b6a2208200120036a220941046a2802002206411874200a41087672360200200841046a200941086a2802002204411874200641087672360200200841086a2009410c6a28020022064118742004410876723602002008410c6a200941106a280200220a411874200641087672360200200341106a2103200c41706a210c0c010b0b2002417f6a2007416d2007416d4b1b200f6a4170716b20056b2106200120036a41016a2104200020036a41016a21030c030b2006210403402004411049450440200020056a2203200120056a2202290200370200200341086a200241086a290200370200200541106a2105200441706a21040c010b0b027f2006410871450440200120056a2104200020056a0c010b200020056a2202200120056a2201290200370200200141086a2104200241086a0b21052006410471044020052004280200360200200441046a2104200541046a21050b20064102710440200520042f00003b0000200441026a2104200541026a21050b2006410171450d03200520042d00003a000020000f0b2003200120056a2206280200220a3a0000200341016a200641016a2f00003b0000200041036a210b200220056b417d6a210720052103034020074111494504402003200b6a2208200120036a220941046a2802002206410874200a41187672360200200841046a200941086a2802002204410874200641187672360200200841086a2009410c6a28020022064108742004411876723602002008410c6a200941106a280200220a410874200641187672360200200341106a2103200741706a21070c010b0b2002417d6a200c416f200c416f4b1b200d6a4170716b20056b2106200120036a41036a2104200020036a41036a21030c010b2003200120056a2206280200220d3a0000200341016a200641016a2d00003a0000200041026a210b200220056b417e6a210720052103034020074112494504402003200b6a2208200120036a220941046a2802002206411074200d41107672360200200841046a200941086a2802002204411074200641107672360200200841086a2009410c6a28020022064110742004411076723602002008410c6a200941106a280200220d411074200641107672360200200341106a2103200741706a21070c010b0b2002417e6a200e416e200e416e4b1b200a6a4170716b20056b2106200120036a41026a2104200020036a41026a21030b20064110710440200320042d00003a00002003200428000136000120032004290005370005200320042f000d3b000d200320042d000f3a000f200441106a2104200341106a21030b2006410871044020032004290000370000200441086a2104200341086a21030b2006410471044020032004280000360000200441046a2104200341046a21030b20064102710440200320042f00003b0000200441026a2104200341026a21030b2006410171450d00200320042d00003a00000b20000bfc0202027f017e02402002450d00200020013a0000200020026a2203417f6a20013a000020024103490d00200020013a0002200020013a00012003417d6a20013a00002003417e6a20013a000020024107490d00200020013a00032003417c6a20013a000020024109490d002000410020006b41037122046a2203200141ff017141818284086c22013602002003200220046b417c7122046a2202417c6a200136020020044109490d002003200136020820032001360204200241786a2001360200200241746a200136020020044119490d002003200136021820032001360214200320013602102003200136020c200241706a20013602002002416c6a2001360200200241686a2001360200200241646a20013602002001ad220542208620058421052004200341047141187222016b2102200120036a2101034020024120490d0120012005370300200141186a2005370300200141106a2005370300200141086a2005370300200141206a2101200241606a21020c000b000b20000b3501017f230041106b220041808b0436020c418408200028020c41076a41787122003602004180082000360200418c083f003602000b9f0101047f230041106b220224002002200036020c027f02400240024020000440418c08200041086a22014110762200418c082802006a2203360200418408200141840828020022016a41076a4178712204360200200341107420044d0d0120000d020c030b41000c030b418c08200341016a360200200041016a21000b200040000d0010000b20012002410c6a4104100a1a200141086a0b200241106a24000b2f01027f2000410120001b2100034002402000100d22010d004190082802002202450d0020021103000c010b0b20010b7a01027f41e80a2100024003402000410371044020002d0000450d02200041016a21000c010b0b2000417c6a21000340200041046a22002802002201417f73200141fffdfb776a7141808182847871450d000b0340200141ff0171450d01200041016a2d00002101200041016a21000c000b000b200041e80a6b0bc10301067f024020002001460d00027f02400240200120006b20026b410020024101746b4b044020002001734103712103200020014f0d012003450d0220000c030b200020012002100a0f0b024020030d002001417f6a21030340200020026a220441037104402002450d052004417f6a200220036a2d00003a00002002417f6a21020c010b0b2000417c6a21032001417c6a2104034020024104490d01200220036a200220046a2802003602002002417c6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200241046a21062002417f73210503400240200120046a2107200020046a2208410371450d0020022004460d03200820072d00003a00002006417f6a2106200541016a2105200441016a21040c010b0b200220046b21014100210303402001410449450440200320086a200320076a280200360200200341046a21032001417c6a21010c010b0b200320076a210120022005417c2005417c4b1b20066a417c716b20046b2102200320086a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b20000b3f01027f024003402002044020002d0000220320012d00002204470d02200141016a2101200041016a21002002417f6a21020c010b0b41000f0b200320046b0b0a0041940841013602000b0a0041940841003602000b4d01017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b200020012802082001280204101520000b6401027f2002417049044002402002410a4d0440200020024101743a0000200041016a21030c010b200241106a4170712204100e21032000200236020420002004410172360200200020033602080b2003200120021016200220036a41003a00000f0b000b13002002047f200020012002100a0520000b1a0b130020002d0000410171044020002802081a0b0b3401017f2000200147044020002001280208200141016a20012d0000220041017122021b2001280204200041017620021b10190b0b9e0101037f410a2103027f0240027f024020002d00002205410171220404402000280200417e71417f6a21030b2003200249044020040d0120054101760c020b20040d02200041016a0c030b20002802040b210420002003200220036b2004200420022001101a0f0b20002802080b220320012002101b200220036a41003a000020002d0000410171450440200020024101743a00000f0b200020023602040bb70101027f416e20016b20024f0440027f200041016a20002d0000410171450d001a20002802080b2108027f416f200141e6ffffff074b0d001a410b20014101742207200120026a220220022007491b2202410b490d001a200241106a4170710b2207100e21022005044020022006200510160b200320046b220322060440200220056a200420086a200610160b200020023602082000200320056a220136020420002007410172360200200120026a41003a00000f0b000b13002002047f20002001200210100520000b1a0bf40101037f027f20002d00002203410171220445044020034101760c010b20002802040b220341004f0440410a2102027f02400240200404402000280200417e71417f6a21020b200220036b200149044020002002200120036a20026b20034100200141e80a101a0c010b2001450d0020040d01200041016a0c020b20000f0b20002802080b22022003047f200120026a20022003101b200141e80a6a41e80a200220036a41e80a4b1b41e80a200241e80a4d1b0541e80a0b2001101b200120036a2101024020002d0000410171450440200020014101743a00000c010b200020013602040b200120026a41003a000020000f0b000be30201057f027f20002d00002205410171220445044020054101760c010b20002802040b220641004f04402006200120062001491b2101410a2105200404402000280200417e71417f6a21050b200120066b20056a200349044020002005200320066a20016b20056b2006200120032002101a0f0b2004047f200028020805200041016a0b21040240024020012003460440200321010c010b200620016b2207450d00200120034b0440200420022003101b200320046a200120046a2007101b0c020b0240200420066a20024d200420024f720d00200120046a20024b0440200420022001101b200320016b200220036a2102200121084100210121030c010b2002200320016b6a21020b200420086a220520036a200120056a2007101b0b200420086a20022003101b0b200320016b20066a2101024020002d0000410171450440200020014101743a00000c010b200020013602040b200120046a41003a00000f0b000b6d01027f2001417049044002402001410a4d0440200020014101743a0000200041016a21020c010b200141106a4170712203100e21022000200136020420002003410172360200200020023602080b2001047f200241302001100b0520020b1a200120026a41003a00000f0b000b2301017f03402001410c46450440200020016a4100360200200141046a21010c010b0b0b190020004200370200200041086a41003602002000101f20000b7601037f101241980828020021000340200004400340419c08419c082802002201417f6a22023602002001410148450440200020024102746a22004184016a280200200041046a2802001013110100101241980828020021000c010b0b419c084120360200419808200028020022003602000c010b0b0b960101027f1012419808280200220145044041980841a00836020041a00821010b0240419c0828020022024120460440418402100d2201044020014100418402100b1a0b2001450d0120014198082802003602004198082001360200419c084100360200410021020b419c08200241016a360200200120024102746a22014184016a4100360200200141046a200036020010130f0b10130b070041a40a10170b780020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000102520012802044f0d002002410471450440200042003702000c010b10000b024002402002411071450d002000102520012802044d0d0020024104710d01200042003702000b20000f0b100020000b290002402000280204044020002802002c0000417f4c0d0141010f0b41000f0b20001026200010276a0b240002402000280204450d0020002802002c0000417f4c0d0041000f0b2000102c41016a0b8a0301047f02400240200028020404402000102d4101210220002802002c00002201417f4c0d010c020b41000f0b200141ff0171220241b7014d0440200241807f6a0f0b02400240200141ff0171220141bf014d04400240200041046a22042802002201200241c97e6a22034d047f100020042802000520010b4102490d0020002802002d00010d0010000b200341054f044010000b20002802002d000145044010000b410021024100210103402001200346450440200028020020016a41016a2d00002002410874722102200141016a21010c010b0b200241384f0d010c020b200141f7014d0440200241c07e6a0f0b0240200041046a22042802002201200241897e6a22034d047f100020042802000520010b4102490d0020002802002d00010d0010000b200341054f044010000b20002802002d000145044010000b410021024100210103402001200346450440200028020020016a41016a2d00002002410874722102200141016a21010c010b0b20024138490d010b200241ff7d490d010b100020020f0b20020b3902017f017e230041306b2201240020012000290200220237031020012002370308200141186a200141086a411410241025200141306a24000b5e01027f2000027f027f2001280200220504404100200220036a200128020422014b2001200249720d011a410020012003490d021a200220056a2104200120026b20032003417f461b0c020b41000b210441000b360204200020043602000b2101017f20011027220220012802044b044010000b2000200120011026200210290b900302097f017e230041406a220324002001280208220520024b0440200341386a2001102a200320032903383703182001200341186a102836020c200341306a2001102a410021052001027f410020032802302206450d001a410020032802342208200128020c2207490d001a200820072007417f461b210420060b360210200141146a2004360200200141086a41003602000b200141106a2109200141146a21072001410c6a2106200141086a210803400240200520024f0d002007280200450d00200341306a2001102a41002105027f2003280230220a044041002003280234220b20062802002204490d011a200b20046b21052004200a6a0c010b41000b210420072005360200200920043602002003200536022c2003200436022820032003290328370310200341306a20094100200341106a1028102920092003290330220c37020020062006280200200c422088a76a3602002008200828020041016a22053602000c010b0b20032009290200220c3703202003200c3703082000200341086a411410241a200341406b24000b4101017f02402000280204450d0020002802002d0000220041bf014d0440200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b4401017f200028020445044010000b0240200028020022012d0000418101470d00200041046a28020041014d047f100020002802000520010b2c00014100480d0010000b0b9f0101037f0240200028020404402000102d200028020022022c000022014100480d0120014100470f0b41000f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200041046a28020041014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a200041046a280200200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b2c002000200220016b22021031200028020020002802046a20012002100a1a2000200028020420026a3602040b9e0201077f02402001450d002000410c6a2107200041106a2105200041046a21060340200528020022022007280200460d01200241786a28020020014904401000200528020021020b200241786a2203200328020020016b220136020020010d01200520033602002000410120062802002002417c6a28020022016b22021032220341016a20024138491b2204200628020022086a10332004200120002802006a22046a2004200820016b10101a0240200241374d0440200028020020016a200241406a3a00000c010b200341f7016a220441ff014d0440200028020020016a20043a00002000280200200120036a6a210103402002450d02200120023a0000200241087621022001417f6a21010c000b000b10000b410121010c000b000b0b1b00200028020420016a220120002802084b04402000200110340b0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b0f00200020011034200020013602040b3901017f200028020820014904402001100d220220002802002000280204100a1a20002802001a200041086a2001360200200020023602000b0b2500200041011031200028020020002802046a20013a00002000200028020441016a3602040b2b01027f20011032220241b7016a22034180024e044010000b2000200341ff0171103520002001200210370b3d002000200028020420026a1033200028020020002802046a417f6a2100034020010440200020013a0000200141087621012000417f6a21000c010b0b0ba00101037f230041106b2202240020012802002103024002400240024020012802042201410146044020032c000022044100480d012000200441ff017110350c040b200141374b0d010b200020014180017341ff017110350c010b2000200110360b2002200136020c2002200336020820022002290308370300200020022802002201200120022802046a102f2000410010300b200041011030200241106a24000b830101037f02400240200150450440200142ff00560d0120002001a741ff017110350c020b200041800110350c010b02402001103a220241374d0440200020024180017341ff017110350c010b20021032220341b7016a22044180024f044010000b2000200441ff0171103520002002200310370b200020012002103b0b2000410110300b3202017f017e034020002002845045044020024238862000420888842100200141016a2101200242088821020c010b0b20010b5101017e2000200028020420026a1033200028020020002802046a417f6a21000340200120038450450440200020013c0000200342388620014208888421012000417f6a2100200342088821030c010b0b0b070041b00a10170b2701017f03402001411446450440200020016a41003a0000200141016a21010c010b0b200010010b3a01017f230041306b22012400200141086a103d200141206a200141086a103f200041186a200141206a1040200141206a1017200141306a24000bc80201067f230041206b22022400200241186a410036020020024200370310200241106a41eb0a41001015200241086a4100360200200242003703002002200228021420022d0010220341017620034101711b220341286a101e027f20022d00102204410171450440200241106a410172210520044101760c010b2002280218210520022802140b2106410021042002200320052006101d20024101722105200241086a21060340200441144704402006280200200520022d00004101711b20036a200120046a22072d000041047641ec0a6a2d00003a00002006280200200520022d00004101711b20036a41016a20072d0000410f7141ec0a6a2d00003a0000200441016a2104200341026a21030c010b0b200241106a101720002002100f101c2201290200370200200041086a200141086a2802003602002001101f20021017200241206a24000b6400024020002d0000410171450440200041003b01000c010b200028020841003a00002000410036020420002d0000410171450d00200041086a2802001a200041003602000b20002001290200370200200041086a200141086a2802003602002001101f0bfe0101057f230041306b22012400200141186a103d200141086a200141186a103f027f02400240024002402000411c6a28020020002d001822024101762204200241017122051b2203200128020c20012d00082202410176200241017122021b470d002001280210200141086a41017220021b210220050d01410020046b2103200041196a210003402003450d0320002d000020022d0000470d01200341016a2103200241016a2102200041016a21000c000b000b200141086a1017417f0c030b2003450d00200041206a2802002002200310112100200141086a1017417f20000d021a0c010b200141086a10170b200141186a1002450b200141306a24000bc80402057f017e23004190016b22002400100910032201100d22021004200041d8006a200041086a20022001104322034100102b200041d8006a102d02400240200041d8006a102e450d00200028025c450d0020002802582d000041c001490d010b10000b200041386a200041d8006a1044200028023c220141094f044010000b200028023821020340200104402001417f6a210120023100002005420886842105200241016a21020c010b0b024002402005500d0041bc0a10452005510440200041d8006a1046200041d8006a103e10470c020b41c10a10452005510440200041d8006a1046200041d8006a10412103200041206a10482101200041d0006a4100360200200041c8006a4200370300200041406b420037030020004200370338200041386a2003ac22054201862005423f87852205104920002802382103200041386a410472104a20012003104b20012005104c200128020c200141106a28020047044010000b2001280200200128020410052001104d10470c020b41d20a10452005510440200041206a10202101200041d8006a20034101102b200041d8006a2001104e200041d8006a1046200041f0006a200041386a2001101422031018200310171047200110170c020b41dd0a10452005520d00200041d8006a104620004180016a200041f0006a10142102200041386a104822012002104f104b2001200041206a200210142204105020041017200128020c200141106a28020047044010000b2001280200200128020410052001104d2002101710470c010b10000b20004190016a24000b3401017f230041106b220324002003200236020c200320013602082003200329030837030020002003411c1024200341106a24000be60101047f200110272204200128020422024b04401000200141046a28020021020b200128020021052000027f024002400240027f0240200204404100210120052c00002203417f4c0d012005450d030c040b41000c010b200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a210120050d010b410021030c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010bf601010a7f230041406a22012400200010202107200042afb59bdd9e8485b9f800370310200041186a10202105200141286a104822032000290310104c200328020c200341106a28020047044010000b0240200328020022082003280204220910062202044020014100360220200142003703182002417f4c0d01200141206a2002100e41002002100b220620026a220a3602002001200a36021c2001200636021820082009200620021007417f47044020012001280218220441016a200128021c2004417f736a10432005104e200221040b200141186a10510b2003104d20044504402005200710180b200141406b240020000f0b000bbd03010b7f230041e0006b22012400200141286a10482105200141d8006a4100360200200141d0006a4200370300200141c8006a420037030020014200370340200141406b2000290310104920012802402102200141406b410472104a20052002104b20052000290310104c200528020c200541106a28020047044010000b200528020421082005280200200141406b10482102200041186a2207104f21064101100e220341fe013a0000200120033602182001200341016a22043602202001200436021c200228020c200241106a2802004704401000200128021c2104200128021821030b200420036b22042002280204220a6a220b20022802084b047f2002200b1052200241046a28020005200a0b20022802006a20032004100a1a200241046a2203200328020020046a3602002002200128021c20066a20012802186b104b2002200141086a20071014220310502003101702402002410c6a2204280200200241106a220628020047044010002002280200210320042802002006280200460d0110000c010b200228020021030b20082003200241046a2802001008200141186a10512002104d2005104d2007101720001017200141e0006a24000b29002000410036020820004200370200200041001052200041146a41003602002000420037020c20000b7902017f017e4101210220014280015a044041002102034020012003845045044020034238862001420888842101200241016a2102200342088821030c010b0b200241384f047f2002103220026a0520020b41016a21020b200041186a2802000440200041046a105421000b2000200028020020026a3602000bdc0201067f200028020422012000280210220241087641fcffff07716a2103027f2001200028020822054704402001200028021420026a220441087641fcffff07716a280200200441ff07714102746a2106200041146a21042003280200200241ff07714102746a0c010b200041146a210441000b2102034020022006470440200241046a220220032802006b418020470d0120032802042102200341046a21030c010b0b20044100360200200041086a21020340200520016b410275220341034f044020012802001a200041046a2201200128020041046a2201360200200228020021050c010b0b0240200041106a027f2003410147044020034102470d024180080c010b4180040b3602000b03402001200547044020012802001a200141046a21010c010b0b200041086a22032802002101200041046a280200210203402001200247044020032001417c6a22013602000c010b0b20002802001a0b1300200028020820014904402000200110520b0b08002000200110390b1c01017f200028020c22010440200041106a20013602000b200010530bad0201057f230041206b22022400024002400240024002402000280204450d0020002802002d000041c0014f0d00200241186a20001044200010272103200228021822000440200228021c220420034f0d020b41002100200241106a410036020020024200370308410021030c020b200241086a10201a0c030b200241106a410036020020024200370308200420032003417f461b22034170490440200020036a21052003410a4d0d01200341106a4170712206100e21042002200336020c20022006410172360208200220043602100c020b000b200220034101743a0008200241086a41017221040b034020002005470440200420002d00003a0000200441016a2104200041016a21000c010b0b200441003a00000b2001200241086a1040200241086a1017200241206a24000bba0101047f230041306b22012400200141286a4100360200200141206a4200370300200141186a420037030020014200370310410121020240200120001014220328020420032d00002200410176200041017122041b2200450d0002400240200041014604402003280208200341016a20041b2c0000417f4a0d030c010b200041374b0d010b200041016a21020c010b2000103220006a41016a21020b2001200236021020031017200141106a410472104a200141306a240020020b5201037f230041106b2202240020022001280208200141016a20012d0000220341017122041b36020820022001280204200341017620041b36020c20022002290308370300200020021038200241106a24000b1501017f200028020022010440200020013602040b0b3601017f200028020820014904402001100d20002802002000280204100a210220001053200041086a2001360200200020023602000b0b080020002802001a0b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0b0b47010041bc0a0b40696e697400646973746f72795f636f6e7472616374007365745f737472696e67006765745f737472696e67003078000030313233343536373839616263646566";

    public static String BINARY = BINARY_0;

    public static final String FUNC_SET_STRING = "set_string";

    public static final String FUNC_GET_STRING = "get_string";

    public static final String FUNC_DISTORY_CONTRACT = "distory_contract";

    protected ContractDistoryWithPermissionCheck(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ContractDistoryWithPermissionCheck(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<TransactionReceipt> set_string(String name) {
        final WasmFunction function = new WasmFunction(FUNC_SET_STRING, Arrays.asList(name), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> set_string(String name, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SET_STRING, Arrays.asList(name), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static RemoteCall<ContractDistoryWithPermissionCheck> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDistoryWithPermissionCheck.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ContractDistoryWithPermissionCheck> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDistoryWithPermissionCheck.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ContractDistoryWithPermissionCheck> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDistoryWithPermissionCheck.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<ContractDistoryWithPermissionCheck> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDistoryWithPermissionCheck.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<String> get_string() {
        final WasmFunction function = new WasmFunction(FUNC_GET_STRING, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<TransactionReceipt> distory_contract() {
        final WasmFunction function = new WasmFunction(FUNC_DISTORY_CONTRACT, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> distory_contract(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_DISTORY_CONTRACT, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static ContractDistoryWithPermissionCheck load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new ContractDistoryWithPermissionCheck(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static ContractDistoryWithPermissionCheck load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new ContractDistoryWithPermissionCheck(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
