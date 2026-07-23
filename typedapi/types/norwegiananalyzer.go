package types

type NorwegianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *NorwegianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s NorwegianAnalyzer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewNorwegianAnalyzer() *NorwegianAnalyzer { _ = "STUB: not implemented"; return nil }

type NorwegianAnalyzerVariant interface {
	NorwegianAnalyzerCaster() *NorwegianAnalyzer
}

func (s *NorwegianAnalyzer) NorwegianAnalyzerCaster() *NorwegianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *NorwegianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
