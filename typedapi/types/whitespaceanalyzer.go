package types

type WhitespaceAnalyzer struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *WhitespaceAnalyzer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s WhitespaceAnalyzer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewWhitespaceAnalyzer() *WhitespaceAnalyzer { _ = "STUB: not implemented"; return nil }

type WhitespaceAnalyzerVariant interface {
	WhitespaceAnalyzerCaster() *WhitespaceAnalyzer
}

func (s *WhitespaceAnalyzer) WhitespaceAnalyzerCaster() *WhitespaceAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *WhitespaceAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
