package types

type SerbianNormalizationTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *SerbianNormalizationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SerbianNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSerbianNormalizationTokenFilter() *SerbianNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type SerbianNormalizationTokenFilterVariant interface {
	SerbianNormalizationTokenFilterCaster() *SerbianNormalizationTokenFilter
}

func (s *SerbianNormalizationTokenFilter) SerbianNormalizationTokenFilterCaster() *SerbianNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *SerbianNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
