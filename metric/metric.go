package metric

// Metric encapsulates metric stuff!
type Metric struct {
	Count        uint64
	Total        float64
	Min          float64
	Max          float64
	SumOfSquares float64
}
