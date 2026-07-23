package types

type DataframeAnalyticsSource struct {
	Index []string `json:"index"`

	Query *Query `json:"query,omitempty"`

	RuntimeMappings RuntimeFields `json:"runtime_mappings,omitempty"`

	Source_ *DataframeAnalysisAnalyzedFields `json:"_source,omitempty"`
}

func (s *DataframeAnalyticsSource) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsSource() *DataframeAnalyticsSource { _ = "STUB: not implemented"; return nil }

type DataframeAnalyticsSourceVariant interface {
	DataframeAnalyticsSourceCaster() *DataframeAnalyticsSource
}

func (s *DataframeAnalyticsSource) DataframeAnalyticsSourceCaster() *DataframeAnalyticsSource {
	_ = "STUB: not implemented"
	return nil
}
