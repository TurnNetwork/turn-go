package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Uint32;
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
public class ReferenceDataTypeArrayFuncContract extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001480d60027f7f0060017f0060017f017f60000060037f7f7f017f60037f7f7f0060027f7f017f60047f7f7f7f017f60047f7f7f7f0060027f7e006000017f60027f7e017f60017f017e02a9010703656e760c706c61746f6e5f70616e6963000303656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000a03656e7610706c61746f6e5f6765745f696e707574000103656e7617706c61746f6e5f6765745f73746174655f6c656e677468000603656e7610706c61746f6e5f6765745f7374617465000703656e7610706c61746f6e5f7365745f7374617465000803656e760d706c61746f6e5f72657475726e000003424103020101010100060000030204050102000c0102020901000b010100030702060109020006000002040000000004000000000102000407010405030505030202080405017001050505030100020608017f0141a08b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070f5f5f66756e63735f6f6e5f65786974002306696e766f6b650011090a010041010b0409090c090aa75841100041e40810081a4101100a104110440b190020004200370200200041086a41003602002000100b20000b0300010b970101027f41f008410136020041f408280200220145044041f40841fc0836020041fc0821010b024041f8082802002202412046044041840210122201450d0120014184021026220141f40828020036020041f408200136020041f8084100360200410021020b41f808200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41f00841003602000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0b390020004180016a418008100d2000418c016a418208100d20004198016a418408100d200041a4016a418608100d200041b0016a418808100d0b7e01027f20012102024003402002410371044020022d0000450d02200241016a21020c010b0b2002417c6a21020340200241046a22022802002203417f73200341fffdfb776a7141808182847871450d000b0340200341ff0171450d01200241016a2d00002103200241016a21020c000b000b20002001200220016b10430ba10101037f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b20012802082103024020012802042201410a4d0440200020014101743a0000200041016a21020c010b200141106a4170712204102921022000200136020420002004410172360200200020023602080b2002200320011042200120026a41003a000020000b2f01017f20004180016a2102410021000340200041f80046450440200020026a200110102000410c6a21000c010b0b0b3401017f2000200147044020002001280208200141016a20012d0000220041017122021b2001280204200041017620021b10430b0bbe0502047f017e230041e0026b22002400100710012201101222021002200041d8006a200041086a200220011013220341001014200041d8006a101502400240200041d8006a1016450d00200028025c450d0020002802582d000041c001490d010b10000b200041386a200041d8006a1017200028023c220141094f044010000b200028023821020340200104402001417f6a210120023100002004420886842104200241016a21020c010b0b024002402004500d00418a0810182004510440410210190c020b418f0810182004510440410310190c020b41a00810182004510440200041d8006a101a200041206a101b2101200041d0006a4100360200200041c8006a4200370300200041406b420037030020004200370338200041386a4200101c20002802382103200041386a410472101d20012003101e20014200101f220128020c200141106a28020047044010000b200128020020012802041006200128020c22030440200120033602100b10200c020b41b00810182004510440200041d8006a200341011014200041d8006a101502400240200041d8006a1016450d00200028025c450d0020002802582d000041c001490d010b10000b200041386a200041d8006a1017200028023c220141054f044010000b41002103200028023821020340200104402001417f6a210120022d00002003410874722103200241016a21020c010b0b200041d8006a101a20002003360220200041386a2000200041206a280200410c6c6a41d8016a100e1a200041386a102110200c020b41c30810182004510440200041d8006a101a200041386a200041d8016a100e102110200c020b41d60810182004520d00200041d0026a10082101200041d8006a200341011014200041d8006a20011022200041d8006a101a200041d8006a200041386a200041206a2001100e100e100f10200c010b10000b1023200041e0026a24000b970101047f230041106b220124002001200036020c2000047f41940b200041086a2202411076220041940b2802006a220336020041900b200241900b28020022026a41076a417871220436020002400240200341107420044d044041940b200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104103441086a0541000b200141106a24000b0c00200020012002411c10240bc90202077f017e230041106b220324002001280208220520024b0440200341086a2001102a20012003280208200328020c102b36020c20032001102a410021052001027f410020032802002206450d001a410020032802042208200128020c2207490d001a200820072007417f461b210420060b360210200141146a2004360200200141003602080b200141106a210903402001280214210402402005200249044020040d01410021040b200020092802002004411410241a200341106a24000f0b20032001102a41002104027f410020032802002207450d001a410020032802042208200128020c2206490d001a200820066b2104200620076a0b2105200120043602142001200536021020032009410020052004102b104720012003290300220a3702102001200128020c200a422088a76a36020c2001200128020841016a22053602080c000b000b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0b980101037f200028020445044041000f0b20001015200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0bd40101047f200110252204200128020422024b04401000200128020421020b200128020021052000027f02400240200204404100210120052c00002203417f4a0d01027f200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a21010c010b4101210120050d000c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b2901027f23004180026b22012400200141086a101a200141086a2000110100102020014180026a24000bf30301087f230041e0006b22022400200041f8001026220510272005428aa4bedd86c3aeee5837037820054180016a1027200241206a101b220620052903781028200628020c200641106a28020047044010000b0240200628020022042006280204220010032208450d002008102921030340200120036a41003a00002008200141016a2201470d000b20042000200320011004417f460440410021010c010b0240024002400240200241086a200341016a200120036a2003417f736a10132207280204450d0020072802002d000041c001490d00200241386a21010c010b1000200241386a21012007280204450d010b20072802002d000041c001490d0020012007102a20012802042101410a2103034020010440200241002002280238220020002001102b22046a20004520012004497222001b3602384100200120046b20001b21012003417f6a21030c010b0b2003450d010b10000b20054180016a21004100210103402001410a46450440200241d0006a10082104200241386a200720011014200241386a2004102220002004102c2000410c6a2100200141016a21010c010b0b200821010b200628020c22000440200620003602100b024020010d00410021010340200141f800460d01200120056a22004180016a200010102001410c6a21010c000b000b200241e0006a240020050b2900200041003602082000420037020020004100102d200041146a41003602002000420037020c20000b840102027f017e4101210320014280015a0440034020012004845045044020044238862001420888842101200241016a2102200442088821040c010b0b200241384f047f2002102e20026a0520020b41016a21030b200041186a28020022020440200041086a280200200041146a2802002002102f21000b2000200028020020036a3602000bea0101047f230041106b22042400200028020422012000280210220341087641fcffff07716a2102027f410020012000280208460d001a2002280200200341ff07714102746a0b2101200441086a20001030200428020c210303400240200120034604402000410036021420002802082102200028020421010340200220016b41027522034103490d022000200141046a22013602040c000b000b200141046a220120022802006b418020470d0120022802042101200241046a21020c010b0b2003417f6a220241014d04402000418004418008200241016b1b3602100b200020011031200441106a24000b13002000280208200149044020002001102d0b0bba0202037f037e02402001500440200041800110360c010b20014280015a044020012107034020062007845045044020064238862007420888842107200241016a2102200642088821060c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff017110362000200028020420046a1037200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff017110360b2000200028020420026a1037200028020420002802006a417f6a210203402001200584500d02200220013c0000200542388620014208888421012002417f6a2102200542088821050c000b000b20002001a741ff017110360b2000103920000bdf05010b7f230041e0006b22032400200341186a101b2106200341c8006a22024100360200200341406b22044200370300200341386a2205420037030020034200370330200341306a2000290378101c20032802302101200341306a410472101d20062001101e200620002903781028200628020c200641106a28020047044010000b2006280204210920062802002003101b210120024100360200200442003703002005420037030020034200370330200341306a4100103241800121020340200241f801470440200341306a200341d0006a200020026a100e10332002410c6a21020c010b0b200341306a4101103220032802302104200341306a410472101d41011029220241fe013a0000200128020c200141106a28020047044010000b2001280204220541016a220720012802084b047f20012007102d20012802040520050b20012802006a2002410110341a2001200128020441016a3602042001200241016a200420026b6a101e20012802042102024020012802102204200141146a280200220549044020042002ad422086420a843702002001200128021041086a3602100c010b027f41002004200128020c22046b410375220741016a2208200520046b2204410275220520052008491b41ffffffff01200441037541ffffffff00491b2204450d001a200441037410290b2105200520074103746a22072002ad422086420a8437020020072001280210200128020c220b6b22026b2108200241014e04402008200b200210341a0b2001200520044103746a3602142001200741086a3602102001200836020c0b41800121020340200241f8014704402001200341306a200020026a100e10352002410c6a21020c010b0b0240200128020c2001280210460440200128020021020c010b100020012802002102200128020c2001280210460d0010000b2009200220012802041005200128020c22000440200120003602100b200628020c22000440200620003602100b200341e0006a24000bab0101037f230041e0006b22012400200141186a101b2102200141d8006a4100360200200141d0006a4200370300200141c8006a420037030020014200370340200141406b200141306a2000100e103320012802402103200141406b410472101d20022003101e2002200141086a2000100e1035200228020c200241106a28020047044010000b200228020020022802041006200228020c22000440200220003602100b200141e0006a24000ba10201057f230041206b22022400024002402000280204044020002802002d000041c001490d010b200241086a10081a0c010b200241186a200010172000102521030240024002400240200228021822000440200228021c220420034f0d010b41002100200241106a410036020020024200370308410021040c010b200241106a4100360200200242003703082000200420032003417f461b22046a21052004410a4b0d010b200220044101743a0008200241086a41017221030c010b200441106a4170712206102921032002200436020c20022006410172360208200220033602100b03402000200546450440200320002d00003a0000200341016a2103200041016a21000c010b0b200341003a00000b2001200241086a102c200241206a24000b880101037f41f008410136020041f4082802002100034020000440034041f80841f8082802002201417f6a2202360200200141014845044041f0084100360200200020024102746a22004184016a280200200041046a28020011010041f008410136020041f40828020021000c010b0b41f808412036020041f408200028020022003602000c010b0b0b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000104520024f0d002003410471044010000c010b200042003702000b02402003411071450d002000104520024d0d0020034104710440100020000f0b200042003702000b20000bff0201037f200028020445044041000f0b2000101541012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020be10201027f02402001450d00200041003a0000200020016a2202417f6a41003a000020014103490d00200041003a0002200041003a00012002417d6a41003a00002002417e6a41003a000020014107490d00200041003a00032002417c6a41003a000020014109490d002000410020006b41037122036a220241003602002002200120036b417c7122036a2201417c6a410036020020034109490d002002410036020820024100360204200141786a4100360200200141746a410036020020034119490d002002410036021820024100360214200241003602102002410036020c200141706a41003602002001416c6a4100360200200141686a4100360200200141646a41003602002003200241047141187222036b2101200220036a2102034020014120490d0120024200370300200241186a4200370300200241106a4200370300200241086a4200370300200241206a2102200141606a21010c000b000b20000b2201027f200041f8006a21010340200010082000410c6a2100410c6a2001470d000b0b090020002001101f1a0b0b002000410120001b10120b2101017f20011025220220012802044b044010000b2000200120011046200210470b2301017f230041206b22022400200241086a20002001411410241045200241206a24000b5b00024020002d0000410171450440200041003b01000c010b200028020841003a00002000410036020420002d0000410171450d00200041003602000b20002001290200370200200041086a200141086a2802003602002001100b0b2f01017f2000280208200149044020011012200028020020002802041034210220002001360208200020023602000b0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b25002000200120026a417f6a220141087641fcffff07716a280200200141ff07714102746a0b4f01037f20012802042203200128021020012802146a220441087641fcffff07716a21022000027f410020032001280208460d001a2002280200200441ff07714102746a0b360204200020023602000b2501017f200028020821020340200120024645044020002002417c6a22023602080c010b0b0bbd0c02077f027e230041306b22052400200041046a2107024020014101460440200041086a280200200041146a280200200041186a22022802002204102f280200210120022004417f6a3602002007103a4180104f044020072000410c6a280200417c6a10310b200141384f047f2001102e20016a0520010b41016a2101200041186a2802002202450d01200041086a280200200041146a2802002002102f21000c010b02402007103a0d00200041146a28020022014180084f0440200020014180786a360214200041086a2201280200220228020021042001200241046a360200200520043602182007200541186a103b0c010b2000410c6a2802002202200041086a2802006b4102752204200041106a2203280200220620002802046b220141027549044041802010292104200220064704400240200028020c220120002802102206470d0020002802082202200028020422034b04402000200220012002200220036b41027541016a417e6d41027422036a103c220136020c2000200028020820036a3602080c010b200541186a200620036b2201410175410120011b22012001410276200041106a103d2102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c200220093702082002103e200028020c21010b200120043602002000200028020c41046a36020c0c020b02402000280208220120002802042206470d00200028020c2202200028021022034904402000200120022002200320026b41027541016a41026d41027422036a103f22013602082000200028020c20036a36020c0c010b200541186a200320066b2201410175410120011b2201200141036a410276200041106a103d2102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c200220093702082002103e200028020821010b2001417c6a2004360200200020002802082201417c6a22023602082002280200210220002001360208200520023602182007200541186a103b0c010b20052001410175410120011b20042003103d210241802010292106024020022802082201200228020c2208470d0020022802042204200228020022034b04402002200420012004200420036b41027541016a417e6d41027422036a103c22013602082002200228020420036a3602040c010b200541186a200820036b2201410175410120011b22012001410276200241106a280200103d21042002280208210320022802042101034020012003470440200428020820012802003602002004200428020841046a360208200141046a21010c010b0b2002290200210920022004290200370200200420093702002002290208210920022004290208370208200420093702082004103e200228020821010b200120063602002002200228020841046a360208200028020c2104034020002802082004460440200028020421012000200228020036020420022001360200200228020421012002200436020420002001360208200029020c21092000200229020837020c200220093702082002103e052004417c6a210402402002280204220120022802002208470d0020022802082203200228020c22064904402002200120032003200620036b41027541016a41026d41027422066a103f22013602042002200228020820066a3602080c010b200541186a200620086b2201410175410120011b2201200141036a4102762002280210103d2002280208210620022802042101034020012006470440200528022020012802003602002005200528022041046a360220200141046a21010c010b0b20022902002109200220052903183702002002290208210a20022005290320370208200520093703182005200a370320103e200228020421010b2001417c6a200428020036020020022002280204417c6a3602040c010b0b0b200541186a20071030200528021c4100360200200041186a2100410121010b2000200028020020016a360200200541306a24000b9a0101037f41012103024002400240200128020420012d00002202410176200241017122041b220241014d0440200241016b0d032001280208200141016a20041b2c0000417f4c0d010c030b200241374b0d010b200241016a21030c010b2002102e20026a41016a21030b200041186a28020022010440200041086a280200200041146a2802002001102f21000b2000200028020020036a3602000bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000b810201047f410121022001280208200141016a20012d0000220341017122051b210402400240024002402001280204200341017620051b2203410146044020042c000022014100480d012000200141ff017110360c040b200341374b0d01200321020b200020024180017341ff017110360c010b2003102e220141b7016a22024180024e044010000b2000200241ff017110362000200028020420016a1037200028020420002802006a417f6a210220032101037f2001047f200220013a0000200141087621012002417f6a21020c010520030b0b21020b200020021038200028020020002802046a2004200210341a2000200028020420026a3602040b200010390b2500200041011038200028020020002802046a20013a00002000200028020441016a3602040b0f0020002001102d200020013602040b1b00200028020420016a220120002802084b044020002001102d0b0bf80101057f0340024020002802102201200028020c460d00200141786a28020041014904401000200028021021010b200141786a2202200228020041016b220436020020040d002000200236021020004101200028020422032001417c6a28020022026b2201102e220441016a20014138491b220520036a1037200220002802006a220320056a2003200110400240200141374d0440200028020020026a200141406a3a00000c010b200441f7016a220341ff014d0440200028020020026a20033a00002000280200200220046a6a210203402001450d02200220013a0000200141087621012002417f6a21020c000b000b10000b0c010b0b0b2801017f200028020820002802046b2201410874417f6a410020011b200028021420002802106a6b0ba10202057f017e230041206b22052400024020002802082202200028020c2206470d0020002802042203200028020022044b04402000200320022003200320046b41027541016a417e6d41027422046a103c22023602082000200028020420046a3602040c010b200541086a200620046b2202410175410120021b220220024102762000410c6a103d2103200028020821042000280204210203402002200446450440200328020820022802003602002003200328020841046a360208200241046a21020c010b0b2000290200210720002003290200370200200320073702002000290208210720002003290208370208200320073702082003103e200028020821020b200220012802003602002000200028020841046a360208200541206a24000b2501017f200120006b220141027521032001044020022000200110400b200220034102746a0b4f01017f2000410036020c200041106a2003360200200104402001410274102921040b200020043602002000200420024102746a22023602082000200420014102746a36020c2000200236020420000b2b01027f200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b0b1b00200120006b22010440200220016b22022000200110400b20020b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210341a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041a08b0436020c418c0b200028020c41076a417871220036020041900b200036020041940b3f003602000b10002002044020002001200210341a0b0bd40101027f410a210320002d0000410171220404402000280200417e71417f6a21030b200320024f0440027f2004044020002802080c010b200041016a0b21032002044020032001200210400b200220036a41003a000020002d00004101710440200020023602040f0b200020024101743a00000f0b416f2104200341e6ffffff074d0440410b200341017422032002200320024b1b220341106a4170712003410b491b21040b200410292203200120021042200020023602042000200441017236020020002003360208200220036a41003a00000b3801017f41800b420037020041880b4100360200417421000340200004402000418c0b6a4100360200200041046a21000c010b0b4104100a0b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001046200010256a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b0b6901004180080b6261006200630064006500696e697400736574496e69744172726179446174650067657441727261794973456d70747900676574417272617956616c7565496e646578006765744172726179466972737456616c756500736574417272617946696c6c";

    public static String BINARY = BINARY_0;

    public static final String FUNC_SETINITARRAYDATE = "setInitArrayDate";

    public static final String FUNC_GETARRAYFIRSTVALUE = "getArrayFirstValue";

    public static final String FUNC_GETARRAYISEMPTY = "getArrayIsEmpty";

    public static final String FUNC_GETARRAYVALUEINDEX = "getArrayValueIndex";

    public static final String FUNC_SETARRAYFILL = "setArrayFill";

    protected ReferenceDataTypeArrayFuncContract(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ReferenceDataTypeArrayFuncContract(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<ReferenceDataTypeArrayFuncContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeArrayFuncContract.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ReferenceDataTypeArrayFuncContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeArrayFuncContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ReferenceDataTypeArrayFuncContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeArrayFuncContract.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<ReferenceDataTypeArrayFuncContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeArrayFuncContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<TransactionReceipt> setInitArrayDate() {
        final WasmFunction function = new WasmFunction(FUNC_SETINITARRAYDATE, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> setInitArrayDate(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SETINITARRAYDATE, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<String> getArrayFirstValue() {
        final WasmFunction function = new WasmFunction(FUNC_GETARRAYFIRSTVALUE, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<Boolean> getArrayIsEmpty() {
        final WasmFunction function = new WasmFunction(FUNC_GETARRAYISEMPTY, Arrays.asList(), Boolean.class);
        return executeRemoteCall(function, Boolean.class);
    }

    public RemoteCall<String> getArrayValueIndex(Uint32 index) {
        final WasmFunction function = new WasmFunction(FUNC_GETARRAYVALUEINDEX, Arrays.asList(index), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<TransactionReceipt> setArrayFill(String str) {
        final WasmFunction function = new WasmFunction(FUNC_SETARRAYFILL, Arrays.asList(str), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> setArrayFill(String str, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SETARRAYFILL, Arrays.asList(str), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static ReferenceDataTypeArrayFuncContract load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new ReferenceDataTypeArrayFuncContract(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static ReferenceDataTypeArrayFuncContract load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new ReferenceDataTypeArrayFuncContract(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
