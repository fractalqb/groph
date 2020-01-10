package tests

import (
	"testing"

	"git.fractalqb.de/fractalqb/groph"
)

const BoolProbe bool = true

func IsWGboolTest(t *testing.T, g groph.WGbool) {
	t.Run("generic set and del", func(t *testing.T) {
		GenericSetDelTest(t, g, BoolProbe)
	})
	for i := 0; i < g.Order(); i++ {
		for j := 0; j < g.Order(); j++ {
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

func IsWUboolTest(t *testing.T, g groph.WUbool) {
	t.Run("is groph.WGbool", func(t *testing.T) { IsWGboolTest(t, g) })
	for i := 0; i < g.Order(); i++ {
		for j := 0; j <= i; j++ {
			g.SetEdge(i, j, true)
			g.SetEdge(j, i, true)
			g.SetEdgeU(i, j, false)
			if g.Edge(i, j) != false || g.Edge(j, i) != false {
				t.Errorf("set edge (%d,%d) false does not return 2x dir false", i, j)
			}
			if g.EdgeU(i, j) != false {
				t.Errorf("set edge (%d,%d) false does not return true", i, j)
			}
			if g.Weight(i, j) != nil || g.Weight(j, i) != nil {
				t.Errorf("set edge (%d,%d) false does not return 2x dir nil", i, j)
			}
			if g.WeightU(i, j) != nil {
				t.Errorf("set edge (%d,%d) false does not return nil", i, j)
			}
			g.SetEdgeU(i, j, true)
			if g.Edge(i, j) != true || g.Edge(j, i) != true {
				t.Errorf("set edge (%d,%d) true does not return 2x dir true", i, j)
			}
			if g.EdgeU(i, j) != true {
				t.Errorf("set edge (%d,%d) true does not return true", i, j)
			}
			if g.Weight(i, j) == nil || g.Weight(j, i) == nil {
				t.Errorf("set edge (%d,%d) true does not return 2x dir non-nil", i, j)
			}
			if g.WeightU(i, j) == nil {
				t.Errorf("set edge (%d,%d) true does not return non-nil", i, j)
			}
			g.SetWeightU(i, j, nil)
			if g.Edge(i, j) != false || g.Edge(j, i) != false {
				t.Errorf("set weight (%d,%d) nil does not return 2x dir false", i, j)
			}
			if g.EdgeU(i, j) != false {
				t.Errorf("set weight (%d,%d) nil does not return false", i, j)
			}
			if g.Weight(i, j) != nil || g.Weight(j, i) != nil {
				t.Errorf("set weight (%d,%d) nil does not return 2x dir nil", i, j)
			}
			if g.WeightU(i, j) != nil {
				t.Errorf("set weight (%d,%d) nil does not return nil", i, j)
			}
		}
	}
}
