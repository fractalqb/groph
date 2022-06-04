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

package groph

func Set[W any](g WGraph[W], w W, vs ...VIdx) {
	for i := 0; i < len(vs); i += 2 {
		j := i + 1
		if j >= len(vs) {
			return
		}
		g.SetEdge(vs[i], vs[j], w)
	}
}

func Del[W any](g WGraph[W], vs ...VIdx) {
	for i := 0; i < len(vs); i++ {
		j := i + 1
		if j >= len(vs) {
			return
		}
		g.DelEdge(vs[i], vs[j])
	}
}

// func Revert[W any](ls []W) {
// 	for i, j := 0, len(ls)-1; i < j; i, j = i+1, j-1 {
// 		ls[i], ls[j] = ls[j], ls[i]
// 	}
// }
