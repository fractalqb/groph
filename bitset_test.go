package groph

import (
	"testing"
)

func bitSetCap(bs []bitsWord) uint { return uint(len(bs)) * wordBits }

func TestBitset(t *testing.T) {
	s := make([]uint64, 3)
	if bitSetCap(s) != 3*wordBits {
		t.Fatalf("got bitset size %d, expected %d", bitSetCap(s), 3*wordBits)
	}
	for i := uint(0); i < bitSetCap(s); i++ {
		if bitSetGet(s, i) {
			t.Errorf("initial state of bit %d is true", i)
		}
		bitSetSet(s, i)
		for j := uint(0); j < bitSetCap(s); j++ {
			if i == j {
				if !bitSetGet(s, i) {
					t.Fatalf("failed to set bit %d", i)
				}
			} else if bitSetGet(s, j) {
				t.Fatalf("setting bit %d also sets %d", i, j)
			}
		}
		bitSetUnset(s, i)
		for j := uint(0); j < bitSetCap(s); j++ {
			if i == j {
				if bitSetGet(s, i) {
					t.Fatalf("failed to unset bit %d", i)
				}
			} else if bitSetGet(s, j) {
				t.Fatalf("unsetting bit %d also sets %d", i, j)
			}
		}
	}
}
