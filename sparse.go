package groph

import (
	"math"
)

type smpro = map[VIdx]interface{}

type SpMap struct {
	sz VIdx
	ws map[VIdx]smpro
}

var (
	_ WGraph          = (*SpMap)(nil)
	_ NeighbourLister = (*SpMap)(nil)
)

func NewSpMap(vertexNo VIdx, reuse *SpMap) *SpMap {
	if reuse == nil {
		return &SpMap{
			sz: vertexNo,
			ws: make(map[VIdx]smpro),
		}
	} else {
		reuse.Clear(vertexNo)
		return reuse
	}
}

func (g *SpMap) VertexNo() VIdx { return g.sz }

func (g *SpMap) Directed() bool {
	return true
}

func (g *SpMap) Weight(edgeFrom, edgeTo VIdx) interface{} {
	row, ok := g.ws[edgeFrom]
	if !ok {
		return nil
	}
	return row[edgeTo]
}

func (g *SpMap) SetWeight(edgeFrom, edgeTo VIdx, w interface{}) {
	g.sz = maxSize(g.sz, edgeFrom, edgeTo)
	row, rok := g.ws[edgeFrom]
	if w == nil {
		if rok {
			delete(row, edgeTo)
			if len(row) == 0 {
				delete(g.ws, edgeFrom)
			}
		}
	} else {
		if !rok {
			row = make(smpro)
			g.ws[edgeFrom] = row
		}
		row[edgeTo] = w
	}
}

func (g *SpMap) Clear(vertexNo VIdx) {
	g.sz = vertexNo
	g.ws = make(map[VIdx]smpro)
}

func (g *SpMap) EachNeighbour(v VIdx, do func(VIdx, bool, interface{})) {
	for a, ns := range g.ws {
		if a == v {
			for b, w := range ns {
				do(b, true, w)
			}
		} else if w, ok := ns[v]; ok {
			do(a, false, w)
		}
	}
}

type spmrof32 = map[VIdx]float32

type SpMapf32 struct {
	sz VIdx
	ws map[VIdx]spmrof32
}

var (
	_ WGf32           = (*SpMapf32)(nil)
	_ NeighbourLister = (*SpMapf32)(nil)
)

var nan32 = float32(math.NaN())

func NewSpMapf32(vertexNo VIdx, reuse *SpMapf32) *SpMapf32 {
	if reuse == nil {
		return &SpMapf32{
			sz: vertexNo,
			ws: make(map[VIdx]spmrof32),
		}
	} else {
		reuse.Clear(vertexNo)
		return reuse
	}
}

func (g *SpMapf32) VertexNo() VIdx { return g.sz }

func (g *SpMapf32) Directed() bool { return true }

func (g *SpMapf32) Edge(edgeFrom, edgeTo VIdx) (weight float32) {
	row, ok := g.ws[edgeFrom]
	if !ok {
		return nan32
	}
	weight, ok = row[edgeTo]
	if ok {
		return weight
	} else {
		return nan32
	}
}

func (g *SpMapf32) SetEdge(edgeFrom, edgeTo VIdx, weight float32) {
	g.sz = maxSize(g.sz, edgeFrom, edgeTo)
	row, rok := g.ws[edgeFrom]
	if math.IsNaN(float64(weight)) {
		if rok {
			delete(row, edgeTo)
			if len(row) == 0 {
				delete(g.ws, edgeFrom)
			}
		}
	} else {
		if !rok {
			row = make(spmrof32)
			g.ws[edgeFrom] = row
		}
		row[edgeTo] = weight
	}
}

func (g *SpMapf32) Weight(edgeFrom, edgeTo VIdx) interface{} {
	tmp := g.Edge(edgeFrom, edgeTo)
	if math.IsNaN(float64(tmp)) {
		return nil
	}
	return tmp
}

func (g *SpMapf32) SetWeight(edgeFrom, edgeTo VIdx, w interface{}) {
	if w == nil {
		g.SetEdge(edgeFrom, edgeTo, nan32)
	} else {
		g.SetEdge(edgeFrom, edgeTo, w.(float32))
	}
}

func (g *SpMapf32) Clear(vertexNo VIdx) {
	g.sz = vertexNo
	g.ws = make(map[VIdx]spmrof32)
}

func (g *SpMapf32) EachNeighbour(v VIdx, do func(VIdx, bool, interface{})) {
	for a, ns := range g.ws {
		if a == v {
			for b, w := range ns {
				do(b, true, w)
			}
		} else if w, ok := ns[v]; ok {
			do(a, false, w)
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

type spmroi32 = map[VIdx]int32

type SpMapi32 struct {
	sz VIdx
	ws map[VIdx]spmroi32
}

var (
	_ WGi32           = (*SpMapi32)(nil)
	_ NeighbourLister = (*SpMapi32)(nil)
)

func NewSpMapi32(vertexNo VIdx, reuse *SpMapi32) *SpMapi32 {
	if reuse == nil {
		return &SpMapi32{
			sz: vertexNo,
			ws: make(map[VIdx]spmroi32),
		}
	} else {
		reuse.Clear(vertexNo)
		return reuse
	}
}

func (g *SpMapi32) VertexNo() VIdx { return g.sz }

func (g *SpMapi32) Directed() bool { return true }

func (g *SpMapi32) Edge(edgeFrom, edgeTo VIdx) (weight int32, exists bool) {
	row, ok := g.ws[edgeFrom]
	if !ok {
		return 0, false
	}
	weight, ok = row[edgeTo]
	if ok {
		return weight, true
	} else {
		return 0, false
	}
}

func (g *SpMapi32) SetEdge(edgeFrom, edgeTo VIdx, weight int32) {
	g.sz = maxSize(g.sz, edgeFrom, edgeTo)
	row, rok := g.ws[edgeFrom]
	if math.IsNaN(float64(weight)) {
		if rok {
			delete(row, edgeTo)
			if len(row) == 0 {
				delete(g.ws, edgeFrom)
			}
		}
	} else {
		if !rok {
			row = make(spmroi32)
			g.ws[edgeFrom] = row
		}
		row[edgeTo] = weight
	}
}

func (g *SpMapi32) DelEdge(edgeFrom, edgeTo VIdx) {
	row, ok := g.ws[edgeFrom]
	if ok {
		delete(row, edgeTo)
	}
}

func (g *SpMapi32) Weight(edgeFrom, edgeTo VIdx) interface{} {
	w, ok := g.Edge(edgeFrom, edgeTo)
	if ok {
		return w
	}
	return nil
}

func (g *SpMapi32) SetWeight(edgeFrom, edgeTo VIdx, w interface{}) {
	if w == nil {
		g.DelEdge(edgeFrom, edgeTo)
	} else {
		g.SetEdge(edgeFrom, edgeTo, w.(int32))
	}
}

func (g *SpMapi32) Clear(vertexNo VIdx) {
	g.sz = vertexNo
	g.ws = make(map[VIdx]spmroi32)
}

func (g *SpMapi32) EachNeighbour(v VIdx, do func(VIdx, bool, interface{})) {
	for a, ns := range g.ws {
		if a == v {
			for b, w := range ns {
				do(b, true, w)
			}
		} else if w, ok := ns[v]; ok {
			do(a, false, w)
		}
	}
}
