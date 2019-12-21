package tsp

import (
	"fmt"
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	test "git.fractalqb.de/fractalqb/groph/internal"
)

//func dist(u, v [2]float32) float32 {
//	d1, d2 := u[0]-v[0], u[1]-v[1]
//	d := math.Sqrt(float64(d1*d1 + d2*d2))
//	return float32(d)
//}

func showMatrix(ps [][2]float32) {
	fmt.Println("Matrix:")
	for i, p := range ps {
		fmt.Printf("%2d: ", i)
		for j, q := range ps {
			if j > 0 {
				fmt.Print(", ")
			}
			d := test.Dist(p, q)
			fmt.Printf("%5.2f", d)
		}
		fmt.Println()
	}
}

var exmp1 = [][2]float32{
	[2]float32{0, 0},
	[2]float32{10, 10},
	[2]float32{2, 9},
	[2]float32{4, 5},
	[2]float32{8, 3},
}

func ExampleAsymGreedy() {
	adp, err := groph.NewSliceNMeasure(exmp1, test.Dist, false).Check()
	if err != nil {
		fmt.Println(err)
	}
	am := groph.CpWeights(groph.NewAdjMxDf32(adp.VertexNo(), nil), adp).(*groph.AdjMxDf32)
	showMatrix(exmp1)
	w, l := GreedyAdjMxDf32(am)
	fmt.Printf("%v %.2f", w, l)
	// Output:
	// Matrix:
	//  0:  0.00, 14.14,  9.22,  6.40,  8.54
	//  1: 14.14,  0.00,  8.06,  7.81,  7.28
	//  2:  9.22,  8.06,  0.00,  4.47,  8.49
	//  3:  6.40,  7.81,  4.47,  0.00,  4.47
	//  4:  8.54,  7.28,  8.49,  4.47,  0.00
	// [0 3 2 1 4] 34.76
}

var exmp2 = [][2]float32{
	[2]float32{0, 0},
	[2]float32{10, 10},
	[2]float32{2, 9},
	[2]float32{4, 5},
	[2]float32{8, 3},
	[2]float32{9, 2},
	[2]float32{5, 4},
	[2]float32{3, 8},
}

func BenchmarkTspGreedyAMf32(b *testing.B) {
	am := groph.NewAdjMxDf32(uint(len(exmp2)), nil)
	groph.CpWeights(am, groph.NewSliceNMeasure(exmp2, test.Dist, false).Verify())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GreedyAdjMxDf32(am)
	}
}

func BenchmarkTspGreedyGenf32(b *testing.B) {
	am := groph.NewAdjMxDf32(uint(len(exmp2)), nil)
	groph.CpWeights(am, groph.NewSliceNMeasure(exmp2, test.Dist, false).Verify())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Greedyf32(am)
	}
}

//func BenchmarkGreedy_Alt(b *testing.B) {
//	solver := AsymSolverf32(AsymGreedyf32_Alt)
//	solver(dist, exmp2)
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		solver(dist, exmp1)
//	}
//}

// AsymGreedyf32_Alt tries to improve AsymGreedyf32 from the fact
// that Heap Algorithm to produce all permutations only performs
// one swap from one to the next permutation. I.e. the complete
// length can be updated for "only" that swap insted on computing
// the complete sum. – Updating for one swap seem to be too coplex
// to pay off.
//func AsymGreedyf32_Alt(size uint, m *AdjMxAf32) (path []uint, plen float64) {
//	switch size {
//	case 0:
//		return nil, 0
//	case 1:
//		return []uint{0}, 0
//	}
//	L := size - 1
//	swap := func(w float32, perm []uint, i, j uint) float32 {
//		// TODO does not work correct
//		if j < i {
//			i, j = j, i
//		}
//		pi, pj := perm[i], perm[j]
//		if i+1 == j { // …pi → pj…
//			w -= m.Get(pi, pj)
//			w += m.Get(pj, pi)
//		} else { // …pi → p … p → pj…
//			p := perm[i+1]
//			w -= m.Get(pi, p)
//			w += m.Get(pj, p)
//			p = perm[j-1]
//			w -= m.Get(p, pj)
//			w += m.Get(p, pi)
//		}
//		// …p → pi…
//		if i == 0 {
//			w -= m.Get(L, pi)
//			w += m.Get(L, pj)
//		} else {
//			p := perm[i-1]
//			w -= m.Get(p, pi)
//			w += m.Get(p, pj)
//		}
//		// …pj → p…
//		if j+1 == L {
//			w -= m.Get(pj, L)
//			w += m.Get(pi, L)
//		} else {
//			p := perm[j+1]
//			w -= m.Get(pj, p)
//			w += m.Get(pi, p)
//		}
//		perm[j], perm[i] = pi, pj
//		return w
//	}
//	path = make([]uint, size)
//	// start with L → 0 → 1 → … → L
//	path[L] = L
//	best := m.Get(L, 0)
//	for k := uint(0); k < L; k++ {
//		path[k] = k
//		best += m.Get(k, k+1)
//	}
//	//	fmt.Printf("best: %.5f\n", best)
//	//	pplen(m, path)
//	perm := make([]uint, L)
//	copy(perm, path[:L])
//	c := make([]uint, L) // automatic set to 0 (go!)
//	i := uint(0)
//	for i < L {
//		if c[i] < i {
//			var curl float32
//			if (i & 1) == 0 {
//				curl = swap(best, perm, 0, i)
//			} else {
//				curl = swap(best, perm, c[i], i)
//			}
//			if curl < best {
//				copy(path[:L], perm)
//				best = curl
//				//				fmt.Printf("best: %.5f\n", best)
//				//				pplen(m, path)
//			}
//			c[i]++
//			i = 0
//		} else {
//			c[i] = 0
//			i++
//		}
//	}
//	return path, float64(best)
//}
