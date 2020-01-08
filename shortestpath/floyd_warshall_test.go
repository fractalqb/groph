package shortestpath

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/util"
)

func ExampleFloydWarshalli32() {
	fwres, _ := groph.AsAdjMxDi32(nil, 0, []int32{
		0, 8, 0, 1,
		0, 0, 1, 0,
		4, 0, 0, 0,
		0, 2, 9, 0,
	})
	FloydWarshalli32(fwres)
	sz := fwres.Order()
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			e, _ := fwres.Edge(i, j)
			if j == 0 {
				fmt.Printf("%d", e)
			} else {
				fmt.Printf(" %d", e)
			}
		}
		fmt.Println()
	}
	// Output:
	// 8 3 4 1
	// 5 8 1 6
	// 4 7 8 5
	// 7 2 3 8
}

func ExampleFloydWarshallf32() {
	graph := util.NewWeightsSlice([]int{
		0, 8, 0, 1,
		0, 0, 1, 0,
		4, 0, 0, 0,
		0, 2, 9, 0,
	}).Must()
	sz := graph.Order()
	fwres := groph.NewAdjMxDf32(sz, nil)
	groph.CpXWeights(fwres, graph, func(i interface{}) interface{} {
		e := i.(int)
		if i == 0 {
			return float32(math.Inf(1))
		}
		return float32(e)
	})
	FloydWarshallAdjMxDf32(fwres)
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			if j == 0 {
				fmt.Printf("%d", int(fwres.Edge(i, j)))
			} else {
				fmt.Printf(" %d", int(fwres.Edge(i, j)))
			}
		}
		fmt.Println()
	}
	// Output:
	// 8 3 4 1
	// 5 8 1 6
	// 4 7 8 5
	// 7 2 3 8
}

func TestFloydWarshallDirEqUndir(t *testing.T) {
	const VNO = 7
	mu := groph.NewAdjMxUf32(VNO, nil)
	md := groph.NewAdjMxDf32(mu.Order(), nil)
	for i := 0; i < VNO; i++ {
		mu.SetEdge(i, i, 0)
		for j := i + 1; j < VNO; j++ {
			w := rand.Float32()
			mu.SetEdge(i, j, w)
		}
	}
	groph.CpWeights(md, mu)
	FloydWarshallf32(md)
	FloydWarshallf32(mu)
	for i := 0; i < VNO; i++ {
		for j := 0; j < VNO; j++ {
			if md.Edge(i, j) != mu.Edge(i, j) {
				t.Errorf("differ@ %d,%d: %f != %f", i, j, md.Edge(i, j), mu.Edge(i, j))
			}
		}
	}
}
