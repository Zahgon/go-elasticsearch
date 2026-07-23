package types

type TruncateTokenFilter struct {
	Length  *int    `json:"length,omitempty"`
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *TruncateTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s TruncateTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTruncateTokenFilter() *TruncateTokenFilter { _ = "STUB: not implemented"; return nil }

type TruncateTokenFilterVariant interface {
	TruncateTokenFilterCaster() *TruncateTokenFilter
}

func (s *TruncateTokenFilter) TruncateTokenFilterCaster() *TruncateTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *TruncateTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
