package sliceofmaps

import "git.fractalqb.de/fractalqb/groph"

type spmroi32 map[groph.VIdx]int32

type DInt32 struct {
	ws []spmroi32
}

func NewDInt32(order int, reuse *DInt32) *DInt32 {
	if reuse == nil {
		return &DInt32{ws: make([]spmroi32, order)}
	}
	reuse.Reset(order)
	return reuse
}

func (g *DInt32) Order() int { return len(g.ws) }

func (m *DInt32) Edge(u, v groph.VIdx) (w int32, exists bool) {
	row := m.ws[u]
	if row == nil {
		return 0, false
	}
	if res, ok := row[v]; ok {
		return res, true
	}
	return 0, false
}

func (g *DInt32) Weight(u, v groph.VIdx) interface{} {
	if res, ok := g.Edge(u, v); ok {
		return res
	}
	return nil
}

func (g *DInt32) SetEdge(u, v groph.VIdx, w int32) {
	row := g.ws[u]
	if row == nil {
		row = make(spmroi32)
		g.ws[u] = row
	}
	row[v] = w
}

func (g *DInt32) SetWeight(u, v groph.VIdx, w interface{}) {
	if w == nil {
		delete(g.ws[u], v)
	} else {
		g.SetEdge(u, v, w.(int32))
	}
}

func (g *DInt32) Reset(order int) {
	if g.ws == nil || cap(g.ws) < order {
		g.ws = make([]spmroi32, order)
	} else {
		g.ws = g.ws[:order]
		for i := range g.ws {
			g.ws[i] = nil
		}
	}
}

func (g *DInt32) EachOutgoing(from groph.VIdx, onDest groph.VisitVertex) (stopped bool) {
	if row := g.ws[from]; row != nil {
		for n := range row {
			if onDest(n) {
				return true
			}
		}
	}
	return false
}

func (g *DInt32) OutDegree(v groph.VIdx) int {
	row := g.ws[v]
	if row == nil {
		return 0
	}
	return len(row)
}

type UInt32 struct {
	DInt32
}

func NewUInt32(order int, reuse *UInt32) *UInt32 {
	if reuse == nil {
		reuse = new(UInt32)
	}
	NewDInt32(order, &reuse.DInt32)
	return reuse
}

func (g *UInt32) EdgeU(u, v groph.VIdx) (w int32, exists bool) {
	return g.DInt32.Edge(u, v)
}

func (g *UInt32) Edge(u, v groph.VIdx) (w int32, exists bool) {
	if u > v {
		return g.EdgeU(u, v)
	}
	return g.EdgeU(v, u)
}

func (g *UInt32) SetEdgeU(u, v groph.VIdx, w int32) {
	g.DInt32.SetEdge(u, v, w)
}

func (g *UInt32) SetEdge(u, v groph.VIdx, w int32) {
	if u > v {
		g.SetEdgeU(u, v, w)
	} else {
		g.SetEdgeU(v, u, w)
	}
}

func (g *UInt32) WeightU(u, v groph.VIdx) interface{} {
	return g.DInt32.Weight(u, v)
}

func (g *UInt32) Weight(u, v groph.VIdx) interface{} {
	if u > v {
		return g.WeightU(u, v)
	}
	return g.WeightU(v, u)
}

func (g *UInt32) SetWeightU(u, v groph.VIdx, w interface{}) {
	g.DInt32.SetWeight(u, v, w)
}

func (g *UInt32) SetWeight(u, v groph.VIdx, w interface{}) {
	if u > v {
		g.SetWeightU(u, v, w)
	} else {
		g.SetWeightU(v, u, w)
	}
}
