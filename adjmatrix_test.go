package groph

import (
	"fmt"
	"math"
	"testing"
)

func TestAdjMxDbool_SetUset(t *testing.T) {
	m := NewAdjMxDbool(testSizeSetUnset, nil)
	testDSetUnset(t, m,
		func(i, j uint) interface{} { m.SetEdge(i, j, true); return true },
		func(i, j uint) { m.SetEdge(i, j, false) },
		func(i, j uint) interface{} { return m.Edge(i, j) },
		func(w interface{}) bool { return w.(bool) == false },
	)
}

func BenchmarkAdjMxDbool(b *testing.B) {
	m := NewAdjMxDbool(testSizeSetUnset, nil)
	max := m.VertexNo()
	w := false
	for n := 0; n < b.N; n++ {
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				r := m.Edge(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
		w = !w
	}
}

func BenchmarkAdjMxDbool_generic(b *testing.B) {
	m := NewAdjMxDbool(testSizeSetUnset, nil)
	max := m.VertexNo()
	w := false
	for n := 0; n < b.N; n++ {
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				m.SetWeight(i, j, w)
			}
		}
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				r := m.Weight(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
		w = !w
	}
}

func TestAdjMxDi32_SetUset(t *testing.T) {
	m := NewAdjMxDi32(testSizeSetUnset, nil)
	const w32 = int32(4711)
	testDSetUnset(t, m,
		func(i, j uint) interface{} { m.SetEdge(i, j, w32); return w32 },
		func(i, j uint) { m.SetEdge(i, j, m.Cleared) },
		func(i, j uint) interface{} { return m.Weight(i, j) },
		func(w interface{}) bool { return w == nil },
	)
}

func BenchmarkAdjMxDi32(b *testing.B) {
	m := NewAdjMxDi32(testSizeSetUnset, nil)
	max := m.VertexNo()
	w := int32(0)
	for n := 0; n < b.N; n++ {
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				r, _ := m.Edge(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
		w++
	}
}

func BenchmarkAdjMxDi32_generic(b *testing.B) {
	m := NewAdjMxDi32(testSizeSetUnset, nil)
	max := m.VertexNo()
	w := int32(0)
	for n := 0; n < b.N; n++ {
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				m.SetWeight(i, j, w)
			}
		}
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				r := m.Weight(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
		w++
	}
}

func TestAdjMxDf32_SetUset(t *testing.T) {
	m := NewAdjMxDf32(testSizeSetUnset, nil)
	const w32 = float32(3.1415)
	testDSetUnset(t, m,
		func(i, j uint) interface{} { m.SetEdge(i, j, w32); return w32 },
		func(i, j uint) { m.SetEdge(i, j, nan32) },
		func(i, j uint) interface{} { return m.Edge(i, j) },
		func(w interface{}) bool { return math.IsNaN(float64(w.(float32))) },
	)
}

func BenchmarkAdjMxDf32(b *testing.B) {
	m := NewAdjMxDf32(testSizeSetUnset, nil)
	max := m.VertexNo()
	w := float32(0)
	for n := 0; n < b.N; n++ {
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				m.SetEdge(i, j, w)
			}
		}
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				r := m.Edge(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
		w += 0.1
	}
}

func BenchmarkAdjMxDf32_generic(b *testing.B) {
	m := NewAdjMxDf32(testSizeSetUnset, nil)
	max := m.VertexNo()
	w := float32(0)
	for n := 0; n < b.N; n++ {
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				m.SetWeight(i, j, w)
			}
		}
		for i := uint(0); i < max; i++ {
			for j := uint(0); j < max; j++ {
				r := m.Weight(i, j)
				if r != w {
					b.Fatal("unexpected read", w, r)
				}
			}
		}
		w += 0.1
	}
}

func ExampleNaNs() {
	nan := math.NaN()
	fmt.Printf("0 < NaN: %t\n", 0 < nan)
	fmt.Printf("0 > NaN: %t\n", 0 > nan)
	fmt.Printf("nan isNan(): %t\n", math.IsNaN(nan))
	tmp := 3.14 + nan
	fmt.Println("tmp := 3.14 + nan")
	fmt.Printf("tmp isNan(): %t\n", math.IsNaN(tmp))
	fmt.Printf("tmp == nan : %t\n", tmp == nan)
	// Output:
	// 0 < NaN: false
	// 0 > NaN: false
	// nan isNan(): true
	// tmp := 3.14 + nan
	// tmp isNan(): true
	// tmp == nan : false
}
