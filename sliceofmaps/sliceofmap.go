// Package sliceofmaps provides an implementations for graph interfaces suitable
// for sparse graphs.
package sliceofmaps

import "git.fractalqb.de/fractalqb/groph"

type spmro map[groph.VIdx]interface{}

type SoMD struct {
	ws []spmro
}

func NewSoMD(order int, reuse *SoMD) *SoMD {
	if reuse == nil {
		return &SoMD{make([]spmro, order)}
	}
	reuse.Reset(order)
	return reuse
}

func (g *SoMD) Order() int { return len(g.ws) }

func (g *SoMD) Weight(u, v groph.VIdx) interface{} {
	row := g.ws[u]
	if row == nil {
		return nil
	}
	return row[v]
}

func (g *SoMD) SetWeight(u, v groph.VIdx, w interface{}) {
	row := g.ws[u]
	if w == nil {
		delete(row, v)
	} else {
		if row == nil {
			row = make(spmro)
			g.ws[u] = row
		}
		row[v] = w
	}
}

func (g *SoMD) Reset(order int) {
	if g.ws == nil || cap(g.ws) < order {
		g.ws = make([]spmro, order)
	} else {
		g.ws = g.ws[:order]
		for i := range g.ws {
			g.ws[i] = nil
		}
	}
}

func (g *SoMD) Reorder(order int) {
	o := g.Order()
	switch {
	case order == g.Order():
		return
	case order < o:
		g.ws = g.ws[:order]
	default:
		tmp := make([]spmro, order)
		copy(tmp, g.ws)
		g.ws = tmp
	}
}

func (g *SoMD) EachOutgoing(from groph.VIdx, onDest groph.VisitVertex) (stopped bool) {
	if row := g.ws[from]; row != nil {
		for n := range row {
			if onDest(n) {
				return true
			}
		}
	}
	return false
}

func (g *SoMD) OutDegree(v groph.VIdx) int {
	row := g.ws[v]
	if row == nil {
		return 0
	}
	return len(row)
}

type SoMU struct {
	SoMD
}

func NewSoMU(order int, reuse *SoMU) *SoMU {
	if reuse == nil {
		reuse = new(SoMU)
	}
	NewSoMD(order, &reuse.SoMD)
	return reuse
}

func (g *SoMU) WeightU(u, v groph.VIdx) interface{} {
	return g.SoMD.Weight(u, v)
}

func (g *SoMU) Weight(u, v groph.VIdx) interface{} {
	if u > v {
		return g.WeightU(u, v)
	}
	return g.WeightU(v, u)
}

func (g *SoMU) SetWeightU(u, v groph.VIdx, w interface{}) {
	g.SoMD.SetWeight(u, v, w)
}

func (g *SoMU) SetWeight(u, v groph.VIdx, w interface{}) {
	if u > v {
		g.SetWeightU(u, v, w)
	} else {
		g.SetWeightU(v, u, w)
	}
}
