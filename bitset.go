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

type bitset []uint

func (bs bitset) size() uint { return uint(len(bs)) * wordBits }

func (bs bitset) get(i uint) bool {
	w, b := i/wordBits, i&wordMask
	return bs[w]&(1<<b) != 0
}

func (bs bitset) set(i uint) {
	w, b := i/wordBits, i&wordMask
	bs[w] |= 1 << b
}

func (bs bitset) unset(i uint) {
	w, b := i/wordBits, i&wordMask
	bs[w] &= ^(1 << b)
}
