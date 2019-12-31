package groph

import (
	"fmt"
	"testing"
)

func ExampleSet() {
	g := NewAdjMxDbool(3, nil)
	type E = Edge
	Set(g, true, E{0, 1}, E{1, 2}, E{2, 0})
	fmt.Println("graph size:", Size(g))
	// Output:
	// graph size: 3
}

func TestSize_undir(t *testing.T) {
	var g WUndirected = NewAdjMxUi32(4, I32Del, nil)
	if sz := Size(g); sz != 0 {
		t.Fatalf("new graph size not 0: size=%d", sz)
	}
	type E = Edge
	Set(g, int32(1), E{0, 0}, E{1, 2}, E{2, 1}, E{2, 3})
	if sz := Size(g); sz != 3 {
		t.Fatalf("unexpected graph size: want 3, got %d", sz)
	}
}

func TestSize_undir_outlister(t *testing.T) {
	var g WUndirected = NewSoMUi32(4, nil)
	if _, ok := g.(OutLister); !ok {
		t.Fatal("graph is not an out lister")
	}
	type E = Edge
	Set(g, int32(1), E{0, 0}, E{1, 2}, E{2, 1}, E{2, 3}, E{2, 0})
	if sz := Size(g); sz != 4 {
		t.Fatalf("unexpected graph size: want 4, got %d", sz)
	}
}

// TODO more test cases for Size()
