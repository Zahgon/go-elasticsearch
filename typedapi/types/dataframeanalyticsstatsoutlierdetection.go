package types

type DataframeAnalyticsStatsOutlierDetection struct {
	Parameters OutlierDetectionParameters `json:"parameters"`

	Timestamp int64 `json:"timestamp"`

	TimingStats TimingStats `json:"timing_stats"`
}

func (s *DataframeAnalyticsStatsOutlierDetection) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsStatsOutlierDetection() *DataframeAnalyticsStatsOutlierDetection {
	_ = "STUB: not implemented"
	return nil
}
