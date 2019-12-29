package groph

type spmro map[VIdx]interface{}

type SpSoM struct {
	ws []spmro
}

func NewSpSoM(vertexNo VIdx, reuse *SpSoM) *SpSoM {
	if reuse == nil {
		return &SpSoM{make([]spmro, vertexNo)}
	}
	reuse.Reset(vertexNo)
	return reuse
}

func (g *SpSoM) VertexNo() VIdx { return len(g.ws) }

func (g *SpSoM) Weight(u, v VIdx) interface{} {
	row := g.ws[u]
	if row == nil {
		return nil
	}
	return row[v]
}

func (g *SpSoM) SetWeight(u, v VIdx, w interface{}) {
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

func (g *SpSoM) Reset(vertexNo VIdx) {
	if g.ws == nil || cap(g.ws) < vertexNo {
		g.ws = make([]spmro, vertexNo)
	} else {
		g.ws = g.ws[:vertexNo]
		for i := range g.ws {
			g.ws[i] = nil
		}
	}
}

func (g *SpSoM) EachNeighbour(v VIdx, do VisitNode) {
	if row := g.ws[v]; row != nil {
		for n := range row {
			do(n)
		}
	}
}

type spmroi32 map[VIdx]int32

type SpSoMi32 struct {
	ws  []spmroi32
	Del int32
}

func NewSpSoMi32(vertexNo VIdx, reuse *SpSoMi32) *SpSoMi32 {
	if reuse == nil {
		return &SpSoMi32{
			ws:  make([]spmroi32, vertexNo),
			Del: i32cleared,
		}
	}
	reuse.Reset(vertexNo)
	return reuse
}

func (g *SpSoMi32) VertexNo() VIdx { return len(g.ws) }

func (m *SpSoMi32) Edge(u, v VIdx) (w int32, exists bool) {
	row := m.ws[u]
	if row == nil {
		return 0, false
	}
	if res, ok := row[v]; ok {
		return res, true
	}
	return 0, false
}

func (g *SpSoMi32) Weight(u, v VIdx) interface{} {
	if res, ok := g.Edge(u, v); ok {
		return res
	}
	return nil
}

func (g *SpSoMi32) SetEdge(u, v VIdx, w int32) {
	row := g.ws[u]
	if row == nil {
		row = make(spmroi32)
		g.ws[u] = row
	}
	row[v] = w
}

func (g *SpSoMi32) DelEdge(u, v VIdx) { delete(g.ws[u], v) }

func (g *SpSoMi32) SetWeight(u, v VIdx, w interface{}) {
	if w == nil {
		g.DelEdge(u, v)
	} else {
		g.SetEdge(u, v, w.(int32))
	}
}

func (g *SpSoMi32) Reset(vertexNo VIdx) {
	if g.ws == nil || cap(g.ws) < vertexNo {
		g.ws = make([]spmroi32, vertexNo)
	} else {
		g.ws = g.ws[:vertexNo]
		for i := range g.ws {
			g.ws[i] = nil
		}
	}
}

func (g *SpSoMi32) EachNeighbour(v VIdx, do VisitNode) {
	if row := g.ws[v]; row != nil {
		for n := range row {
			do(n)
		}
	}
}

type spmrof32 map[VIdx]float32

type SpSoMf32 struct {
	ws []spmrof32
}

func NewSpSoMf32(vertexNo VIdx, reuse *SpSoMf32) *SpSoMf32 {
	if reuse == nil {
		return &SpSoMf32{make([]spmrof32, vertexNo)}
	}
	reuse.Reset(vertexNo)
	return reuse
}

func (g *SpSoMf32) VertexNo() VIdx { return len(g.ws) }

func (m *SpSoMf32) Edge(u, v VIdx) (w float32, exists bool) {
	row := m.ws[u]
	if row == nil {
		return 0, false
	}
	if res, ok := row[v]; ok {
		return res, true
	}
	return 0, false
}

func (g *SpSoMf32) Weight(u, v VIdx) interface{} {
	if res, ok := g.Edge(u, v); ok {
		return res
	}
	return nil
}

func (g *SpSoMf32) SetEdge(u, v VIdx, w float32) {
	row := g.ws[u]
	if row == nil {
		row = make(spmrof32)
		g.ws[u] = row
	}
	row[v] = w
}

func (g *SpSoMf32) DelEdge(u, v VIdx) { delete(g.ws[u], v) }

func (g *SpSoMf32) SetWeight(u, v VIdx, w interface{}) {
	if w == nil {
		g.DelEdge(u, v)
	} else {
		g.SetEdge(u, v, w.(float32))
	}
}

func (g *SpSoMf32) Reset(vertexNo VIdx) {
	if g.ws == nil || cap(g.ws) < vertexNo {
		g.ws = make([]spmrof32, vertexNo)
	} else {
		g.ws = g.ws[:vertexNo]
		for i := range g.ws {
			g.ws[i] = nil
		}
	}
}

func (g *SpSoMf32) EachNeighbour(v VIdx, do VisitNode) {
	if row := g.ws[v]; row != nil {
		for n := range row {
			do(n)
		}
	}
}
