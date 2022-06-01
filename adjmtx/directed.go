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
	"errors"
	"math"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/gimpl"
	"git.fractalqb.de/fractalqb/groph/internal"
)

const (
	IntNoEdge int = minInt

	maxUint = ^uint(0)
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

type adjMtx struct {
	ord int
}

func (m *adjMtx) Order() int { return m.ord }

type Directed[W comparable] struct {
	adjMtx
	ws  []W
	noe W
}

func NewDirected[W comparable](order int, notEdge W, reuse *Directed[W]) *Directed[W] {
	if reuse == nil {
		reuse = new(Directed[W])
	}
	reuse.noe = notEdge
	reuse.Reset(order)
	return reuse
}

func AsDirected[W comparable](reuse *Directed[W], notEdge W, weights ...W) (*Directed[W], error) {
	ord, err := dOrdFromLen(len(weights))
	if err != nil {
		return nil, err
	}
	if reuse == nil {
		reuse = new(Directed[W])
	}
	reuse.ord = ord
	reuse.ws = weights
	reuse.noe = notEdge
	return reuse, nil
}

func (g *Directed[W]) Edge(u, v groph.VIdx) W { return g.ws[g.ord*u+v] }

func (g *Directed[W]) NotEdge() W { return g.noe }

func (g *Directed[W]) IsEdge(w W) bool { return w != g.noe }

func (g *Directed[W]) Reset(order int) {
	g.ord = order
	g.ws = internal.Slice(g.ws, order*order)
	var zeroW W
	if g.noe != zeroW { // Skip if .noe is the zero value of W
		for i := range g.ws {
			g.ws[i] = g.noe
		}
	}
}

func (g *Directed[W]) SetEdge(u, v groph.VIdx, w W) { g.ws[g.ord*u+v] = w }

func (g *Directed[W]) DelEdge(u, v groph.VIdx) { g.SetEdge(u, v, g.noe) }

func (g *Directed[W]) OutDegree(v groph.VIdx) int {
	return gimpl.DOutDegree[W](g, v)
}

func (g *Directed[W]) EachOut(from groph.VIdx, onDest groph.VisitVertex) error {
	return gimpl.DEachOut[W](g, from, onDest)
}

func (g *Directed[W]) InDegree(v groph.VIdx) int {
	return gimpl.DInDegree[W](g, v)
}

func (g *Directed[W]) EachIn(to groph.VIdx, onSource groph.VisitVertex) error {
	return gimpl.DEachIn[W](g, to, onSource)
}

func (g *Directed[W]) Size() (s int) {
	// return gimpl.DSize[W](g) – speeds up with local impl
	for _, w := range g.ws {
		if g.IsEdge(w) {
			s++
		}
	}
	return s
}

func (g *Directed[W]) EachEdge(onEdge groph.VisitEdge[W]) error {
	// return gimpl.DEachEdge[W](g, onEdge) – speeds up with local impl
	var u, v groph.VIdx
	ord := g.Order()
	for _, w := range g.ws {
		if g.IsEdge(w) {
			if err := onEdge(u, v, w); err != nil {
				return err
			}
		}
		v++ // Depends on g.Edge()
		if v == ord {
			u++
			v = 0
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

func (g *Directed[W]) LeafCount() int {
	return gimpl.DLeafCount[W](g)
}

func (g *Directed[W]) EachLeaf(onEdge groph.VisitVertex) error {
	return gimpl.DEachLeaf[W](g, onEdge)
}

func dOrdFromLen(sliceLen int) (order int, err error) {
	order = int(math.Sqrt(float64(sliceLen)))
	if order*order != sliceLen {
		return order, errors.New("weights slice is not square")
	}
	return order, nil
}
