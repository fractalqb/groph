package util

type bitsWord = uint64

const (
	wordBits  = 64
	wordMask  = 0x3f
	wordShift = 6
)

type BitSet []bitsWord

func BitSetWords(setSize int) int {
	return (setSize + (wordBits - 1)) / wordBits
}

func (bs BitSet) Cap() int { return len(bs) * wordBits }

func (bs BitSet) Get(i int) bool {
	w, b := i>>wordShift, i&wordMask
	return bs[w]&(1<<b) != 0
}

func (bs BitSet) Set(i int) {
	w, b := i>>wordShift, i&wordMask
	bs[w] |= 1 << b
}

func (bs BitSet) Unset(i int) {
	w, b := i>>wordShift, i&wordMask
	bs[w] &= ^(1 << b)
}
