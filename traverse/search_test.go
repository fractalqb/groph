package traverse

import (
	"fmt"
	"testing"

	"git.fractalqb.de/fractalqb/groph/adjmatrix"

	"git.fractalqb.de/fractalqb/groph"
)

// Just to make go vet happy! I like
// 	type E = groph.Edge
//  E{u, v} better
func e(u, v groph.VIdx) groph.Edge { return groph.Edge{U: u, V: v} }

func ExampleSearch_Depth1stAt() {
	g := adjmatrix.NewUBool(7, nil)
	groph.Set(g, true,
		e(0, 1), e(0, 2), e(0, 3),
		e(1, 4), e(1, 5),
		e(2, 5),
		e(3, 6),
	)
	search := NewSearch(g)
	search.SortBy = VIdxOrder
	hits, _ := search.AdjDepth1stAt(0, func(u, v groph.VIdx, c bool) bool {
		if c {
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

func ExampleSearch_Breadth1stAt() {
	g := adjmatrix.NewUBool(7, nil)
	groph.Set(g, true,
		e(0, 1), e(0, 2), e(0, 3),
		e(1, 4), e(1, 5),
		e(2, 5),
		e(3, 6),
	)
	hits, _ := NewSearch(g).AdjBreadth1stAt(0, func(u, v groph.VIdx, c bool) bool {
		if c {
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
		search.AdjDepth1st(false, func(_, _ groph.VIdx, ci bool, cl int) bool {
			if ci {
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

func TestSearch_Breadth1st_dir_find_clusters(t *testing.T) {
	g := adjmatrix.NewDBitmap(2, nil)
	search := NewSearch(g)
	test := func() {
		search.Reset(g)
		hits, circ := 0, 0
		clusters := make(map[int]bool)
		search.AdjBreadth1st(false, func(_, _ groph.VIdx, ci bool, cl int) bool {
			if ci {
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
	srch := NewSearch(g)
	srch.SortBy = VIdxOrder
	srch.AdjDepth1stAt(0, func(u, v int, circ bool) bool {
		fmt.Println(u, v, circ, srch.Hits(v))
		return false
	})
	// Output:
	// -1 0 false 0
	// 0 1 false 0
	// 1 2 false 0
	// 2 0 true 1
	// 2 3 false 0
	// 3 0 true 2
	// 0 2 true 1
	// 0 3 true 1
}

func ExampleSearch_ugd1st_loop_not_parent() {
	g := adjmatrix.NewUBool(2, nil)
	groph.Set(g, true, e(0, 1), e(1, 1))
	search := NewSearch(g)
	search.SortBy = VIdxOrder
	search.AdjDepth1stAt(0, func(u, v int, circ bool) bool {
		fmt.Println(u, v, circ, search.Hits(v))
		return false
	})
	// Output:
	// -1 0 false 0
	// 0 1 false 0
	// 1 1 true 1
}

func ExampleSearch_ugb1st_loop_not_parent() {
	g := adjmatrix.NewUBool(2, nil)
	groph.Set(g, true, e(0, 1), e(1, 1))
	search := NewSearch(g)
	search.SortBy = VIdxOrder
	search.AdjBreadth1stAt(0, func(u, v int, circ bool) bool {
		fmt.Println(u, v, circ, search.Hits(v))
		return false
	})
	// Output:
	// -1 0 false 0
	// 0 1 false 0
	// 1 1 true 1
}
