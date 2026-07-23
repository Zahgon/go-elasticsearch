package types

type PersianNormalizationTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *PersianNormalizationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PersianNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPersianNormalizationTokenFilter() *PersianNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type PersianNormalizationTokenFilterVariant interface {
	PersianNormalizationTokenFilterCaster() *PersianNormalizationTokenFilter
}

func (s *PersianNormalizationTokenFilter) PersianNormalizationTokenFilterCaster() *PersianNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *PersianNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
