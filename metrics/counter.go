package metrics

// counter defines the counter. counter is report to each external Sink-able system.
type counter struct {
	name string
}

// Incr increases counter by one.
func (c *counter) Incr() {
	c.IncrBy(1)
}

// IncrBy increases counter by v and reports for each external Sink-able systems.
func (c *counter) IncrBy(v float64) {
	rec := NewSingleDimensionMetrics(c.name, v, PolicySUM)
	for _, sink := range metricsSinks {
		sink.Report(rec)
	}
}
