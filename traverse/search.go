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

func VIdxOrder(_, v1, v2 groph.VIdx) bool { return v1 < v2 }

func NewSearch(g groph.RGraph) *Search {
	res := &Search{g: g}
	res.visit.reset(g.Order())
	return res
}

// Reset prepares the Seach instance for use with graph g.
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
	return df.depth1stAt(start, do, groph.EachAdjacent)
}

func (df *Search) OutDepth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.depth1stAt(start, do, groph.EachOutgoing)
}

func (df *Search) InDepth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.depth1stAt(start, do, groph.EachIncoming)
}

func (df *Search) depth1stAt(
	start groph.VIdx,
	do VisitVertex,
	eachNext stepFn,
) (hits int, stopped bool) {
	if df.mem != nil {
		df.mem = df.mem[:0]
	}
	step := groph.Edge{U: -1, V: start}
	df.push(step)
	var selectNext groph.VisitVertex
	if groph.Directed(df.g) {
		selectNext = func(d groph.VIdx) bool {
			df.push(groph.Edge{U: step.V, V: d})
			return false
		}
	} else {
		selectNext = func(d groph.VIdx) bool {
			if d != step.U {
				df.push(groph.Edge{U: step.V, V: d})
			}
			return false
		}
	}
	for len(df.mem) > 0 {
		step = df.pop()
		h := df.visit.hits(step.V)
		if h > 0 {
			stopped = do(step.U, step.V, true)
			df.visit.setHits(step.V, h+1)
			if stopped {
				return hits, true
			}
		} else {
			stopped = do(step.U, step.V, false)
			df.visit.setHits(step.V, 1)
			hits++
			if stopped {
				return hits, true
			}
			sortStart := len(df.mem)
			eachNext(df.g, step.V, selectNext)
			if df.SortBy != nil {
				sort.Slice(df.mem[sortStart:], func(v1, v2 int) bool {
					return !df.SortBy(step.V, v1, v2)
				})
			}
		}
	}
	return hits, false
}

func (df *Search) AdjDepth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.depth1st(stopToNextCluster, do, groph.EachAdjacent)
}

func (df *Search) OutDepth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.depth1st(stopToNextCluster, do, groph.EachOutgoing)
}

func (df *Search) InDepth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.depth1st(stopToNextCluster, do, groph.EachIncoming)
}

func (df *Search) depth1st(
	stopToNextCluster bool,
	do VisitInCluster,
	eachNext stepFn,
) (stopped bool) {
	cluster := 0
	cdo := func(p, n groph.VIdx, c bool) bool { return do(p, n, c, cluster) }
	hits, stop := df.depth1stAt(0, cdo, eachNext)
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
		n, stop := df.depth1stAt(start, cdo, eachNext)
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
	return df.breadth1stAt(start, do, groph.EachAdjacent)
}

func (df *Search) OutBreadth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.breadth1stAt(start, do, groph.EachOutgoing)
}

func (df *Search) InBreadth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.breadth1stAt(start, do, groph.EachIncoming)
}

func (df *Search) breadth1stAt(
	start groph.VIdx,
	do VisitVertex,
	eachNext stepFn,
) (hits int, stopped bool) {
	if df.mem != nil {
		df.mem = df.mem[:0]
	}
	step := groph.Edge{U: -1, V: start}
	df.push(step)
	df.tail = 0
	var selectNext groph.VisitVertex
	if groph.Directed(df.g) {
		selectNext = func(d groph.VIdx) bool {
			df.push(groph.Edge{U: step.V, V: d})
			return false
		}
	} else {
		selectNext = func(d groph.VIdx) bool {
			if d != step.U {
				df.push(groph.Edge{U: step.V, V: d})
			}
			return false
		}
	}
	for df.tail < len(df.mem) {
		step = df.take()
		h := df.visit.hits(step.V)
		if h > 0 {
			stopped = do(step.U, step.V, true)
			df.visit.setHits(step.V, h+1)
			if stopped {
				return hits, true
			}
		} else {
			stopped = do(step.U, step.V, false)
			df.visit.setHits(step.V, 1)
			hits++
			if stopped {
				return hits, true
			}
			sortStart := len(df.mem)
			eachNext(df.g, step.V, selectNext)
			if df.SortBy != nil {
				sort.Slice(df.mem[sortStart:], func(v1, v2 int) bool {
					return !df.SortBy(step.V, v1, v2)
				})
			}
		}
	}
	return hits, false
}

func (df *Search) AdjBreadth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.breadth1st(stopToNextCluster, do, groph.EachAdjacent)
}

func (df *Search) OutBreadth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.breadth1st(stopToNextCluster, do, groph.EachOutgoing)
}

func (df *Search) InBreadth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.breadth1st(stopToNextCluster, do, groph.EachIncoming)
}

func (df *Search) breadth1st(
	stopToNextCluster bool,
	do VisitInCluster,
	eachNext stepFn,
) (stopped bool) {
	cluster := 0
	cdo := func(p, n groph.VIdx, c bool) bool { return do(p, n, c, cluster) }
	hits, stop := df.breadth1stAt(0, cdo, eachNext)
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
		n, stop := df.breadth1stAt(start, cdo, eachNext)
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

// Hits returns how often the vertex v of graph g of this seach was hit by
// traversal operations.
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

// LeatsHits returns one of the vertices in graph g of the Search that was least
// frequently hit by traversal operations.
func (df *Search) LeastHits() (v groph.VIdx, hits int) {
	if len(df.visit.is) == 0 {
		return -1, -1
	}
	item := df.visit.is[0]
	return item.v, item.hits
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
