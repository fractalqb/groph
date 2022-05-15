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
	"math"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/gtests"
)

var _ groph.WGraph[int] = (*Directed[int])(nil)

func TestDirected(t *testing.T) {
	const noEdge = math.MinInt32
	g := NewDirected[int32](gtests.SetDelSize, noEdge, nil)
	gtests.TestDCleared[int32](t, g)
	gtests.TestDSetDel[int32](t, g, 1, func(a, b int32) bool { return a == b })
}

func BenchmarkDirected(b *testing.B) {
	const noEdge = math.MinInt32
	m := NewDirected[int32](gtests.SetDelSize, noEdge, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := int32(n)
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				r := m.Edge(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}
