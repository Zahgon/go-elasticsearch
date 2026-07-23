package types

type LithuanianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *LithuanianAnalyzer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s LithuanianAnalyzer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLithuanianAnalyzer() *LithuanianAnalyzer { _ = "STUB: not implemented"; return nil }

type LithuanianAnalyzerVariant interface {
	LithuanianAnalyzerCaster() *LithuanianAnalyzer
}

func (s *LithuanianAnalyzer) LithuanianAnalyzerCaster() *LithuanianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *LithuanianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
