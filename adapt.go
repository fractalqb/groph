package groph

import (
	"fmt"
	"reflect"
)

type Vertex = interface{}

// SliceNMeasure implements a metric RGraph based on a slice of vertices of
// some type V and a function f(u V, v V) → W that compute the weight of type W
// for any edge (u, v).
//
// E.g. use CpWeights or CpXWeights to initialize an other WGraph.
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
func (g *SliceNMeasure) Verify() *SliceNMeasure {
	if _, err := g.Check(); err != nil {
		panic(err)
	}
	return g
}

func (g *SliceNMeasure) VertexNo() VIdx {
	return VIdx(g.slc.Len())
}

func (g *SliceNMeasure) Vertex(idx VIdx) Vertex {
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
