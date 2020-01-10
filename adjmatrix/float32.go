package adjmatrix

import (
	"math"

	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type DFloat32 struct {
	adjMx
	ws []float32
}

func NewDFloat32(order int, reuse *DFloat32) *DFloat32 {
	if reuse == nil {
		reuse = &DFloat32{
			adjMx: adjMx{ord: order},
			ws:    make([]float32, order*order),
		}
	} else {
		reuse.ord = order
		reuse.ws = iutil.F32Slice(reuse.ws, int(order*order))
	}
	return reuse
}

func AsDFloat32(reuse *DFloat32, weights []float32) (*DFloat32, error) {
	sz, err := dOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(DFloat32)
	}
	reuse.ord = sz
	reuse.ws = weights
	return reuse, nil
}

func (m *DFloat32) Init(w float32) *DFloat32 {
	for i := range m.ws {
		m.ws[i] = w
	}
	return m
}

func (m *DFloat32) Reset(order int) {
	NewDFloat32(order, m)
	m.Init(groph.NaN32())
}

func (m *DFloat32) Weight(u, v groph.VIdx) interface{} {
	w := m.Edge(u, v)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *DFloat32) SetWeight(i, j groph.VIdx, w interface{}) {
	if w == nil {
		m.ws[m.ord*i+j] = groph.NaN32()
	} else {
		m.ws[m.ord*i+j] = w.(float32)
	}
}

func (m *DFloat32) Edge(i, j groph.VIdx) (w float32) {
	return m.ws[m.ord*i+j]
}

func (m *DFloat32) SetEdge(i, j groph.VIdx, w float32) {
	m.ws[m.ord*i+j] = w
}

type UFloat32 struct {
	adjMx
	ws []float32
}

func NewUFloat32(order int, reuse *UFloat32) *UFloat32 {
	if reuse == nil {
		reuse = &UFloat32{
			adjMx: adjMx{ord: order},
			ws:    make([]float32, nSum(order)),
		}
	} else {
		reuse.ord = order
		reuse.ws = iutil.F32Slice(reuse.ws, int(nSum(order)))
	}
	return reuse
}

func AsUFloat32(reuse *UFloat32, weights []float32) (*UFloat32, error) {
	sz, err := uOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(UFloat32)
	}
	reuse.ord = sz
	reuse.ws = weights
	return reuse, nil
}

func (m *UFloat32) Init(w float32) *UFloat32 {
	for i := range m.ws {
		m.ws[i] = w
	}
	return m
}

func (m *UFloat32) Reset(order int) {
	NewUFloat32(order, m)
	m.Init(groph.NaN32())
}

func (m *UFloat32) Weight(i, j groph.VIdx) interface{} {
	w := m.Edge(i, j)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *UFloat32) WeightU(i, j groph.VIdx) interface{} {
	w := m.EdgeU(i, j)
	if math.IsNaN(float64(w)) {
		return nil
	} else {
		return w
	}
}

func (m *UFloat32) SetWeight(i, j groph.VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(i, j, groph.NaN32())
	} else {
		m.SetEdge(i, j, w.(float32))
	}
}

func (m *UFloat32) SetWeightU(i, j groph.VIdx, w interface{}) {
	if w == nil {
		m.SetEdgeU(i, j, groph.NaN32())
	} else {
		m.SetEdgeU(i, j, w.(float32))
	}
}

func (m *UFloat32) Edge(i, j groph.VIdx) (w float32) {
	if i >= j {
		return m.ws[uIdx(i, j)]
	} else {
		return m.ws[uIdx(j, i)]
	}
}

// EdgeU is used iff i >= j
func (m *UFloat32) EdgeU(i, j groph.VIdx) (w float32) {
	return m.ws[uIdx(i, j)]
}

func (m *UFloat32) SetEdge(i, j groph.VIdx, w float32) {
	if i >= j {
		m.ws[uIdx(i, j)] = w
	} else {
		m.ws[uIdx(j, i)] = w
	}
}

// SetEdgeU is used iff i >= j
func (m *UFloat32) SetEdgeU(i, j groph.VIdx, w float32) {
	m.ws[uIdx(i, j)] = w
}
