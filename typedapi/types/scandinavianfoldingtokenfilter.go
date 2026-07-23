package types

type ScandinavianFoldingTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *ScandinavianFoldingTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ScandinavianFoldingTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewScandinavianFoldingTokenFilter() *ScandinavianFoldingTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type ScandinavianFoldingTokenFilterVariant interface {
	ScandinavianFoldingTokenFilterCaster() *ScandinavianFoldingTokenFilter
}

func (s *ScandinavianFoldingTokenFilter) ScandinavianFoldingTokenFilterCaster() *ScandinavianFoldingTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ScandinavianFoldingTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
