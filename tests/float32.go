package tests

import (
	"testing"

	"git.fractalqb.de/fractalqb/groph"
)

const F32Probe float32 = 3.1415

func IsWGf32Test(t *testing.T, g groph.WGf32) {
	t.Run("generic set and del", func(t *testing.T) {
		GenericSetDelTest(t, g, F32Probe)
	})
	for i := 0; i < g.Order(); i++ {
		for j := 0; j < g.Order(); j++ {
			g.SetEdge(i, j, F32Probe)
			g.SetEdge(i, j, groph.NaN32())
			if w := g.Edge(i, j); !groph.IsNaN32(w) {
				t.Errorf("set edge (%d,%d) NaN returns edge weight %f", i, j, w)
			}
			if w := g.Weight(i, j); w != nil {
				t.Errorf("set edge (%d,%d) NaN returns non-nil weight %v", i, j, w)
			}
			g.SetEdge(i, j, F32Probe)
			if w := g.Edge(i, j); w != F32Probe {
				t.Errorf("set edge (%d,%d) returns wrong weight %f", i, j, w)
			}
			if w := g.Weight(i, j); w == nil {
				t.Errorf("set edge (%d,%d) does return nil weight", i, j)
			}
			g.SetWeight(i, j, nil)
			if w := g.Edge(i, j); !groph.IsNaN32(w) {
				t.Errorf("set edge (%d,%d) nil returns edge weight %f", i, j, w)
			}
			if w := g.Weight(i, j); w != nil {
				t.Errorf("set edge (%d,%d) nil returns non-nil weight %v", i, j, w)
			}
		}
	}
}

func IsWUf32Test(t *testing.T, g groph.WUf32) {
	t.Run("is WGf32", func(t *testing.T) { IsWGf32Test(t, g) })
	// TODO undirected
}
