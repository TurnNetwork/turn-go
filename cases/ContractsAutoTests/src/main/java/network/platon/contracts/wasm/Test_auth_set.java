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
public class Test_auth_set extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001560f60000060017f0060027f7f0060027f7f017f60047f7f7f7f017f6000017f60047f7f7f7f0060037f7f7f017f60017f017f60037f7f7f0060037f7e7e017f60027e7e017f60047f7e7e7f0060027f7e0060017f017e02bd010803656e760c706c61746f6e5f70616e6963000003656e7617706c61746f6e5f6765745f73746174655f6c656e677468000303656e7610706c61746f6e5f6765745f7374617465000403656e760d706c61746f6e5f6f726967696e000103656e760d706c61746f6e5f63616c6c6572000103656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000503656e7610706c61746f6e5f6765745f696e707574000103656e7610706c61746f6e5f7365745f737461746500060357560000070708080703010008010701010103090708010801000701070808080806020908010807090202080202020909070a0b08090c01000801010803080802080808080d030702080101020102000e020202010802000405017001050505030100020615037f0141e08a040b7f0041e08a040b7f0041d90a0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300080b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974001f06696e766f6b650055090a010041010b04215254410ae55f5608001011103e105d0b02000bc60a010b7f2002410f6a210341002104410020026b21052002410e6a2106410120026b21072002410d6a2108410220026b210902400340200020046a210b200120046a210a20022004460d01200a410371450d01200b200a2d00003a00002003417f6a2103200541016a21052006417f6a2106200741016a21072008417f6a2108200941016a2109200441016a21040c000b0b200220046b210c02400240024002400240200b410371220d450d00200c4120490d03200d4101460d01200d4102460d02200d4103470d03200b200120046a28020022063a0000200041016a210c200220046b417f6a21092004210b0240034020094113490d01200c200b6a220a2001200b6a220741046a2802002208411874200641087672360200200a41046a200741086a2802002206411874200841087672360200200a41086a2007410c6a2802002208411874200641087672360200200a410c6a200741106a2802002206411874200841087672360200200b41106a210b200941706a21090c000b0b2002417f6a2005416d2005416d4b1b20036a4170716b20046b210c2001200b6a41016a210a2000200b6a41016a210b0c030b200c210a02400340200a4110490d01200020046a220b200120046a2207290200370200200b41086a200741086a290200370200200441106a2104200a41706a210a0c000b0b02400240200c4108710d00200120046a210a200020046a21040c010b200020046a220b200120046a2204290200370200200441086a210a200b41086a21040b0240200c410471450d002004200a280200360200200a41046a210a200441046a21040b0240200c410271450d002004200a2f00003b0000200441026a2104200a41026a210a0b200c410171450d032004200a2d00003a000020000f0b200b200120046a220a28020022063a0000200b41016a200a41016a2f00003b0000200041036a210c200220046b417d6a21052004210b0240034020054111490d01200c200b6a220a2001200b6a220741046a2802002203410874200641187672360200200a41046a200741086a2802002206410874200341187672360200200a41086a2007410c6a2802002203410874200641187672360200200a410c6a200741106a2802002206410874200341187672360200200b41106a210b200541706a21050c000b0b2002417d6a2009416f2009416f4b1b20086a4170716b20046b210c2001200b6a41036a210a2000200b6a41036a210b0c010b200b200120046a220a28020022083a0000200b41016a200a41016a2d00003a0000200041026a210c200220046b417e6a21052004210b0240034020054112490d01200c200b6a220a2001200b6a220941046a2802002203411074200841107672360200200a41046a200941086a2802002208411074200341107672360200200a41086a2009410c6a2802002203411074200841107672360200200a410c6a200941106a2802002208411074200341107672360200200b41106a210b200541706a21050c000b0b2002417e6a2007416e2007416e4b1b20066a4170716b20046b210c2001200b6a41026a210a2000200b6a41026a210b0b0240200c411071450d00200b200a2d00003a0000200b200a280001360001200b200a290005370005200b200a2f000d3b000d200b200a2d000f3a000f200b41106a210b200a41106a210a0b0240200c410871450d00200b200a290000370000200b41086a210b200a41086a210a0b0240200c410471450d00200b200a280000360000200b41046a210b200a41046a210a0b0240200c410271450d00200b200a2f00003b0000200b41026a210b200a41026a210a0b200c410171450d00200b200a2d00003a00000b20000bfb0202027f017e02402002450d00200020013a0000200020026a2203417f6a20013a000020024103490d00200020013a0002200020013a00012003417d6a20013a00002003417e6a20013a000020024107490d00200020013a00032003417c6a20013a000020024109490d002000410020006b41037122046a2203200141ff017141818284086c22013602002003200220046b417c7122046a2202417c6a200136020020044109490d002003200136020820032001360204200241786a2001360200200241746a200136020020044119490d002003200136021820032001360214200320013602102003200136020c200241706a20013602002002416c6a2001360200200241686a2001360200200241646a20013602002001ad220542208620058421052004200341047141187222016b2102200320016a2101034020024120490d0120012005370300200141186a2005370300200141106a2005370300200141086a2005370300200141206a2101200241606a21020c000b0b20000b3a01017f23808080800041106b220141e08a84800036020c2000200128020c41076a41787122013602042000200136020020003f0036020c20000b120041808880800020004108108e808080000bc10101067f23808080800041106b22032480808080002003200136020c024002402001450d002000200028020c200241036a410020026b220471220520016a220641107622076a220836020c200020022000280204220120066a6a417f6a20047122023602040240200841107420024b0d002000410c6a200841016a360200200741016a21070b0240200740000d001080808080000b20012003410c6a4104108a808080001a200120056a21000c010b410021000b200341106a24808080800020000b2e000240418088808000200120006c22004108108e808080002201450d00200141002000108b808080001a0b20010b02000b0f00418088808000108c808080001a0b3a01027f2000410120001b2101024003402001108d8080800022020d014100280290888080002200450d012000118080808000000c000b0b20020b0a0020001090808080000bce0301067f024020002001460d000240024002400240200120006b20026b410020024101746b4d0d0020012000734103712103200020014f0d012003450d02200021030c030b200020012002108a808080000f0b024020030d002001417f6a210402400340200020026a2203410371450d012002450d052003417f6a200420026a2d00003a00002002417f6a21020c000b0b2000417c6a21032001417c6a2104034020024104490d01200320026a200420026a2802003602002002417c6a21020c000b0b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b0b200241046a21052002417f7321064100210402400340200120046a2107200020046a2208410371450d0120022004460d03200820072d00003a00002005417f6a2105200641016a2106200441016a21040c000b0b200220046b2101410021030240034020014104490d01200820036a200720036a280200360200200341046a21032001417c6a21010c000b0b200720036a2101200820036a210320022006417c2006417c4b1b20056a417c716b20046b21020b03402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b0b20000b0900200041013602000b0900200041003602000b0900108980808000000b5201017f20004200370200200041086a22024100360200024020012d00004101710d00200020012902003702002002200141086a28020036020020000f0b20002001280208200128020410998080800020000b7701027f0240200241704f0d00024002402002410a4b0d00200020024101743a0000200041016a21030c010b200241106a417071220410928080800021032000200236020420002004410172360200200020033602080b200320012002109a808080001a200320026a41003a00000f0b108980808000000b1a0002402002450d00200020012002108a8080800021000b20000b1d00024020002d0000410171450d0020002802081093808080000b20000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b1d0020004200370200200041086a41003602002000109c8080800020000b0900108980808000000bb60101037f4194888080001095808080004100280298888080002100024003402000450d01024003404100410028029c888080002202417f6a220136029c8880800020024101480d01200020014102746a22004184016a2802002102200041046a2802002100419488808000109680808000200220001181808080000041948880800010958080800041002802988880800021000c000b0b4100412036029c88808000410020002802002200360298888080000c000b0b0bcd0101027f419488808000109580808000024041002802988880800022030d0041a0888080002103410041a088808000360298888080000b02400240410028029c8880800022044120470d004184024101108f808080002203450d0141002104200341002802988880800036020041002003360298888080004100410036029c888080000b4100200441016a36029c88808000200320044102746a22034184016a2001360200200341046a200036020041948880800010968080800041000f0b419488808000109680808000417f0b0f0041a48a808000109b808080001a0b89010020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d00200010a38080800020012802044f0d00024020024104710d00200042003702000c010b1080808080000b024002402002411071450d00200010a38080800020012802044d0d0020024104710d01200042003702000b20000f0b10808080800020000b3400024002402000280204450d0020002802002c0000417f4c0d0141010f0b41000f0b200010a480808000200010a5808080006a0b280002402000280204450d0020002802002c0000417f4c0d0041000f0b200010aa8080800041016a0bc90301047f0240024002402000280204450d00200010ab808080004101210120002802002c00002202417f4c0d010c020b41000f0b0240200241ff0171220141b7014b0d00200141807f6a0f0b024002400240200241ff0171220241bf014b0d000240200041046a22032802002202200141c97e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b200141384f0d010c020b0240200241f7014b0d00200141c07e6a0f0b0240200041046a22032802002202200141897e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b20014138490d010b200141ff7d490d010b10808080800020010f0b20010b5102017f017e23808080800041306b220124808080800020012000290200220237031020012002370308200141186a200141086a411410a28080800010a3808080002100200141306a24808080800020000b6a01037f02400240024020012802002204450d0041002105200320026a200128020422064b0d0120062002490d014100210120062003490d02200620026b20032003417f461b2101200420026a21050c020b410021050b410021010b20002001360204200020053602000b3301017f0240200110a580808000220220012802044d0d001080808080000b20002001200110a480808000200210a7808080000bd003020a7f017e23808080800041c0006b220324808080800002402001280208220420024d0d00200341386a200110a880808000200320032903383703182001200341186a10a68080800036020c200341306a200110a880808000410021044100210541002106024020032802302207450d00410021054100210620032802342208200128020c2209490d00200820092009417f461b2105200721060b20012006360210200141146a2005360200200141086a41003602000b200141106a2106200141146a21092001410c6a2107200141086a210802400340200420024f0d012009280200450d01200341306a200110a88080800041002104024002402003280230220a450d00410021052003280234220b2007280200220c490d01200a200c6a2105200b200c6b21040c010b410021050b20092004360200200620053602002003200436022c2003200536022820032003290328370310200341306a20064100200341106a10a68080800010a78080800020062003290330220d37020020072007280200200d422088a76a3602002008200828020041016a22043602000c000b0b20032006290200220d3703202003200d3703082000200341086a411410a2808080001a200341c0006a2480808080000b4701017f4100210102402000280204450d00024020002802002d0000220041bf014b0d00200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b5401017f024020002802040d001080808080000b0240200028020022012d0000418101470d000240200041046a28020041014b0d00108080808000200028020021010b20012c00014100480d001080808080000b0bbc0101047f024002402000280204450d00200010ab80808000200028020022012c000022024100480d0120024100470f0b41000f0b410121030240200241807f460d000240200241ff0171220441b7014b0d000240200041046a28020041014b0d00108080808000200028020021010b20012d00014100470f0b41002103200441bf014b0d000240200041046a280200200241ff017141ca7e6a22024b0d00108080808000200028020021010b200120026a2d000041004721030b20030b2701017f200020012802002203200320012802046a10ae808080002000200210af8080800020000b34002000200220016b220210b080808000200028020020002802046a20012002108a808080001a2000200028020420026a3602040bb60201087f02402001450d002000410c6a2102200041106a2103200041046a21040340200328020022052002280200460d010240200541786a28020020014f0d00108080808000200328020021050b200541786a2206200628020020016b220136020020010d01200320063602002000410120042802002005417c6a28020022016b220510b180808000220741016a20054138491b2206200428020022086a10b280808000200120002802006a220920066a2009200820016b1094808080001a02400240200541374b0d00200028020020016a200541406a3a00000c010b0240200741f7016a220641ff014b0d00200028020020016a20063a00002000280200200720016a6a210103402005450d02200120053a0000200541087621052001417f6a21010c000b0b1080808080000b410121010c000b0b0b21000240200028020420016a220120002802084d0d002000200110b3808080000b0b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b13002000200110b380808000200020013602040b4501017f0240200028020820014f0d002001108d80808000220220002802002000280204108a808080001a200010bd80808000200041086a2001360200200020023602000b0b29002000410110b080808000200028020020002802046a20013a00002000200028020441016a3602040b3c01017f0240200110b180808000220320026a2202418002480d001080808080000b2000200241ff017110b48080800020002001200310b6808080000b44002000200028020420026a10b280808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0bf90101037f23808080800041106b22032480808080002001280200210420012802042105024002402002450d004100210102400340200420016a2102200120054f0d0120022d00000d01200141016a21010c000b0b200520016b21050c010b200421020b0240024002400240024020054101470d0020022c000022014100480d012000200141ff017110b4808080000c040b200541374b0d010b200020054180017341ff017110b4808080000c010b2000200541b70110b5808080000b2003200536020c200320023602082003200329030837030020002003410010ad808080001a0b2000410110af80808000200341106a24808080800020000bc40101037f02400240024020012002844200510d00200142ff005620024200522002501b0d0120002001a741ff017110b4808080000c020b200041800110b4808080000c010b024002402001200210b980808000220341374b0d00200020034180017341ff017110b4808080000c010b0240200310ba80808000220441b7016a2205418002490d001080808080000b2000200541ff017110b48080800020002003200410bb808080000b200020012002200310bc808080000b2000410110af8080800020000b3501017f41002102024003402000200184500d0120004208882001423886842100200241016a2102200142088821010c000b0b20020b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b44002000200028020420026a10b280808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0b54002000200028020420036a10b280808000200028020020002802046a417f6a2100024003402001200284500d01200020013c0000200142088820024238868421012000417f6a2100200242088821020c000b0b0b1700024020002802002200450d0020001090808080000b0b240041a48a808000109d808080001a418180808000410041808880800010a0808080001a0b1d0020004200370200200041086a4100360200200010c08080800020000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b0f0041b08a808000109b808080001a0b2a01017f410021010240034020014114460d01200020016a41003a0000200141016a21010c000b0b20000bfc0401077f23808080800041306b22022480808080004100210302402001280204220420012d000022054101762206200541017122051b22074102490d00410021032001280208200141016a20051b22082d00004130470d0020082d000141f8004641017421030b20024100360210200242003703080240200741016a20036b4101762207450d00200241286a200241106a3602002002200710928080800022053602082002200536020c20024200370320200242003703182002200520076a360210200241186a10c4808080001a20012d00002205410176210620054101712105200141046a28020021040b02400240024002402004200620051b410171450d002001280208200141016a20051b20036a2c000010c5808080002205417f460d01200220053a001820034101722103200241086a200241186a10c6808080000b200141016a2107200141046a2106200141086a21080240024003402003200628020020012d00002205410176200541017122051b4f0d012008280200200720051b20036a22042c000010c5808080002105200441016a2c000010c58080800021042005417f460d022004417f460d022002200420054104746a3a0018200341026a2103200241086a200241186a10c6808080000c000b0b2002200228020822033602182002200228020c220536021c2002200241106a2802003602202002420037020c200241086a21040c030b20024200370318200241206a21040c010b20024200370318200241206a21040b41002105410021030b20044100360200200241086a10c7808080001a0240200520036b2205450d002000200320051094808080001a0b200241186a10c7808080001a200241306a24808080800020000b4b01037f2000280208210120002802042102200041086a21030240034020022001460d0120032001417f6a22013602000c000b0b024020002802002201450d0020011093808080000b20000b4b01017f4150210102400240200041506a41ff0171410a490d0041a97f21012000419f7f6a41ff017141064f0d010b200120006a0f0b200041496a417f200041bf7f6a41ff01714106491b0bef0101047f23808080800041206b2202248080808000024002402000280204220320002802084f0d00200320012d00003a0000200041046a2200200028020041016a3602000c010b2000200341016a20002802006b10cc808080002104200241186a200041086a3602004100210320024100360214200041046a28020020002802006b210502402004450d00200410928080800021030b20022003360208200241146a200320046a360200200320056a220320012d00003a00002002200336020c2002200341016a3602102000200241086a10d880808000200241086a10c4808080001a0b200241206a2480808080000b2201017f024020002802002201450d002000200136020420011093808080000b20000b2901017f0240024020002d000022014101710d00200141017621000c010b200028020421000b2000450b9006010b7f23808080800041e0006b2201248080808000200010c2808080001a200042befadbd3c388a6d88a7f370318200041206a10c2808080002102200141286a10ca808080002203200029031810cb808080000240200328020c200341106a280200460d001080808080000b0240024020032802002204200328020422051081808080002206450d002001410036022020014200370318200141186a200610cc808080002107200141106a200141186a41086a3602002001410036020c200128021c20012802186b21084100210902402007450d00200710928080800021090b200120093602002001410c6a200920076a3602002001200920086a220936020820012009360204200141086a21070340200941003a00002007200728020041016a22093602002006417f6a22060d000b20012001280204200128021c2208200128021822066b220a6b22073602040240200a41004c0d0020072006200a108a808080001a200141086a2802002109200128021c210820012802182106200128020421070b200141086a2008360200200141186a41086a2208280200210a20082001410c6a220b280200360200200b200a36020020012006360204200120073602182001200936021c20012006360200200110c4808080001a2004200520012802182209200128021c20096b1082808080001a20012001280218220941016a200128021c2009417f736a10cd80808000220610ab80808000200141d8006a200610ce80808000200128025c2109024002402006280204450d00200941144b0d0020062802002d000041ff017141c001490d010b1080808080000b200141c0006a10c2808080002009411420094114491b22096b41146a20012802582009108a808080001a200241106a200141c0006a41106a280200360000200241086a200141c0006a41086a29030037000020022001290340370000200141186a10c7808080001a200310cf808080001a0c010b200310cf808080001a200241106a200041106a280000360000200241086a200041086a290000370000200220002900003700000b200141e0006a24808080800020000b2d0020004100360208200042003702002000410010d980808000200041146a41003602002000420037020c20000b0f0020002001420010b8808080001a0b4c01017f02402001417f4c0d0041ffffffff0721020240200028020820002802006b220041feffffff034b0d0020012000410174220020002001491b21020b20020f0b2000109e80808000000b4801017f23808080800041106b22032480808080002003200236020c200320013602082003200329030837030020002003411c10a2808080002100200341106a24808080800020000b870201057f0240200110a5808080002202200128020422034d0d00108080808000200141046a28020021030b2001280200210402400240024002400240024002402003450d004100210120042c00002205417f4c0d012004450d030c040b410021010c010b0240200541ff0171220641bf014b0d0041002101200541ff017141b801490d01200641c97e6a21010c010b41002101200541ff017141f801490d00200641897e6a21010b200141016a210120040d010b410021050c010b41002105200120026a20034b0d0020032001490d004100210620032002490d01200320016b20022002417f461b2106200420016a21050c010b410021060b20002006360204200020053602000b2d01017f0240200028020c2201450d00200041106a200136020020011093808080000b200010da8080800020000b8d0101017f23808080800041d0006b2201248080808000200141386a200010c3808080001a0240200010c880808000450d00200110c2808080001a2001108380808000200141386a41106a200141106a280200360200200141386a41086a200141086a290300370300200120012903003703380b200110c98080800010d180808000200141d0006a2480808080000b820705037f017e017f017e067f2380808080004180016b2201248080808000200141286a10ca80808000210241002103200141c8006a41186a410036020042002104200141c8006a41106a4200370300200141d0006a42003703002001420037034841012105024020002903182206428001540d00024003402006200484500d0120064208882004423886842106200341016a2103200442088821040c000b0b024020034138490d002003210503402005450d01200341016a2103200541087621050c000b0b200341016a21050b20012005360248200141c8006a41047210db808080001a2002200510dc808080002002200041186a29030010cb808080000240200228020c200241106a280200460d001080808080000b2002280204210720022802002108200141106a10ca80808000210941002103200141e0006a4100360200200141c8006a41106a4200370300200141c8006a41086a420037030020014200370348200141e8006a41106a200041306a280000360200200141e8006a41086a200041286a29000037030020012000290020370368200041206a21004101210a0240034020034114460d01200141e8006a20036a2105200341016a210320052d0000450d000b4115210a0b2001200a360248200141c8006a41047210db808080001a4101109280808000220341fe013a0000200120033602682001200341016a22053602702001200536026c0240200928020c200941106a280200460d00108080808000200128026c2105200128026821030b0240200520036b22052009280204220b6a220c20092802084d0d002009200c10d980808000200941046a280200210b0b2009280200200b6a20032005108a808080001a200941046a2203200328020020056a3602002009200128026c200a6a20012802686b10dc80808000200141c8006a41106a200041106a280000360200200141c8006a41086a200041086a29000037030020014114360244200120002900003703482001200141c8006a360240200120012903403703082009200141086a410010b7808080001a024002402009410c6a2205280200200941106a2200280200460d001080808080002009280200210320052802002000280200460d011080808080000c010b200928020021030b200820072003200941046a280200108780808000200141e8006a10c7808080001a200910cf808080001a200210cf808080001a20014180016a2480808080000bbb0101017f23808080800041d0006b2202248080808000200241386a200110c3808080001a0240200110c880808000450d00200210c2808080001a2002108480808000200241386a41106a200241106a280200360200200241386a41086a200241086a290300370300200220022903003703380b200210c980808000220141306a200241c8006a280200360000200141286a200241c0006a29030037000020012002290338370020200110d180808000200241d0006a2480808080000b3901027f23808080800041106b2201248080808000200110bf80808000220210d0808080002002109b808080001a200141106a2480808080000b0a00200110d0808080000ba20302047f017e23808080800041e0006b22002480808080001088808080001085808080002201108d808080002202108680808000200041286a200041c0006a2002200110cd808080002203410010a980808000200041286a10ab8080800002400240200041286a10ac80808000450d00200028022c450d0020002802282d000041c001490d010b1080808080000b200041d8006a200041286a10ce808080000240200028025c22014109490d001080808080000b4200210420002802582102024003402001450d012001417f6a210120044208862002310000842104200241016a21020c000b0b0240024020044200510d000240200441bc8a80800010d680808000520d00200041003602242000418280808000360220200020002903203703082003200041086a10d7808080000c020b0240200441c18a80800010d680808000520d00200041286a10d3808080000c020b200441cc8a80800010d680808000520d002000410036021c2000418380808000360218200020002903183703102003200041106a10d7808080000c010b1080808080000b200041e0006a2480808080000b3a01027e42a5c688a1c89ca7f94b21010240034020003000002202500d01200041016a2100200142b383808080207e20028521010c000b0b20010bcf0401077f23808080800041c0006b22022480808080002001280204210320012802002104200210bf808080002105200241106a2000410110a980808000024002400240024002402002280214450d0020022802102d000041c0014f0d00200241386a200241106a10ce80808000200241106a10a5808080002100024020022802382201450d00200228023c220620004f0d020b41002101200241306a41003602002002420037032841002107410021060c020b200241286a10bf808080001a0c030b200241306a4100360200200242003703280240200620002000417f461b220741704f0d00200120076a21062007410a4d0d01200741106a417071220810928080800021002002200736022c20022008410172360228200220003602300c020b200241286a109780808000000b200220074101743a0028200241286a41017221000b0240034020062001460d01200020012d00003a0000200041016a2100200141016a21010c000b0b200041003a00000b0240024020022d00004101710d00200241003b01000c010b200228020841003a00002002410036020420022d0000410171450d00200241086a280200109380808000200241003602000b200241086a200241286a41086a28020036020020022002290328370300200241286a20034101756a2100200241286a10c080808000200241286a109b808080001a200241106a2005109880808000210102402003410171450d00200028020020046a28020021040b200020012004118280808000002001109b808080001a2005109b808080001a200241c0006a2480808080000b9a0101037f200120012802042000280204200028020022026b22036b22043602040240200341004c0d00200420022003108a808080001a200141046a28020021040b2000280200210320002004360200200141046a22042003360200200041046a220328020021022003200128020836020020012002360208200028020821032000200128020c3602082001200336020c200120042802003602000b4401017f0240200028020820014f0d002001108d8080800020002802002000280204108a808080002102200010da80808000200041086a2001360200200020023602000b0b0d0020002802001090808080000b8f0301067f200028020422012000280210220241087641fcffff07716a210302400240200028020822042001460d002001200028021420026a220541087641fcffff07716a280200200541ff07714102746a2105200041146a21062003280200200241ff07714102746a21020c010b200041146a210641002102410021050b0240034020052002460d01200241046a220220032802006b418020470d0020032802042102200341046a21030c000b0b20064100360200200041086a210302400340200420016b41027522024103490d012001280200109380808000200041046a2201200128020041046a2201360200200328020021040c000b0b02400240024020024101460d0020024102470d0241800821020c010b41800421020b200041106a20023602000b0240034020042001460d012001280200109380808000200141046a21010c000b0b200041086a22022802002101200041046a28020021040240034020042001460d0120022001417c6a22013602000c000b0b024020002802002201450d0020011093808080000b20000b19000240200028020820014f0d002000200110d9808080000b0b240041b08a80800010bf808080001a418480808000410041808880800010a0808080001a0b0be70202004180080bbc02000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041bc0a0b1d696e697400746573745f6f776e657200746573745f6f776e65725f7000";

    public static String BINARY = BINARY_0;

    public static final String FUNC_TEST_OWNER_P = "test_owner_p";

    public static final String FUNC_TEST_OWNER = "test_owner";

    protected Test_auth_set(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected Test_auth_set(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<TransactionReceipt> test_owner_p(String addr) {
        final WasmFunction function = new WasmFunction(FUNC_TEST_OWNER_P, Arrays.asList(addr), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> test_owner_p(String addr, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_TEST_OWNER_P, Arrays.asList(addr), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static RemoteCall<Test_auth_set> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, String addr) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList(addr));
        return deployRemoteCall(Test_auth_set.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<Test_auth_set> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, String addr) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList(addr));
        return deployRemoteCall(Test_auth_set.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<Test_auth_set> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue, String addr) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList(addr));
        return deployRemoteCall(Test_auth_set.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<Test_auth_set> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue, String addr) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList(addr));
        return deployRemoteCall(Test_auth_set.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<TransactionReceipt> test_owner() {
        final WasmFunction function = new WasmFunction(FUNC_TEST_OWNER, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> test_owner(BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_TEST_OWNER, Arrays.asList(), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static Test_auth_set load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new Test_auth_set(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static Test_auth_set load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new Test_auth_set(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
