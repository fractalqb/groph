package groph

import (
	"math"
	"reflect"
)

// VIdx is the type used to represent vertices in the graph implementations
// provided by the groph package.
type VIdx = int

const V0 VIdx = 0

// Edge represents a graphs edge between vertices U and V. For directed graphs
// its the edge from U to V.
type Edge struct {
	U, V VIdx
}

// RGraph represents graph that allows read only access to the egde
// weights.
//
// For graphs that can change be modified see WGraph. For undirected
// graphs see also RUndirected.
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

// RUndirected represents an undirected graph that allows read only
// access to the edge weights.
type RUndirected interface {
	RGraph
	// Weight must only be called when u ≥ v.  Otherwise WeightU's
	// behaviour is unspecified, it even might crash.  In many
	// implementations this can be way more efficient than the
	// general case, see method Weight().
	WeightU(u, v VIdx) interface{}
}

// Directed returns true, iff g is a directed graph and false otherwise.
func Directed(g RGraph) bool {
	_, ok := g.(RUndirected)
	return !ok
}

type VisitVertex = func(neighbour VIdx)

// OutLister is implemented by graph implementations that can easily iterate
// over all outgoing edges of one node.
//
// See also traverse.EachOutgoing function.
type OutLister interface {
	EachOutgoing(from VIdx, onDest VisitVertex)
	OutDegree(v VIdx) int
}

// InLister is implemented by graph implementations that can easily iterate
// over all incoming edges of one node.
//
// See also EachIncoming function.
type InLister interface {
	EachIncoming(to VIdx, onSource VisitVertex)
	InDegree(v VIdx) int
}

// WGraph represents graph that allows read and write access to the
// egde weights.
//
// For undirected graphs see also WUndirected.
type WGraph interface {
	RGraph
	// Reset resizes the graph to vertexNo and reinitializes it. Implementations
	// are expected to reuse memory.
	Reset(vertexNo VIdx)
	// SetWeight sets the edge weight for the edge starting at vertex u and
	// ending at vertex v. Passing w==nil removes the edge.
	SetWeight(u, v VIdx, w interface{})
}

// WUndirected represents an undirected graph that allows read and
// write access to the egde weights.
type WUndirected interface {
	WGraph
	// See RUndirected.WeightU
	WeightU(u, v VIdx) interface{}
	// SetWeightU must only be called when u ≥ v.  Otherwise
	// SetWeightU's behaviour is unspecified, it even might crash.
	//
	// See also RUndirected.WeightU
	SetWeightU(u, v VIdx, w interface{})
}

// Reset clears a WGraph while keeping the original vertexNo.
func Reset(g WGraph) { g.Reset(g.VertexNo()) }

// RGbool represents a RGraph with boolean edge weights.
type RGbool interface {
	RGraph
	// Edge returns true, iff the edge (u,v) is in the graph.
	Edge(u, v VIdx) bool
}

// RUbool represents a RUndirected with boolean edge weights.
type RUbool interface {
	RGbool
	EdgeU(u, v VIdx) bool
}

// WGbool represents a WGraph with boolean edge weights.
type WGbool interface {
	WGraph
	// see RGbool
	Edge(u, v VIdx) bool
	// SetEdge removes the edge (u,v) from the graph when flag == bool.
	// Otherwise it adds the edge (u,v) to the graph.
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
	// Edge returns ok == true, iff the edge (u,v) is in the graph. Then it will
	// also return the weight of the edge. Otherwise the value of weight is
	// unspecified.
	Edge(u, v VIdx) (weight int32, ok bool)
}

type RUi32 interface {
	RGi32
	EdgeU(u, v VIdx) (weight int32, ok bool)
}

// An WGi32 is to WGraph what RGi32 is to RGraph.
type WGi32 interface {
	WGraph
	// see RGi32
	Edge(u, v VIdx) (weight int32, ok bool)
	// SetEdge sets the weight of the edge (u,v). If the edge (u,v) was not in
	// the graph before, it is implicitly added.
	SetEdge(u, v VIdx, weight int32)
	// DelEdge deletes the edge (u,v) from the graph.
	DelEdge(u, v VIdx)
}

type WUi32 interface {
	WGi32
	EdgeU(u, v VIdx) (weight int32, ok bool)
	SetEdgeU(u, v VIdx, weight int32)
}

var nan32 = float32(math.NaN())

func NaN32() float32 { return nan32 }

func IsNaN32(x float32) bool { return math.IsNaN(float64(x)) }

// An RGf32 is a RGraph with type safe access to the edge weight of type
// float32. Besides type safety this avoids boxing/unboxing of the Weight
// method for performance reasons.
type RGf32 interface {
	RGraph
	// Edge returns Nan32() when the edge (u,v) is not in the graph. Otherwise
	// it returns the weight of the edge.
	Edge(u, v VIdx) (weight float32)
}

type RUf32 interface {
	RGf32
	EdgeU(u, v VIdx) (weight float32)
}

// An WGf32 is to WGraph what RGf32 is to RGraph.
type WGf32 interface {
	WGraph
	// see RGf32
	Edge(u, v VIdx) (weight float32)
	// SetEdge removes the edge (u,v) from the graph, iff weight is NaN32().
	// Othwerwise it sets the weight of the edge to weight.
	SetEdge(u, v VIdx, weight float32)
}

type WUf32 interface {
	WGf32
	EdgeU(u, v VIdx) (weight float32)
	SetEdgeU(u, v VIdx, weight float32)
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
