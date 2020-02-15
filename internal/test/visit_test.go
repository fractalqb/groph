package test

import (
	"fmt"
	"reflect"
	"testing"

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

func TestSize_dir(t *testing.T) {
	const (
		ord    = 3
		expect = ord * ord
	)
	g := testRDFull(ord)
	if !groph.Directed(g) {
		t.Fatal("undirected graph detected")
	}
	const msgpat = "unexpected graph size %d, want %d"
	if s := groph.Size(g); s != expect {
		t.Errorf(msgpat, s, expect)
	}
	t.Run("out lister", func(t *testing.T) {
		if s := groph.Size(testRDols{g}); s != expect {
			t.Errorf(msgpat, s, expect)
		}
	})
	t.Run("in lister", func(t *testing.T) {
		if s := groph.Size(testRDils{g}); s != expect {
			t.Errorf(msgpat, s, expect)
		}
	})
	t.Run("in+out lister", func(t *testing.T) {
		if s := groph.Size(testRDiols{testRDols{g}}); s != expect {
			t.Errorf(msgpat, s, expect)
		}
	})
	t.Run("edge lister", func(t *testing.T) {
		if s := groph.Size(testRDels{g}); s != expect {
			t.Errorf(msgpat, s, expect)
		}
	})
}

func TestSize_undir(t *testing.T) {
	const (
		ord    = 3
		expect = (ord * (ord + 1)) / 2
	)
	g := testRUFull(ord)
	if groph.Directed(g) {
		t.Fatal("directed graph detected")
	}
	const msgpat = "unexpected graph size %d, want %d"
	if s := groph.Size(g); s != expect {
		t.Errorf(msgpat, s, expect)
	}
	t.Run("out lister", func(t *testing.T) {
		if s := groph.Size(testRUols{g}); s != expect {
			t.Errorf(msgpat, s, expect)
		}
	})
	t.Run("in lister", func(t *testing.T) {
		if s := groph.Size(testRUils{g}); s != expect {
			t.Errorf(msgpat, s, expect)
		}
	})
	t.Run("in+out lister", func(t *testing.T) {
		if s := groph.Size(testRUiols{testRUols{g}}); s != expect {
			t.Errorf(msgpat, s, expect)
		}
	})
	t.Run("edge lister", func(t *testing.T) {
		if s := groph.Size(testRUels{g}); s != expect {
			t.Errorf(msgpat, s, expect)
		}
	})
}
