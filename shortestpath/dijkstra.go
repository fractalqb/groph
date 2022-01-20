package shortestpath

import (
	"container/heap"
	"math"

	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type pqItemBool struct {
	v groph.VIdx
	p int
}

type pqbool struct {
	v2i []int
	is  []pqItemBool
}

func (pq *pqbool) Len() int { return len(pq.is) }

func (pq *pqbool) Less(i, j int) bool {
	pi, pj := pq.is[i].p, pq.is[j].p
	if pi < 0 {
		return false
	}
	if pj < 0 {
		return true
	}
	return pi < pj
}

func (pq *pqbool) Swap(i, j int) {
	ii, ij := pq.is[i], pq.is[j]
	pq.is[i], pq.is[j] = ij, ii
	pq.v2i[ii.v] = j
	pq.v2i[ij.v] = i
}

func (pq *pqbool) Push(x interface{}) {
	item := x.(pqItemBool)
	pq.v2i[item.v] = len(pq.is)
	pq.is = append(pq.is, item)
}

func (pq *pqbool) Pop() interface{} {
	// TODO what about pq.v2i ?
	n := len(pq.is) - 1
	res := pq.is[n]
	pq.is = pq.is[:n]
	return res
}

func (dij *pqbool) init(ord int) {
	dij.v2i = iutil.IntSlice(dij.v2i, ord)
	if dij.is == nil || cap(dij.is) < ord {
		dij.is = make([]pqItemBool, 0, ord)
	} else {
		dij.is = dij.is[:0]
	}
}

func (pq *pqbool) update(v groph.VIdx, priority int) {
	i := pq.v2i[v]
	pq.is[i].p = priority
	heap.Fix(pq, i)
}

type DijkstraBool struct {
	pq pqbool
}

func (dij *DijkstraBool) init(ord int) { dij.pq.init(ord) }

func (dij *DijkstraBool) On(
	g groph.RGbool,
	start groph.VIdx,
	distReuse []int,
	treeReuse []groph.VIdx,
) ([]int, groph.Tree) {
	order := g.Order()
	distReuse = iutil.IntSlice(distReuse, order)
	if treeReuse != nil {
		treeReuse = iutil.VIdxSlice(treeReuse, order)
	}
	dij.init(order)
	distReuse[start] = 0
	for v := 0; v < g.Order(); v++ {
		if v != start {
			distReuse[v] = -1
		}
		if treeReuse != nil {
			treeReuse[v] = -1
		}
		heap.Push(&dij.pq, pqItemBool{v, distReuse[v]})
	}
	for dij.pq.Len() != 0 {
		u := heap.Pop(&dij.pq).(pqItemBool).v
		groph.EachOutgoing(g, u, func(n groph.VIdx) (stop bool) {
			alt := distReuse[u]
			if alt < 0 {
				return false
			}
			e := g.Edge(u, n)
			if !e {
				return false
			}
			alt++
			if distReuse[n] < 0 || alt < distReuse[n] {
				distReuse[n] = alt
				if treeReuse != nil {
					treeReuse[n] = u
				}
				dij.pq.update(n, alt)
			}
			return false
		})
	}
	return distReuse, treeReuse
}

type pqItemI32 struct {
	v groph.VIdx
	p int32
}

type pqi32 struct {
	v2i []int
	is  []pqItemI32
}

func (pq *pqi32) Len() int { return len(pq.is) }

func (pq *pqi32) Less(i, j int) bool {
	pi, pj := pq.is[i].p, pq.is[j].p
	if pi < 0 {
		return false
	}
	if pj < 0 {
		return true
	}
	return pi < pj
}

func (pq *pqi32) Swap(i, j int) {
	ii, ij := pq.is[i], pq.is[j]
	pq.is[i], pq.is[j] = ij, ii
	pq.v2i[ii.v] = j
	pq.v2i[ij.v] = i
}

func (pq *pqi32) Push(x interface{}) {
	item := x.(pqItemI32)
	pq.v2i[item.v] = len(pq.is)
	pq.is = append(pq.is, item)
}

func (pq *pqi32) Pop() interface{} {
	// TODO what about pq.v2i ?
	n := len(pq.is) - 1
	res := pq.is[n]
	pq.is = pq.is[:n]
	return res
}

func (dij *pqi32) init(ord int) {
	dij.v2i = iutil.IntSlice(dij.v2i, ord)
	if dij.is == nil || cap(dij.is) < ord {
		dij.is = make([]pqItemI32, 0, ord)
	} else {
		dij.is = dij.is[:0]
	}
}

func (pq *pqi32) update(v groph.VIdx, priority int32) {
	i := pq.v2i[v]
	pq.is[i].p = priority
	heap.Fix(pq, i)
}

type DijkstraI32 struct {
	pq pqi32
}

func (dij *DijkstraI32) init(ord int) { dij.pq.init(ord) }

func (dij *DijkstraI32) On(
	g groph.RGi32,
	start groph.VIdx,
	distReuse []int32,
	prevReuse []groph.VIdx,
) ([]int32, groph.Tree) {
	order := g.Order()
	distReuse = iutil.I32Slice(distReuse, order)
	if prevReuse != nil {
		prevReuse = iutil.VIdxSlice(prevReuse, order)
	}
	dij.init(order)
	distReuse[start] = 0
	for v := 0; v < g.Order(); v++ {
		if v != start {
			distReuse[v] = -1
		}
		if prevReuse != nil {
			prevReuse[v] = -1
		}
		heap.Push(&dij.pq, pqItemI32{v, distReuse[v]})
	}
	for dij.pq.Len() != 0 {
		u := heap.Pop(&dij.pq).(pqItemI32).v
		groph.EachOutgoing(g, u, func(n groph.VIdx) (stop bool) {
			alt := distReuse[u]
			if alt < 0 {
				return false
			}
			e, ok := g.Edge(u, n)
			if !ok {
				return false
			}
			alt += e
			if distReuse[n] < 0 || alt < distReuse[n] {
				distReuse[n] = alt
				if prevReuse != nil {
					prevReuse[n] = u
				}
				dij.pq.update(n, alt)
			}
			return false
		})
	}
	return distReuse, prevReuse
}

type pqItemF32 struct {
	v groph.VIdx
	p float32
}

type pqf32 struct {
	v2i []int
	is  []pqItemF32
}

func (pq *pqf32) Len() int { return len(pq.is) }

func (pq *pqf32) Less(i, j int) bool { return pq.is[i].p < pq.is[j].p }

func (pq *pqf32) Swap(i, j int) {
	ii, ij := pq.is[i], pq.is[j]
	pq.v2i[ii.v], pq.v2i[ij.v] = j, i
	pq.is[i], pq.is[j] = ij, ii
}

func (pq *pqf32) Push(x interface{}) {
	item := x.(pqItemF32)
	pq.v2i[item.v] = len(pq.is)
	pq.is = append(pq.is, item)
}

func (pq *pqf32) Pop() interface{} {
	n := len(pq.is) - 1
	res := pq.is[n]
	pq.is = pq.is[:n]
	return res
}

func (dij *pqf32) init(ord int) {
	dij.v2i = iutil.IntSlice(dij.v2i, ord)
	if dij.is == nil || cap(dij.is) < ord {
		dij.is = make([]pqItemF32, 0, ord)
	} else {
		dij.is = dij.is[:0]
	}
}

func (pq *pqf32) update(v groph.VIdx, priority float32) {
	i := pq.v2i[v]
	pq.is[i].p = priority
	heap.Fix(pq, i)
}

type DijkstraF32 struct {
	pq pqf32
}

func (dij *DijkstraF32) init(ord int) { dij.pq.init(ord) }

func (dij *DijkstraF32) On(
	g groph.RGf32,
	start groph.VIdx,
	distReuse []float32,
	treeReuse []groph.VIdx,
) ([]float32, groph.Tree) {
	order := g.Order()
	distReuse = iutil.F32Slice(distReuse, order)
	if treeReuse != nil {
		treeReuse = iutil.VIdxSlice(treeReuse, order)
	}
	dij.init(order)
	distReuse[start] = 0
	for v := 0; v < g.Order(); v++ {
		if v != start {
			distReuse[v] = float32(math.Inf(1))
		}
		if treeReuse != nil {
			treeReuse[v] = -1
		}
		heap.Push(&dij.pq, pqItemF32{v, distReuse[v]})
	}
	for dij.pq.Len() != 0 {
		u := heap.Pop(&dij.pq).(pqItemF32).v
		groph.EachOutgoing(g, u, func(n groph.VIdx) (stop bool) {
			alt := distReuse[u] + g.Edge(u, n) // TODO EdgeU?
			if alt < distReuse[n] {
				distReuse[n] = alt
				if treeReuse != nil {
					treeReuse[n] = u
				}
				dij.pq.update(n, alt)
			}
			return false
		})
	}
	return distReuse, treeReuse
}
