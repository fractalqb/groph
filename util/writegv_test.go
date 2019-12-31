package util

import (
	"os"

	"git.fractalqb.de/fractalqb/groph"
)

func ExampleGraphViz_Write() {
	g := groph.MustCp(groph.CpWeights(
		groph.NewAdjMxDi32(4, 0, nil),
		groph.NewWeightsSlice([]int32{
			0, 8, 0, 1,
			0, 0, 1, 0,
			4, 0, 7, 0,
			0, 2, 9, 0,
		}).Must(),
	))
	gv := GraphViz{}
	gv.Write(os.Stdout, g, "Groph")
	// Output:
	// digraph Groph {
	//   0 -> 1 [label="8"];
	//   0 -> 3 [label="1"];
	//   1 -> 2 [label="1"];
	//   2 -> 0 [label="4"];
	//   2 -> 2 [label="7"];
	//   3 -> 1 [label="2"];
	//   3 -> 2 [label="9"];
	// }
}
