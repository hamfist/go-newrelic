package agent

import (
	"github.com/meatballhat/go-newrelic/metric"
)

// Agent does the heavy lifting with the sending of the things
type Agent struct {
	StatsEngine *statsEngine
}

// NewAgent creates an *Agent given some *Options.  Wow!
func NewAgent(opts *Options) *Agent {
	return &Agent{}
}

// RecordMetric sends a metric!
func (a *Agent) RecordMetric(metricName string, value *metric.Metric) error {
	stats, err := NewStatsFromMetric(value)
	if err != nil {
		return err
	}

	return a.StatsEngine.RecordMetrics(metricName, stats)
}
