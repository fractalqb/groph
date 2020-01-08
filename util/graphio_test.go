package util

import (
	"os"
	"strings"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
)

func ExampleWriteSparse() {
	g := groph.NewSoMDi32(4, nil)
	g.SetEdge(0, 1, 1)
	g.SetEdge(1, 2, 2)
	g.SetEdge(2, 3, 3)
	g.SetEdge(3, 0, 0)
	g.SetEdge(0, 2, -1)
	WriteSparse(os.Stdout, g)
	// Output:
	// groph directed edges order=4
	// 0 1 1
	// 0 2 -1
	// 1 2 2
	// 2 3 3
	// 3 0 0
}

func TestRead_sparse(t *testing.T) {
	buf := strings.NewReader(`groph directed edges order=4
	0 1 1
	0 2 -1
	1 2 2
	2 3 3
	3 0 0`)
	g := groph.NewAdjMxDi32(0, groph.I32Del, nil)
	err := ReadGraph(g, buf, ParseI32)
	if err != nil {
		t.Fatal("failed to read graph:", err)
	}
	if g.Order() != 4 {
		t.Fatalf("read wrong order %d, want 4", g.Order())
	}
	expect := map[groph.Edge]int32{
		groph.Edge{U: 0, V: 1}: 1,
		groph.Edge{U: 0, V: 2}: -1,
		groph.Edge{U: 1, V: 2}: 2,
		groph.Edge{U: 2, V: 3}: 3,
		groph.Edge{U: 3, V: 0}: 0,
	}
	for i := 0; i < g.Order(); i++ {
		for j := 0; j < g.Order(); j++ {
			if w, ok := expect[groph.Edge{U: i, V: j}]; ok {
				if gw, ok := g.Edge(i, j); ok {
					if gw != w {
						t.Errorf("wrong weight %d of (%d,%d), want %d", gw, i, j, w)
					}
				} else {
					t.Errorf("missing edge (%d,%d)", i, j)
				}
			} else if gw, ok := g.Edge(i, j); ok {
				t.Errorf("must not have edge (%d,%d):%d", i, j, gw)
			}
		}
	}
}
