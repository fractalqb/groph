package groph

import (
	"reflect"
)

type SliceNMeasure struct {
	slc reflect.Value
	m   reflect.Value
	dir bool
}

func NewSliceNMeasure(
	vertexSlice interface{},
	measure interface{},
	directed bool) *SliceNMeasure {
	res := &SliceNMeasure{
		slc: reflect.ValueOf(vertexSlice),
		m:   reflect.ValueOf(measure),
		dir: directed,
	}
	return res
}

func (g *SliceNMeasure) VertexNo() uint {
	return uint(g.slc.Len())
}

func (g *SliceNMeasure) Vertex(idx uint) Vertex {
	return g.slc.Index(int(idx)).Interface()
}

func (g *SliceNMeasure) Directed() bool {
	return g.dir
}

func (g *SliceNMeasure) Weight(fromIdx, toIdx uint) interface{} {
	f, t := g.slc.Index(int(fromIdx)), g.slc.Index(int(toIdx))
	d := g.m.Call([]reflect.Value{f, t})
	return d[0].Interface()
}
