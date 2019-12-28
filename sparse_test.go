package groph

import (
	"testing"
)

var (
	_ WGraph          = (*SpMap)(nil)
	_ NeighbourLister = (*SpMap)(nil)
	_ WGf32           = (*SpMapf32)(nil)
	_ NeighbourLister = (*SpMapf32)(nil)
	_ WGi32           = (*SpMapi32)(nil)
	_ NeighbourLister = (*SpMapi32)(nil)
)

func TestSpMap_SetUset(t *testing.T) {
	m := NewSpMap(testSizeSetUnset, nil)
	testGenericSetUnset(t, m, 4)
}

func BenchmarkSpMap_generic(b *testing.B) {
	m := NewSpMap(testSizeSetUnset, nil)
	max := m.VertexNo()
	w := 0
	for n := 0; n < b.N; n++ {
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
		w++
	}
}

func TestSpMapi32_SetUset(t *testing.T) {
	m := NewSpMapi32(testSizeSetUnset, nil)
	testGenericSetUnset(t, m, int32(4))
}

func BenchmarkSpMapf32(b *testing.B) {
	m := NewSpMapf32(testSizeSetUnset, nil)
	max := m.VertexNo()
	w := float32(0)
	for n := 0; n < b.N; n++ {
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
		w += 0.1
	}
}

func TestSpMapf32_SetUset(t *testing.T) {
	m := NewSpMapf32(testSizeSetUnset, nil)
	testGenericSetUnset(t, m, float32(2.7182))
}
