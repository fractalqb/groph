package groph

import (
	"testing"
)

func TestBitset(t *testing.T) {
	s := make([]uint, 3)
	if BitSetCap(s) != 3*wordBits {
		t.Fatalf("got bitset size %d, expected %d", BitSetCap(s), 3*wordBits)
	}
	for i := uint(0); i < BitSetCap(s); i++ {
		if BitSetGet(s, i) {
			t.Errorf("initial state of bit %d is true", i)
		}
		BitSetSet(s, i)
		for j := uint(0); j < BitSetCap(s); j++ {
			if i == j {
				if !BitSetGet(s, i) {
					t.Fatalf("failed to set bit %d", i)
				}
			} else if BitSetGet(s, j) {
				t.Fatalf("setting bit %d also sets %d", i, j)
			}
		}
		BitSetUnset(s, i)
		for j := uint(0); j < BitSetCap(s); j++ {
			if i == j {
				if BitSetGet(s, i) {
					t.Fatalf("failed to unset bit %d", i)
				}
			} else if BitSetGet(s, j) {
				t.Fatalf("unsetting bit %d also sets %d", i, j)
			}
		}
	}
}
