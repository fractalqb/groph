package paths

import (
	"reflect"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmtx"
	"git.fractalqb.de/fractalqb/groph/graphs"
)

type e = graphs.Edge

func TestAStarU(t *testing.T) {
	const (
		frankfurt = iota
		würzburg
		ludwigshafen
		kaiserslautern
		heilbronn
		saarbrücken
		karlsruhe
	)
	g := adjmtx.NewUndirected(7, -1, nil)
	groph.CopyX[int, any](g, graphs.UndirectedEdges{
		e{U: frankfurt, V: würzburg, W: 116},
		e{U: frankfurt, V: kaiserslautern, W: 103},
		e{U: würzburg, V: ludwigshafen, W: 183},
		e{U: würzburg, V: heilbronn, W: 102},
		e{U: kaiserslautern, V: ludwigshafen, W: 53},
		e{U: kaiserslautern, V: saarbrücken, W: 70},
		e{U: saarbrücken, V: karlsruhe, W: 145},
		e{U: heilbronn, V: karlsruhe, W: 84},
	}, func(w any) (int, error) {
		if w == nil {
			return -1, nil
		}
		return w.(int), nil
	})
	saarbrücken2würzburg := []int{96, 0, 108, 158, 87, 222, 140}
	p, l := AStarU[int](g, saarbrücken, würzburg, func(u, v groph.VIdx) int {
		return saarbrücken2würzburg[u]
	})
	bestPath := []int{saarbrücken, kaiserslautern, frankfurt, würzburg}
	if !reflect.DeepEqual(p, bestPath) {
		t.Error("wrong path", p, "want", bestPath)
	}
	if l != 289 {
		t.Error("wrong path len", l)
	}
}
