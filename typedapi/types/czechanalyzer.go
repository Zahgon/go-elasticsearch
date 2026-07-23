package types

type CzechAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *CzechAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s CzechAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewCzechAnalyzer() *CzechAnalyzer { _ = "STUB: not implemented"; return nil }

type CzechAnalyzerVariant interface {
	CzechAnalyzerCaster() *CzechAnalyzer
}

func (s *CzechAnalyzer) CzechAnalyzerCaster() *CzechAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *CzechAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
