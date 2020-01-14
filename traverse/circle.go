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
