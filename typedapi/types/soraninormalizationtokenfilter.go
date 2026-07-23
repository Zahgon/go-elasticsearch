package types

type SoraniNormalizationTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *SoraniNormalizationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SoraniNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSoraniNormalizationTokenFilter() *SoraniNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type SoraniNormalizationTokenFilterVariant interface {
	SoraniNormalizationTokenFilterCaster() *SoraniNormalizationTokenFilter
}

func (s *SoraniNormalizationTokenFilter) SoraniNormalizationTokenFilterCaster() *SoraniNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *SoraniNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
