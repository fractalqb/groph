package groph

import (
	"math"
)

type smpro = map[VIdx]interface{}

type SpMap struct {
	sz VIdx
	ws map[VIdx]smpro
}

func NewSpMap(vertexNo VIdx, reuse *SpMap) *SpMap {
	if reuse == nil {
		return &SpMap{
			sz: vertexNo,
			ws: make(map[VIdx]smpro),
		}
	} else {
		reuse.Reset(vertexNo)
		return reuse
	}
}

func (g *SpMap) VertexNo() VIdx { return g.sz }

func (g *SpMap) Weight(u, v VIdx) interface{} {
	row, ok := g.ws[u]
	if !ok {
		return nil
	}
	return row[v]
}

func (g *SpMap) SetWeight(u, v VIdx, w interface{}) {
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
			row = make(smpro)
			g.ws[u] = row
		}
		row[v] = w
	}
}

func (g *SpMap) Reset(vertexNo VIdx) {
	g.sz = vertexNo
	g.ws = make(map[VIdx]smpro)
}

func (g *SpMap) EachNeighbour(v VIdx, do VisitNeighbour) {
	row, ok := g.ws[v]
	if !ok {
		return
	}
	for n := range row {
		do(n)
	}
}

type spmrof32 = map[VIdx]float32

type SpMapf32 struct {
	sz VIdx
	ws map[VIdx]spmrof32
}

var nan32 = float32(math.NaN())

func NewSpMapf32(vertexNo VIdx, reuse *SpMapf32) *SpMapf32 {
	if reuse == nil {
		return &SpMapf32{
			sz: vertexNo,
			ws: make(map[VIdx]spmrof32),
		}
	} else {
		reuse.Reset(vertexNo)
		return reuse
	}
}

func (g *SpMapf32) VertexNo() VIdx { return g.sz }

func (g *SpMapf32) Edge(u, v VIdx) (weight float32) {
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

func (g *SpMapf32) SetEdge(u, v VIdx, weight float32) {
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

func (g *SpMapf32) Weight(u, v VIdx) interface{} {
	tmp := g.Edge(u, v)
	if math.IsNaN(float64(tmp)) {
		return nil
	}
	return tmp
}

func (g *SpMapf32) SetWeight(u, v VIdx, w interface{}) {
	if w == nil {
		g.SetEdge(u, v, nan32)
	} else {
		g.SetEdge(u, v, w.(float32))
	}
}

func (g *SpMapf32) Reset(vertexNo VIdx) {
	g.sz = vertexNo
	g.ws = make(map[VIdx]spmrof32)
}

func (g *SpMapf32) EachNeighbour(v VIdx, do VisitNeighbour) {
	row, ok := g.ws[v]
	if !ok {
		return
	}
	for n := range row {
		do(n)
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

type spmroi32 = map[VIdx]int32

type SpMapi32 struct {
	sz VIdx
	ws map[VIdx]spmroi32
}

func NewSpMapi32(vertexNo VIdx, reuse *SpMapi32) *SpMapi32 {
	if reuse == nil {
		return &SpMapi32{
			sz: vertexNo,
			ws: make(map[VIdx]spmroi32),
		}
	} else {
		reuse.Reset(vertexNo)
		return reuse
	}
}

func (g *SpMapi32) VertexNo() VIdx { return g.sz }

func (g *SpMapi32) Edge(u, v VIdx) (w int32, ok bool) {
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

func (g *SpMapi32) SetEdge(u, v VIdx, weight int32) {
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

func (g *SpMapi32) DelEdge(u, v VIdx) {
	row, ok := g.ws[u]
	if ok {
		delete(row, v)
	}
}

func (g *SpMapi32) Weight(u, v VIdx) interface{} {
	w, ok := g.Edge(u, v)
	if ok {
		return w
	}
	return nil
}

func (g *SpMapi32) SetWeight(u, v VIdx, w interface{}) {
	if w == nil {
		g.DelEdge(u, v)
	} else {
		g.SetEdge(u, v, w.(int32))
	}
}

func (g *SpMapi32) Reset(vertexNo VIdx) {
	g.sz = vertexNo
	g.ws = make(map[VIdx]spmroi32)
}

func (g *SpMapi32) EachNeighbour(v VIdx, do VisitNeighbour) {
	row, ok := g.ws[v]
	if !ok {
		return
	}
	for n := range row {
		do(n)
	}
}
