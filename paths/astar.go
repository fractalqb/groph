package paths

import (
	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/internal"
	"golang.org/x/exp/constraints"
)

type astKnown[W constraints.Ordered] struct {
	back groph.VIdx
	g, h W
	done bool
}

func AStarU[W constraints.Ordered](
	g groph.RUndirected[W],
	from, to groph.VIdx,
	h func(u, v groph.VIdx) W,
) (p []groph.VIdx, l W) {
	var zero W
	var todo internal.VHeap[W]
	todo.AddVertex(to, zero)
	known := map[groph.VIdx]*astKnown[W]{to: {back: -1}}
	for todo.Len() > 0 {
		cn, _ := todo.PopVertex()
		if cn == from {
			l = known[cn].g
			for cn >= 0 {
				p = append(p, cn)
				ck := known[cn]
				cn = ck.back
			}
			return p, l
		}
		ck := known[cn]
		ck.done = true
		g.EachAdjacent(cn, func(v groph.VIdx) error {
			vk := known[v]
			switch {
			case vk == nil:
				vk := &astKnown[W]{
					back: cn,
					g:    ck.g + g.Edge(cn, v),
					h:    h(v, from),
				}
				known[v] = vk
				todo.PushVertex(v, vk.g+vk.h)
			case !vk.done:
				g := ck.g + g.Edge(cn, v)
				if g < vk.g {
					vk.back = cn
					vk.g = g
					todo.Set(v, g+vk.h)
				}
			}
			return nil
		})
	}
	return nil, zero
}
