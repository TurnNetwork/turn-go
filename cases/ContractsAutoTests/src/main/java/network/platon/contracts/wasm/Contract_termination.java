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
public class Contract_termination extends WasmContract {
    private static String BINARY_0 = "0x0061736d01000000017f1160047f7f7f7f0060017f0060027f7f0060017f017f60000060037f7f7f0060027f7f017f60037f7f7f017f60027f7e0060047f7f7f7f017f60017f017e60087f7f7f7f7f7f7f7f006000017f600a7f7f7f7f7f7f7f7f7f7f017f600b7f7f7f7f7f7f7f7f7f7f7f017f600a7f7f7f7f7e7f7e7f7f7f017f60037f7f7e017f02d0010903656e760c706c61746f6e5f6465627567000203656e760c706c61746f6e5f70616e6963000403656e760d706c61746f6e5f726576657274000403656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000c03656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e000203656e7617706c61746f6e5f6765745f73746174655f6c656e677468000603656e7610706c61746f6e5f6765745f7374617465000903656e7610706c61746f6e5f7365745f737461746500000349480403010110070a040a03010203080102010302010202010803030301090000030d0f0e0707060504030306050102050b0201010204010107030303010006020503010202020202080405017001050505030100020608017f0141908d040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300090f5f5f66756e63735f6f6e5f65786974003d06696e766f6b650010090a010041010b040c26273f0aac6b48180041d40a100a1a4101103e103041f00c100a1a4104103e0b190020004200370200200041086a41003602002000100b20000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0b070041d40a10350be60101037f230041206b22032400200242e30058044041d40a41f20941f2091023103741ce09103a41f809103a418008103a41f809103a41e009103a41f809103a419c09103a41f809103a41e609103a41f809103a200341106a103b027f20032d0010220441017104402003280218210520032802140c010b200341106a410172210520044101760b2104200520041039200341106a103541f809103a41ec09103a41f809103a418d08103a41f809103a41ac09103a200341dc0a28020041d50a41d40a2d00004101711b3602002003102410020b200041186a20011036200341206a240041000b3401017f230041106b220324002003200236020c200320013602082003200329030837030020002003411c1040200341106a24000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010be10302067f017e230041b0016b2200240010091003220110312202100420004188016a200020022001100e2201410010480240024020004188016a10112206500d0041b609100f200651044020004188016a101210130c020b419c09100f2006510440200041186a100a21022000420037032820004188016a20014101104820004188016a2002101420004188016a200141021048200020004188016a101137032820004188016a1012200041306a2002103321032000290328210620004188016a200041406b2003103322012006100d210520011035200041d0006a1015210120004180016a4100360200200041f8006a4200370300200041f0006a420037030020004200370368200041e8006a2005ad2206101620002802682105200041e8006a4104721017200120051018200120061050200128020c200141106a28020047044010010b20012802002001280204100520011019200310351013200210350c020b41bb09100f2006520d0020004188016a1012200041186a200041a0016a10332102200041e8006a101522012002101a10182001200041d0006a200210332204101b20041035200128020c200141106a28020047044010010b200128020020012802041005200110192002103510130c010b10010b103d200041b0016a24000b850102027f017e230041106b22012400200010440240024020001049450d002000280204450d0020002802002d000041c001490d010b10010b200141086a2000101d200128020c220041094f044010010b200128020821020340200004402000417f6a210020023100002003420886842103200241016a21020c010b0b200141106a240020030bde02010c7f230041406a220124002000100a210a200042afb59bdd9e8485b9f800370310200041186a100a2108200141286a1015220420002903101020200428020c200441106a28020047044010010b024002402004280200220b2004280204220c10062203450440410021030c010b20014100360220200142003703182003417f4c0d012003103221050340200220056a41003a00002003200241016a2202470d000b200220056a21072005200128021c200128021822066b22096b2102200941014e0440200220062009102d1a200128021821060b2001200320056a3602202001200736021c200120023602180240200b200c20060440200128021c2107200128021821020b2002200720026b1007417f460440410021030c010b20012001280218220241016a200128021c2002417f736a100e200810140b200141186a101c0b2004101920034504402008200a10360b200141406b240020000f0b000bb103010c7f230041e0006b22012400200141286a10152104200141d8006a4100360200200141d0006a4200370300200141c8006a420037030020014200370340200141406b2000290310101620012802402103200141406b4104721017200420031018200420002903101020200428020c200441106a28020047044010010b200428020421092004280200200141406b10152102200041186a2207101a210b41011032220341fe013a0000200120033602182001200341016a22053602202001200536021c200228020c200241106a2802004704401001200128021c2105200128021821030b2003210620022802042208200520036b22056a220c20022802084b04402002200c101e20022802042108200128021821060b200228020020086a20032005102d1a2002200228020420056a3602042002200128021c200b20066b6a10182002200141086a200710332203101b200310350240200228020c2002280210460440200228020021030c010b100120022802002103200228020c2002280210460d0010010b2009200320022802041008200141186a101c20021019200410192007103520001035200141e0006a24000b8c0301057f230041206b220224000240024002402000280204044020002802002d000041c001490d010b200241086a100a1a0c010b200241186a2000101d2000104321030240024002400240200228021822000440200228021c220420034f0d010b41002100200241106a410036020020024200370308410021030c010b200241106a410036020020024200370308200420032003417f461b220341704f0d04200020036a21052003410a4b0d010b200220034101743a0008200241086a41017221040c010b200341106a4170712206103221042002200336020c20022006410172360208200220043602100b034020002005470440200420002d00003a0000200441016a2104200041016a21000c010b0b200441003a00000b024020012d0000410171450440200141003b01000c010b200128020841003a00002001410036020420012d0000410171450d0020012802081a200141003602000b20012002290308370200200141086a200241106a280200360200200241086a100b200241086a1035200241206a24000f0b000b2900200041003602082000420037020020004100101e200041146a41003602002000420037020c20000b7502027f017e4101210320014280015a0440034020012004845045044020044238862001420888842101200241016a2102200442088821040c010b0b200241384f047f2002102120026a0520020b41016a21030b200041186a2802000440200041046a102221000b2000200028020020036a3602000bc40201067f200028020422012000280210220341087641fcffff07716a2102027f200120002802082205460440200041146a210441000c010b2001200028021420036a220441087641fcffff07716a280200200441ff07714102746a2106200041146a21042002280200200341ff07714102746a0b21030340024020032006460440200441003602000340200520016b41027522024103490d0220012802001a2000200028020441046a2201360204200028020821050c000b000b200341046a220320022802006b418020470d0120022802042103200241046a21020c010b0b2002417f6a220241014d04402000418004418008200241016b1b3602100b03402001200547044020012802001a200141046a21010c010b0b200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b20002802001a0b13002000280208200149044020002001101e0b0b1c01017f200028020c22010440200041106a20013602000b2000101f0bbc0101047f230041306b22012400200141286a4100360200200141206a4200370300200141186a42003703002001420037031041012102024002400240200120001033220328020420032d00002200410176200041017122041b220041014d0440200041016b0d032003280208200341016a20041b2c0000417f4c0d010c030b200041374b0d010b200041016a21020c010b2000102120006a41016a21020b2001200236021020031035200141106a4104721017200141306a240020020b5201037f230041106b2202240020022001280208200141016a20012d0000220341017122041b36020820022001280204200341017620041b36020c2002200229030837030020002002104f200241106a24000b1501017f200028020022010440200020013602040b0bd60101047f200110432204200128020422024b04401001200128020421020b20012802002105027f027f41002002450d001a410020052c00002203417f4a0d011a200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a0b21012000027f02402005450440410021030c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b3401017f2000280208200149044020011031220220002802002000280204102d1a2000101f20002001360208200020023602000b0b080020002802001a0b08002000200110500b1e01017f03402000044020004108762100200141016a21010c010b0b20010b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0b7801027f20002101024003402001410371044020012d0000450d02200141016a21010c010b0b2001417c6a21010340200141046a22012802002202417f73200241fffdfb776a7141808182847871450d000b0340200241ff0171450d01200141016a2d00002102200141016a21010c000b000b200120006b0b3b01017f23004190086b220124002001200036020c200141106a41800841f309200010251a200141106a200141106a1023100020014190086a24000be316030f7f027e037c230041306b220824002008200236020c4102410320001b210b2008410f6a21100340410020076b210a0240034020022d00002206450d01200641254704402006411874411875200020072001200b1100002008200241016a220236020c200a417f6a210a200741016a21070c010b0b2008200241016a220236020c410021040340024002400240024020022c0000220641556a220541054b0440200641606a220541034b0d0102400240200541016b0e03030301000b2008200241016a220236020c200441087221040c060b2008200241016a220236020c200441107221040c050b200541016b0e050002000003010b0240200641506a41ff017141094d04402008410c6a1028210d200828020c21020c010b4100210d2006412a470d00200328020021062008200241016a220236020c2004410272200420064100481b210420062006411f7522056a200573210d200341046a21030b41002109024020022d0000412e470d002008200241016a220636020c200441800872210420022d0001220541506a41ff017141094d04402008410c6a10282109200828020c21020c010b2005412a470440200621020c010b200328020021062008200241026a220236020c20064100200641004a1b2109200341046a21030b0240024020022c000041987f6a411f77220641094b0d000240024002400240200641016b0e09020004040403040403010b2008200241016a220636020c20022d0001220541ec004704402004418002722104200621020c050b2008200241026a220236020c20044180067221040c030b2008200241016a220636020c20022d0001220541e8004704402004418001722104200621020c040b2008200241026a220236020c200441c0017221040c020b2008200241016a220236020c20044180047221040c010b2008200241016a220236020c20044180027221040b20022d000021050b02400240024002400240024002400240024020054118744118752206419e7f6a220c41164b044020064125470440200641c600460d07200641d800470d020c080b4125200020072001200b1100000c020b200c41016b0e15040600050000060000000000060200000300060000060b2006200020072001200b1100000b2008200241016a220236020c200741016a21070c0c0b200b200020072001200328020041004110200941082004412172102921072008200241016a220236020c200341046a21030c0b0b20032802002205417f6a210a0340200a41016a220a2d00000d000b200a20056b2206200920062009491b2006200441800871220c410a761b210602402004410271220f0440200721040c010b4100210a03402007200a6a21042006200a6a220e200d4f4504404120200020042001200b110000200a41016a210a0c010b0b200e41016a21060b200341046a21030340024020052d00002207450d00200c04402009450d012009417f6a21090b2007411874411875200020042001200b110000200441016a2104200541016a21050c010b0b200f450440200421070c050b4100210a03402004200a6a21072006200a6a200d4f0d054120200020072001200b110000200a41016a210a0c000b000b4101210a0240200441027122060440200721050c010b410121040340200420076a417f6a21052004200d4f450440200441016a21044120200020052001200b1100000c010b0b200441016a210a0b20032c0000200020052001200b110000200541016a21072006450d020340200a200d4f0d034120200020072001200b110000200741016a2107200a41016a210a0c000b000b200941062004418008711b220941037441800a6a2106200341076a41787122112b030021154100210503402009410a492005411f4b72450440200841106a20056a41303a0000200641786a21062009417f6a2109200541016a21050c010b0b027f4400000000000000002015a12015201544000000000000000063220e1b22159944000000000000e0416304402015aa0c010b4180808080780b2103027f20152003b7a120062b03002217a2221644000000000000f041632016440000000000000000667104402016ab0c010b41000b210c02402016200cb8a1221644000000000000e03f644101734504402017200c41016a220cb8654101730d01200341016a21034100210c0c010b201644000000000000e03f620d00200c45200c41017172200c6a210c0b4100210602402015440000c0ffffffdf41640d00024020090440200520096a41606a2106034002402005412046044041202105200621090c010b200841106a20056a200c200c410a6e220f41766c6a4130723a0000200541016a21052009417f6a2109200c41094b200f210c0d010b0b03402005411f4b220620094572450440200841106a20056a41303a0000200541016a21052009417f6a21090c010b0b20060d01200841106a20056a412e3a0000200541016a21050c010b20152003b7a1221544000000000000e03f64410173450440200341016a21030c010b20032003201544000000000000e03f61716a21030b03402005411f4d0440200841106a20056a20032003410a6d220641766c6a41306a3a0000200541016a2105200341096a2006210341124b0d010b0b20044103712103034020034101472005411f4b722005200d4f72450440200841106a20056a41303a0000200541016a21050c010b0b20044101712106200441027121030240200d2004410c71410047200e726b20052005200d461b2205411f4b0d000240200e410173450440200841106a20056a412d3a00000c010b20044104710440200841106a20056a412b3a00000c010b2004410871450d01200841106a20056a41203a00000b200541016a21050b024020032006720440200721060c010b410021040340200420076a2106200420056a200d4f0d014120200020062001200b110000200441016a21040c000b000b2006200a6a2107034020050440200520106a2c0000200020062001200b110000200741016a21072005417f6a2105200641016a21060c010b0b2003450d0003402007200d4f0d014120200020062001200b110000200741016a2107200641016a21060c000b000b201141086a21032008200241016a220236020c200621070c080b41102106027f0240200541ff0171220541d80046220c200541f8004672450440200541ef00460440410821060c020b200541e200460440410221060c020b2004416f712104410a21060b20044120722004200c1b2204200541e40046200541e90046720d011a0b20044173710b2204417e7120042004418008711b2104200541e900474100200541e400471b4504402004418004710440200b200020072001200341076a417871220329030022132013423f8722147c2014852013423f88a72006ad2009200d2004102a2107200341086a21030c030b2004418002710440200b200020072001200328020022072007411f7522056a2005732007411f7620062009200d2004102921070c020b200b200020072001027f200441c00071044020032c00000c010b2003280200220541107441107520052004418001711b0b220a411f752207200a6a200773200a411f7620062009200d2004102921070c010b2004418004710440200b200020072001200341076a417871220329030041002006ad2009200d2004102a2107200341086a21030c020b2004418002710440200b2000200720012003280200410020062009200d2004102921070c010b200b200020072001027f200441c00071044020032d00000c010b2003280200220541ffff037120052004418001711b0b410020062009200d2004102921070b200341046a21030b2008200241016a220236020c0c050b2008200241016a220236020c200441047221040c020b2008200241016a220236020c200441027221040c010b2008200241016a220236020c200441017221040c000b000b0b4100200020072001417f6a20072001491b2001200b110000200841306a240020070b140020022003490440200120026a20003a00000b0b0300010b4501037f20002802002101034020012d000041506a41ff017141094b4504402000200141016a220336020020012c00002002410a6c6a41506a2102200321010c010b0b20020ba50101067f230041206b220b240020092009416f7120041b210c02402004450440200c418008710d010b200c41207141e1007341f6016a210d4100210903402009200b6a2004200420066e220e20066c6b220a4130200d200a41187441808080d000481b6a3a0000200941016a210a2009411e4b0d01200420064f200a2109200e21040d000b0b2000200120022003200b200a2005200620072008200c102b200b41206a24000bae0102057f017e230041206b220b240020092009416f71200442005222091b210c02402009450440200c418008710d010b200c41207141e1007341f6016a210d4100210903402009200b6a4130200d20042004200680220f20067e7da7220a41187441808080d000481b200a6a3a0000200941016a210a2009411e4b0d01200420065a200a2109200f21040d000b0b2000200120022003200b200a20052006a720072008200c102b200b41206a24000be70401037f200a410271210c2005210b02400340200c0d01200b411f4b200b20084f724504402004200b6a41303a0000200b41016a210b0c010b0b200b21050b200a410371410147210d2005210b02400340200d0d01200b411f4b200b20094f724504402004200b6a41303a0000200b41016a210b0c010b0b200b21050b200a410171210d024002400240200a4110710440200545200a41800871722005200847410020052009471b724504402005417e6a2005417f6a220520051b200520074110461b21050b024020074110460440200a41207122072005411f4b72450440200420056a41f8003a0000200541016a21050c020b2007452005411f4b720d01200420056a41d8003a0000200541016a21050c010b20074102472005411f4b720d00200420056a41e2003a0000200541016a21050b2005411f4b0d01200420056a41303a0000200541016a21050c010b20050d00410021050c010b2009200a410c714100472006726b200520052009461b2205411f4b0d010b20060440200420056a412d3a0000200541016a21050c010b200a4104710440200420056a412b3a0000200541016a21050c010b200a410871450d00200420056a41203a0000200541016a21050b2002210b0240200c200d720d00200521060340200620094f0d0141202001200b20032000110000200641016a2106200b41016a210b0c000b000b2004417f6a2104037f2005047f200420056a2c00002001200b200320001100002005417f6a2105200b41016a210b0c01050240200c450d00410020026b210203402002200b6a20094f0d0141202001200b20032000110000200b41016a210b0c000b000b200b0b0b0b2601017f230041106b220324002003200236020c2000200141d00a20021025200341106a24000bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000be10201027f02402001450d00200041003a0000200020016a2202417f6a41003a000020014103490d00200041003a0002200041003a00012002417d6a41003a00002002417e6a41003a000020014107490d00200041003a00032002417c6a41003a000020014109490d002000410020006b41037122036a220241003602002002200120036b417c7122036a2201417c6a410036020020034109490d002002410036020820024100360204200141786a4100360200200141746a410036020020034119490d002002410036021820024100360214200241003602102002410036020c200141706a41003602002001416c6a4100360200200141686a4100360200200141646a41003602002003200241047141187222036b2101200220036a2102034020014120490d0120024200370300200241186a4200370300200241106a4200370300200241086a4200370300200241206a2102200141606a21010c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d0440200020012002102d1a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041908d0436020c41fc0c200028020c41076a417871220036020041800d200036020041840d3f003602000b970101047f230041106b220124002001200036020c2000047f41840d200041086a2202411076220041840d2802006a220336020041800d200241800d28020022026a41076a417871220436020002400240200341107420044d044041840d200341016a360200200041016a21000c010b2000450d010b200040000d0010010b20022001410c6a4104102d41086a0541000b200141106a24000b0b002000410120001b10310ba10101037f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b20012802082103024020012802042201410a4d0440200020014101743a0000200041016a21020c010b200141106a4170712204103221022000200136020420002004410172360200200020023602080b2002200320011034200120026a41003a000020000b100020020440200020012002102d1a0b0b130020002d0000410171044020002802081a0b0b3401017f2000200147044020002001280208200141016a20012d0000220041017122021b2001280204200041017620021b10370b0ba00101037f410a210320002d00002205410171220404402000280200417e71417f6a21030b200320024f0440027f2004044020002802080c010b200041016a0b210320020440200320012002102f0b200220036a41003a000020002d00004101710440200020023602040f0b200020024101743a00000f0b20002003200220036b027f2004044020002802040c010b20054101760b2200410020002002200110380bc30101027f027f20002d0000410171044020002802080c010b200041016a0b2109416f2108200141e6ffffff074d0440410b20014101742208200120026a220220022008491b220241106a4170712002410b491b21080b2008103221022004044020022009200410340b20060440200220046a2007200610340b200320056b220320046b22070440200220046a20066a200420096a20056a200710340b200020023602082000200320066a220136020420002008410172360200200120026a41003a00000bb60101037f0240027f41d40a2d000022024101712204044041d80a280200210241d40a280200417e71417f6a0c010b20024101762102410a0b220320026b20014f04402001450d01027f2004044041dc0a2802000c010b41d50a0b220320026a200020011034200120026a2100024041d40a2d0000410171044041d80a20003602000c010b41d40a20004101743a00000b200020036a41003a00000f0b41d40a2003200120026a20036b2002200241002001200010380b0b0a0020002000102310390bcf0101067f230041206b22032400200341106a100a22012001280200417e71417f6a410a20012d00004101711b103c200128020420012d0000220241017620024101711b2104200141016a210503400240200128020821062003410f3602002001027f2006200520024101711b200441016a2003102c220241004e0440200220044d0d0220020c010b20044101744101720b2204103c20012d000021020c010b0b20012002103c200041086a200141086a280200360200200020012902003702002001100b20011035200341206a24000be50201047f0240027f20002d000022024101712203044020002802040c010b20024101760b22042001490440200120046b2204450d01027f2003044020002802002202417e71417f6a210320002802040c010b410a210320024101760b2101200320016b2004490440027f2002410171044020002802080c010b200041016a0b2105416f2102200341e6ffffff074d0440410b20034101742202200120046a220320032002491b220241106a4170712002410b491b21020b2002103221032001044020032005200110340b200020033602082000200241017222023602000b027f2002410171044020002802080c010b200041016a0b220220016a2004102e1a200120046a2101024020002d00004101710440200020013602040c010b200020014101743a00000b200120026a41003a00000f0b20030440200028020820016a41003a0000200020013602040f0b200020016a41016a41003a0000200020014101743a00000b0b880101037f41e00a410136020041e40a2802002100034020000440034041e80a41e80a2802002201417f6a2202360200200141014845044041e00a4100360200200020024102746a22004184016a280200200041046a28020011010041e00a410136020041e40a28020021000c010b0b41e80a412036020041e40a200028020022003602000c010b0b0b970101027f41e00a410136020041e40a280200220145044041e40a41ec0a36020041ec0a21010b024041e80a2802002202412046044041840210312201450d012001418402102e220141e40a28020036020041e40a200136020041e80a4100360200410021020b41e80a200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41e00a41003602000b070041f00c10350b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000104120012802044f0d002002410471044010010c010b200042003702000b02402002411071450d002000104120012802044d0d0020024104710440100120000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001042200010436a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000104441012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100120002802040520010b4102490d0020002802002d00010d0010010b200241054f044010010b20002802002d000145044010010b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100120002802040520010b4102490d0020002802002d00010d0010010b200241054f044010010b20002802002d000145044010010b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10010b20020b4101017f200028020445044010010b0240200028020022012d0000418101470d00200028020441014d047f100120002802000520010b2c00014100480d0010010b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a411410401041200241306a24000b2101017f20011043220220012802044b044010010b2000200120011042200210450bd60202077f017e230041206b220324002001280208220420024b0440200341186a2001104720012003280218200328021c104636020c200341106a20011047410021042001027f410020032802102206450d001a410020032802142208200128020c2207490d001a200820072007417f461b210520060b360210200141146a2005360200200141003602080b200141106a210903400240200420024f0d002001280214450d00200341106a2001104741002104027f410020032802102207450d001a410020032802142208200128020c2206490d001a200820066b2104200620076a0b21052001200436021420012005360210200341106a20094100200520041046104520012003290310220a3702102001200128020c200a422088a76a36020c2001200128020841016a22043602080c010b0b20032009290200220a3703082003200a37030020002003411410401a200341206a24000b980101037f200028020445044041000f0b20001044200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100120002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100120002802000520020b20016a2d00004100470b0bf80101057f0340024020002802102201200028020c460d00200141786a28020041014904401001200028021021010b200141786a2202200228020041016b220436020020040d002000200236021020004101200028020422032001417c6a28020022026b22011021220441016a20014138491b220520036a104b200220002802006a220320056a20032001102f0240200141374d0440200028020020026a200141406a3a00000c010b200441f7016a220341ff014d0440200028020020026a20033a00002000280200200220046a6a210203402001450d02200220013a0000200141087621012002417f6a21020c000b000b10010b0c010b0b0b0f0020002001104c200020013602040b2f01017f200028020820014904402001103120002802002000280204102d210220002001360208200020023602000b0b1b00200028020420016a220120002802084b044020002001104c0b0b250020004101104d200028020020002802046a20013a00002000200028020441016a3602040be70101037f2001280200210441012102024002400240024020012802042201410146044020042c000022014100480d012000200141ff0171104e0c040b200141374b0d01200121020b200020024180017341ff0171104e0c010b20011021220241b7016a22034180024e044010010b2000200341ff0171104e2000200028020420026a104b200028020420002802006a417f6a210320012102037f2002047f200320023a0000200241087621022003417f6a21030c010520010b0b21020b20002002104d200028020020002802046a20042002102d1a2000200028020420026a3602040b2000104a0bb80202037f037e024020015004402000418001104e0c010b20014280015a044020012105034020052006845045044020064238862005420888842105200241016a2102200642088821060c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010010b2000200441b77f6a41ff0171104e2000200028020420046a104b200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff0171104e0b2000200028020420026a104b200028020420002802006a417f6a210203402001200784500d02200220013c0000200742388620014208888421012002417f6a2102200742088821070c000b000b20002001a741ff0171104e0b2000104a0b0bd30202004180080bf90176616c7565203e3d20313030002f686f6d652f6a757a69782f71637869616f2f6175746f746573742f506c61744f4e2d476f2f63617365732f436f6e7472616374734175746f54657374732f7372632f746573742f7265736f75726365732f636f6e7472616374732f7761736d2f636f6e74726163745f7465726d696e6174696f6e2f436f6e74726163745f7465726d696e6174696f6e2e637070007472616e736665725f617373657274006261642076616c756500696e6974006765745f737472696e675f73746f7261676500417373657274696f6e206661696c65643a0066756e633a006c696e653a0066696c653a00002573090a00200041860a0b4cf03f000000000000244000000000000059400000000000408f40000000000088c34000000000006af8400000000080842e4100000000d01263410000000084d797410000000065cdcd412575";

    private static String BINARY = BINARY_0;

    public static final String FUNC_TRANSFER_ASSERT = "transfer_assert";

    public static final String FUNC_GET_STRING_STORAGE = "get_string_storage";

    protected Contract_termination(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected Contract_termination(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<Contract_termination> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(Contract_termination.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<Contract_termination> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(Contract_termination.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public RemoteCall<TransactionReceipt> transfer_assert(String name, Long value) {
        final WasmFunction function = new WasmFunction(FUNC_TRANSFER_ASSERT, Arrays.asList(name,value), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<String> get_string_storage() {
        final WasmFunction function = new WasmFunction(FUNC_GET_STRING_STORAGE, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static Contract_termination load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new Contract_termination(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static Contract_termination load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new Contract_termination(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
