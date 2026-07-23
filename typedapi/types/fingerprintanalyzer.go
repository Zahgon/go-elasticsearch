package types

type FingerprintAnalyzer struct {
	MaxOutputSize *int `json:"max_output_size,omitempty"`

	Separator *string `json:"separator,omitempty"`

	Stopwords StopWords `json:"stopwords,omitempty"`

	StopwordsPath *string `json:"stopwords_path,omitempty"`
	Type          string  `json:"type,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *FingerprintAnalyzer) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FingerprintAnalyzer) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFingerprintAnalyzer() *FingerprintAnalyzer { _ = "STUB: not implemented"; return nil }

type FingerprintAnalyzerVariant interface {
	FingerprintAnalyzerCaster() *FingerprintAnalyzer
}

func (s *FingerprintAnalyzer) FingerprintAnalyzerCaster() *FingerprintAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *FingerprintAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
