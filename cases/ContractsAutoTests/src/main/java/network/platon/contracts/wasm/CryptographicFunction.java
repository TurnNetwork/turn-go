package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.WasmAddress;
import java.math.BigInteger;
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
 * <p>Please use the <a href="https://github.com/PlatONnetwork/client-sdk-java/releases">platon-web3j command line tools</a>,
 * or the org.web3j.codegen.WasmFunctionWrapperGenerator in the 
 * <a href="https://github.com/PlatONnetwork/client-sdk-java/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with platon-web3j version 0.9.1.2-SNAPSHOT.
 */
public class CryptographicFunction extends WasmContract {
    private static String BINARY_0 = "0x0061736d01000000013d0b60000060017f0060047f7f7f7f017f60037f7f7f006000017f60027f7f0060037f7f7f017f60017f017f60027f7f017f60047f7f7f7f0060017f017e029f010703656e760c706c61746f6e5f70616e6963000003656e7610706c61746f6e5f65637265636f766572000203656e7610706c61746f6e5f726970656d64313630000303656e760d706c61746f6e5f736861323536000303656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000403656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e00050345440000060607070608010007010601010701070100060106070707070905030701070603050507050505030306010001070709030300050a050801070507070507050501000405017001030305030100020615037f0141808b040b7f0041808b040b7f0041800b0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974001a06696e766f6b65003b0908010041010b021c350ad04a44080010101034104a0b02000bc60a010b7f2002410f6a210341002104410020026b21052002410e6a2106410120026b21072002410d6a2108410220026b210902400340200020046a210b200120046a210a20022004460d01200a410371450d01200b200a2d00003a00002003417f6a2103200541016a21052006417f6a2106200741016a21072008417f6a2108200941016a2109200441016a21040c000b0b200220046b210c02400240024002400240200b410371220d450d00200c4120490d03200d4101460d01200d4102460d02200d4103470d03200b200120046a28020022063a0000200041016a210c200220046b417f6a21092004210b0240034020094113490d01200c200b6a220a2001200b6a220741046a2802002208411874200641087672360200200a41046a200741086a2802002206411874200841087672360200200a41086a2007410c6a2802002208411874200641087672360200200a410c6a200741106a2802002206411874200841087672360200200b41106a210b200941706a21090c000b0b2002417f6a2005416d2005416d4b1b20036a4170716b20046b210c2001200b6a41016a210a2000200b6a41016a210b0c030b200c210a02400340200a4110490d01200020046a220b200120046a2207290200370200200b41086a200741086a290200370200200441106a2104200a41706a210a0c000b0b02400240200c4108710d00200120046a210a200020046a21040c010b200020046a220b200120046a2204290200370200200441086a210a200b41086a21040b0240200c410471450d002004200a280200360200200a41046a210a200441046a21040b0240200c410271450d002004200a2f00003b0000200441026a2104200a41026a210a0b200c410171450d032004200a2d00003a000020000f0b200b200120046a220a28020022063a0000200b41016a200a41016a2f00003b0000200041036a210c200220046b417d6a21052004210b0240034020054111490d01200c200b6a220a2001200b6a220741046a2802002203410874200641187672360200200a41046a200741086a2802002206410874200341187672360200200a41086a2007410c6a2802002203410874200641187672360200200a410c6a200741106a2802002206410874200341187672360200200b41106a210b200541706a21050c000b0b2002417d6a2009416f2009416f4b1b20086a4170716b20046b210c2001200b6a41036a210a2000200b6a41036a210b0c010b200b200120046a220a28020022083a0000200b41016a200a41016a2d00003a0000200041026a210c200220046b417e6a21052004210b0240034020054112490d01200c200b6a220a2001200b6a220941046a2802002203411074200841107672360200200a41046a200941086a2802002208411074200341107672360200200a41086a2009410c6a2802002203411074200841107672360200200a410c6a200941106a2802002208411074200341107672360200200b41106a210b200541706a21050c000b0b2002417e6a2007416e2007416e4b1b20066a4170716b20046b210c2001200b6a41026a210a2000200b6a41026a210b0b0240200c411071450d00200b200a2d00003a0000200b200a280001360001200b200a290005370005200b200a2f000d3b000d200b200a2d000f3a000f200b41106a210b200a41106a210a0b0240200c410871450d00200b200a290000370000200b41086a210b200a41086a210a0b0240200c410471450d00200b200a280000360000200b41046a210b200a41046a210a0b0240200c410271450d00200b200a2f00003b0000200b41026a210b200a41026a210a0b200c410171450d00200b200a2d00003a00000b20000bfb0202027f017e02402002450d00200020013a0000200020026a2203417f6a20013a000020024103490d00200020013a0002200020013a00012003417d6a20013a00002003417e6a20013a000020024107490d00200020013a00032003417c6a20013a000020024109490d002000410020006b41037122046a2203200141ff017141818284086c22013602002003200220046b417c7122046a2202417c6a200136020020044109490d002003200136020820032001360204200241786a2001360200200241746a200136020020044119490d002003200136021820032001360214200320013602102003200136020c200241706a20013602002002416c6a2001360200200241686a2001360200200241646a20013602002001ad220542208620058421052004200341047141187222016b2102200320016a2101034020024120490d0120012005370300200141186a2005370300200141106a2005370300200141086a2005370300200141206a2101200241606a21020c000b0b20000b3a01017f23808080800041106b220141808b84800036020c2000200128020c41076a41787122013602042000200136020020003f0036020c20000b120041808880800020004108108d808080000bca0101067f23808080800041106b22032480808080002003200136020c024002400240024002402001450d002000200028020c200241036a410020026b220471220520016a220641107622016a220736020c200020022000280204220820066a6a417f6a2004712202360204200741107420024d0d0120010d020c030b410021000c030b2000410c6a200741016a360200200141016a21010b200140000d001080808080000b20082003410c6a41041089808080001a200820056a21000b200341106a24808080800020000b2e000240418088808000200120006c22004108108d808080002201450d00200141002000108a808080001a0b20010b02000b0f00418088808000108b808080001a0b3a01027f2000410120001b2101024003402001108c8080800022020d014100280290888080002200450d012000118080808000000c000b0b20020b0a002000108f808080000bce0301067f024020002001460d000240024002400240200120006b20026b410020024101746b4d0d0020012000734103712103200020014f0d012003450d02200021030c030b2000200120021089808080000f0b024020030d002001417f6a210402400340200020026a2203410371450d012002450d052003417f6a200420026a2d00003a00002002417f6a21020c000b0b2000417c6a21032001417c6a2104034020024104490d01200320026a200420026a2802003602002002417c6a21020c000b0b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b0b200241046a21052002417f7321064100210402400340200120046a2107200020046a2208410371450d0120022004460d03200820072d00003a00002005417f6a2105200641016a2106200441016a21040c000b0b200220046b2101410021030240034020014104490d01200820036a200720036a280200360200200341046a21032001417c6a21010c000b0b200720036a2101200820036a210320022006417c2006417c4b1b20056a417c716b20046b21020b03402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b0b20000b0900200041013602000b0900200041003602000b1d00024020002d0000410171450d0020002802081092808080000b20000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b1d0020004200370200200041086a4100360200200010978080800020000b0900108880808000000bb60101037f4194888080001094808080004100280298888080002100024003402000450d01024003404100410028029c888080002202417f6a220136029c8880800020024101480d01200020014102746a22004184016a2802002102200041046a2802002100419488808000109580808000200220001181808080000041948880800010948080800041002802988880800021000c000b0b4100412036029c88808000410020002802002200360298888080000c000b0b0bcd0101027f419488808000109480808000024041002802988880800022030d0041a0888080002103410041a088808000360298888080000b02400240410028029c8880800022044120470d004184024101108e808080002203450d0141002104200341002802988880800036020041002003360298888080004100410036029c888080000b4100200441016a36029c88808000200320044102746a22034184016a2001360200200341046a200036020041948880800010958080800041000f0b419488808000109580808000417f0b0f0041a48a8080001096808080001a0b89010020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000109e8080800020012802044f0d00024020024104710d00200042003702000c010b1080808080000b024002402002411071450d002000109e8080800020012802044d0d0020024104710d01200042003702000b20000f0b10808080800020000b3400024002402000280204450d0020002802002c0000417f4c0d0141010f0b41000f0b2000109f80808000200010a0808080006a0b280002402000280204450d0020002802002c0000417f4c0d0041000f0b200010a58080800041016a0bc90301047f0240024002402000280204450d00200010a6808080004101210120002802002c00002202417f4c0d010c020b41000f0b0240200241ff0171220141b7014b0d00200141807f6a0f0b024002400240200241ff0171220241bf014b0d000240200041046a22032802002202200141c97e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b200141384f0d010c020b0240200241f7014b0d00200141c07e6a0f0b0240200041046a22032802002202200141897e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b20014138490d010b200141ff7d490d010b10808080800020010f0b20010b5102017f017e23808080800041306b220124808080800020012000290200220237031020012002370308200141186a200141086a4114109d80808000109e808080002100200141306a24808080800020000b6a01037f02400240024020012802002204450d0041002105200320026a200128020422064b0d0120062002490d014100210120062003490d02200620026b20032003417f461b2101200420026a21050c020b410021050b410021010b20002001360204200020053602000b3301017f0240200110a080808000220220012802044d0d001080808080000b200020012001109f80808000200210a2808080000bd003020a7f017e23808080800041c0006b220324808080800002402001280208220420024d0d00200341386a200110a380808000200320032903383703182001200341186a10a18080800036020c200341306a200110a380808000410021044100210541002106024020032802302207450d00410021054100210620032802342208200128020c2209490d00200820092009417f461b2105200721060b20012006360210200141146a2005360200200141086a41003602000b200141106a2106200141146a21092001410c6a2107200141086a210802400340200420024f0d012009280200450d01200341306a200110a38080800041002104024002402003280230220a450d00410021052003280234220b2007280200220c490d01200a200c6a2105200b200c6b21040c010b410021050b20092004360200200620053602002003200436022c2003200536022820032003290328370310200341306a20064100200341106a10a18080800010a28080800020062003290330220d37020020072007280200200d422088a76a3602002008200828020041016a22043602000c000b0b20032006290200220d3703202003200d3703082000200341086a4114109d808080001a200341c0006a2480808080000b4701017f4100210102402000280204450d00024020002802002d0000220041bf014b0d00200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b5401017f024020002802040d001080808080000b0240200028020022012d0000418101470d000240200041046a28020041014b0d00108080808000200028020021010b20012c00014100480d001080808080000b0bbc0101047f024002402000280204450d00200010a680808000200028020022012c000022024100480d0120024100470f0b41000f0b410121030240200241807f460d000240200241ff0171220441b7014b0d000240200041046a28020041014b0d00108080808000200028020021010b20012d00014100470f0b41002103200441bf014b0d000240200041046a280200200241ff017141ca7e6a22024b0d00108080808000200028020021010b200120026a2d000041004721030b20030b2701017f200020012802002203200320012802046a10a9808080002000200210aa8080800020000b34002000200220016b220210ab80808000200028020020002802046a200120021089808080001a2000200028020420026a3602040bb60201087f02402001450d002000410c6a2102200041106a2103200041046a21040340200328020022052002280200460d010240200541786a28020020014f0d00108080808000200328020021050b200541786a2206200628020020016b220136020020010d01200320063602002000410120042802002005417c6a28020022016b220510ac80808000220741016a20054138491b2206200428020022086a10ad80808000200120002802006a220920066a2009200820016b1093808080001a02400240200541374b0d00200028020020016a200541406a3a00000c010b0240200741f7016a220641ff014b0d00200028020020016a20063a00002000280200200720016a6a210103402005450d02200120053a0000200541087621052001417f6a21010c000b0b1080808080000b410121010c000b0b0b21000240200028020420016a220120002802084d0d002000200110ae808080000b0b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b13002000200110ae80808000200020013602040b4501017f0240200028020820014f0d002001108c808080002202200028020020002802041089808080001a200010b380808000200041086a2001360200200020023602000b0b29002000410110ab80808000200028020020002802046a20013a00002000200028020441016a3602040b3c01017f0240200110ac80808000220320026a2202418002480d001080808080000b2000200241ff017110af8080800020002001200310b1808080000b44002000200028020420026a10ad80808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0bf90101037f23808080800041106b22032480808080002001280200210420012802042105024002402002450d004100210102400340200420016a2102200120054f0d0120022d00000d01200141016a21010c000b0b200520016b21050c010b200421020b0240024002400240024020054101470d0020022c000022014100480d012000200141ff017110af808080000c040b200541374b0d010b200020054180017341ff017110af808080000c010b2000200541b70110b0808080000b2003200536020c200320023602082003200329030837030020002003410010a8808080001a0b2000410110aa80808000200341106a24808080800020000b1700024020002802002200450d002000108f808080000b0b240041a48a8080001098808080001a4181808080004100418088808000109b808080001a0b0f0041b08a8080001096808080001a0b2a01017f410021010240034020014114460d01200020016a41003a0000200141016a21010c000b0b20000b2a01017f410021010240034020014120460d01200020016a41003a0000200141016a21010c000b0b20000b2c01017f200010b6808080002100200220032802002204200328020420046b41ff017120001081808080001a0b2501017f200010b680808000210020022802002203200228020420036b20001082808080000b2501017f200010b780808000210020022802002203200228020420036b20001083808080000b9f0b03047f017e067f23808080800041e0016b22002480808080001087808080001084808080002201108c808080002202108580808000200020013602342000200236023020002000290330370310200041306a200041186a200041106a411c109d808080002203410010a480808000200041306a10a68080800002400240200041306a10a780808000450d002000280234450d0020002802302d000041c001490d010b1080808080000b200041a8016a200041306a10bc80808000024020002802ac0122014109490d001080808080000b4200210420002802a8012102024003402001450d012001417f6a210120044208862002310000842104200241016a21020c000b0b0240024020044200510d00200441bc8a80800010bd80808000510d010240200441c18a80800010bd80808000520d00200041306a10b7808080001a2000410036025820004200370350200041e0006a2003410110a480808000200041e0006a10a680808000200041c8016a200041e0006a10bc80808000200041306a41206a210120002802cc012102024002402000280264450d00200241204b0d0020002802602d000041ff017141c001490d010b1080808080000b200041a8016a10b7808080002002412020024120491b22026b41206a20002802c80120021089808080001a200041306a41186a2202200041a8016a41186a2205290300370300200041306a41106a2206200041a8016a41106a2207290300370300200041306a41086a2208200041a8016a41086a2209290300370300200020002903a801370330200041a8016a2003410210a480808000200041a8016a200110be80808000200041e0006a41186a22032002290300370300200041e0006a41106a220a2006290300370300200041e0006a41086a220620082903003703002000200029033037036020004198016a200110bf808080002102200520032903003703002007200a29030037030020092006290300370300200020002903603703a801200041c8016a20004188016a200041a8016a200210b880808000200041c8016a10c080808000200210c1808080001a200110c1808080001a0c020b0240200441d78a80800010bd80808000520d0020004100360268200042003703602003200041e0006a10c280808000200041306a200041c8016a200041a8016a200041e0006a10bf80808000220110b980808000200041306a10c080808000200110c1808080001a200041e0006a10c1808080001a0c020b200441ed8a80800010bd80808000520d004100210120004100360290012000420037038801200320004188016a10c280808000200041a8016a20004180016a20004198016a20004188016a10bf80808000220610ba80808000200041c8016a10c3808080002103200041e0006a41186a4100360200200041e0006a41106a4200370300200041e0006a41086a420037030020004200370360200041306a41186a200041a8016a41186a290300370300200041306a41106a200041a8016a41106a290300370300200041306a41086a200041a8016a41086a290300370300200020002903a801370330410121050240034020014120460d01200041306a20016a2102200141016a210120022d0000450d000b412121050b20002005360260200041e0006a41047210c4808080001a2003200510c580808000200041306a41186a200041a8016a41186a290300370300200041306a41106a200041a8016a41106a290300370300200041306a41086a200041a8016a41086a29030037030020004120360264200020002903a8013703302000200041306a360260200020002903603703082003200041086a410010b2808080001a0240200328020c200341106a280200460d001080808080000b20032802002003280204108680808000200310c6808080001a200610c1808080001a20004188016a10c1808080001a0c010b1080808080000b200041e0016a2480808080000b870201057f0240200110a0808080002202200128020422034d0d00108080808000200141046a28020021030b2001280200210402400240024002400240024002402003450d004100210120042c00002205417f4c0d012004450d030c040b410021010c010b0240200541ff0171220641bf014b0d0041002101200541ff017141b801490d01200641c97e6a21010c010b41002101200541ff017141f801490d00200641897e6a21010b200141016a210120040d010b410021050c010b41002105200120026a20034b0d0020032001490d004100210620032002490d01200320016b20022002417f461b2106200420016a21050c010b410021060b20002006360204200020053602000b3a01027e42a5c688a1c89ca7f94b21010240034020003000002202500d01200041016a2100200142b383808080207e20028521010c000b0b20010b9c0201037f23808080800041206b2202248080808000024002402000280204450d0020002802002d000041c0014f0d00200241186a200010bc8080800020022802182103200241106a200010bc8080800020022802102104200010a08080800021002002410036020820024200370300200420006a20036b2200450d012002200010c78080800020004101480d012002280204200320001089808080001a2002200228020420006a3602040c010b20024100360208200242003703000b024020012802002200450d0020012000360204200010928080800020014100360208200142003702000b20012002280200360200200120022902043702042002410036020820024200370300200210c1808080001a200241206a2480808080000b6a01027f20004200370200200041003602080240200128020420012802006b2202450d002000200210c780808000200141046a280200200128020022036b22014101480d00200041046a2202280200200320011089808080001a2002200228020020016a3602000b20000bed0201057f23808080800041e0006b2201248080808000200141106a10c380808000210241002103200141c0006a4100360200200141286a41106a4200370300200141286a41086a420037030020014200370328200141c8006a41106a200041106a280000360200200141c8006a41086a200041086a29000037030020012000290000370348410121040240034020034114460d01200141c8006a20036a2105200341016a210320052d0000450d000b411521040b20012004360228200141286a41047210c4808080001a2002200410c580808000200141286a41106a200041106a280000360200200141286a41086a200041086a2900003703002001411436024c200120002900003703282001200141286a360248200120012903483703082002200141086a410010b2808080001a0240200228020c200241106a280200460d001080808080000b20022802002002280204108680808000200210c6808080001a200141e0006a2480808080000b2201017f024020002802002201450d002000200136020420011092808080000b20000b3c01017f23808080800041206b2202248080808000200241086a2000410110a480808000200241086a200110be80808000200241206a2480808080000b2d0020004100360208200042003702002000410010c880808000200041146a41003602002000420037020c20000b8f0301067f200028020422012000280210220241087641fcffff07716a210302400240200028020822042001460d002001200028021420026a220541087641fcffff07716a280200200541ff07714102746a2105200041146a21062003280200200241ff07714102746a21020c010b200041146a210641002102410021050b0240034020052002460d01200241046a220220032802006b418020470d0020032802042102200341046a21030c000b0b20064100360200200041086a210302400340200420016b41027522024103490d012001280200109280808000200041046a2201200128020041046a2201360200200328020021040c000b0b02400240024020024101460d0020024102470d0241800821020c010b41800421020b200041106a20023602000b0240034020042001460d012001280200109280808000200141046a21010c000b0b200041086a22022802002101200041046a28020021040240034020042001460d0120022001417c6a22013602000c000b0b024020002802002201450d0020011092808080000b20000b19000240200028020820014f0d002000200110c8808080000b0b2d01017f0240200028020c2201450d00200041106a200136020020011092808080000b200010c98080800020000b3801017f02402001417f4c0d00200020011091808080002202360200200020023602042000200220016a3602080f0b2000109980808000000b4401017f0240200028020820014f0d002001108c80808000200028020020002802041089808080002102200010c980808000200041086a2001360200200020023602000b0b0d002000280200108f808080000b5501017f410042003702b08a808000410041003602b88a80800041742100024003402000450d01200041bc8a8080006a4100360200200041046a21000c000b0b4182808080004100418088808000109b808080001a0b0b8e0302004180080bbc02000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041bc0a0b44696e69740063616c6c5f706c61746f6e5f65637265636f7665720063616c6c5f706c61746f6e5f726970656d643136300063616c6c5f706c61746f6e5f73686132353600";

    public static String BINARY = BINARY_0;

    public static final String FUNC_CALL_PLATON_ECRECOVER = "call_platon_ecrecover";

    public static final String FUNC_CALL_PLATON_RIPEMD160 = "call_platon_ripemd160";

    public static final String FUNC_CALL_PLATON_SHA256 = "call_platon_sha256";

    protected CryptographicFunction(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected CryptographicFunction(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<WasmAddress> call_platon_ecrecover(byte[] hash, byte[] signature) {
        final WasmFunction function = new WasmFunction(FUNC_CALL_PLATON_ECRECOVER, Arrays.asList(hash,signature), WasmAddress.class);
        return executeRemoteCall(function, WasmAddress.class);
    }

    public RemoteCall<WasmAddress> call_platon_ripemd160(byte[] data) {
        final WasmFunction function = new WasmFunction(FUNC_CALL_PLATON_RIPEMD160, Arrays.asList(data, Void.class), WasmAddress.class);
        return executeRemoteCall(function, WasmAddress.class);
    }

    public RemoteCall<byte[]> call_platon_sha256(byte[] data) {
        final WasmFunction function = new WasmFunction(FUNC_CALL_PLATON_SHA256, Arrays.asList(data, Void.class), byte[].class);
        return executeRemoteCall(function, byte[].class);
    }

    public static RemoteCall<CryptographicFunction> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CryptographicFunction.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<CryptographicFunction> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CryptographicFunction.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<CryptographicFunction> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CryptographicFunction.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<CryptographicFunction> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CryptographicFunction.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static CryptographicFunction load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new CryptographicFunction(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static CryptographicFunction load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new CryptographicFunction(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
