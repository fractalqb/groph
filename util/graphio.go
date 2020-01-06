package util

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"git.fractalqb.de/fractalqb/groph"
)

func WriteSparse(wr io.Writer, g groph.RGraph) error {
	fmt.Fprint(wr, "groph ")
	if groph.Directed(g) {
		fmt.Fprint(wr, "directed")
	} else {
		fmt.Fprint(wr, "undirected")
	}
	fmt.Fprintf(wr, " edges order=%d\n", g.Order())
	groph.EachEdge(g, func(u, v groph.VIdx) bool {
		if w := g.Weight(u, v); w != nil {
			fmt.Fprintf(wr, "%d %d %v\n", u, v, w)
		}
		return false
	})
	return nil // TODO detect errors
}

func readOdrer(prop string) (order int) {
	if !strings.HasPrefix(prop, "order=") {
		return -1
	}
	order, err := strconv.Atoi(prop[6:])
	if err != nil {
		return -1
	}
	return order
}

type WeightParse = func(string) (interface{}, error)

func ParseI32(s string) (interface{}, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	return int32(i), err
}

func ReadGraph(into groph.WGraph, rd io.Reader, wParse WeightParse) error {
	scn := bufio.NewScanner(rd)
	if !scn.Scan() {
		return errors.New("empty input")
	}
	hdr := strings.Split(scn.Text(), " ")
	if len(hdr) != 4 {
		return fmt.Errorf("synax error in graph header: has %d fields, not 4", len(hdr))
	}
	if hdr[0] != "groph" {
		return errors.New("no groph file indicator")
	}
	if hdr[1] == "directed" && !groph.Directed(into) {
		return fmt.Errorf(
			"cannot read directed graph data into undirected graph implementaton")
	}
	ord := readOdrer(hdr[3])
	if ord < 0 {
		return errors.New("cannot read odrer of graph")
	}
	into.Reset(ord)
	switch hdr[2] {
	case "edges":
		return readEdges(into, scn, wParse)
	default:
		return fmt.Errorf("unknown graph format: '%s'", hdr[2])
	}
	return nil
}

func readEdges(g groph.WGraph, scn *bufio.Scanner, wParse WeightParse) error {
	count := 0
	for scn.Scan() {
		count++
		fields := strings.Split(scn.Text(), " ")
		if len(fields) != 3 {
			return fmt.Errorf("syntax error in edge %d: has %d fields, not 3",
				count,
				len(fields))
		}
		u, err := strconv.Atoi(strings.TrimSpace(fields[0]))
		if err != nil {
			return fmt.Errorf("syntax error in edge %d: %s", count, err)
		}
		v, err := strconv.Atoi(strings.TrimSpace(fields[1]))
		if err != nil {
			return fmt.Errorf("syntax error in edge %d: %s", count, err)
		}
		var w interface{} = fields[2]
		if wParse != nil {
			w, err = wParse(fields[2])
		}
		g.SetWeight(u, v, w)
	}
	return nil
}
