package trie

import (
	"container/list"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/cespare/xxhash"
	"github.com/panjf2000/ants/v2"
)

type Vertex2 struct {
	inDegree uint32
	outEdge  []uint64
}

type DAG2 struct {
	vtxs     map[uint64]*Vertex2
	topLevel *list.List

	lock sync.Mutex
	cv   *sync.Cond

	totalVertexs  uint32
	totalConsumed uint32
}

func NewDAG2() *DAG2 {
	dag := &DAG2{
		vtxs:          make(map[uint64]*Vertex2),
		topLevel:      list.New(),
		totalConsumed: 0,
	}
	dag.cv = sync.NewCond(&dag.lock)

	return dag
}

func (d *DAG2) addVertex(id uint64) {
	if _, ok := d.vtxs[id]; !ok {
		d.vtxs[id] = &Vertex2{
			inDegree: 0,
			outEdge:  make([]uint64, 0),
		}
		d.totalVertexs++
	}
}

func (d *DAG2) delVertex(id uint64) {
	if _, ok := d.vtxs[id]; ok {
		d.totalVertexs--
		delete(d.vtxs, id)
	}
}

func (d *DAG2) addEdge(from, to uint64) {
	if _, ok := d.vtxs[from]; !ok {
		d.vtxs[from] = &Vertex2{
			inDegree: 0,
			outEdge:  make([]uint64, 0),
		}
		d.totalVertexs++
	}
	vtx := d.vtxs[from]
	found := false
	for _, t := range vtx.outEdge {
		if t == to {
			found = true
			break
		}
	}
	if !found {
		vtx.outEdge = append(vtx.outEdge, to)
	}

	/*
		if _, ok := d.vtxs[to]; !ok {
			d.vtxs[to] = &Vertex{
				inDegree: 0,
				outEdge:  make([]uint64, 0),
			}
			d.totalVertexs++
		}
		d.vtxs[to].inDegree += 1
	*/
}

func (d *DAG2) delEdge(id uint64) {
	if _, ok := d.vtxs[id]; ok {
		for _, k := range d.vtxs[id].outEdge {
			if vtx, found := d.vtxs[k]; found {
				vtx.inDegree--
			}
		}
	}
}

func (d *DAG2) generate() {
	for id, v := range d.vtxs {
		for _, pid := range v.outEdge {
			if d.vtxs[pid] == nil {
				panic(fmt.Sprintf("not found, id: %d, pid: %d", id, pid))
			}
			d.vtxs[pid].inDegree++
		}
	}

	for k, v := range d.vtxs {
		if v.inDegree == 0 {
			d.topLevel.PushBack(k)
		}
	}
}

func (d *DAG2) waitPop() uint64 {
	if d.hasFinished() {
		return invalidId
	}

	d.cv.L.Lock()
	defer d.cv.L.Unlock()
	for d.topLevel.Len() == 0 && !d.hasFinished() {
		d.cv.Wait()
	}

	if d.hasFinished() || d.topLevel.Len() == 0 {
		return invalidId
	}

	el := d.topLevel.Front()
	id := el.Value.(uint64)
	d.topLevel.Remove(el)
	return id
}

func (d *DAG2) hasFinished() bool {
	return d.totalConsumed >= d.totalVertexs
}

func (d *DAG2) consume(id uint64) uint64 {
	producedNum := 0
	var nextID uint64 = invalidId
	var degree uint32 = 0

	for _, k := range d.vtxs[id].outEdge {
		vtx := d.vtxs[k]
		degree = atomic.AddUint32(&vtx.inDegree, ^uint32(0))
		//fmt.Printf("id: %d k: %d degree: %d consumed: %d total: %d\n", id, k, degree, atomic.LoadUint32(&d.totalConsumed), d.totalVertexs)
		if degree == 0 {
			producedNum += 1
			if producedNum == 1 {
				nextID = k
			} else {
				d.lock.Lock()
				d.topLevel.PushBack(k)
				d.lock.Unlock()
			}
		}
	}

	if atomic.AddUint32(&d.totalConsumed, 1) == d.totalVertexs {
		d.cv.L.Lock()
		d.cv.Broadcast()
		d.cv.L.Unlock()
	}

	//fmt.Printf("id: %d nextId: %d consumed: %d total: %d\n", id, nextID, atomic.LoadUint32(&d.totalConsumed), d.totalVertexs)
	return nextID
}

func (d *DAG2) clear() {
	d.vtxs = make(map[uint64]*Vertex2)
	d.topLevel = list.New()
	d.totalConsumed = 0
	d.totalVertexs = 0
}

type DAGNode2 struct {
	collapsed node
	cached    node
	pid       uint64
	idx       int
}

// TrieDAGV2
type TrieDAGV2 struct {
	nodes map[uint64]*DAGNode2

	dag *DAG2

	cachegen   uint16
	cachelimit uint16
}

func newTrieDAGV2() *TrieDAGV2 {
	return &TrieDAGV2{
		nodes: make(map[uint64]*DAGNode2),
		dag:   NewDAG2(),
	}
}

func (td *TrieDAGV2) addVertexAndEdge(pprefix, prefix []byte, n node) {
	var pid uint64
	if len(pprefix) > 0 {
		pid = xxhash.Sum64(pprefix)
	}

	switch nc := n.(type) {
	case *shortNode:
		collapsed, cached := nc.copy(), nc.copy()
		collapsed.Key = hexToCompact(nc.Key)
		cached.Key = common.CopyBytes(nc.Key)

		id := xxhash.Sum64(append(prefix, nc.Key...))
		td.nodes[id] = &DAGNode2{
			collapsed: collapsed,
			cached:    cached,
			pid:       pid,
		}
		if len(prefix) > 0 {
			td.nodes[id].idx = int(prefix[len(prefix)-1])
		}
		td.dag.addVertex(id)

		if pid > 0 {
			td.dag.addEdge(id, pid)
		}

	case *fullNode:
		collapsed, cached := nc.copy(), nc.copy()

		dagNode := &DAGNode2{
			collapsed: collapsed,
			cached:    cached,
			pid:       pid,
		}
		if len(prefix) > 0 {
			dagNode.idx = int(prefix[len(prefix)-1])
		}

		id := xxhash.Sum64(append(prefix, fullNodeSuffix...))
		td.nodes[id] = dagNode
		td.dag.addVertex(id)
		if pid > 0 {
			td.dag.addEdge(id, pid)
		}
	}
}

func (td *TrieDAGV2) delVertexAndEdge(key []byte) {
	id := xxhash.Sum64(key)
	td.dag.delEdge(id)
	td.dag.delVertex(id)
	delete(td.nodes, id)
}

func (td *TrieDAGV2) replaceEdge(old, new []byte) {
	opid := xxhash.Sum64(old)
	npid := xxhash.Sum64(new)

	for id, vtx := range td.dag.vtxs {
		for _, pid := range vtx.outEdge {
			if opid == pid {
				vtx.outEdge = make([]uint64, 0)
				vtx.outEdge = append(vtx.outEdge, npid)
				td.nodes[id].pid = npid
			}
		}
	}
}

func (td *TrieDAGV2) clear() {
	td.dag.clear()
	td.nodes = make(map[uint64]*DAGNode2)
}

func (td *TrieDAGV2) hash(db *Database, force bool, onleaf LeafCallback) (node, node, error) {
	var wg sync.WaitGroup
	var errDone common.AtomicBool
	var e atomic.Value // error
	var resHash node = hashNode{}
	var newRoot node
	numCPU := runtime.NumCPU()

	cachedHash := func(n, c node) (node, node, bool) {
		if hash, dirty := n.cache(); len(hash) != 0 {
			if db == nil {
				return hash, c, true
			}

			if n.canUnload(td.cachegen, td.cachelimit) {
				cacheUnloadCounter.Inc(1)
				return hash, hash, true
			}
			if !dirty {
				return hash, c, true
			}
		}
		return n, n, false
	}

	process := func() {
		hasher := newHasher(td.cachegen, td.cachelimit, onleaf)

		id := td.dag.waitPop()
		if id == invalidId {
			returnHasherToPool(hasher)
			wg.Done()
			return
		}

		var hashed node
		var cached node
		var err error
		var hasCache bool
		for id != invalidId {
			n := td.nodes[id]

			tmpForce := false
			if n.pid == 0 {
				tmpForce = force
			}

			hashed, cached, hasCache = cachedHash(n.collapsed, n.cached)
			if !hasCache {
				hashed, err = hasher.store(n.collapsed, db, tmpForce)
				if err != nil {
					e.Store(err)
					errDone.Set(true)
					break
				}
				cached = n.cached
			}

			if n.pid > 0 {
				p := td.nodes[n.pid]
				switch ptype := p.collapsed.(type) {
				case *shortNode:
					ptype.Val = hashed
				case *fullNode:
					ptype.Children[n.idx] = hashed
				}

				switch nc := p.cached.(type) {
				case *shortNode:
					nc.Val = cached
				case *fullNode:
					nc.Children[n.idx] = cached
				}
			}

			cachedHash, _ := hashed.(hashNode)
			switch cn := n.cached.(type) {
			case *shortNode:
				*cn.flags.hash = cachedHash
				if db != nil {
					*cn.flags.dirty = false
				}
			case *fullNode:
				*cn.flags.hash = cachedHash
				if db != nil {
					*cn.flags.dirty = false
				}
			}

			id = td.dag.consume(id)
			if n.pid == 0 {
				resHash = hashed
				newRoot = n.cached
				break
			}

			if errDone.IsSet() {
				break
			}

			if id == invalidId && !td.dag.hasFinished() {
				id = td.dag.waitPop()
			}
		}
		returnHasherToPool(hasher)
		wg.Done()
	}

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		_ = ants.Submit(process)
	}

	wg.Wait()

	if e.Load() != nil && e.Load().(error) != nil {
		return hashNode{}, nil, e.Load().(error)
	}
	return resHash, newRoot, nil
}

func (td *TrieDAGV2) init() {
	td.dag.generate()
}
