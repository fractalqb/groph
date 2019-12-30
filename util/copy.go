package util

import (
	"errors"
	"fmt"

	"git.fractalqb.de/fractalqb/groph"
)

type CpClipped int

func (cc CpClipped) Error() string {
	if cc < 0 {
		return fmt.Sprintf("%d vertices ignored from source", -cc)
	}
	return fmt.Sprintf("%d vertices not covered in destination", cc)
}

func MustCp(g groph.WGraph, err error) groph.WGraph {
	if err != nil {
		panic(err)
	}
	return g
}

// CpWeights copies the edge weights from one graph to another.
// Vertices are identified by their index, i.e. the user has to take care of
// the vertex order. If the number of vertices in the graph differs the smaller
// graph determines how many edge weights are copied.
func CpWeights(dst groph.WGraph, src groph.RGraph) (groph.WGraph, error) {
	sz := dst.VertexNo()
	if src.VertexNo() < sz {
		sz = src.VertexNo()
	}
	if udst, ok := dst.(groph.WUndirected); ok {
		if usrc, ok := src.(groph.RUndirected); ok {
			for i := groph.V0; i < sz; i++ {
				udst.SetWeightU(i, i, usrc.WeightU(i, i))
				for j := groph.V0; j < i; j++ {
					udst.SetWeightU(i, j, usrc.WeightU(i, j))
				}
			}
		} else {
			return dst, errors.New("cannot copy from directed to undirected graph")
		}
	} else if usrc, ok := src.(groph.RUndirected); ok {
		for i := groph.V0; i < sz; i++ {
			dst.SetWeight(i, i, usrc.WeightU(i, i))
			for j := groph.V0; j < i; j++ {
				w := usrc.WeightU(i, j)
				dst.SetWeight(i, j, w)
				dst.SetWeight(j, i, w)
			}
		}
	} else {
		for i := groph.V0; i < sz; i++ {
			for j := groph.V0; j < sz; j++ {
				dst.SetWeight(i, j, src.Weight(i, j))
			}
		}
	}
	vnd := dst.VertexNo() - src.VertexNo()
	if vnd == 0 {
		return dst, nil
	}
	return dst, CpClipped(vnd)
}

// CpXWeights “transfers” the edge weights from src Graph to dst Graph
// with the same vertex restirctions as CpWeights. CpXWeights applies
// the transformation function xf() to each edge weight.
func CpXWeights(
	dst groph.WGraph,
	src groph.RGraph,
	xf func(in interface{}) interface{},
) (groph.WGraph, error) {
	sz := dst.VertexNo()
	if src.VertexNo() < sz {
		sz = src.VertexNo()
	}
	var w interface{}
	if udst, ok := dst.(groph.WUndirected); ok {
		if usrc, ok := src.(groph.RUndirected); ok {
			for i := groph.V0; i < sz; i++ {
				w = usrc.WeightU(i, i)
				udst.SetWeightU(i, i, xf(w))
				for j := groph.V0; j < i; j++ {
					w = usrc.WeightU(i, j)
					udst.SetWeightU(i, j, xf(w))
				}
			}
		} else {
			return dst, errors.New("cannot copy from directed to undirected graph")
		}
	} else if usrc, ok := src.(groph.RUndirected); ok {
		for i := groph.V0; i < sz; i++ {
			w = usrc.WeightU(i, i)
			dst.SetWeight(i, i, xf(w))
			for j := groph.V0; j < i; j++ {
				w := xf(usrc.WeightU(i, j))
				dst.SetWeight(i, j, w)
				dst.SetWeight(j, i, w)
			}
		}
	} else {
		for i := groph.V0; i < sz; i++ {
			for j := groph.V0; j < sz; j++ {
				w = src.Weight(i, j)
				dst.SetWeight(i, j, xf(w))
			}
		}
	}
	vnd := dst.VertexNo() - src.VertexNo()
	if vnd == 0 {
		return dst, nil
	}
	return dst, CpClipped(vnd)
}
