package util

type bitsWord = uint64

const (
	wordBits  = 64
	wordMask  = 0x3f
	wordShift = 6
)

type BitSet []bitsWord

func BitSetWords(size int) int {
	return (size + (wordBits - 1)) / wordBits
}

func BitSetCap(bs []bitsWord) int { return len(bs) * wordBits }

func BitSetGet(bs BitSet, i int) bool {
	w, b := i>>wordShift, i&wordMask
	return bs[w]&(1<<b) != 0
}

func BitSetSet(bs BitSet, i int) {
	w, b := i>>wordShift, i&wordMask
	bs[w] |= 1 << b
}

func BitSetUnset(bs BitSet, i int) {
	w, b := i>>wordShift, i&wordMask
	bs[w] &= ^(1 << b)
}
