package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Int32;
import com.platon.rlp.datatypes.Int64;
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
public class BubbleSort extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001560f60000060017f006000017f60027f7f0060027f7f017f60047f7f7f7f017f60047f7f7f7f0060037f7f7f017f60017f017f60037f7f7f0060037f7e7e017f60027e7e017f60047f7e7e7f0060017f017e60027f7e0002a9010703656e760c706c61746f6e5f70616e6963000003656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000203656e7610706c61746f6e5f6765745f696e707574000103656e760d706c61746f6e5f72657475726e000303656e7617706c61746f6e5f6765745f73746174655f6c656e677468000403656e7610706c61746f6e5f6765745f7374617465000503656e7610706c61746f6e5f7365745f73746174650006036d6c000007070808070401000801070101080108010007010708080808080607030908010808070903030803030403040803040309090a0b08090c0104050308040601070300010901040300070d0d0808030808080303080803010e0e08080803030804050308030803020508000405017001030305030100020615037f0141d08a040b7f0041d08a040b7f0041d00a0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974001a06696e766f6b6500500908010041010b021c4b0ab0796c08001010104a10720b02000bc60a010b7f2002410f6a210341002104410020026b21052002410e6a2106410120026b21072002410d6a2108410220026b210902400340200020046a210b200120046a210a20022004460d01200a410371450d01200b200a2d00003a00002003417f6a2103200541016a21052006417f6a2106200741016a21072008417f6a2108200941016a2109200441016a21040c000b0b200220046b210c02400240024002400240200b410371220d450d00200c4120490d03200d4101460d01200d4102460d02200d4103470d03200b200120046a28020022063a0000200041016a210c200220046b417f6a21092004210b0240034020094113490d01200c200b6a220a2001200b6a220741046a2802002208411874200641087672360200200a41046a200741086a2802002206411874200841087672360200200a41086a2007410c6a2802002208411874200641087672360200200a410c6a200741106a2802002206411874200841087672360200200b41106a210b200941706a21090c000b0b2002417f6a2005416d2005416d4b1b20036a4170716b20046b210c2001200b6a41016a210a2000200b6a41016a210b0c030b200c210a02400340200a4110490d01200020046a220b200120046a2207290200370200200b41086a200741086a290200370200200441106a2104200a41706a210a0c000b0b02400240200c4108710d00200120046a210a200020046a21040c010b200020046a220b200120046a2204290200370200200441086a210a200b41086a21040b0240200c410471450d002004200a280200360200200a41046a210a200441046a21040b0240200c410271450d002004200a2f00003b0000200441026a2104200a41026a210a0b200c410171450d032004200a2d00003a000020000f0b200b200120046a220a28020022063a0000200b41016a200a41016a2f00003b0000200041036a210c200220046b417d6a21052004210b0240034020054111490d01200c200b6a220a2001200b6a220741046a2802002203410874200641187672360200200a41046a200741086a2802002206410874200341187672360200200a41086a2007410c6a2802002203410874200641187672360200200a410c6a200741106a2802002206410874200341187672360200200b41106a210b200541706a21050c000b0b2002417d6a2009416f2009416f4b1b20086a4170716b20046b210c2001200b6a41036a210a2000200b6a41036a210b0c010b200b200120046a220a28020022083a0000200b41016a200a41016a2d00003a0000200041026a210c200220046b417e6a21052004210b0240034020054112490d01200c200b6a220a2001200b6a220941046a2802002203411074200841107672360200200a41046a200941086a2802002208411074200341107672360200200a41086a2009410c6a2802002203411074200841107672360200200a410c6a200941106a2802002208411074200341107672360200200b41106a210b200541706a21050c000b0b2002417e6a2007416e2007416e4b1b20066a4170716b20046b210c2001200b6a41026a210a2000200b6a41026a210b0b0240200c411071450d00200b200a2d00003a0000200b200a280001360001200b200a290005370005200b200a2f000d3b000d200b200a2d000f3a000f200b41106a210b200a41106a210a0b0240200c410871450d00200b200a290000370000200b41086a210b200a41086a210a0b0240200c410471450d00200b200a280000360000200b41046a210b200a41046a210a0b0240200c410271450d00200b200a2f00003b0000200b41026a210b200a41026a210a0b200c410171450d00200b200a2d00003a00000b20000bfb0202027f017e02402002450d00200020013a0000200020026a2203417f6a20013a000020024103490d00200020013a0002200020013a00012003417d6a20013a00002003417e6a20013a000020024107490d00200020013a00032003417c6a20013a000020024109490d002000410020006b41037122046a2203200141ff017141818284086c22013602002003200220046b417c7122046a2202417c6a200136020020044109490d002003200136020820032001360204200241786a2001360200200241746a200136020020044119490d002003200136021820032001360214200320013602102003200136020c200241706a20013602002002416c6a2001360200200241686a2001360200200241646a20013602002001ad220542208620058421052004200341047141187222016b2102200320016a2101034020024120490d0120012005370300200141186a2005370300200141106a2005370300200141086a2005370300200141206a2101200241606a21020c000b0b20000b3a01017f23808080800041106b220141d08a84800036020c2000200128020c41076a41787122013602042000200136020020003f0036020c20000b120041808880800020004108108d808080000bca0101067f23808080800041106b22032480808080002003200136020c024002400240024002402001450d002000200028020c200241036a410020026b220471220520016a220641107622016a220736020c200020022000280204220820066a6a417f6a2004712202360204200741107420024d0d0120010d020c030b410021000c030b2000410c6a200741016a360200200141016a21010b200140000d001080808080000b20082003410c6a41041089808080001a200820056a21000b200341106a24808080800020000b2e000240418088808000200120006c22004108108d808080002201450d00200141002000108a808080001a0b20010b02000b0f00418088808000108b808080001a0b3a01027f2000410120001b2101024003402001108c8080800022020d014100280290888080002200450d012000118080808000000c000b0b20020b0a002000108f808080000bce0301067f024020002001460d000240024002400240200120006b20026b410020024101746b4d0d0020012000734103712103200020014f0d012003450d02200021030c030b2000200120021089808080000f0b024020030d002001417f6a210402400340200020026a2203410371450d012002450d052003417f6a200420026a2d00003a00002002417f6a21020c000b0b2000417c6a21032001417c6a2104034020024104490d01200320026a200420026a2802003602002002417c6a21020c000b0b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b0b200241046a21052002417f7321064100210402400340200120046a2107200020046a2208410371450d0120022004460d03200820072d00003a00002005417f6a2105200641016a2106200441016a21040c000b0b200220046b2101410021030240034020014104490d01200820036a200720036a280200360200200341046a21032001417c6a21010c000b0b200720036a2101200820036a210320022006417c2006417c4b1b20056a417c716b20046b21020b03402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b0b20000b0900200041013602000b0900200041003602000b1d00024020002d0000410171450d0020002802081092808080000b20000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b1d0020004200370200200041086a4100360200200010978080800020000b0900108880808000000bb60101037f4194888080001094808080004100280298888080002100024003402000450d01024003404100410028029c888080002202417f6a220136029c8880800020024101480d01200020014102746a22004184016a2802002102200041046a2802002100419488808000109580808000200220001181808080000041948880800010948080800041002802988880800021000c000b0b4100412036029c88808000410020002802002200360298888080000c000b0b0bcd0101027f419488808000109480808000024041002802988880800022030d0041a0888080002103410041a088808000360298888080000b02400240410028029c8880800022044120470d004184024101108e808080002203450d0141002104200341002802988880800036020041002003360298888080004100410036029c888080000b4100200441016a36029c88808000200320044102746a22034184016a2001360200200341046a200036020041948880800010958080800041000f0b419488808000109580808000417f0b0f0041a48a8080001096808080001a0b89010020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000109e8080800020012802044f0d00024020024104710d00200042003702000c010b1080808080000b024002402002411071450d002000109e8080800020012802044d0d0020024104710d01200042003702000b20000f0b10808080800020000b3400024002402000280204450d0020002802002c0000417f4c0d0141010f0b41000f0b2000109f80808000200010a0808080006a0b280002402000280204450d0020002802002c0000417f4c0d0041000f0b200010a78080800041016a0bc90301047f0240024002402000280204450d00200010a8808080004101210120002802002c00002202417f4c0d010c020b41000f0b0240200241ff0171220141b7014b0d00200141807f6a0f0b024002400240200241ff0171220241bf014b0d000240200041046a22032802002202200141c97e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b200141384f0d010c020b0240200241f7014b0d00200141c07e6a0f0b0240200041046a22032802002202200141897e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b20014138490d010b200141ff7d490d010b10808080800020010f0b20010bf90102057f017e23808080800041206b2201248080808000200041046a21020240024020002802002203450d0041002104024020022802002205450d002005200041086a2802006a21040b200041086a2003360200200041046a2203200436020020012003290200220637031020012006370308200141186a20024100200141086a10a28080800010a3808080002003200129031822063702002000200028020022022006422088a722032002200220034b1b6b3602000c010b41002103024020022802002202450d002002200041086a2802006a21030b200041086a4100360200200041046a20033602000b200141206a24808080800020000b5102017f017e23808080800041306b220124808080800020012000290200220237031020012002370308200141186a200141086a4114109d80808000109e808080002100200141306a24808080800020000b6a01037f02400240024020012802002204450d0041002105200320026a200128020422064b0d0120062002490d014100210120062003490d02200620026b20032003417f461b2101200420026a21050c020b410021050b410021010b20002001360204200020053602000b9e0202037f017e23808080800041206b220324808080800020004200370200200041086a22044100360200200128020421050240024002402002450d00410021022005450d012005210220012802002d000041c001490d01200341186a200110a58080800020032003290318220637031020032006370308200341086a10a2808080002101200041086a4100200328021c220220012001417f461b200328021822054520022001497222011b2204360200200041046a4100200520011b3602002000200220046b3602000c020b200521020b20012802002105200128020421012000410036020020044100200220016b20054520022001497222021b360200200041046a4100200520016a20021b3602000b200341206a24808080800020000b3301017f0240200110a080808000220220012802044d0d001080808080000b200020012001109f80808000200210a3808080000bd003020a7f017e23808080800041c0006b220324808080800002402001280208220420024d0d00200341386a200110a580808000200320032903383703182001200341186a10a28080800036020c200341306a200110a580808000410021044100210541002106024020032802302207450d00410021054100210620032802342208200128020c2209490d00200820092009417f461b2105200721060b20012006360210200141146a2005360200200141086a41003602000b200141106a2106200141146a21092001410c6a2107200141086a210802400340200420024f0d012009280200450d01200341306a200110a58080800041002104024002402003280230220a450d00410021052003280234220b2007280200220c490d01200a200c6a2105200b200c6b21040c010b410021050b20092004360200200620053602002003200436022c2003200536022820032003290328370310200341306a20064100200341106a10a28080800010a38080800020062003290330220d37020020072007280200200d422088a76a3602002008200828020041016a22043602000c000b0b20032006290200220d3703202003200d3703082000200341086a4114109d808080001a200341c0006a2480808080000b4701017f4100210102402000280204450d00024020002802002d0000220041bf014b0d00200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b5401017f024020002802040d001080808080000b0240200028020022012d0000418101470d000240200041046a28020041014b0d00108080808000200028020021010b20012c00014100480d001080808080000b0bbc0101047f024002402000280204450d00200010a880808000200028020022012c000022024100480d0120024100470f0b41000f0b410121030240200241807f460d000240200241ff0171220441b7014b0d000240200041046a28020041014b0d00108080808000200028020021010b20012d00014100470f0b41002103200441bf014b0d000240200041046a280200200241ff017141ca7e6a22024b0d00108080808000200028020021010b200120026a2d000041004721030b20030bd20103027f017e047f23808080800041206b22012480808080004100210202402000280204450d0020002802002d000041c001490d00200141186a200010a58080800041002102200128021c210003402000450d012001200129031822033703102001200337030841002100200141086a10a28080800021040240024020012802182205450d0041002106200128021c22072004490d01200520046a2106200720046b21000c010b410021060b2001200036021c20012006360218200241016a21020c000b0b200141206a24808080800020020b2701017f200020012802002203200320012802046a10ac808080002000200210ad8080800020000b34002000200220016b220210ae80808000200028020020002802046a200120021089808080001a2000200028020420026a3602040bb60201087f02402001450d002000410c6a2102200041106a2103200041046a21040340200328020022052002280200460d010240200541786a28020020014f0d00108080808000200328020021050b200541786a2206200628020020016b220136020020010d01200320063602002000410120042802002005417c6a28020022016b220510af80808000220741016a20054138491b2206200428020022086a10b080808000200120002802006a220920066a2009200820016b1093808080001a02400240200541374b0d00200028020020016a200541406a3a00000c010b0240200741f7016a220641ff014b0d00200028020020016a20063a00002000280200200720016a6a210103402005450d02200120053a0000200541087621052001417f6a21010c000b0b1080808080000b410121010c000b0b0b21000240200028020420016a220120002802084d0d002000200110b1808080000b0b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b13002000200110b180808000200020013602040b4501017f0240200028020820014f0d002001108c808080002202200028020020002802041089808080001a200010c080808000200041086a2001360200200020023602000b0b6f01017f23808080800041106b2202248080808000024002402001450d0020022001360200200220002802043602042000410c6a200210b3808080000c010b20024100360208200242003703002000200210b4808080001a200210b5808080001a0b200241106a24808080800020000b3d01017f02402000280204220220002802084f0d0020022001290200370200200041046a2200200028020041086a3602000f0b2000200110b6808080000b5101027f23808080800041106b22022480808080002002200128020022033602082002200128020420036b36020c200220022903083703002000200210b7808080002101200241106a24808080800020010b2201017f024020002802002201450d002000200136020420011092808080000b20000b840101027f23808080800041206b2202248080808000200241086a2000200028020420002802006b41037541016a10c180808000200028020420002802006b410375200041086a10c280808000220328020820012902003702002003200328020841086a3602082000200310c380808000200310c4808080001a200241206a2480808080000b7702027f017e23808080800041106b2202248080808000024002402001280204220341374b0d002000200341406a41ff017110b8808080000c010b2000200341f70110b9808080000b2002200129020022043703082002200437030020002002410110ab808080002100200241106a24808080800020000b29002000410110ae80808000200028020020002802046a20013a00002000200028020441016a3602040b3c01017f0240200110af80808000220320026a2202418002480d001080808080000b2000200241ff017110b88080800020002001200310ba808080000b44002000200028020420026a10b080808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0bc40101037f02400240024020012002844200510d00200142ff005620024200522002501b0d0120002001a741ff017110b8808080000c020b200041800110b8808080000c010b024002402001200210bc80808000220341374b0d00200020034180017341ff017110b8808080000c010b0240200310bd80808000220441b7016a2205418002490d001080808080000b2000200541ff017110b88080800020002003200410be808080000b200020012002200310bf808080000b2000410110ad8080800020000b3501017f41002102024003402000200184500d0120004208882001423886842100200241016a2102200142088821010c000b0b20020b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b44002000200028020420026a10b080808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0b54002000200028020420036a10b080808000200028020020002802046a417f6a2100024003402001200284500d01200020013c0000200142088820024238868421012000417f6a2100200242088821020c000b0b0b1700024020002802002200450d002000108f808080000b0b5301017f024020014180808080024f0d0041ffffffff0121020240200028020820002802006b220041037541feffffff004b0d0020012000410275220020002001491b21020b20020f0b2000109980808000000b5c01017f410021042000410036020c200041106a200336020002402001450d002003200110c58080800021040b200020043602002000200420024103746a22033602082000410c6a200420014103746a3602002000200336020420000b7001017f200041086a20002802002000280204200141046a10c680808000200028020021022000200128020436020020012002360204200028020421022000200128020836020420012002360208200028020821022000200128020c3602082001200236020c200120012802043602000b2301017f200010c780808000024020002802002201450d0020011092808080000b20000b0e0020002001410010c8808080000b2f01017f20032003280200200220016b22026b2204360200024020024101480d002004200120021089808080001a0b0b0f002000200028020410c9808080000b2300024020014180808080024f0d0020014103741091808080000f0b108880808000000b2d01017f20002802082102200041086a21000240034020012002460d012000200241786a22023602000c000b0b0b240041a48a8080001098808080001a4181808080004100418088808000109b808080001a0b0f0041b08a8080001096808080001a0bbb0104057f017e017f017e2002417f6a2103200128020021044100210502400340200320054c0d01200320056b2106410021072004210202400340200720064e0d01024020022903002208200241086a2209290300220a570d002002200a370300200241086a20084220864220873703000b200741016a2107200921020c000b0b200541016a21050c000b0b200041186a10cd8080800020002001290200370218200041206a200128020836020020014100360208200142003702000b2e01017f024020002802002201450d0020002001360204200110928080800020004100360208200042003702000b0b6d01027f20004200370200200041003602080240200128020420012802006b2202450d002000200241037510cf80808000200141046a280200200128020022036b22014101480d00200041046a2202280200200320011089808080001a2002200228020020016a3602000b20000b3f01017f024020014180808080024f0d002000200110e7808080002202360200200020023602042000200220014103746a3602080f0b2000109980808000000b880403037f017e017f23808080800041f0006b22002480808080001087808080001081808080002201108c808080002202108280808000200041206a200041086a2002200110d1808080002201410010a68080800002400240200041206a10d28080800022034200510d000240200341bc8a80800010d380808000520d00200041206a10d48080800010d5808080001a0c020b0240200341c18a80800010d380808000520d00200041e0006a420037030020004200370358200041206a2001410110a680808000200041206a200041d8006a10d680808000200041206a2001410210a6808080002000200041206a10d2808080002203420188420020034201837d853e0264200041206a10d4808080002101200041206a200041c8006a200041d8006a10ce808080002202200028026410cc80808000200210d7808080001a200110d5808080001a200041d8006a10d7808080001a0c020b200341c68a80800010d380808000520d00200041206a10d4808080002104200041c8006a200041386a10ce808080002102200041d8006a10d8808080002201200210d98080800010da808080002001200210db808080000240200128020c200141106a280200460d001080808080000b20012802002001280204108380808000200110dc808080001a200210d7808080001a200410d5808080001a0c010b1080808080000b200041f0006a2480808080000b4801017f23808080800041106b22032480808080002003200236020c200320013602082003200329030837030020002003411c109d808080002100200341106a24808080800020000bda0202057f017e200010a88080800002400240200010a980808000450d002000280204450d0020002802002d000041c001490d010b1080808080000b0240200010a0808080002201200028020422024d0d00108080808000200041046a28020021020b200028020021030240024002400240024002402002450d004100210420032c00002200417f4a0d03200041ff0171220441bf014b0d0141002105200041ff017141b801490d02200441c97e6a21050c020b4101210420030d02410021000c030b41002105200041ff017141f801490d00200441897e6a21050b200541016a21040b41002100200420016a20024b0d0020022001490d004100210520022004490d01200320046a2105200220046b20012001417f461b22004109490d011080808080000c010b410021050b42002106024003402000450d012000417f6a210020064208862005310000842106200541016a21050c000b0b20060b3a01027e42a5c688a1c89ca7f94b21010240034020003000002202500d01200041016a2100200142b383808080207e20028521010c000b0b20010bd00401087f23808080800041c0006b220124808080800020004200370218200042b3aed899858e8cbc123703102000410036020820004200370200200041206a4100360200200141286a10d8808080002202200029031010e0808080000240200228020c200241106a280200460d001080808080000b200041186a210302400240024020022802002204200228020422051084808080002206450d0020014100360220200142003703182006417f4c0d01200141206a200610918080800041002006108a80808000220720066a22083602002001200836021c2001200736021820042005200720061085808080001a20012001280218220641016a200128021c2006417f736a10d180808000200310d680808000200141186a10dd808080001a200210dc808080001a0c020b200210dc808080001a0240200041046a2802002205200028020022026b22074103752206200041206a280200200041186a28020022046b4103754d0d00200310cd8080800020032003200610e88080800010cf8080800020074101480d022000411c6a2206280200200220071089808080001a2006200628020020076a3602000c020b024020022000411c6a28020020046b22036a20052006200341037522074b1b220820026b2203450d002004200220031093808080001a0b0240200620074d0d00200520086b22024101480d022000411c6a2206280200200820021089808080001a2006200628020020026a3602000c020b2000411c6a200420034103754103746a3602000c010b200141186a109980808000000b200141c0006a24808080800020000bb104010b7f23808080800041d0006b2201248080808000200141186a10d8808080002102200141306a41186a4100360200200141306a41106a4200370300200141386a420037030020014200370330200141306a200029031010e18080800020012802302103200141306a41047210e4808080001a2002200310da808080002002200029031010e080808000200041186a21040240200228020c200241106a280200460d001080808080000b2002280204210520022802002106200141306a10d8808080002103200410d98080800021074101109180808000220841fe013a0000200120083602082001200841016a22093602102001200936020c0240200328020c200341106a280200460d00108080808000200128020c2109200128020821080b0240200920086b22092003280204220a6a220b20032802084d0d002003200b10de80808000200341046a280200210a0b2003280200200a6a200820091089808080001a200341046a2208200828020020096a3602002003200128020c20076a20012802086b10da808080002003200410db80808000024002402003410c6a2209280200200341106a2204280200460d001080808080002003280200210820092802002004280200460d011080808080000c010b200328020021080b200620052008200341046a280200108680808000200141086a10dd808080001a200310dc808080001a200210dc808080001a200041186a10d7808080001a200010d7808080001a200141d0006a24808080800020000bca0303057f017e017f23808080800041e0006b2202248080808000024002402000280204450d0020002802002d000041c001490d00200010aa80808000210302402001280208200128020022046b41037520034f0d002001200241106a2003200128020420046b410375200141086a10e980808000220310ea80808000200310eb808080001a0b200241386a2000410110a4808080002103200141086a2105200241286a2000410010a480808000210603400240200341046a2200280200200641046a280200470d00200341086a280200200641086a280200460d030b20022000290200220737034820022007370308200241106a200241086a411c109d8080800010d2808080002207420188420020074201837d85210702400240200141046a2204280200220020052802004f0d00200020073703002004200041086a3602000c010b200241c8006a2001200020012802006b41037541016a10e880808000200428020020012802006b410375200510e9808080002100200241c8006a41086a2204280200220820073703002004200841086a3602002001200010ea80808000200010eb808080001a0b200310a1808080001a0c000b0b1080808080000b200241e0006a2480808080000b2201017f024020002802002201450d002000200136020420011092808080000b20000b2d0020004100360208200042003702002000410010de80808000200041146a41003602002000420037020c20000bcf0102027f017e23808080800041206b2201248080808000200141186a4100360200200141106a4200370300200141086a4200370300200142003703000240024020002802002000280204460d002001410010ec80808000200041046a2802002102200028020021000240034020022000460d012001200029030022034201862003423f878510e180808000200041086a21000c000b0b2001410110ec80808000200128020021000c010b41012100200141013602000b200141047210e4808080001a200141206a24808080800020000b19000240200028020820014f0d002000200110de808080000b0b5902017f017e2000200128020420012802006b41037510b2808080001a20012802042102200128020021010240034020022001460d012000200129030022034201862003423f878510e080808000200141086a21010c000b0b0b2d01017f0240200028020c2201450d00200041106a200136020020011092808080000b200010df8080800020000b2201017f024020002802002201450d002000200136020420011092808080000b20000b4401017f0240200028020820014f0d002001108c80808000200028020020002802041089808080002102200010df80808000200041086a2001360200200020023602000b0b0d002000280200108f808080000b0f0020002001420010bb808080001a0b8e0102017f017e4101210202402001428001540d004200210341002102024003402001200384500d0120014208882003423886842101200241016a2102200342088821030c000b0b024020024138490d00200210e28080800020026a21020b200241016a21020b0240200041186a280200450d00200041046a10e38080800021000b2000200028020020026a3602000b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b2e002000280204200028021420002802106a417f6a220041087641fcffff07716a280200200041ff07714102746a0bea0201067f23808080800041106b2201248080808000200028020422022000280210220341087641fcffff07716a21040240024020002802082002460d002004280200200341ff07714102746a21020c010b410021020b200141086a200010e580808000200128020c21030240034020032002460d01200241046a220220042802006b418020470d0020042802042102200441046a21040c000b0b20004100360214200041046a22032802002102200041086a2105024003402005280200220420026b41027522064103490d0120022802001092808080002003200328020041046a22023602000c000b0b02400240024020064101460d0020064102470d0241800821030c010b41800421030b200041106a20033602000b0240034020042002460d012002280200109280808000200241046a21020c000b0b2000200041046a28020010e680808000024020002802002202450d0020021092808080000b200141106a24808080800020000b5901037f20012802042202200128021020012802146a220341087641fcffff07716a21040240024020012802082002460d002004280200200341ff07714102746a21010c010b410021010b20002001360204200020043602000b2d01017f20002802082102200041086a21000240034020012002460d0120002002417c6a22023602000c000b0b0b2300024020004180808080024f0d0020004103741091808080000f0b108880808000000b5301017f024020014180808080024f0d0041ffffffff0121020240200028020820002802006b220041037541feffffff004b0d0020012000410275220020002001491b21020b20020f0b2000109980808000000b5a01017f410021042000410036020c200041106a200336020002402001450d00200110e78080800021040b200020043602002000200420024103746a22023602082000410c6a200420014103746a3602002000200236020420000b9a0101037f200120012802042000280204200028020022026b22036b22043602040240200341004c0d002004200220031089808080001a200141046a28020021040b2000280200210320002004360200200141046a22042003360200200041046a220328020021022003200128020836020020012002360208200028020821032000200128020c3602082001200336020c200120042802003602000b4b01037f2000280208210120002802042102200041086a21030240034020022001460d012003200141786a22013602000c000b0b024020002802002201450d0020011092808080000b20000b961004097f017e027f017e23808080800041306b2202248080808000200041046a21030240024020014101470d00200310e3808080002802002101200041186a22042004280200417f6a3602000240200310ed80808000418010490d002000410c6a2204280200417c6a28020010928080800020032004280200417c6a10e6808080000b024020014138490d00200110e28080800020016a21010b200141016a2101200041186a280200450d01200310e38080800021000c010b0240200310ed808080000d000240200041146a22012802002204418008490d00200120044180786a360200200041086a2201280200220428020021052001200441046a360200200220053602182003200241186a10ee808080000c010b0240024002400240024002400240024002402000410c6a2802002204200041086a2802006b4102752205200041106a2206280200220720002802046b22014102754f0d0010ef80808000210820072004470d01200041086a22072802002204200041046a2802002201460d02200421010c080b20022001410175410120011b2005200610f080808000210810ef80808000210720082802082201200828020c2205470d0320082802042204200828020022064d0d02200120046b220141027521092004200420066b41027541016a417e6d41027422066a210502402001450d002005200420011093808080001a200841046a28020021040b200841046a200420066a360200200841086a200520094102746a22013602000c030b2000410c6a22072802002201200041106a2802002205470d04200041086a22092802002204200041046a28020022064d0d03200120046b220141027521092004200420066b41027541016a417e6d41027422066a210502402001450d002005200420011093808080001a200041086a28020021040b200041086a200420066a3602002000410c6a200520094102746a22013602000c040b2000410c6a22092802002205200041106a220a28020022064f0d042005200620056b41027541016a41026d41027422096a21010240200520046b2206450d00200120066b2201200420061093808080001a2000410c6a28020021050b200041086a20013602002000410c6a200520096a3602000c050b200241186a200520066b2201410175410120011b22012001410276200841106a28020010f0808080002105200841086a2802002106200841046a28020021010240034020062001460d01200541086a220428020020012802003602002004200428020041046a360200200141046a21010c000b0b2008290200210b200820052902003702002005200b370200200841086a2201290200210b2001200541086a22042902003702002004200b370200200510f1808080001a200128020021010b20012007360200200841086a2209200928020041046a3602002000410c6a2802002106200841106a210c024003402006200041086a280200460d012006417c6a210602400240200841046a2207280200220420082802002201460d00200421010c010b0240200928020022052008410c6a280200220a4f0d002005200a20056b41027541016a41026d410274220d6a21010240200520046b220a450d002001200a6b22012004200a1093808080001a200928020021050b2007200136020020092005200d6a3602000c010b200241186a200a20016b2201410175410120011b2201200141036a410276200c28020010f080808000210a20092802002105200728020021010240034020052001460d01200241186a41086a220428020020012802003602002004200428020041046a360200200141046a21010c000b0b2008290200210b200820022903183702002009290200210e2009200241186a41086a22012903003702002001200e3703002002200b370318200a10f1808080001a200728020021010b2001417c6a200628020036020020072007280200417c6a3602000c000b0b200041046a220128020021042001200828020036020020082004360200200841046a2201280200210420012006360200200041086a20043602002000410c6a2201290200210b2001200841086a22042902003702002004200b370200200810f1808080001a0c040b200241186a200520066b2201410175410120011b22012001410276200041106a10f08080800021052000410c6a2802002106200928020021010240034020062001460d01200541086a220428020020012802003602002004200428020041046a360200200141046a21010c000b0b200041046a2201290200210b200120052902003702002005200b3702002000410c6a2201290200210b2001200541086a22042902003702002004200b370200200510f1808080001a200128020021010b200120083602002007200728020041046a3602000c020b200241186a200620016b2201410175410120011b2201200141036a410276200a10f080808000210520092802002106200041086a28020021010240034020062001460d01200541086a220428020020012802003602002004200428020041046a360200200141046a21010c000b0b200041046a2201290200210b200120052902003702002005200b3702002000410c6a2201290200210b2001200541086a22042902003702002004200b370200200510f1808080001a200041086a28020021010b2001417c6a2008360200200720072802002201417c6a22043602002004280200210420072001360200200220043602182003200241186a10ee808080000b200241186a200310e580808000200228021c4100360200200041186a2100410121010b2000200028020020016a360200200241306a2480808080000b2801017f200028020820002802046b2201410874417f6a410020011b200028021420002802106a6b0b850302067f017e23808080800041206b2202248080808000024020002802082203200028020c2204470d00024020002802042205200028020022064d0d00200320056b220341027521072005200520066b41027541016a417e6d41027422066a210402402003450d002004200520031093808080001a200041046a28020021050b200041046a200520066a360200200041086a200420074102746a22033602000c010b200241086a200420066b2203410175410120031b220320034102762000410c6a10f0808080002104200041086a2802002106200041046a28020021030240034020062003460d01200441086a220528020020032802003602002005200528020041046a360200200341046a21030c000b0b200029020021082000200429020037020020042008370200200041086a220329020021082003200441086a220529020037020020052008370200200410f1808080001a200328020021030b20032001280200360200200041086a2203200328020041046a360200200241206a2480808080000b0b004180201091808080000b7301017f410021042000410036020c200041106a2003360200024002402001450d0020014180808080044f0d01200141027410918080800021040b200020043602002000200420024102746a22023602082000410c6a200420014102746a3602002000200236020420000f0b108880808000000b4b01037f2000280208210120002802042102200041086a21030240034020022001460d0120032001417c6a22013602000c000b0b024020002802002201450d0020011092808080000b20000b5501017f410042003702b08a808000410041003602b88a80800041742100024003402000450d01200041bc8a8080006a4100360200200041046a21000c000b0b4182808080004100418088808000109b808080001a0b0bde0202004180080bbc02000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041bc0a0b14696e697400736f7274006765745f617272617900";

    public static String BINARY = BINARY_0;

    public static final String FUNC_GET_ARRAY = "get_array";

    public static final String FUNC_SORT = "sort";

    protected BubbleSort(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected BubbleSort(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Int64[]> get_array() {
        final WasmFunction function = new WasmFunction(FUNC_GET_ARRAY, Arrays.asList(), Int64[].class);
        return executeRemoteCall(function, Int64[].class);
    }

    public static RemoteCall<BubbleSort> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(BubbleSort.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<BubbleSort> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(BubbleSort.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<BubbleSort> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(BubbleSort.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<BubbleSort> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(BubbleSort.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public RemoteCall<TransactionReceipt> sort(Int64[] arr, Int32 n) {
        final WasmFunction function = new WasmFunction(FUNC_SORT, Arrays.asList(arr,n), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> sort(Int64[] arr, Int32 n, BigInteger vonValue) {
        final WasmFunction function = new WasmFunction(FUNC_SORT, Arrays.asList(arr,n), Void.class);
        return executeRemoteCallTransaction(function, vonValue);
    }

    public static BubbleSort load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new BubbleSort(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static BubbleSort load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new BubbleSort(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
