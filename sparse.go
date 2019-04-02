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

func (g *SpMap) Weight(fromIdx, toIdx uint) interface{} {
	row, ok := g.ws[fromIdx]
	if !ok {
		return nil
	}
	return row[toIdx]
}

func (g *SpMap) SetWeight(fromIdx, toIdx uint, w interface{}) {
	g.sz = maxSize(g.sz, fromIdx, toIdx)
	row, rok := g.ws[fromIdx]
	if w == nil {
		if rok {
			delete(row, toIdx)
			if len(row) == 0 {
				delete(g.ws, fromIdx)
			}
		}
	} else {
		if !rok {
			row = make(smpro)
			g.ws[fromIdx] = row
		}
		row[toIdx] = w
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

func (g *SpMapf32) Edge(fromIdx, toIdx uint) (weight float32) {
	row, ok := g.ws[fromIdx]
	if !ok {
		return nan32
	}
	weight, ok = row[toIdx]
	if ok {
		return weight
	} else {
		return nan32
	}
}

func (g *SpMapf32) SetEdge(fromIdx, toIdx uint, weight float32) {
	g.sz = maxSize(g.sz, fromIdx, toIdx)
	row, rok := g.ws[fromIdx]
	if math.IsNaN(float64(weight)) {
		if rok {
			delete(row, toIdx)
			if len(row) == 0 {
				delete(g.ws, fromIdx)
			}
		}
	} else {
		if !rok {
			row = make(spmrof32)
			g.ws[fromIdx] = row
		}
		row[toIdx] = weight
	}
}

func (g *SpMapf32) Weight(fromIdx, toIdx uint) interface{} {
	tmp := g.Edge(fromIdx, toIdx)
	if math.IsNaN(float64(tmp)) {
		return nil
	}
	return tmp
}

func (g *SpMapf32) SetWeight(fromIdx, toIdx uint, w interface{}) {
	if w == nil {
		g.SetEdge(fromIdx, toIdx, nan32)
	} else {
		g.SetEdge(fromIdx, toIdx, w.(float32))
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

func (g *SpMapi32) Edge(fromIdx, toIdx uint) (weight int32, exists bool) {
	row, ok := g.ws[fromIdx]
	if !ok {
		return 0, false
	}
	weight, ok = row[toIdx]
	if ok {
		return weight, true
	} else {
		return 0, false
	}
}

func (g *SpMapi32) SetEdge(fromIdx, toIdx uint, weight int32) {
	g.sz = maxSize(g.sz, fromIdx, toIdx)
	row, rok := g.ws[fromIdx]
	if math.IsNaN(float64(weight)) {
		if rok {
			delete(row, toIdx)
			if len(row) == 0 {
				delete(g.ws, fromIdx)
			}
		}
	} else {
		if !rok {
			row = make(spmroi32)
			g.ws[fromIdx] = row
		}
		row[toIdx] = weight
	}
}

func (g *SpMapi32) DelEdge(fromIdx, toIdx uint) {
	row, ok := g.ws[fromIdx]
	if ok {
		delete(row, toIdx)
	}
}

func (g *SpMapi32) Weight(fromIdx, toIdx uint) interface{} {
	w, ok := g.Edge(fromIdx, toIdx)
	if ok {
		return w
	}
	return nil
}

func (g *SpMapi32) SetWeight(fromIdx, toIdx uint, w interface{}) {
	if w == nil {
		g.DelEdge(fromIdx, toIdx)
	} else {
		g.SetEdge(fromIdx, toIdx, w.(int32))
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
