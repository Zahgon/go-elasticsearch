package types

type DutchAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *DutchAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s DutchAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewDutchAnalyzer() *DutchAnalyzer { _ = "STUB: not implemented"; return nil }

type DutchAnalyzerVariant interface {
	DutchAnalyzerCaster() *DutchAnalyzer
}

func (s *DutchAnalyzer) DutchAnalyzerCaster() *DutchAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *DutchAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
