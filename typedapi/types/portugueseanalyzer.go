package types

type PortugueseAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *PortugueseAnalyzer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PortugueseAnalyzer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPortugueseAnalyzer() *PortugueseAnalyzer { _ = "STUB: not implemented"; return nil }

type PortugueseAnalyzerVariant interface {
	PortugueseAnalyzerCaster() *PortugueseAnalyzer
}

func (s *PortugueseAnalyzer) PortugueseAnalyzerCaster() *PortugueseAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *PortugueseAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
