package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Uint64;
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
 * <p>Generated with platon-web3j version 0.9.1.1-SNAPSHOT.
 */
public class ContractCrossCallStorageString extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001751260000060017f0060077f7f7f7f7f7f7f017f6000017f60027f7f0060027f7f017f60047f7f7f7f017f60047f7f7f7f0060037f7f7f017f60017f017f60037f7f7f0060087f7f7f7f7f7f7f7f0060037f7e7e017f60027e7e017f60047f7e7e7f0060057f7f7f7e7e017e60027f7e0060017f017e02bb010803656e760c706c61746f6e5f70616e6963000003656e760b706c61746f6e5f63616c6c000203656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000303656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e000403656e7617706c61746f6e5f6765745f73746174655f6c656e677468000503656e7610706c61746f6e5f6765745f7374617465000603656e7610706c61746f6e5f7365745f73746174650007036f6e00000808090908050100090108010101050a080905080b08010901000801080909090907040a090109080a0404090404050405090405040a0a080c0d090a0e01050604090507010804000901010f09100404100404090909090409100008111109090404040504040401090904000405017001030305030100020615037f0141f08a040b7f0041f08a040b7f0041e70a0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300080b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974002306696e766f6b6500640908010041010b0225540ac1736e08001011105110750b02000bc60a010b7f2002410f6a210341002104410020026b21052002410e6a2106410120026b21072002410d6a2108410220026b210902400340200020046a210b200120046a210a20022004460d01200a410371450d01200b200a2d00003a00002003417f6a2103200541016a21052006417f6a2106200741016a21072008417f6a2108200941016a2109200441016a21040c000b0b200220046b210c02400240024002400240200b410371220d450d00200c4120490d03200d4101460d01200d4102460d02200d4103470d03200b200120046a28020022063a0000200041016a210c200220046b417f6a21092004210b0240034020094113490d01200c200b6a220a2001200b6a220741046a2802002208411874200641087672360200200a41046a200741086a2802002206411874200841087672360200200a41086a2007410c6a2802002208411874200641087672360200200a410c6a200741106a2802002206411874200841087672360200200b41106a210b200941706a21090c000b0b2002417f6a2005416d2005416d4b1b20036a4170716b20046b210c2001200b6a41016a210a2000200b6a41016a210b0c030b200c210a02400340200a4110490d01200020046a220b200120046a2207290200370200200b41086a200741086a290200370200200441106a2104200a41706a210a0c000b0b02400240200c4108710d00200120046a210a200020046a21040c010b200020046a220b200120046a2204290200370200200441086a210a200b41086a21040b0240200c410471450d002004200a280200360200200a41046a210a200441046a21040b0240200c410271450d002004200a2f00003b0000200441026a2104200a41026a210a0b200c410171450d032004200a2d00003a000020000f0b200b200120046a220a28020022063a0000200b41016a200a41016a2f00003b0000200041036a210c200220046b417d6a21052004210b0240034020054111490d01200c200b6a220a2001200b6a220741046a2802002203410874200641187672360200200a41046a200741086a2802002206410874200341187672360200200a41086a2007410c6a2802002203410874200641187672360200200a410c6a200741106a2802002206410874200341187672360200200b41106a210b200541706a21050c000b0b2002417d6a2009416f2009416f4b1b20086a4170716b20046b210c2001200b6a41036a210a2000200b6a41036a210b0c010b200b200120046a220a28020022083a0000200b41016a200a41016a2d00003a0000200041026a210c200220046b417e6a21052004210b0240034020054112490d01200c200b6a220a2001200b6a220941046a2802002203411074200841107672360200200a41046a200941086a2802002208411074200341107672360200200a41086a2009410c6a2802002203411074200841107672360200200a410c6a200941106a2802002208411074200341107672360200200b41106a210b200541706a21050c000b0b2002417e6a2007416e2007416e4b1b20066a4170716b20046b210c2001200b6a41026a210a2000200b6a41026a210b0b0240200c411071450d00200b200a2d00003a0000200b200a280001360001200b200a290005370005200b200a2f000d3b000d200b200a2d000f3a000f200b41106a210b200a41106a210a0b0240200c410871450d00200b200a290000370000200b41086a210b200a41086a210a0b0240200c410471450d00200b200a280000360000200b41046a210b200a41046a210a0b0240200c410271450d00200b200a2f00003b0000200b41026a210b200a41026a210a0b200c410171450d00200b200a2d00003a00000b20000bfb0202027f017e02402002450d00200020013a0000200020026a2203417f6a20013a000020024103490d00200020013a0002200020013a00012003417d6a20013a00002003417e6a20013a000020024107490d00200020013a00032003417c6a20013a000020024109490d002000410020006b41037122046a2203200141ff017141818284086c22013602002003200220046b417c7122046a2202417c6a200136020020044109490d002003200136020820032001360204200241786a2001360200200241746a200136020020044119490d002003200136021820032001360214200320013602102003200136020c200241706a20013602002002416c6a2001360200200241686a2001360200200241646a20013602002001ad220542208620058421052004200341047141187222016b2102200320016a2101034020024120490d0120012005370300200141186a2005370300200141106a2005370300200141086a2005370300200141206a2101200241606a21020c000b0b20000b3a01017f23808080800041106b220141f08a84800036020c2000200128020c41076a41787122013602042000200136020020003f0036020c20000b120041808880800020004108108e808080000bc10101067f23808080800041106b22032480808080002003200136020c024002402001450d002000200028020c200241036a410020026b220471220520016a220641107622076a220836020c200020022000280204220120066a6a417f6a20047122023602040240200841107420024b0d002000410c6a200841016a360200200741016a21070b0240200740000d001080808080000b20012003410c6a4104108a808080001a200120056a21000c010b410021000b200341106a24808080800020000b2e000240418088808000200120006c22004108108e808080002201450d00200141002000108b808080001a0b20010b02000b0f00418088808000108c808080001a0b3a01027f2000410120001b2101024003402001108d8080800022020d014100280290888080002200450d012000118080808000000c000b0b20020b0a0020001090808080000bce0301067f024020002001460d000240024002400240200120006b20026b410020024101746b4d0d0020012000734103712103200020014f0d012003450d02200021030c030b200020012002108a808080000f0b024020030d002001417f6a210402400340200020026a2203410371450d012002450d052003417f6a200420026a2d00003a00002002417f6a21020c000b0b2000417c6a21032001417c6a2104034020024104490d01200320026a200420026a2802003602002002417c6a21020c000b0b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b0b200241046a21052002417f7321064100210402400340200120046a2107200020046a2208410371450d0120022004460d03200820072d00003a00002005417f6a2105200641016a2106200441016a21040c000b0b200220046b2101410021030240034020014104490d01200820036a200720036a280200360200200341046a21032001417c6a21010c000b0b200720036a2101200820036a210320022006417c2006417c4b1b20056a417c716b20046b21020b03402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b0b20000b0900200041013602000b0900200041003602000b0900108980808000000b5201017f20004200370200200041086a22024100360200024020012d00004101710d00200020012902003702002002200141086a28020036020020000f0b20002001280208200128020410998080800020000b7701027f0240200241704f0d00024002402002410a4b0d00200020024101743a0000200041016a21030c010b200241106a417071220410928080800021032000200236020420002004410172360200200020033602080b200320012002109a808080001a200320026a41003a00000f0b108980808000000b1a0002402002450d00200020012002108a8080800021000b20000b1d00024020002d0000410171450d0020002802081093808080000b20000b3d01027f024020002001460d0020002001280208200141016a20012d0000220241017122031b2001280204200241017620031b109d808080001a0b20000bbb0101037f410a2103024020002d000022044101712205450d002000280200417e71417f6a21030b02400240024002400240200320024f0d0020050d01200441017621050c020b20050d02200041016a21030c030b200028020421050b20002003200220036b20054100200520022001109e8080800020000f0b200028020821030b200320012002109f808080001a200320026a41003a0000024020002d00004101710d00200020024101743a000020000f0b2000200236020420000b8f0201037f0240416e20016b2002490d000240024020002d00004101710d00200041016a21080c010b200028020821080b416f21090240200141e6ffffff074b0d00410b21092001410174220a200220016a22022002200a491b2202410b490d00200241106a41707121090b2009109280808000210202402004450d00200220082004109a808080001a0b02402006450d00200220046a20072006109a808080001a0b0240200320056b220320046b2207450d00200220046a20066a200820046a20056a2007109a808080001a0b02402001410a460d0020081093808080000b200020023602082000200320066a220436020420002009410172360200200220046a41003a00000f0b108980808000000b1a0002402002450d0020002001200210948080800021000b20000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b1d0020004200370200200041086a4100360200200010a08080800020000b0900108980808000000bb60101037f4194888080001095808080004100280298888080002100024003402000450d01024003404100410028029c888080002202417f6a220136029c8880800020024101480d01200020014102746a22004184016a2802002102200041046a2802002100419488808000109680808000200220001181808080000041948880800010958080800041002802988880800021000c000b0b4100412036029c88808000410020002802002200360298888080000c000b0b0bcd0101027f419488808000109580808000024041002802988880800022030d0041a0888080002103410041a088808000360298888080000b02400240410028029c8880800022044120470d004184024101108f808080002203450d0141002104200341002802988880800036020041002003360298888080004100410036029c888080000b4100200441016a36029c88808000200320044102746a22034184016a2001360200200341046a200036020041948880800010968080800041000f0b419488808000109680808000417f0b0f0041a48a808000109b808080001a0b89010020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d00200010a78080800020012802044f0d00024020024104710d00200042003702000c010b1080808080000b024002402002411071450d00200010a78080800020012802044d0d0020024104710d01200042003702000b20000f0b10808080800020000b3400024002402000280204450d0020002802002c0000417f4c0d0141010f0b41000f0b200010a880808000200010a9808080006a0b280002402000280204450d0020002802002c0000417f4c0d0041000f0b200010ae8080800041016a0bc90301047f0240024002402000280204450d00200010af808080004101210120002802002c00002202417f4c0d010c020b41000f0b0240200241ff0171220141b7014b0d00200141807f6a0f0b024002400240200241ff0171220241bf014b0d000240200041046a22032802002202200141c97e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b200141384f0d010c020b0240200241f7014b0d00200141c07e6a0f0b0240200041046a22032802002202200141897e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b20014138490d010b200141ff7d490d010b10808080800020010f0b20010b5102017f017e23808080800041306b220124808080800020012000290200220237031020012002370308200141186a200141086a411410a68080800010a7808080002100200141306a24808080800020000b6a01037f02400240024020012802002204450d0041002105200320026a200128020422064b0d0120062002490d014100210120062003490d02200620026b20032003417f461b2101200420026a21050c020b410021050b410021010b20002001360204200020053602000b3301017f0240200110a980808000220220012802044d0d001080808080000b20002001200110a880808000200210ab808080000bd003020a7f017e23808080800041c0006b220324808080800002402001280208220420024d0d00200341386a200110ac80808000200320032903383703182001200341186a10aa8080800036020c200341306a200110ac80808000410021044100210541002106024020032802302207450d00410021054100210620032802342208200128020c2209490d00200820092009417f461b2105200721060b20012006360210200141146a2005360200200141086a41003602000b200141106a2106200141146a21092001410c6a2107200141086a210802400340200420024f0d012009280200450d01200341306a200110ac8080800041002104024002402003280230220a450d00410021052003280234220b2007280200220c490d01200a200c6a2105200b200c6b21040c010b410021050b20092004360200200620053602002003200436022c2003200536022820032003290328370310200341306a20064100200341106a10aa8080800010ab8080800020062003290330220d37020020072007280200200d422088a76a3602002008200828020041016a22043602000c000b0b20032006290200220d3703202003200d3703082000200341086a411410a6808080001a200341c0006a2480808080000b4701017f4100210102402000280204450d00024020002802002d0000220041bf014b0d00200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b5401017f024020002802040d001080808080000b0240200028020022012d0000418101470d000240200041046a28020041014b0d00108080808000200028020021010b20012c00014100480d001080808080000b0bbc0101047f024002402000280204450d00200010af80808000200028020022012c000022024100480d0120024100470f0b41000f0b410121030240200241807f460d000240200241ff0171220441b7014b0d000240200041046a28020041014b0d00108080808000200028020021010b20012d00014100470f0b41002103200441bf014b0d000240200041046a280200200241ff017141ca7e6a22024b0d00108080808000200028020021010b200120026a2d000041004721030b20030b2701017f200020012802002203200320012802046a10b2808080002000200210b38080800020000b34002000200220016b220210b480808000200028020020002802046a20012002108a808080001a2000200028020420026a3602040bb60201087f02402001450d002000410c6a2102200041106a2103200041046a21040340200328020022052002280200460d010240200541786a28020020014f0d00108080808000200328020021050b200541786a2206200628020020016b220136020020010d01200320063602002000410120042802002005417c6a28020022016b220510b580808000220741016a20054138491b2206200428020022086a10b680808000200120002802006a220920066a2009200820016b1094808080001a02400240200541374b0d00200028020020016a200541406a3a00000c010b0240200741f7016a220641ff014b0d00200028020020016a20063a00002000280200200720016a6a210103402005450d02200120053a0000200541087621052001417f6a21010c000b0b1080808080000b410121010c000b0b0b21000240200028020420016a220120002802084d0d002000200110b7808080000b0b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b13002000200110b780808000200020013602040b4501017f0240200028020820014f0d002001108d80808000220220002802002000280204108a808080001a200010c780808000200041086a2001360200200020023602000b0b6f01017f23808080800041106b2202248080808000024002402001450d0020022001360200200220002802043602042000410c6a200210b9808080000c010b20024100360208200242003703002000200210ba808080001a200210bb808080001a0b200241106a24808080800020000b3d01017f02402000280204220220002802084f0d0020022001290200370200200041046a2200200028020041086a3602000f0b2000200110bc808080000b5101027f23808080800041106b22022480808080002002200128020022033602082002200128020420036b36020c200220022903083703002000200210bd808080002101200241106a24808080800020010b2201017f024020002802002201450d002000200136020420011093808080000b20000b840101027f23808080800041206b2202248080808000200241086a2000200028020420002802006b41037541016a10c880808000200028020420002802006b410375200041086a10c980808000220328020820012902003702002003200328020841086a3602082000200310ca80808000200310cb808080001a200241206a2480808080000b7702027f017e23808080800041106b2202248080808000024002402001280204220341374b0d002000200341406a41ff017110be808080000c010b2000200341f70110bf808080000b2002200129020022043703082002200437030020002002410110b1808080002100200241106a24808080800020000b29002000410110b480808000200028020020002802046a20013a00002000200028020441016a3602040b3c01017f0240200110b580808000220320026a2202418002480d001080808080000b2000200241ff017110be8080800020002001200310c0808080000b44002000200028020420026a10b680808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0bf90101037f23808080800041106b22032480808080002001280200210420012802042105024002402002450d004100210102400340200420016a2102200120054f0d0120022d00000d01200141016a21010c000b0b200520016b21050c010b200421020b0240024002400240024020054101470d0020022c000022014100480d012000200141ff017110be808080000c040b200541374b0d010b200020054180017341ff017110be808080000c010b2000200541b70110bf808080000b2003200536020c200320023602082003200329030837030020002003410010b1808080001a0b2000410110b380808000200341106a24808080800020000bc40101037f02400240024020012002844200510d00200142ff005620024200522002501b0d0120002001a741ff017110be808080000c020b200041800110be808080000c010b024002402001200210c380808000220341374b0d00200020034180017341ff017110be808080000c010b0240200310c480808000220441b7016a2205418002490d001080808080000b2000200541ff017110be8080800020002003200410c5808080000b200020012002200310c6808080000b2000410110b38080800020000b3501017f41002102024003402000200184500d0120004208882001423886842100200241016a2102200142088821010c000b0b20020b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b44002000200028020420026a10b680808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0b54002000200028020420036a10b680808000200028020020002802046a417f6a2100024003402001200284500d01200020013c0000200142088820024238868421012000417f6a2100200242088821020c000b0b0b1700024020002802002200450d0020001090808080000b0b5301017f024020014180808080024f0d0041ffffffff0121020240200028020820002802006b220041037541feffffff004b0d0020012000410275220020002001491b21020b20020f0b200010a280808000000b5c01017f410021042000410036020c200041106a200336020002402001450d002003200110cc8080800021040b200020043602002000200420024103746a22033602082000410c6a200420014103746a3602002000200336020420000b7001017f200041086a20002802002000280204200141046a10cd80808000200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b2301017f200010ce80808000024020002802002201450d0020011093808080000b20000b0e0020002001410010cf808080000b2f01017f20032003280200200220016b22026b2204360200024020024101480d00200420012002108a808080001a0b0b0f002000200028020410d0808080000b2300024020014180808080024f0d0020014103741092808080000f0b108980808000000b2d01017f20002802082102200041086a21000240034020012002460d012000200241786a22023602000c000b0b0b240041a48a80800010a1808080001a418180808000410041808880800010a4808080001a0b1d0020004200370200200041086a4100360200200010d38080800020000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b0f0041b08a808000109b808080001a0baf0a03037f017e037f2380808080004180016b2205248080808000200541106a410036020020054200370308200541086a41bc8a808000410a1099808080000240024020052d000822064101710d0020064101762106200541086a41017221070c010b200528020c2106200528021021070b42a5c688a1c89ca7f94b2108024003402006450d012006417f6a2106200842b383808080207e2007300000852108200741016a21070c000b0b200541d8006a10d6808080002107200541c8006a200210988080800021022007410210b8808080001a200541c0006a4100360200200541286a41106a4200370300200541306a420037030020054200370328200541286a200810d780808000200541286a200541f0006a2002109880808000220610d8808080002006109b808080001a2007200528022810d9808080002007200810da808080002007200541f0006a2002109880808000220610db808080002006109b808080001a024002400240200728020c200741106a2206280200460d00108080808000200741046a21092007280200210a2007410c6a2802002006280200460d011080808080002007280200210b0c020b200741046a21092007280200210a0b200a210b0b4100210620054100360220200542003703180240200b20092802006a200a6b2209450d00200541186a200910dc8080800020094101480d00200528021c200a2009108a808080001a2005200528021c20096a36021c0b200541286a41047210dd808080001a2002109b808080001a200710de808080001a200541086a109b808080001a02402001280204220220012d000022074101762209200741017122071b220a4102490d002001280208200141016a20071b220b2d00004130470d00200b2d000141f8004641017421060b20054100360260200542003703580240200a41016a20066b410176220a450d00200541386a200541e0006a3602002005200a10928080800022073602582005200736025c200542003703302005420037032820052007200a6a360260200541286a10df808080001a20012d00002207410176210920074101712107200141046a28020021020b02400240024002402002200920071b410171450d002001280208200141016a20071b20066a2c000010e0808080002207417f460d01200520073a002820064101722106200541d8006a200541286a10e1808080000b200141016a210a200141046a2109200141086a210b0240024003402006200928020020012d00002207410176200741017122071b4f0d01200b280200200a20071b20066a22022c000010e0808080002107200241016a2c000010e08080800021022007417f460d022002417f460d022005200220074104746a3a0028200641026a2106200541d8006a200541286a10e1808080000c000b0b2005200528025822063602702005200528025c22073602742005200541e0006a2802003602782005420037025c200541d8006a21020c030b20054200370370200541f8006a21020c010b20054200370370200541f8006a21020b41002107410021060b20024100360200200541d8006a10e2808080001a0240200720066b2207450d00200541286a200620071094808080001a0b200541f0006a10e2808080001a200541d8006a200310e380808000200541f0006a200410e380808000200541286a20052802182206200528021c20066b20052802582206200528025c20066b20052802702206200528027420066b1081808080001a200541f0006a10e2808080001a200541d8006a10e2808080001a200541186a10e2808080001a20054180016a24808080800042000b2d0020004100360208200042003702002000410010f080808000200041146a41003602002000420037020c20000b8e0102017f017e4101210202402001428001540d004200210341002102024003402001200384500d0120014208882003423886842101200241016a2102200342088821030c000b0b024020024138490d00200210f28080800020026a21020b200241016a21020b0240200041186a280200450d00200041046a10f38080800021000b2000200028020020026a3602000b9a0101037f410121020240200128020420012d00002203410176200341017122041b2203450d004101210202400240024020034101470d002001280208200141016a20041b2c0000417f4a0d030c010b200341374b0d010b200341016a21020c010b2003200310f2808080006a41016a21020b0240200041186a280200450d00200041046a10f38080800021000b2000200028020020026a3602000b19000240200028020820014f0d002000200110f0808080000b0b0f0020002001420010c2808080001a0b6501037f23808080800041106b220224808080800020022001280208200141016a20012d0000220341017122041b36020820022001280204200341017620041b36020c2002200229030837030020002002410010c1808080001a200241106a2480808080000b3801017f02402001417f4c0d00200020011092808080002202360200200020023602042000200220016a3602080f0b200010a280808000000b8f0301067f200028020422012000280210220241087641fcffff07716a210302400240200028020822042001460d002001200028021420026a220541087641fcffff07716a280200200541ff07714102746a2105200041146a21062003280200200241ff07714102746a21020c010b200041146a210641002102410021050b0240034020052002460d01200241046a220220032802006b418020470d0020032802042102200341046a21030c000b0b20064100360200200041086a210302400340200420016b41027522024103490d012001280200109380808000200041046a2201200128020041046a2201360200200328020021040c000b0b02400240024020024101460d0020024102470d0241800821020c010b41800421020b200041106a20023602000b0240034020042001460d012001280200109380808000200141046a21010c000b0b200041086a22022802002101200041046a28020021040240034020042001460d0120022001417c6a22013602000c000b0b024020002802002201450d0020011093808080000b20000b2d01017f0240200028020c2201450d00200041106a200136020020011093808080000b200010f18080800020000b4b01037f2000280208210120002802042102200041086a21030240034020022001460d0120032001417f6a22013602000c000b0b024020002802002201450d0020011093808080000b20000b4b01017f4150210102400240200041506a41ff0171410a490d0041a97f21012000419f7f6a41ff017141064f0d010b200120006a0f0b200041496a417f200041bf7f6a41ff01714106491b0bef0101047f23808080800041206b2202248080808000024002402000280204220320002802084f0d00200320012d00003a0000200041046a2200200028020041016a3602000c010b2000200341016a20002802006b10ed808080002104200241186a200041086a3602004100210320024100360214200041046a28020020002802006b210502402004450d00200410928080800021030b20022003360208200241146a200320046a360200200320056a220320012d00003a00002002200336020c2002200341016a3602102000200241086a10ee80808000200241086a10df808080001a0b200241206a2480808080000b2201017f024020002802002201450d002000200136020420011093808080000b20000b6d02017f017e4100210220012103024003402003500d0120034208882103200241016a21020c000b0b20004100360208200042003702002000200210ec808080002000280204417f6a2102024003402001500d01200220013c00002002417f6a2102200142088821010c000b0b0b9d0503037f017e047f23808080800041a0016b22002480808080001088808080001082808080002201108d808080002202108380808000200041e8006a20002002200110e5808080002201410010ad8080800002400240200041e8006a10e68080800022034200510d000240200341c78a80800010e780808000520d00200041e8006a10e88080800010e9808080001a0c020b0240200341cc8a80800010e780808000520d00200041e8006a10d2808080002102200041f4006a10d28080800021042000420037038801200042003703800120004101360250200020013602182000200041d0006a36021c200041186a200210ea80808000200041186a200410ea80808000200041186a20004180016a10eb80808000200041186a20004188016a10eb80808000200041186a10e8808080002105200041186a20004190016a20021098808080002206200041c0006a2004109880808000220720002903800120002903880110d5808080002103200041d0006a10d6808080002201200310da808080000240200128020c200141106a280200460d001080808080000b20012802002001280204108480808000200110de808080001a2007109b808080001a2006109b808080001a200510e9808080001a2004109b808080001a2002109b808080001a0c020b200341dc8a80800010e780808000520d00200041e8006a10e880808000210420004190016a20004180016a1098808080002102200041186a10d6808080002201200041d0006a2002109880808000220510db808080002005109b808080001a0240200128020c200141106a280200460d001080808080000b20012802002001280204108480808000200110de808080001a2002109b808080001a200410e9808080001a0c010b1080808080000b200041a0016a2480808080000b4801017f23808080800041106b22032480808080002003200236020c200320013602082003200329030837030020002003411c10a6808080002100200341106a24808080800020000bb00103017f017e017f23808080800041106b2201248080808000200010af8080800002400240200010b080808000450d002000280204450d0020002802002d000041c001490d010b1080808080000b200141086a200010ef808080000240200128020c22004109490d001080808080000b4200210220012802082103024003402000450d012000417f6a210020024208862003310000842102200341016a21030c000b0b200141106a24808080800020020b3a01027e42a5c688a1c89ca7f94b21010240034020003000002202500d01200041016a2100200142b383808080207e20028521010c000b0b20010b9f0201077f23808080800041c0006b2201248080808000200010d2808080002102200042e299efdb8683ebcf58370310200041186a10d2808080002103200141286a10d6808080002204200029031010da808080000240200428020c200441106a280200460d001080808080000b0240024020042802002205200428020422061085808080002207450d002001410036022020014200370318200141186a200710ec808080002005200620012802182207200128021c20076b1086808080001a20012001280218220541016a200128021c2005417f736a10e580808000200310f480808000200141186a10e2808080001a200410de808080001a0c010b200410de808080001a20032002109c808080001a0b200141c0006a24808080800020000b9505010b7f23808080800041e0006b2201248080808000200141186a10d6808080002102200141c0006a41186a22034100360200200141c0006a41106a22044200370300200141c8006a2205420037030020014200370340200141c0006a200029031010d78080800020012802402106200141c0006a410472220710dd808080001a2002200610d9808080002002200029031010da80808000200041186a21080240200228020c200241106a280200460d001080808080000b200228020421092002280200210a200110d680808000210620034100360200200442003703002005420037030020014200370340200141c0006a200141306a2008109880808000220310d8808080002003109b808080001a2001280240210b200710dd808080001a4101109280808000220341fe013a0000200120033602402001200341016a2204360248200120043602440240200628020c200641106a280200460d0010808080800020012802442104200128024021030b0240200420036b2204200628020422056a220720062802084d0d002006200710f080808000200641046a28020021050b200628020020056a20032004108a808080001a200641046a2203200328020020046a36020020062001280244200b6a20012802406b10d9808080002006200141306a2008109880808000220310db808080002003109b808080001a024002402006410c6a2204280200200641106a2205280200460d001080808080002006280200210320042802002005280200460d011080808080000c010b200628020021030b200a20092003200641046a280200108780808000200141c0006a10e2808080001a200610de808080001a200210de808080001a2008109b808080001a2000109b808080001a200141e0006a24808080800020000b5701017f23808080800041206b2202248080808000200241086a2000280200200028020428020010ad80808000200241086a200110f48080800020002802042200200028020041016a360200200241206a2480808080000b5a01017f23808080800041206b2202248080808000200241086a2000280200200028020428020010ad808080002001200241086a10e68080800037030020002802042200200028020041016a360200200241206a2480808080000bd00201057f23808080800041206b220224808080800002400240024020002802042203200028020022046b220520014f0d00200028020820036b200120056b4f0d012000200110ed808080002104200241186a200041086a36020020024100360214200041046a28020020002802006b21064100210302402004450d00200410928080800021030b20022003360208200241146a200320046a3602002002200320066a22033602102002200336020c200520016b2101200241086a41086a21050340200341003a00002005200528020041016a2203360200200141016a22010d000b2000200241086a10ee80808000200241086a10df808080001a0c020b200520014d0d01200041046a200420016a3602000c010b200520016b2101200041046a21050340200341003a00002005200528020041016a2203360200200141016a22010d000b0b200241206a2480808080000b4c01017f02402001417f4c0d0041ffffffff0721020240200028020820002802006b220041feffffff034b0d0020012000410174220020002001491b21020b20020f0b200010a280808000000b9a0101037f200120012802042000280204200028020022026b22036b22043602040240200341004c0d00200420022003108a808080001a200141046a28020021040b2000280200210320002004360200200141046a22042003360200200041046a220328020021022003200128020836020020012002360208200028020821032000200128020c3602082001200336020c200120042802003602000b870201057f0240200110a9808080002202200128020422034d0d00108080808000200141046a28020021030b2001280200210402400240024002400240024002402003450d004100210120042c00002205417f4c0d012004450d030c040b410021010c010b0240200541ff0171220641bf014b0d0041002101200541ff017141b801490d01200641c97e6a21010c010b41002101200541ff017141f801490d00200641897e6a21010b200141016a210120040d010b410021050c010b41002105200120026a20034b0d0020032001490d004100210620032002490d01200320016b20022002417f461b2106200420016a21050c010b410021060b20002006360204200020053602000b4401017f0240200028020820014f0d002001108d8080800020002802002000280204108a808080002102200010f180808000200041086a2001360200200020023602000b0b0d0020002802001090808080000b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0bcd0301057f23808080800041206b2202248080808000024002400240024002402000280204450d0020002802002d000041c0014f0d00200241186a200010ef80808000200010a9808080002103024020022802182200450d00200228021c220420034f0d020b41002100200241106a41003602002002420037030841002105410021040c020b200241086a10d2808080001a0c030b200241106a4100360200200242003703080240200420032003417f461b220541704f0d00200020056a21042005410a4d0d01200541106a417071220610928080800021032002200536020c20022006410172360208200220033602100c020b200241086a109780808000000b200220054101743a0008200241086a41017221030b0240034020042000460d01200320002d00003a0000200341016a2103200041016a21000c000b0b200341003a00000b0240024020012d00004101710d00200141003b01000c010b200128020841003a00002001410036020420012d0000410171450d00200141086a280200109380808000200141003602000b20012002290308370200200141086a200241086a41086a280200360200200241086a10d380808000200241086a109b808080001a200241206a2480808080000b240041b08a80800010d2808080001a418280808000410041808880800010a4808080001a0b0bf50202004180080bbc02000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041bc0a0b2b7365745f737472696e6700696e69740063616c6c5f7365745f737472696e67006765745f737472696e6700";

    public static String BINARY = BINARY_0;

    public static final String FUNC_CALL_SET_STRING = "call_set_string";

    public static final String FUNC_GET_STRING = "get_string";

    protected ContractCrossCallStorageString(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected ContractCrossCallStorageString(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<ContractCrossCallStorageString> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractCrossCallStorageString.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<ContractCrossCallStorageString> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(ContractCrossCallStorageString.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public RemoteCall<TransactionReceipt> call_set_string(String target_address, String name, Uint64 value, Uint64 gas) {
        final WasmFunction function = new WasmFunction(FUNC_CALL_SET_STRING, Arrays.asList(target_address,name,value,gas), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<String> get_string() {
        final WasmFunction function = new WasmFunction(FUNC_GET_STRING, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static ContractCrossCallStorageString load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new ContractCrossCallStorageString(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static ContractCrossCallStorageString load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new ContractCrossCallStorageString(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
