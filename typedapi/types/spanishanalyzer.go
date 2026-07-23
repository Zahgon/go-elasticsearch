package types

type SpanishAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *SpanishAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s SpanishAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewSpanishAnalyzer() *SpanishAnalyzer { _ = "STUB: not implemented"; return nil }

type SpanishAnalyzerVariant interface {
	SpanishAnalyzerCaster() *SpanishAnalyzer
}

func (s *SpanishAnalyzer) SpanishAnalyzerCaster() *SpanishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *SpanishAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
