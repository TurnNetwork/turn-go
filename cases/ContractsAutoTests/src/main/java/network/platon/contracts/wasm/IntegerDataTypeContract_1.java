package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Int16;
import com.platon.rlp.datatypes.Int32;
import com.platon.rlp.datatypes.Int64;
import com.platon.rlp.datatypes.Uint32;
import com.platon.rlp.datatypes.Uint64;
import com.platon.rlp.datatypes.Uint8;
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
public class IntegerDataTypeContract_1 extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001510f60027f7f0060017f0060017f017f60000060037f7f7f0060027f7e0060037f7e7e0060047f7f7f7f0060037f7f7e0060037f7f7f017f60017f017e60037f7e7f006000017f60027f7f017f60017e017f025d0403656e760c706c61746f6e5f70616e6963000303656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000c03656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e00000342410309010302020403030d04040107010003060606030101090202020207000402010204000002000000000400050e0b05020108030a0a02050100010000000001020405017001050505030100020615037f0141808b040b7f0041808b040b7f0041fa0a0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300040b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974001806696e766f6b650037090a010041010b041a3636350ae34a412400100741a40a420037020041ac0a410036020010144101101941b00a10341a410410190ba20a010d7f2002410f6a210f410020026b21072002410e6a210a410120026b210e2002410d6a210d410220026b210c0340200020056a2103200120056a220441037145200220054672450440200320042d00003a0000200f417f6a210f200741016a2107200a417f6a210a200e41016a210e200d417f6a210d200c41016a210c200541016a21050c010b0b200220056b210602400240024002402003410371220b044020064120490d03200b4101460d01200b4102460d02200b4103470d032003200120056a280200220a3a0000200041016a210b200220056b417f6a210c200521030340200c4113494504402003200b6a2208200120036a220941046a2802002206411874200a41087672360200200841046a200941086a2802002204411874200641087672360200200841086a2009410c6a28020022064118742004410876723602002008410c6a200941106a280200220a411874200641087672360200200341106a2103200c41706a210c0c010b0b2002417f6a2007416d2007416d4b1b200f6a4170716b20056b2106200120036a41016a2104200020036a41016a21030c030b2006210403402004411049450440200020056a2203200120056a2202290200370200200341086a200241086a290200370200200541106a2105200441706a21040c010b0b027f2006410871450440200120056a2104200020056a0c010b200020056a2202200120056a2201290200370200200141086a2104200241086a0b21052006410471044020052004280200360200200441046a2104200541046a21050b20064102710440200520042f00003b0000200441026a2104200541026a21050b2006410171450d03200520042d00003a000020000f0b2003200120056a2206280200220a3a0000200341016a200641016a2f00003b0000200041036a210b200220056b417d6a210720052103034020074111494504402003200b6a2208200120036a220941046a2802002206410874200a41187672360200200841046a200941086a2802002204410874200641187672360200200841086a2009410c6a28020022064108742004411876723602002008410c6a200941106a280200220a410874200641187672360200200341106a2103200741706a21070c010b0b2002417d6a200c416f200c416f4b1b200d6a4170716b20056b2106200120036a41036a2104200020036a41036a21030c010b2003200120056a2206280200220d3a0000200341016a200641016a2d00003a0000200041026a210b200220056b417e6a210720052103034020074112494504402003200b6a2208200120036a220941046a2802002206411074200d41107672360200200841046a200941086a2802002204411074200641107672360200200841086a2009410c6a28020022064110742004411076723602002008410c6a200941106a280200220d411074200641107672360200200341106a2103200741706a21070c010b0b2002417e6a200e416e200e416e4b1b200a6a4170716b20056b2106200120036a41026a2104200020036a41026a21030b20064110710440200320042d00003a00002003200428000136000120032004290005370005200320042f000d3b000d200320042d000f3a000f200441106a2104200341106a21030b2006410871044020032004290000370000200441086a2104200341086a21030b2006410471044020032004280000360000200441046a2104200341046a21030b20064102710440200320042f00003b0000200441026a2104200341026a21030b2006410171450d00200320042d00003a00000b20000bc70201027f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122016a22004100360200200041840220016b417c7122026a2201417c6a4100360200024020024109490d002000410036020820004100360204200141786a4100360200200141746a410036020020024119490d002000410036021820004100360214200041003602102000410036020c200141706a41003602002001416c6a4100360200200141686a4100360200200141646a41003602002002200041047141187222026b2101200020026a2100034020014120490d0120004200370300200041186a4200370300200041106a4200370300200041086a4200370300200041206a2100200141606a21010c000b000b0b3501017f230041106b220041808b0436020c418408200028020c41076a41787122003602004180082000360200418c083f003602000b9f0101047f230041106b220224002002200036020c027f02400240024020000440418c08200041086a22014110762200418c082802006a2203360200418408200141840828020022016a41076a4178712204360200200341107420044d0d0120000d020c030b41000c030b418c08200341016a360200200041016a21000b200040000d0010000b20012002410c6a410410051a200141086a0b200241106a24000b2f01027f2000410120001b2100034002402000100822010d004190082802002202450d0020021103000c010b0b20010bbf0301057f024020002001460d00027f02400240200120006b20026b410020024101746b4b044020002001734103712103200020014f0d012003450d0220000c030b20002001200210051a0f0b024020030d002001417f6a21040340200020026a220341037104402002450d052003417f6a200220046a2d00003a00002002417f6a21020c010b0b2000417c6a21042001417c6a2103034020024104490d01200220046a200220036a2802003602002002417c6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200241046a21052002417f73210303400240200120046a2106200020046a2207410371450d0020022004460d03200720062d00003a00002005417f6a2105200341016a2103200441016a21040c010b0b200220046b2101410021000340200141044f0440200020076a200020066a280200360200200041046a21002001417c6a21010c010b0b200020066a210120022003417c2003417c4b1b20056a417c716b20046b2102200020076a0b210003402002450d01200020012d00003a00002002417f6a2102200041016a2100200141016a21010c000b000b0b0a0041940841013602000b0a0041940841003602000b4d01017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b200020012802082001280204100e20000b6401027f2002417049044002402002410a4d0440200020024101743a0000200041016a21030c010b200241106a4170712204100921032000200236020420002004410172360200200020033602080b200320012002100f200220036a41003a00000f0b000b13002002047f20002001200210050520000b1a0b130020002d0000410171044020002802081a0b0b9f0101027f416f20016b41014f0440027f200041016a20002d0000410171450d001a20002802080b2105027f416f200141e6ffffff074b0d001a410b20014101742204200141016a220120012004491b2204410b490d001a200441106a4170710b22011009210420030440200420052003100f0b200220036b22020440200320046a200320056a2002100f0b20002004360208200020014101723602000f0b000bfc0101067f027f410a21010240027f20002d0000220241017145044020024101762103410a0c010b2000280204210320002802002202417e71417f6a0b220520034128200341284b1b2204410b4f0440200441106a417071417f6a21010b20014704402001410a460440200041016a210420002802080c030b200141016a10092204200120054b720d010b0f0b20002d0000220241017145044041012106200041016a0c010b4101210620002802080b2105200420052002410171047f200028020405200241fe01714101760b41016a100f2006044020002004360208200020033602042000200141016a4101723602000f0b200020034101743a00000b950101037f027f20002d0000220241017122040440200028020421032000280200417e71417f6a0c010b20024101762103410a0b2102027f02400240200220034604402000200220022002101120002d0000410171450d010c020b20040d010b2000200341017441026a3a0000200041016a0c010b2000200341016a36020420002802080b20036a220041003a0001200020013a00000b2301017f03402000410c470440200041a40a6a4100360200200041046a21000c010b0b0bd00202027f047e027e027e02400240027e024020025004400c010b41fd00200279a76b220341c000470d0220010c010b2001420a800c030b21072002210541c00021030c010b2003413f4d0440200241c00020036bad22078620012003ad2206888421052002200688210820012007862107420021060c010b200241800120036bad2206862001200341406aad22058884210720022005882105200120068621060b03402003044020084201862005423f8884220220054201862007423f88842201427f852205420a7c200554ad2002427f8542007c7c423f8722024200837d20012002420a83220554ad7d2108200120057d210520074201862006423f888421072003417f6a21032004ad20064201868421062002a741017121040c010b0b2004ad2006420186427e8384210120074201862006423f88840c010b210142000b210220002001370300200020023703080b3701017f230041106b220324002003200120021015200329030021012000200341086a29030037030820002001370300200341106a24000b7701017e20002001427f7e200242767e7c2001422088220242ffffffff0f7e7c200242f6ffffff0f7e200142ffffffff0f83220142f6ffffff0f7e22024220887c22034220887c200142ffffffff0f7e200342ffffffff0f837c22014220887c3703082000200242ffffffff0f832001422086843703000b7601037f100b41980828020021000340200004400340419c08419c082802002201417f6a22023602002001410148450440200020024102746a22004184016a280200200041046a280200100c110100100b41980828020021000c010b0b419c084120360200419808200028020022003602000c010b0b0b900101027f100b419808280200220145044041980841a00836020041a00821010b0240419c0828020022024120460440418402100822010440200110060b2001450d0120014198082802003602004198082001360200419c084100360200410021020b419c08200241016a360200200120024102746a22014184016a4100360200200141046a2000360200100c0f0b100c0b070041a40a10100b780020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000101c20012802044f0d002002410471450440200042003702000c010b10000b024002402002411071450d002000101c20012802044d0d0020024104710d01200042003702000b20000f0b100020000b290002402000280204044020002802002c0000417f4c0d0141010f0b41000f0b2000101d2000101e6a0b240002402000280204450d0020002802002c0000417f4c0d0041000f0b2000102341016a0b8a0301047f0240024020002802040440200010244101210220002802002c00002201417f4c0d010c020b41000f0b200141ff0171220241b7014d0440200241807f6a0f0b02400240200141ff0171220141bf014d04400240200041046a22042802002201200241c97e6a22034d047f100020042802000520010b4102490d0020002802002d00010d0010000b200341054f044010000b20002802002d000145044010000b410021024100210103402001200346450440200028020020016a41016a2d00002002410874722102200141016a21010c010b0b200241384f0d010c020b200141f7014d0440200241c07e6a0f0b0240200041046a22042802002201200241897e6a22034d047f100020042802000520010b4102490d0020002802002d00010d0010000b200341054f044010000b20002802002d000145044010000b410021024100210103402001200346450440200028020020016a41016a2d00002002410874722102200141016a21010c010b0b20024138490d010b200241ff7d490d010b100020020f0b20020b3902017f017e230041306b2201240020012000290200220237031020012002370308200141186a200141086a4114101b101c200141306a24000b5e01027f2000027f027f2001280200220504404100200220036a200128020422014b2001200249720d011a410020012003490d021a200220056a2104200120026b20032003417f461b0c020b41000b210441000b360204200020043602000b2101017f2001101e220220012802044b044010000b200020012001101d200210200b900302097f017e230041406a220324002001280208220520024b0440200341386a20011021200320032903383703182001200341186a101f36020c200341306a20011021410021052001027f410020032802302206450d001a410020032802342208200128020c2207490d001a200820072007417f461b210420060b360210200141146a2004360200200141086a41003602000b200141106a2109200141146a21072001410c6a2106200141086a210803400240200520024f0d002007280200450d00200341306a2001102141002105027f2003280230220a044041002003280234220b20062802002204490d011a200b20046b21052004200a6a0c010b41000b210420072005360200200920043602002003200536022c2003200436022820032003290328370310200341306a20094100200341106a101f102020092003290330220c37020020062006280200200c422088a76a3602002008200828020041016a22053602000c010b0b20032009290200220c3703202003200c3703082000200341086a4114101b1a200341406b24000b4101017f02402000280204450d0020002802002d0000220041bf014d0440200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b4401017f200028020445044010000b0240200028020022012d0000418101470d00200041046a28020041014d047f100020002802000520010b2c00014100480d0010000b0b9f0101037f02402000280204044020001024200028020022022c000022014100480d0120014100470f0b41000f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200041046a28020041014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a200041046a280200200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b2c002000200220016b22021028200028020020002802046a2001200210051a2000200028020420026a3602040b9d0201077f02402001450d002000410c6a2107200041106a2105200041046a21060340200528020022022007280200460d01200241786a28020020014904401000200528020021020b200241786a2203200328020020016b220136020020010d01200520033602002000410120062802002002417c6a28020022016b22021029220341016a20024138491b2204200628020022086a102a2004200120002802006a22046a2004200820016b100a0240200241374d0440200028020020016a200241406a3a00000c010b200341f7016a220441ff014d0440200028020020016a20043a00002000280200200120036a6a210103402002450d02200120023a0000200241087621022001417f6a21010c000b000b10000b410121010c000b000b0b1b00200028020420016a220120002802084b044020002001102b0b0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b0f0020002001102b200020013602040b3901017f200028020820014904402001100822022000280200200028020410051a20002802001a200041086a2001360200200020023602000b0b2500200041011028200028020020002802046a20013a00002000200028020441016a3602040b2b01027f20011029220241b7016a22034180024e044010000b2000200341ff0171102c200020012002102e0b3d002000200028020420026a102a200028020020002802046a417f6a2100034020010440200020013a0000200141087621012000417f6a21000c010b0b0ba00101037f230041106b2202240020012802002103024002400240024020012802042201410146044020032c000022044100480d012000200441ff0171102c0c040b200141374b0d010b200020014180017341ff0171102c0c010b20002001102d0b2002200136020c2002200336020820022002290308370300200020022802002201200120022802046a10262000410010270b200041011027200241106a24000b830101037f02400240200150450440200142ff00560d0120002001a741ff0171102c0c020b2000418001102c0c010b024020011031220241374d0440200020024180017341ff0171102c0c010b20021029220341b7016a22044180024f044010000b2000200441ff0171102c200020022003102e0b20002001200210320b2000410110270b3202017f017e034020002002845045044020024238862000420888842100200141016a2101200242088821020c010b0b20010b5101017e2000200028020420026a102a200028020020002802046a417f6a21000340200120038450450440200020013c0000200342388620014208888421012000417f6a2100200342088821030c010b0b0bda0102057f037e230041206b220324002000103422001012200341186a21020340200341106a2001200710162003200329031022082002290300220910172000200329030020017ca741bc0a6a2c000010132001420956200742005220075020082101200921071b0d000b0240200028020420002d00002202410176200241017122021b2204450d002000280208200041016a20021b220020046a417f6a21020340200020024f0d0120002d00002104200020022d00003a0000200220043a00002002417f6a2102200041016a21000c000b000b200341206a24000b3601017f20004200370200200041086a410036020003402001410c46450440200020016a4100360200200141046a21010c010b0b20000b070041b00a10100b08002000200210330bb50902047f017e23004180016b22002400100410012201100822021002200020013602642000200236026020002000290360370310200041e0006a200041286a200041106a411c101b22014100102202400240200041e0006a10382204500d0041c70a10392004510d0141cc0a10392004510440200041c8006a103a2101200041f8006a4100360200200041f0006a4200370300200041e8006a420037030020004200370360200041e0006a4206103b20002802602102200041e0006a410472103c20012002103d200142061030200128020c200141106a28020047044010000b2001280200200128020410032001103e0c020b41d10a10392004510440200041c8006a103a2101200041f8006a4100360200200041f0006a4200370300200041e8006a420037030020004200370360200041e0006a429003103b20002802602102200041e0006a410472103c20012002103d20014290031030200128020c200141106a28020047044010000b2001280200200128020410032001103e0c020b41d70a10392004510440200041e0006a200141011022200041e0006a102402400240200041e0006a1025450d002000280264450d0020002802602d000041c001490d010b10000b200041c8006a200041e0006a103f200028024c220141024f044010000b20002802482102420021040340200104402001417f6a210120022d0000410174ad42fe01832104200241016a21020c010b0b200041c8006a103a2101200041f8006a4100360200200041f0006a4200370300200041e8006a420037030020004200370360200041e0006a2004103b20002802602102200041e0006a410472103c20012002103d200120041030200128020c200141106a28020047044010000b2001280200200128020410032001103e0c020b41de0a10392004510440200041e0006a200141011022200041e0006a102402400240200041e0006a1025450d002000280264450d0020002802602d000041c001490d010b10000b200041c8006a200041e0006a103f200028024c220141054f044010000b200028024821020340200104402001417f6a210120022d00002003410874722103200241016a21020c010b0b200041c8006a103a2101200041f8006a4100360200200041f0006a4200370300200041e8006a420037030020004200370360200041e0006a2003410174ad2204103b20002802602102200041e0006a410472103c20012002103d200120041030200128020c200141106a28020047044010000b2001280200200128020410032001103e0c020b41e60a10392004510440200042003703402001200041406b104020002903402104200041c8006a103a2101200041f8006a4100360200200041f0006a4200370300200041e8006a420037030020004200370360200041e0006a20044201862204103b20002802602102200041e0006a410472103c20012002103d200120041030200128020c200141106a28020047044010000b2001280200200128020410032001103e0c020b41ee0a103920045104402000410036022420004102360220200020002903203703002001200010410c020b41f40a10392004520d002000410036021c20004103360218200020002903183703082001200041086a10410c010b10000b20004180016a24000b850102027f017e230041106b22012400200010240240024020001025450d002000280204450d0020002802002d000041c001490d010b10000b200141086a2000103f200128020c220041094f044010000b200128020821020340200004402000417f6a210020023100002003420886842103200241016a21020c010b0b200141106a240020030b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b29002000410036020820004200370200200041001042200041146a41003602002000420037020c20000b7502027f017e4101210320014280015a0440034020012004845045044020044238862001420888842101200241016a2102200442088821040c010b0b200241384f047f2002102920026a0520020b41016a21030b200041186a2802000440200041046a104421000b2000200028020020036a3602000bdc0201067f200028020422012000280210220241087641fcffff07716a2103027f2001200028020822054704402001200028021420026a220441087641fcffff07716a280200200441ff07714102746a2106200041146a21042003280200200241ff07714102746a0c010b200041146a210441000b2102034020022006470440200241046a220220032802006b418020470d0120032802042102200341046a21030c010b0b20044100360200200041086a21020340200520016b410275220341034f044020012802001a200041046a2201200128020041046a2201360200200228020021050c010b0b0240200041106a027f2003410147044020034102470d024180080c010b4180040b3602000b03402001200547044020012802001a200141046a21010c010b0b200041086a22032802002101200041046a280200210203402001200247044020032001417c6a22013602000c010b0b20002802001a0b1300200028020820014904402000200110420b0b1c01017f200028020c22010440200041106a20013602000b200010430be60101047f2001101e2204200128020422024b04401000200141046a28020021020b200128020021052000027f024002400240027f0240200204404100210120052c00002203417f4c0d012005450d030c040b41000c010b200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a210120050d010b410021030c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b2b01017f230041206b22022400200241086a2000410110222001200241086a1038370300200241206a24000b840301047f230041f0006b220224002001280200210320012802042101200242003703102000200241106a1040200241186a200241086a20014101756a20022903102003110800200241286a103a2101200241e8006a4100360200200241e0006a4200370300200241d8006a420037030020024200370350410121000240200241406b200241186a100d220428020420022d00402203410176200341017122051b2203450d0002400240200341014604402004280208200441016a20051b2c0000417f4a0d030c010b200341374b0d010b200341016a21000c010b2003102920036a41016a21000b2002200036025020041010200241d0006a410472103c20012000103d2002200241d0006a200241186a100d2200280208200041016a20022d0050220341017122041b36024020022000280204200341017620041b3602442002200229034037030020012002102f20001010200128020c200141106a28020047044010000b2001280200200128020410032001103e200241186a1010200241f0006a24000b3601017f2000280208200149044020011008200028020020002802041005210220001043200041086a2001360200200020023602000b0b080020002802001a0b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0b0b44010041bc0a0b3d3031323334353637383900696e697400696e743800696e7436340075696e7438740075696e743332740075696e74363474007531323874007532353674";

    public static String BINARY = BINARY_0;

    public static final String FUNC_INT32 = "int32";

    public static final String FUNC_INT64 = "int64";

    public static final String FUNC_U256T = "u256t";

    public static final String FUNC_UINT32T = "uint32t";

    public static final String FUNC_UINT64T = "uint64t";

    public static final String FUNC_U128T = "u128t";

    public static final String FUNC_INT8 = "int8";

    public static final String FUNC_UINT8T = "uint8t";

    protected IntegerDataTypeContract_1(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected IntegerDataTypeContract_1(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Int32> int32() {
        final WasmFunction function = new WasmFunction(FUNC_INT32, Arrays.asList(), Int32.class);
        return executeRemoteCall(function, Int32.class);
    }

    public RemoteCall<Int64> int64() {
        final WasmFunction function = new WasmFunction(FUNC_INT64, Arrays.asList(), Int64.class);
        return executeRemoteCall(function, Int64.class);
    }

    public RemoteCall<String> u256t(Uint64 input) {
        final WasmFunction function = new WasmFunction(FUNC_U256T, Arrays.asList(input), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<Uint32> uint32t(Uint32 input) {
        final WasmFunction function = new WasmFunction(FUNC_UINT32T, Arrays.asList(input), Uint32.class);
        return executeRemoteCall(function, Uint32.class);
    }

    public RemoteCall<Uint64> uint64t(Uint64 input) {
        final WasmFunction function = new WasmFunction(FUNC_UINT64T, Arrays.asList(input), Uint64.class);
        return executeRemoteCall(function, Uint64.class);
    }

    public RemoteCall<String> u128t(Uint64 input) {
        final WasmFunction function = new WasmFunction(FUNC_U128T, Arrays.asList(input), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<Int16> int8() {
        final WasmFunction function = new WasmFunction(FUNC_INT8, Arrays.asList(), Int16.class);
        return executeRemoteCall(function, Int16.class);
    }

    public RemoteCall<Uint8> uint8t(Uint8 input) {
        final WasmFunction function = new WasmFunction(FUNC_UINT8T, Arrays.asList(input), Uint8.class);
        return executeRemoteCall(function, Uint8.class);
    }

    public static RemoteCall<IntegerDataTypeContract_1> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(IntegerDataTypeContract_1.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<IntegerDataTypeContract_1> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(IntegerDataTypeContract_1.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<IntegerDataTypeContract_1> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(IntegerDataTypeContract_1.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<IntegerDataTypeContract_1> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(IntegerDataTypeContract_1.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static IntegerDataTypeContract_1 load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new IntegerDataTypeContract_1(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static IntegerDataTypeContract_1 load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new IntegerDataTypeContract_1(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
