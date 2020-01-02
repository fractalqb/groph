package traverse

import (
	"sort"

	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type VisitVertex = func(v groph.VIdx) (stop bool)

type VisitInCluster = func(n groph.VIdx, cluster int) (stop bool)

type Search struct {
	g    groph.RGraph
	mem  []groph.VIdx
	tail int
	// If not nil SortBy is used to sort the neighbours v of node u. SortBy
	// returns true if the edge (u,v1) is less than (u,v2).
	SortBy  func(u, v1, v2 groph.VIdx) bool
	Visited groph.BitSet
}

func NewSearch(g groph.RGraph) *Search {
	res := &Search{
		g:       g,
		Visited: groph.NewBitSet(g.Order()),
	}
	return res
}

func (df *Search) Reset(g groph.RGraph) {
	df.g = g
	df.Visited = iutil.U64Slice(df.Visited, groph.BitSetWords(g.Order()))
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

func (df *Search) Depth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	if df.Visited.Get(start) {
		return 0, false
	}
	if df.mem != nil {
		df.mem = df.mem[:0]
	}
	df.push(start)
	df.Visited.Set(start)
	for len(df.mem) > 0 {
		start = df.pop()
		if do(start) {
			return hits + 1, true
		}
		hits++
		sortStart := len(df.mem)
		groph.EachOutgoing(df.g, start, func(n groph.VIdx) {
			if !df.Visited.Get(n) {
				df.push(n)
				df.Visited.Set(n)
			}
		})
		if df.SortBy != nil {
			sort.Slice(df.mem[sortStart:], func(v1, v2 int) bool {
				return !df.SortBy(start, v1, v2)
			})
		}
	}
	return hits, false
}

func (df *Search) Depth1st(do VisitInCluster, stopToNextCluster bool) (stopped bool) {
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
		start := df.Visited.FirstUnset()
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

func (df *Search) Breadth1stAt(start groph.VIdx, do VisitVertex) (hits int, stopped bool) {
	if df.Visited.Get(start) {
		return 0, false
	}
	if df.mem != nil {
		df.mem = df.mem[:0]
	}
	df.tail = 0
	df.push(start)
	df.Visited.Set(start)
	for df.tail < len(df.mem) {
		start = df.take()
		if do(start) {
			return hits + 1, true
		}
		hits++
		sortStart := len(df.mem)
		groph.EachOutgoing(df.g, start, func(n groph.VIdx) {
			if !df.Visited.Get(n) {
				df.push(n)
				df.Visited.Set(n)
			}
		})
		if df.SortBy != nil {
			sort.Slice(df.mem[sortStart:], func(v1, v2 int) bool {
				return df.SortBy(start, v1, v2)
			})
		}
	}
	return hits, false
}

func (df *Search) Breadth1st(do VisitInCluster, stopToNextCluster bool) (stopped bool) {
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
		start := df.Visited.FirstUnset()
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
