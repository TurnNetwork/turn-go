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
public class ReferenceDataTypeVectorContract extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001420c60027f7f0060017f0060017f017f60037f7f7f017f60000060037f7f7f0060027f7f017f60047f7f7f7f0060047f7f7f7f017f60027f7e006000017f60017f017e02a9010703656e760c706c61746f6e5f70616e6963000403656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000a03656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e000003656e7617706c61746f6e5f6765745f73746174655f6c656e677468000603656e7610706c61746f6e5f6765745f7374617465000803656e7610706c61746f6e5f7365745f737461746500070355540402010100060800010000030b04000201000002000100000102090100010202000002030500010001020003080103030205040202060501000401040103020202010106070300050202070000000000000000090405017001030305030100020608017f0141908b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070f5f5f66756e63735f6f6e5f65786974004006696e766f6b6500140908010041010b020a430a936754100041d80810081a41011041103910420b190020004200370200200041086a41003602002000100920000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0b070041d808103e0b940101047f230041206b2202240002402000411c6a2802002203200041206a220428020047044020032001103c1a2000200028021c410c6a36021c0c010b200241086a200041186a2205200320002802186b410c6d41016a100c200028021c20002802186b410c6d2004100d22002802082001103c1a20002000280208410c6a36020820052000100e2000100f0b200241206a24000b3e01017f200141d6aad5aa014f0440000b2001200028020820002802006b410c6d2200410174220220022001491b41d5aad5aa01200041aad5aad500491b0b4c01017f2000410036020c200041106a2003360200200104402001102921040b20002004360200200020042002410c6c6a2202360208200020042001410c6c6a36020c2000200236020420000baa0101037f200028020421022000280200210303402002200346450440200128020441746a2204200241746a2202290200370200200441086a200241086a280200360200200210092001200128020441746a3602040c010b0b200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b3301027f2000280204210203402002200028020822014704402000200141746a22013602082001103e0c010b0b20002802001a0b2501017f230041106b2202240020022001103c200041186a20021011103e200241106a24000bba0101037f230041206b22032400024020002802042202200028020849044020022001290200370200200241086a200141086a2802003602002001100920002000280204410c6a3602040c010b200341086a2000200220002802006b410c6d41016a100c200028020420002802006b410c6d200041086a100d220228020822042001290200370200200441086a200141086a2802003602002001100920022002280208410c6a36020820002002100e2002100f0b200341206a24000b3401017f230041106b220324002003200236020c200320013602082003200329030837030020002003411c1044200341106a24000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010bb30502057f017e230041a0016b22002400100710012201103a22021002200041f8006a200041086a20022001101222034100104e200041f8006a104802400240200041f8006a104f450d00200028027c450d0020002802782d000041c001490d010b10000b200041d8006a200041f8006a1015200028025c220141094f044010000b200028025821020340200104402001417f6a210120023100002005420886842105200241016a21020c010b0b024002402005500d0041800810132005510440200041f8006a101610170c020b41850810132005510440200041306a10082101200041f8006a20034101104e200041f8006a200041306a1018200041f8006a1016200041d8006a2001103c200041f8006a200041d8006a100b103e10172001103e0c020b41980810132005510440200041306a10082101200041f8006a20034101104e200041f8006a20011019200041f8006a1016200041f8006a200041d8006a2001103c220310102003103e10172001103e0c020b41ab0810132005510440200041f8006a1016200041206a200028029001103c2102200041306a101a2101200041f0006a4100360200200041e8006a4200370300200041e0006a420037030020004200370358200041d8006a200041c8006a2002103c2204101b2004103e20002802582104200041d8006a410472101c20012004101d2001200041d8006a2002103c2204101e2004103e200128020c200141106a28020047044010000b2001280200200128020410032001101f2002103e10170c020b41c00810132005520d00200041f8006a1016200020004194016a2802002000280290016b410c6dad2205370330200041d8006a101a2201200041306a1020101d200120051021200128020c200141106a28020047044010000b2001280200200128020410032001101f10170c010b10000b1040200041a0016a24000bd60101047f200110472204200128020422024b04401000200128020421020b20012802002105027f027f41002002450d001a410020052c00002203417f4a0d011a200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a0b21012000027f02402005450440410021030c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b8c06020b7f017e23004190016b22022400200042003702182000428debc585c3a7f9dbf7003703102000410036020820004200370200200041206a4100360200200241306a101a220620002903101021200628020c200641106a28020047044010000b200041186a2107024002402006280200220a2006280204220b10042204044020024100360228200242003703202004417f4c0d012004103b21030340200120036a41003a00002004200141016a2201470d000b200120036a210520032002280224200228022022086b22096b2101200941014e044020012008200910361a200228022021080b2002200320046a36022820022005360224200220013602200240200a200b2008044020022802242105200228022021010b2001200520016b1005417f460440410021040c010b0240200241086a2002280220220141016a20022802242001417f736a10122201280204450d0020012802002d000041c001490d002001105021032000280220200028021822056b410c6d20034904402007200241d8006a2003200028021c20056b410c6d200041206a100d2203100e2003100f0b20024180016a20014101104c2103200241f0006a20014100104c210103402003280204200128020446044020032802082001280208460d030b20022003290204220c3703482002200c370300200241d8006a2002411c1044200241c8006a10082108200241c8006a10182007200241c8006a10112008103e200310490c000b000b10000b200241206a1022200421010b2006101f024020010d002000411c6a210620002802042203200028020022016b410c6d22052000280220200028021822046b410c6d4d04402005200628020020046b410c6d22054b0440200120012005410c6c6a22012004102a1a200120032006102b0c020b2007200120032004102a102c0c010b200404402007102d20002802181a20004100360220200042003702180b20072005100c220441d6aad5aa014f0d0220002004102922073602182000200736021c200020072004410c6c6a360220200120032006102b0b20024190016a240020000f0b000b000ba605010b7f23004180016b22012400200141286a101a2206200041106a1020101d200620002903101021200628020c200641106a28020047044010000b200628020421082006280200200141106a101a2102200141f8006a4100360200200141f0006a4200370300200141e8006a420037030020014200370360027f20002802182000411c6a2802004604402001410136026041010c010b200141e0006a4100102e200028021c210520002802182103037f2003200546047f200141e0006a4101102e200128026005200141e0006a4100102e200141e0006a200141d0006a2003103c2204101b2004103e200141e0006a4101102e2003410c6a21030c010b0b0b210a200141e0006a410472101c4101103b220341fe013a0000200120033602002001200341016a220436020820012004360204200228020c200241106a280200470440100020012802042104200128020021030b2003210520022802042207200420036b22046a220b20022802084b04402002200b102320022802042107200128020021050b200228020020076a2003200410361a2002200228020420046a36020420022001280204200a20056b6a101d2002200028021c20002802186b410c6d1056200141e0006a4104722105200028021c21042000280218210303402003200447044020024101105620014100360278200142003703702001420037036820014200370360200141e0006a200141d0006a2003103c2207101b2007103e20022001280260101d2002200141406b2003103c2207101e2007103e2005101c2003410c6a21030c010b0b0240200228020c2002280210460440200228020021030c010b100020022802002103200228020c2002280210460d0010000b2008200320022802041006200110222002101f2006101f200041186a102f2000102f20014180016a24000b2801017f230041206b22022400200241086a20004100104e200241086a20011019200241206a24000b8c0301057f230041206b220224000240024002402000280204044020002802002d000041c001490d010b200241086a10081a0c010b200241186a200010152000104721030240024002400240200228021822000440200228021c220420034f0d010b41002100200241106a410036020020024200370308410021030c010b200241106a410036020020024200370308200420032003417f461b220341704f0d04200020036a21052003410a4b0d010b200220034101743a0008200241086a41017221040c010b200341106a4170712206103b21042002200336020c20022006410172360208200220043602100b034020002005470440200420002d00003a0000200441016a2104200041016a21000c010b0b200441003a00000b024020012d0000410171450440200141003b01000c010b200128020841003a00002001410036020420012d0000410171450d0020012802081a200141003602000b20012002290308370200200141086a200241106a280200360200200241086a1009200241086a103e200241206a24000f0b000b29002000410036020820004200370200200041001023200041146a41003602002000420037020c20000b890101037f410121030240200128020420012d00002202410176200241017122041b2202450d0002400240200241014604402001280208200141016a20041b2c0000417f4c0d010c030b200241374b0d010b200241016a21030c010b2002102520026a41016a21030b200041186a2802000440200041046a102621000b2000200028020020036a3602000b940201047f230041106b22042400200028020422012000280210220241087641fcffff07716a2103027f410020012000280208460d001a2003280200200241ff07714102746a0b2101200441086a20001027200428020c21020340024020012002460440200041003602142000280204210103402000280208220320016b41027522024103490d0220012802001a2000200028020441046a22013602040c000b000b200141046a220120032802006b418020470d0120032802042101200341046a21030c010b0b2002417f6a220241014d04402000418004418008200241016b1b3602100b03402001200347044020012802001a200141046a21010c010b0b20002000280204102820002802001a200441106a24000b1300200028020820014904402000200110230b0b5201037f230041106b2202240020022001280208200141016a20012d0000220341017122041b36020820022001280204200341017620041b36020c20022002290308370300200020021059200241106a24000b1c01017f200028020c22010440200041106a20013602000b200010240b9e0102037f027e230041206b22012400200141186a4100360200200141106a4200370300200141086a42003703002001420037030041012103200029030022054280015a0440034020042005845045044020044238862005420888842105200241016a2102200442088821040c010b0b200241384f047f2002102520026a0520020b41016a21030b200120033602002001410472101c200141206a240020030b080020002001105a0b1501017f200028020022010440200020013602040b0b3401017f200028020820014904402001103a22022000280200200028020410361a2000102420002001360208200020023602000b0b080020002802001a0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0b4f01037f20012802042203200128021020012802146a220441087641fcffff07716a21022000027f410020032001280208460d001a2002280200200441ff07714102746a0b360204200020023602000b2501017f200028020821020340200120024645044020002002417c6a22023602080c010b0b0b1600200041d6aad5aa014f0440000b2000410c6c103b0b26000340200020014645044020022000103f2002410c6a21022000410c6a21000c010b0b20020b2e000340200020014645044020022802002000103c1a20022002280200410c6a3602002000410c6a21000c010b0b0b2901017f2000280204210203402001200246450440200241746a2202103e0c010b0b200020013602040b0b0020002000280200102c0ba40c02077f027e230041306b22052400200041046a2107024020014101460440200710262802002101200041186a22022002280200417f6a360200200710304180104f04402000410c6a2202280200417c6a2802001a20072002280200417c6a10280b200141384f047f2001102520016a0520010b41016a21012000280218450d012007102621000c010b0240200710300d00200041146a28020022014180084f0440200020014180786a360214200041086a2201280200220228020021042001200241046a360200200520043602182007200541186a10310c010b2000410c6a2802002202200041086a2802006b4102752204200041106a2203280200220620002802046b2201410275490440418020103b2104200220064704400240200028020c220120002802102206470d0020002802082202200028020422034b04402000200220012002200220036b41027541016a417e6d41027422036a1032220136020c2000200028020820036a3602080c010b200541186a200620036b2201410175410120011b22012001410276200041106a10332102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021034200028020c21010b200120043602002000200028020c41046a36020c0c020b02402000280208220120002802042206470d00200028020c2202200028021022034904402000200120022002200320026b41027541016a41026d41027422036a103522013602082000200028020c20036a36020c0c010b200541186a200320066b2201410175410120011b2201200141036a410276200041106a10332102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021034200028020821010b2001417c6a2004360200200020002802082201417c6a22023602082002280200210220002001360208200520023602182007200541186a10310c010b20052001410175410120011b2004200310332102418020103b2106024020022802082201200228020c2208470d0020022802042204200228020022034b04402002200420012004200420036b41027541016a417e6d41027422036a103222013602082002200228020420036a3602040c010b200541186a200820036b2201410175410120011b22012001410276200241106a280200103321042002280208210320022802042101034020012003470440200428020820012802003602002004200428020841046a360208200141046a21010c010b0b20022902002109200220042902003702002004200937020020022902082109200220042902083702082004200937020820041034200228020821010b200120063602002002200228020841046a360208200028020c2104034020002802082004460440200028020421012000200228020036020420022001360200200228020421012002200436020420002001360208200029020c21092000200229020837020c2002200937020820021034052004417c6a210402402002280204220120022802002208470d0020022802082203200228020c22064904402002200120032003200620036b41027541016a41026d41027422066a103522013602042002200228020820066a3602080c010b200541186a200620086b2201410175410120011b2201200141036a410276200228021010332002280208210620022802042101034020012006470440200528022020012802003602002005200528022041046a360220200141046a21010c010b0b20022902002109200220052903183702002002290208210a20022005290320370208200520093703182005200a3703201034200228020421010b2001417c6a200428020036020020022002280204417c6a3602040c010b0b0b200541186a20071027200528021c4100360200200041186a2100410121010b2000200028020020016a360200200541306a24000b1400200028020004402000102d20002802001a0b0b2801017f200028020820002802046b2201410874417f6a410020011b200028021420002802106a6b0ba10202057f017e230041206b22052400024020002802082202200028020c2206470d0020002802042203200028020022044b04402000200320022003200320046b41027541016a417e6d41027422046a103222023602082000200028020420046a3602040c010b200541086a200620046b2202410175410120021b220220024102762000410c6a10332103200028020821042000280204210203402002200446450440200328020820022802003602002003200328020841046a360208200241046a21020c010b0b20002902002107200020032902003702002003200737020020002902082107200020032902083702082003200737020820031034200028020821020b200220012802003602002000200028020841046a360208200541206a24000b2501017f200120006b220141027521032001044020022000200110380b200220034102746a0b5f01017f2000410036020c200041106a200336020002402001044020014180808080044f0d012001410274103b21040b200020043602002000200420024102746a22023602082000200420014102746a36020c2000200236020420000f0b000b3101027f200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b20002802001a0b1b00200120006b22010440200220016b22022000200110380b20020bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210361a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041908b0436020c41800b200028020c41076a417871220036020041840b200036020041880b3f003602000b970101047f230041106b220124002001200036020c2000047f41880b200041086a2202411076220041880b2802006a220336020041840b200241840b28020022026a41076a417871220436020002400240200341107420044d044041880b200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104103641086a0541000b200141106a24000b0b002000410120001b103a0ba10101037f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b20012802082103024020012802042201410a4d0440200020014101743a0000200041016a21020c010b200141106a4170712204103b21022000200136020420002004410172360200200020023602080b200220032001103d200120026a41003a000020000b10002002044020002001200210361a0b0b130020002d0000410171044020002802081a0b0ba10201047f20002001470440200128020420012d00002202410176200241017122031b2102200141016a21052001280208410a2101200520031b210520002d000022034101712204044020002802002203417e71417f6a21010b200220014d0440027f2004044020002802080c010b200041016a0b21012002044020012005200210380b200120026a41003a000020002d00004101710440200020023602040f0b200020024101743a00000f0b027f2003410171044020002802080c010b41000b1a416f2103200141e6ffffff074d0440410b20014101742203200220022003491b220341106a4170712003410b491b21030b2003103b220420052002103d200020023602042000200436020820002003410172360200200220046a41003a00000b0b880101037f41e408410136020041e8082802002100034020000440034041ec0841ec082802002201417f6a2202360200200141014845044041e4084100360200200020024102746a22004184016a280200200041046a28020011010041e408410136020041e80828020021000c010b0b41ec08412036020041e808200028020022003602000c010b0b0b940101027f41e408410136020041e808280200220145044041e80841f00836020041f00821010b024041ec0828020022024120460440418402103a2201450d0120011037220141e80828020036020041e808200136020041ec084100360200410021020b41ec08200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41e40841003602000b3801017f41f40a420037020041fc0a410036020041742100034020000440200041800b6a4100360200200041046a21000c010b0b410210410b070041f40a103e0b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000104520012802044f0d002002410471044010000c010b200042003702000b02402002411071450d002000104520012802044d0d0020024104710440100020000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001046200010476a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000104841012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bb50102057f017e230041106b22022400200041046a210102402000280200220304402001280200220504402005200041086a2802006a21040b20002004360204200041086a2003360200200241086a2001410020042003104a104b20002002290308220637020420004100200028020022002006422088a76b2201200120004b1b3602000c010b200020012802002201047f2001200041086a2802006a0541000b360204200041086a41003602000b200241106a24000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a411410441045200241306a24000b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000be70101037f230041106b2204240020004200370200200041086a410036020020012802042103024002402002450440200321020c010b410021022003450d002003210220012802002d000041c001490d00200441086a2001104d20004100200428020c2201200428020822022001104a22032003417f461b20024520012003497222031b220536020820004100200220031b3602042000200120056b3602000c010b20012802002103200128020421012000410036020020004100200220016b20034520022001497222021b36020820004100200120036a20021b3602040b200441106a240020000b2101017f20011047220220012802044b044010000b20002001200110462002104b0bd60202077f017e230041206b220324002001280208220420024b0440200341186a2001104d20012003280218200328021c104a36020c200341106a2001104d410021042001027f410020032802102206450d001a410020032802142208200128020c2207490d001a200820072007417f461b210520060b360210200141146a2005360200200141003602080b200141106a210903400240200420024f0d002001280214450d00200341106a2001104d41002104027f410020032802102207450d001a410020032802142208200128020c2206490d001a200820066b2104200620076a0b21052001200436021420012005360210200341106a2009410020052004104a104b20012003290310220a3702102001200128020c200a422088a76a36020c2001200128020841016a22043602080c010b0b20032009290200220a3703082003200a37030020002003411410441a200341206a24000b980101037f200028020445044041000f0b20001048200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b800101047f230041106b2201240002402000280204450d0020002802002d000041c001490d00200141086a2000104d200128020c210003402000450d01200141002001280208220320032000104a22046a20034520002004497222031b3602084100200020046b20031b2100200241016a21020c000b000b200141106a240020020b2d00200020021052200028020020002802046a2001200210361a2000200028020420026a3602042000200310530b1b00200028020420016a220120002802084b04402000200110550b0b820201047f02402001450d00034020002802102202200028020c460d01200241786a28020020014904401000200028021021020b200241786a2203200328020020016b220136020020010d012000200336021020004101200028020422042002417c6a28020022016b22021025220341016a20024138491b220520046a1054200120002802006a220420056a2004200210380240200241374d0440200028020020016a200241406a3a00000c010b200341f7016a220441ff014d0440200028020020016a20043a00002000280200200120036a6a210103402002450d02200120023a0000200241087621022001417f6a21010c000b000b10000b410121010c000b000b0b0f00200020011055200020013602040b2f01017f200028020820014904402001103a200028020020002802041036210220002001360208200020023602000b0b8d0201057f02402001044020002802042104200041106a2802002202200041146a280200220349044020022001ad2004ad422086843702002000200028021041086a3602100f0b027f41002002200028020c22026b410375220541016a2206200320026b2202410275220320032006491b41ffffffff01200241037541ffffffff00491b2202450d001a2002410374103b0b2103200320054103746a22052001ad2004ad4220868437020020052000280210200028020c22016b22046b2106200441014e044020062001200410361a200028020c21010b2000200320024103746a3602142000200541086a3602102000200636020c2001450d010f0b200041c0011057200041004100410110510b0b2500200041011052200028020020002802046a20013a00002000200028020441016a3602040b5e01027f20011025220241b7016a22034180024e044010000b2000200341ff017110572000200028020420026a1054200028020420002802006a417f6a2100034020010440200020013a0000200141087621012000417f6a21000c010b0b0b7501027f2001280200210341012102024002400240024020012802042201410146044020032c000022014100480d012000200141ff017110570c040b200141374b0d01200121020b200020024180017341ff017110570c010b200020011058200121020b200020032002410010510b2000410110530bba0202037f037e02402001500440200041800110570c010b20014280015a044020012105034020052006845045044020064238862005420888842105200241016a2102200642088821060c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff017110572000200028020420046a1054200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff017110570b2000200028020420026a1054200028020420002802006a417f6a210203402001200784500d02200220013c0000200742388620014208888421012002417f6a2102200742088821070c000b000b20002001a741ff017110570b2000410110530b0b5c01004180080b55696e697400736574436c6f74686573436f6c6f724f6e6500736574436c6f74686573436f6c6f7254776f00676574436c6f74686573436f6c6f72496e64657800676574436c6f74686573436f6c6f724c656e677468";

    public static String BINARY = BINARY_0;

    public static final String FUNC_SETCLOTHESCOLORONE = "setClothesColorOne";

    public static final String FUNC_SETCLOTHESCOLORTWO = "setClothesColorTwo";

    public static final String FUNC_GETCLOTHESCOLORINDEX = "getClothesColorIndex";

    public static final String FUNC_GETCLOTHESCOLORLENGTH = "getClothesColorLength";

    protected ReferenceDataTypeVectorContract(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ReferenceDataTypeVectorContract(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<ReferenceDataTypeVectorContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeVectorContract.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ReferenceDataTypeVectorContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeVectorContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ReferenceDataTypeVectorContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeVectorContract.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<ReferenceDataTypeVectorContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeVectorContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<TransactionReceipt> setClothesColorOne(Clothes myClothes) {
        final WasmFunction function = new WasmFunction(FUNC_SETCLOTHESCOLORONE, Arrays.asList(myClothes), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> setClothesColorOne(Clothes myClothes, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SETCLOTHESCOLORONE, Arrays.asList(myClothes), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<TransactionReceipt> setClothesColorTwo(String my_color) {
        final WasmFunction function = new WasmFunction(FUNC_SETCLOTHESCOLORTWO, Arrays.asList(my_color), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> setClothesColorTwo(String my_color, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SETCLOTHESCOLORTWO, Arrays.asList(my_color), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<String> getClothesColorIndex() {
        final WasmFunction function = new WasmFunction(FUNC_GETCLOTHESCOLORINDEX, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<Uint64> getClothesColorLength() {
        final WasmFunction function = new WasmFunction(FUNC_GETCLOTHESCOLORLENGTH, Arrays.asList(), Uint64.class);
        return executeRemoteCall(function, Uint64.class);
    }

    public static ReferenceDataTypeVectorContract load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new ReferenceDataTypeVectorContract(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static ReferenceDataTypeVectorContract load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new ReferenceDataTypeVectorContract(contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static class Clothes {
        public String color;
    }
}
