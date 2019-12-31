package groph

import (
	"math"

	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type adjMx struct {
	sz VIdx
}

func (m *adjMx) Order() VIdx { return m.sz }

// AdjMxDbitmap is a directed WGraph with boolean edge weights implemented as
// bitmap. This sacrifices runtime performance for lesser memory usage.
type AdjMxDbitmap struct {
	adjMx
	bs BitSet
}

func NewAdjMxDbitmap(order VIdx, reuse *AdjMxDbitmap) *AdjMxDbitmap {
	sz := order * order
	sz = BitSetWords(sz)
	if reuse == nil {
		reuse = &AdjMxDbitmap{
			adjMx: adjMx{sz: order},
			bs:    make(BitSet, sz),
		}
	} else {
		reuse.sz = order
		reuse.bs = iutil.U64Slice(reuse.bs, int(sz))
	}
	return reuse
}

func (m *AdjMxDbitmap) Init(flag bool) *AdjMxDbitmap {
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

func (m *AdjMxDbitmap) Reset(order VIdx) {
	NewAdjMxDbitmap(order, m)
	m.Init(false)
}

func (m *AdjMxDbitmap) Weight(u, v VIdx) interface{} {
	if m.Edge(u, v) {
		return true
	}
	return nil
}

func (m *AdjMxDbitmap) SetWeight(i, j VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(i, j, false)
	} else {
		m.SetEdge(i, j, w.(bool))
	}
}

func (m *AdjMxDbitmap) Edge(i, j VIdx) (w bool) {
	w = m.bs.Get(m.sz*i + j)
	return w
}

func (m *AdjMxDbitmap) SetEdge(i, j VIdx, w bool) {
	if w {
		m.bs.Set(m.sz*i + j)
	} else {
		m.bs.Unset(m.sz*i + j)
	}
}

type AdjMxDbool struct {
	adjMx
	bs []bool
}

func NewAdjMxDbool(order VIdx, reuse *AdjMxDbool) *AdjMxDbool {
	sz := order * order
	if reuse == nil {
		reuse = &AdjMxDbool{
			adjMx: adjMx{sz: order},
			bs:    make([]bool, sz),
		}
	} else {
		reuse.sz = order
		reuse.bs = iutil.BoolSlice(reuse.bs, int(sz))
	}
	return reuse
}

func (m *AdjMxDbool) Init(flag bool) *AdjMxDbool {
	for i := range m.bs {
		m.bs[i] = flag
	}
	return m
}

func (m *AdjMxDbool) Reset(order VIdx) {
	NewAdjMxDbool(order, m)
	m.Init(false)
}

func (m *AdjMxDbool) Weight(u, v VIdx) interface{} {
	if m.Edge(u, v) {
		return true
	}
	return nil
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
	w   []int32
	del int32
}

func NewAdjMxDi32(order VIdx, del int32, reuse *AdjMxDi32) *AdjMxDi32 {
	if reuse == nil {
		reuse = &AdjMxDi32{
			adjMx: adjMx{sz: order},
			w:     make([]int32, order*order),
			del:   del,
		}
	} else {
		reuse.sz = order
		reuse.w = iutil.I32Slice(reuse.w, int(order*order))
	}
	reuse.Init(reuse.del)
	return reuse
}

func (m *AdjMxDi32) Init(w int32) *AdjMxDi32 {
	for i := range m.w {
		m.w[i] = w
	}
	return m
}

func (m *AdjMxDi32) Reset(order VIdx) { NewAdjMxDi32(order, m.del, m) }

func (m *AdjMxDi32) Weight(u, v VIdx) interface{} {
	res, ok := m.Edge(u, v)
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
	return w, w != m.del
}

func (m *AdjMxDi32) SetEdge(i, j VIdx, w int32) {
	m.w[m.sz*i+j] = w
}

func (m *AdjMxDi32) DelEdge(i, j VIdx) {
	m.SetEdge(i, j, m.del)
}

type AdjMxDf32 struct {
	adjMx
	w []float32
}

func NewAdjMxDf32(order VIdx, reuse *AdjMxDf32) *AdjMxDf32 {
	if reuse == nil {
		reuse = &AdjMxDf32{
			adjMx: adjMx{sz: order},
			w:     make([]float32, order*order),
		}
	} else {
		reuse.sz = order
		reuse.w = iutil.F32Slice(reuse.w, int(order*order))
	}
	reuse.Init(NaN32())
	return reuse
}

func (m *AdjMxDf32) Init(w float32) *AdjMxDf32 {
	for i := range m.w {
		m.w[i] = w
	}
	return m
}

func (m *AdjMxDf32) Reset(order VIdx) { NewAdjMxDf32(order, m) }

func (m *AdjMxDf32) Weight(u, v VIdx) interface{} {
	w := m.Edge(u, v)
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

func NewAdjMxUf32(order VIdx, reuse *AdjMxUf32) *AdjMxUf32 {
	if reuse == nil {
		reuse = &AdjMxUf32{
			adjMx: adjMx{sz: order},
			w:     make([]float32, nSum(order)),
		}
	} else {
		reuse.sz = order
		reuse.w = iutil.F32Slice(reuse.w, int(nSum(order)))
	}
	reuse.Init(NaN32())
	return reuse
}

func (m *AdjMxUf32) Init(w float32) *AdjMxUf32 {
	for i := range m.w {
		m.w[i] = w
	}
	return m
}

func (m *AdjMxUf32) Reset(order VIdx) { NewAdjMxUf32(order, m) }

// uIdx computes the index into the weight slice of an undirected matrix
func uIdx(i, j VIdx) VIdx { return nSum(i) + j }

func (m *AdjMxUf32) Weight(i, j VIdx) interface{} {
	w := m.Edge(i, j)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *AdjMxUf32) WeightU(i, j VIdx) interface{} {
	w := m.EdgeU(i, j)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *AdjMxUf32) SetWeight(i, j VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(i, j, NaN32())
	} else {
		m.SetEdge(i, j, w.(float32))
	}
}

func (m *AdjMxUf32) SetWeightU(i, j VIdx, w interface{}) {
	if w == nil {
		m.SetEdgeU(i, j, NaN32())
	} else {
		m.SetEdgeU(i, j, w.(float32))
	}
}

func (m *AdjMxUf32) Edge(i, j VIdx) (w float32) {
	if i >= j {
		return m.w[uIdx(i, j)]
	} else {
		return m.w[uIdx(j, i)]
	}
}

// EdgeU is used iff i >= j
func (m *AdjMxUf32) EdgeU(i, j VIdx) (w float32) {
	return m.w[uIdx(i, j)]
}

func (m *AdjMxUf32) SetEdge(i, j VIdx, w float32) {
	if i >= j {
		m.w[uIdx(i, j)] = w
	} else {
		m.w[uIdx(j, i)] = w
	}
}

// SetEdgeU is used iff i >= j
func (m *AdjMxUf32) SetEdgeU(i, j VIdx, w float32) {
	m.w[uIdx(i, j)] = w
}

type AdjMxUi32 struct {
	adjMx
	w   []int32
	del int32
}

func NewAdjMxUi32(order VIdx, del int32, reuse *AdjMxUi32) *AdjMxUi32 {
	if reuse == nil {
		reuse = &AdjMxUi32{
			adjMx: adjMx{sz: order},
			w:     make([]int32, nSum(order)),
			del:   del,
		}
	} else {
		reuse.sz = order
		reuse.w = iutil.I32Slice(reuse.w, int(nSum(order)))
	}
	reuse.Init(reuse.del)
	return reuse
}

func (m *AdjMxUi32) Init(w int32) *AdjMxUi32 {
	for i := range m.w {
		m.w[i] = w
	}
	return m
}

func (m *AdjMxUi32) Reset(order VIdx) { NewAdjMxUi32(order, m.del, m) }

func (m *AdjMxUi32) Weight(u, v VIdx) interface{} {
	if w, ok := m.Edge(u, v); ok {
		return w
	}
	return nil
}

func (m *AdjMxUi32) WeightU(u, v VIdx) interface{} {
	w, ok := m.EdgeU(u, v)
	if ok {
		return w
	}
	return nil
}

func (m *AdjMxUi32) SetWeight(u, v VIdx, w interface{}) {
	if w == nil {
		m.DelEdge(u, v)
	} else {
		m.SetEdge(u, v, w.(int32))
	}
}

func (m *AdjMxUi32) SetWeightU(u, v VIdx, w interface{}) {
	if w == nil {
		m.DelEdgeU(u, v)
	} else {
		m.SetEdgeU(u, v, w.(int32))
	}
}

func (m *AdjMxUi32) DelEdge(u, v VIdx) {
	m.SetEdge(u, v, m.del)
}

func (m *AdjMxUi32) DelEdgeU(u, v VIdx) {
	m.SetEdgeU(u, v, m.del)
}

func (m *AdjMxUi32) Edge(u, v VIdx) (w int32, ok bool) {
	if u >= v {
		w = m.w[uIdx(u, v)]
	} else {
		w = m.w[uIdx(v, u)]
	}
	return w, w != m.del
}

// EdgeU is used iff i >= j
func (m *AdjMxUi32) EdgeU(u, v VIdx) (w int32, ok bool) {
	w = m.w[uIdx(u, v)]
	return w, w != m.del
}

func (m *AdjMxUi32) SetEdge(i, j VIdx, w int32) {
	if i >= j {
		m.w[uIdx(i, j)] = w
	} else {
		m.w[uIdx(j, i)] = w
	}
}

// SetEdgeU is used iff i >= j
func (m *AdjMxUi32) SetEdgeU(u, v VIdx, w int32) {
	m.w[uIdx(u, v)] = w
}
