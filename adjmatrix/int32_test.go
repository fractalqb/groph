package adjmatrix

import (
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/tests"
)

var (
	_ groph.WGi32 = (*DInt32)(nil)
	_ groph.WUi32 = (*UInt32)(nil)
)

func TestDInt32(t *testing.T) {
	g := NewDInt32(tests.SetDelSize, I32Del, nil)
	t.Run("is WGi32", func(t *testing.T) { tests.IsWGi32Test(t, g) })
	tests.DSetDelTest(t, g,
		func(i, j groph.VIdx) { g.SetEdge(i, j, g.del) },
		func(w interface{}) bool { return w == nil },
		func(i, j groph.VIdx) interface{} {
			g.SetEdge(i, j, tests.I32Probe)
			return tests.I32Probe
		},
		func(i, j groph.VIdx) interface{} { return g.Weight(i, j) },
	)
}

func TestUInt32(t *testing.T) {
	u := NewUInt32(tests.SetDelSize, I32Del, nil)
	t.Run("is WUi32", func(t *testing.T) { tests.IsWUi32Test(t, u) })
	tests.USetDelTest(t, u,
		func(i, j groph.VIdx) { u.SetEdgeU(i, j, u.del) },
		func(w interface{}) bool { return w.(int32) == u.del },
		func(i, j groph.VIdx) interface{} {
			u.SetEdgeU(i, j, tests.I32Probe)
			return tests.I32Probe
		},
		func(i, j groph.VIdx) interface{} {
			w, _ := u.Edge(i, j)
			return w
		},
	)
}

func BenchmarkDInt32(b *testing.B) {
	m := NewDInt32(tests.SetDelSize, I32Del, nil)
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

func BenchmarkDInt32_generic(b *testing.B) {
	m := NewDInt32(tests.SetDelSize, I32Del, nil)
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
