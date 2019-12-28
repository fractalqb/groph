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
	I, J VIdx
}

// RGraph represents a set of graph data that allows read only access to the
// egde weights.
//
// For graphs that can change edged see WGraph.
type RGraph interface {
	// VertextNo return the numer of vertices in the graph.
	VertexNo() VIdx
	// Directed returns true if the graph is a directed graph and false
	// otherwiese.
	Directed() bool
	// Returns the weight of the edge that connects the vertex with index
	// edgeFrom with the vertex with index edgeTo. If there is no such edge
	// it returns nil.
	Weight(edgeFrom, edgeTo VIdx) interface{}
}

func WeightOr(g RGraph, edgeFrom, edgeTo VIdx, or interface{}) interface{} {
	if res := g.Weight(edgeFrom, edgeTo); res != nil {
		return res
	}
	return or
}

type VisitNeighbour = func(neighbour VIdx)

// NeighbourLister is implemented by graph implementations that can easily
// iterate over all neighbous of one node.
type NeighbourLister interface {
	EachNeighbour(v VIdx, do VisitNeighbour)
}

// Guarantees to call (i,j) with i <= j on undirected graphs
func EachNeighbour(g RGraph, v VIdx, do VisitNeighbour) {
	if ln, ok := g.(NeighbourLister); ok {
		ln.EachNeighbour(v, do)
	} else if g.Directed() {
		vno := g.VertexNo()
		for n := VIdx(0); n < vno; n++ {
			if w := g.Weight(v, n); w != nil {
				do(n)
			}
		}
	} else {
		vno := g.VertexNo()
		n := VIdx(0)
		for n < v {
			if w := g.Weight(n, v); w != nil {
				do(n)
			}
			n++
		}
		for n < vno {
			if w := g.Weight(v, n); w != nil {
				do(n)
			}
			n++
		}
	}
}

func CheckDirected(g RGraph) bool {
	vno := g.VertexNo()
	for i := VIdx(0); i < vno; i++ {
		for j := i + 1; j < vno; i++ {
			w1 := g.Weight(i, j)
			w2 := g.Weight(j, i)
			if !reflect.DeepEqual(w1, w2) {
				return true
			}
		}
	}
	return false
}

// WGraph represents a set of graph data tha allows read and write access to
// the egde weights.
type WGraph interface {
	RGraph
	// Clear resizes the graph to vertexNo and reinitializes it. Implementations
	// are expected to reuse memory.
	Clear(vertexNo VIdx)
	// SetWeight sets the edge weight for the edge starting at vertex edgeFrom
	// and ending at vertex edgeTo. Passing w==nil removes the edge.
	SetWeight(edgeFrom, edgeTo VIdx, w interface{})
}

// Clear clears a WGraph while keeping the original vertexNo.
func Clear(g WGraph) { g.Clear(g.VertexNo()) }

// RGbool represents a RGraph with boolean edge weights.
type RGbool interface {
	RGraph
	Edge(edgeFrom, edgeTo VIdx) bool
}

// WGbool represents a WGraph with boolean edge weights.
type WGbool interface {
	WGraph
	Edge(edgeFrom, edgeTo VIdx) bool
	SetEdge(edgeFrom, edgeTo VIdx, flag bool)
}

// An RGi32 is a RGraph with type safe access to the edge weight of type
// int32. Besides type safety this avoids boxing/unboxing of the Weight
// method for performance reasons.
type RGi32 interface {
	RGraph
	Edge(edgeFrom, edgeTo VIdx) (weight int32)
}

// An WGi32 is to WGraph what RGi32 is to RGraph.
type WGi32 interface {
	WGraph
	Edge(edgeFrom, edgeTo VIdx) (weight int32, exists bool)
	SetEdge(edgeFrom, edgeTo VIdx, weight int32)
	DelEdge(edgeFrom, edgeTo VIdx)
}

// An RGf32 is a RGraph with type safe access to the edge weight of type
// float32. Besides type safety this avoids boxing/unboxing of the Weight
// method for performance reasons.
type RGf32 interface {
	RGraph
	Edge(edgeFrom, edgeTo VIdx) (weight float32)
}

// An WGf32 is to WGraph what RGf32 is to RGraph.
type WGf32 interface {
	WGraph
	Edge(edgeFrom, edgeTo VIdx) (weight float32)
	SetEdge(edgeFrom, edgeTo VIdx, weight float32)
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
	if dst.Directed() {
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
	if dst.Directed() {
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
