package main

import (
	"fmt"
	"io"
	"os"

	"git.fractalqb.de/fractalqb/groph"
	"git.fractalqb.de/fractalqb/groph/adjmatrix"
	"git.fractalqb.de/fractalqb/groph/shortestpath"
	"git.fractalqb.de/fractalqb/groph/util/graphviz"
)

func writePlain(wr io.Writer) {
	g := adjmatrix.NewDBool(9, nil)
	type E = groph.Edge
	groph.Set(g, true, E{0, 1}, E{1, 3}, E{3, 2}, E{2, 0}, E{4, 3},
		E{4, 5}, E{5, 6}, E{6, 4}, E{7, 4}, E{8, 7})
	graphviz.Writer{}.Write(wr, g, "")
}

func writeFancy(wr io.Writer) {
	g := adjmatrix.NewDBool(9, nil)
	type E = groph.Edge
	groph.Set(g, true, E{0, 1}, E{1, 3}, E{3, 2}, E{2, 0}, E{4, 3},
		E{4, 5}, E{5, 6}, E{6, 4}, E{7, 4}, E{8, 7})

	// Compute distances and minimal spanning tree starting at vertex 8
	dists, mst := (&shortestpath.DijkstraBool{}).On(g, 8, nil, []groph.VIdx{})

	// Tell Graphviz writer how to set the correct node and edge attributes
	dot := graphviz.Writer{
		GraphAtts: graphviz.AttMap(graphviz.Attributes{"rankdir": "LR"}),
		NodeAtts:  graphviz.AttMap(graphviz.Attributes{"shape": "box"}),
		PerNodeAtts: func(g groph.RGraph, v groph.VIdx) graphviz.Attributes {
			atts := graphviz.Attributes{"label": fmt.Sprintf("%c / %d", 'a'+v, v)}
			if v == mst.Root() {
				atts["shape"] = "circle"
			} else if groph.Degree(mst, v) == 1 {
				atts["shape"] = "doublecircle"
			}
			return atts
		},
		PerEdgeAtts: func(g groph.RGraph, u, v groph.VIdx) (atts graphviz.Attributes) {
			if mst.Edge(v, u) {
				atts = graphviz.Attributes{
					"label": fmt.Sprint(dists[v]),
					"color": "blue"}
			} else {
				atts = graphviz.Attributes{
					"label": graphviz.NoLabel,
					"color": "gray"}
			}
			return atts
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
