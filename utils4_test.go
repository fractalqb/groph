package groph

// test utilities have a separate file to not pollute godoc examples

import "testing"

const (
	testSizeSetDel         = 11
	testProbeGen   string  = "probe"
	testProbeBool  bool    = true
	testProbeI32   int32   = 4711
	testProbeF32   float32 = 3.1415
)

func testGenericSetDel(t *testing.T, g WGraph, probeWeight interface{}) {
	genClear := func(i, j VIdx) { g.SetWeight(i, j, nil) }
	genIsClear := func(w interface{}) bool { return w == nil }
	genRead := func(i, j VIdx) interface{} { return g.Weight(i, j) }
	if u, ok := g.(WUndirected); ok {
		testUSetDel(t, u, genClear, genIsClear,
			func(i, j VIdx) interface{} {
				u.SetWeightU(i, j, probeWeight)
				return probeWeight
			},
			genRead,
		)
	} else {
		testDSetDel(t, g, genClear, genIsClear,
			func(i, j VIdx) interface{} {
				g.SetWeight(i, j, probeWeight)
				return probeWeight
			},
			genRead)
	}
}

func testDSetDel(
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

func testUSetDel(
	t *testing.T,
	g WUndirected,
	clear func(i, j VIdx),
	isCleared func(w interface{}) bool,
	setU func(i, j VIdx) interface{},
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
			w := setU(wi, wj)
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

func testIsWGbool(t *testing.T, g WGbool) {
	t.Run("generic set and del", func(t *testing.T) {
		testGenericSetDel(t, g, testProbeBool)
	})
	for i := VIdx(0); i < g.VertexNo(); i++ {
		for j := VIdx(0); j < g.VertexNo(); j++ {
			g.SetEdge(i, j, true)
			g.SetEdge(i, j, false)
			if g.Edge(i, j) != false {
				t.Errorf("set edge (%d,%d) false does not return false", i, j)
			}
			if w := g.Weight(i, j); w != nil {
				t.Errorf("set edge (%d,%d) false does non-nil weight %v", i, j, w)
			}
			g.SetEdge(i, j, true)
			if g.Edge(i, j) != true {
				t.Errorf("set edge (%d,%d) true does not return true", i, j)
			}
			if g.Weight(i, j) == nil {
				t.Errorf("set edge (%d,%d) true does return nil weight", i, j)
			}
			g.SetWeight(i, j, nil)
			if g.Edge(i, j) != false {
				t.Errorf("set weight (%d,%d) nil does not return false", i, j)
			}
			if w := g.Weight(i, j); w != nil {
				t.Errorf("set weight (%d,%d) nil returns non-nil weight %v", i, j, w)
			}
		}
	}
}

func testIsWUbool(t *testing.T, g WUbool) {
	t.Run("is WGbool", func(t *testing.T) { testIsWGbool(t, g) })
	// TODO undirected
}

func testIsWGi32(t *testing.T, g WGi32) {
	t.Run("generic set and del", func(t *testing.T) {
		testGenericSetDel(t, g, testProbeI32)
	})
	for i := VIdx(0); i < g.VertexNo(); i++ {
		for j := VIdx(0); j < g.VertexNo(); j++ {
			g.SetEdge(i, j, 4711)
			g.DelEdge(i, j)
			if w, ok := g.Edge(i, j); ok {
				t.Errorf("del edge (%d,%d) returns edge weight %d", i, j, w)
			}
			if w := g.Weight(i, j); w != nil {
				t.Errorf("del edge (%d,%d) returns non-nil weight %v", i, j, w)
			}
			g.SetEdge(i, j, 4711)
			if w, ok := g.Edge(i, j); !ok {
				t.Errorf("set edge (%d,%d) does not return edge (%d)", i, j, w)
			} else if w != 4711 {
				t.Errorf("set edge (%d,%d) does not return wrong weight %d", i, j, w)
			}
			if w := g.Weight(i, j); w == nil {
				t.Errorf("set edge (%d,%d) does return nil weight", i, j)
			}
			g.SetWeight(i, j, nil)
			if w, ok := g.Edge(i, j); ok {
				t.Errorf("set edge (%d,%d) nil returns edge weight %d", i, j, w)
			}
			if w := g.Weight(i, j); w != nil {
				t.Errorf("set edge (%d,%d) nil returns non-nil weight %v", i, j, w)
			}
		}
	}
}

func testIsWUi32(t *testing.T, g WUi32) {
	t.Run("is WGi32", func(t *testing.T) { testIsWGi32(t, g) })
	// TODO undirected
}

func testIsWGf32(t *testing.T, g WGf32) {
	t.Run("generic set and del", func(t *testing.T) {
		testGenericSetDel(t, g, testProbeF32)
	})
	for i := VIdx(0); i < g.VertexNo(); i++ {
		for j := VIdx(0); j < g.VertexNo(); j++ {
			g.SetEdge(i, j, testProbeF32)
			g.SetEdge(i, j, NaN32())
			if w := g.Edge(i, j); !IsNaN32(w) {
				t.Errorf("set edge (%d,%d) NaN returns edge weight %f", i, j, w)
			}
			if w := g.Weight(i, j); w != nil {
				t.Errorf("set edge (%d,%d) NaN returns non-nil weight %v", i, j, w)
			}
			g.SetEdge(i, j, testProbeF32)
			if w := g.Edge(i, j); w != testProbeF32 {
				t.Errorf("set edge (%d,%d) returns wrong weight %f", i, j, w)
			}
			if w := g.Weight(i, j); w == nil {
				t.Errorf("set edge (%d,%d) does return nil weight", i, j)
			}
			g.SetWeight(i, j, nil)
			if w := g.Edge(i, j); !IsNaN32(w) {
				t.Errorf("set edge (%d,%d) nil returns edge weight %f", i, j, w)
			}
			if w := g.Weight(i, j); w != nil {
				t.Errorf("set edge (%d,%d) nil returns non-nil weight %v", i, j, w)
			}
		}
	}
}

func testIsWUf32(t *testing.T, g WUf32) {
	t.Run("is WGf32", func(t *testing.T) { testIsWGf32(t, g) })
	// TODO undirected
}
