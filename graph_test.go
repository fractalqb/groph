package groph

import (
	"testing"
)

func testSetUnset(g WGraph, w interface{}, t *testing.T) {
	if g.Directed() {
		testDSetUnset(g, w, t)
	} else {
		t.Fatal("testing undirected graphs NYI")
	}
}

func testDSetUnset(g WGraph, w interface{}, t *testing.T) {
	vno := g.VertexNo()
	for wi := uint(0); wi < vno; wi++ {
		for wj := uint(0); wj < vno; wj++ {
			g.SetWeight(wi, wj, nil)
		}
	}
	for ri := uint(0); ri < vno; ri++ {
		for rj := uint(0); rj < vno; rj++ {
			if w := g.Weight(ri, rj); w != nil {
				t.Errorf("read non-nil value [%v] @%d,%d after clear",
					w,
					ri, rj)
			}
		}
	}
	for wi := uint(0); wi < vno; wi++ {
		for wj := uint(0); wj < vno; wj++ {
			g.SetWeight(wi, wi, w)
			for ri := uint(0); ri < vno; ri++ {
				for rj := uint(0); rj < vno; rj++ {
					r := g.Weight(ri, rj)
					if ri == wi && ri == rj {
						if r != w {
							t.Errorf("read wrong value [%v] (expect [%v]) @%d,%d",
								r, w,
								wi, wj)
						}
					} else if r != nil {
						t.Errorf("read non-nil value [%v] @%d,%d after write @%d,%d",
							w,
							ri, rj,
							wi, wj)
					}
				}
			}
			g.SetWeight(wi, wi, nil)
		}
	}
}
