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
)

func DOutDegree[W any, G groph.RDirected[W]](g G, v groph.VIdx) (d int) {
	ord := g.Order()
	for u := groph.VIdx(0); u < ord; u++ {
		if g.IsEdge(g.Edge(v, u)) {
			d++
		}
	}
	return d
}

func DEachOut[W any, G groph.RDirected[W]](g G, from groph.VIdx, onDest groph.VisitVertex) error {
	ord := g.Order()
	for to := groph.VIdx(0); to < ord; to++ {
		if g.IsEdge(g.Edge(from, to)) {
			if err := onDest(to); err != nil {
				return err
			}
		}
	}
	return nil
}

func DInDegree[W any, G groph.RDirected[W]](g G, v groph.VIdx) (d int) {
	ord := g.Order()
	for u := 0; u < ord; u++ {
		if g.IsEdge(g.Edge(u, v)) {
			d++
		}
	}
	return d
}

func DEachIn[W any, G groph.RDirected[W]](g G, to groph.VIdx, onSource groph.VisitVertex) error {
	ord := g.Order()
	for from := groph.VIdx(0); from < ord; from++ {
		if g.IsEdge(g.Edge(from, to)) {
			if err := onSource(from); err != nil {
				return err
			}
		}
	}
	return nil
}

func DEachRoot[W any, G groph.RDirected[W]](g G, onRoot groph.VisitVertex) error {
	ord := g.Order()
	for i := groph.VIdx(0); i < ord; i++ {
		if g.InDegree(i) == 0 {
			if err := onRoot(i); err != nil {
				return err
			}
		}
	}
	return nil
}

func DEachLeaf[W any, G groph.RDirected[W]](g G, onLeaf groph.VisitVertex) error {
	ord := g.Order()
	for i := groph.VIdx(0); i < ord; i++ {
		if g.OutDegree(i) == 0 {
			if err := onLeaf(i); err != nil {
				return err
			}
		}
	}
	return nil
}

func UDegree[W any](g groph.RUndirected[W], v groph.VIdx) (d int) {
	ord := g.Order()
	if g.IsEdge(g.EdgeU(v, v)) {
		d = 2
	}
	for i := groph.VIdx(0); i < v; i++ {
		if g.IsEdge(g.EdgeU(v, i)) {
			d++
		}
	}
	for i := v + 1; i < ord; i++ {
		if g.IsEdge(g.EdgeU(i, v)) {
			d++
		}
	}
	return d
}

func UEachAdjacent[W any, G groph.RUndirected[W]](g G, of groph.VIdx, onNeighbour groph.VisitVertex) error {
	n, ord := groph.VIdx(0), g.Order()
	for n < of {
		if g.IsEdge(g.EdgeU(of, n)) {
			if err := onNeighbour(n); err != nil {
				return err
			}
		}
		n++
	}
	for n < ord { // n >= of
		if g.IsEdge(g.EdgeU(n, of)) {
			if err := onNeighbour(n); err != nil {
				return err
			}
		}
		n++
	}
	return nil
}
