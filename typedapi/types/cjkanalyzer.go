package types

type CjkAnalyzer struct {
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *CjkAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s CjkAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewCjkAnalyzer() *CjkAnalyzer { _ = "STUB: not implemented"; return nil }

type CjkAnalyzerVariant interface {
	CjkAnalyzerCaster() *CjkAnalyzer
}

func (s *CjkAnalyzer) CjkAnalyzerCaster() *CjkAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *CjkAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
