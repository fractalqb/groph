package search

import (
	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/internal"
)

type search struct {
	todo []int
	done internal.BitSet
}

func (s *search) init(order int) {
	if s.todo != nil {
		s.todo = s.todo[:0]
	}
	s.done = internal.NewBitSet(order, s.done)
}

func (s *search) push(v groph.VIdx) {
	s.todo = append(s.todo, v)
}

func (s *search) pop() (v groph.VIdx) {
	lm1 := len(s.todo) - 1
	v = s.todo[lm1]
	s.todo = s.todo[:lm1]
	return v
}

type DepthFirst[W any] struct {
	search
}

func (dfs *DepthFirst[W]) Directed(g groph.RDirected[W], start groph.VIdx, do groph.VisitVertex) bool {
	dfs.init(g.Order())
	dfs.push(start)
	for len(dfs.todo) > 0 {
		start = dfs.pop()
		if !dfs.done.Get(start) {
			dfs.done.Set(start)
			if do(start) {
				return true
			}
			g.EachOut(start, func(v groph.VIdx) bool {
				dfs.push(v)
				return false
			})
		}
	}
	return false
}

func (dfs *DepthFirst[W]) Undirected(g groph.RUndirected[W], start groph.VIdx, do groph.VisitVertex) bool {
	dfs.init(g.Order())
	dfs.push(start)
	for len(dfs.todo) > 0 {
		start = dfs.pop()
		if !dfs.done.Get(start) {
			dfs.done.Set(start)
			if do(start) {
				return true
			}
			g.EachAdjacent(start, func(v groph.VIdx) bool {
				dfs.push(v)
				return false
			})
		}
	}
	return false
}
