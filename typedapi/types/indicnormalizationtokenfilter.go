package types

type IndicNormalizationTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *IndicNormalizationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s IndicNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIndicNormalizationTokenFilter() *IndicNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type IndicNormalizationTokenFilterVariant interface {
	IndicNormalizationTokenFilterCaster() *IndicNormalizationTokenFilter
}

func (s *IndicNormalizationTokenFilter) IndicNormalizationTokenFilterCaster() *IndicNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *IndicNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
