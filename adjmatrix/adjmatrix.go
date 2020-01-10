// Package adjmatrix provides adjacency matrix implementations for graph
// interfaces.
package adjmatrix

import (
	"errors"
	"math"

	"git.fractalqb.de/fractalqb/groph"
)

// I32Del is used by default in adjacency matrices with edge weight
// type 'int32' to mark edges that do not exist.
//
// See AdjMxDi32 and AdjMxUi32
const I32Del = -2147483648 // min{ int32 }

// nSum computes the sum of the n 1st integers, i.e. 1+2+3+…+n
func nSum(n int) int { return n * (n + 1) / 2 }

func nSumRev(n int) float64 {
	r := math.Sqrt(0.25 + 2*float64(n))
	return r - 0.5
}

// uIdx computes the index into the weight slice of an undirected matrix
func uIdx(i, j groph.VIdx) groph.VIdx { return nSum(i) + j }

func dOrdFromLen(sliceLen int) (order int, err error) {
	order = int(math.Sqrt(float64(sliceLen)))
	if order*order != sliceLen {
		return order, errors.New("weights slice is not square")
	}
	return order, nil
}

func uOrdFromLen(sliceLen int) (order int, err error) {
	order = int(nSumRev(sliceLen))
	if nSum(order) != sliceLen {
		return order, errors.New("weights slice len is not sum(1,…,n)")
	}
	return order, nil
}

type adjMx struct {
	ord int
}

func (m *adjMx) Order() int { return m.ord }
