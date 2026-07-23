package types

type BengaliAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *BengaliAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s BengaliAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewBengaliAnalyzer() *BengaliAnalyzer { _ = "STUB: not implemented"; return nil }

type BengaliAnalyzerVariant interface {
	BengaliAnalyzerCaster() *BengaliAnalyzer
}

func (s *BengaliAnalyzer) BengaliAnalyzerCaster() *BengaliAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *BengaliAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
