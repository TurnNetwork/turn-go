package wallet

import (
	"github.com/bubblenet/bubble/crypto/bls"

	"github.com/bubblenet/bubble/common"
)

type Wallet interface {
	PublicKey() *bls.PublicKey
	Sign(hash common.Hash) []byte
}
type LocalWallet struct {
	sk *bls.SecretKey
}

func FromFile(file string) (*LocalWallet, error) {
	return nil, nil
}

func FromHexSk(pri string) (*LocalWallet, error) {
	return nil, nil
}
func FromSk(sk *bls.SecretKey) *LocalWallet {
	return &LocalWallet{
		sk: sk,
	}
}

func (w *LocalWallet) Sign(hash common.Hash) []byte {
	sign := w.sk.Sign(string(hash[:]))
	return sign.Serialize()
}

func (w *LocalWallet) PublicKey() *bls.PublicKey {
	return w.sk.GetPublicKey()
}
