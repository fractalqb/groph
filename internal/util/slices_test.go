package util

import (
	"reflect"
	"testing"
)

func assertLenCap(t *testing.T, s interface{}, l, c int, hint string) {
	sval := reflect.ValueOf(s)
	if sval.Cap() != c {
		t.Errorf("%s: cap expected %d, got %d", hint, c, sval.Cap())
	}
	if sval.Len() != l {
		t.Errorf("%s: len expected %d, got %d", hint, l, sval.Len())
	}
}

func TestIntSlice(t *testing.T) {
	s := IntSlice(nil, 8)
	assertLenCap(t, s, 8, 8, "from nil")
	s = IntSlice(s, 5)
	assertLenCap(t, s, 5, 8, "from nil")
	s = IntSlice(s, 7)
	assertLenCap(t, s, 7, 8, "from nil")
	s = IntSlice(s, 12)
	assertLenCap(t, s, 12, 12, "from nil")
}
