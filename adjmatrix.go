package groph

import (
	"math"
)

type adjMx struct {
	sz uint
}

func (m *adjMx) VertexNo() uint { return m.sz }

type AdjMxDf32 struct {
	adjMx
	w []float32
}

var _ WGf32 = (*AdjMxDf32)(nil)

func NewAdjMxDf32(size uint, reuse *AdjMxDf32) *AdjMxDf32 {
	if reuse == nil {
		reuse = &AdjMxDf32{
			adjMx: adjMx{sz: size},
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

func (m *AdjMxDf32) Weight(fromIdx, toIdx uint) interface{} {
	w := m.Edge(fromIdx, toIdx)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *AdjMxDf32) SetWeight(i, j uint, w interface{}) {
	m.w[m.sz*i+j] = w.(float32)
}

func (m *AdjMxDf32) Directed() bool { return true }

func (m *AdjMxDf32) Edge(i, j uint) (w float32) {
	return m.w[m.sz*i+j]
}

func (m *AdjMxDf32) SetEdge(i, j uint, w float32) {
	m.w[m.sz*i+j] = w
}

func nSum(n uint) uint { return n * (n + 1) / 2 }

type AdjMxUf32 struct {
	adjMx
	w []float32
}

var _ WGf32 = (*AdjMxUf32)(nil)

func NewAdjMxUf32(size uint, reuse *AdjMxUf32) *AdjMxUf32 {
	if reuse == nil {
		reuse = &AdjMxUf32{
			adjMx: adjMx{sz: size},
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

func (m *AdjMxUf32) Weight(fromIdx, toIdx uint) interface{} {
	w := m.Edge(fromIdx, toIdx)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *AdjMxUf32) Directed() bool { return false }

func (m *AdjMxUf32) SetWeight(i, j uint, w interface{}) {
	if i < j {
		m.w[m.sz*i+j] = w.(float32)
	} else {
		m.w[m.sz*j+i] = w.(float32)
	}
}

func (m *AdjMxUf32) Edge(i, j uint) (w float32) {
	if i <= j {
		return m.w[m.sz*i+j]
	} else {
		return m.w[m.sz*j+i]
	}
}

func (m *AdjMxUf32) SetEdge(i, j uint, w float32) {
	if i <= j {
		m.w[m.sz*i+j] = w
	} else {
		m.w[m.sz*j+i] = w
	}
}
