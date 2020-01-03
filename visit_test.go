package groph

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEachOutgoing_directed(t *testing.T) {
	g, _ := AsAdjMxDbool(nil, []bool{
		false, true, false,
		false, false, false,
		true, false, false,
	})
	var ns []VIdx
	EachOutgoing(g, 0, func(d VIdx) { ns = append(ns, d) })
	if !reflect.DeepEqual(ns, []VIdx{1}) {
		t.Errorf("expected [1], got %v", ns)
	}
}

func TestEachOutgoing_undirected(t *testing.T) {
	g, _ := AsAdjMxUbool(nil, []bool{
		false,
		true, false,
		false, true, false,
	})
	var ns []VIdx
	EachOutgoing(g, 1, func(d VIdx) { ns = append(ns, d) })
	if !reflect.DeepEqual(ns, []VIdx{0, 2}) {
		t.Errorf("expected [0 2], got %v", ns)
	}
}

func TestEachIncoming_directed(t *testing.T) {
	g, _ := AsAdjMxDbool(nil, []bool{
		false, true, false,
		false, false, false,
		true, false, false,
	})
	var ns []VIdx
	EachIncoming(g, 0, func(d VIdx) { ns = append(ns, d) })
	if !reflect.DeepEqual(ns, []VIdx{2}) {
		t.Errorf("expected [1], got %v", ns)
	}
}

func TestEachIncoming_undirected(t *testing.T) {
	g, _ := AsAdjMxUbool(nil, []bool{
		false,
		true, false,
		false, true, false,
	})
	var ns []VIdx
	EachIncoming(g, 1, func(d VIdx) { ns = append(ns, d) })
	if !reflect.DeepEqual(ns, []VIdx{0, 2}) {
		t.Errorf("expected [0 2], got %v", ns)
	}
}

func ExampleEachEdge_directed() {
	g, _ := AsAdjMxDbool(nil, []bool{
		false, true, false,
		false, false, false,
		true, false, false,
	})
	var ns []Edge
	EachEdge(g, func(u, v VIdx) { ns = append(ns, Edge{u, v}) })
	fmt.Println(ns)
	// Output:
	// [{0 1} {2 0}]
}

func ExampleEachEdge_undirected() {
	g, _ := AsAdjMxUbool(nil, []bool{
		false,
		true, false,
		false, true, false,
	})
	var ns []Edge
	EachEdge(g, func(u, v VIdx) { ns = append(ns, Edge{u, v}) })
	fmt.Println(ns)
	// Output:
	// [{1 0} {2 1}]
}
