// Package tests provides functions to verify the conformance of graph
// implementations.
package tests

import (
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	//"git.fractalqb.de/fractalqb/groph/util"
)

const (
	SetDelSize        = 11
	GenProbe   string = "probe"
)

func GenericSetDelTest(t *testing.T, g groph.WGraph, probeWeight interface{}) {
	genClear := func(i, j groph.VIdx) { g.SetWeight(i, j, nil) }
	genIsClear := func(w interface{}) bool { return w == nil }
	genSet := func(i, j groph.VIdx) interface{} {
		g.SetWeight(i, j, probeWeight)
		return probeWeight
	}
	genRead := func(i, j groph.VIdx) interface{} { return g.Weight(i, j) }
	if u, ok := g.(groph.WUndirected); ok {
		undirClear := func(i, j groph.VIdx) { u.SetWeightU(i, j, nil) }
		undirSet := func(i, j groph.VIdx) interface{} {
			if i > j {
				u.SetWeightU(i, j, probeWeight)
			} else {
				u.SetWeightU(j, i, probeWeight)
			}
			return probeWeight
		}
		undirRead := func(i, j groph.VIdx) interface{} {
			if i > j {
				return u.WeightU(i, j)
			}
			return u.WeightU(j, i)
		}
		USetDelTest(t, u, undirClear, genIsClear, genSet, genRead)
		USetDelTest(t, u, undirClear, genIsClear, genSet, undirRead)
		USetDelTest(t, u, undirClear, genIsClear, undirSet, genRead)
		USetDelTest(t, u, undirClear, genIsClear, undirSet, undirRead)
	} else {
		DSetDelTest(t, g, genClear, genIsClear, genSet, genRead)
	}
}

func DSetDelTest(
	t *testing.T,
	g groph.WGraph,
	clear func(i, j groph.VIdx),
	isCleared func(w interface{}) bool,
	set func(i, j groph.VIdx) interface{},
	read func(i, j groph.VIdx) interface{},
) {
	if _, ok := g.(groph.RUndirected); ok {
		t.Fatal("graph is not directed")
	}
	vno := g.Order()
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj < vno; wj++ {
			clear(wi, wj)
		}
	}
	for ri := 0; ri < vno; ri++ {
		for rj := 0; rj < vno; rj++ {
			if w := read(ri, rj); !isCleared(w) {
				t.Errorf("read non-cleared value [%v] @%d,%d after clear",
					w,
					ri, rj)
			}
		}
	}
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj < vno; wj++ {
			w := set(wi, wj)
			for ri := 0; ri < vno; ri++ {
				for rj := 0; rj < vno; rj++ {
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
	for i := 0; i < vno; i++ {
		for j := 0; j < vno; j++ {
			set(i, j)
		}
	}
	g.Reset(g.Order())
	if g.Order() != vno {
		t.Fatal("Reset changed graph size")
	}
	for i := 0; i < vno; i++ {
		for j := 0; j < vno; j++ {
			if g.Weight(i, j) != nil {
				t.Fatalf("Reset() did not clear the graph at (%d,%d)", i, j)
			}
		}
	}
}

func USetDelTest(
	t *testing.T,
	g groph.WUndirected,
	clear func(i, j groph.VIdx),
	isCleared func(w interface{}) bool,
	set func(i, j groph.VIdx) interface{},
	read func(i, j groph.VIdx) interface{},
) {
	vno := g.Order()
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj <= wi; wj++ {
			clear(wi, wj)
		}
	}
	for ri := 0; ri < vno; ri++ {
		for rj := ri; rj < vno; rj++ {
			if w := read(ri, rj); !isCleared(w) {
				t.Errorf("read non-cleared value [%v] @%d,%d after clear",
					w,
					ri, rj)
			}
		}
	}
	for wi := 0; wi < vno; wi++ {
		for wj := 0; wj <= wi; wj++ {
			w := set(wi, wj)
			for ri := 0; ri < vno; ri++ {
				for rj := 0; rj < vno; rj++ {
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
				}
			}
			clear(wi, wj)
		}
	}
	for i := 0; i < vno; i++ {
		for j := 0; j <= i; j++ {
			set(i, j)
		}
	}
	g.Reset(g.Order())
	if g.Order() != vno {
		t.Fatal("Reset changed graph size")
	}
	for i := 0; i < vno; i++ {
		for j := 0; j <= i; j++ {
			if g.Weight(i, j) != nil {
				t.Fatalf("Reset() did not clear the graph at (%d,%d)", i, j)
			}
		}
	}
}
