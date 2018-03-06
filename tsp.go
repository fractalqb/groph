package graph

type Solver func(d Measure, points []interface{}) (path []uint, pLen float64)

type MatrixSolverf32 func(size uint, m *AdjMxAf32) (path []uint, pLen float64)

func AsymSolverf32(solver MatrixSolverf32) Solver {
	return func(d Measure, points []interface{}) ([]uint, float64) {
		m := NewAdjMxAf32(uint(len(points)), nil)
		SetMetric(m, func(a, b Vertex) interface{} { return d(a, b) })
		return solver(uint(len(points)), m)
	}
}
