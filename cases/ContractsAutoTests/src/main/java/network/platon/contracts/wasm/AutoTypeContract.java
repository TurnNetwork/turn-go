package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Int32;
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
 * <p>Generated with platon-web3j version 0.13.0.6.
 */
public class AutoTypeContract extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001480d60017f017f60027f7f0060017f0060037f7f7f017f60000060037f7f7f0060027f7f017f60047f7f7f7f017f60047f7f7f7f0060027f7e006000017f60027f7e017f60017f017e02a9010703656e760c706c61746f6e5f70616e6963000403656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000a03656e7610706c61746f6e5f6765745f696e707574000203656e7617706c61746f6e5f6765745f73746174655f6c656e677468000603656e7610706c61746f6e5f6765745f7374617465000703656e7610706c61746f6e5f7365745f7374617465000803656e760d706c61746f6e5f72657475726e0001034d4c0400020202000000000006010506070102000004000301060702010c020202040000020009020109090b0303050102010003010101010301080101020001030702030505000404000008010104050170010a0a05030100020608017f0141b08b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070f5f5f66756e63735f6f6e5f65786974002606696e766f6b65001a090f010041010b0909090c0d0e0f0910090abf5f4c100041f40810081a4101100a104c104d0b190020004200370200200041086a41003602002000100b20000b0300010b940101027f41800941013602004184092802002201450440418409418c09360200418c0921010b024041880928020022024120460440418402101b2201450d012001104b220141840928020036020041840920013602004188094100360200410021020b418809200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41800941003602000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0b040041050b040041760b0400411e0b0400410a0b4f01027f230041106b22012400200041186a2202200141800810111012200220014183081011101220022001418608101110122000411c6a2802002000280218200141106a24006b410c6e41ff01710b910101027f20004200370200200041086a410036020020012102024003402002410371044020022d0000450d02200241016a21020c010b0b2002417c6a21020340200241046a22022802002203417f73200341fffdfb776a7141808182847871450d000b0340200341ff0171450d01200241016a2d00002103200241016a21020c000b000b20002001200220016b101320000bb00101037f230041206b22032400024020002802042202200028020849044020022001290200370200200241086a200141086a2802003602002001100b20002000280204410c6a3602040c010b200341086a2000200220002802006b410c6d220241016a10142002200041086a1015220228020822042001290200370200200441086a200141086a2802003602002001100b20022002280208410c6a360208200020021016200210170b200341206a24000b5a01027f02402002410a4d0440200020024101743a0000200041016a21030c010b200241106a4170712204101921032000200236020420002004410172360200200020033602080b200320012002104a200220036a41003a00000b3101017f2001200028020820002802006b410c6d2200410174220220022001491b41d5aad5aa01200041aad5aad500491b0b4c01017f2000410036020c200041106a2003360200200104402001101821040b20002004360200200020042002410c6c6a2202360208200020042001410c6c6a36020c2000200236020420000baa0101037f200028020421022000280200210303402002200346450440200128020441746a2204200241746a2202290200370200200441086a200241086a2802003602002002100b2001200128020441746a3602040c010b0b200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b2b01027f20002802082101200028020421020340200120024704402000200141746a22013602080c010b0b0b09002000410c6c10190b0b002000410120001b101b0bba0402067f017e230041406a22032400100710012200101b220210020240200341206a20022000101c2201280208450440200141146a2802002104200128021021000c010b200341386a2001101d20012003280238200328023c101e36020c200341086a2001101d2001027f410020032802082202450d001a4100200328020c2205200128020c2200490d001a200520002000417f461b210420020b2200360210200141146a2004360200200141003602080b200341086a200020044114101f22021020024002402002280204450d00200210200240200228020022002c0000220141004e044020010d010c020b200141807f460d00200141ff0171220441b7014d0440200228020441014d04401000200228020021000b20002d00010d010c020b200441bf014b0d012002280204200141ff017141ca7e6a22014d04401000200228020021000b200020016a2d0000450d010b2002280204450d0020002d000041c001490d010b10000b200341386a20021021200328023c220041094f044010000b200328023821040340200004402000417f6a210020043100002006420886842106200441016a21040c010b0b024002402006500d0041890810222006510440410210230c020b418e0810222006510440410310240c020b419b0810222006510440410410240c020b41aa0810222006510440410510240c020b41bc0810222006510440410610250c020b41cd0810222006510440410710230c020b41df0810222006520d00410810250c010b10000b1026200341406b24000b970101047f230041106b220124002001200036020c2000047f41a40b200041086a2202411076220041a40b2802006a220336020041a00b200241a00b28020022026a41076a417871220436020002400240200341107420044d044041a40b200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104103d41086a0541000b200141106a24000b0c00200020012002411c101f0b2101017f20011027220220012802044b044010000b200020012001104f200210500b2301017f230041206b22022400200241086a200020014114101f104e200241206a24000b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000104e20024f0d002003410471044010000c010b200042003702000b02402003411071450d002000104e20024d0d0020034104710440100020000f0b200042003702000b20000b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bd40101047f200110272204200128020422024b04401000200128020421020b200128020021052000027f02400240200204404100210120052c00002203417f4a0d01027f200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a21010c010b4101210120050d000c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b2701027f230041306b22012400200141086a1028200141086a20001102001029200141306a24000bbc0102037f017e230041e0006b2201240020011028200120001100002100200141286a102a2102200141d8006a4100360200200141d0006a4200370300200141c8006a420037030020014200370340200141406b2000ac22044201862004423f87852204102b20012802402100200141406b410472102c20022000102d20022004102e200228020c200241106a28020047044010000b200228020020022802041006200228020c22000440200220003602100b1029200141e0006a24000bb10102037f017e230041e0006b2201240020011028200120001100002100200141286a102a2102200141d8006a4100360200200141d0006a4200370300200141c8006a420037030020014200370340200141406b2000ad2204102f20012802402100200141406b410472102c20022000102d200220041030220228020c200241106a28020047044010000b200228020020022802041006200228020c22000440200220003602100b1029200141e0006a24000b880101037f4180094101360200418409280200210003402000044003404188094188092802002201417f6a220236020020014101484504404180094100360200200020024102746a22004184016a280200200041046a280200110200418009410136020041840928020021000c010b0b4188094120360200418409200028020022003602000c010b0b0bff0201037f200028020445044041000f0b2000102041012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020bb409010d7f23004190016b22032400200042003702182000428debc585c3a7f9dbf7003703102000410036020820004200370200200041206a4100360200200341186a102a22082000290310102e200828020c200841106a28020047044010000b200041186a210a0240200828020022062008280204220410032207450d002007101921020340200120026a41003a00002007200141016a2201470d000b20062004200220011004417f460440410021010c010b024002402003200241016a200120026a2002417f736a101c2202280204450d0020022802002d000041c001490d00200341406b2002101d2003280244210141002106034020010440200341002003280240220420042001101e22056a20044520012005497222041b3602404100200120056b20041b2101200641016a21060c010b0b2000280220200028021822016b410c6d2006490440200a200341406b2006200028021c20016b410c6d200041206a101522011016200110170b200341e8006a2002410110312105200341d8006a200241001031210c20052802042101200341f8006a41017221060340200c280204200146410020052802082202200c280208461b0d02200341406b20012002411c101f2101200341306a1008210b024002402003280244044020032802402d000041c001490d010b200341f8006a10081a0c010b20034188016a20011021200110272102024002400240024020032802880122010440200328028c01220420024f0d010b4100210120034180016a41003602002003420037037841002104410021090c010b20034180016a4100360200200342003703782001200420022002417f461b22046a21092004410a4b0d010b200320044101743a0078200621020c010b200441106a417071220d101921022003200436027c2003200d41017236027820032002360280010b03402001200946450440200220012d00003a0000200241016a2102200141016a21010c010b0b200241003a00000b024020032d0030410171450440200341003b01300c010b200328023841003a00002003410036023420032d0030410171450d00200341003602300b200341386a20034180016a28020036020020032003290378370330200341f8006a100b200a200b101220052005280204220120052802086a410020011b22013602042005280200220204402005200236020820012002101e21092005027f200528020422044504404100210241000c010b4100210241002005280208220b2009490d001a200b20092009417f461b210220040b2201ad2002ad42208684370204200541002005280200220420026b2202200220044b1b36020005200541003602080b0c000b000b10000b200721010b200828020c22020440200820023602100b024020010d002000411c6a210720002802042206200028020022016b410c6d22042000280220200028021822026b410c6d4d04402004200728020020026b410c6d22044b0440200120012004410c6c6a2201200210321a20012006200710330c020b200a200120062002103210340c010b20020440200a103520004100360220200042003702180b2000200a200410142204101822023602182000200236021c200020022004410c6c6a36022020012006200710330b20034190016a240020000bf908010b7f230041e0006b22032400200341186a102a2107200341c8006a22044100360200200341406b22024200370300200341386a2205420037030020034200370330200341306a2000290310102b20032802302101200341306a410472102c20072001102d20072000290310102e200728020c200741106a28020047044010000b2007280204210a20072802002003102a210120044100360200200242003703002005420037030020034200370330027f20002802182000411c6a2802004604402003410136023041010c010b200341306a4100103b200341d0006a4101722106200028021c210820002802182104037f2004200846047f200341306a4101103b200328023005200341d0006a2004103c41012105024002400240200328025420032d00502202410176200241017122091b220241014d0440200241016b0d032003280258200620091b2c0000417f4c0d010c030b200241374b0d010b200241016a21050c010b2002103720026a41016a21050b027f200341306a20032802482202450d001a2003280238200220032802446a417f6a220241087641fcffff07716a280200200241ff07714102746a0b2202200228020020056a3602002004410c6a21040c010b0b0b2102200341306a410472102c41011019220441fe013a0000200128020c200141106a28020047044010000b2001280204220541016a220620012802084b047f20012006103620012802040520050b20012802006a20044101103d1a2001200128020441016a3602042001200441016a200220046b6a102d0240200028021c20002802186b220404402004410c6d21042001280204210220012802102205200141146a280200220649044020052004ad2002ad422086843702002001200128021041086a3602100c020b027f41002005200128020c22056b410375220841016a2209200620056b2205410275220620062009491b41ffffffff01200541037541ffffffff00491b2205450d001a200541037410190b2106200620084103746a22082004ad2002ad4220868437020020082001280210200128020c22096b22046b2102200441014e0440200220092004103d1a0b2001200620054103746a3602142001200841086a3602102001200236020c0c010b200141c001103e2001410041004101103f0b200341306a4101722106200028021c210820002802182104034020042008470440200341306a2004103c2003280238200620032d0030220241017122091b210502400240024002402003280234200241017620091b2202410146044020052c000022024100480440410121020c020b2001200241ff0171103e0c040b200241374b0d010b200120024180017341ff0171103e0c010b2001200210400b2001200520024100103f0b2001410110412004410c6a21040c010b0b0240200128020c2001280210460440200128020021040c010b100020012802002104200128020c2001280210460d0010000b200a200420012802041005200128020c22040440200120043602100b200728020c22010440200720013602100b200041186a104220001042200341e0006a24000b29002000410036020820004200370200200041001036200041146a41003602002000420037020c20000b080020002001102f0bea0101047f230041106b22042400200028020422012000280210220341087641fcffff07716a2102027f410020012000280208460d001a2002280200200341ff07714102746a0b2101200441086a20001039200428020c210303400240200120034604402000410036021420002802082102200028020421010340200220016b41027522034103490d022000200141046a22013602040c000b000b200141046a220120022802006b418020470d0120022802042101200241046a21020c010b0b2003417f6a220241014d04402000418004418008200241016b1b3602100b20002001103a200441106a24000b1300200028020820014904402000200110360b0b09002000200110301a0b840102027f017e4101210320014280015a0440034020012004845045044020044238862001420888842101200241016a2102200442088821040c010b0b200241384f047f2002103720026a0520020b41016a21030b200041186a28020022020440200041086a280200200041146a2802002002103821000b2000200028020020036a3602000bbc0202037f037e024020015004402000418001103e0c010b20014280015a044020012107034020062007845045044020064238862007420888842107200241016a2102200642088821060c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff0171103e2000200028020420046a1052200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff0171103e0b2000200028020420026a1052200028020420002802006a417f6a210203402001200584500d02200220013c0000200542388620014208888421012002417f6a2102200542088821050c000b000b20002001a741ff0171103e0b20004101104120000be70101037f230041106b2204240020004200370200200041086a410036020020012802042103024002402002450440200321020c010b410021022003450d002003210220012802002d000041c001490d00200441086a2001101d20004100200428020c2201200428020822022001101e22032003417f461b20024520012003497222031b220536020820004100200220031b3602042000200120056b3602000c010b20012802002103200128020421012000410036020020004100200220016b20034520022001497222021b36020820004100200120036a20021b3602040b200441106a240020000bc80201067f200120006b2108410021010340200120026a2105200120084645044002402005200020016a2206460d00200641046a28020020062d00002204410176200441017122071b2104410a2103200641086a280200200641016a20071b210720052d0000410171220604402005280200417e71417f6a21030b200420034d0440027f20060440200541086a2802000c010b200541016a0b21032004044020032007200410490b200320046a41003a000020052d00004101710440200541046a20043602000c020b200520044101743a00000c010b416f2106200341e6ffffff074d0440410b20034101742203200420042003491b220341106a4170712003410b491b21060b20061019220320072004104a200541046a200436020020052006410172360200200541086a2003360200200320046a41003a00000b2001410c6a21010c010b0b20050b2d000340200020014645044020022802002000103c20022002280200410c6a3602002000410c6a21000c010b0b0b0900200020013602040b0b002000200028020010340b2f01017f200028020820014904402001101b20002802002000280204103d210220002001360208200020023602000b0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b25002000200120026a417f6a220141087641fcffff07716a280200200141ff07714102746a0b4f01037f20012802042203200128021020012802146a220441087641fcffff07716a21022000027f410020032001280208460d001a2002280200200441ff07714102746a0b360204200020023602000b2501017f200028020821020340200120024645044020002002417c6a22023602080c010b0b0bbd0c02077f027e230041306b22052400200041046a2107024020014101460440200041086a280200200041146a280200200041186a220228020022041038280200210120022004417f6a360200200710434180104f044020072000410c6a280200417c6a103a0b200141384f047f2001103720016a0520010b41016a2101200041186a2802002202450d01200041086a280200200041146a2802002002103821000c010b0240200710430d00200041146a28020022014180084f0440200020014180786a360214200041086a2201280200220228020021042001200241046a360200200520043602182007200541186a10440c010b2000410c6a2802002202200041086a2802006b4102752204200041106a2203280200220620002802046b220141027549044041802010192104200220064704400240200028020c220120002802102206470d0020002802082202200028020422034b04402000200220012002200220036b41027541016a417e6d41027422036a1045220136020c2000200028020820036a3602080c010b200541186a200620036b2201410175410120011b22012001410276200041106a10462102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021047200028020c21010b200120043602002000200028020c41046a36020c0c020b02402000280208220120002802042206470d00200028020c2202200028021022034904402000200120022002200320026b41027541016a41026d41027422036a104822013602082000200028020c20036a36020c0c010b200541186a200320066b2201410175410120011b2201200141036a410276200041106a10462102200028020c210320002802082101034020012003470440200228020820012802003602002002200228020841046a360208200141046a21010c010b0b200029020421092000200229020037020420022009370200200029020c21092000200229020837020c2002200937020820021047200028020821010b2001417c6a2004360200200020002802082201417c6a22023602082002280200210220002001360208200520023602182007200541186a10440c010b20052001410175410120011b200420031046210241802010192106024020022802082201200228020c2208470d0020022802042204200228020022034b04402002200420012004200420036b41027541016a417e6d41027422036a104522013602082002200228020420036a3602040c010b200541186a200820036b2201410175410120011b22012001410276200241106a280200104621042002280208210320022802042101034020012003470440200428020820012802003602002004200428020841046a360208200141046a21010c010b0b20022902002109200220042902003702002004200937020020022902082109200220042902083702082004200937020820041047200228020821010b200120063602002002200228020841046a360208200028020c2104034020002802082004460440200028020421012000200228020036020420022001360200200228020421012002200436020420002001360208200029020c21092000200229020837020c2002200937020820021047052004417c6a210402402002280204220120022802002208470d0020022802082203200228020c22064904402002200120032003200620036b41027541016a41026d41027422066a104822013602042002200228020820066a3602080c010b200541186a200620086b2201410175410120011b2201200141036a410276200228021010462002280208210620022802042101034020012006470440200528022020012802003602002005200528022041046a360220200141046a21010c010b0b20022902002109200220052903183702002002290208210a20022005290320370208200520093703182005200a3703201047200228020421010b2001417c6a200428020036020020022002280204417c6a3602040c010b0b0b200541186a20071039200528021c4100360200200041186a2100410121010b2000200028020020016a360200200541306a24000b4901017f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a2802003602000f0b20002001280208200128020410130bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000b2500200041011051200028020020002802046a20013a00002000200028020441016a3602040b2d00200020021051200028020020002802046a20012002103d1a2000200028020420026a3602042000200310410b5e01027f20011037220241b7016a22034180024e044010000b2000200341ff0171103e2000200028020420026a1052200028020420002802006a417f6a2100034020010440200020013a0000200141087621012000417f6a21000c010b0b0b820201047f02402001450d00034020002802102202200028020c460d01200241786a28020020014904401000200028021021020b200241786a2203200328020020016b220136020020010d012000200336021020004101200028020422042002417c6a28020022016b22021037220341016a20024138491b220520046a1052200120002802006a220420056a2004200210490240200241374d0440200028020020016a200241406a3a00000c010b200341f7016a220441ff014d0440200028020020016a20043a00002000280200200120036a6a210103402002450d02200120023a0000200241087621022001417f6a21010c000b000b10000b410121010c000b000b0b0e0020002802000440200010350b0b2801017f200028020820002802046b2201410874417f6a410020011b200028021420002802106a6b0ba10202057f017e230041206b22052400024020002802082202200028020c2206470d0020002802042203200028020022044b04402000200320022003200320046b41027541016a417e6d41027422046a104522023602082000200028020420046a3602040c010b200541086a200620046b2202410175410120021b220220024102762000410c6a10462103200028020821042000280204210203402002200446450440200328020820022802003602002003200328020841046a360208200241046a21020c010b0b20002902002107200020032902003702002003200737020020002902082107200020032902083702082003200737020820031047200028020821020b200220012802003602002000200028020841046a360208200541206a24000b2501017f200120006b220141027521032001044020022000200110490b200220034102746a0b4f01017f2000410036020c200041106a2003360200200104402001410274101921040b200020043602002000200420024102746a22023602082000200420014102746a36020c2000200236020420000b2b01027f200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b0b1b00200120006b22010440200220016b22022000200110490b20020b8d0301037f024020002001460d00200120006b20026b410020024101746b4d0440200020012002103d1a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b100020020440200020012002103d1a0b0bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b3501017f230041106b220041b08b0436020c419c0b200028020c41076a417871220036020041a00b200036020041a40b3f003602000b3801017f41900b420037020041980b4100360200417421000340200004402000419c0b6a4100360200200041046a21000c010b0b4109100a0b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f2000104f200010276a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b1b00200028020420016a220120002802084b04402000200110360b0b0f00200020011036200020013602040b0b7701004180080b70763100763200763300696e6974006765745f616e746f5f696e74006765745f616e746f5f696e743332006765745f616e746f5f6d756c7469706c65006765745f616e746f5f75696e74385f74007365745f616e746f5f636172655f6f6e65006765745f616e746f5f6974657261746f72";

    public static String BINARY = BINARY_0;

    public static final String FUNC_GET_ANTO_INT = "get_anto_int";

    public static final String FUNC_GET_ANTO_ITERATOR = "get_anto_iterator";

    public static final String FUNC_GET_ANTO_INT32 = "get_anto_int32";

    public static final String FUNC_GET_ANTO_MULTIPLE = "get_anto_multiple";

    public static final String FUNC_GET_ANTO_UINT8_T = "get_anto_uint8_t";

    public static final String FUNC_SET_ANTO_CARE_ONE = "set_anto_care_one";

    protected AutoTypeContract(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    protected AutoTypeContract(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }

    public RemoteCall<Int32> get_anto_int() {
        final WasmFunction function = new WasmFunction(FUNC_GET_ANTO_INT, Arrays.asList(), Int32.class);
        return executeRemoteCall(function, Int32.class);
    }

    public RemoteCall<Uint8> get_anto_iterator() {
        final WasmFunction function = new WasmFunction(FUNC_GET_ANTO_ITERATOR, Arrays.asList(), Uint8.class);
        return executeRemoteCall(function, Uint8.class);
    }

    public static RemoteCall<AutoTypeContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(AutoTypeContract.class, web3j, credentials, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<AutoTypeContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(AutoTypeContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<AutoTypeContract> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(AutoTypeContract.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public static RemoteCall<AutoTypeContract> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(AutoTypeContract.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public RemoteCall<Int32> get_anto_int32() {
        final WasmFunction function = new WasmFunction(FUNC_GET_ANTO_INT32, Arrays.asList(), Int32.class);
        return executeRemoteCall(function, Int32.class);
    }

    public RemoteCall<Int32> get_anto_multiple() {
        final WasmFunction function = new WasmFunction(FUNC_GET_ANTO_MULTIPLE, Arrays.asList(), Int32.class);
        return executeRemoteCall(function, Int32.class);
    }

    public RemoteCall<Uint8> get_anto_uint8_t() {
        final WasmFunction function = new WasmFunction(FUNC_GET_ANTO_UINT8_T, Arrays.asList(), Uint8.class);
        return executeRemoteCall(function, Uint8.class);
    }

    public RemoteCall<TransactionReceipt> set_anto_care_one() {
        final WasmFunction function = new WasmFunction(FUNC_SET_ANTO_CARE_ONE, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> set_anto_care_one(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SET_ANTO_CARE_ONE, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static AutoTypeContract load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        return new AutoTypeContract(contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    public static AutoTypeContract load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        return new AutoTypeContract(contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }
}
