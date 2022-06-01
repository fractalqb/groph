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

package paths

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmtx"
	"git.fractalqb.de/fractalqb/groph/graphs"
)

func randomPoints(n groph.VIdx, ps graphs.Euclidean) graphs.Euclidean {
	if groph.VIdx(cap(ps)) >= n {
		ps = ps[:n-1]
	} else {
		ps = make(graphs.Euclidean, n)
	}
	for i := range ps {
		ps[i] = point{rand.Float64(), rand.Float64()}
	}
	return ps
}

func pathEq(p1, p2 []groph.VIdx) (bool, string) {
	plen := groph.VIdx(len(p1))
	if plen != groph.VIdx(len(p2)) {
		return false, fmt.Sprintf("length differs: %d / %d",
			plen,
			groph.VIdx(len(p2)))
	}
	s2 := 0
	for s2 < plen {
		if p1[0] == p2[s2] {
			break
		}
		s2++
	}
	if s2 >= plen {
		return false, fmt.Sprintf("no start %d in p2=%v", p1[0], p2)
	}
	pidx, nidx := func(i groph.VIdx) groph.VIdx {
		if i == 0 {
			return plen - 1
		}
		return i - 1
	},
		func(i groph.VIdx) groph.VIdx {
			i++
			if i >= plen {
				i = 0
			}
			return i
		}
	s1 := groph.VIdx(1)
	s2 = nidx(s2)
	if p1[s1] == p2[s2] {
		s1++
		s2 = nidx(s2)
		for s1 < plen {
			if p1[s1] != p2[s2] {
				return false, fmt.Sprintf("difference in pos %d / %d: %v %v",
					s1, s2,
					p1, p2)
			}
			s1++
			s2 = nidx(s2)
		}
	} else {
		s2 = pidx(pidx(s2))
		for s1 < plen {
			if p1[s1] != p2[s2] {
				return false, fmt.Sprintf("difference in pos %d / %d: %v %v",
					s1, s2,
					p1, p2)
			}
			s1++
			s2 = pidx(s2)
		}
	}
	return true, ""
}

func Test2OptDAgaintsGreedy(t *testing.T) {
	var points graphs.Euclidean
	var am *adjmtx.Directed[float64]
	for sz := groph.VIdx(4); sz < 12; sz++ {
		points = randomPoints(sz, points)
		am = adjmtx.NewDirected(sz, math.Inf(1), am)
		groph.Copy[float64](am, points)
		gPath, gWeight := GreedyTSP[float64](am)
		tPath, tWeight := TwoOptTSP[float64](am)
		if tWeight/gWeight > 1.01 {
			t.Errorf("size %d: different path length: greedy=%f / 2-opt=%f",
				sz,
				gWeight,
				tWeight)
		}
		if ok, msg := pathEq(gPath, tPath); !ok {
			t.Logf("size %d: %s", sz, msg)
		}
	}
}

func Test2OptUAgaintsGreedy(t *testing.T) {
	var points graphs.Euclidean
	var am *adjmtx.Undirected[float64]
	var dm *adjmtx.Directed[float64]
	for sz := groph.VIdx(4); sz < 12; sz++ {
		points = randomPoints(sz, points)
		dm = adjmtx.NewDirected(sz, math.Inf(1), dm)
		am = adjmtx.NewUndirected(sz, math.Inf(1), am)
		groph.Copy[float64](dm, points)
		groph.Copy[float64](am, points)
		gPath, gWeight := GreedyTSP[float64](dm)
		tPath, tWeight := TwoOptTSP[float64](am)
		if tWeight/gWeight > 1.052 {
			t.Errorf("size %d: different path length: greedy=%f / 2-opt=%f",
				sz,
				gWeight,
				tWeight)
		}
		if ok, msg := pathEq(gPath, tPath); !ok {
			t.Logf("size %d: %s", sz, msg)
		}
	}
}

const twoOptBenchSize = 120

func BenchmarkTwoOptTSP_directed(b *testing.B) {
	points := randomPoints(twoOptBenchSize, nil)
	am := adjmtx.NewDirected(twoOptBenchSize, math.Inf(1), nil)
	groph.Copy[float64](am, points)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoOptTSP[float64](am)
	}
}

func BenchmarkTwoOptTSP_undirected(b *testing.B) {
	points := randomPoints(twoOptBenchSize, nil)
	am := adjmtx.NewUndirected(twoOptBenchSize, math.Inf(1), nil)
	groph.Copy[float64](am, points)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoOptTSP[float64](am)
	}
}
