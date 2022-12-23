package params

import (
	"fmt"
	"github.com/bubblenet/bubble/crypto/bls"
	"github.com/bubblenet/bubble/p2p/enode"
)

type CbftNode struct {
	Node      *enode.Node   `json:"node"`
	BlsPubKey bls.PublicKey `json:"blsPubKey"`
	RPC       string
}

type InitNode struct {
	Enode     string
	BlsPubkey string
	RPC       string
}

// AuthConfig frp server authenticate struct:Enable token authentication
type AuthConfig struct {
	Method       string // authentication method
	HeartBeats   bool   // authenticate heartbeats
	NewWorkConns bool   // authenticate_new_work_conns
	Token        string // token authentication
}

// FrpsConfig frp server configuration structure
type FrpsConfig struct {
	ServerIP    string      // frp server ip
	ServerPort  int         // frp server port
	StunAddress string      // stun server address
	Auth        *AuthConfig // Enable authentication mode
}

type CbftConfig struct {
	Period        uint64     `json:"period,omitempty"`        // Number of seconds between blocks to enforce
	Amount        uint32     `json:"amount,omitempty"`        //The maximum number of blocks generated per cycle
	InitialNodes  []CbftNode `json:"initialNodes,omitempty"`  //Genesis consensus node
	ValidatorMode string     `json:"validatorMode,omitempty"` //Validator mode for easy testing
}

func ConvertNodeUrl(initialNodes []InitNode) []CbftNode {
	bls.Init(bls.BLS12_381)
	NodeList := make([]CbftNode, 0, len(initialNodes))
	for _, n := range initialNodes {

		cbftNode := new(CbftNode)

		if node, err := enode.Parse(enode.ValidSchemes, n.Enode); nil == err {
			cbftNode.Node = node
		}

		if n.BlsPubkey != "" {
			var blsPk bls.PublicKey
			if err := blsPk.UnmarshalText([]byte(n.BlsPubkey)); nil == err {
				cbftNode.BlsPubKey = blsPk
			}
		}

		NodeList = append(NodeList, *cbftNode)
	}
	return NodeList
}

// String implements the fmt.Stringer interface.
func (c *CbftConfig) String() string {
	initialNodes := make([]InitNode, 0)
	for _, node := range c.InitialNodes {
		initialNodes = append(initialNodes, InitNode{
			Enode: node.Node.String(),
			//	BlsPubkey: node.BlsPubKey.GetHexString(),
		})
	}

	return fmt.Sprintf("{period: %v  amount: %v initialNodes: %v validatorMode: %v}",
		c.Period,
		c.Amount,
		initialNodes,
		c.ValidatorMode,
	)
}
