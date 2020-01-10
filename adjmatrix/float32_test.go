package adjmatrix

import (
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/tests"
)

var _ groph.WGf32 = (*DFloat32)(nil)
var _ groph.WUf32 = (*UFloat32)(nil)

func TestDFloat32(t *testing.T) {
	m := NewDFloat32(tests.SetDelSize, nil)
	t.Run("is WGf32", func(t *testing.T) { tests.IsWGf32Test(t, m) })
	const w32 = float32(3.1415)
	tests.DSetDelTest(t, m,
		func(i, j groph.VIdx) { m.SetEdge(i, j, groph.NaN32()) },
		func(w interface{}) bool { return groph.IsNaN32(w.(float32)) },
		func(i, j groph.VIdx) interface{} { m.SetEdge(i, j, w32); return w32 },
		func(i, j groph.VIdx) interface{} { return m.Edge(i, j) },
	)
}

func BenchmarkDFloat32(b *testing.B) {
	m := NewDFloat32(tests.SetDelSize, nil)
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
				r := m.Edge(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func BenchmarkDFloat32_generic(b *testing.B) {
	m := NewDFloat32(tests.SetDelSize, nil)
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

func TestUBool(t *testing.T) {
	u := NewUBool(tests.SetDelSize, nil)
	t.Run("is WUbool", func(t *testing.T) { tests.IsWUboolTest(t, u) })
	tests.USetDelTest(t, u,
		func(i, j groph.VIdx) { u.SetEdgeU(i, j, false) },
		func(w interface{}) bool { return w == nil },
		func(i, j groph.VIdx) interface{} {
			u.SetEdgeU(i, j, tests.BoolProbe)
			return tests.BoolProbe
		},
		func(i, j groph.VIdx) interface{} { return u.Weight(i, j) },
	)
}

func TestUFloat32(t *testing.T) {
	m := NewUFloat32(tests.SetDelSize, nil)
	t.Run("is WUf32", func(t *testing.T) { tests.IsWUf32Test(t, m) })
	const w32 = float32(3.1415)
	tests.USetDelTest(t, m,
		func(i, j groph.VIdx) { m.SetEdgeU(i, j, groph.NaN32()) },
		func(w interface{}) bool { return groph.IsNaN32(w.(float32)) },
		func(i, j groph.VIdx) interface{} { m.SetEdgeU(i, j, w32); return w32 },
		func(i, j groph.VIdx) interface{} { return m.Edge(i, j) },
	)
}
