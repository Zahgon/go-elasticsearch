package types

type DataframeAnalyticsStatsMemoryUsage struct {
	MemoryReestimateBytes *int64 `json:"memory_reestimate_bytes,omitempty"`

	PeakUsageBytes int64 `json:"peak_usage_bytes"`

	Status string `json:"status"`

	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (s *DataframeAnalyticsStatsMemoryUsage) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsStatsMemoryUsage() *DataframeAnalyticsStatsMemoryUsage {
	_ = "STUB: not implemented"
	return nil
}
