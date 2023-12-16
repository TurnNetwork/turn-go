package receiptproof

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/bubblenet/bubble/common"
	types2 "github.com/bubblenet/bubble/consensus/cbft/types"
	"github.com/bubblenet/bubble/core/types"
	"github.com/bubblenet/bubble/ethclient"
	"github.com/bubblenet/bubble/ethdb/memorydb"
	"github.com/bubblenet/bubble/rlp"
	"github.com/bubblenet/bubble/trie"
)

type Proof struct {
	Header *types.Header      `json:"header"`
	QC     *types2.QuorumCert `json:"qc"`
	Index  uint64             `json:"index"`
	Path   [][2][]byte        `json:"path"`
}

type Prover struct {
	cli *ethclient.Client
}

func NewProver(url string) (*Prover, error) {
	cli, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Prover{
		cli: cli,
	}, nil
}

func (p *Prover) Prove(txHash common.Hash) (*Proof, error) {
	ctx := context.Background()
	receipt, err := p.cli.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}
	block, err := p.cli.BlockByHash(ctx, receipt.BlockHash)
	if err != nil {
		return nil, err
	}
	qc, err := p.cli.GetBlockQuorumCertByHash(ctx, block.Hash())
	if err != nil {
		return nil, err
	}
	var receipts types.Receipts
	index := uint64(0)
	for i, tx := range block.Transactions() {
		if tx.Hash() == txHash {
			index = uint64(i)
		}
		r, err := p.cli.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, err
		}
		receipts = append(receipts, r)
	}
	root, path, err := p.TrieProof(receipts, index)
	if err != nil {
		return nil, err
	}
	if root != block.ReceiptHash() {
		return nil, errors.New("invalid receipt root")
	}
	return &Proof{
		Header: block.Header(),
		QC:     qc,
		Path:   path,
		Index:  index,
	}, nil
}

func (p *Prover) TrieProof(receipts types.Receipts, index uint64) (common.Hash, [][2][]byte, error) {
	mdb := memorydb.New()
	t := new(trie.Trie)
	root := DeriveShaTrie(receipts, t)
	keybuf := new(bytes.Buffer)
	keybuf.Reset()
	rlp.Encode(keybuf, uint(index))
	err := t.Prove(keybuf.Bytes(), 0, mdb)
	if err != nil {
		return common.Hash{}, nil, err
	}
	var res [][2][]byte
	it := mdb.NewIterator(nil, nil)
	for it.Next() {
		res = append(res, [2][]byte{it.Key(), it.Value()})
	}
	return root, res, nil
}

func (p *Proof) Bytes() ([]byte, error) {
	buf, err := rlp.EncodeToBytes(p)
	return buf, err
}

func (p *Proof) MarshalJSON() ([]byte, error) {
	var pathStr [][2]string
	for _, k := range p.Path {
		pathStr = append(pathStr, [2]string{
			hex.EncodeToString(k[0]),
			hex.EncodeToString(k[1]),
		})
	}
	type proof struct {
		Header *types.Header
		QC     *types2.QuorumCert
		Index  uint64
		Path   [][2]string
	}
	pf := &proof{
		Header: p.Header,
		QC:     p.QC,
		Index:  p.Index,
		Path:   pathStr,
	}
	return json.Marshal(pf)
}
func (p *Proof) UnMarshalJSON(input []byte) error {
	type proof struct {
		Header *types.Header
		QC     *types2.QuorumCert
		Index  uint64
		Path   [][2]string
	}
	var pf proof
	err := json.Unmarshal(input, &pf)
	if err != nil {
		return err
	}
	var path [][2][]byte
	for _, k := range pf.Path {
		v1, err := hex.DecodeString(k[0])
		if err != nil {
			return err
		}
		v2, err := hex.DecodeString(k[1])
		if err != nil {
			return err
		}
		path = append(path, [2][]byte{
			v1, v2,
		})
	}
	p.Header = pf.Header
	p.QC = pf.QC
	p.Index = pf.Index
	p.Path = path
	return nil
}
func (p *Proof) Verify(verifyHeader func(header *types.Header, qc *types2.QuorumCert) error) (*types.Receipt, error) {
	if err := verifyHeader(p.Header, p.QC); err != nil {
		return nil, err
	}
	keybuf := new(bytes.Buffer)
	keybuf.Reset()
	rlp.Encode(keybuf, uint(p.Index))
	mdb := memorydb.New()
	for _, k := range p.Path {
		mdb.Put(k[0], k[1])
	}
	value, err := trie.VerifyProof(p.Header.ReceiptHash, keybuf.Bytes(), mdb)
	if err != nil {
		return nil, err
	}
	var receipt types.Receipt
	err = receipt.DecodeRLP(rlp.NewStream(bytes.NewReader(value), 0))
	if err != nil {
		return nil, err
	}
	return &receipt, nil
}

func (p *Proof) VerifyLog(verifyHeader func(header *types.Header, qc *types2.QuorumCert) error, logIndex uint64) (*types.Log, error) {
	receipt, err := p.Verify(verifyHeader)
	if err != nil {
		return nil, err
	}
	if uint64(len(receipt.Logs)) <= logIndex {
		return nil, errors.New("invalid log index")
	}
	return receipt.Logs[logIndex], nil
}
func DeriveShaTrie(list types.DerivableList, trie *trie.Trie) common.Hash {
	keybuf := new(bytes.Buffer)
	for i := 0; i < list.Len(); i++ {
		keybuf.Reset()
		rlp.Encode(keybuf, uint(i))
		trie.Update(keybuf.Bytes(), list.GetRlp(i))
	}
	return trie.Hash()
}
