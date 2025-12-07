package correlation_test

import (
	"fmt"

	"github.com/itsubaki/correlation"
)

func ExampleSpearman() {
	x := []float64{1, 2, 3, 4, 5} // rx = [1,2,3,4,5]
	y := []float64{5, 6, 7, 8, 7} // ry = [1, 2, 3.5, 5, 3.5]

	r := correlation.Spearman(x, y)
	fmt.Printf("%.4f\n", r)

	// Output: 0.8250
}
