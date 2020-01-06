package main

import (
	"fmt"
	"io"
	"os"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/shortestpath"
	"git.fractalqb.de/fractalqb/groph/util/graphviz"
)

func writePlain(wr io.Writer) {
	g := groph.NewAdjMxDbool(9, nil)
	type E = groph.Edge
	groph.Set(g, true, E{0, 1}, E{1, 3}, E{3, 2}, E{2, 0}, E{4, 3},
		E{4, 5}, E{5, 6}, E{6, 4}, E{7, 4}, E{8, 7})
	graphviz.Writer{}.Write(wr, g, "")
}

func writeFancy(wr io.Writer) {
	g := groph.NewAdjMxDbool(9, nil)
	type E = groph.Edge
	groph.Set(g, true, E{0, 1}, E{1, 3}, E{3, 2}, E{2, 0}, E{4, 3},
		E{4, 5}, E{5, 6}, E{6, 4}, E{7, 4}, E{8, 7})

	// Compute distances and minimal spanning tree starting at vertex 8
	dists, mst := (&shortestpath.DijkstraBool{}).On(g, 8, nil, []groph.VIdx{})

	// Tell Graphviz writer how to set the correct node and edge attributes
	dot := graphviz.Writer{
		GraphAtts: graphviz.AttMap(graphviz.Attributes{"rankdir": "LR"}),
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

	// Write the dot file
	dot.Write(wr, g, "")
}

func main() {
	wr, _ := os.Create("plain.dot")
	writePlain(wr)
	wr.Close()
	wr, _ = os.Create("fancy.dot")
	writeFancy(wr)
	wr.Close()
}
