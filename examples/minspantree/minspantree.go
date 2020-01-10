package main

import (
	"flag"
	"fmt"
	"os"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmatrix"
	"git.fractalqb.de/fractalqb/groph/shortestpath"
	"git.fractalqb.de/fractalqb/groph/util/graphviz"
)

type E = groph.Edge

var edges = []E{
	E{0, 1}, E{1, 3}, E{3, 2}, E{2, 0}, E{4, 3},
	E{4, 5}, E{5, 6}, E{6, 4}, E{7, 4}, E{8, 7},
}

func main() {
	undir := flag.Bool("u", false, "with undirected graph")
	flag.Parse()

	var mst = groph.Tree{}
	var dists []int
	dot := graphviz.Writer{
		PerNodeAtts: func(g groph.RGraph, v groph.VIdx) graphviz.Attributes {
			res := graphviz.Attributes{"label": fmt.Sprintf("%c / %d", 'a'+v, v)}
			if v == mst.Root() {
				res["shape"] = "diamond"
			}
			return res
		},
		PerEdgeAtts: func(g groph.RGraph, u, v groph.VIdx) (res graphviz.Attributes) {
			if mst.Edge(u, v) {
				res = graphviz.Attributes{"label": fmt.Sprint(dists[v])}
				res["color"] = "blue"
			} else {
				res = graphviz.Attributes{"label": graphviz.NoLabel}
				res["color"] = "gray"
			}
			return res
		},
	}

	const mstStart = 8
	if *undir {
		ug := adjmatrix.NewUBool(9, nil)
		groph.Set(ug, true, edges...)
		dists, mst = (&shortestpath.DijkstraBool{}).On(ug, mstStart, dists, mst)
		dot.Write(os.Stdout, ug, "")
	} else {
		dg := adjmatrix.NewDBool(9, nil)
		groph.Set(dg, true, edges...)
		dists, mst = (&shortestpath.DijkstraBool{}).On(dg, mstStart, dists, mst)
		dot.Write(os.Stdout, dg, "")
	}
}
