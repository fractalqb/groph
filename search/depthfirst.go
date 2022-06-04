// Copyright 2022 Marcus Perlick
// This file is part of Go module git.fractalqb.de/fractalqb/groph
//
// groph is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// groph is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with groph.  If not, see <http://www.gnu.org/licenses/>.

package search

import (
	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/internal"
)

type dfs struct {
	todo []edge
	done internal.BitSet
}

func (s *dfs) push(e edge) {
	s.todo = append(s.todo, e)
}

func (s *dfs) pop() (e edge) {
	lm1 := len(s.todo) - 1
	e = s.todo[lm1]
	s.todo = s.todo[:lm1]
	return e
}

type DirectedDFS[W any] struct {
	dfs
	g groph.RDirected[W]
}

func NewDirecredDFS[W any](g groph.RDirected[W]) *DirectedDFS[W] {
	res := new(DirectedDFS[W])
	return res.Reset(g)
}

func (dfs *DirectedDFS[W]) Reset(g groph.RDirected[W]) *DirectedDFS[W] {
	dfs.done = internal.NewBitSet(g.Order(), dfs.done)
	dfs.g = g
	return dfs
}

func (dfs *DirectedDFS[W]) NextStart() groph.VIdx {
	n := dfs.done.FirstUnset()
	if n >= dfs.g.Order() {
		return -1
	}
	return n
}

func (dfs *DirectedDFS[W]) Forward(start groph.VIdx, do groph.VisitEdge) error {
	dfs.todo = dfs.todo[:0]
	dfs.push(edge{-1, start})
	for len(dfs.todo) > 0 {
		e := dfs.pop()
		if !dfs.done.Get(e.v) {
			dfs.done.Set(e.v)
			if err := do(e.u, e.v); err != nil {
				return err
			}
			dfs.g.EachOut(e.v, func(v groph.VIdx) error {
				dfs.push(edge{e.v, v})
				return nil
			})
		}
	}
	return nil
}

func (dfs *DirectedDFS[W]) Backward(start groph.VIdx, do groph.VisitEdge) error {
	dfs.todo = dfs.todo[:0]
	dfs.push(edge{-1, start})
	for len(dfs.todo) > 0 {
		e := dfs.pop()
		if !dfs.done.Get(e.v) {
			dfs.done.Set(e.v)
			if err := do(e.u, e.v); err != nil {
				return err
			}
			dfs.g.EachIn(e.v, func(v groph.VIdx) error {
				dfs.push(edge{e.v, v})
				return nil
			})
		}
	}
	return nil
}

func (dfs *DirectedDFS[W]) HasCycle(start groph.VIdx) bool {
	dfs.todo = dfs.todo[:0]
	dfs.push(edge{-1, start})
	for len(dfs.todo) > 0 {
		e := dfs.pop()
		if dfs.done.Get(e.v) {
			return true
		}
		dfs.done.Set(e.v)
		dfs.g.EachOut(e.v, func(v groph.VIdx) error {
			dfs.push(edge{e.v, v})
			return nil
		})
		dfs.g.EachIn(e.v, func(v groph.VIdx) error {
			if v != e.u {
				dfs.push(edge{e.v, v})
			}
			return nil
		})
	}
	return false
}

type UndirectedDFS[W any] struct {
	dfs
	g groph.RUndirected[W]
}

func NewUndirecredDFS[W any](g groph.RUndirected[W]) *UndirectedDFS[W] {
	res := new(UndirectedDFS[W])
	res.Reset(g)
	return res
}

func (dfs *UndirectedDFS[W]) Reset(g groph.RUndirected[W]) *UndirectedDFS[W] {
	dfs.done = internal.NewBitSet(g.Order(), dfs.done)
	dfs.g = g
	return dfs
}

func (dfs *UndirectedDFS[W]) NextStart() groph.VIdx {
	n := dfs.done.FirstUnset()
	if n >= dfs.g.Order() {
		return -1
	}
	return n
}

func (dfs *UndirectedDFS[W]) Start(start groph.VIdx, do groph.VisitEdge) error {
	dfs.todo = dfs.todo[:0]
	dfs.push(edge{-1, start})
	for len(dfs.todo) > 0 {
		e := dfs.pop()
		if !dfs.done.Get(e.v) {
			dfs.done.Set(e.v)
			if err := do(e.u, e.v); err != nil {
				return err
			}
			dfs.g.EachAdjacent(e.v, func(v groph.VIdx) error {
				dfs.push(edge{e.v, v})
				return nil
			})
		}
	}
	return nil
}

func (dfs *UndirectedDFS[W]) HasCycle(start groph.VIdx) bool {
	dfs.todo = dfs.todo[:0]
	dfs.push(edge{-1, start})
	for len(dfs.todo) > 0 {
		e := dfs.pop()
		if dfs.done.Get(e.v) {
			return true
		}
		dfs.done.Set(e.v)
		dfs.g.EachAdjacent(e.v, func(v groph.VIdx) error {
			if v != e.u {
				dfs.push(edge{e.v, v})
			}
			return nil
		})
	}
	return false
}
