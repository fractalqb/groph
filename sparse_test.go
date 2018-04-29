package groph

import (
	"testing"
)

func TestSpMap_SetUset(t *testing.T) {
	m := NewSpMap(7, nil)
	testSetUnset(m, 4, t)
}

func TestSpMapf32_SetUset(t *testing.T) {
	m := NewSpMap(7, nil)
	testSetUnset(m, 4, t)
}
