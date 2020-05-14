package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Uint64;
import com.platon.rlp.datatypes.Uint8;
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
public class ReferenceDataTypeSetContract extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001420c60017f0060027f7f0060017f017f60037f7f7f017f60000060037f7f7f0060047f7f7f7f0060027f7e0060027f7f017f60047f7f7f7f017f6000017f60017f017e02a9010703656e760c706c61746f6e5f70616e6963000403656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000a03656e7610706c61746f6e5f6765745f696e707574000003656e760d706c61746f6e5f72657475726e000103656e7617706c61746f6e5f6765745f73746174655f6c656e677468000803656e7610706c61746f6e5f6765745f7374617465000903656e7610706c61746f6e5f7365745f737461746500060358570400000002020306000501000202080102000000030b00020001020401020201070001000107000100020201010200030100020103090003040302050402020004000400030202020000080603010502000101010101070405017001090905030100020608017f0141a08b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070f5f5f66756e63735f6f6e5f65786974004706696e766f6b650022090e010041010b08090a1112161a084a0ae26a570800103f104310490b070041e00810460b0300010be60101057f230041106b22022400200241850c3b000620024181848c20360002200041186a21040240200041206a280200450440200241026a21000c010b2004100b210103400240200241026a20036a21002001452003410646720d00200120002d00003a000d2001100c2004200241086a2001410d6a100d22052802004504402004200228020820052001100e0b200341016a210321010c010b0b2001450d000340200128020822030440200321010c010b0b2001100f0b200241086a210103402000200146450440200241086a200420001010200041016a21000c010b0b200241106a24000b3901027f200028020021012000200041046a36020020004100360208200028020420004100360204410036020820012802042200200120001b0b4c01027f2000280208220145044041000f0b02402000200128020022024604402001410036020020012802042200450d01200010340f0b200141003602042002450d002002103421010b20010b8d0101027f200041046a2103024020002802042200044020022d000021040240034002400240200420002d000d2202490440200028020022020d012001200036020020000f0b200220044f0d03200041046a210320002802042202450d01200321000b20002103200221000c010b0b2001200036020020030f0b200120003602000c010b200120033602000b20030ba40201027f20032001360208200342003702002002200336020020002802002802002201044020002001360200200228020021030b2003200320002802042205463a000c03400240024020032005460d00200328020822012d000c0d002001200128020822022802002204460440024020022802042204450d0020042d000c0d000c030b20032001280200470440200110182001280208220128020821020b200141013a000c200241003a000c200210190c010b02402004450d0020042d000c0d000c020b20032001280200460440200110192001280208220128020821020b200141013a000c200241003a000c200210180b2000200028020841016a3602080f0b2004410c6a200141013a000c200220022005463a000c41013a0000200221030c000b000b1500200004402000280200100f2000280204100f0b0b7601037f230041106b22032400200020012003410c6a2002100d22052802002204047f410005411010452104200341086a41013a0000200420022d00003a000d2003200141046a3602042001200328020c20052004100e200341003602002003103541010b3a000420002004360200200341106a24000b2201017f230041106b22022400200241086a200041186a20011010200241106a24000b2501017f2000411c6a2101200028021821000340200020014704402000101321000c010b0b0b2d01017f02402000280204220104402001101721000c010b0340200020002802082200280200470d000b0b20000b4001027f230041106b22012400200141033a000f027f4100200041186a2001410f6a101522022000411c6a460d001a20022d000d0b200141106a240041ff01710b6101057f20012d000022042105200041046a22012103200121000340200328020022020440200241046a200220022d000d20054922061b21032000200220061b21000c010b0b024020002001470440200420002d000d4f0d010b200121000b20000bc40601057f200041186a22022001101522032000411c6a470440200320022802004604402002200310133602000b200041206a22012001280200417f6a3602002000411c6a2802002101027f0240024020032802002205450440200321020c010b20032802042200450440200321020c020b20001017220228020022050d010b200228020422050d004100210541000c010b2005200228020836020841010b210602402002200228020822042802002200460440200420053602002001200246044041002100200521010c020b200428020421000c010b200420053602040b024020022d000c4520022003470440200220032802082204360208200420032802082802002003474102746a20023602002003280200220420023602082002200436020020022003280204220436020420040440200420023602080b200220032d000c3a000c2002200120012003461b21010b200145720d002006450440034020002d000c21030240200020002802082202280200470440024002402003450440200041013a000c200241003a000c2002101820002001200120002802002200461b2101200028020421000b20002802002203044020032d000c450d010b20002802042202044020022d000c450d020b200041003a000c0240200120002802082200460440200121000c010b20002d000c0d040b200041013a000c0c060b20002802042202044020022d000c450d010b200341013a000c200041003a000c200010192000280208220028020421020b2000200028020822002d000c3a000c200041013a000c200241013a000c200010180c040b02402003450440200041013a000c200241003a000c2002101920002001200120002802042200461b2101200028020021000b20002802002202044020022d000c450d010b024020002802042203044020032d000c450d010b200041003a000c20012000280208220047044020002d000c0d030b200041013a000c0c050b2002044020022d000c450d010b200341013a000c200041003a000c200010182000280208220028020021020b2000200028020822002d000c3a000c200041013a000c200241013a000c200010190c030b2000280208220320032802002000464102746a28020021000c000b000b200541013a000c0b0b0b1401017f03402000220128020022000d000b20010b5101027f200020002802042201280200220236020420020440200220003602080b200120002802083602082000280208220220022802002000474102746a200136020020002001360208200120003602000b5101027f200020002802002201280204220236020020020440200220003602080b200120002802083602082000280208220220022802002000474102746a200136020020002001360208200120003602040b1e01017f2000411c6a2201280200100f20002001360218200142003702000b3401017f230041106b220324002003200236020c200320013602082003200329030837030020002003411c104b200341106a24000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b5101047f230041306b22012400200141086a200028020422034101756a210220002802002100200141086a101e20022003410171047f200228020020006a2802000520000b110000101f200141306a24000bb606020c7f017e230041a0016b2201240020004200370204200042c2c285f3c4ee90c6e4003703102000411c6a220342003702002000200041046a220a36020020002003360218200141386a1024220520002903101027200528020c200541106a28020047044010000b200041186a2106024002402005280200220b2005280204220c10042203450440410021030c010b20014100360230200142003703282003417f4c0d012003104521040340200220046a41003a00002003200241016a2202470d000b200220046a21082004200128022c200128022822076b22096b2102200941014e044020022007200910401a200128022821070b2001200320046a3602302001200836022c200120023602280240200b200c20070440200128022c2108200128022821020b2002200820026b1005417f460440410021030c010b0240200141106a2001280228220241016a200128022c2002417f736a101b2204280204450d0020042802002d000041c001490d0020014180016a20044101105321022000411c6a2108200141f0006a200441001053210403402002280204200428020446044020022802082004280208460d030b20012002290204220d370390012001200d3703082001200141d8006a200141086a411c104b102122073a005720062001419c016a200141d7006a100d220928020045044041101045220b20073a000d2001200836029401200141013a0098012006200128029c012009200b100e200141003602900120014190016a10350b200210500c000b000b10000b200141286a102d0b20051028024020030d002000280200210202402000280220450d002006100b210303402003450d012002200a470440200320022d000d3a000d2003100c2006200141d8006a2003410d6a103621042006200128025820042003100e20021013210221030c010b0b0340200328020822050440200521030c010b0b2003100f200a21020b2000411c6a210503402002200a460d0141101045220320022d000d3a000d2001200536025c200141013a00602006200141386a2003410d6a103621042006200128023820042003100e20014100360258200141d8006a10352002101321020c000b000b200141a0016a240020000f0b000b9504010b7f230041d0006b22012400200141186a10242205200041106a10251026200520002903101027200528020c200541106a28020047044010000b200528020421082005280200200110242102200141c8006a4100360200200141406b4200370300200141386a420037030020014200370330027f200041206a2802004504402001410136023041010c010b200141306a410010372000411c6a210420002802182103037f2003200446047f200141306a41011037200128023005200141306a20032d000d10292003101321030c010b0b0b210a200141306a410472102a41011045220341fe013a0000200120033602302001200341016a220636023820012006360234200228020c200241106a280200470440100020012802342106200128023021030b2003210420022802042207200620036b22066a220b20022802084b04402002200b102e20022802042107200128023021040b200228020020076a2003200610401a2002200228020420066a36020420022001280234200a20046b6a102620022000280220105b2000411c6a210420002802182103034020032004470440200220032d000d102b2003101321030c010b0b0240200228020c2002280210460440200228020021030c010b100020022802002103200228020c2002280210460d0010000b2008200320022802041006200141306a102d2002102820051028200041186a103820001038200141d0006a24000b6501037f230041306b22022400200128020021032001280204210120022000410110552002102121002002101e200220003a002f200220014101756a22002002412f6a2001410171047f200028020020036a2802000520030b110100101f200241306a24000b7d01037f230041106b220124002000104f0240024020001056450d002000280204450d0020002802002d000041c001490d010b10000b200141086a20001023200128020c220041024f044010000b200128020821020340200004402000417f6a210020022d00002103200241016a21020c010b0b200141106a240020030b880702057f017e230041b0016b2200240010071001220110442202100220004188016a200041386a20022001101b22034100105520004188016a104f0240024020004188016a1056450d00200028028c01450d002000280288012d000041c001490d010b10000b200041e8006a20004188016a1023200028026c220141094f044010000b200028026821020340200104402001417f6a210120023100002005420886842105200241016a21020c010b0b024002402005500d00418008101c20055104402000410036028c0120004101360288012000200029038801370308200041086a101d0c020b418508101c20055104402000410036028c0120004102360288012000200029038801370310200041106a101d0c020b418e08101c20055104402000410036028c01200041033602880120002000290388013703182003200041186a10200c020b419908101c200551044020004188016a101e2000200041a8016a3502002205370350200041e8006a10242201200041d0006a10251026200120051027200128020c200141106a28020047044010000b20012802002001280204100320011028101f0c020b41a608101c20055104402000410036028c0120004104360288012000200029038801370320200041206a101d0c020b41b308101c200551044020004188016a101e20004188016a10142102200041d0006a1024210120004180016a4100360200200041f8006a4200370300200041f0006a420037030020004200370368200041e8006a2002102920002802682104200041e8006a410472102a20012004102620012002102b200128020c200141106a28020047044010000b20012802002001280204100320011028101f0c020b41bc08101c20055104402000410036028c01200041053602880120002000290388013703282003200041286a10200c020b41c608101c200551044020004188016a101e200041a8016a2802002103200041d0006a1024210120004180016a4100360200200041f8006a4200370300200041f0006a420037030020004200370368200041e8006a200345ad2205102c20002802682103200041e8006a410472102a20012003102620012005105d200128020c200141106a28020047044010000b20012802002001280204100320011028101f0c020b41d408101c2005520d002000410036028c0120004106360288012000200029038801370330200041306a101d0c010b10000b1047200041b0016a24000bd60101047f2001104e2204200128020422024b04401000200128020421020b20012802002105027f027f41002002450d001a410020052c00002203417f4a0d011a200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a0b21012000027f02402005450440410021030c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b2900200041003602082000420037020020004100102e200041146a41003602002000420037020c20000b4e01017f230041206b22012400200141186a4100360200200141106a4200370300200141086a42003703002001420037030020012000290300102c20012802002001410472102a200141206a24000b13002000280208200149044020002001102e0b0b080020002001105d0b1c01017f200028020c22010440200041106a20013602000b2000102f0b090020002001ad102c0b940201047f230041106b22042400200028020422012000280210220241087641fcffff07716a2103027f410020012000280208460d001a2003280200200241ff07714102746a0b2101200441086a20001032200428020c21020340024020012002460440200041003602142000280204210103402000280208220320016b41027522024103490d0220012802001a2000200028020441046a22013602040c000b000b200141046a220120032802006b418020470d0120032802042101200341046a21030c010b0b2002417f6a220241014d04402000418004418008200241016b1b3602100b03402001200347044020012802001a200141046a21010c010b0b20002000280204103320002802001a200441106a24000b090020002001ad105d0b7502027f017e4101210320014280015a0440034020012004845045044020044238862001420888842101200241016a2102200442088821040c010b0b200241384f047f2002103020026a0520020b41016a21030b200041186a2802000440200041046a103121000b2000200028020020036a3602000b1501017f200028020022010440200020013602040b0b3401017f200028020820014904402001104422022000280200200028020410401a2000102f20002001360208200020023602000b0b080020002802001a0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0b4f01037f20012802042203200128021020012802146a220441087641fcffff07716a21022000027f410020032001280208460d001a2002280200200441ff07714102746a0b360204200020023602000b2501017f200028020821020340200120024645044020002002417c6a22023602080c010b0b0b1d01017f03402000220128020022000d00200128020422000d000b20010b0f0020002802001a200041003602000b6201017f024020002802042203044020022d0000210203400240200220032d000d49044020032802002200450d040c010b200328020422000d0020012003360200200341046a0f0b200021030c000b000b200041046a21030b2001200336020020030ba40c02077f027e230041306b22052400200041046a2107024020014101460440200710312802002101200041186a22022002280200417f6a360200200710394180104f04402000410c6a2202280200417c6a2802001a20072002280200417c6a10330b200141384f047f2001103020016a0520010b41016a21012000280218450d012007103121000c010b0240200710390d00200041146a28020022014180084f0440200020014180786a360214200041086a2201280200220228020021042001200241046a360200200520043602182007200541186a103a0c010b2000410c6a2802002202200041086a2802006b4102752204200041106a2203280200220620002802046b220141027549044041802010452104200220064704400240200028020c220120002802102206470d0020002802082202200028020422034b04402000200220012002200220036b41027541016a417e6d41027422036a103b220136020c2000200028020820036a3602080c010b200541186a200620036b2201410175410120011b22012001410276200041106a103c2102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c200220093702082002103d200028020c21010b200120043602002000200028020c41046a36020c0c020b02402000280208220120002802042206470d00200028020c2202200028021022034904402000200120022002200320026b41027541016a41026d41027422036a103e22013602082000200028020c20036a36020c0c010b200541186a200320066b2201410175410120011b2201200141036a410276200041106a103c2102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c200220093702082002103d200028020821010b2001417c6a2004360200200020002802082201417c6a22023602082002280200210220002001360208200520023602182007200541186a103a0c010b20052001410175410120011b20042003103c210241802010452106024020022802082201200228020c2208470d0020022802042204200228020022034b04402002200420012004200420036b41027541016a417e6d41027422036a103b22013602082002200228020420036a3602040c010b200541186a200820036b2201410175410120011b22012001410276200241106a280200103c21042002280208210320022802042101034020012003470440200428020820012802003602002004200428020841046a360208200141046a21010c010b0b2002290200210920022004290200370200200420093702002002290208210920022004290208370208200420093702082004103d200228020821010b200120063602002002200228020841046a360208200028020c2104034020002802082004460440200028020421012000200228020036020420022001360200200228020421012002200436020420002001360208200029020c21092000200229020837020c200220093702082002103d052004417c6a210402402002280204220120022802002208470d0020022802082203200228020c22064904402002200120032003200620036b41027541016a41026d41027422066a103e22013602042002200228020820066a3602080c010b200541186a200620086b2201410175410120011b2201200141036a4102762002280210103c2002280208210620022802042101034020012006470440200528022020012802003602002005200528022041046a360220200141046a21010c010b0b20022902002109200220052903183702002002290208210a20022005290320370208200520093703182005200a370320103d200228020421010b2001417c6a200428020036020020022002280204417c6a3602040c010b0b0b200541186a20071032200528021c4100360200200041186a2100410121010b2000200028020020016a360200200541306a24000b09002000280204100f0b2801017f200028020820002802046b2201410874417f6a410020011b200028021420002802106a6b0ba10202057f017e230041206b22052400024020002802082202200028020c2206470d0020002802042203200028020022044b04402000200320022003200320046b41027541016a417e6d41027422046a103b22023602082000200028020420046a3602040c010b200541086a200620046b2202410175410120021b220220024102762000410c6a103c2103200028020821042000280204210203402002200446450440200328020820022802003602002003200328020841046a360208200241046a21020c010b0b2000290200210720002003290200370200200320073702002000290208210720002003290208370208200320073702082003103d200028020821020b200220012802003602002000200028020841046a360208200541206a24000b2501017f200120006b220141027521032001044020022000200110420b200220034102746a0b5f01017f2000410036020c200041106a200336020002402001044020014180808080044f0d012001410274104521040b200020043602002000200420024102746a22023602082000200420014102746a36020c2000200236020420000f0b000b3101027f200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b20002802001a0b1b00200120006b22010440200220016b22022000200110420b20020b3801017f41e008420037020041e808410036020041742100034020000440200041ec086a4100360200200041046a21000c010b0b410710480bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210401a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041a08b0436020c41880b200028020c41076a4178712200360200418c0b200036020041900b3f003602000b970101047f230041106b220124002001200036020c2000047f41900b200041086a2202411076220041900b2802006a2203360200418c0b2002418c0b28020022026a41076a417871220436020002400240200341107420044d044041900b200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104104041086a0541000b200141106a24000b0b002000410120001b10440b130020002d0000410171044020002802081a0b0b880101037f41ec08410136020041f0082802002100034020000440034041f40841f4082802002201417f6a2202360200200141014845044041ec084100360200200020024102746a22004184016a280200200041046a28020011000041ec08410136020041f00828020021000c010b0b41f408412036020041f008200028020022003602000c010b0b0b940101027f41ec08410136020041f008280200220145044041f00841f80836020041f80821010b024041f4082802002202412046044041840210442201450d0120011041220141f00828020036020041f008200136020041f4084100360200410021020b41f408200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41ec0841003602000b3801017f41fc0a420037020041840b410036020041742100034020000440200041880b6a4100360200200041046a21000c010b0b410810480b070041fc0a10460b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000104c20012802044f0d002002410471044010000c010b200042003702000b02402002411071450d002000104c20012802044d0d0020024104710440100020000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f2000104d2000104e6a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000104f41012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bb50102057f017e230041106b22022400200041046a210102402000280200220304402001280200220504402005200041086a2802006a21040b20002004360204200041086a2003360200200241086a20014100200420031051105220002002290308220637020420004100200028020022002006422088a76b2201200120004b1b3602000c010b200020012802002201047f2001200041086a2802006a0541000b360204200041086a41003602000b200241106a24000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a4114104b104c200241306a24000b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000be70101037f230041106b2204240020004200370200200041086a410036020020012802042103024002402002450440200321020c010b410021022003450d002003210220012802002d000041c001490d00200441086a2001105420004100200428020c2201200428020822022001105122032003417f461b20024520012003497222031b220536020820004100200220031b3602042000200120056b3602000c010b20012802002103200128020421012000410036020020004100200220016b20034520022001497222021b36020820004100200120036a20021b3602040b200441106a240020000b2101017f2001104e220220012802044b044010000b200020012001104d200210520bd60202077f017e230041206b220324002001280208220420024b0440200341186a2001105420012003280218200328021c105136020c200341106a20011054410021042001027f410020032802102206450d001a410020032802142208200128020c2207490d001a200820072007417f461b210520060b360210200141146a2005360200200141003602080b200141106a210903400240200420024f0d002001280214450d00200341106a2001105441002104027f410020032802102207450d001a410020032802142208200128020c2206490d001a200820066b2104200620076a0b21052001200436021420012005360210200341106a20094100200520041051105220012003290310220a3702102001200128020c200a422088a76a36020c2001200128020841016a22043602080c010b0b20032009290200220a3703082003200a370300200020034114104b1a200341206a24000b980101037f200028020445044041000f0b2000104f200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0bf50101057f0340024020002802102201200028020c460d00200141786a2802004504401000200028021021010b200141786a22022002280200417f6a220436020020040d002000200236021020004101200028020422032001417c6a28020022026b22011030220441016a20014138491b220520036a1058200220002802006a220320056a200320011042200141374d0440200028020020026a200141406a3a00000c020b200441f7016a220341ff014d0440200028020020026a20033a00002000280200200220046a6a210203402001450d03200220013a0000200141087621012002417f6a21020c000b000510000c020b000b0b0b0f00200020011059200020013602040b2f01017f2000280208200149044020011044200028020020002802041040210220002001360208200020023602000b0b1b00200028020420016a220120002802084b04402000200110590b0b9f0201057f02402001044020002802042104200041106a2802002202200041146a280200220349044020022001ad2004ad422086843702002000200028021041086a3602100f0b027f41002002200028020c22026b410375220541016a2206200320026b2202410275220320032006491b41ffffffff01200241037541ffffffff00491b2202450d001a200241037410450b2103200320054103746a22052001ad2004ad4220868437020020052000280210200028020c22016b22046b2106200441014e044020062001200410401a200028020c21010b2000200320024103746a3602142000200541086a3602102000200636020c2001450d010f0b200041c001105c20004100105a200028020020002802046a4100410010401a200010570b0b250020004101105a200028020020002802046a20013a00002000200028020441016a3602040bb80202037f037e024020015004402000418001105c0c010b20014280015a044020012105034020052006845045044020064238862005420888842105200241016a2102200642088821060c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff0171105c2000200028020420046a1058200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff0171105c0b2000200028020420026a1058200028020420002802006a417f6a210203402001200784500d02200220013c0000200742388620014208888421012002417f6a2102200742088821070c000b000b20002001a741ff0171105c0b200010570b0b6401004180080b5d696e697400696e69745f73657400696e736572745f736574006765745f7365745f73697a65006974657261746f725f7365740066696e645f7365740065726173655f736574006765745f7365745f656d70747900636c6561725f736574";

    public static String BINARY = BINARY_0;

    public static final String FUNC_FIND_SET = "find_set";

    public static final String FUNC_ITERATOR_SET = "iterator_set";

    public static final String FUNC_INSERT_SET = "insert_set";

    public static final String FUNC_INIT_SET = "init_set";

    public static final String FUNC_ERASE_SET = "erase_set";

    public static final String FUNC_GET_SET_EMPTY = "get_set_empty";

    public static final String FUNC_CLEAR_SET = "clear_set";

    public static final String FUNC_GET_SET_SIZE = "get_set_size";

    protected ReferenceDataTypeSetContract(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ReferenceDataTypeSetContract(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Uint8> find_set() {
        final WasmFunction function = new WasmFunction(FUNC_FIND_SET, Arrays.asList(), Uint8.class);
        return executeRemoteCall(function, Uint8.class);
    }

    public RemoteCall<TransactionReceipt> iterator_set() {
        final WasmFunction function = new WasmFunction(FUNC_ITERATOR_SET, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> iterator_set(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_ITERATOR_SET, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<TransactionReceipt> insert_set(Uint8 value) {
        final WasmFunction function = new WasmFunction(FUNC_INSERT_SET, Arrays.asList(value), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> insert_set(Uint8 value, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_INSERT_SET, Arrays.asList(value), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static RemoteCall<ReferenceDataTypeSetContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeSetContract.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ReferenceDataTypeSetContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeSetContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ReferenceDataTypeSetContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeSetContract.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<ReferenceDataTypeSetContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeSetContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<TransactionReceipt> init_set() {
        final WasmFunction function = new WasmFunction(FUNC_INIT_SET, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> init_set(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_INIT_SET, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<TransactionReceipt> erase_set(Uint8 value) {
        final WasmFunction function = new WasmFunction(FUNC_ERASE_SET, Arrays.asList(value), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> erase_set(Uint8 value, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_ERASE_SET, Arrays.asList(value), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<Boolean> get_set_empty() {
        final WasmFunction function = new WasmFunction(FUNC_GET_SET_EMPTY, Arrays.asList(), Boolean.class);
        return executeRemoteCall(function, Boolean.class);
    }

    public RemoteCall<TransactionReceipt> clear_set() {
        final WasmFunction function = new WasmFunction(FUNC_CLEAR_SET, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> clear_set(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_CLEAR_SET, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<Uint64> get_set_size() {
        final WasmFunction function = new WasmFunction(FUNC_GET_SET_SIZE, Arrays.asList(), Uint64.class);
        return executeRemoteCall(function, Uint64.class);
    }

    public static ReferenceDataTypeSetContract load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new ReferenceDataTypeSetContract(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static ReferenceDataTypeSetContract load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new ReferenceDataTypeSetContract(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
