package types

type DataframeAnalyticsStatsDataCounts struct {
	SkippedDocsCount int `json:"skipped_docs_count"`

	TestDocsCount int `json:"test_docs_count"`

	TrainingDocsCount int `json:"training_docs_count"`
}

func (s *DataframeAnalyticsStatsDataCounts) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsStatsDataCounts() *DataframeAnalyticsStatsDataCounts {
	_ = "STUB: not implemented"
	return nil
}
