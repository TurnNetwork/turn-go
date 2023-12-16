package simulation

import (
	"testing"

	"github.com/bubblenet/bubble/p2p"
)

func TestServer(t *testing.T) {
	config := p2p.Config{
		Name:       "test",
		MaxPeers:   10,
		ListenAddr: "127.0.0.1:0",
		PrivateKey: newkey(),
	}
	srv := NewServer(config)
	srv.Start()
}
