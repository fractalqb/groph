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

package internal

import (
	"reflect"
	"testing"
)

func assertLenCap(t *testing.T, s interface{}, l, c int, hint string) {
	sval := reflect.ValueOf(s)
	if sval.Cap() != c {
		t.Errorf("%s: cap expected %d, got %d", hint, c, sval.Cap())
	}
	if sval.Len() != l {
		t.Errorf("%s: len expected %d, got %d", hint, l, sval.Len())
	}
}

func TestSlice(t *testing.T) {
	s, _ := Slice[int](nil, 8)
	assertLenCap(t, s, 8, 8, "from nil")
	s, _ = Slice(s, 5)
	assertLenCap(t, s, 5, 8, "from nil")
	s, _ = Slice(s, 7)
	assertLenCap(t, s, 7, 8, "from nil")
	s, _ = Slice(s, 12)
	assertLenCap(t, s, 12, 12, "from nil")
}
