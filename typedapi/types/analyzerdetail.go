package types

type AnalyzerDetail struct {
	Name   string                `json:"name"`
	Tokens []ExplainAnalyzeToken `json:"tokens"`
}

func (s *AnalyzerDetail) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAnalyzerDetail() *AnalyzerDetail { _ = "STUB: not implemented"; return nil }
