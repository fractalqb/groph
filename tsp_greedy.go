package graph

func TspGreedyf32(m *AdjMxAf32) (path []uint, plen float64) {
	size := m.VertexNo()
	switch size {
	case 0:
		return nil, 0
	case 1:
		return []uint{0}, 0
	}
	L := size - 1
	path = make([]uint, size)
	// start with L → 0 → 1 → … → L
	path[L] = L
	_, best := m.Edge(L, 0)
	for k := uint(0); k < L; k++ {
		path[k] = k
		_, w := m.Edge(k, k+1)
		best += w
	}
	//	fmt.Printf("best: %.5f\n", best)
	//	pplen(m, path)
	perm := make([]uint, L)
	copy(perm, path[:L])
	c := make([]uint, L) // automatic set to 0 (go!)
	i := uint(0)
	for i < L {
		if c[i] < i {
			if (i & 1) == 0 {
				perm[0], perm[i] = perm[i], perm[0]
			} else {
				perm[c[i]], perm[i] = perm[i], perm[c[i]]
			}
			_, curl := m.Edge(L, perm[0])
			_, w := m.Edge(perm[L-1], L)
			curl += w
			for i := uint(0); i+1 < L; i++ {
				_, w = m.Edge(perm[i], perm[i+1])
				curl += w
			}
			if curl < best {
				copy(path[:L], perm)
				best = curl
				//				fmt.Printf("best: %.5f\n", best)
				//				pplen(m, path)
			}
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return path, float64(best)
}
