package internal

import (
	"container/heap"

	"golang.org/x/exp/constraints"
)

type vhItem[W constraints.Ordered] struct {
	v int
	w W
}

type VHeap[W constraints.Ordered] struct {
	v2idx []int
	items []vhItem[W]
}

func NewVHeap[W constraints.Ordered](cap int) *VHeap[W] {
	return &VHeap[W]{
		v2idx: make([]int, 0, cap),
		items: make([]vhItem[W], 0, cap),
	}
}

func (h *VHeap[W]) AddVertex(v int, w W) { h.Push(vhItem[W]{v: v, w: w}) }

func (h *VHeap[W]) AddNotVertex(v int) {
	if len(h.v2idx) <= v {
		tmp, cp := Slice(h.v2idx, v+1)
		if cp {
			copy(tmp, h.v2idx)
		}
		for i := len(h.v2idx); i < v; i++ {
			tmp[i] = -1
		}
		h.v2idx = tmp
	}
	h.v2idx[v] = -1
}

func (h *VHeap[W]) PushVertex(v int, w W) {
	heap.Push(h, vhItem[W]{v: v, w: w})
}

func (h *VHeap[W]) PopVertex() (v int, w W) {
	itm := heap.Pop(h).(vhItem[W])
	return itm.v, itm.w
}

func (h *VHeap[W]) Has(v int) bool { return h.v2idx[v] >= 0 }

func (h *VHeap[W]) Peek(v int) W {
	i := h.v2idx[v]
	return h.items[i].w
}

func (h *VHeap[W]) Set(v int, w W) {
	i := h.v2idx[v]
	h.items[i].w = w
	heap.Fix(h, i)
}

func (h *VHeap[W]) Len() int { return len(h.items) }

func (h *VHeap[W]) Less(i, j int) bool {
	return h.items[i].w < h.items[j].w
}

func (h *VHeap[W]) Swap(i, j int) {
	h.items[j], h.items[i] = h.items[i], h.items[j]
	h.v2idx[h.items[i].v] = i
	h.v2idx[h.items[j].v] = j
}

func (h *VHeap[W]) Push(x interface{}) {
	itm := x.(vhItem[W])
	if len(h.v2idx) <= itm.v {
		tmp, cp := Slice(h.v2idx, itm.v+1)
		if cp {
			copy(tmp, h.v2idx)
		}
		for i := len(h.v2idx); i < itm.v; i++ {
			tmp[i] = -1
		}
		h.v2idx = tmp
	}
	h.v2idx[itm.v] = len(h.items)
	h.items = append(h.items, itm)
}

func (h *VHeap[W]) Pop() interface{} {
	lm1 := len(h.items) - 1
	itm := h.items[lm1]
	h.items = h.items[:lm1]
	h.v2idx[itm.v] = -1
	return itm
}
