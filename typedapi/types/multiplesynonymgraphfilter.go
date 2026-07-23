package types

type MultipleSynonymGraphFilter struct {
	AnalyzerCount *int `json:"analyzer_count,omitempty"`

	IndexCount *int `json:"index_count,omitempty"`
}

func (s *MultipleSynonymGraphFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMultipleSynonymGraphFilter() *MultipleSynonymGraphFilter {
	_ = "STUB: not implemented"
	return nil
}
