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
	Edge(fromIdx, toIdx uint) (exists bool, weight int)
}

type Gf32 interface {
	Graph
	Edge(fromIdx, toIdx uint) (exists bool, weight float32)
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
