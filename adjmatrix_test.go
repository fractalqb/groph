package groph

import (
	"fmt"
	"math"
	"testing"
)

func TestAdjMxDbool_SetUset(t *testing.T) {
	m := NewAdjMxDbool(66, nil)
	testSetUnset(m, true, t)
}

func TestAdjMxDf32_SetUset(t *testing.T) {
	m := NewAdjMxDf32(7, nil)
	testSetUnset(m, float32(3.14), t)
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
