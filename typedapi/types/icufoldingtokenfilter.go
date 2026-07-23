package types

type IcuFoldingTokenFilter struct {
	Type             string  `json:"type,omitempty"`
	UnicodeSetFilter string  `json:"unicode_set_filter"`
	Version          *string `json:"version,omitempty"`
}

func (s *IcuFoldingTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s IcuFoldingTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIcuFoldingTokenFilter() *IcuFoldingTokenFilter { _ = "STUB: not implemented"; return nil }

type IcuFoldingTokenFilterVariant interface {
	IcuFoldingTokenFilterCaster() *IcuFoldingTokenFilter
}

func (s *IcuFoldingTokenFilter) IcuFoldingTokenFilterCaster() *IcuFoldingTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *IcuFoldingTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
