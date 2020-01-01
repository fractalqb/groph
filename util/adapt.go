package util

import (
	"fmt"
	"math"
	"reflect"

	"git.fractalqb.de/fractalqb/groph"
)

type WeightsSlice struct {
	slc reflect.Value
	sz  groph.VIdx
}

func NewWeightsSlice(edgeSlice interface{}) *WeightsSlice {
	res := &WeightsSlice{slc: reflect.ValueOf(edgeSlice)}
	res.sz = groph.VIdx(math.Sqrt(float64(res.slc.Len())))
	return res
}

func (g *WeightsSlice) Check() (*WeightsSlice, error) {
	if g.slc.Kind() != reflect.Slice {
		return g, fmt.Errorf("edges have to be a slice, got %s",
			g.slc.Type().String())
	}
	if g.sz*g.sz != groph.VIdx(g.slc.Len()) {
		return g, fmt.Errorf("slice len is not quadratic")
	}
	return g, nil
}

func (g *WeightsSlice) Must() *WeightsSlice {
	var err error
	g, err = g.Check()
	if err != nil {
		panic(err)
	}
	return g
}

func (g *WeightsSlice) Order() groph.VIdx { return g.sz }

func (g *WeightsSlice) Weight(u, v groph.VIdx) interface{} {
	return g.slc.Index(int(g.sz*u + v)).Interface()
}

// PointsNDist implements a metric RUndirected graph based on a slice of
// vertices of some type V and a symmetric function f(u V, v V) → W that
// computes the weight of type W for any edge (u, v).
//
// From this use e.g. CpWeights or CpXWeights to initialize an other WGraph.
type PointsNDist struct {
	ps reflect.Value
	d  reflect.Value
}

func NewPointsNDist(vertexSlice interface{}, measure interface{}) *PointsNDist {
	res := &PointsNDist{
		ps: reflect.ValueOf(vertexSlice),
		d:  reflect.ValueOf(measure),
	}
	return res
}

// Check does type checking on g. It always returns g. Only if everything is OK
// the returned error is nil.
func (g *PointsNDist) Check() (*PointsNDist, error) {
	if g.ps.Kind() != reflect.Slice {
		return g, fmt.Errorf("vertex set has to be a slice, got %s",
			g.ps.Type().String())
	}
	if g.d.Kind() != reflect.Func {
		return g, fmt.Errorf("edge weight measure must be a function, got: %s",
			g.d.Type().String())
	} // TODO more precise checking: func signature
	return g, nil
}

// Verify call Check on g and panics if Check returns an error.
func (g *PointsNDist) Must() *PointsNDist {
	if _, err := g.Check(); err != nil {
		panic(err)
	}
	return g
}

func (g *PointsNDist) Order() groph.VIdx {
	return groph.VIdx(g.ps.Len())
}

func (g *PointsNDist) Vertex(idx groph.VIdx) interface{} {
	return g.ps.Index(int(idx)).Interface()
}

func (g *PointsNDist) Weight(u, v groph.VIdx) interface{} {
	f, t := g.ps.Index(int(u)), g.ps.Index(int(v))
	d := g.d.Call([]reflect.Value{f, t})
	return d[0].Interface()
}

func (g *PointsNDist) WeightU(u, v groph.VIdx) interface{} { return g.Weight(u, v) }
