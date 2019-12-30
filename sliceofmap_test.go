package groph

import (
	"testing"
)

var (
	_ WGraph    = (*SpSoM)(nil)
	_ OutLister = (*SpSoM)(nil)
	_ WGi32     = (*SpSoMi32)(nil)
	_ OutLister = (*SpSoMi32)(nil)
	_ WGf32     = (*SpSoMf32)(nil)
	_ OutLister = (*SpSoMf32)(nil)
)

func TestSpSoM(t *testing.T) {
	g := NewSpSoM(testSizeSetDel, nil)
	testGenericSetDel(t, g, testProbeGen)
}

func TestSpSoM_undir(t *testing.T) {
	//u := AsWUndir(NewSpSoM(testSizeSetDel, nil))
	t.Skip("NYI!")
}

func BenchmarkSpSoM_generic(b *testing.B) {
	m := NewSpSoM(testSizeSetDel, nil)
	max := m.VertexNo()
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

func TestSpSoMi32(t *testing.T) {
	m := NewSpSoMi32(testSizeSetDel, nil)
	t.Run("is WGi32", func(t *testing.T) { testIsWGi32(t, m) })
	// TODO typed access

}

func TestSpSoMi32_undir(t *testing.T) {
	u := AsWUndir(NewSpSoMi32(testSizeSetDel, nil))
	t.Run("generic access", func(t *testing.T) {
		testGenericSetDel(t, u, testProbeI32)
	})
	// TODO is WUi32 & typed access

}

func BenchmarkSpSoMi32_generic(b *testing.B) {
	m := NewSpSoMi32(testSizeSetDel, nil)
	max := m.VertexNo()
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

func BenchmarkSpSoMi32(b *testing.B) {
	m := NewSpSoMi32(testSizeSetDel, nil)
	max := m.VertexNo()
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

func TestSpSoMf32(t *testing.T) {
	g := NewSpSoMf32(testSizeSetDel, nil)
	t.Run("is WGf32", func(t *testing.T) { testIsWGf32(t, g) })
	// TODO typed access
}

func TestSpSoMf32_undir(t *testing.T) {
	u := AsWUndir(NewSpSoMf32(testSizeSetDel, nil))
	t.Run("generic access", func(t *testing.T) {
		testGenericSetDel(t, u, testProbeF32)
	})
	// TODO is WUf32 & typed access
}

func BenchmarkSpSoMf32_generic(b *testing.B) {
	m := NewSpSoMf32(testSizeSetDel, nil)
	max := m.VertexNo()
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

func BenchmarkSpSoMf32(b *testing.B) {
	m := NewSpSoMf32(testSizeSetDel, nil)
	max := m.VertexNo()
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
