package util

import (
	"testing"
)

func TestBitset(t *testing.T) {
	s := make(BitSet, 3)
	if s.Cap() != 3*wordBits {
		t.Fatalf("got bitset size %d, expected %d", s.Cap(), 3*wordBits)
	}
	for i := 0; i < s.Cap(); i++ {
		if s.Get(i) {
			t.Errorf("initial state of bit %d is true", i)
		}
		s.Set(i)
		for j := 0; j < s.Cap(); j++ {
			if i == j {
				if !s.Get(i) {
					t.Fatalf("failed to set bit %d", i)
				}
			} else if s.Get(j) {
				t.Fatalf("setting bit %d also sets %d", i, j)
			}
		}
		s.Unset(i)
		for j := 0; j < s.Cap(); j++ {
			if i == j {
				if s.Get(i) {
					t.Fatalf("failed to unset bit %d", i)
				}
			} else if s.Get(j) {
				t.Fatalf("unsetting bit %d also sets %d", i, j)
			}
		}
	}
}
