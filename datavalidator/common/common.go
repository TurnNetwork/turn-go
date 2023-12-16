package common

import (
	"fmt"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/datavalidator/types"
)

func VerifySignature(id []byte, signature *types.Signature, key *bls.PublicKey) bool {
	var sign bls.Sign
	sign.Deserialize(signature.Signature)
	return sign.Verify(key, string(id))
}

//func VerifySignature(id []byte, signature *types.Signature, address common.Address) bool {
//	pub, err := crypto.SigToPub(id, signature.Signature)
//	if err != nil {
//		return false
//	}
//	addr := crypto.PubkeyToAddress(*pub)
//	return addr == address
//}

func PeerID(id []byte) string {
	return fmt.Sprintf("%x", id[:8])
}
