package groph

import (
	"testing"
)

var (
	_ WGraph          = (*SpMoM)(nil)
	_ NeighbourLister = (*SpMoM)(nil)
	_ WGf32           = (*SpMoMf32)(nil)
	_ NeighbourLister = (*SpMoMf32)(nil)
	_ WGi32           = (*SpMoMi32)(nil)
	_ NeighbourLister = (*SpMoMi32)(nil)
)

func TestSpMoM_SetUset(t *testing.T) {
	m := NewSpMoM(testSizeSetUnset, nil)
	testGenericSetUnset(t, m, 4)
}

func BenchmarkSpMoM_generic(b *testing.B) {
	m := NewSpMoM(testSizeSetUnset, nil)
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

func TestSpMoMi32_SetUset(t *testing.T) {
	m := NewSpMoMi32(testSizeSetUnset, nil)
	testGenericSetUnset(t, m, int32(4))
}

func BenchmarkSpMoMf32(b *testing.B) {
	m := NewSpMoMf32(testSizeSetUnset, nil)
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
				r := m.Edge(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
	}
}

func TestSpMoMf32_SetUset(t *testing.T) {
	m := NewSpMoMf32(testSizeSetUnset, nil)
	testGenericSetUnset(t, m, float32(2.7182))
}
