package util

import (
	"fmt"
	"os"

	"git.fractalqb.de/fractalqb/groph"
)

func ExampleGraphViz_Write() {
	g, _ := groph.AsAdjMxDi32(nil, 0, []int32{
		0, 8, 0, 1,
		0, 0, 1, 0,
		4, 0, 7, 0,
		0, 2, 9, 0,
	})
	gv := GraphViz{
		PerNodeAtts: func(g groph.RGraph, v groph.VIdx) map[string]interface{} {
			return map[string]interface{}{"label": fmt.Sprintf("Node %d", v)}
		},
		PerEdgeAtts: func(_ groph.RGraph, u, v groph.VIdx) map[string]interface{} {
			if u == 0 {
				return map[string]interface{}{"label": ""}
			}
			return map[string]interface{}{"label": fmt.Sprintf("w=%d", g.Weight(u, v))}
		},
	}
	gv.Write(os.Stdout, g, "Groph", "Zero-Node")
	// Output:
	// digraph Groph {
	//   0 [label="Zero-Node"];
	//   1 [label="Node 1"];
	//   2 [label="Node 2"];
	//   3 [label="Node 3"];
	//   0 -> 1;
	//   0 -> 3;
	//   1 -> 2 [label="w=1"];
	//   2 -> 0 [label="w=4"];
	//   2 -> 2 [label="w=7"];
	//   3 -> 1 [label="w=2"];
	//   3 -> 2 [label="w=9"];
	// }
}
