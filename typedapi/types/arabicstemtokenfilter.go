package types

type ArabicStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *ArabicStemTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ArabicStemTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewArabicStemTokenFilter() *ArabicStemTokenFilter { _ = "STUB: not implemented"; return nil }

type ArabicStemTokenFilterVariant interface {
	ArabicStemTokenFilterCaster() *ArabicStemTokenFilter
}

func (s *ArabicStemTokenFilter) ArabicStemTokenFilterCaster() *ArabicStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ArabicStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
