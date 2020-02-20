package network.platon.contracts.wasm;

import com.platon.rlp.Int32;
import java.util.Arrays;
import org.web3j.abi.WasmFunctionEncoder;
import org.web3j.abi.datatypes.WasmFunction;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.RemoteCall;
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
 * <p>Generated with web3j version 0.7.5.5-SNAPSHOT.
 */
public class MemoryFunctionInt extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001410b60017f0060000060027f7f006000017f60017f017f60027f7f017f60037f7f7f017f60037f7f7f0060057f7f7f7f7f017f60047f7f7f7f0060047f7f7f7f017f02700503656e760c706c61746f6e5f6465627567000203656e760c706c61746f6e5f70616e6963000103656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000303656e7610706c61746f6e5f6765745f696e707574000003656e760d706c61746f6e5f72657475726e0002037b7a01040000050105010106060604000001040006060007060408060500020504050400040005000604000404040902070400060a02050a06040909060402020502050205020702070602020504050404040707050505050502090405020202090002020205050609050a02040509000602050705010404060500010405017001040405030100020615037f0141d08c040b7f0041d08c040b7f0041c40c0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300050b5f5f686561705f6261736503010a5f5f646174615f656e64030206696e766f6b65000a0f5f5f66756e63735f6f6e5f65786974000d0909010041010b0307262a0ae6717a0800100c1078107e0b2201017f024020002802002201450d002000200136020420011096808080000b20000b0f00418088808000109c808080001a0b4401027f0240024020002d000022014101710d00200041016a2102200141017621000c010b20002802082102200028020421000b200220001080808080001081808080000b270020004200370200200041086a4100360200200020012001109180808000109a8080800020000bf30e02077f017e2380808080004190026b2200248080808000200041b0016a4100360200200042003703a8014100210102404100410c460d000340200041a8016a20016a4100360200200141046a2201410c470d000b0b200041003602a0012000420037039801024002400240024002401082808080002201450d002001417f4c0d02200041a0016a2001109580808000220241002001109080808000220320016a22013602002000200136029c012000200336029801200321030c010b4100210141002102410021030b2003108380808000200020023602d0012000200120026b3602d401200020002903d001370340200041e8006a20004180016a200041c0006a411c10ab80808000410010b38080800002400240024002400240024002400240200028026c450d0020002802682d000041c0014f0d000240200041e8006a10af808080002204200028026c22024d0d00200041d0016a41938b80800010898080800022011088808080002001109c808080001a200028026c21020b200028026821052002450d014100210320052c00002201417f4a0d04200141ff0171220641bf014b0d0241002103200141ff017141b801490d03200641c97e6a21030c030b200041d8016a4100360200200042003703d001410021014100410c460d060340200041d0016a20016a4100360200200141046a2201410c470d000c070b0b4101210320050d02410021060c030b41002103200141ff017141f801490d00200641897e6a21030b200341016a21030b41002106200320046a20024b0d0020022004490d004100210120022003490d01200220036b20042004417f461b2106200520036a21010c010b410021010b200041e8006a10af8080800021020240024002402001450d0020062002490d00200041d8016a4100360200200042003703d001200620022002417f461b220441704f0d05200120046a21032004410a4d0d01200441106a41707122061095808080002102200020043602d401200020064101723602d001200020023602d8010c020b41002104200041d8016a4100360200200042003703d00141002101410021030b200020044101743a00d001200041d0016a41017221020b024020032001460d000340200220012d00003a0000200241016a21022003200141016a2201470d000b0b200241003a00000b0240024020002d00a8014101710d00200041003b01a8010c010b20002802b00141003a0000200041003602ac0120002d00a801410171450d00200041b0016a280200109680808000200041003602a8010b200041a8016a41086a200041d0016a41086a280200360200200020002903d0013703a8014100210102404100410c460d000340200041d0016a20016a4100360200200141046a2201410c470d000b0b200041d0016a109c808080001a0240024002400240024020002802ac0120002d00a801220141017620014101711b450d00200041a8016a41d78a808000108b808080000d04200041a8016a41dc8a808000108b80808000450d0142002107200041b8016a41106a4200370300200041b8016a41086a4200370300200042003703b80120004188026a420037030020004180026a4200370300200041f8016a4200370300200041f0016a4200370300200041e8016a4200370300200041d0016a41106a4200370300200041d0016a41086a4200370300200042003703d0012000418f026a21010240420042c000510d000340200142012007883c00002001417f6a2101200742087c220742c000520d000b0b4137210102404137417f460d000340200041d0016a20016a41003a00002001417f6a2201417f470d000b0b200041386a200041d0016a41386a290300370300200041306a200041d0016a41306a290300370300200041286a200041d0016a41286a290300370300200041206a200041d0016a41206a290300370300200041186a200041d0016a41186a290300370300200041106a200041d0016a41106a290300370300200041086a200041d0016a41086a290300370300200020002903d001370300200041b8016a200010cf808080001a024020002802c401200041b8016a41106a280200460d00200041d0016a41fc8a80800010898080800022011088808080002001109c808080001a0b41002103200041003602d801200042003703d00120002802bc0120002802b80122046b2202450d022002417f4c0d07200041d8016a2002109580808000220120026a2203360200200020013602d001200020013602d401200120042002108f808080001a200020033602d4010c030b200041d8006a41c98a80800010898080800022011088808080002001109c808080001a0c030b200041c8006a41e98a80800010898080800022011088808080002001109c808080001a0c020b410021010b2001200320016b108480808000200041d0016a1086808080001a0240200041c4016a2802002201450d00200041c8016a200136020020011096808080000b200041b8016a1086808080001a0b20004198016a1086808080001a200041a8016a109c808080001a20004190026a2480808080000f0b20004198016a10a080808000000b200041d0016a109980808000000b200041d0016a10a080808000000b4201037f41002102024020011091808080002203200028020420002d0000220441017620044101711b470d0020004100417f20012003109d808080004521020b20020b5501017f410042003702808880800041004100360288888080004174210002404174450d0003402000418c888080006a4100360200200041046a22000d000b0b4181808080004100418088808000108e808080001a0bb60101037f418c888080001092808080004100280290888080002100024003402000450d010240034041004100280294888080002202417f6a22013602948880800020024101480d01200020014102746a22004184016a2802002102200041046a2802002100418c888080001093808080002002200011808080800000418c8880800010928080800041002802908880800021000c000b0b4100412036029488808000410020002802002200360290888080000c000b0b0bcd0101027f418c88808000109280808000024041002802908880800022030d0041988880800021034100419888808000360290888080000b0240024041002802948880800022044120470d00418402410110fc808080002203450d01410021042003410028029088808000360200410020033602908880800041004100360294888080000b4100200441016a36029488808000200320044102746a22034184016a2001360200200341046a2000360200418c8880800010938080800041000f0b418c88808000109380808000417f0bc60a010b7f2002410f6a210341002104410020026b21052002410e6a2106410120026b21072002410d6a2108410220026b210902400340200020046a210b200120046a210a20022004460d01200a410371450d01200b200a2d00003a00002003417f6a2103200541016a21052006417f6a2106200741016a21072008417f6a2108200941016a2109200441016a21040c000b0b200220046b210c02400240024002400240200b410371220d450d00200c4120490d03200d4101460d01200d4102460d02200d4103470d03200b200120046a28020022063a0000200041016a210c200220046b417f6a21092004210b0240034020094113490d01200c200b6a220a2001200b6a220741046a2802002208411874200641087672360200200a41046a200741086a2802002206411874200841087672360200200a41086a2007410c6a2802002208411874200641087672360200200a410c6a200741106a2802002206411874200841087672360200200b41106a210b200941706a21090c000b0b2002417f6a2005416d2005416d4b1b20036a4170716b20046b210c2001200b6a41016a210a2000200b6a41016a210b0c030b200c210a02400340200a4110490d01200020046a220b200120046a2207290200370200200b41086a200741086a290200370200200441106a2104200a41706a210a0c000b0b02400240200c4108710d00200120046a210a200020046a21040c010b200020046a220b200120046a2204290200370200200441086a210a200b41086a21040b0240200c410471450d002004200a280200360200200a41046a210a200441046a21040b0240200c410271450d002004200a2f00003b0000200441026a2104200a41026a210a0b200c410171450d032004200a2d00003a000020000f0b200b200120046a220a28020022063a0000200b41016a200a41016a2f00003b0000200041036a210c200220046b417d6a21052004210b0240034020054111490d01200c200b6a220a2001200b6a220741046a2802002203410874200641187672360200200a41046a200741086a2802002206410874200341187672360200200a41086a2007410c6a2802002203410874200641187672360200200a410c6a200741106a2802002206410874200341187672360200200b41106a210b200541706a21050c000b0b2002417d6a2009416f2009416f4b1b20086a4170716b20046b210c2001200b6a41036a210a2000200b6a41036a210b0c010b200b200120046a220a28020022083a0000200b41016a200a41016a2d00003a0000200041026a210c200220046b417e6a21052004210b0240034020054112490d01200c200b6a220a2001200b6a220941046a2802002203411074200841107672360200200a41046a200941086a2802002208411074200341107672360200200a41086a2009410c6a2802002203411074200841107672360200200a410c6a200941106a2802002208411074200341107672360200200b41106a210b200541706a21050c000b0b2002417e6a2007416e2007416e4b1b20066a4170716b20046b210c2001200b6a41026a210a2000200b6a41026a210b0b0240200c411071450d00200b200a2d00003a0000200b200a280001360001200b200a290005370005200b200a2f000d3b000d200b200a2d000f3a000f200b41106a210b200a41106a210a0b0240200c410871450d00200b200a290000370000200b41086a210b200a41086a210a0b0240200c410471450d00200b200a280000360000200b41046a210b200a41046a210a0b0240200c410271450d00200b200a2f00003b0000200b41026a210b200a41026a210a0b200c410171450d00200b200a2d00003a00000b20000bfb0202027f017e02402002450d00200020013a0000200020026a2203417f6a20013a000020024103490d00200020013a0002200020013a00012003417d6a20013a00002003417e6a20013a000020024107490d00200020013a00032003417c6a20013a000020024109490d002000410020006b41037122046a2203200141ff017141818284086c22013602002003200220046b417c7122046a2202417c6a200136020020044109490d002003200136020820032001360204200241786a2001360200200241746a200136020020044119490d002003200136021820032001360214200320013602102003200136020c200241706a20013602002002416c6a2001360200200241686a2001360200200241646a20013602002001ad220542208620058421052004200341047141187222016b2102200320016a2101034020024120490d0120012005370300200141186a2005370300200141106a2005370300200141086a2005370300200141206a2101200241606a21020c000b0b20000b7a01027f200021010240024003402001410371450d0120012d0000450d02200141016a21010c000b0b2001417c6a21010340200141046a22012802002202417f73200241fffdfb776a7141808182847871450d000b0340200241ff0171450d01200141016a2d00002102200141016a21010c000b0b200120006b0b0900200041013602000b0900200041003602000b02000b3a01027f2000410120001b210102400340200110fa8080800022020d01410028029c8a8080002200450d012000118180808000000c000b0b20020b0a00200010fd808080000bce0301067f024020002001460d000240024002400240200120006b20026b410020024101746b4d0d0020012000734103712103200020014f0d012003450d02200021030c030b200020012002108f808080000f0b024020030d002001417f6a210402400340200020026a2203410371450d012002450d052003417f6a200420026a2d00003a00002002417f6a21020c000b0b2000417c6a21032001417c6a2104034020024104490d01200320026a200420026a2802003602002002417c6a21020c000b0b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b0b200241046a21052002417f7321064100210402400340200120046a2107200020046a2208410371450d0120022004460d03200820072d00003a00002005417f6a2105200641016a2106200441016a21040c000b0b200220046b2101410021030240034020014104490d01200820036a200720036a280200360200200341046a21032001417c6a21010c000b0b200720036a2101200820036a210320022006417c2006417c4b1b20056a417c716b20046b21020b03402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b0b20000b4201027f0240024003402002450d0120002d0000220320012d00002204470d02200141016a2101200041016a21002002417f6a21020c000b0b41000f0b200320046b0b0900109480808000000b7701027f0240200241704f0d00024002402002410a4b0d00200020024101743a0000200041016a21030c010b200241106a417071220410958080800021032000200236020420002004410172360200200020033602080b200320012002109b808080001a200320026a41003a00000f0b109480808000000b1a0002402002450d00200020012002108f8080800021000b20000b1d00024020002d0000410171450d0020002802081096808080000b20000b9a0101027f0240024020002d0000220541017122060d00200541017621050c010b200028020421050b02402004417f460d0020052001490d00200520016b2205200220052002491b21020240024020060d00200041016a21000c010b200028020821000b0240200020016a200320042002200220044b22001b109e808080002201450d0020010f0b417f200020022004491b0f0b109480808000000b190002402002450d002000200120021098808080000f0b41000b270020004200370200200041086a4100360200200020012001109180808000109a8080800020000b0900109480808000000b6001017f23808080800041206b2202248080808000200241186a420037030020024200370310200242003703082000200241086a200110a28080800010a38080800010a4808080001a200241086a10a5808080001a200241206a2480808080000b4101017f23808080800041106b2202248080808000200020022001109f80808000220110e88080800021002001109c808080001a200241106a24808080800020000b5401027f23808080800041106b22012480808080000240200028020c200041106a280200460d00200141a48c808000109f80808000220210ad808080002002109c808080001a0b200141106a24808080800020000b4e01017f20004200370200200041003602080240200128020420012802006b2202450d002000200210dc80808000200041086a2001280200200141046a280200200041046a10dd808080000b20000b19002000410c6a10de808080001a200010a7808080001a20000b0f0041a08a80800010a7808080001a0b2201017f024020002802002201450d002000200136020420011096808080000b20000b4701027f23808080800041206b22012480808080002000200141086a410010a980808000220210a38080800010a4808080001a200210a5808080001a200141206a2480808080000b24002000420037020820004200370200200041106a42003702002000200110c3808080000b0f0041ac8a80800010a7808080001a0bdb0101027f23808080800041206b220324808080800020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d00200010ac8080800020012802044f0d00024020024104710d00200042003702000c010b200341106a419d8b808000109f80808000220410ad808080002004109c808080001a0b02402002411071450d00200010ac8080800020012802044d0d00024020024104710d00200042003702000c010b200341ab8b808000109f80808000220210ad808080002002109c808080001a0b200341206a24808080800020000b3400024002402000280204450d0020002802002c0000417f4c0d0141010f0b41000f0b200010ae80808000200010af808080006a0b4401027f0240024020002d000022014101710d00200041016a2102200141017621000c010b20002802082102200028020421000b200220001080808080001081808080000b280002402000280204450d0020002802002c0000417f4c0d0041000f0b200010b48080800041016a0bab0601057f23808080800041b0016b22012480808080000240024002402000280204450d00200010b5808080004101210220002802002c00002203417f4a0d02200341ff0171220241b7014b0d01200241807f6a21020c020b410021020c010b02400240200341ff0171220341bf014b0d000240200041046a22042802002203200241c97e6a22054b0d00200141a0016a41ba8b808000109f80808000220210ad808080002002109c808080001a200428020021030b024020034102490d0020002802002d00010d0020014190016a41ba8b808000109f80808000220210ad808080002002109c808080001a0b024020054105490d0020014180016a41ab8b808000109f80808000220210ad808080002002109c808080001a0b024020002802002d00010d00200141f0006a41ba8b808000109f80808000220210ad808080002002109c808080001a0b41002102410021030240034020052003460d012002410874200028020020036a41016a2d0000722102200341016a21030c000b0b200241384f0d01200141e0006a41ba8b808000109f80808000220310ad808080002003109c808080001a0c020b0240200341f7014b0d00200241c07e6a21020c020b0240200041046a22042802002203200241897e6a22054b0d00200141d0006a41ba8b808000109f80808000220210ad808080002002109c808080001a200428020021030b024020034102490d0020002802002d00010d00200141c0006a41ba8b808000109f80808000220210ad808080002002109c808080001a0b024020054105490d00200141306a41ab8b808000109f80808000220210ad808080002002109c808080001a0b024020002802002d00010d00200141206a41ba8b808000109f80808000220210ad808080002002109c808080001a0b41002102410021030240034020052003460d012002410874200028020020036a41016a2d0000722102200341016a21030c000b0b200241384f0d00200141106a41ba8b808000109f80808000220310ad808080002003109c808080001a0c010b200241ff7d490d00200141ab8b808000109f80808000220310ad808080002003109c808080001a0b200141b0016a24808080800020020b5102017f017e23808080800041306b220124808080800020012000290200220237031020012002370308200141186a200141086a411410ab8080800010ac808080002100200141306a24808080800020000b6a01037f02400240024020012802002204450d0041002105200320026a200128020422064b0d0120062002490d014100210120062003490d02200620026b20032003417f461b2101200420026a21050c020b410021050b410021010b20002001360204200020053602000b6801037f23808080800041106b22022480808080000240200110af80808000220320012802044d0d00200241bb8c808000109f80808000220410ad808080002004109c808080001a0b20002001200110ae80808000200310b180808000200241106a2480808080000bd003020a7f017e23808080800041c0006b220324808080800002402001280208220420024d0d00200341386a200110b280808000200320032903383703182001200341186a10b08080800036020c200341306a200110b280808000410021044100210541002106024020032802302207450d00410021054100210620032802342208200128020c2209490d00200820092009417f461b2105200721060b20012006360210200141146a2005360200200141086a41003602000b200141106a2106200141146a21092001410c6a2107200141086a210802400340200420024f0d012009280200450d01200341306a200110b28080800041002104024002402003280230220a450d00410021052003280234220b2007280200220c490d01200a200c6a2105200b200c6b21040c010b410021050b20092004360200200620053602002003200436022c2003200536022820032003290328370310200341306a20064100200341106a10b08080800010b18080800020062003290330220d37020020072007280200200d422088a76a3602002008200828020041016a22043602000c000b0b20032006290200220d3703202003200d3703082000200341086a411410ab808080001a200341c0006a2480808080000b4701017f4100210102402000280204450d00024020002802002d0000220041bf014b0d00200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010bc10101027f23808080800041306b2201248080808000024020002802040d00200141206a41ba8b808000109f80808000220210ad808080002002109c808080001a0b0240200028020022022d0000418101470d000240200041046a28020041014b0d00200141106a41ba8b808000109f80808000220210ad808080002002109c808080001a200028020021020b20022c00014100480d00200141ba8b808000109f80808000220010ad808080002000109c808080001a0b200141306a2480808080000b2d01017f2000200028020420012802002203200320012802046a10b7808080001a2000200210b88080800020000b970201057f23808080800041206b22042480808080000240200320026b22054101480d00024020052000280208200028020422066b4c0d00200441086a2000200520066a20002802006b10b980808000200120002802006b200041086a10ba8080800021060240034020032002460d01200641086a220528020020022d00003a00002005200528020041016a360200200241016a21020c000b0b20002006200110bb808080002101200610bc808080001a0c010b024002402005200620016b22074c0d00200041086a200220076a22082003200041046a10bd80808000200741014e0d010c020b200321080b200020012006200120056a10be8080800020022008200110bf808080001a0b200441206a24808080800020010b950301097f23808080800041206b220224808080800002402001450d002000410c6a2103200041106a2104200041046a21050340200428020022062003280200460d010240200641786a28020020014f0d00200241106a41c28b808000109f80808000220610ad808080002006109c808080001a200428020021060b200641786a2207200728020020016b220136020020010d0120042007360200200528020020002802006b2006417c6a28020022016b220610c08080800021072000200528020020002802006b22084101200741016a20064138491b22096a10c180808000200120002802006a220a20096a200a200820016b1097808080001a02400240200641374b0d00200028020020016a200641406a3a00000c010b0240200741f7016a220841ff014b0d00200028020020016a20083a00002000280200200720016a6a210103402006450d02200120063a0000200641087621062001417f6a21010c000b0b200241d68b808000109f80808000220610ad808080002006109c808080001a0b410121010c000b0b200241206a2480808080000b4c01017f02402001417f4c0d0041ffffffff0721020240200028020820002802006b220041feffffff034b0d0020012000410174220020002001491b21020b20020f0b200010a080808000000b5401017f410021042000410036020c200041106a200336020002402001450d00200110958080800021040b200020043602002000200420026a22023602082000410c6a200420016a3602002000200236020420000b8c0101027f20012802042103200041086a220420002802002002200141046a10e380808000200420022000280204200141086a10eb80808000200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c2001200128020436020020030b2301017f200010e480808000024020002802002201450d0020011096808080000b20000b2e000240200220016b22024101480d00200328020020012002108f808080001a2003200328020020026a3602000b0b5c01037f200041046a21042000280204220521062001200520036b6a2203210002400340200020024f0d01200620002d00003a00002004200428020041016a2206360200200041016a21000c000b0b20012003200510ea808080001a0b21000240200120006b2201450d002002200020011097808080001a0b200220016a0b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b4001027f02402000280204200028020022026b220320014f0d002000200120036b10c2808080000f0b0240200320014d0d00200041046a200220016a3602000b0b920101027f23808080800041206b2202248080808000024002402000280208200028020422036b20014f0d00200241086a2000200320016a20002802006b10b980808000200041046a28020020002802006b200041086a10ba808080002203200110e6808080002000200310e280808000200310bc808080001a0c010b2000200110e7808080000b200241206a2480808080000b7501017f23808080800041106b2202248080808000024002402001450d00200220013602002002200028020420002802006b3602042000410c6a200210c4808080000c010b20024100360208200242003703002000200210c5808080001a200210a7808080001a0b200241106a24808080800020000b3d01017f02402000280204220220002802084f0d0020022001290200370200200041046a2200200028020041086a3602000f0b2000200110c6808080000b5101027f23808080800041106b22022480808080002002200128020022033602082002200128020420036b36020c200220022903083703002000200210c7808080002101200241106a24808080800020010b840101027f23808080800041206b2202248080808000200241086a2000200028020420002802006b41037541016a10ec80808000200028020420002802006b410375200041086a10ed80808000220328020820012902003702002003200328020841086a3602082000200310ee80808000200310ef808080001a200241206a2480808080000b800102027f017e23808080800041206b2202248080808000024002402001280204220341374b0d002002200341406a3a001f20002002411f6a10c8808080000c010b2000200341f70110c9808080000b200220012902002204370310200220043703082000200241086a410110b6808080002100200241206a24808080800020000b3d01017f02402000280204220220002802084f0d00200220012d00003a0000200041046a2200200028020041016a3602000f0b2000200110ca808080000b7a01037f23808080800041206b22032480808080000240200110c080808000220420026a2202418002480d00200341106a418c8c808000109f80808000220510ad808080002005109c808080001a0b200320023a000f20002003410f6a10c88080800020002001200410cb80808000200341206a2480808080000b7e01027f23808080800041206b2202248080808000200241086a2000200028020441016a20002802006b10b980808000200028020420002802006b200041086a10ba80808000220328020820012d00003a00002003200328020841016a3602082000200310e280808000200310bc808080001a200241206a2480808080000b44002000200028020420026a20002802006b10c1808080002000280204417f6a2100024003402001450d01200020013a00002000417f6a2100200141087621010c000b0b0bfc0101037f23808080800041206b22032480808080002001280200210420012802042105024002402002450d004100210102400340200420016a2102200120054f0d0120022d00000d01200141016a21010c000b0b200520016b21050c010b200421020b0240024002400240024020054101470d0020022c00004100480d012000200210cd808080000c040b200541374b0d010b20032005418001733a001f20002003411f6a10c8808080000c010b2000200541b70110c9808080000b2003200536021420032002360210200320032903103703082000200341086a410010b6808080001a0b2000410110b880808000200341206a24808080800020000b3d01017f0240200028020422022000280208460d00200220012d00003a0000200041046a2200200028020041016a3602000f0b2000200110ce808080000b7e01027f23808080800041206b2202248080808000200241086a2000200028020441016a20002802006b10b980808000200028020420002802006b200041086a10ba80808000220328020820012d00003a00002003200328020841016a3602082000200310e280808000200310bc808080001a200241206a2480808080000bf20201057f23808080800041a0026b2202248080808000024002400240200110d080808000450d00200141f28b80800010d180808000450d012002200110d2808080003a009f0220002002419f026a10c8808080000c020b200041f28b80800010cd808080000c010b200241d8016a200141c000108f808080001a200241c8006a200241d8016a41c000108f808080001a02400240200241c8006a10d380808000220341374b0d0020022003418001733a009f0220002002419f026a10c8808080000c010b0240200310d480808000220441b7016a2205418002490d00200241c8016a41f38b808000109f80808000220610ad808080002006109c808080001a0b200220053a009f0220002002419f026a10c88080800020002003200410d5808080000b20024188016a200141c000108f808080001a200241086a20024188016a41c000108f808080001a2000200241086a200310d6808080000b2000410110b880808000200241a0026a24808080800020000b3b01017f23808080800041106b22012480808080002001410036020c20002001410c6a10d7808080002100200141106a24808080800020004101730b5101017f2380808080004180016b2202248080808000200241c0006a200041c000108f808080001a200241c0006a200220012d000010d88080800010d980808000210120024180016a24808080800020010b3701027f410021012000413f6a210241012100024003402000410171450d01200120022d0000722101410021000c000b0b200141ff01710b5701027f23808080800041106b220124808080800041002102024003402001410036020c20002001410c6a10da80808000450d012000410810db808080001a200241016a21020c000b0b200141106a24808080800020020b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b44002000200028020420026a20002802006b10c1808080002000280204417f6a2100024003402001450d01200020013a00002000417f6a2100200141087621010c000b0b0b54002000200028020420026a20002802006b10c1808080002000280204417f6a210002400340200110d080808000450d012000200110d2808080003a00002001410810db808080001a2000417f6a21000c000b0b0b6d01047f23808080800041c0006b22022480808080002002200128020010df808080001a4100210141012103024003402001413f4b0d01200020016a2104200220016a2105200141016a210120042d000020052d0000460d000b410021030b200241c0006a24808080800020030b1b002000410041c0001090808080002200200110e18080800020000b7501047f23808080800041c0006b22022480808080002002200141c000108f8080800021034100210441002101024003402001413f4b0d01200020016a2102200320016a2105200141016a210120022d0000220220052d00002205460d000b200220054921040b200341c0006a24808080800020040b5101017f2380808080004180016b2202248080808000200241c0006a200041c000108f808080001a200241c0006a2002200128020010df8080800010f580808000210120024180016a24808080800020010b3f01017f23808080800041c0006b220224808080800020022000200110f6808080002000200241c000108f808080002100200241c0006a24808080800020000b3801017f02402001417f4c0d00200020011095808080002202360200200020023602042000200220016a3602080f0b200010a080808000000b2e000240200220016b22024101480d00200328020020012002108f808080001a2003200328020020026a3602000b0b2201017f024020002802002201450d002000200136020420011096808080000b20000b1b002000410041c0001090808080002200200110e08080800020000b6802017f027e2000413f6a21022001ac2103420021040240034020044220510d01200220032004873c00002002417f6a2102200442087c21040c000b0b2001411f752101413b2102024003402002417f460d01200020026a20013a00002002417f6a21020c000b0b0b2d00200020013a003f413e2101024003402001417f460d01200020016a41003a00002001417f6a21010c000b0b0b7001017f200041086a20002802002000280204200141046a10e380808000200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b2f01017f20032003280200200220016b22026b2204360200024020024101480d00200420012002108f808080001a0b0b0f002000200028020410e5808080000b2d01017f20002802082102200041086a21000240034020012002460d0120002002417f6a22023602000c000b0b0b3401017f20002802082102200041086a21000340200241003a00002000200028020041016a22023602002001417f6a22010d000b0b3401017f20002802042102200041046a21000340200241003a00002000200028020041016a22023602002001417f6a22010d000b0b4501017f23808080800041106b22022480808080002002200241086a200110e98080800029020037030020002002410010cc808080002100200241106a24808080800020000b360020002001280208200141016a20012d00004101711b3602002000200128020420012d0000220141017620014101711b36020420000b23000240200120006b2201450d00200220016b2202200020011097808080001a0b20020b2e000240200220016b22024101480d00200328020020012002108f808080001a2003200328020020026a3602000b0b5301017f024020014180808080024f0d0041ffffffff0121020240200028020820002802006b220041037541feffffff004b0d0020012000410275220020002001491b21020b20020f0b200010a080808000000b5c01017f410021042000410036020c200041106a200336020002402001450d002003200110f08080800021040b200020043602002000200420024103746a22033602082000410c6a200420014103746a3602002000200336020420000b7001017f200041086a20002802002000280204200141046a10f180808000200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b2301017f200010f280808000024020002802002201450d0020011096808080000b20000b0e0020002001410010f3808080000b2f01017f20032003280200200220016b22026b2204360200024020024101480d00200420012002108f808080001a0b0b0f002000200028020410f4808080000b2300024020014180808080024f0d0020014103741095808080000f0b109480808000000b2d01017f20002802082102200041086a21000240034020012002460d012000200241786a22023602000c000b0b0b0f002000200110f7808080004101730bc10201087f23808080800041c0006b2203248080808000024002402002418004490d002000410010df808080001a0c010b02402002450d002003200141c000108f8080800021040240024020024107712205450d00200420042d003f20057622063a003f410820056b210741002101024003402001413e6a4100480d01200420016a2208413e6a220920092d00002209200576220a3a00002008413f6a20062009200774723a00002001417f6a2101200a21060c000b0b200220056b2202450d010b2004200241086d22066b2108413f21010240034020012006480d01200420016a200820016a2d00003a00002001417f6a21010c000b0b034020014100480d01200420016a41003a00002001417f6a21010c000b0b2000200441c000108f808080001a0c010b2000200141c000108f808080001a0b200341c0006a2480808080000b6e01047f23808080800041c0006b22022480808080002002200141c000108f8080800021034100210141012104024003402001413f4b0d01200020016a2102200320016a2105200141016a210120022d000020052d0000460d000b410021040b200341c0006a24808080800020040b4a0041a08a808000419c8b80800010a1808080004182808080004100418088808000108e808080001a41ac8a80800010a8808080004183808080004100418088808000108e808080001a0b3901017f23808080800041106b2201410036020c2000200128020c28020041076a41787122013602042000200136020020003f0036020c20000b120041b88a8080002000410810fb808080000bb20101037f0240024002402001450d0041002d00c88a808000450d01200028020c2103200028020421040c020b41000f0b20003f0041107422043602043f002103410041013a00c88a8080002000200336020c0b20002003200141107622056a220336020c20002002200420016a6a417f6a410020026b7122013602040240200341107420014b0d002000410c6a200341016a360200200541016a21050b200540001a200441002000410c6a2802003f00461b0b2e00024041b88a808000200120006c2200410810fb808080002201450d002001410020001090808080001a0b20010b02000b0f0041b88a80800010f9808080001a0b0bd30402004180080bc90200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041c90a0bfb0176616c6964206d6574686f640a00696e6974006765746d616c6c6f63696e74006e6f206d6574686f6420746f2063616c6c0a006c697374537461636b206973206e6f7420656d70747900626164206361737400006f7665722073697a6520726c7000756e6465722073697a6520726c700062616420726c70006974656d436f756e7420746f6f206c61726765006974656d436f756e7420746f6f206c6172676520666f7220524c5000804e756d62657220746f6f206c6172676520666f7220524c5000436f756e7420746f6f206c6172676520666f7220524c50006c697374537461636b206973206e6f7420656d70747900626164206361737400";

    public static String BINARY = BINARY_0;

    public static final String FUNC_GETMALLOCINT = "getmallocint";

    protected MemoryFunctionInt(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected MemoryFunctionInt(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<MemoryFunctionInt> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(MemoryFunctionInt.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<MemoryFunctionInt> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(MemoryFunctionInt.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public RemoteCall<Int32> getmallocint() {
        final WasmFunction function = new WasmFunction(FUNC_GETMALLOCINT, Arrays.asList(), Int32.class);
        return executeRemoteCall(function, Int32.class);
    }

    public static MemoryFunctionInt load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new MemoryFunctionInt(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static MemoryFunctionInt load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new MemoryFunctionInt(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
