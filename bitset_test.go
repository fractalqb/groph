package groph

import (
	"testing"
)

func TestBitset(t *testing.T) {
	s := bitset(make([]uint, 3))
	if s.size() != 3*wordBits {
		t.Fatalf("got bitset size %d, expected %d", s.size(), 3*wordBits)
	}
	for i := uint(0); i < s.size(); i++ {
		if s.get(i) {
			t.Errorf("initial state of bit %d is true", i)
		}
		s.set(i)
		for j := uint(0); j < s.size(); j++ {
			if i == j {
				if !s.get(i) {
					t.Fatalf("failed to set bit %d", i)
				}
			} else if s.get(j) {
				t.Fatalf("setting bit %d also sets %d", i, j)
			}
		}
		s.unset(i)
		for j := uint(0); j < s.size(); j++ {
			if i == j {
				if s.get(i) {
					t.Fatalf("failed to unset bit %d", i)
				}
			} else if s.get(j) {
				t.Fatalf("unsetting bit %d also sets %d", i, j)
			}
		}
	}
}
