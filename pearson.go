package correlation

import "math"

func Pearson(x, y []float64) float64 {
	xh, yh := Mean(x), Mean(y)       // Xhat, Yhat
	xc, yc := Sub(x, xh), Sub(y, yh) // X - Xhat, Y - Yhat
	nom := Sum(Mul(xc, yc))          // sum((X - Xhat)*(Y - Yhat))

	// sum((X - Xhat)^2) * sum((Y - Yhat)^2)
	var sx2, sy2 float64
	for i := range x {
		sx2 += xc[i] * xc[i]
		sy2 += yc[i] * yc[i]
	}
	denom := sx2 * sy2

	if isClose(denom, 0) {
		return 0
	}

	return nom / math.Sqrt(denom)
}
