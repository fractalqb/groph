package test

import (
	"fmt"
	"reflect"
	"testing"

	"git.fractalqb.de/fractalqb/groph/sliceofmaps"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmatrix"
)

func TestEachOutgoing_directed(t *testing.T) {
	g, _ := adjmatrix.AsDBool(nil, []bool{
		false, true, false,
		false, false, false,
		true, false, false,
	})
	var ns []groph.VIdx
	groph.EachOutgoing(g, 0, func(d groph.VIdx) bool {
		ns = append(ns, d)
		return false
	})
	if !reflect.DeepEqual(ns, []groph.VIdx{1}) {
		t.Errorf("expected [1], got %v", ns)
	}
}

func TestEachOutgoing_undirected(t *testing.T) {
	g, _ := adjmatrix.AsUBool(nil, []bool{
		false,
		true, false,
		false, true, false,
	})
	var ns []groph.VIdx
	groph.EachOutgoing(g, 1, func(d groph.VIdx) bool {
		ns = append(ns, d)
		return false
	})
	if !reflect.DeepEqual(ns, []groph.VIdx{0, 2}) {
		t.Errorf("expected [0 2], got %v", ns)
	}
}

func TestEachIncoming_directed(t *testing.T) {
	g, _ := adjmatrix.AsDBool(nil, []bool{
		false, true, false,
		false, false, false,
		true, false, false,
	})
	var ns []groph.VIdx
	groph.EachIncoming(g, 0, func(d groph.VIdx) bool {
		ns = append(ns, d)
		return false
	})
	if !reflect.DeepEqual(ns, []groph.VIdx{2}) {
		t.Errorf("expected [1], got %v", ns)
	}
}

func TestEachIncoming_undirected(t *testing.T) {
	g, _ := adjmatrix.AsUBool(nil, []bool{
		false,
		true, false,
		false, true, false,
	})
	var ns []groph.VIdx
	groph.EachIncoming(g, 1, func(d groph.VIdx) bool {
		ns = append(ns, d)
		return false
	})
	if !reflect.DeepEqual(ns, []groph.VIdx{0, 2}) {
		t.Errorf("expected [0 2], got %v", ns)
	}
}

func ExampleEachEdge_directed() {
	g, _ := adjmatrix.AsDBool(nil, []bool{
		false, true, false,
		false, false, false,
		true, false, false,
	})
	var ns []groph.Edge
	groph.EachEdge(g, func(u, v groph.VIdx) bool {
		ns = append(ns, groph.Edge{U: u, V: v})
		return false
	})
	fmt.Println(ns)
	// Output:
	// [{0 1} {2 0}]
}

func ExampleEachEdge_undirected() {
	g, _ := adjmatrix.AsUBool(nil, []bool{
		false,
		true, false,
		false, true, false,
	})
	var ns []groph.Edge
	groph.EachEdge(g, func(u, v groph.VIdx) bool {
		ns = append(ns, groph.Edge{U: u, V: v})
		return false
	})
	fmt.Println(ns)
	// Output:
	// [{1 0} {2 1}]
}

func TestSize_undir(t *testing.T) {
	var g groph.WUndirected = adjmatrix.NewUInt32(4, adjmatrix.I32Del, nil).
		Init(adjmatrix.I32Del)
	if sz := groph.Size(g); sz != 0 {
		t.Fatalf("new graph size not 0: size=%d", sz)
	}
	type E = groph.Edge
	groph.Set(g, int32(1), E{0, 0}, E{1, 2}, E{2, 1}, E{2, 3})
	if sz := groph.Size(g); sz != 3 {
		t.Fatalf("unexpected graph size: want 3, got %d", sz)
	}
}

func TestSize_undir_outlister(t *testing.T) {
	var g groph.WUndirected = sliceofmaps.NewUInt32(4, nil)
	if _, ok := g.(groph.OutLister); !ok {
		t.Fatal("graph is not an out lister")
	}
	type E = groph.Edge
	groph.Set(g, int32(1), E{0, 0}, E{1, 2}, E{2, 1}, E{2, 3}, E{2, 0})
	if sz := groph.Size(g); sz != 4 {
		t.Fatalf("unexpected graph size: want 4, got %d", sz)
	}
}

// TODO more test cases for Size()
