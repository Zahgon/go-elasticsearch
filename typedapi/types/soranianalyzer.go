package types

type SoraniAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *SoraniAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s SoraniAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewSoraniAnalyzer() *SoraniAnalyzer { _ = "STUB: not implemented"; return nil }

type SoraniAnalyzerVariant interface {
	SoraniAnalyzerCaster() *SoraniAnalyzer
}

func (s *SoraniAnalyzer) SoraniAnalyzerCaster() *SoraniAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *SoraniAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
