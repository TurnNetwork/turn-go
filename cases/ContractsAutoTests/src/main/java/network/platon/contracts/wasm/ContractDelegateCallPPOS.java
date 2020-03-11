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
public class ContractDelegateCallPPOS extends WasmContract {
    private static String BINARY_0 = "0x0061736d01000000016d1160000060017f0060057f7f7f7f7f017f6000017f60027f7f0060037f7f7f017f60017f017f60027f7f017f60037f7f7f0060087f7f7f7f7f7f7f7f0060047f7f7f7f0060037f7e7e017f60027e7e017f60047f7e7e7f0060047f7f7f7e017e60057f7f7f7f7e0060017f017e02b9010703656e760c706c61746f6e5f70616e6963000003656e7614706c61746f6e5f64656c65676174655f63616c6c000203656e761d706c61746f6e5f6765745f63616c6c5f6f75747075745f6c656e677468000303656e7616706c61746f6e5f6765745f63616c6c5f6f7574707574000103656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000303656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e00040357560000050506060507010006010501010107080506090505020801060100050105060606060a040806010605080404060404040808050b0c06080d01000601010704060e05060604040f060704001010060406060404000405017001030305030100020615037f0141908b040b7f0041908b040b7f0041840b0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974002306696e766f6b6500530908010041010b0225450aed5e56080010101042105c0b02000bc60a010b7f2002410f6a210341002104410020026b21052002410e6a2106410120026b21072002410d6a2108410220026b210902400340200020046a210b200120046a210a20022004460d01200a410371450d01200b200a2d00003a00002003417f6a2103200541016a21052006417f6a2106200741016a21072008417f6a2108200941016a2109200441016a21040c000b0b200220046b210c02400240024002400240200b410371220d450d00200c4120490d03200d4101460d01200d4102460d02200d4103470d03200b200120046a28020022063a0000200041016a210c200220046b417f6a21092004210b0240034020094113490d01200c200b6a220a2001200b6a220741046a2802002208411874200641087672360200200a41046a200741086a2802002206411874200841087672360200200a41086a2007410c6a2802002208411874200641087672360200200a410c6a200741106a2802002206411874200841087672360200200b41106a210b200941706a21090c000b0b2002417f6a2005416d2005416d4b1b20036a4170716b20046b210c2001200b6a41016a210a2000200b6a41016a210b0c030b200c210a02400340200a4110490d01200020046a220b200120046a2207290200370200200b41086a200741086a290200370200200441106a2104200a41706a210a0c000b0b02400240200c4108710d00200120046a210a200020046a21040c010b200020046a220b200120046a2204290200370200200441086a210a200b41086a21040b0240200c410471450d002004200a280200360200200a41046a210a200441046a21040b0240200c410271450d002004200a2f00003b0000200441026a2104200a41026a210a0b200c410171450d032004200a2d00003a000020000f0b200b200120046a220a28020022063a0000200b41016a200a41016a2f00003b0000200041036a210c200220046b417d6a21052004210b0240034020054111490d01200c200b6a220a2001200b6a220741046a2802002203410874200641187672360200200a41046a200741086a2802002206410874200341187672360200200a41086a2007410c6a2802002203410874200641187672360200200a410c6a200741106a2802002206410874200341187672360200200b41106a210b200541706a21050c000b0b2002417d6a2009416f2009416f4b1b20086a4170716b20046b210c2001200b6a41036a210a2000200b6a41036a210b0c010b200b200120046a220a28020022083a0000200b41016a200a41016a2d00003a0000200041026a210c200220046b417e6a21052004210b0240034020054112490d01200c200b6a220a2001200b6a220941046a2802002203411074200841107672360200200a41046a200941086a2802002208411074200341107672360200200a41086a2009410c6a2802002203411074200841107672360200200a410c6a200941106a2802002208411074200341107672360200200b41106a210b200541706a21050c000b0b2002417e6a2007416e2007416e4b1b20066a4170716b20046b210c2001200b6a41026a210a2000200b6a41026a210b0b0240200c411071450d00200b200a2d00003a0000200b200a280001360001200b200a290005370005200b200a2f000d3b000d200b200a2d000f3a000f200b41106a210b200a41106a210a0b0240200c410871450d00200b200a290000370000200b41086a210b200a41086a210a0b0240200c410471450d00200b200a280000360000200b41046a210b200a41046a210a0b0240200c410271450d00200b200a2f00003b0000200b41026a210b200a41026a210a0b200c410171450d00200b200a2d00003a00000b20000bfb0202027f017e02402002450d00200020013a0000200020026a2203417f6a20013a000020024103490d00200020013a0002200020013a00012003417d6a20013a00002003417e6a20013a000020024107490d00200020013a00032003417c6a20013a000020024109490d002000410020006b41037122046a2203200141ff017141818284086c22013602002003200220046b417c7122046a2202417c6a200136020020044109490d002003200136020820032001360204200241786a2001360200200241746a200136020020044119490d002003200136021820032001360214200320013602102003200136020c200241706a20013602002002416c6a2001360200200241686a2001360200200241646a20013602002001ad220542208620058421052004200341047141187222016b2102200320016a2101034020024120490d0120012005370300200141186a2005370300200141106a2005370300200141086a2005370300200141206a2101200241606a21020c000b0b20000b3a01017f23808080800041106b220141908b84800036020c2000200128020c41076a41787122013602042000200136020020003f0036020c20000b120041808880800020004108108d808080000bc10101067f23808080800041106b22032480808080002003200136020c024002402001450d002000200028020c200241036a410020026b220471220520016a220641107622076a220836020c200020022000280204220120066a6a417f6a20047122023602040240200841107420024b0d002000410c6a200841016a360200200741016a21070b0240200740000d001080808080000b20012003410c6a41041089808080001a200120056a21000c010b410021000b200341106a24808080800020000b2e000240418088808000200120006c22004108108d808080002201450d00200141002000108a808080001a0b20010b02000b0f00418088808000108b808080001a0b3a01027f2000410120001b2101024003402001108c8080800022020d014100280290888080002200450d012000118080808000000c000b0b20020b0a002000108f808080000bce0301067f024020002001460d000240024002400240200120006b20026b410020024101746b4d0d0020012000734103712103200020014f0d012003450d02200021030c030b2000200120021089808080000f0b024020030d002001417f6a210402400340200020026a2203410371450d012002450d052003417f6a200420026a2d00003a00002002417f6a21020c000b0b2000417c6a21032001417c6a2104034020024104490d01200320026a200420026a2802003602002002417c6a21020c000b0b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b0b200241046a21052002417f7321064100210402400340200120046a2107200020046a2208410371450d0120022004460d03200820072d00003a00002005417f6a2105200641016a2106200441016a21040c000b0b200220046b2101410021030240034020014104490d01200820036a200720036a280200360200200341046a21032001417c6a21010c000b0b200720036a2101200820036a210320022006417c2006417c4b1b20056a417c716b20046b21020b03402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b0b20000b0900200041013602000b0900200041003602000b0900108880808000000b5201017f20004200370200200041086a22024100360200024020012d00004101710d00200020012902003702002002200141086a28020036020020000f0b20002001280208200128020410988080800020000b7701027f0240200241704f0d00024002402002410a4b0d00200020024101743a0000200041016a21030c010b200241106a417071220410918080800021032000200236020420002004410172360200200020033602080b2003200120021099808080001a200320026a41003a00000f0b108880808000000b1a0002402002450d0020002001200210898080800021000b20000b1d00024020002d0000410171450d0020002802081092808080000b20000b8f0201037f0240416e20016b2002490d000240024020002d00004101710d00200041016a21080c010b200028020821080b416f21090240200141e6ffffff074b0d00410b21092001410174220a200220016a22022002200a491b2202410b490d00200241106a41707121090b2009109180808000210202402004450d002002200820041099808080001a0b02402006450d00200220046a200720061099808080001a0b0240200320056b220320046b2207450d00200220046a20066a200820046a20056a20071099808080001a0b02402001410a460d0020081092808080000b200020023602082000200320066a220436020420002009410172360200200220046a41003a00000f0b108880808000000b1a0002402002450d0020002001200210938080800021000b20000b1e0002402001450d002000200241ff01712001108a8080800021000b20000bbb0301047f0240024020002d0000220541017122060d00200541017621050c010b200028020421050b024020052001490d00200520016b2207200220072002491b2102410a210802402006450d002000280200417e71417f6a21080b0240200220056b20086a20044f0d0020002008200520046a20026b20086b20052001200220042003109b8080800020000f0b0240024020060d00200041016a21080c010b200028020821080b02400240024020022004470d00200421020c010b200720026b2207450d00200820016a21060240200220044d0d00200620032004109c808080001a200620046a200620026a2007109c808080001a0c020b0240200620034f0d00200820056a20034d0d000240200620026a20034d0d00200620032002109c808080001a200420026b2106200320046a2103200220016a210141002102200621040c010b2003200420026b6a21030b200820016a220620046a200620026a2007109c808080001a0b200820016a20032004109c808080001a0b200420026b20056a21050240024020002d00004101710d00200020054101743a00000c010b200020053602040b200820056a41003a000020000f0b108880808000000b7701027f0240200141704f0d00024002402001410a4b0d00200020014101743a0000200041016a21030c010b200141106a417071220410918080800021032000200136020420002004410172360200200020033602080b200320012002109d808080001a200320016a41003a00000f0b108880808000000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b1d0020004200370200200041086a4100360200200010a08080800020000b0900108880808000000bb60101037f4194888080001094808080004100280298888080002100024003402000450d01024003404100410028029c888080002202417f6a220136029c8880800020024101480d01200020014102746a22004184016a2802002102200041046a2802002100419488808000109580808000200220001181808080000041948880800010948080800041002802988880800021000c000b0b4100412036029c88808000410020002802002200360298888080000c000b0b0bcd0101027f419488808000109480808000024041002802988880800022030d0041a0888080002103410041a088808000360298888080000b02400240410028029c8880800022044120470d004184024101108e808080002203450d0141002104200341002802988880800036020041002003360298888080004100410036029c888080000b4100200441016a36029c88808000200320044102746a22034184016a2001360200200341046a200036020041948880800010958080800041000f0b419488808000109580808000417f0b0f0041a48a808000109a808080001a0b89010020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d00200010a78080800020012802044f0d00024020024104710d00200042003702000c010b1080808080000b024002402002411071450d00200010a78080800020012802044d0d0020024104710d01200042003702000b20000f0b10808080800020000b3400024002402000280204450d0020002802002c0000417f4c0d0141010f0b41000f0b200010a880808000200010a9808080006a0b280002402000280204450d0020002802002c0000417f4c0d0041000f0b200010ae8080800041016a0bc90301047f0240024002402000280204450d00200010af808080004101210120002802002c00002202417f4c0d010c020b41000f0b0240200241ff0171220141b7014b0d00200141807f6a0f0b024002400240200241ff0171220241bf014b0d000240200041046a22032802002202200141c97e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b200141384f0d010c020b0240200241f7014b0d00200141c07e6a0f0b0240200041046a22032802002202200141897e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b20014138490d010b200141ff7d490d010b10808080800020010f0b20010b5102017f017e23808080800041306b220124808080800020012000290200220237031020012002370308200141186a200141086a411410a68080800010a7808080002100200141306a24808080800020000b6a01037f02400240024020012802002204450d0041002105200320026a200128020422064b0d0120062002490d014100210120062003490d02200620026b20032003417f461b2101200420026a21050c020b410021050b410021010b20002001360204200020053602000b3301017f0240200110a980808000220220012802044d0d001080808080000b20002001200110a880808000200210ab808080000bd003020a7f017e23808080800041c0006b220324808080800002402001280208220420024d0d00200341386a200110ac80808000200320032903383703182001200341186a10aa8080800036020c200341306a200110ac80808000410021044100210541002106024020032802302207450d00410021054100210620032802342208200128020c2209490d00200820092009417f461b2105200721060b20012006360210200141146a2005360200200141086a41003602000b200141106a2106200141146a21092001410c6a2107200141086a210802400340200420024f0d012009280200450d01200341306a200110ac8080800041002104024002402003280230220a450d00410021052003280234220b2007280200220c490d01200a200c6a2105200b200c6b21040c010b410021050b20092004360200200620053602002003200436022c2003200536022820032003290328370310200341306a20064100200341106a10aa8080800010ab8080800020062003290330220d37020020072007280200200d422088a76a3602002008200828020041016a22043602000c000b0b20032006290200220d3703202003200d3703082000200341086a411410a6808080001a200341c0006a2480808080000b4701017f4100210102402000280204450d00024020002802002d0000220041bf014b0d00200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b5401017f024020002802040d001080808080000b0240200028020022012d0000418101470d000240200041046a28020041014b0d00108080808000200028020021010b20012c00014100480d001080808080000b0bbc0101047f024002402000280204450d00200010af80808000200028020022012c000022024100480d0120024100470f0b41000f0b410121030240200241807f460d000240200241ff0171220441b7014b0d000240200041046a28020041014b0d00108080808000200028020021010b20012d00014100470f0b41002103200441bf014b0d000240200041046a280200200241ff017141ca7e6a22024b0d00108080808000200028020021010b200120026a2d000041004721030b20030b2701017f200020012802002203200320012802046a10b2808080002000200210b38080800020000b34002000200220016b220210b480808000200028020020002802046a200120021089808080001a2000200028020420026a3602040bb60201087f02402001450d002000410c6a2102200041106a2103200041046a21040340200328020022052002280200460d010240200541786a28020020014f0d00108080808000200328020021050b200541786a2206200628020020016b220136020020010d01200320063602002000410120042802002005417c6a28020022016b220510b580808000220741016a20054138491b2206200428020022086a10b680808000200120002802006a220920066a2009200820016b1093808080001a02400240200541374b0d00200028020020016a200541406a3a00000c010b0240200741f7016a220641ff014b0d00200028020020016a20063a00002000280200200720016a6a210103402005450d02200120053a0000200541087621052001417f6a21010c000b0b1080808080000b410121010c000b0b0b21000240200028020420016a220120002802084d0d002000200110b7808080000b0b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b13002000200110b780808000200020013602040b4501017f0240200028020820014f0d002001108c808080002202200028020020002802041089808080001a200010c180808000200041086a2001360200200020023602000b0b29002000410110b480808000200028020020002802046a20013a00002000200028020441016a3602040b3c01017f0240200110b580808000220320026a2202418002480d001080808080000b2000200241ff017110b88080800020002001200310ba808080000b44002000200028020420026a10b680808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0bf90101037f23808080800041106b22032480808080002001280200210420012802042105024002402002450d004100210102400340200420016a2102200120054f0d0120022d00000d01200141016a21010c000b0b200520016b21050c010b200421020b0240024002400240024020054101470d0020022c000022014100480d012000200141ff017110b8808080000c040b200541374b0d010b200020054180017341ff017110b8808080000c010b2000200541b70110b9808080000b2003200536020c200320023602082003200329030837030020002003410010b1808080001a0b2000410110b380808000200341106a24808080800020000bc40101037f02400240024020012002844200510d00200142ff005620024200522002501b0d0120002001a741ff017110b8808080000c020b200041800110b8808080000c010b024002402001200210bd80808000220341374b0d00200020034180017341ff017110b8808080000c010b0240200310be80808000220441b7016a2205418002490d001080808080000b2000200541ff017110b88080800020002003200410bf808080000b200020012002200310c0808080000b2000410110b38080800020000b3501017f41002102024003402000200184500d0120004208882001423886842100200241016a2102200142088821010c000b0b20020b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b44002000200028020420026a10b680808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0b54002000200028020420036a10b680808000200028020020002802046a417f6a2100024003402001200284500d01200020013c0000200142088820024238868421012000417f6a2100200242088821020c000b0b0b1700024020002802002200450d002000108f808080000b0b240041a48a80800010a1808080001a418180808000410041808880800010a4808080001a0b1d0020004200370200200041086a4100360200200010c48080800020000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b0f0041b08a808000109a808080001a0b5701027f23808080800041106b22022480808080002002200110c78080800002402002280204200228020022016b2203450d002000200120031093808080001a0b200210c8808080001a200241106a24808080800020000bb80401077f23808080800041306b22022480808080004100210302402001280204220420012d000022054101762206200541017122051b22074102490d00410021032001280208200141016a20051b22082d00004130470d0020082d000141f8004641017421030b20024100360210200242003703080240200741016a20036b4101762207450d00200241286a200241106a3602002002200710918080800022053602082002200536020c20024200370320200242003703182002200520076a360210200241186a10cb808080001a20012d00002205410176210620054101712105200141046a28020021040b0240024002402004200620051b410171450d002001280208200141016a20051b20036a2c000010cc808080002205417f460d01200220053a001820034101722103200241086a200241186a10cd808080000b200141016a2107200141046a2106200141086a21080240024003402003200628020020012d00002205410176200541017122051b4f0d012008280200200720051b20036a22042c000010cc808080002105200441016a2c000010cc8080800021042005417f460d022004417f460d022002200420054104746a3a0018200341026a2103200241086a200241186a10cd808080000c000b0b200020022802083602002000200229020c3702042002420037020c200241086a21030c020b20004200370200200041086a21030c010b20004200370200200041086a21030b20034100360200200241086a10c8808080001a200241306a2480808080000b2201017f024020002802002201450d002000200136020420011092808080000b20000b6201017f23808080800041306b220424808080800020042003370328200441186a200210c7808080002004200110c680808000200441186a200441286a10ca808080002101200441186a10c8808080001a200441306a2480808080002001410173ad0bc20102027f027e23808080800041106b220324808080800041002104200229030022052106024003402006500d0120064208882106200441016a21040c000b0b20034100360208200342003703002003200410ce808080002003280204417f6a2104024003402005500d01200420053c00002004417f6a2104200542088821050c000b0b200020012802002204200128020420046b20032802002204200328020420046b1081808080002104200310c8808080001a200341106a2480808080002004450b4b01037f2000280208210120002802042102200041086a21030240034020022001460d0120032001417f6a22013602000c000b0b024020002802002201450d0020011092808080000b20000b4b01017f4150210102400240200041506a41ff0171410a490d0041a97f21012000419f7f6a41ff017141064f0d010b200120006a0f0b200041496a417f200041bf7f6a41ff01714106491b0bef0101047f23808080800041206b2202248080808000024002402000280204220320002802084f0d00200320012d00003a0000200041046a2200200028020041016a3602000c010b2000200341016a20002802006b10d1808080002104200241186a200041086a3602004100210320024100360214200041046a28020020002802006b210502402004450d00200410918080800021030b20022003360208200241146a200320046a360200200320056a220320012d00003a00002002200336020c2002200341016a3602102000200241086a10d280808000200241086a10cb808080001a0b200241206a2480808080000bd00201057f23808080800041206b220224808080800002400240024020002802042203200028020022046b220520014f0d00200028020820036b200120056b4f0d012000200110d1808080002104200241186a200041086a36020020024100360214200041046a28020020002802006b21064100210302402004450d00200410918080800021030b20022003360208200241146a200320046a3602002002200320066a22033602102002200336020c200520016b2101200241086a41086a21050340200341003a00002005200528020041016a2203360200200141016a22010d000b2000200241086a10d280808000200241086a10cb808080001a0c020b200520014d0d01200041046a200420016a3602000c010b200520016b2101200041046a21050340200341003a00002005200528020041016a2203360200200141016a22010d000b0b200241206a2480808080000bc40301057f23808080800041c0006b220524808080800020052004370338200541286a200310c7808080000240024002400240200541106a200210c680808000200541286a200541386a10ca80808000450d002005410036020820054200370300200510828080800010ce8080800020052802001083808080002005280204210620052802002102200541106a10d0808080002107200041086a4100360200200042003702002000200728020420052d0010220341017620034101711b2203200620026b4101746a4130109f8080800020052d001022084101710d0120084101762108200741016a21090c020b200010d0808080001a0c020b200741046a2802002108200728020821090b20004100200320092008109e808080001a200041016a2108200041086a21090240034020062002460d012009280200200820002d00004101711b20036a20022d000041047641f38a8080006a2d00003a00002009280200200820002d00004101711b20036a41016a20022d0000410f7141f38a8080006a2d00003a0000200241016a2102200341026a21030c000b0b2007109a808080001a200510c8808080001a0b200541286a10c8808080001a200541c0006a2480808080000b250020004200370200200041086a4100360200200041bc8a808000410010988080800020000b4c01017f02402001417f4c0d0041ffffffff0721020240200028020820002802006b220041feffffff034b0d0020012000410174220020002001491b21020b20020f0b200010a280808000000b9a0101037f200120012802042000280204200028020022026b22036b22043602040240200341004c0d002004200220031089808080001a200141046a28020021040b2000280200210320002004360200200141046a22042003360200200041046a220328020021022003200128020836020020012002360208200028020821032000200128020c3602082001200336020c200120042802003602000ba20603037f017e057f23808080800041c0016b22002480808080001087808080001084808080002201108c808080002202108580808000200020013602342000200236023020002000290330370308200041306a200041106a200041086a411c10a6808080002201410010ad8080800002400240200041306a10d48080800022034200510d00200341bd8a80800010d580808000510d010240200341c28a80800010d580808000520d00200041306a10d68080800021042001200041306a10d78080800020004180016a200041306a1097808080002101200041f0006a2000413c6a109780808000210220002903482103200041e0006a20004190016a200110978080800022052002200310c98080800021032005109a808080001a200041b0016a22054200370300200041a8016a4200370300200042003703a001200041a0016a2003420010bc808080001a024020002802ac012005280200460d001080808080000b20002802a00120002802a401108680808000200041a0016a10d8808080001a2002109a808080001a2001109a808080001a200410d9808080001a0c020b200341da8a80800010d580808000520d00200041306a10d68080800021052001200041306a10d780808000200041e0006a200041306a1097808080002102200041d0006a2000413c6a10978080800021042000290348210320004180016a200041286a200041f0006a200210978080800022012004200310cf808080002001109a808080001a200041b0016a22064200370300200041a8016a4200370300200042003703a001200020004190016a20004180016a1097808080002201280208200141016a20002d009001220741017122081b3602b80120002001280204200741017620081b3602bc01200020002903b801370300200041a0016a2000410010bb808080001a2001109a808080001a024020002802ac012006280200460d001080808080000b20002802a00120002802a401108680808000200041a0016a10d8808080001a20004180016a109a808080001a2004109a808080001a2002109a808080001a200510d9808080001a0c010b1080808080000b200041c0016a2480808080000bb00103017f017e017f23808080800041106b2201248080808000200010af8080800002400240200010b080808000450d002000280204450d0020002802002d000041c001490d010b1080808080000b200141086a200010da808080000240200128020c22004109490d001080808080000b4200210220012802082103024003402000450d012000417f6a210020024208862003310000842102200341016a21030c000b0b200141106a24808080800020020b3a01027e42a5c688a1c89ca7f94b21010240034020003000002202500d01200041016a2100200142b383808080207e20028521010c000b0b20010b2000200010c3808080001a2000410c6a10c3808080001a2000420037031820000b7701017f23808080800041306b220224808080800020024101360214200220003602082002200241146a36020c200241086a200110db80808000200241086a2001410c6a10db80808000200241186a2000200228021410ad808080002001200241186a10d480808000370318200241306a2480808080000b3001017f0240200028020c2201450d00200041106a200136020020011092808080000b2000280200108f8080800020000b19002000410c6a109a808080001a2000109a808080001a20000b870201057f0240200110a9808080002202200128020422034d0d00108080808000200141046a28020021030b2001280200210402400240024002400240024002402003450d004100210120042c00002205417f4c0d012004450d030c040b410021010c010b0240200541ff0171220641bf014b0d0041002101200541ff017141b801490d01200641c97e6a21010c010b41002101200541ff017141f801490d00200641897e6a21010b200141016a210120040d010b410021050c010b41002105200120026a20034b0d0020032001490d004100210620032002490d01200320016b20022002417f461b2106200420016a21050c010b410021060b20002006360204200020053602000bf70301067f23808080800041306b220224808080800020022000280200200028020428020010ad80808000024002400240024002402002280204450d0020022802002d000041c0014f0d00200241286a200210da80808000200210a9808080002103024020022802282204450d00200228022c220520034f0d020b41002104200241206a41003602002002420037031841002106410021050c020b200241186a10c3808080001a0c030b200241206a4100360200200242003703180240200520032003417f461b220641704f0d00200420066a21052006410a4d0d01200641106a417071220710918080800021032002200636021c20022007410172360218200220033602200c020b200241186a109680808000000b200220064101743a0018200241186a41017221030b0240034020052004460d01200320042d00003a0000200341016a2103200441016a21040c000b0b200341003a00000b0240024020012d00004101710d00200141003b01000c010b200128020841003a00002001410036020420012d0000410171450d00200141086a280200109280808000200141003602000b20012002290318370200200141086a200241186a41086a280200360200200241186a10c480808000200241186a109a808080001a200041046a2802002204200428020041016a360200200241306a2480808080000b240041b08a80800010c3808080001a418280808000410041808880800010a4808080001a0b0b920302004180080bbc02000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041bc0a0b4800696e69740064656c65676174655f63616c6c5f70706f735f73656e640064656c65676174655f63616c6c5f70706f735f7175657279003031323334353637383961626364656600";

    public static String BINARY = BINARY_0;

    public static final String FUNC_DELEGATE_CALL_PPOS_QUERY = "delegate_call_ppos_query";

    public static final String FUNC_DELEGATE_CALL_PPOS_SEND = "delegate_call_ppos_send";

    protected ContractDelegateCallPPOS(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ContractDelegateCallPPOS(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<String> delegate_call_ppos_query(String target_addr, String in, Uint64 gas) {
        final WasmFunction function = new WasmFunction(FUNC_DELEGATE_CALL_PPOS_QUERY, Arrays.asList(target_addr,in,gas), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static RemoteCall<ContractDelegateCallPPOS> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDelegateCallPPOS.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ContractDelegateCallPPOS> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDelegateCallPPOS.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ContractDelegateCallPPOS> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDelegateCallPPOS.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<ContractDelegateCallPPOS> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractDelegateCallPPOS.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<TransactionReceipt> delegate_call_ppos_send(String target_addr, String in, Uint64 gas) {
        final WasmFunction function = new WasmFunction(FUNC_DELEGATE_CALL_PPOS_SEND, Arrays.asList(target_addr,in,gas), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> delegate_call_ppos_send(String target_addr, String in, Uint64 gas, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_DELEGATE_CALL_PPOS_SEND, Arrays.asList(target_addr,in,gas), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static ContractDelegateCallPPOS load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new ContractDelegateCallPPOS(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static ContractDelegateCallPPOS load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new ContractDelegateCallPPOS(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
