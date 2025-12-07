package correlation_test

import (
	"fmt"

	"github.com/itsubaki/correlation"
)

func ExampleChatterjee() {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{5, 6, 7, 8, 7}

	r := correlation.Chatterjee(x, y)
	fmt.Printf("%.4f\n", r)

	// Output: 0.3125
}
