package types

type HindiNormalizationTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *HindiNormalizationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s HindiNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewHindiNormalizationTokenFilter() *HindiNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type HindiNormalizationTokenFilterVariant interface {
	HindiNormalizationTokenFilterCaster() *HindiNormalizationTokenFilter
}

func (s *HindiNormalizationTokenFilter) HindiNormalizationTokenFilterCaster() *HindiNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *HindiNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
