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
	"golang.org/x/exp/constraints"

	"git.fractalqb.de/fractalqb/groph"
)

func FloydWarshallD[W constraints.Ordered, G groph.WDirected[W]](g G) {
	vno := g.Order()
	for k := 0; k < vno; k++ {
		for i := 0; i < vno; i++ {
			for j := 0; j < vno; j++ {
				ds := g.Edge(i, k)
				if !g.IsEdge(ds) {
					continue
				}
				if tmp := g.Edge(k, j); g.IsEdge(tmp) {
					ds += tmp
				} else {
					continue
				}
				if d := g.Edge(i, j); !g.IsEdge(d) || ds < d {
					g.SetEdge(i, j, ds)
				}
			}
		}
	}
}

func FloydWarshallU[W constraints.Ordered, G groph.WUndirected[W]](g G) {
	vno := g.Order()
	for k := 0; k < vno; k++ {
		for i := 0; i < vno; i++ {
			for j := 0; j < i; j++ {
				ds := g.Edge(i, k)
				if !g.IsEdge(ds) {
					continue
				}
				if tmp := g.Edge(k, j); g.IsEdge(tmp) {
					ds += tmp
				} else {
					continue
				}
				if d := g.EdgeU(i, j); !g.IsEdge(d) || ds < d {
					g.SetEdgeU(i, j, ds)
				}
			}
		}
	}
}
