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

package shortpath

import (
	"container/heap"

	"golang.org/x/exp/constraints"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/gimpl"
)

type minVItem[W constraints.Ordered] struct {
	v groph.VIdx
	d W
}

type minVertex[W constraints.Ordered] struct {
	v2idx []int
	itms  []minVItem[W]
	g     groph.RGraph[W]
}

func (mv *minVertex[W]) init(g groph.RGraph[W], v groph.VIdx) {
	ord := g.Order()
	mv.v2idx = make([]int, ord)
	mv.itms = make([]minVItem[W], 0, ord-1)
	mv.g = g
	mv.v2idx[v] = -1
	for w := groph.VIdx(0); w < ord; w++ {
		if w == v {
			continue
		}
		mv.v2idx[w] = len(mv.itms)
		mv.itms = append(mv.itms, minVItem[W]{
			v: w,
			d: g.Edge(v, w),
		})
	}
	heap.Init(mv)
}

func (mv *minVertex[W]) Len() int { return len(mv.itms) }

func (mv *minVertex[W]) Less(i, j int) bool {
	if !mv.g.IsEdge(mv.itms[i].d) {
		return false
	} else if !mv.g.IsEdge(mv.itms[j].d) {
		return true
	}
	return mv.itms[i].d < mv.itms[j].d
}

func (mv *minVertex[W]) Swap(i, j int) {
	mv.itms[j], mv.itms[i] = mv.itms[i], mv.itms[j]
	mv.v2idx[mv.itms[i].v] = i
	mv.v2idx[mv.itms[j].v] = j
}

func (mv *minVertex[W]) Push(x interface{}) {
	panic("minVertex.Push must not be called")
	// v := len(mv.v2idx)
	// mv.v2idx = append(mv.v2idx, v)
	// mv.itms = append(mv.itms, minVItem[W]{w: x.(W), v: v})
}

func (mv *minVertex[W]) Pop() interface{} {
	lm1 := len(mv.itms) - 1
	itm := mv.itms[lm1]
	mv.itms = mv.itms[:lm1]
	mv.v2idx[itm.v] = -1
	return itm
}

func DijkstraD[W constraints.Ordered, G groph.RDirected[W]](g G, v groph.VIdx) *gimpl.InForest[W] {
	var minv minVertex[W]
	minv.init(g, v)
	res := gimpl.NewInForest(g.Order(), g.NotEdge())
	for minv.Len() > 0 {
		top := heap.Pop(&minv).(minVItem[W])
		if !g.IsEdge(top.d) {
			break
		}
		res.SetEdge(v, top.v, top.d)
		g.EachOut(top.v, func(n groph.VIdx) bool {
			if minv.v2idx[n] < 0 {
				return false
			}
			d := g.Edge(top.v, n)
			if g.IsEdge(d) {
				vi := minv.v2idx[n]
				nv := &minv.itms[vi]
				nv.d = top.d + d
				heap.Fix(&minv, vi)
			}
			return false
		})
		v = top.v
	}
	return res
}

func DijkstraU[W constraints.Ordered, G groph.RUndirected[W]](g G, v groph.VIdx) *gimpl.InForest[W] {
	var minv minVertex[W]
	minv.init(g, v)
	res := gimpl.NewInForest(g.Order(), g.NotEdge())
	for minv.Len() > 0 {
		top := heap.Pop(&minv).(minVItem[W])
		if !g.IsEdge(top.d) {
			break
		}
		res.SetEdge(v, top.v, top.d)
		g.EachAdjacent(top.v, func(n groph.VIdx) bool {
			if minv.v2idx[n] < 0 {
				return false
			}
			d := g.Edge(top.v, n)
			if g.IsEdge(d) {
				vi := minv.v2idx[n]
				nv := &minv.itms[vi]
				nv.d = top.d + d
				heap.Fix(&minv, vi)
			}
			return false
		})
		v = top.v
	}
	return res
}
