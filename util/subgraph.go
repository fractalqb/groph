package util

import "git.fractalqb.de/fractalqb/groph"

type RSubgraph struct {
	G    groph.RGraph
	VMap []groph.VIdx
}

func (g RSubgraph) Order() groph.VIdx {
	return groph.VIdx(len(g.VMap))
}

func (g RSubgraph) Weight(u, v groph.VIdx) interface{} {
	return g.Weight(g.VMap[u], g.VMap[v])
}

type WSubgraph struct {
	G    groph.WGraph
	VMap []groph.VIdx
}

func (g WSubgraph) Order() groph.VIdx {
	return groph.VIdx(len(g.VMap))
}

func (g WSubgraph) Weight(u, v groph.VIdx) interface{} {
	return g.Weight(g.VMap[u], g.VMap[v])
}

func (g WSubgraph) Reset(order groph.VIdx) {
	panic("must not clear WSubgraph")
}

func (g WSubgraph) SetWeight(u, v groph.VIdx, w interface{}) {
	g.SetWeight(g.VMap[u], g.VMap[v], w)
}

// TODO Must sort vls to handle WeightU and SetWeightU?
// type RSubUndir struct {
// 	g   RUndirected
// 	vls []groph.VIdx
// }

// func (g RSubUndir) Order() groph.VIdx {
// 	return groph.VIdx(len(g.vls))
// }

// func (g RSubUndir) Weight(u, v groph.VIdx) interface{} {
// 	return g.Weight(g.vls[u], g.vls[v])
// }

// func (g RSubUndir) WeightU(u, v groph.VIdx) interface{} {
// 	return g.WeightU(g.vls[u], g.vls[v])
// }

// type WSubUndir struct {
// 	g   WGraph
// 	vls []groph.VIdx
// }

// func (g WSubUndir) Order() groph.VIdx {
// 	return groph.VIdx(len(g.vls))
// }

// func (g WSubUndir) Weight(u, v groph.VIdx) interface{} {
// 	return g.Weight(g.vls[u], g.vls[v])
// }

// func (g WSubUndir) WeightU(u, v groph.VIdx) interface{} {
// 	return g.WeightU(g.vls[u], g.vls[v])
// }

// func (g WSubUndir) Reset(order groph.VIdx) {
// 	panic("must not clear WSubgraph")
// }

// func (g WSubUndir) SetWeight(u, v groph.VIdx, w interface{}) {
// 	g.SetWeight(g.vls[u], g.vls[v], w)
// }

// func (g WSubUndir) SetWeightU(u, v groph.VIdx, w interface{}) {
// 	g.SetWeight(g.vls[u], g.vls[v], w)
// }
