package types

type PersianStemTokenFilter struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *PersianStemTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s PersianStemTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewPersianStemTokenFilter() *PersianStemTokenFilter { _ = "STUB: not implemented"; return nil }

type PersianStemTokenFilterVariant interface {
	PersianStemTokenFilterCaster() *PersianStemTokenFilter
}

func (s *PersianStemTokenFilter) PersianStemTokenFilterCaster() *PersianStemTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *PersianStemTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
