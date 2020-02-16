package groph

// OutDegree returns the number of outgoing edges of vertex v in graph
// g. Note that for undirected graphs each edge is also considered to
// be an outgoing edge.
func OutDegree(g RGraph, v VIdx) (res int) {
	incRes := func(_ VIdx) bool { res++; return false }
	switch gi := g.(type) {
	case OutLister:
		return gi.OutDegree(v)
	case RUndirected:
		eachUAdj(gi, v, incRes)
	default:
		eachDOut(g, v, incRes)
	}
	return res
}

// InDegree returns the number of incoming edges of vertex v in graph
// g. Note that for undirected graphs each edge is also considered to
// be an incoming edge.
func InDegree(g RGraph, v VIdx) (res int) {
	incRes := func(_ VIdx) bool { res++; return false }
	switch gi := g.(type) {
	case InLister:
		return gi.InDegree(v)
	case RUndirected:
		eachUAdj(gi, v, incRes)
	default:
		eachDIn(g, v, incRes)
	}
	return res
}

func Degree(g RGraph, v VIdx) (res int) {
	incRes := func(_ VIdx) bool { res++; return false }
	switch tg := g.(type) {
	case RUndirected:
		switch ls := tg.(type) {
		case OutLister:
			return ls.OutDegree(v)
		case InLister:
			return ls.InDegree(v)
		default:
			eachUAdj(tg, v, incRes)
		}
	case OutLister:
		res = tg.OutDegree(v)
		if il, ok := g.(InLister); ok {
			res += il.InDegree(v)
		} else {
			eachDIn(g, v, incRes)
		}
		if g.Weight(v, v) != nil {
			res--
		}
	case InLister:
		res = tg.InDegree(v)
		if ol, ok := g.(OutLister); ok {
			res += ol.OutDegree(v)
		} else {
			eachDOut(g, v, incRes)
		}
		if g.Weight(v, v) != nil {
			res--
		}
	default:
		eachDAdj(g, v, incRes)
	}
	return res
}

// Size returns the number of edges in the graph g.
func Size(g RGraph) (res int) {
	switch xl := g.(type) {
	case EdgeLister:
		return xl.Size()
	case RUndirected:
		ord := g.Order()
		switch ls := g.(type) {
		case OutLister:
			for i := 0; i < ord; i++ {
				res += ls.OutDegree(i)
				if g.Weight(i, i) != nil {
					res++
				}
			}
			res /= 2
		case InLister:
			for i := 0; i < ord; i++ {
				res += ls.InDegree(i)
				if g.Weight(i, i) != nil {
					res++
				}
			}
			res /= 2
		default:
			for i := 0; i < ord; i++ {
				for j := 0; j <= i; j++ {
					if xl.WeightU(i, j) != nil {
						res++
					}
				}
			}
		}
	case OutLister:
		ord := g.Order()
		for i := 0; i < ord; i++ {
			res += xl.OutDegree(i)
		}
	case InLister:
		ord := g.Order()
		for i := 0; i < ord; i++ {
			res += xl.InDegree(i)
		}
	default:
		ord := g.Order()
		for i := 0; i < ord; i++ {
			for j := 0; j < ord; j++ {
				if g.Weight(i, j) != nil {
					res++
				}
			}
		}
	}
	return res
}

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
	for n := 0; n < vno; n++ {
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
	for n := 0; n < vno; n++ {
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
	for u := 0; u < ord; u++ {
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
	n := 0
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
func EachEdge(g RGraph, onEdge VisitEdge) (stopped bool) {
	// TODO optimize with In-/OutLister
	switch ge := g.(type) {
	case EdgeLister:
		return ge.EachEdge(onEdge)
	case RUndirected:
		vno := ge.Order()
		for i := 0; i < vno; i++ {
			for j := 0; j <= i; j++ {
				if w := ge.WeightU(i, j); w != nil {
					if onEdge(i, j) {
						return true
					}
				}
			}
		}
	default:
		vno := g.Order()
		for i := 0; i < vno; i++ {
			for j := 0; j < vno; j++ {
				if w := g.Weight(i, j); w != nil {
					if onEdge(i, j) {
						return true
					}
				}
			}
		}
	}
	return false
}
