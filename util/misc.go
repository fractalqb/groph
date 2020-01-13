package util

import (
	"reflect"

	"git.fractalqb.de/fractalqb/groph"
)

// WeightOr returns parameter 'or' when the edge (u,v) is not in graph g.
// Otherwise the weight of edge (u,v) is returned.
func WeightOr(g groph.RGraph, u, v groph.VIdx, or interface{}) interface{} {
	if res := g.Weight(u, v); res != nil {
		return res
	}
	return or
}

// TODO can this be done in place?
func ReorderPath(slice interface{}, path []groph.VIdx) {
	slv := reflect.ValueOf(slice)
	if slv.Len() == 0 {
		return
	}
	tmp := make([]interface{}, slv.Len())
	for i := 0; i < slv.Len(); i++ {
		tmp[i] = slv.Index(i).Interface()
	}
	for w, r := range path {
		v := tmp[r]
		slv.Index(w).Set(reflect.ValueOf(v))
	}
}
