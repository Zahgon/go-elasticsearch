package types

type LatvianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *LatvianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s LatvianAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewLatvianAnalyzer() *LatvianAnalyzer { _ = "STUB: not implemented"; return nil }

type LatvianAnalyzerVariant interface {
	LatvianAnalyzerCaster() *LatvianAnalyzer
}

func (s *LatvianAnalyzer) LatvianAnalyzerCaster() *LatvianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *LatvianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
