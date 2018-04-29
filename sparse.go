package groph

import (
	"math"
)

type smpro = map[uint]interface{}

type SpMap struct {
	sz uint
	ws map[uint]smpro
}

var _ WGraph = (*SpMap)(nil)

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

type spmrof32 = map[uint]float32

type SpMapf32 struct {
	sz uint
	ws map[uint]spmrof32
}

var _ WGf32 = (*SpMapf32)(nil)
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

func (g *SpMapf32) Directed() bool {
	return true
}

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
			row = make(map[uint]float32)
			g.ws[fromIdx] = row
		}
		row[toIdx] = weight
	}
}

func (g *SpMapf32) Weight(fromIdx, toIdx uint) interface{} {
	return g.Edge(fromIdx, toIdx)
}

func (g *SpMapf32) SetWeight(fromIdx, toIdx uint, w interface{}) {
	g.SetEdge(fromIdx, toIdx, w.(float32))
}

func (g *SpMapf32) Clear(vertexNo uint) {
	g.sz = vertexNo
	g.ws = make(map[uint]spmrof32)
}
