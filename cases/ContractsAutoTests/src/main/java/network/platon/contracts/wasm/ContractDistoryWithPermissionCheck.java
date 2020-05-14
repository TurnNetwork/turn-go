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
    private static String BINARY_0 = "0x0061736d0100000001420c60027f7f0060017f0060017f017f60000060037f7f7f0060027f7f017f60037f7f7f017f60047f7f7f7f0060027f7e006000017f60047f7f7f7f017f60017f017e02d2010903656e760c706c61746f6e5f70616e6963000303656e760d706c61746f6e5f6f726967696e000103656e760e706c61746f6e5f64657374726f79000203656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000903656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e000003656e7617706c61746f6e5f6765745f73746174655f6c656e677468000503656e7610706c61746f6e5f6765745f7374617465000a03656e7610706c61746f6e5f7365745f73746174650007034b4a0302010100000100040001010002060b030002010208010008010002000504010000010202090602040603020205040401000000070301030106020202010705000402010000000000080405017001030305030100020608017f0141908b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300090f5f5f66756e63735f6f6e5f65786974003e06696e766f6b6500190908010041010b020c410acb5b4a100041d408100a1a4101103f103310400b190020004200370200200041086a41003602002000100b20000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0b070041d40810390b9f0801057f230041e0006b22022400200241086a410036020020024200370300200241a108102e1037200241003602382002420037033020024114103522053602502002200541146a2206360258200520014114102f1a200220063602544100210103402004200620056b4f044002402003044020022001410520036b74411f713a0040200241306a200241406b100e0b200241d0006a100f410021032002410036024820024200370340200241406b200228020420022d0000220141017620014101711b41017441017210102002410172210503402003200228020420022d00002201410176200141017122011b22044f0d01200228024020036a2002280208200520011b20036a2d000022014105763a00002002280240200228020420022d0000220441017620044101711b6a20036a41016a2001411f713a0000200341016a21030c000b000b05200420056a2d0000200141087441801e71722101200341086a210303402003410548450440200220012003417b6a220376411f713a0040200241306a200241406b100e0c010b0b200441016a210420022802502105200228025421060c010b0b200228024020046a41003a0000200241d0006a200241406b200241306a1011200241406b100f200241d0006a200228025420022802506b41066a10102002280254200228025022046b21014101210303402001044020042d000041002003411d764101716b41b3c5d1d0027141002003411c764101716b41dde788ea037141002003411b764101716b41fab384f5017141002003411a764101716b41ed9cc2b20271410020034119764101716b41b2afa9db0371200341057441e0ffffff037173737373737321032001417f6a2101200441016a21040c010b0b410021012002410036022820024200370320200241206a41061010200341017321044119210303402003417b46450440200228022020016a2004200376411f713a00002003417b6a2103200141016a21010c010b0b200241d0006a100f4100210320024100360218200242003703100240200228023420022802306b2201450d00200241106a200110122002280234200228023022046b22014101480d00200228021420042001102f1a2002200228021420016a3602140b200241d0006a200241106a200241206a1011200241106a100f200041086a4100360200200042003702002000100b20002002280208200520022d0000220141017122041b2002280204200141017620041b2201200141016a103d20004131103c200020022802542204200228025022016b200028020420002d0000220541017620054101711b6a103b03402003200420016b4f4504402000200120036a2d00004180086a2c0000103c200341016a210320022802502101200228025421040c010b0b200241d0006a100f200241206a100f200241306a100f20021039200241e0006a24000bc20101047f230041206b220224000240200028020422032000280208490440200320012d00003a00002000200028020441016a3602040c010b2000200320002802006b41016a10262105200241186a200041086a3602004100210320024100360214200028020420002802006b2104200504402005103521030b20022003360208200320046a220420012d00003a00002002200320056a3602142002200436020c2002200441016a3602102000200241086a1029200241086a10280b200241206a24000b1501017f200028020022010440200020013602040b0b870201047f230041206b22022400024020002802042203200028020022056b22042001490440200028020820036b200120046b22044f04400340200341003a00002000200028020441016a22033602042004417f6a22040d000c030b000b2000200110262105200241186a200041086a36020020024100360214200028020420002802006b210341002101200504402005103521010b200220013602082002200120036a22033602102002200120056a3602142002200336020c0340200341003a00002002200228021041016a22033602102004417f6a22040d000b2000200241086a1029200241086a10280c010b200420014d0d002000200120056a3602040b200241206a24000bda0301067f230041206b2203240020012802042105024020022802042208200228020022026b22044101480d002004200128020820056b4c0440034020022008460d02200520022d00003a00002001200128020441016a2205360204200241016a21020c000b000b2001200420056a20012802006b10262107200341186a200141086a3602004100210420034100360214200520012802006b2106200704402007103521040b200320043602082003200420066a22063602102003200420076a3602142003200636020c200341086a410472210403402002200846450440200620022d00003a00002003200328021041016a2206360210200241016a21020c010b0b200128020020052004102702402001280204220420056b220241004c0440200328021021020c010b2003280210220420052002102f1a2003200220046a2202360210200128020421040b20012002360204200128020021022001200328020c3602002001280208210520012003280214360208200320043602102003200236020c2003200536021420032002360208200341086a1028200128020421050b20002005360204200141003602042000200128020036020020012802082102200141003602082000200236020820014100360200200341206a24000b2901017f2001417f4c0440000b2000200110352202360200200020023602042000200120026a3602080b2701017f03402001411446450440200020016a41003a0000200141016a21010c010b0b200010010b3a01017f230041306b22012400200141086a1013200141206a200141086a100d200041186a200141206a1015200141206a1039200141306a24000b6100024020002d0000410171450440200041003b01000c010b200028020841003a00002000410036020420002d0000410171450d0020002802081a200041003602000b20002001290200370200200041086a200141086a2802003602002001100b0bf90101057f230041306b22012400200141186a1013200141086a200141186a100d027f024002400240024002402000411c6a28020020002d001822024101762203200241017122051b2204200128020c20012d00082202410176200241017122021b470d002001280210200141086a41017220021b210220050d01200041196a210003402003450d0320002d000020022d0000470d01200241016a2102200041016a21002003417f6a21030c000b000b200141086a1039417f0c040b20040d010b200141086a10390c010b200041206a2802002002200410322100200141086a1039417f20000d011a0b200141186a1002450b200141306a24000b3401017f230041106b220324002003200236020c200320013602082003200329030837030020002003411c1042200341106a24000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010bca0402057f017e23004190016b22002400100910032201103422021004200041d8006a200041086a20022001101722034100104a200041d8006a104602400240200041d8006a104b450d00200028025c450d0020002802582d000041c001490d010b10000b200041386a200041d8006a101a200028023c220141094f044010000b200028023821020340200104402001417f6a210120023100002005420886842105200241016a21020c010b0b024002402005500d0041a50810182005510440200041d8006a101b200041d8006a1014101c0c020b41aa0810182005510440200041d8006a101b200041d8006a10162103200041206a101d2101200041d0006a4100360200200041c8006a4200370300200041406b420037030020004200370338200041386a2003ac22054201862005423f87852205101e20002802382103200041386a410472101f200120031020200120051021200128020c200141106a28020047044010000b20012802002001280204100520011022101c0c020b41bb0810182005510440200041206a100a2101200041d8006a20034101104a200041d8006a20011023200041d8006a101b200041f0006a200041386a200110362203103a20031039101c200110390c020b41c60810182005520d00200041d8006a101b20004180016a200041f0006a10362102200041386a101d22012002102410202001200041206a200210362204102520041039200128020c200141106a28020047044010000b2001280200200128020410052001102220021039101c0c010b10000b103e20004190016a24000bd60101047f200110452204200128020422024b04401000200128020421020b20012802002105027f027f41002002450d001a410020052c00002203417f4a0d011a200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a0b21012000027f02402005450440410021030c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000bdd0101087f230041406a220124002000100a2107200042afb59bdd9e8485b9f800370310200041186a100a2105200141286a101d220320002903101021200328020c200341106a28020047044010000b02402003280200220420032802042208100622064504400c010b2001410036022020014200370318200141186a200610102004200820012802182204200128021c20046b1007417f47044020012001280218220241016a200128021c2002417f736a101720051023200621020b200141186a100f0b20031022200245044020052007103a0b200141406b240020000bb103010c7f230041e0006b22012400200141286a101d2104200141d8006a4100360200200141d0006a4200370300200141c8006a420037030020014200370340200141406b2000290310101e20012802402103200141406b410472101f200420031020200420002903101021200428020c200441106a28020047044010000b200428020421092004280200200141406b101d2102200041186a22071024210b41011035220341fe013a0000200120033602182001200341016a22053602202001200536021c200228020c200241106a2802004704401000200128021c2105200128021821030b2003210620022802042208200520036b22056a220c20022802084b04402002200c102a20022802042108200128021821060b200228020020086a20032005102f1a2002200228020420056a3602042002200128021c200b20066b6a10202002200141086a2007103622031025200310390240200228020c2002280210460440200228020021030c010b100020022802002103200228020c2002280210460d0010000b2009200320022802041008200141186a100f20021022200410222007103920001039200141e0006a24000b2900200041003602082000420037020020004100102a200041146a41003602002000420037020c20000b7902017f017e4101210220014280015a044041002102034020012003845045044020034238862001420888842101200241016a2102200342088821030c010b0b200241384f047f2002102c20026a0520020b41016a21020b200041186a2802000440200041046a102d21000b2000200028020020026a3602000bc40201067f200028020422012000280210220341087641fcffff07716a2102027f200120002802082205460440200041146a210441000c010b2001200028021420036a220441087641fcffff07716a280200200441ff07714102746a2106200041146a21042002280200200341ff07714102746a0b21030340024020032006460440200441003602000340200520016b41027522024103490d0220012802001a2000200028020441046a2201360204200028020821050c000b000b200341046a220320022802006b418020470d0120022802042103200241046a21020c010b0b2002417f6a220241014d04402000418004418008200241016b1b3602100b03402001200547044020012802001a200141046a21010c010b0b200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b20002802001a0b13002000280208200149044020002001102a0b0b08002000200110520b1c01017f200028020c22010440200041106a20013602000b2000102b0bb30201057f230041206b220224000240024002402000280204044020002802002d000041c001490d010b200241086a100a1a0c010b200241186a2000101a2000104521030240024002400240200228021822000440200228021c220420034f0d010b41002100200241106a410036020020024200370308410021030c010b200241106a410036020020024200370308200420032003417f461b220341704f0d04200020036a21052003410a4b0d010b200220034101743a0008200241086a41017221040c010b200341106a4170712206103521042002200336020c20022006410172360208200220043602100b034020002005470440200420002d00003a0000200441016a2104200041016a21000c010b0b200441003a00000b2001200241086a1015200241086a1039200241206a24000f0b000bbc0101047f230041306b22012400200141286a4100360200200141206a4200370300200141186a42003703002001420037031041012102024002400240200120001036220328020420032d00002200410176200041017122041b220041014d0440200041016b0d032003280208200341016a20041b2c0000417f4c0d010c030b200041374b0d010b200041016a21020c010b2000102c20006a41016a21020b2001200236021020031039200141106a410472101f200141306a240020020b5201037f230041106b2202240020022001280208200141016a20012d0000220341017122041b36020820022001280204200341017620041b36020c20022002290308370300200020021051200241106a24000b3701017f2001417f4c0440000b2001200028020820002802006b2200410174220220022001491b41ffffffff07200041ffffffff03491b0b270020022002280200200120006b22016b2202360200200141014e0440200220002001102f1a0b0b3101027f200028020821012000280204210203402001200247044020002001417f6a22013602080c010b0b20002802001a0b6701017f20002802002000280204200141046a1027200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b3401017f2000280208200149044020011034220220002802002000280204102f1a2000102b20002001360208200020023602000b0b080020002802001a0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0b7a01027f41a1082100024003402000410371044020002d0000450d02200041016a21000c010b0b2000417c6a21000340200041046a22002802002201417f73200141fffdfb776a7141808182847871450d000b0340200141ff0171450d01200041016a2d00002101200041016a21000c000b000b200041a1086b0bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d0440200020012002102f1a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3e01027f0340200245044041000f0b20002d0000220320012d00002204460440200141016a2101200041016a21002002417f6a21020c010b0b200320046b0b3501017f230041106b220041908b0436020c41fc0a200028020c41076a417871220036020041800b200036020041840b3f003602000b970101047f230041106b220124002001200036020c2000047f41840b200041086a2202411076220041840b2802006a220336020041800b200241800b28020022026a41076a417871220436020002400240200341107420044d044041840b200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104102f41086a0541000b200141106a24000b0b002000410120001b10340b4d01017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b200020012802082001280204103720000b5a01027f02402002410a4d0440200020024101743a0000200041016a21030c010b200241106a4170712204103521032000200236020420002004410172360200200020033602080b2003200120021038200220036a41003a00000b100020020440200020012002102f1a0b0b130020002d0000410171044020002802081a0b0ba10201047f20002001470440200128020420012d00002202410176200241017122031b2102200141016a21052001280208410a2101200520031b210520002d000022034101712204044020002802002203417e71417f6a21010b200220014d0440027f2004044020002802080c010b200041016a0b21012002044020012005200210310b200120026a41003a000020002d00004101710440200020023602040f0b200020024101743a00000f0b027f2003410171044020002802080c010b41000b1a416f2103200141e6ffffff074d0440410b20014101742203200220022003491b220341106a4170712003410b491b21030b200310352204200520021038200020023602042000200436020820002003410172360200200220046a41003a00000b0bf80101057f0240027f20002d0000220241017104402000280204210320002802002202417e71417f6a0c010b20024101762103410a0b220420032001200320014b1b220141106a417071417f6a410a2001410a4b1b2201460d00027f2001410a460440200041016a210420002802080c010b4100200120044d200141016a103522041b0d0120002d0000220241017104404101210520002802080c010b41012105200041016a0b210620042006027f2002410171044020002802040c010b200241fe01714101760b41016a10382005044020002004360208200020033602042000200141016a4101723602000f0b200020034101743a00000b0bf40101057f024002400240027f20002d00002202410171220345044020024101762104410a0c010b2000280204210420002802002202417e71417f6a0b22052004460440027f2002410171044020002802080c010b200041016a0b2106416f2103200541e6ffffff074d0440410b20054101742202200541016a220320032002491b220241106a4170712002410b491b21030b20031035220220062005103820002002360208200020034101723602000c010b2003450d01200028020821020b2000200441016a3602040c010b2000200441017441026a3a0000200041016a21020b200220046a220041003a0001200020013a00000b5a01017f02402003410a4d0440200020024101743a0000200041016a21030c010b200341106a4170712204103521032000200236020420002004410172360200200020033602080b2003200120021038200220036a41003a00000b880101037f41e008410136020041e4082802002100034020000440034041e80841e8082802002201417f6a2202360200200141014845044041e0084100360200200020024102746a22004184016a280200200041046a28020011010041e008410136020041e40828020021000c010b0b41e808412036020041e408200028020022003602000c010b0b0b940101027f41e008410136020041e408280200220145044041e40841ec0836020041ec0821010b024041e8082802002202412046044041840210342201450d0120011030220141e40828020036020041e408200136020041e8084100360200410021020b41e808200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41e00841003602000b3801017f41f00a420037020041f80a410036020041742100034020000440200041fc0a6a4100360200200041046a21000c010b0b4102103f0b070041f00a10390b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000104320012802044f0d002002410471044010000c010b200042003702000b02402002411071450d002000104320012802044d0d0020024104710440100020000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001044200010456a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000104641012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a411410421043200241306a24000b2101017f20011045220220012802044b044010000b2000200120011044200210470bd60202077f017e230041206b220324002001280208220420024b0440200341186a2001104920012003280218200328021c104836020c200341106a20011049410021042001027f410020032802102206450d001a410020032802142208200128020c2207490d001a200820072007417f461b210520060b360210200141146a2005360200200141003602080b200141106a210903400240200420024f0d002001280214450d00200341106a2001104941002104027f410020032802102207450d001a410020032802142208200128020c2206490d001a200820066b2104200620076a0b21052001200436021420012005360210200341106a20094100200520041048104720012003290310220a3702102001200128020c200a422088a76a36020c2001200128020841016a22043602080c010b0b20032009290200220a3703082003200a37030020002003411410421a200341206a24000b980101037f200028020445044041000f0b20001046200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0bf80101057f0340024020002802102201200028020c460d00200141786a28020041014904401000200028021021010b200141786a2202200228020041016b220436020020040d002000200236021020004101200028020422032001417c6a28020022026b2201102c220441016a20014138491b220520036a104d200220002802006a220320056a2003200110310240200141374d0440200028020020026a200141406a3a00000c010b200441f7016a220341ff014d0440200028020020026a20033a00002000280200200220046a6a210203402001450d02200220013a0000200141087621012002417f6a21020c000b000b10000b0c010b0b0b0f0020002001104e200020013602040b2f01017f200028020820014904402001103420002802002000280204102f210220002001360208200020023602000b0b1b00200028020420016a220120002802084b044020002001104e0b0b250020004101104f200028020020002802046a20013a00002000200028020441016a3602040be70101037f2001280200210441012102024002400240024020012802042201410146044020042c000022014100480d012000200141ff017110500c040b200141374b0d01200121020b200020024180017341ff017110500c010b2001102c220241b7016a22034180024e044010000b2000200341ff017110502000200028020420026a104d200028020420002802006a417f6a210320012102037f2002047f200320023a0000200241087621022003417f6a21030c010520010b0b21020b20002002104f200028020020002802046a20042002102f1a2000200028020420026a3602040b2000104c0bb80202037f037e02402001500440200041800110500c010b20014280015a044020012105034020052006845045044020064238862005420888842105200241016a2102200642088821060c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff017110502000200028020420046a104d200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff017110500b2000200028020420026a104d200028020420002802006a417f6a210203402001200784500d02200220013c0000200742388620014208888421012002417f6a2102200742088821070c000b000b20002001a741ff017110500b2000104c0b0b5701004180080b5071707a7279397838676632747664773073336a6e35346b686365366d7561376c006c617800696e697400646973746f72795f636f6e7472616374007365745f737472696e67006765745f737472696e67";

    public static String BINARY = BINARY_0;

    public static final String FUNC_DISTORY_CONTRACT = "distory_contract";

    public static final String FUNC_SET_STRING = "set_string";

    public static final String FUNC_GET_STRING = "get_string";

    protected ContractDistoryWithPermissionCheck(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ContractDistoryWithPermissionCheck(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
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

    public RemoteCall<TransactionReceipt> distory_contract() {
        final WasmFunction function = new WasmFunction(FUNC_DISTORY_CONTRACT, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> distory_contract(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_DISTORY_CONTRACT, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<TransactionReceipt> set_string(String name) {
        final WasmFunction function = new WasmFunction(FUNC_SET_STRING, Arrays.asList(name), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> set_string(String name, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SET_STRING, Arrays.asList(name), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<String> get_string() {
        final WasmFunction function = new WasmFunction(FUNC_GET_STRING, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static ContractDistoryWithPermissionCheck load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new ContractDistoryWithPermissionCheck(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static ContractDistoryWithPermissionCheck load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new ContractDistoryWithPermissionCheck(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
