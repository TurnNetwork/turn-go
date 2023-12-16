package mock

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/bubblenet/bubble/crypto/bls"
	"sync"

	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/datavalidator/types"
)

type InnerContract struct {
	sync.Mutex
	vs      []*types.Validator
	bubbles map[uint64]*types.ChildChainValidator
}

func NewInnerContract(sks []*ecdsa.PrivateKey, blsPrivate []*bls.SecretKey, bubble map[uint64][]uint64) *InnerContract {
	var vs []*types.Validator
	for i, sk := range sks {
		//fmt.Println("inner validator", "index", i, "PubId", common.BlsID(blsPrivate[i].GetPublicKey()))
		//wallet := wallet2.FromSk(blsPrivate[i])
		pub := hex.EncodeToString(crypto.FromECDSAPub(&sk.PublicKey)[1:])
		vs = append(vs, &types.Validator{
			Index:       uint64(i),
			Address:     crypto.PubkeyToAddress(sk.PublicKey),
			BlsPubKey:   blsPrivate[i].GetPublicKey(),
			P2pUrl:      fmt.Sprintf("enode://%s@172.17.0.1:1800%d", pub, i+1),
			BlockNumber: 0,
		})
	}
	children := make(map[uint64]*types.ChildChainValidator)
	for id, indexs := range bubble {
		var group []*types.Validator
		for _, i := range indexs {
			group = append(group, vs[i])
		}
		children[id] = &types.ChildChainValidator{
			Validators: group,
			Threshold:  len(group),
		}
	}
	return &InnerContract{
		vs:      vs,
		bubbles: children,
	}
}
func (i InnerContract) ValidatorSet() []*types.Validator {
	return i.vs
}
func (i InnerContract) GetValidator(addr *bls.PublicKey) *types.Validator {
	for _, v := range i.vs {
		if v.BlsPubKey.IsEqual(addr) {
			return v
		}
	}
	return nil
}
func (i *InnerContract) GetValidators(chainId uint64) map[uint64]*types.Validator {
	i.Lock()
	defer i.Unlock()
	vs := i.bubbles[chainId]
	if vs != nil {
		s := make(map[uint64]*types.Validator)
		for _, v := range vs.Validators {
			//fmt.Println("validator:", common.BlsID(v.BlsPubKey))
			s[v.Index] = v
		}
		return s
	}
	return nil
}

func (i InnerContract) IsValidator(id string) bool {
	return true
}
func (i *InnerContract) GetBubbleId() []uint64 {
	i.Lock()
	defer i.Unlock()
	var ids []uint64
	for id, _ := range i.bubbles {
		ids = append(ids, id)
	}
	return ids
}

func (i InnerContract) GetBubbleValidator(chainId uint64) *types.ChildChainValidator {
	return i.bubbles[chainId]
}
