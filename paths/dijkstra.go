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

package paths

import (
	"container/heap"

	"golang.org/x/exp/constraints"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/graphs"
	"git.fractalqb.de/fractalqb/groph/internal"
)

func dijkstraVHeapD[W constraints.Ordered](g groph.RDirected[W], v groph.VIdx) *internal.VHeap[W] {
	h := internal.NewVHeap[W](g.Order())
	ord := g.Order()
	for u := 0; u < ord; u++ {
		if u == v {
			h.AddNotVertex(u)
		}
		if w := g.Edge(v, u); g.IsEdge(w) {
			h.AddVertex(u, w)
		} else {
			h.AddNotVertex(u)
		}
	}
	heap.Init(h)
	return h
}

func DijkstraD[W constraints.Ordered, G groph.RDirected[W]](g G, v groph.VIdx) *graphs.InForest[W] {
	minv := *dijkstraVHeapD[W](g, v)
	res := graphs.NewInForest(g.Order(), g.NotEdge())
	for minv.Len() > 0 {
		u, w := minv.PopVertex()
		res.SetEdge(v, u, w)
		g.EachOut(u, func(n groph.VIdx) error {
			if !minv.Has(n) {
				return nil
			}
			d := g.Edge(u, n)
			if g.IsEdge(d) {
				minv.Set(n, w+d)
			}
			return nil
		})
		v = u
	}
	return res
}

func dijkstraVHeapU[W constraints.Ordered](g groph.RUndirected[W], v groph.VIdx) *internal.VHeap[W] {
	h := internal.NewVHeap[W](g.Order())
	for u := 0; u < v; u++ {
		if w := g.EdgeU(v, u); g.IsEdge(w) {
			h.AddVertex(u, w)
		} else {
			h.AddNotVertex(u)
		}
	}
	h.AddNotVertex(v)
	ord := g.Order()
	for u := v + 1; u < ord; u++ {
		if w := g.EdgeU(u, v); g.IsEdge(w) {
			h.AddVertex(u, w)
		} else {
			h.AddNotVertex(u)
		}
	}
	heap.Init(h)
	return h
}

func DijkstraU[W constraints.Ordered, G groph.RUndirected[W]](g G, v groph.VIdx) *graphs.InForest[W] {
	minv := *dijkstraVHeapU[W](g, v)
	res := graphs.NewInForest(g.Order(), g.NotEdge())
	for minv.Len() > 0 {
		u, w := minv.PopVertex()
		res.SetEdge(v, u, w)
		g.EachAdjacent(u, func(n groph.VIdx) error {
			if !minv.Has(n) {
				return nil
			}
			d := g.Edge(u, n)
			if g.IsEdge(d) {
				minv.Set(n, w+d)
			}
			return nil
		})
		v = u
	}
	return res
}
