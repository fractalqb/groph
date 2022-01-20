package traverse

import (
	"container/heap"
	"sort"

	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type VisitVertex = func(pred, v groph.VIdx, vHits int) (stop bool)

type VisitInCluster = func(pred, v groph.VIdx, vHits int, cluster int) (stop bool)

// Search performs depth-first or breadth-first searches of groph.RGraph
// objects. A Search is created for the graph to be searched. The Search must
// not be used concurrently. It can be resued for other graphs with the Reset
// methods.
type Search struct {
	g    groph.RGraph
	mem  []groph.Edge
	tail int
	// Heap with per-node hit counts. A root with hit count 0 is a good start
	// for the next tree in a forest.
	visit hitPrioQ
	// If not nil SortBy is used to sort the neighbours v of node u. SortBy
	// returns true if the edge (u,v1) is less than (u,v2).
	SortBy func(u, v1, v2 groph.VIdx) bool
}

func VIdxOrder(_, v1, v2 groph.VIdx) bool { return v1 < v2 }

func NewSearch(g groph.RGraph) (res *Search) {
	if g == nil {
		res = new(Search)
	} else {
		res = &Search{g: g}
		res.visit.reset(g.Order())
	}
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
	return df.depth1stAt(start, do, groph.EachAdjacent, false)
}

func (df *Search) OutDepth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.depth1stAt(start, do, groph.EachOutgoing, groph.Directed(df.g))
}

func (df *Search) InDepth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	return df.depth1stAt(start, do, groph.EachIncoming, groph.Directed(df.g))
}

func (df *Search) depth1stAt(
	start groph.VIdx,
	do VisitVertex,
	eachNext stepFn,
	selectDir bool,
) (hits int, stopped bool) {
	if df.g.Order() == 0 {
		return 0, false
	}
	if df.mem != nil {
		df.mem = df.mem[:0]
	}
	step := groph.Edge{U: -1, V: start}
	df.push(step)
	var selectNext groph.VisitVertex
	if selectDir {
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
		if step.V < 0 { // finish marker for backtracking
			df.visit.setFini(step.U, true)
			continue
		}
		h := df.visit.hits(step.V)
		if h > 0 {
			stopped = do(step.U, step.V, h)
			df.visit.setHits(step.V, h+1)
			if stopped {
				return hits, true
			}
		} else {
			stopped = do(step.U, step.V, 0)
			df.visit.setHits(step.V, 1)
			hits++
			if stopped {
				return hits, true
			}
			df.push(groph.Edge{U: step.V, V: -1})
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
	return df.depth1st(stopToNextCluster, do, groph.EachAdjacent, false)
}

func (df *Search) OutDepth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.depth1st(stopToNextCluster, do, groph.EachOutgoing, groph.Directed(df.g))
}

func (df *Search) InDepth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	return df.depth1st(stopToNextCluster, do, groph.EachIncoming, groph.Directed(df.g))
}

func (df *Search) depth1st(
	stopToNextCluster bool,
	do VisitInCluster,
	eachNext stepFn,
	selectDir bool,
) (stopped bool) {
	cluster := 0
	cdo := func(p, n groph.VIdx, nh int) bool { return do(p, n, nh, cluster) }
	hits, stop := df.depth1stAt(0, cdo, eachNext, selectDir)
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
		n, stop := df.depth1stAt(start, cdo, eachNext, selectDir)
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
	if df.g.Order() == 0 {
		return 0, false
	}
	if df.mem != nil {
		df.mem = df.mem[:0]
	}
	lastU := -1
	step := groph.Edge{U: lastU, V: start}
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
			stopped = do(step.U, step.V, h)
			df.visit.setHits(step.V, h+1)
			if stopped {
				return hits, true
			}
		} else {
			stopped = do(step.U, step.V, 0)
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
	cdo := func(p, n groph.VIdx, nh int) bool { return do(p, n, nh, cluster) }
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

// State returns the state of a vertex caused by a search. The state is the
// number of hits and its finished state. The finished state of vertext v is set
// true only during depth-first searchs when all vertices reached through v have
// been visited.
//
// State may be called during searches in VisitVertex or VisitInCluster
// callback functions. Note that the hit count of a vertex is updated after
// the callback.
func (df *Search) State(v groph.VIdx) (hits int, finished bool) {
	if v >= len(df.visit.vtx2item) {
		return 0, false
	}
	pqi := df.visit.items[df.visit.vtx2item[v]]
	return pqi.hits(), pqi.fini()
}

// LeatsHits returns one of the vertices in graph g of the Search that was least
// frequently hit by traversal operations.
func (df *Search) LeastHits() (v groph.VIdx, hits int) {
	if len(df.visit.items) == 0 {
		return -1, -1
	}
	item := df.visit.items[0]
	return item.v, item.hits()
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

func (df *Search) top() (step groph.Edge) {
	l := len(df.mem)
	if l == 0 {
		return groph.Edge{U: -1, V: -1}
	}
	return df.mem[l-1]
}

func (df *Search) take() (step groph.Edge) {
	step = df.mem[df.tail]
	df.tail++
	return step
}

type hitPqItem struct {
	finHits int
	v       groph.VIdx
}

const (
	MaxUInt = ^uint(0)
	MaxInt  = int(MaxUInt >> 1)
	MinInt  = -MaxInt - 1
)

func (pqi *hitPqItem) fini() bool { return pqi.finHits&MinInt != 0 }

func (pqi *hitPqItem) setFini(f bool) {
	hits := pqi.hits()
	if f {
		pqi.finHits = hits | MinInt
	} else {
		pqi.finHits = hits
	}
}

func (pqi *hitPqItem) hits() int { return pqi.finHits & MaxInt }

func (pqi *hitPqItem) setHits(h int) {
	fini := pqi.finHits & MinInt
	pqi.finHits = fini | h
}

type hitPrioQ struct {
	vtx2item []int
	items    []hitPqItem
}

func (pq *hitPrioQ) reset(ord groph.VIdx) {
	pq.vtx2item = iutil.IntSlice(pq.vtx2item, ord)
	if pq.items == nil || cap(pq.items) < ord {
		pq.items = make([]hitPqItem, ord)
	} else {
		pq.items = pq.items[:ord]
	}
	for i := range pq.items {
		pq.vtx2item[i] = i
		pq.items[i] = hitPqItem{finHits: 0, v: i}
	}
}

func (pq *hitPrioQ) top() groph.VIdx { return pq.items[0].v }

func (pq *hitPrioQ) hits(v groph.VIdx) int {
	return pq.items[pq.vtx2item[v]].hits()
}

func (pq *hitPrioQ) setHits(v groph.VIdx, hits int) {
	ii := pq.vtx2item[v]
	pq.items[ii].setHits(hits)
	heap.Fix(pq, ii)
}

func (pq *hitPrioQ) setFini(v groph.VIdx, f bool) {
	ii := pq.vtx2item[v]
	pq.items[ii].setFini(f)
}

func (pq *hitPrioQ) Len() int { return len(pq.items) }

func (pq *hitPrioQ) Less(i, j int) bool {
	return pq.items[i].hits() < pq.items[j].hits()
}

func (pq *hitPrioQ) Swap(i, j int) {
	ii, ji := pq.items[i], pq.items[j]
	pq.items[i], pq.items[j] = ji, ii
	pq.vtx2item[ii.v] = j
	pq.vtx2item[ji.v] = i
}

func (pq *hitPrioQ) Push(x interface{}) {
	panic("must not be called")
	// pqItem := x.(hitPqItem)
	// for pqItem.v >= len(pq.v2i) {
	// 	pq.v2i = append(pq.v2i, -1)
	// }
	// pq.v2i[pqItem.v] = len(pq.is)
	// pq.is = append(pq.is, pqItem)
}

func (pq *hitPrioQ) Pop() interface{} {
	panic("must not be called")
	// lm1 := len(pq.is) - 1
	// res := pq.is[lm1]
	// pq.is = pq.is[:lm1]
	// pq.v2i[res.v] = -1
	// return res
}
