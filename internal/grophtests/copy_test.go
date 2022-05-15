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

package grophtests

import (
	"math"
	"fmt"
	"math/rand"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmtx"
)
const intNoE = -1

var f32NoE = float32(math.NaN())

func tstGSame[V, W any](g1 groph.RGraph[V], g2 groph.RGraph[W], sameOrd bool, wEq func(V, W) bool) error {
	ord := g1.Order()
	if sameOrd && g2.Order() != ord {
		return fmt.Errorf("graphs order differs: g1=%d, g2=%d", ord, g2.Order())
	}
	if g2.Order() < ord {
		ord = g2.Order()
	}
	for i := 0; i < ord; i++ {
		for j := 0; j < ord; j++ {
			w1, w2 := g1.Edge(i, j), g2.Edge(i, j)
			if !wEq(w1, w2) {
				return fmt.Errorf("differ at (%d,%d): w1=%v, w2=%v", i, j, w1, w2)
			}
		}
	}
	return nil
}

func TestCpWeights_from_directed(t *testing.T) {
	src := adjmtx.NewDirected[int](11, intNoE, nil)
	for i := 0; i < src.Order(); i++ {
		for j := 0; j < src.Order(); j++ {
			if rand.Intn(100) < 75 {
				src.SetEdge(i, j, rand.Int())
			}
		}
	}
	wEq := func(w1, w2 int) bool { return w1 == w2 }
	t.Run("to directed", func(t *testing.T) {
		dst := adjmtx.NewDirected[int](src.Order(), intNoE, nil)
		err := groph.Copy[int](dst, src)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame[int, int](src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
	t.Run("to undirected", func(t *testing.T) {
		dst := adjmtx.NewUndirected[int](src.Order(), intNoE, nil)
		err := groph.Copy[int](dst, src)
		if err == nil {
			t.Fatal("copy from directed to undirected did not return error")
		}
		if err.Error() != "cannot copy from directed to undirected graph" {
			t.Error("wrong error:", err)
		}
	})
}

func TestCpWeights_from_undirected(t *testing.T) {
	src := adjmtx.NewUndirected[int](11, intNoE, nil)
	for i := 0; i < src.Order(); i++ {
		for j := 0; j <= i; j++ {
			if rand.Intn(100) < 75 {
				src.SetEdge(i, j, rand.Int())
			}
		}
	}
	wEq := func(w1, w2 int) bool { return w1 == w2 }
	t.Run("to directed", func(t *testing.T) {
		dst := adjmtx.NewDirected[int](src.Order(), intNoE, nil)
		err := groph.Copy[int](dst, src)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame[int, int](src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
	t.Run("to undirected", func(t *testing.T) {
		dst := adjmtx.NewUndirected[int](src.Order(), intNoE, nil)
		err := groph.Copy[int](dst, src)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame[int, int](src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
}

func TestCpXWeights_from_directed(t *testing.T) {
	src := adjmtx.NewDirected[int](11, intNoE, nil)
	for i := 0; i < src.Order(); i++ {
		for j := 0; j < src.Order(); j++ {
			if rand.Intn(100) < 75 {
				src.SetEdge(i, j, rand.Int())
			}
		}
	}
	xFn := func(w int) (float32, error) { return float32(w), nil }
	wEq := func(w1 int, w2 float32) bool { return float32(w1) == w2	}
	t.Run("to directed", func(t *testing.T) {
		dst := adjmtx.NewDirected[float32](src.Order(), f32NoE, nil)
		err := groph.CopyX[float32, int](dst, src, xFn)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame[int, float32](src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
	t.Run("to undirected", func(t *testing.T) {
		dst := adjmtx.NewUndirected[float32](src.Order(), f32NoE, nil)
		err := groph.CopyX[float32, int](dst, src, xFn)
		if err == nil {
			t.Fatal("copy from directed to undirected did not return error")
		}
		if err.Error() != "cannot copy from directed to undirected graph" {
			t.Error("wrong error:", err)
		}
	})
}

func TestCpXWeights_from_undirected(t *testing.T) {
	src := adjmtx.NewUndirected[int](11, intNoE, nil)
	for i := 0; i < src.Order(); i++ {
		for j := 0; j <= i; j++ {
			if rand.Intn(100) < 75 {
				src.SetEdge(i, j, rand.Int())
			}
		}
	}
	xFn := func(w int) (float32, error) { return float32(w), nil }
	wEq := func(w1 int, w2 float32) bool { return float32(w1) == w2	}
	t.Run("to directed", func(t *testing.T) {
		dst := adjmtx.NewDirected[float32](src.Order(), f32NoE, nil)
		err := groph.CopyX[float32, int](dst, src, xFn)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame[int, float32](src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
	t.Run("to undirected", func(t *testing.T) {
		dst := adjmtx.NewUndirected[float32](src.Order(), f32NoE, nil)
		err := groph.CopyX[float32, int](dst, src, xFn)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame[int, float32](src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
}
