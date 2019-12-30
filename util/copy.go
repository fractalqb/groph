package util

import (
	"git.fractalqb.de/fractalqb/groph"
)

// CpWeights copies the edge weights from one graph to another.
// Vertices are identified by their index, i.e. the user has to take care of
// the vertex order. If the number of vertices in the graph differs the smaller
// graph determines how many edge weights are copied.
func CpWeights(dst groph.WGraph, src groph.RGraph) groph.WGraph {
	sz := dst.VertexNo()
	if src.VertexNo() < sz {
		sz = src.VertexNo()
	}
	if groph.Directed(dst) {
		for i := groph.VIdx(0); i < sz; i++ {
			for j := groph.VIdx(0); j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, w)
			}
		}
	} else {
		for i := groph.VIdx(0); i < sz; i++ {
			for j := i; j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, w)
			}
		}
	}
	return dst
}

// CpXWeights “transfers” the edge weights from src Graph to dst Graph
// with the same vertex restirctions as CpWeights. CpXWeights applies
// the transformation function xf() to each edge weight.
func CpXWeights(
	dst groph.WGraph,
	src groph.RGraph,
	xf func(in interface{}) interface{},
) groph.WGraph {
	sz := dst.VertexNo()
	if src.VertexNo() < sz {
		sz = src.VertexNo()
	}
	if groph.Directed(dst) {
		for i := groph.VIdx(0); i < sz; i++ {
			for j := groph.VIdx(0); j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, xf(w))
			}
		}
	} else {
		for i := groph.VIdx(0); i < sz; i++ {
			for j := i; j < sz; j++ {
				w := src.Weight(i, j)
				dst.SetWeight(i, j, xf(w))
			}
		}
	}
	return dst
}
