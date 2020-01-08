package groph

type VisitVertex = func(v VIdx) (stop bool)

type VisitEdge = func(u, v VIdx) (stop bool)

// OutLister is implemented by graph implementations that can easily iterate
// over all outgoing edges of one node.
//
// See also OutDegree and traverse.EachOutgoing function.
type OutLister interface {
	EachOutgoing(from VIdx, onDest VisitVertex) (stopped bool)
	OutDegree(v VIdx) int
}

// InLister is implemented by graph implementations that can easily iterate
// over all incoming edges of one node.
//
// See also InDegree and traverse.EachIncoming function.
type InLister interface {
	EachIncoming(to VIdx, onSource VisitVertex) (stopped bool)
	InDegree(v VIdx) int
}

// InLister is implemented by graph implementations that can easily iterate
// over all edges of the graph.
//
// See also Size and traverse.EachEdge function.
type EdgeLister interface {
	EachEdge(onEdge VisitEdge) (stop bool)
	Size() int
}

type RootsLister interface {
	EachRoot(onEdge VisitVertex) (stop bool)
	RootCount() int
}

type LeavesLister interface {
	EachLeaf(onEdge VisitVertex) (stop bool)
	LeafCount() int
}
