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

type testRDFull int

var _ groph.RGraph = testRDFull(0)

func (g testRDFull) Order() int { return int(g) }

func (g testRDFull) Weight(u, v groph.VIdx) interface{} { return true }

type testRDols struct{ testRDFull }

var _ groph.OutLister = testRDols{}

func (g testRDols) EachOutgoing(from groph.VIdx, onDest groph.VisitVertex) bool {
	for i := 0; i < g.Order(); i++ {
		if onDest(i) {
			return true
		}
	}
	return false
}

func (g testRDols) OutDegree(v groph.VIdx) int { return g.Order() }

type testRDils struct{ testRDFull }

var _ groph.InLister = testRDils{}

func (g testRDils) EachIncoming(from groph.VIdx, onDest groph.VisitVertex) bool {
	for i := 0; i < g.Order(); i++ {
		if onDest(i) {
			return true
		}
	}
	return false
}

func (g testRDils) InDegree(v groph.VIdx) int { return g.Order() }

type testRDiols struct{ testRDols }

var (
	_ groph.OutLister = testRDiols{}
	_ groph.InLister  = testRDiols{}
)

func (g testRDiols) EachIncoming(from groph.VIdx, onDest groph.VisitVertex) bool {
	for i := 0; i < g.Order(); i++ {
		if onDest(i) {
			return true
		}
	}
	return false
}

func (g testRDiols) InDegree(v groph.VIdx) int { return g.Order() }

type testRDels struct{ testRDFull }

var _ groph.EdgeLister = testRDels{}

func (g testRDels) EachEdge(onEdge groph.VisitEdge) (stop bool) {
	o := g.Order()
	for i := 0; i < o; i++ {
		for j := 0; j < o; j++ {
			if onEdge(i, j) {
				return true
			}
		}
	}
	return false
}

func (g testRDels) Size() int {
	o := g.Order()
	return o * o
}

type testRUFull int

var _ groph.RGraph = testRDFull(0)

func (g testRUFull) Order() int { return int(g) }

func (g testRUFull) Weight(u, v groph.VIdx) interface{} { return true }

func (g testRUFull) WeightU(u, v groph.VIdx) interface{} { return true }

type testRUols struct{ testRUFull }

var _ groph.OutLister = testRUols{}

func (g testRUols) EachOutgoing(from groph.VIdx, onDest groph.VisitVertex) bool {
	for i := 0; i < g.Order(); i++ {
		if onDest(i) {
			return true
		}
	}
	return false
}

func (g testRUols) OutDegree(v groph.VIdx) int { return g.Order() }

type testRUils struct{ testRUFull }

var _ groph.InLister = testRUils{}

func (g testRUils) EachIncoming(from groph.VIdx, onDest groph.VisitVertex) bool {
	for i := 0; i < g.Order(); i++ {
		if onDest(i) {
			return true
		}
	}
	return false
}

func (g testRUils) InDegree(v groph.VIdx) int { return g.Order() }

type testRUiols struct{ testRUols }

var (
	_ groph.OutLister = testRUiols{}
	_ groph.InLister  = testRUiols{}
)

func (g testRUiols) EachIncoming(from groph.VIdx, onDest groph.VisitVertex) bool {
	for i := 0; i < g.Order(); i++ {
		if onDest(i) {
			return true
		}
	}
	return false
}

func (g testRUiols) InDegree(v groph.VIdx) int { return g.Order() }

type testRUels struct{ testRUFull }

var _ groph.EdgeLister = testRUels{}

func (g testRUels) EachEdge(onEdge groph.VisitEdge) (stop bool) {
	o := g.Order()
	for i := 0; i < o; i++ {
		for j := 0; j < o; j++ {
			if onEdge(i, j) {
				return true
			}
		}
	}
	return false
}

func (g testRUels) Size() int {
	o := g.Order()
	return (o * (o + 1)) / 2
}
