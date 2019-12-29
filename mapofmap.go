package groph

import (
	"math"
)

// Deprecated for SpSoM.
type SpMoM struct {
	sz VIdx
	ws map[VIdx]spmro
}

// Deprecated for NewSpSoM.
func NewSpMoM(vertexNo VIdx, reuse *SpMoM) *SpMoM {
	if reuse == nil {
		return &SpMoM{
			sz: vertexNo,
			ws: make(map[VIdx]spmro),
		}
	} else {
		reuse.Reset(vertexNo)
		return reuse
	}
}

func (g *SpMoM) VertexNo() VIdx { return g.sz }

func (g *SpMoM) Weight(u, v VIdx) interface{} {
	row, ok := g.ws[u]
	if !ok {
		return nil
	}
	return row[v]
}

func (g *SpMoM) SetWeight(u, v VIdx, w interface{}) {
	g.sz = maxSize(g.sz, u, v)
	row, rok := g.ws[u]
	if w == nil {
		if rok {
			delete(row, v)
			if len(row) == 0 {
				delete(g.ws, u)
			}
		}
	} else {
		if !rok {
			row = make(spmro)
			g.ws[u] = row
		}
		row[v] = w
	}
}

func (g *SpMoM) Reset(vertexNo VIdx) {
	g.sz = vertexNo
	g.ws = make(map[VIdx]spmro)
}

func (g *SpMoM) EachNeighbour(v VIdx, do VisitNode) {
	if row, ok := g.ws[v]; ok {
		for n := range row {
			do(n)
		}
	}
}

// Deprecated for SpSoMf32.
type SpMoMf32 struct {
	sz VIdx
	ws map[VIdx]spmrof32
}

// Deprecated for NewSpSoMf32.
func NewSpMoMf32(vertexNo VIdx, reuse *SpMoMf32) *SpMoMf32 {
	if reuse == nil {
		return &SpMoMf32{
			sz: vertexNo,
			ws: make(map[VIdx]spmrof32),
		}
	} else {
		reuse.Reset(vertexNo)
		return reuse
	}
}

func (g *SpMoMf32) VertexNo() VIdx { return g.sz }

func (g *SpMoMf32) Edge(u, v VIdx) (weight float32) {
	row, ok := g.ws[u]
	if !ok {
		return nan32
	}
	weight, ok = row[v]
	if ok {
		return weight
	} else {
		return nan32
	}
}

func (g *SpMoMf32) SetEdge(u, v VIdx, weight float32) {
	g.sz = maxSize(g.sz, u, v)
	row, rok := g.ws[u]
	if math.IsNaN(float64(weight)) {
		if rok {
			delete(row, v)
			if len(row) == 0 {
				delete(g.ws, u)
			}
		}
	} else {
		if !rok {
			row = make(spmrof32)
			g.ws[u] = row
		}
		row[v] = weight
	}
}

func (g *SpMoMf32) Weight(u, v VIdx) interface{} {
	tmp := g.Edge(u, v)
	if math.IsNaN(float64(tmp)) {
		return nil
	}
	return tmp
}

func (g *SpMoMf32) SetWeight(u, v VIdx, w interface{}) {
	if w == nil {
		g.SetEdge(u, v, nan32)
	} else {
		g.SetEdge(u, v, w.(float32))
	}
}

func (g *SpMoMf32) Reset(vertexNo VIdx) {
	g.sz = vertexNo
	g.ws = make(map[VIdx]spmrof32)
}

func (g *SpMoMf32) EachNeighbour(v VIdx, do VisitNode) {
	if row, ok := g.ws[v]; ok {
		for n := range row {
			do(n)
		}
	}
}

func maxSize(currentSize, newIdx1, newIdx2 VIdx) VIdx {
	if s := newIdx1 + 1; s > currentSize {
		if newIdx2 > newIdx1 {
			return newIdx2 + 1
		} else {
			return s
		}
	}
	if s := newIdx2 + 1; s > currentSize {
		return s
	}
	return currentSize
}

// Deprecated for SpSoMi32.
type SpMoMi32 struct {
	sz VIdx
	ws map[VIdx]spmroi32
}

// Deprecated for NewSpSoMi32.
func NewSpMoMi32(vertexNo VIdx, reuse *SpMoMi32) *SpMoMi32 {
	if reuse == nil {
		return &SpMoMi32{
			sz: vertexNo,
			ws: make(map[VIdx]spmroi32),
		}
	} else {
		reuse.Reset(vertexNo)
		return reuse
	}
}

func (g *SpMoMi32) VertexNo() VIdx { return g.sz }

func (g *SpMoMi32) Edge(u, v VIdx) (w int32, ok bool) {
	row, ok := g.ws[u]
	if !ok {
		return 0, false
	}
	w, ok = row[v]
	if ok {
		return w, true
	} else {
		return 0, false
	}
}

func (g *SpMoMi32) SetEdge(u, v VIdx, weight int32) {
	g.sz = maxSize(g.sz, u, v)
	row, rok := g.ws[u]
	if math.IsNaN(float64(weight)) {
		if rok {
			delete(row, v)
			if len(row) == 0 {
				delete(g.ws, u)
			}
		}
	} else {
		if !rok {
			row = make(spmroi32)
			g.ws[u] = row
		}
		row[v] = weight
	}
}

func (g *SpMoMi32) DelEdge(u, v VIdx) {
	row, ok := g.ws[u]
	if ok {
		delete(row, v)
	}
}

func (g *SpMoMi32) Weight(u, v VIdx) interface{} {
	w, ok := g.Edge(u, v)
	if ok {
		return w
	}
	return nil
}

func (g *SpMoMi32) SetWeight(u, v VIdx, w interface{}) {
	if w == nil {
		g.DelEdge(u, v)
	} else {
		g.SetEdge(u, v, w.(int32))
	}
}

func (g *SpMoMi32) Reset(vertexNo VIdx) {
	g.sz = vertexNo
	g.ws = make(map[VIdx]spmroi32)
}

func (g *SpMoMi32) EachNeighbour(v VIdx, do VisitNode) {
	if row, ok := g.ws[v]; ok {
		for n := range row {
			do(n)
		}
	}
}
