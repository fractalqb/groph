// Copyright 2022 Marcus Perlick
// This file is part of Go module git.fractalqb.de/fractalqb/groph
//
// groph is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// groph is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with groph.  If not, see <http://www.gnu.org/licenses/>.

package paths

import (
	"fmt"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmtx"
	"git.fractalqb.de/fractalqb/groph/graphs"
)

type point = graphs.Point[float64]

var tspExample1 = graphs.Euclidean{
	point{0, 0},
	point{10, 10},
	point{2, 9},
	point{4, 5},
	point{8, 3},
}

func ExampleGreedyTSP() {
	am := adjmtx.NewUndirected(tspExample1.Order(), tspExample1.NotEdge(), nil)
	groph.Copy[float64](am, tspExample1)
	testShowMatrix(am)
	w, l := GreedyTSP[float64](am)
	fmt.Printf("%v %.2f", w, l)
	// Output:
	// Matrix:
	//  0:  0.00, 14.14,  9.22,  6.40,  8.54
	//  1: 14.14,  0.00,  8.06,  7.81,  7.28
	//  2:  9.22,  8.06,  0.00,  4.47,  8.49
	//  3:  6.40,  7.81,  4.47,  0.00,  4.47
	//  4:  8.54,  7.28,  8.49,  4.47,  0.00
	// [0 3 2 1 4] 34.76
}

func testShowMatrix(g groph.RGraph[float64]) {
	fmt.Println("Matrix:")
	ord := g.Order()
	for i := groph.VIdx(0); i < ord; i++ {
		fmt.Printf("%2d: ", i)
		for j := groph.VIdx(0); j < ord; j++ {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%5.2f", g.Edge(i, j))
		}
		fmt.Println()
	}
}

var tspExample2 = graphs.Euclidean{
	point{0, 0},
	point{10, 10},
	point{2, 9},
	point{4, 5},
	point{8, 3},
	point{9, 2},
	point{5, 4},
	point{3, 8},
}

func BenchmarkGreedyTSP(b *testing.B) {
	am := adjmtx.NewUndirected(tspExample2.Order(), tspExample2.NotEdge(), nil)
	groph.Copy[float64](am, tspExample2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GreedyTSP[float64](am)
	}
}
