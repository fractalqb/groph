package adjmatrix

import (
	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type DInt32 struct {
	adjMx
	ws  []int32
	del int32
}

func NewDInt32(order int, del int32, reuse *DInt32) *DInt32 {
	if reuse == nil {
		reuse = &DInt32{
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

func AsDInt32(reuse *DInt32, del int32, weights []int32) (*DInt32, error) {
	sz, err := dOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(DInt32)
	}
	reuse.ord = sz
	reuse.ws = weights
	reuse.del = del
	return reuse, nil
}

func (m *DInt32) Del() int32 { return m.del }

func (m *DInt32) Init(w int32) *DInt32 {
	for i := range m.ws {
		m.ws[i] = w
	}
	return m
}

func (m *DInt32) Reset(order int) {
	NewDInt32(order, m.del, m)
	m.Init(m.del)
}

func (m *DInt32) Weight(u, v groph.VIdx) interface{} {
	res, ok := m.Edge(u, v)
	if ok {
		return res
	}
	return nil
}

func (m *DInt32) SetWeight(i, j groph.VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(i, j, m.del)
	} else {
		m.SetEdge(i, j, w.(int32))
	}
}

func (m *DInt32) Edge(i, j groph.VIdx) (w int32, exists bool) {
	w = m.ws[m.ord*i+j]
	return w, w != m.del
}

func (m *DInt32) SetEdge(i, j groph.VIdx, w int32) {
	m.ws[m.ord*i+j] = w
}

type UInt32 struct {
	adjMx
	ws  []int32
	del int32
}

func NewUInt32(order int, del int32, reuse *UInt32) *UInt32 {
	if reuse == nil {
		reuse = &UInt32{
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

func AsUInt32(reuse *UInt32, del int32, weights []int32) (*UInt32, error) {
	sz, err := uOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(UInt32)
	}
	reuse.ord = sz
	reuse.ws = weights
	reuse.del = del
	return reuse, nil
}

func (m *UInt32) Del() int32 { return m.del }

func (m *UInt32) Init(w int32) *UInt32 {
	for i := range m.ws {
		m.ws[i] = w
	}
	return m
}

func (m *UInt32) Reset(order int) {
	NewUInt32(order, m.del, m)
	m.Init(m.del)
}

func (m *UInt32) Weight(u, v groph.VIdx) interface{} {
	if w, ok := m.Edge(u, v); ok {
		return w
	}
	return nil
}

func (m *UInt32) WeightU(u, v groph.VIdx) interface{} {
	w, ok := m.EdgeU(u, v)
	if ok {
		return w
	}
	return nil
}

func (m *UInt32) SetWeight(u, v groph.VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(u, v, m.del)
	} else {
		m.SetEdge(u, v, w.(int32))
	}
}

func (m *UInt32) SetWeightU(u, v groph.VIdx, w interface{}) {
	if w == nil {
		m.SetEdgeU(u, v, m.del)
	} else {
		m.SetEdgeU(u, v, w.(int32))
	}
}

func (m *UInt32) Edge(u, v groph.VIdx) (w int32, ok bool) {
	if u >= v {
		w = m.ws[uIdx(u, v)]
	} else {
		w = m.ws[uIdx(v, u)]
	}
	return w, w != m.del
}

func (m *UInt32) EdgeU(u, v groph.VIdx) (w int32, ok bool) {
	w = m.ws[uIdx(u, v)]
	return w, w != m.del
}

func (m *UInt32) SetEdge(i, j groph.VIdx, w int32) {
	if i >= j {
		m.ws[uIdx(i, j)] = w
	} else {
		m.ws[uIdx(j, i)] = w
	}
}

func (m *UInt32) SetEdgeU(u, v groph.VIdx, w int32) {
	m.ws[uIdx(u, v)] = w
}
