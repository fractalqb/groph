package util

type bitsWord = uint64

const (
	wordAll = 0xffff_ffff_ffff_ffff
	wordBits  = 64
	wordMask  = 0x3f
	wordShift = 6
)

type BitSet []bitsWord

func BitSetWords(setSize int) int {
	return (setSize + (wordBits - 1)) / wordBits
}

func NewBitSet(setSize int) BitSet { return make(BitSet, BitSetWords(setSize))}

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

func (bs BitSet) Clear() {
	for i := range bs {
		bs[i] = 0
	}
}

func (bs BitSet) FirstUnset() (res int) {
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