package mon

import (
	"context"
)

type metricHook struct {
	writer      MetricWriter
	application string
}

func NewMetricHook() *metricHook {
	defaults := getDefaultMetrics()
	writer := NewMetricDaemonWriter(defaults...)

	return &metricHook{
		writer: writer,
	}
}

func (h metricHook) Fire(level string, msg string, err error, fields Fields, contextFields ContextFields, tags Tags, configValues ConfigValues, context context.Context, ecsMetadata EcsMetadata) error {
	if level != Warn && level != Error {
		return nil
	}

	h.writer.WriteOne(&MetricDatum{
		Priority:   PriorityHigh,
		MetricName: level,
		Unit:       UnitCount,
		Value:      1.0,
	})

	return nil
}

func getDefaultMetrics() MetricData {
	return MetricData{
		{
			Priority:   PriorityHigh,
			MetricName: Warn,
			Unit:       UnitCount,
			Value:      0.0,
		},
		{
			Priority:   PriorityHigh,
			MetricName: Error,
			Unit:       UnitCount,
			Value:      0.0,
		},
	}
}
