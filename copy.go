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

package groph

import (
	"errors"
	"fmt"

	"git.fractalqb.de/fractalqb/groph/internal"
)

// TODO What to do with w → !g.IsEdge(w) ??? Use CopyX ???
func Copy[W any](dst WGraph[W], src RGraph[W]) error {
	sz := dst.Order()
	if src.Order() < sz {
		sz = src.Order()
	}
	if udst, ok := dst.(WUndirected[W]); ok {
		if usrc, ok := src.(RUndirected[W]); ok {
			for i := 0; i < sz; i++ {
				for j := 0; j <= i; j++ {
					udst.SetEdgeU(i, j, usrc.EdgeU(i, j))
				}
			}
		} else {
			return errors.New("cannot copy from directed to undirected graph")
		}
	} else if usrc, ok := src.(RUndirected[W]); ok {
		for i := 0; i < sz; i++ {
			dst.SetEdge(i, i, usrc.EdgeU(i, i))
			for j := 0; j < i; j++ {
				w := usrc.EdgeU(i, j)
				dst.SetEdge(i, j, w)
				dst.SetEdge(j, i, w)
			}
		}
	} else {
		for i := 0; i < sz; i++ {
			for j := 0; j < sz; j++ {
				dst.SetEdge(i, j, src.Edge(i, j))
			}
		}
	}
	sderr := SrcDstError{Src: internal.ErrState(src), Dst: internal.ErrState(dst)}
	if sderr.Src != nil || sderr.Dst != nil {
		return sderr
	}
	vnd := dst.Order() - src.Order()
	if vnd == 0 {
		return nil
	}
	return ClipError(vnd)
}

func CopyX[V, W any](dst WGraph[V], src RGraph[W], x func(W) (V, error)) error {
	sz := dst.Order()
	if src.Order() < sz {
		sz = src.Order()
	}
	if udst, ok := dst.(WUndirected[V]); ok {
		if usrc, ok := src.(RUndirected[W]); ok {
			for i := 0; i < sz; i++ {
				for j := 0; j <= i; j++ {
					v, err := x(usrc.EdgeU(i, j))
					if err != nil {
						return UVError{U: i, V: j, err: err}
					}
					udst.SetEdgeU(i, j, v)
				}
			}
		} else {
			return errors.New("cannot copy from directed to undirected graph")
		}
	} else if usrc, ok := src.(RUndirected[W]); ok {
		for i := 0; i < sz; i++ {
			v, err := x(usrc.EdgeU(i, i))
			if err != nil {
				return UVError{U: i, V: i, err: err}
			}
			dst.SetEdge(i, i, v)
			for j := 0; j < i; j++ {
				v, err := x(usrc.EdgeU(i, j))
				if err != nil {
					return UVError{U: i, V: j, err: err}
				}
				dst.SetEdge(i, j, v)
				dst.SetEdge(j, i, v)
			}
		}
	} else {
		for i := 0; i < sz; i++ {
			for j := 0; j < sz; j++ {
				v, err := x(src.Edge(i, j))
				if err != nil {
					return UVError{U: i, V: j, err: err}
				}
				dst.SetEdge(i, j, v)
			}
		}
	}
	vnd := dst.Order() - src.Order()
	if vnd == 0 {
		return nil
	}
	return ClipError(vnd)
}

type UVError struct {
	U, V VIdx
	err  error
}

func (e UVError) Error() string {
	return fmt.Sprintf("at %d, %d: %s", e.U, e.V, e.err)
}

func (e UVError) Unwrap() error { return e.err }

// ClipError is returned when the order of src and dst in Cp*Weights does not
// match. If err is < 0 there were -err vertices ignored from src. If err > 0
// then err vertices in dst were not covered.
type ClipError int

func (err ClipError) Error() string {
	if err < 0 {
		return fmt.Sprintf("%d vertices ignored from source", -err)
	}
	return fmt.Sprintf("%d vertices not covered in destination", err)
}

// Clipped returns the clipping if err is a ClipError. Otherwise it returns 0.
func Clipped(err error) int {
	if ce, ok := err.(ClipError); ok {
		return int(ce)
	}
	return 0
}

// SrcDstError is returned when an error occurred in the src or dst graph during
// Cp*Weights.
type SrcDstError struct {
	Src error
	Dst error
}

func (err SrcDstError) Error() string {
	if err.Src != nil {
		if err.Dst != nil {
			return fmt.Sprintf("source: %s; dest: %s", err.Src, err.Dst)
		}
		return fmt.Sprintf("source: %s", err.Src)
	} else if err.Dst != nil {
		return fmt.Sprintf("dest: %s", err.Dst)
	}
	return "unspecific error"
}
