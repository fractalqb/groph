package minspantree

import (
	"fmt"

	"git.fractalqb.de/fractalqb/groph"
)

func ExampleKruskal() {
	// See https://de.wikipedia.org/wiki/Algorithmus_von_Kruskal
	g := groph.NewAdjMxUf32(7, nil).Init(groph.NaN32())
	g.SetEdge(0, 1, 7)
	g.SetEdge(0, 3, 5)
	g.SetEdge(1, 2, 8)
	g.SetEdge(1, 3, 9)
	g.SetEdge(1, 4, 7)
	g.SetEdge(2, 4, 5)
	g.SetEdge(3, 4, 15)
	g.SetEdge(3, 5, 6)
	g.SetEdge(4, 5, 8)
	g.SetEdge(4, 6, 9)
	g.SetEdge(5, 6, 11)
	mst, err := Kruskalf32(g, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len: %d\n", len(mst))
	for i, e := range mst {
		fmt.Printf("%d: %d -- %d\n", i, e.U, e.V)
	}
	// Output:
	// len: 6
	// 0: 0 -- 3
	// 1: 2 -- 4
	// 2: 3 -- 5
	// 3: 0 -- 1
	// 4: 1 -- 4
	// 5: 4 -- 6
}
