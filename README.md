# groph

[![Test Coverage](https://img.shields.io/badge/coverage-58%25-orange.svg)](file:coverage.html)
[![Go Report Card](https://goreportcard.com/badge/codeberg.org/fractalqb/groph)](https://goreportcard.com/report/codeberg.org/fractalqb/groph)
[![GoDoc](https://godoc.org/codeberg.org/fractalqb/groph?status.svg)](https://godoc.org/codeberg.org/fractalqb/groph)

`import "git.fractalqb.de/fractalqb/groph"`

---

Yet another graph library. More info at [godoc.org](https://godoc.org/codeberg.org/fractalqb/groph).

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
| [Shortest path](https://godoc.org/codeberg.org/fractalqb/groph/sp)| Floyd Warshall |
| [Shortest path](https://godoc.org/codeberg.org/fractalqb/groph/sp)| Dijkstra |
| [Minimal spannig Tree](https://godoc.org/codeberg.org/fractalqb/groph/sp) | Dijkstra |
| [Minimal spannig Tree](https://godoc.org/codeberg.org/fractalqb/groph/mst) | Kruskal |
| [Travelling Salesman](https://godoc.org/codeberg.org/fractalqb/groph/tsp) | greedy |
| [Travelling Salesman](https://godoc.org/codeberg.org/fractalqb/groph/tsp) | 2-Opt |

## Performance compared to other Libraries

Comparison benchmarks can be found in the separate project [groph-cmpbench](https://codeberg.org/fractalqb/groph-cmpbench).

### Access Performance

| Library | Test | t-Factor |
|---------|------|-------:|
| groph   | access directed int32 | 1 |
| [yourbasic graph](https://github.com/yourbasic/graph) | access directed int64 | 43.56 |
_Factor: smaller is better_