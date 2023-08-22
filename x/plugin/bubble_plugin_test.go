package plugin

import (
	"fmt"
	"github.com/bubblenet/bubble/common/hexutil"
	"math/big"
	"testing"
)

func TestVRF(t *testing.T) {
	queue := VRFQueue{
		{w: big.NewInt(int64(4))},
		{w: big.NewInt(int64(15))},
		{w: big.NewInt(int64(20))},
		//{w: big.NewInt(int64(12))},
		//{w: big.NewInt(int64(31))},
		//{w: big.NewInt(int64(1))},
		//{w: big.NewInt(int64(7))},
		//{w: big.NewInt(int64(8))},
		//{w: big.NewInt(int64(6))},
		//{w: big.NewInt(int64(59))},
		//{w: big.NewInt(int64(5))},
		//{w: big.NewInt(int64(3))},
		//{w: big.NewInt(int64(4))},
		//{w: big.NewInt(int64(17))},
		//{w: big.NewInt(int64(46))},
		//{w: big.NewInt(int64(22))},
		//{w: big.NewInt(int64(16))},
		//{w: big.NewInt(int64(35))},
	}
	curNonce := hexutil.MustDecode("0x024c6378c176ef6c717cd37a74c612c9abd615d13873ff6651e3d352b31cb0b2e1")
	preNoce := [][]byte{
		hexutil.MustDecode("0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23"),
		hexutil.MustDecode("0x024c6378c176ef6c717cd37a74c612c9abd615d13873ff6651e3d351b31cb0b2e1"),
		hexutil.MustDecode("0x032c6378c170c1f9472a6925a3ea2f951a606222ab231a31310e15d1b31cb0b2e1"),
		hexutil.MustDecode("0x0caccce29d07a066eae8fe4a34049cbd2d68d92bf02b615d13873ff6651e3d351b"),
		//hexutil.MustDecode("0x0b9015c94a2fea0c1f9472a6925a3ea2f951a606222ab231a31310e15d8c2725"),
		//hexutil.MustDecode("0x05cce29d07a066eae8fe4a34049cbd2d68d92bf02bc618ae2d43f0b355ad678c"),
		//hexutil.MustDecode("0xa5880c17478c173f407e38c997c4defe290779d15dda19a1e1bd5ca66429006c"),
	}
	for i := 0; i < 100; i++ {
		fmt.Println("=============================")
		if vrfed, err := VRF(queue, 1, curNonce, preNoce); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(vrfed)
		}
	}

}
