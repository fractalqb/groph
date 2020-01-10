package tsp

import (
	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmatrix"
)

func Greedyf32(m groph.RGf32) (path []groph.VIdx, plen float32) {
	size := m.Order()
	switch size {
	case 0:
		return nil, 0
	case 1:
		return []groph.VIdx{0}, 0
	}
	L := size - 1
	path = make([]groph.VIdx, size)
	// start with L → 0 → 1 → … → L
	path[L] = L
	best := m.Edge(L, 0)
	for k := 0; k < L; k++ {
		path[k] = k
		best += m.Edge(k, k+1)
	}
	perm := make([]groph.VIdx, L)
	copy(perm, path)
	c := make([]groph.VIdx, L) // automatic set to 0 (go!)
	i := 0
	for i < L {
		if c[i] < i {
			if (i & 1) == 0 {
				perm[0], perm[i] = perm[i], perm[0]
			} else {
				perm[c[i]], perm[i] = perm[i], perm[c[i]]
			}
			curl := m.Edge(L, perm[0])
			curl += m.Edge(perm[L-1], L)
			for i := 0; i+1 < L; i++ {
				curl += m.Edge(perm[i], perm[i+1])
			}
			if curl < best {
				copy(path[:L], perm)
				best = curl
			}
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return path, best
}

func GreedyAdjMxDf32(m *adjmatrix.DFloat32) (path []groph.VIdx, plen float32) {
	size := m.Order()
	switch size {
	case 0:
		return nil, 0
	case 1:
		return []groph.VIdx{0}, 0
	}
	L := size - 1
	path = make([]groph.VIdx, size)
	// start with L → 0 → 1 → … → L
	path[L] = L
	best := m.Edge(L, 0)
	for k := 0; k < L; k++ {
		path[k] = k
		best += m.Edge(k, k+1)
	}
	perm := make([]groph.VIdx, L)
	copy(perm, path)
	c := make([]groph.VIdx, L) // automatic set to 0 (go!)
	i := 0
	for i < L {
		if c[i] < i {
			if (i & 1) == 0 {
				perm[0], perm[i] = perm[i], perm[0]
			} else {
				perm[c[i]], perm[i] = perm[i], perm[c[i]]
			}
			curl := m.Edge(L, perm[0])
			curl += m.Edge(perm[L-1], L)
			for i := 0; i+1 < L; i++ {
				curl += m.Edge(perm[i], perm[i+1])
			}
			if curl < best {
				copy(path[:L], perm)
				best = curl
			}
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return path, best
}
