package groph

func Set[W any](g WGraph[W], w W, vs ...VIdx) {
	for i := 0; i < len(vs); i += 2 {
		j := i + 1
		if j >= len(vs) {
			return
		}
		g.SetEdge(vs[i], vs[j], w)
	}
}

func Del[W any](g WGraph[W], vs ...VIdx) {
	for i := 0; i < len(vs); i++ {
		j := i + 1
		if j >= len(vs) {
			return
		}
		g.DelEdge(vs[i], vs[j])
	}
}
