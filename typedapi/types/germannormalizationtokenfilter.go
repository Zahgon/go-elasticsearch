package types

type GermanNormalizationTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *GermanNormalizationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s GermanNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGermanNormalizationTokenFilter() *GermanNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type GermanNormalizationTokenFilterVariant interface {
	GermanNormalizationTokenFilterCaster() *GermanNormalizationTokenFilter
}

func (s *GermanNormalizationTokenFilter) GermanNormalizationTokenFilterCaster() *GermanNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *GermanNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
