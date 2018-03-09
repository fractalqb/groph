package groph

import (
	"reflect"
)

// http://graphics.stanford.edu/~seander/bithacks.html

var wordBits uint
var wordMask uint

func init() {
	wordBits = uint(reflect.TypeOf(uint(0)).Bits())
	wordMask = wordBits - 1
}

func BitSetCap(bs []uint) uint { return uint(len(bs)) * wordBits }

func BitSetGet(bs []uint, i uint) bool {
	w, b := i/wordBits, i&wordMask
	return bs[w]&(1<<b) != 0
}

func BitSetSet(bs []uint, i uint) {
	w, b := i/wordBits, i&wordMask
	bs[w] |= 1 << b
}

func BitSetUnset(bs []uint, i uint) {
	w, b := i/wordBits, i&wordMask
	bs[w] &= ^(1 << b)
}

type BitSet struct {
	Raw  []uint
	Size uint
}
