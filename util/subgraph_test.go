package util

import "git.fractalqb.de/fractalqb/groph"

var (
	_ groph.RGraph = RSubgraph{}
	_ groph.WGraph = WSubgraph{}
	// _ RUndirected = RSubUndir{}
	// _ WUndirected = WSubUndir{}
)
