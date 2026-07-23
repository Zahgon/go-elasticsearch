package types

type AnalysisConfig struct {
	BucketSpan Duration `json:"bucket_span,omitempty"`

	CategorizationAnalyzer CategorizationAnalyzer `json:"categorization_analyzer,omitempty"`

	CategorizationFieldName *string `json:"categorization_field_name,omitempty"`

	CategorizationFilters []string `json:"categorization_filters,omitempty"`

	Detectors []Detector `json:"detectors"`

	Influencers []string `json:"influencers,omitempty"`

	Latency Duration `json:"latency,omitempty"`

	ModelPruneWindow Duration `json:"model_prune_window,omitempty"`

	MultivariateByFields *bool `json:"multivariate_by_fields,omitempty"`

	PerPartitionCategorization *PerPartitionCategorization `json:"per_partition_categorization,omitempty"`

	SummaryCountFieldName *string `json:"summary_count_field_name,omitempty"`
}

func (s *AnalysisConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAnalysisConfig() *AnalysisConfig { _ = "STUB: not implemented"; return nil }

type AnalysisConfigVariant interface {
	AnalysisConfigCaster() *AnalysisConfig
}

func (s *AnalysisConfig) AnalysisConfigCaster() *AnalysisConfig {
	_ = "STUB: not implemented"
	return nil
}
