package groph

type RSubgraph struct {
	G    RGraph
	VMap []VIdx
}

func (g RSubgraph) Order() VIdx {
	return VIdx(len(g.VMap))
}

func (g RSubgraph) Weight(u, v VIdx) interface{} {
	return g.Weight(g.VMap[u], g.VMap[v])
}

type WSubgraph struct {
	G    WGraph
	VMap []VIdx
}

func (g WSubgraph) Order() VIdx {
	return VIdx(len(g.VMap))
}

func (g WSubgraph) Weight(u, v VIdx) interface{} {
	return g.Weight(g.VMap[u], g.VMap[v])
}

func (g WSubgraph) Reset(order VIdx) {
	panic("must not clear WSubgraph")
}

func (g WSubgraph) SetWeight(u, v VIdx, w interface{}) {
	g.SetWeight(g.VMap[u], g.VMap[v], w)
}

// TODO Must sort vls to handle WeightU and SetWeightU?
// type RSubUndir struct {
// 	g   RUndirected
// 	vls []VIdx
// }

// func (g RSubUndir) Order() VIdx {
// 	return VIdx(len(g.vls))
// }

// func (g RSubUndir) Weight(u, v VIdx) interface{} {
// 	return g.Weight(g.vls[u], g.vls[v])
// }

// func (g RSubUndir) WeightU(u, v VIdx) interface{} {
// 	return g.WeightU(g.vls[u], g.vls[v])
// }

// type WSubUndir struct {
// 	g   WGraph
// 	vls []VIdx
// }

// func (g WSubUndir) Order() VIdx {
// 	return VIdx(len(g.vls))
// }

// func (g WSubUndir) Weight(u, v VIdx) interface{} {
// 	return g.Weight(g.vls[u], g.vls[v])
// }

// func (g WSubUndir) WeightU(u, v VIdx) interface{} {
// 	return g.WeightU(g.vls[u], g.vls[v])
// }

// func (g WSubUndir) Reset(order VIdx) {
// 	panic("must not clear WSubgraph")
// }

// func (g WSubUndir) SetWeight(u, v VIdx, w interface{}) {
// 	g.SetWeight(g.vls[u], g.vls[v], w)
// }

// func (g WSubUndir) SetWeightU(u, v VIdx, w interface{}) {
// 	g.SetWeight(g.vls[u], g.vls[v], w)
// }
