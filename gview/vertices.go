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

import "git.fractalqb.de/fractalqb/groph"

type Vertices[V comparable] struct {
	byIdx []V
	byVtx map[V]groph.VIdx
}

func NewVertices[V comparable](vs ...V) *Vertices[V] {
	res := &Vertices[V]{
		byIdx: vs,
		byVtx: make(map[V]groph.VIdx),
	}
	for i, v := range res.byIdx {
		res.byVtx[v] = i
	}
	return res
}

func (vs *Vertices[V]) Order() int { return len(vs.byIdx) }

func (vs *Vertices[V]) IndexOf(v V) groph.VIdx {
	i, ok := vs.byVtx[v]
	if !ok {
		return -1
	}
	return i
}

func (vs *Vertices[V]) VertexOf(i groph.VIdx) V { return vs.byIdx[i] }

type vertRD[V comparable, W any] struct {
	groph.RDirected[W]
	*Vertices[V]
}

func VertexedRD[V comparable, W any](vs *Vertices[V], g groph.RDirected[W]) vertRD[V, W] {
	return vertRD[V, W]{g, vs}
}

func (g *vertRD[V, W]) Order() int { return g.RDirected.Order() }

func (g *vertRD[V, W]) VEdge(u, v V) W {
	return g.Edge(g.IndexOf(u), g.IndexOf(v))
}

type vertWD[V comparable, W any] struct {
	groph.WDirected[W]
	*Vertices[V]
}

func VertexedWD[V comparable, W any](vs *Vertices[V], g groph.WDirected[W]) vertWD[V, W] {
	return vertWD[V, W]{g, vs}
}

func (g *vertWD[V, W]) Order() int { return g.WDirected.Order() }

func (g *vertWD[V, W]) VEdge(u, v V) W {
	return g.Edge(g.IndexOf(u), g.IndexOf(v))
}

func (g *vertWD[V, W]) VSetEdge(u, v V, w W) {
	g.SetEdge(g.IndexOf(u), g.IndexOf(v), w)
}
