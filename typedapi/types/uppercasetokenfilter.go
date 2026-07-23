package types

type UppercaseTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *UppercaseTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s UppercaseTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewUppercaseTokenFilter() *UppercaseTokenFilter { _ = "STUB: not implemented"; return nil }

type UppercaseTokenFilterVariant interface {
	UppercaseTokenFilterCaster() *UppercaseTokenFilter
}

func (s *UppercaseTokenFilter) UppercaseTokenFilterCaster() *UppercaseTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *UppercaseTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
