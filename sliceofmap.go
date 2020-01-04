package groph

type spmro map[VIdx]interface{}

type SoMD struct {
	ws []spmro
}

func NewSoMD(order VIdx, reuse *SoMD) *SoMD {
	if reuse == nil {
		return &SoMD{make([]spmro, order)}
	}
	reuse.Reset(order)
	return reuse
}

func (g *SoMD) Order() VIdx { return len(g.ws) }

func (g *SoMD) Weight(u, v VIdx) interface{} {
	row := g.ws[u]
	if row == nil {
		return nil
	}
	return row[v]
}

func (g *SoMD) SetWeight(u, v VIdx, w interface{}) {
	row := g.ws[u]
	if w == nil {
		delete(row, v)
	} else {
		if row == nil {
			row = make(spmro)
			g.ws[u] = row
		}
		row[v] = w
	}
}

func (g *SoMD) Reset(order VIdx) {
	if g.ws == nil || cap(g.ws) < order {
		g.ws = make([]spmro, order)
	} else {
		g.ws = g.ws[:order]
		for i := range g.ws {
			g.ws[i] = nil
		}
	}
}

func (g *SoMD) EachOutgoing(from VIdx, onDest VisitVertex) (stopped bool) {
	if row := g.ws[from]; row != nil {
		for n := range row {
			if onDest(n) {
				return true
			}
		}
	}
	return false
}

func (g *SoMD) OutDegree(v VIdx) int {
	row := g.ws[v]
	if row == nil {
		return 0
	}
	return len(row)
}

type SoMU struct {
	SoMD
}

func NewSoMU(order VIdx, reuse *SoMU) *SoMU {
	if reuse == nil {
		reuse = new(SoMU)
	}
	NewSoMD(order, &reuse.SoMD)
	return reuse
}

func (g *SoMU) WeightU(u, v VIdx) interface{} {
	return g.SoMD.Weight(u, v)
}

func (g *SoMU) Weight(u, v VIdx) interface{} {
	if u > v {
		return g.WeightU(u, v)
	}
	return g.WeightU(v, u)
}

func (g *SoMU) SetWeightU(u, v VIdx, w interface{}) {
	g.SoMD.SetWeight(u, v, w)
}

func (g *SoMU) SetWeight(u, v VIdx, w interface{}) {
	if u > v {
		g.SetWeightU(u, v, w)
	} else {
		g.SetWeightU(v, u, w)
	}
}

type spmroi32 map[VIdx]int32

type SoMDi32 struct {
	ws  []spmroi32
	Del int32
}

func NewSoMDi32(order VIdx, reuse *SoMDi32) *SoMDi32 {
	if reuse == nil {
		return &SoMDi32{
			ws:  make([]spmroi32, order),
			Del: I32Del,
		}
	}
	reuse.Reset(order)
	return reuse
}

func (g *SoMDi32) Order() VIdx { return len(g.ws) }

func (m *SoMDi32) Edge(u, v VIdx) (w int32, exists bool) {
	row := m.ws[u]
	if row == nil {
		return 0, false
	}
	if res, ok := row[v]; ok {
		return res, true
	}
	return 0, false
}

func (g *SoMDi32) Weight(u, v VIdx) interface{} {
	if res, ok := g.Edge(u, v); ok {
		return res
	}
	return nil
}

func (g *SoMDi32) SetEdge(u, v VIdx, w int32) {
	row := g.ws[u]
	if row == nil {
		row = make(spmroi32)
		g.ws[u] = row
	}
	row[v] = w
}

func (g *SoMDi32) SetWeight(u, v VIdx, w interface{}) {
	if w == nil {
		delete(g.ws[u], v)
	} else {
		g.SetEdge(u, v, w.(int32))
	}
}

func (g *SoMDi32) Reset(order VIdx) {
	if g.ws == nil || cap(g.ws) < order {
		g.ws = make([]spmroi32, order)
	} else {
		g.ws = g.ws[:order]
		for i := range g.ws {
			g.ws[i] = nil
		}
	}
}

func (g *SoMDi32) EachOutgoing(from VIdx, onDest VisitVertex) (stopped bool) {
	if row := g.ws[from]; row != nil {
		for n := range row {
			if onDest(n) {
				return true
			}
		}
	}
	return false
}

func (g *SoMDi32) OutDegree(v VIdx) int {
	row := g.ws[v]
	if row == nil {
		return 0
	}
	return len(row)
}

type SoMUi32 struct {
	SoMDi32
}

func NewSoMUi32(order VIdx, reuse *SoMUi32) *SoMUi32 {
	if reuse == nil {
		reuse = new(SoMUi32)
	}
	NewSoMDi32(order, &reuse.SoMDi32)
	return reuse
}

func (g *SoMUi32) EdgeU(u, v VIdx) (w int32, exists bool) {
	return g.SoMDi32.Edge(u, v)
}

func (g *SoMUi32) Edge(u, v VIdx) (w int32, exists bool) {
	if u > v {
		return g.EdgeU(u, v)
	}
	return g.EdgeU(v, u)
}

func (g *SoMUi32) SetEdgeU(u, v VIdx, w int32) {
	g.SoMDi32.SetEdge(u, v, w)
}

func (g *SoMUi32) SetEdge(u, v VIdx, w int32) {
	if u > v {
		g.SetEdgeU(u, v, w)
	} else {
		g.SetEdgeU(v, u, w)
	}
}

func (g *SoMUi32) WeightU(u, v VIdx) interface{} {
	return g.SoMDi32.Weight(u, v)
}

func (g *SoMUi32) Weight(u, v VIdx) interface{} {
	if u > v {
		return g.WeightU(u, v)
	}
	return g.WeightU(v, u)
}

func (g *SoMUi32) SetWeightU(u, v VIdx, w interface{}) {
	g.SoMDi32.SetWeight(u, v, w)
}

func (g *SoMUi32) SetWeight(u, v VIdx, w interface{}) {
	if u > v {
		g.SetWeightU(u, v, w)
	} else {
		g.SetWeightU(v, u, w)
	}
}

type spmrof32 map[VIdx]float32

type SoMDf32 struct {
	ws []spmrof32
}

func NewSoMDf32(order VIdx, reuse *SoMDf32) *SoMDf32 {
	if reuse == nil {
		return &SoMDf32{make([]spmrof32, order)}
	}
	reuse.Reset(order)
	return reuse
}

func (g *SoMDf32) Order() VIdx { return len(g.ws) }

func (m *SoMDf32) Edge(u, v VIdx) (w float32) {
	row := m.ws[u]
	if row == nil {
		return NaN32()
	}
	if res, ok := row[v]; ok {
		return res
	}
	return NaN32()
}

func (g *SoMDf32) Weight(u, v VIdx) interface{} {
	if res := g.Edge(u, v); IsNaN32(res) {
		return nil
	} else {
		return res
	}
}

func (g *SoMDf32) SetEdge(u, v VIdx, w float32) {
	row := g.ws[u]
	if row == nil {
		row = make(spmrof32)
		g.ws[u] = row
	}
	row[v] = w
}

func (g *SoMDf32) SetWeight(u, v VIdx, w interface{}) {
	if w == nil {
		delete(g.ws[u], v)
	} else {
		g.SetEdge(u, v, w.(float32))
	}
}

func (g *SoMDf32) Reset(order VIdx) {
	if g.ws == nil || cap(g.ws) < order {
		g.ws = make([]spmrof32, order)
	} else {
		g.ws = g.ws[:order]
		for i := range g.ws {
			g.ws[i] = nil
		}
	}
}

func (g *SoMDf32) EachOutgoing(from VIdx, onDest VisitVertex) (stopped bool) {
	if row := g.ws[from]; row != nil {
		for n := range row {
			if onDest(n) {
				return true
			}
		}
	}
	return false
}

func (g *SoMDf32) OutDegree(v VIdx) int {
	row := g.ws[v]
	if row == nil {
		return 0
	}
	return len(row)
}

type SoMUf32 struct {
	SoMDf32
}

func NewSoMUf32(order VIdx, reuse *SoMUf32) *SoMUf32 {
	if reuse == nil {
		reuse = new(SoMUf32)
	}
	NewSoMDf32(order, &reuse.SoMDf32)
	return reuse
}

func (g *SoMUf32) EdgeU(u, v VIdx) float32 {
	return g.SoMDf32.Edge(u, v)
}

func (g *SoMUf32) Edge(u, v VIdx) float32 {
	if u > v {
		return g.EdgeU(u, v)
	}
	return g.EdgeU(v, u)
}

func (g *SoMUf32) SetEdgeU(u, v VIdx, w float32) {
	g.SoMDf32.SetEdge(u, v, w)
}

func (g *SoMUf32) SetEdge(u, v VIdx, w float32) {
	if u > v {
		g.SetEdgeU(u, v, w)
	} else {
		g.SetEdgeU(v, u, w)
	}
}

func (g *SoMUf32) WeightU(u, v VIdx) interface{} {
	return g.SoMDf32.Weight(u, v)
}

func (g *SoMUf32) Weight(u, v VIdx) interface{} {
	if u > v {
		return g.WeightU(u, v)
	}
	return g.WeightU(v, u)
}

func (g *SoMUf32) SetWeightU(u, v VIdx, w interface{}) {
	g.SoMDf32.SetWeight(u, v, w)
}

func (g *SoMUf32) SetWeight(u, v VIdx, w interface{}) {
	if u > v {
		g.SetWeightU(u, v, w)
	} else {
		g.SetWeightU(v, u, w)
	}
}
