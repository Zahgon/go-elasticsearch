package types

type GreekAnalyzer struct {
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *GreekAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GreekAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGreekAnalyzer() *GreekAnalyzer { _ = "STUB: not implemented"; return nil }

type GreekAnalyzerVariant interface {
	GreekAnalyzerCaster() *GreekAnalyzer
}

func (s *GreekAnalyzer) GreekAnalyzerCaster() *GreekAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *GreekAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
