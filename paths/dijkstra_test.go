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
	"math/rand"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmtx"
)

func TestDijkstraiD_toFW(t *testing.T) {
	const NoE = 0
	const VNo = 11
	g := adjmtx.NewDirected[int32](VNo, NoE, nil)
	for i := 0; i < VNo; i++ {
		for j := 0; j < VNo; j++ {
			if i == j || rand.Intn(100) < 20 {
				g.DelEdge(i, j)
			} else {
				g.SetEdge(i, j, int32(1+rand.Intn(20)))
			}
		}
	}
	fwDs := adjmtx.NewDirected[int32](VNo, NoE, nil)
	groph.Copy[int32](fwDs, g)
	FloydWarshallD[int32](fwDs)
	for start := 0; start < VNo; start++ {
		dist := DijkstraD[int32](g, start)
		for dest := 0; dest < VNo; dest++ {
			if start == dest {
				continue
			}
			dfw := fwDs.Edge(start, dest)
			ddj := dist.Edge(start, dest)
			if !fwDs.IsEdge(dfw) {
				if dist.IsEdge(ddj) {
					t.Errorf("Unexpected edge %d->%d in Disjkstra", start, dest)
				}
			} else if !dist.IsEdge(ddj) {
				if dist.IsEdge(ddj) {
					t.Errorf("Missing edge %d->%d in Disjkstra", start, dest)
				}
			} else if ddj != dfw {
				t.Errorf("dist %d->%d: F/W=%d, Dijkstra=%d", start, dest, dfw, ddj)
			}
		}
	}
}

func TestDijkstraiU_toFW(t *testing.T) {
	const NoE = 0
	const VNo = 11
	g := adjmtx.NewUndirected[int32](VNo, NoE, nil)
	for i := 0; i < VNo; i++ {
		g.DelEdge(i, i)
		for j := i + 1; j < VNo; j++ {
			if rand.Intn(100) < 20 {
				g.DelEdge(i, j)
			} else {
				g.SetEdge(i, j, int32(1+rand.Intn(20)))
			}
		}
	}
	fwDs := adjmtx.NewUndirected[int32](VNo, NoE, nil)
	groph.Copy[int32](fwDs, g)
	FloydWarshallU[int32](fwDs)
	for start := 0; start < VNo; start++ {
		dist := DijkstraU[int32](g, start)
		for dest := 0; dest < VNo; dest++ {
			if start == dest {
				continue
			}
			dfw := fwDs.Edge(start, dest)
			ddj := dist.Edge(start, dest)
			if !fwDs.IsEdge(dfw) {
				if dist.IsEdge(ddj) {
					t.Errorf("Unexpected edge %d->%d in Disjkstra", start, dest)
				}
			} else if !dist.IsEdge(ddj) {
				if dist.IsEdge(ddj) {
					t.Errorf("Missing edge %d->%d in Disjkstra", start, dest)
				}
			} else if ddj != dfw {
				t.Errorf("dist %d->%d: F/W=%d, Dijkstra=%d", start, dest, dfw, ddj)
			}
		}
	}
}

// func TestDijkstrai32_paths(t *testing.T) {
// 	const VNo = 11
// 	g := adjmatrix.NewUInt32(VNo, adjmatrix.I32Del, nil)
// 	for i := 0; i < VNo; i++ {
// 		g.SetEdge(i, i, 0)
// 		for j := i + 1; j < VNo; j++ {
// 			if rand.Intn(100) < 20 {
// 				g.SetWeight(i, j, nil)
// 			} else {
// 				g.SetEdge(i, j, int32(1+rand.Intn(20)))
// 			}
// 		}
// 	}
// 	var (
// 		dijkstra DijkstraI32
// 		dist     []int32
// 		path     = []groph.VIdx{}
// 	)
// 	for start := 0; start < VNo; start++ {
// 		dist, path = dijkstra.On(g, start, dist, path)
// 		for dest := 0; dest < VNo; dest++ {
// 			if dest == start {
// 				continue
// 			}
// 			current := dest
// 			len := int32(0)
// 			for pred := path[current]; pred >= 0; pred = path[current] {
// 				d, ok := g.Edge(pred, current)
// 				if !ok {
// 					t.Fatalf("inexistent edge in path")
// 				}
// 				len += d
// 				current = pred
// 			}
// 			if len != dist[dest] {
// 				t.Errorf("len(%d;%d): expect %d, got %d", start, dest, dist[dest], len)
// 			}
// 		}
// 	}
// }
