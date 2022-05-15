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

package shortpath

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmtx"
)

func ExampleFloydWarshallD() {
	fwres, _ := adjmtx.AsDirected[int32](nil, 0,
		0, 8, 0, 1,
		0, 0, 1, 0,
		4, 0, 0, 0,
		0, 2, 9, 0,
	)
	FloydWarshallD[int32](fwres)
	sz := fwres.Order()
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			e := fwres.Edge(i, j)
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

func TestFloydWarshallDirEqUndir(t *testing.T) {
	const VNO = 7
	nan := float32(math.NaN())
	mu := adjmtx.NewUndirected[float32](VNO, nan, nil)
	md := adjmtx.NewDirected[float32](mu.Order(), nan, nil)
	for i := 0; i < VNO; i++ {
		mu.SetEdge(i, i, 0)
		for j := i + 1; j < VNO; j++ {
			w := rand.Float32()
			mu.SetEdge(i, j, w)
		}
	}
	groph.Copy[float32](md, mu)
	FloydWarshallD[float32](md)
	FloydWarshallU[float32](mu)
	for i := 0; i < VNO; i++ {
		for j := 0; j < VNO; j++ {
			if md.Edge(i, j) != mu.Edge(i, j) {
				t.Errorf("differ@ %d,%d: %f != %f", i, j, md.Edge(i, j), mu.Edge(i, j))
			}
		}
	}
}
