package traverse

import (
	"fmt"

	"git.fractalqb.de/fractalqb/groph"
)

func ExampleTraversal_Depth1stAt() {
	g := groph.NewAdjMxDbool(7, nil)
	type E = groph.Edge
	groph.Set(g, true,
		E{0, 1}, E{0, 2}, E{0, 3},
		E{1, 4}, E{1, 5},
		E{2, 5},
		E{3, 6},
	)
	hits := NewTraversal(g).Depth1stAt(0, func(v groph.VIdx) {
		fmt.Printf(" %d", v)
	})
	fmt.Println("\nhits:", hits)
	// Output:
	// 0 3 6 2 5 1 4
	// hits: 7
}

func ExampleTraversal_Breadth1stAt() {
	g := groph.NewAdjMxDbool(7, nil)
	type E = groph.Edge
	groph.Set(g, true,
		E{0, 1}, E{0, 2}, E{0, 3},
		E{1, 4}, E{1, 5},
		E{2, 5},
		E{3, 6},
	)
	hits := NewTraversal(g).Breadth1stAt(0, func(v groph.VIdx) {
		fmt.Printf(" %d", v)
	})
	fmt.Println("\nhits:", hits)
	// Output:
	// 0 1 2 3 4 5 6
	// hits: 7
}
