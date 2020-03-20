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
    private static String BINARY_0 = "0x0061736d01000000016a1160000060017f0060017f017f6000017f60027f7f0060027f7f017f60047f7f7f7f017f60047f7f7f7f0060037f7f7f017f60037f7f7f0060087f7f7f7f7f7f7f7f0060057f7f7f7f7f017f60037f7e7e017f60027e7e017f60047f7e7e7f0060017f017e60027f7e0002d2010903656e760c706c61746f6e5f70616e6963000003656e760d706c61746f6e5f6f726967696e000103656e760e706c61746f6e5f64657374726f79000203656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000303656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e000403656e7617706c61746f6e5f6765745f73746174655f6c656e677468000503656e7610706c61746f6e5f6765745f7374617465000603656e7610706c61746f6e5f7365745f7374617465000703605f0000080802020805010002010208080101010509080205080a080806080b09010201000801080202020207040902010208090404020404040909080c0d02090e010002010101010404020008040f02020210020410020402040204010202000405017001030305030100020615037f0141808b040b7f0041808b040b7f0041fd0a0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300090b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974002b06696e766f6b6500530908010041010b022d4d0ad9665f08001012104a10670b02000bc60a010b7f2002410f6a210341002104410020026b21052002410e6a2106410120026b21072002410d6a2108410220026b210902400340200020046a210b200120046a210a20022004460d01200a410371450d01200b200a2d00003a00002003417f6a2103200541016a21052006417f6a2106200741016a21072008417f6a2108200941016a2109200441016a21040c000b0b200220046b210c02400240024002400240200b410371220d450d00200c4120490d03200d4101460d01200d4102460d02200d4103470d03200b200120046a28020022063a0000200041016a210c200220046b417f6a21092004210b0240034020094113490d01200c200b6a220a2001200b6a220741046a2802002208411874200641087672360200200a41046a200741086a2802002206411874200841087672360200200a41086a2007410c6a2802002208411874200641087672360200200a410c6a200741106a2802002206411874200841087672360200200b41106a210b200941706a21090c000b0b2002417f6a2005416d2005416d4b1b20036a4170716b20046b210c2001200b6a41016a210a2000200b6a41016a210b0c030b200c210a02400340200a4110490d01200020046a220b200120046a2207290200370200200b41086a200741086a290200370200200441106a2104200a41706a210a0c000b0b02400240200c4108710d00200120046a210a200020046a21040c010b200020046a220b200120046a2204290200370200200441086a210a200b41086a21040b0240200c410471450d002004200a280200360200200a41046a210a200441046a21040b0240200c410271450d002004200a2f00003b0000200441026a2104200a41026a210a0b200c410171450d032004200a2d00003a000020000f0b200b200120046a220a28020022063a0000200b41016a200a41016a2f00003b0000200041036a210c200220046b417d6a21052004210b0240034020054111490d01200c200b6a220a2001200b6a220741046a2802002203410874200641187672360200200a41046a200741086a2802002206410874200341187672360200200a41086a2007410c6a2802002203410874200641187672360200200a410c6a200741106a2802002206410874200341187672360200200b41106a210b200541706a21050c000b0b2002417d6a2009416f2009416f4b1b20086a4170716b20046b210c2001200b6a41036a210a2000200b6a41036a210b0c010b200b200120046a220a28020022083a0000200b41016a200a41016a2d00003a0000200041026a210c200220046b417e6a21052004210b0240034020054112490d01200c200b6a220a2001200b6a220941046a2802002203411074200841107672360200200a41046a200941086a2802002208411074200341107672360200200a41086a2009410c6a2802002203411074200841107672360200200a410c6a200941106a2802002208411074200341107672360200200b41106a210b200541706a21050c000b0b2002417e6a2007416e2007416e4b1b20066a4170716b20046b210c2001200b6a41026a210a2000200b6a41026a210b0b0240200c411071450d00200b200a2d00003a0000200b200a280001360001200b200a290005370005200b200a2f000d3b000d200b200a2d000f3a000f200b41106a210b200a41106a210a0b0240200c410871450d00200b200a290000370000200b41086a210b200a41086a210a0b0240200c410471450d00200b200a280000360000200b41046a210b200a41046a210a0b0240200c410271450d00200b200a2f00003b0000200b41026a210b200a41026a210a0b200c410171450d00200b200a2d00003a00000b20000bfb0202027f017e02402002450d00200020013a0000200020026a2203417f6a20013a000020024103490d00200020013a0002200020013a00012003417d6a20013a00002003417e6a20013a000020024107490d00200020013a00032003417c6a20013a000020024109490d002000410020006b41037122046a2203200141ff017141818284086c22013602002003200220046b417c7122046a2202417c6a200136020020044109490d002003200136020820032001360204200241786a2001360200200241746a200136020020044119490d002003200136021820032001360214200320013602102003200136020c200241706a20013602002002416c6a2001360200200241686a2001360200200241646a20013602002001ad220542208620058421052004200341047141187222016b2102200320016a2101034020024120490d0120012005370300200141186a2005370300200141106a2005370300200141086a2005370300200141206a2101200241606a21020c000b0b20000b3a01017f23808080800041106b220141808b84800036020c2000200128020c41076a41787122013602042000200136020020003f0036020c20000b120041808880800020004108108f808080000bca0101067f23808080800041106b22032480808080002003200136020c024002400240024002402001450d002000200028020c200241036a410020026b220471220520016a220641107622016a220736020c200020022000280204220820066a6a417f6a2004712202360204200741107420024d0d0120010d020c030b410021000c030b2000410c6a200741016a360200200141016a21010b200140000d001080808080000b20082003410c6a4104108b808080001a200820056a21000b200341106a24808080800020000b2e000240418088808000200120006c22004108108f808080002201450d00200141002000108c808080001a0b20010b02000b0f00418088808000108d808080001a0b3a01027f2000410120001b2101024003402001108e8080800022020d014100280290888080002200450d012000118080808000000c000b0b20020b0a0020001091808080000b7a01027f200021010240024003402001410371450d0120012d0000450d02200141016a21010c000b0b2001417c6a21010340200141046a22012802002202417f73200241fffdfb776a7141808182847871450d000b0340200241ff0171450d01200141016a2d00002102200141016a21010c000b0b200120006b0bce0301067f024020002001460d000240024002400240200120006b20026b410020024101746b4d0d0020012000734103712103200020014f0d012003450d02200021030c030b200020012002108b808080000f0b024020030d002001417f6a210402400340200020026a2203410371450d012002450d052003417f6a200420026a2d00003a00002002417f6a21020c000b0b2000417c6a21032001417c6a2104034020024104490d01200320026a200420026a2802003602002002417c6a21020c000b0b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b0b200241046a21052002417f7321064100210402400340200120046a2107200020046a2208410371450d0120022004460d03200820072d00003a00002005417f6a2105200641016a2106200441016a21040c000b0b200220046b2101410021030240034020014104490d01200820036a200720036a280200360200200341046a21032001417c6a21010c000b0b200720036a2101200820036a210320022006417c2006417c4b1b20056a417c716b20046b21020b03402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b0b20000b4201027f0240024003402002450d0120002d0000220320012d00002204470d02200141016a2101200041016a21002002417f6a21020c000b0b41000f0b200320046b0b0900200041013602000b0900200041003602000b0900108a80808000000b5201017f20004200370200200041086a22024100360200024020012d00004101710d00200020012902003702002002200141086a28020036020020000f0b200020012802082001280204109c8080800020000b7701027f0240200241704f0d00024002402002410a4b0d00200020024101743a0000200041016a21030c010b200241106a417071220410938080800021032000200236020420002004410172360200200020033602080b200320012002109d808080001a200320026a41003a00000f0b108a80808000000b1a0002402002450d00200020012002108b8080800021000b20000b1d00024020002d0000410171450d0020002802081094808080000b20000b3d01027f024020002001460d0020002001280208200141016a20012d0000220241017122031b2001280204200241017620031b10a0808080001a0b20000bbb0101037f410a2103024020002d000022044101712205450d002000280200417e71417f6a21030b02400240024002400240200320024f0d0020050d01200441017621050c020b20050d02200041016a21030c030b200028020421050b20002003200220036b2005410020052002200110a18080800020000f0b200028020821030b20032001200210a2808080001a200320026a41003a0000024020002d00004101710d00200020024101743a000020000f0b2000200236020420000b8f0201037f0240416e20016b2002490d000240024020002d00004101710d00200041016a21080c010b200028020821080b416f21090240200141e6ffffff074b0d00410b21092001410174220a200220016a22022002200a491b2202410b490d00200241106a41707121090b2009109380808000210202402004450d00200220082004109d808080001a0b02402006450d00200220046a20072006109d808080001a0b0240200320056b220320046b2207450d00200220046a20066a200820046a20056a2007109d808080001a0b02402001410a460d0020081094808080000b200020023602082000200320066a220436020420002009410172360200200220046a41003a00000f0b108a80808000000b1a0002402002450d0020002001200210968080800021000b20000b1e0002402001450d002000200241ff01712001108c8080800021000b20000ba50201047f0240024020002d0000220441017122050d00200441017621040c010b200028020421040b024020042001490d00410a210602402005450d002000280200417e71417f6a21060b0240024002400240200620046b20034f0d0020002006200420036a20066b2004200141002003200210a1808080000c010b2003450d0020050d01200041016a21060c020b20000f0b200028020821060b0240200420016b2207450d00200620016a220520036a2005200710a2808080001a200220036a2002200620046a20024b1b2002200520024d1b21020b200620016a2002200310a2808080001a200420036a21040240024020002d00004101710d00200020044101743a00000c010b200020043602040b200620046a41003a000020000f0b108a80808000000b1600200020012002200210958080800010a4808080000bbb0301047f0240024020002d0000220541017122060d00200541017621050c010b200028020421050b024020052001490d00200520016b2207200220072002491b2102410a210802402006450d002000280200417e71417f6a21080b0240200220056b20086a20044f0d0020002008200520046a20026b20086b2005200120022004200310a18080800020000f0b0240024020060d00200041016a21080c010b200028020821080b02400240024020022004470d00200421020c010b200720026b2207450d00200820016a21060240200220044d0d0020062003200410a2808080001a200620046a200620026a200710a2808080001a0c020b0240200620034f0d00200820056a20034d0d000240200620026a20034d0d0020062003200210a2808080001a200420026b2106200320046a2103200220016a210141002102200621040c010b2003200420026b6a21030b200820016a220620046a200620026a200710a2808080001a0b200820016a2003200410a2808080001a0b200420026b20056a21050240024020002d00004101710d00200020054101743a00000c010b200020053602040b200820056a41003a000020000f0b108a80808000000b7701027f0240200141704f0d00024002402001410a4b0d00200020014101743a0000200041016a21030c010b200141106a417071220410938080800021032000200136020420002004410172360200200020033602080b20032001200210a3808080001a200320016a41003a00000f0b108a80808000000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b1d0020004200370200200041086a4100360200200010a88080800020000b0900108a80808000000bb60101037f4194888080001098808080004100280298888080002100024003402000450d01024003404100410028029c888080002202417f6a220136029c8880800020024101480d01200020014102746a22004184016a2802002102200041046a2802002100419488808000109980808000200220001181808080000041948880800010988080800041002802988880800021000c000b0b4100412036029c88808000410020002802002200360298888080000c000b0b0bcd0101027f419488808000109880808000024041002802988880800022030d0041a0888080002103410041a088808000360298888080000b02400240410028029c8880800022044120470d0041840241011090808080002203450d0141002104200341002802988880800036020041002003360298888080004100410036029c888080000b4100200441016a36029c88808000200320044102746a22034184016a2001360200200341046a200036020041948880800010998080800041000f0b419488808000109980808000417f0b0f0041a48a808000109e808080001a0b89010020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d00200010af8080800020012802044f0d00024020024104710d00200042003702000c010b1080808080000b024002402002411071450d00200010af8080800020012802044d0d0020024104710d01200042003702000b20000f0b10808080800020000b3400024002402000280204450d0020002802002c0000417f4c0d0141010f0b41000f0b200010b080808000200010b1808080006a0b280002402000280204450d0020002802002c0000417f4c0d0041000f0b200010b68080800041016a0bc90301047f0240024002402000280204450d00200010b7808080004101210120002802002c00002202417f4c0d010c020b41000f0b0240200241ff0171220141b7014b0d00200141807f6a0f0b024002400240200241ff0171220241bf014b0d000240200041046a22032802002202200141c97e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b200141384f0d010c020b0240200241f7014b0d00200141c07e6a0f0b0240200041046a22032802002202200141897e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b20014138490d010b200141ff7d490d010b10808080800020010f0b20010b5102017f017e23808080800041306b220124808080800020012000290200220237031020012002370308200141186a200141086a411410ae8080800010af808080002100200141306a24808080800020000b6a01037f02400240024020012802002204450d0041002105200320026a200128020422064b0d0120062002490d014100210120062003490d02200620026b20032003417f461b2101200420026a21050c020b410021050b410021010b20002001360204200020053602000b3301017f0240200110b180808000220220012802044d0d001080808080000b20002001200110b080808000200210b3808080000bd003020a7f017e23808080800041c0006b220324808080800002402001280208220420024d0d00200341386a200110b480808000200320032903383703182001200341186a10b28080800036020c200341306a200110b480808000410021044100210541002106024020032802302207450d00410021054100210620032802342208200128020c2209490d00200820092009417f461b2105200721060b20012006360210200141146a2005360200200141086a41003602000b200141106a2106200141146a21092001410c6a2107200141086a210802400340200420024f0d012009280200450d01200341306a200110b48080800041002104024002402003280230220a450d00410021052003280234220b2007280200220c490d01200a200c6a2105200b200c6b21040c010b410021050b20092004360200200620053602002003200436022c2003200536022820032003290328370310200341306a20064100200341106a10b28080800010b38080800020062003290330220d37020020072007280200200d422088a76a3602002008200828020041016a22043602000c000b0b20032006290200220d3703202003200d3703082000200341086a411410ae808080001a200341c0006a2480808080000b4701017f4100210102402000280204450d00024020002802002d0000220041bf014b0d00200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b5401017f024020002802040d001080808080000b0240200028020022012d0000418101470d000240200041046a28020041014b0d00108080808000200028020021010b20012c00014100480d001080808080000b0bbc0101047f024002402000280204450d00200010b780808000200028020022012c000022024100480d0120024100470f0b41000f0b410121030240200241807f460d000240200241ff0171220441b7014b0d000240200041046a28020041014b0d00108080808000200028020021010b20012d00014100470f0b41002103200441bf014b0d000240200041046a280200200241ff017141ca7e6a22024b0d00108080808000200028020021010b200120026a2d000041004721030b20030b2701017f200020012802002203200320012802046a10ba808080002000200210bb8080800020000b34002000200220016b220210bc80808000200028020020002802046a20012002108b808080001a2000200028020420026a3602040bb60201087f02402001450d002000410c6a2102200041106a2103200041046a21040340200328020022052002280200460d010240200541786a28020020014f0d00108080808000200328020021050b200541786a2206200628020020016b220136020020010d01200320063602002000410120042802002005417c6a28020022016b220510bd80808000220741016a20054138491b2206200428020022086a10be80808000200120002802006a220920066a2009200820016b1096808080001a02400240200541374b0d00200028020020016a200541406a3a00000c010b0240200741f7016a220641ff014b0d00200028020020016a20063a00002000280200200720016a6a210103402005450d02200120053a0000200541087621052001417f6a21010c000b0b1080808080000b410121010c000b0b0b21000240200028020420016a220120002802084d0d002000200110bf808080000b0b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b13002000200110bf80808000200020013602040b4501017f0240200028020820014f0d002001108e80808000220220002802002000280204108b808080001a200010c980808000200041086a2001360200200020023602000b0b29002000410110bc80808000200028020020002802046a20013a00002000200028020441016a3602040b3c01017f0240200110bd80808000220320026a2202418002480d001080808080000b2000200241ff017110c08080800020002001200310c2808080000b44002000200028020420026a10be80808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0bf90101037f23808080800041106b22032480808080002001280200210420012802042105024002402002450d004100210102400340200420016a2102200120054f0d0120022d00000d01200141016a21010c000b0b200520016b21050c010b200421020b0240024002400240024020054101470d0020022c000022014100480d012000200141ff017110c0808080000c040b200541374b0d010b200020054180017341ff017110c0808080000c010b2000200541b70110c1808080000b2003200536020c200320023602082003200329030837030020002003410010b9808080001a0b2000410110bb80808000200341106a24808080800020000bc40101037f02400240024020012002844200510d00200142ff005620024200522002501b0d0120002001a741ff017110c0808080000c020b200041800110c0808080000c010b024002402001200210c580808000220341374b0d00200020034180017341ff017110c0808080000c010b0240200310c680808000220441b7016a2205418002490d001080808080000b2000200541ff017110c08080800020002003200410c7808080000b200020012002200310c8808080000b2000410110bb8080800020000b3501017f41002102024003402000200184500d0120004208882001423886842100200241016a2102200142088821010c000b0b20020b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b44002000200028020420026a10be80808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0b54002000200028020420036a10be80808000200028020020002802046a417f6a2100024003402001200284500d01200020013c0000200142088820024238868421012000417f6a2100200242088821020c000b0b0b1700024020002802002200450d0020001091808080000b0b240041a48a80800010a9808080001a418180808000410041808880800010ac808080001a0b1d0020004200370200200041086a4100360200200010cc8080800020000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b0f0041b08a808000109e808080001a0b3001017f410021010240034020014114460d01200020016a41003a0000200141016a21010c000b0b20001081808080000b5701017f23808080800041306b2201248080808000200141086a10ce80808000200141206a200141086a10d080808000200041186a200141206a10d180808000200141206a109e808080001a200141306a2480808080000b8e0301067f23808080800041206b2202248080808000200241106a41086a410036020020024200370310200241106a41eb8a8080004100109c80808000200241086a4100360200200242003703002002200228021420022d0010220341017620034101711b220341286a413010a7808080000240024020022d001022044101710d0020044101762105200241106a41017221060c010b20022802142105200228021821060b410021042002410020032006200510a6808080001a20024101722105200241086a21060240034020044114460d012006280200200520022d00004101711b20036a200120046a22072d000041047641ec8a8080006a2d00003a00002006280200200520022d00004101711b20036a41016a20072d0000410f7141ec8a8080006a2d00003a0000200441016a2104200341026a21030c000b0b200241106a109e808080001a20002002410041e88a80800010a5808080002203290200370200200041086a200341086a280200360200200310cc808080002002109e808080001a200241206a2480808080000b6e000240024020002d00004101710d00200041003b01000c010b200028020841003a00002000410036020420002d0000410171450d00200041086a280200109480808000200041003602000b20002001290200370200200041086a200141086a280200360200200110cc808080000bb20201057f23808080800041306b2201248080808000200141186a10ce80808000200141086a200141186a10d080808000024002400240024002402000411c6a28020020002d001822024101762203200241017122041b2205200128020c20012d00082202410176200241017122021b470d002001280210200141086a41017220021b210220040d01410020036b2104200041186a41016a210003402004450d0320002d000020022d0000470d01200441016a2104200241016a2102200041016a21000c000b0b200141086a109e808080001a417f21020c030b2005450d00200041206a280200200220051097808080002104200141086a109e808080001a417f210220040d020c010b200141086a109e808080001a0b200141186a1082808080004521020b200141306a24808080800020020be60603047f017e017f2380808080004190016b22002480808080001089808080001083808080002201108e808080002202108480808000200041d8006a200041086a2002200110d4808080002203410010b580808000200041d8006a10b78080800002400240200041d8006a10b880808000450d00200028025c450d0020002802582d000041c001490d010b1080808080000b200041386a200041d8006a10d5808080000240200028023c22014109490d001080808080000b4200210420002802382102024003402001450d012001417f6a210120044208862002310000842104200241016a21020c000b0b0240024020044200510d000240200441bc8a80800010d680808000520d00200041d8006a10d7808080002101200041d8006a10cf80808000200110d8808080001a0c020b0240200441c18a80800010d680808000520d00200041d8006a10d7808080002102200041d8006a10d2808080002103200041206a10d9808080002101200041d0006a4100360200200041386a41106a4200370300200041c0006a420037030020004200370338200041386a2003ac22044201862004423f8785220410da8080800020002802382103200041386a41047210db808080001a2001200310dc808080002001200410dd808080000240200128020c200141106a280200460d001080808080000b20012802002001280204108580808000200110de808080001a200210d8808080001a0c020b0240200441d28a80800010d680808000520d00200041206a10cb808080002101200041d8006a2003410110b580808000200041d8006a200110df80808000200041d8006a10d7808080002102200041f0006a200041386a2001109b808080002203109f808080001a2003109e808080001a200210d8808080001a2001109e808080001a0c020b200441dd8a80800010d680808000520d00200041d8006a10d780808000210320004180016a200041f0006a109b808080002102200041386a10d9808080002201200210e08080800010dc808080002001200041206a2002109b80808000220510e1808080002005109e808080001a0240200128020c200141106a280200460d001080808080000b20012802002001280204108580808000200110de808080001a2002109e808080001a200310d8808080001a0c010b1080808080000b20004190016a2480808080000b4801017f23808080800041106b22032480808080002003200236020c200320013602082003200329030837030020002003411c10ae808080002100200341106a24808080800020000b870201057f0240200110b1808080002202200128020422034d0d00108080808000200141046a28020021030b2001280200210402400240024002400240024002402003450d004100210120042c00002205417f4c0d012004450d030c040b410021010c010b0240200541ff0171220641bf014b0d0041002101200541ff017141b801490d01200641c97e6a21010c010b41002101200541ff017141f801490d00200641897e6a21010b200141016a210120040d010b410021050c010b41002105200120026a20034b0d0020032001490d004100210620032002490d01200320016b20022002417f461b2106200420016a21050c010b410021060b20002006360204200020053602000b3a01027e42a5c688a1c89ca7f94b21010240034020003000002202500d01200041016a2100200142b383808080207e20028521010c000b0b20010bce0201087f23808080800041c0006b2201248080808000200010cb808080002102200042afb59bdd9e8485b9f800370310200041186a10cb808080002103200141286a10d9808080002204200029031010dd808080000240200428020c200441106a280200460d001080808080000b02400240024020042802002205200428020422061086808080002207450d0020014100360220200142003703182007417f4c0d02200141206a200710938080800041002007108c80808000220220076a22083602002001200836021c2001200236021820052006200220071087808080001a20012001280218220741016a200128021c2007417f736a10d480808000200310df80808000200141186a10e2808080001a200410de808080001a0c010b200410de808080001a20032002109f808080001a0b200141c0006a24808080800020000f0b200141186a10aa80808000000bca04010b7f23808080800041e0006b2201248080808000200141286a10d9808080002102200141c0006a41186a4100360200200141c0006a41106a4200370300200141c8006a420037030020014200370340200141c0006a200029031010da8080800020012802402103200141c0006a41047210db808080001a2002200310dc808080002002200029031010dd80808000200041186a21040240200228020c200241106a280200460d001080808080000b2002280204210520022802002106200141c0006a10d9808080002103200410e08080800021074101109380808000220841fe013a0000200120083602182001200841016a22093602202001200936021c0240200328020c200341106a280200460d00108080808000200128021c2109200128021821080b0240200920086b22092003280204220a6a220b20032802084d0d002003200b10e380808000200341046a280200210a0b2003280200200a6a20082009108b808080001a200341046a2208200828020020096a3602002003200128021c20076a20012802186b10dc808080002003200141086a2004109b80808000220810e1808080002008109e808080001a024002402003410c6a2209280200200341106a220a280200460d00108080808000200328020021082009280200200a280200460d011080808080000c010b200328020021080b200620052008200341046a280200108880808000200141186a10e2808080001a200310de808080001a200210de808080001a2004109e808080001a2000109e808080001a200141e0006a24808080800020000b2d0020004100360208200042003702002000410010e380808000200041146a41003602002000420037020c20000b8e0102017f017e4101210202402001428001540d004200210341002102024003402001200384500d0120014208882003423886842101200241016a2102200342088821030c000b0b024020024138490d00200210e58080800020026a21020b200241016a21020b0240200041186a280200450d00200041046a10e68080800021000b2000200028020020026a3602000b8f0301067f200028020422012000280210220241087641fcffff07716a210302400240200028020822042001460d002001200028021420026a220541087641fcffff07716a280200200541ff07714102746a2105200041146a21062003280200200241ff07714102746a21020c010b200041146a210641002102410021050b0240034020052002460d01200241046a220220032802006b418020470d0020032802042102200341046a21030c000b0b20064100360200200041086a210302400340200420016b41027522024103490d012001280200109480808000200041046a2201200128020041046a2201360200200328020021040c000b0b02400240024020024101460d0020024102470d0241800821020c010b41800421020b200041106a20023602000b0240034020042001460d012001280200109480808000200141046a21010c000b0b200041086a22022802002101200041046a28020021040240034020042001460d0120022001417c6a22013602000c000b0b024020002802002201450d0020011094808080000b20000b19000240200028020820014f0d002000200110e3808080000b0b0f0020002001420010c4808080001a0b2d01017f0240200028020c2201450d00200041106a200136020020011094808080000b200010e48080800020000be80201057f23808080800041206b2202248080808000024002400240024002402000280204450d0020002802002d000041c0014f0d00200241186a200010d580808000200010b1808080002103024020022802182200450d00200228021c220420034f0d020b41002100200241106a41003602002002420037030841002105410021040c020b200241086a10cb808080001a0c030b200241106a4100360200200242003703080240200420032003417f461b220541704f0d00200020056a21042005410a4d0d01200541106a417071220610938080800021032002200536020c20022006410172360208200220033602100c020b200241086a109a80808000000b200220054101743a0008200241086a41017221030b0240034020042000460d01200320002d00003a0000200341016a2103200041016a21000c000b0b200341003a00000b2001200241086a10d180808000200241086a109e808080001a200241206a2480808080000bde0101047f23808080800041306b2201248080808000200141286a4100360200200141206a4200370300200141186a42003703002001420037031041012102024020012000109b80808000220028020420002d00002203410176200341017122041b2203450d004101210202400240024020034101470d002000280208200041016a20041b2c0000417f4a0d030c010b200341374b0d010b200341016a21020c010b2003200310e5808080006a41016a21020b200120023602102000109e808080001a200141106a41047210db808080001a200141306a24808080800020020b6501037f23808080800041106b220224808080800020022001280208200141016a20012d0000220341017122041b36020820022001280204200341017620041b36020c2002200229030837030020002002410010c3808080001a200241106a2480808080000b2201017f024020002802002201450d002000200136020420011094808080000b20000b4401017f0240200028020820014f0d002001108e8080800020002802002000280204108b808080002102200010e480808000200041086a2001360200200020023602000b0b0d0020002802001091808080000b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0b240041b08a80800010cb808080001a418280808000410041808880800010ac808080001a0b0b8b0302004180080bbc02000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041bc0a0b41696e697400646973746f72795f636f6e7472616374007365745f737472696e67006765745f737472696e6700307800003031323334353637383961626364656600";

    public static String BINARY = BINARY_0;

    public static final String FUNC_SET_STRING = "set_string";

    public static final String FUNC_DISTORY_CONTRACT = "distory_contract";

    public static final String FUNC_GET_STRING = "get_string";

    protected ContractDistoryWithPermissionCheck(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ContractDistoryWithPermissionCheck(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<TransactionReceipt> set_string(String name) {
        final WasmFunction function = new WasmFunction(FUNC_SET_STRING, Arrays.asList(name), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> set_string(String name, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SET_STRING, Arrays.asList(name), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
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
