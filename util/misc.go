package util

import "git.fractalqb.de/fractalqb/groph"

// WeightOr returns parameter 'or' when the edge (u,v) is not in graph g.
// Otherwise the weight of edge (u,v) is returned.
func WeightOr(g groph.RGraph, u, v groph.VIdx, or interface{}) interface{} {
	if res := g.Weight(u, v); res != nil {
		return res
	}
	return or
}
