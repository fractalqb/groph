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

package graphs

import (
	"math"

	"git.fractalqb.de/fractalqb/groph"
	"golang.org/x/exp/constraints"
)

type Distancer interface {
	Distance(to Distancer) float64
}

type Point[W constraints.Float] []W

func (p Point[W]) Distance(to Distancer) (d float64) {
	if q, ok := to.(Point[W]); ok {
		l := len(p)
		if len(q) < l {
			l = len(q)
		}
		for i := 0; i < l; i++ {
			tmp := p[i] - q[i]
			d += float64(tmp * tmp)
		}
		return math.Sqrt(d)
	}
	return math.NaN()
}

type Euclidean []Distancer

var _ groph.RUndirected[float64] = Euclidean{}

func (g Euclidean) Order() int { return len(g) }

func (g Euclidean) Edge(u, v groph.VIdx) (weight float64) {
	return g[u].Distance(g[v])
}

func (g Euclidean) IsEdge(weight float64) bool {
	return !math.IsNaN(weight)
}

func (g Euclidean) NotEdge() float64 { return math.NaN() }

func (g Euclidean) Size() int {
	o := g.Order()
	return o * o
}

func (g Euclidean) EachEdge(onEdge groph.VisitEdgeW[float64]) error {
	ord := g.Order()
	for i := groph.VIdx(0); i < ord; i++ {
		for j := 0; j < ord; j++ {
			err := onEdge(i, j, g.Edge(i, j))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g Euclidean) EdgeU(u, v groph.VIdx) (weight float64) { return g.Edge(u, v) }

func (g Euclidean) Degree(v groph.VIdx) int { return g.Order() }

func (g Euclidean) EachAdjacent(of groph.VIdx, onNeighbour groph.VisitVertex) error {
	ord := g.Order()
	for i := groph.VIdx(0); i < ord; i++ {
		if err := onNeighbour(i); err != nil {
			return err
		}
	}
	return nil
}
