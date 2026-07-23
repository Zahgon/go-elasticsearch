package types

type DataframeAnalyticsStatsProgress struct {
	Phase string `json:"phase"`

	ProgressPercent int `json:"progress_percent"`
}

func (s *DataframeAnalyticsStatsProgress) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsStatsProgress() *DataframeAnalyticsStatsProgress {
	_ = "STUB: not implemented"
	return nil
}
