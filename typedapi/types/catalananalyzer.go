package types

type CatalanAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *CatalanAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s CatalanAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewCatalanAnalyzer() *CatalanAnalyzer { _ = "STUB: not implemented"; return nil }

type CatalanAnalyzerVariant interface {
	CatalanAnalyzerCaster() *CatalanAnalyzer
}

func (s *CatalanAnalyzer) CatalanAnalyzerCaster() *CatalanAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *CatalanAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
