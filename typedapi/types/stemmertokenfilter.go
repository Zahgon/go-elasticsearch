package types

type StemmerTokenFilter struct {
	Language *string `json:"language,omitempty"`
	Type     string  `json:"type,omitempty"`
	Version  *string `json:"version,omitempty"`
}

func (s *StemmerTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s StemmerTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewStemmerTokenFilter() *StemmerTokenFilter { _ = "STUB: not implemented"; return nil }

type StemmerTokenFilterVariant interface {
	StemmerTokenFilterCaster() *StemmerTokenFilter
}

func (s *StemmerTokenFilter) StemmerTokenFilterCaster() *StemmerTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *StemmerTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
