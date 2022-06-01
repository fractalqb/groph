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

package graphs

import "git.fractalqb.de/fractalqb/groph"

// TODO sort edges for faster access
type Edge struct {
	U, V groph.VIdx
	W    any
}

type EdgeList []Edge

func (g EdgeList) Order() int {
	ord := -1
	for _, e := range g {
		if e.U > ord {
			ord = e.U
		}
		if e.V > ord {
			ord = e.V
		}
	}
	return ord + 1
}

func (g EdgeList) Edge(u, v groph.VIdx) (weight any) {
	for _, e := range g {
		if e.U == u && e.V == v {
			return e.W
		}
	}
	return nil
}

func (g EdgeList) IsEdge(weight any) bool { return weight != nil }

func (g EdgeList) NotEdge() any { return nil }

func (g EdgeList) Size() int { return len(g) }

func (g EdgeList) EachEdge(onEdge groph.VisitEdge[any]) error {
	for _, e := range g {
		if err := onEdge(e.U, e.V, e.W); err != nil {
			return err
		}
	}
	return nil
}
