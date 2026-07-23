package types

type StandardAnalyzer struct {
	MaxTokenLength *int `json:"max_token_length,omitempty"`

	Stopwords StopWords `json:"stopwords,omitempty"`

	StopwordsPath *string `json:"stopwords_path,omitempty"`
	Type          string  `json:"type,omitempty"`
}

func (s *StandardAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s StandardAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewStandardAnalyzer() *StandardAnalyzer { _ = "STUB: not implemented"; return nil }

type StandardAnalyzerVariant interface {
	StandardAnalyzerCaster() *StandardAnalyzer
}

func (s *StandardAnalyzer) StandardAnalyzerCaster() *StandardAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *StandardAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
