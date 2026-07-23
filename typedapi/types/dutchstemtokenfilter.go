package types

type DutchStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *DutchStemTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s DutchStemTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDutchStemTokenFilter() *DutchStemTokenFilter { _ = "STUB: not implemented"; return nil }

type DutchStemTokenFilterVariant interface {
	DutchStemTokenFilterCaster() *DutchStemTokenFilter
}

func (s *DutchStemTokenFilter) DutchStemTokenFilterCaster() *DutchStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *DutchStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
