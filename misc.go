package groph

import (
	"math"
)

const I32Del = -2147483648 // min{ int32 }

var nan32 = float32(math.NaN())

func NaN32() float32 { return nan32 }

func IsNaN32(x float32) bool { return math.IsNaN(float64(x)) }

func errState(v interface{}) error {
	if es, ok := v.(interface{ ErrState() error }); ok {
		return es.ErrState()
	}
	return nil
}
