package groph

import (
	"math"
)

type adjMx struct {
	sz VIdx
}

func (m *adjMx) VertexNo() VIdx { return m.sz }

// AdjMxDbitmap is a directed WGraph with boolean edge weights implemented as
// bitmap. This sacrifices runtime performance for lesser memory usage.
type AdjMxDbitmap struct {
	adjMx
	bs []uint
}

func NewAdjMxDbitmap(vertexNo VIdx, reuse *AdjMxDbitmap) *AdjMxDbitmap {
	sz := uint(vertexNo * vertexNo)
	sz = (sz + (wordBits - 1)) / wordBits
	if reuse == nil {
		reuse = &AdjMxDbitmap{
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

func (m *AdjMxDbitmap) Init(flag bool) *AdjMxDbitmap {
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

func (m *AdjMxDbitmap) Directed() bool { return true }

func (m *AdjMxDbitmap) Clear(vertexNo VIdx) {
	NewAdjMxDbitmap(vertexNo, m)
	m.Init(false)
}

func (m *AdjMxDbitmap) Weight(edgeFrom, edgeTo VIdx) interface{} {
	return m.Edge(edgeFrom, edgeTo)
}

func (m *AdjMxDbitmap) SetWeight(i, j VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(i, j, false)
	} else {
		m.SetEdge(i, j, w.(bool))
	}
}

func (m *AdjMxDbitmap) Edge(i, j VIdx) (w bool) {
	w = BitSetGet(m.bs, uint(m.sz*i+j))
	return w
}

func (m *AdjMxDbitmap) SetEdge(i, j VIdx, w bool) {
	if w {
		BitSetSet(m.bs, uint(m.sz*i+j))
	} else {
		BitSetUnset(m.bs, uint(m.sz*i+j))
	}
}

type AdjMxDbool struct {
	adjMx
	bs []bool
}

func NewAdjMxDbool(vertexNo VIdx, reuse *AdjMxDbool) *AdjMxDbool {
	sz := vertexNo * vertexNo
	if reuse == nil {
		reuse = &AdjMxDbool{
			adjMx: adjMx{sz: vertexNo},
			bs:    make([]bool, sz),
		}
	} else if VIdx(cap(reuse.bs)) >= sz {
		reuse.sz = vertexNo
		reuse.bs = reuse.bs[:sz]
	} else {
		reuse.sz = vertexNo
		reuse.bs = make([]bool, sz)
	}
	return reuse
}

func (m *AdjMxDbool) Init(flag bool) *AdjMxDbool {
	for i := range m.bs {
		m.bs[i] = flag
	}
	return m
}

func (m *AdjMxDbool) Directed() bool { return true }

func (m *AdjMxDbool) Clear(vertexNo VIdx) {
	NewAdjMxDbool(vertexNo, m)
	m.Init(false)
}

func (m *AdjMxDbool) Weight(edgeFrom, edgeTo VIdx) interface{} {
	return m.Edge(edgeFrom, edgeTo)
}

func (m *AdjMxDbool) SetWeight(i, j VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(i, j, false)
	} else {
		m.SetEdge(i, j, w.(bool))
	}
}

func (m *AdjMxDbool) Edge(i, j VIdx) (w bool) {
	return m.bs[m.sz*i+j]
}

func (m *AdjMxDbool) SetEdge(i, j VIdx, w bool) {
	m.bs[m.sz*i+j] = w
}

type AdjMxDi32 struct {
	adjMx
	w       []int32
	Cleared int32
}

const I32cleared = -2147483648

func NewAdjMxDi32(vertexNo VIdx, reuse *AdjMxDi32) *AdjMxDi32 {
	if reuse == nil {
		reuse = &AdjMxDi32{
			adjMx:   adjMx{sz: vertexNo},
			w:       make([]int32, vertexNo*vertexNo),
			Cleared: I32cleared,
		}
	} else if VIdx(cap(reuse.w)) >= vertexNo*vertexNo {
		reuse.sz = vertexNo
		reuse.w = reuse.w[:vertexNo*vertexNo]
	} else {
		reuse.sz = vertexNo
		reuse.w = make([]int32, vertexNo*vertexNo)
	}
	return reuse
}

func (m *AdjMxDi32) Init(w int32) *AdjMxDi32 {
	for i := range m.w {
		m.w[i] = w
	}
	return m
}

func (m *AdjMxDi32) Directed() bool { return true }

func (m *AdjMxDi32) Clear(vertexNo VIdx) {
	c := m.Cleared
	NewAdjMxDi32(vertexNo, m)
	m.Cleared = c
	m.Init(m.Cleared)
}

func (m *AdjMxDi32) Weight(edgeFrom, edgeTo VIdx) interface{} {
	res, ok := m.Edge(edgeFrom, edgeTo)
	if ok {
		return res
	}
	return nil
}

func (m *AdjMxDi32) SetWeight(i, j VIdx, w interface{}) {
	if w == nil {
		m.DelEdge(i, j)
	} else {
		m.w[m.sz*i+j] = w.(int32)
	}
}

func (m *AdjMxDi32) Edge(i, j VIdx) (w int32, exists bool) {
	w = m.w[m.sz*i+j]
	return w, w != m.Cleared
}

func (m *AdjMxDi32) SetEdge(i, j VIdx, w int32) {
	m.w[m.sz*i+j] = w
}

func (m *AdjMxDi32) DelEdge(i, j VIdx) {
	m.SetEdge(i, j, m.Cleared)
}

type AdjMxDf32 struct {
	adjMx
	w []float32
}

func NewAdjMxDf32(vertexNo VIdx, reuse *AdjMxDf32) *AdjMxDf32 {
	if reuse == nil {
		reuse = &AdjMxDf32{
			adjMx: adjMx{sz: vertexNo},
			w:     make([]float32, vertexNo*vertexNo),
		}
	} else if VIdx(cap(reuse.w)) >= vertexNo*vertexNo {
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

func (m *AdjMxDf32) Clear(vertexNo VIdx) {
	NewAdjMxDf32(vertexNo, m)
	m.Init(nan32)
}

func (m *AdjMxDf32) Weight(edgeFrom, edgeTo VIdx) interface{} {
	w := m.Edge(edgeFrom, edgeTo)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *AdjMxDf32) SetWeight(i, j VIdx, w interface{}) {
	if w == nil {
		m.w[m.sz*i+j] = nan32
	} else {
		m.w[m.sz*i+j] = w.(float32)
	}
}

func (m *AdjMxDf32) Edge(i, j VIdx) (w float32) {
	return m.w[m.sz*i+j]
}

func (m *AdjMxDf32) SetEdge(i, j VIdx, w float32) {
	m.w[m.sz*i+j] = w
}

// uSum computes the sum of the n 1st integers, i.e. 1+2+3+…+n
func nSum(n VIdx) VIdx { return n * (n + 1) / 2 }

func nSumRev(n VIdx) float64 {
	r := math.Sqrt(0.25 + 2*float64(n))
	return r - 0.5
}

type AdjMxUf32 struct {
	adjMx
	w []float32
}

func NewAdjMxUf32(vertexNo VIdx, reuse *AdjMxUf32) *AdjMxUf32 {
	if reuse == nil {
		reuse = &AdjMxUf32{
			adjMx: adjMx{sz: vertexNo},
			w:     make([]float32, nSum(vertexNo)),
		}
	} else if VIdx(cap(reuse.w)) >= nSum(vertexNo) {
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

func (m *AdjMxUf32) Weight(edgeFrom, edgeTo VIdx) interface{} {
	w := m.Edge(edgeFrom, edgeTo)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *AdjMxUf32) Directed() bool { return false }

func (m *AdjMxUf32) Clear(vertexNo VIdx) {
	NewAdjMxUf32(vertexNo, m)
	m.Init(nan32)
}

// uIdx computes the index into the weight slice of an undirected matrix
func uIdx(sz, i, j VIdx) VIdx {
	j -= i
	i = nSum(sz - i - 1)
	return i + j
}

func (m *AdjMxUf32) SetWeight(i, j VIdx, w interface{}) {
	if i < j {
		m.w[uIdx(m.sz, i, j)] = w.(float32)
	} else {
		m.w[uIdx(m.sz, j, i)] = w.(float32)
	}
}

func (m *AdjMxUf32) Edge(i, j VIdx) (w float32) {
	if i <= j {
		return m.w[uIdx(m.sz, i, j)]
	} else {
		return m.w[uIdx(m.sz, j, i)]
	}
}

func (m *AdjMxUf32) SetEdge(i, j VIdx, w float32) {
	if i <= j {
		m.w[uIdx(m.sz, i, j)] = w
	} else {
		m.w[uIdx(m.sz, j, i)] = w
	}
}
