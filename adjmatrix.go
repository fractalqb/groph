package groph

import (
	"errors"
	"math"

	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type adjMx struct {
	ord VIdx
}

func (m *adjMx) Order() VIdx { return m.ord }

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
			adjMx: adjMx{ord: order},
			bs:    make(BitSet, sz),
		}
	} else {
		reuse.ord = order
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
	w = m.bs.Get(m.ord*i + j)
	return w
}

func (m *AdjMxDbitmap) SetEdge(i, j VIdx, w bool) {
	if w {
		m.bs.Set(m.ord*i + j)
	} else {
		m.bs.Unset(m.ord*i + j)
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
			adjMx: adjMx{ord: order},
			bs:    make([]bool, sz),
		}
	} else {
		reuse.ord = order
		reuse.bs = iutil.BoolSlice(reuse.bs, int(sz))
	}
	return reuse
}

func AsAdjMxDbool(reuse *AdjMxDbool, weights []bool) (*AdjMxDbool, error) {
	sz, err := dOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(AdjMxDbool)
	}
	reuse.ord = sz
	reuse.bs = weights
	return reuse, nil
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
	return m.bs[m.ord*i+j]
}

func (m *AdjMxDbool) SetEdge(i, j VIdx, w bool) {
	m.bs[m.ord*i+j] = w
}

type AdjMxDi32 struct {
	adjMx
	ws  []int32
	del int32
}

func NewAdjMxDi32(order VIdx, del int32, reuse *AdjMxDi32) *AdjMxDi32 {
	if reuse == nil {
		reuse = &AdjMxDi32{
			adjMx: adjMx{ord: order},
			ws:    make([]int32, order*order),
			del:   del,
		}
	} else {
		reuse.ord = order
		reuse.ws = iutil.I32Slice(reuse.ws, int(order*order))
	}
	return reuse
}

func AsAdjMxDi32(reuse *AdjMxDi32, del int32, weights []int32) (*AdjMxDi32, error) {
	sz, err := dOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(AdjMxDi32)
	}
	reuse.ord = sz
	reuse.ws = weights
	reuse.del = del
	return reuse, nil
}

func (m *AdjMxDi32) Init(w int32) *AdjMxDi32 {
	for i := range m.ws {
		m.ws[i] = w
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
		m.SetEdge(i, j, m.del)
	} else {
		m.SetEdge(i, j, w.(int32))
	}
}

func (m *AdjMxDi32) Edge(i, j VIdx) (w int32, exists bool) {
	w = m.ws[m.ord*i+j]
	return w, w != m.del
}

func (m *AdjMxDi32) SetEdge(i, j VIdx, w int32) {
	m.ws[m.ord*i+j] = w
}

type AdjMxDf32 struct {
	adjMx
	ws []float32
}

func NewAdjMxDf32(order VIdx, reuse *AdjMxDf32) *AdjMxDf32 {
	if reuse == nil {
		reuse = &AdjMxDf32{
			adjMx: adjMx{ord: order},
			ws:    make([]float32, order*order),
		}
	} else {
		reuse.ord = order
		reuse.ws = iutil.F32Slice(reuse.ws, int(order*order))
	}
	return reuse
}

func AsAdjMxDf32(reuse *AdjMxDf32, weights []float32) (*AdjMxDf32, error) {
	sz, err := dOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(AdjMxDf32)
	}
	reuse.ord = sz
	reuse.ws = weights
	return reuse, nil
}

func (m *AdjMxDf32) Init(w float32) *AdjMxDf32 {
	for i := range m.ws {
		m.ws[i] = w
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
		m.ws[m.ord*i+j] = nan32
	} else {
		m.ws[m.ord*i+j] = w.(float32)
	}
}

func (m *AdjMxDf32) Edge(i, j VIdx) (w float32) {
	return m.ws[m.ord*i+j]
}

func (m *AdjMxDf32) SetEdge(i, j VIdx, w float32) {
	m.ws[m.ord*i+j] = w
}

type AdjMxUbool struct {
	adjMx
	ws []bool
}

func NewAdjMxUbool(order VIdx, reuse *AdjMxUbool) *AdjMxUbool {
	if reuse == nil {
		reuse = &AdjMxUbool{
			adjMx: adjMx{ord: order},
			ws:    make([]bool, nSum(order)),
		}
	} else {
		reuse.ord = order
		reuse.ws = iutil.BoolSlice(reuse.ws, int(nSum(order)))
	}
	return reuse
}

func AsAdjMxUbool(reuse *AdjMxUbool, weights []bool) (*AdjMxUbool, error) {
	sz, err := uOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(AdjMxUbool)
	}
	reuse.ord = sz
	reuse.ws = weights
	return reuse, nil
}

func (m *AdjMxUbool) Init(w bool) *AdjMxUbool {
	for i := range m.ws {
		m.ws[i] = w
	}
	return m
}

func (m *AdjMxUbool) Reset(order VIdx) { NewAdjMxUbool(order, m) }

func (m *AdjMxUbool) Weight(u, v VIdx) interface{} {
	if m.Edge(u, v) {
		return true
	}
	return nil
}

func (m *AdjMxUbool) WeightU(u, v VIdx) interface{} {
	if m.EdgeU(u, v) {
		return true
	}
	return nil
}

func (m *AdjMxUbool) SetWeight(u, v VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(u, v, false)
	} else {
		m.SetEdge(u, v, w.(bool))
	}
}

func (m *AdjMxUbool) SetWeightU(u, v VIdx, w interface{}) {
	if w == nil {
		m.SetEdgeU(u, v, false)
	} else {
		m.SetEdgeU(u, v, w.(bool))
	}
}

func (m *AdjMxUbool) Edge(u, v VIdx) (w bool) {
	if u > v {
		return m.ws[uIdx(u, v)]
	}
	return m.ws[uIdx(v, u)]
}

// EdgeU is used iff i >= j
func (m *AdjMxUbool) EdgeU(u, v VIdx) (w bool) { return m.ws[uIdx(u, v)] }

func (m *AdjMxUbool) SetEdge(i, j VIdx, w bool) {
	if i >= j {
		m.ws[uIdx(i, j)] = w
	} else {
		m.ws[uIdx(j, i)] = w
	}
}

// SetEdgeU is used iff i >= j
func (m *AdjMxUbool) SetEdgeU(u, v VIdx, w bool) { m.ws[uIdx(u, v)] = w }

type AdjMxUi32 struct {
	adjMx
	ws  []int32
	del int32
}

func NewAdjMxUi32(order VIdx, del int32, reuse *AdjMxUi32) *AdjMxUi32 {
	if reuse == nil {
		reuse = &AdjMxUi32{
			adjMx: adjMx{ord: order},
			ws:    make([]int32, nSum(order)),
			del:   del,
		}
	} else {
		reuse.ord = order
		reuse.ws = iutil.I32Slice(reuse.ws, int(nSum(order)))
	}
	return reuse
}

func AsAdjMxUi32(reuse *AdjMxUi32, del int32, weights []int32) (*AdjMxUi32, error) {
	sz, err := uOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(AdjMxUi32)
	}
	reuse.ord = sz
	reuse.ws = weights
	reuse.del = del
	return reuse, nil
}

func (m *AdjMxUi32) Init(w int32) *AdjMxUi32 {
	for i := range m.ws {
		m.ws[i] = w
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
		m.SetEdge(u, v, m.del)
	} else {
		m.SetEdge(u, v, w.(int32))
	}
}

func (m *AdjMxUi32) SetWeightU(u, v VIdx, w interface{}) {
	if w == nil {
		m.SetEdgeU(u, v, m.del)
	} else {
		m.SetEdgeU(u, v, w.(int32))
	}
}

func (m *AdjMxUi32) Edge(u, v VIdx) (w int32, ok bool) {
	if u >= v {
		w = m.ws[uIdx(u, v)]
	} else {
		w = m.ws[uIdx(v, u)]
	}
	return w, w != m.del
}

// EdgeU is used iff i >= j
func (m *AdjMxUi32) EdgeU(u, v VIdx) (w int32, ok bool) {
	w = m.ws[uIdx(u, v)]
	return w, w != m.del
}

func (m *AdjMxUi32) SetEdge(i, j VIdx, w int32) {
	if i >= j {
		m.ws[uIdx(i, j)] = w
	} else {
		m.ws[uIdx(j, i)] = w
	}
}

// SetEdgeU is used iff i >= j
func (m *AdjMxUi32) SetEdgeU(u, v VIdx, w int32) {
	m.ws[uIdx(u, v)] = w
}

type AdjMxUf32 struct {
	adjMx
	ws []float32
}

func NewAdjMxUf32(order VIdx, reuse *AdjMxUf32) *AdjMxUf32 {
	if reuse == nil {
		reuse = &AdjMxUf32{
			adjMx: adjMx{ord: order},
			ws:    make([]float32, nSum(order)),
		}
	} else {
		reuse.ord = order
		reuse.ws = iutil.F32Slice(reuse.ws, int(nSum(order)))
	}
	return reuse
}

func AsAdjMxUf32(reuse *AdjMxUf32, weights []float32) (*AdjMxUf32, error) {
	sz, err := uOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(AdjMxUf32)
	}
	reuse.ord = sz
	reuse.ws = weights
	return reuse, nil
}

func (m *AdjMxUf32) Init(w float32) *AdjMxUf32 {
	for i := range m.ws {
		m.ws[i] = w
	}
	return m
}

func (m *AdjMxUf32) Reset(order VIdx) { NewAdjMxUf32(order, m) }

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
		return m.ws[uIdx(i, j)]
	} else {
		return m.ws[uIdx(j, i)]
	}
}

// EdgeU is used iff i >= j
func (m *AdjMxUf32) EdgeU(i, j VIdx) (w float32) {
	return m.ws[uIdx(i, j)]
}

func (m *AdjMxUf32) SetEdge(i, j VIdx, w float32) {
	if i >= j {
		m.ws[uIdx(i, j)] = w
	} else {
		m.ws[uIdx(j, i)] = w
	}
}

// SetEdgeU is used iff i >= j
func (m *AdjMxUf32) SetEdgeU(i, j VIdx, w float32) {
	m.ws[uIdx(i, j)] = w
}

// uIdx computes the index into the weight slice of an undirected matrix
func uIdx(i, j VIdx) VIdx { return nSum(i) + j }

func dOrdFromLen(sliceLen int) (order int, err error) {
	order = int(math.Sqrt(float64(sliceLen)))
	if order*order != sliceLen {
		return order, errors.New("weights slice is not square")
	}
	return order, nil
}

func uOrdFromLen(sliceLen int) (order int, err error) {
	order = int(nSumRev(sliceLen))
	if nSum(order) != sliceLen {
		return order, errors.New("weights slice len is not sum(1,…,n)")
	}
	return order, nil
}
