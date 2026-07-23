package types

type ReverseTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *ReverseTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ReverseTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewReverseTokenFilter() *ReverseTokenFilter { _ = "STUB: not implemented"; return nil }

type ReverseTokenFilterVariant interface {
	ReverseTokenFilterCaster() *ReverseTokenFilter
}

func (s *ReverseTokenFilter) ReverseTokenFilterCaster() *ReverseTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ReverseTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
