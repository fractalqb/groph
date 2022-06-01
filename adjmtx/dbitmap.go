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

package adjmtx

import (
	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/gimpl"
	"git.fractalqb.de/fractalqb/groph/internal"
)

// AdjMxDbitmap implements WGraph[bool] as a bitmap based adjacency
// matrix. Compared to AdjMxbool, this sacrifices runtime performance
// for lesser memory usage.
type DBitmap struct {
	adjMtx
	bs internal.BitSet
}

func NewDBitmap(order int, reuse *DBitmap) *DBitmap {
	// TODO better reuse memory
	if reuse == nil {
		reuse = new(DBitmap)
	}
	*reuse = DBitmap{
		adjMtx: adjMtx{ord: order},
		bs:     internal.NewBitSet(order*order, nil),
	}
	return reuse
}

func (m *DBitmap) Init(flag bool) *DBitmap {
	if flag {
		for i := range m.bs {
			m.bs[i] = ^uint64(0)
		}
	} else {
		for i := range m.bs {
			m.bs[i] = 0
		}
	}
	return m
}

func (m *DBitmap) Reset(order int) {
	NewDBitmap(order, m)
	m.Init(false)
}

func (g *DBitmap) Edge(u, v groph.VIdx) bool { return g.bs.Get(g.ord*u + v) }

func (g *DBitmap) NotEdge() bool { return false }

func (g *DBitmap) IsEdge(w bool) bool { return w }

func (g *DBitmap) SetEdge(u, v groph.VIdx, w bool) {
	if w {
		g.bs.Set(g.ord*u + v)
	} else {
		g.bs.Unset(g.ord*u + v)
	}
}

func (g *DBitmap) DelEdge(u, v groph.VIdx) { g.SetEdge(u, v, false) }

func (g *DBitmap) OutDegree(v groph.VIdx) int {
	return gimpl.DOutDegree[bool](g, v)
}

func (g *DBitmap) EachOut(from groph.VIdx, onDest groph.VisitVertex) error {
	return gimpl.DEachOut[bool](g, from, onDest)
}

func (g *DBitmap) InDegree(v groph.VIdx) int {
	return gimpl.DInDegree[bool](g, v)
}

func (g *DBitmap) EachIn(to groph.VIdx, onSource groph.VisitVertex) error {
	return gimpl.DEachIn[bool](g, to, onSource)
}

func (g *DBitmap) Size() int {
	return gimpl.DSize[bool](g)
}

func (g *DBitmap) EachEdge(onEdge groph.VisitEdge[bool]) error {
	return gimpl.DEachEdge(g, onEdge)
}

func (g *DBitmap) RootCount() int {
	return gimpl.DRootCount[bool](g)
}

func (g *DBitmap) EachRoot(onEdge groph.VisitVertex) error {
	return gimpl.DEachRoot[bool](g, onEdge)
}

func (g *DBitmap) LeafCount() int {
	return gimpl.DLeafCount[bool](g)
}

func (g *DBitmap) EachLeaf(onEdge groph.VisitVertex) error {
	return gimpl.DEachLeaf[bool](g, onEdge)
}
