package types

type ClassicTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *ClassicTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ClassicTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewClassicTokenFilter() *ClassicTokenFilter { _ = "STUB: not implemented"; return nil }

type ClassicTokenFilterVariant interface {
	ClassicTokenFilterCaster() *ClassicTokenFilter
}

func (s *ClassicTokenFilter) ClassicTokenFilterCaster() *ClassicTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClassicTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
