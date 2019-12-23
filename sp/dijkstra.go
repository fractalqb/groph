package sp

import (
	"container/heap"
	"math"

	"git.fractalqb.de/fractalqb/groph"
)

type pqItemf32 struct {
	v groph.VIdx
	p float32
}

type prioQf32 struct {
	v2i []int
	is  []pqItemf32
}

func newPrioQf32(size groph.VIdx) *prioQf32 {
	res := &prioQf32{
		v2i: make([]int, size),
		is:  make([]pqItemf32, 0, size),
	}
	return res
}

func (pq *prioQf32) Len() int { return len(pq.is) }

func (pq *prioQf32) Less(i, j int) bool { return pq.is[i].p < pq.is[j].p }

func (pq *prioQf32) Swap(i, j int) {
	ii, ij := pq.is[i], pq.is[j]
	pq.v2i[ii.v], pq.v2i[ij.v] = j, i
	pq.is[i], pq.is[j] = ij, ii
}

func (pq *prioQf32) Push(x interface{}) {
	item := x.(pqItemf32)
	pq.v2i[item.v] = len(pq.is)
	pq.is = append(pq.is, item)
}

func (pq *prioQf32) Pop() interface{} {
	n := len(pq.is) - 1
	res := pq.is[n]
	pq.is = pq.is[:n]
	return res
}

func (pq *prioQf32) update(v groph.VIdx, priority float32) {
	i := pq.v2i[v]
	pq.is[i].p = priority
	heap.Fix(pq, i)
}

func Dijkstraf32(g groph.RGf32, start groph.VIdx) (dist []float32, prev []groph.VIdx) {
	vertexNo := g.VertexNo()
	dist = make([]float32, vertexNo)
	prev = make([]groph.VIdx, vertexNo)
	q := *newPrioQf32(vertexNo)
	for v := groph.VIdx(0); v < g.VertexNo(); v++ {
		if v != start {
			dist[v] = float32(math.Inf(1))
		}
		prev[v] = -1
		heap.Push(&q, pqItemf32{v, dist[v]})
	}
	for q.Len() != 0 {
		u := heap.Pop(&q).(pqItemf32).v
		groph.EachNeighbour(g, u, func(v groph.VIdx, _ bool, w interface{}) {
			alt := dist[u] + w.(float32)
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				q.update(v, alt)
			}
		})
	}
	return dist, prev
}
