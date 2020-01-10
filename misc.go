package groph

// Directed returns true, iff g is a directed graph and false otherwise.
func Directed(g RGraph) bool {
	_, ok := g.(RUndirected)
	return !ok
}

// Set sets the weight of all passed edges to w.
func Set(g WGraph, w interface{}, edges ...Edge) {
	for _, e := range edges {
		g.SetWeight(e.U, e.V, w)
	}
}

// Reset clears a WGraph while keeping the original order. This is the same as
// calling g.Reset(g.Order()).
func Reset(g WGraph) { g.Reset(g.Order()) }
