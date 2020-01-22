package traverse

import (
	"fmt"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmatrix"
)

func TestHasCycle_dgraph(t *testing.T) {
	g := adjmatrix.NewDBool(3, nil)
	groph.Set(g, true, e(0, 1), e(1, 2))
	if HasCycle(g, nil) {
		t.Error("unexpected cycle detected")
	}
	g.SetEdge(2, 0, true)
	if !HasCycle(g, nil) {
		t.Error("no cycle in cyclic graph")
	}
}

func TestHasCycle_ugraph(t *testing.T) {
	g := adjmatrix.NewUBool(3, nil)
	groph.Set(g, true, e(0, 1), e(1, 2))
	if HasCycle(g, nil) {
		t.Error("unexpected cycle detected")
	}
	g.SetEdge(2, 0, true)
	if !HasCycle(g, nil) {
		t.Error("no cycle in cyclic graph")
	}
}

func ExampleCircle() {
	g := adjmatrix.NewUBool(4, nil)
	groph.Set(g, true, e(0, 1), e(1, 2), e(2, 3), e(3, 0), e(0, 2))
	srch := NewSearch(g)
	srch.SortBy = VIdxOrder
	circ := Circle{
		OnFind: func(vs []groph.VIdx) bool {
			fmt.Println(vs)
			return false
		},
	}
	srch.AdjDepth1st(false, circ.Search)
	// Output:
	// [0 1 2]
	// [0 1 2 3]
}
