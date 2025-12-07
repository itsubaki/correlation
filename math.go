package correlation

import "math"

func Mean(x []float64) float64 {
	return Sum(x) / float64(len(x))
}

func Sum(x []float64) float64 {
	var sum float64
	for i := range x {
		sum += x[i]
	}

	return sum
}

func Mul(x, y []float64) []float64 {
	z := make([]float64, len(x))
	for i := range x {
		z[i] = x[i] * y[i]
	}

	return z
}

func Sub(x []float64, y float64) []float64 {
	z := make([]float64, len(x))
	for i := range x {
		z[i] = x[i] - y
	}

	return z
}

func isClose(a, b float64, tol ...float64) bool {
	atol, rtol := func() (float64, float64) {
		if len(tol) == 0 {
			return 1e-8, 1e-5
		}

		if len(tol) == 1 {
			return tol[0], tol[0]
		}

		return tol[0], tol[1]
	}()

	return math.Abs(a-b) <= atol+rtol*math.Max(math.Abs(a), math.Abs(b))
}
