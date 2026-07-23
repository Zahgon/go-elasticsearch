package types

type AnalysisConfigRead struct {
	BucketSpan Duration `json:"bucket_span"`

	CategorizationAnalyzer CategorizationAnalyzer `json:"categorization_analyzer,omitempty"`

	CategorizationFieldName *string `json:"categorization_field_name,omitempty"`

	CategorizationFilters []string `json:"categorization_filters,omitempty"`

	Detectors []DetectorRead `json:"detectors"`

	Influencers []string `json:"influencers"`

	Latency Duration `json:"latency,omitempty"`

	ModelPruneWindow Duration `json:"model_prune_window,omitempty"`

	MultivariateByFields *bool `json:"multivariate_by_fields,omitempty"`

	PerPartitionCategorization *PerPartitionCategorization `json:"per_partition_categorization,omitempty"`

	SummaryCountFieldName *string `json:"summary_count_field_name,omitempty"`
}

func (s *AnalysisConfigRead) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAnalysisConfigRead() *AnalysisConfigRead { _ = "STUB: not implemented"; return nil }
