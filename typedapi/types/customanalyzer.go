package types

type CustomAnalyzer struct {
	CharFilter           []string `json:"char_filter,omitempty"`
	Filter               []string `json:"filter,omitempty"`
	PositionIncrementGap *int     `json:"position_increment_gap,omitempty"`
	PositionOffsetGap    *int     `json:"position_offset_gap,omitempty"`
	Tokenizer            string   `json:"tokenizer"`
	Type                 string   `json:"type,omitempty"`
}

func (s *CustomAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s CustomAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewCustomAnalyzer() *CustomAnalyzer { _ = "STUB: not implemented"; return nil }

type CustomAnalyzerVariant interface {
	CustomAnalyzerCaster() *CustomAnalyzer
}

func (s *CustomAnalyzer) CustomAnalyzerCaster() *CustomAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *CustomAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
