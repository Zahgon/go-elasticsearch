package types

type BasqueAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *BasqueAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s BasqueAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewBasqueAnalyzer() *BasqueAnalyzer { _ = "STUB: not implemented"; return nil }

type BasqueAnalyzerVariant interface {
	BasqueAnalyzerCaster() *BasqueAnalyzer
}

func (s *BasqueAnalyzer) BasqueAnalyzerCaster() *BasqueAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *BasqueAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
