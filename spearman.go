package correlation

import "sort"

func Spearman(x, y []float64) float64 {
	n := len(x)
	rx, ry := rank(x), rank(y)

	var d2 float64
	for i := range n {
		d := rx[i] - ry[i]
		d2 += d * d
	}

	denom := float64(n*n*n) - float64(n) // n^3 - n
	if isClose(denom, 0) {
		return 0
	}

	return 1 - (6*d2)/denom
}

func rank(data []float64) []float64 {
	n := len(data)
	type pair struct {
		value float64
		index int
	}

	tmp := make([]pair, n)
	for i := range n {
		tmp[i] = pair{data[i], i}
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].value < tmp[j].value
	})

	ranks := make([]float64, n)
	var i int
	for i < n {
		j := i + 1
		for j < n && tmp[j].value == tmp[i].value {
			j++
		}

		avg := 0.5 * (float64(i+1) + float64(j))
		for k := i; k < j; k++ {
			ranks[tmp[k].index] = avg
		}

		i = j
	}

	return ranks
}
