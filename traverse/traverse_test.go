package traverse

import (
	"fmt"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
)

func ExampleSearch_Depth1stAt() {
	g := groph.NewAdjMxUbool(7, nil)
	type E = groph.Edge
	groph.Set(g, true,
		E{0, 1}, E{0, 2}, E{0, 3},
		E{1, 4}, E{1, 5},
		E{2, 5},
		E{3, 6},
	)
	t := NewSearch(g)
	t.SortBy = func(u, v1, v2 groph.VIdx) bool { return v1 < v2 }
	hits, _ := t.Depth1stAt(0, func(v groph.VIdx) bool {
		fmt.Printf(" %d", v)
		return false
	})
	fmt.Println("\nhits:", hits)
	// Output:
	// 0 1 4 5 2 3 6
	// hits: 7
}

func ExampleSearch_Breadth1stAt() {
	g := groph.NewAdjMxUbool(7, nil)
	type E = groph.Edge
	groph.Set(g, true,
		E{0, 1}, E{0, 2}, E{0, 3},
		E{1, 4}, E{1, 5},
		E{2, 5},
		E{3, 6},
	)
	hits, _ := NewSearch(g).Breadth1stAt(0, func(v groph.VIdx) bool {
		fmt.Printf(" %d", v)
		return false
	})
	fmt.Println("\nhits:", hits)
	// Output:
	// 0 1 2 3 4 5 6
	// hits: 7
}

func TestSearch_Depth1st(t *testing.T) {
	g := groph.NewAdjMxUbool(2, nil)
	type E = groph.Edge
	groph.Set(g, true, E{0, 1}, E{1, 1})
	search := NewSearch(g)
	stopped := search.Depth1st(false,
		func(v groph.VIdx, c int) bool { return false })
	if stopped {
		t.Fatal("search was stopped unexpectedly")
	}
	for i := groph.V0; i < g.Order(); i++ {
		if h := search.Hits(i); h > 1 {
			t.Errorf("%d has hist %d > 1", i, h)
		}
	}
}

func TestSearch_Breadth1st(t *testing.T) {
	g := groph.NewAdjMxUbool(2, nil)
	type E = groph.Edge
	groph.Set(g, true, E{0, 1}, E{1, 1})
	search := NewSearch(g)
	stopped := search.Breadth1st(false,
		func(v groph.VIdx, c int) bool { return false })
	if stopped {
		t.Fatal("search was stopped unexpectedly")
	}
	for i := groph.V0; i < g.Order(); i++ {
		if h := search.Hits(i); h > 1 {
			t.Errorf("%d has hist %d > 1", i, h)
		}
	}
}
