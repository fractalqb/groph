package adjmatrix

import (
	"git.fractalqb.de/fractalqb/groph"
	iutil "git.fractalqb.de/fractalqb/groph/internal/util"
)

type DBool struct {
	adjMx
	bs []bool
}

func NewDBool(order int, reuse *DBool) *DBool {
	sz := order * order
	if reuse == nil {
		reuse = &DBool{
			adjMx: adjMx{ord: order},
			bs:    make([]bool, sz),
		}
	} else {
		reuse.ord = order
		reuse.bs = iutil.BoolSlice(reuse.bs, int(sz))
	}
	return reuse
}

func AsDBool(reuse *DBool, weights []bool) (*DBool, error) {
	sz, err := dOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(DBool)
	}
	reuse.ord = sz
	reuse.bs = weights
	return reuse, nil
}

func (m *DBool) Init(flag bool) *DBool {
	for i := range m.bs {
		m.bs[i] = flag
	}
	return m
}

func (m *DBool) Reset(order int) {
	NewDBool(order, m)
	m.Init(false)
}

func (m *DBool) Weight(u, v groph.VIdx) interface{} {
	if m.Edge(u, v) {
		return true
	}
	return nil
}

func (m *DBool) SetWeight(i, j groph.VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(i, j, false)
	} else {
		m.SetEdge(i, j, w.(bool))
	}
}

func (m *DBool) Edge(i, j groph.VIdx) (w bool) {
	return m.bs[m.ord*i+j]
}

func (m *DBool) SetEdge(i, j groph.VIdx, w bool) {
	m.bs[m.ord*i+j] = w
}

type UBool struct {
	adjMx
	ws []bool
}

func NewUBool(order int, reuse *UBool) *UBool {
	if reuse == nil {
		reuse = &UBool{
			adjMx: adjMx{ord: order},
			ws:    make([]bool, nSum(order)),
		}
	} else {
		reuse.ord = order
		reuse.ws = iutil.BoolSlice(reuse.ws, int(nSum(order)))
	}
	return reuse
}

func AsUBool(reuse *UBool, weights []bool) (*UBool, error) {
	sz, err := uOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(UBool)
	}
	reuse.ord = sz
	reuse.ws = weights
	return reuse, nil
}

func (m *UBool) Init(w bool) *UBool {
	for i := range m.ws {
		m.ws[i] = w
	}
	return m
}

func (m *UBool) Reset(order int) {
	NewUBool(order, m)
	m.Init(false)
}

func (m *UBool) Weight(u, v groph.VIdx) interface{} {
	if m.Edge(u, v) {
		return true
	}
	return nil
}

func (m *UBool) WeightU(u, v groph.VIdx) interface{} {
	if m.EdgeU(u, v) {
		return true
	}
	return nil
}

func (m *UBool) SetWeight(u, v groph.VIdx, w interface{}) {
	if w == nil {
		m.SetEdge(u, v, false)
	} else {
		m.SetEdge(u, v, w.(bool))
	}
}

func (m *UBool) SetWeightU(u, v groph.VIdx, w interface{}) {
	if w == nil {
		m.SetEdgeU(u, v, false)
	} else {
		m.SetEdgeU(u, v, w.(bool))
	}
}

func (m *UBool) Edge(u, v groph.VIdx) (w bool) {
	if u > v {
		return m.ws[uIdx(u, v)]
	}
	return m.ws[uIdx(v, u)]
}

// EdgeU is used iff i >= j
func (m *UBool) EdgeU(u, v groph.VIdx) (w bool) { return m.ws[uIdx(u, v)] }

func (m *UBool) SetEdge(i, j groph.VIdx, w bool) {
	if i >= j {
		m.ws[uIdx(i, j)] = w
	} else {
		m.ws[uIdx(j, i)] = w
	}
}

// SetEdgeU is used iff i >= j
func (m *UBool) SetEdgeU(u, v groph.VIdx, w bool) { m.ws[uIdx(u, v)] = w }
