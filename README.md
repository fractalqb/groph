# groph

[![Build Status](https://travis-ci.org/fractalqb/groph.svg)](https://travis-ci.org/fractalqb/groph)
[![codecov](https://codecov.io/gh/fractalqb/groph/branch/master/graph/badge.svg)](https://codecov.io/gh/fractalqb/groph)
[![Go Report Card](https://goreportcard.com/badge/git.fractalqb.de/fractalqb/groph)](https://goreportcard.com/report/git.fractalqb.de/fractalqb/groph)
[![GoDoc](https://godoc.org/git.fractalqb.de/fractalqb/groph?status.svg)](https://godoc.org/git.fractalqb.de/fractalqb/groph)

`import "git.fractalqb.de/fractalqb/groph"`

---

A pure Go library of graphs and their algorithms. More info at [godoc.org](https://godoc.org/git.fractalqb.de/fractalqb/groph).

## Catchy Examples… hopefully
On can create the Graphviz .dot file for a directed graph:

![Simple graph](./examples/4readme/plain.png)

with this lines of code:

```go
func writePlain(wr io.Writer) {
	g := groph.NewAdjMxDbool(9, nil)
	type E = groph.Edge
	groph.Set(g, true, E{0, 1}, E{1, 3}, E{3, 2}, E{2, 0}, E{4, 3},
		E{4, 5}, E{5, 6}, E{6, 4}, E{7, 4}, E{8, 7})
	graphviz.Writer{}.Write(wr, g, "")
}
```

If you want it more fancy and show the embedded MST in the graph…

![Fancy graph](./examples/4readme/fancy.png)

use Dijkstra's algorithm and configure the Graphviz writer to highlight it:

```go
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
```

## Graph Implementations
The following table shows the supported graph implementations along
with their relative access performance, i.e. read & write. Access
performance is the factor relative to the fastest implementation –
the one with 1 in the `t/*` column.
Each implementation can be accessed through their `WGraph` interface
_(t/gen)_ or through their specifically typed interface _(t/typed)_.

| Implementation   | Weight Type       | Dir/Undir | t/typed | t/gen |
|------------------|:-----------------:|:---------:|--------:|------:|
| Adjacency matrix | bool (bitmap)     | D         | 1.51    | 3.68  |
| Adjacency matrix | bool              | D         | 1       | 2.01  |
| Adjacency matrix | int32             | D, U      | 1.08    | 8.18  |
| Adjacency matrix | float32           | D, U      | 1.12    | 8.35  |
| Slice of Maps    | interface\{\}     | D, U      | –       | 26.07 |
| Slice of Maps    | int32             | D, U      | 15.92   | 25.22 |
| Slice of Maps    | float32           | D, U      | 15.90   | 25.90 |

_Note:_ Performance losses for generic access are mainly due to the type cast to a
specific type after calling `g.Weight()`. Because this is probably relevant for real
applications, this remains part of the benchmarks.

## Algorithms

| Algorithm | Problem | Weight Types |
|-----------|---------|--------------|
| Depth First | Traversal (tvr) | interface\{\} |
| Floyd Warshall | Shortest path (sp) | int32, float32 |
| Dijkstra | Shortest path (sp), Minimal spannig Tree | int32, float32 |
| Kruskal | Minimal spannig Tree (msp) | int32, float32 |
| TSP greedy | Travelling Salesman (tsp) | float32 |
| 2-Opt | Travelling Salesman (tsp) | float32 |

## Performance compared to other Libraries

Comparison benchmarks can be found in the separate project [groph-cmpbench](https://codeberg.org/fractalqb/groph-cmpbench).

### Access Performance

Access performance is measured by writing and reading graph edges of a graph with a fixed number of vertices. While this does not tell anything about the quality of the provided algorithms theses operations are frequently called in inner loops of most algorithms. I.e. access performance will make a factor of algorithm's performance.

Other graph implementations are run against the groph top-speed graph. Use the numbers from groph's internal benchmark to estimate other comparisons.

_Feedback on the benchmark project is very welcome to improve the validity of the comparison!_

| Library | Test | t-Factor |
|---------|------|-------:|
| groph   | access directed int32 | 1 |
| [yourbasic graph](https://github.com/yourbasic/graph) | directed int64 | 44 |
| [goraph](https://github.com/gyuho/goraph) | directed float64 | 737 |
| [thcyron graphs](https://github.com/thcyron/graphs) | directed float64 | 1145 |

_t-Factor: smaller is better_
