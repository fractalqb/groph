package groph

import (
	"reflect"
)

// VIdx is the type used to represent vertices in the graph implementations
// provided by the groph package.
type VIdx = int

// Edge represents a graphs edge between vertices I and J. For directed graphs
// its the edge from I to J.
type Edge struct {
	U, V VIdx
}

// RGraph represents a set of graph data that allows read only access to the
// egde weights.
//
// For graphs that can change edged see WGraph.
type RGraph interface {
	// VertextNo return the numer of vertices in the graph.
	VertexNo() VIdx
	// Returns the weight of the edge that connects the vertex with index
	// u with the vertex with index v. If there is no such edge it returns nil.
	Weight(u, v VIdx) interface{}
}

func WeightOr(g RGraph, u, v VIdx, or interface{}) interface{} {
	if res := g.Weight(u, v); res != nil {
		return res
	}
	return or
}

type RUndirected interface {
	RGraph
	WeightU(u, v VIdx) interface{}
}

func Directed(g RGraph) bool {
	_, ok := g.(RUndirected)
	return !ok
}

type VisitNeighbour = func(neighbour VIdx)

// NeighbourLister is implemented by graph implementations that can easily
// iterate over all neighbous of one node.
type NeighbourLister interface {
	EachNeighbour(v VIdx, do VisitNeighbour)
}

// Guarantees to call (i,j) with i >= j on undirected graphs
func EachNeighbour(g RGraph, v VIdx, do VisitNeighbour) {
	switch gi := g.(type) {
	case NeighbourLister:
		gi.EachNeighbour(v, do)
	case RUndirected:
		vno := gi.VertexNo()
		n := VIdx(0)
		for n < v {
			if w := gi.WeightU(v, n); w != nil {
				do(n)
			}
			n++
		}
		for n < vno {
			if w := gi.WeightU(n, v); w != nil {
				do(n)
			}
			n++
		}
	default:
		vno := g.VertexNo()
		for n := VIdx(0); n < vno; n++ {
			if w := g.Weight(v, n); w != nil {
				do(n)
			}
		}
	}
}

// WGraph represents a set of graph data tha allows read and write access to
// the egde weights.
type WGraph interface {
	RGraph
	// Clear resizes the graph to vertexNo and reinitializes it. Implementations
	// are expected to reuse memory.
	Reset(vertexNo VIdx)
	// SetWeight sets the edge weight for the edge starting at vertex u and
	// ending at vertex v. Passing w==nil removes the edge.
	SetWeight(u, v VIdx, w interface{})
}

type WUndirected interface {
	WGraph
	WeightU(u, v VIdx) interface{}
	SetWeightU(u, v VIdx, w interface{})
}

// Clear clears a WGraph while keeping the original vertexNo.
func Reset(g WGraph) { g.Reset(g.VertexNo()) }

// RGbool represents a RGraph with boolean edge weights.
type RGbool interface {
	RGraph
	Edge(u, v VIdx) bool
}

type RUbool interface {
	RGbool
	EdgeU(u, v VIdx) bool
}

// WGbool represents a WGraph with boolean edge weights.
type WGbool interface {
	WGraph
	Edge(u, v VIdx) bool
	SetEdge(u, v VIdx, flag bool)
}

type WUbool interface {
	WGbool
	EdgeU(u, v VIdx) bool
	SetEdgeU(u, v VIdx, flag bool)
}

// An RGi32 is a RGraph with type safe access to the edge weight of type
// int32. Besides type safety this avoids boxing/unboxing of the Weight
// method for performance reasons.
type RGi32 interface {
	RGraph
	Edge(u, v VIdx) (weight int32, ok bool)
}

type RUi32 interface {
	RGi32
	EdgeU(u, v VIdx) (weight int32, ok bool)
}

// An WGi32 is to WGraph what RGi32 is to RGraph.
type WGi32 interface {
	WGraph
	Edge(u, v VIdx) (weight int32, ok bool)
	SetEdge(u, v VIdx, weight int32)
	DelEdge(u, v VIdx)
}

type WUi32 interface {
	WGi32
	EdgeU(u, v VIdx) (weight int32, ok bool)
	SetEdgeU(u, v VIdx, weight int32)
}

// An RGf32 is a RGraph with type safe access to the edge weight of type
// float32. Besides type safety this avoids boxing/unboxing of the Weight
// method for performance reasons.
type RGf32 interface {
	RGraph
	Edge(u, v VIdx) (weight float32)
}

type RUf32 interface {
	RGf32
	EdgeU(u, v VIdx) (weight float32)
}

// An WGf32 is to WGraph what RGf32 is to RGraph.
type WGf32 interface {
	WGraph
	Edge(u, v VIdx) (weight float32)
	SetEdge(u, v VIdx, weight float32)
}

type WUf32 interface {
	WGf32
	EdgeU(u, v VIdx) (weight float32)
	SetEdgeU(u, v VIdx, weight float32)
}

// CpWeights copies the edge weights from one grap to another.
// Vertices are identified by their index, i.e. the user has to take care of
// the vertex order. If the number of vertices in the graph differs the smaller
// graph determines how many edge weights are copied.
func CpWeights(dst WGraph, src RGraph) WGraph {
	sz := dst.VertexNo()
	if src.VertexNo() < sz {
		sz = src.VertexNo()
	}
	if Directed(dst) {
		for i := VIdx(0); i < sz; i++ {
			for j := VIdx(0); j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, w)
			}
		}
	} else {
		for i := VIdx(0); i < sz; i++ {
			for j := i; j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, w)
			}
		}
	}
	return dst
}

// CpXWeights “transfers” the edge weights from src Graph to dst Graph
// with the same vertex restirctions as CpWeights. CpXWeights applies
// the transformation function xf() to each edge weight.
func CpXWeights(dst WGraph, src RGraph, xf func(in interface{}) interface{}) WGraph {
	sz := dst.VertexNo()
	if src.VertexNo() < sz {
		sz = src.VertexNo()
	}
	if Directed(dst) {
		for i := VIdx(0); i < sz; i++ {
			for j := VIdx(0); j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, xf(w))
			}
		}
	} else {
		for i := VIdx(0); i < sz; i++ {
			for j := i; j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, xf(w))
			}
		}
	}
	return dst
}

// TODO can this be done in place?
func ReorderPath(slice interface{}, path []VIdx) {
	slv := reflect.ValueOf(slice)
	if slv.Len() == 0 {
		return
	}
	tmp := make([]interface{}, slv.Len())
	for i := 0; i < slv.Len(); i++ {
		tmp[i] = slv.Index(i).Interface()
	}
	for w, r := range path {
		v := tmp[r]
		slv.Index(w).Set(reflect.ValueOf(v))
	}
}

const i32cleared = -2147483648 // min{ int32 }
