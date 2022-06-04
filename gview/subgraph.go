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

package gview

import (
	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/gimpl"
)

type subgraph[W any, G groph.RGraph[W]] struct {
	G G
	V []groph.VIdx
}

func (g subgraph[W, G]) Order() int { return len(g.V) }

func (g subgraph[W, G]) Edge(u, v groph.VIdx) (weight W) {
	gu, gv := g.V[u], g.V[v]
	return g.G.Edge(gu, gv)
}

func (g subgraph[W, G]) NotEdge() W { return g.G.NotEdge() }

func (g subgraph[W, G]) IsEdge(weight W) bool { return g.G.IsEdge(weight) }

type DirectedSub[W any] struct {
	subgraph[W, groph.RDirected[W]]
}

func NewDirectedSub[W any](g groph.RDirected[W], v ...groph.VIdx) DirectedSub[W] {
	return DirectedSub[W]{
		subgraph: subgraph[W, groph.RDirected[W]]{
			G: g,
			V: v,
		},
	}
}

func (g DirectedSub[W]) Size() int {
	return gimpl.DSize[W](g)
}

func (g DirectedSub[W]) EachEdge(onEdge groph.VisitEdgeW[W]) error {
	return gimpl.DEachEdge(g, onEdge)
}

func (g DirectedSub[W]) OutDegree(v groph.VIdx) int {
	return gimpl.DOutDegree[W](g, v)
}

func (g DirectedSub[W]) EachOut(from groph.VIdx, onDest groph.VisitVertex) error {
	return gimpl.DEachOut[W](g, from, onDest)
}

func (g DirectedSub[W]) InDegree(v groph.VIdx) int {
	return gimpl.DInDegree[W](g, v)
}

func (g DirectedSub[W]) EachIn(to groph.VIdx, onSource groph.VisitVertex) error {
	return gimpl.DEachIn[W](g, to, onSource)
}

func (g DirectedSub[W]) RootCount() int {
	return gimpl.DRootCount[W](g)
}

func (g DirectedSub[W]) EachRoot(onEdge groph.VisitVertex) error {
	return gimpl.DEachRoot[W](g, onEdge)
}

func (g DirectedSub[W]) LeafCount() int {
	return gimpl.DLeafCount[W](g)
}

func (g DirectedSub[W]) EachLeaf(onEdge groph.VisitVertex) error {
	return gimpl.DEachLeaf[W](g, onEdge)
}

type UndirectedSub[W any] struct {
	subgraph[W, groph.RUndirected[W]]
}

func NewUndirectedSub[W any](g groph.RUndirected[W], v ...groph.VIdx) UndirectedSub[W] {
	return UndirectedSub[W]{
		subgraph: subgraph[W, groph.RUndirected[W]]{
			G: g,
			V: v,
		},
	}
}

func (g UndirectedSub[W]) Size() int {
	return gimpl.USize[W](g)
}

func (g UndirectedSub[W]) EachEdge(onEdge groph.VisitEdgeW[W]) error {
	return gimpl.UEachEdge(g, onEdge)
}

func (g UndirectedSub[W]) EdgeU(u, v groph.VIdx) (weight W) {
	gu, gv := g.V[u], g.V[v]
	return g.G.EdgeU(gu, gv)
}

func (g UndirectedSub[W]) Degree(v groph.VIdx) int {
	return gimpl.UDegree[W](g, v)
}

func (g UndirectedSub[W]) EachAdjacent(of groph.VIdx, onNeighbour groph.VisitVertex) error {
	return gimpl.UEachAdjacent[W](g, of, onNeighbour)
}
