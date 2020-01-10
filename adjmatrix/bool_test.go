package adjmatrix

import (
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/tests"
)

var (
	_ groph.WGbool = (*DBool)(nil)
	_ groph.WGbool = (*UBool)(nil)
)

func TestDBool(t *testing.T) {
	m := NewDBool(3, nil)
	t.Run("is WGbool", func(t *testing.T) { tests.IsWGboolTest(t, m) })
	tests.DSetDelTest(t, m,
		func(i, j groph.VIdx) { m.SetEdge(i, j, false) },
		func(w interface{}) bool { return w.(bool) == false },
		func(i, j groph.VIdx) interface{} { m.SetEdge(i, j, true); return true },
		func(i, j groph.VIdx) interface{} { return m.Edge(i, j) },
	)
}

func BenchmarkDBool(b *testing.B) {
	m := NewDBool(tests.SetDelSize, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := true
		if n&1 == 0 {
			w = false
		}
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				r := m.Edge(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func BenchmarkDBool_generic(b *testing.B) {
	m := NewDBool(tests.SetDelSize, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := true
		if n&1 == 0 {
			w = false
		}
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				m.SetWeight(i, j, w)
			}
		}
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				r := m.Weight(i, j)
				if (w && r == nil) || (!w && r != nil) {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}
