package groph

import (
	"fmt"
	"testing"
)

func pathEq(p1, p2 []uint) (bool, string) {
	plen := uint(len(p1))
	if plen != uint(len(p2)) {
		return false, fmt.Sprintf("length differs: %d / %d",
			plen,
			uint(len(p2)))
	}
	s2 := uint(0)
	for s2 < plen {
		if p1[0] == p2[s2] {
			break
		}
		s2++
	}
	if s2 >= plen {
		return false, fmt.Sprintf("no start %d in p2=%v", p1[0], p2)
	}
	if p1[1] == p2[(s2+1)%plen] {
		if s2 += 2; s2 >= plen {
			s2 = 0
		}
		for s1 := uint(2); s1 < plen; s1++ {
			if p1[s1] != p2[s2] {
				return false, fmt.Sprintf("difference in pos %d / %d: %v %v",
					s1, s2,
					p1, p2)
			}
			if s2 += 1; s2 >= plen {
				s2 = 0
			}
		}
	} else {
		if s2 == 0 {
			s2 = plen - 1
		} else {
			s2--
		}
		for s1 := uint(1); s1 < plen; s1++ {
			if p1[s1] != p2[s2] {
				return false, fmt.Sprintf("difference in pos %d / %d: %v %v",
					s1, s2,
					p1, p2)
			}
			if s2 == 0 {
				s2 = plen - 1
			} else {
				s2--
			}
		}
	}
	return true, ""
}

func Test2OptDAgaintsGreedy(t *testing.T) {
	var points []point
	var am *AdjMxDf32
	for sz := uint(4); sz < 12; sz++ {
		points = randomPoints(sz, points)
		am = CpWeights(
			NewAdjMxDf32(sz, am),
			NewSliceNMeasure(points, dist, false).Verify(),
		).(*AdjMxDf32)
		gPath, gWeight := am.TspGreedy()
		tPath, tWeight := Tsp2Optf32(am)
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
	var points []point
	var am *AdjMxUf32
	var dm *AdjMxDf32
	for sz := uint(4); sz < 12; sz++ {
		points = randomPoints(sz, points)
		am = CpWeights(
			NewAdjMxUf32(sz, am),
			NewSliceNMeasure(points, dist, false).Verify(),
		).(*AdjMxUf32)
		dm = CpWeights(
			NewAdjMxDf32(sz, dm),
			NewSliceNMeasure(points, dist, false).Verify(),
		).(*AdjMxDf32)
		gPath, gWeight := dm.TspGreedy()
		tPath, tWeight := Tsp2Optf32(am)
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

func BenchmarkTsp2OptGenf32D(b *testing.B) {
	sz := uint(120)
	points := randomPoints(sz, nil)
	am := CpWeights(
		NewAdjMxDf32(sz, nil),
		NewSliceNMeasure(points, dist, false).Verify(),
	).(*AdjMxDf32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Tsp2Optf32(am)
	}
}

func BenchmarkTsp2OptGenf32U(b *testing.B) {
	sz := uint(120)
	points := randomPoints(sz, nil)
	am := CpWeights(
		NewAdjMxUf32(sz, nil),
		NewSliceNMeasure(points, dist, false).Verify(),
	).(*AdjMxUf32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Tsp2Optf32(am)
	}
}
