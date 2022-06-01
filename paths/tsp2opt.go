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
	"git.fractalqb.de/fractalqb/groph"
	"golang.org/x/exp/constraints"
)

func TwoOptTSP[W constraints.Float](g groph.RGraph[W]) (path []groph.VIdx, plen W) {
	diff2opt := diff2optU[W]
	if _, ok := g.(groph.RDirected[W]); ok {
		diff2opt = diff2optD[W]
	}
	vno := g.Order()
	path = make([]groph.VIdx, vno)
	for i := 0; i+1 < vno; i++ {
		path[i] = i
		plen += g.Edge(i, i+1)
	}
	path[vno-1] = vno - 1
	plen += g.Edge(vno-1, 0)
	for {
		be0, be1 := vno, vno
		var bdiff W
		for e0 := 0; e0 < vno; e0++ {
			for e1 := e0 + 1; e1 < vno; e1++ {
				diff := diff2opt(g, path, e0, e1)
				if diff < bdiff {
					be0, be1 = e0, e1
					bdiff = diff
				}
			}
		}
		if bdiff < 0 {
			apply2opt(path, be0, be1)
			plen += bdiff
		} else {
			break
		}
	}
	return path, plen
}

// d2optU computes the difference in weight sum for a specific 2-opt operation
// that swaps e0 / e1 for undirected graphs.
func diff2optU[W constraints.Float](g groph.RGraph[W], p []groph.VIdx, e0, e1 groph.VIdx) (wdiff W) {
	lenp := groph.VIdx(len(p))
	wdiff = -g.Edge(p[e0], p[e0+1])
	wdiff += g.Edge(p[e0], p[e1])
	if e1+1 == lenp {
		wdiff -= g.Edge(p[e1], p[0])
		wdiff += g.Edge(p[e0+1], p[0])
	} else {
		wdiff -= g.Edge(p[e1], p[e1+1])
		wdiff += g.Edge(p[e0+1], p[e1+1])
	}
	return wdiff
}

// d2optD computes the difference in weight sum for a specific 2-opt operation
// that swaps e0 / e1 for directed graphs.
func diff2optD[W constraints.Float](g groph.RGraph[W], p []groph.VIdx, e0, e1 groph.VIdx) (wdiff W) {
	wdiff = diff2optU(g, p, e0, e1)
	for i := e0 + 1; i < e1; i++ {
		wdiff -= g.Edge(p[i], p[i+1])
		wdiff += g.Edge(p[i+1], p[i])
	}
	return wdiff
}

func apply2opt(p []groph.VIdx, e0, e1 groph.VIdx) {
	e0++
	for e0 < e1 {
		p[e0], p[e1] = p[e1], p[e0]
		e0++
		e1--
	}
}
