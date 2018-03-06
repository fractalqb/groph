package graph

type Vertex = interface{}

type Measure func(from, to Vertex) interface{}

type RGraph interface {
	VertexNo() uint
	Vertex(idx uint) Vertex
	Directed() bool
	Weight(fromIdx, toIdx uint) interface{}
}

type WGraph interface {
	RGraph
	SetWeight(fromIdx, toIdx uint, w interface{})
}

// type Gbool extends a Graph with a specific access method to edge
// weight. This shall avoid boxing/unboxing of the Graph Weight method
// for performance reasons.
type RGbool interface {
	RGraph
	Edge(fromIdx, toIdx uint) (exists bool)
}

// type Gint extends a Graph with a specific access method to edge
// weight. This shall avoid boxing/unboxing of the Graph Weight method
// for performance reasons.
type RGint interface {
	RGraph
	Edge(fromIdx, toIdx uint) (weight int)
}

// type Gbool extends a Graph with a specific access method to edge
// weight. This shall avoid boxing/unboxing of the Graph Weight method
// for performance reasons.
type RGf32 interface {
	RGraph
	Edge(fromIdx, toIdx uint) (weight float32)
}

func SetMetric(g WGraph, d Measure) {
	if g.Directed() {
		vno := g.VertexNo()
		for i := uint(0); i < vno; i++ {
			v1 := g.Vertex(i)
			for j := uint(0); j < vno; j++ {
				v2 := g.Vertex(j)
				g.SetWeight(i, j, d(v1, v2))
			}
		}
	} else {
		vno := g.VertexNo()
		for i := uint(0); i < vno; i++ {
			v1 := g.Vertex(i)
			for j := i; j < vno; j++ {
				v2 := g.Vertex(j)
				g.SetWeight(i, j, d(v1, v2))
			}
		}
	}
}

// CpWeights copies the edge weights from one grap to
// another. Vertices are identified by their index, i.e. the user has
// to take care of the vertex order. If the number of vertices in the
// graph differs the smaller graph determines how many edge weights
// are copied.
func CpWeights(dst WGraph, src RGraph) {
	sz := dst.VertexNo()
	if src.VertexNo() < sz {
		sz = src.VertexNo()
	}
	if dst.Directed() || src.Directed() {
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
}

// CpXWeights “transfers” the edge weights from src Graph to dst Graph
// with the same vertex restirctions as CpWeights. CpXWeights allpies
// the transformation function xf() to each edge weight.
func CpXWeights(dst WGraph, src RGraph, xf func(in interface{}) interface{}) {
	sz := dst.VertexNo()
	if src.VertexNo() < sz {
		sz = src.VertexNo()
	}
	if dst.Directed() || src.Directed() {
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
}
