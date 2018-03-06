package graph

import (
	"math"
)

type adjMx struct {
	sz             uint
	VertexProvider func(idx uint) Vertex
}

func (m *adjMx) VertexNo() uint { return m.sz }

func (m *adjMx) Vertex(idx uint) Vertex { return m.VertexProvider(idx) }

type AdjMxAbool struct {
	// TODO
}

type AdjMxAint struct {
	adjMx
	nx int
	w  []int
}

var _ Gint = (*AdjMxAint)(nil)

func NewAdjMxAint(size uint, notExist int, reuse *AdjMxAint) *AdjMxAint {
	if reuse == nil {
		reuse = &AdjMxAint{
			adjMx: adjMx{sz: size},
			w:     make([]int, size*size),
		}
	} else if uint(cap(reuse.w)) >= size*size {
		reuse.sz = size
		reuse.w = reuse.w[:size*size]
	} else {
		reuse.sz = size
		reuse.w = make([]int, reuse.sz)
	}
	reuse.nx = notExist
	return reuse
}

func (m *AdjMxAint) Weight(fromIdx, toIdx uint) interface{} {
	x, w := m.Edge(fromIdx, toIdx)
	if x {
		return w
	} else {
		return nil
	}
}

func (m *AdjMxAint) Directed() bool { return true }

func (m *AdjMxAint) Edge(fromIdx, toIdx uint) (exists bool, weight int) {
	weight = m.w[m.sz*fromIdx+toIdx]
	return weight != m.nx, weight
}

func (m *AdjMxAint) SetWeight(i, j uint, w interface{}) {
	m.w[m.sz*i+j] = w.(int)
}

type AdjMxAf32 struct {
	adjMx
	sz uint
	w  []float32
}

var _ Gf32 = (*AdjMxAf32)(nil)

func NewAdjMxAf32(size uint, reuse *AdjMxAf32) *AdjMxAf32 {
	if reuse == nil {
		reuse = &AdjMxAf32{
			sz: size,
			w:  make([]float32, size*size),
		}
	} else if uint(cap(reuse.w)) >= size*size {
		reuse.sz = size
		reuse.w = reuse.w[:size*size]
	} else {
		reuse.sz = size
		reuse.w = make([]float32, reuse.sz)
	}
	return reuse
}

func (m *AdjMxAf32) Weight(fromIdx, toIdx uint) interface{} {
	x, w := m.Edge(fromIdx, toIdx)
	if x {
		return w
	} else {
		return nil
	}
}

func (m *AdjMxAf32) Directed() bool { return true }

func (m *AdjMxAf32) Edge(i, j uint) (ex bool, w float32) {
	w = m.w[m.sz*i+j]
	return !math.IsNaN(float64(w)), w
}

func (m *AdjMxAf32) SetWeight(i, j uint, w interface{}) {
	m.w[m.sz*i+j] = w.(float32)
}

func nSum(n uint) uint { return n * (n + 1) / 2 }

type AdjMxSf32 struct {
	sz uint
	w  []float32
}

func NewAdjMxSf32(size uint, reuse *AdjMxSf32) *AdjMxSf32 {
	if reuse == nil {
		reuse = &AdjMxSf32{
			sz: size,
			w:  make([]float32, nSum(size)),
		}
	} else if uint(cap(reuse.w)) >= nSum(size) {
		reuse.sz = size
		reuse.w = reuse.w[:nSum(size)]
	} else {
		reuse.sz = size
		reuse.w = make([]float32, nSum(size))
	}
	return reuse
}
