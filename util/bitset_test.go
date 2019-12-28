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

func TestBitset_firstUnset(t *testing.T) {
	s := make(BitSet, 3)
	for i := range s {
		s[i] = wordAll
	}
	if s1 := s.FirstUnset(); s1 >= 0 {
		t.Errorf("found unset at %d in complete set", s1)
	}
	exp := len(s)*wordBits - 1
	s.Unset(exp)
	if s1 := s.FirstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = (len(s) - 1) * wordBits
	s.Unset(exp)
	if s1 := s.FirstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = (len(s)-1)*wordBits - 1
	s.Unset(exp)
	if s1 := s.FirstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = wordBits + (wordBits >> 1)
	s.Unset(exp)
	if s1 := s.FirstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = wordBits
	s.Unset(exp)
	if s1 := s.FirstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = wordBits >> 1
	s.Unset(exp)
	if s1 := s.FirstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
	exp = 0
	s.Unset(exp)
	if s1 := s.FirstUnset(); s1 != exp {
		t.Errorf("found unset at %d instead of %d", s1, exp)
	}
}
