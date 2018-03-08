package groph

import (
	"math"
)

type spmrof32 = map[uint]float32

type SpMapf32 struct {
	sz uint
	vp func(idx uint) Vertex
	ws map[uint]spmrof32
}

var _ WGf32 = (*SpMapf32)(nil)
var nan32 = float32(math.NaN())

func (g *SpMapf32) VertexNo() uint { return g.sz }

func (g *SpMapf32) Vertex(idx uint) Vertex {
	return g.vp(idx)
}

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
