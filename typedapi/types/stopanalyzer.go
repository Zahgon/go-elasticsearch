package types

type StopAnalyzer struct {
	Stopwords StopWords `json:"stopwords,omitempty"`

	StopwordsPath *string `json:"stopwords_path,omitempty"`
	Type          string  `json:"type,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *StopAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s StopAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewStopAnalyzer() *StopAnalyzer { _ = "STUB: not implemented"; return nil }

type StopAnalyzerVariant interface {
	StopAnalyzerCaster() *StopAnalyzer
}

func (s *StopAnalyzer) StopAnalyzerCaster() *StopAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *StopAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
