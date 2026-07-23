package types

type FrenchAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *FrenchAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s FrenchAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewFrenchAnalyzer() *FrenchAnalyzer { _ = "STUB: not implemented"; return nil }

type FrenchAnalyzerVariant interface {
	FrenchAnalyzerCaster() *FrenchAnalyzer
}

func (s *FrenchAnalyzer) FrenchAnalyzerCaster() *FrenchAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *FrenchAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
