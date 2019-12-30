package groph

import (
	"reflect"
	"testing"
)

var (
	_ RGraph = (*RSubgraph)(nil)
	_ WGraph = (*WSubgraph)(nil)
)

func TestSliceNMeasure(t *testing.T) {
	slc := []float64{1, 2, 3}
	a := NewPointsNDist(slc, func(a, b float64) float64 {
		return b - a
	})
	if vn := a.Order(); vn != VIdx(len(slc)) {
		t.Fatalf("unexpected vertex no: %d (expected %d)", vn, len(slc))
	}
	for i := 0; i < len(slc); i++ {
		v, ok := a.Vertex(VIdx(i)).(float64)
		if !ok {
			t.Fatalf("unexpected vertex type[%d]: %s", i, reflect.TypeOf(a.Vertex(VIdx(i))))
		}
		if v != slc[i] {
			t.Fatalf("unexpected vertex value[%d]: %f", i, v)
		}
	}
	w := a.Weight(0, 1)
	if w != 1.0 {
		t.Fatalf("unexpected edge weight: %f", w)
	}
}
