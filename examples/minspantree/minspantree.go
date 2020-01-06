package main

import (
	"flag"
	"fmt"
	"os"

	"git.fractalqb.de/fractalqb/groph/shortestpath"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/util"
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
	dot := util.GraphViz{
		PerNodeAtts: func(g groph.RGraph, v groph.VIdx) util.GvAtts {
			res := util.GvAtts{"label": fmt.Sprintf("%c / %d", 'a'+v, v)}
			if v == mst.Root() {
				res["shape"] = "diamond"
			}
			return res
		},
		PerEdgeAtts: func(g groph.RGraph, u, v groph.VIdx) (res util.GvAtts) {
			if mst.Edge(u, v) {
				res = util.GvAtts{"label": fmt.Sprint(dists[v])}
				res["color"] = "blue"
			} else {
				res = util.GvAtts{"label": util.GvNoLabel}
				res["color"] = "gray"
			}
			return res
		},
	}

	const mstStart = 8
	if *undir {
		ug := groph.NewAdjMxUbool(9, nil)
		groph.Set(ug, true, edges...)
		dists, mst = (&shortestpath.DijkstraBool{}).On(ug, mstStart, dists, mst)
		dot.Write(os.Stdout, ug, "")
	} else {
		dg := groph.NewAdjMxDbool(9, nil)
		groph.Set(dg, true, edges...)
		dists, mst = (&shortestpath.DijkstraBool{}).On(dg, mstStart, dists, mst)
		dot.Write(os.Stdout, dg, "")
	}
}
