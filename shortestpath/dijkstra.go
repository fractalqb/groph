package shortestpath

import (
	"container/heap"
	"math"

	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type pqItemi32 struct {
	v groph.VIdx
	p int32
}

type pqi32 struct {
	v2i []int
	is  []pqItemi32
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
	item := x.(pqItemi32)
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
		dij.is = make([]pqItemi32, 0, ord)
	} else {
		dij.is = dij.is[:0]
	}
}

func (pq *pqi32) update(v groph.VIdx, priority int32) {
	i := pq.v2i[v]
	pq.is[i].p = priority
	heap.Fix(pq, i)
}

type Dijkstrai32 struct {
	pq pqi32
}

func (dij *Dijkstrai32) init(ord int) { dij.pq.init(ord) }

func (dij *Dijkstrai32) On(
	g groph.RGi32,
	start groph.VIdx,
	dist []int32,
	prev []groph.VIdx,
) ([]int32, groph.Tree) {
	order := g.Order()
	dist = iutil.I32Slice(dist, order)
	if prev != nil {
		prev = iutil.VIdxSlice(prev, order)
	}
	dij.init(order)
	dist[start] = 0
	for v := groph.V0; v < g.Order(); v++ {
		if v != start {
			dist[v] = -1
		}
		if prev != nil {
			prev[v] = -1
		}
		heap.Push(&dij.pq, pqItemi32{v, dist[v]})
	}
	for dij.pq.Len() != 0 {
		u := heap.Pop(&dij.pq).(pqItemi32).v
		groph.EachOutgoing(g, u, func(n groph.VIdx) (stop bool) {
			alt := dist[u]
			if alt < 0 {
				return false
			}
			e, _ := g.Edge(u, n) // TODO EdgeD?
			alt += e
			if dist[n] < 0 || alt < dist[n] {
				dist[n] = alt
				if prev != nil {
					prev[n] = u
				}
				dij.pq.update(n, alt)
			}
			return false
		})
	}
	return dist, prev
}

type pqItemf32 struct {
	v groph.VIdx
	p float32
}

type pqf32 struct {
	v2i []int
	is  []pqItemf32
}

func (pq *pqf32) Len() int { return len(pq.is) }

func (pq *pqf32) Less(i, j int) bool { return pq.is[i].p < pq.is[j].p }

func (pq *pqf32) Swap(i, j int) {
	ii, ij := pq.is[i], pq.is[j]
	pq.v2i[ii.v], pq.v2i[ij.v] = j, i
	pq.is[i], pq.is[j] = ij, ii
}

func (pq *pqf32) Push(x interface{}) {
	item := x.(pqItemf32)
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
		dij.is = make([]pqItemf32, 0, ord)
	} else {
		dij.is = dij.is[:0]
	}
}

func (pq *pqf32) update(v groph.VIdx, priority float32) {
	i := pq.v2i[v]
	pq.is[i].p = priority
	heap.Fix(pq, i)
}

type Dijkstraf32 struct {
	pq pqf32
}

func (dij *Dijkstraf32) init(ord int) { dij.pq.init(ord) }

func (dij *Dijkstraf32) On(
	g groph.RGf32,
	start groph.VIdx,
	dist []float32,
	prev []groph.VIdx,
) ([]float32, groph.Tree) {
	order := g.Order()
	dist = iutil.F32Slice(dist, order)
	if prev != nil {
		prev = iutil.VIdxSlice(prev, order)
	}
	dij.init(order)
	dist[start] = 0
	for v := groph.V0; v < g.Order(); v++ {
		if v != start {
			dist[v] = float32(math.Inf(1))
		}
		if prev != nil {
			prev[v] = -1
		}
		heap.Push(&dij.pq, pqItemf32{v, dist[v]})
	}
	for dij.pq.Len() != 0 {
		u := heap.Pop(&dij.pq).(pqItemf32).v
		groph.EachOutgoing(g, u, func(n groph.VIdx) (stop bool) {
			alt := dist[u] + g.Edge(u, n) // TODO EdgeU?
			if alt < dist[n] {
				dist[n] = alt
				if prev != nil {
					prev[n] = u
				}
				dij.pq.update(n, alt)
			}
			return false
		})
	}
	return dist, prev
}
