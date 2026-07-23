package types

type ExponentialAverageCalculationContext struct {
	IncrementalMetricValueMs     Float64 `json:"incremental_metric_value_ms"`
	LatestTimestamp              *int64  `json:"latest_timestamp,omitempty"`
	PreviousExponentialAverageMs Float64 `json:"previous_exponential_average_ms,omitempty"`
}

func (s *ExponentialAverageCalculationContext) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExponentialAverageCalculationContext() *ExponentialAverageCalculationContext {
	_ = "STUB: not implemented"
	return nil
}
