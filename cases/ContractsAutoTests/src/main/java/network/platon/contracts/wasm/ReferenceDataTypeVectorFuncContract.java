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
 * <p>Generated with platon-web3j version 0.13.0.6.
 */
public class ReferenceDataTypeVectorFuncContract extends WasmContract {
    private static String BINARY_0 = "0x0061736d01000000014d0e60027f7f0060017f0060017f017f60037f7f7f017f60037f7f7f0060000060047f7f7f7f0060027f7f017f60047f7f7f7f017f60027f7e0060017f017e6000017f60027f7e017f60017e017f02a9010703656e760c706c61746f6e5f70616e6963000503656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000b03656e7610706c61746f6e5f6765745f696e707574000103656e7617706c61746f6e5f6765745f73746174655f6c656e677468000703656e7610706c61746f6e5f6765745f7374617465000803656e7610706c61746f6e5f7365745f7374617465000603656e760d706c61746f6e5f72657475726e00000357560502010101000706000708030104040000010001010100020402050203040a0a01000201020d0009010109010c05080100020007000303040000000300060001020300000000020003080103040205040502020600000405017001090905030100020608017f0141f08b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070f5f5f66756e63735f6f6e5f65786974003406696e766f6b650021090e010041010b0809091617181a1b090a8b6c56100041b40910081a4101100a105510570b190020004200370200200041086a41003602002000100b20000b0300010b940101027f41c009410136020041c409280200220145044041c40941cc0936020041cc0921010b024041c8092802002202412046044041840210222201450d0120011054220141c40928020036020041c409200136020041c8094100360200410021020b41c809200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41c00941003602000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0bd20302077f017e230041306b22062400200041186a21072000280218210402402000411c6a2802002202200041206a22032802004904402002200446044020042001100d1a2000200028021c410c6a36021c0c020b2007200420022004410c6a100e2004200420014d047f2001410c6a2001200028021c20014b1b0520010b100f0c010b024020062007200220046b410c6d41016a101041002003101122022802082200200228020c2208470d0020022802042203200228020022054b04402002200320002003200320056b410c6d41016a417e6d410c6c22056a101222003602082002200228020420056a3602040c010b200641186a200820056b2200410c6d410174410120001b22002000410276200241106a28020010112103200228020821052002280204210003402000200546450440200328020822082000290200370200200841086a200041086a2802003602002000100b20032003280208410c6a3602082000410c6a21000c010b0b20022902002109200220032902003702002003200937020020022902082109200220032902083702082003200937020820031013200228020821000b20002001100d1a20022002280208410c6a3602082007200220041014200210130b200641306a24000ba10101037f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b20012802082103024020012802042201410a4d0440200020014101743a0000200041016a21020c010b200141106a4170712204102021022000200136020420002004410172360200200020023602080b2002200320011056200120026a41003a000020000b970101037f20012000280204220520036b410c6d2206410c6c6a2103200521040340200320024f0440200541746a21042006410c6c2103200141746a21000340200304402004200020036a101d200441746a2104200341746a21030c010b0b0520042003290200370200200441086a200341086a2802003602002003100b20002000280204410c6a22043602042003410c6a21030c010b0b0b880201047f20002001470440200128020420012d00002202410176200241017122031b2102200141016a21042001280208410a2101200420031b210420002d0000410171220304402000280200417e71417f6a21010b200220014d0440027f2003044020002802080c010b200041016a0b21012002044020012004200210530b200120026a41003a000020002d00004101710440200020023602040f0b200020024101743a00000f0b416f2103200141e6ffffff074d0440410b20014101742201200220022001491b220141106a4170712001410b491b21030b200310202201200420021056200020023602042000200341017236020020002001360208200120026a41003a00000b0b3101017f2001200028020820002802006b410c6d2200410174220220022001491b41d5aad5aa01200041aad5aad500491b0b4c01017f2000410036020c200041106a2003360200200104402001101e21040b20002004360200200020042002410c6c6a2202360208200020042001410c6c6a36020c2000200236020420000b26000340200020014645044020022000101d2002410c6a21022000410c6a21000c010b0b20020b2b01027f20002802082101200028020421020340200120024704402000200141746a22013602080c010b0b0bb00101027f20002802002002200141046a101f2000280204210303402002200346450440200128020822042002290200370200200441086a200241086a2802003602002002100b20012001280208410c6a3602082002410c6a21020c010b0b200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000bb802010a7f230041206b22082400024020012802002204450d00200041186a210620002802182101200041206a22032802002000411c6a28020022056b410c6d20044f0440200520016b220b410c6d220920042203490440200920046b21032005210a0340200a2002100d1a2000200028021c410c6a220a36021c200341016a220720034f200721030d000b20092103200b450d020b20062001200520012004410c6c22076a100e200120024d0440200220076a2002200028021c20024b1b21020b03402003450d0220012002100f2001410c6a21012003417f6a21030c000b000b200841086a2006200520016b410c6d20046a101041002003101122032802082100034020002002100d1a20032003280208410c6a22003602082004417f6a22040d000b2006200320011014200310130b200841206a24000b0c0020002001280218100d1a0b120020002001411c6a28020041746a100d1a0b1400200041186a2000411c6a28020041746a10190b0900200020013602040b2101017f200041186a20002802182201410c6a2000411c6a2802002001101210190b0900200041186a101c0b0b002000200028020010190b5b00024020002d0000410171450440200041003b01000c010b200028020841003a00002000410036020420002d0000410171450d00200041003602000b20002001290200370200200041086a200141086a2802003602002001100b0b09002000410c6c10200b4a01017f03402000200146450440200228020041746a2203200141746a2201290200370200200341086a200141086a2802003602002001100b2002200228020041746a3602000c010b0b0b0b002000410120001b10220bf50502057f017e23004180016b22002400100710012201102222021002200041206a200020022001102322014100102402400240200041206a10252205500d0041800810262005510440410210270c020b41850810262005510440200041c8006a10082102200041206a200141011024200041206a20021028200041206a1029200041206a200041e0006a2002100d100c102a0c020b4197081026200551044020004200370360200041e8006a10082102200041206a2001410110242000200041206a1025370360200041206a200141021024200041206a20021028200041206a102920002903602105200041c8006a2002100d210220002005370318200041206a200041186a20021015102a0c020b41ad0810262005510440200041206a10292000413c6a280200210320002802382104200041e0006a102b2201200320046b410c6dad2205102c102d20012005102e200128020c200141106a28020047044010000b200128020020012802041006200128020c22030440200120033602100b102a0c020b41bd0810262005510440200041206a200141011024200041206a10252105200041206a102920002005370348200041e0006a200041206a280218200041c8006a280200410c6c6a100d1a200041e0006a102f102a0c020b41ca0810262005510440410310300c020b41da0810262005510440410410300c020b41e90810262005510440410510270c020b41fd0810262005510440410610270c020b418f0910262005510440410710270c020b41a10910262005520d00200041206a10292000413c6a280200210320002802382104200041c8006a102b2101200041f8006a4100360200200041f0006a4200370300200041e8006a420037030020004200370360200041e0006a2003200446ad2205103120002802602103200041e0006a410472103220012003102d200120051033220128020c200141106a28020047044010000b200128020020012802041006200128020c22030440200120033602100b102a0c010b10000b103420004180016a24000b970101047f230041106b220124002001200036020c2000047f41e40b200041086a2202411076220041e40b2802006a220336020041e00b200241e00b28020022026a41076a417871220436020002400240200341107420044d044041e40b200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104104241086a0541000b200141106a24000b0c00200020012002411c10350bc90202077f017e230041106b220324002001280208220520024b0440200341086a2001103920012003280208200328020c103a36020c200320011039410021052001027f410020032802002206450d001a410020032802042208200128020c2207490d001a200820072007417f461b210420060b360210200141146a2004360200200141003602080b200141106a210903402001280214210402402005200249044020040d01410021040b200020092802002004411410351a200341106a24000f0b20032001103941002104027f410020032802002207450d001a410020032802042208200128020c2206490d001a200820066b2104200620076a0b2105200120043602142001200536021020032009410020052004103a105a20012003290300220a3702102001200128020c200a422088a76a36020c2001200128020841016a22053602080c000b000b870202047f017e230041106b2203240020001036024002402000280204450d00200010360240200028020022012c0000220241004e044020020d010c020b200241807f460d00200241ff0171220441b7014d0440200028020441014d04401000200028020021010b20012d00010d010c020b200441bf014b0d012000280204200241ff017141ca7e6a22024d04401000200028020021010b200120026a2d0000450d010b2000280204450d0020012d000041c001490d010b10000b200341086a20001037200328020c220041094f044010000b200328020821010340200004402000417f6a210020013100002005420886842105200141016a21010c010b0b200341106a240020050b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b2701027f230041306b22012400200141086a1029200141086a2000110100102a200141306a24000ba10201057f230041206b22022400024002402000280204044020002802002d000041c001490d010b200241086a10081a0c010b200241186a200010372000103821030240024002400240200228021822000440200228021c220420034f0d010b41002100200241106a410036020020024200370308410021040c010b200241106a4100360200200242003703082000200420032003417f461b22046a21052004410a4b0d010b200220044101743a0008200241086a41017221030c010b200441106a4170712206102021032002200436020c20022006410172360208200220033602100b03402000200546450440200320002d00003a0000200341016a2103200041016a21000c010b0b200341003a00000b2001200241086a101d200241206a24000be207010c7f23004190016b22042400200042003702182000428debc585c3a7f9dbf7003703102000410036020820004200370200200041206a4100360200200441186a102b22082000290310102e200828020c200841106a28020047044010000b200041186a21072000411c6a210a0240200828020022022008280204220510032206450d002006102021030340200120036a41003a00002006200141016a2201470d000b20022005200320011004417f460440410021010c010b024002402004200341016a200120036a2003417f736a10232203280204450d0020032802002d000041c001490d00200441406b200310392004280244210141002102034020010440200441002004280240220520052001103a22096a20054520012009497222051b3602404100200120096b20051b2101200241016a21020c010b0b2000280220200028021822016b410c6d20024904402007200441406b2002200028021c20016b410c6d200041206a10112201103b200110130b200441e8006a20034101103c2102200441d8006a20034100103c210b200041206a210c200228020421010340200b280204200146410020022802082203200b280208461b0d02200441406b20012003411c1035200441306a1008220310280240200028021c2201200028022049044020012004290330370200200141086a200441386a2802003602002003100b200a200a280200410c6a3602000c010b200441f8006a2007200120072802006b410c6d220141016a10102001200c1011210120042802800122052004290330370200200541086a200441386a2802003602002003100b2004200428028001410c6a3602800120072001103b200110130b20022002280204220120022802086a410020011b22013602042002280200220304402002200336020820012003103a21092002027f200228020422034504404100210541000c010b410021054100200228020822012009490d001a200120092009417f461b210520030b2201ad2005ad42208684370204200241002002280200220320056b2205200520034b1b36020005200241003602080b0c000b000b10000b200621010b200828020c22060440200820063602100b024020010d0020002802042203200028020022016b410c6d22022000280220200028021822066b410c6d4d04402002200a28020020066b410c6d22024b0440200120012002410c6c6a22012006103d1a20012003200a103e0c020b2007200120032006103d10190c010b200604402007101c20004100360220200042003702180b20002007200210102202101e22063602182000200636021c200020062002410c6c6a36022020012003200a103e0b20044190016a240020000b8406010b7f230041e0006b22032400200341186a102b22072000290310102c102d20072000290310102e200728020c200741106a28020047044010000b2007280204210a20072802002003102b2101200341c8006a4100360200200341406b4200370300200341386a420037030020034200370330027f20002802182000411c6a2802004604402003410136023041010c010b200341306a4100103f200028021c210420002802182102037f2002200446047f200341306a4101103f200328023005200341306a200341d0006a2002100d10402002410c6a21020c010b0b0b2104200341306a410472103241011020220241fe013a0000200128020c200141106a28020047044010000b2001280204220541016a220620012802084b047f20012006104120012802040520050b20012802006a2002410110421a2001200128020441016a3602042001200241016a200420026b6a102d0240200028021c20002802186b220204402002410c6d21022001280204210420012802102205200141146a280200220649044020052002ad2004ad422086843702002001200128021041086a3602100c020b027f41002005200128020c22056b410375220841016a2209200620056b2205410275220620062009491b41ffffffff01200541037541ffffffff00491b2205450d001a200541037410200b2106200620084103746a22082002ad2004ad4220868437020020082001280210200128020c22096b22026b2104200241014e044020042009200210421a0b2001200620054103746a3602142001200841086a3602102001200436020c0c010b200141c0011043200141004100410110440b200028021c2104200028021821020340200220044704402001200341306a2002100d10452002410c6a21020c010b0b0240200128020c2001280210460440200128020021020c010b100020012802002102200128020c2001280210460d0010000b200a200220012802041005200128020c22020440200120023602100b200728020c22010440200720013602100b200041186a104620001046200341e0006a24000b29002000410036020820004200370200200041001041200041146a41003602002000420037020c20000b4b01027f230041206b22012400200141186a4100360200200141106a4200370300200141086a420037030020014200370300200120001031200128020020014104721032200141206a24000b1300200028020820014904402000200110410b0b09002000200110331a0bab0101037f230041e0006b22012400200141186a102b2102200141d8006a4100360200200141d0006a4200370300200141c8006a420037030020014200370340200141406b200141306a2000100d104020012802402103200141406b410472103220022003102d2002200141086a2000100d1045200228020c200241106a28020047044010000b200228020020022802041006200228020c22000440200220003602100b200141e0006a24000b3301027f230041406a22012400200141086a1029200141306a200141086a2000110000200141306a102f102a200141406b24000b840102027f017e4101210320014280015a0440034020012004845045044020044238862001420888842101200241016a2102200442088821040c010b0b200241384f047f2002104720026a0520020b41016a21030b200041186a28020022020440200041086a280200200041146a2802002002104821000b2000200028020020036a3602000bea0101047f230041106b22042400200028020422012000280210220341087641fcffff07716a2102027f410020012000280208460d001a2002280200200341ff07714102746a0b2101200441086a20001049200428020c210303400240200120034604402000410036021420002802082102200028020421010340200220016b41027522034103490d022000200141046a22013602040c000b000b200141046a220120022802006b418020470d0120022802042101200241046a21020c010b0b2003417f6a220241014d04402000418004418008200241016b1b3602100b20002001104a200441106a24000bbc0202037f037e02402001500440200041800110430c010b20014280015a044020012107034020062007845045044020064238862007420888842107200241016a2102200642088821060c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff017110432000200028020420046a105c200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff017110430b2000200028020420026a105c200028020420002802006a417f6a210203402001200584500d02200220013c0000200542388620014208888421012002417f6a2102200542088821050c000b000b20002001a741ff017110430b20004101104c20000b880101037f41c009410136020041c4092802002100034020000440034041c80941c8092802002201417f6a2202360200200141014845044041c0094100360200200020024102746a22004184016a280200200041046a28020011010041c009410136020041c40928020021000c010b0b41c809412036020041c409200028020022003602000c010b0b0b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000105820024f0d002003410471044010000c010b200042003702000b02402003411071450d002000105820024d0d0020034104710440100020000f0b200042003702000b20000b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bd40101047f200110382204200128020422024b04401000200128020421020b200128020021052000027f02400240200204404100210120052c00002203417f4a0d01027f200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a21010c010b4101210120050d000c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000bff0201037f200028020445044041000f0b2000103641012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b2101017f20011038220220012802044b044010000b20002001200110592002105a0b2301017f230041206b22022400200241086a20002001411410351058200241206a24000b6701017f20002802002000280204200141046a101f200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000be70101037f230041106b2204240020004200370200200041086a410036020020012802042103024002402002450440200321020c010b410021022003450d002003210220012802002d000041c001490d00200441086a2001103920004100200428020c2201200428020822022001103a22032003417f461b20024520012003497222031b220536020820004100200220031b3602042000200120056b3602000c010b20012802002103200128020421012000410036020020004100200220016b20034520022001497222021b36020820004100200120036a20021b3602040b200441106a240020000b26000340200020014645044020022000100f2002410c6a21022000410c6a21000c010b0b20020b2e000340200020014645044020022802002000100d1a20022002280200410c6a3602002000410c6a21000c010b0b0bbd0c02077f027e230041306b22052400200041046a2107024020014101460440200041086a280200200041146a280200200041186a220228020022041048280200210120022004417f6a3602002007104d4180104f044020072000410c6a280200417c6a104a0b200141384f047f2001104720016a0520010b41016a2101200041186a2802002202450d01200041086a280200200041146a2802002002104821000c010b02402007104d0d00200041146a28020022014180084f0440200020014180786a360214200041086a2201280200220228020021042001200241046a360200200520043602182007200541186a104e0c010b2000410c6a2802002202200041086a2802006b4102752204200041106a2203280200220620002802046b220141027549044041802010202104200220064704400240200028020c220120002802102206470d0020002802082202200028020422034b04402000200220012002200220036b41027541016a417e6d41027422036a104f220136020c2000200028020820036a3602080c010b200541186a200620036b2201410175410120011b22012001410276200041106a10502102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021051200028020c21010b200120043602002000200028020c41046a36020c0c020b02402000280208220120002802042206470d00200028020c2202200028021022034904402000200120022002200320026b41027541016a41026d41027422036a105222013602082000200028020c20036a36020c0c010b200541186a200320066b2201410175410120011b2201200141036a410276200041106a10502102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021051200028020821010b2001417c6a2004360200200020002802082201417c6a22023602082002280200210220002001360208200520023602182007200541186a104e0c010b20052001410175410120011b200420031050210241802010202106024020022802082201200228020c2208470d0020022802042204200228020022034b04402002200420012004200420036b41027541016a417e6d41027422036a104f22013602082002200228020420036a3602040c010b200541186a200820036b2201410175410120011b22012001410276200241106a280200105021042002280208210320022802042101034020012003470440200428020820012802003602002004200428020841046a360208200141046a21010c010b0b20022902002109200220042902003702002004200937020020022902082109200220042902083702082004200937020820041051200228020821010b200120063602002002200228020841046a360208200028020c2104034020002802082004460440200028020421012000200228020036020420022001360200200228020421012002200436020420002001360208200029020c21092000200229020837020c2002200937020820021051052004417c6a210402402002280204220120022802002208470d0020022802082203200228020c22064904402002200120032003200620036b41027541016a41026d41027422066a105222013602042002200228020820066a3602080c010b200541186a200620086b2201410175410120011b2201200141036a410276200228021010502002280208210620022802042101034020012006470440200528022020012802003602002005200528022041046a360220200141046a21010c010b0b20022902002109200220052903183702002002290208210a20022005290320370208200520093703182005200a3703201051200228020421010b2001417c6a200428020036020020022002280204417c6a3602040c010b0b0b200541186a20071049200528021c4100360200200041186a2100410121010b2000200028020020016a360200200541306a24000b9a0101037f41012103024002400240200128020420012d00002202410176200241017122041b220241014d0440200241016b0d032001280208200141016a20041b2c0000417f4c0d010c030b200241374b0d010b200241016a21030c010b2002104720026a41016a21030b200041186a28020022010440200041086a280200200041146a2802002001104821000b2000200028020020036a3602000b2f01017f2000280208200149044020011022200028020020002802041042210220002001360208200020023602000b0bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000b250020004101105b200028020020002802046a20013a00002000200028020441016a3602040b2d0020002002105b200028020020002802046a2001200210421a2000200028020420026a36020420002003104c0b8f0101047f410121022001280208200141016a20012d0000220441017122051b210302400240024002402001280204200441017620051b2201410146044020032c000022014100480d012000200141ff017110430c040b200141374b0d01200121020b200020024180017341ff017110430c010b20002001104b200121020b200020032002410010440b20004101104c0b0e00200028020004402000101c0b0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b25002000200120026a417f6a220141087641fcffff07716a280200200141ff07714102746a0b4f01037f20012802042203200128021020012802146a220441087641fcffff07716a21022000027f410020032001280208460d001a2002280200200441ff07714102746a0b360204200020023602000b2501017f200028020821020340200120024645044020002002417c6a22023602080c010b0b0b5e01027f20011047220241b7016a22034180024e044010000b2000200341ff017110432000200028020420026a105c200028020420002802006a417f6a2100034020010440200020013a0000200141087621012000417f6a21000c010b0b0b820201047f02402001450d00034020002802102202200028020c460d01200241786a28020020014904401000200028021021020b200241786a2203200328020020016b220136020020010d012000200336021020004101200028020422042002417c6a28020022016b22021047220341016a20024138491b220520046a105c200120002802006a220420056a2004200210530240200241374d0440200028020020016a200241406a3a00000c010b200341f7016a220441ff014d0440200028020020016a20043a00002000280200200120036a6a210103402002450d02200120023a0000200241087621022001417f6a21010c000b000b10000b410121010c000b000b0b2801017f200028020820002802046b2201410874417f6a410020011b200028021420002802106a6b0ba10202057f017e230041206b22052400024020002802082202200028020c2206470d0020002802042203200028020022044b04402000200320022003200320046b41027541016a417e6d41027422046a104f22023602082000200028020420046a3602040c010b200541086a200620046b2202410175410120021b220220024102762000410c6a10502103200028020821042000280204210203402002200446450440200328020820022802003602002003200328020841046a360208200241046a21020c010b0b20002902002107200020032902003702002003200737020020002902082107200020032902083702082003200737020820031051200028020821020b200220012802003602002000200028020841046a360208200541206a24000b2501017f200120006b220141027521032001044020022000200110530b200220034102746a0b4f01017f2000410036020c200041106a2003360200200104402001410274102021040b200020043602002000200420024102746a22023602082000200420014102746a36020c2000200236020420000b2b01027f200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b0b1b00200120006b22010440200220016b22022000200110530b20020b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210421a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b3501017f230041106b220041f08b0436020c41dc0b200028020c41076a417871220036020041e00b200036020041e40b3f003602000b10002002044020002001200210421a0b0b3801017f41d00b420037020041d80b410036020041742100034020000440200041dc0b6a4100360200200041046a21000c010b0b4108100a0b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001059200010386a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b1b00200028020420016a220120002802084b04402000200110410b0b0f00200020011041200020013602040b0bb80101004180080bb001696e697400696e73657274566563746f7256616c756500696e73657274566563746f724d616e6756616c756500676574566563746f724c656e6774680066696e64566563746f7241740066696e64566563746f7246726f6e740066696e64566563746f724261636b0064656c657465566563746f72506f704261636b0064656c657465566563746f7245726173650064656c657465566563746f72436c6561720066696e64566563746f72456d707479";

    public static String BINARY = BINARY_0;

    public static final String FUNC_INSERTVECTORMANGVALUE = "insertVectorMangValue";

    public static final String FUNC_INSERTVECTORVALUE = "insertVectorValue";

    public static final String FUNC_GETVECTORLENGTH = "getVectorLength";

    public static final String FUNC_FINDVECTORAT = "findVectorAt";

    public static final String FUNC_FINDVECTORFRONT = "findVectorFront";

    public static final String FUNC_FINDVECTORBACK = "findVectorBack";

    public static final String FUNC_DELETEVECTORPOPBACK = "deleteVectorPopBack";

    public static final String FUNC_DELETEVECTORERASE = "deleteVectorErase";

    public static final String FUNC_DELETEVECTORCLEAR = "deleteVectorClear";

    public static final String FUNC_FINDVECTOREMPTY = "findVectorEmpty";

    protected ReferenceDataTypeVectorFuncContract(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    protected ReferenceDataTypeVectorFuncContract(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }

    public RemoteCall<TransactionReceipt> insertVectorMangValue(Uint64 num, String my_value) {
        final WasmFunction function = new WasmFunction(FUNC_INSERTVECTORMANGVALUE, Arrays.asList(num,my_value), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> insertVectorMangValue(Uint64 num, String my_value, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_INSERTVECTORMANGVALUE, Arrays.asList(num,my_value), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<TransactionReceipt> insertVectorValue(String my_value) {
        final WasmFunction function = new WasmFunction(FUNC_INSERTVECTORVALUE, Arrays.asList(my_value), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> insertVectorValue(String my_value, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_INSERTVECTORVALUE, Arrays.asList(my_value), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static RemoteCall<ReferenceDataTypeVectorFuncContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeVectorFuncContract.class, web3j, credentials, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<ReferenceDataTypeVectorFuncContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeVectorFuncContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<ReferenceDataTypeVectorFuncContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeVectorFuncContract.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public static RemoteCall<ReferenceDataTypeVectorFuncContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ReferenceDataTypeVectorFuncContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public RemoteCall<Uint64> getVectorLength() {
        final WasmFunction function = new WasmFunction(FUNC_GETVECTORLENGTH, Arrays.asList(), Uint64.class);
        return executeRemoteCall(function, Uint64.class);
    }

    public RemoteCall<String> findVectorAt(Uint64 index) {
        final WasmFunction function = new WasmFunction(FUNC_FINDVECTORAT, Arrays.asList(index), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<String> findVectorFront() {
        final WasmFunction function = new WasmFunction(FUNC_FINDVECTORFRONT, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<String> findVectorBack() {
        final WasmFunction function = new WasmFunction(FUNC_FINDVECTORBACK, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<TransactionReceipt> deleteVectorPopBack() {
        final WasmFunction function = new WasmFunction(FUNC_DELETEVECTORPOPBACK, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> deleteVectorPopBack(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_DELETEVECTORPOPBACK, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<TransactionReceipt> deleteVectorErase() {
        final WasmFunction function = new WasmFunction(FUNC_DELETEVECTORERASE, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> deleteVectorErase(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_DELETEVECTORERASE, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<TransactionReceipt> deleteVectorClear() {
        final WasmFunction function = new WasmFunction(FUNC_DELETEVECTORCLEAR, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> deleteVectorClear(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_DELETEVECTORCLEAR, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<Boolean> findVectorEmpty() {
        final WasmFunction function = new WasmFunction(FUNC_FINDVECTOREMPTY, Arrays.asList(), Boolean.class);
        return executeRemoteCall(function, Boolean.class);
    }

    public static ReferenceDataTypeVectorFuncContract load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        return new ReferenceDataTypeVectorFuncContract(contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    public static ReferenceDataTypeVectorFuncContract load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        return new ReferenceDataTypeVectorFuncContract(contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }
}
