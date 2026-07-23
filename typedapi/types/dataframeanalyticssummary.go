package types

type DataframeAnalyticsSummary struct {
	AllowLazyStart *bool                            `json:"allow_lazy_start,omitempty"`
	Analysis       DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields *DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`

	Authorization    *DataframeAnalyticsAuthorization `json:"authorization,omitempty"`
	CreateTime       *int64                           `json:"create_time,omitempty"`
	Description      *string                          `json:"description,omitempty"`
	Dest             DataframeAnalyticsDestination    `json:"dest"`
	Id               string                           `json:"id"`
	MaxNumThreads    *int                             `json:"max_num_threads,omitempty"`
	Meta_            Metadata                         `json:"_meta,omitempty"`
	ModelMemoryLimit *string                          `json:"model_memory_limit,omitempty"`
	Source           DataframeAnalyticsSource         `json:"source"`
	Version          *string                          `json:"version,omitempty"`
}

func (s *DataframeAnalyticsSummary) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsSummary() *DataframeAnalyticsSummary {
	_ = "STUB: not implemented"
	return nil
}
