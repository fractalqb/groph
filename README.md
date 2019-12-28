# groph

[![Test Coverage](https://img.shields.io/badge/coverage-58%25-orange.svg)](file:coverage.html)
[![Go Report Card](https://goreportcard.com/badge/codeberg.org/fractalqb/groph)](https://goreportcard.com/report/codeberg.org/fractalqb/groph)
[![GoDoc](https://godoc.org/git.fractalqb.de/fractalqb/groph?status.svg)](https://godoc.org/git.fractalqb.de/fractalqb/groph)

`import "git.fractalqb.de/fractalqb/groph"`

---

A pure Go library of graphs and algorithms. More info at [godoc.org](https://godoc.org/git.fractalqb.de/fractalqb/groph).

## Graph Implementations
The following table shows the supported graph implementations along
with their relative access performance, i.e. read & write. Access
performance is the factor relative to the fastest implementation –
currently the typed access on adjacency matrix with `int32` edges.
Each implementation can be accessed through their `WGraph` interface
_(t/gen)_ or through their specifically typed interface _(t/typed)_.

| Implementation   | Weight Type       | Dir/Undir | t/typed | t/gen |
|------------------|:-----------------:|:---------:|--------:|------:|
| Adjacency matrix | bool (bitmap)     | D         | 1.48    | 3.16  |
| Adjacency matrix | bool              | D         | 1.06    | 1.29  |
| Adjacency matrix | int32             | D, U      | 1       | 8.05  |
| Adjacency matrix | float32           | D, U      | 1.09    | 8.14  |
| Sparse / Go map  | interface\{\}     | D         | –       | 48.26 |
| Sparse / Go map  | int32             | D         |         |       |
| Sparse / Go map  | float32           | D         | 33.83   |       |

## Algorithms

| Problem | Algorithm |
|---------|-----------|
| [Traversal](https://godoc.org/codeberg.org/fractalqb/groph/trv)| Depth First |
| [Shortest path](https://godoc.org/codeberg.org/fractalqb/groph/sp)| Floyd Warshall |
| [Shortest path](https://godoc.org/codeberg.org/fractalqb/groph/sp)| Dijkstra |
| [Minimal spannig Tree](https://godoc.org/codeberg.org/fractalqb/groph/sp) | Dijkstra |
| [Minimal spannig Tree](https://godoc.org/codeberg.org/fractalqb/groph/mst) | Kruskal |
| [Travelling Salesman](https://godoc.org/codeberg.org/fractalqb/groph/tsp) | greedy |
| [Travelling Salesman](https://godoc.org/codeberg.org/fractalqb/groph/tsp) | 2-Opt |

## Performance compared to other Libraries

Comparison benchmarks can be found in the separate project [groph-cmpbench](https://codeberg.org/fractalqb/groph-cmpbench).

### Access Performance

Access performance is measured by writing and reading graph edges of a graph with a fixed number of vertices. While this does not tell anything about the quality of the provided algorithms theses operations are frequently called in inner loops of most algorithms. I.e. access performance will make factor of algorithm's performance.

Other graph implementations are run against the groph top-speed graph. Use the numbers from the internal benchmark to estimate other comparisons.

_Feedback on the benchmark project is very welcome to improve the validity of the comparison!_

| Library | Test | t-Factor |
|---------|------|-------:|
| groph   | access directed int32 | 1 |
| [yourbasic graph](https://github.com/yourbasic/graph) | directed int64 | 44 |
| [goraph](https://github.com/gyuho/goraph) | directed float64 | 737 |
| [thcyron graphs](https://github.com/thcyron/graphs) | directed float64 | 1145 |

_t-Factor: smaller is better_