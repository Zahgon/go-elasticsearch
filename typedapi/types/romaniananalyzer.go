package types

type RomanianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *RomanianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s RomanianAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewRomanianAnalyzer() *RomanianAnalyzer { _ = "STUB: not implemented"; return nil }

type RomanianAnalyzerVariant interface {
	RomanianAnalyzerCaster() *RomanianAnalyzer
}

func (s *RomanianAnalyzer) RomanianAnalyzerCaster() *RomanianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *RomanianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
