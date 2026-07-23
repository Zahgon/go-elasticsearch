package types

type BrazilianStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *BrazilianStemTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s BrazilianStemTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBrazilianStemTokenFilter() *BrazilianStemTokenFilter { _ = "STUB: not implemented"; return nil }

type BrazilianStemTokenFilterVariant interface {
	BrazilianStemTokenFilterCaster() *BrazilianStemTokenFilter
}

func (s *BrazilianStemTokenFilter) BrazilianStemTokenFilterCaster() *BrazilianStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *BrazilianStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
