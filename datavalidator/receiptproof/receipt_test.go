package receiptproof

import (
	"encoding/json"
	"fmt"
	"github.com/bubblenet/bubble/trie"
	"math/big"
	"testing"

	types2 "github.com/bubblenet/bubble/consensus/cbft/types"
	"github.com/bubblenet/bubble/rlp"

	"github.com/bubblenet/bubble/common"
	"github.com/bubblenet/bubble/core/types"
	"github.com/stretchr/testify/require"
)

func TestReceiptProof(t *testing.T) {
	receipts := types.Receipts{
		types.NewReceipt(nil, false, 100),
		types.NewReceipt(nil, false, 101),
		types.NewReceipt(nil, false, 102),
		types.NewReceipt(nil, false, 103),
		types.NewReceipt(nil, false, 104),
		types.NewReceipt(nil, false, 105),
		types.NewReceipt(nil, false, 106),
		types.NewReceipt(nil, false, 107),
	}
	//create receipt proof of index  2
	receipts[2].Logs = []*types.Log{
		&types.Log{
			Address:     common.BigToAddress(big.NewInt(22)),
			Topics:      []common.Hash{},
			Data:        []byte{},
			BlockNumber: 33,
			TxHash:      [32]byte{},
			TxIndex:     0,
			BlockHash:   [32]byte{},
			Index:       0,
			Removed:     false,
		},
	}
	header := types.Header{
		ParentHash:  common.BigToHash(big.NewInt(1)),
		Coinbase:    common.BigToAddress(big.NewInt(2)),
		Root:        common.BigToHash(big.NewInt(3)),
		TxHash:      common.BigToHash(big.NewInt(4)),
		ReceiptHash: types.DeriveSha(receipts, new(trie.Trie)),
		Bloom:       [256]byte{},
		Number:      &big.Int{},
		GasLimit:    0,
		GasUsed:     0,
		Time:        0,
		Extra:       []byte{},
		Nonce:       [81]byte{},
	}
	prover := &Prover{}
	root, path, err := prover.TrieProof(receipts, 2)
	require.Nil(t, err)
	require.Equal(t, header.ReceiptHash, root)
	//create log proof, proof contains header qc proof, receipt header proof, log in receipt
	proof := &Proof{
		Header: &header,
		QC:     nil,
		Index:  2,
		Path:   path,
	}
	r, err := proof.Verify(func(header *types.Header, qc *types2.QuorumCert) error {
		return nil
	})
	require.Nil(t, err)
	require.True(t, r.CumulativeGasUsed == 102)
	l, err := proof.VerifyLog(func(header *types.Header, qc *types2.QuorumCert) error {
		return nil
	}, 0)
	require.Nil(t, err)
	require.True(t, l.Address == common.BigToAddress(big.NewInt(22)))
	s2, _ := rlp.EncodeToBytes(proof)
	fmt.Println("len:", len(s2))

	s, _ := json.Marshal(proof)
	fmt.Println("len:", len(s))
	fmt.Println(string(s))
	// mdb := memorydb.New()
	// mdb2 := memorydb.New()
	// tdb := trie.NewDatabase(mdb)
	// tr, _ := trie.New(common.Hash{}, tdb)
	// receiptSha := types.DeriveShaTrie(receipts, tr)
	// keybuf := new(bytes.Buffer)
	// keybuf.Reset()
	// rlp.Encode(keybuf, uint(2))
	// err := tr.Prove(keybuf.Bytes(), 0, mdb2)
	// require.Nil(t, err)
	// it := mdb2.NewIterator(nil, nil)
	// for it.Next() {
	// 	mdb2.Put(it.Key(), []byte{1})
	// }
	// value, err := trie.VerifyProof(receiptSha, keybuf.Bytes(), mdb2)
	// require.Nil(t, err)
	// var r types.Receipt
	// rlp.DecodeBytes(value, &r)
	// require.True(t, r.CumulativeGasUsed == 102)
}
