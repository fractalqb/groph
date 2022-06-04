// Copyright 2022 Marcus Perlick
// This file is part of Go module git.fractalqb.de/fractalqb/groph
//
// groph is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// groph is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with groph.  If not, see <http://www.gnu.org/licenses/>.

package groph

const (
	IntNoEdge int = minInt

	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

type VisitVertex = func(v VIdx) error

type VisitEdge = func(u, v VIdx) error

type VisitEdgeW[W any] func(u, v VIdx, w W) error

type stop int

func (stop) Error() string { return "stopped" }

const Stopped stop = 0

// VIdx is the type used to represent vertices in the graph implementations
// provided by the groph package.
type VIdx = int

// RGraph represents graph that allows read only access to the egde
// weights.
//
// For graphs that can change be modified see WGraph. For undirected
// graphs see also RUndirected.
type RGraph[W any] interface {
	// Order return the numer of vertices in the graph.
	Order() int

	// Returns the weight of the edge that connects the vertex with index
	// u with the vertex with index v. If there is no such edge it returns nil.
	Edge(u, v VIdx) (weight W)

	// IsEdge returns true if and only if weight denotes an existing edge.
	IsEdge(weight W) bool

	// NotEdge return a W typed value that denotes a non-existing edge, i.e.
	// IsEdge(NotEdge()) == false.
	NotEdge() W

	// Size returns the number of edges of the graph.
	Size() int

	// EachEdge calls onEdge for each edge in the graph. When onEdge returns
	// true EachEdge stops immediately and returns true. Otherwise EachEdge
	// returns false.
	EachEdge(onEdge VisitEdgeW[W]) error
}

type RDirected[W any] interface {
	RGraph[W]

	// OutDegree returns the number of outgoing edges of vertex v in graph
	// g. Note that for undirected graphs each edge is also considered to
	// be an outgoing edge.
	OutDegree(v VIdx) int

	EachOut(from VIdx, onDest VisitVertex) error

	// InDegree returns the number of incoming edges of vertex v in graph
	// g. Note that for undirected graphs each edge is also considered to
	// be an incoming edge.
	InDegree(v VIdx) int

	EachIn(to VIdx, onSource VisitVertex) error

	RootCount() int

	EachRoot(onEdge VisitVertex) error

	LeafCount() int

	EachLeaf(onEdge VisitVertex) error
}

// RUndirected represents an undirected graph that allows read only
// access to the edge weights. For undirected graphs each edge (u,v) is
// considered outgiong as well as incoming for both, vertex u and vertext v.
type RUndirected[W any] interface {
	RGraph[W]
	// Weight must only be called when u ≥ v.  Otherwise WeightU's
	// behaviour is unspecified, it even might crash.  In many
	// implementations this can be way more efficient than the
	// general case, see method Weight().
	EdgeU(u, v VIdx) (weight W)

	Degree(v VIdx) int

	EachAdjacent(of VIdx, onNeighbour VisitVertex) error
}

// WGraph represents graph that allows read and write access to the
// egde weights.
//
// For undirected graphs see also WUndirected.
type WGraph[W any] interface {
	RGraph[W]
	// Reset resizes the graph to order and reinitializes it. Implementations
	// are expected to reuse memory.
	Reset(order int)
	// SetWeight sets the edge weight for the edge starting at vertex u and
	// ending at vertex v. Passing w==nil removes the edge.
	SetEdge(u, v VIdx, weight W)

	DelEdge(u, v VIdx)
}

type WDirected[W any] interface {
	WGraph[W]
	RDirected[W]
}

// WUndirected represents an undirected graph that allows read and
// write access to the egde weights.
type WUndirected[W any] interface {
	WGraph[W]
	RUndirected[W]

	// SetWeightU must only be called when u ≥ v.  Otherwise
	// SetWeightU's behaviour is unspecified, it even might crash.
	//
	// See also RUndirected.WeightU
	SetEdgeU(u, v VIdx, w W)

	DelEdgeU(u, v VIdx)
}
