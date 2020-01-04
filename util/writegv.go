package util

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"git.fractalqb.de/fractalqb/groph"
)

type GraphViz struct {
	GraphAtts   func(g groph.RGraph) map[string]interface{}
	NodeAtts    func(g groph.RGraph) map[string]interface{}
	EdgeAtts    func(g groph.RGraph) map[string]interface{}
	PerNodeAtts func(g groph.RGraph, v groph.VIdx) map[string]interface{}
	PerEdgeAtts func(g groph.RGraph, u, v groph.VIdx) map[string]interface{}
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
	var atts map[string]interface{}
	if gv.PerNodeAtts != nil {
		atts = gv.PerNodeAtts(g, v)
	}
	var sb strings.Builder
	if label != nil {
		fmt.Fprintf(&sb, "label=\"%s\"", label)
		for k, att := range atts {
			if k == "label" {
				continue
			}
			sb.WriteString(", ")
			fmt.Fprintf(&sb, ", %s=\"%s\"", k, att)
		}
	} else {
		sep := false
		for k, att := range atts {
			if sep {
				sb.WriteString(", ")
			}
			fmt.Fprintf(&sb, "%s=\"%s\"", k, att)
			sep = true
		}
	}
	return sb.String()
}

func (gv *GraphViz) eAtts(g groph.RGraph, u, v groph.VIdx) string {
	w := g.Weight(u, v)
	if gv.PerEdgeAtts == nil {
		return fmt.Sprintf("label=\"%s\"", w)
	}
	atts := gv.PerEdgeAtts(g, u, v)
	var sb strings.Builder
	sep := false
	if lb, ok := atts["label"]; ok {
		if lb != "" {
			fmt.Fprintf(&sb, "label=\"%s\"", lb)
			sep = true
		}
	} else {
		fmt.Fprintf(&sb, "label=\"%s\"", w)
		sep = true
	}
	for k, att := range atts {
		if k == "label" {
			continue
		}
		if sep {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%s=\"%s\"", k, att)
		sep = true
	}
	return sb.String()
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
	gvWriteAtts(tw, g, "graph", gv.GraphAtts)
	gvWriteAtts(tw, g, "node", gv.NodeAtts)
	gvWriteAtts(tw, g, "edge", gv.EdgeAtts)
	if len(vlabels) > 0 || gv.PerNodeAtts != nil {
		for i := groph.V0; i < g.Order(); i++ {
			atts := gv.nAtts(g, i, vlabels)
			if atts != "" {
				fmt.Fprintf(tw, "\t%d\t[%s];\n", i, atts)
			}
		}
	}
	groph.EachEdge(g, func(u, v groph.VIdx) bool {
		atts := gv.eAtts(g, u, v)
		if atts == "" {
			fmt.Fprintf(tw, "\t%d -> %d;\n", u, v)
		} else {
			fmt.Fprintf(tw, "\t%d -> %d [%s];\n", u, v, atts)
		}
		return false
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
	gvWriteAtts(tw, g, "graph", gv.GraphAtts)
	gvWriteAtts(tw, g, "node", gv.NodeAtts)
	gvWriteAtts(tw, g, "edge", gv.EdgeAtts)
	if len(vlabels) > 0 || gv.PerNodeAtts != nil {
		for i := groph.V0; i < g.Order(); i++ {
			atts := gv.nAtts(g, i, vlabels)
			if atts != "" {
				fmt.Fprintf(tw, "\t%d\t[%s];\n", i, atts)
			}
		}
	}
	groph.EachEdge(g, func(u, v groph.VIdx) bool {
		fmt.Fprintf(tw, "\t%d -- %d [label=\"%v\"];\n", u, v, g.Weight(u, v))
		return false
	})
	fmt.Fprintln(tw, "}")
}

func gvWriteAtts(
	wr io.Writer,
	g groph.RGraph,
	tag string,
	attFn func(groph.RGraph) map[string]interface{},
) {
	if attFn == nil {
		return
	}
	atts := attFn(g)
	if len(atts) > 0 {
		fmt.Fprintf(wr, "\t%s\t[", tag)
		sep := false
		for k, att := range atts {
			if sep {
				fmt.Fprint(wr, ", ")
			}
			fmt.Fprintf(wr, "%s=\"%s\"", k, att)
		}
		fmt.Fprintln(wr, "];")
	}
}
