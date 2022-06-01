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

package adjls

import (
	"math"
	"testing"

	"git.fractalqb.de/fractalqb/groph/gimpl"
)

func TestDirected(t *testing.T) {
	g := NewDirected(11, math.Inf(1), nil)
	gimpl.TestDCleared[float64](t, g)
	gimpl.TestDSetDel[float64](t, g, 4711.0, func(a, b float64) bool { return a == b })
}
