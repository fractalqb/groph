package groph

import (
	"fmt"
	"testing"
)

const testSizeSetUnset = 11

func testGenericSetUnset(t *testing.T, g WGraph, w interface{}) {
	testSetUnset(t, g,
		func(i, j uint) interface{} { g.SetWeight(i, j, w); return w },
		func(i, j uint) { g.SetWeight(i, j, nil) },
		func(i, j uint) interface{} { return g.Weight(i, j) },
		func(w interface{}) bool { return w == nil },
	)
}

func testSetUnset(
	t *testing.T,
	g WGraph,
	set func(i, j uint) interface{},
	clear func(i, j uint),
	read func(i, j uint) interface{},
	isCleared func(w interface{}) bool,
) {
	if g.Directed() {
		testDSetUnset(t, g, set, clear, read, isCleared)
	} else {
		t.Fatal("testing undirected graphs NYI")
	}
}

func testDSetUnset(
	t *testing.T,
	g WGraph,
	set func(i, j uint) interface{},
	clear func(i, j uint),
	read func(i, j uint) interface{},
	isCleared func(w interface{}) bool,
) {
	vno := g.VertexNo()
	for wi := uint(0); wi < vno; wi++ {
		for wj := uint(0); wj < vno; wj++ {
			clear(wi, wj)
		}
	}
	for ri := uint(0); ri < vno; ri++ {
		for rj := uint(0); rj < vno; rj++ {
			if w := read(ri, rj); !isCleared(w) {
				t.Errorf("read non-cleared value [%v] @%d,%d after clear",
					w,
					ri, rj)
			}
		}
	}
	for wi := uint(0); wi < vno; wi++ {
		for wj := uint(0); wj < vno; wj++ {
			w := set(wi, wj)
			for ri := uint(0); ri < vno; ri++ {
				for rj := uint(0); rj < vno; rj++ {
					r := read(ri, rj)
					if ri == wi && rj == wj {
						if r != w {
							t.Fatalf("read wrong value [%v] (expect [%v]) @%d,%d",
								r, w,
								wi, wj)
						}
					} else if !isCleared(r) {
						t.Fatalf("read non-cleared value [%v] @%d,%d after write @%d,%d",
							w,
							ri, rj,
							wi, wj)
					}
				}
			}
			clear(wi, wj)
		}
	}
}

func ExampleReorderPath() {
	data := []string{"a", "b", "c", "d", "e"}
	path := []uint{1, 3, 0, 4, 2}
	ReorderPath(data, path)
	fmt.Println(data)
	// Output:
	// [b d a e c]
}
