package util

import (
	"fmt"
	"io"
	"text/tabwriter"

	"git.fractalqb.de/fractalqb/groph/traverse"

	"git.fractalqb.de/fractalqb/groph"
)

type GraphViz struct{}

func (gv *GraphViz) Write(
	wr io.Writer,
	g groph.RGraph,
	name string,
	vlabels ...interface{},
) {
	if u, ok := g.(groph.RUndirected); ok {
		gv.uwrite(wr, u, name, vlabels)
	} else {
		gv.dwrite(wr, g, name, vlabels)
	}
}

func (gv *GraphViz) dwrite(
	wr io.Writer,
	g groph.RGraph,
	name string,
	vlabels []interface{},
) {
	if name == "" {
		name = "G"
	}
	tw := tabwriter.NewWriter(wr, 2, 0, 1, ' ', 0)
	fmt.Fprintf(tw, "digraph %s {\n", name)
	if len(vlabels) > 0 {
		n := g.Order()
		if len(vlabels) < n {
			n = len(vlabels)
		}
		for i := 0; i < n; i++ {
			fmt.Fprintf(tw, "\t%d\t[label=\"%s\"];\n", i, vlabels[i])
		}
	}
	traverse.EachEdge(g, func(u, v groph.VIdx) {
		fmt.Fprintf(tw, "\t%d -> %d [label=\"%v\"];\n", u, v, g.Weight(u, v))
	})
	fmt.Fprintln(tw, "}")
}

func (gv *GraphViz) uwrite(
	wr io.Writer,
	g groph.RUndirected,
	name string,
	vlabels []interface{},
) {
	if name == "" {
		name = "G"
	}
	tw := tabwriter.NewWriter(wr, 2, 0, 1, ' ', 0)
	fmt.Fprintf(tw, "graph %s {\n", name)
	if len(vlabels) > 0 {
		n := g.Order()
		if len(vlabels) < n {
			n = len(vlabels)
		}
		for i := 0; i < n; i++ {
			fmt.Fprintf(tw, "\t%d\t[label=\"%s\"];\n", i, vlabels[i])
		}
	}
	traverse.EachEdge(g, func(u, v groph.VIdx) {
		fmt.Fprintf(tw, "\t%d -- %d [label=\"%v\"];\n", u, v, g.Weight(u, v))
	})
	fmt.Fprintln(tw, "}")
}
