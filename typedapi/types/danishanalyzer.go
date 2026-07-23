package types

type DanishAnalyzer struct {
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *DanishAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s DanishAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDanishAnalyzer() *DanishAnalyzer { _ = "STUB: not implemented"; return nil }

type DanishAnalyzerVariant interface {
	DanishAnalyzerCaster() *DanishAnalyzer
}

func (s *DanishAnalyzer) DanishAnalyzerCaster() *DanishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *DanishAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
