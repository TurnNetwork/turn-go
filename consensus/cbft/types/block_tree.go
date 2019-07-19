package types

import (
	"time"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
)

//BlockTree used to store blocks that are not currently written to disk， Block of QC, LockQC. Every time you submit to blockTree, it is possible to start QC changes.
type BlockTree struct {
	// The highest block that has been written to disk, root will grow with each commit
	root *BlockExt
	// Contains blocks generated by multiple views, all blocks stored are not committed
	blocks map[uint64]map[common.Hash]*BlockExt
}

type BlockExt struct {
	// Block belongs to the view
	viewNumber uint64
	block      *types.Block

	// Time of receipt of the block
	rcvTime int64

	// blockExt only store aggregated signatures,
	qc *QuorumCert

	// Point to the parent block
	parent *BlockExt

	// There may be more than one sub-block, and the block will not be deleted if it is not in the state of LockQC.
	children map[common.Hash]*BlockExt
}

// Insert a block that has reached the QC state, returns the LockQC, Commit block based on the height of the inserted block
func (b *BlockTree) InsertQCBlock(block *types.Block, qc *QuorumCert) (*types.Block, *types.Block) {
	ext := &BlockExt{
		viewNumber: qc.ViewNumber,
		block:      block,
		rcvTime:    time.Now().Unix(),
		qc:         qc,
		parent:     nil,
		children:   make(map[common.Hash]*BlockExt),
	}

	return b.insertBlock(ext)
}

// Delete invalid branch block
func (b *BlockTree) PruneBlock(hash common.Hash, number uint64, clearFn func(*types.Block)) {
	if extMap, ok := b.blocks[number]; ok {
		for h, ext := range extMap {
			if h != hash {
				delete(extMap, h)
				b.pruneBranch(ext, clearFn)
			} else {
				b.root = ext
			}
		}
	}
}

// FindBlockAndQC find the specified block and its QC.
func (b *BlockTree) FindBlockAndQC(hash common.Hash, number uint64) (*types.Block, *QuorumCert) {
	if extMap, ok := b.blocks[number]; ok {
		for h, ext := range extMap {
			if hash == h {
				return ext.block, ext.qc
			}
		}
	}
	return nil, nil
}

// FindBlockByHash find the specified block by hash.
func (b *BlockTree) FindBlockByHash(hash common.Hash) *types.Block {
	for _, extMap := range b.blocks {
		for h, ext := range extMap {
			if h == hash {
				return ext.block
			}
		}
	}
	return nil
}

func (b *BlockTree) pruneBranch(ext *BlockExt, clearFn func(*types.Block)) {
	for h, e := range ext.children {
		if extMap, ok := b.blocks[e.block.NumberU64()]; ok {
			if clearFn != nil {
				clearFn(e.block)
			}
			delete(extMap, h)
			b.pruneBranch(ext, clearFn)
		}
	}
}

func (b *BlockTree) insertBlock(ext *BlockExt) (*types.Block, *types.Block) {
	number := ext.block.NumberU64()
	hash := ext.block.Hash()
	if extMap, ok := b.blocks[number]; ok {
		if ext, ok := extMap[hash]; !ok {
			extMap[hash] = ext
		}
	} else {
		extMap := make(map[common.Hash]*BlockExt)
		extMap[hash] = ext
		b.blocks[number] = extMap
	}

	b.fixTree(ext)

	return b.commitBlock(b.maxBlock(ext))
}

// Return LockQC, Commit Blocks
func (b *BlockTree) commitBlock(ext *BlockExt) (*types.Block, *types.Block) {
	lock := ext.parent
	var commit *BlockExt
	if lock != nil {
		commit = lock.parent
	}
	return lock.block, commit.block
}

// Returns the maximum view number block for a given height
func (b *BlockTree) maxBlock(ext *BlockExt) *BlockExt {
	max := ext
	if extMap, ok := b.blocks[ext.block.NumberU64()]; ok {
		for _, e := range extMap {
			if max.viewNumber < e.viewNumber {
				max = e
			}
		}
	}
	return max
}

// Connect parent and child blocks
func (b *BlockTree) fixTree(ext *BlockExt) {
	parent := b.findParent(ext.block.ParentHash(), ext.block.NumberU64())
	child := b.findChild(ext.block.Hash(), ext.block.NumberU64())

	if parent != nil {
		parent.children[ext.block.Hash()] = ext
		ext.parent = parent
	}

	if child != nil {
		child.parent = ext
		ext.children[child.block.Hash()] = child
	}
}

func (b *BlockTree) findParent(hash common.Hash, number uint64) *BlockExt {
	if extMap, ok := b.blocks[number-1]; ok {
		for _, v := range extMap {
			if v.block != nil {
				if v.block.Hash() == hash {
					return v
				}
			}
		}
	}
	return nil
}

func (b *BlockTree) findChild(hash common.Hash, number uint64) *BlockExt {
	if extMap, ok := b.blocks[number+1]; ok {
		for _, v := range extMap {
			if v.block != nil {
				if v.block.ParentHash() == hash {
					return v
				}
			}
		}
	}
	return nil
}
