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

package gimpl

import (
	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/internal"
)

var _ groph.WDirected[int] = (*InForest[int])(nil)

type uplink[W comparable] struct {
	up groph.VIdx
	w  W
}

type InForest[W comparable] struct {
	es  []uplink[W]
	noe W
}

func NewInForest[W comparable](order int, noe W) *InForest[W] {
	res := &InForest[W]{make([]uplink[W], order), noe}
	for i := range res.es {
		res.es[i].w = noe
	}
	return res
}

func (f *InForest[W]) Order() int { return len(f.es) }

func (f *InForest[W]) Edge(u, v groph.VIdx) (weight W) {
	if f.es[u].up == v {
		return f.es[u].w
	}
	return f.noe
}

func (f *InForest[W]) IsEdge(weight W) bool { return weight != f.noe }

func (f *InForest[W]) NotEdge() W { return f.noe }

func (f *InForest[W]) Size() (res int) {
	for _, l := range f.es {
		if l.up >= 0 && l.w != f.noe {
			res++
		}
	}
	return res
}

func (f *InForest[W]) EachEdge(onEdge groph.VisitEdge[W]) (stopped bool) {
	for u, l := range f.es {
		if l.up >= 0 && l.w != f.noe {
			if onEdge(u, l.up, l.w) {
				return true
			}
		}
	}
	return false
}

func (f *InForest[W]) Reset(order int) {
	f.es = internal.Slice(f.es, order)
	for i := range f.es {
		f.es[i].up = -1
	}
}

func (f *InForest[W]) SetEdge(u, v groph.VIdx, weight W) {
	f.es[u] = uplink[W]{v, weight}
}

func (f *InForest[W]) DelEdge(u, v groph.VIdx) {
	if f.es[u].up == v {
		f.es[u].w = f.noe
	}
}

func (f *InForest[W]) OutDegree(v groph.VIdx) int {
	if f.es[v].w == f.noe || f.es[v].up < 0 {
		return 0
	}
	return 1
}

func (f *InForest[W]) EachOut(from groph.VIdx, onDest groph.VisitVertex) (stopped bool) {
	e := f.es[from]
	if e.w == f.noe || e.up < 0 {
		return false
	}
	return onDest(e.up)
}

func (f *InForest[W]) InDegree(v groph.VIdx) int {
	for _, e := range f.es {
		if e.up == v && e.w != f.noe {
			return 1
		}
	}
	return 0
}

func (f *InForest[W]) EachIn(to groph.VIdx, onSource groph.VisitVertex) (stopped bool) {
	for from, e := range f.es {
		if e.up == to && e.w != f.noe {
			return onSource(from)
		}
	}
	return false
}

func (f *InForest[W]) RootCount() int { return DRootCount[W](f) }

func (f *InForest[W]) EachRoot(onEdge groph.VisitVertex) (stopped bool) {
	return DEachRoot[W](f, onEdge)
}

func (f *InForest[W]) LeafCount() (res int) {
	for _, e := range f.es {
		if e.w == f.noe || e.up < 0 {
			res++
		}
	}
	return res
}

func (f *InForest[W]) EachLeaf(onEdge groph.VisitVertex) (stopped bool) {
	for v, e := range f.es {
		if e.w == f.noe || e.up < 0 {
			if onEdge(v) {
				return true
			}
		}
	}
	return false
}
