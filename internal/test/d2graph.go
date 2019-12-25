package test

import (
	"math"
	"math/rand"

	"git.fractalqb.de/fractalqb/groph"
)

type Point [2]float32

func Dist(p, q Point) float32 {
	d1 := p[0] - q[0]
	d2 := p[1] - q[1]
	return float32(math.Sqrt(float64(d1*d1 + d2*d2)))
}

func RandomPoints(n groph.VIdx, ps []Point) []Point {
	if groph.VIdx(cap(ps)) >= n {
		ps = ps[:n-1]
	} else {
		ps = make([]Point, n)
	}
	for i := range ps {
		ps[i][0] = rand.Float32()
		ps[i][1] = rand.Float32()
	}
	return ps
}
