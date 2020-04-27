package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Uint64;
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
public class ContractDelegateCallStorageString extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001671160027f7f0060017f0060017f017f60000060037f7f7f0060037f7f7f017f60027f7f017f60027f7e0060047f7f7f7f0060047f7f7f7f017f60017f017e60077f7f7f7f7f7f7f0060037f7e7f006000017f60057f7f7f7f7f017f60017e017f60037f7f7e017e02c4010803656e760c706c61746f6e5f70616e6963000303656e7614706c61746f6e5f64656c65676174655f63616c6c000e03656e7617706c61746f6e5f6765745f73746174655f6c656e677468000603656e7610706c61746f6e5f6765745f7374617465000903656e7610706c61746f6e5f7365745f7374617465000803656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000d03656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e00000356550305010302020503030604040100040b01020301010502020202080004020102050400000200000100010000040400070f0c0609000400011002070000070000010101020000010500020200030a0a0006000001020405017001030305030100020615037f0141f08a040b7f0041f08a040b7f0041f00a0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300080b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974001a06696e766f6b6500540908010041010b021c3f0aff5a551800100b41a40a10191a4101101b41b00a10191a4102101b0ba20a010d7f2002410f6a210f410020026b21072002410e6a210a410120026b210e2002410d6a210d410220026b210c0340200020056a2103200120056a220441037145200220054672450440200320042d00003a0000200f417f6a210f200741016a2107200a417f6a210a200e41016a210e200d417f6a210d200c41016a210c200541016a21050c010b0b200220056b210602400240024002402003410371220b044020064120490d03200b4101460d01200b4102460d02200b4103470d032003200120056a280200220a3a0000200041016a210b200220056b417f6a210c200521030340200c4113494504402003200b6a2208200120036a220941046a2802002206411874200a41087672360200200841046a200941086a2802002204411874200641087672360200200841086a2009410c6a28020022064118742004410876723602002008410c6a200941106a280200220a411874200641087672360200200341106a2103200c41706a210c0c010b0b2002417f6a2007416d2007416d4b1b200f6a4170716b20056b2106200120036a41016a2104200020036a41016a21030c030b2006210403402004411049450440200020056a2203200120056a2202290200370200200341086a200241086a290200370200200541106a2105200441706a21040c010b0b027f2006410871450440200120056a2104200020056a0c010b200020056a2202200120056a2201290200370200200141086a2104200241086a0b21052006410471044020052004280200360200200441046a2104200541046a21050b20064102710440200520042f00003b0000200441026a2104200541026a21050b2006410171450d03200520042d00003a000020000f0b2003200120056a2206280200220a3a0000200341016a200641016a2f00003b0000200041036a210b200220056b417d6a210720052103034020074111494504402003200b6a2208200120036a220941046a2802002206410874200a41187672360200200841046a200941086a2802002204410874200641187672360200200841086a2009410c6a28020022064108742004411876723602002008410c6a200941106a280200220a410874200641187672360200200341106a2103200741706a21070c010b0b2002417d6a200c416f200c416f4b1b200d6a4170716b20056b2106200120036a41036a2104200020036a41036a21030c010b2003200120056a2206280200220d3a0000200341016a200641016a2d00003a0000200041026a210b200220056b417e6a210720052103034020074112494504402003200b6a2208200120036a220941046a2802002206411074200d41107672360200200841046a200941086a2802002204411074200641107672360200200841086a2009410c6a28020022064110742004411076723602002008410c6a200941106a280200220d411074200641107672360200200341106a2103200741706a21070c010b0b2002417e6a200e416e200e416e4b1b200a6a4170716b20056b2106200120036a41026a2104200020036a41026a21030b20064110710440200320042d00003a00002003200428000136000120032004290005370005200320042f000d3b000d200320042d000f3a000f200441106a2104200341106a21030b2006410871044020032004290000370000200441086a2104200341086a21030b2006410471044020032004280000360000200441046a2104200341046a21030b20064102710440200320042f00003b0000200441026a2104200341026a21030b2006410171450d00200320042d00003a00000b20000bc70201027f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122016a22004100360200200041840220016b417c7122026a2201417c6a4100360200024020024109490d002000410036020820004100360204200141786a4100360200200141746a410036020020024119490d002000410036021820004100360214200041003602102000410036020c200141706a41003602002001416c6a4100360200200141686a4100360200200141646a41003602002002200041047141187222026b2101200020026a2100034020014120490d0120004200370300200041186a4200370300200041106a4200370300200041086a4200370300200041206a2100200141606a21010c000b000b0b3501017f230041106b220041f08a0436020c418408200028020c41076a41787122003602004180082000360200418c083f003602000b9f0101047f230041106b220224002002200036020c027f02400240024020000440418c08200041086a22014110762200418c082802006a2203360200418408200141840828020022016a41076a4178712204360200200341107420044d0d0120000d020c030b41000c030b418c08200341016a360200200041016a21000b200040000d0010000b20012002410c6a410410091a200141086a0b200241106a24000b2f01027f2000410120001b2100034002402000100c22010d004190082802002202450d0020021103000c010b0b20010bc10301067f024020002001460d00027f02400240200120006b20026b410020024101746b4b044020002001734103712103200020014f0d012003450d0220000c030b20002001200210090f0b024020030d002001417f6a21030340200020026a220441037104402002450d052004417f6a200220036a2d00003a00002002417f6a21020c010b0b2000417c6a21032001417c6a2104034020024104490d01200220036a200220046a2802003602002002417c6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200241046a21062002417f73210503400240200120046a2107200020046a2208410371450d0020022004460d03200820072d00003a00002006417f6a2106200541016a2105200441016a21040c010b0b200220046b21014100210303402001410449450440200320086a200320076a280200360200200341046a21032001417c6a21010c010b0b200320076a210120022005417c2005417c4b1b20066a417c716b20046b2102200320086a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b20000b0a0041940841013602000b0a0041940841003602000b4d01017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b200020012802082001280204101220000b6401027f2002417049044002402002410a4d0440200020024101743a0000200041016a21030c010b200241106a4170712204100d21032000200236020420002004410172360200200020033602080b2003200120021013200220036a41003a00000f0b000b13002002047f20002001200210090520000b1a0b130020002d0000410171044020002802081a0b0b3401017f2000200147044020002001280208200141016a20012d0000220041017122021b2001280204200041017620021b10160b0bab0101037f410a2103027f0240027f024020002d00002205410171220404402000280200417e71417f6a21030b2003200249044020040d0120054101760c020b20040d02200041016a0c030b20002802040b210420002003200220036b200420042002200110170f0b20002802080b220421032002047f200320012002100e0520030b1a200220046a41003a000020002d0000410171450440200020024101743a00000f0b200020023602040bb70101027f416e20016b20024f0440027f200041016a20002d0000410171450d001a20002802080b2108027f416f200141e6ffffff074b0d001a410b20014101742207200120026a220220022007491b2202410b490d001a200241106a4170710b2207100d21022005044020022006200510130b200320046b220322060440200220056a200420086a200610130b200020023602082000200320056a220136020420002007410172360200200120026a41003a00000f0b000b2301017f03402001410c46450440200020016a4100360200200141046a21010c010b0b0b190020004200370200200041086a41003602002000101820000b7601037f100f41980828020021000340200004400340419c08419c082802002201417f6a22023602002001410148450440200020024102746a22004184016a280200200041046a2802001010110100100f41980828020021000c010b0b419c084120360200419808200028020022003602000c010b0b0b900101027f100f419808280200220145044041980841a00836020041a00821010b0240419c0828020022024120460440418402100c220104402001100a0b2001450d0120014198082802003602004198082001360200419c084100360200410021020b419c08200241016a360200200120024102746a22014184016a4100360200200141046a200036020010100f0b10100b070041a40a10140b780020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000101e20012802044f0d002002410471450440200042003702000c010b10000b024002402002411071450d002000101e20012802044d0d0020024104710d01200042003702000b20000f0b100020000b290002402000280204044020002802002c0000417f4c0d0141010f0b41000f0b2000101f200010206a0b240002402000280204450d0020002802002c0000417f4c0d0041000f0b2000102541016a0b8a0301047f0240024020002802040440200010264101210220002802002c00002201417f4c0d010c020b41000f0b200141ff0171220241b7014d0440200241807f6a0f0b02400240200141ff0171220141bf014d04400240200041046a22042802002201200241c97e6a22034d047f100020042802000520010b4102490d0020002802002d00010d0010000b200341054f044010000b20002802002d000145044010000b410021024100210103402001200346450440200028020020016a41016a2d00002002410874722102200141016a21010c010b0b200241384f0d010c020b200141f7014d0440200241c07e6a0f0b0240200041046a22042802002201200241897e6a22034d047f100020042802000520010b4102490d0020002802002d00010d0010000b200341054f044010000b20002802002d000145044010000b410021024100210103402001200346450440200028020020016a41016a2d00002002410874722102200141016a21010c010b0b20024138490d010b200241ff7d490d010b100020020f0b20020b3902017f017e230041306b2201240020012000290200220237031020012002370308200141186a200141086a4114101d101e200141306a24000b5e01027f2000027f027f2001280200220504404100200220036a200128020422014b2001200249720d011a410020012003490d021a200220056a2104200120026b20032003417f461b0c020b41000b210441000b360204200020043602000b2101017f20011020220220012802044b044010000b200020012001101f200210220b900302097f017e230041406a220324002001280208220520024b0440200341386a20011023200320032903383703182001200341186a102136020c200341306a20011023410021052001027f410020032802302206450d001a410020032802342208200128020c2207490d001a200820072007417f461b210420060b360210200141146a2004360200200141086a41003602000b200141106a2109200141146a21072001410c6a2106200141086a210803400240200520024f0d002007280200450d00200341306a2001102341002105027f2003280230220a044041002003280234220b20062802002204490d011a200b20046b21052004200a6a0c010b41000b210420072005360200200920043602002003200536022c2003200436022820032003290328370310200341306a20094100200341106a1021102220092003290330220c37020020062006280200200c422088a76a3602002008200828020041016a22053602000c010b0b20032009290200220c3703202003200c3703082000200341086a4114101d1a200341406b24000b4101017f02402000280204450d0020002802002d0000220041bf014d0440200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b4401017f200028020445044010000b0240200028020022012d0000418101470d00200041046a28020041014d047f100020002802000520010b2c00014100480d0010000b0b9f0101037f02402000280204044020001026200028020022022c000022014100480d0120014100470f0b41000f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200041046a28020041014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a200041046a280200200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b1f01017f200020012802002203200320012802046a102920002002102a20000b2c002000200220016b2202102b200028020020002802046a2001200210091a2000200028020420026a3602040b9e0201077f02402001450d002000410c6a2107200041106a2105200041046a21060340200528020022022007280200460d01200241786a28020020014904401000200528020021020b200241786a2203200328020020016b220136020020010d01200520033602002000410120062802002002417c6a28020022016b2202102c220341016a20024138491b2204200628020022086a102d2004200120002802006a22046a2004200820016b100e1a0240200241374d0440200028020020016a200241406a3a00000c010b200341f7016a220441ff014d0440200028020020016a20043a00002000280200200120036a6a210103402002450d02200120023a0000200241087621022001417f6a21010c000b000b10000b410121010c000b000b0b1b00200028020420016a220120002802084b044020002001102e0b0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b0f0020002001102e200020013602040b3901017f200028020820014904402001100c22022000280200200028020410091a20002802001a200041086a2001360200200020023602000b0b2e01017f230041106b2201240020014102360200200120002802043602042000410c6a20011030200141106a24000b3701017f20002802042202200028020849044020022001290200370200200041046a2200200028020041086a3602000f0b2000200110320b1501017f200028020022010440200020013602040b0b7201027f230041206b22032400200341086a2000200028020420002802006b41037541016a103a200028020420002802006b410375200041086a103b220228020820012902003702002002200228020841086a36020820002002103c20022002280204103e20022802001a200341206a24000b250020004101102b200028020020002802046a20013a00002000200028020441016a3602040b2a01017f20022001102c22026a22034180024e044010000b2000200341ff0171103320002001200210350b3d002000200028020420026a102d200028020020002802046a417f6a2100034020010440200020013a0000200141087621012000417f6a21000c010b0b0b930101037f230041106b2202240020012802002103024002400240024020012802042201410146044020032c000022044100480d012000200441ff017110330c040b200141374b0d010b200020014180017341ff017110330c010b2000200141b70110340b2002200136020c200220033602082002200229030837030020002002410010281a0b20004101102a200241106a24000b830101037f02400240200150450440200142ff00560d0120002001a741ff017110330c020b200041800110330c010b024020011038220241374d0440200020024180017341ff017110330c010b2002102c220341b7016a22044180024f044010000b2000200441ff0171103320002002200310350b20002001200210390b20004101102a0b3202017f017e034020002002845045044020024238862000420888842100200141016a2101200242088821020c010b0b20010b5101017e2000200028020420026a102d200028020020002802046a417f6a21000340200120038450450440200020013c0000200342388620014208888421012000417f6a2100200342088821030c010b0b0b40002001418080808002490440200028020820002802006b220041037541feffffff004d047f20012000410275220020002001491b0541ffffffff010b0f0b000b6401017f2000410036020c200041106a200336020020010440027f20014180808080024904402001410374100d0c010b000b21040b200020043602002000200420024103746a22023602082000410c6a200420014103746a3602002000200236020420000b6701017f20002802002000280204200141046a103d200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b270020022002280200200120006b22016b2202360200200141014e044020022000200110091a0b0b2c01017f20002802082102200041086a2100034020012002464504402000200241786a22023602000c010b0b0b070041b00a10140bc00902067f017e23004180016b22032400200341106a410036020020034200370308200341086a41bc0a410a1012027f20032d0008220441017145044020044101762104200341086a4101720c010b200328020c210420032802100b210642a5c688a1c89ca7f94b21090340200404402004417f6a21042006300000200942b383808080207e852109200641016a21060c010b0b200341d8006a10412105200341c8006a2001101121082005102f200341406b4100360200200341386a4200370300200341306a420037030020034200370328200341286a20091042200341286a200341f0006a2008101122011043200110142005200328022810442005200910452005200341f0006a200810112201104620011014027f027f200528020c200541106a22012802004704401000200541046a2107200528020022062005410c6a2802002001280200460d011a100020052802000c020b200541046a210720052802000b22060b21014100210420034100360220200342003703180240200120072802006a20066b2201450d00200341186a2001104720014101480d00200328021c2006200110091a2003200328021c20016a36021c0b200341286a41047210482008101420051049200341086a101402402000280204220120002d000022054101762207200541017122061b22054102490d002000280208200041016a20061b22082d00004130470d0020082d000141f8004641017421040b2003410036026020034200370358027f0240200541016a20046b4101762205047f200341386a200341e0006a36020020032005100d22013602582003200136025c20034200370330200342003703282003200120056a360260200341286a104a20002d00002205410176210720054101712106200041046a2802000520010b200720061b41017104402000280208200041016a20061b20046a2c0000104b2201417f460d01200320013a0028200341d8006a200341286a104c200441017221040b200041016a2101200041046a2105200041086a2106024003402004200528020020002d00002207410176200741017122071b4904402006280200200120071b20046a22072c0000104b2208417f46200741016a2c0000104b2207417f46720d022003200720084104746a3a0028200441026a2104200341d8006a200341286a104c0c010b0b200320032802583602702003200329025c3702742003420037025c200341d8006a0c020b20034200370370200341f8006a0c010b20034200370370200341f8006a0b4100360200200341d8006a103141002104034020044114470440200341286a20046a41003a0000200441016a21040c010b0b2003280274200328027022006b21014100210403402001200446200441134b72450440200341286a20046a200020046a2d00003a0000200441016a21040c010b0b200341f0006a10314100210420022109034020095045044020094208882109200441016a21040c010b0b2003410036026020034200370358200341d8006a2004104d200328025c417f6a21040340200250450440200420023c00002004417f6a2104200242088821020c010b0b200341286a20032802182200200328021c20006b20032802582200200328025c20006b10011a200341d8006a1031200341186a103120034180016a240042000b29002000410036020820004200370200200041001053200041146a41003602002000420037020c20000b7902017f017e4101210220014280015a044041002102034020012003845045044020034238862001420888842101200241016a2102200342088821030c010b0b200241384f047f2002102c20026a0520020b41016a21020b200041186a2802000440200041046a105c21000b2000200028020020026a3602000b890101037f410121030240200128020420012d00002202410176200241017122041b2202450d0002400240200241014604402001280208200141016a20041b2c0000417f4a0d030c010b200241374b0d010b200241016a21030c010b2002102c20026a41016a21030b200041186a2802000440200041046a105c21000b2000200028020020036a3602000b1300200028020820014904402000200110530b0b08002000200110370b5201037f230041106b2202240020022001280208200141016a20012d0000220341017122041b36020820022001280204200341017620041b36020c20022002290308370300200020021036200241106a24000b2a01017f2001417f4a044020002001100d2202360200200020023602042000200120026a3602080f0b000bdc0201067f200028020422012000280210220241087641fcffff07716a2103027f2001200028020822054704402001200028021420026a220441087641fcffff07716a280200200441ff07714102746a2106200041146a21042003280200200241ff07714102746a0c010b200041146a210441000b2102034020022006470440200241046a220220032802006b418020470d0120032802042102200341046a21030c010b0b20044100360200200041086a21020340200520016b410275220341034f044020012802001a200041046a2201200128020041046a2201360200200228020021050c010b0b0240200041106a027f2003410147044020034102470d024180080c010b4180040b3602000b03402001200547044020012802001a200141046a21010c010b0b200041086a22032802002101200041046a280200210203402001200247044020032001417c6a22013602000c010b0b20002802001a0b1c01017f200028020c22010440200041106a20013602000b2000105b0b3801037f2000280208210120002802042102200041086a210303402001200247044020032001417f6a22013602000c010b0b20002802001a0b4901017f415021010240200041506a41ff0171410a4f044041a97f21012000419f7f6a41ff017141064f0d010b200020016a0f0b200041496a417f200041bf7f6a41ff01714106491b0bcd0101047f230041206b220324000240200028020422022000280208490440200220012d00003a0000200041046a2200200028020041016a3602000c010b2000200241016a20002802006b10582104200341186a200041086a3602004100210220034100360214200041046a28020020002802006b2105200404402004100d21020b20032002360208200341146a200220046a360200200220056a220220012d00003a00002003200236020c2003200241016a3602102000200341086a1059200341086a104a0b200341206a24000bab0201057f230041206b220324000240024020002802042202200028020022056b22042001490440200028020820026b200120046b4f0d012000200110582105200341186a200041086a36020020034100360214200041046a28020020002802006b210641002102200504402005100d21020b20032002360208200341146a200220056a3602002003200220066a22023602102003200236020c200420016b2101200341106a21040340200241003a00002004200428020041016a2202360200200141016a22010d000b2000200341086a1059200341086a104a0c020b200420014d0d01200041046a200120056a3602000c010b200420016b2101200041046a21000340200241003a00002000200028020041016a2202360200200141016a22010d000b0b200341206a24000bcf04010b7f23004180016b22012400200141086a10192109200142e299efdb8683ebcf58370318200141206a10192105200141e8006a1041220320012903181045200328020c200341106a28020047044010000b024020032802002206200328020422071002220404402001410036024820014200370340200141406b2004104d2006200720012802402206200128024420066b1003417f470440200141d0006a2001280240220241016a20012802442002417f736a104f20051050200421020b200141406b10310c010b0b2003104920024504402005200910150b2000200510111a200141e8006a10412204200141186a105110442004200141186a2903001045200428020c200441106a28020047044010000b200428020421062004280200200141d0006a104121022005105221084101100d220041fe013a0000200120003602402001200041016a220336024820012003360244200228020c200241106a280200470440100020012802442103200128024021000b200320006b22032002280204220a6a220b20022802084b047f2002200b1053200241046a28020005200a0b20022802006a2000200310091a200241046a2200200028020020036a3602002002200128024420086a20012802406b10442002200141306a20051011220010462000101402402002410c6a2203280200200241106a220828020047044010002002280200210020032802002008280200460d0110000c010b200228020021000b20062000200241046a2802001004200141406b10312002104920041049200510142009101420014180016a24000b3401017f230041106b220324002003200236020c200320013602082003200329030837030020002003411c101d200341106a24000b890301057f230041206b22022400024002400240024002402000280204450d0020002802002d000041c0014f0d00200241186a2000105a200010202103200228021822000440200228021c220420034f0d020b41002100200241106a410036020020024200370308410021030c020b200241086a10191a0c030b200241106a410036020020024200370308200420032003417f461b22034170490440200020036a21052003410a4d0d01200341106a4170712206100d21042002200336020c20022006410172360208200220043602100c020b000b200220034101743a0008200241086a41017221040b034020002005470440200420002d00003a0000200441016a2104200041016a21000c010b0b200441003a00000b024020012d0000410171450440200141003b01000c010b200128020841003a00002001410036020420012d0000410171450d00200141086a2802001a200141003602000b20012002290308370200200141086a200241106a280200360200200241086a1018200241086a1014200241206a24000b4e01017f230041206b22012400200141186a4100360200200141106a4200370300200141086a420037030020014200370300200120002903001042200128020020014104721048200141206a24000b5b01017f230041306b22012400200141286a4100360200200141206a4200370300200141186a420037030020014200370310200141106a20012000101122001043200010142001280210200141106a4104721048200141306a24000b3601017f200028020820014904402001100c20002802002000280204100921022000105b200041086a2001360200200020023602000b0b9f0302067f017e23004180016b22002400100810052201100c22021006200041206a200020022001104f22014100102402400240200041206a10552206500d0041c70a10562006510d0141cc0a10562006510440200041206a101921022000412c6a101921032000420037033820004101360240200020013602702000200041406b360274200041f0006a20021057200041f0006a20031057200041d8006a2001200028024010242000200041d8006a10553703382000200041f0006a200210112204200041406b200310112205200029033810402206370350200041d8006a10412201200041d0006a10511044200120061045200128020c200141106a28020047044010000b20012802002001280204100720011049200510142004101420031014200210140c020b41e50a10562006520d00200041f0006a104e200041206a10412201200041f0006a105210442001200041d8006a200041f0006a10112202104620021014200128020c200141106a28020047044010000b20012802002001280204100720011049200041f0006a10140c010b10000b20004180016a24000b850102027f017e230041106b22012400200010260240024020001027450d002000280204450d0020002802002d000041c001490d010b10000b200141086a2000105a200128020c220041094f044010000b200128020821020340200004402000417f6a210020023100002003420886842103200241016a21020c010b0b200141106a240020030b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b4301017f230041206b22022400200241086a200028020020002802042802001024200241086a2001105020002802042200200028020041016a360200200241206a24000b39002001417f4a0440200028020820002802006b220041feffffff034d047f20012000410174220020002001491b0541ffffffff070b0f0b000b940101037f200120012802042000280204200028020022046b22026b2203360204200241004a044020032004200210091a200141046a28020021030b2000280200210220002003360200200141046a22032002360200200041046a220228020021042002200128020836020020012004360208200028020821022000200128020c3602082001200236020c200120032802003602000be60101047f200110202204200128020422024b04401000200141046a28020021020b200128020021052000027f024002400240027f0240200204404100210120052c00002203417f4c0d012005450d030c040b41000c010b200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a210120050d010b410021030c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b080020002802001a0b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0b0b3a010041bc0a0b337365745f737472696e6700696e69740064656c65676174655f63616c6c5f7365745f737472696e67006765745f737472696e67";

    public static String BINARY = BINARY_0;

    public static final String FUNC_DELEGATE_CALL_SET_STRING = "delegate_call_set_string";

    public static final String FUNC_GET_STRING = "get_string";

    protected ContractDelegateCallStorageString(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ContractDelegateCallStorageString(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<ContractDelegateCallStorageString> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDelegateCallStorageString.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ContractDelegateCallStorageString> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDelegateCallStorageString.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ContractDelegateCallStorageString> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDelegateCallStorageString.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<ContractDelegateCallStorageString> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDelegateCallStorageString.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<TransactionReceipt> delegate_call_set_string(String target_address, String name, Uint64 gas) {
        final WasmFunction function = new WasmFunction(FUNC_DELEGATE_CALL_SET_STRING, Arrays.asList(target_address,name,gas), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> delegate_call_set_string(String target_address, String name, Uint64 gas, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_DELEGATE_CALL_SET_STRING, Arrays.asList(target_address,name,gas), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<String> get_string() {
        final WasmFunction function = new WasmFunction(FUNC_GET_STRING, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static ContractDelegateCallStorageString load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new ContractDelegateCallStorageString(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static ContractDelegateCallStorageString load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new ContractDelegateCallStorageString(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
