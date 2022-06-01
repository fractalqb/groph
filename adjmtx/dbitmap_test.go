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
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/gimpl"
)

var _ groph.WGraph[bool] = (*DBitmap)(nil)

func TestDBitmap(t *testing.T) {
	m := NewDBitmap(3, nil)
	gimpl.TestDSetDel[bool](t, m, true, func(a, b bool) bool { return a == b })
}

func BenchmarkDBitmap(b *testing.B) {
	m := NewDBitmap(gimpl.SetDelSize, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := true
		if n&1 == 0 {
			w = false
		}
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
