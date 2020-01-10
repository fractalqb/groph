package adjmatrix

import (
	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

// AdjMxDbitmap implements WGbool as a bitmap based adjacency
// matrix. Compared to AdjMxbool, this sacrifices runtime performance
// for lesser memory usage.
type DBitmap struct {
	adjMx
	bs iutil.BitSet
}

func NewDBitmap(order int, reuse *DBitmap) *DBitmap {
	sz := order * order
	sz = iutil.BitSetWords(sz)
	if reuse == nil {
		reuse = &DBitmap{
			adjMx: adjMx{ord: order},
			bs:    make(iutil.BitSet, sz),
		}
	} else {
		reuse.ord = order
		reuse.bs = iutil.U64Slice(reuse.bs, int(sz))
	}
	return reuse
}

func (m *DBitmap) Init(flag bool) *DBitmap {
	if flag {
		for i := range m.bs {
			m.bs[i] = ^uint64(0)
		}
	} else {
		for i := range m.bs {
			m.bs[i] = 0
		}
	}
	return m
}

func (m *DBitmap) Reset(order int) {
	NewDBitmap(order, m)
	m.Init(false)
}

func (m *DBitmap) Weight(u, v groph.VIdx) interface{} {
	if m.Edge(u, v) {
		return true
	}
	return nil
}

func (m *DBitmap) SetWeight(i, j groph.VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(i, j, false)
	} else {
		m.SetEdge(i, j, w.(bool))
	}
}

func (m *DBitmap) Edge(i, j groph.VIdx) (w bool) {
	w = m.bs.Get(m.ord*i + j)
	return w
}

func (m *DBitmap) SetEdge(i, j groph.VIdx, w bool) {
	if w {
		m.bs.Set(m.ord*i + j)
	} else {
		m.bs.Unset(m.ord*i + j)
	}
}
