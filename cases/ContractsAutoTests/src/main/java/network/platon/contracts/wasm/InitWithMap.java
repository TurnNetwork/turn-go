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
public class InitWithMap extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001420c60027f7f0060017f017f60017f0060027f7f017f60037f7f7f017f60000060037f7f7f0060047f7f7f7f0060047f7f7f7f017f60027f7e006000017f60017f017e02a9010703656e760c706c61746f6e5f70616e6963000503656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000a03656e7610706c61746f6e5f6765745f696e707574000203656e7617706c61746f6e5f6765745f73746174655f6c656e677468000303656e7610706c61746f6e5f6765745f7374617465000803656e7610706c61746f6e5f7365745f7374617465000703656e760d706c61746f6e5f72657475726e00000349480501020202060303030407030102020501040602000b00000101030200000205080100000904000301040100000104000000000700030403000100040802040601050605010107000405017001050505030100020608017f0141d08a040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070f5f5f66756e63735f6f6e5f65786974002606696e766f6b650016090a010041010b04090c0c090aa36448100041980810081a4101100a1048104a0b190020004200370200200041086a41003602002000100b20000b0300010b940101027f41a408410136020041a808280200220145044041a80841b00836020041b00821010b024041ac082802002202412046044041840210172201450d0120011047220141a80828020036020041a808200136020041ac084100360200410021020b41ac08200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41a40841003602000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0b1000200041186a2001100d2002100e1a0b5201037f230041106b2203240020002003410c6a200110102204280200220245044041281013220241106a2001100f1a2002411c6a10081a2000200328020c2004200210110b200341106a24002002411c6a0b8e0201047f20002001470440200128020420012d00002202410176200241017122031b2102200141016a21042001280208410a2101200420031b210420002d0000410171220304402000280200417e71417f6a21010b200220014d0440027f2003044020002802080c010b200041016a0b21012002044020012004200210460b200120026a41003a000020002d000041017104402000200236020420000f0b200020024101743a000020000f0b416f2103200141e6ffffff074d0440410b20014101742201200220022001491b220141106a4170712001410b491b21030b200310132201200420021049200020023602042000200341017236020020002001360208200120026a41003a00000b20000ba10101037f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b20012802082103024020012802042201410a4d0440200020014101743a0000200041016a21020c010b200141106a4170712204101321022000200136020420002004410172360200200020023602080b2002200320011049200120026a41003a000020000b890101027f200041046a2103024020002802042200044002400340024002402002200041106a220410120440200028020022040d012001200036020020000f0b200420021012450d03200041046a210320002802042204450d01200321000b20002103200421000c010b0b2001200036020020030f0b200120003602000c010b200120033602000b20030ba40201027f20032001360208200342003702002002200336020020002802002802002201044020002001360200200228020021030b2003200320002802042205463a000c03400240024020032005460d00200328020822012d000c0d002001200128020822022802002204460440024020022802042204450d0020042d000c0d000c030b20032001280200470440200110142001280208220128020821020b200141013a000c200241003a000c200210150c010b02402004450d0020042d000c0d000c020b20032001280200460440200110152001280208220128020821020b200141013a000c200241003a000c200210140b2000200028020841016a3602080f0b2004410c6a200141013a000c200220022005463a000c41013a0000200221030c000b000bb20101067f02400240200128020420012d00002202410176200241017122031b2205200028020420002d00002202410176200241017122041b2206200520064922071b2202450d002000280208200041016a20041b21002001280208200141016a20031b210103402002450d0120002d0000220320012d00002204460440200141016a2101200041016a21002002417f6a21020c010b0b200320046b22020d010b417f200720062005491b21020b2002411f760b0b002000410120001b10170b5101027f200020002802042201280200220236020420020440200220003602080b200120002802083602082000280208220220022802002000474102746a200136020020002001360208200120003602000b5101027f200020002802002201280204220236020020020440200220003602080b200120002802083602082000280208220220022802002000474102746a200136020020002001360208200120003602040bd50402057f017e230041c0016b22002400100710012201101722021002200041206a200041086a200220011018220341001019200041206a101a024002402000280224450d00200041206a101a0240200028022022012c0000220241004e044020020d010c020b200241807f460d00200241ff0171220441b7014d0440200028022441014d04401000200028022021010b20012d00010d010c020b200441bf014b0d012000280224200241ff017141ca7e6a22024d04401000200028022021010b200120026a2d0000450d010b2000280224450d0020012d000041c001490d010b10000b200041a0016a200041206a101b20002802a401220241094f044010000b20002802a00121010340200204402002417f6a210220013100002005420886842105200141016a21010c010b0b024002402005500d00418008101c200551044020034102101d0c020b418508101c200551044020034103101d0c020b418d08101c2005520d00200041c8006a10082101200041206a200341011019200041206a2001101e200041206a101f200041e8006a200041386a200041d8006a2001100f100d100f1a200041f8006a10202101200041b8016a4100360200200041b0016a4200370300200041a8016a4200370300200042003703a001200041a0016a20004190016a200041e8006a100f102120002802a001210441046a10222001200410232001200041a0016a200041e8006a100f1024200128020c200141106a28020047044010000b200128020020012802041006200128020c22030440200120033602100b10250c010b10000b1026200041c0016a24000b970101047f230041106b220124002001200036020c2000047f41c80a200041086a2202411076220041c80a2802006a220336020041c40a200241c40a28020022026a41076a417871220436020002400240200341107420044d044041c80a200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104103d41086a0541000b200141106a24000b0c00200020012002411c10270bc90202077f017e230041106b220324002001280208220520024b0440200341086a2001102d20012003280208200328020c102e36020c20032001102d410021052001027f410020032802002206450d001a410020032802042208200128020c2207490d001a200820072007417f461b210420060b360210200141146a2004360200200141003602080b200141106a210903402001280214210402402005200249044020040d01410021040b200020092802002004411410271a200341106a24000f0b20032001102d41002104027f410020032802002207450d001a410020032802042208200128020c2206490d001a200820066b2104200620076a0b2105200120043602142001200536021020032009410020052004102e104d20012003290300220a3702102001200128020c200a422088a76a36020c2001200128020841016a22053602080c000b000b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bd40101047f200110282204200128020422024b04401000200128020421020b200128020021052000027f02400240200204404100210120052c00002203417f4a0d01027f200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a21010c010b4101210120050d000c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b6d01037f230041e0006b22022400200241286a10082103200241346a1008210420024101360250200220003602002002200241d0006a3602042002200310292002200410292002101f2002200241d0006a2003100f200241406b2004100f20011106001025200241e0006a24000ba10201057f230041206b22022400024002402000280204044020002802002d000041c001490d010b200241086a10081a0c010b200241186a2000101b2000102821030240024002400240200228021822000440200228021c220420034f0d010b41002100200241106a410036020020024200370308410021040c010b200241106a4100360200200242003703082000200420032003417f461b22046a21052004410a4b0d010b200220044101743a0008200241086a41017221030c010b200441106a4170712206101321032002200436020c20022006410172360208200220033602100b03402000200546450440200320002d00003a0000200341016a2103200041016a21000c010b0b200341003a00000b2001200241086a102a200241206a24000b8f0901107f230041c0016b220424002000420037020420004297a9bef8bdd0fea8413703102000411c6a220d42003702002000200041046a220c3602002000200d360218200441206a102022052000290310102b200528020c200541106a28020047044010000b200041186a210702400240200528020022032005280204220910032208450d002008101321020340200120026a41003a00002008200141016a2201470d000b20032009200220011004417f460d000240200441086a200241016a200120026a2002417f736a10182201280204450d0020012802002d000041c001490d00200441f8006a20014101102c2103200441e8006a20014100102c210e20032802042101200441c4006a21090340200e280204200146410020032802082202200e280208461b0d03200441d0006a20012002411c10272102200441386a1008210620091008210f0240024002402004280254044020042802502d000041bf014b0d010b10002004280254450d010b20042802502d000041c001490d0020044198016a2002102d4102210a200428029c01210103402001044020044100200428029801220b200b2001102e22106a200b45200120104972220b1b360298014100200120106b200b1b2101200a417f6a210a0c010b0b200a450d010b10000b200441b0016a1008210120044198016a20024100101920044198016a2001101e20062001102a20044188016a1008210120044198016a20024101101920044198016a2001101e200f2001102a200720044198016a2006101022022802004504404128101322012004290338370210200141186a200441406b2802003602002006100b200141246a200941086a2802003602002001411c6a2009290200370200200f100b20072004280298012002200110110b20032003280204220120032802086a410020011b22013602042003280200220204402003200236020820012002102e21012003027f200328020422064504404100210241000c010b4100210241002003280208220a2001490d001a200a20012001417f461b210220060b2201ad2002ad42208684370204200341002003280200220620026b2202200220064b1b36020005200341003602080b0c000b000b10000c010b410021080b200528020c22010440200520013602100b024020080d002000280200210102402000280220450d00200028021821022000200d36021820004100360220200028021c2000410036021c410036020820022802042203200220031b2103034020032202450d012001200c470440200241106a200141106a100e21082002411c6a2001411c6a100e1a024020022802082203450440410021030c010b2002200328020022054604402003410036020020032802042205450d012005102f21030c010b200341003602042005450d002005102f21030b200720044198016a20081030210520072004280298012005200210112001103121010c010b0b0340200228020822020d000b200c21010b03402001200c460d0141281013220241106a2203200141106a1032200720044198016a20031030210320072004280298012003200210112001103121010c000b000b200441c0016a240020000b29002000410036020820004200370200200041001033200041146a41003602002000420037020c20000ba10101037f41012103024002400240200128020420012d00002202410176200241017122041b220241014d0440200241016b0d032001280208200141016a20041b2c0000417f4c0d010c030b200241374b0d010b200241016a21030c010b2002103420026a41016a21030b027f200041186a28020022010440200041086a280200200041146a280200200110350c010b20000b2201200128020020036a36020020000bea0101047f230041106b22042400200028020422012000280210220341087641fcffff07716a2102027f410020012000280208460d001a2002280200200341ff07714102746a0b2101200441086a20001036200428020c210303400240200120034604402000410036021420002802082102200028020421010340200220016b41027522034103490d022000200141046a22013602040c000b000b200141046a220120022802006b418020470d0120022802042101200241046a21020c010b0b2003417f6a220241014d04402000418004418008200241016b1b3602100b200020011037200441106a24000b1300200028020820014904402000200110330b0b8f0101047f410121022001280208200141016a20012d0000220441017122051b210302400240024002402001280204200441017620051b2201410146044020032c000022014100480d012000200141ff017110380c040b200141374b0d01200121020b200020024180017341ff017110380c010b200020011039200121020b2000200320024100103a0b20004101103b0bc20502097f027e23004190016b22012400200141206a10202104200141d0006a4100360200200141c8006a4200370300200141406b420037030020014200370338410121022000290310220a4280015a04400340200a200b8450450440200b423886200a42088884210a200341016a2103200b420888210b0c010b0b200341384f047f2003103420036a0520030b41016a21020b20012002360238200141386a410472102220042002102320042000290310102b200428020c200441106a28020047044010000b200428020421082004280200200141086a10202102200141d0006a4100360200200141c8006a4200370300200141406b420037030020014200370338027f200041206a2802004504402001410136023841010c010b2000411c6a2106200141386a4100103c2105200141e4006a210720002802182103037f2003200646047f20054101103c1a200128023805200141d8006a200341106a103220054100103c20014180016a200141d8006a100f1021200141f0006a2007100f10214101103c1a2003103121030c010b0b0b2105200141386a410472102241011013220341fe013a0000200228020c200241106a28020047044010000b2002280204220641016a220720022802084b047f20022007103320022802040520060b20022802006a20034101103d1a2002200228020441016a3602042002200341016a200520036b6a102320022000280220103e21052000411c6a21062000280218210303402003200647044020054102103e2200200141386a200341106a100f10242000200141d8006a2003411c6a100f10242003103121030c010b0b0240200228020c2002280210460440200228020021030c010b100020022802002103200228020c2002280210460d0010000b2008200320022802041005200228020c22000440200220003602100b200428020c22000440200420003602100b20014190016a24000b880101037f41a408410136020041a8082802002100034020000440034041ac0841ac082802002201417f6a2202360200200141014845044041a4084100360200200020024102746a22004184016a280200200041046a28020011020041a408410136020041a80828020021000c010b0b41ac08412036020041a808200028020022003602000c010b0b0b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000104b20024f0d002003410471044010000c010b200042003702000b02402003411071450d002000104b20024d0d0020034104710440100020000f0b200042003702000b20000bff0201037f200028020445044041000f0b2000101a41012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4301017f230041206b22022400200241086a200028020020002802042802001019200241086a2001101e20002802042200200028020041016a360200200241206a24000b5b00024020002d0000410171450440200041003b01000c010b200028020841003a00002000410036020420002d0000410171450d00200041003602000b20002001290200370200200041086a200141086a2802003602002001100b0bbe0202037f027e02402001500440200041800110380c010b20014280015a044020012106034020052006845045044020054238862006420888842106200241016a2102200542088821050c010b0b0240200241384f04402002210403402004044020044108762104200341016a21030c010b0b200341c9004f044010000b2000200341b77f6a41ff017110382000200028020420036a103f200028020420002802006a417f6a21032002210403402004450d02200320043a0000200441087621042003417f6a21030c000b000b200020024180017341ff017110380b2000200028020420026a103f200028020420002802006a417f6a21024200210503402001200584500d02200220013c0000200542388620014208888421012002417f6a2102200542088821050c000b000b20002001a741ff017110380b20004101103b0be70101037f230041106b2204240020004200370200200041086a410036020020012802042103024002402002450440200321020c010b410021022003450d002003210220012802002d000041c001490d00200441086a2001102d20004100200428020c2201200428020822022001102e22032003417f461b20024520012003497222031b220536020820004100200220031b3602042000200120056b3602000c010b20012802002103200128020421012000410036020020004100200220016b20034520022001497222021b36020820004100200120036a20021b3602040b200441106a240020000b2101017f20011028220220012802044b044010000b200020012001104c2002104d0b2301017f230041206b22022400200241086a2000200141141027104b200241206a24000b1d01017f03402000220128020022000d00200128020422000d000b20010b5c01017f0240200028020422030440034002402002200341106a1012044020032802002200450d040c010b200328020422000d0020012003360200200341046a0f0b200021030c000b000b200041046a21030b2001200336020020030b3601017f024020002802042201044003402001220028020022010d000c020b000b0340200020002802082200280200470d000b0b20000b160020002001100f1a2000410c6a2001410c6a100f1a0b2f01017f200028020820014904402001101720002802002000280204103d210220002001360208200020023602000b0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b25002000200120026a417f6a220141087641fcffff07716a280200200141ff07714102746a0b4f01037f20012802042203200128021020012802146a220441087641fcffff07716a21022000027f410020032001280208460d001a2002280200200441ff07714102746a0b360204200020023602000b2501017f200028020821020340200120024645044020002002417c6a22023602080c010b0b0b250020004101104e200028020020002802046a20013a00002000200028020441016a3602040b5e01027f20011034220241b7016a22034180024e044010000b2000200341ff017110382000200028020420026a103f200028020420002802006a417f6a2100034020010440200020013a0000200141087621012000417f6a21000c010b0b0b2d0020002002104e200028020020002802046a20012002103d1a2000200028020420026a36020420002003103b0b820201047f02402001450d00034020002802102202200028020c460d01200241786a28020020014904401000200028021021020b200241786a2203200328020020016b220136020020010d012000200336021020004101200028020422042002417c6a28020022016b22021034220341016a20024138491b220520046a103f200120002802006a220420056a2004200210460240200241374d0440200028020020016a200241406a3a00000c010b200341f7016a220441ff014d0440200028020020016a20043a00002000280200200120036a6a210103402002450d02200120023a0000200241087621022001417f6a21010c000b000b10000b410121010c000b000b0bbf0c02077f027e230041306b22052400200041046a2107027f20014101460440200041086a280200200041146a280200200041186a220228020022041035280200210120022004417f6a360200200710404180104f044020072000410c6a280200417c6a10370b200141384f047f2001103420016a0520010b41016a2101200041186a28020022020440200041086a280200200041146a280200200210350c020b20000c010b0240200710400d00200041146a28020022014180084f0440200020014180786a360214200041086a2201280200220228020021042001200241046a360200200520043602182007200541186a10410c010b2000410c6a2802002202200041086a2802006b4102752204200041106a2203280200220620002802046b220141027549044041802010132104200220064704400240200028020c220120002802102206470d0020002802082202200028020422034b04402000200220012002200220036b41027541016a417e6d41027422036a1042220136020c2000200028020820036a3602080c010b200541186a200620036b2201410175410120011b22012001410276200041106a10432102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021044200028020c21010b200120043602002000200028020c41046a36020c0c020b02402000280208220120002802042206470d00200028020c2202200028021022034904402000200120022002200320026b41027541016a41026d41027422036a104522013602082000200028020c20036a36020c0c010b200541186a200320066b2201410175410120011b2201200141036a410276200041106a10432102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021044200028020821010b2001417c6a2004360200200020002802082201417c6a22023602082002280200210220002001360208200520023602182007200541186a10410c010b20052001410175410120011b200420031043210241802010132106024020022802082201200228020c2208470d0020022802042204200228020022034b04402002200420012004200420036b41027541016a417e6d41027422036a104222013602082002200228020420036a3602040c010b200541186a200820036b2201410175410120011b22012001410276200241106a280200104321042002280208210320022802042101034020012003470440200428020820012802003602002004200428020841046a360208200141046a21010c010b0b20022902002109200220042902003702002004200937020020022902082109200220042902083702082004200937020820041044200228020821010b200120063602002002200228020841046a360208200028020c2104034020002802082004460440200028020421012000200228020036020420022001360200200228020421012002200436020420002001360208200029020c21092000200229020837020c2002200937020820021044052004417c6a210402402002280204220120022802002208470d0020022802082203200228020c22064904402002200120032003200620036b41027541016a41026d41027422066a104522013602042002200228020820066a3602080c010b200541186a200620086b2201410175410120011b2201200141036a410276200228021010432002280208210620022802042101034020012006470440200528022020012802003602002005200528022041046a360220200141046a21010c010b0b20022902002109200220052903183702002002290208210a20022005290320370208200520093703182005200a3703201044200228020421010b2001417c6a200428020036020020022002280204417c6a3602040c010b0b0b200541186a20071036200528021c410036020041012101200041186a0b2202200228020020016a360200200541306a240020000bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000b840201057f2001044020002802042104200041106a2802002202200041146a280200220349044020022001ad2004ad422086843702002000200028021041086a36021020000f0b027f41002002200028020c22026b410375220541016a2206200320026b2202410275220320032006491b41ffffffff01200241037541ffffffff00491b2203450d001a200341037410130b2102200220054103746a22052001ad2004ad4220868437020020052000280210200028020c22066b22016b2104200141014e0440200420062001103d1a0b2000200220034103746a3602142000200541086a3602102000200436020c20000f0b200041c00110382000410041004101103a20000b0f00200020011033200020013602040b2801017f200028020820002802046b2201410874417f6a410020011b200028021420002802106a6b0ba10202057f017e230041206b22052400024020002802082202200028020c2206470d0020002802042203200028020022044b04402000200320022003200320046b41027541016a417e6d41027422046a104222023602082000200028020420046a3602040c010b200541086a200620046b2202410175410120021b220220024102762000410c6a10432103200028020821042000280204210203402002200446450440200328020820022802003602002003200328020841046a360208200241046a21020c010b0b20002902002107200020032902003702002003200737020020002902082107200020032902083702082003200737020820031044200028020821020b200220012802003602002000200028020841046a360208200541206a24000b2501017f200120006b220141027521032001044020022000200110460b200220034102746a0b4f01017f2000410036020c200041106a2003360200200104402001410274101321040b200020043602002000200420024102746a22023602082000200420014102746a36020c2000200236020420000b2b01027f200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b0b1b00200120006b22010440200220016b22022000200110460b20020b8d0301037f024020002001460d00200120006b20026b410020024101746b4d0440200020012002103d1a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b3501017f230041106b220041d08a0436020c41c00a200028020c41076a417871220036020041c40a200036020041c80a3f003602000b100020020440200020012002103d1a0b0b3801017f41b40a420037020041bc0a410036020041742100034020000440200041c00a6a4100360200200041046a21000c010b0b4104100a0b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f2000104c200010286a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b1b00200028020420016a220120002802084b04402000200110330b0b0b1b01004180080b14696e6974007365745f6d6170006765745f6d6170";

    public static String BINARY = BINARY_0;

    public static final String FUNC_SET_MAP = "set_map";

    public static final String FUNC_GET_MAP = "get_map";

    protected InitWithMap(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected InitWithMap(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<InitWithMap> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, String key, String value) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList(key,value));
        return deployRemoteCall(InitWithMap.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<InitWithMap> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, String key, String value) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList(key,value));
        return deployRemoteCall(InitWithMap.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<InitWithMap> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue, String key, String value) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList(key,value));
        return deployRemoteCall(InitWithMap.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<InitWithMap> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue, String key, String value) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList(key,value));
        return deployRemoteCall(InitWithMap.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<TransactionReceipt> set_map(String key, String value) {
        final WasmFunction function = new WasmFunction(FUNC_SET_MAP, Arrays.asList(key,value), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> set_map(String key, String value, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SET_MAP, Arrays.asList(key,value), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<String> get_map(String key) {
        final WasmFunction function = new WasmFunction(FUNC_GET_MAP, Arrays.asList(key), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static InitWithMap load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new InitWithMap(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static InitWithMap load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new InitWithMap(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
