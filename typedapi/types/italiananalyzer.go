package types

type ItalianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *ItalianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ItalianAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewItalianAnalyzer() *ItalianAnalyzer { _ = "STUB: not implemented"; return nil }

type ItalianAnalyzerVariant interface {
	ItalianAnalyzerCaster() *ItalianAnalyzer
}

func (s *ItalianAnalyzer) ItalianAnalyzerCaster() *ItalianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *ItalianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
