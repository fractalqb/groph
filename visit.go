package groph

// EachOutgoing calls onDest on each vertex d where the edge (from,d) is in
// graph g.
//
// For undirected graphs EachOutgoing calls WeightU with v ≥ u to detect
// neighbours.
func EachOutgoing(g RGraph, from VIdx, onDest VisitVertex) {
	switch gi := g.(type) {
	case OutLister:
		gi.EachOutgoing(from, onDest)
	case RUndirected:
		vno := gi.Order()
		n := V0
		for n < from {
			if w := gi.WeightU(from, n); w != nil {
				onDest(n)
			}
			n++
		}
		for n < vno {
			if w := gi.WeightU(n, from); w != nil {
				onDest(n)
			}
			n++
		}
	default:
		vno := g.Order()
		for n := V0; n < vno; n++ {
			if w := g.Weight(from, n); w != nil {
				onDest(n)
			}
		}
	}
}

// EachIncoming calls onSource on each vertex s where the edge (s,to) is in
// graph g.
//
// For undirected graphs EachIncoming calls WeightU with v ≥ u to detect
// neighbours.
func EachIncoming(g RGraph, to VIdx, onSource VisitVertex) {
	switch gi := g.(type) {
	case InLister:
		gi.EachIncoming(to, onSource)
	case RUndirected:
		vno := gi.Order()
		n := V0
		for n < to {
			if w := gi.WeightU(to, n); w != nil {
				onSource(n)
			}
			n++
		}
		for n < vno {
			if w := gi.WeightU(n, to); w != nil {
				onSource(n)
			}
			n++
		}
	default:
		vno := g.Order()
		for n := V0; n < vno; n++ {
			if w := g.Weight(n, to); w != nil {
				onSource(n)
			}
		}
	}
}

func EachEdge(g RGraph, onEdge VisitEdge) {
	switch ge := g.(type) {
	case EdgeLister:
		ge.EachEdge(onEdge)
	case RUndirected:
		vno := ge.Order()
		for i := V0; i < vno; i++ {
			for j := V0; j <= i; j++ {
				if w := ge.WeightU(i, j); w != nil {
					onEdge(i, j)
				}
			}
		}
	default:
		vno := g.Order()
		for i := V0; i < vno; i++ {
			for j := V0; j < vno; j++ {
				if w := g.Weight(i, j); w != nil {
					onEdge(i, j)
				}
			}
		}
	}
}
