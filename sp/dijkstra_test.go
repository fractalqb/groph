package sp

import (
	"math/rand"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
)

func TestDijkstraf32_toFW(t *testing.T) {
	const VNo = 5
	g := groph.NewAdjMxUf32(VNo, nil)
	for i := 0; i < VNo; i++ {
		g.SetEdge(i, i, 0)
		for j := i + 1; j < VNo; j++ {
			g.SetEdge(i, j, 1+20*rand.Float32())
		}
	}
	fwDs := groph.NewAdjMxDf32(VNo, nil)
	groph.CpWeights(fwDs, g)
	FloydWarshallAdjMxDf32(fwDs)
	for start := groph.VIdx(0); start < VNo; start++ {
		ds, _ := Dijkstraf32(g, start)
		for dest := groph.VIdx(0); dest < VNo; dest++ {
			if start == dest {
				continue
			}
			dfw := fwDs.Edge(start, dest)
			ddj := ds[dest]
			if ddj != dfw {
				t.Errorf("dist %d->%d: F/W=%f, Dijkstra=%f", start, dest, dfw, ddj)
			}
		}
	}
}
