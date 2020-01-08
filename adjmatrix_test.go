package groph

import (
	"testing"
)

var _ WGbool = (*AdjMxDbitmap)(nil)
var _ WGbool = (*AdjMxDbool)(nil)
var _ WGi32 = (*AdjMxDi32)(nil)
var _ WGf32 = (*AdjMxDf32)(nil)
var _ WUbool = (*AdjMxUbool)(nil)
var _ WUi32 = (*AdjMxUi32)(nil)
var _ WUf32 = (*AdjMxUf32)(nil)

func TestAdjMxDbitmap(t *testing.T) {
	m := NewAdjMxDbitmap(3, nil)
	t.Run("is WGbool", func(t *testing.T) { testIsWGbool(t, m) })
	testDSetDel(t, m,
		func(i, j VIdx) { m.SetEdge(i, j, false) },
		func(w interface{}) bool { return w.(bool) == false },
		func(i, j VIdx) interface{} { m.SetEdge(i, j, true); return true },
		func(i, j VIdx) interface{} { return m.Edge(i, j) },
	)
}

func BenchmarkAdjMxDbitmap(b *testing.B) {
	m := NewAdjMxDbitmap(testSizeSetDel, nil)
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

func BenchmarkAdjMxDbitmap_generic(b *testing.B) {
	m := NewAdjMxDbitmap(testSizeSetDel, nil)
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

func TestAdjMxDbool(t *testing.T) {
	m := NewAdjMxDbool(3, nil)
	t.Run("is WGbool", func(t *testing.T) { testIsWGbool(t, m) })
	testDSetDel(t, m,
		func(i, j VIdx) { m.SetEdge(i, j, false) },
		func(w interface{}) bool { return w.(bool) == false },
		func(i, j VIdx) interface{} { m.SetEdge(i, j, true); return true },
		func(i, j VIdx) interface{} { return m.Edge(i, j) },
	)
}

func BenchmarkAdjMxDbool(b *testing.B) {
	m := NewAdjMxDbool(testSizeSetDel, nil)
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

func BenchmarkAdjMxDbool_generic(b *testing.B) {
	m := NewAdjMxDbool(testSizeSetDel, nil)
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

func TestAdjMxDi32(t *testing.T) {
	g := NewAdjMxDi32(testSizeSetDel, I32Del, nil)
	t.Run("is WGi32", func(t *testing.T) { testIsWGi32(t, g) })
	testDSetDel(t, g,
		func(i, j VIdx) { g.SetEdge(i, j, g.del) },
		func(w interface{}) bool { return w == nil },
		func(i, j VIdx) interface{} { g.SetEdge(i, j, testProbeI32); return testProbeI32 },
		func(i, j VIdx) interface{} { return g.Weight(i, j) },
	)
}

func BenchmarkAdjMxDi32(b *testing.B) {
	m := NewAdjMxDi32(testSizeSetDel, I32Del, nil)
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

func BenchmarkAdjMxDi32_generic(b *testing.B) {
	m := NewAdjMxDi32(testSizeSetDel, I32Del, nil)
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

func TestAdjMxDf32(t *testing.T) {
	m := NewAdjMxDf32(testSizeSetDel, nil)
	t.Run("is WGf32", func(t *testing.T) { testIsWGf32(t, m) })
	const w32 = float32(3.1415)
	testDSetDel(t, m,
		func(i, j VIdx) { m.SetEdge(i, j, NaN32()) },
		func(w interface{}) bool { return IsNaN32(w.(float32)) },
		func(i, j VIdx) interface{} { m.SetEdge(i, j, w32); return w32 },
		func(i, j VIdx) interface{} { return m.Edge(i, j) },
	)
}

func BenchmarkAdjMxDf32(b *testing.B) {
	m := NewAdjMxDf32(testSizeSetDel, nil)
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

func BenchmarkAdjMxDf32_generic(b *testing.B) {
	m := NewAdjMxDf32(testSizeSetDel, nil)
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

func TestAdjMxUbool(t *testing.T) {
	u := NewAdjMxUbool(testSizeSetDel, nil)
	t.Run("is WUbool", func(t *testing.T) { testIsWUbool(t, u) })
	testUSetDel(t, u,
		func(i, j VIdx) { u.SetEdgeU(i, j, false) },
		func(w interface{}) bool { return w == nil },
		func(i, j VIdx) interface{} { u.SetEdgeU(i, j, testProbeBool); return testProbeBool },
		func(i, j VIdx) interface{} { return u.Weight(i, j) },
	)
}

func TestAdjMxUi32(t *testing.T) {
	u := NewAdjMxUi32(testSizeSetDel, I32Del, nil)
	t.Run("is WUi32", func(t *testing.T) { testIsWUi32(t, u) })
	testUSetDel(t, u,
		func(i, j VIdx) { u.SetEdgeU(i, j, u.del) },
		func(w interface{}) bool { return w.(int32) == u.del },
		func(i, j VIdx) interface{} { u.SetEdgeU(i, j, testProbeI32); return testProbeI32 },
		func(i, j VIdx) interface{} {
			w, _ := u.Edge(i, j)
			return w
		},
	)
}

func TestAdjMxUf32(t *testing.T) {
	m := NewAdjMxUf32(testSizeSetDel, nil)
	t.Run("is WUf32", func(t *testing.T) { testIsWUf32(t, m) })
	const w32 = float32(3.1415)
	testUSetDel(t, m,
		func(i, j VIdx) { m.SetEdgeU(i, j, NaN32()) },
		func(w interface{}) bool { return IsNaN32(w.(float32)) },
		func(i, j VIdx) interface{} { m.SetEdgeU(i, j, w32); return w32 },
		func(i, j VIdx) interface{} { return m.Edge(i, j) },
	)
}
