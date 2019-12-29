package groph

import (
	"fmt"
	"math"
	"testing"
)

var _ WGbool = (*AdjMxDbitmap)(nil)
var _ WGbool = (*AdjMxDbool)(nil)
var _ WGi32 = (*AdjMxDi32)(nil)
var _ WGf32 = (*AdjMxDf32)(nil)
var _ WUi32 = (*AdjMxUi32)(nil)
var _ WUf32 = (*AdjMxUf32)(nil)

func TestAdjMxDbitmap_SetUset(t *testing.T) {
	m := NewAdjMxDbitmap(testSizeSetUnset, nil)
	testDSetUnset(t, m,
		func(i, j VIdx) { m.SetEdge(i, j, false) },
		func(w interface{}) bool { return w.(bool) == false },
		func(i, j VIdx) interface{} { m.SetEdge(i, j, true); return true },
		func(i, j VIdx) interface{} { return m.Edge(i, j) },
	)
}

func BenchmarkAdjMxDbitmap(b *testing.B) {
	m := NewAdjMxDbitmap(testSizeSetUnset, nil)
	max := m.VertexNo()
	for n := 0; n < b.N; n++ {
		w := true
		if n&1 == 0 {
			w = false
		}
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

func BenchmarkAdjMxDbitmap_generic(b *testing.B) {
	m := NewAdjMxDbitmap(testSizeSetUnset, nil)
	max := m.VertexNo()
	for n := 0; n < b.N; n++ {
		w := true
		if n&1 == 0 {
			w = false
		}
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

func TestAdjMxDbool_SetUset(t *testing.T) {
	m := NewAdjMxDbool(testSizeSetUnset, nil)
	testDSetUnset(t, m,
		func(i, j VIdx) { m.SetEdge(i, j, false) },
		func(w interface{}) bool { return w.(bool) == false },
		func(i, j VIdx) interface{} { m.SetEdge(i, j, true); return true },
		func(i, j VIdx) interface{} { return m.Edge(i, j) },
	)
}

func BenchmarkAdjMxDbool(b *testing.B) {
	m := NewAdjMxDbool(testSizeSetUnset, nil)
	max := m.VertexNo()
	for n := 0; n < b.N; n++ {
		w := true
		if n&1 == 0 {
			w = false
		}
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

func BenchmarkAdjMxDbool_generic(b *testing.B) {
	m := NewAdjMxDbool(testSizeSetUnset, nil)
	max := m.VertexNo()
	for n := 0; n < b.N; n++ {
		w := true
		if n&1 == 0 {
			w = false
		}
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

func TestAdjMxDi32_SetUset(t *testing.T) {
	m := NewAdjMxDi32(testSizeSetUnset, nil)
	const w32 = int32(4711)
	testDSetUnset(t, m,
		func(i, j VIdx) { m.SetEdge(i, j, m.Del) },
		func(w interface{}) bool { return w == nil },
		func(i, j VIdx) interface{} { m.SetEdge(i, j, w32); return w32 },
		func(i, j VIdx) interface{} { return m.Weight(i, j) },
	)
}

func BenchmarkAdjMxDi32(b *testing.B) {
	m := NewAdjMxDi32(testSizeSetUnset, nil)
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

func BenchmarkAdjMxDi32_generic(b *testing.B) {
	m := NewAdjMxDi32(testSizeSetUnset, nil)
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

func TestAdjMxDf32_SetUset(t *testing.T) {
	m := NewAdjMxDf32(testSizeSetUnset, nil)
	const w32 = float32(3.1415)
	testDSetUnset(t, m,
		func(i, j VIdx) { m.SetEdge(i, j, NaN32()) },
		func(w interface{}) bool { return IsNaN32(w.(float32)) },
		func(i, j VIdx) interface{} { m.SetEdge(i, j, w32); return w32 },
		func(i, j VIdx) interface{} { return m.Edge(i, j) },
	)
}

func BenchmarkAdjMxDf32(b *testing.B) {
	m := NewAdjMxDf32(testSizeSetUnset, nil)
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

func BenchmarkAdjMxDf32_generic(b *testing.B) {
	m := NewAdjMxDf32(testSizeSetUnset, nil)
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

func TestAdjMxUf32_SetUset(t *testing.T) {
	m := NewAdjMxUf32(testSizeSetUnset, nil)
	const w32 = float32(3.1415)
	testUSetUnset(t, m,
		func(i, j VIdx) { m.SetEdgeU(i, j, NaN32()) },
		func(w interface{}) bool { return IsNaN32(w.(float32)) },
		func(i, j VIdx) interface{} { m.SetEdgeU(i, j, w32); return w32 },
		func(i, j VIdx) interface{} { return m.Edge(i, j) },
	)
}

func TestAdjMxUi32_SetUset(t *testing.T) {
	m := NewAdjMxUi32(testSizeSetUnset, nil)
	const w32 int32 = 31415
	testUSetUnset(t, m,
		func(i, j VIdx) { m.SetEdgeU(i, j, m.Del) },
		func(w interface{}) bool { return w.(int32) == m.Del },
		func(i, j VIdx) interface{} { m.SetEdgeU(i, j, w32); return w32 },
		func(i, j VIdx) interface{} {
			w, _ := m.Edge(i, j)
			return w
		},
	)
}

func ExampleNaN64() {
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

func ExampleNaN32() {
	nan := NaN32()
	fmt.Printf("0 < NaN: %t\n", 0 < nan)
	fmt.Printf("0 > NaN: %t\n", 0 > nan)
	fmt.Printf("nan isNan(): %t\n", IsNaN32(nan))
	tmp := 3.14 + nan
	fmt.Println("tmp := 3.14 + nan")
	fmt.Printf("tmp isNan(): %t\n", IsNaN32(tmp))
	fmt.Printf("tmp == nan : %t\n", tmp == nan)
	// Output:
	// 0 < NaN: false
	// 0 > NaN: false
	// nan isNan(): true
	// tmp := 3.14 + nan
	// tmp isNan(): true
	// tmp == nan : false
}
