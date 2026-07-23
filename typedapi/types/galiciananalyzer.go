package types

type GalicianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *GalicianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s GalicianAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewGalicianAnalyzer() *GalicianAnalyzer { _ = "STUB: not implemented"; return nil }

type GalicianAnalyzerVariant interface {
	GalicianAnalyzerCaster() *GalicianAnalyzer
}

func (s *GalicianAnalyzer) GalicianAnalyzerCaster() *GalicianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *GalicianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
