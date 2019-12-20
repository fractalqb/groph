package groph

import (
	"math"
)

type smpro = map[uint]interface{}

type SpMap struct {
	sz uint
	ws map[uint]smpro
}

var (
	_ WGraph          = (*SpMap)(nil)
	_ ListNeightbours = (*SpMap)(nil)
)

func NewSpMap(vertexNo uint, reuse *SpMap) *SpMap {
	if reuse == nil {
		return &SpMap{
			sz: vertexNo,
			ws: make(map[uint]smpro),
		}
	} else {
		reuse.Clear(vertexNo)
		return reuse
	}
}

func (g *SpMap) VertexNo() uint { return g.sz }

func (g *SpMap) Directed() bool {
	return true
}

func (g *SpMap) Weight(edgeFrom, edgeTo uint) interface{} {
	row, ok := g.ws[edgeFrom]
	if !ok {
		return nil
	}
	return row[edgeTo]
}

func (g *SpMap) SetWeight(edgeFrom, edgeTo uint, w interface{}) {
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

func (g *SpMap) Clear(vertexNo uint) {
	g.sz = vertexNo
	g.ws = make(map[uint]smpro)
}

func (g *SpMap) EachNeighbour(v uint, do func(uint, bool, interface{})) {
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

type spmrof32 = map[uint]float32

type SpMapf32 struct {
	sz uint
	ws map[uint]spmrof32
}

var (
	_ WGf32           = (*SpMapf32)(nil)
	_ ListNeightbours = (*SpMapf32)(nil)
)

var nan32 = float32(math.NaN())

func NewSpMapf32(vertexNo uint, reuse *SpMapf32) *SpMapf32 {
	if reuse == nil {
		return &SpMapf32{
			sz: vertexNo,
			ws: make(map[uint]spmrof32),
		}
	} else {
		reuse.Clear(vertexNo)
		return reuse
	}
}

func (g *SpMapf32) VertexNo() uint { return g.sz }

func (g *SpMapf32) Directed() bool { return true }

func (g *SpMapf32) Edge(edgeFrom, edgeTo uint) (weight float32) {
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

func (g *SpMapf32) SetEdge(edgeFrom, edgeTo uint, weight float32) {
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

func (g *SpMapf32) Weight(edgeFrom, edgeTo uint) interface{} {
	tmp := g.Edge(edgeFrom, edgeTo)
	if math.IsNaN(float64(tmp)) {
		return nil
	}
	return tmp
}

func (g *SpMapf32) SetWeight(edgeFrom, edgeTo uint, w interface{}) {
	if w == nil {
		g.SetEdge(edgeFrom, edgeTo, nan32)
	} else {
		g.SetEdge(edgeFrom, edgeTo, w.(float32))
	}
}

func (g *SpMapf32) Clear(vertexNo uint) {
	g.sz = vertexNo
	g.ws = make(map[uint]spmrof32)
}

func (g *SpMapf32) EachNeighbour(v uint, do func(uint, bool, interface{})) {
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

func maxSize(currentSize, newIdx1, newIdx2 uint) uint {
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

type spmroi32 = map[uint]int32

type SpMapi32 struct {
	sz uint
	ws map[uint]spmroi32
}

var (
	_ WGi32           = (*SpMapi32)(nil)
	_ ListNeightbours = (*SpMapi32)(nil)
)

func NewSpMapi32(vertexNo uint, reuse *SpMapi32) *SpMapi32 {
	if reuse == nil {
		return &SpMapi32{
			sz: vertexNo,
			ws: make(map[uint]spmroi32),
		}
	} else {
		reuse.Clear(vertexNo)
		return reuse
	}
}

func (g *SpMapi32) VertexNo() uint { return g.sz }

func (g *SpMapi32) Directed() bool { return true }

func (g *SpMapi32) Edge(edgeFrom, edgeTo uint) (weight int32, exists bool) {
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

func (g *SpMapi32) SetEdge(edgeFrom, edgeTo uint, weight int32) {
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

func (g *SpMapi32) DelEdge(edgeFrom, edgeTo uint) {
	row, ok := g.ws[edgeFrom]
	if ok {
		delete(row, edgeTo)
	}
}

func (g *SpMapi32) Weight(edgeFrom, edgeTo uint) interface{} {
	w, ok := g.Edge(edgeFrom, edgeTo)
	if ok {
		return w
	}
	return nil
}

func (g *SpMapi32) SetWeight(edgeFrom, edgeTo uint, w interface{}) {
	if w == nil {
		g.DelEdge(edgeFrom, edgeTo)
	} else {
		g.SetEdge(edgeFrom, edgeTo, w.(int32))
	}
}

func (g *SpMapi32) Clear(vertexNo uint) {
	g.sz = vertexNo
	g.ws = make(map[uint]spmroi32)
}

func (g *SpMapi32) EachNeighbour(v uint, do func(uint, bool, interface{})) {
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
