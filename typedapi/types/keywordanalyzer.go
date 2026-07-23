package types

type KeywordAnalyzer struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *KeywordAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s KeywordAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewKeywordAnalyzer() *KeywordAnalyzer { _ = "STUB: not implemented"; return nil }

type KeywordAnalyzerVariant interface {
	KeywordAnalyzerCaster() *KeywordAnalyzer
}

func (s *KeywordAnalyzer) KeywordAnalyzerCaster() *KeywordAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *KeywordAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
