package groph

import (
	"fmt"
)

func ExampleReorderPath() {
	data := []string{"a", "b", "c", "d", "e"}
	path := []VIdx{1, 3, 0, 4, 2}
	ReorderPath(data, path)
	fmt.Println(data)
	// Output:
	// [b d a e c]
}
