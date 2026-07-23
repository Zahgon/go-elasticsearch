package types

type IrishAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *IrishAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s IrishAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewIrishAnalyzer() *IrishAnalyzer { _ = "STUB: not implemented"; return nil }

type IrishAnalyzerVariant interface {
	IrishAnalyzerCaster() *IrishAnalyzer
}

func (s *IrishAnalyzer) IrishAnalyzerCaster() *IrishAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *IrishAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
