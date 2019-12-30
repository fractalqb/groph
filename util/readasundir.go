package util

import (
	"fmt"

	"git.fractalqb.de/fractalqb/groph"
)

type MergeError struct {
	u, v groph.VIdx
	err  error
}

func (e MergeError) Error() string {
	return fmt.Sprintf("edges %d, %d: %s", e.u, e.v, e.err)
}

func (e MergeError) Unwrap() error { return e.err }

type RUadapter struct {
	G     groph.RGraph
	Merge func(u, v groph.VIdx) (merged interface{}, err error)
	Err   error
}

func (rua *RUadapter) Order() groph.VIdx { return rua.G.Order() }

func (rua *RUadapter) Weight(u, v groph.VIdx) interface{} {
	return rua.WeightU(u, v)
}

func (rua *RUadapter) WeightU(u, v groph.VIdx) (res interface{}) {
	var err error
	if res, err = rua.Merge(u, v); err != nil {
		rua.Err = MergeError{u, v, err}
	}
	return res
}

func MergeWeights(
	g groph.RGraph,
	merge func(w1, w2 interface{}) (merged interface{}, err error),
) func(u, v groph.VIdx) (interface{}, error) {
	return func(u, v groph.VIdx) (interface{}, error) {
		w1 := g.Weight(u, v)
		w2 := g.Weight(v, u)
		return merge(w1, w2)
	}
}

func MergeEqual(w1, w2 interface{}) (interface{}, error) {
	if w1 == w2 {
		return w1, nil
	}
	return nil, fmt.Errorf("not equal: '%v' / '%v'", w1, w2)
}
