package types

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
)

type blockExt struct {
	// Block belongs to the view
	ViewNumber uint64
	Block      *types.Block

	// Time of receipt of the Block
	RcvTime time.Time

	// blockExt only store aggregated signatures,
	QC *QuorumCert

	// Point to the Parent Block
	Parent *blockExt

	// There may be more than one sub-Block, and the Block will not be deleted if it is not in the state of LockQC.
	Children map[common.Hash]*blockExt
}

func (b *blockExt) clearParent() {
	if b.Parent != nil {
		b.Parent.Children = nil
		b.Parent = nil
	}
}

func (b *blockExt) clearChildren() {
	if b.Children != nil {
		b.Children = nil
	}
}

func (b *blockExt) MarshalJSON() ([]byte, error) {
	type BlockExt struct {
		ViewNumber  uint64        `json:"view_number"`
		BlockHash   common.Hash   `json:"block_hash"`
		BlockNumber uint64        `json:"block_number"`
		RcvTime     time.Time     `json:"receive_time"`
		QC          *QuorumCert   `json:"qc"`
		ParentHash  common.Hash   `json:"parent_hash"`
		Children    []common.Hash `json:"children_hash"`
	}
	ext := &BlockExt{
		ViewNumber:  b.ViewNumber,
		BlockHash:   b.Block.Hash(),
		BlockNumber: b.Block.NumberU64(),
		RcvTime:     b.RcvTime,
		QC:          b.QC,
		Children:    make([]common.Hash, 0),
	}
	if b.Parent != nil {
		ext.ParentHash = b.Parent.Block.Hash()
	}

	for h := range b.Children {
		ext.Children = append(ext.Children, h)
	}
	return json.Marshal(ext)
}

// BlockTree used to store blocks that are not currently written to disk， Block of QC, LockQC. Every time you submit to blockTree, it is possible to start QC changes.
type BlockTree struct {
	// The highest Block that has been written to disk, root will grow with each commit
	root *blockExt
	// Contains blocks generated by multiple views, all blocks stored are not committed
	blocks map[uint64]map[common.Hash]*blockExt
}

func NewBlockTree(root *types.Block, qc *QuorumCert) *BlockTree {
	blockTree := &BlockTree{
		root: &blockExt{
			Block:    root,
			RcvTime:  time.Now(),
			QC:       qc,
			Parent:   nil,
			Children: make(map[common.Hash]*blockExt),
		},
		blocks: make(map[uint64]map[common.Hash]*blockExt),
	}

	if root.NumberU64() == 0 {
		blockTree.root.ViewNumber = math.MaxUint64
	} else {
		blockTree.root.ViewNumber = qc.ViewNumber
	}

	blocks := make(map[common.Hash]*blockExt)
	blocks[root.Hash()] = blockTree.root
	blockTree.blocks[root.NumberU64()] = blocks
	return blockTree
}

// Insert a Block that has reached the QC state, returns the LockQC, Commit Block based on the height of the inserted Block
func (b *BlockTree) InsertQCBlock(block *types.Block, qc *QuorumCert) (*types.Block, *types.Block) {
	ext := &blockExt{
		Block:    block,
		RcvTime:  time.Now(),
		QC:       qc,
		Parent:   nil,
		Children: make(map[common.Hash]*blockExt),
	}
	if block.NumberU64() == 0 {
		ext.ViewNumber = math.MaxUint64
	} else {
		ext.ViewNumber = qc.ViewNumber
	}

	return b.insertBlock(ext)
}

// Delete invalid branch Block
func (b *BlockTree) PruneBlock(hash common.Hash, number uint64, clearFn func(*types.Block)) {
	for num := b.root.Block.NumberU64(); num < number; num++ {
		delete(b.blocks, num)
	}
	if extMap, ok := b.blocks[number]; ok {
		for h, ext := range extMap {
			ext.clearParent()
			if h != hash {
				delete(extMap, h)
				b.pruneBranch(ext, clearFn)
			} else {
				b.root = ext
			}
		}
	}
}

func (b *BlockTree) NewRoot(block *types.Block) {
	hash, number := block.Hash(), block.NumberU64()
	for i := b.root.Block.NumberU64(); i < block.NumberU64(); i++ {
		delete(b.blocks, i)
	}
	b.root = b.findBlockExt(hash, number)
}

// FindBlockAndQC find the specified Block and its QC.
func (b *BlockTree) FindBlockAndQC(hash common.Hash, number uint64) (*types.Block, *QuorumCert) {
	ext := b.findBlockExt(hash, number)
	if ext != nil {
		return ext.Block, ext.QC
	}
	return nil, nil
}

func (b *BlockTree) findBlockExt(hash common.Hash, number uint64) *blockExt {
	if extMap, ok := b.blocks[number]; ok {
		for h, ext := range extMap {
			if hash == h {
				return ext
			}
		}
	}
	return nil
}

// FindBlockByHash find the specified Block by hash.
func (b *BlockTree) FindBlockByHash(hash common.Hash) *types.Block {
	for _, extMap := range b.blocks {
		for h, ext := range extMap {
			if h == hash {
				return ext.Block
			}
		}
	}
	return nil
}

func (b *BlockTree) pruneBranch(ext *blockExt, clearFn func(*types.Block)) {
	for h, e := range ext.Children {
		if extMap, ok := b.blocks[e.Block.NumberU64()]; ok {
			if clearFn != nil {
				clearFn(e.Block)
			}
			e.clearParent()
			delete(extMap, h)
			b.pruneBranch(e, clearFn)
		}
	}
	ext.clearChildren()
}

func (b *BlockTree) insertBlock(ext *blockExt) (*types.Block, *types.Block) {
	number := ext.Block.NumberU64()
	hash := ext.Block.Hash()
	if extMap, ok := b.blocks[number]; ok {
		if _, ok := extMap[hash]; !ok {
			extMap[hash] = ext
		} else {
			return nil, nil
		}
	} else {
		extMap := make(map[common.Hash]*blockExt)
		extMap[hash] = ext
		b.blocks[number] = extMap
	}

	b.fixTree(ext)

	return b.commitBlock(b.maxBlock(ext))
}

// Return LockQC, Commit Blocks
func (b *BlockTree) commitBlock(ext *blockExt) (*types.Block, *types.Block) {
	lock := ext.Parent
	if lock == nil {
		panic(fmt.Sprintf("%s,%d", ext.Block.Hash().String(), ext.Block.NumberU64()))
	}
	if lock.Block.Hash() == b.root.Block.Hash() {
		return b.root.Block, b.root.Block
	}
	commit := lock.Parent
	return lock.Block, commit.Block
}

// Returns the maximum view number Block for a given height
func (b *BlockTree) maxBlock(ext *blockExt) *blockExt {
	max := ext
	if extMap, ok := b.blocks[ext.Block.NumberU64()]; ok {
		for _, e := range extMap {
			//genesis Block is max.Uint64
			if max.ViewNumber+1 < e.ViewNumber+1 && e.Parent != nil {
				max = e
			}
		}
	}
	return max
}

// Connect Parent and child blocks
func (b *BlockTree) fixTree(ext *blockExt) {
	parent := b.findParent(ext.Block.ParentHash(), ext.Block.NumberU64())
	child := b.findChild(ext.Block.Hash(), ext.Block.NumberU64())

	if parent != nil {
		parent.Children[ext.Block.Hash()] = ext
		ext.Parent = parent
	}

	if child != nil {
		child.Parent = ext
		ext.Children[child.Block.Hash()] = child
	}
}

func (b *BlockTree) findParent(hash common.Hash, number uint64) *blockExt {
	if extMap, ok := b.blocks[number-1]; ok {
		for _, v := range extMap {
			if v.Block != nil {
				if v.Block.Hash() == hash {
					return v
				}
			}
		}
	}
	return nil
}

func (b *BlockTree) findChild(hash common.Hash, number uint64) *blockExt {
	if extMap, ok := b.blocks[number+1]; ok {
		for _, v := range extMap {
			if v.Block != nil {
				if v.Block.ParentHash() == hash {
					return v
				}
			}
		}
	}
	return nil
}

func (b *BlockTree) MarshalJSON() ([]byte, error) {
	type blockTree struct {
		Root *blockExt `json:"root"`
		// Contains blocks generated by multiple views, all blocks stored are not committed
		Blocks map[uint64]map[common.Hash]*blockExt `json:"blocks"`
	}

	tree := &blockTree{
		Root:   b.root,
		Blocks: b.blocks,
	}
	return json.Marshal(tree)
}

func (b *BlockTree) Reset(root *types.Block, qc *QuorumCert) {
	b.root = &blockExt{
		Block:    root,
		RcvTime:  time.Now(),
		QC:       qc,
		Parent:   nil,
		Children: make(map[common.Hash]*blockExt),
	}

	if root.NumberU64() == 0 {
		b.root.ViewNumber = math.MaxUint64
	} else {
		b.root.ViewNumber = qc.ViewNumber
	}

	b.blocks = make(map[uint64]map[common.Hash]*blockExt)
	blocks := make(map[common.Hash]*blockExt)
	blocks[root.Hash()] = b.root
	b.blocks[root.NumberU64()] = blocks
}
