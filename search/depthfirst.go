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
	todo []int
	done internal.BitSet
}

func (s *dfs) push(v groph.VIdx) {
	s.todo = append(s.todo, v)
}

func (s *dfs) pop() (v groph.VIdx) {
	lm1 := len(s.todo) - 1
	v = s.todo[lm1]
	s.todo = s.todo[:lm1]
	return v
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

func (dfs *DirectedDFS[W]) Forward(start groph.VIdx, do groph.VisitVertex) error {
	dfs.todo = dfs.todo[:0]
	dfs.push(start)
	for len(dfs.todo) > 0 {
		start = dfs.pop()
		if !dfs.done.Get(start) {
			dfs.done.Set(start)
			if err := do(start); err != nil {
				return err
			}
			dfs.g.EachOut(start, func(v groph.VIdx) error {
				dfs.push(v)
				return nil
			})
		}
	}
	return nil
}

func (dfs *DirectedDFS[W]) Backward(start groph.VIdx, do groph.VisitVertex) error {
	dfs.todo = dfs.todo[:0]
	dfs.push(start)
	for len(dfs.todo) > 0 {
		start = dfs.pop()
		if !dfs.done.Get(start) {
			dfs.done.Set(start)
			if err := do(start); err != nil {
				return err
			}
			dfs.g.EachIn(start, func(v groph.VIdx) error {
				dfs.push(v)
				return nil
			})
		}
	}
	return nil
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

func (dfs *UndirectedDFS[W]) Start(start groph.VIdx, do groph.VisitVertex) error {
	dfs.todo = dfs.todo[:0]
	dfs.push(start)
	for len(dfs.todo) > 0 {
		start = dfs.pop()
		if !dfs.done.Get(start) {
			dfs.done.Set(start)
			if err := do(start); err != nil {
				return err
			}
			dfs.g.EachAdjacent(start, func(v groph.VIdx) error {
				dfs.push(v)
				return nil
			})
		}
	}
	return nil
}
