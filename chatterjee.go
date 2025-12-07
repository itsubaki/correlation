package correlation

import (
	"math"
	"sort"
)

func Chatterjee(x, y []float64) float64 {
	type pair struct {
		x float64
		y float64
	}

	// Sort by x
	n := len(x)
	tmp := make([]pair, n)
	for i := range n {
		tmp[i] = pair{x[i], y[i]}
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].x < tmp[j].x
	})

	// Extract sorted y values
	ySorted := make([]float64, n)
	for i := range n {
		ySorted[i] = tmp[i].y
	}

	// Rank the sorted y values
	r := rank(ySorted)

	// Compute the sum of absolute differences of consecutive ranks
	var sum float64
	for i := range n - 1 {
		sum += math.Abs(r[i+1] - r[i])
	}

	denom := float64(n*n) - 1
	if IsClose(denom, 0) {
		return 0
	}

	return 1 - (3*sum)/denom
}
