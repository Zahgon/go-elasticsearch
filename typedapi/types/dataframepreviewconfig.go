package types

type DataframePreviewConfig struct {
	Analysis         DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields   *DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	MaxNumThreads    *int                             `json:"max_num_threads,omitempty"`
	ModelMemoryLimit *string                          `json:"model_memory_limit,omitempty"`
	Source           DataframeAnalyticsSource         `json:"source"`
}

func (s *DataframePreviewConfig) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframePreviewConfig() *DataframePreviewConfig { _ = "STUB: not implemented"; return nil }

type DataframePreviewConfigVariant interface {
	DataframePreviewConfigCaster() *DataframePreviewConfig
}

func (s *DataframePreviewConfig) DataframePreviewConfigCaster() *DataframePreviewConfig {
	_ = "STUB: not implemented"
	return nil
}
