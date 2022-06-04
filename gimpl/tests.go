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
	"testing"

	"git.fractalqb.de/fractalqb/groph"
)

const SetDelSize = 11

func TestDCleared[W any](t *testing.T, g groph.RGraph[W], what string) {
	vno := g.Order()
	for ri := 0; ri < vno; ri++ {
		for rj := 0; rj < vno; rj++ {
			if w := g.Edge(ri, rj); g.IsEdge(w) {
				t.Errorf("%s: read non-cleared value [%v] @%d,%d",
					what,
					w,
					ri, rj)
			}
		}
	}
	if s := g.Size(); s != 0 {
		t.Errorf("%s: cleared graph has non-zero size %d", what, s)
	}
}

func TestUCleared[W any](t *testing.T, g groph.RUndirected[W], what string) {
	vno := g.Order()
	for ri := 0; ri < vno; ri++ {
		for rj := ri; rj < vno; rj++ {
			if w := g.Edge(ri, rj); g.IsEdge(w) {
				t.Errorf("%s: read non-cleared value [%v] @%d,%d",
					what,
					w,
					ri, rj)
			}
		}
	}
	if s := g.Size(); s != 0 {
		t.Errorf("%s: cleared graph has non-zero size %d", what, s)
	}
}

// TODO Break down tests into t.Run() calls
type SetDelTest[W any] struct {
	Probe    W
	EqWeight func(a, b W) bool
	LazySize bool
}

func (tst SetDelTest[W]) Directed(t *testing.T, g groph.WGraph[W]) {
	if _, ok := g.(groph.RUndirected[W]); ok {
		t.Fatal("graph is not directed")
	}
	vno := g.Order()
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj < vno; wj++ {
			g.DelEdge(wi, wj)
		}
	}
	TestDCleared[W](t, g, "del all edges")
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj < vno; wj++ {
			g.SetEdge(wi, wj, tst.Probe)
			if s := g.Size(); s != 1 {
				t.Errorf("Setting single edge has wrong size: %d", s)
			}
			for ri := 0; ri < vno; ri++ {
				for rj := 0; rj < vno; rj++ {
					r := g.Edge(ri, rj)
					if ri == wi && rj == wj {
						if !tst.EqWeight(r, tst.Probe) {
							t.Fatalf("read wrong value [%v] (expect [%v]) @%d,%d",
								r, tst.Probe,
								wi, wj)
						}
					} else if g.IsEdge(r) {
						t.Fatalf("read non-cleared value [%v] @%d,%d after write @%d,%d",
							tst.Probe,
							ri, rj,
							wi, wj)
					}
				}
			}
			g.DelEdge(wi, wj)
			if s := g.Size(); s != 0 {
				t.Errorf("Deleting single edge has wrong size: %d", s)
			}
		}
	}
	TestDCleared[W](t, g, "2x flip each edge")
	size := 0
	for i := 0; i < vno; i++ {
		for j := 0; j < vno; j++ {
			g.SetEdge(i, j, tst.Probe)
			size++
			if s := g.Size(); !tst.LazySize && s != size {
				t.Errorf("Unexpected size %d when filling graph, want %d", s, size)
			}
		}
	}
	g.Reset(g.Order())
	if g.Order() != vno {
		t.Fatal("Reset changed graph size")
	}
	TestDCleared[W](t, g, "Reset()")
}

func (tst SetDelTest[W]) Undirected(t *testing.T, g groph.WUndirected[W]) {
	vno := g.Order()
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj <= wi; wj++ {
			g.DelEdge(wi, wj)
		}
	}
	TestUCleared[W](t, g, "del all edges")
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj <= wi; wj++ {
			g.SetEdge(wi, wj, tst.Probe)
			if s := g.Size(); s != 1 {
				t.Errorf("Setting single edge has wrong size: %d", s)
			}
			for ri := 0; ri < vno; ri++ {
				for rj := 0; rj < vno; rj++ {
					expectSet := (ri == wi && rj == wj) || (ri == wj && rj == wi)
					r := g.Edge(ri, rj)
					if expectSet {
						if !tst.EqWeight(r, tst.Probe) {
							t.Fatalf("read wrong value [%v] (expect [%v]) @%d,%d",
								r, tst.Probe,
								ri, rj)
						}
					} else if g.IsEdge(r) {
						t.Fatalf("read non-cleared value [%v] @%d,%d after write @%d,%d",
							tst.Probe,
							ri, rj,
							wi, wj)
					}
					r = g.Edge(rj, ri)
					if expectSet {
						if !tst.EqWeight(r, tst.Probe) {
							t.Fatalf("read wrong value [%v] (expect [%v]) @%d,%d",
								r, tst.Probe,
								ri, rj)
						}
					} else if g.IsEdge(r) {
						t.Fatalf("read non-cleared value [%v] @%d,%d after write @%d,%d",
							tst.Probe,
							ri, rj,
							wi, wj)
					}
				}
			}
			g.DelEdge(wi, wj)
			if s := g.Size(); s != 0 {
				t.Errorf("Deleting single edge has wrong size: %d", s)
			}
		}
	}
	TestUCleared[W](t, g, "2x flip each edge")
	size := 0
	for i := 0; i < vno; i++ {
		for j := 0; j <= i; j++ {
			g.SetEdge(i, j, tst.Probe)
			size++
			if s := g.Size(); !tst.LazySize && s != size {
				t.Errorf("Unexpected size %d when filling graph, want %d", s, size)
			}
		}
	}
	g.Reset(g.Order())
	if g.Order() != vno {
		t.Fatal("Reset changed graph size")
	}
	TestUCleared[W](t, g, "Reset()")
}
