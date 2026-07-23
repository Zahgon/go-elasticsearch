package types

type FinnishAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *FinnishAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s FinnishAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewFinnishAnalyzer() *FinnishAnalyzer { _ = "STUB: not implemented"; return nil }

type FinnishAnalyzerVariant interface {
	FinnishAnalyzerCaster() *FinnishAnalyzer
}

func (s *FinnishAnalyzer) FinnishAnalyzerCaster() *FinnishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *FinnishAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
