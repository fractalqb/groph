package groph

func FloydWarshallf32(g WGf32) {
	// TODO optimize for undirected graphs
	vno := g.VertexNo()
	for k := uint(0); k < vno; k++ {
		for i := uint(0); i < vno; i++ {
			for j := uint(0); j < vno; j++ {
				ds := g.Edge(i, k) + g.Edge(k, j)
				if g.Edge(i, j) > ds {
					g.SetEdge(i, j, ds)
				}
			}
		}
	}
}

func (g *AdjMxDf32) FloydWarshall() {
	vno := g.VertexNo()
	for k := uint(0); k < vno; k++ {
		for i := uint(0); i < vno; i++ {
			for j := uint(0); j < vno; j++ {
				ds := g.Edge(i, k) + g.Edge(k, j)
				if g.Edge(i, j) > ds {
					g.SetEdge(i, j, ds)
				}
			}
		}
	}
}
