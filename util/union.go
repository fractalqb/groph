package util

import (
	"git.fractalqb.de/fractalqb/groph"
)

func FirstEdge(ofIdx int, akku, w interface{}) (sum interface{}, stop bool) {
	if w != nil {
		return w, true
	}
	return nil, false
}

type RUnion struct {
	Of        []groph.RGraph
	AddWeight func(ofIdx int, akku, w interface{}) (sum interface{}, stop bool)
}

func (g RUnion) Order() (res groph.VIdx) {
	for _, e := range g.Of {
		if o := e.Order(); o > res {
			res = o
		}
	}
	return res
}

func (g RUnion) Weight(u, v groph.VIdx) (w interface{}) {
	var stop bool
	for i, e := range g.Of {
		if o := e.Order(); u >= o || v >= o {
			continue
		}
		w, stop = g.AddWeight(i, w, e.Weight(u, v))
		if stop {
			break
		}
	}
	return w
}
