package traverse

import (
	"fmt"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
)

// Just to make go vet happy! I like
// 	type E = groph.Edge
//  E{u, v} better
func e(u, v groph.VIdx) groph.Edge { return groph.Edge{U: u, V: v} }

func ExampleSearch_Depth1stAt() {
	g := groph.NewAdjMxUbool(7, nil)
	groph.Set(g, true,
		e(0, 1), e(0, 2), e(0, 3),
		e(1, 4), e(1, 5),
		e(2, 5),
		e(3, 6),
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
	groph.Set(g, true,
		e(0, 1), e(0, 2), e(0, 3),
		e(1, 4), e(1, 5),
		e(2, 5),
		e(3, 6),
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

func TestSearch_Depth1st_avoid_loop_and_parent(t *testing.T) {
	g := groph.NewAdjMxUbool(2, nil)
	groph.Set(g, true, e(0, 1), e(1, 1))
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

func TestSearch_Breadth1st_avoid_loop_and_parent(t *testing.T) {
	g := groph.NewAdjMxUbool(2, nil)
	groph.Set(g, true, e(0, 1), e(1, 1))
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

func TestSearch_Depth1st_dir_find_clusters(t *testing.T) {
	g := groph.NewAdjMxDbitmap(2, nil)
	search := NewSearch(g)
	test := func() {
		search.Reset(g)
		hits := 0
		clusters := make(map[int]bool)
		search.Depth1st(false, func(v groph.VIdx, c int) bool {
			hits++
			clusters[c] = true
			return false
		})
		if hits != 2 {
			t.Errorf("unexpected number of hits: %d, want 2", hits)
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
	g := groph.NewAdjMxDbitmap(2, nil)
	search := NewSearch(g)
	test := func() {
		search.Reset(g)
		hits := 0
		clusters := make(map[int]bool)
		search.Breadth1st(false, func(v groph.VIdx, c int) bool {
			hits++
			clusters[c] = true
			return false
		})
		if hits != 2 {
			t.Errorf("unexpected number of hits: %d, want 2", hits)
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
