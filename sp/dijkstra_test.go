package sp

import (
	"math"
	"math/rand"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
)

func TestDijkstrai32_toFW(t *testing.T) {
	const VNo = 11
	g := groph.NewAdjMxUi32(VNo, nil)
	for i := 0; i < VNo; i++ {
		g.SetEdge(i, i, 0)
		for j := i + 1; j < VNo; j++ {
			if rand.Intn(100) < 20 {
				g.SetEdge(i, j, g.Del)
			} else {
				g.SetEdge(i, j, int32(1+rand.Intn(20)))
			}
		}
	}
	fwDs := groph.NewAdjMxDi32(VNo, nil)
	fwDs.Del = g.Del
	groph.CpWeights(fwDs, g)
	FloydWarshallAdjMxDi32(fwDs)
	var (
		dijkstra Dijkstrai32
		dist     []int32
	)
	for start := groph.VIdx(0); start < VNo; start++ {
		dist, _ := dijkstra.On(g, start, dist, nil)
		for dest := groph.VIdx(0); dest < VNo; dest++ {
			if start == dest {
				continue
			}
			dfw, _ := fwDs.Edge(start, dest)
			ddj := dist[dest]
			if ddj != dfw {
				t.Errorf("dist %d->%d: F/W=%d, Dijkstra=%d", start, dest, dfw, ddj)
			}
		}
	}
}

func TestDijkstrai32_paths(t *testing.T) {
	const VNo = 11
	g := groph.NewAdjMxUi32(VNo, nil)
	for i := 0; i < VNo; i++ {
		g.SetEdge(i, i, 0)
		for j := i + 1; j < VNo; j++ {
			if rand.Intn(100) < 20 {
				g.SetEdge(i, j, g.Del)
			} else {
				g.SetEdge(i, j, int32(1+rand.Intn(20)))
			}
		}
	}
	var (
		dijkstra Dijkstrai32
		dist     []int32
		path     = []groph.VIdx{}
	)
	for start := groph.VIdx(0); start < VNo; start++ {
		dist, path = dijkstra.On(g, start, dist, path)
		for dest := groph.VIdx(0); dest < VNo; dest++ {
			if dest == start {
				continue
			}
			current := dest
			len := int32(0)
			for pred := path[current]; pred >= 0; pred = path[current] {
				d, ok := g.Edge(pred, current)
				if !ok {
					t.Fatalf("inexistent edge in path")
				}
				len += d
				current = pred
			}
			if len != dist[dest] {
				t.Errorf("len(%d;%d): expect %d, got %d", start, dest, dist[dest], len)
			}
		}
	}
}

func TestDijkstraf32_toFW(t *testing.T) {
	const VNo = 11
	g := groph.NewAdjMxUf32(VNo, nil)
	for i := 0; i < VNo; i++ {
		g.SetEdge(i, i, 0)
		for j := i + 1; j < VNo; j++ {
			if rand.Intn(100) < 20 {
				g.SetEdge(i, j, float32(math.Inf(1)))
			} else {
				g.SetEdge(i, j, 1+20*rand.Float32())
			}
		}
	}
	fwDs := groph.NewAdjMxDf32(VNo, nil)
	groph.CpWeights(fwDs, g)
	FloydWarshallAdjMxDf32(fwDs)
	var (
		dijkstra Dijkstraf32
		dist     []float32
	)
	for start := groph.VIdx(0); start < VNo; start++ {
		dist, _ := dijkstra.On(g, start, dist, nil)
		for dest := groph.VIdx(0); dest < VNo; dest++ {
			if start == dest {
				continue
			}
			dfw := fwDs.Edge(start, dest)
			ddj := dist[dest]
			if math.Abs(float64(ddj-dfw)) > 0.00001 {
				t.Errorf("dist %d->%d: F/W=%f, Dijkstra=%f", start, dest, dfw, ddj)
			}
		}
	}
}

func TestDijkstraf32_paths(t *testing.T) {
	const VNo = 11
	g := groph.NewAdjMxUf32(VNo, nil)
	for i := 0; i < VNo; i++ {
		g.SetEdge(i, i, 0)
		for j := i + 1; j < VNo; j++ {
			if rand.Intn(100) < 20 {
				g.SetEdge(i, j, float32(math.Inf(1)))
			} else {
				g.SetEdge(i, j, 1+20*rand.Float32())
			}
		}
	}
	var (
		dijkstra Dijkstraf32
		dist     []float32
		path     = []groph.VIdx{}
	)
	for start := groph.VIdx(0); start < VNo; start++ {
		dist, path = dijkstra.On(g, start, dist, path)
		for dest := groph.VIdx(0); dest < VNo; dest++ {
			if dest == start {
				continue
			}
			current := dest
			len := float32(0)
			for pred := path[current]; pred >= 0; pred = path[current] {
				len += g.Edge(pred, current)
				current = pred
			}
			if math.Abs(float64(len-dist[dest])) > 0.00001 {
				t.Errorf("len(%d;%d): expect %f, got %f", start, dest, dist[dest], len)
			}
		}
	}
}
