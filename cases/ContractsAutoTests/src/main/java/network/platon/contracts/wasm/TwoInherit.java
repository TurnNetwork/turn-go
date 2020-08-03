package network.platon.contracts.wasm;

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
 * <p>Generated with platon-web3j version 0.13.1.1.
 */
public class TwoInherit extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001540f60027f7f0060017f0060017f017f60027f7f017f60037f7f7f0060037f7f7f017f60000060047f7f7f7f017f60047f7f7f7f0060027f7e0060037e7e7f006000017f60027f7e017f60027e7e017f60017f017e02a9020d03656e760c706c61746f6e5f70616e6963000603656e760d726c705f6c6973745f73697a65000203656e760f706c61746f6e5f726c705f6c697374000403656e760e726c705f62797465735f73697a65000303656e7610706c61746f6e5f726c705f6279746573000403656e760d726c705f753132385f73697a65000d03656e760f706c61746f6e5f726c705f75313238000a03656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000b03656e7610706c61746f6e5f6765745f696e707574000103656e7617706c61746f6e5f6765745f73746174655f6c656e677468000303656e7610706c61746f6e5f6765745f7374617465000703656e7610706c61746f6e5f7365745f7374617465000803656e760d706c61746f6e5f72657475726e0000034d4c0600030307000104030402000102060205040102000e01020101020200020901000c0006070209050305040001030003000503000301010000020500000001020005070105040006040202080405017001030305030100020608017f01418089040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f7273000d0f5f5f66756e63735f6f6e5f65786974003006696e766f6b65001b0908010041010b0214160aa2634c040010540b8a0101047f230041206b2202240002402000411c6a2802002203200041206a220428020047044020032001100f1a2000200028021c413c6a36021c0c010b200241086a200041186a2205200320002802186b413c6d220041016a101020002004101122002802082001100f1a20002000280208413c6a360208200520001012200010130b200241206a24000b3f002000200110151a2000410c6a2001410c6a10151a200041186a200141186a10151a200041246a200141246a10151a200041306a200141306a10151a20000b2f01017f2001200028020820002802006b413c6d2200410174220220022001491b41c4889122200041a2c48811491b0b4c01017f2000410036020c200041106a2003360200200104402001101721040b20002004360200200020042002413c6c6a2202360208200020042001413c6c6a36020c2000200236020420000b900101027f200028020421022000280200210303402002200346450440200128020441446a200241446a220210182001200128020441446a3602040c010b0b200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b2b01027f20002802082101200028020421020340200120024704402000200141446a22013602080c010b0b0b1200200020012802182002413c6c6a10151a0ba10101037f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b20012802082103024020012802042201410a4d0440200020014101743a0000200041016a21020c010b200141106a4170712204101a21022000200136020420002004410172360200200020023602080b2002200320011055200120026a41003a000020000b1500200020012802182002413c6c6a41246a10151a0b09002000413c6c101a0ba4010020002001290200370200200041086a200141086a28020036020020011019200041146a200141146a2802003602002000200129020c37020c2001410c6a1019200041206a200141206a28020036020020002001290218370218200141186a10192000412c6a2001412c6a28020036020020002001290224370224200141246a1019200041386a200141386a28020036020020002001290230370230200141306a10190b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0b0b002000410120001b101c0ba50402057f017e230041c0016b22002400105410072201101c2203100820004180016a200020032001101d22024100101e20004180016a101f0240024020004180016a1020450d00200028028401450d002000280280012d000041c001490d010b10000b200041406b20004180016a10212000280244220141094f044010000b200028024021030340200104402001417f6a210120033100002005420886842105200341016a21030c010b0b024002402005500d00418008102220055104402002102320004180016a102410250c020b41850810222005510440200041406b1026200041406b1027210120021028410247044010000b20004180016a20024101101e20004180016a20011029200041186a10242103200041186a20004180016a2001100f100e200310250c020b419808102220055104402002102320004180016a102421032000419c016a28020021022000280298012104200041186a102a2101200041d8006a4100360200200041d0006a4200370300200041c8006a420037030020004200370340200041406b200220046b413c6d41ff0171ad2205102b20002802402102200041406b410472102c20012002102d20012005102e220128020c200141106a28020047044010000b20012802002001280204100c200128020c22020440200120023602100b200310250c020b41b0081022200551044020024101102f0c020b41c80810222005520d0020024102102f0c010b10000b1030200041c0016a24000b9b0101047f230041106b220124002001200036020c2000047f41f408200041086a2202411076220041f4082802006a220336020041f00841f008280200220420026a41076a417871220236020002400240200341107420024d044041f408200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20042001410c6a4104103e41086a0541000b2100200141106a240020000b0c00200020012002411c10310bc90202067f017e230041106b220324002001280208220520024b0440200341086a2001104420012003280208200328020c103536020c200320011044410021052001027f410020032802002207450d001a410020032802042208200128020c2206490d001a200820062006417f461b210420070b360210200141146a2004360200200141003602080b200141106a210603402001280214210402402005200249044020040d01410021040b200020062802002004411410311a200341106a24000f0b20032001104441002104027f410020032802002205450d001a410020032802042208200128020c2207490d001a200820076b2104200520076a0b2105200120043602142001200536021020032006410020052004103510582001200329030022093702102001200128020c2009422088a76a36020c2001200128020841016a22053602080c000b000b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0b980101037f200028020445044041000f0b2000101f200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0bd50101047f200110322204200128020422024b04401000200128020421020b200128020021052000027f02400240200204404100210120052c00002203417f4a0d01027f200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a21010c010b4101210120050d000c010b41002103200120046a20024b0d0020022001490d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b0e0020001028410147044010000b0be606010c7f230041c0016b2204240020004200370218200042b1fdc48289b7f690b67f3703102000410036020820004200370200200041206a4100360200200441186a102a220720002903101033200728020c200741106a28020047044010000b200041186a21052000411c6a21090240200728020022022007280204220810092206450d002006101a21030340200120036a41003a00002006200141016a2201470d000b2002200820032001100a417f460440410021010c010b024002402004200341016a200120036a2003417f736a101d2203280204450d0020032802002d000041c001490d002003102821012000280220200028021822026b413c6d20014904402005200441306a2001200028021c20026b413c6d200041206a101122011012200110130b20044198016a200341011034210120044188016a2003410010342108200041206a210b20012802042103034020082802042003464100200128020822022008280208461b0d02200441f0006a20032002411c1031200441306a1027220310290240200028021c2202200028022049044020022003101820092009280200413c6a3602000c010b200441a8016a2005200220052802006b413c6d220241016a10102002200b1011210220042802b00120031018200420042802b001413c6a3602b001200520021012200210130b20012001280204220320012802086a410020031b220336020420012802002202044020012002360208200320021035210a2001027f2001280204220c4504404100210241000c010b41002102410020012802082203200a490d001a2003200a200a417f461b2102200c0b2203ad2002ad42208684370204200141002001280200220a20026b22022002200a4b1b3602000c0105200141003602080c010b000b000b10000b200621010b200728020c22030440200720033602100b024020010d0020002802042206200028020022036b413c6d22022000280220200028021822016b413c6d4d04402002200928020020016b413c6d22084b0440200320032008413c6c6a2202200110361a20022006200910370c020b2005200320062001103610380c010b200104402005103920004100360220200042003702180b20002005200210102202101722013602182000200136021c200020012002413c6c6a36022020032006200910370b200441c0016a240020000bf20701107f230041f0016b22012400200141186a102a2106200141e8016a22024100360200200141e0016a22054200370300200141d8016a22044200370300200142003703d001200141d0016a2000290310102b20012802d0012107200141d0016a410472102c20062007102d200620002903101033200628020c200641106a28020047044010000b200628020421092006280200210a2001102a2103200241003602002005420037030020044200370300200142003703d001027f20002802182000411c6a280200460440200141013602d00141010c010b200141d0016a4100103a2107200028021c200028021822026b2105037f2005047f20074100103a22042002103b200420014190016a200241246a1015103c200141d0006a200241306a1015103c4101103a1a200541446a21052002413c6a21020c010520074101103a1a20012802d0010b0b0b2104200141d0016a410472102c4101101a220241fe013a0000200328020c200341106a28020047044010000b200241016a21072003280204220541016a220820032802084b047f20032008103d20032802040520050b20032802006a20024101103e1a2003200328020441016a3602042003200420026b20076a102d2003200028021c20002802186b413c6d103f210b200028021c200028021822026b2105200141d0006a410472210c20014190016a410472210d200141d0016a410472210e034020050440200b4103103f210420014100360268200142003703602001420037035820014200370350200141d0006a2002103b200141d0006a200141406b200241246a22071015103c200141306a200241306a22081015103c1a20042001280250102d20044103103f2104200141003602a801200142003703a0012001420037039801200142003703900120014190016a2002104020014190016a20014180016a2002410c6a220f1015103c200141f0006a200241186a22101015103c1a2004200128029001102d20044101103f2104200141003602e801200142003703e001200142003703d801200142003703d001200141d0016a200141c0016a20021015103c1a200420012802d001102d2004200141b0016a2002101510412104200e102c2004200141d0016a200f10151041200141c0016a2010101510412104200d102c2004200141d0016a20071015104120014190016a2008101510411a200c102c200541446a21052002413c6a21020c010b0b0240200328020c2003280210460440200328020021020c010b100020032802002102200328020c2003280210460d0010000b200a200920022003280204100b200328020c22020440200320023602100b200628020c22020440200620023602100b200041186a104220001042200141f0016a24000bc50201027f200041003a00002000413c6a2202417f6a41003a0000200041003a0002200041003a00012002417d6a41003a00002002417e6a41003a0000200041003a00032002417c6a41003a00002000410020006b41037122016a220241003602002002413c20016b417c7122016a2200417c6a4100360200024020014109490d002002410036020820024100360204200041786a4100360200200041746a410036020020014119490d002002410036021820024100360214200241003602102002410036020c200041706a41003602002000416c6a4100360200200041686a4100360200200041646a41003602002001200241047141187222016b2100200120026a2101034020004120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200041606a21000c000b000b0b2400200010432000410c6a1043200041186a1043200041246a1043200041306a104320000b800101047f230041106b2201240002402000280204450d0020002802002d000041c001490d00200141086a20001044200128020c210003402000450d01200141002001280208220320032000103522046a20034520002004497222031b3602084100200020046b20031b2100200241016a21020c000b000b200141106a240020020ba80101017f230041d0006b22022400200241086a20004100101e200241206a200241086a4100101e200241386a200241206a4100101e200241386a20011045200241386a200241086a4101101e200241386a2001410c6a1045200241386a200241086a4102101e200241386a200141186a1045200241386a20004101101e200241386a200141246a1045200241386a20004102101e200241386a200141306a1045200241d0006a24000b2900200041003602082000420037020020004100103d200041146a41003602002000420037020c20000b840102027f017e4101210320014280015a0440034020012004845045044020044238862001420888842101200241016a2102200442088821040c010b0b200241384f047f2002104620026a0520020b41016a21030b200041186a28020022020440200041086a280200200041146a2802002002104721000b2000200028020020036a3602000bea0101047f230041106b22042400200028020422012000280210220241087641fcffff07716a2103027f410020012000280208460d001a2003280200200241ff07714102746a0b2101200441086a20001048200428020c210203400240200120024604402000410036021420002802082103200028020421010340200320016b41027522024103490d022000200141046a22013602040c000b000b200141046a220120032802006b418020470d0120032802042101200341046a21030c010b0b2002417f6a220241014d04402000418004418008200241016b1b3602100b200020011049200441106a24000b13002000280208200149044020002001103d0b0b2a01017f2000420020011005200028020422026a104a42002001200220002802006a10062000104b20000bcb0201037f23004180016b2202240020001028410247044010000b200220004101101e2002101f0240024020021020450d002002280204450d0020022802002d000041c001490d010b10000b200241e0006a200210212002280264220041024f044010000b200228026021030340200004402000417f6a210020032d00002104200341016a21030c010b0b200210242103200241286a2002200441ff01712001110400200241386a102a2100200241f8006a4100360200200241f0006a4200370300200241e8006a420037030020024200370360200241e0006a200241d0006a200241286a1015103c210420022802602101200441046a102c20002001102d2000200241e0006a200241286a10151041220028020c200041106a28020047044010000b20002802002000280204100c200028020c22040440200020043602100b2003102520024180016a24000b880101037f41e008410136020041e4082802002100034020000440034041e80841e8082802002201417f6a2202360200200141014845044041e0084100360200200020024102746a22004184016a280200200041046a28020011010041e008410136020041e40828020021000c010b0b41e808412036020041e408200028020022003602000c010b0b0b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000105620024f0d002003410471044010000c010b200042003702000b02402003411071450d002000105620024d0d0020034104710440100020000f0b200042003702000b20000bff0201037f200028020445044041000f0b2000101f41012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b090020002001102e1a0be70101037f230041106b2204240020004200370200200041086a410036020020012802042103024002402002450440200321020c010b410021022003450d002003210220012802002d000041c001490d00200441086a2001104420004100200428020c2201200428020822022001103522032003417f461b20024520012003497222031b220536020820004100200220031b3602042000200120056b3602000c010b20012802002103200128020421012000410036020020004100200220016b20034520022001497222021b36020820004100200120036a20021b3602040b200441106a240020000b2701017f230041206b22022400200241086a200020014114103110562100200241206a240020000b6901037f200120006b21054100210103402001200546450440200120026a2203200020016a220410532003410c6a2004410c6a1053200341186a200441186a1053200341246a200441246a1053200341306a200441306a10532001413c6a21010c010b0b200120026a0b2e000340200020014645044020022802002000100f1a20022002280200413c6a3602002000413c6a21000c010b0b0b0900200020013602040b0b002000200028020010380bc30c02077f027e230041306b22042400200041046a2107027f20014101460440200041086a280200200041146a280200200041186a220228020022031047280200210120022003417f6a3602002007104c4180104f044020072000410c6a280200417c6a10490b200141384f047f2001104620016a0520010b41016a2102200041186a28020022010440200041086a280200200041146a280200200110470c020b20000c010b02402007104c0d00200041146a28020022014180084f0440200020014180786a360214200041086a2201280200220228020021032001200241046a360200200420033602182007200441186a104d0c010b2000410c6a2802002202200041086a2802006b4102752203200041106a2205280200220620002802046b2201410275490440418020101a2105200220064704400240200028020c220120002802102202470d0020002802082203200028020422064b04402000200320012003200320066b41027541016a417e6d41027422026a104e220136020c2000200028020820026a3602080c010b200441186a200220066b2201410175410120011b22012001410276200041106a104f2102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021050200028020c21010b200120053602002000200028020c41046a36020c0c020b02402000280208220120002802042202470d00200028020c2203200028021022064904402000200120032003200620036b41027541016a41026d41027422026a105122013602082000200028020c20026a36020c0c010b200441186a200620026b2201410175410120011b2201200141036a410276200041106a104f2102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021050200028020821010b2001417c6a2005360200200020002802082201417c6a22023602082002280200210220002001360208200420023602182007200441186a104d0c010b20042001410175410120011b20032005104f2102418020101a2106024020022802082201200228020c2203470d0020022802042205200228020022084b04402002200520012005200520086b41027541016a417e6d41027422036a104e22013602082002200228020420036a3602040c010b200441186a200320086b2201410175410120011b22012001410276200241106a280200104f21032002280208210520022802042101034020012005470440200328020820012802003602002003200328020841046a360208200141046a21010c010b0b20022902002109200220032902003702002003200937020020022902082109200220032902083702082003200937020820031050200228020821010b200120063602002002200228020841046a360208200028020c2105034020002802082005460440200028020421012000200228020036020420022001360200200228020421012002200536020420002001360208200029020c21092000200229020837020c2002200937020820021050052005417c6a210502402002280204220120022802002203470d0020022802082206200228020c22084904402002200120062006200820066b41027541016a41026d41027422036a105122013602042002200228020820036a3602080c010b200441186a200820036b2201410175410120011b2201200141036a4102762002280210104f21062002280208210320022802042101034020012003470440200428022020012802003602002004200428022041046a360220200141046a21010c010b0b20022902002109200220042903183702002002290208210a20022004290320370208200420093703182004200a37032020061050200228020421010b2001417c6a200528020036020020022002280204417c6a3602040c010b0b0b200441186a20071048200428021c410036020041012102200041186a0b2201200128020020026a360200200441306a240020000b4001017f230041206b2202240020004100103a2200200110402000200241106a2001410c6a1015103c2002200141186a1015103c4101103a1a200241206a24000ba10101037f41012103024002400240200128020420012d00002202410176200241017122041b220241014d0440200241016b0d032001280208200141016a20041b2c0000417f4c0d010c030b200241374b0d010b200241016a21030c010b2002104620026a41016a21030b027f200041186a28020022010440200041086a280200200041146a280200200110470c010b20000b2201200128020020036a36020020000b2f01017f200028020820014904402001101c20002802002000280204103e210220002001360208200020023602000b0bfc0801067f03400240200020046a2105200120046a210320022004460d002003410371450d00200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220745044003402006411049450440200020046a2203200120046a2205290200370200200341086a200541086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2205200120046a2204290200370200200441086a2103200541086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002007417f6a220741024b0d00024002400240024002400240200741016b0e020102000b2005200120046a220328020022073a0000200541016a200341016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2203200120046a220541046a2802002202410874200741187672360200200341046a200541086a2802002207410874200241187672360200200341086a2005410c6a28020022024108742007411876723602002003410c6a200541106a2802002207410874200241187672360200200441106a2104200641706a21060c000b000b2005200120046a220328020022073a0000200541016a200341016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2203200120046a220541046a2802002202411074200741107672360200200341046a200541086a2802002207411074200241107672360200200341086a2005410c6a28020022024110742007411076723602002003410c6a200541106a2802002207411074200241107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022073a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2203200120046a220541046a2802002202411874200741087672360200200341046a200541086a2802002207411874200241087672360200200341086a2005410c6a28020022024118742007410876723602002003410c6a200541106a2802002207411874200241087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000b9d0201057f2001044020002802042105200041106a2802002202200041146a280200220349044020022001ad2005ad422086843702002000200028021041086a36021020000f0b027f41002002200028020c22046b410375220641016a2202200320046b2203410275220420042002491b41ffffffff01200341037541ffffffff00491b2204450d001a2004410374101a0b2102200220064103746a22032001ad2005ad4220868437020020032000280210200028020c22066b22016b2105200220044103746a2102200341086a2103200141014e0440200520062001103e1a0b20002002360214200020033602102000200536020c20000f0b200041001001200028020422016a104a41004100200120002802006a10022000104b20000b2701017f230041106b2202240020004100103a200220011015103c4101103a1a200241106a24000b4e01037f20002001280208200141016a20012d0000220241017122031b22042001280204200241017620031b22011003200028020422026a104a20042001200220002802006a10042000104b20000b0e0020002802000440200010390b0b170020004200370200200041086a4100360200200010190b2101017f20011032220220012802044b044010000b2000200120011057200210580bf30201057f230041206b22022400024002402000280204044020002802002d000041c001490d010b200241086a10430c010b200241186a200010212000103221030240024002400240200228021822000440200228021c220520034f0d010b41002100200241106a410036020020024200370308410021050c010b200241106a4100360200200242003703082000200520032003417f461b22046a21052004410a4b0d010b200220044101743a0008200241086a41017221030c010b200441106a4170712206101a21032002200436020c20022006410172360208200220033602100b03402000200546450440200320002d00003a0000200341016a2103200041016a21000c010b0b200341003a00000b024020012d0000410171450440200141003b01000c010b200128020841003a00002001410036020420012d0000410171450d00200141003602000b20012002290308370200200141086a200241106a280200360200200241086a1019200241206a24000b1e01017f03402000044020004108762100200141016a21010c010b0b20010b25002000200120026a417f6a220241087641fcffff07716a280200200241ff07714102746a0b4f01037f20012802042203200128021020012802146a220441087641fcffff07716a21022000027f410020032001280208460d001a2002280200200441ff07714102746a0b360204200020023602000b2501017f200028020821020340200120024645044020002002417c6a22023602080c010b0b0b3601017f200028020820014904402001101c20002802002000280204103e210220002001360208200020023602000b200020013602040b7a01037f0340024020002802102201200028020c460d00200141786a2802004504401000200028021021010b200141786a22022002280200417f6a220336020020030d002000200236021020002001417c6a2802002201200028020420016b220210016a104a200120002802006a22012002200110020c010b0b0b2801017f200028020820002802046b2201410874417f6a410020011b200028021420002802106a6b0ba10202057f017e230041206b22052400024020002802082202200028020c2203470d0020002802042204200028020022064b04402000200420022004200420066b41027541016a417e6d41027422036a104e22023602082000200028020420036a3602040c010b200541086a200320066b2202410175410120021b220220024102762000410c6a104f2103200028020821042000280204210203402002200446450440200328020820022802003602002003200328020841046a360208200241046a21020c010b0b20002902002107200020032902003702002003200737020020002902082107200020032902083702082003200737020820031050200028020821020b200220012802003602002000200028020841046a360208200541206a24000b2501017f200120006b220141027521032001044020022000200110520b200220034102746a0b4f01017f2000410036020c200041106a2003360200200104402001410274101a21040b200020043602002000200420024102746a22023602082000200420014102746a36020c2000200236020420000b2b01027f200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b0b1b00200120006b22010440200220016b22022000200110520b20020b8d0301037f024020002001460d00200120006b20026b410020024101746b4d0440200020012002103e1a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2105200020036a2204410371450440200220036b210241002103034020024104490d04200320046a200320056a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200420052d00003a0000200341016a21030c000b000b024020030d002001417f6a21040340200020026a22034103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042003417f6a200220046a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320056a2101200320046a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b8c0201047f20002001470440200128020420012d00002202410176200241017122041b2102200141016a210320012802082105410a21012005200320041b210420002d0000410171220304402000280200417e71417f6a21010b200220014d0440027f2003044020002802080c010b200041016a0b21012002044020012004200210520b200120026a41003a000020002d00004101710440200020023602040f0b200020024101743a00000f0b416f2103200141e6ffffff074d0440410b20014101742201200220022001491b220141106a4170712001410b491b21030b2003101a2201200420021055200020023602042000200341017236020020002001360208200120026a41003a00000b0b3501017f230041106b22004180890436020c41ec08200028020c41076a417871220036020041f008200036020041f4083f003602000b100020020440200020012002103e1a0b0b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001057200010326a0520010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0b5b01027f2000027f0240200128020022054504400c010b200220036a200128020422014b0d0020012002490d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b0b6601004180080b5f696e6974006164645f7375625f6d795f6d657373616765006765745f7375625f6d795f6d6573736167655f73697a65006765745f7375625f6d795f6d6573736167655f68656164006765745f7375625f6d795f6d6573736167655f66726f6d";

    public static String BINARY = BINARY_0;

    public static final String FUNC_ADD_SUB_MY_MESSAGE = "add_sub_my_message";

    public static final String FUNC_GET_SUB_MY_MESSAGE_SIZE = "get_sub_my_message_size";

    public static final String FUNC_GET_SUB_MY_MESSAGE_FROM = "get_sub_my_message_from";

    public static final String FUNC_GET_SUB_MY_MESSAGE_HEAD = "get_sub_my_message_head";

    protected TwoInherit(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    protected TwoInherit(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }

    public static RemoteCall<TwoInherit> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(TwoInherit.class, web3j, credentials, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<TwoInherit> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(TwoInherit.class, web3j, transactionManager, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<TwoInherit> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(TwoInherit.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public static RemoteCall<TwoInherit> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(TwoInherit.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public RemoteCall<TransactionReceipt> add_sub_my_message(Sub_my_message sub_one_message) {
        final WasmFunction function = new WasmFunction(FUNC_ADD_SUB_MY_MESSAGE, Arrays.asList(sub_one_message), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> add_sub_my_message(Sub_my_message sub_one_message, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_ADD_SUB_MY_MESSAGE, Arrays.asList(sub_one_message), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<Uint8> get_sub_my_message_size() {
        final WasmFunction function = new WasmFunction(FUNC_GET_SUB_MY_MESSAGE_SIZE, Arrays.asList(), Uint8.class);
        return executeRemoteCall(function, Uint8.class);
    }

    public RemoteCall<String> get_sub_my_message_from(Uint8 index) {
        final WasmFunction function = new WasmFunction(FUNC_GET_SUB_MY_MESSAGE_FROM, Arrays.asList(index), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<String> get_sub_my_message_head(Uint8 index) {
        final WasmFunction function = new WasmFunction(FUNC_GET_SUB_MY_MESSAGE_HEAD, Arrays.asList(index), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static TwoInherit load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        return new TwoInherit(contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    public static TwoInherit load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        return new TwoInherit(contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }

    public static class Message {
        public String head;
    }

    public static class My_message {
        public Message baseClass;

        public String body;

        public String end;
    }

    public static class Sub_my_message {
        public My_message baseClass;

        public String from;

        public String to;
    }
}
