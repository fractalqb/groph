package groph

import (
	"testing"
)

func TestBitset(t *testing.T) {
	s := make(bitSet, 3)
	if s.cap() != 3*wordBits {
		t.Fatalf("got bitset size %d, expected %d", s.cap(), 3*wordBits)
	}
	for i := 0; i < s.cap(); i++ {
		if s.get(i) {
			t.Errorf("initial state of bit %d is true", i)
		}
		s.set(i)
		for j := 0; j < s.cap(); j++ {
			if i == j {
				if !s.get(i) {
					t.Fatalf("failed to set bit %d", i)
				}
			} else if s.get(j) {
				t.Fatalf("setting bit %d also sets %d", i, j)
			}
		}
		s.unset(i)
		for j := 0; j < s.cap(); j++ {
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

func TestBitset_firstUnset(t *testing.T) {
	s := make(bitSet, 3)
	for i := range s {
		s[i] = wordAll
	}
	if s1 := s.firstUnset(); s1 >= 0 {
		t.Errorf("found unset at %d in complete set", s1)
	}
	exp := len(s)*wordBits - 1
	s.unset(exp)
	if s1 := s.firstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = (len(s) - 1) * wordBits
	s.unset(exp)
	if s1 := s.firstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = (len(s)-1)*wordBits - 1
	s.unset(exp)
	if s1 := s.firstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = wordBits + (wordBits >> 1)
	s.unset(exp)
	if s1 := s.firstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = wordBits
	s.unset(exp)
	if s1 := s.firstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = wordBits >> 1
	s.unset(exp)
	if s1 := s.firstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = 0
	s.unset(exp)
	if s1 := s.firstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
}
