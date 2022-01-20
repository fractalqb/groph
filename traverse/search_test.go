package traverse

import (
	"fmt"
	"testing"

	"git.fractalqb.de/fractalqb/groph/adjmatrix"

	"git.fractalqb.de/fractalqb/groph"
)

func e(u, v groph.VIdx) groph.Edge { return groph.Edge{U: u, V: v} }

func TestSearch_emptyGraphs(t *testing.T) {
	var visited bool
	visit := func(_, _ groph.VIdx, _ int, _ int) bool {
		visited = true
		return false
	}
	s := NewSearch(nil)
	g := adjmatrix.NewDBool(0, nil)
	test := func(name string, call func(bool, VisitInCluster) bool) {
		t.Run(name, func(t *testing.T) {
			s.Reset(g)
			visited = false
			if call(true, visit) {
				t.Error("visit was stopped")
			}
			if visited {
				t.Error("visited something in empty graph")
			}
		})
	}
	test("InDepth1st", s.InDepth1st)
	test("InBreadth1st", s.InBreadth1st)
	test("AdjDepth1st", s.AdjDepth1st)
	test("AdjBreadth1st", s.AdjBreadth1st)
	test("OutDepth1st", s.OutDepth1st)
	test("OutBreadth1st", s.OutBreadth1st)
}

func ExampleSearch_AdjDepth1stAt() {
	g := adjmatrix.NewUBool(7, nil)
	groph.Set(g, true,
		e(0, 1), e(0, 2), e(0, 3),
		e(1, 4), e(1, 5),
		e(2, 5),
		e(3, 6),
	)
	search := NewSearch(g)
	search.SortBy = VIdxOrder
	hits, _ := search.AdjDepth1stAt(0, func(u, v groph.VIdx, vHits int) bool {
		if vHits > 0 {
			fmt.Printf("[%d %d]", u, v)
		} else {
			fmt.Printf("(%d %d)", u, v)
		}
		return false
	})
	fmt.Println("\nhits:", hits)
	// Output:
	// (-1 0)(0 1)(1 4)(1 5)(5 2)[2 0][0 2](0 3)(3 6)
	// hits: 7
}

func ExampleSearch_AdjBreadth1stAt() {
	g := adjmatrix.NewUBool(7, nil)
	groph.Set(g, true,
		e(0, 1), e(0, 2), e(0, 3),
		e(1, 4), e(1, 5),
		e(2, 5),
		e(3, 6),
	)
	hits, _ := NewSearch(g).AdjBreadth1stAt(0, func(u, v groph.VIdx, vHits int) bool {
		if vHits > 0 {
			fmt.Printf("[%d %d]", u, v)
		} else {
			fmt.Printf("(%d %d)", u, v)
		}
		return false
	})
	fmt.Println("\nhits:", hits)
	// Output:
	// (-1 0)(0 1)(0 2)(0 3)(1 4)(1 5)[2 5](3 6)[5 2]
	// hits: 7
}

func TestSearch_Depth1st_dir_find_clusters(t *testing.T) {
	g := adjmatrix.NewDBitmap(2, nil)
	search := NewSearch(g)
	test := func() {
		search.Reset(g)
		hits := 0
		circ := 0
		clusters := make(map[int]bool)
		search.AdjDepth1st(false, func(_, _ groph.VIdx, h int, cl int) bool {
			if h > 0 {
				circ++
			} else {
				hits++
			}
			clusters[cl] = true
			return false
		})
		if hits != 2 {
			t.Errorf("unexpected number of hits: %d, want 2", hits)
		}
		if circ != 0 {
			t.Errorf("unexpected number of circles: %d, want 0", circ)
		}
		if len(clusters) != 1 {
			t.Errorf("found wrong clusters: %v", clusters)
		} else if _, ok := clusters[0]; !ok {
			t.Errorf("found wrong clusters: %v", clusters)
		}
	}
	g.SetEdge(0, 1, true)
	test()
	g.SetWeight(0, 1, false)
	g.SetEdge(1, 0, true)
	test()
}

func TestSearch_Breadth1st_dir_find_clusters(t *testing.T) {
	g := adjmatrix.NewDBitmap(2, nil)
	search := NewSearch(g)
	test := func() {
		search.Reset(g)
		hits, circ := 0, 0
		clusters := make(map[int]bool)
		search.AdjBreadth1st(false, func(_, _ groph.VIdx, h int, cl int) bool {
			if h > 0 {
				circ++
			} else {
				hits++
			}
			clusters[cl] = true
			return false
		})
		if hits != 2 {
			t.Errorf("unexpected number of hits: %d, want 2", hits)
		}
		if circ != 1 {
			t.Errorf("unexpected number of circles: %d, want 1", circ)
		}
		if len(clusters) != 1 {
			t.Errorf("found wrong clusters: %v", clusters)
		} else if _, ok := clusters[0]; !ok {
			t.Errorf("found wrong clusters: %v", clusters)
		}
	}
	g.SetEdge(0, 1, true)
	test()
	g.SetWeight(0, 1, false)
	g.SetEdge(1, 0, true)
	test()
}

func ExampleSearch_ugraph() {
	g := adjmatrix.NewUBool(4, nil)
	groph.Set(g, true, e(0, 1), e(1, 2), e(2, 3), e(3, 0), e(0, 2))
	search := NewSearch(g)
	search.SortBy = VIdxOrder
	search.AdjDepth1stAt(0, func(u, v int, vHits int) bool {
		_, fin := search.State(v)
		fmt.Println(u, v, vHits, fin)
		return false
	})
	// Output:
	// -1 0 0 false
	// 0 1 0 false
	// 1 2 0 false
	// 2 0 1 false
	// 2 3 0 false
	// 3 0 2 false
	// 0 2 1 true
	// 0 3 1 true

}

func ExampleSearch_ugd1st_loop_not_parent() {
	g := adjmatrix.NewUBool(2, nil)
	groph.Set(g, true, e(0, 1), e(1, 1))
	search := NewSearch(g)
	search.SortBy = VIdxOrder
	search.AdjDepth1stAt(0, func(u, v int, vHits int) bool {
		_, fin := search.State(v)
		fmt.Println(u, v, vHits, fin)
		return false
	})
	// Output:
	// -1 0 0 false
	// 0 1 0 false
	// 1 1 1 false
}

func ExampleSearch_ugb1st_loop_not_parent() {
	g := adjmatrix.NewUBool(2, nil)
	groph.Set(g, true, e(0, 1), e(1, 1))
	search := NewSearch(g)
	search.SortBy = VIdxOrder
	search.AdjBreadth1stAt(0, func(u, v int, vHits int) bool {
		_, fin := search.State(v)
		fmt.Println(u, v, vHits, fin)
		return false
	})
	// Output:
	// -1 0 0 false
	// 0 1 0 false
	// 1 1 1 false
}
