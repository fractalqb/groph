package groph

import (
	"testing"
)

func TestNaN32(t *testing.T) {
	if !IsNaN32(NaN32()) {
		t.Errorf("cannot detect NaN of type float32")
	}
}
