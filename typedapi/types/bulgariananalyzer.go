package types

type BulgarianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *BulgarianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s BulgarianAnalyzer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBulgarianAnalyzer() *BulgarianAnalyzer { _ = "STUB: not implemented"; return nil }

type BulgarianAnalyzerVariant interface {
	BulgarianAnalyzerCaster() *BulgarianAnalyzer
}

func (s *BulgarianAnalyzer) BulgarianAnalyzerCaster() *BulgarianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *BulgarianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
