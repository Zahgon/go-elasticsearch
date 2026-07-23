package types

type IndonesianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *IndonesianAnalyzer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s IndonesianAnalyzer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIndonesianAnalyzer() *IndonesianAnalyzer { _ = "STUB: not implemented"; return nil }

type IndonesianAnalyzerVariant interface {
	IndonesianAnalyzerCaster() *IndonesianAnalyzer
}

func (s *IndonesianAnalyzer) IndonesianAnalyzerCaster() *IndonesianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *IndonesianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
