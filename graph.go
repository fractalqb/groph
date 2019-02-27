package groph

import (
	"reflect"
)

type Vertex = interface{}

type Edge struct {
	I, J uint
}

// An RGraph represents a set of graph data that allows read only access to the
// egde's weight data.
type RGraph interface {
	// VertextNo return the numer of vertices in the graph.
	VertexNo() uint
	// Directed returns true if the graph is a directed graph and false
	// otherwiese.
	Directed() bool
	// Returns the weight of the edge that connects the vertex with index
	// fromIdx with the vertex with index toIdx.
	Weight(fromIdx, toIdx uint) interface{}
}

func WeightOr(g RGraph, fromIdx, toIdx uint, or interface{}) interface{} {
	if res := g.Weight(fromIdx, toIdx); res != nil {
		return res
	}
	return or
}

type ListNeightbours interface {
	EachNeighbour(v uint, do func(a, b uint, w interface{}))
}

func EachNeighbour(g RGraph, v uint, do func(a, b uint, w interface{})) {
	if ln, ok := g.(ListNeightbours); ok {
		ln.EachNeighbour(v, do)
	} else if g.Directed() {
		for n := uint(0); n < g.VertexNo(); n++ {
			if w := g.Weight(v, n); w != nil {
				do(v, n, w)
			}
			if w := g.Weight(n, v); w != nil {
				do(n, v, w)
			}
		}
	} else {
		for n := uint(0); n < g.VertexNo(); n++ {
			if w := g.Weight(v, n); w != nil {
				do(v, n, w)
			}
		}
	}
}

func CheckDirected(g RGraph) bool {
	vno := g.VertexNo()
	for i := uint(0); i < vno; i++ {
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

// An RGraph represents a set of graph data tha allows read and write access
// to the egde's weight data.
type WGraph interface {
	RGraph
	Clear(vertexNo uint)
	SetWeight(fromIdx, toIdx uint, w interface{})
}

func Clear(g WGraph) { g.Clear(g.VertexNo()) }

type RGbool interface {
	RGraph
	Edge(fromIdx, toIdx uint) bool
}

type WGbool interface {
	WGraph
	Edge(fromIdx, toIdx uint) bool
	SetEdge(fromIdx, toIdx uint, flag bool)
}

// An RGi32 is a RGraph with type safe access to the edge weight of type
// int32. Besides type safety this avoids boxing/unboxing of the Weight
// method for performance reasons.
type RGi32 interface {
	RGraph
	Edge(fromIdx, toIdx uint) (weight int32)
}

// An WGi32 is to WGraph what RGi32 is to RGraph.
type WGi32 interface {
	WGraph
	Edge(fromIdx, toIdx uint) (weight int32, exists bool)
	SetEdge(fromIdx, toIdx uint, weight int32)
	DelEdge(fromIdx, toIdx uint)
}

// An RGf32 is a RGraph with type safe access to the edge weight of type
// float32. Besides type safety this avoids boxing/unboxing of the Weight
// method for performance reasons.
type RGf32 interface {
	RGraph
	Edge(fromIdx, toIdx uint) (weight float32)
}

// An WGf32 is to WGraph what RGf32 is to RGraph.
type WGf32 interface {
	WGraph
	Edge(fromIdx, toIdx uint) (weight float32)
	SetEdge(fromIdx, toIdx uint, weight float32)
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
		for i := uint(0); i < sz; i++ {
			for j := uint(0); j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, w)
			}
		}
	} else {
		for i := uint(0); i < sz; i++ {
			for j := i; j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, w)
			}
		}
	}
	return dst
}

// CpXWeights “transfers” the edge weights from src Graph to dst Graph
// with the same vertex restirctions as CpWeights. CpXWeights allpies
// the transformation function xf() to each edge weight.
func CpXWeights(dst WGraph, src RGraph, xf func(in interface{}) interface{}) WGraph {
	sz := dst.VertexNo()
	if src.VertexNo() < sz {
		sz = src.VertexNo()
	}
	if dst.Directed() {
		for i := uint(0); i < sz; i++ {
			for j := uint(0); j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, xf(w))
			}
		}
	} else {
		for i := uint(0); i < sz; i++ {
			for j := i; j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, xf(w))
			}
		}
	}
	return dst
}

// TODO can this be done in place?
func ReorderPath(slice interface{}, path []uint) {
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
