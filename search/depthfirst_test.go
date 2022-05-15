package search

import (
	"fmt"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmtx"
)

func ExampleDepthFirst_Directed() {
	g := adjmtx.NewDirected(7, false, nil)
	groph.Set[bool](g, true,
		0, 1, 0, 2, 0, 3,
		1, 4, 1, 5,
		2, 5,
		3, 6,
	)
	var dfs DepthFirst[bool]
	dfs.Directed(g, 0, func(v groph.VIdx) bool {
		fmt.Printf(" %d", v)
		return false
	})
	fmt.Println()
	// Output:
	// 0 3 6 2 5 1 4
}

func ExampleDepthFirst_Undirected() {
	g := adjmtx.NewUndirected(7, false, nil)
	groph.Set[bool](g, true,
		0, 1, 0, 2, 0, 3,
		1, 4, 1, 5,
		2, 5,
		3, 6,
	)
	var dfs DepthFirst[bool]
	dfs.Undirected(g, 0, func(v groph.VIdx) bool {
		fmt.Printf(" %d", v)
		return false
	})
	fmt.Println()
	// Output:
	// 0 3 6 2 5 1 4
}
