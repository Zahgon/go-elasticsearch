package types

type RussianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *RussianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s RussianAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewRussianAnalyzer() *RussianAnalyzer { _ = "STUB: not implemented"; return nil }

type RussianAnalyzerVariant interface {
	RussianAnalyzerCaster() *RussianAnalyzer
}

func (s *RussianAnalyzer) RussianAnalyzerCaster() *RussianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *RussianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
