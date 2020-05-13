package network.platon.contracts.wasm;

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
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 0.7.5.3-SNAPSHOT.
 */
public class Contract_panic extends WasmContract {
    private static String BINARY_0 = "0x0061736d0100000001420c60027f7f0060017f0060017f017f60000060027f7f017f60037f7f7f0060037f7f7f017f60047f7f7f7f0060027f7e0060017f017e6000017f60047f7f7f7f017f02a9010703656e760c706c61746f6e5f70616e6963000303656e760d706c61746f6e5f72657475726e000003656e7617706c61746f6e5f6765745f696e7075745f6c656e677468000a03656e7610706c61746f6e5f6765745f696e707574000103656e7617706c61746f6e5f6765745f73746174655f6c656e677468000403656e7610706c61746f6e5f6765745f7374617465000b03656e7610706c61746f6e5f7365745f73746174650007033c3b03020101000006090102020200000101030900010000010802010400060205030202040501000301030106020202010704000502010000000000080405017001050505030100020608017f0141908b040b073904066d656d6f72790200115f5f7761736d5f63616c6c5f63746f727300070f5f5f66756e63735f6f6e5f65786974002d06696e766f6b650017090a010041010b040b0c0a300aa6433b100041d00810081a4103102e1026102f0b190020004200370200200041086a41003602002000100920000b2201017f03402001410c470440200020016a4100360200200141046a21010c010b0b0b070041d008102b0b0c002000200141186a10291a0b0c002000200141406b10291a0b3401017f230041106b220324002003200236020c200320013602082003200329030837030020002003411c1031200341106a24000b3901027e42a5c688a1c89ca7f94b210103402000300000220250450440200041016a2100200142b383808080207e20028521010c010b0b20010bb00101047f23004190016b22012400200141086a200028020422034101756a210220002802002100200141086a1010200141d8006a20022003410171047f200228020020006a2802000520000b110000200141f8006a10112200200141d8006a101210132000200141e8006a200141d8006a1029220210142002102b200028020c200041106a28020047044010000b20002802002000280204100120001015200141d8006a102b101620014190016a24000b6b01027f200010082102200042afb59bdd9e8485b9f800370310200041106a200041186a10082201102145044020012002102c0b200041286a10082102200041386a220142ecae96e4b694b9d9d0003703002001200041406b10082201102145044020012002102c0b20000b2900200041003602082000420037020020004100101c200041146a41003602002000420037020c20000bbc0101047f230041306b22012400200141286a4100360200200141206a4200370300200141186a42003703002001420037031041012102024002400240200120001029220328020420032d00002200410176200041017122041b220041014d0440200041016b0d032003280208200341016a20041b2c0000417f4c0d010c030b200041374b0d010b200041016a21020c010b2000101f20006a41016a21020b200120023602102003102b200141106a4104721020200141306a240020020b13002000280208200149044020002001101c0b0b5201037f230041106b2202240020022001280208200141016a20012d0000220341017122041b36020820022001280204200341017620041b36020c20022002290308370300200020021040200241106a24000b1c01017f200028020c22010440200041106a20013602000b2000101d0b3301017f200041386a200041406b220110222001102b200041286a102b200041106a200041186a220110222001102b2000102b0ba50302067f017e230041b0016b22002400100710022201102722021003200041d0006a200041106a20022001100d22014100103902400240200041d0006a10182206500d00418008100e2006510440200041d0006a101010160c020b418508100e2006510440200041286a1008210220004200370338200041d0006a200141011039200041d0006a20021019200041d0006a2001410210392000200041d0006a1018370338200041d0006a10102101200041406b20021029210320002903381a200141186a200041a0016a2003102922042205102c200141406b2005102c2004102b2003102b200110162002102b0c020b419408100e2006510440200041a0016a10082102200041d0006a200141011039200041d0006a20021019200041d0006a1010220141186a200041286a200210292203102c2003102b200110162002102b0c020b41a708100e20065104402000410036025420004101360250200020002903503703002000100f0c020b41ba08100e2006520d00200041003602542000410236025020002000290350370308200041086a100f0c010b10000b102d200041b0016a24000b850102027f017e230041106b2201240020001035024002402000103a450d002000280204450d0020002802002d000041c001490d010b10000b200141086a2000101b200128020c220041094f044010000b200128020821020340200004402000417f6a210020023100002003420886842103200241016a21020c010b0b200141106a240020030b8c0301057f230041206b220224000240024002402000280204044020002802002d000041c001490d010b200241086a10081a0c010b200241186a2000101b2000103421030240024002400240200228021822000440200228021c220420034f0d010b41002100200241106a410036020020024200370308410021030c010b200241106a410036020020024200370308200420032003417f461b220341704f0d04200020036a21052003410a4b0d010b200220034101743a0008200241086a41017221040c010b200341106a4170712206102821042002200336020c20022006410172360208200220043602100b034020002005470440200420002d00003a0000200441016a2104200041016a21000c010b0b200441003a00000b024020012d0000410171450440200141003b01000c010b200128020841003a00002001410036020420012d0000410171450d0020012802081a200141003602000b20012002290308370200200141086a200241106a280200360200200241086a1009200241086a102b200241206a24000f0b000b1501017f200028020022010440200020013602040b0bd60101047f200110342204200128020422024b04401000200128020421020b20012802002105027f027f41002002450d001a410020052c00002203417f4a0d011a200341ff0171220141bf014d04404100200341ff017141b801490d011a200141c97e6a0c010b4100200341ff017141f801490d001a200141897e6a0b41016a0b21012000027f02402005450440410021030c010b410021032002200149200120046a20024b720d00410020022004490d011a200120056a2103200220016b20042004417f461b0c010b41000b360204200020033602000b3401017f200028020820014904402001102722022000280200200028020410231a2000101d20002001360208200020023602000b0b080020002802001a0b08002000200110410b1e01017f03402000044020004108762100200141016a21010c010b0b20010bc40201067f200028020422012000280210220341087641fcffff07716a2102027f200120002802082205460440200041146a210441000c010b2001200028021420036a220441087641fcffff07716a280200200441ff07714102746a2106200041146a21042002280200200341ff07714102746a0b21030340024020032006460440200441003602000340200520016b41027522024103490d0220012802001a2000200028020441046a2201360204200028020821050c000b000b200341046a220320022802006b418020470d0120022802042103200241046a21020c010b0b2002417f6a220241014d04402000418004418008200241016b1b3602100b03402001200547044020012802001a200141046a21010c010b0b200028020821012000280204210203402001200247044020002001417c6a22013602080c010b0b20002802001a0bb70201097f230041406a22022400200241286a101122042000290300101e200428020c200441106a28020047044010000b02400240200428020022092004280204220a10042203450440410021030c010b20024100360220200242003703182003417f4c0d01200310282105410021000340200020056a41003a00002003200041016a2200470d000b200020056a21072005200228021c200228021822066b22086b2100200841014e044020002006200810231a200228021821060b2002200320056a3602202002200736021c2002200036021802402009200a20060440200228021c2107200228021821000b2000200720006b1005417f460440410021030c010b20022002280218220041016a200228021c2000417f736a100d200110190b200241186a101a0b20041015200241406b240020030f0b000bed03020a7f027e230041e0006b22032400200341286a10112106200341d8006a4100360200200341d0006a4200370300200341c8006a420037030020034200370340410121042000290300220d4280015a04400340200c200d8450450440200c423886200d42088884210d200541016a2105200c420888210c0c010b0b200541384f047f2005101f20056a0520050b41016a21040b20032004360240200341406b410472102020062004101320062000290300101e200628020c200641106a28020047044010000b200628020421082006280200200341406b1011210220011012210a41011028220441fe013a0000200320043602182003200441016a22003602202003200036021c200228020c200241106a280200470440100020032802182104200328021c21000b2004210520022802042207200020046b22006a220b20022802084b04402002200b101c20022802042107200328021821050b200228020020076a2004200010231a2002200228020420006a3602042002200328021c200a6a20056b10132002200341086a20011029220010142000102b0240200228020c2002280210460440200228020021040c010b100020022802002104200228020c2002280210460d0010000b2008200420022802041006200341186a101a2002101520061015200341e0006a24000bf80801067f0340200020046a2105200120046a220341037145200220044672450440200520032d00003a0000200441016a21040c010b0b200220046b210602402005410371220845044003402006411049450440200020046a2202200120046a2203290200370200200241086a200341086a290200370200200441106a2104200641706a21060c010b0b027f2006410871450440200120046a2103200020046a0c010b200020046a2202200120046a2201290200370200200141086a2103200241086a0b21042006410471044020042003280200360200200341046a2103200441046a21040b20064102710440200420032f00003b0000200341026a2103200441026a21040b2006410171450d01200420032d00003a000020000f0b024020064120490d002008417f6a220841024b0d00024002400240024002400240200841016b0e020102000b2005200120046a220628020022033a0000200541016a200641016a2f00003b0000200041036a2108200220046b417d6a2106034020064111490d03200420086a2202200120046a220541046a2802002207410874200341187672360200200241046a200541086a2802002203410874200741187672360200200241086a2005410c6a28020022074108742003411876723602002002410c6a200541106a2802002203410874200741187672360200200441106a2104200641706a21060c000b000b2005200120046a220628020022033a0000200541016a200641016a2d00003a0000200041026a2108200220046b417e6a2106034020064112490d03200420086a2202200120046a220541046a2802002207411074200341107672360200200241046a200541086a2802002203411074200741107672360200200241086a2005410c6a28020022074110742003411076723602002002410c6a200541106a2802002203411074200741107672360200200441106a2104200641706a21060c000b000b2005200120046a28020022033a0000200041016a21082004417f7320026a2106034020064113490d03200420086a2202200120046a220541046a2802002207411874200341087672360200200241046a200541086a2802002203411874200741087672360200200241086a2005410c6a28020022074118742003410876723602002002410c6a200541106a2802002203411874200741087672360200200441106a2104200641706a21060c000b000b200120046a41036a2103200020046a41036a21050c020b200120046a41026a2103200020046a41026a21050c010b200120046a41016a2103200020046a41016a21050b20064110710440200520032d00003a00002005200328000136000120052003290005370005200520032f000d3b000d200520032d000f3a000f200541106a2105200341106a21030b2006410871044020052003290000370000200541086a2105200341086a21030b2006410471044020052003280000360000200541046a2105200341046a21030b20064102710440200520032f00003b0000200541026a2105200341026a21030b2006410171450d00200520032d00003a00000b20000bc90201037f200041003a000020004184026a2201417f6a41003a0000200041003a0002200041003a00012001417d6a41003a00002001417e6a41003a0000200041003a00032001417c6a41003a00002000410020006b41037122026a22014100360200200141840220026b417c7122036a2202417c6a4100360200024020034109490d002001410036020820014100360204200241786a4100360200200241746a410036020020034119490d002001410036021820014100360214200141003602102001410036020c200241706a41003602002002416c6a4100360200200241686a4100360200200241646a41003602002003200141047141187222036b2102200120036a2101034020024120490d0120014200370300200141186a4200370300200141106a4200370300200141086a4200370300200141206a2101200241606a21020c000b000b20000b8d0301037f024020002001460d00200120006b20026b410020024101746b4d044020002001200210231a0c010b20002001734103712103027f024020002001490440200020030d021a410021030340200120036a2104200020036a2205410371450440200220036b210241002103034020024104490d04200320056a200320046a280200360200200341046a21032002417c6a21020c000b000b20022003460d04200520042d00003a0000200341016a21030c000b000b024020030d002001417f6a21030340200020026a22044103714504402001417c6a21032000417c6a2104034020024104490d03200220046a200220036a2802003602002002417c6a21020c000b000b2002450d042004417f6a200220036a2d00003a00002002417f6a21020c000b000b2001417f6a210103402002450d03200020026a417f6a200120026a2d00003a00002002417f6a21020c000b000b200320046a2101200320056a0b210303402002450d01200320012d00003a00002002417f6a2102200341016a2103200141016a21010c000b000b0b3501017f230041106b220041908b0436020c41f80a200028020c41076a417871220036020041fc0a200036020041800b3f003602000b970101047f230041106b220124002001200036020c2000047f41800b200041086a2202411076220041800b2802006a220336020041fc0a200241fc0a28020022026a41076a417871220436020002400240200341107420044d044041800b200341016a360200200041016a21000c010b2000450d010b200040000d0010000b20022001410c6a4104102341086a0541000b200141106a24000b0b002000410120001b10270ba10101037f20004200370200200041086a2202410036020020012d0000410171450440200020012902003702002002200141086a28020036020020000f0b20012802082103024020012802042201410a4d0440200020014101743a0000200041016a21020c010b200141106a4170712204102821022000200136020420002004410172360200200020023602080b200220032001102a200120026a41003a000020000b10002002044020002001200210231a0b0b130020002d0000410171044020002802081a0b0ba10201047f20002001470440200128020420012d00002202410176200241017122031b2102200141016a21052001280208410a2101200520031b210520002d000022034101712204044020002802002203417e71417f6a21010b200220014d0440027f2004044020002802080c010b200041016a0b21012002044020012005200210250b200120026a41003a000020002d00004101710440200020023602040f0b200020024101743a00000f0b027f2003410171044020002802080c010b41000b1a416f2103200141e6ffffff074d0440410b20014101742203200220022003491b220341106a4170712003410b491b21030b20031028220420052002102a200020023602042000200436020820002003410172360200200220046a41003a00000b0b880101037f41dc08410136020041e0082802002100034020000440034041e40841e4082802002201417f6a2202360200200141014845044041dc084100360200200020024102746a22004184016a280200200041046a28020011010041dc08410136020041e00828020021000c010b0b41e408412036020041e008200028020022003602000c010b0b0b940101027f41dc08410136020041e008280200220145044041e00841e80836020041e80821010b024041e4082802002202412046044041840210272201450d0120011024220141e00828020036020041e008200136020041e4084100360200410021020b41e408200241016a360200200120024102746a22014184016a4100360200200141046a20003602000b41dc0841003602000b3801017f41ec0a420037020041f40a410036020041742100034020000440200041f80a6a4100360200200041046a21000c010b0b4104102e0b070041ec0a102b0b750020004200370210200042ffffffff0f3702082000200129020037020002402002410871450d002000103220012802044f0d002002410471044010000c010b200042003702000b02402002411071450d002000103220012802044d0d0020024104710440100020000f0b200042003702000b20000b2e01017f200028020445044041000f0b4101210120002802002c0000417f4c047f20001033200010346a0541010b0b5b00027f027f41002000280204450d001a410020002802002c0000417f4a0d011a20002802002d0000220041bf014d04404100200041b801490d011a200041c97e6a0c010b4100200041f801490d001a200041897e6a0b41016a0b0bff0201037f200028020445044041000f0b2000103541012102024020002802002c00002201417f4a0d00200141ff0171220341b7014d0440200341807f6a0f0b02400240200141ff0171220141bf014d0440024020002802042201200341c97e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241b7012101034020012003460440200241384f0d030c0405200028020020016a41ca7e6a2d00002002410874722102200141016a21010c010b000b000b200141f7014d0440200341c07e6a0f0b024020002802042201200341897e6a22024d047f100020002802040520010b4102490d0020002802002d00010d0010000b200241054f044010000b20002802002d000145044010000b4100210241f701210103402001200346044020024138490d0305200028020020016a418a7e6a2d00002002410874722102200141016a21010c010b0b0b200241ff7d490d010b10000b20020b4101017f200028020445044010000b0240200028020022012d0000418101470d00200028020441014d047f100020002802000520010b2c00014100480d0010000b0b5a01027f2000027f0240200128020022054504400c010b200220036a200128020422014b2001200249720d00410020012003490d011a200220056a2104200120026b20032003417f461b0c010b41000b360204200020043602000b3c01017f230041306b22022400200220013602142002200036021020022002290310370308200241186a200241086a411410311032200241306a24000b2101017f20011034220220012802044b044010000b2000200120011033200210360bd60202077f017e230041206b220324002001280208220420024b0440200341186a2001103820012003280218200328021c103736020c200341106a20011038410021042001027f410020032802102206450d001a410020032802142208200128020c2207490d001a200820072007417f461b210520060b360210200141146a2005360200200141003602080b200141106a210903400240200420024f0d002001280214450d00200341106a2001103841002104027f410020032802102207450d001a410020032802142208200128020c2206490d001a200820066b2104200620076a0b21052001200436021420012005360210200341106a20094100200520041037103620012003290310220a3702102001200128020c200a422088a76a36020c2001200128020841016a22043602080c010b0b20032009290200220a3703082003200a37030020002003411410311a200341206a24000b980101037f200028020445044041000f0b20001035200028020022022c0000220141004e044020014100470f0b027f4101200141807f460d001a200141ff0171220341b7014d0440200028020441014d047f100020002802000520020b2d00014100470f0b4100200341bf014b0d001a2000280204200141ff017141ca7e6a22014d047f100020002802000520020b20016a2d00004100470b0bf80101057f0340024020002802102201200028020c460d00200141786a28020041014904401000200028021021010b200141786a2202200228020041016b220436020020040d002000200236021020004101200028020422032001417c6a28020022026b2201101f220441016a20014138491b220520036a103c200220002802006a220320056a2003200110250240200141374d0440200028020020026a200141406a3a00000c010b200441f7016a220341ff014d0440200028020020026a20033a00002000280200200220046a6a210203402001450d02200220013a0000200141087621012002417f6a21020c000b000b10000b0c010b0b0b0f0020002001103d200020013602040b2f01017f2000280208200149044020011027200028020020002802041023210220002001360208200020023602000b0b1b00200028020420016a220120002802084b044020002001103d0b0b250020004101103e200028020020002802046a20013a00002000200028020441016a3602040be70101037f2001280200210441012102024002400240024020012802042201410146044020042c000022014100480d012000200141ff0171103f0c040b200141374b0d01200121020b200020024180017341ff0171103f0c010b2001101f220241b7016a22034180024e044010000b2000200341ff0171103f2000200028020420026a103c200028020420002802006a417f6a210320012102037f2002047f200320023a0000200241087621022003417f6a21030c010520010b0b21020b20002002103e200028020020002802046a2004200210231a2000200028020420026a3602040b2000103b0bb80202037f037e024020015004402000418001103f0c010b20014280015a044020012105034020052006845045044020064238862005420888842105200241016a2102200642088821060c010b0b0240200241384f04402002210303402003044020034108762103200441016a21040c010b0b200441c9004f044010000b2000200441b77f6a41ff0171103f2000200028020420046a103c200028020420002802006a417f6a21042002210303402003450d02200420033a0000200341087621032004417f6a21040c000b000b200020024180017341ff0171103f0b2000200028020420026a103c200028020420002802006a417f6a210203402001200784500d02200220013c0000200742388620014208888421012002417f6a2102200742088821070c000b000b20002001a741ff0171103f0b2000103b0b0b5401004180080b4d696e69740070616e69635f636f6e7472616374007365745f737472696e675f73746f72616765006765745f737472696e675f73746f72616765006765745f737472696e675f73746f7261676531";

    private static String BINARY = BINARY_0;

    public static final String FUNC_PANIC_CONTRACT = "panic_contract";

    public static final String FUNC_SET_STRING_STORAGE = "set_string_storage";

    public static final String FUNC_GET_STRING_STORAGE = "get_string_storage";

    public static final String FUNC_GET_STRING_STORAGE1 = "get_string_storage1";

    protected Contract_panic(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    protected Contract_panic(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static RemoteCall<Contract_panic> deploy(Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(Contract_panic.class, web3j, credentials, contractGasProvider, encodedConstructor);
    }

    public static RemoteCall<Contract_panic> deploy(Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        String encodedConstructor = WasmFunctionEncoder.encodeConstructor(BINARY, Arrays.asList());
        return deployRemoteCall(Contract_panic.class, web3j, transactionManager, contractGasProvider, encodedConstructor);
    }

    public RemoteCall<TransactionReceipt> panic_contract(String name, Long value) {
        final WasmFunction function = new WasmFunction(FUNC_PANIC_CONTRACT, Arrays.asList(name,value), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> set_string_storage(String name) {
        final WasmFunction function = new WasmFunction(FUNC_SET_STRING_STORAGE, Arrays.asList(name), Void.class);
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<String> get_string_storage() {
        final WasmFunction function = new WasmFunction(FUNC_GET_STRING_STORAGE, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public RemoteCall<String> get_string_storage1() {
        final WasmFunction function = new WasmFunction(FUNC_GET_STRING_STORAGE1, Arrays.asList(), String.class);
        return executeRemoteCall(function, String.class);
    }

    public static Contract_panic load(String contractAddress, Web3j web3j, Credentials credentials, GasProvider contractGasProvider) {
        return new Contract_panic(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static Contract_panic load(String contractAddress, Web3j web3j, TransactionManager transactionManager, GasProvider contractGasProvider) {
        return new Contract_panic(contractAddress, web3j, transactionManager, contractGasProvider);
    }
}
