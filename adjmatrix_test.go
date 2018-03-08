package groph

import (
	"fmt"
	"math"
	"testing"
)

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

func BenchmarkIsNaNf64(b *testing.B) {
	v1, v2 := 3.14, math.NaN()
	for i := 0; i < b.N; i++ {
		if math.IsNaN(v1) == math.IsNaN(v2) {
			panic("epic fail!")
		}
	}
}

func BenchmarkIsNaNf32(b *testing.B) {
	var v1, v2 float32 = float32(3.14), float32(math.NaN())
	for i := 0; i < b.N; i++ {
		if math.IsNaN(float64(v1)) == math.IsNaN(float64(v2)) {
			panic("epic fail!")
		}
	}
}
