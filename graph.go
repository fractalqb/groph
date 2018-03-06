package graph

type Vertex = interface{}

type Measure func(from, to Vertex) interface{}

type Graph interface {
	VertexNo() uint
	Vertex(idx uint) Vertex
	Directed() bool
	Weight(fromIdx, toIdx uint) interface{}
	SetWeight(fromIdx, toIdx uint, w interface{})
}

type Gbool interface {
	Graph
	Edge(fromIdx, toIdx uint) (exists bool)
}

type Gint interface {
	Graph
	Edge(fromIdx, toIdx uint) (weight int)
}

type Gf32 interface {
	Graph
	Edge(fromIdx, toIdx uint) (weight float32)
}

func SetMetric(g Graph, d Measure) {
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

func CpWeights(dst, src Graph) {
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

func CpXWeights(dst, src Graph, xf func(in interface{}) interface{}) {
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
