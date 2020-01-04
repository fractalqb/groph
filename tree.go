package groph

type Tree []VIdx

func (t Tree) Root() (res VIdx) {
	max := len(t)
	for n := t[res]; n >= 0; n = t[res] {
		res = n
		if max--; max == 0 {
			return -1
		}
	}
	return res
}

func (t Tree) Order() VIdx { return len(t) }

func (t Tree) EdgeU(u, v VIdx) bool {
	return t[u] == v || t[v] == u
}

func (t Tree) Edge(u, v VIdx) bool {
	if u > v {
		return t.EdgeU(u, v)
	}
	return t.EdgeU(v, u)
}

func (t Tree) WeightU(u, v VIdx) interface{} {
	if t.EdgeU(u, v) {
		return true
	}
	return nil
}

func (t Tree) Weight(u, v VIdx) interface{} {
	if t.Edge(u, v) {
		return true
	}
	return nil
}

func (t Tree) EachOutgoing(from VIdx, onDest VisitVertex) (stop bool) {
	if dest := t[from]; dest >= 0 {
		if onDest(dest) {
			return true
		}
	}
	return false
}

func (t Tree) OutDegree(v VIdx) int {
	if t[v] < 0 {
		return 0
	}
	return 1
}

func (t Tree) EachEdge(onEdge VisitEdge) (stop bool) {
	for u, v := range t {
		if v >= 0 {
			if onEdge(u, v) {
				return true
			}
		}
	}
	return false
}

func (t Tree) Size() int { return len(t) - 1 }
