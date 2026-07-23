package types

type PorterStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *PorterStemTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PorterStemTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPorterStemTokenFilter() *PorterStemTokenFilter { _ = "STUB: not implemented"; return nil }

type PorterStemTokenFilterVariant interface {
	PorterStemTokenFilterCaster() *PorterStemTokenFilter
}

func (s *PorterStemTokenFilter) PorterStemTokenFilterCaster() *PorterStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *PorterStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
