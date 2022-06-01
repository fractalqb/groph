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
)

type Undirected[W comparable] struct {
	adjMtx
	ws  []W
	noe W
}

func NewUndirected[W comparable](order int, notEdge W, reuse *Undirected[W]) *Undirected[W] {
	if reuse == nil {
		reuse = new(Undirected[W])
	}
	reuse.noe = notEdge
	reuse.Reset(order)
	return reuse
}

func (g *Undirected[W]) Reset(order int) {
	g.ord = order
	g.ws = make([]W, nSum(order))
	var zeroW W
	if g.noe != zeroW {
		for i := range g.ws {
			g.ws[i] = g.noe
		}
	}
}

func (g *Undirected[W]) Edge(u, v groph.VIdx) W {
	// return gimpl.UEdge[W](g, u, v) – speeds up with local impl
	if u < v {
		return g.EdgeU(v, u)
	}
	return g.EdgeU(u, v)
}

func (g *Undirected[W]) EdgeU(u, v groph.VIdx) W {
	return g.ws[uIdx(u, v)]
}

func (g *Undirected[W]) NotEdge() W { return g.noe }

func (g *Undirected[W]) IsEdge(w W) bool { return w != g.noe }

func (g *Undirected[W]) SetEdge(u, v groph.VIdx, w W) {
	// gimpl.USetEdge(g, u, v, w) – speeds up with local impl
	if u < v {
		g.SetEdgeU(v, u, w)
	} else {
		g.SetEdgeU(u, v, w)
	}
}

func (g *Undirected[W]) SetEdgeU(u, v groph.VIdx, w W) { g.ws[uIdx(u, v)] = w }

func (g *Undirected[W]) DelEdge(u, v groph.VIdx) {
	// gimpl.USetEdge(g, u, v, g.noe) – speeds up with local impl
	if u < v {
		g.SetEdgeU(v, u, g.noe)
	} else {
		g.SetEdgeU(u, v, g.noe)
	}
}

func (g *Undirected[W]) DelEdgeU(u, v groph.VIdx) { g.SetEdgeU(u, v, g.noe) }

func (g *Undirected[W]) Size() int {
	// TODO Simply iterate over g.ws and use g.IsEgde
	return gimpl.USize[W](g)
}

func (g *Undirected[W]) EachEdge(onEdge groph.VisitEdge[W]) error {
	return gimpl.UEachEdge(g, onEdge)
}

func (g *Undirected[W]) Degree(v groph.VIdx) int {
	return gimpl.UDegree[W](g, v)
}

func (g *Undirected[W]) EachAdjacent(of groph.VIdx, onNeighbour groph.VisitVertex) error {
	return gimpl.UEachAdjacent[W](g, of, onNeighbour)
}

// nSum computes the sum of the n 1st integers, i.e. 1+2+3+…+n
func nSum(n int) int { return n * (n + 1) / 2 }

func nSumRev(n int) float64 {
	r := math.Sqrt(0.25 + 2*float64(n))
	return r - 0.5
}

// uIdx computes the index into the weight slice of an undirected matrix
func uIdx(i, j groph.VIdx) groph.VIdx { return nSum(i) + j }

func uOrdFromLen(sliceLen int) (order int, err error) {
	order = int(nSumRev(sliceLen))
	if nSum(order) != sliceLen {
		return order, errors.New("weights slice len is not sum(1,…,n)")
	}
	return order, nil
}
