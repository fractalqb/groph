package groph

import (
	"fmt"
	"math"
	"reflect"
)

type Slice struct {
	slc reflect.Value
	dir bool
	sz  VIdx
}

func NewSlice(directed bool, edgeSlice interface{}) *Slice {
	res := &Slice{
		slc: reflect.ValueOf(edgeSlice),
		dir: directed,
	}
	if directed {
		res.sz = VIdx(math.Sqrt(float64(res.slc.Len())))
	} else {
		res.sz = VIdx(nSumRev(VIdx(res.slc.Len())))
	}
	return res
}

func (g *Slice) Check() (*Slice, error) {
	if g.slc.Kind() != reflect.Slice {
		return g, fmt.Errorf("edges have to be a slice, got %s",
			g.slc.Type().String())
	}
	if g.dir {
		if g.sz*g.sz != VIdx(g.slc.Len()) {
			return g, fmt.Errorf("slice len is not quadratic")
		}
	} else if nSum(g.sz) != VIdx(g.slc.Len()) {
		return g, fmt.Errorf("slice len is not Sum(1, ..., n)")
	}
	return g, nil
}

func (g *Slice) Must() *Slice {
	var err error
	g, err = g.Check()
	if err != nil {
		panic(err)
	}
	return g
}

func (g *Slice) VertexNo() VIdx { return g.sz }

func (g *Slice) Directed() bool { return g.dir }

func (g *Slice) Weight(edgeFrom, edgeTo VIdx) interface{} {
	if g.dir {
		return g.slc.Index(int(g.sz*edgeFrom + edgeTo)).Interface()
	} else if edgeFrom < edgeTo {
		return g.slc.Index(int(uIdx(g.sz, edgeFrom, edgeTo))).Interface()
	}
	return g.slc.Index(int(uIdx(g.sz, edgeTo, edgeFrom))).Interface()
}

// SliceNMeasure implements a metric RGraph based on a slice of vertices of
// some type V and a function f(u V, v V) → W that computes the weight of type
// W for any edge (u, v).
//
// From this use e.g. CpWeights or CpXWeights to initialize an other WGraph.
type SliceNMeasure struct {
	slc reflect.Value
	m   reflect.Value
	dir bool
}

func NewSliceNMeasure(
	vertexSlice interface{},
	measure interface{},
	directed bool) *SliceNMeasure {
	res := &SliceNMeasure{
		slc: reflect.ValueOf(vertexSlice),
		m:   reflect.ValueOf(measure),
		dir: directed,
	}
	return res
}

// Check does type checking on g. It always returns g. Only if everything is OK
// the returned error is nil.
func (g *SliceNMeasure) Check() (*SliceNMeasure, error) {
	if g.slc.Kind() != reflect.Slice {
		return g, fmt.Errorf("vertex set has to be a slice, got %s",
			g.slc.Type().String())
	}
	if g.m.Kind() != reflect.Func {
		return g, fmt.Errorf("edge weight measure must be a function, got: %s",
			g.m.Type().String())
	} // TODO more precise checking
	return g, nil
}

// Verify call Check on g and panics if Check returns an error.
func (g *SliceNMeasure) Must() *SliceNMeasure {
	if _, err := g.Check(); err != nil {
		panic(err)
	}
	return g
}

func (g *SliceNMeasure) VertexNo() VIdx {
	return VIdx(g.slc.Len())
}

func (g *SliceNMeasure) Vertex(idx VIdx) interface{} {
	return g.slc.Index(int(idx)).Interface()
}

func (g *SliceNMeasure) Directed() bool {
	return g.dir
}

func (g *SliceNMeasure) Weight(edgeFrom, edgeTo VIdx) interface{} {
	f, t := g.slc.Index(int(edgeFrom)), g.slc.Index(int(edgeTo))
	d := g.m.Call([]reflect.Value{f, t})
	return d[0].Interface()
}

type RSubgraph struct {
	g   RGraph
	vls []VIdx
}

var _ RGraph = (*RSubgraph)(nil)

func (g *RSubgraph) VertexNo() VIdx {
	return VIdx(len(g.vls))
}

func (g *RSubgraph) Directed() bool {
	return g.g.Directed()
}

func (g *RSubgraph) Weight(edgeFrom, edgeTo VIdx) interface{} {
	edgeFrom = g.vls[edgeFrom]
	edgeTo = g.vls[edgeTo]
	return g.Weight(edgeFrom, edgeTo)
}

type WSubgraph struct {
	g   WGraph
	vls []VIdx
}

var _ WGraph = (*WSubgraph)(nil)

func (g *WSubgraph) VertexNo() VIdx {
	return VIdx(len(g.vls))
}

func (g *WSubgraph) Directed() bool {
	return g.g.Directed()
}

func (g *WSubgraph) Weight(edgeFrom, edgeTo VIdx) interface{} {
	edgeFrom = g.vls[edgeFrom]
	edgeTo = g.vls[edgeTo]
	return g.Weight(edgeFrom, edgeTo)
}

func (g *WSubgraph) Clear(vertexNo VIdx) {
	panic("must not clear WSubgraph")
}

func (g *WSubgraph) SetWeight(edgeFrom, edgeTo VIdx, w interface{}) {
	edgeFrom = g.vls[edgeFrom]
	edgeTo = g.vls[edgeTo]
	g.SetWeight(edgeFrom, edgeTo, w)
}
