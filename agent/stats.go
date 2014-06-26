package agent

import (
	"fmt"

	"github.com/meatballhat/go-newrelic/metric"
)

var (
	errInvalidValue = fmt.Errorf("invalid value")
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

// Record handles float64, named apdex, or *Stats instances, recording them internally
func (st *Stats) Record(value interface{}, aux interface{}) error {
	switch value.(type) {
	case float64:
		st.recordDataPoint(value.(float64), aux.(float64))
	case string:
		if value == "apdex_s" || value == "apdex_t" || value == "apdex_f" {
			st.recordApdex(value.(string), aux.(float64))
		}
	case *Stats:
		st.merge(value.(*Stats))
	default:
		return errInvalidValue
	}

	return nil
}

func (st *Stats) recordDataPoint(value float64, exclusiveTime float64) {
	st.CallCount++
	st.TotalCallTime += value
	if value < st.MinCallTime || st.CallCount == 1 {
		st.MinCallTime = value
	}
	if value > st.MaxCallTime {
		st.MaxCallTime = value
	}
	st.TotalExclusiveTime += exclusiveTime
	st.SumOfSquares += (value * value)
}

func (st *Stats) recordApdex(bucket string, apdexT float64) {
	switch bucket {
	case "apdex_s":
		st.CallCount++
	case "apdex_t":
		st.TotalCallTime++
	case "apdex_f":
		st.TotalExclusiveTime++
	}
	if apdexT != 0 {
		st.MinCallTime = apdexT
		st.MaxCallTime = apdexT
	}
}

func (st *Stats) merge(other *Stats) {
	if st.isMinTimeLess(other) {
		st.MinCallTime = other.MinCallTime
	}

	if other.MaxCallTime > st.MaxCallTime {
		st.MaxCallTime = other.MaxCallTime
	}

	st.SumOfSquares += other.SumOfSquares
	st.CallCount += other.CallCount
	st.TotalCallTime += other.TotalCallTime
	st.TotalExclusiveTime += other.TotalExclusiveTime
}

func (st *Stats) isMinTimeLess(other *Stats) bool {
	if (other.MinCallTime < st.MinCallTime && other.CallCount > 0) || st.CallCount == 0 {
		return true
	}
	return false
}
