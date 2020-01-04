package groph

type bitsWord = uint64

const (
	wordAll = 0xffff_ffff_ffff_ffff
	wordBits  = 64
	wordMask  = 0x3f
	wordShift = 6
)

type bitSet []bitsWord

func bitSetWords(setSize int) int {
	return (setSize + (wordBits - 1)) / wordBits
}

func newBitSet(setSize int) bitSet { return make(bitSet, bitSetWords(setSize))}

func (bs bitSet) cap() int { return len(bs) * wordBits }

func (bs bitSet) get(i int) bool {
	w, b := i>>wordShift, i&wordMask
	return bs[w]&(1<<b) != 0
}

func (bs bitSet) set(i int) {
	w, b := i>>wordShift, i&wordMask
	bs[w] |= 1 << b
}

func (bs bitSet) unset(i int) {
	w, b := i>>wordShift, i&wordMask
	bs[w] &= ^(1 << b)
}

func (bs bitSet) clear() {
	for i := range bs {
		bs[i] = 0
	}
}

func (bs bitSet) firstUnset() (res int) {
	var w bitsWord
	for res, w = range bs {
		if w != wordAll {
			res *= wordBits			
			for b := bitsWord(1); w&b == b; b <<= 1 {
				res++
			}
			return res
		}
	}
	return -1
}