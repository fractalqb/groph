package sp

import (
	"container/heap"
	"math"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/internal/util"
)

type pqItemi32 struct {
	v groph.VIdx
	p int32
}

type Dijkstrai32 struct {
	v2i []int
	is  []pqItemi32
}

func (dij *Dijkstrai32) init(vno int) {
	dij.v2i = util.IntSlice(dij.v2i, vno)
	if dij.is == nil || cap(dij.is) < vno {
		dij.is = make([]pqItemi32, 0, vno)
	} else {
		dij.is = dij.is[:0]
	}
}

func (pq *Dijkstrai32) Len() int { return len(pq.is) }

func (pq *Dijkstrai32) Less(i, j int) bool {
	pi, pj := pq.is[i].p, pq.is[j].p
	if pi < 0 {
		return false
	}
	if pj < 0 {
		return true
	}
	return pi < pj
}

func (pq *Dijkstrai32) Swap(i, j int) {
	ii, ij := pq.is[i], pq.is[j]
	pq.v2i[ii.v], pq.v2i[ij.v] = j, i
	pq.is[i], pq.is[j] = ij, ii
}

func (pq *Dijkstrai32) Push(x interface{}) {
	item := x.(pqItemi32)
	pq.v2i[item.v] = len(pq.is)
	pq.is = append(pq.is, item)
}

func (pq *Dijkstrai32) Pop() interface{} {
	n := len(pq.is) - 1
	res := pq.is[n]
	pq.is = pq.is[:n]
	return res
}

func (pq *Dijkstrai32) update(v groph.VIdx, priority int32) {
	i := pq.v2i[v]
	pq.is[i].p = priority
	heap.Fix(pq, i)
}

func (dij *Dijkstrai32) On(
	g groph.RGi32,
	start groph.VIdx,
	dist []int32,
	prev []groph.VIdx,
) ([]int32, []groph.VIdx) {
	vertexNo := g.VertexNo()
	dist = util.I32Slice(dist, vertexNo)
	if prev != nil {
		prev = util.VIdxSlice(prev, vertexNo)
	}
	dij.init(vertexNo)
	dist[start] = 0
	for v := groph.VIdx(0); v < g.VertexNo(); v++ {
		if v != start {
			dist[v] = -1
		}
		if prev != nil {
			prev[v] = -1
		}
		heap.Push(dij, pqItemi32{v, dist[v]})
	}
	for dij.Len() != 0 {
		u := heap.Pop(dij).(pqItemi32).v
		groph.EachNeighbour(g, u, func(n groph.VIdx) {
			alt := dist[u]
			if alt < 0 {
				return
			}
			alt += g.Edge(u, n) // TODO EdgeD?
			if dist[n] < 0 || alt < dist[n] {
				dist[n] = alt
				if prev != nil {
					prev[n] = u
				}
				dij.update(n, alt)
			}
		})
	}
	return dist, prev
}

type pqItemf32 struct {
	v groph.VIdx
	p float32
}

type Dijkstraf32 struct {
	v2i []int
	is  []pqItemf32
}

func (dij *Dijkstraf32) init(vno int) {
	dij.v2i = util.IntSlice(dij.v2i, vno)
	if dij.is == nil || cap(dij.is) < vno {
		dij.is = make([]pqItemf32, 0, vno)
	} else {
		dij.is = dij.is[:0]
	}
}

func (pq *Dijkstraf32) Len() int { return len(pq.is) }

func (pq *Dijkstraf32) Less(i, j int) bool { return pq.is[i].p < pq.is[j].p }

func (pq *Dijkstraf32) Swap(i, j int) {
	ii, ij := pq.is[i], pq.is[j]
	pq.v2i[ii.v], pq.v2i[ij.v] = j, i
	pq.is[i], pq.is[j] = ij, ii
}

func (pq *Dijkstraf32) Push(x interface{}) {
	item := x.(pqItemf32)
	pq.v2i[item.v] = len(pq.is)
	pq.is = append(pq.is, item)
}

func (pq *Dijkstraf32) Pop() interface{} {
	n := len(pq.is) - 1
	res := pq.is[n]
	pq.is = pq.is[:n]
	return res
}

func (pq *Dijkstraf32) update(v groph.VIdx, priority float32) {
	i := pq.v2i[v]
	pq.is[i].p = priority
	heap.Fix(pq, i)
}

func (dij *Dijkstraf32) On(
	g groph.RGf32,
	start groph.VIdx,
	dist []float32,
	prev []groph.VIdx,
) ([]float32, []groph.VIdx) {
	vertexNo := g.VertexNo()
	dist = util.F32Slice(dist, vertexNo)
	if prev != nil {
		prev = util.VIdxSlice(prev, vertexNo)
	}
	dij.init(vertexNo)
	dist[start] = 0
	for v := groph.VIdx(0); v < g.VertexNo(); v++ {
		if v != start {
			dist[v] = float32(math.Inf(1))
		}
		if prev != nil {
			prev[v] = -1
		}
		heap.Push(dij, pqItemf32{v, dist[v]})
	}
	for dij.Len() != 0 {
		u := heap.Pop(dij).(pqItemf32).v
		groph.EachNeighbour(g, u, func(n groph.VIdx) {
			alt := dist[u] + g.Edge(u, n) // TOOD EdgeD?
			if alt < dist[n] {
				dist[n] = alt
				if prev != nil {
					prev[n] = u
				}
				dij.update(n, alt)
			}
		})
	}
	return dist, prev
}
