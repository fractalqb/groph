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

package graphs

import (
	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/gimpl"
)

// TODO sort edges for faster access
type Edge struct {
	U, V groph.VIdx
	W    any
}

var _ groph.RDirected[any] = DirectedEdges{}

type DirectedEdges []Edge

func (g DirectedEdges) Order() int {
	ord := -1
	for _, e := range g {
		if e.U > ord {
			ord = e.U
		}
		if e.V > ord {
			ord = e.V
		}
	}
	return ord + 1
}

func (g DirectedEdges) Edge(u, v groph.VIdx) (weight any) {
	for _, e := range g {
		if e.U == u && e.V == v {
			return e.W
		}
	}
	return nil
}

func (g DirectedEdges) IsEdge(weight any) bool { return weight != nil }

func (g DirectedEdges) NotEdge() any { return nil }

func (g DirectedEdges) Size() int { return len(g) }

func (g DirectedEdges) EachEdge(onEdge groph.VisitEdgeW[any]) error {
	for _, e := range g {
		if err := onEdge(e.U, e.V, e.W); err != nil {
			return err
		}
	}
	return nil
}

func (g DirectedEdges) OutDegree(v groph.VIdx) int {
	return gimpl.DOutDegree[any](g, v)
}

func (g DirectedEdges) EachOut(from groph.VIdx, onDest groph.VisitVertex) error {
	return gimpl.DEachOut[any](g, from, onDest)
}

func (g DirectedEdges) InDegree(v groph.VIdx) int {
	return gimpl.DInDegree[any](g, v)
}

func (g DirectedEdges) EachIn(to groph.VIdx, onSource groph.VisitVertex) error {
	return gimpl.DEachIn[any](g, to, onSource)
}

func (g DirectedEdges) RootCount() int {
	return gimpl.DRootCount[any](g)
}

func (g DirectedEdges) EachRoot(onEdge groph.VisitVertex) error {
	return gimpl.DEachRoot[any](g, onEdge)
}

func (g DirectedEdges) LeafCount() int {
	return gimpl.DLeafCount[any](g)
}

func (g DirectedEdges) EachLeaf(onEdge groph.VisitVertex) error {
	return gimpl.DEachLeaf[any](g, onEdge)
}

var _ groph.RUndirected[any] = UndirectedEdges{}

type UndirectedEdges []Edge

func (g UndirectedEdges) Order() int {
	ord := -1
	for _, e := range g {
		if e.U > ord {
			ord = e.U
		}
		if e.V > ord {
			ord = e.V
		}
	}
	return ord + 1
}

func (g UndirectedEdges) Edge(u, v groph.VIdx) (weight any) {
	for _, e := range g {
		if (e.U == u && e.V == v) || (e.U == v && e.V == u) {
			return e.W
		}
	}
	return nil
}

func (g UndirectedEdges) IsEdge(weight any) bool { return weight != nil }

func (g UndirectedEdges) NotEdge() any { return nil }

func (g UndirectedEdges) Size() int { return len(g) }

func (g UndirectedEdges) EachEdge(onEdge groph.VisitEdgeW[any]) error {
	for _, e := range g {
		if err := onEdge(e.U, e.V, e.W); err != nil {
			return err
		}
	}
	return nil
}

func (g UndirectedEdges) EdgeU(u, v groph.VIdx) (weight any) {
	return g.Edge(u, v)
}

func (g UndirectedEdges) Degree(v groph.VIdx) int {
	return gimpl.UDegree[any](g, v)
}

func (g UndirectedEdges) EachAdjacent(of groph.VIdx, onNeighbour groph.VisitVertex) error {
	return gimpl.UEachAdjacent[any](g, of, onNeighbour)
}
