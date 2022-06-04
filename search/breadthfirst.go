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

type edge struct{ u, v groph.VIdx }

type bfs struct {
	todo []edge
	tail int
	done internal.BitSet
}

func (s *bfs) enq(v edge) {
	if 2*s.tail > len(s.todo) { // TODO wild guess
		copy(s.todo, s.todo[s.tail:])
		s.todo = s.todo[:len(s.todo)-s.tail]
		s.tail = 0
	}
	s.todo = append(s.todo, v)
}

func (s *bfs) deq() edge {
	v := s.todo[s.tail]
	s.tail++
	return v
}

func (s *bfs) qlen() int { return len(s.todo) - s.tail }

type DirectedBFS[W any] struct {
	bfs
	g groph.RDirected[W]
}

func NewDirectedBFS[W any](g groph.RDirected[W]) *DirectedBFS[W] {
	s := new(DirectedBFS[W])
	return s.Reset(g)
}

func (bfs *DirectedBFS[W]) Reset(g groph.RDirected[W]) *DirectedBFS[W] {
	bfs.done = internal.NewBitSet(g.Order(), bfs.done)
	bfs.g = g
	return bfs
}

func (bfs *DirectedBFS[W]) NextStart() groph.VIdx {
	n := bfs.done.FirstUnset()
	if n >= bfs.g.Order() {
		return -1
	}
	return n
}

func (bfs *DirectedBFS[W]) Forward(start groph.VIdx, do groph.VisitEdge) error {
	bfs.todo = bfs.todo[:0]
	bfs.tail = 0
	bfs.done.Set(start)
	bfs.enq(edge{-1, start})
	for bfs.qlen() > 0 {
		e := bfs.deq()
		if err := do(e.u, e.v); err != nil {
			return err
		}
		bfs.g.EachOut(start, func(n groph.VIdx) error {
			if !bfs.done.Get(n) {
				bfs.done.Set(n)
				bfs.enq(edge{e.v, n})
			}
			return nil
		})
	}
	return nil
}

func (bfs *DirectedBFS[W]) Backward(start groph.VIdx, do groph.VisitEdge) error {
	bfs.todo = bfs.todo[:0]
	bfs.tail = 0
	bfs.done.Set(start)
	bfs.enq(edge{-1, start})
	for bfs.qlen() > 0 {
		e := bfs.deq()
		if err := do(e.u, e.v); err != nil {
			return err
		}
		bfs.g.EachIn(start, func(n groph.VIdx) error {
			if !bfs.done.Get(n) {
				bfs.done.Set(n)
				bfs.enq(edge{e.v, n})
			}
			return nil
		})
	}
	return nil
}

type UndirectedBFS[W any] struct {
	bfs
	g groph.RUndirected[W]
}

func NewUndirectedBFS[W any](g groph.RUndirected[W]) *UndirectedBFS[W] {
	s := new(UndirectedBFS[W])
	return s.Reset(g)
}

func (bfs *UndirectedBFS[W]) Reset(g groph.RUndirected[W]) *UndirectedBFS[W] {
	bfs.done = internal.NewBitSet(g.Order(), bfs.done)
	bfs.g = g
	return bfs
}

func (bfs *UndirectedBFS[W]) NextStart() groph.VIdx {
	n := bfs.done.FirstUnset()
	if n >= bfs.g.Order() {
		return -1
	}
	return n
}

func (bfs *UndirectedBFS[W]) Start(start groph.VIdx, do groph.VisitEdge) error {
	bfs.todo = bfs.todo[:0]
	bfs.tail = 0
	bfs.done.Set(start)
	bfs.enq(edge{-1, start})
	for bfs.qlen() > 0 {
		e := bfs.deq()
		if err := do(e.u, e.v); err != nil {
			return err
		}
		bfs.g.EachAdjacent(e.v, func(n groph.VIdx) error {
			if !bfs.done.Get(n) {
				bfs.done.Set(n)
				bfs.enq(edge{e.v, n})
			}
			return nil
		})
	}
	return nil
}
