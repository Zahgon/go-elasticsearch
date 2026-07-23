package types

type PersianAnalyzer struct {
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *PersianAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s PersianAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewPersianAnalyzer() *PersianAnalyzer { _ = "STUB: not implemented"; return nil }

type PersianAnalyzerVariant interface {
	PersianAnalyzerCaster() *PersianAnalyzer
}

func (s *PersianAnalyzer) PersianAnalyzerCaster() *PersianAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *PersianAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
