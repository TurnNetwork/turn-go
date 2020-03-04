package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Int16;
import com.platon.rlp.datatypes.Int32;
import com.platon.rlp.datatypes.Int64;
import com.platon.rlp.datatypes.Uint32;
import com.platon.rlp.datatypes.Uint64;
import com.platon.rlp.datatypes.Uint8;
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
 * <p>Generated with platon-web3j version 0.9.1.1-SNAPSHOT.
 */
public class IntegerDataTypeContract_1 extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001701260000060017f0060037f7f7e006000017f60027f7f0060037f7f7f017f60017f017f60027f7f017f60037f7f7f0060077f7f7f7f7f7f7f0060067f7e7e7e7e7f0060057f7e7e7e7e0060047f7f7f7f0060037f7e7e017f60027e7e017f60047f7e7e7f0060037f7e7e0060017f017e025d0403656e760c706c61746f6e5f70616e6963000003656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000303656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e00040349480000050506060507010006010501010708050609040401060a0b0b00050105060606060c040806010605080404060404040808050d0e06080f0100100601020200111106040404000405017001050505030100020615037f0141808b040b7f0041808b040b7f0041fa0a0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300040b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974001f06696e766f6b650044090a010041010b04214243410a9b56480800100d103e104b0b02000bc60a010b7f2002410f6a210341002104410020026b21052002410e6a2106410120026b21072002410d6a2108410220026b210902400340200020046a210b200120046a210a20022004460d01200a410371450d01200b200a2d00003a00002003417f6a2103200541016a21052006417f6a2106200741016a21072008417f6a2108200941016a2109200441016a21040c000b0b200220046b210c02400240024002400240200b410371220d450d00200c4120490d03200d4101460d01200d4102460d02200d4103470d03200b200120046a28020022063a0000200041016a210c200220046b417f6a21092004210b0240034020094113490d01200c200b6a220a2001200b6a220741046a2802002208411874200641087672360200200a41046a200741086a2802002206411874200841087672360200200a41086a2007410c6a2802002208411874200641087672360200200a410c6a200741106a2802002206411874200841087672360200200b41106a210b200941706a21090c000b0b2002417f6a2005416d2005416d4b1b20036a4170716b20046b210c2001200b6a41016a210a2000200b6a41016a210b0c030b200c210a02400340200a4110490d01200020046a220b200120046a2207290200370200200b41086a200741086a290200370200200441106a2104200a41706a210a0c000b0b02400240200c4108710d00200120046a210a200020046a21040c010b200020046a220b200120046a2204290200370200200441086a210a200b41086a21040b0240200c410471450d002004200a280200360200200a41046a210a200441046a21040b0240200c410271450d002004200a2f00003b0000200441026a2104200a41026a210a0b200c410171450d032004200a2d00003a000020000f0b200b200120046a220a28020022063a0000200b41016a200a41016a2f00003b0000200041036a210c200220046b417d6a21052004210b0240034020054111490d01200c200b6a220a2001200b6a220741046a2802002203410874200641187672360200200a41046a200741086a2802002206410874200341187672360200200a41086a2007410c6a2802002203410874200641187672360200200a410c6a200741106a2802002206410874200341187672360200200b41106a210b200541706a21050c000b0b2002417d6a2009416f2009416f4b1b20086a4170716b20046b210c2001200b6a41036a210a2000200b6a41036a210b0c010b200b200120046a220a28020022083a0000200b41016a200a41016a2d00003a0000200041026a210c200220046b417e6a21052004210b0240034020054112490d01200c200b6a220a2001200b6a220941046a2802002203411074200841107672360200200a41046a200941086a2802002208411074200341107672360200200a41086a2009410c6a2802002203411074200841107672360200200a410c6a200941106a2802002208411074200341107672360200200b41106a210b200541706a21050c000b0b2002417e6a2007416e2007416e4b1b20066a4170716b20046b210c2001200b6a41026a210a2000200b6a41026a210b0b0240200c411071450d00200b200a2d00003a0000200b200a280001360001200b200a290005370005200b200a2f000d3b000d200b200a2d000f3a000f200b41106a210b200a41106a210a0b0240200c410871450d00200b200a290000370000200b41086a210b200a41086a210a0b0240200c410471450d00200b200a280000360000200b41046a210b200a41046a210a0b0240200c410271450d00200b200a2f00003b0000200b41026a210b200a41026a210a0b200c410171450d00200b200a2d00003a00000b20000bfb0202027f017e02402002450d00200020013a0000200020026a2203417f6a20013a000020024103490d00200020013a0002200020013a00012003417d6a20013a00002003417e6a20013a000020024107490d00200020013a00032003417c6a20013a000020024109490d002000410020006b41037122046a2203200141ff017141818284086c22013602002003200220046b417c7122046a2202417c6a200136020020044109490d002003200136020820032001360204200241786a2001360200200241746a200136020020044119490d002003200136021820032001360214200320013602102003200136020c200241706a20013602002002416c6a2001360200200241686a2001360200200241646a20013602002001ad220542208620058421052004200341047141187222016b2102200320016a2101034020024120490d0120012005370300200141186a2005370300200141106a2005370300200141086a2005370300200141206a2101200241606a21020c000b0b20000b3a01017f23808080800041106b220141808b84800036020c2000200128020c41076a41787122013602042000200136020020003f0036020c20000b120041808880800020004108108a808080000bc10101067f23808080800041106b22032480808080002003200136020c024002402001450d002000200028020c200241036a410020026b220471220520016a220641107622076a220836020c200020022000280204220120066a6a417f6a20047122023602040240200841107420024b0d002000410c6a200841016a360200200741016a21070b0240200740000d001080808080000b20012003410c6a41041086808080001a200120056a21000c010b410021000b200341106a24808080800020000b2e000240418088808000200120006c22004108108a808080002201450d002001410020001087808080001a0b20010b02000b0f004180888080001088808080001a0b3a01027f2000410120001b210102400340200110898080800022020d014100280290888080002200450d012000118080808000000c000b0b20020b0a002000108c808080000bce0301067f024020002001460d000240024002400240200120006b20026b410020024101746b4d0d0020012000734103712103200020014f0d012003450d02200021030c030b2000200120021086808080000f0b024020030d002001417f6a210402400340200020026a2203410371450d012002450d052003417f6a200420026a2d00003a00002002417f6a21020c000b0b2000417c6a21032001417c6a2104034020024104490d01200320026a200420026a2802003602002002417c6a21020c000b0b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b0b200241046a21052002417f7321064100210402400340200120046a2107200020046a2208410371450d0120022004460d03200820072d00003a00002005417f6a2105200641016a2106200441016a21040c000b0b200220046b2101410021030240034020014104490d01200820036a200720036a280200360200200341046a21032001417c6a21010c000b0b200720036a2101200820036a210320022006417c2006417c4b1b20056a417c716b20046b21020b03402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b0b20000b0900200041013602000b0900200041003602000b5201017f20004200370200200041086a22024100360200024020012d00004101710d00200020012902003702002002200141086a28020036020020000f0b20002001280208200128020410948080800020000b7701027f0240200241704f0d00024002402002410a4b0d00200020024101743a0000200041016a21030c010b200241106a4170712204108e8080800021032000200236020420002004410172360200200020033602080b2003200120021095808080001a200320026a41003a00000f0b108580808000000b1a0002402002450d0020002001200210868080800021000b20000b1d00024020002d0000410171450d002000280208108f808080000b20000bdf0101037f0240416f20016b2002490d000240024020002d00004101710d00200041016a21070c010b200028020821070b416f21080240200141e6ffffff074b0d00410b210820014101742209200220016a220220022009491b2202410b490d00200241106a41707121080b2008108e80808000210202402004450d002002200720041095808080001a0b0240200320056b20046b2203450d00200220046a20066a200720046a20056a20031095808080001a0b02402001410a460d002007108f808080000b20002002360208200020084101723602000f0b108580808000000bda0201067f02400240200141704f0d000240024020002d000022024101710d0020024101762103410a21040c010b20002802002202417e71417f6a2104200028020421030b410a2105024020032001200320014b1b2201410b490d00200141106a417071417f6a21050b0240024020052004460d0002402005410a470d0041012106200041016a210120002802082104410021070c040b200541016a108e808080002101200520044b0d0120010d010b0f0b024020002d000022024101710d0041012107200041016a2104410021060c020b2000280208210441012106410121070c010b108580808000000b0240024020024101710d00200241fe017141017621020c010b200028020421020b20012004200241016a1095808080001a02402006450d002004108f808080000b02402007450d0020002001360208200020033602042000200541016a4101723602000f0b200020034101743a00000bac0101037f0240024020002d000022024101712203450d002000280200417e71417f6a2104200028020421020c010b20024101762102410a21040b024002400240024020022004470d00200020044101200420044100410010978080800020002d0000410171450d010c020b20030d010b2000200241017441026a3a0000200041016a21000c010b2000200241016a360204200028020821000b200020026a220041003a0001200020013a00000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b1d0020004200370200200041086a41003602002000109a8080800020000be60604017f017e017f037e024002400240024002400240024002400240024002400240024002400240024002400240024020024200510d0020034200510d0120044200510d02200479a7200279a76b220641c000490d040c0f0b20044200510d022005450d0f420021022005420037030820052001370300420021010c110b20044200510d0420014200510d062004427f7c22072004834200510d0a200479a7200279a76b2208413f4f0d0d2002413f20086bad2207862001200841016a2208ad22098884210a2002200988210b200120078621090c090b2003427f7c22072003834200510d064200210b200379a741c1006a200279a76b220841c000470d0720012109420021070c020b02402005450d0020054200370308200520012003823703000b200120038021010c0d0b42002107200641016a220841c000470d02200121094200210b0b2002210a41c00021080c080b2005450d094200210120054200370308200542003703000c0a0b2002413f20066bad22098620012008ad220b8884210a2002200b88210b200120098621090c060b02402005450d0020054200370300200520022004823703080b200220048021010c080b02402005450d0020054200370308200520072001833703000b20034201510d08200242c00020037a22077d42ffffffff0f83862001200788842101200220078821020c080b2008413f4b0d02200241c00020086bad22078620012008ad22098884210a2002200988210b200120078621090b420021070c020b02402005450d0020052001370300200520072002833703080b200220047a8821010c040b200241800120086bad2207862001200841406aad220a888421092002200a88210a200120078621070b41002106024003402008450d01200a423f88200b4201868422022002427f8520047c200a4201862009423f8884220a427f85220220037c200254ad7c423f8722022004837d200a2002200383220154ad7d210b200a20017d210a2007423f8820094201868421092008417f6a210820074201862006ad8421072002a741017121060c000b0b2007423f8820094201868421022007420186210702402005450d002005200b3703082005200a3703000b2007427e832006ad8421010c030b2005450d0020052001370300200520023703080b420021010b420021020b20002001370300200020023703080b4d01017f23808080800041106b2205248080808000200520012002200320044100109c80808000200529030021012000200541086a29030037030820002001370300200541106a2480808080000b7501017e2000200420017e200220037e7c20034220882204200142208822027e7c200342ffffffff0f832203200142ffffffff0f8322017e2205422088200320027e7c22034220887c200342ffffffff0f83200420017e7c22034220887c37030820002003422086200542ffffffff0f83843703000bb60101037f4194888080001091808080004100280298888080002100024003402000450d01024003404100410028029c888080002202417f6a220136029c8880800020024101480d01200020014102746a22004184016a2802002102200041046a2802002100419488808000109280808000200220001181808080000041948880800010918080800041002802988880800021000c000b0b4100412036029c88808000410020002802002200360298888080000c000b0b0bcd0101027f419488808000109180808000024041002802988880800022030d0041a0888080002103410041a088808000360298888080000b02400240410028029c8880800022044120470d004184024101108b808080002203450d0141002104200341002802988880800036020041002003360298888080004100410036029c888080000b4100200441016a36029c88808000200320044102746a22034184016a2001360200200341046a200036020041948880800010928080800041000f0b419488808000109280808000417f0b0f0041a48a8080001096808080001a0b89010020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d00200010a38080800020012802044f0d00024020024104710d00200042003702000c010b1080808080000b024002402002411071450d00200010a38080800020012802044d0d0020024104710d01200042003702000b20000f0b10808080800020000b3400024002402000280204450d0020002802002c0000417f4c0d0141010f0b41000f0b200010a480808000200010a5808080006a0b280002402000280204450d0020002802002c0000417f4c0d0041000f0b200010aa8080800041016a0bc90301047f0240024002402000280204450d00200010ab808080004101210120002802002c00002202417f4c0d010c020b41000f0b0240200241ff0171220141b7014b0d00200141807f6a0f0b024002400240200241ff0171220241bf014b0d000240200041046a22032802002202200141c97e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b200141384f0d010c020b0240200241f7014b0d00200141c07e6a0f0b0240200041046a22032802002202200141897e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b20014138490d010b200141ff7d490d010b10808080800020010f0b20010b5102017f017e23808080800041306b220124808080800020012000290200220237031020012002370308200141186a200141086a411410a28080800010a3808080002100200141306a24808080800020000b6a01037f02400240024020012802002204450d0041002105200320026a200128020422064b0d0120062002490d014100210120062003490d02200620026b20032003417f461b2101200420026a21050c020b410021050b410021010b20002001360204200020053602000b3301017f0240200110a580808000220220012802044d0d001080808080000b20002001200110a480808000200210a7808080000bd003020a7f017e23808080800041c0006b220324808080800002402001280208220420024d0d00200341386a200110a880808000200320032903383703182001200341186a10a68080800036020c200341306a200110a880808000410021044100210541002106024020032802302207450d00410021054100210620032802342208200128020c2209490d00200820092009417f461b2105200721060b20012006360210200141146a2005360200200141086a41003602000b200141106a2106200141146a21092001410c6a2107200141086a210802400340200420024f0d012009280200450d01200341306a200110a88080800041002104024002402003280230220a450d00410021052003280234220b2007280200220c490d01200a200c6a2105200b200c6b21040c010b410021050b20092004360200200620053602002003200436022c2003200536022820032003290328370310200341306a20064100200341106a10a68080800010a78080800020062003290330220d37020020072007280200200d422088a76a3602002008200828020041016a22043602000c000b0b20032006290200220d3703202003200d3703082000200341086a411410a2808080001a200341c0006a2480808080000b4701017f4100210102402000280204450d00024020002802002d0000220041bf014b0d00200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b5401017f024020002802040d001080808080000b0240200028020022012d0000418101470d000240200041046a28020041014b0d00108080808000200028020021010b20012c00014100480d001080808080000b0bbc0101047f024002402000280204450d00200010ab80808000200028020022012c000022024100480d0120024100470f0b41000f0b410121030240200241807f460d000240200241ff0171220441b7014b0d000240200041046a28020041014b0d00108080808000200028020021010b20012d00014100470f0b41002103200441bf014b0d000240200041046a280200200241ff017141ca7e6a22024b0d00108080808000200028020021010b200120026a2d000041004721030b20030b2701017f200020012802002203200320012802046a10ae808080002000200210af8080800020000b34002000200220016b220210b080808000200028020020002802046a200120021086808080001a2000200028020420026a3602040bb60201087f02402001450d002000410c6a2102200041106a2103200041046a21040340200328020022052002280200460d010240200541786a28020020014f0d00108080808000200328020021050b200541786a2206200628020020016b220136020020010d01200320063602002000410120042802002005417c6a28020022016b220510b180808000220741016a20054138491b2206200428020022086a10b280808000200120002802006a220920066a2009200820016b1090808080001a02400240200541374b0d00200028020020016a200541406a3a00000c010b0240200741f7016a220641ff014b0d00200028020020016a20063a00002000280200200720016a6a210103402005450d02200120053a0000200541087621052001417f6a21010c000b0b1080808080000b410121010c000b0b0b21000240200028020420016a220120002802084d0d002000200110b3808080000b0b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b13002000200110b380808000200020013602040b4501017f0240200028020820014f0d0020011089808080002202200028020020002802041086808080001a200010bd80808000200041086a2001360200200020023602000b0b29002000410110b080808000200028020020002802046a20013a00002000200028020441016a3602040b3c01017f0240200110b180808000220320026a2202418002480d001080808080000b2000200241ff017110b48080800020002001200310b6808080000b44002000200028020420026a10b280808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0bf90101037f23808080800041106b22032480808080002001280200210420012802042105024002402002450d004100210102400340200420016a2102200120054f0d0120022d00000d01200141016a21010c000b0b200520016b21050c010b200421020b0240024002400240024020054101470d0020022c000022014100480d012000200141ff017110b4808080000c040b200541374b0d010b200020054180017341ff017110b4808080000c010b2000200541b70110b5808080000b2003200536020c200320023602082003200329030837030020002003410010ad808080001a0b2000410110af80808000200341106a24808080800020000bc40101037f02400240024020012002844200510d00200142ff005620024200522002501b0d0120002001a741ff017110b4808080000c020b200041800110b4808080000c010b024002402001200210b980808000220341374b0d00200020034180017341ff017110b4808080000c010b0240200310ba80808000220441b7016a2205418002490d001080808080000b2000200541ff017110b48080800020002003200410bb808080000b200020012002200310bc808080000b2000410110af8080800020000b3501017f41002102024003402000200184500d0120004208882001423886842100200241016a2102200142088821010c000b0b20020b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b44002000200028020420026a10b280808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0b54002000200028020420036a10b280808000200028020020002802046a417f6a2100024003402001200284500d01200020013c0000200142088820024238868421012000417f6a2100200242088821020c000b0b0b1700024020002802002200450d002000108c808080000b0b240041a48a808000109b808080001a418180808000410041808880800010a0808080001a0b940203037f027e027f23808080800041206b2203248080808000200010c08080800022044128109880808000200341186a21050340200341106a20012002420a4200109d80808000200320032903102206200529030022074276427f109e808080002004200329030020017ca741bc8a8080006a2c00001099808080002001420956210020024200522108200250210920062101200721022000200820091b0d000b0240200428020420042d00002200410176200041017122001b2208450d002004280208200441016a20001b220020086a417f6a21080340200020084f0d0120002d00002109200020082d00003a0000200820093a00002008417f6a2108200041016a21000c000b0b200341206a2480808080000b3b01017f20004200370200200041086a410036020041002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b20000b0f0041b08a8080001096808080001a0b0e0020002002420010bf808080000b0e0020002002420010bf808080000bf50903037f017e017f23808080800041e0006b2200248080808000108480808000108180808000220110898080800022021082808080002000200136024c2000200236024820002000290348370310200041c8006a200041286a200041106a411c10a2808080002201410010a98080800002400240200041c8006a10c58080800022034200510d00200341c78a80800010c680808000510d010240200341cc8a80800010c680808000520d00200041d8006a22014200370300200041d0006a420037030020004200370348200041c8006a4206420010b8808080001a024020002802542001280200460d001080808080000b2000280248200028024c108380808000200041c8006a10c7808080001a0c020b0240200341d18a80800010c680808000520d00200041d8006a22014200370300200041d0006a420037030020004200370348200041c8006a429003420010b8808080001a024020002802542001280200460d001080808080000b2000280248200028024c108380808000200041c8006a10c7808080001a0c020b0240200341d78a80800010c680808000520d00200041c8006a2001410110a980808000200041c8006a10ab8080800002400240200041c8006a10ac80808000450d00200028024c450d0020002802482d000041c001490d010b1080808080000b200041c0006a200041c8006a10c8808080000240200028024422014102490d001080808080000b2000280240210242002103024003402001450d012001417f6a210120022d0000410174ad42fe01832103200241016a21020c000b0b200041d8006a22014200370300200041d0006a420037030020004200370348200041c8006a2003420010b8808080001a024020002802542001280200460d001080808080000b2000280248200028024c108380808000200041c8006a10c7808080001a0c020b0240200341de8a80800010c680808000520d00200041c8006a2001410110a980808000200041c8006a10ab8080800002400240200041c8006a10ac80808000450d00200028024c450d0020002802482d000041c001490d010b1080808080000b200041c0006a200041c8006a10c8808080000240200028024422014105490d001080808080000b4100210420002802402102024003402001450d012001417f6a2101200441087420022d0000722104200241016a21020c000b0b200041d8006a22014200370300200041d0006a420037030020004200370348200041c8006a2004410174ad420010b8808080001a024020002802542001280200460d001080808080000b2000280248200028024c108380808000200041c8006a10c7808080001a0c020b0240200341e68a80800010c680808000520d00200042003703402001200041c0006a10c98080800020002903402103200041d8006a22014200370300200041d0006a420037030020004200370348200041c8006a2003420186420010b8808080001a024020002802542001280200460d001080808080000b2000280248200028024c108380808000200041c8006a10c7808080001a0c020b0240200341ee8a80800010c680808000520d00200041003602242000418280808000360220200020002903203703002001200010ca808080000c020b200341f48a80800010c680808000520d002000410036021c2000418380808000360218200020002903183703082001200041086a10ca808080000c010b1080808080000b200041e0006a2480808080000bb00103017f017e017f23808080800041106b2201248080808000200010ab8080800002400240200010ac80808000450d002000280204450d0020002802002d000041c001490d010b1080808080000b200141086a200010c8808080000240200128020c22004109490d001080808080000b4200210220012802082103024003402000450d012000417f6a210020024208862003310000842102200341016a21030c000b0b200141106a24808080800020020b3a01027e42a5c688a1c89ca7f94b21010240034020003000002202500d01200041016a2100200142b383808080207e20028521010c000b0b20010b3001017f0240200028020c2201450d00200041106a20013602002001108f808080000b2000280200108c8080800020000b870201057f0240200110a5808080002202200128020422034d0d00108080808000200141046a28020021030b2001280200210402400240024002400240024002402003450d004100210120042c00002205417f4c0d012004450d030c040b410021010c010b0240200541ff0171220641bf014b0d0041002101200541ff017141b801490d01200641c97e6a21010c010b41002101200541ff017141f801490d00200641897e6a21010b200141016a210120040d010b410021050c010b41002105200120026a20034b0d0020032001490d004100210620032002490d01200320016b20022002417f461b2106200420016a21050c010b410021060b20002006360204200020053602000b3f01017f23808080800041206b2202248080808000200241086a2000410110a9808080002001200241086a10c580808000370300200241206a2480808080000ba50201037f23808080800041e0006b22022480808080002001280200210320012802042101200242003703182000200241186a10c980808000200241206a200241106a20014101756a2002290318200311828080800000200241d0006a22004200370300200241c8006a4200370300200242003703402002200241306a200241206a1093808080002201280208200141016a20022d0030220341017122041b36025820022001280204200341017620041b36025c20022002290358370308200241c0006a200241086a410010b7808080001a20011096808080001a0240200228024c2000280200460d001080808080000b20022802402002280244108380808000200241c0006a10c7808080001a200241206a1096808080001a200241e0006a2480808080000b240041b08a80800010c0808080001a418480808000410041808880800010a0808080001a0b0b880302004180080bbc02000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041bc0a0b3e3031323334353637383900696e697400696e743800696e7436340075696e7438740075696e743332740075696e7436347400753132387400753235367400";

    public static String BINARY = BINARY_0;

    public static final String FUNC_U256T = "u256t";

    public static final String FUNC_UINT32T = "uint32t";

    public static final String FUNC_U128T = "u128t";

    public static final String FUNC_INT8 = "int8";

    public static final String FUNC_INT32 = "int32";

    public static final String FUNC_INT64 = "int64";

    public static final String FUNC_UINT8T = "uint8t";

    public static final String FUNC_UINT64T = "uint64t";

    protected IntegerDataTypeContract_1(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected IntegerDataTypeContract_1(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<String> u256t(Uint64 input) {
        final WasmFunction function = new WasmFunction(FUNC_U256T, Arrays.asList(input), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<Uint32> uint32t(Uint32 input) {
        final WasmFunction function = new WasmFunction(FUNC_UINT32T, Arrays.asList(input), Uint32.class);
        return executeRemoteCall(function, Uint32.class);
    }

    public RemoteCall<String> u128t(Uint64 input) {
        final WasmFunction function = new WasmFunction(FUNC_U128T, Arrays.asList(input), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static RemoteCall<IntegerDataTypeContract_1> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(IntegerDataTypeContract_1.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<IntegerDataTypeContract_1> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(IntegerDataTypeContract_1.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public RemoteCall<Int16> int8() {
        final WasmFunction function = new WasmFunction(FUNC_INT8, Arrays.asList(), Int16.class);
        return executeRemoteCall(function, Int16.class);
    }

    public RemoteCall<Int32> int32() {
        final WasmFunction function = new WasmFunction(FUNC_INT32, Arrays.asList(), Int32.class);
        return executeRemoteCall(function, Int32.class);
    }

    public RemoteCall<Int64> int64() {
        final WasmFunction function = new WasmFunction(FUNC_INT64, Arrays.asList(), Int64.class);
        return executeRemoteCall(function, Int64.class);
    }

    public RemoteCall<Uint8> uint8t(Uint8 input) {
        final WasmFunction function = new WasmFunction(FUNC_UINT8T, Arrays.asList(input), Uint8.class);
        return executeRemoteCall(function, Uint8.class);
    }

    public RemoteCall<Uint64> uint64t(Uint64 input) {
        final WasmFunction function = new WasmFunction(FUNC_UINT64T, Arrays.asList(input), Uint64.class);
        return executeRemoteCall(function, Uint64.class);
    }

    public static IntegerDataTypeContract_1 load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new IntegerDataTypeContract_1(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static IntegerDataTypeContract_1 load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new IntegerDataTypeContract_1(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
