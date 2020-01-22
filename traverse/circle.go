package traverse

import "git.fractalqb.de/fractalqb/groph"

func HasCycle(g groph.RGraph, reuse *Search) bool {
	if reuse == nil {
		reuse = NewSearch(g)
	} else {
		reuse.Reset(g)
	}
	return reuse.OutDepth1st(false,
		func(pred, v groph.VIdx, circle bool, cluster int) bool {
			if circle {
				return true
			}
			return false
		})
}

type Circle struct {
	path   []groph.VIdx
	OnFind func(vs []groph.VIdx) (stop bool)
}

func (c *Circle) Search(u, v groph.VIdx, circle bool, cluster int) bool {
	if u < 0 {
		if c.path == nil {
			c.path = []groph.VIdx{v}
		} else {
			c.path = c.path[:1]
			c.path[0] = v
		}
		return false
	}
	if ui := c.backTo(u); ui < 0 {
		if c.path != nil {
			c.path = c.path[:0]
		}
	} else {
		c.path = c.path[:ui+1]
	}
	if circle {
		if vi := c.backTo(v); vi >= 0 {
			return c.OnFind(c.path[vi:])
		}
	}
	c.path = append(c.path, v)
	return false
}

func (c *Circle) backTo(v groph.VIdx) int {
	if v < 0 {
		return -1
	}
	for i := len(c.path) - 1; i >= 0; i-- {
		if c.path[i] == v {
			return i
		}
	}
	return -1
}
