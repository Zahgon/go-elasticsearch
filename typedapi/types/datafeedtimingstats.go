package types

type DatafeedTimingStats struct {
	AverageSearchTimePerBucketMs Float64 `json:"average_search_time_per_bucket_ms,omitempty"`

	BucketCount                          int64                                 `json:"bucket_count"`
	ExponentialAverageCalculationContext *ExponentialAverageCalculationContext `json:"exponential_average_calculation_context,omitempty"`

	ExponentialAverageSearchTimePerHourMs Float64 `json:"exponential_average_search_time_per_hour_ms"`

	JobId string `json:"job_id"`

	SearchCount int64 `json:"search_count"`

	TotalSearchTimeMs Float64 `json:"total_search_time_ms"`
}

func (s *DatafeedTimingStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDatafeedTimingStats() *DatafeedTimingStats { _ = "STUB: not implemented"; return nil }
