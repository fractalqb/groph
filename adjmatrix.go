package graph

import (
	"math"
)

type adjMx struct {
	sz uint
	vp func(idx uint) Vertex
}

func vpId(i uint) Vertex { return i }

func (m *adjMx) VertexNo() uint { return m.sz }

func (m *adjMx) Vertex(idx uint) Vertex { return m.vp(idx) }

type AdjMxAbool struct {
	adjMx
	// TODO
}

type AdjMxAint struct {
	adjMx
	nx int
	w  []int
}

var _ RGint = (*AdjMxAint)(nil)

func NewAdjMxAint(size uint, notExist int,
	vertexProvider func(idx uint) Vertex,
	reuse *AdjMxAint) *AdjMxAint {
	if vertexProvider == nil {
		vertexProvider = vpId
	}
	if reuse == nil {
		reuse = &AdjMxAint{
			adjMx: adjMx{sz: size, vp: vertexProvider},
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
	w := m.Edge(fromIdx, toIdx)
	if w == m.nx {
		return nil
	} else {
		return w
	}
}

func (m *AdjMxAint) Directed() bool { return true }

func (m *AdjMxAint) Edge(fromIdx, toIdx uint) (weight int) {
	return m.w[m.sz*fromIdx+toIdx]
}

func (m *AdjMxAint) SetWeight(i, j uint, w interface{}) {
	m.w[m.sz*i+j] = w.(int)
}

type AdjMxAf32 struct {
	adjMx
	w []float32
}

var _ RGf32 = (*AdjMxAf32)(nil)

func NewAdjMxAf32(size uint,
	vertexProvider func(idx uint) Vertex,
	reuse *AdjMxAf32) *AdjMxAf32 {
	if vertexProvider == nil {
		vertexProvider = vpId
	}
	if reuse == nil {
		reuse = &AdjMxAf32{
			adjMx: adjMx{sz: size, vp: vertexProvider},
			w:     make([]float32, size*size),
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
	w := m.Edge(fromIdx, toIdx)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *AdjMxAf32) Directed() bool { return true }

func (m *AdjMxAf32) Edge(i, j uint) (w float32) {
	return m.w[m.sz*i+j]
}

func (m *AdjMxAf32) SetWeight(i, j uint, w interface{}) {
	m.w[m.sz*i+j] = w.(float32)
}

func nSum(n uint) uint { return n * (n + 1) / 2 }

type AdjMxSf32 struct {
	adjMx
	w []float32
}

var _ RGf32 = (*AdjMxSf32)(nil)

func NewAdjMxSf32(size uint,
	vertexProvider func(idx uint) Vertex,
	reuse *AdjMxSf32) *AdjMxSf32 {
	if vertexProvider == nil {
		vertexProvider = vpId
	}
	if reuse == nil {
		reuse = &AdjMxSf32{
			adjMx: adjMx{sz: size, vp: vertexProvider},
			w:     make([]float32, nSum(size)),
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

func (m *AdjMxSf32) Weight(fromIdx, toIdx uint) interface{} {
	w := m.Edge(fromIdx, toIdx)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *AdjMxSf32) Directed() bool { return false }

func (m *AdjMxSf32) Edge(i, j uint) (w float32) {
	if i <= j {
		return m.w[m.sz*i+j]
	} else {
		return m.w[m.sz*j+i]
	}
}

func (m *AdjMxSf32) SetWeight(i, j uint, w interface{}) {
	if i < j {
		m.w[m.sz*i+j] = w.(float32)
	} else {
		m.w[m.sz*j+i] = w.(float32)
	}
}
