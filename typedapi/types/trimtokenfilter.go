package types

type TrimTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *TrimTokenFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s TrimTokenFilter) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewTrimTokenFilter() *TrimTokenFilter { _ = "STUB: not implemented"; return nil }

type TrimTokenFilterVariant interface {
	TrimTokenFilterCaster() *TrimTokenFilter
}

func (s *TrimTokenFilter) TrimTokenFilterCaster() *TrimTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *TrimTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
