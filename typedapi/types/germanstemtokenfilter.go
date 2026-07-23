package types

type GermanStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *GermanStemTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s GermanStemTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGermanStemTokenFilter() *GermanStemTokenFilter { _ = "STUB: not implemented"; return nil }

type GermanStemTokenFilterVariant interface {
	GermanStemTokenFilterCaster() *GermanStemTokenFilter
}

func (s *GermanStemTokenFilter) GermanStemTokenFilterCaster() *GermanStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *GermanStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
