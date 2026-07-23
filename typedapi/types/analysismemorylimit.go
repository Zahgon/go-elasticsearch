package types

type AnalysisMemoryLimit struct {
	ModelMemoryLimit string `json:"model_memory_limit"`
}

func (s *AnalysisMemoryLimit) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAnalysisMemoryLimit() *AnalysisMemoryLimit { _ = "STUB: not implemented"; return nil }

type AnalysisMemoryLimitVariant interface {
	AnalysisMemoryLimitCaster() *AnalysisMemoryLimit
}

func (s *AnalysisMemoryLimit) AnalysisMemoryLimitCaster() *AnalysisMemoryLimit {
	_ = "STUB: not implemented"
	return nil
}
