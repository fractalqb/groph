package groph

import (
	"math/rand"
	"testing"
)

func TestFloydWarshallDirEqUndir(t *testing.T) {
	const VNO = 7
	mu := NewAdjMxUf32(VNO, nil)
	md := NewAdjMxDf32(mu.VertexNo(), nil)
	for i := uint(0); i < VNO; i++ {
		mu.SetEdge(i, i, 0)
		for j := i + 1; j < VNO; j++ {
			w := rand.Float32()
			mu.SetEdge(i, j, w)
		}
	}
	CpWeights(md, mu)
	FloydWarshallf32(md)
	FloydWarshallf32(mu)
	for i := uint(0); i < VNO; i++ {
		for j := uint(0); j < VNO; j++ {
			if md.Edge(i, j) != mu.Edge(i, j) {
				t.Errorf("differ@ %d,%d: %f != %f", i, j, md.Edge(i, j), mu.Edge(i, j))
			}
		}
	}
}
