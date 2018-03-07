package groph

import (
	"math"
	"math/rand"
)

type point [2]float32

func dist(p, q point) float32 {
	d1 := p[0] - q[0]
	d2 := p[1] - q[1]
	return float32(math.Sqrt(float64(d1*d1 + d2*d2)))
}

func randomPoints(n uint, ps []point) []point {
	if uint(cap(ps)) >= n {
		ps = ps[:n-1]
	} else {
		ps = make([]point, n)
	}
	for i := range ps {
		ps[i][0] = rand.Float32()
		ps[i][1] = rand.Float32()
	}
	return ps
}
