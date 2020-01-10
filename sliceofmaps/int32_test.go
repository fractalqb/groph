package sliceofmaps

import (
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/tests"
)

var (
	_ groph.WGi32     = (*DInt32)(nil)
	_ groph.OutLister = (*DInt32)(nil)
	_ groph.WUi32     = (*UInt32)(nil)
	_ groph.OutLister = (*UInt32)(nil)
)

func TestSoMDi32(t *testing.T) {
	g := NewDInt32(tests.SetDelSize, nil)
	t.Run("is WGi32", func(t *testing.T) { tests.IsWGi32Test(t, g) })
	tests.DSetDelTest(t, g,
		func(i, j groph.VIdx) { g.SetWeight(i, j, nil) },
		func(w interface{}) bool { return w == nil },
		func(i, j groph.VIdx) interface{} {
			g.SetEdge(i, j, tests.I32Probe)
			return tests.I32Probe
		},
		func(i, j groph.VIdx) interface{} { return g.Weight(i, j) },
	)
}

func TestSoMUi32(t *testing.T) {
	u := NewUInt32(tests.SetDelSize, nil)
	t.Run("is WUi32", func(t *testing.T) { tests.IsWUi32Test(t, u) })
	tests.USetDelTest(t, u,
		func(i, j groph.VIdx) { u.SetWeightU(i, j, nil) },
		func(w interface{}) bool { return w == nil },
		func(i, j groph.VIdx) interface{} {
			u.SetEdgeU(i, j, tests.I32Probe)
			return tests.I32Probe
		},
		func(i, j groph.VIdx) interface{} {
			if w, ok := u.Edge(i, j); ok {
				return w
			}
			return nil
		},
	)
}

func BenchmarkSoMDi32(b *testing.B) {
	m := NewDInt32(tests.SetDelSize, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := int32(n)
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				r, _ := m.Edge(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func BenchmarkSoMDi32_generic(b *testing.B) {
	m := NewDInt32(tests.SetDelSize, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := int32(n)
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
