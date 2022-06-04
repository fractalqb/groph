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
)

var _ groph.WDirected[int] = (*Directed[int])(nil)

type Directed[W comparable] struct {
	adjlist[W]
}

func NewDirected[W comparable](order int, notEdge W, reuse *Directed[W]) *Directed[W] {
	if reuse == nil {
		reuse = &Directed[W]{*newAdjList(notEdge)}
	}
	reuse.Reset(order)
	return reuse
}

func (g *Directed[W]) Edge(u, v groph.VIdx) (weight W) {
	ls := g.es[u]
	for _, c := range ls {
		if c.v == v {
			return c.w
		}
	}
	return g.noe
}

func (g *Directed[W]) SetEdge(u, v groph.VIdx, w W) { g.set(u, v, w) }

func (g *Directed[W]) DelEdge(u, v groph.VIdx) { g.del(u, v) }

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
