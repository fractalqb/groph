package shortestpath

import "git.fractalqb.de/fractalqb/groph"

func FloydWarshallf32(g groph.WGf32) {
	vno := g.Order()
	if u, ok := g.(groph.WUf32); ok {
		for k := groph.V0; k < vno; k++ {
			for i := groph.V0; i < vno; i++ {
				for j := 0; j < i; j++ {
					ds := u.Edge(i, k) + u.Edge(k, j)
					if u.EdgeU(i, j) > ds {
						u.SetEdgeU(i, j, ds)
					}
				}
			}
		}
	} else {
		for k := groph.V0; k < vno; k++ {
			for i := groph.V0; i < vno; i++ {
				for j := groph.V0; j < vno; j++ {
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
	vno := g.Order()
	if u, ok := g.(groph.WUi32); ok {
		for k := groph.V0; k < vno; k++ {
			for i := groph.V0; i < vno; i++ {
				for j := 0; j < i; j++ {
					ds, ok := u.Edge(i, k)
					if !ok {
						continue
					}
					if tmp, ok := u.Edge(k, j); ok {
						ds += tmp
					} else {
						continue
					}
					if d, ok := u.EdgeU(i, j); !ok || d > ds {
						u.SetEdgeU(i, j, ds)
					}
				}
			}
		}
	} else {
		for k := groph.V0; k < vno; k++ {
			for i := groph.V0; i < vno; i++ {
				for j := groph.V0; j < vno; j++ {
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
	vno := g.Order()
	for k := groph.V0; k < vno; k++ {
		for i := groph.V0; i < vno; i++ {
			for j := groph.V0; j < vno; j++ {
				ds := g.Edge(i, k) + g.Edge(k, j)
				if g.Edge(i, j) > ds {
					g.SetEdge(i, j, ds)
				}
			}
		}
	}
}

func FloydWarshallAdjMxDi32(g *groph.AdjMxDi32) {
	vno := g.Order()
	for k := groph.V0; k < vno; k++ {
		for i := groph.V0; i < vno; i++ {
			for j := groph.V0; j < vno; j++ {
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
