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

func USize[W any, G groph.RUndirected[W]](g G) (res int) {
	ord := g.Order()
	for i := 0; i < ord; i++ {
		res += g.Degree(i)
	}
	return res / 2
}

func DSize[W any, G groph.RDirected[W]](g G) (res int) {
	ord := g.Order()
	for i := 0; i < ord; i++ {
		res += g.InDegree(i)
	}
	return res
}

func DRootCount[W any, G groph.RDirected[W]](g G) (res int) {
	ord := g.Order()
	for i := 0; i < ord; i++ {
		if g.InDegree(i) == 0 {
			res++
		}
	}
	return res
}

func DLeafCount[W any, G groph.RDirected[W]](g G) (res int) {
	ord := g.Order()
	for i := 0; i < ord; i++ {
		if g.OutDegree(i) == 0 {
			res++
		}
	}
	return res
}