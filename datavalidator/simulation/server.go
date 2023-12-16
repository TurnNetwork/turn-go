package simulation

import (
	"crypto/ecdsa"

	"github.com/bubblenet/bubble/crypto"
	"github.com/bubblenet/bubble/p2p"
)

func newkey() *ecdsa.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic("couldn't generate key: " + err.Error())
	}
	return key
}

func NewServer(config p2p.Config) *p2p.Server {
	return &p2p.Server{
		Config: config,
	}
}
