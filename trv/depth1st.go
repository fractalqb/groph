package trv

import (
	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
	"git.fractalqb.de/fractalqb/groph/util"
)

type DepthFirst struct {
	g       groph.RGraph
	s       []groph.VIdx
	Visited util.BitSet
}

func NewDepthFirst(g groph.RGraph) *DepthFirst {
	res := &DepthFirst{
		g:       g,
		Visited: util.NewBitSet(g.VertexNo()),
	}
	return res
}

func (df *DepthFirst) Reset(g groph.RGraph) {
	df.g = g
	df.Visited = iutil.U64Slice(df.Visited, util.BitSetWords(g.VertexNo()))
}

func (df *DepthFirst) push(v groph.VIdx) {
	df.s = append(df.s, v)
}

func (df *DepthFirst) pop() (res groph.VIdx) {
	l := len(df.s) - 1
	res = df.s[l]
	df.s = df.s[:l]
	return res
}

func (df *DepthFirst) Cluster(start groph.VIdx, do groph.VisitVertex) int {
	if df.Visited.Get(start) {
		return 0
	}
	if df.s != nil {
		df.s = df.s[:0]
	}
	df.push(start)
	count := 0
	for len(df.s) > 0 {
		start = df.pop()
		df.Visited.Set(start)
		do(start)
		count++
		groph.EachOutgoing(df.g, start, func(n groph.VIdx) {
			if !df.Visited.Get(n) {
				df.push(n)
			}
		})
	}
	return count
}

func (df *DepthFirst) Complete(do func(n groph.VIdx, cluster int)) {
	cluster := 0
	cdo := func(n groph.VIdx) { do(n, cluster) }
	count := df.Cluster(0, cdo)
	for count < df.g.VertexNo() {
		cluster++
		start := df.Visited.FirstUnset()
		count += df.Cluster(start, cdo)
	}
}
