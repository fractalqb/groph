package util

import (
	"fmt"

	"git.fractalqb.de/fractalqb/groph"
)

var _ groph.RUndirected = (*RUadapter)(nil)

func ExampleRUadapter() {
	ws := groph.NewWeightsSlice([]int32{1, 2, 3, 4})
	u := groph.NewAdjMxUi32(ws.VertexNo(), nil)
	_, err := CpWeights(u, ws)
	fmt.Println("copy error:", err)
	rua := RUadapter{G: ws, Merge: MergeWeights(ws, MergeEqual)}
	_, err = CpWeights(u, &rua)
	fmt.Println("copy error:", err)
	fmt.Println("rua error:", rua.Err)
	// Output:
	// copy error: cannot copy from directed to undirected graph
	// copy error: <nil>
	// rua error: edges 1, 0: not equal: '3' / '2'
}
