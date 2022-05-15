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

package gview

import (
	"fmt"

	"git.fractalqb.de/fractalqb/groph/adjmtx"
)

func ExampleDirectedSub() {
	g, _ := adjmtx.AsDirected(nil, 0,
		0, 8, 0, 1,
		0, 0, 1, 0,
		4, 0, 0, 0,
		0, 2, 9, 0,
	)
	s := NewDirectedSub[int](g, 1, 3)
	s.EachEdge(func(u, v, w int) bool {
		fmt.Printf("%d -> %d: %d\n", u, v, w)
		return false
	})
	// Output:
	// 1 -> 0: 2
}
