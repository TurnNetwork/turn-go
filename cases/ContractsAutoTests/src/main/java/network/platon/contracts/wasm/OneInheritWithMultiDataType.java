package network.platon.contracts.wasm;

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
public class OneInheritWithMultiDataType extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001480d60027f7f0060017f0060017f017f60027f7f017f60037f7f7f0060037f7f7f017f60000060047f7f7f7f0060047f7f7f7f017f60027f7e0060027f7e017f60017f017e6000017f02a9010703656e760c706c61746f6e5f70616e6963000603656e760d706c61746f6e5f72657475726e000003656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000c03656e7610706c61746f6e5f6765745f696e707574000103656e7617706c61746f6e5f6765745f73746174655f6c656e677468000303656e7610706c61746f6e5f6765745f7374617465000803656e7610706c61746f6e5f7365745f73746174650007036160060101010003030800010404050b00000202030100030101060b0200010a03030209010000010902020000020005040001000300010200050801050503040602020304010006010601050202020101030705000402020700000000000000030a0405017001050505030100020608017f0141b08b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070f5f5f66756e63735f6f6e5f65786974004c06696e766f6b65001f090a010041010b0411120a4f0abb72600f0041f80810084103104d1045104e0b170020004200370200200041086a4100360200200010090b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0b070041f808104a0b940101047f230041206b2202240002402000411c6a2802002203200041206a220428020047044020032001100c1a2000200028021c41306a36021c0c010b200241086a200041186a2205200320002802186b41306d41016a100d200028021c20002802186b41306d2004100e22002802082001100c1a2000200028020841306a36020820052000100f200010100b200241206a24000b3f002000200110481a200041146a200141146a2802003602002000200129020c37020c200041186a200141186a10481a200041246a200141246a10481a20000b3b01017f200141d6aad52a4f0440000b2001200028020820002802006b41306d2200410174220220022001491b41d5aad52a200041aad5aa15491b0b4c01017f2000410036020c200041106a2003360200200104402001103221040b2000200436020020002004200241306c6a220236020820002004200141306c6a36020c2000200236020420000b900101027f200028020421022000280200210303402002200346450440200128020441506a200241506a220210332001200128020441506a3602040c010b0b200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b3301027f2000280204210203402002200028020822014704402000200141506a2201360208200110230c010b0b20002802001a0b120020002001280218200241306c6a10481a0b150020002001280218200241306c6a41186a10481a0b3401017f230041106b220324002003200236020c200320013602082003200329030837030020002003411c1050200341106a24000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b8e0201037f23004190016b220224002001280200210320012802042101200241003a00302000200241306a1016200241086a1017200241386a200241086a20014101756a220020022d00302001410171047f200028020020036a2802000520030b110400200241c8006a1018210020024188016a410036020020024180016a4200370300200241f8006a420037030020024200370370200241f0006a200241e0006a200241386a1048220110192001104a2002280270210141046a101a20002001101b2000200241f0006a200241386a10482201101c1a2001104a200028020c200041106a28020047044010000b2000280200200028020410012000101d200241386a104a101e20024190016a24000b8a0101037f230041206b22022400200220004101105a20021054024002402002105b450d002002280204450d0020022802002d000041c001490d010b10000b200241186a2002102a200228021c220041024f044010000b200228021821030340200004402000417f6a210020032d00002104200341016a21030c010b0b200120043a0000200241206a24000bfc06020c7f017e230041d0016b2202240020004200370218200042c5cb94caeffff0b2a67f3703102000410036020820004200370200200041206a4100360200200241386a1018220720002903101028200728020c200741106a28020047044010000b200041186a21052000411c6a2109024002402007280200220b2007280204220c10042204044020024100360230200242003703282004417f4c0d012004104721030340200120036a41003a00002004200141016a2201470d000b200120036a21082003200228022c200228022822066b220a6b2101200a41014e044020012006200a10421a200228022821060b2002200320046a3602302002200836022c200220013602280240200b200c20060440200228022c2108200228022821010b2001200820016b1005417f460440410021040c010b0240200241106a2002280228220141016a200228022c2001417f736a10132201280204450d0020012802002d000041c001490d002001105c21032000280220200028021822086b41306d20034904402005200241d0006a2003200028021c20086b41306d200041206a100e2203100f200310100b200241a8016a2001410110582103200041206a210a20024198016a200141001058210803402003280204200828020446044020032802082008280208460d030b20022003290204220d3703502002200d37030820024180016a200241086a411c1050200241d0006a1021220110220240200028021c220620002802204904402006200110332009200928020041306a3602000c010b200241b8016a2005200620052802006b41306d41016a100d200928020020052802006b41306d200a100e210620022802c00120011033200220022802c00141306a3602c00120052006100f200610100b20011023200310550c000b000b10000b200241286a1029200421010b2007101d024020010d0020002802042207200028020022016b41306d22032000280220200028021822046b41306d4d04402003200928020020046b41306d22034b044020012001200341306c6a2201200410341a20012007200910350c020b2005200120072004103410360c010b200404402005103720002802181a20004100360220200042003702180b20052003100d220441d6aad52a4f0d0220002004103222053602182000200536021c20002005200441306c6a36022020012007200910350b200241d0016a240020000f0b000b000b2900200041003602082000420037020020004100102b200041146a41003602002000420037020c20000b900101037f410121030240200128020420012d00002202410176200241017122041b2202450d0002400240200241014604402001280208200141016a20041b2c0000417f4c0d010c030b200241374b0d010b200241016a21030c010b2002102e20026a41016a21030b027f200041186a2802000440200041046a102f0c010b20000b2201200128020020036a36020020000b940201047f230041106b22042400200028020422012000280210220241087641fcffff07716a2103027f410020012000280208460d001a2003280200200241ff07714102746a0b2101200441086a20001030200428020c21020340024020012002460440200041003602142000280204210103402000280208220320016b41027522024103490d0220012802001a2000200028020441046a22013602040c000b000b200141046a220120032802006b418020470d0120032802042101200341046a21030c010b0b2002417f6a220241014d04402000418004418008200241016b1b3602100b03402001200347044020012802001a200141046a21010c010b0b20002000280204103120002802001a200441106a24000b13002000280208200149044020002001102b0b0b5201037f230041106b2202240020022001280208200141016a20012d0000220341017122041b36020820022001280204200341017620041b36020c20022002290308370300200020021065200241106a24000b1c01017f200028020c22010440200041106a20013602000b2000102c0b8407010d7f230041c0016b22012400200141286a10182209200041106a1027101b200920002903101028200928020c200941106a28020047044010000b2009280204210a2009280200200141106a10182103200141b8016a4100360200200141b0016a4200370300200141a8016a4200370300200142003703a001027f20002802182000411c6a280200460440200141013602a00141010c010b200141a0016a410010392106200028021c200028021822026b2104037f2004047f20064100103922072002103a2007200141e0006a200241186a10482207101920014190016a200241246a1048220810192008104a2007104a410110391a200441506a2104200241306a21020c01052006410110391a20012802a0010b0b0b2108200141a0016a410472101a41011047220241fe013a0000200120023602002001200241016a220436020820012004360204200328020c200341106a280200470440100020012802042104200128020021020b2002210620032802042207200420026b22046a220520032802084b044020032005102b20032802042107200128020021060b200328020020076a2002200410421a2003200328020420046a36020420032001280204200820066b6a101b2003200028021c20002802186b41306d1062200028021c200028021822026b2104200141e0006a4104722106200141a0016a410472210703402004044020034103106220014100360278200142003703702001420037036820014200370360200141e0006a2002103a200141e0006a200141d0006a200241186a2208104822051019200141406b200241246a220c1048220d10191a200d104a2005104a20032001280260101b200341031062200141003602b801200142003703b001200142003703a801200142003703a001200141a0016a20014190016a2002104822051019200228020c10252002290310102d2005104a200320012802a001101b200320014180016a200210482205101c200228020c1026200229031010282005104a2007101a2003200141a0016a200810482208101c20014190016a200c10482205101c1a2005104a2008104a2006101a200441506a2104200241306a21020c010b0b0240200328020c2003280210460440200328020021020c010b100020032802002102200328020c2003280210460d0010000b200a200220032802041006200110292003101d2009101d200041186a103b2000103b200141c0016a24000b880602067f017e230041c0016b2200240010071002220110462202100320004188016a200041186a20022001101322014100105a0240024020004188016a10202206500d004180081014200651044020004188016a1017101e0c020b41850810142006510440200041d8006a413010431a200041d8006a1021210220004188016a20014101105a20004188016a20021022200041306a1017200041306a20004188016a2002100c2203100b20031023101e200210230c020b4194081014200651044020004188016a1017200041a4016a280200210320002802a0012104200041306a10182101200041f0006a4100360200200041e8006a4200370300200041e0006a420037030020004200370358200041d8006a200320046b41306d41ff0171ad220610242000280258210441046a101a20012004101b2001200610661a200128020c200141106a28020047044010000b2001280200200128020410012001101d101e0c020b41a808101420065104402000410036028c01200041013602880120002000290388013703082001200041086a10150c020b41bc0810142006510440200041003a00b8012001200041b8016a101620004188016a101720002802a00120002d00b80141306c6a28020c2102200041306a10182101200041f0006a4100360200200041e8006a4200370300200041e0006a420037030020004200370358200041d8006a200210252000280258210541046a101a20012005101b2001200210261a200128020c200141106a28020047044010000b2001280200200128020410012001101d101e0c020b41cf0810142006510440200041003a00b8012001200041b8016a101620004188016a1017200020002802a00120002d00b80141306c6a2903102206370330200041d8006a10182201200041306a1027101b200120061028200128020c200141106a28020047044010000b2001280200200128020410012001101d101e0c020b41e40810142006520d002000410036028c01200041023602880120002000290388013703102001200041106a10150c010b10000b104c200041c0016a24000b850102027f017e230041106b2201240020001054024002402000105b450d002000280204450d0020002802002d000041c001490d010b10000b200141086a2000102a200128020c220041094f044010000b200128020821020340200004402000417f6a210020023100002003420886842103200241016a21020c010b0b200141106a240020030b160020001008200041186a1008200041246a100820000b890201047f230041406a22022400200241086a20004100105a200241206a200241086a4100105a200241206a20011038200241206a200241086a4101105a200241206a105402400240200241206a105b450d002002280224450d0020022802202d000041c001490d010b10000b200241386a200241206a102a200228023c220341054f044010000b200228023821050340200304402003417f6a210320052d00002004410874722104200541016a21050c010b0b2001200436020c200241206a200241086a4102105a2001200241206a1020370310200241206a20004101105a200241206a200141186a1038200241206a20004102105a200241206a200141246a1038200241406b24000b1400200041246a104a200041186a104a2000104a0b7c02027f017e4101210320014280015a0440034020012004845045044020044238862001420888842101200241016a2102200442088821040c010b0b200241384f047f2002102e20026a0520020b41016a21030b027f200041186a2802000440200041046a102f0c010b20000b2202200228020020036a36020020000b090020002001ad10240b090020002001ad10660b4e01017f230041206b22012400200141186a4100360200200141106a4200370300200141086a42003703002001420037030020012000290300102d20012802002001410472101a200141206a24000b09002000200110661a0b1501017f200028020022010440200020013602040b0bd60101047f200110532204200128020422024b04401000200128020421020b20012802002105027f027f41002002450d001a410020052c00002203417f4a0d011a200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a0b21012000027f02402005450440410021030c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b3401017f200028020820014904402001104622022000280200200028020410421a2000102c20002001360208200020023602000b0b080020002802001a0b09002000200110241a0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0b4f01037f20012802042203200128021020012802146a220441087641fcffff07716a21022000027f410020032001280208460d001a2002280200200441ff07714102746a0b360204200020023602000b2501017f200028020821020340200120024645044020002002417c6a22023602080c010b0b0b1500200041d6aad52a4f0440000b200041306c10470b7c0020002001290200370200200041086a200141086a28020036020020011009200041146a200141146a2802003602002000200129020c37020c20002001290218370218200041206a200141206a280200360200200141186a10092000412c6a2001412c6a28020036020020002001290224370224200141246a10090b6500200120006b210103402001044020022000104b200241146a200041146a2802003602002002410c6a200029020c370200200241186a200041186a104b200241246a200041246a104b200241306a2102200141506a2101200041306a21000c010b0b20020b2e000340200020014645044020022802002000100c1a2002200228020041306a360200200041306a21000c010b0b0b2901017f2000280204210203402001200246450440200241506a220210230c010b0b200020013602040b0b002000200028020010360b8b0301057f230041206b220224000240024002402000280204044020002802002d000041c001490d010b200241086a10080c010b200241186a2000102a2000105321030240024002400240200228021822000440200228021c220420034f0d010b41002100200241106a410036020020024200370308410021030c010b200241106a410036020020024200370308200420032003417f461b220341704f0d04200020036a21052003410a4b0d010b200220034101743a0008200241086a41017221040c010b200341106a4170712206104721042002200336020c20022006410172360208200220043602100b034020002005470440200420002d00003a0000200441016a2104200041016a21000c010b0b200441003a00000b024020012d0000410171450440200141003b01000c010b200128020841003a00002001410036020420012d0000410171450d0020012802081a200141003602000b20012002290308370200200141086a200241106a280200360200200241086a1009200241086a104a200241206a24000f0b000ba60c02077f027e230041306b22052400200041046a2107027f200141014604402007102f2802002101200041186a22022002280200417f6a3602002007103c4180104f04402000410c6a2202280200417c6a2802001a20072002280200417c6a10310b200141384f047f2001102e20016a0520010b41016a2101200028021804402007102f0c020b20000c010b02402007103c0d00200041146a28020022014180084f0440200020014180786a360214200041086a2201280200220228020021042001200241046a360200200520043602182007200541186a103d0c010b2000410c6a2802002202200041086a2802006b4102752204200041106a2203280200220620002802046b220141027549044041802010472104200220064704400240200028020c220120002802102206470d0020002802082202200028020422034b04402000200220012002200220036b41027541016a417e6d41027422036a103e220136020c2000200028020820036a3602080c010b200541186a200620036b2201410175410120011b22012001410276200041106a103f2102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021040200028020c21010b200120043602002000200028020c41046a36020c0c020b02402000280208220120002802042206470d00200028020c2202200028021022034904402000200120022002200320026b41027541016a41026d41027422036a104122013602082000200028020c20036a36020c0c010b200541186a200320066b2201410175410120011b2201200141036a410276200041106a103f2102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021040200028020821010b2001417c6a2004360200200020002802082201417c6a22023602082002280200210220002001360208200520023602182007200541186a103d0c010b20052001410175410120011b20042003103f210241802010472106024020022802082201200228020c2208470d0020022802042204200228020022034b04402002200420012004200420036b41027541016a417e6d41027422036a103e22013602082002200228020420036a3602040c010b200541186a200820036b2201410175410120011b22012001410276200241106a280200103f21042002280208210320022802042101034020012003470440200428020820012802003602002004200428020841046a360208200141046a21010c010b0b20022902002109200220042902003702002004200937020020022902082109200220042902083702082004200937020820041040200228020821010b200120063602002002200228020841046a360208200028020c2104034020002802082004460440200028020421012000200228020036020420022001360200200228020421012002200436020420002001360208200029020c21092000200229020837020c2002200937020820021040052004417c6a210402402002280204220120022802002208470d0020022802082203200228020c22064904402002200120032003200620036b41027541016a41026d41027422066a104122013602042002200228020820066a3602080c010b200541186a200620086b2201410175410120011b2201200141036a4102762002280210103f2002280208210620022802042101034020012006470440200528022020012802003602002005200528022041046a360220200141046a21010c010b0b20022902002109200220052903183702002002290208210a20022005290320370208200520093703182005200a3703201040200228020421010b2001417c6a200428020036020020022002280204417c6a3602040c010b0b0b200541186a20071030200528021c410036020041012101200041186a0b2202200228020020016a360200200541306a240020000b3f01027f230041106b2202240020004100103920022001104822001019200128020c102522032001290310102d2000104a2003410110391a200241106a24000b1400200028020004402000103720002802001a0b0b2801017f200028020820002802046b2201410874417f6a410020011b200028021420002802106a6b0ba10202057f017e230041206b22052400024020002802082202200028020c2206470d0020002802042203200028020022044b04402000200320022003200320046b41027541016a417e6d41027422046a103e22023602082000200028020420046a3602040c010b200541086a200620046b2202410175410120021b220220024102762000410c6a103f2103200028020821042000280204210203402002200446450440200328020820022802003602002003200328020841046a360208200241046a21020c010b0b20002902002107200020032902003702002003200737020020002902082107200020032902083702082003200737020820031040200028020821020b200220012802003602002000200028020841046a360208200541206a24000b2501017f200120006b220141027521032001044020022000200110440b200220034102746a0b5f01017f2000410036020c200041106a200336020002402001044020014180808080044f0d012001410274104721040b200020043602002000200420024102746a22023602082000200420014102746a36020c2000200236020420000f0b000b3101027f200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b20002802001a0b1b00200120006b22010440200220016b22022000200110440b20020bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000be10201027f02402001450d00200041003a0000200020016a2202417f6a41003a000020014103490d00200041003a0002200041003a00012002417d6a41003a00002002417e6a41003a000020014107490d00200041003a00032002417c6a41003a000020014109490d002000410020006b41037122036a220241003602002002200120036b417c7122036a2201417c6a410036020020034109490d002002410036020820024100360204200141786a4100360200200141746a410036020020034119490d002002410036021820024100360214200241003602102002410036020c200141706a41003602002001416c6a4100360200200141686a4100360200200141646a41003602002003200241047141187222036b2101200220036a2102034020014120490d0120024200370300200241186a4200370300200241106a4200370300200241086a4200370300200241206a2102200141606a21010c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210421a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041b08b0436020c41a00b200028020c41076a417871220036020041a40b200036020041a80b3f003602000b970101047f230041106b220124002001200036020c2000047f41a80b200041086a2202411076220041a80b2802006a220336020041a40b200241a40b28020022026a41076a417871220436020002400240200341107420044d044041a80b200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104104241086a0541000b200141106a24000b0b002000410120001b10460ba10101037f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b20012802082103024020012802042201410a4d0440200020014101743a0000200041016a21020c010b200141106a4170712204104721022000200136020420002004410172360200200020023602080b2002200320011049200120026a41003a000020000b10002002044020002001200210421a0b0b130020002d0000410171044020002802081a0b0ba10201047f20002001470440200128020420012d00002202410176200241017122031b2102200141016a21052001280208410a2101200520031b210520002d000022034101712204044020002802002203417e71417f6a21010b200220014d0440027f2004044020002802080c010b200041016a0b21012002044020012005200210440b200120026a41003a000020002d00004101710440200020023602040f0b200020024101743a00000f0b027f2003410171044020002802080c010b41000b1a416f2103200141e6ffffff074d0440410b20014101742203200220022003491b220341106a4170712003410b491b21030b200310472204200520021049200020023602042000200436020820002003410172360200200220046a41003a00000b0b880101037f418409410136020041880928020021000340200004400340418c09418c092802002201417f6a220236020020014101484504404184094100360200200020024102746a22004184016a280200200041046a280200110100418409410136020041880928020021000c010b0b418c094120360200418809200028020022003602000c010b0b0b970101027f4184094101360200418809280200220145044041880941900936020041900921010b0240418c092802002202412046044041840210462201450d012001418402104322014188092802003602004188092001360200418c094100360200410021020b418c09200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41840941003602000b3801017f41940b4200370200419c0b410036020041742100034020000440200041a00b6a4100360200200041046a21000c010b0b4104104d0b070041940b104a0b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000105120012802044f0d002002410471044010000c010b200042003702000b02402002411071450d002000105120012802044d0d0020024104710440100020000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001052200010536a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000105441012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bb50102057f017e230041106b22022400200041046a210102402000280200220304402001280200220504402005200041086a2802006a21040b20002004360204200041086a2003360200200241086a20014100200420031056105720002002290308220637020420004100200028020022002006422088a76b2201200120004b1b3602000c010b200020012802002201047f2001200041086a2802006a0541000b360204200041086a41003602000b200241106a24000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a411410501051200241306a24000b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000be70101037f230041106b2204240020004200370200200041086a410036020020012802042103024002402002450440200321020c010b410021022003450d002003210220012802002d000041c001490d00200441086a2001105920004100200428020c2201200428020822022001105622032003417f461b20024520012003497222031b220536020820004100200220031b3602042000200120056b3602000c010b20012802002103200128020421012000410036020020004100200220016b20034520022001497222021b36020820004100200120036a20021b3602040b200441106a240020000b2101017f20011053220220012802044b044010000b2000200120011052200210570bd60202077f017e230041206b220324002001280208220420024b0440200341186a2001105920012003280218200328021c105636020c200341106a20011059410021042001027f410020032802102206450d001a410020032802142208200128020c2207490d001a200820072007417f461b210520060b360210200141146a2005360200200141003602080b200141106a210903400240200420024f0d002001280214450d00200341106a2001105941002104027f410020032802102207450d001a410020032802142208200128020c2206490d001a200820066b2104200620076a0b21052001200436021420012005360210200341106a20094100200520041056105720012003290310220a3702102001200128020c200a422088a76a36020c2001200128020841016a22043602080c010b0b20032009290200220a3703082003200a37030020002003411410501a200341206a24000b980101037f200028020445044041000f0b20001054200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0b800101047f230041106b2201240002402000280204450d0020002802002d000041c001490d00200141086a20001059200128020c210003402000450d01200141002001280208220320032000105622046a20034520002004497222031b3602084100200020046b20031b2100200241016a21020c000b000b200141106a240020020b2d0020002002105e200028020020002802046a2001200210421a2000200028020420026a36020420002003105f0b1b00200028020420016a220120002802084b04402000200110610b0b820201047f02402001450d00034020002802102202200028020c460d01200241786a28020020014904401000200028021021020b200241786a2203200328020020016b220136020020010d012000200336021020004101200028020422042002417c6a28020022016b2202102e220341016a20024138491b220520046a1060200120002802006a220420056a2004200210440240200241374d0440200028020020016a200241406a3a00000c010b200341f7016a220441ff014d0440200028020020016a20043a00002000280200200120036a6a210103402002450d02200120023a0000200241087621022001417f6a21010c000b000b10000b410121010c000b000b0b0f00200020011061200020013602040b2f01017f2000280208200149044020011046200028020020002802041042210220002001360208200020023602000b0b8d0201057f02402001044020002802042104200041106a2802002202200041146a280200220349044020022001ad2004ad422086843702002000200028021041086a3602100f0b027f41002002200028020c22026b410375220541016a2206200320026b2202410275220320032006491b41ffffffff01200241037541ffffffff00491b2202450d001a200241037410470b2103200320054103746a22052001ad2004ad4220868437020020052000280210200028020c22016b22046b2106200441014e044020062001200410421a200028020c21010b2000200320024103746a3602142000200541086a3602102000200636020c2001450d010f0b200041c00110632000410041004101105d0b0b250020004101105e200028020020002802046a20013a00002000200028020441016a3602040b5e01027f2001102e220241b7016a22034180024e044010000b2000200341ff017110632000200028020420026a1060200028020420002802006a417f6a2100034020010440200020013a0000200141087621012000417f6a21000c010b0b0b7701027f2001280200210341012102024002400240024020012802042201410146044020032c000022014100480d012000200141ff017110630c040b200141374b0d01200121020b200020024180017341ff017110630c010b200020011064200121020b2000200320024100105d0b20004101105f20000bbc0202037f037e02402001500440200041800110630c010b20014280015a044020012107034020062007845045044020064238862007420888842107200241016a2102200642088821060c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff017110632000200028020420046a1060200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff017110630b2000200028020420026a1060200028020420002802006a417f6a210203402001200584500d02200220013c0000200542388620014208888421012002417f6a2102200542088821050c000b000b20002001a741ff017110630b20004101105f20000b0b7e01004180080b77696e6974006164645f6d795f6d657373616765006765745f6d795f6d6573736167655f73697a65006765745f6d795f6d6573736167655f68656164006765745f6d795f6d6573736167655f616765006765745f6d795f6d6573736167655f6d6f6e6579006765745f6d795f6d6573736167655f626f6479";

    public static String BINARY = BINARY_0;

    public static final String FUNC_ADD_MY_MESSAGE = "add_my_message";

    public static final String FUNC_GET_MY_MESSAGE_SIZE = "get_my_message_size";

    public static final String FUNC_GET_MY_MESSAGE_HEAD = "get_my_message_head";

    public static final String FUNC_GET_MY_MESSAGE_AGE = "get_my_message_age";

    public static final String FUNC_GET_MY_MESSAGE_MONEY = "get_my_message_money";

    public static final String FUNC_GET_MY_MESSAGE_BODY = "get_my_message_body";

    protected OneInheritWithMultiDataType(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected OneInheritWithMultiDataType(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<OneInheritWithMultiDataType> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(OneInheritWithMultiDataType.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<OneInheritWithMultiDataType> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(OneInheritWithMultiDataType.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<OneInheritWithMultiDataType> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(OneInheritWithMultiDataType.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<OneInheritWithMultiDataType> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(OneInheritWithMultiDataType.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<TransactionReceipt> add_my_message(My_message one_message) {
        final WasmFunction function = new WasmFunction(FUNC_ADD_MY_MESSAGE, Arrays.asList(one_message), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> add_my_message(My_message one_message, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_ADD_MY_MESSAGE, Arrays.asList(one_message), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public RemoteCall<Uint8> get_my_message_size() {
        final WasmFunction function = new WasmFunction(FUNC_GET_MY_MESSAGE_SIZE, Arrays.asList(), Uint8.class);
        return executeRemoteCall(function, Uint8.class);
    }

    public RemoteCall<String> get_my_message_head(Uint8 index) {
        final WasmFunction function = new WasmFunction(FUNC_GET_MY_MESSAGE_HEAD, Arrays.asList(index), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<Uint32> get_my_message_age(Uint8 index) {
        final WasmFunction function = new WasmFunction(FUNC_GET_MY_MESSAGE_AGE, Arrays.asList(index), Uint32.class);
        return executeRemoteCall(function, Uint32.class);
    }

    public RemoteCall<Uint64> get_my_message_money(Uint8 index) {
        final WasmFunction function = new WasmFunction(FUNC_GET_MY_MESSAGE_MONEY, Arrays.asList(index), Uint64.class);
        return executeRemoteCall(function, Uint64.class);
    }

    public RemoteCall<String> get_my_message_body(Uint8 index) {
        final WasmFunction function = new WasmFunction(FUNC_GET_MY_MESSAGE_BODY, Arrays.asList(index), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static OneInheritWithMultiDataType load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new OneInheritWithMultiDataType(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static OneInheritWithMultiDataType load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new OneInheritWithMultiDataType(contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static class Message {
        public String head;

        public Uint32 age;

        public Uint64 money;
    }

    public static class My_message {
        public Message baseClass;

        public String body;

        public String end;
    }
}
