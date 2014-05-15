package agent

import (
	"github.com/meatballhat/go-newrelic/metric"
)

// Stats contains bits that are sent through the stats engine
type Stats struct {
	CallCount          uint64
	TotalCallTime      float64
	TotalExclusiveTime float64
	MinCallTime        float64
	MaxCallTime        float64
	SumOfSquares       float64
}

// NewStatsFromMetric generates *Stats from *Metric
func NewStatsFromMetric(m *metric.Metric) (*Stats, error) {
	return nil, nil
}
