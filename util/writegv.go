package util

import (
	"fmt"
	"io"
	"text/tabwriter"

	"git.fractalqb.de/fractalqb/groph/traverse"

	"git.fractalqb.de/fractalqb/groph"
)

type GraphViz struct {
	NodeAtts func(g groph.RGraph, v groph.VIdx) (atts string)
	EdgeAtts func(g groph.RGraph, u, v groph.VIdx) (atts string, hasLabel bool)
}

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

func (gv *GraphViz) nAtts(g groph.RGraph, v groph.VIdx, vlbs []interface{}) string {
	var label interface{}
	if v < len(vlbs) {
		label = vlbs[v]
	}
	var atts string
	if gv.NodeAtts != nil {
		atts = gv.NodeAtts(g, v)
	}
	if label != nil {
		if atts == "" {
			return fmt.Sprintf("label=\"%s\"", label)
		} else {
			return fmt.Sprintf("label=\"%s\", %s", label, atts)
		}
	}
	return atts
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
	if len(vlabels) > 0 || gv.NodeAtts != nil {
		for i := groph.V0; i < g.Order(); i++ {
			atts := gv.nAtts(g, i, vlabels)
			if atts != "" {
				fmt.Fprintf(tw, "\t%d\t[%s];\n", i, atts)
			}
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
	if len(vlabels) > 0 || gv.NodeAtts != nil {
		for i := groph.V0; i < g.Order(); i++ {
			atts := gv.nAtts(g, i, vlabels)
			if atts != "" {
				fmt.Fprintf(tw, "\t%d\t[%s];\n", i, atts)
			}
		}
	}
	traverse.EachEdge(g, func(u, v groph.VIdx) {
		fmt.Fprintf(tw, "\t%d -- %d [label=\"%v\"];\n", u, v, g.Weight(u, v))
	})
	fmt.Fprintln(tw, "}")
}
