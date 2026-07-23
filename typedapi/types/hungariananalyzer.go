package types

type HungarianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *HungarianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s HungarianAnalyzer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHungarianAnalyzer() *HungarianAnalyzer { _ = "STUB: not implemented"; return nil }

type HungarianAnalyzerVariant interface {
	HungarianAnalyzerCaster() *HungarianAnalyzer
}

func (s *HungarianAnalyzer) HungarianAnalyzerCaster() *HungarianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *HungarianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
