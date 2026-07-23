package types

type ArmenianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *ArmenianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ArmenianAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewArmenianAnalyzer() *ArmenianAnalyzer { _ = "STUB: not implemented"; return nil }

type ArmenianAnalyzerVariant interface {
	ArmenianAnalyzerCaster() *ArmenianAnalyzer
}

func (s *ArmenianAnalyzer) ArmenianAnalyzerCaster() *ArmenianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *ArmenianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
