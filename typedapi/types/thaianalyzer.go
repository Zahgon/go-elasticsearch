package types

type ThaiAnalyzer struct {
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *ThaiAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ThaiAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewThaiAnalyzer() *ThaiAnalyzer { _ = "STUB: not implemented"; return nil }

type ThaiAnalyzerVariant interface {
	ThaiAnalyzerCaster() *ThaiAnalyzer
}

func (s *ThaiAnalyzer) ThaiAnalyzerCaster() *ThaiAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *ThaiAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
