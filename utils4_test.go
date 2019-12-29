package groph

// test utilities have a separate file to not pollute godoc examples

import "testing"

const testSizeSetUnset = 11

func testGenericSetUnset(t *testing.T, g WGraph, w interface{}) {
	genClear := func(i, j VIdx) { g.SetWeight(i, j, nil) }
	genIsClear := func(w interface{}) bool { return w == nil }
	genRead := func(i, j VIdx) interface{} { return g.Weight(i, j) }
	if u, ok := g.(WUndirected); ok {
		testUSetUnset(t, u, genClear, genIsClear,
			func(i, j VIdx) interface{} { u.SetWeightU(i, j, w); return w },
			genRead,
		)
	} else {
		testDSetUnset(t, g, genClear, genIsClear,
			func(i, j VIdx) interface{} { g.SetWeight(i, j, w); return w },
			genRead,
		)
	}
}

func testDSetUnset(
	t *testing.T,
	g WGraph,
	clear func(i, j VIdx),
	isCleared func(w interface{}) bool,
	set func(i, j VIdx) interface{},
	read func(i, j VIdx) interface{},
) {
	if !Directed(g) {
		t.Fatal("graph is not directed")
	}
	vno := g.VertexNo()
	for wi := VIdx(0); wi < vno; wi++ {
		for wj := VIdx(0); wj < vno; wj++ {
			clear(wi, wj)
		}
	}
	for ri := VIdx(0); ri < vno; ri++ {
		for rj := VIdx(0); rj < vno; rj++ {
			if w := read(ri, rj); !isCleared(w) {
				t.Errorf("read non-cleared value [%v] @%d,%d after clear",
					w,
					ri, rj)
			}
		}
	}
	for wi := VIdx(0); wi < vno; wi++ {
		for wj := VIdx(0); wj < vno; wj++ {
			w := set(wi, wj)
			for ri := VIdx(0); ri < vno; ri++ {
				for rj := VIdx(0); rj < vno; rj++ {
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

func testUSetUnset(
	t *testing.T,
	g WUndirected,
	clear func(i, j VIdx),
	isCleared func(w interface{}) bool,
	set func(i, j VIdx) interface{},
	read func(i, j VIdx) interface{},
) {
	vno := g.VertexNo()
	for wi := VIdx(0); wi < vno; wi++ {
		for wj := 0; wj <= wi; wj++ {
			clear(wi, wj)
		}
	}
	for ri := VIdx(0); ri < vno; ri++ {
		for rj := ri; rj < vno; rj++ {
			if w := read(ri, rj); !isCleared(w) {
				t.Errorf("read non-cleared value [%v] @%d,%d after clear",
					w,
					ri, rj)
			}
		}
	}
	for wi := VIdx(0); wi < vno; wi++ {
		for wj := VIdx(0); wj <= wi; wj++ {
			w := set(wi, wj)
			for ri := VIdx(0); ri < vno; ri++ {
				for rj := VIdx(0); rj < vno; rj++ {
					expectSet := (ri == wi && rj == wj) || (ri == wj && rj == wi)
					r := read(ri, rj)
					if expectSet {
						if r != w {
							t.Fatalf("read wrong value [%v] (expect [%v]) @%d,%d",
								r, w,
								ri, rj)
						}
					} else if !isCleared(r) {
						t.Fatalf("read non-cleared value [%v] @%d,%d after write @%d,%d",
							w,
							ri, rj,
							wi, wj)
					}
					r = read(rj, ri)
					if expectSet {
						if r != w {
							t.Fatalf("read wrong value [%v] (expect [%v]) @%d,%d",
								r, w,
								ri, rj)
						}
					} else if !isCleared(r) {
						t.Fatalf("read non-cleared value [%v] @%d,%d after write @%d,%d",
							w,
							ri, rj,
							wi, wj)
					}
					// TODO optimized read
				}
			}
			clear(wi, wj)
		}
	}
}
