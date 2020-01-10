package sliceofmaps

import (
	"testing"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/tests"
)

var (
	_ groph.WGraph      = (*SoMD)(nil)
	_ groph.OutLister   = (*SoMD)(nil)
	_ groph.WUndirected = (*SoMU)(nil)
	_ groph.OutLister   = (*SoMU)(nil)
)

func TestSoMD(t *testing.T) {
	g := NewSoMD(tests.SetDelSize, nil)
	tests.GenericSetDelTest(t, g, tests.GenProbe)
}

func TestSoMU(t *testing.T) {
	g := NewSoMU(tests.SetDelSize, nil)
	tests.GenericSetDelTest(t, g, tests.GenProbe)
}

func BenchmarkSoMD_generic(b *testing.B) {
	m := NewSoMD(tests.SetDelSize, nil)
	max := m.Order()
	for n := 0; n < b.N; n++ {
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				m.SetWeight(i, j, n)
			}
		}
		for i := 0; i < max; i++ {
			for j := 0; j < max; j++ {
				r := m.Weight(i, j)
				if r != n {
					b.Fatal("unexpected read", n, r)
				}
			}
		}
	}
}
