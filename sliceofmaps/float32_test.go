package sliceofmaps

import (
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/tests"
)

var (
	_ groph.WGf32     = (*SoMDf32)(nil)
	_ groph.OutLister = (*SoMDf32)(nil)
	_ groph.WGf32     = (*SoMUf32)(nil)
	_ groph.OutLister = (*SoMUf32)(nil)
)

func TestSoMDf32(t *testing.T) {
	g := NewSoMDf32(tests.SetDelSize, nil)
	t.Run("is WGf32", func(t *testing.T) { tests.IsWGf32Test(t, g) })
	const w32 = float32(3.1415)
	tests.DSetDelTest(t, g,
		func(i, j groph.VIdx) { g.SetEdge(i, j, groph.NaN32()) },
		func(w interface{}) bool { return groph.IsNaN32(w.(float32)) },
		func(i, j groph.VIdx) interface{} { g.SetEdge(i, j, w32); return w32 },
		func(i, j groph.VIdx) interface{} { return g.Edge(i, j) },
	)
}

func TestSoMUf32(t *testing.T) {
	u := NewSoMUf32(tests.SetDelSize, nil)
	t.Run("is WUf32", func(t *testing.T) { tests.IsWUf32Test(t, u) })
	const w32 = float32(3.1415)
	tests.USetDelTest(t, u,
		func(i, j groph.VIdx) { u.SetEdgeU(i, j, groph.NaN32()) },
		func(w interface{}) bool { return groph.IsNaN32(w.(float32)) },
		func(i, j groph.VIdx) interface{} { u.SetEdgeU(i, j, w32); return w32 },
		func(i, j groph.VIdx) interface{} { return u.Edge(i, j) },
	)
}

func BenchmarkSoMDf32(b *testing.B) {
	m := NewSoMDf32(tests.SetDelSize, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := float32(n)
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				if r := m.Edge(i, j); r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func BenchmarkSoMDf32_generic(b *testing.B) {
	m := NewSoMDf32(tests.SetDelSize, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := float32(n)
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				m.SetWeight(i, j, w)
			}
		}
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				r := m.Weight(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}
