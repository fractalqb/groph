package groph

import (
	"testing"
)

var (
	_ WGraph          = (*SpSoM)(nil)
	_ NeighbourLister = (*SpSoM)(nil)
	_ WGi32           = (*SpSoMi32)(nil)
	_ NeighbourLister = (*SpSoMi32)(nil)
	_ WGf32           = (*SpSoMf32)(nil)
	_ NeighbourLister = (*SpSoMf32)(nil)
)

func TestSpSoM_SetUset(t *testing.T) {
	m := NewSpSoM(testSizeSetUnset, nil)
	testGenericSetUnset(t, m, 4)
}

func TestSpSoM_SetUsetUndir(t *testing.T) {
	u := AsWUndir(NewSpSoM(testSizeSetUnset, nil))
	testGenericSetUnset(t, u, 4)
}

func BenchmarkSpSoM_generic(b *testing.B) {
	m := NewSpSoM(testSizeSetUnset, nil)
	max := m.VertexNo()
	for n := 0; n < b.N; n++ {
		for i := VIdx(0); i < max; i++ {
			for j := VIdx(0); j < max; j++ {
				m.SetWeight(i, j, n)
			}
		}
		for i := VIdx(0); i < max; i++ {
			for j := VIdx(0); j < max; j++ {
				r := m.Weight(i, j)
				if r != n {
					b.Fatal("unexpected read", n, r)
				}
			}
		}
	}
}

func TestSpSoMi32_SetUset(t *testing.T) {
	m := NewSpSoMi32(testSizeSetUnset, nil)
	testGenericSetUnset(t, m, int32(4))
}

func TestSpSoMi32_SetUsetUndir(t *testing.T) {
	u := AsWUndir(NewSpSoMi32(testSizeSetUnset, nil))
	testGenericSetUnset(t, u, int32(4))
}

func BenchmarkSpSoMi32_generic(b *testing.B) {
	m := NewSpSoMi32(testSizeSetUnset, nil)
	max := m.VertexNo()
	for n := 0; n < b.N; n++ {
		w := int32(n)
		for i := VIdx(0); i < max; i++ {
			for j := VIdx(0); j < max; j++ {
				m.SetWeight(i, j, w)
			}
		}
		for i := VIdx(0); i < max; i++ {
			for j := VIdx(0); j < max; j++ {
				r := m.Weight(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func BenchmarkSpSoMi32(b *testing.B) {
	m := NewSpSoMi32(testSizeSetUnset, nil)
	max := m.VertexNo()
	for n := 0; n < b.N; n++ {
		w := int32(n)
		for i := VIdx(0); i < max; i++ {
			for j := VIdx(0); j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := VIdx(0); i < max; i++ {
			for j := VIdx(0); j < max; j++ {
				r, _ := m.Edge(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func TestSpSoMf32_SetUset(t *testing.T) {
	m := NewSpSoMf32(testSizeSetUnset, nil)
	testGenericSetUnset(t, m, float32(4))
}

func TestSpSoMf32_SetUsetUndir(t *testing.T) {
	u := AsWUndir(NewSpSoMf32(testSizeSetUnset, nil))
	testGenericSetUnset(t, u, float32(4))
}

func BenchmarkSpSoMf32_generic(b *testing.B) {
	m := NewSpSoMf32(testSizeSetUnset, nil)
	max := m.VertexNo()
	for n := 0; n < b.N; n++ {
		w := float32(n)
		for i := VIdx(0); i < max; i++ {
			for j := VIdx(0); j < max; j++ {
				m.SetWeight(i, j, w)
			}
		}
		for i := VIdx(0); i < max; i++ {
			for j := VIdx(0); j < max; j++ {
				r := m.Weight(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func BenchmarkSpSoMf32(b *testing.B) {
	m := NewSpSoMf32(testSizeSetUnset, nil)
	max := m.VertexNo()
	for n := 0; n < b.N; n++ {
		w := float32(n)
		for i := VIdx(0); i < max; i++ {
			for j := VIdx(0); j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := VIdx(0); i < max; i++ {
			for j := VIdx(0); j < max; j++ {
				if r := m.Edge(i, j); r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}
