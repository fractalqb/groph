package groph

import "math"

// VIdx is the type used to represent vertices in the graph implementations
// provided by the groph package.
type VIdx = int

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
	// Order return the numer of vertices in the graph.
	Order() int
	// Returns the weight of the edge that connects the vertex with index
	// u with the vertex with index v. If there is no such edge it returns nil.
	Weight(u, v VIdx) interface{}
}

// RUndirected represents an undirected graph that allows read only
// access to the edge weights. For undirected graphs each edge (u,v) is
// considered outgiong as well as incoming for both, vertex u and vertext v.
type RUndirected interface {
	RGraph
	// Weight must only be called when u ≥ v.  Otherwise WeightU's
	// behaviour is unspecified, it even might crash.  In many
	// implementations this can be way more efficient than the
	// general case, see method Weight().
	WeightU(u, v VIdx) interface{}
}

// WGraph represents graph that allows read and write access to the
// egde weights.
//
// For undirected graphs see also WUndirected.
type WGraph interface {
	RGraph
	// Reset resizes the graph to order and reinitializes it. Implementations
	// are expected to reuse memory.
	Reset(order int)
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
	WeightU(u, v VIdx) interface{}
	SetWeightU(u, v VIdx, w interface{})
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
}

type WUi32 interface {
	WGi32
	WeightU(u, v VIdx) interface{}
	SetWeightU(u, v VIdx, w interface{})
	EdgeU(u, v VIdx) (weight int32, ok bool)
	SetEdgeU(u, v VIdx, weight int32)
}

var nan32 = float32(math.NaN())

// NaN32 is used by adjacency matrices with edge weight type 'float32'
// to mark edges that do not exist.
//
// See AdjMxDf32 and AdjMxUf32
func NaN32() float32 { return nan32 }

// IsNaN32 test is x is NaN (no a number). See also NaN32.
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
	WeightU(u, v VIdx) interface{}
	SetWeightU(u, v VIdx, w interface{})
	EdgeU(u, v VIdx) (weight float32)
	SetEdgeU(u, v VIdx, weight float32)
}
