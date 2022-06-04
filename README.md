# groph

[![Build Status](https://travis-ci.org/fractalqb/groph.svg)](https://travis-ci.org/fractalqb/groph)
[![codecov](https://codecov.io/gh/fractalqb/groph/branch/master/graph/badge.svg)](https://codecov.io/gh/fractalqb/groph)
[![Go Report Card](https://goreportcard.com/badge/git.fractalqb.de/fractalqb/groph)](https://goreportcard.com/report/git.fractalqb.de/fractalqb/groph)
[![GoDoc](https://godoc.org/git.fractalqb.de/fractalqb/groph?status.svg)](https://pkg.go.dev/git.fractalqb.de/fractalqb/groph)

`import "git.fractalqb.de/fractalqb/groph"`

---

A pure Go library of graphs and their algorithms.

*The library is currently rewritten for Go generics.*

## Graph implementations:

- [**Adjacency matrix**](https://en.wikipedia.org/wiki/Adjacency_matrix): Uses a continuous region of memory, i.e. a slice
- [**Adjacency list**](https://en.wikipedia.org/wiki/Adjacency_list): Uses a slice of slices
- **Edgelist**: A slice of {u, v, w}
- **Euclidean**: Computes the euclidean distance as weight for each edge where vertices have to implement the Distancer interface
- **Forest**: A compact representation for trees and forests

## Algorithms

### Computing Paths
- Floyd Warshall for shortest paths
- Dijkstra's Algorithm for shortest paths
- A greedy implementation for the TSP
- A 2-opt based implementation for the TSP
- A* algorithm for undirected graphs
