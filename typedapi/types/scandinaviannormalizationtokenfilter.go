package types

type ScandinavianNormalizationTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *ScandinavianNormalizationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ScandinavianNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewScandinavianNormalizationTokenFilter() *ScandinavianNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type ScandinavianNormalizationTokenFilterVariant interface {
	ScandinavianNormalizationTokenFilterCaster() *ScandinavianNormalizationTokenFilter
}

func (s *ScandinavianNormalizationTokenFilter) ScandinavianNormalizationTokenFilterCaster() *ScandinavianNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ScandinavianNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
