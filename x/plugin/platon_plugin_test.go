package plugin_test

import (
	"errors"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	cvm "github.com/PlatONnetwork/PlatON-Go/common/vm"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/core/snapshotdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/ethdb"
	"github.com/PlatONnetwork/PlatON-Go/p2p/discover"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/x/gov"
	"github.com/PlatONnetwork/PlatON-Go/x/plugin"
	"github.com/PlatONnetwork/PlatON-Go/x/restricting"
	"github.com/PlatONnetwork/PlatON-Go/x/staking"
	"github.com/PlatONnetwork/PlatON-Go/x/xcom"
	"github.com/PlatONnetwork/PlatON-Go/x/xutil"
	"math/big"
	"testing"
)



var (
	nodeIdArr = []discover.NodeID{
		discover.MustHexID("0x1f3a8672348ff6b789e416762ad53e69063138b8eb4d8780101658f24b2369f1a8e09499226b467d8bc0c4e03e1dc903df857eeb3c67733d21b6aaee2840e429"),
		discover.MustHexID("0xa6ef31a2006f55f5039e23ccccef343e735d56699bde947cfe253d441f5f291561640a8e2bbaf8a85a8a367b939efcef6f80ae28d2bd3d0b21bdac01c3aa6f2f"),
		discover.MustHexID("0xc7fc34d6d8b3d894a35895aaf2f788ed445e03b7673f7ce820aa6fdc02908eeab6982b7eb97e983cc708bcec093b3bc512b0b1fbf668e6ab94cd91f2d642e591"),
		discover.MustHexID("0x97e424be5e58bfd4533303f8f515211599fd4ffe208646f7bfdf27885e50b6dd85d957587180988e76ae77b4b6563820a27b16885419e5ba6f575f19f6cb36b0"),

		discover.MustHexID("0x77adbb9cec2eeb02081de51a2f99e570552e8d879ad329877d073ffaa62609e37008c0de584e1eea982fe9cfea5e622c614971f50d46185d4cc45cfe7c98a575"),
		discover.MustHexID("0x0987d9a181dfb7802dbc2ae45b6b6c7d835a3c50eb1c3f462cdd6c5517e75156e49ac886d5fddb843e92b75bd1442bf47464d84d006238026cb151280620cd48"),
		discover.MustHexID("0x9b033b576400dcc936d2cf62ac5e38a6cff14105f903feb052a59386f70588e5db7a2356329a9ece198dd407bd976a27e0c1f89dfae436e1a00b00560be85b70"),
		discover.MustHexID("0xd5a4d4f85404f92e5027dc915446bb4142a03f96cb3955ed39c742cc0960b927a2c274cdbef1a91efe29cd7fb496c4cdf1785cb666a7f0a232f453862864e734"),
		discover.MustHexID("0x374cf83c15f69ba2c9cad810d724a9906fb91217c9b8cc8f19bfd36c65b6bbb91fc3d65f3a30c7ea3e99439a74db3eb7840a691655ff68188bcdbd3231f01fd0"),
		discover.MustHexID("0x74006a2f25cea77122d3f729db958cb1cb89aec07e598af2ec97b7890d7e2fc9d60f1a9e5c41b9d2d327d29212b26a4149d67362cc8da89002eaf4baa9e1c7f5"),
		discover.MustHexID("0x2f8abfa60df6f6cac1fd2bb1f33db8f2a2953f186912565aa3b627f429825f9ae85b48f27213fb618586c03a7824e9f9c30e10f6ef23d26434ae8ed1f654d214"),
		discover.MustHexID("0xa268f2bd2399360c422752b9dff2a5875051d9ffa210aa303bd67b61d923b4aa08c91758b51d490888ac31048a4f0ee89bec44b8a9fe20dfef66e6f648433383"),
		discover.MustHexID("0x1bf795bcdc327866be93f9d6b2087fa9ed7fb68dadade1631e842a27b471c3e76004ded8df6152fd2a3c938182ff31477797e8ee7328f1fa5c12884dc8550d37"),
		discover.MustHexID("0x757106dfb4c1a17647a2f7817137d4cafd87d91d26dabc5047d6ee9eb00eac4ce030221aad8d6eaa20ee46d57f656dda2d2909113f1f566a483138d54547b3ec"),
		discover.MustHexID("0x7314b2d4c55b1769e76c879b2952be4da8ab0cc837319441a0f105f57833ce156dd36cb4ed37501dddfb6e1c336da0e1bb29391ba589ba0ce9ca3ff5a8520598"),
		discover.MustHexID("0x5bc1eef7aa549a50969c4dff7af626d000da5509dcd2d61f99dbbe88aa663d6c9af9bf9d1a070be90ba13add0b76b6bddd3a427bca359ae002cd9c06f66b9044"),
		discover.MustHexID("0xe81f25d15a09b3e8d46a4d755a019dfcb93d1e2af3ecf0bc64d94a57078a3d7b48f2c0b67a66fb1a34a69d40fd852b6bac9aecada8873795d6bbceb247288a88"),
		discover.MustHexID("0x712ff45cb8308b4d370e3f8c74fefc261372bdfe32f1c14175b867f6101c5eb05657ed7c7df1b0338bc08716cb9d8f7409df5a1ce7b2d1367a515774d002aa7b"),
		discover.MustHexID("0x5ccdd94a4369a0a9fe911601007d9da35e8d14b4abdbf360dd0727726fa7089f1f3c4b8363e00e2e574eb421d20273f1011de771f17759eed5d53b3807359665"),
		discover.MustHexID("0xf20a258de956c632c5f625bf62aa87a6e42d79bac4c9a05a2ff9b9c99af87e7619912ca2b2b6e023f178aa360d5aae20515be1d1c7a22dad28d54bb002327f9d"),
		discover.MustHexID("0x3f26dadae44317ace8ffef3638906e593225a3621d09c7a8afc1635a20a5f14538e0c6923d92a1799d47a2901fa319ab5b6225d7c95f333ee6d9d7d365a0a04c"),
		discover.MustHexID("0xa607b602735e50d2fbba7bd569df0392b4c8aa95835b502caf6d9f90fa1fd5c452c2a3ee71229161f990b175ed46c6e1e1ec0bcc43af9795f5e927597d7930d8"),
		discover.MustHexID("0x0f5c12e8431243dc49c50b694c16b4d526f75673b17afacb5e79efd7bbbd1b41733283c9a547a4c215ca652e91dbb614ee21b7521cb179cf7b0af12f7049b573"),
		discover.MustHexID("0x6048883096c3cf3f31a9be302650a7fdeb089a0096f41600ea8be49a3cb62b649dd42d3a20743c0f8e547fdf688f353204c711711c751fe2fa2d6b3c5886dcda"),
		discover.MustHexID("0x7bae841405067598bf65e7260ca693a964316e752249c4970085c805dbee738fdb41fc434e96e2b65e8bf1db2f52f05d9300d04c1e6129c26cb5d0f214b49968"),

	}
	addrArr = []common.Address{
		common.HexToAddress("0x740ce31b3fac20dac379db243021a51e80aadd24"),
		common.HexToAddress("0x740ce31b3fac20dac379db243021a51e80444555"),
		common.HexToAddress("0x740ce31b3fac20dac379db243021a51e80eeda12"),
		common.HexToAddress("0x740ce31b3fac20dac379db243021a51e80522345"),

		common.HexToAddress("0xf8136ba2aeDa08BD20239B85Fe0ecB53959605A4"),
		common.HexToAddress("0x50c78829980342444A9eC7195188c8bbaD7F059F"),
		common.HexToAddress("0xA4dB304A178B233E30350aEF6EA10efEab3E39a0"),
		common.HexToAddress("0x687DE3a2c61d93A8e36a702d063596aE68B6b76C"),
		common.HexToAddress("0xe4504BB003D4FF18BA239D863Da0C3b4e5a64015"),
		common.HexToAddress("0xf269718398D6Dfa4505975ec288f6Fedf63446c2"),
		common.HexToAddress("0x025E464b5087ce804Be5B9217d3Ca5c5D9666a8C"),
		common.HexToAddress("0x796573b74F3e585feb75DAf4899909b18011cf9e"),
		common.HexToAddress("0xEAAa15641C357389e9a51fD3c78E34c7035300B3"),
		common.HexToAddress("0x002754FE71b8140fDD84fF34E4D42c1FF7Ac6FB3"),
		common.HexToAddress("0x92BF4dcFfA87F00863Bf4Bf15B7a075B8B82FAa3"),
		common.HexToAddress("0xB02D72F883895575466d37F2A38C11FC061b7D2a"),
		common.HexToAddress("0x12e9dfC6262E189af6e09b18F34C5132bFa2D721"),
		common.HexToAddress("0x492F766bc09028D20B488db6e28a5600B5966Ff8"),
		common.HexToAddress("0x3B1b6a7942f9d70221F584D30C6309BEA12d88ab"),
		common.HexToAddress("0xfed6Ebb71f0685a8901136303F6C0C4d370D90bC"),
		common.HexToAddress("0xb9D0D6f843B8948C1C8f48Dfe8aB12B5dEcaDDAC"),
		common.HexToAddress("0x6A6975e605c5968db4aaF87295E05f611396050E"),
		common.HexToAddress("0xe4a22694827bFa617bF039c937403190477934bF"),
		common.HexToAddress("0x3571089Dc0BE9c992cA20AB3AD91FA98808638eA"),
		common.HexToAddress("0x2540c09C69DA41cB66BC78A5121A7E8FDc892Ac5"),
	}


	blockNumber = big.NewInt(1)
	blockHash = common.HexToHash("9d4fb5346abcf593ad80a0d3d5a371b22c962418ad34189d5b1b39065668d663")

	blockNumber2 = big.NewInt(2)
	blockHash2 = common.HexToHash("c95876b92443d652d7eb7d7a9c0e2c58a95e934c0c1197978c5445180cc60980")

	lastBlockNumber uint64
 	lastBlockHash common.Hash
	lastHeader types.Header

	sender = common.HexToAddress("0xeef233120ce31b3fac20dac379db243021a5234")

	sndb = snapshotdb.Instance()

	sender_balance, _ = new(big.Int).SetString("9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", 10)



	txHashArr = []common.Hash{
		common.HexToHash("0x00000000000000000000000000000000000000886d5ba2d3dfb2e2f6a1814f22"),
		common.HexToHash("0x000000000000000000000000000000005249b59609286f2fa91a2abc8555e887"),
		common.HexToHash("0x000000008dba388834e2515c4d9ccb02a48bae177e73959330e55067211c2456"),
		common.HexToHash("0x0000000000000000000000000000000000009a715a765a72b8a289156f9543c9"),
		common.HexToHash("0x0000e1b4a5508c11772b61f463657585c33b577019e4a23bd359c018a4e306d1"),
		common.HexToHash("0x00fd854f940e2d2af8e74c33e640ea6f75c1d9ee49b816b8a4647611d0c91863"),
		common.HexToHash("0x0000000000001038575739a53385cfe42321585a56050e18f8ea2b3e8dc21966"),
		common.HexToHash("0x0000000000000000000000000000000000000048f3b312dc8d081e1186abe8c2"),
		common.HexToHash("0x000000000000000000000000f5bd37579e7ca954eba8fbe7a65646250e92ab7d"),
		common.HexToHash("0x00000000000000000000000000000000000000001d65a5a69fed6ddb0cb58dff"),
		common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000d2"),
		common.HexToHash("0x0000000000000000000000000000000000000000000000000000f2e8b2706c9e"),
		common.HexToHash("0x00000000000000000000000000e22a393898aac376b079e0894e8e2be6024d03"),
		common.HexToHash("0x000000000000000000000000000000000000000000000000483570dd0679860a"),
		common.HexToHash("0x000000000000000000000000000000000000007fc9e1dc435b5d0064ac50fd4e"),
		common.HexToHash("0x00000000000000000000000000cbeb8f4d51969d7eb70a4f6e8505950d870df7"),
		common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000b4"),
		common.HexToHash("0x000000008fd2abdf28d87efb2c7fa2d37618c8dba97059376d6a58007bee3d8b"),
		common.HexToHash("0x0000000000000000000000003566f3a0adf49d90e610ef3d3548b5a72b1fe199"),
		common.HexToHash("0x00000000000054fa3d19eb57e98aa1dd69d216722054d8539ede4b89c5b77ee9"),

	}


	initProcessVersion = uint32(1<<16 | 0<<8 | 0) // 65536

	balanceStr = []string {

		"9000000000000000000000000",
		"60000000000000000000000000",
		"1300000000000000000000000",
		"1100000000000000000000000",
		"1000000000000000000000000",
		"4879000000000000000000000",
		"1800000000000000000000000",
		"1000000000000000000000000",
		"1000000000000000000000000",
		"70000000000000000000000000",
		"5550000000000000000000000",
		"9000000000000000000000000",
		"60000000000000000000000000",
		"1300000000000000000000000",
		"1100000000000000000000000",
		"1000000000000000000000000",
		"4879000000000000000000000",
		"1800000000000000000000000",
		"1000000000000000000000000",
		"1000000000000000000000000",
		"70000000000000000000000000",
		"5550000000000000000000000",
		"1000000000000000000000000",
		"70000000000000000000000000",
		"5550000000000000000000000",
	}


	nodeNameArr = []string {
		"PlatON",
		"Gavin",
		"Emma",
		"Kally",
		"Juzhen",
		"Baidu",
		"Alibaba",
		"Tencent",
		"ming",
		"hong",
		"gang",
		"guang",
		"hua",
		"PlatON_2",
		"Gavin_2",
		"Emma_2",
		"Kally_2",
		"Juzhen_2",
		"Baidu_2",
		"Alibaba_2",
		"Tencent_2",
		"ming_2",
		"hong_2",
		"gang_2",
		"guang_2",
	}

	chaList = []string{"A", "a", "B", "b", "C", "c", "D", "d", "E", "e", "F", "f", "G", "g", "H", "h", "J", "j", "K", "k", "M", "m",
						"N", "n", "P", "p", "Q", "q", "R", "r", "S", "s", "T", "t", "U", "u", "V", "v", "W", "w", "X", "x", "Y", "y", "Z", "z"}


	specialCharList = []string{
		"☄", "★", "☎", "☻", "♨", "✠", "❝", "♚", "♘", "✎", "♞", "✩", "✪", "❦", "❥", "❣", "웃", "卍", "Ⓞ", "▶", "◙", "⊕", "◌", "⅓", "∭",
		"∮", "╳", "㏒", "㏕", "‱", "㎏", "❶", "Ň", "🅱", "🅾", "𝖋", "𝕻", "𝕼", "𝕽", "お", "な", "ぬ", "㊎", "㊞", "㊮", "✘"}
)


func newEvm(blockNumber *big.Int, blockHash common.Hash, state *state.StateDB) *vm.EVM {
	if nil == state {
		state, _ = newChainState()
	}
	evm := &vm.EVM{
		StateDB:  state,
	}
	context := vm.Context{
		BlockNumber: blockNumber,
		BlockHash: blockHash,
	}
	evm.Context = context

	//set a default active version
	govDB := gov.GovDBInstance()
	govDB.SetActiveVersion(initProcessVersion, state)

	return evm
}


func newPlugins() {
	plugin.GovPluginInstance()
	plugin.StakingInstance()
	plugin.SlashInstance()
	plugin.RestrictingInstance()
	plugin.RewardMgrInstance()

	snapshotdb.Instance()
}


func newChainState() (*state.StateDB, error) {
	var (
		db      = ethdb.NewMemDatabase()
		genesis = new(core.Genesis).MustCommit(db)
	)
	fmt.Println("genesis", genesis)
	// Initialize a fresh chain with only a genesis block
	blockchain, _ := core.NewBlockChain(db, nil, params.AllEthashProtocolChanges, nil, vm.Config{}, nil)

	var state *state.StateDB
	if statedb, err := blockchain.State(); nil != err {
		return nil, errors.New("reference statedb failed" + err.Error())
	} else {
		state = statedb
	}

	// init account balance
	state.AddBalance(sender, sender_balance)
	for i, addr := range addrArr {

		amount, _ := new(big.Int).SetString(balanceStr[len(addrArr) - 1 - i], 10)
		amount = new(big.Int).Mul(common.Big257, amount)
		state.AddBalance(addr, amount)
	}

	return state, nil
}

func build_staking_data (){


	stakingDB := staking.NewStakingDB ()
	sndb.NewBlock(big.NewInt(1), common.ZeroHash, blockHash)
	// MOCK


	nodeId_A := nodeIdArr[0]
	addr_A, _ := xutil.NodeId2Addr(nodeId_A)

	nodeId_B := nodeIdArr[1]
	addr_B, _ := xutil.NodeId2Addr(nodeId_B)

	nodeId_C := nodeIdArr[2]
	addr_C, _ := xutil.NodeId2Addr(nodeId_C)


	//canArr := make(staking.CandidateQueue, 0)


	c1 := &staking.Candidate{
		NodeId: nodeId_A,
		StakingAddress: sender,
		BenifitAddress: addrArr[1],
		StakingTxIndex: uint32(2),
		ProcessVersion: uint32(1),
		Status: staking.Valided,
		StakingEpoch: uint32(1),
		StakingBlockNum: uint64(1),
		Shares:             common.Big256,
		Released:           common.Big2,
		ReleasedHes:        common.Big32,
		RestrictingPlan:    common.Big1,
		RestrictingPlanHes: common.Big257,
		Description: staking.Description{
			ExternalId: "xxccccdddddddd",
			NodeName: "I Am " +fmt.Sprint(1),
			Website: "www.baidu.com",
			Details: "this is  baidu ~~",
		},
	}

	c2 := &staking.Candidate{
		NodeId: nodeId_B,
		StakingAddress: sender,
		BenifitAddress: addrArr[2],
		StakingTxIndex: uint32(3),
		ProcessVersion: uint32(1),
		Status: staking.Valided,
		StakingEpoch: uint32(1),
		StakingBlockNum: uint64(1),
		Shares:             common.Big256,
		Released:           common.Big2,
		ReleasedHes:        common.Big32,
		RestrictingPlan:    common.Big1,
		RestrictingPlanHes: common.Big257,
		Description: staking.Description{
			ExternalId: "SFSFSFSFSFSFSSFS",
			NodeName: "I Am " +fmt.Sprint(2),
			Website: "www.JD.com",
			Details: "this is  JD ~~",
		},
	}



	c3 := &staking.Candidate{
		NodeId: nodeId_C,
		StakingAddress: sender,
		BenifitAddress: addrArr[3],
		StakingTxIndex: uint32(4),
		ProcessVersion: uint32(1),
		Status: staking.Valided,
		StakingEpoch: uint32(1),
		StakingBlockNum: uint64(1),
		Shares:             common.Big256,
		Released:           common.Big2,
		ReleasedHes:        common.Big32,
		RestrictingPlan:    common.Big1,
		RestrictingPlanHes: common.Big257,
		Description: staking.Description{
			ExternalId: "FWAGGDGDGG",
			NodeName: "I Am " +fmt.Sprint(3),
			Website: "www.alibaba.com",
			Details: "this is  alibaba ~~",
		},
	}


	//canArr = append(canArr, c1)
	//canArr = append(canArr, c2)
	//canArr = append(canArr, c3)

	stakingDB.SetCanPowerStore(blockHash, addr_A, c1)
	stakingDB.SetCanPowerStore(blockHash, addr_B, c2)
	stakingDB.SetCanPowerStore(blockHash, addr_C, c3)


	stakingDB.SetCandidateStore(blockHash, addr_A, c1)
	stakingDB.SetCandidateStore(blockHash, addr_B, c2)
	stakingDB.SetCandidateStore(blockHash, addr_C, c3)


	queue := make(staking.ValidatorQueue, 0)

	v1 := &staking.Validator{
		NodeAddress: addr_A,
		NodeId: c1.NodeId,
		StakingWeight: [staking.SWeightItem]string{"1", common.Big256.String(), fmt.Sprint(c1.StakingBlockNum), fmt.Sprint(c1.StakingTxIndex)},
		ValidatorTerm: 0,
	}

	v2 := &staking.Validator{
		NodeAddress: addr_B,
		NodeId: c2.NodeId,
		StakingWeight: [staking.SWeightItem]string{"1", common.Big256.String(), fmt.Sprint(c2.StakingBlockNum), fmt.Sprint(c2.StakingTxIndex)},
		ValidatorTerm: 0,
	}

	v3 := &staking.Validator{
		NodeAddress: addr_C,
		NodeId: c3.NodeId,
		StakingWeight: [staking.SWeightItem]string{"1", common.Big256.String(), fmt.Sprint(c3.StakingBlockNum), fmt.Sprint(c3.StakingTxIndex)},
		ValidatorTerm: 0,
	}

	queue = append(queue, v1)
	queue = append(queue, v2)
	queue = append(queue, v3)


	val_Arr :=  &staking.Validator_array{
		Start: 1,
		End: 22000,
		Arr: queue,
	}

	stakingDB.SetVerfierList(blockHash, val_Arr)
	stakingDB.SetPreValidatorList(blockHash, val_Arr)
	stakingDB.SetCurrentValidatorList(blockHash, val_Arr)

	lastBlockHash = blockHash
	lastBlockNumber =  blockNumber.Uint64()
}

var (
	cList []*staking.Candidate
	vList []*staking.Validator
	addrList []common.Address
	val_Arr *staking.Validator_array
)

func InitPlatONPluginTestData() {
	for idx, nodeId := range  nodeIdArr{
		addr, _ := xutil.NodeId2Addr(nodeId)
		addrList = append(addrList, addr)

		c := &staking.Candidate{
			NodeId:             nodeId,
			StakingAddress:     sender,
			BenifitAddress:     addrArr[idx],
			StakingTxIndex:     uint32(idx),
			ProcessVersion:     uint32(1),
			Status:             staking.Valided,
			StakingEpoch:       uint32(1),
			StakingBlockNum:    uint64(1),
			Shares:             common.Big256,
			Released:           common.Big2,
			ReleasedHes:        common.Big32,
			RestrictingPlan:    common.Big1,
			RestrictingPlanHes: common.Big257,
			Description: staking.Description{
				ExternalId: "xxccccdddddddd",
				NodeName:   "I Am " + fmt.Sprint(idx),
				Website:    "www.baidu.com",
				Details:    "this is  baidu ~~",
			},
		}
		cList = append(cList, c)


		v := &staking.Validator{
			NodeAddress:   addr,
			NodeId:        nodeId,
			StakingWeight: [staking.SWeightItem]string{"1", common.Big256.String(), fmt.Sprint(c.StakingBlockNum), fmt.Sprint(c.StakingTxIndex)},
			ValidatorTerm: 0,
		}
		vList = append(vList, v)
	}

	val_Arr = &staking.Validator_array{
		Start: 1,
		End:   32000,
		Arr:   vList,
	}
}

func buildSnapDBDataNoCommit(blockNum int){

	no := int64(blockNum)
	header := types.Header{
		Number: big.NewInt(no),
	}
	hash := header.Hash()

	stakingDB := staking.NewStakingDB ()
	sndb.NewBlock(big.NewInt(int64(blockNum)), lastBlockHash, hash)

	for i :=0; i<25; i++ {
		stakingDB.SetCanPowerStore(hash, addrList[i], cList[i])
		stakingDB.SetCandidateStore(hash, addrList[i], cList[i])
	}

	stakingDB.SetVerfierList(hash, val_Arr)
	stakingDB.SetPreValidatorList(hash, val_Arr)
	stakingDB.SetCurrentValidatorList(hash, val_Arr)

	lastBlockHash = hash
	lastBlockNumber = uint64(blockNum)
	lastHeader = header
}

func buildSnapDBDataCommitted(start, end int){

	stakingDB := staking.NewStakingDB ()

	for i := start; i <= end; i++ {
		no := int64(i)
		header := types.Header{
			Number: big.NewInt(no),
		}
		hash := header.Hash()

		sndb.NewBlock(big.NewInt(no), lastBlockHash, hash)

		for i :=0; i<25; i++ {
			stakingDB.SetCanPowerStore(hash, addrList[i], cList[i])
			stakingDB.SetCandidateStore(hash, addrList[i], cList[i])
		}

		stakingDB.SetVerfierList(hash, val_Arr)
		stakingDB.SetPreValidatorList(hash, val_Arr)
		stakingDB.SetCurrentValidatorList(hash, val_Arr)

		sndb.Commit(hash)

		lastBlockHash = hash
		lastBlockNumber = uint64(i)
		lastHeader = header
	}
}

func build_gov_data (state *state.StateDB){
	//set a default active version
	govDB := gov.GovDBInstance()
	govDB.SetActiveVersion(initProcessVersion, state)
}


func buildDbRestrictingPlan(t *testing.T, stateDB xcom.StateDB) {
	account := addrArr[0]

	const Epochs = 6
	var list = make([]uint64, Epochs)

	for epoch := 1; epoch < Epochs; epoch++ {
		// build release account record
		releaseAccountKey := restricting.GetReleaseAccountKey(uint64(epoch), 1)
		stateDB.SetState(cvm.RestrictingContractAddr, releaseAccountKey, account.Bytes())

		// build release amount record
		releaseAmount := big.NewInt(int64(1E18))
		releaseAmountKey := restricting.GetReleaseAmountKey(uint64(epoch), account)
		stateDB.SetState(account, releaseAmountKey, releaseAmount.Bytes())

		// build release epoch record
		releaseEpochKey := restricting.GetReleaseEpochKey(uint64(epoch))
		stateDB.SetState(cvm.RestrictingContractAddr, releaseEpochKey, common.Uint64ToBytes(1))

		list = append(list, uint64(epoch))
	}

	// build restricting user info
	var user restrictingInfo
	user.balance = big.NewInt(int64(5E18))
	user.debt = big.NewInt(0)
	user.releaseList = list

	bUser, err := rlp.EncodeToBytes(user)
	if err != nil {
		t.Errorf("failed to rlp encode restricting info: %s", err.Error())
	}

	restrictingKey := restricting.GetRestrictingKey(account)
	stateDB.SetState(cvm.RestrictingContractAddr, restrictingKey, bUser)

	stateDB.AddBalance(sender, sender_balance.Sub(sender_balance, big.NewInt(int64(5E18))))
	stateDB.AddBalance(cvm.RestrictingContractAddr, big.NewInt(int64(5E18)))
}
