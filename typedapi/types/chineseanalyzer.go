package types

type ChineseAnalyzer struct {
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *ChineseAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ChineseAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewChineseAnalyzer() *ChineseAnalyzer { _ = "STUB: not implemented"; return nil }

type ChineseAnalyzerVariant interface {
	ChineseAnalyzerCaster() *ChineseAnalyzer
}

func (s *ChineseAnalyzer) ChineseAnalyzerCaster() *ChineseAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *ChineseAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
