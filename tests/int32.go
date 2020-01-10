package tests

import (
	"testing"

	"git.fractalqb.de/fractalqb/groph"
)

const I32Probe int32 = 4711

func IsWGi32Test(t *testing.T, g groph.WGi32) {
	t.Run("generic set and del", func(t *testing.T) {
		GenericSetDelTest(t, g, I32Probe)
	})
	for i := 0; i < g.Order(); i++ {
		for j := 0; j < g.Order(); j++ {
			g.SetEdge(i, j, 4711)
			g.SetWeight(i, j, nil)
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

func IsWUi32Test(t *testing.T, g groph.WUi32) {
	t.Run("is WGi32", func(t *testing.T) { IsWGi32Test(t, g) })
	// TODO undirected
}
