package minspantree

import (
	"errors"
	"math"
	"sort"

	"git.fractalqb.de/fractalqb/groph"
)

func sortedEdges(g groph.RGf32) (res []groph.Edge) {
	vno := g.Order()
	for i := 0; i < vno; i++ {
		for j := i + 1; j < vno; j++ {
			if !math.IsNaN(float64(g.Edge(i, j))) {
				res = append(res, groph.Edge{U: i, V: j})
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		e1, e2 := &res[i], &res[j]
		return g.Edge(e1.U, e1.V) < g.Edge(e2.U, e2.V)
	})
	return res
}

// TODO more efficient way for bookkeping of connected sets?
func retag(f map[groph.VIdx]groph.VIdx, oldTag, newTag groph.VIdx) {
	for v, t := range f {
		if t == oldTag {
			f[v] = newTag
		}
	}
}

func Kruskalf32(g groph.RGf32, mst []groph.Edge) ([]groph.Edge, error) {
	if groph.Directed(g) {
		return mst, errors.New("cannot apply Kruskal's algorithm to directed graphs")
	}
	mst = mst[:0]
	ebo := sortedEdges(g)
	frs := make(map[groph.VIdx]groph.VIdx)
	vc := 0
	for _, e := range ebo {
		ti, iOk := frs[e.U]
		tj, jOk := frs[e.V]
		if iOk {
			if jOk { // no new vertex
				if ti != tj {
					retag(frs, ti, tj)
					mst = append(mst, e)
				}
			} else { // j is new vertex
				frs[e.V] = ti
				mst = append(mst, e)
				vc++
				if vc == g.Order() {
					return mst, nil
				}
			}
		} else if jOk { // i is new vertex
			frs[e.U] = tj
			mst = append(mst, e)
			vc++
			if vc == g.Order() {
				return mst, nil
			}
		} else { // i & j are new vertices
			frs[e.U] = e.U
			frs[e.V] = e.U
			mst = append(mst, e)
			vc += 2
			if vc == g.Order() {
				return mst, nil
			}
		}
	}
	return mst, nil
}
