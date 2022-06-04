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

// Package adjls provides graph implementations as adjacency lists.
package adjls

import (
	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/internal"
)

type cell[W comparable] struct {
	v groph.VIdx
	w W
}

type adjlist[W comparable] struct {
	es  [][]cell[W]
	sz  int
	noe W
}

func newAdjList[W comparable](noe W) *adjlist[W] {
	return &adjlist[W]{noe: noe}
}

func (g *adjlist[W]) Order() int { return len(g.es) }

func (g *adjlist[W]) IsEdge(weight W) bool { return weight != g.noe }

func (g *adjlist[W]) NotEdge() W { return g.noe }

func (g *adjlist[W]) Size() int { return g.sz }

func (g *adjlist[W]) Reset(order int) {
	g.es, _ = internal.Slice(g.es, order)
	g.sz = 0
	for i := range g.es {
		g.es[i] = g.es[i][:0]
	}
}

func (g *adjlist[W]) EachEdge(onEdge groph.VisitEdgeW[W]) error {
	for u, cls := range g.es {
		for _, c := range cls {
			if err := onEdge(u, c.v, c.w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *adjlist[W]) set(u, v groph.VIdx, w W) {
	if g.IsEdge(w) {
		ls := g.es[u]
		for i := range ls {
			c := &ls[i]
			if c.v == v {
				if c.w != w {
					c.w = w
					g.sz++
				}
				return
			}
		}
		g.es[u] = append(ls, cell[W]{v: v, w: w})
		g.sz++
	} else {
		g.del(u, v)
	}
}

func (g *adjlist[W]) del(u, v groph.VIdx) {
	ls := g.es[u]
	for i, c := range ls {
		if c.v == v {
			copy(ls[i:], ls[i+1:])
			g.es[u] = ls[:len(ls)-1]
			g.sz--
			return
		}
	}
}
