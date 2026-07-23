package types

type AnalysisLimits struct {
	CategorizationExamplesLimit *int64 `json:"categorization_examples_limit,omitempty"`

	ModelMemoryLimit ByteSize `json:"model_memory_limit,omitempty"`
}

func (s *AnalysisLimits) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAnalysisLimits() *AnalysisLimits { _ = "STUB: not implemented"; return nil }

type AnalysisLimitsVariant interface {
	AnalysisLimitsCaster() *AnalysisLimits
}

func (s *AnalysisLimits) AnalysisLimitsCaster() *AnalysisLimits {
	_ = "STUB: not implemented"
	return nil
}
