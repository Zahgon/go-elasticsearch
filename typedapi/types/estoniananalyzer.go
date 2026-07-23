package types

type EstonianAnalyzer struct {
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *EstonianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s EstonianAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewEstonianAnalyzer() *EstonianAnalyzer { _ = "STUB: not implemented"; return nil }

type EstonianAnalyzerVariant interface {
	EstonianAnalyzerCaster() *EstonianAnalyzer
}

func (s *EstonianAnalyzer) EstonianAnalyzerCaster() *EstonianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *EstonianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
