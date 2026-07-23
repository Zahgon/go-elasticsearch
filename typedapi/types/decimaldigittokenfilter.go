package types

type DecimalDigitTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *DecimalDigitTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s DecimalDigitTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDecimalDigitTokenFilter() *DecimalDigitTokenFilter { _ = "STUB: not implemented"; return nil }

type DecimalDigitTokenFilterVariant interface {
	DecimalDigitTokenFilterCaster() *DecimalDigitTokenFilter
}

func (s *DecimalDigitTokenFilter) DecimalDigitTokenFilterCaster() *DecimalDigitTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *DecimalDigitTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
