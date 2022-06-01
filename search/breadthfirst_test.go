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

package search

import "testing"

func TestBfs_queue(t *testing.T) {
	var s bfs
	for i := 0; i < 10; i++ {
		s.enq(i)
	}
	rd := 0
	for i := 0; i < 7; i++ {
		v := s.deq()
		if v != rd {
			t.Fatalf("deq: expect %d, got %d", rd, v)
		}
		rd++
	}
	for i := 10; i < 15; i++ {
		s.enq(i)
	}
	for s.qlen() > 0 {
		v := s.deq()
		if v != rd {
			t.Fatalf("deq: expect %d, got %d", rd, v)
		}
		rd++
	}
	if rd != 15 {
		t.Fatalf("Last read value is %d, want 14", rd-1)
	}
}
