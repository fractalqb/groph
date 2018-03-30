package groph

import (
	"errors"
	"math"
	"sort"
)

func sortedEdges(g RGf32) (res []Edge) {
	vno := g.VertexNo()
	for i := uint(0); i < vno; i++ {
		for j := i + 1; j < vno; j++ {
			if !math.IsNaN(float64(g.Edge(i, j))) {
				res = append(res, Edge{i, j})
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		e1, e2 := &res[i], &res[j]
		return g.Edge(e1.I, e1.J) < g.Edge(e2.I, e2.J)
	})
	return res
}

// TODO more efficient way for bookkeping of connected sets?
func retag(f map[uint]uint, oldTag, newTag uint) {
	for v, t := range f {
		if t == oldTag {
			f[v] = newTag
		}
	}
}

func Kruskalf32(g RGf32, mst []Edge) ([]Edge, error) {
	if g.Directed() {
		return mst, errors.New("cannot apply Kruskal's algorithm to directed graphs")
	}
	mst = mst[:0]
	ebo := sortedEdges(g)
	frs := make(map[uint]uint)
	vc := uint(0)
	for _, e := range ebo {
		ti, iOk := frs[e.I]
		tj, jOk := frs[e.J]
		if iOk {
			if jOk { // no new vertex
				if ti != tj {
					retag(frs, ti, tj)
					mst = append(mst, e)
				}
			} else { // j is new vertex
				frs[e.J] = ti
				mst = append(mst, e)
				vc++
				if vc == g.VertexNo() {
					return mst, nil
				}
			}
		} else if jOk { // i is new vertex
			frs[e.I] = tj
			mst = append(mst, e)
			vc++
			if vc == g.VertexNo() {
				return mst, nil
			}
		} else { // i & j are new vertices
			frs[e.I] = e.I
			frs[e.J] = e.I
			mst = append(mst, e)
			vc += 2
			if vc == g.VertexNo() {
				return mst, nil
			}
		}
	}
	return mst, nil
}
