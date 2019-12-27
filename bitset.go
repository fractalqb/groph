package groph

// http://graphics.stanford.edu/~seander/bithacks.html

type bitsWord = uint64

const (
	wordBits  = 64
	wordMask  = 0x3f
	wordShift = 6
)

func bitSetWords(size VIdx) VIdx {
	return (size + (wordBits - 1)) / wordBits
}

func bitSetGet(bs []bitsWord, i uint) bool {
	w, b := i>>wordShift, i&wordMask
	return bs[w]&(1<<b) != 0
}

func bitSetSet(bs []bitsWord, i uint) {
	w, b := i>>wordShift, i&wordMask
	bs[w] |= 1 << b
}

func bitSetUnset(bs []bitsWord, i uint) {
	w, b := i>>wordShift, i&wordMask
	bs[w] &= ^(1 << b)
}
