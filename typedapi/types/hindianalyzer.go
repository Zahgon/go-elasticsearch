package types

type HindiAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *HindiAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s HindiAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewHindiAnalyzer() *HindiAnalyzer { _ = "STUB: not implemented"; return nil }

type HindiAnalyzerVariant interface {
	HindiAnalyzerCaster() *HindiAnalyzer
}

func (s *HindiAnalyzer) HindiAnalyzerCaster() *HindiAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *HindiAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
