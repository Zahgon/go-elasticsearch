package types

type FrenchStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *FrenchStemTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FrenchStemTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFrenchStemTokenFilter() *FrenchStemTokenFilter { _ = "STUB: not implemented"; return nil }

type FrenchStemTokenFilterVariant interface {
	FrenchStemTokenFilterCaster() *FrenchStemTokenFilter
}

func (s *FrenchStemTokenFilter) FrenchStemTokenFilterCaster() *FrenchStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *FrenchStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
