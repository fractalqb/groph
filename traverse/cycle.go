package traverse

import (
	"git.fractalqb.de/fractalqb/groph"
)

func HasCycle(g groph.RGraph, reuse *Search) bool {
	if reuse == nil {
		reuse = NewSearch(g)
	} else {
		reuse.Reset(g)
	}
	return reuse.OutDepth1st(false,
		func(pred, v groph.VIdx, vHits int, cluster int) bool {
			if vHits == 0 {
				return false
			}
			if _, fin := reuse.State(v); fin {
				return false
			}
			return true
		})
}

type Cycle struct {
	search *Search
	path   []groph.VIdx
}

func NewCycle(reuse *Search) *Cycle {
	if reuse == nil {
		reuse = NewSearch(nil)
	}
	return &Cycle{search: reuse}
}

func (c *Cycle) FindOne(g groph.RGraph, do func(cycle []groph.VIdx)) {
	c.search.Reset(g)
	if c.path != nil {
		c.path = c.path[:0]
	}
	c.search.OutDepth1st(false,
		func(u, v groph.VIdx, vHits int, cluster int) bool {
			if vHits > 0 {
				if _, fin := c.search.State(v); fin {
					return false
				}
				cStart := c.backTo(v)
				do(c.path[cStart:])
				return true
			}
			if btu := c.backTo(u); btu+1 != len(c.path) {
				c.path = c.path[:btu]
			}
			c.path = append(c.path, v)
			return false
		})
}

func (c *Cycle) backTo(v groph.VIdx) int {
	if v < 0 {
		return len(c.path) - 1
	}
	for i := len(c.path) - 1; i >= 0; i-- {
		if c.path[i] == v {
			return i
		}
	}
	panic("unreachable")
}
