package types

type ArabicNormalizationTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *ArabicNormalizationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ArabicNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewArabicNormalizationTokenFilter() *ArabicNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type ArabicNormalizationTokenFilterVariant interface {
	ArabicNormalizationTokenFilterCaster() *ArabicNormalizationTokenFilter
}

func (s *ArabicNormalizationTokenFilter) ArabicNormalizationTokenFilterCaster() *ArabicNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ArabicNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
