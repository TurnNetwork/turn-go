package mock

import (
	"crypto/ecdsa"
	"encoding/hex"
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
		"f958b2708c0a6eae0ea5761edcf0257526a8bbe521cc099e32adbff14b049734",
		"86e71e0d2feb7cb2233aaf084beaca51a1dfa0107d4e8ae4417555786820de1a",
		"079d678e6e61d949c60d43237c5b62d9fb2b4e71747ddd804673358e5bbd253f",
		"6ad2849adea42f1a9f981a7caa3733024c08bcba3849cc7f7ec2fe808debaf5b",
	}
	var sks []*ecdsa.PrivateKey
	var blsKeys []*bls.SecretKey
	for i := 0; i < len(skstr); i++ {
		sk, _ := crypto.HexToECDSA(skstr[i])
		sks = append(sks, sk)
		var key bls.SecretKey
		buf, _ := hex.DecodeString(blsKeyStrs[i])
		key.SetLittleEndian(buf)
		blsKeys = append(blsKeys, &key)
	}
	SystemInnerContract = NewInnerContract(sks, blsKeys, map[uint64][]uint64{
		1: []uint64{0, 1, 2, 3},
	})
	SystemBlockFilter = NewBlockFilter([]uint64{1})
}
