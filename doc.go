// Package groph is yet another graph library for go.
//
// It provides an abstract graph model through some interfaces
// described below, a set of different implementations for those graph
// models—there is no one-size-fits-all implementation—and graph
// algorithms provided in sub packages.
//
// Vertices are always represented by an integer value of type VIdx
// ranging from 0 to n-1 for graphs with n vertices. Mapping different
// vertex types to the VIdx values for groph is currently left to the
// user. Edges are associated with values called “weights” which is a
// useful feature for most graph algorithms. A simple, unweighted
// graph can be considered to have weights of type bool that represent
// whether an edged is in the graph.
//
// The graph interfaces
//
// The graph interfaces are grouped by some specific aspects:
//
// - Graph implementations can be read-only and read-write.
//
// - Graph implementations are for directed or undirected graphs.
//
// - Values associated with edges can be general, i.e. of type
// interface{}, or have some specific type.
//
// The most basic is the RGraph interface that allows read-only access
// to directed and undirected graphs with no specific edge type. The
// WGraph interface extends RGraph with methods to change edges in the
// graph.
//
// The variants of RGraph and WGraph for undirected graphs are
// RUndirected and WUndirected. Any specific graph that does not
// implement RUndirected is considered to be a directed graph. One can
// use the Directed() function to just distinguish directed and
// undirected at runtime. To also downcast to undirected graphs use
// standard Go features.
//
// More specific interfaces are derived from these interfaces and
// adhere to the following naming convention. The first letter is
// either 'R' or 'W' to indicate read-only respectively read-write
// graphs. The second letter is either 'G' or 'U' to distinguish
// general graph interfaces from undirected interfaces. Finally comes
// in indicator for the edge weights, e.g. “bool”, “i32” for int32 or
// “f32” for float32.
package groph
