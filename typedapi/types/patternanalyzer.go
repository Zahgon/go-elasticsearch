package types

type PatternAnalyzer struct {
	Flags *string `json:"flags,omitempty"`

	Lowercase *bool `json:"lowercase,omitempty"`

	Pattern *string `json:"pattern,omitempty"`

	Stopwords StopWords `json:"stopwords,omitempty"`

	StopwordsPath *string `json:"stopwords_path,omitempty"`
	Type          string  `json:"type,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *PatternAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s PatternAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewPatternAnalyzer() *PatternAnalyzer { _ = "STUB: not implemented"; return nil }

type PatternAnalyzerVariant interface {
	PatternAnalyzerCaster() *PatternAnalyzer
}

func (s *PatternAnalyzer) PatternAnalyzerCaster() *PatternAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *PatternAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
