package util

import (
	"fmt"
	"math/rand"
	"testing"

	"git.fractalqb.de/fractalqb/groph/sliceofmaps"

	"git.fractalqb.de/fractalqb/groph/adjmatrix"

	"git.fractalqb.de/fractalqb/groph"
)

func tstGSame(g1, g2 groph.RGraph, sameOrd bool, wEq func(w1, w2 interface{}) bool) error {
	ord := g1.Order()
	if sameOrd && g2.Order() != ord {
		return fmt.Errorf("graphs order differs: g1=%d, g2=%d", ord, g2.Order())
	}
	if g2.Order() < ord {
		ord = g2.Order()
	}
	for i := 0; i < ord; i++ {
		for j := 0; j < ord; j++ {
			w1, w2 := g1.Weight(i, j), g2.Weight(i, j)
			if !wEq(w1, w2) {
				return fmt.Errorf("differ at (%d,%d): w1=%v, w2=%v", i, j, w1, w2)
			}
		}
	}
	return nil
}

func TestCpWeights_from_directed(t *testing.T) {
	src := adjmatrix.NewDInt32(11, adjmatrix.I32Del, nil)
	for i := 0; i < src.Order(); i++ {
		for j := 0; j < src.Order(); j++ {
			if rand.Intn(100) < 75 {
				src.SetEdge(i, j, rand.Int31())
			}
		}
	}
	wEq := func(w1, w2 interface{}) bool { return w1 == w2 }
	t.Run("to directed", func(t *testing.T) {
		dst := sliceofmaps.NewDInt32(src.Order(), nil)
		_, err := CpWeights(dst, src)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame(src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
	t.Run("to undirected", func(t *testing.T) {
		dst := sliceofmaps.NewUInt32(src.Order(), nil)
		_, err := CpWeights(dst, src)
		if err == nil {
			t.Fatal("copy from directed to undirected did not return error")
		}
		if err.Error() != "cannot copy from directed to undirected graph" {
			t.Error("wrong error:", err)
		}
	})
}

func TestCpWeights_from_undirected(t *testing.T) {
	src := adjmatrix.NewUInt32(11, adjmatrix.I32Del, nil)
	for i := 0; i < src.Order(); i++ {
		for j := 0; j <= i; j++ {
			if rand.Intn(100) < 75 {
				src.SetEdge(i, j, rand.Int31())
			}
		}
	}
	wEq := func(w1, w2 interface{}) bool { return w1 == w2 }
	t.Run("to directed", func(t *testing.T) {
		dst := sliceofmaps.NewDInt32(src.Order(), nil)
		_, err := CpWeights(dst, src)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame(src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
	t.Run("to undirected", func(t *testing.T) {
		dst := sliceofmaps.NewUInt32(src.Order(), nil)
		_, err := CpWeights(dst, src)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame(src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
}

func TestCpXWeights_from_directed(t *testing.T) {
	src := adjmatrix.NewDInt32(11, adjmatrix.I32Del, nil)
	for i := 0; i < src.Order(); i++ {
		for j := 0; j < src.Order(); j++ {
			if rand.Intn(100) < 75 {
				src.SetEdge(i, j, rand.Int31())
			}
		}
	}
	xFn := func(w interface{}) interface{} { return float32(w.(int32)) }
	wEq := func(w1, w2 interface{}) bool {
		f1, f2 := float32(w1.(int32)), w2.(float32)
		return f1 == f2
	}
	t.Run("to directed", func(t *testing.T) {
		dst := sliceofmaps.NewSoMDf32(src.Order(), nil)
		_, err := CpXWeights(dst, src, xFn)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame(src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
	t.Run("to undirected", func(t *testing.T) {
		dst := sliceofmaps.NewSoMUf32(src.Order(), nil)
		_, err := CpXWeights(dst, src, xFn)
		if err == nil {
			t.Fatal("copy from directed to undirected did not return error")
		}
		if err.Error() != "cannot copy from directed to undirected graph" {
			t.Error("wrong error:", err)
		}
	})
}

func TestCpXWeights_from_undirected(t *testing.T) {
	src := adjmatrix.NewUInt32(11, adjmatrix.I32Del, nil)
	for i := 0; i < src.Order(); i++ {
		for j := 0; j <= i; j++ {
			if rand.Intn(100) < 75 {
				src.SetEdge(i, j, rand.Int31())
			}
		}
	}
	xFn := func(w interface{}) interface{} { return float32(w.(int32)) }
	wEq := func(w1, w2 interface{}) bool {
		f1, f2 := float32(w1.(int32)), w2.(float32)
		return f1 == f2
	}
	t.Run("to directed", func(t *testing.T) {
		dst := sliceofmaps.NewSoMDf32(src.Order(), nil)
		_, err := CpXWeights(dst, src, xFn)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame(src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
	t.Run("to undirected", func(t *testing.T) {
		dst := sliceofmaps.NewSoMUf32(src.Order(), nil)
		_, err := CpXWeights(dst, src, xFn)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}
		if err = tstGSame(src, dst, true, wEq); err != nil {
			t.Error(err)
		}
	})
}
