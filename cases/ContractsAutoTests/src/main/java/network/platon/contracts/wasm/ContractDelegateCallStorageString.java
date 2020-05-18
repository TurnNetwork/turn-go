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
    private static String BINARY_0 = "0x0061736d0100000001560f60027f7f0060017f0060017f017f60000060037f7f7f0060027f7f017f60037f7f7f017f60047f7f7f7f0060027f7e0060047f7f7f7f017f60017f017e60037f7f7e006000017f60057f7f7f7f7f017f60017e017f02c4010803656e760c706c61746f6e5f70616e6963000303656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000c03656e7610706c61746f6e5f6765745f696e707574000103656e7617706c61746f6e5f6765745f73746174655f6c656e677468000503656e7610706c61746f6e5f6765745f7374617465000903656e7610706c61746f6e5f7365745f7374617465000703656e7614706c61746f6e5f64656c65676174655f63616c6c000d03656e760d706c61746f6e5f72657475726e0000033c3b03020101010b05020502060800000800000100040104010000010600040e020000000105090002020200060403040a0a00030102030302020705000405017001030305030100020608017f0141f08b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300080f5f5f66756e63735f6f6e5f65786974003906696e766f6b6500340908010041010b020a0a0abe5d3b100041bc0910091a4101100b103c103d0b190020004200370200200041086a41003602002000100c20000b0300010b940101027f41c809410136020041cc09280200220145044041cc0941d40936020041d40921010b024041d00928020022024120460440418402102f2201450d012001103b220141cc0928020036020041cc09200136020041d0094100360200410021020b41d009200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41c80941003602000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0bf519020c7f017e23004190016b22032400200341e0006a418008100e220528020420032d00602204410176200441017122081b21042005280208200541016a20081b210642a5c688a1c89ca7f94b210f0340200404402004417f6a21042006300000200f42b383808080207e85210f200641016a21060c010b0b200341386a100f2105200341f0006a200110102101200528020421040240200541106a2802002208200541146a280200220749044020082004ad4220864202843702002005200528021041086a3602100c010b027f41002008200528020c22086b410375220641016a2209200720086b2208410275220720072009491b41ffffffff01200841037541ffffffff00491b2208450d001a200841037410110b2107200720064103746a22062004ad42208642028437020020062005280210200528020c220a6b22046b2109200441014e04402009200a200410121a0b2005200720084103746a3602142005200641086a3602102005200936020c0b200341306a4100360200200341286a4200370300200341206a420037030020034200370318200341186a200f1013200341186a20034180016a2001101010142005200328021810152005200f1016200520034180016a200110101017027f02400240200528020c2005280210460440200528020021010c010b100020052802002101200528020c2005280210470d010b20010c010b100020052802000b210420034100360210200342003703080240200420052802046a20016b2204450d00200341086a2004101820044101480d002003200328020c20012004101220046a36020c0b200341186a4104721019200528020c22010440200520013602100b200341d0006a418b08100e210a2000280208200041016a220c20002d0000220141017122051b21092000280204200141017620051b210741002104410121014100210541002108024003402001410171410020042007491b0440200420096a2d00002206415f6a41ff017141de004921012005200641bf7f6a41ff0171411a4972210520082006419f7f6a41ff0171411a49722108200441016a21040c01050240200120052008714101737121012007450d002007210403402004450d01200420096a2004417f6a22062104417f6a2d00004131470d000b0c030b0b0b417f21060b024002400240200141017145200641076a20074b72200641016a410249200741da004b72720d00410021042003410036028801200342003703800120034180016a20072006417f7322016a101a4101210502400240024003402004200028020420002d00002208410176200841017122081b20016a4f044002402005410171450d05200341f0006a1009220741016a2104410021010340200120064604402007280204220820032d007022014101762200200141017122051b2201200a28020420032d00502206410176200641017122061b470d05200a280208200a41016a20061b210620050d02200021012004210503402001450440200021080c080b20052d000020062d0000470d06200641016a2106200541016a21052001417f6a21010c000b00052000280208200c20002d00004101711b20016a2c000022054120722005200541bf7f6a411a491b210d024002400240027f20032d007022094101712205450440410a210820094101760c010b20072802002209417e71417f6a210820072802040b220b20084604402007280208200420094101711b210e416f2109200841e6ffffff074d0440410b20084101742205200841016a220920092005491b220541106a4170712005410b491b21090b200910112205200e2008101b20072009410172360200200720053602080c010b2005450d01200728020821050b2007200b41016a3602040c010b2003200b41017441026a3a0070200421050b2005200b6a220541003a00012005200d3a0000200141016a21010c010b000b000b0520032802800120046a2000280208200c20081b20066a20046a41016a2d00004190086a2d000022083a00002005200841ff0147712105200441016a21040c010b0b2001450d012007280208210503402001450d0220052d000020062d0000470d01200641016a2106200541016a21052001417f6a21010c000b000b200341386a101c200341286a200341c8006a280200360200200341206a200341406b290300370300200341003a002c200320032903383703182003280280012200450d0320032000360284010c030b410021062003410036024020034200370338200341386a2008410174410172101a03402006200728020420032d00702200410176200041017122001b2201490440200328023820066a2007280208200420001b20066a2d000022004105763a00002003280238200728020420032d0070220141017620014101711b6a20066a41016a2000411f713a0000200641016a21060c010b0b200328023820016a41003a0000200328023c21010240200328028401220020032802800122046b22054101480d0020052003280240220620016b4c0440034020002004460d02200120042d00003a00002003200328023c41016a220136023c200441016a21040c000b000b200341286a200341406b36020041002108200341003602242001200328023822076b2109200120056a20076b2205200620076b2207410174220620062005491b41ffffffff07200741ffffffff03491b220704402007101121080b200320083602182003200820096a22053602202003200720086a3602242003200536021c200341186a4104722108034020002004470440200520042d00003a00002003200328022041016a2205360220200441016a21040c010b0b200328023820012008101d0240200328023c220620016b220041004c0440200328022021010c010b2003200328022020012000101220006a2201360220200328023c21060b200328023821002003200328021c3602382003200036021c2003200136023c2003200636022020032802402101200320032802243602402003200136022420032000360218200341186a101e200328023c21010b2003420037023c2003280238210620034100360238200120066b21014101210403402001044020062d000041002004411d764101716b41b3c5d1d0027141002004411c764101716b41dde788ea037141002004411b764101716b41fab384f5017141002004411a764101716b41ed9cc2b20271410020044119764101716b41b2afa9db0371200441057441e0ffffff037173737373737321042001417f6a2101200641016a21060c010b0b20044101470d004100210020034100360268200342003703602003410036024020034200370338027f4100200328028401417a6a220120032802800122046b2205450d001a200341386a20051018034020012004470440200328023c220020042d00003a00002003200041016a36023c200441016a21040c010b0b20032802382100200328023c0b20006b210a200341e8006a210b410021044100210541002101024003402001200a4604400240200004402003200036023c0b2005410820046b7441ff0171452004410548710d00200328026021040c030b05200020016a2d0000200541057441e01f71722105200441056a21040340200441084e04402005200441786a22047621082003280264220720032802682209490440200720083a00002003200328026441016a360264052003200b36022841002106200341003602242007200328026022076b220c41016a220d200920076b220741017422092009200d491b41ffffffff07200741ffffffff03491b220704402007101121060b200320063602182006200c6a220920083a00002003200620076a3602242003200936021c2003200941016a360220200341e0006a200341186a101f200341186a101e0b0c010b0b200141016a21010c010b0b2003280264200328026022046b4114470d0041002104200341003a0018200341386a200341186a10202003280264200328026022006b210103402001200446200441134b72450440200341386a20046a200020046a2d00003a0000200441016a21040c010b0b200341286a200341c8006a280200360200200341206a200341406b29030037030020032003290338370318200341013a002c20000440200320003602640b2003280280012200044020032000360284010b410021042002210f0340200f50450440200f420888210f200441016a21040c010b0b2003410036024020034200370338200341386a2004101a200328023c417f6a21040340200250450440200420023c00002004417f6a2104200242088821020c010b0b200341186a20032802082204200328020c20046b20032802382200200328023c20006b10061a20032802382200450d042003200036023c0c040b2004450d00200320043602640b2003280280012200450d0020032000360284010b200341386a101c200341286a200341c8006a280200360200200341206a200341406b290300370300200341003a002c200320032903383703180b200328020821040b200404402003200436020c0b20034190016a24000b910101027f20004200370200200041086a410036020020012102024003402002410371044020022d0000450d02200241016a21020c010b0b2002417c6a21020340200241046a22022802002203417f73200341fffdfb776a7141808182847871450d000b0340200341ff0171450d01200241016a2d00002103200241016a21020c000b000b20002001200220016b103320000b29002000410036020820004200370200200041001027200041146a41003602002000420037020c20000b4d01017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b200020012802082001280204103320000b0b002000410120001b102f0bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000b880102027f017e4101210220014280015a044041002102034020012004845045044020044238862001420888842101200241016a2102200442088821040c010b0b200241384f047f2002103020026a0520020b41016a21020b200041186a28020022030440200041086a280200200041146a2802002003103221000b2000200028020020026a3602000b9a0101037f41012103024002400240200128020420012d00002202410176200241017122041b220241014d0440200241016b0d032001280208200141016a20041b2c0000417f4c0d010c030b200241374b0d010b200241016a21030c010b2002103020026a41016a21030b200041186a28020022010440200041086a280200200041146a2802002001103221000b2000200028020020036a3602000b1300200028020820014904402000200110270b0bbc0202037f027e02402001500440200041800110280c010b20014280015a044020012106034020052006845045044020054238862006420888842106200241016a2102200542088821050c010b0b0240200241384f04402002210403402004044020044108762104200341016a21030c010b0b200341c9004f044010000b2000200341b77f6a41ff017110282000200028020420036a1029200028020420002802006a417f6a21032002210403402004450d02200320043a0000200441087621042003417f6a21030c000b000b200020024180017341ff017110280b2000200028020420026a1029200028020420002802006a417f6a21024200210503402001200584500d02200220013c0000200542388620014208888421012002417f6a2102200542088821050c000b000b20002001a741ff017110280b2000102a0b810201047f410121022001280208200141016a20012d0000220341017122051b210402400240024002402001280204200341017620051b2203410146044020042c000022014100480d012000200141ff017110280c040b200341374b0d01200321020b200020024180017341ff017110280c010b20031030220141b7016a22024180024e044010000b2000200241ff017110282000200028020420016a1029200028020420002802006a417f6a210220032101037f2001047f200220013a0000200141087621012002417f6a21020c010520030b0b21020b200020021031200028020020002802046a2004200210121a2000200028020420026a3602040b2000102a0b2001017f2000200110112202360200200020023602042000200120026a3602080b860201067f200028020422032000280210220241087641fcffff07716a2101027f200320002802082204460440200041146a210541000c010b2003200028021420026a220541087641fcffff07716a280200200541ff07714102746a2106200041146a21052001280200200241ff07714102746a0b21020340024020022006460440200541003602000340200420036b41027522014103490d022000200341046a22033602040c000b000b200241046a220220012802006b418020470d0120012802042102200141046a21010c010b0b2001417f6a220141014d04402000418004418008200141016b1b3602100b03402003200447044020002004417c6a22043602080c010b0b0bfa0101057f230041206b22022400024020002802042203200028020022046b22052001490440200028020820036b200120056b22044f04400340200341003a00002000200028020441016a22033602042004417f6a22040d000c030b000b20002001102b2106200241186a200041086a3602002002410036021441002101200604402006101121010b200220013602082002200120056a22033602102002200120066a3602142002200336020c0340200341003a00002002200228021041016a22033602102004417f6a22040d000b2000200241086a101f200241086a101e0c010b200520014d0d002000200120046a3602040b200241206a24000b10002002044020002001200210121a0b0b2401017f230041106b22012400200141003a000f20002001410f6a1020200141106a24000b270020022002280200200120006b22016b2202360200200141014e044020022000200110121a0b0b2b01027f200028020821012000280204210203402001200247044020002001417f6a22013602080c010b0b0b6701017f20002802002000280204200141046a101d200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b2601017f03402002411446450440200020026a20012d00003a0000200241016a21020c010b0b0b960601097f230041f0006b22022400200241086a10091a200242e299efdb8683ebcf58370318200241206a10092106200241d8006a100f220420022903181016200428020c200441106a28020047044010000b02402004280200220520042802042207100322034504400c010b2002410036023820024200370330200241306a2003101a2005200720022802302205200228023420056b1004417f470440200241406b2002280230220141016a20022802342001417f736a102220061023200321010b20022802302203450d00200220033602340b200428020c22030440200420033602100b024020010d002002280210200241086a41017220022d0008220141017122031b2104200228020c200141017620031b22012006280200417e71417f6a410a20062d000041017122031b22054d0440200241286a280200200641016a20031b21032001044020032004200110240b200120036a41003a000020062d00004101710440200241246a20013602000c020b200620014101743a00000c010b416f2103200541e6ffffff074d0440410b20054101742203200120012003491b220341106a4170712003410b491b21030b20031011220520042001101b20062003410172360200200241286a2005360200200241246a2001360200200120056a41003a00000b2000200610101a200241d8006a100f2203200229031810251015200320022903181016200328020c200341106a28020047044010000b200328020421042003280200200241406b100f210120061026210741011011220041fe013a0000200128020c200141106a28020047044010000b2001280204220841016a220920012802084b047f20012009102720012802040520080b20012802006a2000410110121a2001200128020441016a3602042001200041016a200720006b6a10152001200241306a2006101010170240200128020c2001280210460440200128020021000c010b100020012802002100200128020c2001280210460d0010000b2004200020012802041005200128020c22000440200120003602100b200328020c22000440200320003602100b200241f0006a24000b0c00200020012002411c102c0bf40201057f230041206b22022400024002402000280204044020002802002d000041c001490d010b200241086a10091a0c010b200241186a2000102d2000102e21030240024002400240200228021822000440200228021c220420034f0d010b41002100200241106a410036020020024200370308410021040c010b200241106a4100360200200242003703082000200420032003417f461b22046a21052004410a4b0d010b200220044101743a0008200241086a41017221030c010b200441106a4170712206101121032002200436020c20022006410172360208200220033602100b03402000200546450440200320002d00003a0000200341016a2103200041016a21000c010b0b200341003a00000b024020012d0000410171450440200141003b01000c010b200128020841003a00002001410036020420012d0000410171450d00200141003602000b20012002290308370200200141086a200241106a280200360200200241086a100c200241206a24000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210121a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b4b01027f230041206b22012400200141186a4100360200200141106a4200370300200141086a420037030020014200370300200120001013200128020020014104721019200141206a24000b5501017f230041306b22012400200141286a4100360200200141206a4200370300200141186a420037030020014200370310200141106a20012000101010142001280210200141106a4104721019200141306a24000b2f01017f200028020820014904402001102f200028020020002802041012210220002001360208200020023602000b0b2500200041011031200028020020002802046a20013a00002000200028020441016a3602040b0f00200020011027200020013602040bf80101057f0340024020002802102201200028020c460d00200141786a28020041014904401000200028021021010b200141786a2202200228020041016b220436020020040d002000200236021020004101200028020422032001417c6a28020022026b22011030220441016a20014138491b220520036a1029200220002802006a220320056a2003200110240240200141374d0440200028020020026a200141406a3a00000c010b200441f7016a220341ff014d0440200028020020026a20033a00002000280200200220046a6a210203402001450d02200220013a0000200141087621012002417f6a21020c000b000b10000b0c010b0b0b2e01017f2001200028020820002802006b2200410174220220022001491b41ffffffff07200041ffffffff03491b0b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000103e20024f0d002003410471044010000c010b200042003702000b02402003411071450d002000103e20024d0d0020034104710440100020000f0b200042003702000b20000bd40101047f2001102e2204200128020422024b04401000200128020421020b200128020021052000027f02400240200204404100210120052c00002203417f4a0d01027f200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a21010c010b4101210120050d000c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000bff0201037f200028020445044041000f0b2000103a41012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b970101047f230041106b220124002001200036020c2000047f41ec0b200041086a2202411076220041ec0b2802006a220336020041e80b200241e80b28020022026a41076a417871220436020002400240200341107420044d044041ec0b200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104101241086a0541000b200141106a24000b1e01017f03402000044020004108762100200141016a21010c010b0b20010b1b00200028020420016a220120002802084b04402000200110270b0b25002000200120026a417f6a220141087641fcffff07716a280200200141ff07714102746a0b5a01027f02402002410a4d0440200020024101743a0000200041016a21030c010b200241106a4170712204101121032000200236020420002004410172360200200020033602080b200320012002101b200220036a41003a00000b930302047f017e23004180016b22002400100810012201102f22021002200041286a200041086a20022001102222014100103502400240200041286a10362204500d0041900910372004510d0141950910372004510440200041286a10092102200041346a100921032000420037034020004101360248200020013602702000200041c8006a360274200041f0006a20021038200041f0006a20031038200041d8006a2001200028024810352000200041d8006a1036370340200041f0006a20021010200041c8006a200310102000290340100d200041d8006a100f2201420010251015200142001016200128020c200141106a28020047044010000b200128020020012802041007200128020c2202450d02200120023602100c020b41ae0910372004520d00200041f0006a1021200041286a100f2201200041f0006a102610152001200041d8006a200041f0006a10101017200128020c200141106a28020047044010000b200128020020012802041007200128020c2202450d01200120023602100c010b10000b103920004180016a24000bc90202077f017e230041106b220324002001280208220520024b0440200341086a2001104220012003280208200328020c104136020c200320011042410021052001027f410020032802002206450d001a410020032802042208200128020c2207490d001a200820072007417f461b210420060b360210200141146a2004360200200141003602080b200141106a210903402001280214210402402005200249044020040d01410021040b2000200928020020044114102c1a200341106a24000f0b20032001104241002104027f410020032802002207450d001a410020032802042208200128020c2206490d001a200820066b2104200620076a0b21052001200436021420012005360210200320094100200520041041104020012003290300220a3702102001200128020c200a422088a76a36020c2001200128020841016a22053602080c000b000b870202047f017e230041106b220324002000103a024002402000280204450d002000103a0240200028020022012c0000220241004e044020020d010c020b200241807f460d00200241ff0171220441b7014d0440200028020441014d04401000200028020021010b20012d00010d010c020b200441bf014b0d012000280204200241ff017141ca7e6a22024d04401000200028020021010b200120026a2d0000450d010b2000280204450d0020012d000041c001490d010b10000b200341086a2000102d200328020c220041094f044010000b200328020821010340200004402000417f6a210020013100002005420886842105200141016a21010c010b0b200341106a240020050b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b4301017f230041206b22022400200241086a200028020020002802042802001035200241086a2001102320002802042200200028020041016a360200200241206a24000b880101037f41c809410136020041cc092802002100034020000440034041d00941d0092802002201417f6a2202360200200141014845044041c8094100360200200020024102746a22004184016a280200200041046a28020011010041c809410136020041cc0928020021000c010b0b41d009412036020041cc09200028020022003602000c010b0b0b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b3501017f230041106b220041f08b0436020c41e40b200028020c41076a417871220036020041e80b200036020041ec0b3f003602000b3801017f41d80b420037020041e00b410036020041742100034020000440200041e40b6a4100360200200041046a21000c010b0b4102100b0b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f2000103f2000102e6a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b2301017f230041206b22022400200241086a200020014114102c103e200241206a24000b2101017f2001102e220220012802044b044010000b200020012001103f200210400b0bc00101004180080bb8017365745f737472696e67006c61780000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0fff0a1115141a1e0705ffffffffffffff1dff180d19090817ff12161f1b13ff010003100b1c0c0e060402ffffffffffff1dff180d19090817ff12161f1b13ff010003100b1c0c0e060402ffffffffff696e69740064656c65676174655f63616c6c5f7365745f737472696e67006765745f737472696e67";

    public static String BINARY = BINARY_0;

    public static final String FUNC_DELEGATE_CALL_SET_STRING = "delegate_call_set_string";

    public static final String FUNC_GET_STRING = "get_string";

    protected ContractDelegateCallStorageString(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ContractDelegateCallStorageString(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<TransactionReceipt> delegate_call_set_string(String target_address, String name, Uint64 gas) {
        final WasmFunction function = new WasmFunction(FUNC_DELEGATE_CALL_SET_STRING, Arrays.asList(target_address,name,gas), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> delegate_call_set_string(String target_address, String name, Uint64 gas, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_DELEGATE_CALL_SET_STRING, Arrays.asList(target_address,name,gas), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
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
