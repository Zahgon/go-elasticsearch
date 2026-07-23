package types

type BrazilianAnalyzer struct {
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *BrazilianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s BrazilianAnalyzer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBrazilianAnalyzer() *BrazilianAnalyzer { _ = "STUB: not implemented"; return nil }

type BrazilianAnalyzerVariant interface {
	BrazilianAnalyzerCaster() *BrazilianAnalyzer
}

func (s *BrazilianAnalyzer) BrazilianAnalyzerCaster() *BrazilianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *BrazilianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
