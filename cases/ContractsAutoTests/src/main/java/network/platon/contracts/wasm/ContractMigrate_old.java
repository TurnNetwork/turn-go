package network.platon.contracts.wasm;

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
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 0.7.5.3-SNAPSHOT.
 */
public class ContractMigrate_old extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001540e60027f7f0060017f0060017f017f60027f7f017f60000060037f7f7f0060047f7f7f7f0060037f7f7f017f60027f7e0060047f7f7f7f017f60017f017e60047f7f7e7e006000017f60077f7f7f7f7f7f7f017f02e3010a03656e760c706c61746f6e5f70616e6963000403656e760c706c61746f6e5f6576656e74000603656e760e706c61746f6e5f6d696772617465000d03656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000c03656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e000003656e760b706c61746f6e5f73686133000603656e7617706c61746f6e5f6765745f73746174655f6c656e677468000303656e7610706c61746f6e5f6765745f7374617465000903656e7610706c61746f6e5f7365745f737461746500060361600401000307000100050301000100030000020000010001000b08070a00020201040a00000202000305010000010008080202000002000709010702070205040202030505010000060401040107020202010603000502060000000000000000080405017001050505030100020608017f0141908b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f7273000a0f5f5f66756e63735f6f6e5f65786974005206696e766f6b65002a090a010041010b0421210b550abf74601f0041d808420037020041e008410036020041d808101441031053104810540b070041d808104e0b970701067f230041d0006b22022400200241a108100d21052002410036022820024200370320200241406b2001200141146a100e210641002101034020042006280204200628020022076b4f044002402003044020022001410520036b74411f713a0030200241206a200241306a100f0b20061010410021032002410036023820024200370330200241306a200528020420022d0000220141017620014101711b4101744101721011200541016a210603402003200528020420022d00002201410176200141017122011b22044f0d01200228023020036a2005280208200620011b20036a2d000022014105763a00002002280230200528020420022d0000220441017620044101711b6a20036a41016a2001411f713a0000200341016a21030c000b000b05200420076a2d0000200141087441801e71722101200341086a210303402003410548450440200220012003417b6a220376411f713a0030200241206a200241306a100f0c010b0b200441016a21040c010b0b200228023020046a41003a0000200241406b200241306a200241206a1012200241306a1010200241406b200228024420022802406b41066a10112002280244200228024022046b21014101210303402001044020042d000041002003411d764101716b41b3c5d1d0027141002003411c764101716b41dde788ea037141002003411b764101716b41fab384f5017141002003411a764101716b41ed9cc2b20271410020034119764101716b41b2afa9db0371200341057441e0ffffff037173737373737321032001417f6a2101200441016a21040c010b0b410021012002410036021820024200370310200241106a41061011200341017321044119210303402003417b46450440200228021020016a2004200376411f713a00002003417b6a2103200141016a21010c010b0b200241406b1010200241406b200241306a200241206a10132201200241106a10122001101041002103200041086a4100360200200042003702002000101420002005280208200620022d0000220141017122041b2005280204200141017620041b2201200141016a1051200041311050200020022802442204200228024022016b200028020420002d0000220641017620064101711b6a104f03402003200420016b4f4504402000200120036a2d00004180086a2c00001050200341016a210320022802402101200228024421040c010b0b200241406b1010200241106a1010200241206a10102005104e200241d0006a24000b1f0020004200370200200041086a41003602002000200120011044104c20000b450020004100360208200042003702000240200220016b2202450d0020002002103420024101480d0020002802042001200210451a2000200028020420026a3602040b20000bc20101047f230041206b220224000240200028020422032000280208490440200320012d00003a00002000200028020441016a3602040c010b2000200320002802006b41016a10312105200241186a200041086a3602004100210320024100360214200028020420002802006b2104200504402005104a21030b20022003360208200320046a220420012d00003a00002002200320056a3602142002200436020c2002200441016a3602102000200241086a1035200241086a10330b200241206a24000b1501017f200028020022010440200020013602040b0b870201047f230041206b22022400024020002802042203200028020022056b22042001490440200028020820036b200120046b22044f04400340200341003a00002000200028020441016a22033602042004417f6a22040d000c030b000b2000200110312105200241186a200041086a36020020024100360214200028020420002802006b210341002101200504402005104a21010b200220013602082002200120036a22033602102002200120056a3602142002200336020c0340200341003a00002002200228021041016a22033602102004417f6a22040d000b2000200241086a1035200241086a10330c010b200420014d0d002000200120056a3602040b200241206a24000bda0301067f230041206b2203240020012802042105024020022802042208200228020022026b22044101480d002004200128020820056b4c0440034020022008460d02200520022d00003a00002001200128020441016a2205360204200241016a21020c000b000b2001200420056a20012802006b10312107200341186a200141086a3602004100210420034100360214200520012802006b2106200704402007104a21040b200320043602082003200420066a22063602102003200420076a3602142003200636020c200341086a410472210403402002200846450440200620022d00003a00002003200328021041016a2206360210200241016a21020c010b0b200128020020052004103202402001280204220420056b220241004c0440200328021021020c010b200328021022042005200210451a2003200220046a2202360210200128020421040b20012002360204200128020021022001200328020c3602002001280208210520012003280214360208200320043602102003200236020c2003200536021420032002360208200341086a1033200128020421050b20002005360204200141003602042000200128020036020020012802082102200141003602082000200236020820014100360200200341206a24000b5a01017f20004200370200200041003602080240200128020420012802006b2202450d002000200210342001280204200128020022026b22014101480d0020002802042002200110451a2000200028020420016a3602040b20000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0bc10301067f230041c0016b22022400200241a508100d2104200241e8006a1016200241fc006a410036020020024200370274200241e8006a41021065200241d8006a20041017200241c8006a20001017200241406b4100360200200241386a4200370300200241306a420037030020024200370328200241286a200241d8006a1018200241c8006a1018200241e8006a20022802281019200241e8006a200241d8006a101a200241e8006a200241c8006a101a2002280274200241f8006a28020047044010000b200228026c21062002280268200241106a101b210020024180016a2001104b2101200241a8016a4100360200200241a0016a420037030020024198016a4200370300200242003703900120024190016a4100101c20024190016a200241b0016a2001104b2203101d2003104e20024190016a4101101c200228029001210320024190016a410472101e200020031019200041011065200020024190016a2001104b2203101f2003104e2001104e200028020c200041106a28020047044010000b20062000280200200028020410012000102041046a101e200241c8006a1010200241d8006a1010200241e8006a10202004104e200241c0016a24000b160020004100360208200042003702002000410010370b800101037f230041306b22032400200341186a101b22022001102e10192002200341086a2001104b2201101f2001104e200228020c200241106a28020047044010000b2000410036020820004200370200200228020421012002280200200041201011200120002802002201200028020420016b100620021020200341306a24000b860101027f02402001280200220320012802042201460440410121020c010b4101210202400240200120036b2201410146044020032c0000417f4c0d010c030b200141374b0d010b200141016a21020c010b2001103a20016a41016a21020b027f200041186a2802000440200041046a103b0c010b20000b2201200128020020026a36020020000b1300200028020820014904402000200110370b0b3d01027f230041106b220224002002200128020022033602082002200128020420036b36020c20022002290308370300200020021068200241106a24000b190020001016200041146a41003602002000420037020c20000ba40c02077f027e230041306b22052400200041046a21070240200141014604402007103b2802002101200041186a22022002280200417f6a3602002007103e4180104f04402000410c6a2202280200417c6a2802001a20072002280200417c6a103d0b200141384f047f2001103a20016a0520010b41016a21012000280218450d012007103b21000c010b02402007103e0d00200041146a28020022014180084f0440200020014180786a360214200041086a2201280200220228020021042001200241046a360200200520043602182007200541186a103f0c010b2000410c6a2802002202200041086a2802006b4102752204200041106a2203280200220620002802046b2201410275490440418020104a2104200220064704400240200028020c220120002802102206470d0020002802082202200028020422034b04402000200220012002200220036b41027541016a417e6d41027422036a1040220136020c2000200028020820036a3602080c010b200541186a200620036b2201410175410120011b22012001410276200041106a10412102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021042200028020c21010b200120043602002000200028020c41046a36020c0c020b02402000280208220120002802042206470d00200028020c2202200028021022034904402000200120022002200320026b41027541016a41026d41027422036a104322013602082000200028020c20036a36020c0c010b200541186a200320066b2201410175410120011b2201200141036a410276200041106a10412102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021042200028020821010b2001417c6a2004360200200020002802082201417c6a22023602082002280200210220002001360208200520023602182007200541186a103f0c010b20052001410175410120011b2004200310412102418020104a2106024020022802082201200228020c2208470d0020022802042204200228020022034b04402002200420012004200420036b41027541016a417e6d41027422036a104022013602082002200228020420036a3602040c010b200541186a200820036b2201410175410120011b22012001410276200241106a280200104121042002280208210320022802042101034020012003470440200428020820012802003602002004200428020841046a360208200141046a21010c010b0b20022902002109200220042902003702002004200937020020022902082109200220042902083702082004200937020820041042200228020821010b200120063602002002200228020841046a360208200028020c2104034020002802082004460440200028020421012000200228020036020420022001360200200228020421012002200436020420002001360208200029020c21092000200229020837020c2002200937020820021042052004417c6a210402402002280204220120022802002208470d0020022802082203200228020c22064904402002200120032003200620036b41027541016a41026d41027422066a104322013602042002200228020820066a3602080c010b200541186a200620086b2201410175410120011b2201200141036a410276200228021010412002280208210620022802042101034020012006470440200528022020012802003602002005200528022041046a360220200141046a21010c010b0b20022902002109200220052903183702002002290208210a20022005290320370208200520093703182005200a3703201042200228020421010b2001417c6a200428020036020020022002280204417c6a3602040c010b0b0b200541186a2007103c200528021c4100360200200041186a2100410121010b2000200028020020016a360200200541306a24000b890101037f410121030240200128020420012d00002202410176200241017122041b2202450d0002400240200241014604402001280208200141016a20041b2c0000417f4c0d010c030b200241374b0d010b200241016a21030c010b2002103a20026a41016a21030b200041186a2802000440200041046a103b21000b2000200028020020036a3602000b940201047f230041106b22042400200028020422012000280210220241087641fcffff07716a2103027f410020012000280208460d001a2003280200200241ff07714102746a0b2101200441086a2000103c200428020c21020340024020012002460440200041003602142000280204210103402000280208220320016b41027522024103490d0220012802001a2000200028020441046a22013602040c000b000b200141046a220120032802006b418020470d0120032802042101200341046a21030c010b0b2002417f6a220241014d04402000418004418008200241016b1b3602100b03402001200347044020012802001a200141046a21010c010b0b20002000280204103d20002802001a200441106a24000b5201037f230041106b2202240020022001280208200141016a20012d0000220341017122041b36020820022001280204200341017620041b36020c20022002290308370300200020021068200241106a24000b1c01017f200028020c22010440200041106a20013602000b200010360b0900200020013a00100bc70101027f230041d0006b22042400034020054114470440200441186a20056a41003a0000200541016a21050c010b0b200441406b20021023200441306a20031023200441186a20012802002205200128020420056b20042802402201200428024420016b20042802302201200428023420016b10021a200441306a1010200441406b1010200441406b200441186a100c200441086a200441186a100c200441406b200441086a1015200441086a104e200441406b104e2000200441186a100c200441d0006a24000b6302017f017e20012103034020035045044020034208882103200241016a21020c010b0b20004100360208200042003702002000200210112000280204417f6a21020340200150450440200220013c00002002417f6a2102200142088821010c010b0b0b3401017f230041106b220324002003200236020c200320013602082003200329030837030020002003411c1056200341106a24000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b6701047f230041206b220224002001280200210320012802042101200241086a20004101105e200241086a10272104200241086a1028200241086a20014101756a220020042001410171047f200028020020036a2802000520030b1100001029200241206a24000b7d01037f230041106b220124002000105a024002402000105f450d002000280204450d0020002802002d000041c001490d010b10000b200141086a2000102c200128020c220041024f044010000b200128020821020340200004402000417f6a210020022d00002103200241016a21020c010b0b200141106a240020030bd60101067f230041406a22012400200042e4e9dbd7f9a5b1f1e000370308200041003a0000200141286a101b220320002903081038200328020c200341106a28020047044010000b20032802002204200328020422061007220504402001410036022020014200370318200141186a200510112004200620012802182204200128021c20046b1008417f470440200020012001280218220241016a200128021c2002417f736a102410273a0010200521020b200141186a10100b200310202002450440200020002d00003a00100b200141406b240020000b9d03010b7f230041d0006b22022400200241186a101b2103200241c8006a4100360200200241406b4200370300200241386a420037030020024200370330200241306a2000290308103920022802302105200241306a410472101e200320051019200320002903081038200328020c200341106a28020047044010000b200328020421072003280200200241306a101b2101200041106a2209102f210a4101104a220041fe013a0000200220003602082002200041016a22043602102002200436020c200128020c200141106a2802004704401000200228020c2104200228020821000b2000210520012802042206200420006b22046a220b20012802084b04402001200b103720012802042106200228020821050b200128020020066a2000200410451a2001200128020420046a3602042001200228020c200a20056b6a1019200120092d000010300240200128020c2001280210460440200128020021000c010b100020012802002100200128020c2001280210460d0010000b2007200020012802041009200241086a10102001102020031020200241d0006a24000bc80502057f017e230041b0016b22002400100a10032201104922021004200041406b200041106a20022001102422014100105e02400240200041406b102b2205500d0041ae08102520055104402000410036024420004101360240200020002903403703002001200010260c020b41b308102520055104402000420037035820004200370350200041003602482000420037034020004101360264200020013602682000200041e4006a36026c20004198016a20014101105e200041d8006a2101200041d0006a210202400240200028029c0104402000280298012d000041c001490d010b20004100360230200042003703280c010b20004188016a20004198016a102c2000280288012103200041f8006a20004198016a102c200041286a2003200028027820004198016a10596a100e1a20002802402203450d002000200336024420004100360248200042003703400b20002000280228360240200029022c2105200041003602302000200537024420004200370328200041286a10102000200028026441016a360264200041e8006a2002102d200041e8006a2001102d200041286a1028200041f8006a200041e8006a200041406b1013220320002903502000290358102220004198016a101b2201200041f8006a102e1019200120004188016a200041f8006a104b2204101f2004104e200128020c200141106a28020047044010000b20012802002001280204100520011020200041f8006a104e200310101029200041406b10100c020b41c408102520055104402000410036024420004102360240200020002903403703082001200041086a10260c020b41cd0810252005520d0020004198016a1028200020002d00a80122033a0028200041406b101b2201200041286a102f1019200120031030200128020c200141106a28020047044010000b2001280200200128020410052001102010290c010b10000b1052200041b0016a24000b850102027f017e230041106b220124002000105a024002402000105f450d002000280204450d0020002802002d000041c001490d010b10000b200141086a2000102c200128020c220041094f044010000b200128020821020340200004402000417f6a210020023100002003420886842103200241016a21020c010b0b200141106a240020030bd60101047f200110592204200128020422024b04401000200128020421020b20012802002105027f027f41002002450d001a410020052c00002203417f4a0d011a200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a0b21012000027f02402005450440410021030c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b4601017f230041206b22022400200241086a20002802002000280204280200105e2001200241086a102b37030020002802042200200028020041016a360200200241206a24000b5b01017f230041306b22012400200141286a4100360200200141206a4200370300200141186a420037030020014200370310200141106a20012000104b2200101d2000104e2001280210200141106a410472101e200141306a24000b4e01017f230041206b22012400200141186a4100360200200141106a4200370300200141086a42003703002001420037030020012000310000103920012802002001410472101e200141206a24000b090020002001ad10690b3701017f2001417f4c0440000b2001200028020820002802006b2200410174220220022001491b41ffffffff07200041ffffffff03491b0b270020022002280200200120006b22016b2202360200200141014e044020022000200110451a0b0b3101027f200028020821012000280204210203402001200247044020002001417f6a22013602080c010b0b20002802001a0b2901017f2001417f4c0440000b20002001104a2202360200200020023602042000200120026a3602080b6701017f20002802002000280204200141046a1032200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b080020002802001a0b3401017f200028020820014904402001104922022000280200200028020410451a2000103620002001360208200020023602000b0b08002000200110690b7502027f017e4101210320014280015a0440034020012004845045044020044238862001420888842101200241016a2102200442088821040c010b0b200241384f047f2002103a20026a0520020b41016a21030b200041186a2802000440200041046a103b21000b2000200028020020036a3602000b1e01017f03402000044020004108762100200141016a21010c010b0b20010b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0b4f01037f20012802042203200128021020012802146a220441087641fcffff07716a21022000027f410020032001280208460d001a2002280200200441ff07714102746a0b360204200020023602000b2501017f200028020821020340200120024645044020002002417c6a22023602080c010b0b0b2801017f200028020820002802046b2201410874417f6a410020011b200028021420002802106a6b0ba10202057f017e230041206b22052400024020002802082202200028020c2206470d0020002802042203200028020022044b04402000200320022003200320046b41027541016a417e6d41027422046a104022023602082000200028020420046a3602040c010b200541086a200620046b2202410175410120021b220220024102762000410c6a10412103200028020821042000280204210203402002200446450440200328020820022802003602002003200328020841046a360208200241046a21020c010b0b20002902002107200020032902003702002003200737020020002902082107200020032902083702082003200737020820031042200028020821020b200220012802003602002000200028020841046a360208200541206a24000b2501017f200120006b220141027521032001044020022000200110470b200220034102746a0b5f01017f2000410036020c200041106a200336020002402001044020014180808080044f0d012001410274104a21040b200020043602002000200420024102746a22023602082000200420014102746a36020c2000200236020420000f0b000b3101027f200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b20002802001a0b1b00200120006b22010440200220016b22022000200110470b20020b7801027f20002101024003402001410371044020012d0000450d02200141016a21010c010b0b2001417c6a21010340200141046a22012802002202417f73200241fffdfb776a7141808182847871450d000b0340200241ff0171450d01200141016a2d00002102200141016a21010c000b000b200120006b0bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210451a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041908b0436020c41800b200028020c41076a417871220036020041840b200036020041880b3f003602000b970101047f230041106b220124002001200036020c2000047f41880b200041086a2202411076220041880b2802006a220336020041840b200241840b28020022026a41076a417871220436020002400240200341107420044d044041880b200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104104541086a0541000b200141106a24000b0b002000410120001b10490b4d01017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b200020012802082001280204104c20000b5a01027f02402002410a4d0440200020024101743a0000200041016a21030c010b200241106a4170712204104a21032000200236020420002004410172360200200020033602080b200320012002104d200220036a41003a00000b10002002044020002001200210451a0b0b130020002d0000410171044020002802081a0b0bf80101057f0240027f20002d0000220241017104402000280204210320002802002202417e71417f6a0c010b20024101762103410a0b220420032001200320014b1b220141106a417071417f6a410a2001410a4b1b2201460d00027f2001410a460440200041016a210420002802080c010b4100200120044d200141016a104a22041b0d0120002d0000220241017104404101210520002802080c010b41012105200041016a0b210620042006027f2002410171044020002802040c010b200241fe01714101760b41016a104d2005044020002004360208200020033602042000200141016a4101723602000f0b200020034101743a00000b0bf40101057f024002400240027f20002d00002202410171220345044020024101762104410a0c010b2000280204210420002802002202417e71417f6a0b22052004460440027f2002410171044020002802080c010b200041016a0b2106416f2103200541e6ffffff074d0440410b20054101742202200541016a220320032002491b220241106a4170712002410b491b21030b2003104a220220062005104d20002002360208200020034101723602000c010b2003450d01200028020821020b2000200441016a3602040c010b2000200441017441026a3a0000200041016a21020b200220046a220041003a0001200020013a00000b5a01017f02402003410a4d0440200020024101743a0000200041016a21030c010b200341106a4170712204104a21032000200236020420002004410172360200200020033602080b200320012002104d200220036a41003a00000b880101037f41e408410136020041e8082802002100034020000440034041ec0841ec082802002201417f6a2202360200200141014845044041e4084100360200200020024102746a22004184016a280200200041046a28020011010041e408410136020041e80828020021000c010b0b41ec08412036020041e808200028020022003602000c010b0b0b940101027f41e408410136020041e808280200220145044041e80841f00836020041f00821010b024041ec082802002202412046044041840210492201450d0120011046220141e80828020036020041e808200136020041ec084100360200410021020b41ec08200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41e40841003602000b3801017f41f40a420037020041fc0a410036020041742100034020000440200041800b6a4100360200200041046a21000c010b0b410410530b070041f40a104e0b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000105720012802044f0d002002410471044010000c010b200042003702000b02402002411071450d002000105720012802044d0d0020024104710440100020000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001058200010596a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000105a41012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a411410561057200241306a24000b2101017f20011059220220012802044b044010000b20002001200110582002105b0bd60202077f017e230041206b220324002001280208220420024b0440200341186a2001105d20012003280218200328021c105c36020c200341106a2001105d410021042001027f410020032802102206450d001a410020032802142208200128020c2207490d001a200820072007417f461b210520060b360210200141146a2005360200200141003602080b200141106a210903400240200420024f0d002001280214450d00200341106a2001105d41002104027f410020032802102207450d001a410020032802142208200128020c2206490d001a200820066b2104200620076a0b21052001200436021420012005360210200341106a2009410020052004105c105b20012003290310220a3702102001200128020c200a422088a76a36020c2001200128020841016a22043602080c010b0b20032009290200220a3703082003200a37030020002003411410561a200341206a24000b980101037f200028020445044041000f0b2000105a200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b2d00200020021061200028020020002802046a2001200210451a2000200028020420026a3602042000200310620b1b00200028020420016a220120002802084b04402000200110640b0b820201047f02402001450d00034020002802102202200028020c460d01200241786a28020020014904401000200028021021020b200241786a2203200328020020016b220136020020010d012000200336021020004101200028020422042002417c6a28020022016b2202103a220341016a20024138491b220520046a1063200120002802006a220420056a2004200210470240200241374d0440200028020020016a200241406a3a00000c010b200341f7016a220441ff014d0440200028020020016a20043a00002000280200200120036a6a210103402002450d02200120023a0000200241087621022001417f6a21010c000b000b10000b410121010c000b000b0b0f00200020011064200020013602040b2f01017f2000280208200149044020011049200028020020002802041045210220002001360208200020023602000b0b8d0201057f02402001044020002802042104200041106a2802002202200041146a280200220349044020022001ad2004ad422086843702002000200028021041086a3602100f0b027f41002002200028020c22026b410375220541016a2206200320026b2202410275220320032006491b41ffffffff01200241037541ffffffff00491b2202450d001a2002410374104a0b2103200320054103746a22052001ad2004ad4220868437020020052000280210200028020c22016b22046b2106200441014e044020062001200410451a200028020c21010b2000200320024103746a3602142000200541086a3602102000200636020c2001450d010f0b200041c0011066200041004100410110600b0b2500200041011061200028020020002802046a20013a00002000200028020441016a3602040b5e01027f2001103a220241b7016a22034180024e044010000b2000200341ff017110662000200028020420026a1063200028020420002802006a417f6a2100034020010440200020013a0000200141087621012000417f6a21000c010b0b0b7501027f2001280200210341012102024002400240024020012802042201410146044020032c000022014100480d012000200141ff017110660c040b200141374b0d01200121020b200020024180017341ff017110660c010b200020011067200121020b200020032002410010600b2000410110620bba0202037f037e02402001500440200041800110660c010b20014280015a044020012105034020052006845045044020064238862005420888842105200241016a2102200642088821060c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff017110662000200028020420046a1063200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff017110660b2000200028020420026a1063200028020420002802006a417f6a210203402001200784500d02200220013c0000200742388620014208888421012002417f6a2102200742088821070c000b000b20002001a741ff017110660b2000410110620b0b5c01004180080b5571707a7279397838676632747664773073336a6e35346b686365366d7561376c006c6174007472616e7366657200696e6974006d6967726174655f636f6e74726163740073657455696e74380067657455696e7438";

    private static String BINARY = BINARY_0;

    public static final String FUNC_MIGRATE_CONTRACT = "migrate_contract";

    public static final String FUNC_SETUINT8 = "setUint8";

    public static final String FUNC_GETUINT8 = "getUint8";

    protected ContractMigrate_old(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ContractMigrate_old(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<ContractMigrate_old> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Byte input) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList(input));
        return deployRemoteCall(ContractMigrate_old.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ContractMigrate_old> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Byte input) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList(input));
        return deployRemoteCall(ContractMigrate_old.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public RemoteCall<TransactionReceipt> migrate_contract(Byte[] init_arg, Long transfer_value, Long gas_value) {
        final WasmFunction function = new WasmFunction(FUNC_MIGRATE_CONTRACT, Arrays.asList(init_arg,transfer_value,gas_value), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> setUint8(Byte input) {
        final WasmFunction function = new WasmFunction(FUNC_SETUINT8, Arrays.asList(input), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<Byte> getUint8() {
        final WasmFunction function = new WasmFunction(FUNC_GETUINT8, Arrays.asList(), Byte.class);
        return executeRemoteCall(function, Byte.class);
    }

    public static ContractMigrate_old load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new ContractMigrate_old(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static ContractMigrate_old load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new ContractMigrate_old(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
