package sliceofmaps

import "git.fractalqb.de/fractalqb/groph"

type spmrof32 map[groph.VIdx]float32

type SoMDf32 struct {
	ws []spmrof32
}

func NewSoMDf32(order int, reuse *SoMDf32) *SoMDf32 {
	if reuse == nil {
		return &SoMDf32{make([]spmrof32, order)}
	}
	reuse.Reset(order)
	return reuse
}

func (g *SoMDf32) Order() groph.VIdx { return len(g.ws) }

func (m *SoMDf32) Edge(u, v groph.VIdx) (w float32) {
	row := m.ws[u]
	if row == nil {
		return groph.NaN32()
	}
	if res, ok := row[v]; ok {
		return res
	}
	return groph.NaN32()
}

func (g *SoMDf32) Weight(u, v groph.VIdx) interface{} {
	if res := g.Edge(u, v); groph.IsNaN32(res) {
		return nil
	} else {
		return res
	}
}

func (g *SoMDf32) SetEdge(u, v groph.VIdx, w float32) {
	row := g.ws[u]
	if row == nil {
		row = make(spmrof32)
		g.ws[u] = row
	}
	row[v] = w
}

func (g *SoMDf32) SetWeight(u, v groph.VIdx, w interface{}) {
	if w == nil {
		delete(g.ws[u], v)
	} else {
		g.SetEdge(u, v, w.(float32))
	}
}

func (g *SoMDf32) Reset(order int) {
	if g.ws == nil || cap(g.ws) < order {
		g.ws = make([]spmrof32, order)
	} else {
		g.ws = g.ws[:order]
		for i := range g.ws {
			g.ws[i] = nil
		}
	}
}

func (g *SoMDf32) EachOutgoing(from groph.VIdx, onDest groph.VisitVertex) (stopped bool) {
	if row := g.ws[from]; row != nil {
		for n := range row {
			if onDest(n) {
				return true
			}
		}
	}
	return false
}

func (g *SoMDf32) OutDegree(v groph.VIdx) int {
	row := g.ws[v]
	if row == nil {
		return 0
	}
	return len(row)
}

type SoMUf32 struct {
	SoMDf32
}

func NewSoMUf32(order int, reuse *SoMUf32) *SoMUf32 {
	if reuse == nil {
		reuse = new(SoMUf32)
	}
	NewSoMDf32(order, &reuse.SoMDf32)
	return reuse
}

func (g *SoMUf32) EdgeU(u, v groph.VIdx) float32 {
	return g.SoMDf32.Edge(u, v)
}

func (g *SoMUf32) Edge(u, v groph.VIdx) float32 {
	if u > v {
		return g.EdgeU(u, v)
	}
	return g.EdgeU(v, u)
}

func (g *SoMUf32) SetEdgeU(u, v groph.VIdx, w float32) {
	g.SoMDf32.SetEdge(u, v, w)
}

func (g *SoMUf32) SetEdge(u, v groph.VIdx, w float32) {
	if u > v {
		g.SetEdgeU(u, v, w)
	} else {
		g.SetEdgeU(v, u, w)
	}
}

func (g *SoMUf32) WeightU(u, v groph.VIdx) interface{} {
	return g.SoMDf32.Weight(u, v)
}

func (g *SoMUf32) Weight(u, v groph.VIdx) interface{} {
	if u > v {
		return g.WeightU(u, v)
	}
	return g.WeightU(v, u)
}

func (g *SoMUf32) SetWeightU(u, v groph.VIdx, w interface{}) {
	g.SoMDf32.SetWeight(u, v, w)
}

func (g *SoMUf32) SetWeight(u, v groph.VIdx, w interface{}) {
	if u > v {
		g.SetWeightU(u, v, w)
	} else {
		g.SetWeightU(v, u, w)
	}
}
