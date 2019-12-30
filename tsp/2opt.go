package tsp

import "git.fractalqb.de/fractalqb/groph"

// d2optU computes the difference in weight sum for a specific 2-opt operation
// that swaps e0 / e1 for undirected graphs.
func diff2optU(g groph.RGf32, p []groph.VIdx, e0, e1 groph.VIdx) (wdiff float32) {
	lenp := groph.VIdx(len(p))
	wdiff = -g.Edge(p[e0], p[e0+1])
	wdiff += g.Edge(p[e0], p[e1])
	if e1+1 == lenp {
		wdiff -= g.Edge(p[e1], p[0])
		wdiff += g.Edge(p[e0+1], p[0])
	} else {
		wdiff -= g.Edge(p[e1], p[e1+1])
		wdiff += g.Edge(p[e0+1], p[e1+1])
	}
	return wdiff
}

// d2optD computes the difference in weight sum for a specific 2-opt operation
// that swaps e0 / e1 for directed graphs.
func diff2optD(g groph.RGf32, p []groph.VIdx, e0, e1 groph.VIdx) (wdiff float32) {
	wdiff = diff2optU(g, p, e0, e1)
	for i := e0 + 1; i < e1; i++ {
		wdiff -= g.Edge(p[i], p[i+1])
		wdiff += g.Edge(p[i+1], p[i])
	}
	return wdiff
}

func apply2opt(p []groph.VIdx, e0, e1 groph.VIdx) {
	e0++
	for e0 < e1 {
		p[e0], p[e1] = p[e1], p[e0]
		e0++
		e1--
	}
}

func TwoOptf32(g groph.RGf32) (path []groph.VIdx, plen float32) {
	var diff2opt func(groph.RGf32, []groph.VIdx, groph.VIdx, groph.VIdx) float32
	if groph.Directed(g) {
		diff2opt = diff2optD
	} else {
		diff2opt = diff2optU
	}
	vno := g.VertexNo()
	path = make([]groph.VIdx, vno)
	for i := groph.V0; i+1 < vno; i++ {
		path[i] = i
		plen += g.Edge(i, i+1)
	}
	path[vno-1] = vno - 1
	plen += g.Edge(vno-1, 0)
	for {
		be0, be1 := vno, vno
		bdiff := float32(0)
		for e0 := groph.V0; e0 < vno; e0++ {
			for e1 := e0 + 1; e1 < vno; e1++ {
				diff := diff2opt(g, path, e0, e1)
				if diff < bdiff {
					be0, be1 = e0, e1
					bdiff = diff
				}
			}
		}
		if bdiff < float32(0) {
			apply2opt(path, be0, be1)
			plen += bdiff
		} else {
			break
		}
	}
	return path, plen
}

// Benchmark showed that the proformance gain is not worth it
// func (g *AdjMxDf32) Tsp2Opt() (path []uint, plen float32) {
// 	vno := g.VertexNo()
// 	path = make([]uint, vno)
// 	for i := uint(0); i+1 < vno; i++ {
// 		path[i] = i
// 		plen += g.Edge(i, i+1)
// 	}
// 	path[vno-1] = vno - 1
// 	plen += g.Edge(vno-1, 0)
// 	for {
// 		be0, be1 := vno, vno
// 		bdiff := float32(0)
// 		for e0 := uint(0); e0 < vno; e0++ {
// 			for e1 := e0 + 1; e1 < vno; e1++ {
// 				diff := diff2optD(g, path, e0, e1)
// 				if diff < bdiff {
// 					be0, be1 = e0, e1
// 					bdiff = diff
// 				}
// 			}
// 		}
// 		if bdiff < float32(0) {
// 			apply2opt(path, be0, be1)
// 			plen += bdiff
// 		} else {
// 			break
// 		}
// 	}
// 	return path, plen
// }

// Benchmark showed that the proformance gain is not worth it
// func (g *AdjMxUf32) Tsp2Opt() (path []uint, plen float32) {
// 	vno := g.VertexNo()
// 	path = make([]uint, vno)
// 	for i := uint(0); i+1 < vno; i++ {
// 		path[i] = i
// 		plen += g.Edge(i, i+1)
// 	}
// 	path[vno-1] = vno - 1
// 	plen += g.Edge(vno-1, 0)
// 	for {
// 		be0, be1 := vno, vno
// 		bdiff := float32(0)
// 		for e0 := uint(0); e0 < vno; e0++ {
// 			for e1 := e0 + 1; e1 < vno; e1++ {
// 				diff := diff2optU(g, path, e0, e1)
// 				if diff < bdiff {
// 					be0, be1 = e0, e1
// 					bdiff = diff
// 				}
// 			}
// 		}
// 		if bdiff < float32(0) {
// 			apply2opt(path, be0, be1)
// 			plen += bdiff
// 		} else {
// 			break
// 		}
// 	}
// 	return path, plen
// }
