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

func TestDCleared[W any](t *testing.T, g groph.RGraph[W]) {
	vno := g.Order()
	for ri := 0; ri < vno; ri++ {
		for rj := 0; rj < vno; rj++ {
			if w := g.Edge(ri, rj); g.IsEdge(w) {
				t.Errorf("read non-cleared value [%v] @%d,%d after clear",
					w,
					ri, rj)
			}
		}
	}
}

func TestDSetDel[W any](t *testing.T, g groph.WGraph[W], probe W, eq func(a, b W) bool) {
	if _, ok := g.(groph.RUndirected[W]); ok {
		t.Fatal("graph is not directed")
	}
	vno := g.Order()
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj < vno; wj++ {
			g.DelEdge(wi, wj)
		}
	}
	TestDCleared[W](t, g)
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj < vno; wj++ {
			g.SetEdge(wi, wj, probe)
			for ri := 0; ri < vno; ri++ {
				for rj := 0; rj < vno; rj++ {
					r := g.Edge(ri, rj)
					if ri == wi && rj == wj {
						if !eq(r, probe) {
							t.Fatalf("read wrong value [%v] (expect [%v]) @%d,%d",
								r, probe,
								wi, wj)
						}
					} else if g.IsEdge(r) {
						t.Fatalf("read non-cleared value [%v] @%d,%d after write @%d,%d",
							probe,
							ri, rj,
							wi, wj)
					}
				}
			}
			g.DelEdge(wi, wj)
		}
	}
	for i := 0; i < vno; i++ {
		for j := 0; j < vno; j++ {
			g.SetEdge(i, j, probe)
		}
	}
	g.Reset(g.Order())
	if g.Order() != vno {
		t.Fatal("Reset changed graph size")
	}
	for i := 0; i < vno; i++ {
		for j := 0; j < vno; j++ {
			if g.IsEdge(g.Edge(i, j)) {
				t.Fatalf("Reset() did not clear the graph at (%d,%d)", i, j)
			}
		}
	}
}

func TestUCleared[W any](t *testing.T, g groph.RUndirected[W]) {
	vno := g.Order()
	for ri := 0; ri < vno; ri++ {
		for rj := ri; rj < vno; rj++ {
			if w := g.Edge(ri, rj); g.IsEdge(w) {
				t.Errorf("read non-cleared value [%v] @%d,%d after clear",
					w,
					ri, rj)
			}
		}
	}
}

func TestUSetDel[W any](t *testing.T, g groph.WUndirected[W], probe W, eq func(a, b W) bool) {
	vno := g.Order()
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj <= wi; wj++ {
			g.DelEdge(wi, wj)
		}
	}
	TestUCleared[W](t, g)
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj <= wi; wj++ {
			g.SetEdge(wi, wj, probe)
			for ri := 0; ri < vno; ri++ {
				for rj := 0; rj < vno; rj++ {
					expectSet := (ri == wi && rj == wj) || (ri == wj && rj == wi)
					r := g.Edge(ri, rj)
					if expectSet {
						if !eq(r, probe) {
							t.Fatalf("read wrong value [%v] (expect [%v]) @%d,%d",
								r, probe,
								ri, rj)
						}
					} else if g.IsEdge(r) {
						t.Fatalf("read non-cleared value [%v] @%d,%d after write @%d,%d",
							probe,
							ri, rj,
							wi, wj)
					}
					r = g.Edge(rj, ri)
					if expectSet {
						if !eq(r, probe) {
							t.Fatalf("read wrong value [%v] (expect [%v]) @%d,%d",
								r, probe,
								ri, rj)
						}
					} else if g.IsEdge(r) {
						t.Fatalf("read non-cleared value [%v] @%d,%d after write @%d,%d",
							probe,
							ri, rj,
							wi, wj)
					}
				}
			}
			g.DelEdge(wi, wj)
		}
	}
	for i := 0; i < vno; i++ {
		for j := 0; j <= i; j++ {
			g.SetEdge(i, j, probe)
		}
	}
	g.Reset(g.Order())
	if g.Order() != vno {
		t.Fatal("Reset changed graph size")
	}
	for i := 0; i < vno; i++ {
		for j := 0; j <= i; j++ {
			if g.IsEdge(g.Edge(i, j)) {
				t.Fatalf("Reset() did not clear the graph at (%d,%d)", i, j)
			}
		}
	}
}
