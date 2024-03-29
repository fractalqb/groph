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

package graphs

import (
	"math"
	"testing"

	"git.fractalqb.de/fractalqb/groph/gimpl"
)

func TestInForest(t *testing.T) {
	t.Run("cleared", func(t *testing.T) {
		f := NewInForest(11, math.MinInt)
		gimpl.TestDCleared[int](t, f, "new forest")
	})
	t.Run("set-del", func(t *testing.T) {
		f := NewInForest(11, math.MinInt)
		gimpl.SetDelTest[int]{
			Probe:    4,
			EqWeight: func(a, b int) bool { return a == b },
			LazySize: true,
		}.Directed(t, f)
	})
}
