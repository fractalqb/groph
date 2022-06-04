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

func ExampleDirectedDFS() {
	g := adjmtx.NewDirected(7, false, nil)
	groph.Set[bool](g, true,
		0, 1, 0, 2, 0, 3,
		1, 4, 1, 5,
		2, 5,
		3, 6,
	)
	dfs := *NewDirecredDFS[bool](g)
	for s := 0; s >= 0; s = dfs.NextStart() {
		dfs.Forward(s, func(u, v groph.VIdx) error {
			fmt.Printf(" %d->%d", u, v)
			return nil
		})
		fmt.Println()
	}
	// Output:
	// -1->0 0->3 3->6 0->2 2->5 0->1 1->4
}

func ExampleUndirectedDFS() {
	g := adjmtx.NewUndirected(7, false, nil)
	groph.Set[bool](g, true,
		0, 1, 0, 2, 0, 3,
		1, 4, 1, 5,
		2, 5,
		3, 6,
	)
	dfs := *NewUndirecredDFS[bool](g)
	for s := 0; s >= 0; s = dfs.NextStart() {
		dfs.Start(s, func(u, v groph.VIdx) error {
			fmt.Printf(" %d->%d", u, v)
			return nil
		})
		fmt.Println()
	}
	fmt.Println()
	// Output:
	// -1->0 0->3 3->6 0->2 2->5 5->1 1->4
}

func TestDirectedDFS_HasCycle(t *testing.T) {
	g := adjmtx.NewDirected(3, false, nil)
	groph.Set[bool](g, true, 0, 1, 1, 2)
	dfs := *NewDirecredDFS[bool](g)
	if dfs.HasCycle(0) {
		t.Error("unexpected cycle detected")
	}
	g.SetEdge(2, 0, true)
	dfs.Reset(g)
	if !dfs.HasCycle(0) {
		t.Error("no cycle in cyclic graph")
	}
}

func TestUndirectedDFS_HasCycle(t *testing.T) {
	g := adjmtx.NewUndirected(3, false, nil)
	groph.Set[bool](g, true, 0, 1, 1, 2)
	dfs := *NewUndirecredDFS[bool](g)
	if dfs.HasCycle(0) {
		t.Error("unexpected cycle detected")
	}
	g.SetEdge(2, 0, true)
	dfs.Reset(g)
	if !dfs.HasCycle(0) {
		t.Error("no cycle in cyclic graph")
	}
}
