package groph

// EachOutgoing calls onDest on each vertex d where the edge (from,d) is in
// graph g.
func EachOutgoing(g RGraph, from VIdx, onDest VisitVertex) {
	switch gi := g.(type) {
	case OutLister:
		gi.EachOutgoing(from, onDest)
	case RUndirected:
		eachUAdj(gi, from, onDest)
	default:
		eachDOut(g, from, onDest)
	}
}

func eachDOut(g RGraph, from VIdx, onDest VisitVertex) {
	vno := g.Order()
	for n := V0; n < vno; n++ {
		if g.Weight(from, n) != nil {
			onDest(n)
		}
	}
}

// EachIncoming calls onSource on each vertex s where the edge (s,to) is in
// graph g.
func EachIncoming(g RGraph, to VIdx, onSource VisitVertex) {
	switch gi := g.(type) {
	case InLister:
		gi.EachIncoming(to, onSource)
	case RUndirected:
		eachUAdj(gi, to, onSource)
	default:
		eachDIn(g, to, onSource)
	}
}

func eachDIn(g RGraph, to VIdx, onSource VisitVertex) {
	vno := g.Order()
	for n := V0; n < vno; n++ {
		if g.Weight(n, to) != nil {
			onSource(n)
		}
	}
}

// EachAdjacent calls onOther on each vertex o where at least one of the edges
// (this,o) and (o,this) is in graph g.
func EachAdjacent(g RGraph, this VIdx, onOther VisitVertex) {
	switch u := g.(type) {
	case RUndirected:
		switch ls := u.(type) {
		case OutLister:
			ls.EachOutgoing(this, onOther)
		case InLister:
			ls.EachIncoming(this, onOther)
		default:
			eachUAdj(u, this, onOther)
		}
	case OutLister:
		u.EachOutgoing(this, onOther)
		if il, ok := g.(InLister); ok {
			il.EachIncoming(this, onOther)
		} else {
			eachDIn(g, this, onOther)
		}
	case InLister:
		u.EachIncoming(this, onOther)
		if ol, ok := g.(OutLister); ok {
			ol.EachOutgoing(this, onOther)
		} else {
			eachDOut(g, this, onOther)
		}
	default:
		eachDAdj(g, this, onOther)
	}
}

func eachDAdj(g RGraph, v VIdx, on VisitVertex) {
	ord := g.Order()
	for u := V0; u < ord; u++ {
		if g.Weight(v, u) != nil || g.Weight(u, v) != nil {
			on(u)
		}
	}
}

func eachUAdj(u RUndirected, v VIdx, on VisitVertex) {
	vno := u.Order()
	n := V0
	for n < v {
		if u.WeightU(v, n) != nil {
			on(n)
		}
		n++
	}
	for n < vno {
		if u.WeightU(n, v) != nil {
			on(n)
		}
		n++
	}
}

// EachEdge call onEdge for every edge in graph g.
func EachEdge(g RGraph, onEdge VisitEdge) {
	// TODO optimize with In-/OutLister
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
