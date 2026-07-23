package types

type BengaliNormalizationTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *BengaliNormalizationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s BengaliNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBengaliNormalizationTokenFilter() *BengaliNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type BengaliNormalizationTokenFilterVariant interface {
	BengaliNormalizationTokenFilterCaster() *BengaliNormalizationTokenFilter
}

func (s *BengaliNormalizationTokenFilter) BengaliNormalizationTokenFilterCaster() *BengaliNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *BengaliNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
