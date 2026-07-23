package types

type SerbianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *SerbianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s SerbianAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewSerbianAnalyzer() *SerbianAnalyzer { _ = "STUB: not implemented"; return nil }

type SerbianAnalyzerVariant interface {
	SerbianAnalyzerCaster() *SerbianAnalyzer
}

func (s *SerbianAnalyzer) SerbianAnalyzerCaster() *SerbianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *SerbianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
