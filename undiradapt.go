package groph

type RUadapt struct {
	G RGraph
}

func AsRUndir(g RGraph) RUndirected {
	if u, ok := g.(RUndirected); ok {
		return u
	}
	return RUadapt{g}
}

func (g RUadapt) Order() VIdx { return g.G.Order() }

func (g RUadapt) Weight(u, v VIdx) interface{} {
	if u > v {
		return g.G.Weight(u, v)
	}
	return g.G.Weight(v, u)
}

func (g RUadapt) WeightU(u, v VIdx) interface{} { return g.G.Weight(u, v) }

type WUadapt struct {
	G WGraph
}

func AsWUndir(g WGraph) WUndirected {
	if u, ok := g.(WUndirected); ok {
		return u
	}
	return WUadapt{g}
}

func (g WUadapt) Order() VIdx { return g.G.Order() }

func (g WUadapt) Weight(u, v VIdx) interface{} {
	if u > v {
		return g.G.Weight(u, v)
	}
	return g.G.Weight(v, u)
}

func (g WUadapt) WeightU(u, v VIdx) interface{} { return g.G.Weight(u, v) }

func (g WUadapt) Reset(order VIdx) {
	g.G.Reset(order)
}

func (g WUadapt) SetWeight(u, v VIdx, w interface{}) {
	if u > v {
		g.G.SetWeight(u, v, w)
	}
	g.G.SetWeight(v, u, w)
}

func (g WUadapt) SetWeightU(u, v VIdx, w interface{}) { g.G.SetWeight(u, v, w) }
