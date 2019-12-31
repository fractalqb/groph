package groph

import (
	"testing"
)

var (
	_ WGraph      = (*SoMD)(nil)
	_ OutLister   = (*SoMD)(nil)
	_ WUndirected = (*SoMU)(nil)
	_ OutLister   = (*SoMU)(nil)
	_ WGi32       = (*SoMDi32)(nil)
	_ OutLister   = (*SoMDi32)(nil)
	_ WGf32       = (*SoMDf32)(nil)
	_ OutLister   = (*SoMDf32)(nil)
	_ WUi32       = (*SoMUi32)(nil)
	_ OutLister   = (*SoMUi32)(nil)
	_ WGf32       = (*SoMUf32)(nil)
	_ OutLister   = (*SoMUf32)(nil)
)

func TestSoMD(t *testing.T) {
	g := NewSoMD(testSizeSetDel, nil)
	testGenericSetDel(t, g, testProbeGen)
}

func TestSoMU(t *testing.T) {
	g := NewSoMU(testSizeSetDel, nil)
	testGenericSetDel(t, g, testProbeGen)
}

func BenchmarkSoMD_generic(b *testing.B) {
	m := NewSoMD(testSizeSetDel, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		for i := V0; i < max; i++ {
			for j := V0; j < max; j++ {
				m.SetWeight(i, j, n)
			}
		}
		for i := V0; i < max; i++ {
			for j := V0; j < max; j++ {
				r := m.Weight(i, j)
				if r != n {
					b.Fatal("unexpected read", n, r)
				}
			}
		}
	}
}

func TestSoMDi32(t *testing.T) {
	g := NewSoMDi32(testSizeSetDel, nil)
	t.Run("is WGi32", func(t *testing.T) { testIsWGi32(t, g) })
	testDSetDel(t, g,
		func(i, j VIdx) { g.DelEdge(i, j) },
		func(w interface{}) bool { return w == nil },
		func(i, j VIdx) interface{} { g.SetEdge(i, j, testProbeI32); return testProbeI32 },
		func(i, j VIdx) interface{} { return g.Weight(i, j) },
	)
}

func TestSoMUi32(t *testing.T) {
	u := NewSoMUi32(testSizeSetDel, nil)
	t.Run("is WUi32", func(t *testing.T) { testIsWUi32(t, u) })
	testUSetDel(t, u,
		func(i, j VIdx) { u.DelEdgeU(i, j) },
		func(w interface{}) bool { return w == nil },
		func(i, j VIdx) interface{} { u.SetEdgeU(i, j, testProbeI32); return testProbeI32 },
		func(i, j VIdx) interface{} {
			if w, ok := u.Edge(i, j); ok {
				return w
			}
			return nil
		},
	)
}

func BenchmarkSoMDi32_generic(b *testing.B) {
	m := NewSoMDi32(testSizeSetDel, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := int32(n)
		for i := V0; i < max; i++ {
			for j := V0; j < max; j++ {
				m.SetWeight(i, j, w)
			}
		}
		for i := V0; i < max; i++ {
			for j := V0; j < max; j++ {
				r := m.Weight(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func BenchmarkSoMDi32(b *testing.B) {
	m := NewSoMDi32(testSizeSetDel, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := int32(n)
		for i := V0; i < max; i++ {
			for j := V0; j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := V0; i < max; i++ {
			for j := V0; j < max; j++ {
				r, _ := m.Edge(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func TestSoMDf32(t *testing.T) {
	g := NewSoMDf32(testSizeSetDel, nil)
	t.Run("is WGf32", func(t *testing.T) { testIsWGf32(t, g) })
	const w32 = float32(3.1415)
	testDSetDel(t, g,
		func(i, j VIdx) { g.SetEdge(i, j, NaN32()) },
		func(w interface{}) bool { return IsNaN32(w.(float32)) },
		func(i, j VIdx) interface{} { g.SetEdge(i, j, w32); return w32 },
		func(i, j VIdx) interface{} { return g.Edge(i, j) },
	)
}

func TestSoMUf32(t *testing.T) {
	u := NewSoMUf32(testSizeSetDel, nil)
	t.Run("is WUf32", func(t *testing.T) { testIsWUf32(t, u) })
	const w32 = float32(3.1415)
	testUSetDel(t, u,
		func(i, j VIdx) { u.SetEdgeU(i, j, NaN32()) },
		func(w interface{}) bool { return IsNaN32(w.(float32)) },
		func(i, j VIdx) interface{} { u.SetEdgeU(i, j, w32); return w32 },
		func(i, j VIdx) interface{} { return u.Edge(i, j) },
	)
}

func BenchmarkSoMDf32_generic(b *testing.B) {
	m := NewSoMDf32(testSizeSetDel, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := float32(n)
		for i := V0; i < max; i++ {
			for j := V0; j < max; j++ {
				m.SetWeight(i, j, w)
			}
		}
		for i := V0; i < max; i++ {
			for j := V0; j < max; j++ {
				r := m.Weight(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func BenchmarkSoMDf32(b *testing.B) {
	m := NewSoMDf32(testSizeSetDel, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		w := float32(n)
		for i := V0; i < max; i++ {
			for j := V0; j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := V0; i < max; i++ {
			for j := V0; j < max; j++ {
				if r := m.Edge(i, j); r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}
