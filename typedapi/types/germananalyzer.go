package types

type GermanAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *GermanAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GermanAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGermanAnalyzer() *GermanAnalyzer { _ = "STUB: not implemented"; return nil }

type GermanAnalyzerVariant interface {
	GermanAnalyzerCaster() *GermanAnalyzer
}

func (s *GermanAnalyzer) GermanAnalyzerCaster() *GermanAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *GermanAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
