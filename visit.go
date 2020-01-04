package groph

// EachOutgoing calls onDest on each vertex d where the edge (from,d) is in
// graph g.
func EachOutgoing(g RGraph, from VIdx, onDest VisitVertex) (stopped bool) {
	switch gi := g.(type) {
	case OutLister:
		return gi.EachOutgoing(from, onDest)
	case RUndirected:
		return eachUAdj(gi, from, onDest)
	default:
		return eachDOut(g, from, onDest)
	}
}

func eachDOut(g RGraph, from VIdx, onDest VisitVertex) bool {
	vno := g.Order()
	for n := V0; n < vno; n++ {
		if g.Weight(from, n) != nil {
			if onDest(n) {
				return true
			}
		}
	}
	return false
}

// EachIncoming calls onSource on each vertex s where the edge (s,to) is in
// graph g.
func EachIncoming(g RGraph, to VIdx, onSource VisitVertex) (stopped bool) {
	switch gi := g.(type) {
	case InLister:
		return gi.EachIncoming(to, onSource)
	case RUndirected:
		return eachUAdj(gi, to, onSource)
	default:
		return eachDIn(g, to, onSource)
	}
}

func eachDIn(g RGraph, to VIdx, onSource VisitVertex) bool {
	vno := g.Order()
	for n := V0; n < vno; n++ {
		if g.Weight(n, to) != nil {
			if onSource(n) {
				return true
			}
		}
	}
	return false
}

// EachAdjacent calls onOther on each vertex o where at least one of the edges
// (this,o) and (o,this) is in graph g.
func EachAdjacent(g RGraph, this VIdx, onOther VisitVertex) (stopped bool) {
	switch u := g.(type) {
	case RUndirected:
		switch ls := u.(type) {
		case OutLister:
			return ls.EachOutgoing(this, onOther)
		case InLister:
			return ls.EachIncoming(this, onOther)
		default:
			return eachUAdj(u, this, onOther)
		}
	case OutLister:
		if u.EachOutgoing(this, onOther) {
			return true
		}
		if il, ok := g.(InLister); ok {
			return il.EachIncoming(this, onOther)
		} else {
			return eachDIn(g, this, onOther)
		}
	case InLister:
		if u.EachIncoming(this, onOther) {
			return true
		}
		if ol, ok := g.(OutLister); ok {
			return ol.EachOutgoing(this, onOther)
		} else {
			return eachDOut(g, this, onOther)
		}
	default:
		return eachDAdj(g, this, onOther)
	}
}

func eachDAdj(g RGraph, v VIdx, on VisitVertex) bool {
	ord := g.Order()
	for u := V0; u < ord; u++ {
		if g.Weight(v, u) != nil || g.Weight(u, v) != nil {
			if on(u) {
				return true
			}
		}
	}
	return false
}

func eachUAdj(u RUndirected, v VIdx, on VisitVertex) bool {
	vno := u.Order()
	n := V0
	for n < v {
		if u.WeightU(v, n) != nil {
			if on(n) {
				return true
			}
		}
		n++
	}
	for n < vno {
		if u.WeightU(n, v) != nil {
			if on(n) {
				return true
			}
		}
		n++
	}
	return false
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
