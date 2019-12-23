package groph

import (
	"reflect"
	"testing"
)

func TestSliceNMeasure(t *testing.T) {
	slc := []float64{1, 2, 3}
	a := NewSliceNMeasure(slc, func(a, b float64) float64 {
		return b - a
	}, true)
	if vn := a.VertexNo(); vn != VIdx(len(slc)) {
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
