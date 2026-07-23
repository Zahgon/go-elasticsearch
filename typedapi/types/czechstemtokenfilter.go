package types

type CzechStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *CzechStemTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s CzechStemTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCzechStemTokenFilter() *CzechStemTokenFilter { _ = "STUB: not implemented"; return nil }

type CzechStemTokenFilterVariant interface {
	CzechStemTokenFilterCaster() *CzechStemTokenFilter
}

func (s *CzechStemTokenFilter) CzechStemTokenFilterCaster() *CzechStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *CzechStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
