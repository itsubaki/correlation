package correlation

var (
	_ Correlation = Chatterjee
	_ Correlation = Pearson
	_ Correlation = Spearman
)

type Correlation func(x, y []float64) float64
