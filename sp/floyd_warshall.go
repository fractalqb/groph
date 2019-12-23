package sp

import "git.fractalqb.de/fractalqb/groph"

func FloydWarshallf32(g groph.WGf32) {
	vno := g.VertexNo()
	if g.Directed() {
		for k := groph.VIdx(0); k < vno; k++ {
			for i := groph.VIdx(0); i < vno; i++ {
				for j := groph.VIdx(0); j < vno; j++ {
					ds := g.Edge(i, k) + g.Edge(k, j)
					if g.Edge(i, j) > ds {
						g.SetEdge(i, j, ds)
					}
				}
			}
		}
	} else {
		for k := groph.VIdx(0); k < vno; k++ {
			for i := groph.VIdx(0); i+1 <= vno; i++ {
				for j := i + 1; j < vno; j++ {
					ds := g.Edge(i, k) + g.Edge(k, j)
					if g.Edge(i, j) > ds {
						g.SetEdge(i, j, ds)
					}
				}
			}
		}
	}
}

func FloydWarshalli32(g groph.WGi32) {
	vno := g.VertexNo()
	if g.Directed() {
		for k := groph.VIdx(0); k < vno; k++ {
			for i := groph.VIdx(0); i < vno; i++ {
				for j := groph.VIdx(0); j < vno; j++ {
					ds, ok := g.Edge(i, k)
					if !ok {
						continue
					}
					if tmp, ok := g.Edge(k, j); ok {
						ds += tmp
					} else {
						continue
					}
					if d, ok := g.Edge(i, j); !ok || d > ds {
						g.SetEdge(i, j, ds)
					}
				}
			}
		}
	} else {
		for k := groph.VIdx(0); k < vno; k++ {
			for i := groph.VIdx(0); i+1 <= vno; i++ {
				for j := i + 1; j < vno; j++ {
					ds, ok := g.Edge(i, k)
					if !ok {
						continue
					}
					if tmp, ok := g.Edge(k, j); ok {
						ds += tmp
					} else {
						continue
					}
					if d, ok := g.Edge(i, j); !ok || d > ds {
						g.SetEdge(i, j, ds)
					}
				}
			}
		}
	}
}

func FloydWarshallAdjMxDf32(g *groph.AdjMxDf32) {
	vno := g.VertexNo()
	for k := groph.VIdx(0); k < vno; k++ {
		for i := groph.VIdx(0); i < vno; i++ {
			for j := groph.VIdx(0); j < vno; j++ {
				ds := g.Edge(i, k) + g.Edge(k, j)
				if g.Edge(i, j) > ds {
					g.SetEdge(i, j, ds)
				}
			}
		}
	}
}

func FloydWarshallAdjMxDi32(g *groph.AdjMxDi32) {
	vno := g.VertexNo()
	for k := groph.VIdx(0); k < vno; k++ {
		for i := groph.VIdx(0); i < vno; i++ {
			for j := groph.VIdx(0); j < vno; j++ {
				ds, ok := g.Edge(i, k)
				if !ok {
					continue
				}
				if tmp, ok := g.Edge(k, j); ok {
					ds += tmp
				} else {
					continue
				}
				if d, ok := g.Edge(i, j); !ok || d > ds {
					g.SetEdge(i, j, ds)
				}
			}
		}
	}
}
