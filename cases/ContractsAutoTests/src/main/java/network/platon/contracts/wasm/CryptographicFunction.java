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
 * <p>Generated with platon-web3j version 0.13.0.6.
 */
public class CryptographicFunction extends WasmContract {
    private static String BINARY_0 = "0x0061736d01000000013d0b60027f7f0060017f017f60037f7f7f0060000060017f0060027f7f017f60037f7f7f017f60047f7f7f7f017f60047f7f7f7f006000017f60017f017e029f010703656e760c706c61746f6e5f70616e6963000303656e7610706c61746f6e5f65637265636f766572000703656e7610706c61746f6e5f726970656d64313630000203656e760d706c61746f6e5f736861323536000203656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000903656e7610706c61746f6e5f6765745f696e707574000403656e760d706c61746f6e5f72657475726e0000032b2a030304040202000201060301070204000a010005010400060003010000010203030101080500010000000405017001050505030100020608017f0141808b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070f5f5f66756e63735f6f6e5f65786974002006696e766f6b650011090a010041010b04090c0e090aa9392a08001008102610270b3801017f41c408420037020041cc08410036020041742100034020000440200041d0086a4100360200200041046a21000c010b0b4101100a0b0300010b940101027f41d008410136020041d408280200220145044041d40841dc0836020041dc0821010b024041d8082802002202412046044041840210122201450d0120011024220141d40828020036020041d408200136020041d8084100360200410021020b41d808200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41d00841003602000b3c01017f034020034114470440200020036a41003a0000200341016a21030c010b0b200120022802002201200228020420016b41ff0171200010011a0b2c00200041003602082000420037020020004114100d20022802002201200228020420016b200028020010020bf90101067f024020002802042202200028020022056b220320014904402000280208220720026b200120036b22044f04400340200241003a00002000200028020441016a22023602042004417f6a22040d000c030b000b2001200720056b2202410174220520052001491b41ffffffff07200241ffffffff03491b220104402001100f21060b200320066a220321020340200241003a0000200241016a21022004417f6a22040d000b20032000280204200028020022056b22046b2103200441014e044020032005200410101a0b2000200120066a36020820002002360204200020033602000f0b200320014d0d002000200120056a3602040b0b2c00200041003602082000420037020020004120100d20022802002201200228020420016b200028020010030b0b002000410120001b10120bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bed0902097f027e230041d0026b22002400100710042202101222011005200041206a200020012002411c1013220441001014200041206a1015024002402000280224450d00200041206a10150240200028022022032c0000220241004e044020020d010c020b200241807f460d00200241ff0171220141b7014d0440200028022441014d04401000200028022021030b20032d00010d010c020b200141bf014b0d012000280224200241ff017141ca7e6a22014d04401000200028022021030b200120036a2d0000450d010b2000280224450d0020032d000041c001490d010b10000b20004180016a200041206a1016200028028401220541094f044010000b20002802800121030340200504402005417f6a210520033100002009420886842109200341016a21030c010b0b024002402009500d0041800810172009510d0141850810172009510440200041206a10181a2000410036024820004200370340200041e8016a200441011014200041e8016a1015200041b8016a200041e8016a1016200041406b2106024020002802ec014520002802bc01220141204b7245044020002802e8012d000041c001490d010b10000b20004180016a10182001412020014120491b22016b41206a20002802b801200110101a200041286a220220004188016a2207290300370300200041306a220120004190016a2208290300370300200041386a220320004198016a2205290300370300200020002903800122093703502000200937032020004180016a20044102101420004180016a20061019200041d8006a22042002290300370300200041e0006a22022001290300370300200041e8006a2201200329030037030020002000290320370350200041f0006a2006101a21032005200129030037030020082002290300370300200720042903003703002000200029035037038001200041a0016a20004180016a2003100b200041b8016a101b2102200041e0016a200041b0016a2802002201360200200041d8016a200041a8016a290300220a3703004100210520004180026a4100360200200041f8016a4200370300200041f0016a4200370300200020002903a00122093703d001200042003703e80120004190026a200a37030020004198026a2001360200200041a8026a200a370300200041b0026a20013602002000200937038802200020093703a002200041c8026a2001360200200041c0026a200a370300200020093703b802410121040240034020054114460d01200041b8026a20056a200541016a21052d0000450d000b411521040b200020043602e801200041e8016a410472101c20022004101d20004198026a200041b0016a280200220136020020004190026a200041a8016a290300220a370300200020002903a001220937038802200041a8026a200a370300200041b0026a2001360200200041c0026a200a370300200041c8026a2001360200200020093703a002200020093703b802200041f8016a2001360200200041f0016a200a370300200020093703e8012002200041e8016a4114101e220228020c200241106a28020047044010000b200228020020022802041006200228020c22010440200220013602100b200328020022010440200320013602040b20002802402201450d02200020013602440c020b419b081017200951044020044102101f0c020b41b10810172009520d0020044103101f0c010b10000b1020200041d0026a24000b970101047f230041106b220124002001200036020c2000047f41f40a200041086a2202411076220041f40a2802006a220336020041f00a200241f00a28020022026a41076a417871220436020002400240200341107420044d044041f40a200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104101041086a0541000b200141106a24000b730020004200370210200042ffffffff0f370208200020023602042000200136020002402003410871450d002000102820024f0d002003410471044010000c010b200042003702000b02402003411071450d002000102820024d0d0020034104710440100020000f0b200042003702000b20000bc90202077f017e230041106b220324002001280208220520024b0440200341086a2001102c20012003280208200328020c102b36020c20032001102c410021052001027f410020032802002206450d001a410020032802042208200128020c2207490d001a200820072007417f461b210420060b360210200141146a2004360200200141003602080b200141106a210903402001280214210402402005200249044020040d01410021040b200020092802002004411410131a200341106a24000f0b20032001102c41002104027f410020032802002207450d001a410020032802042208200128020c2206490d001a200820066b2104200620076a0b2105200120043602142001200536021020032009410020052004102b102a20012003290300220a3702102001200128020c200a422088a76a36020c2001200128020841016a22053602080c000b000b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0bd40101047f200110212204200128020422024b04401000200128020421020b200128020021052000027f02400240200204404100210120052c00002203417f4a0d01027f200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a21010c010b4101210120050d000c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010b2501017f03402001412046450440200020016a41003a0000200141016a21010c010b0b20000bc70101037f230041206b22022400024002402000280204044020002802002d000041c001490d010b20024100360208200242003703000c010b200241186a2000101620022802182103200241106a200010162002280210210420001021200241003602082002420037030020046a20036b2200450d0020022000102220004101480d002002200228020420032000101020006a3602040b2001280200044020014100360208200142003702000b2001200228020036020020012002290204370204200241206a24000b5a01017f20004200370200200041003602080240200128020420012802006b2202450d002000200210222001280204200128020022026b22014101480d0020002802042002200110101a2000200028020420016a3602040b20000b29002000410036020820004200370200200041001023200041146a41003602002000420037020c20000b860201067f200028020422032000280210220241087641fcffff07716a2101027f200320002802082204460440200041146a210541000c010b2003200028021420026a220541087641fcffff07716a280200200541ff07714102746a2106200041146a21052001280200200241ff07714102746a0b21020340024020022006460440200541003602000340200420036b41027522014103490d022000200341046a22033602040c000b000b200241046a220220012802006b418020470d0120012802042102200141046a21010c010b0b2001417f6a220141014d04402000418004418008200141016b1b3602100b03402003200447044020002004417c6a22043602080c010b0b0b1300200028020820014904402000200110230b0bca0301037f4101210302400240024002402002410146044020012c000022024100480d012000200241ff017110300c040b200241374b0d01200221030b200020034180017341ff017110300c010b2002102d220341b7016a22044180024e044010000b2000200441ff017110302000200028020420036a102e200028020420002802006a417f6a210420022103037f2003047f200420033a0000200341087621032004417f6a21040c010520020b0b21030b20002003102f200028020020002802046a2001200310101a2000200028020420036a3602040b0340024020002802102203200028020c460d00200341786a2802004504401000200028021021030b200341786a22012001280200417f6a220236020020020d002000200136021020004101200028020422042003417c6a28020022016b2203102d220241016a20034138491b220520046a102e200120002802006a220420056a200420031025200341374d0440200028020020016a200341406a3a00000c020b200241f7016a220441ff014d0440200028020020016a20043a00002000280200200120026a6a210403402003450d03200420033a0000200341087621032004417f6a21040c000b000510000c020b000b0b20000bfb0201057f230041f0006b22022400200241003602102002420037030841012103200241d0006a200041011014200241d0006a200241086a1019200241286a2002200241186a200241086a101a22042001110200200241386a101b2105200241e8006a4100360200200241e0006a4200370300200241d8006a420037030020024200370350024020022802282200200228022c2206460d000240200620006b2201410146044020002c0000417f4a0d020c010b20014138490d00200620006b41016a210303402001450d02200341016a2103200141087621010c000b000b200141016a21030b20022003360250200241d0006a410472101c20052003101d200520022802282200200228022c20006b101e220028020c200041106a28020047044010000b200028020020002802041006200028020c22010440200020013602100b2002280228220004402002200036022c0b200428020022000440200420003602040b2002280208220004402002200036020c0b200241f0006a24000b880101037f41d008410136020041d4082802002100034020000440034041d80841d8082802002201417f6a2202360200200141014845044041d0084100360200200020024102746a22004184016a280200200041046a28020011040041d008410136020041d40828020021000c010b0b41d808412036020041d408200028020022003602000c010b0b0bff0201037f200028020445044041000f0b2000101541012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b2001017f20002001100f2202360200200020023602042000200120026a3602080b2f01017f2000280208200149044020011012200028020020002802041010210220002001360208200020023602000b0bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210101a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041808b0436020c41ec0a200028020c41076a417871220036020041f00a200036020041f40a3f003602000b3801017f41e00a420037020041e80a410036020041742100034020000440200041ec0a6a4100360200200041046a21000c010b0b4104100a0b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001029200010216a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b2301017f230041206b22022400200241086a20002001411410131028200241206a24000b2101017f20011021220220012802044b044010000b20002001200110292002102a0b1e01017f03402000044020004108762100200141016a21010c010b0b20010b0f00200020011023200020013602040b1b00200028020420016a220120002802084b04402000200110230b0b250020004101102f200028020020002802046a20013a00002000200028020441016a3602040b0b4a01004180080b43696e69740063616c6c5f706c61746f6e5f65637265636f7665720063616c6c5f706c61746f6e5f726970656d643136300063616c6c5f706c61746f6e5f736861323536";

    public static String BINARY = BINARY_0;

    public static final String FUNC_CALL_PLATON_ECRECOVER = "call_platon_ecrecover";

    public static final String FUNC_CALL_PLATON_RIPEMD160 = "call_platon_ripemd160";

    public static final String FUNC_CALL_PLATON_SHA256 = "call_platon_sha256";

    protected CryptographicFunction(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    protected CryptographicFunction(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }

    public RemoteCall<WasmAddress> call_platon_ecrecover(byte[] hash, byte[] signature) {
        final WasmFunction function = new WasmFunction(FUNC_CALL_PLATON_ECRECOVER, Arrays.asList(hash,signature), WasmAddress.class);
        return executeRemoteCall(function, WasmAddress.class);
    }

    public static RemoteCall<CryptographicFunction> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CryptographicFunction.class, web3j, credentials, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<CryptographicFunction> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CryptographicFunction.class, web3j, transactionManager, contractGasProvider, encodedConstructor, chainId);
    }

    public static RemoteCall<CryptographicFunction> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CryptographicFunction.class, web3j, credentials, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public static RemoteCall<CryptographicFunction> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, BigInteger initialVonValue, Long chainId) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(CryptographicFunction.class, web3j, transactionManager, contractGasProvider, encodedConstructor, initialVonValue, chainId);
    }

    public RemoteCall<byte[]> call_platon_ripemd160(byte[] data) {
        final WasmFunction function = new WasmFunction(FUNC_CALL_PLATON_RIPEMD160, Arrays.asList(data, Void.class), byte[].class);
        return executeRemoteCall(function, byte[].class);
    }

    public RemoteCall<byte[]> call_platon_sha256(byte[] data) {
        final WasmFunction function = new WasmFunction(FUNC_CALL_PLATON_SHA256, Arrays.asList(data, Void.class), byte[].class);
        return executeRemoteCall(function, byte[].class);
    }

    public static CryptographicFunction load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider, Long chainId) {
        return new CryptographicFunction(contractAddress, web3j, credentials, contractGasProvider, chainId);
    }

    public static CryptographicFunction load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider, Long chainId) {
        return new CryptographicFunction(contractAddress, web3j, transactionManager, contractGasProvider, chainId);
    }
}
