package util

import (
	"fmt"

	"git.fractalqb.de/fractalqb/groph"
)

func ExampleReorderPath() {
	data := []string{"a", "b", "c", "d", "e"}
	path := []groph.VIdx{1, 3, 0, 4, 2}
	ReorderPath(data, path)
	fmt.Println(data)
	// Output:
	// [b d a e c]
}
