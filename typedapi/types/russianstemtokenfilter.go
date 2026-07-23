package types

type RussianStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *RussianStemTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s RussianStemTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewRussianStemTokenFilter() *RussianStemTokenFilter { _ = "STUB: not implemented"; return nil }

type RussianStemTokenFilterVariant interface {
	RussianStemTokenFilterCaster() *RussianStemTokenFilter
}

func (s *RussianStemTokenFilter) RussianStemTokenFilterCaster() *RussianStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *RussianStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
