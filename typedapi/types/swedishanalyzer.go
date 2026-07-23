package types

type SwedishAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *SwedishAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s SwedishAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewSwedishAnalyzer() *SwedishAnalyzer { _ = "STUB: not implemented"; return nil }

type SwedishAnalyzerVariant interface {
	SwedishAnalyzerCaster() *SwedishAnalyzer
}

func (s *SwedishAnalyzer) SwedishAnalyzerCaster() *SwedishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *SwedishAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
