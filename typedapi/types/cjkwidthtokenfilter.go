package types

type CjkWidthTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *CjkWidthTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s CjkWidthTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCjkWidthTokenFilter() *CjkWidthTokenFilter { _ = "STUB: not implemented"; return nil }

type CjkWidthTokenFilterVariant interface {
	CjkWidthTokenFilterCaster() *CjkWidthTokenFilter
}

func (s *CjkWidthTokenFilter) CjkWidthTokenFilterCaster() *CjkWidthTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *CjkWidthTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
