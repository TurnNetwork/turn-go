package mock

import (
	"crypto/ecdsa"
	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/crypto/bls"
)

var SystemBlockFilter *BlockFilter
var SystemInnerContract *InnerContract

func init() {
	var skstr = []string{
		"02c3ba761f7a190cc69a13ffe929acbde4bdea64260fa448b86613e794bf5d82",
		"603536aa1efa9f9a0c14411d911526261e78b81ceeab3338085e086954a71c94",
		"02e9ecf58eaba70e3ebc3d1d34031f389cccb1fea0a8df98db40fb0c339a8f81",
		"e89bf32b04bc2a8bdec7bafb514caf1348b7e5dab76dcf5948705c8ac098ba23",
	}
	var blsKeyStrs = []string{
		"2488270d920dd57f0a2c7f69577b019c575480663356df4e94f90d5a1d6a72ec",
		"6b5ac12b43001b4f53f407878de9e0225bbfeb92541ed8a6e66af514bdac05ed",
		"1bf528b31769a7d1bfaf041ade586e2896a503e39e85bd0d463ef19bc45d3aa7",
		"4c2ea01368827230e8c86e27ad4f8a8f7925f32c49a70a98435e50b9e1411495",
	}
	var sks []*ecdsa.PrivateKey
	var blsKeys []*bls.SecretKey
	for i := 0; i < len(skstr); i++ {
		sk, _ := crypto.HexToECDSA(skstr[i])
		sks = append(sks, sk)
		var key bls.SecretKey
		key.SetHexString(blsKeyStrs[i])
		blsKeys = append(blsKeys, &key)
	}
	SystemInnerContract = NewInnerContract(sks, blsKeys, map[uint64][]uint64{
		1: []uint64{0, 1, 2, 3},
	})
	SystemBlockFilter = NewBlockFilter([]uint64{1})
}
