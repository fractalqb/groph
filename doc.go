// Copyright 2022 Marcus Perlick
// This file is part of Go module git.fractalqb.de/fractalqb/groph
//
// groph is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// groph is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with groph.  If not, see <http://www.gnu.org/licenses/>.

// Package groph is a pure Go library for graph data and algorithms.
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
// an indicator for the edge weights, e.g. “i32” for int32 or
// “f32” for float32 or “bool” for just bool.
package groph
