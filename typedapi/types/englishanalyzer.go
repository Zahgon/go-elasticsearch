package types

type EnglishAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *EnglishAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s EnglishAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewEnglishAnalyzer() *EnglishAnalyzer { _ = "STUB: not implemented"; return nil }

type EnglishAnalyzerVariant interface {
	EnglishAnalyzerCaster() *EnglishAnalyzer
}

func (s *EnglishAnalyzer) EnglishAnalyzerCaster() *EnglishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *EnglishAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
