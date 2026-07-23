package types

type ArabicAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *ArabicAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ArabicAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewArabicAnalyzer() *ArabicAnalyzer { _ = "STUB: not implemented"; return nil }

type ArabicAnalyzerVariant interface {
	ArabicAnalyzerCaster() *ArabicAnalyzer
}

func (s *ArabicAnalyzer) ArabicAnalyzerCaster() *ArabicAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *ArabicAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
