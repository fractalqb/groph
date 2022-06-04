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

import (
	"fmt"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmtx"
)

func TestBfs_queue(t *testing.T) {
	var s bfs
	for i := 0; i < 10; i++ {
		s.enq(edge{-1, i})
	}
	rd := 0
	for i := 0; i < 7; i++ {
		e := s.deq()
		if e.v != rd {
			t.Fatalf("deq: expect %d, got %d", rd, e.v)
		}
		rd++
	}
	for i := 10; i < 15; i++ {
		s.enq(edge{-1, i})
	}
	for s.qlen() > 0 {
		e := s.deq()
		if e.v != rd {
			t.Fatalf("deq: expect %d, got %d", rd, e.v)
		}
		rd++
	}
	if rd != 15 {
		t.Fatalf("Last read value is %d, want 14", rd-1)
	}
}

func ExampleUndirectedBFS() {
	g := adjmtx.NewUndirected(7, false, nil)
	groph.Set[bool](g, true,
		0, 1, 0, 2, 0, 3,
		1, 4, 1, 5,
		2, 5,
		3, 6,
	)
	bfs := *NewUndirectedBFS[bool](g)
	for s := bfs.NextStart(); s >= 0; s = bfs.NextStart() {
		bfs.Start(s, func(u, v groph.VIdx) error {
			fmt.Printf(" %d->%d", u, v)
			return nil
		})
		fmt.Println()
	}
	// Output:
	// -1->0 0->1 0->2 0->3 1->4 1->5 3->6
}
