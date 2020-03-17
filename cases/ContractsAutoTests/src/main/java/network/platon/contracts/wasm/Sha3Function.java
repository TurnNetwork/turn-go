package network.platon.contracts.wasm;

import com.platon.rlp.datatypes.Uint32;
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
public class Sha3Function extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001490d60017f0060000060047f7f7f7f006000017f60027f7f0060037f7f7f017f60017f017f60027f7f017f60037f7f7f0060037f7e7e017f60027e7e017f60047f7e7e7f0060017f017e026f0503656e760c706c61746f6e5f70616e6963000103656e760b706c61746f6e5f73686133000203656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000303656e7610706c61746f6e5f6765745f696e707574000003656e760d706c61746f6e5f72657475726e0004033332010505060605070001000500000600060105000506060606020408060006040406040404090a06080b00010006010c0400010405017001030305030100020615037f0141d08a040b7f0041d08a040b7f0041cd0a0b075406066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300050b5f5f686561705f6261736503010a5f5f646174615f656e6403020f5f5f66756e63735f6f6e5f65786974001506696e766f6b6500320908010041010b0217300a8d37320800100d102f10360bc60a010b7f2002410f6a210341002104410020026b21052002410e6a2106410120026b21072002410d6a2108410220026b210902400340200020046a210b200120046a210a20022004460d01200a410371450d01200b200a2d00003a00002003417f6a2103200541016a21052006417f6a2106200741016a21072008417f6a2108200941016a2109200441016a21040c000b0b200220046b210c02400240024002400240200b410371220d450d00200c4120490d03200d4101460d01200d4102460d02200d4103470d03200b200120046a28020022063a0000200041016a210c200220046b417f6a21092004210b0240034020094113490d01200c200b6a220a2001200b6a220741046a2802002208411874200641087672360200200a41046a200741086a2802002206411874200841087672360200200a41086a2007410c6a2802002208411874200641087672360200200a410c6a200741106a2802002206411874200841087672360200200b41106a210b200941706a21090c000b0b2002417f6a2005416d2005416d4b1b20036a4170716b20046b210c2001200b6a41016a210a2000200b6a41016a210b0c030b200c210a02400340200a4110490d01200020046a220b200120046a2207290200370200200b41086a200741086a290200370200200441106a2104200a41706a210a0c000b0b02400240200c4108710d00200120046a210a200020046a21040c010b200020046a220b200120046a2204290200370200200441086a210a200b41086a21040b0240200c410471450d002004200a280200360200200a41046a210a200441046a21040b0240200c410271450d002004200a2f00003b0000200441026a2104200a41026a210a0b200c410171450d032004200a2d00003a000020000f0b200b200120046a220a28020022063a0000200b41016a200a41016a2f00003b0000200041036a210c200220046b417d6a21052004210b0240034020054111490d01200c200b6a220a2001200b6a220741046a2802002203410874200641187672360200200a41046a200741086a2802002206410874200341187672360200200a41086a2007410c6a2802002203410874200641187672360200200a410c6a200741106a2802002206410874200341187672360200200b41106a210b200541706a21050c000b0b2002417d6a2009416f2009416f4b1b20086a4170716b20046b210c2001200b6a41036a210a2000200b6a41036a210b0c010b200b200120046a220a28020022083a0000200b41016a200a41016a2d00003a0000200041026a210c200220046b417e6a21052004210b0240034020054112490d01200c200b6a220a2001200b6a220941046a2802002203411074200841107672360200200a41046a200941086a2802002208411074200341107672360200200a41086a2009410c6a2802002203411074200841107672360200200a410c6a200941106a2802002208411074200341107672360200200b41106a210b200541706a21050c000b0b2002417e6a2007416e2007416e4b1b20066a4170716b20046b210c2001200b6a41026a210a2000200b6a41026a210b0b0240200c411071450d00200b200a2d00003a0000200b200a280001360001200b200a290005370005200b200a2f000d3b000d200b200a2d000f3a000f200b41106a210b200a41106a210a0b0240200c410871450d00200b200a290000370000200b41086a210b200a41086a210a0b0240200c410471450d00200b200a280000360000200b41046a210b200a41046a210a0b0240200c410271450d00200b200a2f00003b0000200b41026a210b200a41026a210a0b200c410171450d00200b200a2d00003a00000b20000bfb0202027f017e02402002450d00200020013a0000200020026a2203417f6a20013a000020024103490d00200020013a0002200020013a00012003417d6a20013a00002003417e6a20013a000020024107490d00200020013a00032003417c6a20013a000020024109490d002000410020006b41037122046a2203200141ff017141818284086c22013602002003200220046b417c7122046a2202417c6a200136020020044109490d002003200136020820032001360204200241786a2001360200200241746a200136020020044119490d002003200136021820032001360214200320013602102003200136020c200241706a20013602002002416c6a2001360200200241686a2001360200200241646a20013602002001ad220542208620058421052004200341047141187222016b2102200320016a2101034020024120490d0120012005370300200141186a2005370300200141106a2005370300200141086a2005370300200141206a2101200241606a21020c000b0b20000b3a01017f23808080800041106b220141d08a84800036020c2000200128020c41076a41787122013602042000200136020020003f0036020c20000b120041808880800020004108108a808080000bca0101067f23808080800041106b22032480808080002003200136020c024002400240024002402001450d002000200028020c200241036a410020026b220471220520016a220641107622016a220736020c200020022000280204220820066a6a417f6a2004712202360204200741107420024d0d0120010d020c030b410021000c030b2000410c6a200741016a360200200141016a21010b200140000d001080808080000b20082003410c6a41041086808080001a200820056a21000b200341106a24808080800020000b2e000240418088808000200120006c22004108108a808080002201450d002001410020001087808080001a0b20010b02000b0f004180888080001088808080001a0b0a002000108c808080000bce0301067f024020002001460d000240024002400240200120006b20026b410020024101746b4d0d0020012000734103712103200020014f0d012003450d02200021030c030b2000200120021086808080000f0b024020030d002001417f6a210402400340200020026a2203410371450d012002450d052003417f6a200420026a2d00003a00002002417f6a21020c000b0b2000417c6a21032001417c6a2104034020024104490d01200320026a200420026a2802003602002002417c6a21020c000b0b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b0b200241046a21052002417f7321064100210402400340200120046a2107200020046a2208410371450d0120022004460d03200820072d00003a00002005417f6a2105200641016a2106200441016a21040c000b0b200220046b2101410021030240034020014104490d01200820036a200720036a280200360200200341046a21032001417c6a21010c000b0b200720036a2101200820036a210320022006417c2006417c4b1b20056a417c716b20046b21020b03402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b0b20000b0900200041013602000b0900200041003602000b1d00024020002d0000410171450d002000280208108e808080000b20000b2801017f41002101024003402001410c460d01200020016a4100360200200141046a21010c000b0b0b1d0020004200370200200041086a4100360200200010938080800020000bb60101037f4190888080001090808080004100280294888080002100024003402000450d010240034041004100280298888080002202417f6a22013602988880800020024101480d01200020014102746a22004184016a2802002102200041046a2802002100419088808000109180808000200220001180808080000041908880800010908080800041002802948880800021000c000b0b4100412036029888808000410020002802002200360294888080000c000b0b0bcd0101027f419088808000109080808000024041002802948880800022030d00419c8880800021034100419c88808000360294888080000b0240024041002802988880800022044120470d004184024101108b808080002203450d01410021042003410028029488808000360200410020033602948880800041004100360298888080000b4100200441016a36029888808000200320044102746a22034184016a2001360200200341046a200036020041908880800010918080800041000f0b419088808000109180808000417f0b0f0041a08a8080001092808080001a0b89010020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d00200010998080800020012802044f0d00024020024104710d00200042003702000c010b1080808080000b024002402002411071450d00200010998080800020012802044d0d0020024104710d01200042003702000b20000f0b10808080800020000b3400024002402000280204450d0020002802002c0000417f4c0d0141010f0b41000f0b2000109a808080002000109b808080006a0b280002402000280204450d0020002802002c0000417f4c0d0041000f0b200010a08080800041016a0bc90301047f0240024002402000280204450d00200010a1808080004101210120002802002c00002202417f4c0d010c020b41000f0b0240200241ff0171220141b7014b0d00200141807f6a0f0b024002400240200241ff0171220241bf014b0d000240200041046a22032802002202200141c97e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b200141384f0d010c020b0240200241f7014b0d00200141c07e6a0f0b0240200041046a22032802002202200141897e6a22044b0d00108080808000200328020021020b024020024102490d0020002802002d00010d001080808080000b024020044105490d001080808080000b024020002802002d00010d001080808080000b41002101410021020240034020042002460d012001410874200028020020026a41016a2d0000722101200241016a21020c000b0b20014138490d010b200141ff7d490d010b10808080800020010f0b20010b5102017f017e23808080800041306b220124808080800020012000290200220237031020012002370308200141186a200141086a41141098808080001099808080002100200141306a24808080800020000b6a01037f02400240024020012802002204450d0041002105200320026a200128020422064b0d0120062002490d014100210120062003490d02200620026b20032003417f461b2101200420026a21050c020b410021050b410021010b20002001360204200020053602000b3301017f02402001109b80808000220220012802044d0d001080808080000b200020012001109a808080002002109d808080000bd003020a7f017e23808080800041c0006b220324808080800002402001280208220420024d0d00200341386a2001109e80808000200320032903383703182001200341186a109c8080800036020c200341306a2001109e80808000410021044100210541002106024020032802302207450d00410021054100210620032802342208200128020c2209490d00200820092009417f461b2105200721060b20012006360210200141146a2005360200200141086a41003602000b200141106a2106200141146a21092001410c6a2107200141086a210802400340200420024f0d012009280200450d01200341306a2001109e8080800041002104024002402003280230220a450d00410021052003280234220b2007280200220c490d01200a200c6a2105200b200c6b21040c010b410021050b20092004360200200620053602002003200436022c2003200536022820032003290328370310200341306a20064100200341106a109c80808000109d8080800020062003290330220d37020020072007280200200d422088a76a3602002008200828020041016a22043602000c000b0b20032006290200220d3703202003200d3703082000200341086a41141098808080001a200341c0006a2480808080000b4701017f4100210102402000280204450d00024020002802002d0000220041bf014b0d00200041b801490d01200041c97e6a0f0b200041f801490d00200041897e6a21010b20010b5401017f024020002802040d001080808080000b0240200028020022012d0000418101470d000240200041046a28020041014b0d00108080808000200028020021010b20012c00014100480d001080808080000b0bbc0101047f024002402000280204450d00200010a180808000200028020022012c000022024100480d0120024100470f0b41000f0b410121030240200241807f460d000240200241ff0171220441b7014b0d000240200041046a28020041014b0d00108080808000200028020021010b20012d00014100470f0b41002103200441bf014b0d000240200041046a280200200241ff017141ca7e6a22024b0d00108080808000200028020021010b200120026a2d000041004721030b20030bb60201087f02402001450d002000410c6a2102200041106a2103200041046a21040340200328020022052002280200460d010240200541786a28020020014f0d00108080808000200328020021050b200541786a2206200628020020016b220136020020010d01200320063602002000410120042802002005417c6a28020022016b220510a580808000220741016a20054138491b2206200428020022086a10a680808000200120002802006a220920066a2009200820016b108f808080001a02400240200541374b0d00200028020020016a200541406a3a00000c010b0240200741f7016a220641ff014b0d00200028020020016a20063a00002000280200200720016a6a210103402005450d02200120053a0000200541087621052001417f6a21010c000b0b1080808080000b410121010c000b0b0b21000240200028020420016a220120002802084d0d002000200110a7808080000b0b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b13002000200110a780808000200020013602040b4501017f0240200028020820014f0d0020011089808080002202200028020020002802041086808080001a200010ae80808000200041086a2001360200200020023602000b0b29002000410110a480808000200028020020002802046a20013a00002000200028020441016a3602040bc40101037f02400240024020012002844200510d00200142ff005620024200522002501b0d0120002001a741ff017110a8808080000c020b200041800110a8808080000c010b024002402001200210aa80808000220341374b0d00200020034180017341ff017110a8808080000c010b0240200310ab80808000220441b7016a2205418002490d001080808080000b2000200541ff017110a88080800020002003200410ac808080000b200020012002200310ad808080000b2000410110a38080800020000b3501017f41002102024003402000200184500d0120004208882001423886842100200241016a2102200142088821010c000b0b20020b2501017f41002101024003402000450d0120004108762100200141016a21010c000b0b20010b44002000200028020420026a10a680808000200028020020002802046a417f6a2100024003402001450d01200020013a0000200141087621012000417f6a21000c000b0b0b54002000200028020420036a10a680808000200028020020002802046a417f6a2100024003402001200284500d01200020013c0000200142088820024238868421012000417f6a2100200242088821020c000b0b0b1700024020002802002200450d002000108c808080000b0b240041a08a8080001094808080001a41818080800041004180888080001096808080001a0b0f0041ac8a8080001092808080001a0b7b01027f23808080800041306b22012480808080002001412c6a41002d00bc8a8080003a0000200141002800b88a808000360228200141186a4200370300200141106a42003703002001420037030820014200370300200141286a41052001412010818080800020012802002102200141306a24808080800020020bb60602077f027e23808080800041c0006b2200248080808000108580808000108280808000220110898080800022021083808080002000200136022c2000200236022820002000290328370300200041286a200041086a2000411c1098808080004100109f80808000200041286a10a18080800002400240200041286a10a280808000450d00200028022c450d0020002802282d000041c001490d010b1080808080000b0240200041286a109b808080002203200028022c22044d0d00108080808000200028022c21040b200028022821050240024002400240024002402004450d004100210620052c00002201417f4a0d03200141ff0171220641bf014b0d0141002102200141ff017141b801490d02200641c97e6a21020c020b4101210620050d02410021010c030b41002102200141ff017141f801490d00200641897e6a21020b200241016a21060b41002101200620036a20044b0d0020042003490d004100210220042006490d01200520066a2102200420066b20032003417f461b22014109490d011080808080000c010b410021020b42002107024003402001450d012001417f6a210120074208862002310000842107200241016a21020c000b0b0240024020074200510d00200741bd8a80800010b380808000510d01200741c28a80800010b380808000520d00200041206a10b180808000210441002101200041003602304200210720004200370328200041286a410010b4808080002000413c6a4100360200200042003702344101210202402004418001490d002004ad2108024003402008200784500d0120084208882007423886842108200141016a2101200742088821070c000b0b024020014138490d002001210203402002450d01200141016a2101200241087621020c000b0b200141016a21020b0240200041306a28020020024f0d00200041286a200210b4808080000b200041286a2004ad420010a9808080001a0240200041346a2201280200200041386a280200460d001080808080000b2000280228200028022c108480808000024020012802002201450d00200041386a20013602002001108e808080000b2000280228108c808080000c010b1080808080000b200041c0006a2480808080000b3a01027e42a5c688a1c89ca7f94b21010240034020003000002202500d01200041016a2100200142b383808080207e20028521010c000b0b20010b4401017f0240200028020820014f0d002001108980808000200028020020002802041086808080002102200010b580808000200041086a2001360200200020023602000b0b0d002000280200108c808080000b5501017f410042003702ac8a808000410041003602b48a80800041742100024003402000450d01200041b88a8080006a4100360200200041046a21000c000b0b41828080800041004180888080001096808080001a0b0bdb0202004180080bb8020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041b80a0b150107102030696e69740053686133526573756c7400";

    public static String BINARY = BINARY_0;

    public static final String FUNC_SHA3RESULT = "Sha3Result";

    protected Sha3Function(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected Sha3Function(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Uint32> Sha3Result() {
        final WasmFunction function = new WasmFunction(FUNC_SHA3RESULT, Arrays.asList(), Uint32.class);
        return executeRemoteCall(function, Uint32.class);
    }

    public static RemoteCall<Sha3Function> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(Sha3Function.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<Sha3Function> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(Sha3Function.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<Sha3Function> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(Sha3Function.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static RemoteCall<Sha3Function> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(Sha3Function.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue);
    }

    public static Sha3Function load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new Sha3Function(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static Sha3Function load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new Sha3Function(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
