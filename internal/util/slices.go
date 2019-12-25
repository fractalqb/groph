package util

var VIdxSlice = IntSlice

func BoolSlice(s []bool, l int) []bool {
	if s == nil || cap(s) < l {
		return make([]bool, l)
	}
	return s[:l]
}

func IntSlice(s []int, l int) []int {
	if s == nil || cap(s) < l {
		return make([]int, l)
	}
	return s[:l]
}

func I32Slice(s []int32, l int) []int32 {
	if s == nil || cap(s) < l {
		return make([]int32, l)
	}
	return s[:l]
}

func UIntSlice(s []uint, l int) []uint {
	if s == nil || cap(s) < l {
		return make([]uint, l)
	}
	return s[:l]
}

func F32Slice(s []float32, l int) []float32 {
	if s == nil || cap(s) < l {
		return make([]float32, l)
	}
	return s[:l]
}
