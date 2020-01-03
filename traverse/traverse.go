package traverse

import (
	"container/heap"
	"sort"

	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type VisitVertex = func(v groph.VIdx) (stop bool)

type VisitInCluster = func(n groph.VIdx, cluster int) (stop bool)

// Search performs depth-first or breadth-breadth searches of
// groph.RGraph objects.
type Search struct {
	g     groph.RGraph
	mem   []groph.VIdx
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

// Depth1stAt performs a depth-first search on the respective
// RGraph. The search starts at vertex start and terminates when no
// further vertices can be reached via an edge of the graph. It
// returns the number of distinct vertex hits during the search and
// an indicator if the visit function 'do' stopped the search.
func (df *Search) Depth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	if h := df.visit.hits(start); h > 0 {
		return 0, false
	}
	if df.mem != nil {
		df.mem = df.mem[:0]
	}
	df.push(start)
	df.visit.setHits(start, 1)
	last := groph.VIdx(-1)
	var onDest func(n groph.VIdx)
	if groph.Directed(df.g) {
		onDest = df.chechDest
	} else {
		onDest = func(n groph.VIdx) {
			if n == last || n == start {
				return
			}
			df.chechDest(n)
		}
	}
	for len(df.mem) > 0 {
		start = df.pop()
		if do(start) {
			return hits + 1, true
		}
		hits++
		sortStart := len(df.mem)
		groph.EachAdjacent(df.g, start, onDest)
		if df.SortBy != nil {
			sort.Slice(df.mem[sortStart:], func(v1, v2 int) bool {
				return !df.SortBy(start, v1, v2)
			})
		}
		last = start
	}
	return hits, false
}

func (df *Search) Depth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	cluster := 0
	cdo := func(n groph.VIdx) bool { return do(n, cluster) }
	hits, stop := df.Depth1stAt(0, cdo)
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
		n, stop := df.Depth1stAt(start, cdo)
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
func (df *Search) Breadth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	if h := df.visit.hits(start); h > 0 {
		return 0, false
	}
	if df.mem != nil {
		df.mem = df.mem[:0]
	}
	df.tail = 0
	df.push(start)
	df.visit.setHits(start, 1)
	last := groph.VIdx(-1)
	var onDest func(n groph.VIdx)
	if groph.Directed(df.g) {
		onDest = df.chechDest
	} else {
		onDest = func(n groph.VIdx) {
			if n == last || n == start {
				return
			}
			df.chechDest(n)
		}
	}
	for df.tail < len(df.mem) {
		start = df.take()
		if do(start) {
			return hits + 1, true
		}
		hits++
		sortStart := len(df.mem)
		groph.EachAdjacent(df.g, start, onDest)
		if df.SortBy != nil {
			sort.Slice(df.mem[sortStart:], func(v1, v2 int) bool {
				return df.SortBy(start, v1, v2)
			})
		}
		last = start
	}
	return hits, false
}

func (df *Search) Breadth1st(stopToNextCluster bool, do VisitInCluster) (stopped bool) {
	cluster := 0
	cdo := func(n groph.VIdx) bool { return do(n, cluster) }
	hits, stop := df.Breadth1stAt(0, cdo)
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
		n, stop := df.Breadth1stAt(start, cdo)
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

func (df *Search) push(v groph.VIdx) {
	df.mem = append(df.mem, v)
}

func (df *Search) pop() (res groph.VIdx) {
	l := len(df.mem) - 1
	res = df.mem[l]
	df.mem = df.mem[:l]
	return res
}

func (df *Search) take() (res groph.VIdx) {
	res = df.mem[df.tail]
	df.tail++
	return res
}

func (df *Search) chechDest(n groph.VIdx) {
	h := df.visit.hits(n)
	if h == 0 {
		df.push(n)
	}
	df.visit.setHits(n, h+1)
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
