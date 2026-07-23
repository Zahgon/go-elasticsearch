package types

type TurkishAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *TurkishAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s TurkishAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewTurkishAnalyzer() *TurkishAnalyzer { _ = "STUB: not implemented"; return nil }

type TurkishAnalyzerVariant interface {
	TurkishAnalyzerCaster() *TurkishAnalyzer
}

func (s *TurkishAnalyzer) TurkishAnalyzerCaster() *TurkishAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *TurkishAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
