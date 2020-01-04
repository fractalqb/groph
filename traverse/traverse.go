package traverse

import (
	"container/heap"
	"sort"

	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type VisitVertex = func(pred, v groph.VIdx, circle bool) (stop bool)

type VisitInCluster = func(pred, v groph.VIdx, circle bool, cluster int) (stop bool)

// Search performs depth-first or breadth-breadth searches of
// groph.RGraph objects.
type Search struct {
	g     groph.RGraph
	mem   []groph.Edge
	tail  int
	visit hitPq
	// If not nil SortBy is used to sort the neighbours v of node u. SortBy
	// returns true if the edge (u,v1) is less than (u,v2).
	SortBy func(u, v1, v2 groph.VIdx) bool
}

func NewSearch(g groph.RGraph) *Search {
	res := &Search{g: g}
	res.visit.reset(g.Order())
	return res
}

func (df *Search) Reset(g groph.RGraph) {
	df.g = g
	df.visit.reset(g.Order())
}

type stepFn = func(g groph.RGraph, v groph.VIdx, on groph.VisitVertex) bool

// Depth1stAt performs a depth-first search on the respective
// RGraph. The search starts at vertex start and terminates when no
// further vertices can be reached via an edge of the graph. It
// returns the number of distinct vertex hits during the search and
// an indicator if the visit function 'do' stopped the search.
func (df *Search) AdjDepth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.d1stAt(start, do, groph.EachAdjacent)
}

func (df *Search) OutDepth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.d1stAt(start, do, groph.EachOutgoing)
}

func (df *Search) InDepth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.d1stAt(start, do, groph.EachIncoming)
}

func (df *Search) d1stAt(
	start groph.VIdx,
	do VisitVertex,
	eachNext stepFn,
) (hits int, stopped bool) {
	if h := df.visit.hits(start); h > 0 {
		return 0, false
	}
	if df.mem != nil {
		df.mem = df.mem[:0]
	}
	step := groph.Edge{U: -1, V: start}
	df.push(step)
	df.visit.setHits(start, 1)
	var onDest func(d groph.VIdx) (stop bool)
	if groph.Directed(df.g) {
		onDest = func(d groph.VIdx) bool { return df.checkDDest(step, d, do) }
	} else {
		onDest = func(d groph.VIdx) bool { return df.checkUDest(step, d, do) }
	}
	for len(df.mem) > 0 {
		step = df.pop()
		if do(step.U, step.V, false) {
			return hits + 1, true
		}
		hits++
		sortStart := len(df.mem)
		if eachNext(df.g, step.V, onDest) {
			return hits, true
		}
		if df.SortBy != nil {
			sort.Slice(df.mem[sortStart:], func(v1, v2 int) bool {
				return !df.SortBy(step.V, v1, v2)
			})
		}
	}
	return hits, false
}

func (df *Search) AdjDepth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.d1st(stopToNextCluster, do, groph.EachAdjacent)
}

func (df *Search) OutDepth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.d1st(stopToNextCluster, do, groph.EachOutgoing)
}

func (df *Search) InDepth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.d1st(stopToNextCluster, do, groph.EachIncoming)
}

func (df *Search) d1st(
	stopToNextCluster bool,
	do VisitInCluster,
	eachNext stepFn,
) (stopped bool) {
	cluster := 0
	cdo := func(p, n groph.VIdx, c bool) bool { return do(p, n, c, cluster) }
	hits, stop := df.d1stAt(0, cdo, eachNext)
	if stop {
		if !stopToNextCluster {
			return true
		}
		cluster = -1
	}
	for hits < df.g.Order() {
		if cluster >= 0 {
			cluster++
		}
		start := df.visit.top() // TODO assert hits(start) == 0
		n, stop := df.d1stAt(start, cdo, eachNext)
		if stop {
			if !stopToNextCluster {
				return true
			}
			cluster = -1
		}
		hits += n
	}
	return false
}

// Breadth1stAt performs a breadth-first search on the respective
// RGraph. The search starts at vertex start and terminates when no
// further vertices can be reached via an edge of the graph. It
// returns the number of distinct vertex hits during the search and an
// indicator if the visit function 'do' stopped the search.

func (df *Search) AdjBreadth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.b1stAt(start, do, groph.EachAdjacent)
}

func (df *Search) OutBreadth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.b1stAt(start, do, groph.EachOutgoing)
}

func (df *Search) InBreadth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.b1stAt(start, do, groph.EachIncoming)
}

func (df *Search) b1stAt(
	start groph.VIdx,
	do VisitVertex,
	eachNext stepFn,
) (hits int, stopped bool) {
	if h := df.visit.hits(start); h > 0 {
		return 0, false
	}
	if df.mem != nil {
		df.mem = df.mem[:0]
	}
	df.tail = 0
	step := groph.Edge{U: -1, V: start}
	df.push(step)
	df.visit.setHits(start, 1)
	var onDest func(d groph.VIdx) (stop bool)
	if groph.Directed(df.g) {
		onDest = func(d groph.VIdx) bool { return df.checkDDest(step, d, do) }
	} else {
		onDest = func(d groph.VIdx) bool { return df.checkUDest(step, d, do) }
	}
	for df.tail < len(df.mem) {
		step = df.take()
		if do(step.U, step.V, false) {
			return hits + 1, true
		}
		hits++
		sortStart := len(df.mem)
		if eachNext(df.g, step.V, onDest) {
			return hits, true
		}
		if df.SortBy != nil {
			sort.Slice(df.mem[sortStart:], func(v1, v2 int) bool {
				return df.SortBy(step.V, v1, v2)
			})
		}
	}
	return hits, false
}

func (df *Search) AdjBreadth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.b1st(stopToNextCluster, do, groph.EachAdjacent)
}

func (df *Search) OutBreadth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.b1st(stopToNextCluster, do, groph.EachOutgoing)
}

func (df *Search) InBreadth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.b1st(stopToNextCluster, do, groph.EachIncoming)
}

func (df *Search) b1st(
	stopToNextCluster bool,
	do VisitInCluster,
	eachNext stepFn,
) (stopped bool) {
	cluster := 0
	cdo := func(p, n groph.VIdx, c bool) bool { return do(p, n, c, cluster) }
	hits, stop := df.b1stAt(0, cdo, eachNext)
	if stop {
		if !stopToNextCluster {
			return true
		}
		cluster = -1
	}
	for hits < df.g.Order() {
		if cluster >= 0 {
			cluster++
		}
		start := df.visit.top() // TODO assert hits(start) == 0
		n, stop := df.b1stAt(start, cdo, eachNext)
		if stop {
			if !stopToNextCluster {
				return true
			}
			cluster = -1
		}
		hits += n
	}
	return false
}

func (df *Search) Hits(v groph.VIdx) int {
	if v >= len(df.visit.v2i) {
		return 0
	}
	ii := df.visit.v2i[v]
	if ii < 0 {
		return 0
	}
	return df.visit.is[ii].hits
}

func (df *Search) LeastHits() (v groph.VIdx, hits int) {
	if len(df.visit.is) == 0 {
		return -1, -1
	}
	item := df.visit.is[0]
	return item.v, item.hits
}

func (df *Search) checkDDest(pre groph.Edge, d groph.VIdx, do VisitVertex) bool {
	h := df.visit.hits(d)
	if h == 0 {
		df.push(groph.Edge{U: pre.V, V: d})
	} else {
		return do(pre.V, d, true)
	}
	df.visit.setHits(d, h+1)
	return false
}

func (df *Search) checkUDest(pre groph.Edge, d groph.VIdx, do VisitVertex) bool {
	if d == pre.U || d == pre.V {
		return false
	}
	return df.checkDDest(pre, d, do)
}

func (df *Search) push(e groph.Edge) {
	df.mem = append(df.mem, e)
}

func (df *Search) pop() (step groph.Edge) {
	l := len(df.mem) - 1
	step = df.mem[l]
	df.mem = df.mem[:l]
	return step
}

func (df *Search) take() (step groph.Edge) {
	step = df.mem[df.tail]
	df.tail++
	return step
}

type hitPqItem struct {
	hits int
	v    groph.VIdx
}

type hitPq struct {
	v2i []int
	is  []hitPqItem
}

func (pq *hitPq) reset(ord groph.VIdx) {
	pq.v2i = iutil.IntSlice(pq.v2i, ord)
	if pq.is == nil || cap(pq.is) < ord {
		pq.is = make([]hitPqItem, ord)
	} else {
		pq.is = pq.is[:ord]
	}
	for i := range pq.is {
		pq.v2i[i] = i
		pq.is[i] = hitPqItem{hits: 0, v: i}
	}
}

func (pq *hitPq) top() groph.VIdx { return pq.is[0].v }

func (pq *hitPq) hits(v groph.VIdx) int { return pq.is[pq.v2i[v]].hits }

func (pq *hitPq) setHits(v groph.VIdx, hits int) {
	ii := pq.v2i[v]
	pq.is[ii].hits = hits
	heap.Fix(pq, ii)
}

func (pq *hitPq) Len() int { return len(pq.is) }

func (pq *hitPq) Less(i, j int) bool { return pq.is[i].hits < pq.is[j].hits }

func (pq *hitPq) Swap(i, j int) {
	ii, ji := pq.is[i], pq.is[j]
	pq.is[i], pq.is[j] = ji, ii
	pq.v2i[ii.v] = j
	pq.v2i[ji.v] = i
}

func (pq *hitPq) Push(x interface{}) {
	panic("must not be called")
	// pqItem := x.(hitPqItem)
	// for pqItem.v >= len(pq.v2i) {
	// 	pq.v2i = append(pq.v2i, -1)
	// }
	// pq.v2i[pqItem.v] = len(pq.is)
	// pq.is = append(pq.is, pqItem)
}

func (pq *hitPq) Pop() interface{} {
	panic("must not be called")
	// lm1 := len(pq.is) - 1
	// res := pq.is[lm1]
	// pq.is = pq.is[:lm1]
	// pq.v2i[res.v] = -1
	// return res
}
