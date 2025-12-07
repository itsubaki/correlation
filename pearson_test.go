package correlation_test

import (
	"fmt"

	"github.com/itsubaki/correlation"
)

func ExamplePearson() {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{5, 6, 7, 8, 7}

	r := correlation.Pearson(x, y)
	fmt.Printf("%.4f\n", r)

	// Output: 0.8321
}
