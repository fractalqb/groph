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

package adjls

import "git.fractalqb.de/fractalqb/groph"

var _ groph.WUndirected[int] = (*Undirected[int])(nil)

type Undirected[W comparable] struct {
	adjlist[W]
}

func NewUndirected[W comparable](order int, notEdge W, reuse *Undirected[W]) *Undirected[W] {
	if reuse == nil {
		reuse = &Undirected[W]{*newAdjList(notEdge)}
	}
	reuse.Reset(order)
	return reuse
}

func (g *Undirected[W]) Edge(u, v groph.VIdx) (weight W) {
	if u < v {
		return g.EdgeU(v, u)
	}
	return g.EdgeU(u, v)
}

func (g *Undirected[W]) EdgeU(u, v groph.VIdx) (weight W) {
	ls := g.es[u]
	for _, c := range ls {
		if c.v == v {
			return c.w
		}
	}
	return g.noe
}

func (g *Undirected[W]) SetEdgeU(u, v groph.VIdx, w W) { g.set(u, v, w) }

func (g *Undirected[W]) SetEdge(u, v groph.VIdx, w W) {
	if u < v {
		g.SetEdgeU(v, u, w)
	} else {
		g.SetEdgeU(u, v, w)
	}
}

func (g *Undirected[W]) DelEdgeU(u, v groph.VIdx) { g.del(u, v) }

func (g *Undirected[W]) DelEdge(u, v groph.VIdx) {
	if u < v {
		g.DelEdgeU(v, u)
	} else {
		g.DelEdgeU(u, v)
	}
}

func (g *Undirected[W]) Degree(v groph.VIdx) int {
	ls := g.es[v]
	d := len(ls)
	for _, c := range ls {
		if c.v == v {
			d++
		}
	}
	ord := g.Order()
	for u := v + 1; u < ord; u++ {
		for _, c := range g.es[u] {
			if c.v == v {
				d++
				break
			}
		}
	}
	return d
}

func (g *Undirected[W]) EachAdjacent(of groph.VIdx, onNeighbour groph.VisitVertex) error {
	// TODO NYI
	return nil
}
