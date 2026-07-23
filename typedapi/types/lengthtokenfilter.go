package types

type LengthTokenFilter struct {
	Max *int `json:"max,omitempty"`

	Min     *int    `json:"min,omitempty"`
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *LengthTokenFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s LengthTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLengthTokenFilter() *LengthTokenFilter { _ = "STUB: not implemented"; return nil }

type LengthTokenFilterVariant interface {
	LengthTokenFilterCaster() *LengthTokenFilter
}

func (s *LengthTokenFilter) LengthTokenFilterCaster() *LengthTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *LengthTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
