package types

type FingerprintTokenFilter struct {
	MaxOutputSize *int `json:"max_output_size,omitempty"`

	Separator *string `json:"separator,omitempty"`
	Type      string  `json:"type,omitempty"`
	Version   *string `json:"version,omitempty"`
}

func (s *FingerprintTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FingerprintTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFingerprintTokenFilter() *FingerprintTokenFilter { _ = "STUB: not implemented"; return nil }

type FingerprintTokenFilterVariant interface {
	FingerprintTokenFilterCaster() *FingerprintTokenFilter
}

func (s *FingerprintTokenFilter) FingerprintTokenFilterCaster() *FingerprintTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *FingerprintTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
