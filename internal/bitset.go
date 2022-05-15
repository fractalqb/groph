// Copyright 2022 Marcus Perlick
// This file is part of Go module git.fractalqb.de/fractalqb/groph
//
// groph is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// groph is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with groph.  If not, see <http://www.gnu.org/licenses/>.

package internal

type bitsWord = uint64

const (
	wordAll   = 0xffff_ffff_ffff_ffff
	wordBits  = 64
	wordMask  = 0x3f
	wordShift = 6
)

type BitSet []bitsWord

func bitSetWords(setSize int) int {
	return (setSize + (wordBits - 1)) / wordBits
}

func NewBitSet(setSize int, reuse BitSet) BitSet {
	bsw := bitSetWords(setSize)
	if reuse == nil || cap(reuse) < bsw {
		return make(BitSet, bitSetWords(setSize))
	}
	reuse = reuse[:bsw]
	reuse.Clear()
	return reuse
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
