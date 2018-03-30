package groph

import (
	"math"
)

type adjMx struct {
	sz uint
}

func (m *adjMx) VertexNo() uint { return m.sz }

type AdjMxDbool struct {
	adjMx
	bs []uint
}

var _ WGbool = (*AdjMxDbool)(nil)

func NewAdjMxDbool(vertexNo uint, reuse *AdjMxDbool) *AdjMxDbool {
	sz := vertexNo * vertexNo
	sz = (sz + (wordBits - 1)) / wordBits
	if reuse == nil {
		reuse = &AdjMxDbool{
			adjMx: adjMx{sz: vertexNo},
			bs:    make([]uint, sz),
		}
	} else if uint(cap(reuse.bs)) >= sz {
		reuse.sz = vertexNo
		reuse.bs = reuse.bs[:sz]
	} else {
		reuse.sz = vertexNo
		reuse.bs = make([]uint, sz)
	}
	return reuse
}

func (m *AdjMxDbool) Init(flag bool) *AdjMxDbool {
	if flag {
		for i := range m.bs {
			m.bs[i] = ^uint(0)
		}
	} else {
		for i := range m.bs {
			m.bs[i] = 0
		}
	}
	return m
}

func (m *AdjMxDbool) Directed() bool { return true }

func (m *AdjMxDbool) Clear(vertexNo uint) {
	NewAdjMxDbool(vertexNo, m)
	m.Init(false)
}

func (m *AdjMxDbool) Weight(fromIdx, toIdx uint) interface{} {
	w := m.Edge(fromIdx, toIdx)
	if w {
		return w
	} else {
		return nil
	}
}

func (m *AdjMxDbool) SetWeight(i, j uint, w interface{}) {
	if w == nil {
		m.SetEdge(i, j, false)
	} else {
		m.SetEdge(i, j, w.(bool))
	}
}

func (m *AdjMxDbool) Edge(i, j uint) (w bool) {
	w = BitSetGet(m.bs, m.sz*i+j)
	return w
}

func (m *AdjMxDbool) SetEdge(i, j uint, w bool) {
	if w {
		BitSetSet(m.bs, m.sz*i+j)
	} else {
		BitSetUnset(m.bs, m.sz*i+j)
	}
}

type AdjMxDf32 struct {
	adjMx
	w []float32
}

var _ WGf32 = (*AdjMxDf32)(nil)

func NewAdjMxDf32(vertexNo uint, reuse *AdjMxDf32) *AdjMxDf32 {
	if reuse == nil {
		reuse = &AdjMxDf32{
			adjMx: adjMx{sz: vertexNo},
			w:     make([]float32, vertexNo*vertexNo),
		}
	} else if uint(cap(reuse.w)) >= vertexNo*vertexNo {
		reuse.sz = vertexNo
		reuse.w = reuse.w[:vertexNo*vertexNo]
	} else {
		reuse.sz = vertexNo
		reuse.w = make([]float32, vertexNo*vertexNo)
	}
	return reuse
}

func (m *AdjMxDf32) Init(w float32) *AdjMxDf32 {
	for i := range m.w {
		m.w[i] = w
	}
	return m
}

func (m *AdjMxDf32) Directed() bool { return true }

func (m *AdjMxDf32) Clear(vertexNo uint) {
	NewAdjMxDf32(vertexNo, m)
	m.Init(nan32)
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
	if w == nil {
		m.w[m.sz*i+j] = nan32
	} else {
		m.w[m.sz*i+j] = w.(float32)
	}
}

func (m *AdjMxDf32) Edge(i, j uint) (w float32) {
	return m.w[m.sz*i+j]
}

func (m *AdjMxDf32) SetEdge(i, j uint, w float32) {
	m.w[m.sz*i+j] = w
}

// uSum computes the sum of the n 1st integers, i.e. 1+2+3+…+n
func nSum(n uint) uint { return n * (n + 1) / 2 }

type AdjMxUf32 struct {
	adjMx
	w []float32
}

var _ WGf32 = (*AdjMxUf32)(nil)

func NewAdjMxUf32(vertexNo uint, reuse *AdjMxUf32) *AdjMxUf32 {
	if reuse == nil {
		reuse = &AdjMxUf32{
			adjMx: adjMx{sz: vertexNo},
			w:     make([]float32, nSum(vertexNo)),
		}
	} else if uint(cap(reuse.w)) >= nSum(vertexNo) {
		reuse.sz = vertexNo
		reuse.w = reuse.w[:nSum(vertexNo)]
	} else {
		reuse.sz = vertexNo
		reuse.w = make([]float32, nSum(vertexNo))
	}
	return reuse
}

func (m *AdjMxUf32) Init(w float32) *AdjMxUf32 {
	for i := range m.w {
		m.w[i] = w
	}
	return m
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

func (m *AdjMxUf32) Clear(vertexNo uint) {
	NewAdjMxUf32(vertexNo, m)
	m.Init(nan32)
}

// uIdx computes the index into the weight slice of an undirected matrix
func uIdx(sz, i, j uint) uint {
	j -= i
	i = nSum(sz - i - 1)
	return i + j
}

func (m *AdjMxUf32) SetWeight(i, j uint, w interface{}) {
	if i < j {
		m.w[uIdx(m.sz, i, j)] = w.(float32)
	} else {
		m.w[uIdx(m.sz, j, i)] = w.(float32)
	}
}

func (m *AdjMxUf32) Edge(i, j uint) (w float32) {
	if i <= j {
		return m.w[uIdx(m.sz, i, j)]
	} else {
		return m.w[uIdx(m.sz, j, i)]
	}
}

func (m *AdjMxUf32) SetEdge(i, j uint, w float32) {
	if i <= j {
		m.w[uIdx(m.sz, i, j)] = w
	} else {
		m.w[uIdx(m.sz, j, i)] = w
	}
}
