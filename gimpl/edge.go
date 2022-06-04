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

func UEdge[W any, G groph.RUndirected[W]](g G, u, v groph.VIdx) W {
	if u < v {
		return g.EdgeU(v, u)
	}
	return g.EdgeU(u, v)
}

func USetEdge[W any, G groph.WUndirected[W]](g G, u, v groph.VIdx, w W) {
	if u < v {
		g.SetEdgeU(v, u, w)
	} else {
		g.SetEdgeU(u, v, w)
	}
}

func UEachEdge[W any, G groph.RUndirected[W]](g G, onEdge groph.VisitEdgeW[W]) error {
	ord := g.Order()
	for u := groph.VIdx(0); u < ord; u++ {
		for v := groph.VIdx(0); v <= u; v++ {
			if w := g.EdgeU(u, v); g.IsEdge(w) {
				if err := onEdge(u, v, w); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func DEachEdge[W any, G groph.RDirected[W]](g G, onEdge groph.VisitEdgeW[W]) error {
	ord := g.Order()
	for u := groph.VIdx(0); u < ord; u++ {
		for v := groph.VIdx(0); v < ord; v++ {
			if w := g.Edge(u, v); g.IsEdge(w) {
				if err := onEdge(u, v, w); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
