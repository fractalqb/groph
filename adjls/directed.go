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

import (
	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/gimpl"
	"git.fractalqb.de/fractalqb/groph/internal"
)

type cell[W comparable] struct {
	v groph.VIdx
	w W
}

type Directed[W comparable] struct {
	es  [][]cell[W]
	sz  int
	noe W
}

var _ groph.WDirected[int] = (*Directed[int])(nil)

func NewDirected[W comparable](order int, notEdge W, reuse *Directed[W]) *Directed[W] {
	return &Directed[W]{
		es:  make([][]cell[W], order),
		noe: notEdge,
	}
}

func (g *Directed[W]) Order() int { return len(g.es) }

func (g *Directed[W]) Edge(u, v groph.VIdx) (weight W) {
	ls := g.es[u]
	for _, c := range ls {
		if c.v == v {
			return c.w
		}
	}
	return g.noe
}

func (g *Directed[W]) IsEdge(weight W) bool { return weight != g.noe }

func (g *Directed[W]) NotEdge() W { return g.noe }

func (g *Directed[W]) Size() int { return g.sz }

func (g *Directed[W]) EachEdge(onEdge groph.VisitEdge[W]) error {
	for u, cls := range g.es {
		for _, c := range cls {
			if err := onEdge(u, c.v, c.w); err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Directed[W]) Reset(order int) {
	g.es = internal.Slice(g.es, order)
	g.sz = 0
	for i := range g.es {
		g.es[i] = g.es[i][:0]
	}
}

func (g *Directed[W]) SetEdge(u, v groph.VIdx, weight W) {
	ls := g.es[u]
	for i := range ls {
		c := &ls[i]
		if c.v == v {
			c.w = weight
			return
		}
	}
	g.es[u] = append(ls, cell[W]{v: v, w: weight})
}

func (g *Directed[W]) DelEdge(u, v groph.VIdx) {
	ls := g.es[u]
	for i, c := range ls {
		if c.v == v {
			copy(ls[i:], ls[i+1:])
			g.es[u] = ls[:len(ls)-1]
			return
		}
	}
}

func (g *Directed[W]) OutDegree(v groph.VIdx) int { return len(g.es[v]) }

func (g *Directed[W]) EachOut(from groph.VIdx, onDest groph.VisitVertex) error {
	for _, c := range g.es[from] {
		if err := onDest(c.v); err != nil {
			return err
		}
	}
	return nil
}

func (g *Directed[W]) InDegree(v groph.VIdx) (d int) {
	for _, ls := range g.es {
		for _, c := range ls {
			if c.v == v {
				d++
				break
			}
		}
	}
	return d
}

func (g *Directed[W]) EachIn(to groph.VIdx, onSource groph.VisitVertex) error {
	for v, ls := range g.es {
		for _, c := range ls {
			if c.v == v {
				if err := onSource(v); err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}

func (g *Directed[W]) RootCount() int {
	return gimpl.DRootCount[W](g)
}

func (g *Directed[W]) EachRoot(onEdge groph.VisitVertex) error {
	return gimpl.DEachRoot[W](g, onEdge)
}

func (g *Directed[W]) LeafCount() (l int) {
	for _, ls := range g.es {
		if len(ls) == 0 {
			l++
		}
	}
	return l
}

func (g *Directed[W]) EachLeaf(onEdge groph.VisitVertex) error {
	for v, ls := range g.es {
		if len(ls) == 0 {
			if err := onEdge(v); err != nil {
				return err
			}
		}
	}
	return nil
}
