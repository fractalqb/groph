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

func GreedyTSP[W constraints.Ordered](g groph.RGraph[W]) (path []groph.VIdx, plen W) {
	ord := g.Order()
	switch ord {
	case 0:
		var zero W
		return nil, zero
	case 1:
		var zero W
		return []groph.VIdx{0}, zero
	}
	l := ord - 1
	path = make([]groph.VIdx, ord)
	path[l] = l
	plen = g.Edge(l, 0)
	for k := 0; k < l; k++ {
		path[k] = k
		plen += g.Edge(k, k+1)
	}
	perm := make([]groph.VIdx, l)
	copy(perm, path)
	c := make([]groph.VIdx, l)
	i := 0
	for i < l {
		if c[i] < i {
			if i&1 == 0 {
				perm[0], perm[i] = perm[i], perm[0]
			} else {
				perm[c[i]], perm[i] = perm[i], perm[c[i]]
			}
			curl := g.Edge(l, perm[0])
			curl += g.Edge(perm[l-1], l)
			for i := 0; i+1 < l; i++ {
				curl += g.Edge(perm[i], perm[i+1])
			}
			if curl < plen {
				copy(path[:l], perm)
				plen = curl
			}
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return path, plen
}
