package agent

type statsEngine struct {
}

func (se *statsEngine) RecordMetrics(metricName string, value *Stats) error {
	return nil
}
